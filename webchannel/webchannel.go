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
	return n
}

func newQWebChannelFromPointer(ptr unsafe.Pointer) *QWebChannel {
	var n = NewQWebChannelFromPointer(ptr)
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

	return newQWebChannelFromPointer(C.QWebChannel_NewQWebChannel(core.PointerFromQObject(parent)))
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

type QWebChannelAbstractTransport struct {
	core.QObject
}

type QWebChannelAbstractTransport_ITF interface {
	core.QObject_ITF
	QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport
}

func PointerFromQWebChannelAbstractTransport(ptr QWebChannelAbstractTransport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannelAbstractTransport_PTR().Pointer()
	}
	return nil
}

func NewQWebChannelAbstractTransportFromPointer(ptr unsafe.Pointer) *QWebChannelAbstractTransport {
	var n = new(QWebChannelAbstractTransport)
	n.SetPointer(ptr)
	return n
}

func newQWebChannelAbstractTransportFromPointer(ptr unsafe.Pointer) *QWebChannelAbstractTransport {
	var n = NewQWebChannelAbstractTransportFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebChannelAbstractTransport_") {
		n.SetObjectName("QWebChannelAbstractTransport_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebChannelAbstractTransport) QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport {
	return ptr
}

func (ptr *QWebChannelAbstractTransport) ConnectMessageReceived(f func(message *core.QJsonObject, transport *QWebChannelAbstractTransport)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ConnectMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageReceived", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectMessageReceived() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DisconnectMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageReceived")
	}
}

//export callbackQWebChannelAbstractTransportMessageReceived
func callbackQWebChannelAbstractTransportMessageReceived(ptr unsafe.Pointer, ptrName *C.char, message unsafe.Pointer, transport unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::messageReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "messageReceived"); signal != nil {
		signal.(func(*core.QJsonObject, *QWebChannelAbstractTransport))(core.NewQJsonObjectFromPointer(message), NewQWebChannelAbstractTransportFromPointer(transport))
	}

}

func (ptr *QWebChannelAbstractTransport) MessageReceived(message core.QJsonObject_ITF, transport QWebChannelAbstractTransport_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_MessageReceived(ptr.Pointer(), core.PointerFromQJsonObject(message), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannelAbstractTransport) SendMessage(message core.QJsonObject_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::sendMessage")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_SendMessage(ptr.Pointer(), core.PointerFromQJsonObject(message))
	}
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransport() {
	defer qt.Recovering("QWebChannelAbstractTransport::~QWebChannelAbstractTransport")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQWebChannelAbstractTransportTimerEvent
func callbackQWebChannelAbstractTransportTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQWebChannelAbstractTransportChildEvent
func callbackQWebChannelAbstractTransportChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQWebChannelAbstractTransportCustomEvent
func callbackQWebChannelAbstractTransportCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
