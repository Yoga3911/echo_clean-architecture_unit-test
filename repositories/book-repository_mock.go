package repositories

import (
	"day-13-orm/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock interface {
	GetBooksRepository() ([]*models.Book, error)
	GetBookRepository(id string) (*models.Book, error)
	CreateRepository(Book models.Book) (*models.Book, error)
	UpdateRepository(id string, BookBody models.Book) (*models.Book, error)
	DeleteRepository(id string) error
}

type IbookRepositoryMock struct {
	Mock mock.Mock
}

func NewBookRepositoryMock(mock mock.Mock) BookRepositoryMock {
	return &IbookRepositoryMock{
		Mock: mock,
	}
}

func (b *IbookRepositoryMock) GetBooksRepository() ([]*models.Book, error) {
	args := b.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	books := args.Get(0).([]*models.Book)

	return books, nil
}

func (b *IbookRepositoryMock) GetBookRepository(id string) (*models.Book, error) {
	args := b.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	book := args.Get(0).(models.Book)

	return &book, nil
}

func (u *IbookRepositoryMock) CreateRepository(bookData models.Book) (*models.Book, error) {
	args := u.Mock.Called(bookData)
	if args.Get(0) == nil {
		return nil, args.Get(1).(error)
	}

	book := args.Get(0).(models.Book)

	return &book, nil
}

func (u *IbookRepositoryMock) UpdateRepository(id string, bookData models.Book) (*models.Book, error) {
	args := u.Mock.Called(id, bookData)
	if args.Get(0) == nil {
		return nil,  args.Get(1).(error)
	}

	book := args.Get(0).(models.Book)

	return &book, nil
}

func (u *IbookRepositoryMock) DeleteRepository(id string) error {
	args := u.Mock.Called(id)
	if args.Get(0) != nil {
		return fmt.Errorf("must nil")
	}

	return nil
}
