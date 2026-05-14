package apperror

import "fmt"

type ErrorType string

const (
	TypeValidation ErrorType = "VALIDATION"
	TypeNotFound   ErrorType = "NOT_FOUND"
	TypeConflict   ErrorType = "CONFLICT"
	TypeInternal   ErrorType = "INTERNAL"
)

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Type, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Type:    TypeValidation,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Type:    TypeNotFound,
		Message: message,
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Type:    TypeConflict,
		Message: message,
	}
}

func NewInternalError(err error) *AppError {
	return &AppError{
		Type:    TypeInternal,
		Message: "Internal server error",
		Err:     err,
	}
}
