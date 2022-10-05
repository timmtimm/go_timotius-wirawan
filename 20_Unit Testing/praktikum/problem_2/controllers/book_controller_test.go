package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"praktikum_section_20/config"
	"praktikum_section_20/models"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// Title         string `json:"title"`
// Writer        string `json:"writer"`
// Publisher     string `json:"publisher"`
// Genre         string `json:"genre"`
// PublishedYear int    `json:"published_year"`
// TotalPage     int    `json:"total_page"`
// Price         int    `json:"price"`

func seedBook() models.Book {
	var book models.Book = models.Book{
		Title:         "Dilanda Rindu",
		Writer:        "Dilan",
		Publisher:     "Gramedia",
		Genre:         "Action",
		PublishedYear: 2003,
		TotalPage:     100,
		Price:         75000,
	}

	if err := config.DB.Create(&book).Error; err != nil {
		panic(err)
	}

	var createdBook models.Book

	config.DB.Last(&createdBook)

	return createdBook
}

func TestGetAllBooks_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/books",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"books\":",
	}}

	config.InitTestDB()
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestGetBook_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/book",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"book\":",
	}}

	config.InitTestDB()
	e := echo.New()

	book := seedBook()
	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodGet, "/book", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestCreateBook_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/book",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"book\":",
	}}

	config.InitTestDB()
	e := echo.New()

	var book models.Book = models.Book{
		Title:         "Dilanda Rindu",
		Writer:        "Dilan",
		Publisher:     "Gramedia",
		Genre:         "Action",
		PublishedYear: 2003,
		TotalPage:     100,
		Price:         75000,
	}

	bodyRequest := []byte(
		`{"title":` + book.Title +
			`"writer":"` + book.Writer +
			`","publisher":"` + book.Publisher +
			`","genre":"` + book.Genre +
			`","published_year":"` + fmt.Sprint(book.PublishedYear) +
			`","total_page":"` + fmt.Sprint(book.TotalPage) +
			`","price":"` + fmt.Sprint(book.Price) +
			`"}`)
	bodyReader := bytes.NewReader(bodyRequest)

	req := httptest.NewRequest(http.MethodPost, "/book", bodyReader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestDeleteBook_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/book",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"message\":",
	}}

	config.InitTestDB()
	e := echo.New()

	book := seedBook()
	bookID := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodPut, "/book", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestUpdateBook_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/book",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"book\":",
	}}

	config.InitTestDB()
	e := echo.New()

	book := seedBook()
	bookID := strconv.Itoa(int(book.ID))

	var updatedbook = models.Book{
		Title:         "Dilanda Rindu",
		Writer:        "Dilan",
		Publisher:     "Gramedia",
		Genre:         "Action",
		PublishedYear: 2003,
		TotalPage:     100,
		Price:         75000,
	}

	bodyRequest := []byte(
		`{"title":` + updatedbook.Title +
			`"writer":"` + updatedbook.Writer +
			`","publisher":"` + updatedbook.Publisher +
			`","genre":"` + updatedbook.Genre +
			`","published_year":"` + fmt.Sprint(updatedbook.PublishedYear) +
			`","total_page":"` + fmt.Sprint(updatedbook.TotalPage) +
			`","price":"` + fmt.Sprint(updatedbook.Price) +
			`"}`)
	bodyReader := bytes.NewReader(bodyRequest)

	req := httptest.NewRequest(http.MethodPut, "/book", bodyReader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookID)

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}
