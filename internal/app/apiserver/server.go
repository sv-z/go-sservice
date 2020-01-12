package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/sv-z/in-scanner/internal/infrastructure"
	"github.com/sv-z/in-scanner/internal/validator"
)

type server struct {
	logger            *logrus.Logger
	router            *mux.Router
	repositoryManager infrastructure.RepositoryManagerInterface
	validator         *validator.Validator
}

func newServer(rm infrastructure.RepositoryManagerInterface) *server {
	s := &server{
		logger:            logrus.New(),
		router:            mux.NewRouter(),
		repositoryManager: rm,
		validator:         validator.New(),
	}

	s.configureRouter()

	return s
}

func (srv *server) setLoggerLevel(logLevel string) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	srv.logger.SetLevel(level)
}

func (srv *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	srv.router.ServeHTTP(writer, request)
}

func (srv *server) configureRouter() {
	srv.router.HandleFunc("/ping", srv.handlePing()).Methods("GET")
	srv.router.HandleFunc("/info", srv.handleInfo()).Methods("GET")
	srv.router.HandleFunc("/", srv.handleInfo()).Methods("GET")

	// json-rps server
	srv.router.Handle("/api/", handleJsonRpcRequest(srv))

	// rest server
	handleRestRequest(srv)
}

func (srv *server) handlePing() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "PONG")
	}
}

func (srv *server) handleInfo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "Server Name: InScanner")
	}
}
