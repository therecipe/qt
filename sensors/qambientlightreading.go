package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAmbientLightReading struct {
	QSensorReading
}

type QAmbientLightReading_ITF interface {
	QSensorReading_ITF
	QAmbientLightReading_PTR() *QAmbientLightReading
}

func PointerFromQAmbientLightReading(ptr QAmbientLightReading_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAmbientLightReading_PTR().Pointer()
	}
	return nil
}

func NewQAmbientLightReadingFromPointer(ptr unsafe.Pointer) *QAmbientLightReading {
	var n = new(QAmbientLightReading)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAmbientLightReading_") {
		n.SetObjectName("QAmbientLightReading_" + qt.Identifier())
	}
	return n
}

func (ptr *QAmbientLightReading) QAmbientLightReading_PTR() *QAmbientLightReading {
	return ptr
}

//QAmbientLightReading::LightLevel
type QAmbientLightReading__LightLevel int64

const (
	QAmbientLightReading__Undefined = QAmbientLightReading__LightLevel(0)
	QAmbientLightReading__Dark      = QAmbientLightReading__LightLevel(1)
	QAmbientLightReading__Twilight  = QAmbientLightReading__LightLevel(2)
	QAmbientLightReading__Light     = QAmbientLightReading__LightLevel(3)
	QAmbientLightReading__Bright    = QAmbientLightReading__LightLevel(4)
	QAmbientLightReading__Sunny     = QAmbientLightReading__LightLevel(5)
)

func (ptr *QAmbientLightReading) LightLevel() QAmbientLightReading__LightLevel {
	defer qt.Recovering("QAmbientLightReading::lightLevel")

	if ptr.Pointer() != nil {
		return QAmbientLightReading__LightLevel(C.QAmbientLightReading_LightLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAmbientLightReading) SetLightLevel(lightLevel QAmbientLightReading__LightLevel) {
	defer qt.Recovering("QAmbientLightReading::setLightLevel")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_SetLightLevel(ptr.Pointer(), C.int(lightLevel))
	}
}

func (ptr *QAmbientLightReading) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAmbientLightReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAmbientLightReading::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAmbientLightReadingTimerEvent
func callbackQAmbientLightReadingTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightReading::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAmbientLightReading) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightReading) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::timerEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAmbientLightReading) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAmbientLightReading::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAmbientLightReading::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAmbientLightReadingChildEvent
func callbackQAmbientLightReadingChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightReading::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAmbientLightReading) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightReading) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::childEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAmbientLightReading) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAmbientLightReading::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAmbientLightReading) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAmbientLightReading::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAmbientLightReadingCustomEvent
func callbackQAmbientLightReadingCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAmbientLightReading::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAmbientLightReadingFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAmbientLightReading) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAmbientLightReading) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAmbientLightReading::customEvent")

	if ptr.Pointer() != nil {
		C.QAmbientLightReading_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
