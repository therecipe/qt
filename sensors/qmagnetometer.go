package sensors

//#include "qmagnetometer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMagnetometer struct {
	QSensor
}

type QMagnetometerITF interface {
	QSensorITF
	QMagnetometerPTR() *QMagnetometer
}

func PointerFromQMagnetometer(ptr QMagnetometerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometerPTR().Pointer()
	}
	return nil
}

func QMagnetometerFromPointer(ptr unsafe.Pointer) *QMagnetometer {
	var n = new(QMagnetometer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMagnetometer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMagnetometer) QMagnetometerPTR() *QMagnetometer {
	return ptr
}

func (ptr *QMagnetometer) Reading() *QMagnetometerReading {
	if ptr.Pointer() != nil {
		return QMagnetometerReadingFromPointer(unsafe.Pointer(C.QMagnetometer_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMagnetometer) ReturnGeoValues() bool {
	if ptr.Pointer() != nil {
		return C.QMagnetometer_ReturnGeoValues(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMagnetometer) SetReturnGeoValues(returnGeoValues bool) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_SetReturnGeoValues(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(returnGeoValues)))
	}
}

func NewQMagnetometer(parent core.QObjectITF) *QMagnetometer {
	return QMagnetometerFromPointer(unsafe.Pointer(C.QMagnetometer_NewQMagnetometer(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QMagnetometer) ConnectReturnGeoValuesChanged(f func(returnGeoValues bool)) {
	if ptr.Pointer() != nil {
		C.QMagnetometer_ConnectReturnGeoValuesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "returnGeoValuesChanged", f)
	}
}

func (ptr *QMagnetometer) DisconnectReturnGeoValuesChanged() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DisconnectReturnGeoValuesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "returnGeoValuesChanged")
	}
}

//export callbackQMagnetometerReturnGeoValuesChanged
func callbackQMagnetometerReturnGeoValuesChanged(ptrName *C.char, returnGeoValues C.int) {
	qt.GetSignal(C.GoString(ptrName), "returnGeoValuesChanged").(func(bool))(int(returnGeoValues) != 0)
}

func (ptr *QMagnetometer) DestroyQMagnetometer() {
	if ptr.Pointer() != nil {
		C.QMagnetometer_DestroyQMagnetometer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
