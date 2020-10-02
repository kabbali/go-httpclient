package gohttp

import "sync"

var mockupServer = mockServer{
	mocks: make(map[string]*Mock),
}

type mockServer struct {
	enable	bool
	serverMutex	sync.Mutex
	mocks map[string]*Mock
}

func StartServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Lock()

	mockupServer.enable = true
}

func StopServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Lock()

	mockupServer.enable = false
}

// mocking feature
func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Lock()

	key := mock.Method + mock.Url + mock.RequestBody
	mockupServer.mocks[key] = &mock
}
