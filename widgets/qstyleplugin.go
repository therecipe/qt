package widgets

//#include "qstyleplugin.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QStylePlugin struct {
	core.QObject
}

type QStylePluginITF interface {
	core.QObjectITF
	QStylePluginPTR() *QStylePlugin
}

func PointerFromQStylePlugin(ptr QStylePluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStylePluginPTR().Pointer()
	}
	return nil
}

func QStylePluginFromPointer(ptr unsafe.Pointer) *QStylePlugin {
	var n = new(QStylePlugin)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QStylePlugin_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QStylePlugin) QStylePluginPTR() *QStylePlugin {
	return ptr
}

func (ptr *QStylePlugin) Create(key string) *QStyle {
	if ptr.Pointer() != nil {
		return QStyleFromPointer(unsafe.Pointer(C.QStylePlugin_Create(C.QtObjectPtr(ptr.Pointer()), C.CString(key))))
	}
	return nil
}

func (ptr *QStylePlugin) DestroyQStylePlugin() {
	if ptr.Pointer() != nil {
		C.QStylePlugin_DestroyQStylePlugin(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
