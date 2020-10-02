package gohttp

import (
	"errors"
	"fmt"
	"sync"
)

var mockupServer = mockServer{
	mocks: make(map[string]*Mock),
}

type mockServer struct {
	enable      bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}

func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enable = true
}

func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enable = false
}

// mocking feature
func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	return method + url + body
}

func (m *mockServer) getMock(method, url, body string) *Mock {
	if !m.enable {
		return nil
	}
	if mock := m.mocks[m.getMockKey(method, url, body)]; mock != nil {
		return mock
	}
	return &Mock{
		Error: errors.New(fmt.Sprintf("no mock matching %s from %s with given body", method, url)),
	}
}
