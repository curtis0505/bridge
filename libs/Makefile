#!/bin/make

generate-all: generate-repository-v1 generate-repository-v2 generate-mock

generate-repository-v1:
	@go run ./internal/cmd.go generate repository --version v1

generate-repository-v2:
	@go run ./internal/cmd.go generate repository --version v2

generate-mock:
	@mockery