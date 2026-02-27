package helpers

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err validator.ValidationErrors) []string {
	var errorList []string

	for _, e := range err {
		log.Println(e.Param())

		errorList = append(errorList, formatFieldValidationError(e))

	}

	return errorList
}

func formatFieldValidationError(e validator.FieldError) string {
	switch e.Tag() {
	case "email":
		return fmt.Sprintf("%s must be an email", e.Field())
	case "required":
		return fmt.Sprintf("%s is a required", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param())

	default:
		return fmt.Sprintf("%s is invalid", e.Field())
	}
}
