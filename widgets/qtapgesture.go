package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTapGesture struct {
	QGesture
}

type QTapGesture_ITF interface {
	QGesture_ITF
	QTapGesture_PTR() *QTapGesture
}

func PointerFromQTapGesture(ptr QTapGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTapGesture_PTR().Pointer()
	}
	return nil
}

func NewQTapGestureFromPointer(ptr unsafe.Pointer) *QTapGesture {
	var n = new(QTapGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTapGesture_") {
		n.SetObjectName("QTapGesture_" + qt.Identifier())
	}
	return n
}

func (ptr *QTapGesture) QTapGesture_PTR() *QTapGesture {
	return ptr
}

func (ptr *QTapGesture) SetPosition(pos core.QPointF_ITF) {
	defer qt.Recovering("QTapGesture::setPosition")

	if ptr.Pointer() != nil {
		C.QTapGesture_SetPosition(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func (ptr *QTapGesture) DestroyQTapGesture() {
	defer qt.Recovering("QTapGesture::~QTapGesture")

	if ptr.Pointer() != nil {
		C.QTapGesture_DestroyQTapGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTapGesture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTapGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTapGesture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTapGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTapGestureTimerEvent
func callbackQTapGestureTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTapGesture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTapGesture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTapGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTapGesture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTapGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTapGestureChildEvent
func callbackQTapGestureChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTapGesture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QTapGesture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTapGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTapGesture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTapGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTapGestureCustomEvent
func callbackQTapGestureCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTapGesture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
