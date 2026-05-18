package errors

import (
	"errors"
	"fmt"
)

// ErrorCode represents a standardized error code.
type ErrorCode string

const (
	ErrNotFound       ErrorCode = "NOT_FOUND"
	ErrAlreadyExists  ErrorCode = "ALREADY_EXISTS"
	ErrInvalidInput   ErrorCode = "INVALID_INPUT"
	ErrUnauthorized   ErrorCode = "UNAUTHORIZED"
	ErrForbidden      ErrorCode = "FORBIDDEN"
	ErrConflict       ErrorCode = "CONFLICT"
	ErrInternal       ErrorCode = "INTERNAL_ERROR"
	ErrNotImplemented ErrorCode = "NOT_IMPLEMENTED"
)

// DomainError is a typed error with a code and message.
type DomainError struct {
	Code    ErrorCode
	Message string
	Details string
}

func (e *DomainError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("[%s] %s: %s", e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// New creates a new DomainError.
func New(code ErrorCode, message string) *DomainError {
	return &DomainError{Code: code, Message: message}
}

// Newf creates a new DomainError with formatted message.
func Newf(code ErrorCode, format string, args ...interface{}) *DomainError {
	return &DomainError{Code: code, Message: fmt.Sprintf(format, args...)}
}

// WithDetails adds details to a DomainError.
func (e *DomainError) WithDetails(details string) *DomainError {
	e.Details = details
	return e
}

// Is checks if an error matches a specific code.
func Is(err error, code ErrorCode) bool {
	var domainErr *DomainError
	if errors.As(err, &domainErr) {
		return domainErr.Code == code
	}
	return false
}

// AsDomainError extracts a DomainError from an error.
func AsDomainError(err error) (*DomainError, bool) {
	var domainErr *DomainError
	if errors.As(err, &domainErr) {
		return domainErr, true
	}
	return nil, false
}

// Predefined common errors
var (
	ErrNotFoundDefault     = New(ErrNotFound, "resource not found")
	ErrAlreadyExistsDefault = New(ErrAlreadyExists, "resource already exists")
	ErrInvalidInputDefault  = New(ErrInvalidInput, "invalid input")
	ErrUnauthorizedDefault  = New(ErrUnauthorized, "unauthorized")
	ErrForbiddenDefault     = New(ErrForbidden, "forbidden")
	ErrConflictDefault      = New(ErrConflict, "resource conflict")
	ErrInternalDefault      = New(ErrInternal, "internal server error")
)
