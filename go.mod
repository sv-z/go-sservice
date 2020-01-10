module github.com/sv-z/sservice

go 1.13

replace github.com/sv-z/logger => ./logger

replace github.com/sv-z/validator => ./validator

replace github.com/sv-z/sservice/api => ./api

require github.com/sv-z/sservice/api v0.0.0-00010101000000-000000000000 // indirect
