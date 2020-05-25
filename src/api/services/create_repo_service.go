package services

import (
	"github.com/aliykh/golang-microservices/src/api/domain/github"
	"github.com/aliykh/golang-microservices/src/api/domain/repositories"
	"github.com/aliykh/golang-microservices/src/api/errors"
	"github.com/aliykh/golang-microservices/src/api/provider/github_provider"
	"github.com/aliykh/golang-microservices/src/api/util"
)

var (
	RepoServiceInterface repoServiceInterface
)

type repoService struct {

}


type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

func init() {
	RepoServiceInterface = &repoService{}
}

func (r *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {


	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Name,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(util.GetAccessToken(), request)
	if err != nil {
		return nil, errors.NewCustomApiError(err.StatusCode, err.Message)
	}

	return &repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}, nil
}


