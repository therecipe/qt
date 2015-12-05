package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	for len(n.ObjectName()) < len("QQmlExtensionPlugin_") {
		n.SetObjectName("QQmlExtensionPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQmlExtensionPlugin) QQmlExtensionPlugin_PTR() *QQmlExtensionPlugin {
	return ptr
}

func (ptr *QQmlExtensionPlugin) InitializeEngine(engine QQmlEngine_ITF, uri string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlExtensionPlugin::initializeEngine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_InitializeEngine(ptr.Pointer(), PointerFromQQmlEngine(engine), C.CString(uri))
	}
}

func (ptr *QQmlExtensionPlugin) RegisterTypes(uri string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlExtensionPlugin::registerTypes")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQmlExtensionPlugin_RegisterTypes(ptr.Pointer(), C.CString(uri))
	}
}
