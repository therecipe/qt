package sensors

//#include "sensors.h"
import "C"
import (
	"log"
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
	return n
}

func (ptr *QSensorGesturePluginInterface) QSensorGesturePluginInterface_PTR() *QSensorGesturePluginInterface {
	return ptr
}

func (ptr *QSensorGesturePluginInterface) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesturePluginInterface::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGesturePluginInterface_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorGesturePluginInterface) SupportedIds() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesturePluginInterface::supportedIds")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesturePluginInterface_SupportedIds(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesturePluginInterface) DestroyQSensorGesturePluginInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorGesturePluginInterface::~QSensorGesturePluginInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(ptr.Pointer())
	}
}
