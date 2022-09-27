package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type CreateUser struct {
	User  *User   `json:"user" form:"user"`
	Token string `json:"token" form:"token"`
}
