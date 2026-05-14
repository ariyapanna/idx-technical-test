package http

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"todolist-backend/internal/domain/apperror"
)

type PaginationMeta struct {
	CurrentPage  int   `json:"current_page"`
	PerPage int   `json:"per_page"`
	TotalItems   int64 `json:"total"`
	TotalPages   int   `json:"total_pages"`
}

type APIResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    any             `json:"data,omitempty"`
	Pagination *PaginationMeta `json:"pagination,omitempty"`
	Error   any             `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func PaginationResponse(c *gin.Context, statusCode int, message string, data any, meta PaginationMeta) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Pagination:    &meta,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	var errMessage any = err.Error()
	finalStatusCode := statusCode

	var ae *apperror.AppError
	if errors.As(err, &ae) {
		errMessage = ae.Message
		switch ae.Type {
		case apperror.TypeValidation:
			finalStatusCode = 422 
		case apperror.TypeNotFound:
			finalStatusCode = 404
		case apperror.TypeConflict:
			finalStatusCode = 409 
		case apperror.TypeInternal:
			finalStatusCode = 500 
		}
	} else {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make(map[string]string)
			for _, fe := range ve {
				out[strings.ToLower(fe.Field())] = getErrorMsg(fe)
			}
			errMessage = out
			finalStatusCode = 400 
		} else if errors.Is(err, io.EOF) {
			errMessage = "Request body is empty"
			finalStatusCode = 400 
		}
	}

	c.JSON(finalStatusCode, APIResponse{
		Success: false,
		Message: message,
		Error:   errMessage,
	})
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	}
	return fmt.Sprintf("Error on tag: %s", fe.Tag())
}
