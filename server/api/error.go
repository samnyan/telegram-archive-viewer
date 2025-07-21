package api

import (
	"fmt"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("AppError: %d: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("AppError: %d: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// WrapError wraps an error with AppError
func WrapErrorRaw(err error, code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
func WrapError(err error, appError *AppError) *AppError {
	return &AppError{
		Code:    appError.Code,
		Message: appError.Message,
		Err:     err,
	}
}

// Common errors:
var (
	ErrElementNotFound = &AppError{Code: 404, Message: "Element not found"}
	ErrTokenInvalid    = &AppError{Code: 403, Message: "Token invalid or expired"}
	ErrInternalServer  = &AppError{Code: 500, Message: "Internal Server Error"}
)

type DatabaseError struct {
	IsNotFound bool
	Err        error
}

func (e *DatabaseError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("DatabaseError: %v", e.Err)
	}
	return "Unknown DatabaseError"
}

func (e *DatabaseError) Unwrap() error {
	return e.Err
}
