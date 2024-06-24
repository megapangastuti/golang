// ToDo :
// 1. Mendeklarasikan nama package pada file book_controller.go
// 2. Mendeklarasikan struct bernama BookController
// 3. Mendeklarasikan function Route
// 4. Membuat detail handlernya
// 5. Membuat constructor bernama NewBookController

package controller

import (
	"net/http"
	"simple-clean-architecture/middleware"
	"simple-clean-architecture/model"
	"simple-clean-architecture/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	useCase usecase.BookUseCase
	rg      *gin.RouterGroup
}

func (b *BookController) Route() {
	b.rg.POST("/books", middleware.AuthMiddleware("admin"), b.createNewBook)
	b.rg.GET("/books", middleware.AuthMiddleware("admin", "user"), b.getAllBook)
	b.rg.GET("/books/:id", middleware.AuthMiddleware("admin", "user"), b.getBookById)
	b.rg.PUT("/books", middleware.AuthMiddleware("admin"), b.updateBookById)
	b.rg.DELETE("/books/:id", middleware.AuthMiddleware("admin"), b.deleteBookById)
}

func (b *BookController) createNewBook(c *gin.Context) {
	var payload model.Book

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	book, err := b.useCase.CreateNewBook(payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to create book data"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (b *BookController) getAllBook(c *gin.Context) {
	books, err := b.useCase.GetAllBook()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to retrieve books data"})
		return
	}

	if len(books) > 0 {
		c.JSON(http.StatusOK, books)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List book is empty"})
}

func (b *BookController) getBookById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := b.useCase.GetBookById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to get boo by ID"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (b *BookController) updateBookById(c *gin.Context) {
	var payload model.Book

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	book, err := b.useCase.UpdateBookById(payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (b *BookController) deleteBookById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := b.useCase.DeleteBookById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewBookController(useCase usecase.BookUseCase, rg *gin.RouterGroup) *BookController {
	return &BookController{useCase: useCase, rg: rg}
}
