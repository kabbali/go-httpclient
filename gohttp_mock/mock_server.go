package gohttp_mock

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var MockupServer = mockServer{
	mocks: make(map[string]*Mock),
}

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}

func StartMockServer() {
	MockupServer.serverMutex.Lock()
	defer MockupServer.serverMutex.Unlock()

	MockupServer.enabled = true
}

func StopMockServer() {
	MockupServer.serverMutex.Lock()
	defer MockupServer.serverMutex.Unlock()

	MockupServer.enabled = false
}

func DeleteMocks() {
	MockupServer.serverMutex.Lock()
	defer MockupServer.serverMutex.Unlock()

	MockupServer.mocks = make(map[string]*Mock)
}
func AddMock(mock Mock) {
	MockupServer.serverMutex.Lock()
	defer MockupServer.serverMutex.Unlock()

	key := MockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	MockupServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	//fmt.Println(fmt.Sprintf("KEY: `%s`", key))
	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	return body
}

func GetMock(method, url, body string) *Mock {
	if !MockupServer.enabled {
		return nil
	}
	if mock := MockupServer.mocks[MockupServer.getMockKey(method, url, body)]; mock != nil {
		return mock
	}
	return &Mock{
		Error: errors.New(fmt.Sprintf("no mock matching %s from `%s` with given body", method, url)),
	}
}
