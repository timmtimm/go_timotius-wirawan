package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	useCase usecase.UserUsecase
}

func NewUserController(userUsecase usecase.UserUsecase) *userController {
	return &userController{
		userUsecase,
	}
}

func (u *userController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := u.useCase.GetAllUsers()

		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err,
			})
		}

		return c.JSON(200, echo.Map{
			"data": users,
		})
	}
}

func (u *userController) CreateUser() echo.HandlerFunc {
	user := model.User{}
	return func(c echo.Context) error {
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err := u.useCase.CreateUser(user)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err,
			})
		}
		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}
