package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMagnetometer struct {
	QSensor
}

type QMagnetometer_ITF interface {
	QSensor_ITF
	QMagnetometer_PTR() *QMagnetometer
}

func PointerFromQMagnetometer(ptr QMagnetometer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMagnetometer_PTR().Pointer()
	}
	return nil
}

func NewQMagnetometerFromPointer(ptr unsafe.Pointer) *QMagnetometer {
	var n = new(QMagnetometer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMagnetometer_") {
		n.SetObjectName("QMagnetometer_" + qt.Identifier())
	}
	return n
}

func (ptr *QMagnetometer) QMagnetometer_PTR() *QMagnetometer {
	return ptr
}

func (ptr *QMagnetometer) Reading() *QMagnetometerReading {
	defer qt.Recovering("QMagnetometer::reading")

	if ptr.Pointer() != nil {
		return NewQMagnetometerReadingFromPointer(C.QMagnetometer_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMagnetometer) ReturnGeoValues() bool {
	defer qt.Recovering("QMagnetometer::returnGeoValues")

	if ptr.Pointer() != nil {
		return C.QMagnetometer_ReturnGeoValues(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMagnetometer) SetReturnGeoValues(returnGeoValues bool) {
	defer qt.Recovering("QMagnetometer::setReturnGeoValues")

	if ptr.Pointer() != nil {
		C.QMagnetometer_SetReturnGeoValues(ptr.Pointer(), C.int(qt.GoBoolToInt(returnGeoValues)))
	}
}

func NewQMagnetometer(parent core.QObject_ITF) *QMagnetometer {
	defer qt.Recovering("QMagnetometer::QMagnetometer")

	return NewQMagnetometerFromPointer(C.QMagnetometer_NewQMagnetometer(core.PointerFromQObject(parent)))
}

func (ptr *QMagnetometer) ConnectReturnGeoValuesChanged(f func(returnGeoValues bool)) {
	defer qt.Recovering("connect QMagnetometer::returnGeoValuesChanged")

	if ptr.Pointer() != nil {
		C.QMagnetometer_ConnectReturnGeoValuesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnGeoValuesChanged", f)
	}
}

func (ptr *QMagnetometer) DisconnectReturnGeoValuesChanged() {
	defer qt.Recovering("disconnect QMagnetometer::returnGeoValuesChanged")

	if ptr.Pointer() != nil {
		C.QMagnetometer_DisconnectReturnGeoValuesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnGeoValuesChanged")
	}
}

//export callbackQMagnetometerReturnGeoValuesChanged
func callbackQMagnetometerReturnGeoValuesChanged(ptrName *C.char, returnGeoValues C.int) {
	defer qt.Recovering("callback QMagnetometer::returnGeoValuesChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "returnGeoValuesChanged")
	if signal != nil {
		signal.(func(bool))(int(returnGeoValues) != 0)
	}

}

func (ptr *QMagnetometer) DestroyQMagnetometer() {
	defer qt.Recovering("QMagnetometer::~QMagnetometer")

	if ptr.Pointer() != nil {
		C.QMagnetometer_DestroyQMagnetometer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
