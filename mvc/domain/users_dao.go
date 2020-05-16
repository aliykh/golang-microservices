package domain

import (
	"fmt"
	"github.com/aliykh/golang-microservices/mvc/utils"
	"log"
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

	UserDaoInterface userDaoInterface
)

type userDaoInterface interface {
	GetUser(userId int64) (*User, *utils.ApplicationError)
}

func init() {
	UserDaoInterface = &userDao{}
}

type userDao struct {}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("Accessing User DAO Layer")
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Status:  "Not found",
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("User with given id = %v does not exist", userId),
	}
}
