package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSensorBackend struct {
	core.QObject
}

type QSensorBackend_ITF interface {
	core.QObject_ITF
	QSensorBackend_PTR() *QSensorBackend
}

func PointerFromQSensorBackend(ptr QSensorBackend_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorBackend_PTR().Pointer()
	}
	return nil
}

func NewQSensorBackendFromPointer(ptr unsafe.Pointer) *QSensorBackend {
	var n = new(QSensorBackend)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSensorBackend_") {
		n.SetObjectName("QSensorBackend_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorBackend) QSensorBackend_PTR() *QSensorBackend {
	return ptr
}

func (ptr *QSensorBackend) AddDataRate(min float64, max float64) {
	defer qt.Recovering("QSensorBackend::addDataRate")

	if ptr.Pointer() != nil {
		C.QSensorBackend_AddDataRate(ptr.Pointer(), C.double(min), C.double(max))
	}
}

func (ptr *QSensorBackend) IsFeatureSupported(feature QSensor__Feature) bool {
	defer qt.Recovering("QSensorBackend::isFeatureSupported")

	if ptr.Pointer() != nil {
		return C.QSensorBackend_IsFeatureSupported(ptr.Pointer(), C.int(feature)) != 0
	}
	return false
}

func (ptr *QSensorBackend) SensorBusy() {
	defer qt.Recovering("QSensorBackend::sensorBusy")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorBusy(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) SensorError(error int) {
	defer qt.Recovering("QSensorBackend::sensorError")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorError(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QSensorBackend) AddOutputRange(min float64, max float64, accuracy float64) {
	defer qt.Recovering("QSensorBackend::addOutputRange")

	if ptr.Pointer() != nil {
		C.QSensorBackend_AddOutputRange(ptr.Pointer(), C.double(min), C.double(max), C.double(accuracy))
	}
}

func (ptr *QSensorBackend) NewReadingAvailable() {
	defer qt.Recovering("QSensorBackend::newReadingAvailable")

	if ptr.Pointer() != nil {
		C.QSensorBackend_NewReadingAvailable(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) Reading() *QSensorReading {
	defer qt.Recovering("QSensorBackend::reading")

	if ptr.Pointer() != nil {
		return NewQSensorReadingFromPointer(C.QSensorBackend_Reading(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensorBackend) Sensor() *QSensor {
	defer qt.Recovering("QSensorBackend::sensor")

	if ptr.Pointer() != nil {
		return NewQSensorFromPointer(C.QSensorBackend_Sensor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSensorBackend) SensorStopped() {
	defer qt.Recovering("QSensorBackend::sensorStopped")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SensorStopped(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) SetDataRates(otherSensor QSensor_ITF) {
	defer qt.Recovering("QSensorBackend::setDataRates")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SetDataRates(ptr.Pointer(), PointerFromQSensor(otherSensor))
	}
}

func (ptr *QSensorBackend) SetDescription(description string) {
	defer qt.Recovering("QSensorBackend::setDescription")

	if ptr.Pointer() != nil {
		C.QSensorBackend_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QSensorBackend) Start() {
	defer qt.Recovering("QSensorBackend::start")

	if ptr.Pointer() != nil {
		C.QSensorBackend_Start(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) Stop() {
	defer qt.Recovering("QSensorBackend::stop")

	if ptr.Pointer() != nil {
		C.QSensorBackend_Stop(ptr.Pointer())
	}
}

func (ptr *QSensorBackend) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensorBackend::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensorBackend) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensorBackend::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorBackendTimerEvent
func callbackQSensorBackendTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSensorBackend::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSensorBackend) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensorBackend::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensorBackend) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensorBackend::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorBackendChildEvent
func callbackQSensorBackendChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSensorBackend::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QSensorBackend) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensorBackend::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensorBackend) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensorBackend::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorBackendCustomEvent
func callbackQSensorBackendCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSensorBackend::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
