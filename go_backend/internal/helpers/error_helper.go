package helpers

import (
	"github.com/go-playground/validator/v10"
)

func ExtractValidationErrors(err error) map[string]string {
	result := map[string]string{}

	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		tag := e.Tag()

		switch tag {
		case "required":
			result[field] = field + " is required"
		case "email":
			result[field] = "Invalid email format"
		case "min":
			result[field] = field + " is too short"
		case "max":
			result[field] = field + " is too long"
		default:
			result[field] = "Invalid value for " + field
		}
	}

	return result
}

func FirstValidationError(err error) string {

	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		tag := e.Tag()

		switch tag {
		case "required":
			return field + " is required."
		case "email":
			return "Invalid email format."
		case "min":
			return field + " is too short."
		case "max":
			return field + " is too long."
		default:
			return "Invalid value for " + field + "."
		}
	}

	return ""
}
