package sensors

//#include "sensors.h"
import "C"
import (
	"log"
	"unsafe"
)

type QSensorChangesInterface struct {
	ptr unsafe.Pointer
}

type QSensorChangesInterface_ITF interface {
	QSensorChangesInterface_PTR() *QSensorChangesInterface
}

func (p *QSensorChangesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorChangesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorChangesInterface(ptr QSensorChangesInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorChangesInterface_PTR().Pointer()
	}
	return nil
}

func NewQSensorChangesInterfaceFromPointer(ptr unsafe.Pointer) *QSensorChangesInterface {
	var n = new(QSensorChangesInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorChangesInterface) QSensorChangesInterface_PTR() *QSensorChangesInterface {
	return ptr
}

func (ptr *QSensorChangesInterface) SensorsChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSensorChangesInterface::sensorsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSensorChangesInterface_SensorsChanged(ptr.Pointer())
	}
}
