package validator

import (
	"encoding/json"
	"fmt"
	"strings"

	valid "github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
)

type Constraint valid.Func

type Validator struct {
	validate *valid.Validate
}

type Errors map[string][]map[string]interface{}

// Error returns the error string of Errors.
func (es Errors) Error() string {
	if len(es) == 0 {
		return ""
	}
	jsonString, err := json.Marshal(es)
	if err != nil {
		panic(fmt.Sprintf("Cannot decode errors to string, due error: %T - %v", err.Error(), err))
	}

	return string(jsonString)
}

func New() *Validator {
	var validator Validator
	validator.validate = valid.New()
	validator.validate.RegisterValidation("notblank", validators.NotBlank)
	validator.validate.RegisterValidation("pass_regex", IsPasswordValid)

	return &validator
}

// AddConstraint ...
func (val *Validator) AddConstraint(name string, message string, code string, fn Constraint) {
	val.validate.RegisterValidation(name, valid.Func(fn))
	addErrorMap(name, message, code)
}

// Validate ...
func (val *Validator) Validate(object interface{}) Errors {
	err := val.validate.Struct(object)

	if err != nil {

		// val check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like val.
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
