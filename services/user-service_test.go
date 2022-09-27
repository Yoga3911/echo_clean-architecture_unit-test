package services

import (
	"day-13-orm/models"
	"day-13-orm/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
var userSMock = NewUserService(userRMock)

func TestGetUsersService(t *testing.T) {
	usersMP := []*models.User{
		{
			Name:     "Mamat",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
		{
			Name:     "Mamat",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
	}

	usersM := []models.User{
		{
			Name:     "Mamat",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
		{
			Name:     "Mamat",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
	}

	userRMock.Mock.On("GetUsersRepository").Return(usersMP)
	users, err := userSMock.GetUsersService()
	

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, usersM[0].Name, users[0].Name)
	assert.Equal(t, usersM[0].Password, users[0].Password)
	assert.Equal(t, usersM[0].Email, users[0].Email)
}

func TestGetUserService(t *testing.T) {
	user := models.User{
		Name:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("GetUserRepository", "1").Return(user, nil)
	users, err := userSMock.GetUserService("1")

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, user.Name, users.Name)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
}
