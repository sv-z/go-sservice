package apiserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"

	"github.com/sv-z/in-scanner/internal/model"
)

func handleJsonRpcRequest(srv *server) *rpc.Server {
	server := rpc.NewServer()
	server.RegisterCodec(json2.NewCodec(), "application/json")
	server.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")

	server.RegisterService(&UserJsonRpcApi{srv: srv}, "user")

	return server
}

// UserJsonRpcApi ...
type UserJsonRpcApi struct {
	srv *server
}

type UserRegisterCommand struct {
	Email    string `validate:"notblank,email"`
	Password string `validate:"notblank,max=20,min=4,pass_regex"`
}

// Api handler - create new user
func (api *UserJsonRpcApi) Register(_ *http.Request, command *UserRegisterCommand, result *bool) error {
	err := api.srv.validator.Validate(command)
	if err != nil {
		*result = false
		return &json2.Error{Message: "Validation error.", Code: json2.E_BAD_PARAMS, Data: err}
	}

	api.srv.logger.Info(fmt.Sprintf("Start register new api \"%s\"", command.Email))

	user := model.CreateNewUser(command.Email, command.Password)

	if err := api.srv.repositoryManager.User().Save(user); err != nil {
		api.srv.logger.Info(fmt.Sprintf("%T - %v", err, err))
		*result = false
		return &json2.Error{
			Message: "User not saved. An infrastructure error occurred. Please try again later.",
			Code:    json2.E_INTERNAL,
		}
	}

	api.srv.logger.Info(fmt.Sprintf("User \"%s\" was created id: %d", command.Email, user.Id))
	*result = true

	return nil
}

type FindQuery struct {
}

// Api handler - Find user by params
func (api *UserJsonRpcApi) Find(_ *http.Request, query *FindQuery, result *[]interface{}) error {
	err := api.srv.validator.Validate(query)
	if err != nil {
		*result = make([]interface{}, 0)
		return &json2.Error{Message: "Validation error.", Code: json2.E_BAD_PARAMS, Data: err}
	}

	users := api.srv.repositoryManager.User().GetAll()
	res := make([]interface{}, len(users))
	for i, user := range users {
		responseRow := make(map[string]interface{})
		responseRow["id"] = user.Id
		responseRow["email"] = user.Email

		res[i] = responseRow
	}

	*result = res

	return nil
}
