package goxios

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// newRequest creates a new HTTP request with the given method, URL, and body.
// It attaches the provided context to the request.
func newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}

// setHeaders sets custom headers on the given HTTP request.
func setHeaders(req *http.Request, headers []Header) {
	for _, header := range headers {
		req.Header.Set(header.Key, fmt.Sprintf("%v", header.Value))
	}
}

// client represents an HTTP client.
type client struct {
	*http.Client
	req *http.Request
	ctx context.Context
}

// NewClient creates a new HTTP client with the given context.
func NewClient(ctx context.Context) *client {
	return &client{
		Client: &http.Client{},
		req:    &http.Request{},
		ctx:    ctx,
	}
}

// Context returns the context associated with the client.
func (c *client) Context() context.Context {
	return c.ctx
}

// Response performs the HTTP request and returns the response.
func (c *client) Response() (*http.Response, error) {
	return c.Do(c.req)
}
