package main

import (
	"github.com/kawan02/handlers"
	"github.com/kawan02/models"
	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()

	models.ConnectDatabase()

	r.GET("/books", handlers.FindBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/:id", handlers.FindBook)
	r.PUT("/books/:id", handlers.UpdateBook) // New
	r.DELETE("/books/:id", handlers.DeleteBook)
	r.DELETE("/books", handlers.DeleteBookTodos)

	r.Start(":8080")

}
