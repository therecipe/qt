package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QMagnetometer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMagnetometer) QMagnetometer_PTR() *QMagnetometer {
	return ptr
}

func (ptr *QMagnetometer) Reading() *QMagnetometerReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometer::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMagnetometerReadingFromPointer(C.QMagnetometer_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMagnetometer) ReturnGeoValues() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometer::returnGeoValues")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QMagnetometer_ReturnGeoValues(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMagnetometer) SetReturnGeoValues(returnGeoValues bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometer::setReturnGeoValues")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMagnetometer_SetReturnGeoValues(ptr.Pointer(), C.int(qt.GoBoolToInt(returnGeoValues)))
	}
}

func NewQMagnetometer(parent core.QObject_ITF) *QMagnetometer {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometer::QMagnetometer")
		}
	}()

	return NewQMagnetometerFromPointer(C.QMagnetometer_NewQMagnetometer(core.PointerFromQObject(parent)))
}

func (ptr *QMagnetometer) ConnectReturnGeoValuesChanged(f func(returnGeoValues bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometer::returnGeoValuesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMagnetometer_ConnectReturnGeoValuesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnGeoValuesChanged", f)
	}
}

func (ptr *QMagnetometer) DisconnectReturnGeoValuesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometer::returnGeoValuesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMagnetometer_DisconnectReturnGeoValuesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnGeoValuesChanged")
	}
}

//export callbackQMagnetometerReturnGeoValuesChanged
func callbackQMagnetometerReturnGeoValuesChanged(ptrName *C.char, returnGeoValues C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometer::returnGeoValuesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "returnGeoValuesChanged").(func(bool))(int(returnGeoValues) != 0)
}

func (ptr *QMagnetometer) DestroyQMagnetometer() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMagnetometer::~QMagnetometer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMagnetometer_DestroyQMagnetometer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
