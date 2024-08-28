package controllers

import (
	domain "LoanTrackerAPI/Domain"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BorrowController struct {
	borrowUsecase domain.BorrowRequestUsecase
	userUsecase   domain.UserUsecase
	ctx           context.Context
}

func NewBorrowController(b domain.BorrowRequestUsecase, ctx context.Context, userusecase domain.UserUsecase) domain.BorrowController {
	return &BorrowController{
		borrowUsecase: b,
		ctx:           ctx,
		userUsecase:   userusecase,
	}
}

// BorrowBook implements domain.BorrowController.
func (b *BorrowController) BorrowBook(c *gin.Context) {
	var borrowRequest *domain.BorrowBookRequest
	err := c.ShouldBind(&borrowRequest)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid request", Status: 400, Error: err.Error()})
		return
	}
	user_id := c.GetString("user_id")
	if user_id == "" {
		HandleResponse(c, domain.ErrorResponse{Message: "User not found", Status: 404})
		return
	}
	user_idObj, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid ID", Status: 400, Error: err.Error()})
		return
	}
	bookIdObj, err := primitive.ObjectIDFromHex(borrowRequest.BookID)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid ID", Status: 400, Error: err.Error()})
		return
	}
	NewRequest:= &domain.BorrowRequest{
		BookID: bookIdObj,
		UserID: user_idObj,
	}
	response := b.borrowUsecase.CreateBorrowRequest(NewRequest)
	HandleResponse(c, response)
}

// DeleteBorrowRequest implements domain.BorrowController.
func (b *BorrowController) DeleteBorrowRequest(c *gin.Context) {
	user_id := c.GetString("user_id")
	if user_id == "" {
		HandleResponse(c, domain.ErrorResponse{Message: "User not found", Status: 404})
		return
	}
	//check if the user is admin
	admin := b.userUsecase.IsAdmin(b.ctx, user_id)
	if !admin {
		HandleResponse(c, domain.ErrorResponse{Message: "Unauthorized", Status: 401})
		return
	}
	id := c.Param("id")
	response := b.borrowUsecase.DeleteBorrowRequest(id)
	HandleResponse(c, response)

}

// UpdateBorrowingStatus implements domain.BorrowController.
func (b *BorrowController) UpdateBorrowingStatus(c *gin.Context) {

	user_id := c.GetString("user_id")
	if user_id == "" {
		HandleResponse(c, domain.ErrorResponse{Message: "User not found", Status: 404})
		return
	}
	//check if the user is admin
	admin := b.userUsecase.IsAdmin(b.ctx, user_id)
	if !admin {
		HandleResponse(c, domain.ErrorResponse{Message: "Unauthorized", Status: 401})
		return
	}

	borrowRequest := domain.BorrowRequest{}
	var request *domain.UpdateBorrowRequest
	err := c.ShouldBind(&request)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid request", Status: 400, Error: err.Error()})
		return
	}
	id := c.Param("id")
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid ID", Status: 400, Error: err.Error()})
		return
	}
	borrowRequest.Status = request.Status
	borrowRequest.ID = idObj

	response := b.borrowUsecase.UpdateBorrowRequest(&borrowRequest)
	HandleResponse(c, response)

}

// ViewAllBorrowingRequests implements domain.BorrowController.
func (b *BorrowController) ViewAllBorrowingRequests(c *gin.Context) {
	user_id := c.GetString("user_id")
	if user_id == "" {
		HandleResponse(c, domain.ErrorResponse{Message: "User not found", Status: 404})
		return
	}
	//check if the user is admin
	admin := b.userUsecase.IsAdmin(b.ctx, user_id)
	if !admin {
		HandleResponse(c, domain.ErrorResponse{Message: "Unauthorized", Status: 401})
		return
	}
	response := b.borrowUsecase.GetAllBorrowRequests()
	HandleResponse(c, response)
}

// ViewBorrowingStatus implements domain.BorrowController.
func (b *BorrowController) ViewBorrowingStatus(c *gin.Context) {
	id := c.Param("id")

	response := b.borrowUsecase.GetBorrowRequestByID(id)
	HandleResponse(c, response)
}
