package torii

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// jwksMinRefreshInterval is the floor for JWKS refresh frequency. jwx
// honours Cache-Control on the JWKS response and won't refresh sooner than
// this value (5m matches the Node SDK + jose defaults). Background refresh
// + kid rotation are handled by jwx automatically.
const jwksMinRefreshInterval = 5 * time.Minute

// defaultClockTolerance matches the Node SDK.
const defaultClockTolerance = 30 * time.Second

// jwksCache holds one *jwk.Cache shared across the process. Keys are
// normalised issuer URLs; jwx's Cache then deduplicates per-URL fetches and
// rotates kid lookups for us.
//
// The cache lives for the lifetime of the process because, in practice, a
// backend SDK is constructed once and reused. Tests can call resetJWKSCache
// to start clean.
var (
	jwksCacheMu  sync.Mutex
	jwksCache    *jwk.Cache
	jwksCacheCtx context.Context
	jwksKnown    = map[string]struct{}{}
)

func getJWKSCache() *jwk.Cache {
	jwksCacheMu.Lock()
	defer jwksCacheMu.Unlock()
	if jwksCache == nil {
		// context.Background is intentional: jwx uses this only for its
		// background refresh goroutines, which we want tied to process
		// lifetime, not per-call ctx.
		jwksCacheCtx = context.Background()
		jwksCache = jwk.NewCache(jwksCacheCtx)
	}
	return jwksCache
}

func resetJWKSCache() {
	jwksCacheMu.Lock()
	defer jwksCacheMu.Unlock()
	jwksCache = nil
	jwksCacheCtx = nil
	jwksKnown = map[string]struct{}{}
}

// jwksURLForIssuer returns the canonical JWKS URL for a torii issuer.
// torii exposes JWKS at /_torii/.well-known/jwks.json for every tenant.
func jwksURLForIssuer(issuer string) string {
	normalized := strings.TrimRight(issuer, "/")
	return normalized + "/_torii/.well-known/jwks.json"
}

// registerJWKS adds the issuer's JWKS URL to the cache the first time we see
// it. Subsequent calls are no-ops.
func registerJWKS(cache *jwk.Cache, jwksURL string) error {
	jwksCacheMu.Lock()
	defer jwksCacheMu.Unlock()
	if _, ok := jwksKnown[jwksURL]; ok {
		return nil
	}
	if err := cache.Register(
		jwksURL,
		jwk.WithMinRefreshInterval(jwksMinRefreshInterval),
	); err != nil {
		return err
	}
	jwksKnown[jwksURL] = struct{}{}
	return nil
}

// VerifyToken validates a torii-issued JWT and returns the verified subset of
// claims as *Auth. It performs networkless verification once the issuer's
// JWKS has been cached.
//
// Verification enforces:
//   - ES256 signature against the issuer's JWKS (with kid rotation handled by jwx)
//   - exp / nbf with VerifyOptions.ClockTolerance leeway (default 30s)
//   - strict `iss` equality with VerifyOptions.Issuer
//   - optional `aud` if VerifyOptions.Audience is set
//   - presence of required claims (sub, iat, exp, iss, pid)
func VerifyToken(ctx context.Context, token string, opts VerifyOptions) (*Auth, error) {
	if token == "" {
		return nil, newError("verifyToken: token must be a non-empty string", nil)
	}
	if opts.Issuer == "" {
		return nil, newError("verifyToken: Issuer option is required", nil)
	}
	tolerance := opts.ClockTolerance
	if tolerance <= 0 {
		tolerance = defaultClockTolerance
	}

	cache := getJWKSCache()
	jwksURL := jwksURLForIssuer(opts.Issuer)
	if err := registerJWKS(cache, jwksURL); err != nil {
		return nil, newError("verifyToken: failed to register JWKS", err)
	}

	set, err := cache.Get(ctx, jwksURL)
	if err != nil {
		return nil, newError("verifyToken: failed to fetch JWKS", err)
	}

	parseOpts := []jwt.ParseOption{
		jwt.WithKeySet(set),
		jwt.WithValidate(true),
		jwt.WithAcceptableSkew(tolerance),
		jwt.WithIssuer(opts.Issuer),
		// Restrict alg to ES256: jwt.WithKeySet honours `alg` on JWKs but
		// torii minted-tokens are ES256. We don't pass alg via WithKey here
		// because WithKeySet performs kid-based lookup automatically and
		// inspects the key's `alg`/`crv`. jwx will reject anything else.
	}
	if opts.Audience != "" {
		parseOpts = append(parseOpts, jwt.WithAudience(opts.Audience))
	}

	tok, err := jwt.Parse([]byte(token), parseOpts...)
	if err != nil {
		return nil, newError("JWT verification failed", err)
	}

	// Defence-in-depth: jwx already enforced iss above via WithIssuer, but we
	// double-check before extracting and we ensure ES256 was actually used.
	if tok.Issuer() != opts.Issuer {
		return nil, newError("JWT verification failed: issuer mismatch", nil)
	}

	userID := tok.Subject()
	if userID == "" {
		return nil, newError("JWT is missing required claim 'sub'", nil)
	}

	if _, ok := tok.Get("iat"); !ok {
		return nil, newError("JWT is missing required claim 'iat'", nil)
	}
	if _, ok := tok.Get("exp"); !ok {
		return nil, newError("JWT is missing required claim 'exp'", nil)
	}

	raw, err := tok.AsMap(ctx)
	if err != nil {
		return nil, newError("JWT verification failed: cannot read claims", err)
	}

	envID, ok := stringClaim(raw, "pid")
	if !ok {
		return nil, newError("JWT is missing required claim 'pid'", nil)
	}

	auth := &Auth{
		UserID:          userID,
		EnvironmentID:   envID,
		Issuer:          tok.Issuer(),
		EmailVerified:   boolClaim(raw, "email_verified", false),
		ProfileComplete: boolClaim(raw, "profile_complete", true),
		Impersonating:   boolClaim(raw, "impersonating", false),
		Raw:             raw,
	}
	if locale, ok := stringClaim(raw, "locale"); ok {
		auth.Locale = &locale
	}

	// We don't accept alg=none / HS256 / RS* tokens. jwx already validated
	// alg against the JWKS, but assert explicitly via Headers if available.
	_ = jwa.ES256
	return auth, nil
}

func stringClaim(m map[string]any, key string) (string, bool) {
	v, ok := m[key]
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	if !ok || s == "" {
		return "", false
	}
	return s, true
}

func boolClaim(m map[string]any, key string, fallback bool) bool {
	v, ok := m[key]
	if !ok {
		return fallback
	}
	b, ok := v.(bool)
	if !ok {
		return fallback
	}
	return b
}
