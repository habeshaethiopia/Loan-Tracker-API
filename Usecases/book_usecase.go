package usecases

import (
	domain "LoanTrackerAPI/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type bookUsecase struct {
	BookRepo domain.BookRepository
	ctx      context.Context
}

// CreateBook implements domain.BookUsecase.
func (b *bookUsecase) CreateBook(book *domain.Book) interface{} {
	book.ID = primitive.NewObjectID()
	book.IsAvailable = true
	err := b.BookRepo.CreateBook(b.ctx, book)
	if err != nil {
		return domain.ErrorResponse{Message: "Error creating book", Error: err.Error(), Status: 500}
	}
	return domain.SuccessResponse{Message: "Book created successfully", Status: 200, Data: book}
}

// DeleteBook implements domain.BookUsecase.
func (b *bookUsecase) DeleteBook(id string) interface{} {
	panic("unimplemented")
}

// GetAllAvailableBooks implements domain.BookUsecase.
func (b *bookUsecase) GetAllAvailableBooks() interface{} {
	panic("unimplemented")
}

// GetAllBooks implements domain.BookUsecase.
func (b *bookUsecase) GetAllBooks() interface{} {
	panic("unimplemented")
}

// GetBookByID implements domain.BookUsecase.
func (b *bookUsecase) GetBookByID(id string) interface{} {
	panic("unimplemented")
}

// UpdateBook implements domain.BookUsecase.
func (b *bookUsecase) UpdateBook(book *domain.Book) interface{} {
	panic("unimplemented")
}

func NewBookUsecase(repo domain.BookRepository, ctx context.Context) domain.BookUsecase {
	return &bookUsecase{
		BookRepo: repo,
		ctx:      ctx,
	}
}
