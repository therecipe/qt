package main

/*
#cgo CPPFLAGS: -I ${SRCDIR}/WinSparkle-0.6.0/include
#cgo windows,386 LDFLAGS: -L ${SRCDIR}/WinSparkle-0.6.0/Release
#cgo windows,amd64 LDFLAGS: -L ${SRCDIR}/WinSparkle-0.6.0/x64/Release
#cgo LDFLAGS: -lWinSparkle

#include "sparkle.h"
*/
import "C"

func sparkle_checkUpdates() { C.sparkle_checkUpdates() }
