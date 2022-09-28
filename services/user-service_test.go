package services

import (
	"day-13-orm/models"
	"day-13-orm/repositories"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = NewUserService(userRMock)
)

func TestGetUsersService_Success(t *testing.T) {
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

	userRMock.Mock.On("GetUsersRepository").Return(usersMP, nil)
	users, err := userSMock.GetUsersService()

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, usersM[0].Name, users[0].Name)
	assert.Equal(t, usersM[0].Password, users[0].Password)
	assert.Equal(t, usersM[0].Email, users[0].Email)
}

func TestGetUsersService_Failure(t *testing.T) {
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = NewUserService(userRMock)
	userRMock.Mock.On("GetUsersRepository").Return(nil, errors.New("get all users failed"))
	users, err := userSMock.GetUsersService()

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestGetUserService_Success(t *testing.T) {
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

func TestGetUserService_Failure(t *testing.T) {
	userRMock.Mock.On("GetUserRepository", "3").Return(nil, fmt.Errorf("user not found"))
	user, err := userSMock.GetUserService("3")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestCreateUserService_Success(t *testing.T) {
	user := models.User{
		Name:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("CreateRepository", user).Return(user, nil)
	users, err := userSMock.CreateService(user)

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, user.Name, users.Name)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
}

func TestCreateUserService_Failure(t *testing.T) {
	user := models.User{
		Name:     "Mamat123",
		Email:    "qwe3123@gmail.com",
		Password: "123456321",
	}

	userRMock.Mock.On("CreateRepository", user).Return(nil, fmt.Errorf("create user failed"))
	users, err := userSMock.CreateService(user)

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestUpdateUserService_Success(t *testing.T) {
	user := models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("UpdateRepository", "1", user).Return(user, nil)
	users, err := userSMock.UpdateService("1", user)

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, uint(1), users.ID)
	assert.Equal(t, user.Name, users.Name)
	assert.Equal(t, user.Password, users.Password)
	assert.Equal(t, user.Email, users.Email)
}

func TestUpdateUserService_Failure(t *testing.T) {
	user := models.User{
		Model: gorm.Model{
			ID: 2,
		},
		Name:     "Mamat123",
		Email:    "q321e@gmail.com",
		Password: "123456321",
	}

	userRMock.Mock.On("UpdateRepository", "2", user).Return(nil, fmt.Errorf("user not found"))
	users, err := userSMock.UpdateService("2", user)

	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestDeleteUserService_Success(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := userSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteUserService_Failure(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("user not found"))
	err := userSMock.DeleteService("2")

	assert.NotNil(t, err)
}
