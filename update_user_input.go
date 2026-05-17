package torii

import "encoding/json"

// UpdateUserInput is the request body for Users.Update. Each field uses
// the tri-state Patch[T any] wrapper so callers can express "leave the
// field unchanged" (zero value), "set the field to v" (SetPatch(v)) and
// "clear the field" (ClearPatch[T]()).
type UpdateUserInput struct {
	Name        Patch[string]
	Phone       Patch[string]
	Locale      Patch[string]
	Address     Patch[string]
	DateOfBirth Patch[string] // ISO date "YYYY-MM-DD"
}

// asJSONBody renders the input to a JSON object containing only the
// fields that are NOT omitted. Set fields appear with their value; Clear
// fields appear with null; Omitted fields are absent from the object.
func (u UpdateUserInput) asJSONBody() ([]byte, error) {
	m := map[string]any{}
	if !u.Name.IsOmitted() {
		m["name"] = jsonValue(u.Name)
	}
	if !u.Phone.IsOmitted() {
		m["phone"] = jsonValue(u.Phone)
	}
	if !u.Locale.IsOmitted() {
		m["locale"] = jsonValue(u.Locale)
	}
	if !u.Address.IsOmitted() {
		m["address"] = jsonValue(u.Address)
	}
	if !u.DateOfBirth.IsOmitted() {
		m["dateOfBirth"] = jsonValue(u.DateOfBirth)
	}
	return json.Marshal(m)
}

func jsonValue[T any](p Patch[T]) any {
	if p.state == patchClear {
		return nil
	}
	return p.value
}
