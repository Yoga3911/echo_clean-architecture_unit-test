package repositories

import (
	"day-13-orm/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsersRepository() ([]models.User, error)
	GetUserRepository(id string) (models.User, error)
	CreateRepository() (models.User, error)
	// UpdateRepository(id string) (models.User, error)
	// DeleteRepository(id string) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (u *userRepository) GetUsersRepository() ([]models.User, error) {
	var users []models.User

	if err := u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetUserRepository(id string) (models.User, error) {
	var user models.User
	
	if err := u.DB.Where("ID = ?", id).Take(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) CreateRepository() (models.User, error) {
	var user models.User
	if err := u.DB.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// func (u *userRepository) UpdateRepository(id string) (models.User, error) 
// func (u *userRepository) DeleteRepository(id string) error
