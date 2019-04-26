package qzxing

/*
#cgo CPPFLAGS: -DQZXING_QML -I ${SRCDIR}/qzxing/src/zxing

#cgo darwin,amd64,!ios LDFLAGS: -L ${SRCDIR}/qzxing/src/darwin
#cgo linux,amd64 LDFLAGS: -L ${SRCDIR}/qzxing/src/linux
#cgo windows,amd64 LDFLAGS: -L ${SRCDIR}/qzxing/src/windows

#cgo ios LDFLAGS: -L ${SRCDIR}/qzxing/src/ios

#cgo android,arm LDFLAGS: -L ${SRCDIR}/qzxing/src/android
#cgo android,386 LDFLAGS: -L ${SRCDIR}/qzxing/src/android_emulator

#cgo LDFLAGS: -lQZXing
#include "qzxing.h"
*/
import "C"
import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)

type stub struct{ core.QObject } //needed to make QtCore available to qzxing/src/QZXing.h

func RegisterQMLTypes() {
	C.QZXing_registerQMLTypes()
}

func RegisterQMLImageProvider(engine qml.QQmlEngine_ITF) {
	C.QZXing_registerQMLImageProvider(engine.QQmlEngine_PTR().Pointer())
}
