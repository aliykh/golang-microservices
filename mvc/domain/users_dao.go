package domain

import (
	"fmt"
	"github.com/aliykh/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "FirstName-123",
			LastName:  "LastName-123",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Status:  "Not found",
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("User with given id = %v does not exist", userId),
	}
}
