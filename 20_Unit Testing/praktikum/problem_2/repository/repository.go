package repository

import (
	"praktikum_section_20/models"
)

type AuthRepository interface {
	Register(input models.UserInput) models.User
	Login(input models.UserInput) string
}

type BookRepository interface {
	GetBooks() (interface{}, error)
	CreateBook(book *models.Book) (interface{}, error)
	GetBook(bookID int) (interface{}, error)
	DeleteBook(bookID int) error
	UpdateBook(bookID int, newBook models.Book) (interface{}, error)
}

type UserRepository interface {
	GetUsers() (interface{}, error)
	CreateUser(user *models.User) (interface{}, error)
	GetUser(userID int) (interface{}, error)
	DeleteUser(userID int) error
	UpdateUser(userID int, newUser models.User) (interface{}, error)
}
