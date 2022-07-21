package main

import (
	"github.com/kawan02/controllers"
	"github.com/kawan02/models"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook) // new
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.DELETE("/books", controllers.DeleteBookTodos)

	r.Start(":8080")

}
