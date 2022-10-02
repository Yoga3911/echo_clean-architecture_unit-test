package controllers

import (
	"day-13-orm/middlewares"
	"day-13-orm/models"
	"day-13-orm/repositories"
	"day-13-orm/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = services.NewUserService(userRMock)
	jwtMock   = &middlewares.IjwtSMock{Mock: mock.Mock{}}
	userCTest = NewUserController(userSMock, jwtMock)
)

func TestGetUsersController_Success(t *testing.T) {
	users := []*models.User{
		{
			Name:     "Mamat2342",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
		{
			Name:     "Mamat",
			Email:    "qwe@gmail.com",
			Password: "123456",
		},
	}

	userRMock.Mock.On("GetUsersRepository").Return(users, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := userCTest.GetUsersController(c)
	assert.Nil(t, err)
}

func TestGetUsersController_Failure(t *testing.T) {
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = services.NewUserService(userRMock)
	userCTest = NewUserController(userSMock, jwtMock)
	userRMock.Mock.On("GetUsersRepository").Return(nil, errors.New("get all users failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := userCTest.GetUsersController(c)
	assert.Nil(t, err)
}

func TestGetUserController_Success(t *testing.T) {
	user := models.User{
		Model: gorm.Model{
			ID: 2,
		},
		Name:     "Mamat2342",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("GetUserRepository", "2").Return(user, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := userCTest.GetUserController(c)
	assert.Nil(t, err)
}

func TestGetUserController_Failure1(t *testing.T) {
	userRMock.Mock.On("GetUserRepository", "qwe").Return(nil, errors.New("get user failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/users/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := userCTest.GetUserController(c)
	assert.Nil(t, err)
}

func TestGetUserController_Failure2(t *testing.T) {
	userRMock.Mock.On("GetUserRepository", "3").Return(nil, fmt.Errorf("user not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/users/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := userCTest.GetUserController(c)
	assert.Nil(t, err)
}

func TestCreateUserController_Success(t *testing.T) {
	user := models.User{
		Name:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("CreateRepository", user).Return(user, nil)
	jwtMock.Mock.On("CreateJWTToken").Return("token123", nil)

	rec := httptest.NewRecorder()

	userByte, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(userByte))

	req := httptest.NewRequest(http.MethodPost, "/users", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = userCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateUserController_Failure1(t *testing.T) {
	user := models.User{}

	userRMock.Mock.On("CreateRepository", user).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	userByte, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(userByte))

	req := httptest.NewRequest(http.MethodPost, "/users", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = userCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateUserController_Failure2(t *testing.T) {
	user := models.Book{}

	userRMock.Mock.On("CreateRepository", user).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/users", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := userCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateUserController_Failure3(t *testing.T) {
	jwtMock = &middlewares.IjwtSMock{Mock: mock.Mock{}}
	userCTest = NewUserController(userSMock, jwtMock)
	user := models.User{
		Name:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("CreateRepository", user).Return(user, nil)
	jwtMock.Mock.On("CreateJWTToken").Return(nil, errors.New("error bang"))

	rec := httptest.NewRecorder()

	userByte, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(userByte))

	req := httptest.NewRequest(http.MethodPost, "/users", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = userCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdateUserController_Success(t *testing.T) {
	user := models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("UpdateRepository", "1", user).Return(user, nil)

	rec := httptest.NewRecorder()

	userByte, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(userByte))

	req := httptest.NewRequest(http.MethodPut, "/users/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = userCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateUserController_Failure1(t *testing.T) {
	user := models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("UpdateRepository", "1", user).Return(user, nil)

	rec := httptest.NewRecorder()

	userByte, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(userByte))

	req := httptest.NewRequest(http.MethodPut, "/users/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = userCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateUserController_Failure2(t *testing.T) {
	user := models.User{}

	userRMock.Mock.On("UpdateRepository", "1", user).Return(user, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/users/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = userCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateUserController_Failure3(t *testing.T) {
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = services.NewUserService(userRMock)
	userCTest = NewUserController(userSMock, jwtMock)
	user := models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "Mamat",
		Email:    "qwe@gmail.com",
		Password: "123456",
	}

	userRMock.Mock.On("UpdateRepository", "1", user).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	userByte, err := json.Marshal(user)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(userByte))

	req := httptest.NewRequest(http.MethodPut, "/users/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = userCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeleteUserController_Success(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/users/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := userCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteUserController_Failure1(t *testing.T) {
	userRMock = &repositories.IuserRepositoryMock{Mock: mock.Mock{}}
	userSMock = services.NewUserService(userRMock)
	userCTest = NewUserController(userSMock, jwtMock)
	userRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("user not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/users/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := userCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteUserController_Failure2(t *testing.T) {
	userRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("user not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := userCTest.DeleteController(c)

	assert.Nil(t, err)
}
