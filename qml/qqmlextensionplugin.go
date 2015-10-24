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

type QQmlExtensionPluginITF interface {
	core.QObjectITF
	QQmlExtensionPluginPTR() *QQmlExtensionPlugin
}

func PointerFromQQmlExtensionPlugin(ptr QQmlExtensionPluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlExtensionPluginPTR().Pointer()
	}
	return nil
}

func QQmlExtensionPluginFromPointer(ptr unsafe.Pointer) *QQmlExtensionPlugin {
	var n = new(QQmlExtensionPlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQmlExtensionPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlExtensionPlugin) QQmlExtensionPluginPTR() *QQmlExtensionPlugin {
	return ptr
}

func (ptr *QQmlExtensionPlugin) InitializeEngine(engine QQmlEngineITF, uri string) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_InitializeEngine(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQQmlEngine(engine)), C.CString(uri))
	}
}

func (ptr *QQmlExtensionPlugin) BaseUrl() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QQmlExtensionPlugin_BaseUrl(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QQmlExtensionPlugin) RegisterTypes(uri string) {
	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_RegisterTypes(C.QtObjectPtr(ptr.Pointer()), C.CString(uri))
	}
}
