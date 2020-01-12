package apiserver

import (
	"net/http"

	"github.com/sv-z/in-scanner/internal/infrastructure"
)

// Start ...
func Start(config *Config) error {

	db, err := newPostgresDB(config.PostgresDatabaseUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	rm := infrastructure.NewRepositoryManager(db)
	srv := newServer(rm)
	srv.setLoggerLevel(config.LogLevel)

	return http.ListenAndServe(config.BindAddr, srv)
}
