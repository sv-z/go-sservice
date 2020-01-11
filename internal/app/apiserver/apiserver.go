package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/sv-z/in-scaner/internal/infrastructure"
)

// APIServer ...
type APIServer struct {
	config            *Config
	logger            *logrus.Logger
	router            *mux.Router
	connectionHolder  *infrastructure.ConnectionHolder
	repositoryManager *infrastructure.RepositoryManager
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config:           config,
		logger:           logrus.New(),
		router:           mux.NewRouter(),
		connectionHolder: infrastructure.New(config.Postgres),
	}
}

func (server *APIServer) Run() error {
	if err := server.configureLogger(); err != nil {
		return err
	}

	server.configureRouter()
	if err := server.configurePostgres(); err != nil {
		return err
	}

	if err := server.configureRepositoryManager(server.connectionHolder); err != nil {
		return err
	}

	server.logger.Info("Starting api server")

	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}

	server.logger.SetLevel(level)

	return nil
}

func (server *APIServer) configureRouter() {
	server.router.Handle("/ping", server.handlePing())
}

func (server *APIServer) configurePostgres() error {
	store := infrastructure.New(server.config.Postgres)
	if err := store.Init(); err != nil {
		return err
	}

	server.connectionHolder = store

	return nil
}

func (server *APIServer) configureRepositoryManager(connectionHolder *infrastructure.ConnectionHolder) error {
	server.repositoryManager = infrastructure.NewRepositoryManager(connectionHolder)

	return nil
}

func (server *APIServer) handlePing() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "pong")
	}
}
