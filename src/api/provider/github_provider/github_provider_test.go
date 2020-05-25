package github_provider

import (
	"errors"
	"github.com/aliykh/golang-microservices/src/api/clients/restclient"
	"github.com/aliykh/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)


func TestCreateRepoConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", authorization)
	assert.EqualValues(t, "token %s", accessToken)
	assert.EqualValues(t, "https://api.github.com/user/repos", repoCreateUrl)
}

func TestGetAuthorizationToken(t *testing.T) {

	token := getAuthorizationToken("abc123")
	assert.EqualValues(t, token, "token abc123")
}

func TestMain(m *testing.M) {
	restclient.StartMockUp()
	os.Exit(m.Run())
}

func TestCreateRepoError(t *testing.T) {

	restclient.FlushMocks()

	restclient.AddMocks(restclient.Mock{
		Err:        errors.New("Invalid rest client response"),
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
	})

	response, err := CreateRepo("token", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid rest client response", err.Message)
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {

	restclient.FlushMocks()

	invalidCloser, _ := os.Open("-asf3")

	restclient.AddMocks(restclient.Mock{
		Err:        nil,
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       invalidCloser,
		},
	})

	response, err := CreateRepo("token", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid response body: invalid argument", err.Message)
}

func TestCreateRepoInvalidJsonResponseBody(t *testing.T) {

	restclient.FlushMocks()

	restclient.AddMocks(restclient.Mock{
		Err:        nil,
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1234}`)),
		},
	})

	response, err := CreateRepo("token", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid json response body: json: cannot unmarshal number into Go struct field GithubErrorResponse.message of type string", err.Message)
}

func TestCreateRepoUnauthorizedResponse(t *testing.T) {

	restclient.FlushMocks()

	restclient.AddMocks(restclient.Mock{
		Err:        nil,
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Bad credentials","documentation_url": "https://developer.github.com/v3"}`)),
		},
	})

	response, err := CreateRepo("token", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Bad credentials", err.Message)
}

func TestCreateRepoSuccessInvalidBody(t *testing.T) {

	restclient.FlushMocks()

	restclient.AddMocks(restclient.Mock{
		Err:        nil,
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})

	response, err := CreateRepo("token", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "Invalid json response body", err.Message)
}

func TestCreateRepoSuccessWithoutErrors(t *testing.T) {

	restclient.FlushMocks()

	restclient.AddMocks(restclient.Mock{
		Err:        nil,
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 264608095,"name": "Hello-World","full_name": "aliykh/Hello-World"}`)),
		},
	})

	response, err := CreateRepo("token", github.CreateRepoRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
}
