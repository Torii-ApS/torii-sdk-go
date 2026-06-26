package torii

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// sampleUserJSON is a minimal-but-valid ServerUserResponse (all generator-
// required fields present) so decode succeeds and tests can focus on the
// outbound request.
const sampleUserJSON = `{
	"id": "00000000-0000-0000-0000-000000000001",
	"environmentId": "00000000-0000-0000-0000-000000000002",
	"status": "active",
	"createdAt": "2024-01-01T00:00:00Z",
	"updatedAt": "2024-01-01T00:00:00Z",
	"publicMetadata": {},
	"privateMetadata": {},
	"unsafeMetadata": {}
}`

// The secret key must reach the wire as `Authorization: Bearer <key>` on every
// call — both the generated path (Create) and the hand-rolled doJSON path
// (Update). And create must omit unset metadata bags so the server defaults
// them to {}.
func TestClient_SendsBearerAuth(t *testing.T) {
	var gotAuth, gotMethod string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		gotMethod = r.Method
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(sampleUserJSON))
	}))
	defer srv.Close()

	client, err := New(Options{SecretKey: "sk_test_abc", APIURL: srv.URL})
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	ctx := context.Background()

	// Create — generated client, bearerAuth scheme.
	email := "ada@example.com"
	if _, err := client.Users().Create(ctx, CreateUserInput{Email: &email}); err != nil {
		t.Fatalf("Create: %v", err)
	}
	if gotAuth != "Bearer sk_test_abc" {
		t.Errorf("create Authorization = %q, want %q", gotAuth, "Bearer sk_test_abc")
	}
	if gotMethod != http.MethodPost {
		t.Errorf("create method = %q, want POST", gotMethod)
	}
	var createBody map[string]any
	if err := json.Unmarshal(gotBody, &createBody); err != nil {
		t.Fatalf("create body not JSON: %v", err)
	}
	for _, bag := range []string{"publicMetadata", "privateMetadata", "unsafeMetadata"} {
		if _, present := createBody[bag]; present {
			t.Errorf("create body should omit unset %s; got %v", bag, createBody)
		}
	}
	if createBody["email"] != "ada@example.com" {
		t.Errorf("create body email = %v, want ada@example.com", createBody["email"])
	}

	// Update — hand-rolled doJSON path.
	gotAuth = ""
	if _, err := client.Users().Update(ctx, "00000000-0000-0000-0000-000000000001", UpdateUserInput{FirstName: SetPatch("Ada")}); err != nil {
		t.Fatalf("Update: %v", err)
	}
	if gotAuth != "Bearer sk_test_abc" {
		t.Errorf("update Authorization = %q, want %q", gotAuth, "Bearer sk_test_abc")
	}
}
