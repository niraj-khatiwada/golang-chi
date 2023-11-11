package utils

import validator "github.com/go-ozzo/ozzo-validation/v4"

type Validation struct {
	Key   string
	Value interface{}
	Rules []validator.Rule
}
type ValidationError struct {
	Path    interface{}
	Message string
}

type ValidationErrors []ValidationError

func ValidateInput(validations []Validation) ValidationErrors {
	var vErrs ValidationErrors
	for _, validation := range validations {
		if err := validator.Validate(validation.Value, validation.Rules...); err != nil {
			ve := ValidationError{Path: validation.Key, Message: err.Error()}
			vErrs = append(vErrs, ve)
		}
	}
	return vErrs
}
