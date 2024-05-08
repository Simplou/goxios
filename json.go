package goxios

import (
	"encoding/json"
	"io"
)

// JSON represents a JSON object.
type JSON map[string]interface{}

// Marshal converts the JSON object to its JSON representation.
func (j JSON) Marshal() ([]byte, error) {
	return json.Marshal(j)
}

func (j JSON) Unmarshal(v any) error {
	b, err := j.Marshal()
	if err != nil {
		return err
	}
	return UnmarshalJSON(b, v)
}

// GenericJSON represents a generic JSON object with keys of type string and values of any type.
// It allows flexibility in defining JSON objects with various value types.
type GenericJSON[T any] map[string]T

// Marshal converts the GenericJSON object to its JSON representation.
func (gj GenericJSON[T]) Marshal() ([]byte, error) {
	return json.Marshal(gj)
}

func (gj GenericJSON[T]) Unmarshal(v any) error {
	b, err := gj.Marshal()
	if err != nil {
		return err
	}
	return UnmarshalJSON(b, v)
}

// DecodeJSON decodes JSON data from the provided io.Reader into the given value.
func DecodeJSON(body io.Reader, v any) error {
	return json.NewDecoder(body).Decode(v)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result in the value pointed to by v.
func UnmarshalJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
