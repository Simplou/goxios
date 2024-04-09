package goxios

import (
	"io"
	"net/http"
)

// Post sends a POST request to the specified URL with optional headers, request body, and query parameters.
// It returns the HTTP response received or an error if any.
func (c *client) Post(url string, headers []Header, body io.Reader, queryParams ...QueryParam) (*http.Response, error) {
	// Append query parameters to the URL
	url = setQueryParams(queryParams, url)

	// Create a new POST request
	req, err := newRequest(c.ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	c.req = req

	// Set request headers
	setHeaders(c.req, headers)

	// Get the HTTP response
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Get sends a GET request to the specified URL with optional headers and query parameters.
// It returns the HTTP response received or an error if any.
func (c *client) Get(url string, headers []Header, queryParams ...QueryParam) (*http.Response, error) {
	// Append query parameters to the URL
	url = setQueryParams(queryParams, url)

	// Create a new GET request
	req, err := newRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	c.req = req

	// Set request headers
	setHeaders(c.req, headers)

	// Get the HTTP response
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Put sends a PUT request to the specified URL with optional headers, request body, and query parameters.
// It returns the HTTP response received or an error if any.
func (c *client) Put(url string, headers []Header, body io.Reader, queryParams ...QueryParam) (*http.Response, error) {
	// Append query parameters to the URL
	url = setQueryParams(queryParams, url)

	// Create a new PUT request
	req, err := newRequest(c.ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	c.req = req

	// Set request headers
	setHeaders(c.req, headers)

	// Get the HTTP response
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Patch sends a PATCH request to the specified URL with optional headers, request body, and query parameters.
// It returns the HTTP response received or an error if any.
func (c *client) Patch(url string, headers []Header, body io.Reader, queryParams ...QueryParam) (*http.Response, error) {
	// Append query parameters to the URL
	url = setQueryParams(queryParams, url)

	// Create a new PATCH request
	req, err := newRequest(c.ctx, http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}
	c.req = req

	// Set request headers
	setHeaders(c.req, headers)

	// Get the HTTP response
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Delete sends a DELETE request to the specified URL with optional headers, request body, and query parameters.
// It returns the HTTP response received or an error if any.
func (c *client) Delete(url string, headers []Header, body io.Reader, queryParams ...QueryParam) (*http.Response, error) {
	// Append query parameters to the URL
	url = setQueryParams(queryParams, url)

	// Create a new DELETE request
	req, err := newRequest(c.ctx, http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}
	c.req = req

	// Set request headers
	setHeaders(c.req, headers)

	// Get the HTTP response
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}
