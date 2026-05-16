package torii

import "net/http"

// VerifyWebhook will, in a future release, verify a torii outbound webhook
// signature and decode the event payload. Today it is a stub — torii's
// webhook subsystem is being designed under #424 Phase 0.5.
//
// The signature is published now so that adopting this function later does
// not break callers. Once Phase 0.5 ships with the final signing scheme
// (Svix-compatible HMAC or homegrown), this body becomes the real verifier.
func VerifyWebhook(secret string, headers http.Header, payload []byte) (*WebhookEvent, error) {
	_ = secret
	_ = headers
	_ = payload
	return nil, newError(
		"verifyWebhook: torii's outbound webhook subsystem has not shipped yet — see #424 Phase 0.5",
		nil,
	)
}
