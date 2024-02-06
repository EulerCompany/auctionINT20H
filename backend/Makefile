SHELL=/bin/bash

MAIN_PACKAGE_PATH := ./cmd/web
BINARY_NAME := auction-backend

.DEFAULT_GOAL := run

.PHONY: build
build:
	GOARCH=amd64 GOOS=linux go build -o=./target/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

.PHONY: build-image
build-image: build
	sudo docker build -t auction-backend -f build/Dockerfile.backend .

.PHONY: run
run: build
	./target/${BINARY_NAME}
	
.PHONY: compose-up
compose-up: build-image
	sudo docker compose -f build/docker-compose.yaml up

.PHONY: compose-down
compose-down:
	sudo docker compose -v -f build/docker-compose.yaml down

.PHONY: mysql
mysql:
	sudo docker run -d --rm -p 3306:3306 -e MYSQL_USER=web -e MYSQL_PASSWORD=pass -e MYSQL_ALLOW_EMPTY_PASSWORD=yes mysql:8.0-bullseye 
.PHONY: clean
clean:
	go clean
	rm -rf ./target/

.PHONY: dep
dep:
	go mod download


