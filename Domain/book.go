package domain

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StatusType represents the status of a book or borrow request
type StatusType string

const (
	StatusPending   StatusType = "pending"
	StatusApproved  StatusType = "approved"
	StatusRejected  StatusType = "rejected"
	StatusAvailable StatusType = "available"
	StatusBorrowed  StatusType = "borrowed"
)

type Book struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Author      string             `json:"author"`
	Description string             `json:"description"`
	IsAvailable bool               `json:"is_available"`
}

type BorrowRequest struct {
	ID     primitive.ObjectID `json:"id"`
	BookID primitive.ObjectID `json:"book_id"`
	UserID primitive.ObjectID `json:"user_id"`
	Status StatusType         `json:"status"`
}
type BorrowBookRequest struct {
	BookID string `json:"book_id"`
}
type BookRepository interface {
	CreateBook(ctx context.Context, book *Book) error
	GetBookByID(ctx context.Context, id primitive.ObjectID) (*Book, error)
	GetAllBooks(ctx context.Context) ([]*Book, error)
	GetAllAvailableBooks(ctx context.Context) ([]*Book, error)
	UpdateBook(ctx context.Context, book *Book) error
	DeleteBook(ctx context.Context, id primitive.ObjectID) error
}
type BorrowRequestRepository interface {
	CreateBorrowRequest(ctx context.Context, request *BorrowRequest) error
	GetBorrowRequestByID(ctx context.Context, id primitive.ObjectID) (*BorrowRequest, error)
	GetAllBorrowRequests(ctx context.Context) ([]*BorrowRequest, error)
	GetBorrowRequestsByUserID(ctx context.Context, userID primitive.ObjectID) ([]*BorrowRequest, error)
	GetBorrowRequestsByBookID(ctx context.Context, bookID primitive.ObjectID) ([]*BorrowRequest, error)
	UpdateBorrowRequest(ctx context.Context, request *BorrowRequest) error
	DeleteBorrowRequest(ctx context.Context, id primitive.ObjectID) error
}
type BookUsecase interface {
	CreateBook(book *Book) interface{}
	GetBookByID(id string) interface{}
	GetAllBooks() interface{}
	GetAllAvailableBooks() interface{}
	UpdateBook(book *Book) interface{}
	DeleteBook(id string) interface{}
}
type BorrowRequestUsecase interface {
	CreateBorrowRequest(request *BorrowRequest) interface{}
	GetBorrowRequestByID(id string) interface{}
	GetAllBorrowRequests() interface{}
	GetBorrowRequestsByUserID(userID string) interface{}
	GetBorrowRequestsByBookID(bookID string) interface{}
	UpdateBorrowRequest(request *BorrowRequest) interface{}
	DeleteBorrowRequest(id string) interface{}
}
type BookController interface {
	CreateBook(c *gin.Context)
	GetBookByID(c *gin.Context)
	GetAllBooks(c *gin.Context)
	GetAllAvailableBooks(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}
type BorrowRequestController interface {
	CreateBorrowRequest(c *gin.Context)
	GetBorrowRequestByID(c *gin.Context)
	GetAllBorrowRequests(c *gin.Context)
	GetBorrowRequestsByUserID(c *gin.Context)
	GetBorrowRequestsByBookID(c *gin.Context)
	UpdateBorrowRequest(c *gin.Context)
	DeleteBorrowRequest(c *gin.Context)
}
type BorrowController interface {
	// Apply to borrow a book
	BorrowBook(c *gin.Context)

	// View the status of a specific borrowing request by the current user
	ViewBorrowingStatus(c *gin.Context)

	// View all borrowing requests (Admin)
	ViewAllBorrowingRequests(c *gin.Context)

	// Approve or reject a borrow request (Admin)
	UpdateBorrowingStatus(c *gin.Context)

	// Delete a borrow request (Admin)
	DeleteBorrowRequest(c *gin.Context)
}
