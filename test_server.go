package goxios

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type testServer struct {
	*httptest.Server
}

type testMethods struct {
	method         string
	url            string
	body           []byte
	expectedStatus int
}

func (ts *testServer) Methods(b []byte) []testMethods {
	testMethods := []testMethods{
		{http.MethodGet, ts.URL + "/get", nil, http.StatusOK},
		{http.MethodPost, ts.URL + "/post", b, http.StatusCreated},
		{http.MethodPut, ts.URL + "/put", b, http.StatusOK},
		{http.MethodPatch, ts.URL + "/patch", b, http.StatusOK},
		{http.MethodDelete, ts.URL + "/delete", b, http.StatusNoContent},
	}
	return testMethods
}

func getTestServer(t *testing.T) *testServer {
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
	return &testServer{ts}
}

func getTestServerWithTimeout(t *testing.T, timeout time.Duration) *testServer {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			time.Sleep(timeout)
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
				time.Sleep(timeout)
				msg = fmt.Sprintf("{'message':'%s created'}", string(b))
				w.WriteHeader(http.StatusCreated)
			case http.MethodPut, http.MethodPatch:
				time.Sleep(timeout)
				msg = fmt.Sprintf("{'message':'%s updated'}", string(b))
				w.WriteHeader(http.StatusOK)
			}
			w.Write([]byte(msg))
		case http.MethodDelete:
			time.Sleep(timeout)
			w.WriteHeader(http.StatusNoContent)
		default:
			time.Sleep(timeout)
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	return &testServer{ts}
}
