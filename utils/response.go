package utils

import (
	"github.com/gin-gonic/gin"
)

// ResponsePayload defines the structure for API responses
type ResponsePayload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// SuccessResponse sends a success response with custom code, message, and optional data
func SuccessResponse(c *gin.Context, code int, message string, data any) {
	c.JSON(CodeSuccess, ResponsePayload{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response with custom code and message
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, ResponsePayload{
		Code:    code,
		Message: message,
	})
}

func BadRequestResponse(c *gin.Context, message string) {
		Logger.WithFields(map[string]any{
		"status":  CodeInternalServerError,
		"path":    c.Request.URL.Path,
		"method":  c.Request.Method,
		"message": message,
	}).Error(message)
	
	c.JSON(CodeBadRequest, ResponsePayload{
		Code:    CodeBadRequest,
		Message: "Bad Request: " + message,
	})
}

func InternalServerErrorResponse(c *gin.Context, code int, message string) {
	// Log ke terminal (logrus sudah diinisialisasi di InitLogger)
	Logger.WithFields(map[string]any{
		"status":  CodeInternalServerError,
		"path":    c.Request.URL.Path,
		"method":  c.Request.Method,
		"message": message,
	}).Error("internal server error")

	// Kirim response ke client
	c.JSON(CodeInternalServerError, ResponsePayload{
		Code:    CodeInternalServerError,
		Message: message,
	})
}
