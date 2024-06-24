package controller

import (
	"net/http"
	"simple-clean-architecture/middleware"
	"simple-clean-architecture/model"
	"simple-clean-architecture/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	useCase        usecase.UserUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

func (b *UserController) Route() {
	b.rg.POST("/users", b.authMiddleware.RequireToken("admin"), b.createUser)
	b.rg.GET("/users", b.authMiddleware.RequireToken("admin"), b.getAllUser)
	b.rg.GET("/users/:id", b.authMiddleware.RequireToken("admin"), b.getUserById)
}

func (b *UserController) createUser(c *gin.Context) {
	var payload model.UserCredential
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	book, err := b.useCase.RegisterNewUser(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (b *UserController) getAllUser(c *gin.Context) {
	books, err := b.useCase.FindAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to retrieve data books"})
		return
	}

	if len(books) > 0 {
		c.JSON(http.StatusOK, books)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "List book empty"})
}

func (b *UserController) getUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := b.useCase.FindUserById(uint32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to get book by ID"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func NewUserController(useCase usecase.UserUseCase, rg *gin.RouterGroup, am middleware.AuthMiddleware) *UserController {
	return &UserController{useCase: useCase, rg: rg, authMiddleware: am}
}
