package dto

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CreatePostRequest struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required,min=10"`
}

func CustomErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "email":
		return "Invalid email address"
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s cannot be more than %s characters", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
}
