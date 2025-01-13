package validate

import (
	validator "github.com/go-playground/validator/v10"
)

var (
	Validator *validator.Validate
)

func init() {
	Validator = validator.New(
		validator.WithRequiredStructEnabled(),
	)
}

func Errors(err error) map[string]string {
	errors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {
		errors[err.Field()] = err.Tag()
	}

	return errors
}

func Struct(s interface{}) error {
	return Validator.Struct(s)
}
