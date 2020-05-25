package errors

import (
	"net/http"
	"time"
)


type ApiError interface {
	Status() int
	Error() string
	GetMessage() string
}

type apiError struct {
	StatusCode int `json:"status"`
	Message string `json:"message"`
	Timestamp string `json:"timestamp"`
	ErrorMessage string `json:"error"`
	Path string `json:"path"`
}

func (err *apiError) Status() int {
	return err.StatusCode
}

func (err *apiError) Error() string {
	return err.ErrorMessage
}

func (err *apiError) GetMessage() string {
	return err.Message
}


func NewCustomApiError(statusCode int, message string) ApiError {
	return &apiError{
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  time.Now().String(),
		ErrorMessage:      "Not found",
		Path:       "",
	}
}


func NotFoundException(message string) ApiError {
	return &apiError{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Timestamp:  time.Now().String(),
		ErrorMessage:      "Not found",
		Path:       "",
	}
}


func InternalServerErrorException(message string) ApiError {
	return &apiError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
		Timestamp:  time.Now().String(),
		ErrorMessage:      "Internal server error",
		Path:       "",
	}
}

func BadRequestException(message string) ApiError {
	return &apiError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Timestamp:  time.Now().String(),
		ErrorMessage:      "Bad request",
		Path:       "",
	}
}

func NotAcceptableException(message string) ApiError {
	return &apiError{
		StatusCode: http.StatusNotAcceptable,
		Message:    message,
		Timestamp:  time.Now().String(),
		ErrorMessage:      "Bad request",
		Path:       "",
	}
}

