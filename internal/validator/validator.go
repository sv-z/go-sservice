package validator

import (
	valid "github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"strings"
)

type Validator struct {
	validate *valid.Validate
}

func New() *Validator {
	var validator Validator
	validator.validate = valid.New()
	validator.validate.RegisterValidation("notblank", validators.NotBlank)
	validator.validate.RegisterValidation("pass_regex", IsPasswordValid)

	return &validator
}

func (this *Validator) Validate(object interface{}) map[string][]map[string]interface{} {
	err := this.validate.Struct(object)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*valid.InvalidValidationError); ok {
			panic(err)
		}

		errors := map[string][]map[string]interface{}{}
		var error = map[string]interface{}{}
		for _, err := range err.(valid.ValidationErrors) {
			code, message, values := getErrMessage(err)
			key := strings.ToLower(err.Field())

			error = map[string]interface{}{
				"code":    code,
				"message": message,
				"data":    values,
			}

			errors[key] = append(errors[key], error)
		}

		return errors
	}

	return nil
}
