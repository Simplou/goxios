package goxios

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// newRequest creates a new HTTP request with the given method, URL, and body.
// It attaches the provided context to the request.
func newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}

// setHeaders sets custom headers on the given HTTP request.
// It takes a request pointer req and a slice of Header structs.
// Each Header struct contains a key-value pair representing a header field and its value.
func setHeaders(req *http.Request, headers []Header) {
	for _, header := range headers {
		req.Header.Set(header.Key, fmt.Sprintf("%v", header.Value))
	}
}

// client represents an HTTP client.
// It embeds an http.Client pointer, a request pointer, and a context.
type client struct {
	*http.Client
	req *http.Request
	ctx context.Context
}

// clientV2 represents an extended version of the HTTP client.
// It embeds a client pointer and an array of Header pointers.
type clientV2 struct {
	headers []Header
	*client
}

// SetHeaders sets custom headers on the client.
// It takes one or more Header pointers as arguments and assigns them to the client's headers field.
// It returns the slice of Header pointers provided.
func (v2 *clientV2) SetHeaders(headers ...Header) []Header {
	if len(headers) > 0 {
		v2.headers = headers
		setHeaders(v2.req, headers)
	}
	return headers
}

// RequestUrl returns the URL of the current request.
// It fetches the URL from the underlying request object and returns it as a string.
func (v2 *clientV2) RequestUrl() string {
	return v2.req.URL.String()
}

// New creates and returns a new HTTP client with the provided context.
// It initializes and returns a new clientV2 pointer with an embedded http.Client pointer,
// an empty request pointer, and the provided context.
func New(ctx context.Context) *clientV2 {
	return &clientV2{
		client: NewClient(ctx),
	}
}

// NewClient creates a new HTTP client with the given context.
// It initializes and returns a new client pointer with an embedded http.Client pointer,
// an empty request pointer, and the provided context.
func NewClient(ctx context.Context) *client {
	return &client{
		Client: &http.Client{},
		req:    &http.Request{Header: make(http.Header)},
		ctx:    ctx,
	}
}

// SetTimeout sets the timeout for the HTTP client.
// If a timeout is set, the client will abort the request
// if the server does not respond within the specified duration.
// It takes a duration argument and sets it as the timeout for the client.
func (c *client) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}

// Context returns the context associated with the client.
// It retrieves and returns the context stored within the client.
func (c *client) Context() context.Context {
	return c.ctx
}

// Response performs the HTTP request and returns the response.
// It sends the request associated with the client and returns the response received from the server.
func (c *client) Response() (*http.Response, error) {
	return c.Do(c.req)
}
