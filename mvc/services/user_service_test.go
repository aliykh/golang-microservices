package services

import (
	"github.com/aliykh/golang-microservices/mvc/domain"
	"github.com/aliykh/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)


//TODO In this test file, we have implemented mocking feature to test DAO layer inside user_service.go file [return domain.UserDaoInterface.GetUser(userId)]

var (
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDaoInterface = &usersServiceMock{}
}

type usersServiceMock struct {

}

func (u *usersServiceMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError)  {
	log.Println("User Mock Layer GetUser()")
	return getUserFunction(userId)
}




func TestUserService_GetUserNil(t *testing.T) {

	getUserFunction = func(userId int64) (user *domain.User, applicationError *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Status:  "Not found",
			Code:    http.StatusNotFound,
			Message: "User with given id = 0 does not exist",
		}
	}


	user, err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Code)
	assert.EqualValues(t, "Not found", err.Status)
	assert.EqualValues(t, "User with given id = 0 does not exist", err.Message)


}


func TestUserService_GetUserNotNil(t *testing.T) {

	getUserFunction = func(userId int64) (user *domain.User, applicationError *utils.ApplicationError) {
		return &domain.User{
			Id:        123,
			FirstName: "FirstName-123",
			LastName:  "LastName-123",
		}, nil
	}

	user, err := UserService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)


}
