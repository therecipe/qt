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
func callbackQMediaAvailabilityControlAvailabilityChanged(ptrName *C.char, availability C.int) {
	defer qt.Recovering("callback QMediaAvailabilityControl::availabilityChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availabilityChanged"); signal != nil {
		signal.(func(QMultimedia__AvailabilityStatus))(QMultimedia__AvailabilityStatus(availability))
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
func callbackQMediaAvailabilityControlTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaAvailabilityControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMediaAvailabilityControlChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaAvailabilityControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMediaAvailabilityControlCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaAvailabilityControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
