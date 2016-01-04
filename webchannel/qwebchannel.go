package webchannel

//#include "webchannel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWebChannel struct {
	core.QObject
}

type QWebChannel_ITF interface {
	core.QObject_ITF
	QWebChannel_PTR() *QWebChannel
}

func PointerFromQWebChannel(ptr QWebChannel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannel_PTR().Pointer()
	}
	return nil
}

func NewQWebChannelFromPointer(ptr unsafe.Pointer) *QWebChannel {
	var n = new(QWebChannel)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWebChannel_") {
		n.SetObjectName("QWebChannel_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebChannel) QWebChannel_PTR() *QWebChannel {
	return ptr
}

func (ptr *QWebChannel) BlockUpdates() bool {
	defer qt.Recovering("QWebChannel::blockUpdates")

	if ptr.Pointer() != nil {
		return C.QWebChannel_BlockUpdates(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebChannel) SetBlockUpdates(block bool) {
	defer qt.Recovering("QWebChannel::setBlockUpdates")

	if ptr.Pointer() != nil {
		C.QWebChannel_SetBlockUpdates(ptr.Pointer(), C.int(qt.GoBoolToInt(block)))
	}
}

func NewQWebChannel(parent core.QObject_ITF) *QWebChannel {
	defer qt.Recovering("QWebChannel::QWebChannel")

	return NewQWebChannelFromPointer(C.QWebChannel_NewQWebChannel(core.PointerFromQObject(parent)))
}

func (ptr *QWebChannel) ConnectBlockUpdatesChanged(f func(block bool)) {
	defer qt.Recovering("connect QWebChannel::blockUpdatesChanged")

	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectBlockUpdatesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "blockUpdatesChanged", f)
	}
}

func (ptr *QWebChannel) DisconnectBlockUpdatesChanged() {
	defer qt.Recovering("disconnect QWebChannel::blockUpdatesChanged")

	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectBlockUpdatesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "blockUpdatesChanged")
	}
}

//export callbackQWebChannelBlockUpdatesChanged
func callbackQWebChannelBlockUpdatesChanged(ptr unsafe.Pointer, ptrName *C.char, block C.int) {
	defer qt.Recovering("callback QWebChannel::blockUpdatesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "blockUpdatesChanged"); signal != nil {
		signal.(func(bool))(int(block) != 0)
	}

}

func (ptr *QWebChannel) BlockUpdatesChanged(block bool) {
	defer qt.Recovering("QWebChannel::blockUpdatesChanged")

	if ptr.Pointer() != nil {
		C.QWebChannel_BlockUpdatesChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(block)))
	}
}

func (ptr *QWebChannel) ConnectTo(transport QWebChannelAbstractTransport_ITF) {
	defer qt.Recovering("QWebChannel::connectTo")

	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectTo(ptr.Pointer(), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannel) DeregisterObject(object core.QObject_ITF) {
	defer qt.Recovering("QWebChannel::deregisterObject")

	if ptr.Pointer() != nil {
		C.QWebChannel_DeregisterObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QWebChannel) DisconnectFrom(transport QWebChannelAbstractTransport_ITF) {
	defer qt.Recovering("QWebChannel::disconnectFrom")

	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectFrom(ptr.Pointer(), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannel) RegisterObject(id string, object core.QObject_ITF) {
	defer qt.Recovering("QWebChannel::registerObject")

	if ptr.Pointer() != nil {
		C.QWebChannel_RegisterObject(ptr.Pointer(), C.CString(id), core.PointerFromQObject(object))
	}
}

func (ptr *QWebChannel) DestroyQWebChannel() {
	defer qt.Recovering("QWebChannel::~QWebChannel")

	if ptr.Pointer() != nil {
		C.QWebChannel_DestroyQWebChannel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebChannel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebChannel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebChannel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebChannel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQWebChannelTimerEvent
func callbackQWebChannelTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebChannelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebChannel) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebChannel::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebChannel_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebChannel) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebChannel::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebChannel_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebChannel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebChannel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebChannel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebChannel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQWebChannelChildEvent
func callbackQWebChannelChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebChannelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebChannel) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebChannel::childEvent")

	if ptr.Pointer() != nil {
		C.QWebChannel_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebChannel) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebChannel::childEvent")

	if ptr.Pointer() != nil {
		C.QWebChannel_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebChannel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebChannel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebChannel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebChannel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQWebChannelCustomEvent
func callbackQWebChannelCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebChannelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebChannel) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebChannel::customEvent")

	if ptr.Pointer() != nil {
		C.QWebChannel_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebChannel) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebChannel::customEvent")

	if ptr.Pointer() != nil {
		C.QWebChannel_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
