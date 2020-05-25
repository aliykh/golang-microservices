package controllers

import (
	"github.com/aliykh/golang-microservices/src/api/domain/repositories"
	"github.com/aliykh/golang-microservices/src/api/errors"
	"github.com/aliykh/golang-microservices/src/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRepo(c *gin.Context){

	var input repositories.CreateRepoRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusNotAcceptable, errors.NotAcceptableException(err.Error()))
		return
	}

	response, err := services.RepoServiceInterface.CreateRepo(input)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}


	c.JSON(http.StatusCreated, response)
}
