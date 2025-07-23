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
	Success(message)
	c.JSON(CodeSuccess, ResponsePayload{
		Message: SUCCESS,
		Data:    data,
	})
}

// ErrorResponse sends an error response with custom code and message
func ErrorResponse(c *gin.Context, code int, message string) {
	Error(message)
	c.JSON(code, ResponsePayload{
		Message: message,
	})
}

func BadRequestResponse(c *gin.Context, message string) {
	Error(message)
	c.JSON(CodeBadRequest, ResponsePayload{
		Message: message,
	})
}

func InternalServerErrorResponse(c *gin.Context) {
	Error("Internal Server Error")
	c.JSON(CodeInternalServerError, ResponsePayload{
		Message: "Internal Server Error",
	})
}
