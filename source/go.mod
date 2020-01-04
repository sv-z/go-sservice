module github.com/sv-z/sservice

replace github.com/sv-z/logger => ./logger

replace github.com/sv-z/validator => ./validator

go 1.13

require (
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/rpc v1.2.0
	github.com/sv-z/logger v0.0.0-00010101000000-000000000000
	github.com/sv-z/validator v0.0.0-00010101000000-000000000000
)
