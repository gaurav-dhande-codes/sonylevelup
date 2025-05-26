package pkg

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserOrGameNotFound = errors.New("user or game not found")
)
