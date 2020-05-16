package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)


func init() {
	UserDaoInterface = &userDao{}
}


func TestGetUserNil(t *testing.T) {

	user, err := UserDaoInterface.GetUser(0)

	assert.NotNil(t,  err)
	assert.Nil(t, user, "User with given id = 0 does not exist")

	assert.EqualValues(t, http.StatusNotFound, err.Code)
	assert.EqualValues(t, "Not found", err.Status)
	assert.EqualValues(t, "User with given id = 0 does not exist", err.Message)
}


func TestGetUser(t *testing.T) {

	user, err := UserDaoInterface.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "FirstName-123", user.FirstName)
	assert.EqualValues(t, "LastName-123", user.LastName)

}