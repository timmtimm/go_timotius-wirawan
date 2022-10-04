package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	gorm.Model
	Title         string `json:"title"`
	Writer        string `json:"writer"`
	Publisher     string `json:"publisher"`
	Genre         string `json:"genre"`
	PublishedYear int    `json:"published_year"`
	TotalPage     int    `json:"total_page"`
	Price         int    `json:"price"`
}

type BookInput struct {
	gorm.Model
	Title         string `json:"title" form:"title" validate:"required"`
	Writer        string `json:"writer" form:"writer" validate:"required"`
	Publisher     string `json:"publisher" form:"publisher" validate:"required"`
	Genre         string `json:"genre" form:"genre" validate:"required"`
	PublishedYear int    `json:"published_year" form:"published_year" validate:"required"`
	TotalPage     int    `json:"total_page" form:"total_page" validate:"required"`
	Price         int    `json:"price" form:"price" validate:"required"`
}

func (input *BookInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
