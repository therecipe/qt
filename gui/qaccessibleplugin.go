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

type QAccessiblePlugin_ITF interface {
	core.QObject_ITF
	QAccessiblePlugin_PTR() *QAccessiblePlugin
}

func PointerFromQAccessiblePlugin(ptr QAccessiblePlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessiblePlugin_PTR().Pointer()
	}
	return nil
}

func NewQAccessiblePluginFromPointer(ptr unsafe.Pointer) *QAccessiblePlugin {
	var n = new(QAccessiblePlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAccessiblePlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAccessiblePlugin) QAccessiblePlugin_PTR() *QAccessiblePlugin {
	return ptr
}
