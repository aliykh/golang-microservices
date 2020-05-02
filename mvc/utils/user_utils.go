package utils

type ApplicationError struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Message string `json:"message"`
}