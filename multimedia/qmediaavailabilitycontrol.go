package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaAvailabilityControl struct {
	QMediaControl
}

type QMediaAvailabilityControl_ITF interface {
	QMediaControl_ITF
	QMediaAvailabilityControl_PTR() *QMediaAvailabilityControl
}

func PointerFromQMediaAvailabilityControl(ptr QMediaAvailabilityControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaAvailabilityControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaAvailabilityControlFromPointer(ptr unsafe.Pointer) *QMediaAvailabilityControl {
	var n = new(QMediaAvailabilityControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaAvailabilityControl_") {
		n.SetObjectName("QMediaAvailabilityControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaAvailabilityControl) QMediaAvailabilityControl_PTR() *QMediaAvailabilityControl {
	return ptr
}

func (ptr *QMediaAvailabilityControl) Availability() QMultimedia__AvailabilityStatus {
	defer qt.Recovering("QMediaAvailabilityControl::availability")

	if ptr.Pointer() != nil {
		return QMultimedia__AvailabilityStatus(C.QMediaAvailabilityControl_Availability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMediaAvailabilityControl) ConnectAvailabilityChanged(f func(availability QMultimedia__AvailabilityStatus)) {
	defer qt.Recovering("connect QMediaAvailabilityControl::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_ConnectAvailabilityChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availabilityChanged", f)
	}
}

func (ptr *QMediaAvailabilityControl) DisconnectAvailabilityChanged() {
	defer qt.Recovering("disconnect QMediaAvailabilityControl::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_DisconnectAvailabilityChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availabilityChanged")
	}
}

//export callbackQMediaAvailabilityControlAvailabilityChanged
func callbackQMediaAvailabilityControlAvailabilityChanged(ptr unsafe.Pointer, ptrName *C.char, availability C.int) {
	defer qt.Recovering("callback QMediaAvailabilityControl::availabilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availabilityChanged"); signal != nil {
		signal.(func(QMultimedia__AvailabilityStatus))(QMultimedia__AvailabilityStatus(availability))
	}

}

func (ptr *QMediaAvailabilityControl) AvailabilityChanged(availability QMultimedia__AvailabilityStatus) {
	defer qt.Recovering("QMediaAvailabilityControl::availabilityChanged")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_AvailabilityChanged(ptr.Pointer(), C.int(availability))
	}
}

func (ptr *QMediaAvailabilityControl) DestroyQMediaAvailabilityControl() {
	defer qt.Recovering("QMediaAvailabilityControl::~QMediaAvailabilityControl")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_DestroyQMediaAvailabilityControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaAvailabilityControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaAvailabilityControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaAvailabilityControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaAvailabilityControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaAvailabilityControlTimerEvent
func callbackQMediaAvailabilityControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAvailabilityControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMediaAvailabilityControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMediaAvailabilityControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaAvailabilityControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaAvailabilityControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaAvailabilityControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaAvailabilityControlChildEvent
func callbackQMediaAvailabilityControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAvailabilityControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMediaAvailabilityControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMediaAvailabilityControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::childEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaAvailabilityControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaAvailabilityControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaAvailabilityControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaAvailabilityControlCustomEvent
func callbackQMediaAvailabilityControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMediaAvailabilityControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMediaAvailabilityControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMediaAvailabilityControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMediaAvailabilityControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMediaAvailabilityControl::customEvent")

	if ptr.Pointer() != nil {
		C.QMediaAvailabilityControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
