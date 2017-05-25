#!/bin/bash
set -ev

qtrcc
cd ./component
qtmoc
qtrcc
sed -i '' -e 's/main/component/' ./rcc_cgo_darwin_darwin_amd64.go
$QT_DIR/5.8/clang_64/bin/rcc -name SomeComponent -o rcc.cpp SomeComponent.qrc
cd ..
go build -o test_qml_go && ./test_qml_go
