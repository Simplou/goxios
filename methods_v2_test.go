package goxios

import (
	"bytes"
	"context"
	"net/http"
	"testing"
)

func TestV2ClientMethods(t *testing.T) {
	ts := getTestServer(t)
	defer ts.Close()

	client := New(context.Background())

	requestJSON := JSON{
		"username": "gabriel",
	}
	b, err := requestJSON.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	testMethods := ts.Methods(b)
	for _, tm := range testMethods {
		var res *http.Response
		var err error
		contentType := Header{Key: "Content-Type", Value: "application/json"}
		requestOptions := &RequestOpts{
			Headers: []Header{contentType},
			Body:    bytes.NewBuffer(tm.body),
		}
		switch tm.method {
		case http.MethodGet:
			res, err = client.Get(tm.url, &RequestOpts{Headers: []Header{}})
		case http.MethodPost:
			res, err = client.Post(tm.url, requestOptions)
		case http.MethodPut:
			res, err = client.Put(tm.url, requestOptions)
		case http.MethodPatch:
			res, err = client.Patch(tm.url, requestOptions)
		case http.MethodDelete:
			res, err = client.Delete(tm.url, requestOptions)
		}
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != tm.expectedStatus {
			t.Errorf("Expected status code %d, got %d", tm.expectedStatus, res.StatusCode)
		}
	}
}
