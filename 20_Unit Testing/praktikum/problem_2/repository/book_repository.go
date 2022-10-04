package repository

import (
	"praktikum_section_20/config"
	"praktikum_section_20/models"
)

type BookRepositoryImpl struct{}

func (b *BookRepositoryImpl) GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (b *BookRepositoryImpl) CreateBook(book *models.Book) (interface{}, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookRepositoryImpl) GetBook(bookID int) (interface{}, error) {
	var book models.Book

	if err := config.DB.First(&book, bookID).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookRepositoryImpl) DeleteBook(bookID int) error {
	var book models.Book

	if err := config.DB.Delete(&book, bookID).Error; err != nil {
		return err
	}

	return nil
}

func (b *BookRepositoryImpl) UpdateBook(bookID int, newBook models.Book) (interface{}, error) {
	var book models.Book

	if err := config.DB.Model(&book).Where("id = ?", bookID).Update(newBook).Error; err != nil {
		return nil, err
	}

	book.ID = uint(bookID)

	return book, nil
}
