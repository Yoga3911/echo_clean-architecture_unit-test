package services

import (
	"day-13-orm/models"
	"day-13-orm/repositories"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	bookRMock = &repositories.IbookRepositoryMock{Mock: mock.Mock{}}
	bookSMock = NewBookService(bookRMock)
)

func TestGetBooksService_Success(t *testing.T) {
	booksMP := []*models.Book{
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

	booksM := []models.Book{
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

	bookRMock.Mock.On("GetBooksRepository").Return(booksMP)
	books, err := bookSMock.GetBooksService()

	assert.Nil(t, err)
	assert.NotNil(t, books)

	assert.Equal(t, booksM[0].Title, books[0].Title)
	assert.Equal(t, booksM[0].Author, books[0].Author)
	assert.Equal(t, booksM[0].Description, books[0].Description)
}

func TestGetBooksService_Failure(t *testing.T) {
	bookRMock = &repositories.IbookRepositoryMock{Mock: mock.Mock{}}
	bookSMock = NewBookService(bookRMock)
	bookRMock.Mock.On("GetBooksRepository").Return(nil, errors.New("get all books failed"))
	books, err := bookSMock.GetBooksService()

	assert.Nil(t, books)
	assert.NotNil(t, err)
}

func TestGetBookService_Success(t *testing.T) {
	book := models.Book{
		Title:       "Batman",
		Author:      "Boy",
		Description: "Buku action",
	}

	bookRMock.Mock.On("GetBookRepository", "1").Return(book, nil)
	books, err := bookSMock.GetBookService("1")

	assert.Nil(t, err)
	assert.NotNil(t, books)

	assert.Equal(t, book.Title, books.Title)
	assert.Equal(t, book.Author, books.Author)
	assert.Equal(t, book.Description, books.Description)
}

func TestGetBookService_Failure(t *testing.T) {
	bookRMock.Mock.On("GetBookRepository", "3").Return(nil, fmt.Errorf("book not found"))
	book, err := bookSMock.GetBookService("3")

	assert.NotNil(t, err)
	assert.Nil(t, book)
}

func TestCreateBookService_Success(t *testing.T) {
	book := models.Book{
		Title:       "Batman",
		Author:      "Boy",
		Description: "dsa",
	}

	bookRMock.Mock.On("CreateRepository", book).Return(book, nil)
	books, err := bookSMock.CreateService(book)

	assert.Nil(t, err)
	assert.NotNil(t, books)

	assert.Equal(t, book.Title, books.Title)
	assert.Equal(t, book.Author, books.Author)
	assert.Equal(t, book.Description, books.Description)
}

func TestCreateBookService_Failure(t *testing.T) {
	book := models.Book{
		Title:       "Batman32",
		Author:      "Bo321y",
		Description: "dsa321",
	}

	bookRMock.Mock.On("CreateRepository", book).Return(nil, fmt.Errorf("create book failed"))
	books, err := bookSMock.CreateService(book)

	assert.Nil(t, books)
	assert.NotNil(t, err)
}

func TestUpdateBookService_Success(t *testing.T) {
	book := models.Book{
		Title:       "Batman",
		Author:      "qwe",
		Description: "dsadass",
	}

	bookRMock.Mock.On("UpdateRepository", "1", book).Return(book, nil)
	books, err := bookSMock.UpdateService("1", book)

	assert.Nil(t, err)
	assert.NotNil(t, books)

	assert.Equal(t, book.ID, books.ID)
	assert.Equal(t, book.Title, books.Title)
	assert.Equal(t, book.Description, books.Description)
	assert.Equal(t, book.Author, books.Author)
}

func TestUpdateBookService_Failure(t *testing.T) {
	book := models.Book{
		Title:       "Batman123",
		Author:      "qwe321",
		Description: "dsadas32",
	}

	bookRMock.Mock.On("UpdateRepository", "2", book).Return(nil, fmt.Errorf("book not found"))
	books, err := bookSMock.UpdateService("2", book)

	assert.Nil(t, books)
	assert.NotNil(t, err)
}

func TestDeleteBookService_Success(t *testing.T) {
	bookRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := bookSMock.DeleteService("1")

	assert.Nil(t, err)
}

func TestDeleteBookService_Failure(t *testing.T) {
	bookRMock.Mock.On("DeleteRepository", "2").Return(fmt.Errorf("book not found"))
	err := bookSMock.DeleteService("2")

	assert.NotNil(t, err)
}
