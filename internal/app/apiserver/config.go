package apiserver

import "github.com/sv-z/in-scaner/internal/infrastructure"

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Postgres *infrastructure.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":1234",
		LogLevel: "debug",
		Postgres: infrastructure.NewConfig(),
	}
}
