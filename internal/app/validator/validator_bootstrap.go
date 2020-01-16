package validator

import (
	"github.com/sv-z/in-scanner/internal/infrastructure"
	valid "github.com/sv-z/in-scanner/internal/validator"
)

func New(rm infrastructure.RepositoryManagerInterface) *valid.Validator {
	vContext := vContext{rm}

	v := valid.New()
	v.AddConstraint(
		"email_not_exist",
		`User with email "{{ value }}" already exists.`,
		"336ca680-f6fa-4cf5-8885-81ae62177b97",
		vContext.configConstraintUserWithEmailNotExist(),
	)

	return v
}
