package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTapSensor struct {
	QSensor
}

type QTapSensor_ITF interface {
	QSensor_ITF
	QTapSensor_PTR() *QTapSensor
}

func PointerFromQTapSensor(ptr QTapSensor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapSensor_PTR().Pointer()
	}
	return nil
}

func NewQTapSensorFromPointer(ptr unsafe.Pointer) *QTapSensor {
	var n = new(QTapSensor)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTapSensor_") {
		n.SetObjectName("QTapSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTapSensor) QTapSensor_PTR() *QTapSensor {
	return ptr
}

func (ptr *QTapSensor) Reading() *QTapReading {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapSensor::reading")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQTapReadingFromPointer(C.QTapSensor_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTapSensor) ReturnDoubleTapEvents() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapSensor::returnDoubleTapEvents")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTapSensor_ReturnDoubleTapEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTapSensor) SetReturnDoubleTapEvents(returnDoubleTapEvents bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapSensor::setReturnDoubleTapEvents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTapSensor_SetReturnDoubleTapEvents(ptr.Pointer(), C.int(qt.GoBoolToInt(returnDoubleTapEvents)))
	}
}

func NewQTapSensor(parent core.QObject_ITF) *QTapSensor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapSensor::QTapSensor")
		}
	}()

	return NewQTapSensorFromPointer(C.QTapSensor_NewQTapSensor(core.PointerFromQObject(parent)))
}

func (ptr *QTapSensor) ConnectReturnDoubleTapEventsChanged(f func(returnDoubleTapEvents bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapSensor::returnDoubleTapEventsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTapSensor_ConnectReturnDoubleTapEventsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "returnDoubleTapEventsChanged", f)
	}
}

func (ptr *QTapSensor) DisconnectReturnDoubleTapEventsChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapSensor::returnDoubleTapEventsChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTapSensor_DisconnectReturnDoubleTapEventsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "returnDoubleTapEventsChanged")
	}
}

//export callbackQTapSensorReturnDoubleTapEventsChanged
func callbackQTapSensorReturnDoubleTapEventsChanged(ptrName *C.char, returnDoubleTapEvents C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapSensor::returnDoubleTapEventsChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "returnDoubleTapEventsChanged").(func(bool))(int(returnDoubleTapEvents) != 0)
}

func (ptr *QTapSensor) DestroyQTapSensor() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTapSensor::~QTapSensor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTapSensor_DestroyQTapSensor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
