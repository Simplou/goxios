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

// DecodeJSON decodes JSON data from the provided io.Reader into the given value.
func DecodeJSON(body io.Reader, v any) error {
	return json.NewDecoder(body).Decode(v)
}

// UnmarshalJSON parses the JSON-encoded data and stores the result in the value pointed to by v.
func UnmarshalJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
