package database

import (
	"praktikum_section_18/config"
	"praktikum_section_18/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(userID int) (interface{}, error) {
	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(userID int) error {
	var user models.User

	if err := config.DB.Delete(&user, userID).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(userID int, newUser models.User) (interface{}, error) {
	var user models.User

	if err := config.DB.Model(&user).Where("id = ?", userID).Update(newUser).Error; err != nil {
		return nil, err
	}

	user.ID = uint(userID)

	return user, nil
}