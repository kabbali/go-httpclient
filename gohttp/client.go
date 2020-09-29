package gohttp

import (
	"net/http"
	"time"
)

type httpClient struct {
	client  *http.Client
	Headers http.Header

	// controls the maximum idle (keep-alive) connections to keep.
	maxIdleConnections int

	// amount of time to wait for a server's response headers after fully
	// writing the request.
	responseTimeout time.Duration

	// connectionTimeout is the maximum amount of time a dial will wait for
	// a connection.
	connectionTimeout time.Duration

	// allow disable client timeouts
	disableTimeouts bool
}

func New() HttpClient {
	return &httpClient{}
}

type HttpClient interface {
	SetMaxIdleConnections(i int)
	SetResponseTimeout(timeout time.Duration)
	SetConnectionTimeout(timeout time.Duration)
	SetHeaders(headers http.Header)
	DisableTimeouts(disable bool)

	Get(url string, headers http.Header) (*http.Response, error)
	Post(url string, headers http.Header, body interface{}) (*http.Response, error)
	Put(url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete(url string, headers http.Header) (*http.Response, error)
}

func (c *httpClient) SetMaxIdleConnections(i int) {
	c.maxIdleConnections = i
}

func (c *httpClient) SetResponseTimeout(timeout time.Duration) {
	c.responseTimeout = timeout
}

func (c *httpClient) SetConnectionTimeout(timeout time.Duration) {
	c.connectionTimeout = timeout
}

func (c *httpClient) DisableTimeouts(disable bool) {
	c.disableTimeouts = disable
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.Headers = headers
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}
