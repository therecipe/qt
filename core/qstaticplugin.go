package core

//#include "qstaticplugin.h"
import "C"
import (
	"unsafe"
)

type QStaticPlugin struct {
	ptr unsafe.Pointer
}

type QStaticPluginITF interface {
	QStaticPluginPTR() *QStaticPlugin
}

func (p *QStaticPlugin) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QStaticPlugin) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQStaticPlugin(ptr QStaticPluginITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStaticPluginPTR().Pointer()
	}
	return nil
}

func QStaticPluginFromPointer(ptr unsafe.Pointer) *QStaticPlugin {
	var n = new(QStaticPlugin)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStaticPlugin) QStaticPluginPTR() *QStaticPlugin {
	return ptr
}

func (ptr *QStaticPlugin) Instance() *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QStaticPlugin_Instance(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
