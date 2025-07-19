package utils

const (
	// Success response code
	CodeSuccess = 200

	// Client error codes
	CodeBadRequest          = 400
	CodeUnauthorized        = 401
	CodeForbidden           = 403
	CodeNotFound            = 404
	CodeMethodNotAllowed    = 405

	// Server error codes
	CodeInternalServerError = 500
	CodeBadGateway          = 502
	CodeServiceUnavailable  = 503
)