.PHONY: build-alpine help run fmt

GO_BIN = go
BIN_NAME=small-service
IMAGE_NAME=small-service

BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')

default: help

help:
	@echo 'Management commands:'
	@echo
	@echo 'make build-alpine    Compile optimized for alpine linux.'
	@echo 'make fmt             Format the go code across the project'
	@echo 'make run             Run ${small-service} locally'
	@echo

build-alpine:
	@echo "Building ${BIN_NAME}"
	CC=$(which musl-gcc) ${GO_BIN} build -ldflags "-linkmode external -extldflags -static" -o bin/${BIN_NAME} ./cmd/${BIN_NAME}

fmt:
	@echo "Formatting ${BIN_NAME}"
	go fmt ./...

run:
	@echo "Running ${BIN_NAME} locally"
	go run cmd/${BIN_NAME}/main.go
