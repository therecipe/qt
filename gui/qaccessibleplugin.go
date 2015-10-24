package gui

//#include "qaccessibleplugin.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessiblePlugin struct {
	core.QObject
}

type QAccessiblePluginITF interface {
	core.QObjectITF
	QAccessiblePluginPTR() *QAccessiblePlugin
}

func PointerFromQAccessiblePlugin(ptr QAccessiblePluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessiblePluginPTR().Pointer()
	}
	return nil
}

func QAccessiblePluginFromPointer(ptr unsafe.Pointer) *QAccessiblePlugin {
	var n = new(QAccessiblePlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAccessiblePlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAccessiblePlugin) QAccessiblePluginPTR() *QAccessiblePlugin {
	return ptr
}
