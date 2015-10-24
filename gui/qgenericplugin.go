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

type QGenericPluginITF interface {
	core.QObjectITF
	QGenericPluginPTR() *QGenericPlugin
}

func PointerFromQGenericPlugin(ptr QGenericPluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGenericPluginPTR().Pointer()
	}
	return nil
}

func QGenericPluginFromPointer(ptr unsafe.Pointer) *QGenericPlugin {
	var n = new(QGenericPlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGenericPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGenericPlugin) QGenericPluginPTR() *QGenericPlugin {
	return ptr
}
