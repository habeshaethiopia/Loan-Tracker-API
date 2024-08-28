package controllers

import (
	domain "LoanTrackerAPI/Domain"
	infrastructure "LoanTrackerAPI/Infrastructure"

	"net/http"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserController struct {
	UserUsecase domain.UserUsecase
	Env         infrastructure.Config
}

// DeleteUserAccount implements domain.UserController.
func (u *UserController) DeleteUserAccount(c *gin.Context) {
	user_id := c.GetString("user_id")
	if user_id == "" {
		HandleResponse(c, domain.ErrorResponse{Message: "User not found", Status: 404})
		return
	}
	id := c.Param("id")
	response := u.UserUsecase.DeleteUserAccount(c, id, user_id)
	HandleResponse(c, response)
}

// GetAllUsers implements domain.UserController.
func (u *UserController) GetAllUsers(c *gin.Context) {
	user_id := c.GetString("user_id")
	if user_id == "" {
		HandleResponse(c, domain.ErrorResponse{Message: "User not found", Status: 404})
		return
	}
	users := u.UserUsecase.GetAllUsers(c, user_id)
	HandleResponse(c, users)

}

// GetUserProfile implements domain.UserController.
func (u *UserController) GetUserProfile(c *gin.Context) {
	user_id := c.GetString("user_id")
	if user_id == "" {
		HandleResponse(c, domain.ErrorResponse{Message: "User not found", Status: 404})
		return
	}
	profile := u.UserUsecase.GetUserProfile(c, user_id)
	HandleResponse(c, profile)
}

// Login implements domain.UserController.
func (u *UserController) Login(c *gin.Context) {
	var loginRequest domain.LoginRequest
	err := c.BindJSON(&loginRequest)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid request", Status: 400, Error: err.Error()})
		return
	}
	response := u.UserUsecase.Login(c, loginRequest)

	HandleResponse(c, response)

}

// RefreshToken implements domain.UserController.
func (u *UserController) RefreshToken(c *gin.Context) {
	var refreshToken domain.RefreshTokenRequest
	err := c.BindJSON(&refreshToken)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid request", Status: 400, Error: err.Error()})
		return
	}
	response := u.UserUsecase.RefreshToken(c, refreshToken.RefreshToken)

	HandleResponse(c, response)
}

// RegisterUser implements domain.UserController.
func (u *UserController) RegisterUser(c *gin.Context) {
	var user domain.User
	var signupRequest domain.RegisterRequest
	err := c.BindJSON(&signupRequest)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = copier.Copy(&user, &signupRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to map fields"})
		return
	}

	response := u.UserUsecase.RegisterUser(c, &user)
	if response != nil {

		HandleResponse(c, response)

	}
}

// RequestPasswordReset implements domain.UserController.
func (u *UserController) RequestPasswordReset(c *gin.Context) {
	var resetRequest domain.PasswordResetRequest
	err := c.BindJSON(&resetRequest)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid request", Status: 400, Error: err.Error()})
		return
	}
	response := u.UserUsecase.RequestPasswordReset(c, resetRequest.Email)
	HandleResponse(c, response)

}

// UpdatePasswordAfterReset implements domain.UserController.
func (u *UserController) UpdatePasswordAfterReset(c *gin.Context) {
	var resetRequest domain.PasswordResetConfirmRequest
	err := c.BindJSON(&resetRequest)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid request", Status: 400, Error: err.Error()})
		return
	}
	token := c.Query("token")
	if token == "" {
		HandleResponse(c, domain.ErrorResponse{Message: "Token is required", Status: 400})
		return
	}
	response := u.UserUsecase.UpdatePasswordAfterReset(c, token, resetRequest.Password)
	HandleResponse(c, response)
}

// VerifyEmail implements domain.UserController.
func (u *UserController) VerifyEmail(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		HandleResponse(c, domain.ErrorResponse{Message: "Token is required", Status: 400})
		return
	}
	claims := &domain.EmailVerificationClaims{}
	res, err := infrastructure.ExtractClaim(tokenString, u.Env.Jwt_secret, claims)
	if err != nil {
		HandleResponse(c, domain.ErrorResponse{Message: "Invalid token123", Status: 400, Error: err.Error()})
		return
	}
	claims = res.(*domain.EmailVerificationClaims)
	color.Red(claims.Email)

	response := u.UserUsecase.VerifyEmail(c, tokenString, claims.Email)
	if response != nil {
		HandleResponse(c, response)
	}

}
