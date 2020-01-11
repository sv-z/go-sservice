package infrastructure

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// ConnectionConfig ...
type ConnectionConfig struct {
	PostgresDatabaseUrl string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *ConnectionConfig {
	return &ConnectionConfig{
		PostgresDatabaseUrl: "",
	}
}

// ConnectionHolder ...
type ConnectionHolder struct {
	postgresDB *sql.DB
	config     *ConnectionConfig
}

// New ...
func New(config *ConnectionConfig) *ConnectionHolder {
	return &ConnectionHolder{
		config: config,
	}
}

// Init all connection...
func (prm *ConnectionHolder) Init() error {
	if err := prm.intPostgres(); err != nil {
		return err
	}

	return nil
}

// Close ...
func (prm *ConnectionHolder) Close() {

}

// postgres postgresDB init
func (prm *ConnectionHolder) intPostgres() error {
	db, err := sql.Open("postgres", prm.config.PostgresDatabaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	prm.postgresDB = db

	return nil
}
