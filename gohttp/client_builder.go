package gohttp

import (
	"net/http"
	"time"
)

type (
	ClientBuilder interface {
		SetMaxIdleConnections(i int) ClientBuilder
		SetResponseTimeout(timeout time.Duration) ClientBuilder
		SetConnectionTimeout(timeout time.Duration) ClientBuilder
		SetHeaders(headers http.Header) ClientBuilder
		DisableTimeouts(disable bool) ClientBuilder
		SetHttpClient(client *http.Client) ClientBuilder
		SetUserAgent(userAgent string) ClientBuilder

		Build() Client
	}
	clientBuilder struct {
		headers http.Header

		// controls the maximum idle (keep-alive) connections to keep.
		maxIdleConnections int

		// amount of time to wait for a server's response headers after fully
		// writing the request.
		responseTimeout time.Duration

		// connectionTimeout is the maximum amount of time a dial will wait for
		// a connection.
		connectionTimeout time.Duration

		// allow to disable client timeouts
		disableTimeouts bool

		client    *http.Client
		userAgent string
	}
)

func NewBuilder() ClientBuilder {
	return &clientBuilder{}
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) SetMaxIdleConnections(i int) ClientBuilder {
	c.maxIdleConnections = i
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	c.userAgent = userAgent
	return c
}
