package util

import "os"

const (
	apiGithubAccessToken  = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetAccessToken() string {
	return githubAccessToken
}
