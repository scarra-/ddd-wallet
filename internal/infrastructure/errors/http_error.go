package errors

import (
	"net/http"

	val "github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Message string              `json:"message"`
	Details []map[string]string `json:"details"`
}

func NewErrorResponse(code int) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusText(code),
		Message: "",
	}
}

func NewCustomErrorResponse(code int, message string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusText(code),
		Message: message,
	}
}

func NewValidationErrorResponse(err val.ValidationErrors) ValidationErrorResponse {
	return ValidationErrorResponse{
		Message: "Validation failed",
		Details: formatValidationErrors(err),
	}
}

func formatValidationErrors(vErr val.ValidationErrors) []map[string]string {
	errors := make([]map[string]string, 0, len(vErr))
	for _, fieldError := range vErr {
		errors = append(errors, map[string]string{
			"field":   fieldError.StructField(),
			"message": fieldError.Tag(),
		})
	}
	return errors
}
