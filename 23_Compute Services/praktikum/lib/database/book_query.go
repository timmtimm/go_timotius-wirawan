package database

import (
	"praktikum_section_18/config"
	"praktikum_section_18/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func CreateBook(book *models.Book) (interface{}, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func GetBook(bookID int) (interface{}, error) {
	var book models.Book

	if err := config.DB.First(&book, bookID).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(bookID int) error {
	var book models.Book

	if err := config.DB.Delete(&book, bookID).Error; err != nil {
		return err
	}

	return nil
}

func UpdateBook(bookID int, newBook models.Book) (interface{}, error) {
	var book models.Book

	if err := config.DB.Model(&book).Where("id = ?", bookID).Update(newBook).Error; err != nil {
		return nil, err
	}

	book.ID = uint(bookID)

	return book, nil
}