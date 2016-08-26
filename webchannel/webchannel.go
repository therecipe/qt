// +build !minimal

package webchannel

//#include <stdint.h>
//#include <stdlib.h>
//#include "webchannel.h"
import "C"
import (
	"fmt"
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

func (p *QWebChannel) QWebChannel_PTR() *QWebChannel {
	return p
}

func (p *QWebChannel) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebChannel) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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
		C.QWebChannel_SetBlockUpdates(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(block))))
	}
}

func NewQWebChannel(parent core.QObject_ITF) *QWebChannel {
	defer qt.Recovering("QWebChannel::QWebChannel")

	var tmpValue = NewQWebChannelFromPointer(C.QWebChannel_NewQWebChannel(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebChannel_BlockUpdatesChanged
func callbackQWebChannel_BlockUpdatesChanged(ptr unsafe.Pointer, block C.char) {
	defer qt.Recovering("callback QWebChannel::blockUpdatesChanged")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::blockUpdatesChanged"); signal != nil {
		signal.(func(bool))(int8(block) != 0)
	}

}

func (ptr *QWebChannel) ConnectBlockUpdatesChanged(f func(block bool)) {
	defer qt.Recovering("connect QWebChannel::blockUpdatesChanged")

	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectBlockUpdatesChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::blockUpdatesChanged", f)
	}
}

func (ptr *QWebChannel) DisconnectBlockUpdatesChanged() {
	defer qt.Recovering("disconnect QWebChannel::blockUpdatesChanged")

	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectBlockUpdatesChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::blockUpdatesChanged")
	}
}

func (ptr *QWebChannel) BlockUpdatesChanged(block bool) {
	defer qt.Recovering("QWebChannel::blockUpdatesChanged")

	if ptr.Pointer() != nil {
		C.QWebChannel_BlockUpdatesChanged(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(block))))
	}
}

//export callbackQWebChannel_ConnectTo
func callbackQWebChannel_ConnectTo(ptr unsafe.Pointer, transport unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::connectTo")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::connectTo"); signal != nil {
		signal.(func(*QWebChannelAbstractTransport))(NewQWebChannelAbstractTransportFromPointer(transport))
	}

}

func (ptr *QWebChannel) ConnectConnectTo(f func(transport *QWebChannelAbstractTransport)) {
	defer qt.Recovering("connect QWebChannel::connectTo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::connectTo", f)
	}
}

func (ptr *QWebChannel) DisconnectConnectTo(transport QWebChannelAbstractTransport_ITF) {
	defer qt.Recovering("disconnect QWebChannel::connectTo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::connectTo")
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

//export callbackQWebChannel_DisconnectFrom
func callbackQWebChannel_DisconnectFrom(ptr unsafe.Pointer, transport unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::disconnectFrom")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::disconnectFrom"); signal != nil {
		signal.(func(*QWebChannelAbstractTransport))(NewQWebChannelAbstractTransportFromPointer(transport))
	}

}

func (ptr *QWebChannel) ConnectDisconnectFrom(f func(transport *QWebChannelAbstractTransport)) {
	defer qt.Recovering("connect QWebChannel::disconnectFrom")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::disconnectFrom", f)
	}
}

func (ptr *QWebChannel) DisconnectDisconnectFrom(transport QWebChannelAbstractTransport_ITF) {
	defer qt.Recovering("disconnect QWebChannel::disconnectFrom")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::disconnectFrom")
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
		var idC = C.CString(id)
		defer C.free(unsafe.Pointer(idC))
		C.QWebChannel_RegisterObject(ptr.Pointer(), idC, core.PointerFromQObject(object))
	}
}

func (ptr *QWebChannel) DestroyQWebChannel() {
	defer qt.Recovering("QWebChannel::~QWebChannel")

	if ptr.Pointer() != nil {
		C.QWebChannel_DestroyQWebChannel(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebChannel_TimerEvent
func callbackQWebChannel_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebChannelFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebChannel) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebChannel::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::timerEvent", f)
	}
}

func (ptr *QWebChannel) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebChannel::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::timerEvent")
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

//export callbackQWebChannel_ChildEvent
func callbackQWebChannel_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebChannelFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebChannel) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebChannel::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::childEvent", f)
	}
}

func (ptr *QWebChannel) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebChannel::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::childEvent")
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

//export callbackQWebChannel_ConnectNotify
func callbackQWebChannel_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebChannelFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebChannel) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebChannel::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::connectNotify", f)
	}
}

