package routers

import (
	"LoanTrackerAPI/Delivery/controllers"
	infrastructure "LoanTrackerAPI/Infrastructure"
	repositories "LoanTrackerAPI/Repositories"
	usecases "LoanTrackerAPI/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRouter(user *gin.RouterGroup, client mongo.Database, config infrastructure.Config) {

	userRepository := repositories.NewUserRepository(client, config.Usercoll)
	userUsecase := usecases.NewUserUseCase(userRepository, time.Duration(10)*time.Minute, config)
	userController := controllers.UserController{
		UserUsecase: userUsecase,
		Env:         config,
	}
	R := user.Group("/users")
	secure := user.Group("")
	secure.Use(infrastructure.AuthenticationMiddleware(config.Jwt_secret))
	// User Registration
	R.POST("/register", userController.RegisterUser)

	// Email Verification
	R.GET("/verify-email", userController.VerifyEmail)

	// User Login
	R.POST("/login", userController.Login)

	// Token Refresh
	R.POST("/token/refresh", userController.RefreshToken)

	// User Profile
	secure.GET("users/profile", userController.GetUserProfile)

	// Password Reset Request
	R.POST("/password-reset", userController.RequestPasswordReset)

	// Password Update After Reset
	R.POST("/password-update", userController.UpdatePasswordAfterReset)
	admin := secure.Group("admin/users")
	// View All Users (Admin)
	admin.GET("/", userController.GetAllUsers)

	// Delete User Account (Admin)
	admin.DELETE("/:id", userController.DeleteUserAccount)
}
