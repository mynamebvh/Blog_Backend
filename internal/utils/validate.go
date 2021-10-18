package utils

import (
	"fmt"

	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	Field   string
	Message string
}

func FormatValidateError(kind string, field string) string {
	switch kind {
	case "required":
		return fmt.Sprintf("Trường %s là bắt buộc", field)
	default:
		return fmt.Sprintf("Trường %s không đúng định dạng", field)
	}
}

func Validate(data interface{}) []ErrorResponse {
	var errors []ErrorResponse
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errs ErrorResponse
			errs.Field = err.Field()
			errs.Message = FormatValidateError(err.Tag(), err.Field())
			errors = append(errors, errs)
		}
	}
	return errors
}
