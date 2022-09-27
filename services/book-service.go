package services

import (
	"day-13-orm/models"
	"day-13-orm/repositories"
)

type BookService interface {
	GetBooksService() ([]*models.Book, error)
	GetBookService(id string) (*models.Book, error)
	CreateService(Book models.Book) (*models.Book, error)
	UpdateService(id string, BookBody models.Book) (*models.Book, error)
	DeleteService(id string) error
}

type bookService struct {
	BookR repositories.BookRepository
}

func NewBookService(BookR repositories.BookRepository) BookService {
	return &bookService{
		BookR: BookR,
	}
}

func (b *bookService) GetBooksService() ([]*models.Book, error) {
	Books, err := b.BookR.GetBooksRepository()
	if err != nil {
		return nil, err
	}

	return Books, nil
}

func (b *bookService) GetBookService(id string) (*models.Book, error) {
	Book, err := b.BookR.GetBookRepository(id)
	if err != nil {
		return nil, err
	}

	return Book, nil
}

func (b *bookService) CreateService(Book models.Book) (*models.Book, error) {
	BookR, err := b.BookR.CreateRepository(Book)
	if err != nil {
		return nil, err
	}

	return BookR, nil
}

func (b *bookService) UpdateService(id string, BookBody models.Book) (*models.Book, error) {
	Book, err := b.BookR.UpdateRepository(id, BookBody)
	if err != nil {
		return Book, err
	}

	return Book, nil
}

func (b *bookService) DeleteService(id string) error {
	err := b.BookR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
