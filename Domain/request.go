package domain

// login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// register request
type RegisterRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// pasword reset request
type PasswordResetRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// password reset confirm request
type PasswordResetConfirmRequest struct {
	// Token    string `json:"token" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// request refresh token
type TokenRefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
//write refreshtofken response and request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}