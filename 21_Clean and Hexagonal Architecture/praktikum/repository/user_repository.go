package repository

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user model.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (u *UserRepositoryImpl) GetAllUsers() ([]model.User, error) {
	users := make([]model.User, 0)

	if err := u.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) CreateUser(user model.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
