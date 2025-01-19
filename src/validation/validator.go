package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/samber/do/v2"
)

func NewValidatorInject(i do.Injector) (*validator.Validate, error) {
	validator := validator.New()

	return validator, nil
}
