package gohttp_mock

import (
	"fmt"
	"github.com/kabbali/go-httpclient/core"
	"net/http"
)

type Mock struct {
	Method      string
	Url         string
	RequestBody string

	// Response
	Error              error
	ResponseBody       string
	ResponseStatusCode int
}

func (m *Mock) GetResponse() (*core.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := core.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
	}
	return &response, nil
}
