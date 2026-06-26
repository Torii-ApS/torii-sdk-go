package torii

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/Torii-ApS/torii-sdk-go/internal/generated"
)

// TestPatchWireParity pins that the generated request models emit the exact wire
// bytes blessed by the shared contract (contract-tests/fixtures/patch-wire), the
// same fixtures the server round-trip test asserts. For each fixture we unmarshal
// expectedBody into the generated struct and marshal it straight back: the
// round-trip must be identical, which proves the SDK preserves the tri-state
// distinctions (an absent key stays absent => leave, an explicit null stays null
// => clear, a value stays a value => set, and nested nulls survive => key delete).
func TestPatchWireParity(t *testing.T) {
	raw, err := os.ReadFile("testdata/patch_wire_fixtures.json")
	if err != nil {
		t.Fatalf("read fixtures: %v", err)
	}
	var manifest struct {
		Fixtures []struct {
			Name         string          `json:"name"`
			Schema       string          `json:"schema"`
			ExpectedBody json.RawMessage `json:"expectedBody"`
		} `json:"fixtures"`
	}
	if err := json.Unmarshal(raw, &manifest); err != nil {
		t.Fatalf("parse fixtures: %v", err)
	}
	if len(manifest.Fixtures) == 0 {
		t.Fatal("no fixtures loaded")
	}

	for _, f := range manifest.Fixtures {
		t.Run(f.Name, func(t *testing.T) {
			model := newModelForSchema(f.Schema)
			if model == nil {
				t.Fatalf("no generated model registered for schema %q", f.Schema)
			}
			if err := json.Unmarshal(f.ExpectedBody, model); err != nil {
				t.Fatalf("unmarshal into %s: %v", f.Schema, err)
			}
			got, err := json.Marshal(model)
			if err != nil {
				t.Fatalf("marshal %s: %v", f.Schema, err)
			}
			if !jsonEqual(got, f.ExpectedBody) {
				t.Fatalf("wire mismatch for %s:\n  want %s\n  got  %s", f.Name, f.ExpectedBody, got)
			}
		})
	}
}

func newModelForSchema(schema string) any {
	switch schema {
	case "UpdateUserRequest":
		return generated.NewUpdateUserRequest()
	case "CreateUserRequest":
		return generated.NewCreateUserRequest()
	case "ServerUserSearchRequest":
		return generated.NewServerUserSearchRequest()
	case "UpdateUserMetadataRequest":
		return generated.NewUpdateUserMetadataRequest()
	default:
		return nil
	}
}

func jsonEqual(a, b []byte) bool {
	var x, y any
	if json.Unmarshal(a, &x) != nil {
		return false
	}
	if json.Unmarshal(b, &y) != nil {
		return false
	}
	return reflect.DeepEqual(x, y)
}
