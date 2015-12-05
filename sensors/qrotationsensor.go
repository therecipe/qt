package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QRotationSensor struct {
	QSensor
}

type QRotationSensor_ITF interface {
	QSensor_ITF
	QRotationSensor_PTR() *QRotationSensor
}

func PointerFromQRotationSensor(ptr QRotationSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRotationSensor_PTR().Pointer()
	}
	return nil
}

func NewQRotationSensorFromPointer(ptr unsafe.Pointer) *QRotationSensor {
	var n = new(QRotationSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRotationSensor_") {
		n.SetObjectName("QRotationSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRotationSensor) QRotationSensor_PTR() *QRotationSensor {
	return ptr
}

func (ptr *QRotationSensor) HasZ() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationSensor::hasZ")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRotationSensor_HasZ(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRotationSensor) Reading() *QRotationReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationSensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQRotationReadingFromPointer(C.QRotationSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func NewQRotationSensor(parent core.QObject_ITF) *QRotationSensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationSensor::QRotationSensor")
		}
	}()

	return NewQRotationSensorFromPointer(C.QRotationSensor_NewQRotationSensor(core.PointerFromQObject(parent)))
}

func (ptr *QRotationSensor) ConnectHasZChanged(f func(hasZ bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationSensor::hasZChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRotationSensor_ConnectHasZChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hasZChanged", f)
	}
}

func (ptr *QRotationSensor) DisconnectHasZChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationSensor::hasZChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRotationSensor_DisconnectHasZChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hasZChanged")
	}
}

//export callbackQRotationSensorHasZChanged
func callbackQRotationSensorHasZChanged(ptrName *C.char, hasZ C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationSensor::hasZChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "hasZChanged").(func(bool))(int(hasZ) != 0)
}

func (ptr *QRotationSensor) SetHasZ(hasZ bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationSensor::setHasZ")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRotationSensor_SetHasZ(ptr.Pointer(), C.int(qt.GoBoolToInt(hasZ)))
	}
}

func (ptr *QRotationSensor) DestroyQRotationSensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRotationSensor::~QRotationSensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRotationSensor_DestroyQRotationSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
