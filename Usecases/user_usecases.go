package usecases

import (
	domain "LoanTrackerAPI/Domain"
	infrastructure "LoanTrackerAPI/Infrastructure"
	"context"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	userRepository domain.UserRepository
	Env            infrastructure.Config
	contextTimeout time.Duration
}

// IsAdmin implements domain.UserUsecase.
func (u *userUsecase) IsAdmin(ctx context.Context, userID string) bool {
	user, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return false
	}
	if user.Role == "admin" {
		return true
	}
	return false
}

// DeleteUserAccount implements domain.UserUsecase.
func (u *userUsecase) DeleteUserAccount(ctx context.Context, userID string, adminId string) interface{} {
	admin, err := u.userRepository.GetUserByID(ctx, adminId)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error fetching user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	if admin.Role != "admin" {
		return domain.ErrorResponse{
			Message: "Unauthorized",
			Error:   "User is not an admin",
			Status:  401,
		}
	}
	err = u.userRepository.DeleteUser(ctx, userID)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error deleting user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	return domain.SuccessResponse{
		Message: "User deleted successfully",
		Status:  200,
	}
}

// GetAllUsers implements domain.UserUsecase.
func (u *userUsecase) GetAllUsers(ctx context.Context, id string) interface{} {
	admin, err := u.userRepository.GetUserByID(ctx, id)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error fetching user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	if admin.Role != "admin" {
		return domain.ErrorResponse{
			Message: "Unauthorized",
			Error:   "User is not an admin",
			Status:  401,
		}
	}

	users, err := u.userRepository.GetAllUsers(ctx)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error fetching users",
			Error:   err.Error(),
			Status:  500,
		}
	}
	return domain.SuccessResponse{
		Message: "Users fetched successfully",
		Status:  200,
		Data:    users,
	}
}

// GetUserProfile implements domain.UserUsecase.
func (u *userUsecase) GetUserProfile(ctx context.Context, userID string) interface{} {
	user, err := u.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error fetching user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	var profile domain.UserProfile
	copier.Copy(&profile, user)
	return domain.SuccessResponse{
		Message: "User profile fetched successfully",
		Status:  200,
		Data:    profile,
	}
}

// Login implements domain.UserUsecase.
func (u *userUsecase) Login(ctx context.Context, request domain.LoginRequest) interface{} {
	user, err := u.userRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error fetching user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	if !infrastructure.ComparePassword(request.Password, user.Password) {
		return domain.ErrorResponse{
			Message: "Invalid credentials",
			Error:   "Invalid password",
			Status:  401,
		}
	}
	if !user.IsVerified {
		return domain.ErrorResponse{
			Message: "Unverified email",
			Error:   "Email not verified",
			Status:  401,
		}
	}
	accessclaims := domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)), // Convert expiration time to *jwt.NumericDate
		},
	}
	refrechclaims := domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Convert expiration time to *jwt.NumericDate
		},
	}

	accesstoken, err := infrastructure.CreateToken(accessclaims, u.Env.Jwt_secret)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error generating token",
			Error:   err.Error(),
			Status:  500,
		}
	}
	refrechtoken, err := infrastructure.CreateToken(refrechclaims, u.Env.Jwt_secret)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error generating token",
			Error:   err.Error(),
			Status:  500,
		}
	}
	user.RefreshToken = refrechtoken
	err = u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error updating user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	return domain.SuccessResponse{
		Message: "Login successful",
		Status:  200,
		Data:    domain.LoginResponse{AccessToken: accesstoken, RefreshToken: refrechtoken},
	}
}

// RefreshToken implements domain.UserUsecase.
func (u *userUsecase) RefreshToken(ctx context.Context, refreshToken string) interface{} {
	claim, err := infrastructure.ExtractClaim(refreshToken, u.Env.Jwt_secret, &domain.JwtCustomClaims{})
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error extracting claim",
			Error:   err.Error(),
			Status:  500,
		}
	}
	claims := claim.(*domain.JwtCustomClaims)

	user, err := u.userRepository.GetUserByID(ctx, claims.ID)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error fetching user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	if user.RefreshToken != refreshToken {
		return domain.ErrorResponse{
			Message: "Invalid token",
			Error:   "Invalid refresh token",
			Status:  401,
		}
	}
	verify, err := infrastructure.VerifyToken(refreshToken, u.Env.Jwt_secret)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error verifying token",
			Error:   err.Error(),
			Status:  500,
		}
	}
	if !verify {
		return domain.ErrorResponse{
			Message: "Invalid token",
			Error:   "Invalid refresh token",
			Status:  401,
		}
	}

	accessclaims := domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)), // Convert expiration time to *jwt.NumericDate
		},
	}
	accessTocken, err := infrastructure.CreateToken(accessclaims, u.Env.Jwt_secret)

	if err != nil {
		return domain.ErrorResponse{
			Message: "Error generating token",
			Error:   err.Error(),
			Status:  500,
		}
	}
	return domain.SuccessResponse{
		Message: "Token refreshed successfully",
		Status:  200,
		Data:    domain.RefreshTokenResponse{AccessToken: accessTocken},
	}

}

