package main

import (
	"github.com/sv-z/logger"
	"github.com/sv-z/sservice/api"
	"github.com/sv-z/validator"
)

func main() {
	validate := validator.New()
	logger, _ := logger.New("api")
	api.Run(logger, validate)
}
