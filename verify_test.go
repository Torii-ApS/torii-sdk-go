package torii

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// Each test gets its own issuer because the JWKS cache is process-wide and
// keyed by URL — reusing an issuer across tests would leak keys.
var issuerCounter uint64

func nextIssuerHost() string {
	n := atomic.AddUint64(&issuerCounter, 1)
	return fmt.Sprintf("test-%d.torii.test", n)
}

// jwksFixture spins an in-process JWKS server and signs JWTs for tests.
type jwksFixture struct {
	priv    *ecdsa.PrivateKey
	jwkPriv jwk.Key
	jwkPub  jwk.Key
	keyID   string
	server  *httptest.Server
	issuer  string
}

func newJWKSFixture(t *testing.T) *jwksFixture {
	t.Helper()
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("generate ecdsa: %v", err)
	}
	privJWK, err := jwk.FromRaw(priv)
	if err != nil {
		t.Fatalf("jwk.FromRaw priv: %v", err)
	}
	const kid = "test-kid"
	_ = privJWK.Set(jwk.KeyIDKey, kid)
	_ = privJWK.Set(jwk.AlgorithmKey, jwa.ES256)

	pubJWK, err := jwk.PublicKeyOf(privJWK)
	if err != nil {
		t.Fatalf("jwk.PublicKeyOf: %v", err)
	}

	pubSet := jwk.NewSet()
	_ = pubSet.AddKey(pubJWK)

	mux := http.NewServeMux()
	mux.HandleFunc("/_torii/.well-known/jwks.json", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "max-age=60")
		_ = json.NewEncoder(w).Encode(pubSet)
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	t.Cleanup(resetJWKSCache)

	// Use a unique issuer per fixture so JWKS cache slots don't overlap with
	// other tests in this run.
	host := nextIssuerHost()
	issuer := srv.URL // we use the httptest URL as the issuer for end-to-end correctness
	_ = host          // (kept around for descriptive logs only)

	return &jwksFixture{
		priv:    priv,
		jwkPriv: privJWK,
		jwkPub:  pubJWK,
		keyID:   kid,
		server:  srv,
		issuer:  issuer,
	}
}

// signToken returns a signed JWT with the given claims, defaulting required
// fields when callers leave them empty.
func (f *jwksFixture) signToken(t *testing.T, claims map[string]any, signer jwk.Key) string {
	t.Helper()
	if signer == nil {
		signer = f.jwkPriv
	}
	tok := jwt.New()
	for k, v := range claims {
		if err := tok.Set(k, v); err != nil {
			t.Fatalf("token.Set(%q): %v", k, err)
		}
	}
	if _, ok := claims["iat"]; !ok {
		_ = tok.Set(jwt.IssuedAtKey, time.Now().Add(-1*time.Minute))
	}
	if _, ok := claims["exp"]; !ok {
		_ = tok.Set(jwt.ExpirationKey, time.Now().Add(5*time.Minute))
	}
	if _, ok := claims["iss"]; !ok {
		_ = tok.Set(jwt.IssuerKey, f.issuer)
	}
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.ES256, signer))
	if err != nil {
		t.Fatalf("jwt.Sign: %v", err)
	}
	return string(signed)
}

