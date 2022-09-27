package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	h "day-13-orm/helpers"
	"day-13-orm/models"
	"day-13-orm/services"
)

type BookController interface {
	GetBooksController(c echo.Context) error
	GetBookController(c echo.Context) error
	CreateController(c echo.Context) error
	UpdateController(c echo.Context) error
	DeleteController(c echo.Context) error
}

type bookController struct {
	BookS services.BookService
}

func NewBookController(BookS services.BookService) BookController {
	return &bookController{
		BookS: BookS,
	}
}

// get all Books
func (b *bookController) GetBooksController(c echo.Context) error {
	Books, err := b.BookS.GetBooksService()
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Books,
		Message: "Get all Books success",
		Status:  true,
	})
}

func (b *bookController) GetBookController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Book *models.Book

	Book, err = b.BookS.GetBookService(id)
	if err != nil {
		return h.Response(c, http.StatusNotFound, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Book,
		Message: "Get Book success",
		Status:  true,
	})
}

func (b *bookController) CreateController(c echo.Context) error {
	var Book *models.Book

	err := c.Bind(&Book)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Book, err = b.BookS.CreateService(*Book)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Book,
		Message: "Create Book success",
		Status:  true,
	})
}

func (b *bookController) UpdateController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	var Book *models.Book

	err = c.Bind(&Book)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	Book, err = b.BookS.UpdateService(id, *Book)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    Book,
		Message: "Update Book success",
		Status:  true,
	})
}

func (b *bookController) DeleteController(c echo.Context) error {
	id := c.Param("id")

	err := h.IsNumber(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	err = b.BookS.DeleteService(id)
	if err != nil {
		return h.Response(c, http.StatusBadRequest, h.ResponseModel{
			Data:    nil,
			Message: err.Error(),
			Status:  false,
		})
	}

	return h.Response(c, http.StatusOK, h.ResponseModel{
		Data:    nil,
		Message: "Delete Book success",
		Status:  true,
	})
}
