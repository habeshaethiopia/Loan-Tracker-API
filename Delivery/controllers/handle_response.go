package controllers

import (
	domain "LoanTrackerAPI/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleResponse(c *gin.Context, response interface{}) {

	switch res := response.(type) {
	case domain.SuccessResponse:
		c.JSON(http.StatusOK, res)
	case domain.ErrorResponse:
		c.JSON(res.Status, res)
	default:
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Internal Server Error unknown response", Status: 500})
	}
}
