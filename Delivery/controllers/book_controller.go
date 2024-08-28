package controllers

import (
	domain "LoanTrackerAPI/Domain"
	"github.com/gin-gonic/gin"
)

type bookController struct {
	bookUsecase domain.BookUsecase
}

// CreateBook implements domain.BookController.
func (b *bookController) CreateBook(c *gin.Context) {
	var book *domain.Book
	err := c.BindJSON(&book)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid request", Status: 400, Error: err.Error()})
		return
	}
	response := b.bookUsecase.CreateBook(book)
	HandleResponse(c, response)
}

// DeleteBook implements domain.BookController.
func (b *bookController) DeleteBook(c *gin.Context) {
	panic("unimplemented")
}

// GetAllAvailableBooks implements domain.BookController.
func (b *bookController) GetAllAvailableBooks(c *gin.Context) {
	panic("unimplemented")
}

// GetAllBooks implements domain.BookController.
func (b *bookController) GetAllBooks(c *gin.Context) {
	panic("unimplemented")
}

// GetBookByID implements domain.BookController.
func (b *bookController) GetBookByID(c *gin.Context) {
	panic("unimplemented")
}

// UpdateBook implements domain.BookController.
func (b *bookController) UpdateBook(c *gin.Context) {
	panic("unimplemented")
}

// NewBookController creates a new instance of bookController.
func NewBookController(b domain.BookUsecase) domain.BookController {
	return &bookController{
		bookUsecase: b,
	}
}
