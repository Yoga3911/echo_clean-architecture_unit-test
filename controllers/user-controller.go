package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "day-13-orm/helpers"
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
}

func NewUserController(userS services.UserService) UserController {
	return &userController{
		userS: userS,
	}
}

// get all users
func (u *userController) GetUsersController(c echo.Context) error {
	users, err := u.userS.GetUsersService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    users,
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
	userId := c.Param("id")
	user, err := u.userS.GetUserService(userId)
	if err != nil {
		return h.Response(c, http.StatusOK, h.ResponseModel{
			Data:    user,
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

	return nil
}

func (u *userController) UpdateController(c echo.Context) error {

	return nil
}

func (u *userController) DeleteController(c echo.Context) error {
	return nil
}

// // get user by id
// func GetUserController(c echo.Context) error {
// 	// your solution here
// 	userId := c.Param("id")
// 	var user models.User

// 	if err := configs.DB.Where("ID = ?", userId).Take(&user).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]interface{}{
// 			"message": "user not found",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get all users",
// 		"users":   user,
// 	})
// }

// // create new user
// func CreateUserController(c echo.Context) error {
// 	user := models.User{}
// 	c.Bind(&user)

// 	if err := configs.DB.Save(&user).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	token, err := middlewares.CreateJWTToken(user.ID, user.Name)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": err,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success create new user",
// 		"user":    user,
// 		"token":   token,
// 	})
// }

// // delete user by id
// func DeleteUserController(c echo.Context) error {
// 	// your solution here
// 	userId := c.Param("id")
// 	var user models.User

// 	if err := configs.DB.Where("ID = ?", userId).Take(&user).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]interface{}{
// 			"message": "user not found",
// 		})
// 	}

// 	if err := configs.DB.Delete(&models.User{}, userId).Error; err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": err,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success delete user",
// 	})
// }

// // update user by id
// func UpdateUserController(c echo.Context) error {
// 	// your solution here
// 	var user models.User
// 	userId := c.Param("id")

// 	if err := configs.DB.Where("ID = ?", userId).Take(&user).Error; err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]interface{}{
// 			"message": "user not found",
// 		})
// 	}

// 	var userBody models.User
// 	err := c.Bind(&userBody)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"message": err,
// 		})
// 	}

// 	err = configs.DB.Where("ID = ?", userId).Updates(models.User{Name: userBody.Name, Email: userBody.Email, Password: userBody.Password}).Error
// 	if err != nil {
// 		return c.JSON(http.StatusNotFound, map[string]interface{}{
// 			"message": err,
// 		})
// 	}

// 	user.Name = userBody.Name
// 	user.Email = userBody.Email
// 	user.Password = userBody.Password

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success update user",
// 		"user":    user,
// 	})
// }
