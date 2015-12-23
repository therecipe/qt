package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMediaVideoProbeControl struct {
	QMediaControl
}

type QMediaVideoProbeControl_ITF interface {
	QMediaControl_ITF
	QMediaVideoProbeControl_PTR() *QMediaVideoProbeControl
}

func PointerFromQMediaVideoProbeControl(ptr QMediaVideoProbeControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaVideoProbeControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaVideoProbeControlFromPointer(ptr unsafe.Pointer) *QMediaVideoProbeControl {
	var n = new(QMediaVideoProbeControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaVideoProbeControl_") {
		n.SetObjectName("QMediaVideoProbeControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaVideoProbeControl) QMediaVideoProbeControl_PTR() *QMediaVideoProbeControl {
	return ptr
}

func (ptr *QMediaVideoProbeControl) ConnectFlush(f func()) {
	defer qt.Recovering("connect QMediaVideoProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_ConnectFlush(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "flush", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectFlush() {
	defer qt.Recovering("disconnect QMediaVideoProbeControl::flush")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DisconnectFlush(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "flush")
	}
}

//export callbackQMediaVideoProbeControlFlush
func callbackQMediaVideoProbeControlFlush(ptrName *C.char) {
	defer qt.Recovering("callback QMediaVideoProbeControl::flush")

	if signal := qt.GetSignal(C.GoString(ptrName), "flush"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMediaVideoProbeControl) DestroyQMediaVideoProbeControl() {
	defer qt.Recovering("QMediaVideoProbeControl::~QMediaVideoProbeControl")

	if ptr.Pointer() != nil {
		C.QMediaVideoProbeControl_DestroyQMediaVideoProbeControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMediaVideoProbeControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMediaVideoProbeControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMediaVideoProbeControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMediaVideoProbeControlTimerEvent
func callbackQMediaVideoProbeControlTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaVideoProbeControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaVideoProbeControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMediaVideoProbeControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMediaVideoProbeControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMediaVideoProbeControlChildEvent
func callbackQMediaVideoProbeControlChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaVideoProbeControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMediaVideoProbeControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMediaVideoProbeControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMediaVideoProbeControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMediaVideoProbeControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMediaVideoProbeControlCustomEvent
func callbackQMediaVideoProbeControlCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMediaVideoProbeControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
