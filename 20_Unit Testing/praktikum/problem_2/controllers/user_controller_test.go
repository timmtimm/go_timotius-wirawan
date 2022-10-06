package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"praktikum_section_20/config"
	"praktikum_section_20/models"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func seedUser() models.User {
	password, _ := bcrypt.GenerateFromPassword([]byte("password_user"), bcrypt.DefaultCost)

	var user models.User = models.User{
		Name:     "testing",
		Email:    "testing@gmail.com",
		Password: string(password),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		panic(err)
	}

	var createdUser models.User

	config.DB.Last(&createdUser)

	createdUser.Password = "password_user"

	return createdUser
}

func TestGetAllUsers_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/users",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"message\":",
	}}

	config.InitDB("testing")
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestGetUser_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/user",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"message\":",
	}}

	config.InitDB("testing")
	e := echo.New()

	user := seedUser()
	userID := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest(http.MethodGet, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userID)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestCreateUser_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/user",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"message\":",
	}}

	config.InitDB("testing")
	e := echo.New()

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("password_user"), bcrypt.DefaultCost)

	user := models.User{
		Name:     "budi",
		Email:    "budi@gmail.com",
		Password: string(encryptedPassword),
	}

	bodyRequest := []byte(`{"name":` + user.Name + `"email":"` + user.Email + `","password":"` + user.Password + `"}`)
	bodyReader := bytes.NewReader(bodyRequest)

	req := httptest.NewRequest(http.MethodPost, "/user", bodyReader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestDeleteUser_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/user",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"message\":",
	}}

	config.InitDB("testing")
	e := echo.New()

	user := seedUser()
	userID := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest(http.MethodPut, "/user", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userID)

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestUpdateUser_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectedStatus      int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/user",
		expectedStatus:      http.StatusOK,
		expectBodyStartWith: "{\"message\":",
	}}

	config.InitDB("testing")
	e := echo.New()

	user := seedUser()
	userID := strconv.Itoa(int(user.ID))

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("password_user_update"), bcrypt.DefaultCost)

	updatedUser := models.User{
		Name:     "budi setiawan",
		Email:    "budisetiawan@gmail.com",
		Password: string(encryptedPassword),
	}

	bodyRequest := []byte(`{"name":` + updatedUser.Name + `"email":"` + updatedUser.Email + `","password":"` + updatedUser.Password + `"}`)
	bodyReader := bytes.NewReader(bodyRequest)

	req := httptest.NewRequest(http.MethodPut, "/user", bodyReader)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userID)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, testCase.expectedStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}