// RegisterUser implements domain.UserUsecase.
func (u *userUsecase) RegisterUser(ctx context.Context, user *domain.User) interface{} {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Invalid input",
			Error:   err.Error(),
			Status:  400,
		}

	}
	err = infrastructure.ValidateEmail(user.Email)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Invalid email",
			Error:   err.Error(),
			Status:  400,
		}
	}
	expirationTime := time.Now().Add(time.Hour * 24)
	claim := domain.EmailVerificationClaims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Convert expiration time to *jwt.NumericDate
		},
	}
	// fmt.Println("env", u.Env.Jwt_secret)
	token, err := infrastructure.CreateToken(claim, u.Env.Jwt_secret)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error generating token",
			Error:   err.Error(),
			Status:  500,
		}
	}
	user.VerifyToken = token
	user.IsVerified = false
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "user"
	user.Password, err = infrastructure.HashPassword(user.Password)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error hashing password",
			Error:   err.Error(),
			Status:  500,
		}
	}
	user.ID = primitive.NewObjectID()
	err = infrastructure.SendVerificationEmail(user.Email, token)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error sending verification email",
			Error:   err.Error(),
			Status:  500,
		}
	}
	err = u.userRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error creating user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	var profile domain.UserProfile
	copier.Copy(&profile, user)
	return domain.SuccessResponse{
		Message: "User created successfully",
		Status:  201,
		Data:    profile,
	}

}

// RequestPasswordReset implements domain.UserUsecase.
func (u *userUsecase) RequestPasswordReset(ctx context.Context, email string) interface{} {
	user, err := u.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error fetching user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	expirationTime := time.Now().Add(time.Hour * 24)
	claim := domain.EmailVerificationClaims{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Convert expiration time to *jwt.NumericDate
		},
	}
	token, err := infrastructure.CreateToken(claim, u.Env.Jwt_secret)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error generating token",
			Error:   err.Error(),
			Status:  500,
		}
	}
	user.ResetToken = token
	err = u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error updating user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	err = infrastructure.SendResetEmail(user.Email, token)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error sending password reset email",
			Error:   err.Error(),
			Status:  500,
		}
	}
	return domain.SuccessResponse{
		Message: "Password reset email sent successfully",
		Status:  200,
	}
}

// UpdatePasswordAfterReset implements domain.UserUsecase.
func (u *userUsecase) UpdatePasswordAfterReset(ctx context.Context, token string, newPassword string) interface{} {
	claim, err := infrastructure.ExtractClaim(token, u.Env.Jwt_secret, &domain.EmailVerificationClaims{})
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error extracting claim",
			Error:   err.Error(),
			Status:  500,
		}
	}
	claims := claim.(*domain.EmailVerificationClaims)
	user, err := u.userRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error fetching user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	user.Password, err = infrastructure.HashPassword(newPassword)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error hashing password",
			Error:   err.Error(),
			Status:  500,
		}
	}
	user.ResetToken = ""
	err = u.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error updating user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	return domain.SuccessResponse{
		Message: "Password updated successfully",
		Status:  200,
	}
}

// VerifyEmail implements domain.UserUsecase.
func (u *userUsecase) VerifyEmail(ctx context.Context, token string, email string) interface{} {
	err := u.userRepository.VerifyEmail(ctx, token, email)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error verifying email",
			Error:   err.Error(),
			Status:  500,
		}
	}
	user, err := u.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.ErrorResponse{
			Message: "Error fetching user",
			Error:   err.Error(),
			Status:  500,
		}
	}
	var profile domain.UserProfile
	copier.Copy(&profile, user)
	return domain.SuccessResponse{
		Message: "User verified successfully",
		Status:  200,
		Data:    profile,
	}

}

func NewUserUseCase(UserRepository domain.UserRepository, timeout time.Duration, env infrastructure.Config) domain.UserUsecase {
	return &userUsecase{
		userRepository: UserRepository,
		contextTimeout: timeout,
		Env:            env,
	}
}
