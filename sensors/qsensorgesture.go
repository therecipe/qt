package sensors

//#include "sensors.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QSensorGesture struct {
	core.QObject
}

type QSensorGesture_ITF interface {
	core.QObject_ITF
	QSensorGesture_PTR() *QSensorGesture
}

func PointerFromQSensorGesture(ptr QSensorGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSensorGesture_PTR().Pointer()
	}
	return nil
}

func NewQSensorGestureFromPointer(ptr unsafe.Pointer) *QSensorGesture {
	var n = new(QSensorGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSensorGesture_") {
		n.SetObjectName("QSensorGesture_" + qt.Identifier())
	}
	return n
}

func (ptr *QSensorGesture) QSensorGesture_PTR() *QSensorGesture {
	return ptr
}

func NewQSensorGesture(ids []string, parent core.QObject_ITF) *QSensorGesture {
	defer qt.Recovering("QSensorGesture::QSensorGesture")

	return NewQSensorGestureFromPointer(C.QSensorGesture_NewQSensorGesture(C.CString(strings.Join(ids, ",,,")), core.PointerFromQObject(parent)))
}

func (ptr *QSensorGesture) GestureSignals() []string {
	defer qt.Recovering("QSensorGesture::gestureSignals")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_GestureSignals(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) InvalidIds() []string {
	defer qt.Recovering("QSensorGesture::invalidIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_InvalidIds(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) IsActive() bool {
	defer qt.Recovering("QSensorGesture::isActive")

	if ptr.Pointer() != nil {
		return C.QSensorGesture_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSensorGesture) StartDetection() {
	defer qt.Recovering("QSensorGesture::startDetection")

	if ptr.Pointer() != nil {
		C.QSensorGesture_StartDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) StopDetection() {
	defer qt.Recovering("QSensorGesture::stopDetection")

	if ptr.Pointer() != nil {
		C.QSensorGesture_StopDetection(ptr.Pointer())
	}
}

func (ptr *QSensorGesture) ValidIds() []string {
	defer qt.Recovering("QSensorGesture::validIds")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSensorGesture_ValidIds(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QSensorGesture) DestroyQSensorGesture() {
	defer qt.Recovering("QSensorGesture::~QSensorGesture")

	if ptr.Pointer() != nil {
		C.QSensorGesture_DestroyQSensorGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSensorGesture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSensorGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSensorGesture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSensorGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSensorGestureTimerEvent
func callbackQSensorGestureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGesture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGesture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSensorGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSensorGesture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSensorGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSensorGesture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSensorGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSensorGestureChildEvent
func callbackQSensorGestureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGesture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGesture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSensorGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSensorGesture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSensorGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSensorGesture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSensorGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSensorGestureCustomEvent
func callbackQSensorGestureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSensorGesture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSensorGestureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSensorGesture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSensorGesture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSensorGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QSensorGesture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
