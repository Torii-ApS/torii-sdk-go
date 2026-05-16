// Package torii is the official Go SDK for torii (https://torii.so).
//
// Use it to:
//   - Verify end-user JWTs minted by torii without a per-request round trip
//     (see VerifyToken / AuthenticateRequest).
//   - Call the torii backend API to manage users and sessions (see New, Client).
//   - React to outbound webhooks (VerifyWebhook — currently a stub awaiting #424
//     Phase 0.5).
package torii

import (
	"fmt"
	"net/http"
	"time"
)

// Options configures the torii backend API Client.
type Options struct {
	// SecretKey is the backend secret (e.g. sk_live_... / sk_test_...). Required.
	SecretKey string

	// APIURL overrides the default backend API base URL.
	// Defaults to "https://api.torii.so".
	APIURL string

	// HTTPClient lets callers inject a custom *http.Client (timeouts, transport,
	// proxies, etc). When nil, http.DefaultClient is used.
	HTTPClient *http.Client
}

// VerifyOptions configures JWT verification.
type VerifyOptions struct {
	// Issuer is the expected JWT `iss` claim — required and strictly enforced.
	// For torii this is the FAPI URL for the environment, e.g.
	// "https://acme.torii.so" or a verified custom domain.
	Issuer string

	// Audience optionally enforces the JWT `aud` claim. torii does not set
	// `aud` today, so leaving this empty skips the check. Reserved for
	// future-compat.
	Audience string

	// ClockTolerance is the leeway applied to exp/nbf checks.
	// Defaults to 30 seconds when zero.
	ClockTolerance time.Duration
}

// Auth is the subset of fields the SDK exposes from a verified torii access
// token. Callers who need raw claims can read Raw.
type Auth struct {
	// UserID is the end-user ID (JWT `sub`).
	UserID string
	// EnvironmentID is the environment this token was issued in (JWT `pid`).
	EnvironmentID string
	// Issuer is the JWT `iss` claim — canonical FAPI URL for this environment.
	Issuer string
	// EmailVerified is true if at least one of the end-user's emails is verified.
	EmailVerified bool
	// ProfileComplete is true if all environment-required profile fields are
	// filled. Defaults to true if the claim is absent.
	ProfileComplete bool
	// Impersonating is true if the token is being used for admin impersonation.
	Impersonating bool
	// Locale is the end-user's preferred locale when set on the profile, else nil.
	Locale *string
	// Raw is the full JWT payload — escape hatch for custom claims, audience
	// checks, etc.
	Raw map[string]any
}

// Error is the SDK's auth / verification error type.
// REST errors from the backend API surface as APIError.
type Error struct {
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

func (e *Error) Unwrap() error { return e.Cause }

func newError(message string, cause error) *Error {
	return &Error{Message: message, Cause: cause}
}

// APIError wraps a non-2xx response from the torii backend API.
type APIError struct {
	Status    int
	Code      string
	SupportID string
	Body      []byte
	Message   string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("torii API error %d: %s", e.Status, e.Message)
	}
	return fmt.Sprintf("torii API error %d", e.Status)
}

// WebhookEvent is the shape of a verified outbound webhook payload.
// The fields are speculative until torii's webhook subsystem ships (#424
// Phase 0.5); they're declared here so that adopting VerifyWebhook later
// doesn't break callers that already type their handler around this struct.
type WebhookEvent struct {
	Type      string
	ID        string
	CreatedAt string
	Data      map[string]any
}
