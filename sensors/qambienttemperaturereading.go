package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAmbientTemperatureReading struct {
	QSensorReading
}

type QAmbientTemperatureReading_ITF interface {
	QSensorReading_ITF
	QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading
}

func PointerFromQAmbientTemperatureReading(ptr QAmbientTemperatureReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientTemperatureReading_PTR().Pointer()
	}
	return nil
}

func NewQAmbientTemperatureReadingFromPointer(ptr unsafe.Pointer) *QAmbientTemperatureReading {
	var n = new(QAmbientTemperatureReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAmbientTemperatureReading_") {
		n.SetObjectName("QAmbientTemperatureReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAmbientTemperatureReading) QAmbientTemperatureReading_PTR() *QAmbientTemperatureReading {
	return ptr
}

func (ptr *QAmbientTemperatureReading) Temperature() float64 {
	defer qt.Recovering("QAmbientTemperatureReading::temperature")

	if ptr.Pointer() != nil {
		return float64(C.QAmbientTemperatureReading_Temperature(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAmbientTemperatureReading) SetTemperature(temperature float64) {
	defer qt.Recovering("QAmbientTemperatureReading::setTemperature")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_SetTemperature(ptr.Pointer(), C.double(temperature))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAmbientTemperatureReadingTimerEvent
func callbackQAmbientTemperatureReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAmbientTemperatureReadingChildEvent
func callbackQAmbientTemperatureReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAmbientTemperatureReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAmbientTemperatureReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAmbientTemperatureReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAmbientTemperatureReadingCustomEvent
func callbackQAmbientTemperatureReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientTemperatureReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientTemperatureReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientTemperatureReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientTemperatureReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientTemperatureReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientTemperatureReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
