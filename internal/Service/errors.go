package service

import "fmt"

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
