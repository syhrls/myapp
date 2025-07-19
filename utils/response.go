package utils

import (
	"github.com/gin-gonic/gin"
)

// ResponsePayload defines the structure for API responses
type ResponsePayload struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// SuccessResponse sends a success response with custom code, message, and optional data
func SuccessResponse(c *gin.Context, message string, data any) {
	c.JSON(CodeSuccess, ResponsePayload{
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response with custom code and message
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, ResponsePayload{
		Message: message,
	})
}

func BadRequestResponse(c *gin.Context, message string) {
	c.JSON(CodeBadRequest, ResponsePayload{
		Message: "Bad Request: " + message,
	})
}

func InternalServerErrorResponse(c *gin.Context) {
	c.JSON(CodeInternalServerError, ResponsePayload{
		Message: "Internal Server Error",
	})
}
