package repositories

import (
	"day-13-orm/models"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock interface {
	GetUsersRepository() ([]*models.User, error)
	GetUserRepository(id string) (*models.User, error)
	CreateRepository(user models.User) (models.User, error)
	UpdateRepository(id string, userBody models.User) (*models.User, error)
	DeleteRepository(id string) error
}

type IuserRepositoryMock struct {
	Mock mock.Mock
}

func NewUserRepositoryMock(mock mock.Mock) UserRepositoryMock {
	return &IuserRepositoryMock{
		Mock: mock,
	}
}

func (u *IuserRepositoryMock) GetUsersRepository() ([]*models.User, error) {
	args := u.Mock.Called()
	if args.Get(0) == nil {
		return nil, nil
	}

	users := args.Get(0).([]*models.User)

	return users, nil
}

func (u *IuserRepositoryMock) GetUserRepository(id string) (*models.User, error) {
	args := u.Mock.Called(id)
	if args.Get(0) == nil {
		return nil, nil
	}

	user := args.Get(0).(models.User)

	return &user, nil
}

func (u *IuserRepositoryMock) CreateRepository(user models.User) (models.User, error) {
	return models.User{}, nil
}
func (u *IuserRepositoryMock) UpdateRepository(id string, userBody models.User) (*models.User, error) {
	return nil, nil
}
func (u *IuserRepositoryMock) DeleteRepository(id string) error {
	return nil
}