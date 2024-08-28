package infrastructure

import (
	domain "LoanTrackerAPI/Domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header required",
			})
			c.Abort()
			return
		}
		auth := strings.Split(authHeader, " ")
		if len(auth) != 2 || strings.ToLower(auth[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authorization header.",
			})
			c.Abort()
			return
		}
		claims, err := ExtractClaim(auth[1], secret, &domain.JwtCustomClaims{})
		if err != nil {
			c.JSON(500, domain.ErrorResponse{
				Message: "Error extracting claim",
				Error:   err.Error(),
				Status:  500,
			})
			c.Abort()
			return
		}
		if jwtClaims, ok := claims.(*domain.JwtCustomClaims); ok {
			c.Set("user_id", jwtClaims.ID)
		} else {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
				Message: "Invalid token claims",
				Error:   "Failed to assert token claims",
				Status:  http.StatusInternalServerError,
			})
			c.Abort()
			return
		}
		c.Next()
		// Can check the expiration time of the token if it is valid or not
	}
}
