package sensors

//#include "qsensorgestureplugininterface.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QSensorGesturePluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorGesturePluginInterfaceITF interface {
	QSensorGesturePluginInterfacePTR() *QSensorGesturePluginInterface
}

func (p *QSensorGesturePluginInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorGesturePluginInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorGesturePluginInterface(ptr QSensorGesturePluginInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesturePluginInterfacePTR().Pointer()
	}
	return nil
}

func QSensorGesturePluginInterfaceFromPointer(ptr unsafe.Pointer) *QSensorGesturePluginInterface {
	var n = new(QSensorGesturePluginInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorGesturePluginInterface) QSensorGesturePluginInterfacePTR() *QSensorGesturePluginInterface {
	return ptr
}

func (ptr *QSensorGesturePluginInterface) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGesturePluginInterface_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSensorGesturePluginInterface) SupportedIds() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesturePluginInterface_SupportedIds(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterface() {
	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
