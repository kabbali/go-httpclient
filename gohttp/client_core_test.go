package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.Headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("We expect 3 headers")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id received")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent")
	}
}

func TestGetRequestBody(t *testing.T) {
	// Initialization
	client := httpClient{}

	t.Run("NilBodyNilResponse", func(t *testing.T) {
		// Execution
		body, err := client.getRequestBody("", nil)

		// Validation
		if err != nil {
			t.Error("no error expected when passing a nil body")
		}

		if body != nil {
			t.Error("no body expected when passing a nil body")
		}
	})

	t.Run("JsonBody", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)
		
		// Validation
		if err != nil {
			t.Error("no error expected when marshaling string slice as json")
		}

		if string(body) != 	`["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("XmlBody", func(t *testing.T) {
		// Execution
		requestBody := []string{"three", "four"}
		body, err := client.getRequestBody("application/xml", requestBody)

		// Validation
		if err != nil {
			t.Error("no error expected when marshaling string slice as xml")
		}

		if string(body) != `<string>three</string><string>four</string>` {
			t.Error("invalid xml body obtained")
		}
	})

	t.Run("DefaultJsonBody", func(t *testing.T) {
		// Execution
		requestBody := []string{"default", "xml", "body"}
		body, err := client.getRequestBody("", requestBody)

		if err != nil {
			t.Error("no error expected when marshaling default case as json")
		}

		if string(body) != `["default","xml","body"]` {
			t.Error("invalid json body obtained")
		}
	})
}
