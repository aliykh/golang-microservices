package services

import (
	"github.com/aliykh/golang-microservices/mvc/domain"
	"github.com/aliykh/golang-microservices/mvc/utils"
	"log"
)

var (
	UserService userService
)

type userService struct {

}

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	log.Println("Accessing User Service Layer")
	return domain.UserDaoInterface.GetUser(userId)
}
