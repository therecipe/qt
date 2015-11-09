package qml

//#include "qqmlextensionplugin.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlExtensionPlugin struct {
	core.QObject
}

type QQmlExtensionPlugin_ITF interface {
	core.QObject_ITF
	QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin
}

func PointerFromQQmlExtensionPlugin(ptr QQmlExtensionPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExtensionPlugin_PTR().Pointer()
	}
	return nil
}

func NewQQmlExtensionPluginFromPointer(ptr unsafe.Pointer) *QQmlExtensionPlugin {
	var n = new(QQmlExtensionPlugin)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QQmlExtensionPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlExtensionPlugin) QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin {
	return ptr
}

func (ptr *QQmlExtensionPlugin) InitializeEngine(engine QQmlEngine_ITF, uri string) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_InitializeEngine(ptr.Pointer(), PointerFromQQmlEngine(engine), C.CString(uri))
	}
}

func (ptr *QQmlExtensionPlugin) RegisterTypes(uri string) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_RegisterTypes(ptr.Pointer(), C.CString(uri))
	}
}
