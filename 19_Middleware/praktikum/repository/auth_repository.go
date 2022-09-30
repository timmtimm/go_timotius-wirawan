package repository

import (
	"praktikum_section_19/auth"
	"praktikum_section_19/config"
	"praktikum_section_19/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryImpl struct{}

func (a *AuthRepositoryImpl) Register(input models.UserInput) models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser models.User = models.User{
		Email:    input.Email,
		Password: string(password),
	}

	var createdUser models.User = models.User{}

	result := config.DB.Create(&newUser)

	result.Last(&createdUser)

	return createdUser
}

func (a *AuthRepositoryImpl) Login(input models.UserInput) string {
	var user models.User = models.User{}

	config.DB.First(&user, "email = ?", input.Email)

	if user.ID == 0 {
		return ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return ""
	}

	token := auth.CreateToken(user.ID)

	return token
}
