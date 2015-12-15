package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QSensorGesturePluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorGesturePluginInterface_ITF interface {
	QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface
}

func (p *QSensorGesturePluginInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorGesturePluginInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorGesturePluginInterface(ptr QSensorGesturePluginInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesturePluginInterface_PTR().Pointer()
	}
	return nil
}

func NewQSensorGesturePluginInterfaceFromPointer(ptr unsafe.Pointer) *QSensorGesturePluginInterface {
	var n = new(QSensorGesturePluginInterface)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSensorGesturePluginInterface_") {
		n.SetObjectNameAbs("QSensorGesturePluginInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGesturePluginInterface) QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface {
	return ptr
}

func (ptr *QSensorGesturePluginInterface) Name() string {
	defer qt.Recovering("QSensorGesturePluginInterface::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGesturePluginInterface_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorGesturePluginInterface) SupportedIds() []string {
	defer qt.Recovering("QSensorGesturePluginInterface::supportedIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesturePluginInterface_SupportedIds(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterface() {
	defer qt.Recovering("QSensorGesturePluginInterface::~QSensorGesturePluginInterface")

	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(ptr.Pointer())
	}
}

func (ptr *QSensorGesturePluginInterface) ObjectNameAbs() string {
	defer qt.Recovering("QSensorGesturePluginInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGesturePluginInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorGesturePluginInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSensorGesturePluginInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
