package goxios

import (
	"encoding/json"
	"io"
)

type JSON map[string]interface{}

func (j JSON) Marshal() ([]byte, error) {
	return json.Marshal(j)
}

func DecodeJSON(body io.Reader, v any) error {
	return json.NewDecoder(body).Decode(v)
}

func UnmarshalJSON(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
