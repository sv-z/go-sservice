.PHONY: build
build:
	go build -v ./cmd/apiserver

.DEFAULT_GOAL := build

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

help:
	./apiserver -help

run_env:
	docker-compose up -d

stop_env:
	docker-compose down && docker-compose down -v && docker-compose rm -f