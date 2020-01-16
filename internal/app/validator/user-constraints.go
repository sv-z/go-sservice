package validator

import (
	valid "github.com/go-playground/validator/v10"

	"github.com/sv-z/in-scanner/internal/infrastructure"
	"github.com/sv-z/in-scanner/internal/validator"
)

type vContext struct {
	rm infrastructure.RepositoryManagerInterface
}

func (v vContext) configConstraintUserWithEmailNotExist() validator.Constraint {
	return func(fl valid.FieldLevel) bool {
		pw := fl.Field().String()
		return v.rm.User().FindByEmail(pw) == nil
	}
}
