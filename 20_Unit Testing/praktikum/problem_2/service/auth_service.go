package service

import (
	"praktikum_section_20/models"
	"praktikum_section_20/repository"
)

type AuthService struct {
	Repository repository.AuthRepository
}

func NewAuthService() AuthService {
	return AuthService{
		Repository: &repository.AuthRepositoryImpl{},
	}
}

func (a *AuthService) Register(input models.UserInput) models.User {
	return a.Repository.Register(input)
}

func (a *AuthService) Login(input models.UserInput) string {
	return a.Repository.Login(input)
}
