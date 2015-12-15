package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QImageIOPlugin struct {
	core.QObject
}

type QImageIOPlugin_ITF interface {
	core.QObject_ITF
	QImageIOPlugin_PTR() *QImageIOPlugin
}

func PointerFromQImageIOPlugin(ptr QImageIOPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageIOPlugin_PTR().Pointer()
	}
	return nil
}

func NewQImageIOPluginFromPointer(ptr unsafe.Pointer) *QImageIOPlugin {
	var n = new(QImageIOPlugin)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QImageIOPlugin_") {
		n.SetObjectName("QImageIOPlugin_" + qt.Identifier())
	}
	return n
}

func (ptr *QImageIOPlugin) QImageIOPlugin_PTR() *QImageIOPlugin {
	return ptr
}

//QImageIOPlugin::Capability
type QImageIOPlugin__Capability int64

const (
	QImageIOPlugin__CanRead            = QImageIOPlugin__Capability(0x1)
	QImageIOPlugin__CanWrite           = QImageIOPlugin__Capability(0x2)
	QImageIOPlugin__CanReadIncremental = QImageIOPlugin__Capability(0x4)
)
