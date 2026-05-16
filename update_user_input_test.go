package torii

import (
	"encoding/json"
	"testing"
)

func TestUpdateUserInput_asJSONBody(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input UpdateUserInput
		want  map[string]any
	}{
		{
			name:  "all omitted yields empty object",
			input: UpdateUserInput{},
			want:  map[string]any{},
		},
		{
			name: "set string field",
			input: UpdateUserInput{
				Name: SetPatch("Alice"),
			},
			want: map[string]any{"name": "Alice"},
		},
		{
			name: "clear string field sends null",
			input: UpdateUserInput{
				Name: ClearPatch[string](),
			},
			want: map[string]any{"name": nil},
		},
		{
			name: "empty string is preserved (set, not omitted)",
			input: UpdateUserInput{
				Name: SetPatch(""),
			},
			want: map[string]any{"name": ""},
		},
		{
			name: "mixed set/clear/omit across fields",
			input: UpdateUserInput{
				Name:        SetPatch("Bob"),
				Phone:       ClearPatch[string](),
				AvatarUrl:   SetPatch("https://example.com/a.png"),
				Locale:      SetPatch("en"),
				Address:     ClearPatch[string](),
				DateOfBirth: SetPatch("1990-01-01"),
			},
			want: map[string]any{
				"name":        "Bob",
				"phone":       nil,
				"avatarUrl":   "https://example.com/a.png",
				"locale":      "en",
				"address":     nil,
				"dateOfBirth": "1990-01-01",
			},
		},
		{
			name: "only locale set",
			input: UpdateUserInput{
				Locale: SetPatch("da"),
			},
			want: map[string]any{"locale": "da"},
		},
		{
			name: "only date of birth cleared",
			input: UpdateUserInput{
				DateOfBirth: ClearPatch[string](),
			},
			want: map[string]any{"dateOfBirth": nil},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotBytes, err := tt.input.asJSONBody()
			if err != nil {
				t.Fatalf("asJSONBody() error = %v", err)
			}
			var got map[string]any
			if err := json.Unmarshal(gotBytes, &got); err != nil {
				t.Fatalf("could not unmarshal output: %v (raw: %s)", err, gotBytes)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("size mismatch: got %d keys, want %d keys (got=%v, want=%v)",
					len(got), len(tt.want), got, tt.want)
			}
			for k, v := range tt.want {
				gv, ok := got[k]
				if !ok {
					t.Errorf("missing key %q in output %v", k, got)
					continue
				}
				if gv != v {
					t.Errorf("key %q: got %v (%T), want %v (%T)", k, gv, gv, v, v)
				}
			}
			for k := range got {
				if _, ok := tt.want[k]; !ok {
					t.Errorf("unexpected extra key %q in output", k)
				}
			}
		})
	}
}

func TestPatch_IsOmitted(t *testing.T) {
	t.Parallel()
	var zero Patch[string]
	if !zero.IsOmitted() {
		t.Errorf("zero-value Patch[string] should be omitted")
	}
	if SetPatch("x").IsOmitted() {
		t.Errorf("SetPatch should not be omitted")
	}
	if ClearPatch[string]().IsOmitted() {
		t.Errorf("ClearPatch should not be omitted")
	}
}

func TestPatch_MarshalJSON(t *testing.T) {
	t.Parallel()
	// Set
	b, err := SetPatch("hi").MarshalJSON()
	if err != nil {
		t.Fatalf("SetPatch MarshalJSON err: %v", err)
	}
	if string(b) != `"hi"` {
		t.Errorf("SetPatch MarshalJSON = %q, want %q", b, `"hi"`)
	}
	// Clear
	b, err = ClearPatch[string]().MarshalJSON()
	if err != nil {
		t.Fatalf("ClearPatch MarshalJSON err: %v", err)
	}
	if string(b) != "null" {
		t.Errorf("ClearPatch MarshalJSON = %q, want %q", b, "null")
	}
	// Omitted — emits nil bytes (caller must use IsOmitted to skip).
	var zero Patch[string]
	b, err = zero.MarshalJSON()
	if err != nil {
		t.Fatalf("Omitted MarshalJSON err: %v", err)
	}
	if b != nil {
		t.Errorf("Omitted MarshalJSON = %q, want nil", b)
	}
}
