package torii

import "encoding/json"

// UpdateUserInput is the request body for Users.Update. Each field uses
// the tri-state Patch[T any] wrapper so callers can express "leave the
// field unchanged" (zero value), "set the field to v" (SetPatch(v)) and
// "clear the field" (ClearPatch[T]()).
type UpdateUserInput struct {
	FirstName Patch[string]
	LastName  Patch[string]
	Locale    Patch[string]
	// UnsafeMetadata is tri-state: omit to leave the server's metadata untouched,
	// set to replace it, clear to null it. Never silently sent as empty.
	UnsafeMetadata Patch[map[string]any]
}

// asJSONBody renders the input to a JSON object containing only the
// fields that are NOT omitted. Set fields appear with their value; Clear
// fields appear with null; Omitted fields are absent from the object.
func (u UpdateUserInput) asJSONBody() ([]byte, error) {
	m := map[string]any{}
	if !u.FirstName.IsOmitted() {
		m["firstName"] = jsonValue(u.FirstName)
	}
	if !u.LastName.IsOmitted() {
		m["lastName"] = jsonValue(u.LastName)
	}
	if !u.Locale.IsOmitted() {
		m["locale"] = jsonValue(u.Locale)
	}
	if !u.UnsafeMetadata.IsOmitted() {
		m["unsafeMetadata"] = jsonValue(u.UnsafeMetadata)
	}
	return json.Marshal(m)
}

func jsonValue[T any](p Patch[T]) any {
	if p.state == patchClear {
		return nil
	}
	return p.value
}
