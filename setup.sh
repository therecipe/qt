#!/bin/sh

go run ./internal/setup/generate.go "$@"
go run ./internal/setup/install.go "$@"
go run ./internal/setup/test.go "$@"
