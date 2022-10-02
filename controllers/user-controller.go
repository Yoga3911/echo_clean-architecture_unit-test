package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "day-13-orm/helpers"
	m "day-13-orm/middlewares"
	"day-13-orm/models"
	"day-13-orm/services"
)

type UserController interface {
	GetUsersController(c echo.Context) error
	GetUserController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type userController struct {
	userS services.UserService
	jwt m.JWTS
}

func NewUserController(userS services.UserService, jwtS m.JWTS) UserController {
	return &userController{
		userS: userS,
		jwt: jwtS,
	}
}

// get all users
func (u *userController) GetUsersController(c echo.Context) error {
	users, err := u.userS.GetUsersService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    users,
		Message: "Get all users success",
		Status:  true,
	})
}

func (u *userController) GetUserController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var user *models.User

	user, err = u.userS.GetUserService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Get user success",
		Status:  true,
	})
}

func (u *userController) CreateController(c echo.Context) error {
	var user models.CreateUser

	err := c.Bind(&user.User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.User, err = u.userS.CreateService(*user.User)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	token, err := u.jwt.CreateJWTToken(user.User.ID, user.User.Name)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user.Token = token
	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Create user success",
		Status:  true,
	})
}

func (u *userController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var user *models.User

	err = c.Bind(&user)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	user, err = u.userS.UpdateService(id, *user)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    user,
		Message: "Update user success",
		Status:  true,
	})
}

func (u *userController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = u.userS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete user success",
		Status:  true,
	})
}
