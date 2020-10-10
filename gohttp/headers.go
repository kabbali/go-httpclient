package gohttp

import (
	"github.com/kabbali/go-httpclient/gomime"
	"net/http"
)

func getHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return http.Header{}
}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	headers := make(http.Header)

	// Add custom Headers from current request (defined in do method)
	for header, value := range requestHeaders {
		if len(value) > 0 {
			headers.Set(header, value[0])
		}
	}

	// Add common Headers from HTTP client instance (defined in httpClient struct)
	for header, value := range c.builder.headers {
		if len(value) > 0 {
			headers.Set(header, value[0])
		}
	}

	// Set User-Agent if none is defined
	if c.builder.userAgent != "" {
		if headers.Get(gomime.HeaderUserAgent) != "" {
			return headers
		}
		headers.Set(gomime.HeaderUserAgent, c.builder.userAgent)
	}
	return headers
}
