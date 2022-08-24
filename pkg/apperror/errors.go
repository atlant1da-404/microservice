package apperror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ErrNotFound = NewAppError("not found", "NF", "")
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func NewAppError(message, code, developerMessage string) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Message:          message,
		DeveloperMessage: developerMessage,
		Code:             code,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) JSON(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(e)
}

func BadRequestError(message string) *AppError {
	return NewAppError(message, "NF", "some wrong with user data")
}

func SystemError(message string) *AppError {
	return NewAppError(message, "SE", "unknown error")
}
