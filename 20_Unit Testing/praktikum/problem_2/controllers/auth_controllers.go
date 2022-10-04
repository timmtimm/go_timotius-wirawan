package controllers

import (
	"net/http"
	"praktikum_section_20/models"
	"praktikum_section_20/service"

	"github.com/labstack/echo"
)

var authService service.AuthService = service.NewAuthService()

func Register(c echo.Context) error {
	var userInput *models.UserInput = new(models.UserInput)

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "please fullfill register form",
		})
	}

	user := authService.Register(*userInput)

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	var userInput *models.UserInput = new(models.UserInput)

	if err := c.Bind(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "missing email or password",
		})
	}

	token := authService.Login(*userInput)

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "wrong email or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
