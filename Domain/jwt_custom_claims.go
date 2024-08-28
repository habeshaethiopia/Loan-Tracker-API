package domain

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type EmailVerificationClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type Clams struct {
}

func (c JwtCustomClaims) Valid() error {
	// You can add custom validation logic here if needed
	return nil
}

// Valid method for UserClaims (Optional if using jwt.RegisteredClaims)
func (c EmailVerificationClaims) Valid() error {
	// You can add custom validation logic here if needed
	return nil
}
