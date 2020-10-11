package core

import (
	"encoding/json"
	"net/http"
)

type (
	Response struct {
		Status     string
		StatusCode int
		Headers    http.Header
		Body       []byte
	}
)

//
//func (r *Response) Status() string {
//	return r.status
//}
//
//func (r *Response) StatusCode() int {
//	return r.statusCode
//}
//
//func (r *Response) Headers() http.Header {
//	return r.headers
//}
//
func (r *Response) Bytes() []byte {
	return r.Body
}

func (r *Response) StringBody() string {
	return string(r.Body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}
