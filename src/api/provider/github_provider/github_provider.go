package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/aliykh/golang-microservices/src/api/clients/restclient"
	"github.com/aliykh/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	authorization = "Authorization"
	accessToken   = "token %s"
	repoCreateUrl = "https://api.github.com/user/repos"
)

func getAuthorizationToken(token string) string {
	return fmt.Sprintf(accessToken, token)
}

func CreateRepo(accessToken string, requestBody github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {

	header := http.Header{}
	header.Set(authorization, getAuthorizationToken(accessToken))

	response, err := restclient.Post(repoCreateUrl, header, requestBody)
	if err != nil {
		log.Println("Error while calling POST method (request) to github url")
		return nil, &github.GithubErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		log.Println("Error while reading response.Body into bytes")
		return nil, &github.GithubErrorResponse{
			Message:    fmt.Sprintf("Invalid response body: %v", err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}

	if response.StatusCode > 299 {
		var errorResponse github.GithubErrorResponse
		fmt.Println(string(bytes))
		err := json.Unmarshal(bytes, &errorResponse)
		if err != nil {
			log.Println("Error when trying to unmarshal error repo response")
			return nil, &github.GithubErrorResponse{
				Message:    fmt.Sprintf("Invalid json response body: %v", err.Error()),
				StatusCode: http.StatusInternalServerError,
			}
		}

		errorResponse.StatusCode = response.StatusCode
		return nil, &errorResponse
	}

	var result github.CreateRepoResponse

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		log.Println("Error when trying to unmarshal create repo response")
		return nil, &github.GithubErrorResponse{
			Message:    "Invalid json response body",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &result, nil
}
