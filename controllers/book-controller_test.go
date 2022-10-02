package controllers

import (
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
	bookRMock = &repositories.IbookRepositoryMock{Mock: mock.Mock{}}
	bookSMock = services.NewBookService(bookRMock)
	bookCTest = NewBookController(bookSMock)
)

func TestGetBooksController_Success(t *testing.T) {
	books := []*models.Book{
		{
			Title:       "Batman",
			Author:      "Boy",
			Description: "Buku action",
		},
		{
			Title:       "Batman",
			Author:      "Boy",
			Description: "Buku action",
		},
	}

	bookRMock.Mock.On("GetBooksRepository").Return(books, nil)

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := bookCTest.GetBooksController(c)
	assert.Nil(t, err)
}

func TestGetBooksController_Failure(t *testing.T) {
	bookRMock = &repositories.IbookRepositoryMock{Mock: mock.Mock{}}
	bookSMock = services.NewBookService(bookRMock)
	bookCTest = NewBookController(bookSMock)
	bookRMock.Mock.On("GetBooksRepository").Return(nil, errors.New("get all Books failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	e := echo.New()

	c := e.NewContext(req, rec)

	err := bookCTest.GetBooksController(c)
	assert.Nil(t, err)
}

func TestGetBookController_Success(t *testing.T) {
	book := models.Book{
		Model: gorm.Model{
			ID: 2,
		},
		Title:       "Batman",
		Author:      "Sugeng",
		Description: "Mantap",
	}

	bookRMock.Mock.On("GetBookRepository", "2").Return(book, nil)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/books/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := bookCTest.GetBookController(c)
	assert.Nil(t, err)
}

func TestGetBookController_Failure1(t *testing.T) {
	bookRMock.Mock.On("GetBookRepository", "qwe").Return(nil, errors.New("get book failed"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/books/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := bookCTest.GetBookController(c)
	assert.Nil(t, err)
}

func TestGetBookController_Failure2(t *testing.T) {
	bookRMock.Mock.On("GetBookRepository", "3").Return(nil, fmt.Errorf("book not found"))

	rec := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/books/3", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("3")

	err := bookCTest.GetBookController(c)
	assert.Nil(t, err)
}

func TestCreateBookController_Success(t *testing.T) {
	book := models.Book{
		Title:       "Batman",
		Author:      "Sugeng",
		Description: "Mantap",
	}

	bookRMock.Mock.On("CreateRepository", book).Return(book, nil)

	rec := httptest.NewRecorder()

	bookByte, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(bookByte))

	req := httptest.NewRequest(http.MethodPost, "/books", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = bookCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateBookController_Failure1(t *testing.T) {
	book := models.Book{}

	bookRMock.Mock.On("CreateRepository", book).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	bookByte, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(bookByte))

	req := httptest.NewRequest(http.MethodPost, "/books", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err = bookCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestCreateBookController_Failure2(t *testing.T) {
	book := models.Book{}

	bookRMock.Mock.On("CreateRepository", book).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPost, "/books", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)

	err := bookCTest.CreateController(c)
	assert.Nil(t, err)
}

func TestUpdateBookController_Success(t *testing.T) {
	book := models.Book{
		Model: gorm.Model{
			ID: 1,
		},
		Title:       "Batman",
		Author:      "Sugeng",
		Description: "Mantap",
	}

	bookRMock.Mock.On("UpdateRepository", "1", book).Return(book, nil)

	rec := httptest.NewRecorder()

	bookByte, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(bookByte))

	req := httptest.NewRequest(http.MethodPut, "/books/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = bookCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateBookController_Failure1(t *testing.T) {
	book := models.Book{
		Model: gorm.Model{
			ID: 1,
		},
		Title:       "Batman",
		Author:      "Sugeng",
		Description: "Mantap",
	}

	bookRMock.Mock.On("UpdateRepository", "1", book).Return(book, nil)

	rec := httptest.NewRecorder()

	bookByte, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(bookByte))

	req := httptest.NewRequest(http.MethodPut, "/books/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err = bookCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateBookController_Failure2(t *testing.T) {
	book := models.Book{}

	bookRMock.Mock.On("UpdateRepository", "1", book).Return(book, nil)

	rec := httptest.NewRecorder()

	_, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string([]byte("qwe")))

	req := httptest.NewRequest(http.MethodPut, "/books/qwe", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = bookCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestUpdateBookController_Failure3(t *testing.T) {
	bookRMock = &repositories.IbookRepositoryMock{Mock: mock.Mock{}}
	bookSMock = services.NewBookService(bookRMock)
	bookCTest = NewBookController(bookSMock)
	book := models.Book{
		Model: gorm.Model{
			ID: 1,
		},
		Title:       "Batman",
		Author:      "Sugeng",
		Description: "Mantap",
	}

	bookRMock.Mock.On("UpdateRepository", "1", book).Return(nil, errors.New("something wrong"))

	rec := httptest.NewRecorder()

	bookByte, err := json.Marshal(book)
	if err != nil {
		t.Error(err)
	}

	reqBody := strings.NewReader(string(bookByte))

	req := httptest.NewRequest(http.MethodPut, "/books/1", reqBody)
	req.Header.Add("Content-type", "application/json")
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err = bookCTest.UpdateController(c)
	assert.Nil(t, err)
}

func TestDeleteBookController_Success(t *testing.T) {
	bookRMock.Mock.On("DeleteRepository", "2").Return(nil)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/books/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := bookCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteBookController_Failure1(t *testing.T) {
	bookRMock = &repositories.IbookRepositoryMock{Mock: mock.Mock{}}
	bookSMock = services.NewBookService(bookRMock)
	bookCTest = NewBookController(bookSMock)
	bookRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("book not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/books/2", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	err := bookCTest.DeleteController(c)

	assert.Nil(t, err)
}

func TestDeleteBookController_Failure2(t *testing.T) {
	bookRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("book not found"))
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/books/qwe", nil)

	e := echo.New()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("qwe")

	err := bookCTest.DeleteController(c)

	assert.Nil(t, err)
}
