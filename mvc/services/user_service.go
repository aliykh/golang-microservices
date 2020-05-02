package services

import (
	"github.com/aliykh/golang-microservices/mvc/domain"
	"github.com/aliykh/golang-microservices/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
