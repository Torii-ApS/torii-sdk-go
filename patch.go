package torii

import "encoding/json"

// Patch is a tri-state wrapper for PATCH body fields. The zero value
// (state==omitted) means "leave field alone"; SetPatch and ClearPatch
// produce "update to v" and "clear" respectively.
type Patch[T any] struct {
	state patchState
	value T
}

type patchState uint8

const (
	patchOmitted patchState = iota
	patchSet
	patchClear
)

// SetPatch returns a Patch that updates the field to v.
func SetPatch[T any](v T) Patch[T] { return Patch[T]{state: patchSet, value: v} }

// ClearPatch returns a Patch that clears the field (sends JSON null).
func ClearPatch[T any]() Patch[T] { return Patch[T]{state: patchClear} }

// IsOmitted reports whether the field should be omitted from the request body.
func (p Patch[T]) IsOmitted() bool { return p.state == patchOmitted }

// MarshalJSON: Patch[T] cannot be the top-level body — it's a field-level
// wrapper consumed by the SDK's update() implementation. This method
// exists so misuse is loud (it shouldn't be serialised directly).
func (p Patch[T]) MarshalJSON() ([]byte, error) {
	switch p.state {
	case patchClear:
		return []byte("null"), nil
	case patchSet:
		return json.Marshal(p.value)
	default: // patchOmitted
		return nil, nil
	}
}
