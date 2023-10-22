SHELL := /bin/bash

-include .env

# set a default value for port
SERVER_PORT ?= 8080

.PHONY: get-packs-number
get-packs-number:
	curl "http://localhost:$(SERVER_PORT)/api/v1/get-packs-number/$(items)"

.PHONY: start
start:
	docker-compose up --build -d

.PHONY: lint
stop:
	docker-compose down

.PHONY: build
build:
	docker-compose build

.PHONY: rebuild-and-restart
rebuild-and-restart:
	make stop && make build && make start

.PHONY: test-unit
test-unit:
	go test ./internal/...

.PHONY: test-integration
test-integration:
	docker-compose up --build -d
	go test -v ./integration-tests/...
	docker-compose down


.PHONY: tests
tests: test-unit test-integration


# you should have golangci-lint installed to use it
.PHONY: lint
lint:
	golangci-lint --config=./golangci.yml run -v