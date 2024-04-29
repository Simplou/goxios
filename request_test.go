package goxios

import (
	"context"
	"io"
	"testing"
)

var requestUrl = "https://api.agify.io/?name=michael"

func TestRequestUrl(t *testing.T) {
	client := New(context.Background())
	res, err := client.Get(requestUrl, &RequestOpts{})
	if err != nil {
		t.Fatal(err)
	}
	url := client.RequestUrl()
	if url != requestUrl {
		t.Fatalf("expected: %s received: %s", requestUrl, url)
	}
	defer res.Body.Close()
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(resBytes))
}

func TestSetHeaders(t *testing.T) {
	client := New(context.Background())
	contentType := "application/json"
	headerNotSet := func() {
		t.Fatal("headers not set")
	}
	headers := client.SetHeaders(
		Header{Key: "Content-Type", Value: contentType},
		Header{Key: "Accept", Value: contentType},
	)
	if len(client.headers) != 2 {
		headerNotSet()
	}
	if client.req.Header.Get(headers[0].Key) != contentType {
		headerNotSet()
	}
	if client.req.Header.Get(headers[len(headers)-1].Key) != contentType {
		headerNotSet()
	}
	acceptHeaderKey := "Accept"
	contentType = "application/xml"
	res, err := client.Get(requestUrl, &RequestOpts{Headers: []Header{{Key: acceptHeaderKey, Value: contentType}}})
	if err != nil {
		t.Fatal(err)
	}
	if res.Request.Header.Get(acceptHeaderKey) != contentType {
		headerNotSet()
	}
	if res.Request.Header.Get(headers[0].Key) != headers[0].Value {
		headerNotSet()
	}

}
