package helpers

import "net/http"

type RequestError interface {
	Error() string
	GetType() string
	GetStatusCode() int
}

// Base error type
type BaseError struct {
	Type       string `json:"type"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (err BaseError) Error() string {
	return err.Message
}

func (err BaseError) GetType() string {
	return err.Type
}

func (err BaseError) GetStatusCode() int {
	return err.StatusCode
}

// Specific error types using the base error as an embedding
type ValidationError struct {
	BaseError
}

func NewValidationError(message string) ValidationError {
	return ValidationError{
		BaseError: BaseError{
			Type:       "validation_error",
			StatusCode: http.StatusBadRequest,
			Message:    message,
		},
	}
}

type InternalServerError struct {
	BaseError
}

func NewInternalServerError(message string) InternalServerError {
	return InternalServerError{
		BaseError: BaseError{
			Type:       "internal_server_error",
			StatusCode: http.StatusInternalServerError,
			Message:    message,
		},
	}
}

type NotFoundError struct {
	BaseError
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		BaseError: BaseError{
			Type:       "not_found",
			StatusCode: http.StatusNotFound,
			Message:    message,
		},
	}
}

type UnauthorizedError struct {
	BaseError
}

func NewUnauthorizedError(message string) UnauthorizedError {
	return UnauthorizedError{
		BaseError: BaseError{
			Type:       "unauthorized",
			StatusCode: http.StatusUnauthorized,
			Message:    message,
		},
	}
}
