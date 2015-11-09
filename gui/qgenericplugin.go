package gui

//#include "qgenericplugin.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGenericPlugin struct {
	core.QObject
}

type QGenericPlugin_ITF interface {
	core.QObject_ITF
	QGenericPlugin_PTR() *QGenericPlugin
}

func PointerFromQGenericPlugin(ptr QGenericPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGenericPlugin_PTR().Pointer()
	}
	return nil
}

func NewQGenericPluginFromPointer(ptr unsafe.Pointer) *QGenericPlugin {
	var n = new(QGenericPlugin)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QGenericPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGenericPlugin) QGenericPlugin_PTR() *QGenericPlugin {
	return ptr
}
