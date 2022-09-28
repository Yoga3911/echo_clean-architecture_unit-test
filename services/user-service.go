package services

import (
	"day-13-orm/models"
	"day-13-orm/repositories"
)

type UserService interface {
	GetUsersService() ([]*models.User, error)
	GetUserService(id string) (*models.User, error)
	CreateService(user models.User) (*models.User, error)
	UpdateService(id string, userBody models.User) (*models.User, error)
	DeleteService(id string) error
}

type userService struct {
	userR repositories.UserRepository
}

func NewUserService(userR repositories.UserRepository) UserService {
	return &userService{
		userR: userR,
	}
}

func (u *userService) GetUsersService() ([]*models.User, error) {
	users, err := u.userR.GetUsersRepository()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) GetUserService(id string) (*models.User, error) {
	user, err := u.userR.GetUserRepository(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) CreateService(user models.User) (*models.User, error) {
	userR, err := u.userR.CreateRepository(user)
	if err != nil {
		return nil, err
	}

	return userR, nil
}

func (u *userService) UpdateService(id string, userBody models.User) (*models.User, error) {
	user, err := u.userR.UpdateRepository(id, userBody)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userService) DeleteService(id string) error {
	err := u.userR.DeleteRepository(id)
	if err != nil {
		return err
	}

	return nil
}
