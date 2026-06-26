package service

import "fmt"

const (
	NoTitle   = "task title not specified"
	NoID      = "ID not specified"
	InvalidID = "invalid ID"
	NoTask    = "task not found"
)

type AppError struct {
	StatusCode int
	Message    string
}

func NewError(code int, message string) error {
	return &AppError{
		StatusCode: code,
		Message:    message,
	}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("status %d: %s", e.StatusCode, e.Message)
}
