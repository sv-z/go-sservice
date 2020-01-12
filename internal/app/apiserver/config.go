package apiserver

// ConnectionConfig ...
type Config struct {
	BindAddr            string `toml:"bind_addr"`
	LogLevel            string `toml:"log_level"`
	PostgresDatabaseUrl string `toml:"postgres_database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:            ":1234",
		LogLevel:            "debug",
		PostgresDatabaseUrl: "host=localhost dbname=local_server",
	}
}
