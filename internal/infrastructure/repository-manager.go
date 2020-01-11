package infrastructure

import (
	"database/sql"
	_ "github.com/lib/pq"
	//"github.com/mitchellh/mapstructure"
)

// Config ...
type Config struct {
	DatabaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseUrl: "",
	}
}

type PostgresRepositoryManager struct {
	db     *sql.DB
	config *Config
}

// New ...
func New(config *Config) *PostgresRepositoryManager {
	return &PostgresRepositoryManager{
		config: config,
	}
}

// Open ...
func (prm *PostgresRepositoryManager) Open() error {
	db, err := sql.Open("postgres", prm.config.DatabaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	prm.db = db

	return nil
}

// Close ...
func (prm *PostgresRepositoryManager) Close() {

}
