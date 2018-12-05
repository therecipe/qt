package main

/*
#cgo CFLAGS: -I ${SRCDIR}/darwin/Contents/Frameworks/Sparkle.framework
#cgo LDFLAGS: -F ${SRCDIR}/darwin/Contents/Frameworks -framework Sparkle

void sparkle_checkUpdates();
*/
import "C"

func sparkle_checkUpdates() { C.sparkle_checkUpdates() }
