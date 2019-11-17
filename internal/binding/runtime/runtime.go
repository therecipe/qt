package runtime

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
)


func init() {
	h := NewHelper(nil)
	h.ConnectWrapperFunction(qml.Z_wrapperFunction)
	qml.Z_engineHelper = h.QObject_PTR()
}

type helper struct {
	core.QObject

	_ func(func())                      `slot:"runOnMain,auto"`
	_ func(*qml.QJSValue) *qml.QJSValue `slot:"wrapperFunction"`
}

func (*helper) runOnMain(f func()) { f() }
