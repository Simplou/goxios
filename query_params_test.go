package goxios

import "testing"

func TestSetQueryParams(t *testing.T) {
	queryParams := []QueryParam{
		{Key: "name", Value: "Gabriel"},
		{Key: "id", Value: 123},
	}
	url := "https://example.com/query"
	expected := url + "?name=Gabriel&id=123"
	url = setQueryParams(queryParams, url)
	if url != expected {
		t.Fatalf("Expected: %s, Result: %s", expected, url)
	}
}
