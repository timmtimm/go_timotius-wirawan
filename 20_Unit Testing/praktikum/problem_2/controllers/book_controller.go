package controllers

import (
	"fmt"
	"net/http"
	"praktikum_section_20/models"
	"praktikum_section_20/service"
	"strconv"

	"github.com/labstack/echo"
)

var bookService service.BookService = service.NewBookService()

// get all books
func GetBooksController(c echo.Context) error {
	books, err := bookService.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	FormBook := models.Book{}
	c.Bind(&FormBook)

	book, err := bookService.CreateBook(&FormBook)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	bookID, _ := strconv.Atoi(c.Param("id"))

	book, err := bookService.GetBook(bookID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success get book with id %d", bookID),
		"book":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	bookID, _ := strconv.Atoi(c.Param("id"))

	err := bookService.DeleteBook(bookID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success delete books with id %d", bookID),
	})
}

// update user by id
func UpdateBookController(c echo.Context) error {
	bookID, _ := strconv.Atoi(c.Param("id"))

	var newBook models.Book
	newBook.Title = c.FormValue("title")
	newBook.Writer = c.FormValue("writer")
	newBook.Publisher = c.FormValue("publisher")
	newBook.Genre = c.FormValue("Genre")
	newBook.PublishedYear, _ = strconv.Atoi(c.FormValue("published_year"))
	newBook.TotalPage, _ = strconv.Atoi(c.FormValue("total_page"))
	newBook.Price, _ = strconv.Atoi(c.FormValue("price"))

	book, err := bookService.UpdateBook(bookID, newBook)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success update user with id %d", bookID),
		"book":    book,
	})
}
