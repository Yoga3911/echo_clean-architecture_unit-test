package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title" form:"title"`
	Author      string `json:"author" form:"author"`
	Description string `json:"description" form:"description"`
}
