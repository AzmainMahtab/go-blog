// Package handlers
package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/AzmainMahtab/go-blog/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Success creates a successful response
func Success(c *gin.Context, data interface{}, meta interface{}, message string) {
	status := http.StatusOK
	if c.Request.Method == http.MethodPost {
		status = http.StatusCreated
	}

	c.JSON(status, responses.Response{
		Success:    true,
		Status:     "success",
		StatusCode: status,
		Message:    message,
		Data:       data,
		Meta:       meta,
	})
}

// Error handles all error responses consistently
func Error(c *gin.Context, err error) {
	var response responses.Response
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		response = handleNotFoundError(err)
	case errors.Is(err, gorm.ErrDuplicatedKey):
		response = handleDuplicateError(err)
	case errors.Is(err, gorm.ErrInvalidData):
		response = handleInvalidDataError(err)
	case errors.Is(err, gorm.ErrInvalidTransaction):
		response = handleTransactionError(err)
	case errors.Is(err, gorm.ErrInvalidField):
		response = handleInvalidFieldError(err)
	default:
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			response = handleValidationError(validationErrs)
		} else {
			response = handleGenericError(err)
		}
	}

	c.JSON(response.StatusCode, response)
}

// Specific error handlers
func handleNotFoundError(err error) responses.Response {
	return responses.Response{
		Success:    false,
		Status:     "fail",
		StatusCode: http.StatusNotFound,
		Message:    "Requested resource not found",
		Errors: []responses.ErrorItem{
			{Code: "RESOURCE_NOT_FOUND", Message: err.Error()},
		},
	}
}

func handleDuplicateError(err error) responses.Response {
	return responses.Response{
		Success:    false,
		Status:     "fail",
		StatusCode: http.StatusConflict,
		Message:    "Resource already exists",
		Errors: []responses.ErrorItem{
			{Code: "DUPLICATE_RESOURCE", Message: err.Error()},
		},
	}
}

func handleValidationError(errs validator.ValidationErrors) responses.Response {
	errors := make([]responses.ErrorItem, len(errs))
	for i, err := range errs {
		errors[i] = responses.ErrorItem{
			Field:   err.Field(),
			Code:    strings.ToUpper(err.Tag()),
			Message: validationMessage(err),
		}
	}

	return responses.Response{
		Success:    false,
		Status:     "fail",
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Validation failed",
		Errors:     errors,
	}
}

// Helper function for validation messages
func validationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return "Value is too short"
	case "max":
		return "Value is too long"
	// Add more cases as needed
	default:
		return err.Error() // default error message
	}
}

func handleGenericError(err error) responses.Response {
	return responses.Response{
		Success:    false,
		Status:     "error",
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal server error",
		Errors: []responses.ErrorItem{
			{Code: "INTERNAL_ERROR", Message: err.Error()},
		},
	}
}

// handleInvalidDataError handles GORM invalid data errors
func handleInvalidDataError(err error) responses.Response {
	return responses.Response{
		Success:    false,
		Status:     "fail",
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid data provided",
		Errors: []responses.ErrorItem{
			{
				Code:    "INVALID_DATA",
				Message: "The provided data is invalid or malformed",
			},
		},
	}
}

// handleTransactionError handles GORM transaction errors
func handleTransactionError(err error) responses.Response {
	return responses.Response{
		Success:    false,
		Status:     "error",
		StatusCode: http.StatusInternalServerError,
		Message:    "Database transaction failed",
		Errors: []responses.ErrorItem{
			{
				Code:    "TRANSACTION_ERROR",
				Message: "A database transaction operation failed",
			},
		},
	}
}

// handleInvalidFieldError handles GORM invalid field errors
func handleInvalidFieldError(err error) responses.Response {
	return responses.Response{
		Success:    false,
		Status:     "fail",
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid field in request",
		Errors: []responses.ErrorItem{
			{
				Code:    "INVALID_FIELD",
				Message: "One or more fields in the request are invalid",
			},
		},
	}
}

// handleNotNullViolationError handles GORM not null violation errors
func handleNotNullViolationError(err error) responses.Response {
	// Extract field name from error message if possible
	field := "unknown"
	if strings.Contains(err.Error(), "column") {
		parts := strings.Split(err.Error(), "column ")
		if len(parts) > 1 {
			field = strings.Split(parts[1], " ")[0]
			field = strings.Trim(field, `"`)
		}
	}

	return responses.Response{
		Success:    false,
		Status:     "fail",
		StatusCode: http.StatusBadRequest,
		Message:    "Required field missing",
		Errors: []responses.ErrorItem{
			{
				Code:    "NULL_CONSTRAINT_VIOLATION",
				Field:   field,
				Message: fmt.Sprintf("Field '%s' cannot be null", field),
			},
		},
	}
}