func TestVerifyToken(t *testing.T) {
	t.Run("success returns full Auth", func(t *testing.T) {
		f := newJWKSFixture(t)
		locale := "da"
		tok := f.signToken(t, map[string]any{
			"sub":              "user-123",
			"pid":              "env-456",
			"iss":              f.issuer,
			"email_verified":   true,
			"profile_complete": true,
			"impersonating":    false,
			"locale":           locale,
		}, nil)

		auth, err := VerifyToken(context.Background(), tok, VerifyOptions{Issuer: f.issuer})
		if err != nil {
			t.Fatalf("VerifyToken: %v", err)
		}
		if auth.UserID != "user-123" {
			t.Errorf("UserID = %q, want user-123", auth.UserID)
		}
		if auth.EnvironmentID != "env-456" {
			t.Errorf("EnvironmentID = %q, want env-456", auth.EnvironmentID)
		}
		if auth.Issuer != f.issuer {
			t.Errorf("Issuer = %q, want %q", auth.Issuer, f.issuer)
		}
		if !auth.EmailVerified {
			t.Error("EmailVerified = false, want true")
		}
		if !auth.ProfileComplete {
			t.Error("ProfileComplete = false, want true")
		}
		if auth.Impersonating {
			t.Error("Impersonating = true, want false")
		}
		if auth.Locale == nil || *auth.Locale != "da" {
			t.Errorf("Locale = %v, want 'da'", auth.Locale)
		}
		if auth.Raw["sub"] != "user-123" {
			t.Errorf("Raw[sub] = %v, want user-123", auth.Raw["sub"])
		}
	})

	t.Run("profile_complete defaults to true when claim absent", func(t *testing.T) {
		f := newJWKSFixture(t)
		tok := f.signToken(t, map[string]any{
			"sub": "u",
			"pid": "e",
			"iss": f.issuer,
		}, nil)
		auth, err := VerifyToken(context.Background(), tok, VerifyOptions{Issuer: f.issuer})
		if err != nil {
			t.Fatalf("VerifyToken: %v", err)
		}
		if !auth.ProfileComplete {
			t.Error("ProfileComplete should default to true when claim absent")
		}
		if auth.Locale != nil {
			t.Errorf("Locale = %v, want nil", auth.Locale)
		}
	})

	t.Run("rejects token signed with the wrong key", func(t *testing.T) {
		f := newJWKSFixture(t)
		otherPriv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		otherJWK, _ := jwk.FromRaw(otherPriv)
		_ = otherJWK.Set(jwk.KeyIDKey, f.keyID) // same kid, different key
		_ = otherJWK.Set(jwk.AlgorithmKey, jwa.ES256)
		tok := f.signToken(t, map[string]any{
			"sub": "u",
			"pid": "e",
			"iss": f.issuer,
		}, otherJWK)
		_, err := VerifyToken(context.Background(), tok, VerifyOptions{Issuer: f.issuer})
		if err == nil {
			t.Fatal("expected verification error for wrong key, got nil")
		}
	})

	t.Run("rejects token with mismatched issuer", func(t *testing.T) {
		f := newJWKSFixture(t)
		tok := f.signToken(t, map[string]any{
			"sub": "u",
			"pid": "e",
			"iss": "https://evil.example.com",
		}, nil)
		_, err := VerifyToken(context.Background(), tok, VerifyOptions{Issuer: f.issuer})
		if err == nil {
			t.Fatal("expected verification error for mismatched iss, got nil")
		}
	})

	t.Run("rejects token missing required claim 'pid'", func(t *testing.T) {
		f := newJWKSFixture(t)
		tok := f.signToken(t, map[string]any{
			"sub": "u",
			"iss": f.issuer,
		}, nil)
		_, err := VerifyToken(context.Background(), tok, VerifyOptions{Issuer: f.issuer})
		if err == nil {
			t.Fatal("expected verification error for missing pid, got nil")
		}
		if !strings.Contains(err.Error(), "pid") {
			t.Errorf("error should mention missing 'pid': %v", err)
		}
	})

	t.Run("rejects expired token", func(t *testing.T) {
		f := newJWKSFixture(t)
		tok := f.signToken(t, map[string]any{
			"sub": "u",
			"pid": "e",
			"iss": f.issuer,
			"iat": time.Now().Add(-2 * time.Hour),
			"exp": time.Now().Add(-1 * time.Hour),
		}, nil)
		_, err := VerifyToken(context.Background(), tok, VerifyOptions{Issuer: f.issuer})
		if err == nil {
			t.Fatal("expected verification error for expired token, got nil")
		}
	})

	t.Run("empty token rejected", func(t *testing.T) {
		_, err := VerifyToken(context.Background(), "", VerifyOptions{Issuer: "https://x"})
		if err == nil {
			t.Fatal("expected error for empty token, got nil")
		}
	})

	t.Run("empty issuer rejected", func(t *testing.T) {
		_, err := VerifyToken(context.Background(), "abc", VerifyOptions{})
		if err == nil {
			t.Fatal("expected error for empty issuer, got nil")
		}
	})
}

func TestAuthenticateRequest(t *testing.T) {
	t.Run("reads bearer header and verifies", func(t *testing.T) {
		f := newJWKSFixture(t)
		tok := f.signToken(t, map[string]any{
			"sub": "u-7",
			"pid": "p-9",
			"iss": f.issuer,
		}, nil)

		h := http.Header{}
		h.Set("Authorization", "Bearer "+tok)

		auth, err := AuthenticateRequest(context.Background(), h, VerifyOptions{Issuer: f.issuer})
		if err != nil {
			t.Fatalf("AuthenticateRequest: %v", err)
		}
		if auth.UserID != "u-7" {
			t.Errorf("UserID = %q, want u-7", auth.UserID)
		}
	})

	t.Run("missing header rejected", func(t *testing.T) {
		_, err := AuthenticateRequest(context.Background(), http.Header{}, VerifyOptions{Issuer: "https://x"})
		if err == nil {
			t.Fatal("expected error for missing header, got nil")
		}
		if !strings.Contains(err.Error(), "Authorization") {
			t.Errorf("error should mention 'Authorization': %v", err)
		}
	})

	t.Run("non-bearer scheme rejected", func(t *testing.T) {
		h := http.Header{}
		h.Set("Authorization", "Basic ZGVtbzpkZW1v")
		_, err := AuthenticateRequest(context.Background(), h, VerifyOptions{Issuer: "https://x"})
		if err == nil {
			t.Fatal("expected error for non-Bearer scheme, got nil")
		}
		if !strings.Contains(err.Error(), "Bearer") {
			t.Errorf("error should mention 'Bearer': %v", err)
		}
	})

	t.Run("bearer with empty value rejected", func(t *testing.T) {
		h := http.Header{}
		h.Set("Authorization", "Bearer ")
		_, err := AuthenticateRequest(context.Background(), h, VerifyOptions{Issuer: "https://x"})
		if err == nil {
			t.Fatal("expected error for empty bearer value, got nil")
		}
	})
}

func TestNew(t *testing.T) {
	t.Run("requires secret key", func(t *testing.T) {
		_, err := New(Options{})
		if err == nil {
			t.Fatal("expected error for empty SecretKey")
		}
	})

	t.Run("rejects malformed API URL", func(t *testing.T) {
		_, err := New(Options{SecretKey: "sk_test", APIURL: "not-a-url"})
		if err == nil {
			t.Fatal("expected error for invalid APIURL")
		}
	})

	t.Run("defaults APIURL when empty", func(t *testing.T) {
		c, err := New(Options{SecretKey: "sk_test"})
		if err != nil {
			t.Fatalf("New: %v", err)
		}
		if c.Users() == nil || c.Sessions() == nil {
			t.Fatal("Users/Sessions should be non-nil")
		}
	})
}

func TestVerifyWebhook_stub(t *testing.T) {
	_, err := VerifyWebhook("whsec_xxx", http.Header{}, []byte("{}"))
	if err == nil {
		t.Fatal("expected VerifyWebhook to return stub error")
	}
	if !strings.Contains(err.Error(), "not yet available") {
		t.Errorf("stub error should mention webhook is not yet available: %v", err)
	}
}