func (ptr *QWebChannel) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebChannel::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::connectNotify")
	}
}

func (ptr *QWebChannel) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebChannel::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebChannel) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebChannel::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebChannel_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebChannel_CustomEvent
func callbackQWebChannel_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebChannelFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebChannel) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebChannel::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::customEvent", f)
	}
}

func (ptr *QWebChannel) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebChannel::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::customEvent")
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

//export callbackQWebChannel_DeleteLater
func callbackQWebChannel_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebChannelFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebChannel) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebChannel::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::deleteLater", f)
	}
}

func (ptr *QWebChannel) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebChannel::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::deleteLater")
	}
}

func (ptr *QWebChannel) DeleteLater() {
	defer qt.Recovering("QWebChannel::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebChannel_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebChannel) DeleteLaterDefault() {
	defer qt.Recovering("QWebChannel::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebChannel_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebChannel_DisconnectNotify
func callbackQWebChannel_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannel::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebChannelFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebChannel) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebChannel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::disconnectNotify", f)
	}
}

func (ptr *QWebChannel) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebChannel::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::disconnectNotify")
	}
}

func (ptr *QWebChannel) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebChannel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebChannel) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebChannel::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebChannel_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebChannel_Event
func callbackQWebChannel_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QWebChannel::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebChannelFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebChannel) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebChannel::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::event", f)
	}
}

func (ptr *QWebChannel) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebChannel::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::event")
	}
}

func (ptr *QWebChannel) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebChannel::event")

	if ptr.Pointer() != nil {
		return C.QWebChannel_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebChannel) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebChannel::event")

	if ptr.Pointer() != nil {
		return C.QWebChannel_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebChannel_EventFilter
func callbackQWebChannel_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QWebChannel::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebChannelFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebChannel) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebChannel::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::eventFilter", f)
	}
}

func (ptr *QWebChannel) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebChannel::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::eventFilter")
	}
}

func (ptr *QWebChannel) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebChannel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebChannel_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebChannel) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebChannel::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebChannel_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebChannel_MetaObject
func callbackQWebChannel_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QWebChannel::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannel::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebChannelFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebChannel) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebChannel::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::metaObject", f)
	}
}

func (ptr *QWebChannel) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebChannel::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannel::metaObject")
	}
}

func (ptr *QWebChannel) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebChannel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebChannel_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebChannel) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebChannel::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebChannel_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebChannelAbstractTransport struct {
	core.QObject
}

type QWebChannelAbstractTransport_ITF interface {
	core.QObject_ITF
	QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport
}

func (p *QWebChannelAbstractTransport) QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport {
	return p
}

func (p *QWebChannelAbstractTransport) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebChannelAbstractTransport) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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
func NewQWebChannelAbstractTransport(parent core.QObject_ITF) *QWebChannelAbstractTransport {
	defer qt.Recovering("QWebChannelAbstractTransport::QWebChannelAbstractTransport")

	var tmpValue = NewQWebChannelAbstractTransportFromPointer(C.QWebChannelAbstractTransport_NewQWebChannelAbstractTransport(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebChannelAbstractTransport_MessageReceived
func callbackQWebChannelAbstractTransport_MessageReceived(ptr unsafe.Pointer, message unsafe.Pointer, transport unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::messageReceived")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::messageReceived"); signal != nil {
		signal.(func(*core.QJsonObject, *QWebChannelAbstractTransport))(core.NewQJsonObjectFromPointer(message), NewQWebChannelAbstractTransportFromPointer(transport))
	}

}

func (ptr *QWebChannelAbstractTransport) ConnectMessageReceived(f func(message *core.QJsonObject, transport *QWebChannelAbstractTransport)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ConnectMessageReceived(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::messageReceived", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectMessageReceived() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DisconnectMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::messageReceived")
	}
}

func (ptr *QWebChannelAbstractTransport) MessageReceived(message core.QJsonObject_ITF, transport QWebChannelAbstractTransport_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_MessageReceived(ptr.Pointer(), core.PointerFromQJsonObject(message), PointerFromQWebChannelAbstractTransport(transport))
	}
}

