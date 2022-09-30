package service

import (
	"praktikum_section_19/models"
	"praktikum_section_19/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func NewUserService() UserService {
	return UserService{
		Repository: &repository.UserRepositoryImpl{},
	}
}

func (u *UserService) GetUsers() (interface{}, error) {
	return u.Repository.GetUsers()
}

func (u *UserService) GetUser(userID int) (interface{}, error) {
	return u.Repository.GetUser(userID)
}

func (u *UserService) CreateUser(user *models.User) (interface{}, error) {
	return u.Repository.CreateUser(user)
}

func (u *UserService) UpdateUser(userID int, newUser models.User) (interface{}, error) {
	return u.Repository.UpdateUser(userID, newUser)
}

func (u *UserService) DeleteUser(userID int) error {
	return u.Repository.DeleteUser(userID)
}
