package torii

import "net/http"

// VerifyWebhook will, in a future release, verify a torii outbound webhook
// signature and decode the event payload. Today it is a stub — torii's
// outbound webhook subsystem is not yet available.
//
// The signature is published now so that adopting this function later does
// not break callers.
func VerifyWebhook(secret string, headers http.Header, payload []byte) (*WebhookEvent, error) {
	_ = secret
	_ = headers
	_ = payload
	return nil, newError(
		"verifyWebhook: torii's outbound webhook subsystem is not yet available.",
		nil,
	)
}
