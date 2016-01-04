package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QSensorGestureRecognizer struct {
	core.QObject
}

type QSensorGestureRecognizer_ITF interface {
	core.QObject_ITF
	QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer
}

func PointerFromQSensorGestureRecognizer(ptr QSensorGestureRecognizer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGestureRecognizer_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureRecognizerFromPointer(ptr unsafe.Pointer) *QSensorGestureRecognizer {
	var n = new(QSensorGestureRecognizer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSensorGestureRecognizer_") {
		n.SetObjectName("QSensorGestureRecognizer_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGestureRecognizer) QSensorGestureRecognizer_PTR() *QSensorGestureRecognizer {
	return ptr
}

func (ptr *QSensorGestureRecognizer) CreateBackend() {
	defer qt.Recovering("QSensorGestureRecognizer::createBackend")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CreateBackend(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) GestureSignals() []string {
	defer qt.Recovering("QSensorGestureRecognizer::gestureSignals")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGestureRecognizer_GestureSignals(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGestureRecognizer) Id() string {
	defer qt.Recovering("QSensorGestureRecognizer::id")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSensorGestureRecognizer_Id(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSensorGestureRecognizer) IsActive() bool {
	defer qt.Recovering("QSensorGestureRecognizer::isActive")

	if ptr.Pointer() != nil {
		return C.QSensorGestureRecognizer_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGestureRecognizer) StartBackend() {
	defer qt.Recovering("QSensorGestureRecognizer::startBackend")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StartBackend(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) StopBackend() {
	defer qt.Recovering("QSensorGestureRecognizer::stopBackend")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_StopBackend(ptr.Pointer())
	}
}

func (ptr *QSensorGestureRecognizer) DestroyQSensorGestureRecognizer() {
	defer qt.Recovering("QSensorGestureRecognizer::~QSensorGestureRecognizer")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGestureRecognizer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensorGestureRecognizer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensorGestureRecognizer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorGestureRecognizerTimerEvent
func callbackQSensorGestureRecognizerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGestureRecognizer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensorGestureRecognizer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensorGestureRecognizer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorGestureRecognizerChildEvent
func callbackQSensorGestureRecognizerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGestureRecognizer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensorGestureRecognizer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensorGestureRecognizer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensorGestureRecognizer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorGestureRecognizerCustomEvent
func callbackQSensorGestureRecognizerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGestureRecognizer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureRecognizerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGestureRecognizer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorGestureRecognizer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGestureRecognizer::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGestureRecognizer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
