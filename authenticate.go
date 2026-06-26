package torii

import (
	"context"
	"net/http"
	"strings"
)

// AuthenticateRequest reads a Bearer token from the request headers and
// verifies it. Returns *Auth on success, *Error on failure.
//
// The header read is "Authorization" by default. Gateways that forward the
// token in a different header should call VerifyToken directly.
func AuthenticateRequest(ctx context.Context, headers http.Header, opts VerifyOptions) (*Auth, error) {
	if headers == nil {
		return nil, newError("authenticateRequest: headers is nil", nil)
	}
	raw := headers.Get("Authorization")
	if raw == "" {
		return nil, newError("authenticateRequest: missing Authorization header", nil)
	}
	// Match "Bearer <token>" case-insensitively. http.Header.Get is already
	// canonical-cased on the *key* but the *value* still needs trimming.
	const prefix = "bearer "
	if len(raw) <= len(prefix) || !strings.EqualFold(raw[:len(prefix)], prefix) {
		return nil, newError("authenticateRequest: Authorization header is not in 'Bearer <token>' form", nil)
	}
	token := strings.TrimSpace(raw[len(prefix):])
	if token == "" {
		return nil, newError("authenticateRequest: Authorization header is not in 'Bearer <token>' form", nil)
	}
	return VerifyToken(ctx, token, opts)
}
