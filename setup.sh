#!/bin/sh
set -e
go run ./internal/setup/check.go "$@"
go run ./internal/setup/generate.go "$@"
go run ./internal/setup/install.go "$@"
go run ./internal/setup/test.go "$@"
