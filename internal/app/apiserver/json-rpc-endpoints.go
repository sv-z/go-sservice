package apiserver

import (
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

func handleJsonRpcRequest(srv *server) *rpc.Server {
	server := rpc.NewServer()
	server.RegisterCodec(json2.NewCodec(), "application/json")
	server.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")

	server.RegisterService(&UserJsonRpcApi{srv: srv}, "user")
	// server.RegisterBeforeFunc(func(i *rpc.RequestInfo){
	// 	methodParts := strings.Split(i.Method, ".")
	// 	i.Method = methodParts[0] + "." + strings.Title(methodParts[0])
	// });

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

func (api *UserJsonRpcApi) Register(_ *http.Request, command *UserRegisterCommand, result *bool) error {
	return nil
}
