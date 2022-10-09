package model

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserInput struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

func (input *UserInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
