package controllers

import (
	"net/http"

	"github.com/kawan02/models"
	"github.com/labstack/echo/v4"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

///FindBooks retornará todos os livros do nosso banco de dados.
func FindBooks(c echo.Context) error {
	var books []models.Book
	models.DB.Find(&books) // Select * From book

	return c.JSON(http.StatusOK, echo.Map{
		"Listando todos os livros que estão em nossa biblioteca": books})
}

// Encontra um livro
func FindBook(c echo.Context) error {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Mensagem": "Registro não encontrado!",
			"error":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"Pesquisa concluída:": book})
}

func CreateBook(c echo.Context) error {
	// Valida input
	var input CreateBookInput

	// Validamos o corpo da solicitação se os dados forem inválidos, ele retornará um erro 400

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Mensagem": "Ocorreu um erro inesperado",
			"Error":    err.Error(),
		})
	}

	//Create book --> cria livro
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	return c.JSON(http.StatusOK, echo.Map{"Livro criado": book})

}

// Atualiza um livro
func UpdateBook(c echo.Context) error {

	// Obtém o modelo se existir
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Mensagem:": "Registro não encontrado!",
			"Error:":    err.Error(),
		})
	}

	// Validar input
	var input UpdateBookInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Mensagem": "Ocorreu um erro inesperado",
			"error":    err.Error(),
		})
	}

	UpdateBookInput := models.Book{Title: input.Title, Author: input.Author}

	models.DB.Model(&book).Updates(&UpdateBookInput)
	return c.JSON(http.StatusOK, echo.Map{"Livro atualizado:": book})

}

// Delete book --> excluir um livro
func DeleteBook(c echo.Context) error {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"Mensagem": "Registro não encontrado!",
			"Error":    err.Error(),
		})

	}

	models.DB.Delete(&book)

	return c.JSON(http.StatusOK, echo.Map{"Livro deletado:": true})
}

// DeleteBookTodos vai excluir todos os livros do nosso banco de dados.
func DeleteBookTodos(c echo.Context) error {
	var book []models.Book
	models.DB.Find(&book)

	models.DB.Delete(&book)
	return c.JSON(http.StatusOK, echo.Map{"Todos os livros foram excluidos com sucesso!": book})
}
