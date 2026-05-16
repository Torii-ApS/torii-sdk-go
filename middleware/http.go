// Package middleware provides a net/http middleware that verifies a torii
// JWT on incoming requests and stores the resulting *torii.Auth in the
// request context.
//
// Stays framework-agnostic on purpose: anything that consumes
// http.Handler (chi, gorilla/mux, the stdlib) can use it directly. Adapters
// for echo, gin, and fiber will land later (see PR #436 follow-ups).
package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	torii "github.com/GOOD-Code-ApS/torii-sdk-go"
)

// authContextKeyType is unexported to make Go's compile-time type system
// guarantee that no other package can collide on the same context key.
type authContextKeyType struct{}

// AuthContextKey is the context key under which Middleware stores the
// verified *torii.Auth. Use AuthFromContext to read it back; we expose the
// key as well for callers that want raw context.Value access.
var AuthContextKey = authContextKeyType{}

// ErrorWriter renders an authentication failure to the response.
// Defaults to a 401 with a small JSON body; override via Options.OnError.
type ErrorWriter func(w http.ResponseWriter, r *http.Request, err error)

// Options configures Middleware.
type Options struct {
	// Verify holds the JWT verification settings (issuer required).
	Verify torii.VerifyOptions
	// OnError, when non-nil, is called on verification failure instead of
	// the default 401 response.
	OnError ErrorWriter
	// Optional: skip middleware for requests where this returns true.
	// Useful for /health, /metrics, etc.
	Skip func(*http.Request) bool
}

// Middleware returns net/http middleware that verifies the Authorization
// header on every request and stores the resulting *torii.Auth in the
// request context under AuthContextKey.
func Middleware(opts Options) func(http.Handler) http.Handler {
	if opts.OnError == nil {
		opts.OnError = defaultOnError
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if opts.Skip != nil && opts.Skip(r) {
				next.ServeHTTP(w, r)
				return
			}
			auth, err := torii.AuthenticateRequest(r.Context(), r.Header, opts.Verify)
			if err != nil {
				opts.OnError(w, r, err)
				return
			}
			ctx := context.WithValue(r.Context(), AuthContextKey, auth)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// AuthFromContext reads the verified *torii.Auth stashed by Middleware.
// The second return value is false when no auth is present (middleware not
// run, or skipped for this route).
func AuthFromContext(ctx context.Context) (*torii.Auth, bool) {
	a, ok := ctx.Value(AuthContextKey).(*torii.Auth)
	return a, ok
}

func defaultOnError(w http.ResponseWriter, _ *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("WWW-Authenticate", `Bearer realm="torii", error="invalid_token"`)
	w.WriteHeader(http.StatusUnauthorized)
	// Deliberately *don't* leak err.Error() in case it contains JWT
	// validation specifics that help attackers. Log it on the server side
	// instead — callers can override OnError to surface more.
	_ = err
	_ = json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
}
