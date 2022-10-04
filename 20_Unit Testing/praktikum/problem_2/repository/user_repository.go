package repository

import (
	"praktikum_section_20/config"
	"praktikum_section_20/models"
)

type UserRepositoryImpl struct{}

func (u *UserRepositoryImpl) GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) CreateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) GetUser(userID int) (interface{}, error) {
	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) DeleteUser(userID int) error {
	var user models.User

	if err := config.DB.Delete(&user, userID).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryImpl) UpdateUser(userID int, newUser models.User) (interface{}, error) {
	var user models.User

	if err := config.DB.Model(&user).Where("id = ?", userID).Update(newUser).Error; err != nil {
		return nil, err
	}

	user.ID = uint(userID)

	return user, nil
}
