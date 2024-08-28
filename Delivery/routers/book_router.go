package routers

import (
	"LoanTrackerAPI/Delivery/controllers"
	infrastructure "LoanTrackerAPI/Infrastructure"
	repositories "LoanTrackerAPI/Repositories"
	usecases "LoanTrackerAPI/Usecases"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func BookRouter(R *gin.Engine, env infrastructure.Config, client mongo.Database) {
	bookRepository := repositories.NewBookRepository(client, env.Bookcoll)
	bookUsecase := usecases.NewBookUsecase(bookRepository, context.TODO())
	bookController := controllers.NewBookController(bookUsecase)

	bookroute := R.Group("/books")
	bookroute.POST("/", bookController.CreateBook)
	// {
	// 	bookroute.POST("/", bookController.CreateBook)
	// 	bookroute.GET("/", bookController.GetAllBooks)
	// 	bookroute.GET("/available", bookController.GetAllAvailableBooks)
	// 	bookroute.GET("/:id", bookController.GetBookByID)
	// 	bookroute.PUT("/:id", bookController.UpdateBook)
	// 	bookroute.DELETE("/:id", bookController.DeleteBook)
	// }
}
