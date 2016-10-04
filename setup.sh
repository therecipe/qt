#!/bin/sh
set -ev

go get golang.org/x/crypto/ssh
go get github.com/emirpasic/gods/lists/arraylist
go run $GOPATH/src/github.com/therecipe/qt/internal/setup/check.go "$@"
go run $GOPATH/src/github.com/therecipe/qt/internal/setup/generate.go "$@"
go run $GOPATH/src/github.com/therecipe/qt/internal/setup/install.go "$@"
go run $GOPATH/src/github.com/therecipe/qt/internal/setup/test.go "$@"
