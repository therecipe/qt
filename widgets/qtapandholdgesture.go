package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTapAndHoldGesture struct {
	QGesture
}

type QTapAndHoldGesture_ITF interface {
	QGesture_ITF
	QTapAndHoldGesture_PTR() *QTapAndHoldGesture
}

func PointerFromQTapAndHoldGesture(ptr QTapAndHoldGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapAndHoldGesture_PTR().Pointer()
	}
	return nil
}

func NewQTapAndHoldGestureFromPointer(ptr unsafe.Pointer) *QTapAndHoldGesture {
	var n = new(QTapAndHoldGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTapAndHoldGesture_") {
		n.SetObjectName("QTapAndHoldGesture_" + qt.Identifier())
	}
	return n
}

func (ptr *QTapAndHoldGesture) QTapAndHoldGesture_PTR() *QTapAndHoldGesture {
	return ptr
}

func (ptr *QTapAndHoldGesture) SetPosition(pos core.QPointF_ITF) {
	defer qt.Recovering("QTapAndHoldGesture::setPosition")

	if ptr.Pointer() != nil {
		C.QTapAndHoldGesture_SetPosition(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func QTapAndHoldGesture_SetTimeout(msecs int) {
	defer qt.Recovering("QTapAndHoldGesture::setTimeout")

	C.QTapAndHoldGesture_QTapAndHoldGesture_SetTimeout(C.int(msecs))
}

func QTapAndHoldGesture_Timeout() int {
	defer qt.Recovering("QTapAndHoldGesture::timeout")

	return int(C.QTapAndHoldGesture_QTapAndHoldGesture_Timeout())
}

func (ptr *QTapAndHoldGesture) DestroyQTapAndHoldGesture() {
	defer qt.Recovering("QTapAndHoldGesture::~QTapAndHoldGesture")

	if ptr.Pointer() != nil {
		C.QTapAndHoldGesture_DestroyQTapAndHoldGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTapAndHoldGesture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTapAndHoldGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTapAndHoldGesture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTapAndHoldGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTapAndHoldGestureTimerEvent
func callbackQTapAndHoldGestureTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTapAndHoldGesture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTapAndHoldGesture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTapAndHoldGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTapAndHoldGesture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTapAndHoldGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTapAndHoldGestureChildEvent
func callbackQTapAndHoldGestureChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTapAndHoldGesture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTapAndHoldGesture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTapAndHoldGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTapAndHoldGesture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTapAndHoldGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTapAndHoldGestureCustomEvent
func callbackQTapAndHoldGestureCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTapAndHoldGesture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
