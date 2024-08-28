package usecases

import (
	domain "LoanTrackerAPI/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BorrowRequestUsecase struct {
	BorrowRequestRepo domain.BorrowRequestRepository
	BookRepo          domain.BookRepository
	Ctx               context.Context
}

func NewBorrowRequestUsecase(repo domain.BorrowRequestRepository, ctx context.Context, bookrepo domain.BookRepository) domain.BorrowRequestUsecase {
	return &BorrowRequestUsecase{
		BorrowRequestRepo: repo,
		BookRepo:          bookrepo,
		Ctx:               ctx,
	}
}

// CreateBorrowRequest implements domain.BorrowRequestUsecase.
func (b *BorrowRequestUsecase) CreateBorrowRequest(request *domain.BorrowRequest) interface{} {
	// color.Green("Creating borrow request", request)
	book, err := b.BookRepo.GetBookByID(b.Ctx, request.BookID)
	if err != nil || !book.IsAvailable {
		return domain.ErrorResponse{Message: "No book found", Error: err.Error(), Status: 404}
	}
	request.ID = primitive.NewObjectID()
	request.Status = domain.StatusPending
	err = b.BorrowRequestRepo.CreateBorrowRequest(b.Ctx, request)
	if err != nil {
		return domain.ErrorResponse{Message: "Error creating borrow request", Error: err.Error(), Status: 500}
	}

	return domain.SuccessResponse{Message: "Borrow request created successfully", Status: 200, Data: request}
}

// DeleteBorrowRequest implements domain.BorrowRequestUsecase.
func (b *BorrowRequestUsecase) DeleteBorrowRequest(id string) interface{} {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrorResponse{Message: "Invalid ID", Error: err.Error(), Status: 400}
	}
	err = b.BorrowRequestRepo.DeleteBorrowRequest(b.Ctx, idObj)

	if err != nil {
		return domain.ErrorResponse{Message: "Error deleting borrow request", Error: err.Error(), Status: 500}
	}

	return domain.SuccessResponse{Message: "Borrow request deleted successfully", Status: 200}

}

// GetAllBorrowRequests implements domain.BorrowRequestUsecase.
func (b *BorrowRequestUsecase) GetAllBorrowRequests() interface{} {
	requests, err := b.BorrowRequestRepo.GetAllBorrowRequests(b.Ctx)
	if err != nil {
		return domain.ErrorResponse{Message: "Error fetching borrow requests", Error: err.Error(), Status: 500}
	}

	return domain.SuccessResponse{Message: "Borrow requests fetched successfully", Status: 200, Data: requests}

}

// GetBorrowRequestByID implements domain.BorrowRequestUsecase.
func (b *BorrowRequestUsecase) GetBorrowRequestByID(id string) interface{} {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.ErrorResponse{Message: "Invalid ID", Error: err.Error(), Status: 400}
	}
	request, err := b.BorrowRequestRepo.GetBorrowRequestByID(b.Ctx, idObj)
	if err != nil {
		return domain.ErrorResponse{Message: "Error fetching borrow request", Error: err.Error(), Status: 500}
	}

	return domain.SuccessResponse{Message: "Borrow request fetched successfully", Status: 200, Data: request}
}

// GetBorrowRequestsByBookID implements domain.BorrowRequestUsecase.
func (b *BorrowRequestUsecase) GetBorrowRequestsByBookID(bookID string) interface{} {
	bookIDObj, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return domain.ErrorResponse{Message: "Invalid ID", Error: err.Error(), Status: 400}
	}
	requests, err := b.BorrowRequestRepo.GetBorrowRequestsByBookID(b.Ctx, bookIDObj)
	if err != nil {
		return domain.ErrorResponse{Message: "Error fetching borrow requests", Error: err.Error(), Status: 500}
	}

	return domain.SuccessResponse{Message: "Borrow requests fetched successfully", Status: 200, Data: requests}
}

// GetBorrowRequestsByUserID implements domain.BorrowRequestUsecase.
func (b *BorrowRequestUsecase) GetBorrowRequestsByUserID(userID string) interface{} {
	userIDObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return domain.ErrorResponse{Message: "Invalid ID", Error: err.Error(), Status: 400}
	}
	requests, err := b.BorrowRequestRepo.GetBorrowRequestsByUserID(b.Ctx, userIDObj)
	if err != nil {
		return domain.ErrorResponse{Message: "Error fetching borrow requests", Error: err.Error(), Status: 500}
	}

	return domain.SuccessResponse{Message: "Borrow requests fetched successfully", Status: 200, Data: requests}
}

// UpdateBorrowRequest implements domain.BorrowRequestUsecase.
func (b *BorrowRequestUsecase) UpdateBorrowRequest(request *domain.BorrowRequest) interface{} {
	borrowBook, err := b.BorrowRequestRepo.GetBorrowRequestByID(b.Ctx, request.ID)
	if err != nil {
		return domain.ErrorResponse{Message: "Error updating borrow request", Error: err.Error(), Status: 500}
	}
	borrowBook.Status = request.Status

	err = b.BorrowRequestRepo.UpdateBorrowRequest(b.Ctx, borrowBook)
	if request.Status == domain.StatusApproved {
		book, err := b.BookRepo.GetBookByID(b.Ctx, request.BookID)
		if err != nil {
			return domain.ErrorResponse{Message: "Error updating borrow request", Error: err.Error(), Status: 500}
		}
		book.IsAvailable = false
		err = b.BookRepo.UpdateBook(b.Ctx, book)
		if err != nil {
			return domain.ErrorResponse{Message: "Error updating borrow request", Error: err.Error(), Status: 500}
		}
	}
	if err != nil {
		return domain.ErrorResponse{Message: "Error updating borrow request", Error: err.Error(), Status: 500}
	}

	return domain.SuccessResponse{Message: "Borrow request updated successfully", Status: 200, Data: request}
}
