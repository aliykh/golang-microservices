package controllers

import (
	"fmt"
	"github.com/aliykh/golang-microservices/mvc/services"
	"github.com/aliykh/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context){

	queryParam, _ := c.Params.Get("user_id")
	userId, err := strconv.ParseInt(queryParam, 10, 64)

	if err != nil {
	 appErr := utils.ApplicationError{
		 Status:  "Bad request",
		 Code:    http.StatusBadRequest,
		 Message: fmt.Sprintf("user id must be a number = %v", queryParam),
	 }
		utils.RespondAppError(c, &appErr)
		return
	}


	user, appErr := services.UserService.GetUser(userId)

	if appErr != nil {
		utils.RespondAppError(c, appErr)
		return
	}


	//	Return to the client
	utils.Respond(c, http.StatusOK, user)
}