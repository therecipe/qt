package sensors

//#include "qtapsensor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTapSensor struct {
	QSensor
}

type QTapSensorITF interface {
	QSensorITF
	QTapSensorPTR() *QTapSensor
}

func PointerFromQTapSensor(ptr QTapSensorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapSensorPTR().Pointer()
	}
	return nil
}

func QTapSensorFromPointer(ptr unsafe.Pointer) *QTapSensor {
	var n = new(QTapSensor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTapSensor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTapSensor) QTapSensorPTR() *QTapSensor {
	return ptr
}

func (ptr *QTapSensor) Reading() *QTapReading {
	if ptr.Pointer() != nil {
		return QTapReadingFromPointer(unsafe.Pointer(C.QTapSensor_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTapSensor) ReturnDoubleTapEvents() bool {
	if ptr.Pointer() != nil {
		return C.QTapSensor_ReturnDoubleTapEvents(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTapSensor) SetReturnDoubleTapEvents(returnDoubleTapEvents bool) {
	if ptr.Pointer() != nil {
		C.QTapSensor_SetReturnDoubleTapEvents(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(returnDoubleTapEvents)))
	}
}

func NewQTapSensor(parent core.QObjectITF) *QTapSensor {
	return QTapSensorFromPointer(unsafe.Pointer(C.QTapSensor_NewQTapSensor(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QTapSensor) ConnectReturnDoubleTapEventsChanged(f func(returnDoubleTapEvents bool)) {
	if ptr.Pointer() != nil {
		C.QTapSensor_ConnectReturnDoubleTapEventsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "returnDoubleTapEventsChanged", f)
	}
}

func (ptr *QTapSensor) DisconnectReturnDoubleTapEventsChanged() {
	if ptr.Pointer() != nil {
		C.QTapSensor_DisconnectReturnDoubleTapEventsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "returnDoubleTapEventsChanged")
	}
}

//export callbackQTapSensorReturnDoubleTapEventsChanged
func callbackQTapSensorReturnDoubleTapEventsChanged(ptrName *C.char, returnDoubleTapEvents C.int) {
	qt.GetSignal(C.GoString(ptrName), "returnDoubleTapEventsChanged").(func(bool))(int(returnDoubleTapEvents) != 0)
}

func (ptr *QTapSensor) DestroyQTapSensor() {
	if ptr.Pointer() != nil {
		C.QTapSensor_DestroyQTapSensor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
