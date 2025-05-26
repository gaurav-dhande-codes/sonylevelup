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

func WriteErrorResponseToResponseWriter(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GetErrorResponseBody(ErrorMapping[err], err.Error()))
}

func GetErrorResponseBody(errorCtx, message string) *ErrorResponse {
	return &ErrorResponse{
		Error:   errorCtx,
		Message: message,
	}
}
