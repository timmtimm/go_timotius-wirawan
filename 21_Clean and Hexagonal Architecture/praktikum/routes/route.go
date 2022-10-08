package routes

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/repository"
	"belajar-go-echo/usecase"
	"belajar-go-echo/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)

	userService := usecase.NewUserUsecase(userRepository)

	userController := controller.NewUserController(userService)

	e.POST("/register", userController.Register())
	e.POST("/login", userController.Login())

	privateRoutes := e.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(util.GetConfig("SECRET_JWT_KEY")),
	}))

	privateRoutes.GET("/users", userController.GetAllUsers())
	privateRoutes.POST("/users", userController.CreateUser())
	e.Start(":8080")
}
