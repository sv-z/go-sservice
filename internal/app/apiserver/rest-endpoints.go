package apiserver

import (
	"io"
	"net/http"
)

func handleRestRequest(srv *server) {
	api := UserRestApi{srv: srv}

	srv.router.HandleFunc("/user", api.handleUserRegister()).Methods("GET")
}

// UserRestApi ...
type UserRestApi struct {
	srv *server
}

func (api *UserRestApi) handleUserRegister() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "user register rest method")
	}
}
