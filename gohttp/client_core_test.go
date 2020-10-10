package gohttp

import (
	"testing"
)

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

		if string(body) != `["one","two"]` {
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
