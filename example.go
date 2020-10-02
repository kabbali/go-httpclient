package main

import (
	"fmt"
	"github.com/kabbali/go-httpclient.git/gohttp"
)

var (
	httpClient = getGithubClient()
)

func getGithubClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()

	//client.SetMaxIdleConnections(5)
	//client.SetConnectionTimeout(2 * time.Second)
	//client.SetResponseTimeout(5 * time.Millisecond)

	//client.DisableTimeouts(true)

	//commonHeaders := make(http.Header)
	//commonHeaders.Set("Authorization", "Bearer ABC-123")

	//client.SetHeaders(commonHeaders)

	return client
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	getUrls()
}

func getUrls() {

	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.StringBody())

	// Using our custom response
	//var user User
	//if err:= response.UnmarshalJson(&user); err != nil {
	//	panic(err)
	//}
	//fmt.Println(user.FirstName)
	//
	//// Using default http.Response
	//fmt.Println(response.StatusCode())
	//
	//bytes, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(bytes))
}

//func createUser(user User) {
//
//	response, err := httpClient.Post("https://api.github.com", nil, user)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(response.StatusCode())
//
//	bytes, err := ioutil.ReadAll(response.Body())
//	fmt.Println(string(bytes))
//}
