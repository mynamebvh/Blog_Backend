package utils

import "github.com/go-playground/validator"

func Validate(data interface{}) []ErrorResponse {
	var errors []ErrorResponse
	validate := validator.New()
	err := validate.Struct(data)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var errs ErrorResponse
			errs.Field = err.Field()
			errs.Message = HandlerError(err.Tag(), err.Field())
			errors = append(errors, errs)
		}
	}
	return errors
}
