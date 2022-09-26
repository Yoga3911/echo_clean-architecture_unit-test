package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"day-13-orm/configs"
	"day-13-orm/models"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := configs.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	userId := c.Param("id")
	var user models.User

	if err := configs.DB.Where("ID = ?", userId).Take(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := configs.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	userId := c.Param("id")
	var user models.User

	if err := configs.DB.Where("ID = ?", userId).Take(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	if err := configs.DB.Delete(&models.User{}, userId).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var user models.User
	userId := c.Param("id")

	if err := configs.DB.Where("ID = ?", userId).Take(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "user not found",
		})
	}

	var userBody models.User
	err := c.Bind(&userBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err,
		})
	}

	err = configs.DB.Where("ID = ?", userId).Updates(models.User{Name: userBody.Name, Email: userBody.Email, Password: userBody.Password}).Error
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err,
		})
	}

	user.Name = userBody.Name
	user.Email = userBody.Email
	user.Password = userBody.Password

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}
