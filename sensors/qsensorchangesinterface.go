package sensors

//#include "qsensorchangesinterface.h"
import "C"
import (
	"unsafe"
)

type QSensorChangesInterface struct {
	ptr unsafe.Pointer
}

type QSensorChangesInterfaceITF interface {
	QSensorChangesInterfacePTR() *QSensorChangesInterface
}

func (p *QSensorChangesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorChangesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorChangesInterface(ptr QSensorChangesInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorChangesInterfacePTR().Pointer()
	}
	return nil
}

func QSensorChangesInterfaceFromPointer(ptr unsafe.Pointer) *QSensorChangesInterface {
	var n = new(QSensorChangesInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorChangesInterface) QSensorChangesInterfacePTR() *QSensorChangesInterface {
	return ptr
}

func (ptr *QSensorChangesInterface) SensorsChanged() {
	if ptr.Pointer() != nil {
		C.QSensorChangesInterface_SensorsChanged(C.QtObjectPtr(ptr.Pointer()))
	}
}
