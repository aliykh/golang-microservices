package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/aliykh/golang-microservices/mvc/services"
	"github.com/aliykh/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUsers(resp http.ResponseWriter, req *http.Request){

	queryParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(queryParam, 10, 64)

	if err != nil {
	// just return bad request to the user
		resp.WriteHeader(http.StatusBadRequest)
		jsonValue, _ := json.Marshal(utils.ApplicationError{
			Status:  "Bad request",
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("user id must be a number = %v", queryParam),
		})
		_, _ = resp.Write(jsonValue)
		return
	}


	user, applicationError := services.GetUser(userId)

	if applicationError != nil {
		// handle user get function error -> return bad request
		jsonValue, _ := json.Marshal(applicationError)
		resp.WriteHeader(applicationError.Code)
		_, _ = resp.Write(jsonValue)
		return
	}


	//	Return to the client
	jsonValue, err := json.Marshal(user)
	_, _ = resp.Write(jsonValue)

}