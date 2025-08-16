package validator

import "github.com/go-playground/validator/v10"

var validatorInstance *validator.Validate = validator.New()

func ValidateStruct(data interface{}) error {
	return validatorInstance.Struct(data)
}
