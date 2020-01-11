package apiserver

import "github.com/sv-z/in-scaner/internal/infrastructure"

// ConnectionConfig ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Postgres *infrastructure.ConnectionConfig
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":1234",
		LogLevel: "debug",
		Postgres: infrastructure.NewConfig(),
	}
}
