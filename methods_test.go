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
			return
		case http.MethodPost:
			defer r.Body.Close()
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			msg := fmt.Sprintf("{'message':'%s created'}", string(b))
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(msg))
			return
		case http.MethodPut:
			defer r.Body.Close()
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			msg := fmt.Sprintf("{'message':'%s updated'}", string(b))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(msg))
		case http.MethodPatch:
			defer r.Body.Close()
			b, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			msg := fmt.Sprintf("{'message':'%s updated'}", string(b))
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(msg))
		case http.MethodDelete:
			w.WriteHeader(http.StatusNoContent)
			return
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
		method string
		url    string
		body   []byte
	}{
		{http.MethodGet, ts.URL + "/get", nil},
		{http.MethodPost, ts.URL + "/post", b},
		{http.MethodPut, ts.URL + "/put", b},
		{http.MethodPatch, ts.URL + "/patch", b},
		{http.MethodDelete, ts.URL + "/delete", b},
	}

	for _, tm := range testMethods {
		var res *http.Response
		var err error
		contentType := Header{Key: "Content-Type", Value: "application/json"}
		switch tm.method {
		case http.MethodGet:
			res, err = client.Get(tm.url, []Header{})
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("Expected status code 200, got %d", res.StatusCode)
			}
		case http.MethodPost:
			res, err = client.Post(tm.url, []Header{contentType}, bytes.NewBuffer(tm.body))
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode != http.StatusCreated {
				t.Errorf("Expected status code %d, got %d", http.StatusCreated, res.StatusCode)
			}
		case http.MethodPut:
			res, err := client.Put(tm.url, []Header{contentType}, bytes.NewBuffer(b))
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("Expected status code 200, got %d", res.StatusCode)
			}
		case http.MethodPatch:
			res, err := client.Patch(tm.url, []Header{contentType}, bytes.NewBuffer(b))
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("Expected status code 200, got %d", res.StatusCode)
			}
		case http.MethodDelete:
			res, err := client.Delete(tm.url, []Header{contentType}, bytes.NewBuffer(b))
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode != http.StatusNoContent {
				t.Errorf("Expected status code %d, got %d", http.StatusNoContent, res.StatusCode)
			}
		}
	}
}
