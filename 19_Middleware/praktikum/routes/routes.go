package routes

import (
	"praktikum_section_19/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	// users
	e.GET("/users", controllers.GetUsersController)
	e.POST("/user", controllers.CreateUserController)
	e.GET("/user/:id", controllers.GetUserController)
	e.DELETE("/user/:id", controllers.DeleteUserController)
	e.PUT("/user/:id", controllers.UpdateUserController)

	// books
	e.GET("/books", controllers.GetBooksController)
	e.POST("/book", controllers.CreateBookController)
	e.GET("/book/:id", controllers.GetBookController)
	e.DELETE("/book/:id", controllers.DeleteBookController)
	e.PUT("/book/:id", controllers.UpdateBookController)

	return e
}
