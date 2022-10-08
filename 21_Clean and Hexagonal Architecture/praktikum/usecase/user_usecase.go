package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user model.User) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (u *userUsecase) GetAllUsers() ([]model.User, error) {
	users, err := u.userRepository.GetAllUsers()

	return users, err
}

func (u *userUsecase) CreateUser(user model.User) error {
	err := u.userRepository.CreateUser(user)

	return err
}
