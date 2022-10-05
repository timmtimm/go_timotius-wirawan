package routes

import (
	"praktikum_section_20/controllers"
	"praktikum_section_20/utils"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// unauthenticated routes
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	e.POST("/user", controllers.CreateUserController)
	e.GET("/books", controllers.GetBooksController)
	e.GET("/book/:id", controllers.GetBookController)

	// authenticated routes
	privateRoutes := e.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(utils.GetConfig("SECRET_JWT")),
	}))

	// users
	privateRoutes.GET("/users", controllers.GetUsersController)
	privateRoutes.GET("/user/:id", controllers.GetUserController)
	privateRoutes.DELETE("/user/:id", controllers.DeleteUserController)
	privateRoutes.PUT("/user/:id", controllers.UpdateUserController)

	// books
	privateRoutes.POST("/book", controllers.CreateBookController)
	privateRoutes.DELETE("/book/:id", controllers.DeleteBookController)
	privateRoutes.PUT("/book/:id", controllers.UpdateBookController)

	return e
}
