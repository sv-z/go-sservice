package apiserver

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/sv-z/in-scaner/internal/infrastructure"
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

// postgres postgresDB init
func newPostgresDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
