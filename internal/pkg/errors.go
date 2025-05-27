package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserOrGameNotFound  = errors.New("user or game not found")
	ErrInvalidUserID       = errors.New("user id provided is invalid")
	ErrInternalServerError = errors.New("an unexpected error occurred. Please try again later")
)

const (
	BadRequest          = "BAD_REQUEST"
	UserNotFound        = "USER_NOT_FOUND"
	UserOrGameNotFound  = "USER_OR_GAME_NOT_FOUND"
	InternalServerError = "INTERNAL_SERVER_ERROR"
)

var (
	ErrorMapping = map[error]string{
		ErrUserNotFound:        UserNotFound,
		ErrUserOrGameNotFound:  UserOrGameNotFound,
		ErrInternalServerError: InternalServerError,
		ErrInvalidUserID:       BadRequest,
	}
)

// WriteErrorResponseToResponseWriter writes a structured JSON error response to the HTTP response writer.
//
// This function sets the response status code, content type, and writes a JSON-encoded
// ErrorResponse based on the provided error and its corresponding error code from ErrorMapping.
//
// Parameters:
//   - w: The http.ResponseWriter to write the response to.
//   - statusCode: The HTTP status code to return (e.g., 400, 404, 500).
//   - err: The error to convert into an ErrorResponse. It must be a known key in ErrorMapping.
//
// Returns:
//   - None. The function writes directly to the HTTP response.
//
// Example:
//
//	pkg.WriteErrorResponseToResponseWriter(w, http.StatusNotFound, pkg.ErrUserNotFound)
func WriteErrorResponseToResponseWriter(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(GetErrorResponseBody(ErrorMapping[err], err.Error()))
}

// GetErrorResponseBody constructs a new ErrorResponse with the specified code and message.
//
// Parameters:
//   - errorCtx: A string identifier for the error context (e.g., "USER_NOT_FOUND").
//   - message: A human-readable error message.
//
// Returns:
//   - A pointer to an ErrorResponse struct populated with the given values.
//
// Example:
//
//	errResp := GetErrorResponseBody(pkg.UserNotFound, "User with ID 42 does not exist")
func GetErrorResponseBody(errorCtx, message string) *ErrorResponse {
	return &ErrorResponse{
		Error:   errorCtx,
		Message: message,
	}
}
