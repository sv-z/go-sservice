package api

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/sv-z/logger"
	"github.com/sv-z/validator"
	"net/http"
)

func Run(logger *logger.Logger, validator *validator.Validator) {
	server := rpc.NewServer()
	server.RegisterCodec(json2.NewCodec(), "application/json")
	server.RegisterCodec(json2.NewCodec(), "application/json;charset=UTF-8")

	server.RegisterService(&UserApi{logger, validator}, "user")

	router := mux.NewRouter()
	router.Handle("/api/", server)

	http.ListenAndServe(":1234", router)
}
