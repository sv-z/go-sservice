package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sv-z/in-scanner/internal/infrastructure/errors"
	"github.com/sv-z/in-scanner/internal/model"
	"github.com/sv-z/in-scanner/internal/validator"
)

func handleRestRequest(srv *server) {
	api := UserRestApi{srv: srv}

	srv.router.HandleFunc("/users", api.handleUserFind()).Methods("GET")
	srv.router.HandleFunc("/users", api.handleUserCreate()).Methods("POST")
	srv.router.HandleFunc("/users/{id:[0-9]+}", api.handleGetUser()).Methods("GET")
}

// UserRestApi ...
type UserRestApi struct {
	srv *server
}

func (api *UserRestApi) error(writer http.ResponseWriter, request *http.Request, code int, err error) {
	api.srv.logger.Info(fmt.Sprintf("%T - %v, http error code %d", err, err, code))

	if vErr, ok := err.(validator.Errors); !ok {
		api.respond(writer, request, code, map[string]string{"error": err.Error()})
	} else {
		api.respond(writer, request, code, map[string]validator.Errors{"error": vErr})
	}
}

func (api *UserRestApi) respond(writer http.ResponseWriter, request *http.Request, code int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(writer).Encode(data); err != nil {
			panic(err)
		}
	}
}

func (api *UserRestApi) handleUserCreate() http.HandlerFunc {

	type requestData struct {
		Email    string `json:"email"validate:"notblank,email,email_not_exist"`
		Password string `json:"password"validate:"notblank,max=20,min=4,pass_regex"`
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		req := &requestData{}

		if err := json.NewDecoder(request.Body).Decode(req); err != nil {
			api.error(writer, request, http.StatusBadRequest, err)
			return
		}

		if err := api.srv.validator.Validate(req); err != nil {
			api.error(writer, request, http.StatusBadRequest, err)
			return
		}

		user := model.CreateNewUser(req.Email, req.Password)
		if err := api.srv.repositoryManager.User().Save(user); err != nil {
			api.error(writer, request, http.StatusUnprocessableEntity, err)
			return
		}

		api.respond(writer, request, http.StatusCreated, nil)
	}
}

func (api *UserRestApi) handleUserFind() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		users := api.srv.repositoryManager.User().GetAll()
		api.respond(writer, request, http.StatusOK, users)
	}
}

func (api *UserRestApi) handleGetUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		idParam := mux.Vars(request)["id"]
		userId, err := strconv.Atoi(idParam)
		if err != nil {
			api.respond(writer, request, http.StatusBadRequest, nil)
		}

		user, err := api.srv.repositoryManager.User().GetById(userId)
		switch err {
		case errors.NoRecord:
			api.respond(writer, request, http.StatusNotFound, nil)
		case nil:
			api.respond(writer, request, http.StatusOK, user)
		default:
			api.respond(writer, request, http.StatusBadRequest, nil)
		}
	}
}
