package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sv-z/in-scaner/internal/app/apiserver"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	apiserver := apiserver.New(config)
	err = apiserver.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}
