package sensors

//#include "qsensorplugininterface.h"
import "C"
import (
	"unsafe"
)

type QSensorPluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorPluginInterfaceITF interface {
	QSensorPluginInterfacePTR() *QSensorPluginInterface
}

func (p *QSensorPluginInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorPluginInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorPluginInterface(ptr QSensorPluginInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorPluginInterfacePTR().Pointer()
	}
	return nil
}

func QSensorPluginInterfaceFromPointer(ptr unsafe.Pointer) *QSensorPluginInterface {
	var n = new(QSensorPluginInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorPluginInterface) QSensorPluginInterfacePTR() *QSensorPluginInterface {
	return ptr
}

func (ptr *QSensorPluginInterface) RegisterSensors() {
	if ptr.Pointer() != nil {
		C.QSensorPluginInterface_RegisterSensors(C.QtObjectPtr(ptr.Pointer()))
	}
}
