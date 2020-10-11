package examples

import (
	"errors"
	"github.com/kabbali/go-httpclient/gohttp_mock"
	"net/http"
	"strings"
	"testing"
)

func TestCreateRepo(t *testing.T) {
	t.Run("TimeoutErrorFromGithub", func(t *testing.T) {
		gohttp_mock.DeleteMocks()
		gohttp_mock.AddMock(gohttp_mock.Mock{
			Method:      http.MethodPost,
			Url:         "https://api.github.com/user/repos",
			RequestBody: `{"name":"test-repo","description":"repository for testing","private":true}`,
			Error:       errors.New("timeout error from github"),
		})

		repository := Repository{
			Name:        "test-repo",
			Description: "repository for testing",
			Private:     true,
		}

		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("no repository expected when we get a timeout from github")
		}

		if err == nil {
			t.Error("error expected when we get a timeout from github")
		}

		if err.Error() != "timeout error from github" {
			t.Error("error message should be: timeout error from github")
		}
	})

	t.Run("NoErrorProcessingGithubErrorResponse", func(t *testing.T) {
		gohttp_mock.DeleteMocks()
		gohttp_mock.AddMock(gohttp_mock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"repository for testing","private":true}`,
			ResponseStatusCode: http.StatusUnauthorized,
			ResponseBody:       `{"message":"Requires authentication"}`,
		})

		repository := Repository{
			Name:        "test-repo",
			Description: "repository for testing",
			Private:     true,
		}

		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("no repository expected when we get a timeout from github")
		}

		if err == nil {
			t.Error("error expected when we get a timeout from github")
		}

		if err.Error() != "Requires authentication" {
			t.Error("error message should be: Requires authentication")
		}
	})

	t.Run("ErrorProcessingGithubErrorResponse", func(t *testing.T) {
		gohttp_mock.DeleteMocks()
		gohttp_mock.AddMock(gohttp_mock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"repository for testing","private":true}`,
			ResponseStatusCode: http.StatusUnauthorized,
			ResponseBody:       `{"message":123}`,
		})

		repository := Repository{
			Name:        "test-repo",
			Description: "repository for testing",
			Private:     true,
		}

		repo, err := CreateRepo(repository)

		if repo != nil {
			t.Error("no repository expected when we get a timeout from github")
		}

		if err == nil {
			t.Error("error expected when we get a timeout from github")
		}

		if !strings.Contains(err.Error(), "error processing github error response when creating a new repo") {
			t.Error("invalid error message received")
		}
	})

	t.Run("NoErrorFromGithub", func(t *testing.T) {
		gohttp_mock.DeleteMocks()
		gohttp_mock.AddMock(gohttp_mock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"repository for testing","private":true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":129,"name":"test-repo"}`,
		})

		repository := Repository{
			Name:        "test-repo",
			Description: "repository for testing",
			Private:     true,
		}

		repo, err := CreateRepo(repository)

		if err != nil {
			t.Error("no error expected when we get a response from github")
		}

		if repo == nil {
			t.Error("repository expected when we get a response from github")
		}

		if repo.Name != repository.Name {
			t.Error("a valid repo was expected at this point")
		}
	})

	t.Run("ErrorUnmarshalResponseBody", func(t *testing.T) {
		gohttp_mock.DeleteMocks()
		gohttp_mock.AddMock(gohttp_mock.Mock{
			Method:             http.MethodPost,
			Url:                "https://api.github.com/user/repos",
			RequestBody:        `{"name":"test-repo","description":"repository for testing","private":true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody:       `{"id":129,"name":123}`,
		})

		repository := Repository{
			Name:        "test-repo",
			Description: "repository for testing",
			Private:     true,
		}

		// Execution
		repo, err := CreateRepo(repository)

		// Validation
		if repo != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field Repository") {
			t.Error("invalid error message received")
		}
	})
}
