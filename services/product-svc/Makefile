VERSION := $(shell cat VERSION)
COMMIT_HASH := $(shell git rev-parse HEAD)
PROJECT_NAME := productsvc

all: generate build-all

.PHONY: generate
generate:
	buf generate api

.PHONY: build
build:
	go build \
		-ldflags "-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH)" \
		-o build/$(PROJECT_NAME) \
		cmd/$(PROJECT_NAME)/*.go

.PHONY: clean
clean:
	rm -rf build/

.PHONY: run
run:
	go run main.go

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: proto
proto:
	protoc \
	-I api/productsvc \
	--go_out=./internal/generated \
	--go-grpc_out=./internal/generated \
	--validate_out="lang=go:./internal/generated" \
	--openapiv2_out=./api/productsvc \
	--grpc-gateway_out ./internal/generated --grpc-gateway_opt generate_unbound_methods=true \
	--experimental_allow_proto3_optional \
	api/productsvc/productsvc.proto
