package gohttp_mock

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClientMock struct {
}

func (c *httpClientMock) Do(request *http.Request) (*http.Response, error) {

	requestBody, err := request.GetBody()
	if err != nil {
		return nil, err
	}
	defer requestBody.Close()

	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}

	// Validate if a mock exists for request
	var response http.Response
	mock := MockupServer.mocks[MockupServer.getMockKey(request.Method, request.URL.String(), string(body))]
	if mock != nil {
		if mock.Error != nil {
			return nil, mock.Error
		}
		// We have a mock at this point so we return a Mock
		response.StatusCode = mock.ResponseStatusCode
		response.Body = ioutil.NopCloser(strings.NewReader(mock.ResponseBody))
		response.ContentLength = int64(len(mock.ResponseBody))
		response.Request = request
		return &response, nil
	}

	// We dont have a mock at this point so we return error
	return nil, errors.New(fmt.Sprintf("no mock matching %s from `%s` with given body", request.Method, request.URL.String()))
}
