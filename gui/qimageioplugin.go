package gui

//#include "qimageioplugin.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QImageIOPlugin struct {
	core.QObject
}

type QImageIOPluginITF interface {
	core.QObjectITF
	QImageIOPluginPTR() *QImageIOPlugin
}

func PointerFromQImageIOPlugin(ptr QImageIOPluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QImageIOPluginPTR().Pointer()
	}
	return nil
}

func QImageIOPluginFromPointer(ptr unsafe.Pointer) *QImageIOPlugin {
	var n = new(QImageIOPlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QImageIOPlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QImageIOPlugin) QImageIOPluginPTR() *QImageIOPlugin {
	return ptr
}

//QImageIOPlugin::Capability
type QImageIOPlugin__Capability int

var (
	QImageIOPlugin__CanRead            = QImageIOPlugin__Capability(0x1)
	QImageIOPlugin__CanWrite           = QImageIOPlugin__Capability(0x2)
	QImageIOPlugin__CanReadIncremental = QImageIOPlugin__Capability(0x4)
)
