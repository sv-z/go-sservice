package api

import (
	"fmt"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/sv-z/logger"
	"github.com/sv-z/validator"
	"net/http"
)

type UserApi struct {
	logger    *logger.Logger
	validator *validator.Validator
}

type UserRegisterCommand struct {
	Email    string `validate:"notblank,email"`
	Password string `validate:"notblank,max=20,min=4,pass_regex"`
}

func (this *UserApi) Register(_ *http.Request, command *UserRegisterCommand, result *bool) error {
	err := this.validator.Validate(command)
	if err != nil {
		*result = false
		return &json2.Error{Message: "Validation error.", Code: json2.E_BAD_PARAMS, Data: err}
	}

	this.logger.Info(fmt.Sprintf(
		"Start register new api \"%s\" with password \"%s\"\n", command.Email, command.Password))

	*result = true
	userId := 0

	this.logger.Info(fmt.Sprintf("User \"%s\" was created id: %d", command.Email, userId))

	return nil
}
