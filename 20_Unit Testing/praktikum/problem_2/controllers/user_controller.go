package controllers

import (
	"fmt"
	"net/http"
	"praktikum_section_20/models"
	"praktikum_section_20/service"
	"strconv"

	"github.com/labstack/echo"
)

var userService service.UserService = service.NewUserService()

// get all users
func GetUsersController(c echo.Context) error {
	users, err := userService.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	FormUser := models.User{}
	c.Bind(&FormUser)

	user, err := userService.CreateUser(&FormUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	user, err := userService.GetUser(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success get user with id %d", userID),
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	err := userService.DeleteUser(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success delete user with id %d", userID),
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("id"))

	var newUser models.User
	newUser.Name = c.FormValue("name")
	newUser.Email = c.FormValue("email")
	newUser.Password = c.FormValue("password")

	user, err := userService.UpdateUser(userID, newUser)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("success update user with id %d", userID),
		"user":    user,
	})
}
