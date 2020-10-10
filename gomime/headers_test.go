package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Error("invalid content type header")
	}

	if HeaderUserAgent != "User-Agent" {
		t.Error("invalid user agent header")
	}

	// possible values of HTTP Content-type header
	// https://www.geeksforgeeks.org/http-headers-content-type
	// Application type
	if ContentTypeJson != "application/json" {
		t.Error("invalid json content type value")
	}

	if ContentTypeXml != "application/xml" {
		t.Error("invalid xml content type value")
	}

	if ContentTypeOctetStream != "application/octet-stream" {
		t.Error("invalid octet-stream content type value")
	}

}
