package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/sv-z/in-scaner/internal/infrastructure"
)

type server struct {
	logger            *logrus.Logger
	router            *mux.Router
	repositoryManager *infrastructure.RepositoryManagerInterface
}

func newServer(rm infrastructure.RepositoryManagerInterface) *server {
	s := &server{
		logger:            logrus.New(),
		router:            mux.NewRouter(),
		repositoryManager: &rm,
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
	srv.router.HandleFunc("/info", srv.handlePing()).Methods("GET")
}

func (srv *server) handlePing() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "pong")
	}
}

func (srv *server) handleInfo() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "ServerName: InScanner")
	}
}
