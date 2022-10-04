package service

import (
	"praktikum_section_20/models"
	"praktikum_section_20/repository"
)

type BookService struct {
	Repository repository.BookRepository
}

func NewBookService() BookService {
	return BookService{
		Repository: &repository.BookRepositoryImpl{},
	}
}

func (b *BookService) GetBooks() (interface{}, error) {
	return b.Repository.GetBooks()
}

func (b *BookService) GetBook(bookID int) (interface{}, error) {
	return b.Repository.GetBook(bookID)
}

func (b *BookService) CreateBook(book *models.Book) (interface{}, error) {
	return b.Repository.CreateBook(book)
}

func (b *BookService) UpdateBook(bookID int, newBook models.Book) (interface{}, error) {
	return b.Repository.UpdateBook(bookID, newBook)
}

func (b *BookService) DeleteBook(bookID int) error {
	return b.Repository.DeleteBook(bookID)
}
