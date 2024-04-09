package goxios

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}

func setHeaders(req *http.Request, headers []Header) {
	for _, header := range headers {
		req.Header.Set(header.Key, fmt.Sprintf("%v", header.Value))
	}
}

type client struct {
	*http.Client
	req *http.Request
	ctx context.Context
}

func NewClient(ctx context.Context) *client {
	return &client{
		Client: &http.Client{},
		req:    &http.Request{},
		ctx:    ctx,
	}
}

func (c *client) Context() context.Context {
	return c.ctx
}

func (c *client) Response() (*http.Response, error) {
	return c.Do(c.req)
}
