package goxios

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientMethods(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{'message':'Welcome'}"))
		case http.MethodPost, http.MethodPut, http.MethodPatch:
			defer r.Body.Close()
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			var msg string
			switch r.Method {
			case http.MethodPost:
				msg = fmt.Sprintf("{'message':'%s created'}", string(b))
				w.WriteHeader(http.StatusCreated)
			case http.MethodPut, http.MethodPatch:
				msg = fmt.Sprintf("{'message':'%s updated'}", string(b))
				w.WriteHeader(http.StatusOK)
			}
			w.Write([]byte(msg))
		case http.MethodDelete:
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	defer ts.Close()

	client := NewClient(context.Background())

	requestJSON := JSON{
		"username": "gabriel",
	}
	b, err := requestJSON.Marshal()
	if err != nil {
		t.Fatal(err)
	}

	testMethods := []struct {
		method         string
		url            string
		body           []byte
		expectedStatus int
	}{
		{http.MethodGet, ts.URL + "/get", nil, http.StatusOK},
		{http.MethodPost, ts.URL + "/post", b, http.StatusCreated},
		{http.MethodPut, ts.URL + "/put", b, http.StatusOK},
		{http.MethodPatch, ts.URL + "/patch", b, http.StatusOK},
		{http.MethodDelete, ts.URL + "/delete", b, http.StatusNoContent},
	}

	for _, tm := range testMethods {
		var res *http.Response
		var err error
		contentType := Header{Key: "Content-Type", Value: "application/json"}
		switch tm.method {
		case http.MethodGet:
			res, err = client.Get(tm.url, []Header{})
		case http.MethodPost:
			res, err = client.Post(tm.url, []Header{contentType}, bytes.NewBuffer(tm.body))
		case http.MethodPut:
			res, err = client.Put(tm.url, []Header{contentType}, bytes.NewBuffer(tm.body))
		case http.MethodPatch:
			res, err = client.Patch(tm.url, []Header{contentType}, bytes.NewBuffer(tm.body))
		case http.MethodDelete:
			res, err = client.Delete(tm.url, []Header{contentType}, bytes.NewBuffer(tm.body))
		}
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != tm.expectedStatus {
			t.Errorf("Expected status code %d, got %d", tm.expectedStatus, res.StatusCode)
		}
	}
}
