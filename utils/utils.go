package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type ErrorResponse struct {
	Field string `json:"field"`
	Error string `json:"error"`
}


// Initialize validator
var Validate = validator.New()

// Parses the request body into the provided struct and validates it
func ParseAndValidate(c *gin.Context, data any) []ErrorResponse {
	// Parse JSON body into struct
	if err := c.Bind(&data); err != nil {
		fmt.Println("Error(ParseAndValidate) ", err.Error())

		return []ErrorResponse{
			{
				Field: "",
				Error: "Invalid request body format." + err.Error(),
			},
		}
	}

	// Initialize validator
	validate := validator.New()

	// Validate struct
	if err := validate.Struct(data); err != nil {
		var errors []ErrorResponse
		for _, err := range err.(validator.ValidationErrors) {
			// Create key for custom error message (e.g., "Name.required")
			key := strings.ToLower(err.Field() + "." + err.Tag())
			// Get custom message or fallback to default
			message, exists := CustomErrorMessages[key]
			if !exists {
				message = err.Error()
			}
			errors = append(errors, ErrorResponse{
				Field: err.Field(),
				Error: message,
			})
		}
		return errors
	}

	return nil
}
