package routers

import (
	"LoanTrackerAPI/Delivery/controllers"
	infrastructure "LoanTrackerAPI/Infrastructure"
	repositories "LoanTrackerAPI/Repositories"
	usecases "LoanTrackerAPI/Usecases"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func BorrowRouter(user *gin.RouterGroup, client mongo.Database, config infrastructure.Config) {

	ctx := context.Background()

	// Initialize repositories
	BorrowRepositories := repositories.NewBookBorrowRepository(client, config.Borrowcoll)
	userRepository := repositories.NewUserRepository(client, config.Usercoll)
	bookrepo := repositories.NewBookRepository(client, config.Bookcoll)

	// Initialize use cases
	BorrowUsecase := usecases.NewBorrowRequestUsecase(BorrowRepositories, ctx, bookrepo)
	userUsecase := usecases.NewUserUseCase(userRepository, time.Duration(10)*time.Minute, config)

	// Initialize controller
	// type BorrowController struct {
	// 	borrowUsecase domain.BorrowRequestUsecase
	// 	userUsecase   domain.UserUsecase
	// 	ctx           context.Context
	// }
	BorrowController := controllers.NewBorrowController(BorrowUsecase, ctx, userUsecase)

	// Unsecured routes (no authentication required)
	// Route to borrow a book

	// Secured routes (authentication required)
	secure := user.Group("/borrow")
	secure.Use(infrastructure.AuthenticationMiddleware(config.Jwt_secret))
	secure.POST("/", BorrowController.BorrowBook)

	// Route to view the status of a specific borrowing request
	secure.GET("/:id", BorrowController.ViewBorrowingStatus)

	// Admin routes (assuming additional middleware for admin is applied)
	admin := user.Group("/admin/borrows")
	admin.Use(infrastructure.AuthenticationMiddleware(config.Jwt_secret)) // This could include an admin-specific middleware
	// Route to view all borrowing requests
	admin.GET("/", BorrowController.ViewAllBorrowingRequests)
	// Route to approve/reject a borrowing request
	admin.PATCH("/:id/status", BorrowController.UpdateBorrowingStatus)
	// Route to delete a borrowing request
	admin.DELETE("/:id", BorrowController.DeleteBorrowRequest)
}
