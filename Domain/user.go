package domain

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	FullName        string    `json:"full_name" bson:"full_name"`
	Email           string    `json:"email" bson:"email"`
	Password        string    `json:"password" bson:"password"`
	Role            string    `json:"role" bson:"role"`
	ProfileImageURL string    `json:"profile_image_url,omitempty" bson:"profile_image_url,omitempty"`
	IsVerified      bool      `json:"is_verified" bson:"is_verified"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" bson:"updated_at"`
	RefreshToken    string    `json:"refresh_token" bson:"refresh_token"`
	ResetToken      string    `json:"reset_token" bson:"reset_token"`
	VerifyToken     string    `json:"verify_token" bson:"verify_token"`
}
type UserProfile struct {
	ID              string    `json:"id" bson:"_id,omitempty"`
	FullName        string    `json:"full_name" bson:"full_name"`
	Email           string    `json:"email" bson:"email"`
	Role            string    `json:"role" bson:"role"`
	ProfileImageURL string    `json:"profile_image_url,omitempty" bson:"profile_image_url,omitempty"`
	IsVerified      bool      `json:"is_verified" bson:"is_verified"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" bson:"updated_at"`
}

// UserController defines the interface for user controller operations
type UserController interface {
	RegisterUser(c *gin.Context)
	VerifyEmail(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	GetUserProfile(c *gin.Context)
	RequestPasswordReset(c *gin.Context)
	UpdatePasswordAfterReset(c *gin.Context)
	GetAllUsers(c *gin.Context)
	DeleteUserAccount(c *gin.Context)
}
type UserUsecase interface {
	RegisterUser(ctx context.Context, user *User) interface{}
	VerifyEmail(ctx context.Context, token, email string) interface{}
	Login(ctx context.Context, request LoginRequest) interface{}
	RefreshToken(ctx context.Context, refreshToken string) interface{}
	GetUserProfile(ctx context.Context, userID string) interface{}
	RequestPasswordReset(ctx context.Context, email string) interface{}
	UpdatePasswordAfterReset(ctx context.Context, token, newPassword string) interface{}
	GetAllUsers(ctx context.Context, id string) interface{}
	DeleteUserAccount(ctx context.Context, userID string, adminId string) interface{}
	IsAdmin(ctx context.Context, userID string) bool
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	VerifyEmail(ctx context.Context, token, email string) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, userID string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, userID string) error
	GetAllUsers(ctx context.Context) ([]*User, error)
}
