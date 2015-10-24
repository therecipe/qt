package sensors

//#include "qaccelerometer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccelerometer struct {
	QSensor
}

type QAccelerometerITF interface {
	QSensorITF
	QAccelerometerPTR() *QAccelerometer
}

func PointerFromQAccelerometer(ptr QAccelerometerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccelerometerPTR().Pointer()
	}
	return nil
}

func QAccelerometerFromPointer(ptr unsafe.Pointer) *QAccelerometer {
	var n = new(QAccelerometer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAccelerometer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAccelerometer) QAccelerometerPTR() *QAccelerometer {
	return ptr
}

//QAccelerometer::AccelerationMode
type QAccelerometer__AccelerationMode int

var (
	QAccelerometer__Combined = QAccelerometer__AccelerationMode(0)
	QAccelerometer__Gravity  = QAccelerometer__AccelerationMode(1)
	QAccelerometer__User     = QAccelerometer__AccelerationMode(2)
)

func (ptr *QAccelerometer) AccelerationMode() QAccelerometer__AccelerationMode {
	if ptr.Pointer() != nil {
		return QAccelerometer__AccelerationMode(C.QAccelerometer_AccelerationMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccelerometer) Reading() *QAccelerometerReading {
	if ptr.Pointer() != nil {
		return QAccelerometerReadingFromPointer(unsafe.Pointer(C.QAccelerometer_Reading(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func NewQAccelerometer(parent core.QObjectITF) *QAccelerometer {
	return QAccelerometerFromPointer(unsafe.Pointer(C.QAccelerometer_NewQAccelerometer(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAccelerometer) ConnectAccelerationModeChanged(f func(accelerationMode QAccelerometer__AccelerationMode)) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_ConnectAccelerationModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "accelerationModeChanged", f)
	}
}

func (ptr *QAccelerometer) DisconnectAccelerationModeChanged() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DisconnectAccelerationModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "accelerationModeChanged")
	}
}

//export callbackQAccelerometerAccelerationModeChanged
func callbackQAccelerometerAccelerationModeChanged(ptrName *C.char, accelerationMode C.int) {
	qt.GetSignal(C.GoString(ptrName), "accelerationModeChanged").(func(QAccelerometer__AccelerationMode))(QAccelerometer__AccelerationMode(accelerationMode))
}

func (ptr *QAccelerometer) SetAccelerationMode(accelerationMode QAccelerometer__AccelerationMode) {
	if ptr.Pointer() != nil {
		C.QAccelerometer_SetAccelerationMode(C.QtObjectPtr(ptr.Pointer()), C.int(accelerationMode))
	}
}

func (ptr *QAccelerometer) DestroyQAccelerometer() {
	if ptr.Pointer() != nil {
		C.QAccelerometer_DestroyQAccelerometer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
