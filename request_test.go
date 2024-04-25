package goxios

import (
	"context"
	"io"
	"testing"
)

func TestRequestUrl(t *testing.T) {
	client := New(context.Background())
	requestUrl := "https://api.agify.io/?name=michael"
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
	headerNotSet := func(){
		t.Fatal("headers not set")	
	}
	headers := client.SetHeaders(
		Header{Key: "Content-Type", Value: contentType},
		Header{Key: "Accept", Value: contentType},
	)
	if len(client.headers) != 2{
		headerNotSet()
	}
	if client.req.Header.Get(headers[0].Key) != contentType{
		headerNotSet()
	}
	if client.req.Header.Get(headers[len(headers)-1].Key) != contentType{
		headerNotSet()
	}
}
