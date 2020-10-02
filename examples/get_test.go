package examples

import (
	"errors"
	"fmt"
	"github.com/kabbali/go-httpclient.git/gohttp"
	"net/http"
	"testing"
)

func TestGetEndpoints(t *testing.T) {
	gohttp.StartServer()

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		// Initialization
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url: "https://api.github.com",
			Error: errors.New("timeout getting github endpoints"),
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// Initialization
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url: "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			RequestBody: `{"current_user_url": 123}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "json unmarshalled error" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		// Initialization
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url: "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			RequestBody: `{"current_user_url": "https://api.github.com/user"}`,
		})

		// Execution
		endpoints, err := GetEndpoints()

		// Validation
		if err != nil {
			t.Error(fmt.Sprintf("no error was expected and we got '%s'", err.Error()))
		}

		if endpoints == nil {
			t.Error("expected endpoints and received nil ")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})


	// Initialization
	// Execution
	endpoints, err := GetEndpoints()

	// Validation
	fmt.Println(err)
	fmt.Println(endpoints)
}