//export callbackQWebChannelAbstractTransport_SendMessage
func callbackQWebChannelAbstractTransport_SendMessage(ptr unsafe.Pointer, message unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::sendMessage")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::sendMessage"); signal != nil {
		signal.(func(*core.QJsonObject))(core.NewQJsonObjectFromPointer(message))
	}

}

func (ptr *QWebChannelAbstractTransport) ConnectSendMessage(f func(message *core.QJsonObject)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::sendMessage")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::sendMessage", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectSendMessage(message core.QJsonObject_ITF) {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::sendMessage")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::sendMessage")
	}
}

func (ptr *QWebChannelAbstractTransport) SendMessage(message core.QJsonObject_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::sendMessage")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_SendMessage(ptr.Pointer(), core.PointerFromQJsonObject(message))
	}
}

//export callbackQWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport
func callbackQWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::~QWebChannelAbstractTransport")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::~QWebChannelAbstractTransport"); signal != nil {
		signal.(func())()
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).DestroyQWebChannelAbstractTransportDefault()
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectDestroyQWebChannelAbstractTransport(f func()) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::~QWebChannelAbstractTransport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::~QWebChannelAbstractTransport", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectDestroyQWebChannelAbstractTransport() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::~QWebChannelAbstractTransport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::~QWebChannelAbstractTransport")
	}
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransport() {
	defer qt.Recovering("QWebChannelAbstractTransport::~QWebChannelAbstractTransport")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransportDefault() {
	defer qt.Recovering("QWebChannelAbstractTransport::~QWebChannelAbstractTransport")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransportDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebChannelAbstractTransport_TimerEvent
func callbackQWebChannelAbstractTransport_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::timerEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::timerEvent", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::timerEvent")
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

//export callbackQWebChannelAbstractTransport_ChildEvent
func callbackQWebChannelAbstractTransport_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::childEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::childEvent", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::childEvent")
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

//export callbackQWebChannelAbstractTransport_ConnectNotify
func callbackQWebChannelAbstractTransport_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::connectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::connectNotify", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::connectNotify")
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebChannelAbstractTransport_CustomEvent
func callbackQWebChannelAbstractTransport_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::customEvent")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::customEvent", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::customEvent")
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

//export callbackQWebChannelAbstractTransport_DeleteLater
func callbackQWebChannelAbstractTransport_DeleteLater(ptr unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::deleteLater")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::deleteLater", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::deleteLater")
	}
}

func (ptr *QWebChannelAbstractTransport) DeleteLater() {
	defer qt.Recovering("QWebChannelAbstractTransport::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebChannelAbstractTransport) DeleteLaterDefault() {
	defer qt.Recovering("QWebChannelAbstractTransport::deleteLater")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebChannelAbstractTransport_DisconnectNotify
func callbackQWebChannelAbstractTransport_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::disconnectNotify")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::disconnectNotify", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::disconnectNotify")
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebChannelAbstractTransport_Event
func callbackQWebChannelAbstractTransport_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	defer qt.Recovering("callback QWebChannelAbstractTransport::event")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebChannelAbstractTransportFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebChannelAbstractTransport) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::event", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::event")
	}
}

func (ptr *QWebChannelAbstractTransport) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebChannelAbstractTransport::event")

	if ptr.Pointer() != nil {
		return C.QWebChannelAbstractTransport_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebChannelAbstractTransport) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebChannelAbstractTransport::event")

	if ptr.Pointer() != nil {
		return C.QWebChannelAbstractTransport_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebChannelAbstractTransport_EventFilter
func callbackQWebChannelAbstractTransport_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	defer qt.Recovering("callback QWebChannelAbstractTransport::eventFilter")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebChannelAbstractTransportFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebChannelAbstractTransport) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::eventFilter", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::eventFilter")
	}
}

func (ptr *QWebChannelAbstractTransport) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebChannelAbstractTransport::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebChannelAbstractTransport_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebChannelAbstractTransport) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebChannelAbstractTransport::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebChannelAbstractTransport_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebChannelAbstractTransport_MetaObject
func callbackQWebChannelAbstractTransport_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QWebChannelAbstractTransport::metaObject")

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebChannelAbstractTransport::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebChannelAbstractTransportFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebChannelAbstractTransport) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::metaObject", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebChannelAbstractTransport::metaObject")
	}
}

func (ptr *QWebChannelAbstractTransport) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebChannelAbstractTransport::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebChannelAbstractTransport_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebChannelAbstractTransport) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebChannelAbstractTransport::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebChannelAbstractTransport_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
