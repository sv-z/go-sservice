package apiserver

import (
	"io"
	"net/http"
)

func (srv *server) handleUserRegister() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "ServerName: InScanner")
	}
}
