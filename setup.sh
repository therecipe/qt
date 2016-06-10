#!/bin/sh
set -e
go get golang.org/x/crypto/ssh
go run ./internal/setup/check.go "$@"
go run $GOPATH/src/github.com/therecipe/qt/internal/setup/generate.go "$@"
go run $GOPATH/src/github.com/therecipe/qt/internal/setup/install.go "$@"
go run $GOPATH/src/github.com/therecipe/qt/internal/setup/test.go "$@"
