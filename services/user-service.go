package services

import (
	"day-13-orm/models"
	"day-13-orm/repositories"
)

type UserService interface {
	GetUsersService() ([]models.User, error)
	GetUserService(id string) (models.User, error)
	CreateService() (models.User, error)
	// UpdateService(id string) (models.User, error)
	// DeleteService(id string) error
}

type userService struct {
	userR repositories.UserRepository
}

func NewUserService(userR repositories.UserRepository) UserService {
	return &userService{
		userR: userR,
	}
}

func (u *userService) GetUsersService() ([]models.User, error) {
	users, err := u.userR.GetUsersRepository()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) GetUserService(id string) (models.User, error) {
	user, err := u.userR.GetUserRepository(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userService) CreateService() (models.User, error) {
	user, err := u.userR.CreateRepository()
	if err != nil {
		return user, err
	}

	return user, nil
}

// func (u *userService) UpdateService(id string) (models.User, error)
// func (u *userService) DeleteService(id string) error