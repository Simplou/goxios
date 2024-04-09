package goxios

import (
	"io"
	"net/http"
)

func (c *client) Post(url string, headers []Header, body io.Reader, queryParams ...QueryParam) (*http.Response, error) {
	url = setQueryParams(queryParams, url)
	req, err := newRequest(c.ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	c.req = req
	setHeaders(c.req, headers)
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *client) Get(url string, headers []Header, queryParams ...QueryParam) (*http.Response, error) {
	url = setQueryParams(queryParams, url)
	req, err := newRequest(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	c.req = req
	setHeaders(c.req, headers)
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *client) Put(url string, headers []Header, body io.Reader, queryParams ...QueryParam) (*http.Response, error) {
	url = setQueryParams(queryParams, url)
	req, err := newRequest(c.ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	c.req = req
	setHeaders(c.req, headers)
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *client) Patch(url string, headers []Header, body io.Reader, queryParams ...QueryParam) (*http.Response, error) {
	url = setQueryParams(queryParams, url)
	req, err := newRequest(c.ctx, http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}
	c.req = req
	setHeaders(c.req, headers)
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *client) Delete(url string, headers []Header, body io.Reader, queryParams ...QueryParam) (*http.Response, error) {
	url = setQueryParams(queryParams, url)
	req, err := newRequest(c.ctx, http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}
	c.req = req
	setHeaders(c.req, headers)
	res, err := c.Response()
	if err != nil {
		return nil, err
	}
	return res, nil
}
