package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSensorPluginInterface struct {
	ptr unsafe.Pointer
}

type QSensorPluginInterface_ITF interface {
	QSensorPluginInterface_PTR() *QSensorPluginInterface
}

func (p *QSensorPluginInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSensorPluginInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSensorPluginInterface(ptr QSensorPluginInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorPluginInterface_PTR().Pointer()
	}
	return nil
}

func NewQSensorPluginInterfaceFromPointer(ptr unsafe.Pointer) *QSensorPluginInterface {
	var n = new(QSensorPluginInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSensorPluginInterface) QSensorPluginInterface_PTR() *QSensorPluginInterface {
	return ptr
}

func (ptr *QSensorPluginInterface) RegisterSensors() {
	defer qt.Recovering("QSensorPluginInterface::registerSensors")

	if ptr.Pointer() != nil {
		C.QSensorPluginInterface_RegisterSensors(ptr.Pointer())
	}
}
