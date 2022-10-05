package controllers

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"praktikum_section_20/config"
	"praktikum_section_20/models"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func seedAuth() models.User {
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

func TestRegister_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectStatus        int
		expectBodyStartWith string
	}{{
		name:                "success_register",
		path:                "/register",
		expectStatus:        http.StatusCreated,
		expectBodyStartWith: "{\"ID\":",
	}}

	config.InitTestDB()
	e := echo.New()

	user := models.UserInput{
		Email:    "setiawan@gmail.com",
		Password: "rahasia_setiawan",
	}

	bodyRequest := []byte(`{"email":"` + user.Email + `","password":"` + user.Password + `"}`)
	bodyReader := bytes.NewReader(bodyRequest)

	req := httptest.NewRequest(http.MethodPost, "/register", bodyReader)
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, Register(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()

			fmt.Println(body)

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestRegister_Failed(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectStatus        int
		expectBodyStartWith string
	}{{
		name:                "failed_invalid_request",
		path:                "/register",
		expectStatus:        http.StatusBadRequest,
		expectBodyStartWith: "{\"message\":",
	}, {
		name:                "failed_form_not_fullfilled",
		path:                "/register",
		expectStatus:        http.StatusBadRequest,
		expectBodyStartWith: "{\"message\":",
	}}

	config.InitTestDB()
	e := echo.New()

	user := models.UserInput{
		Email:    "setiawan@gmail.com",
		Password: "rahasia_setiawan",
	}

	bodyRequest := [][]byte{
		[]byte(`{}`),
		[]byte(`{"email":"` + user.Email + `"}`)}

	for index, testCase := range testCases {
		bodyReader := bytes.NewReader(bodyRequest[index])

		req := httptest.NewRequest(http.MethodGet, "/register", bodyReader)
		req.Header.Add("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		if assert.NoError(t, Register(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestLogin_Success(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectStatus        int
		expectBodyStartWith string
	}{{
		name:                "success",
		path:                "/login",
		expectStatus:        http.StatusOK,
		expectBodyStartWith: "{\"token\":",
	}}

	config.InitTestDB()
	e := echo.New()

	user := seedAuth()

	bodyRequest := []byte(`{"email":"` + user.Email + `","password":"` + user.Password + `"}`)
	bodyReader := bytes.NewReader(bodyRequest)

	req := httptest.NewRequest(http.MethodGet, "/login", bodyReader)
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, Login(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}

func TestLogin_Failed(t *testing.T) {
	var testCases = []struct {
		name                string
		path                string
		expectStatus        int
		expectBodyStartWith string
	}{{
		name:                "failed_invalid_request",
		path:                "/login",
		expectStatus:        http.StatusBadRequest,
		expectBodyStartWith: "{\"message\":",
	}, {
		name:                "failed_missing_body",
		path:                "/login",
		expectStatus:        http.StatusBadRequest,
		expectBodyStartWith: "{\"message\":",
	}, {
		name:                "failed_wrong_email_password",
		path:                "/login",
		expectStatus:        http.StatusUnauthorized,
		expectBodyStartWith: "{\"message\":",
	}}

	config.InitTestDB()
	e := echo.New()

	bodyRequest := [][]byte{
		[]byte(`{}`),
		[]byte(`{"email":` + "" + `"password"` + "" + `}`),
		[]byte(`{"email":"` + "agus@gmail.com" + `","password":"` + "suga123" + `"}`)}

	for index, testCase := range testCases {
		bodyReader := bytes.NewReader(bodyRequest[index])

		req := httptest.NewRequest(http.MethodPost, "/login", bodyReader)
		req.Header.Add("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, Login(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartWith))
		}
	}
}
