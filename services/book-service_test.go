package services

import (
	"day-13-orm/models"
	"day-13-orm/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var bookRMock = &repositories.IbookRepositoryMock{Mock: mock.Mock{}}
var bookSMock = NewBookService(bookRMock)

func TestGetBooksService(t *testing.T) {
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

func TestGetBookService(t *testing.T) {
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

func TestCreateBookService(t *testing.T) {
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

func TestUpdateBookService(t *testing.T) {
	book := models.Book{
		Model: gorm.Model{
			ID: 1,
		},
		Title:       "Batman",
		Author:      "qwe",
		Description: "dsadass",
	}

	bookRMock.Mock.On("UpdateRepository", "1", book).Return(book, nil)
	books, err := bookSMock.UpdateService("1", book)

	assert.Nil(t, err)
	assert.NotNil(t, books)

	assert.Equal(t, uint(1), books.ID)
	assert.Equal(t, book.Title, books.Title)
	assert.Equal(t, book.Description, books.Description)
	assert.Equal(t, book.Author, books.Author)
}

func TestDeleteBookService(t *testing.T) {
	bookRMock.Mock.On("DeleteRepository", "1").Return(nil)
	err := bookSMock.DeleteService("1")

	assert.Nil(t, err)
}
