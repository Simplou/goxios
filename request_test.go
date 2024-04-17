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
