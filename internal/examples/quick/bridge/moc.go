package main

//#include "moc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QmlBridge_ITF interface {
	core.QObject_ITF
	QmlBridge_PTR() *QmlBridge
}

func (p *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return p
}

func (p *QmlBridge) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QmlBridge) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQmlBridge(ptr QmlBridge_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QmlBridge_PTR().Pointer()
	}
	return nil
}

func NewQmlBridgeFromPointer(ptr unsafe.Pointer) *QmlBridge {
	var n = new(QmlBridge)
	n.SetPointer(ptr)
	return n
}

func newQmlBridgeFromPointer(ptr unsafe.Pointer) *QmlBridge {
	var n = NewQmlBridgeFromPointer(ptr)
	for len(n.ObjectName()) < len("QmlBridge_") {
		n.SetObjectName("QmlBridge_" + qt.Identifier())
	}
	return n
}

//export callbackQmlBridge_SendToQml
func callbackQmlBridge_SendToQml(ptr unsafe.Pointer, ptrName *C.char, source *C.char, action *C.char, data *C.char) {
	defer qt.Recovering("callback QmlBridge::sendToQml")

	if signal := qt.GetSignal(C.GoString(ptrName), "sendToQml"); signal != nil {
		signal.(func(string, string, string))(C.GoString(source), C.GoString(action), C.GoString(data))
	}

}

func (ptr *QmlBridge) ConnectSendToQml(f func(source string, action string, data string)) {
	defer qt.Recovering("connect QmlBridge::sendToQml")

	if ptr.Pointer() != nil {
		C.QmlBridge_ConnectSendToQml(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sendToQml", f)
	}
}

func (ptr *QmlBridge) DisconnectSendToQml() {
	defer qt.Recovering("disconnect QmlBridge::sendToQml")

	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectSendToQml(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sendToQml")
	}
}

func (ptr *QmlBridge) SendToQml(source string, action string, data string) {
	defer qt.Recovering("QmlBridge::sendToQml")

	if ptr.Pointer() != nil {
		C.QmlBridge_SendToQml(ptr.Pointer(), C.CString(source), C.CString(action), C.CString(data))
	}
}

//export callbackQmlBridge_SendToGo
func callbackQmlBridge_SendToGo(ptr unsafe.Pointer, ptrName *C.char, source *C.char, action *C.char, data *C.char) {
	defer qt.Recovering("callback QmlBridge::sendToGo")

	if signal := qt.GetSignal(C.GoString(ptrName), "sendToGo"); signal != nil {
		signal.(func(string, string, string))(C.GoString(source), C.GoString(action), C.GoString(data))
	}

}

func (ptr *QmlBridge) ConnectSendToGo(f func(source string, action string, data string)) {
	defer qt.Recovering("connect QmlBridge::sendToGo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sendToGo", f)
	}
}

func (ptr *QmlBridge) DisconnectSendToGo(source string, action string, data string) {
	defer qt.Recovering("disconnect QmlBridge::sendToGo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sendToGo")
	}
}

func (ptr *QmlBridge) SendToGo(source string, action string, data string) {
	defer qt.Recovering("QmlBridge::sendToGo")

	if ptr.Pointer() != nil {
		C.QmlBridge_SendToGo(ptr.Pointer(), C.CString(source), C.CString(action), C.CString(data))
	}
}

//export callbackQmlBridge_RegisterToGo
func callbackQmlBridge_RegisterToGo(ptr unsafe.Pointer, ptrName *C.char, object unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::registerToGo")

	if signal := qt.GetSignal(C.GoString(ptrName), "registerToGo"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(object))
	}

}

func (ptr *QmlBridge) ConnectRegisterToGo(f func(object *core.QObject)) {
	defer qt.Recovering("connect QmlBridge::registerToGo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "registerToGo", f)
	}
}

func (ptr *QmlBridge) DisconnectRegisterToGo(object core.QObject_ITF) {
	defer qt.Recovering("disconnect QmlBridge::registerToGo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "registerToGo")
	}
}

func (ptr *QmlBridge) RegisterToGo(object core.QObject_ITF) {
	defer qt.Recovering("QmlBridge::registerToGo")

	if ptr.Pointer() != nil {
		C.QmlBridge_RegisterToGo(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

//export callbackQmlBridge_DeregisterToGo
func callbackQmlBridge_DeregisterToGo(ptr unsafe.Pointer, ptrName *C.char, objectName *C.char) {
	defer qt.Recovering("callback QmlBridge::deregisterToGo")

	if signal := qt.GetSignal(C.GoString(ptrName), "deregisterToGo"); signal != nil {
		signal.(func(string))(C.GoString(objectName))
	}

}

func (ptr *QmlBridge) ConnectDeregisterToGo(f func(objectName string)) {
	defer qt.Recovering("connect QmlBridge::deregisterToGo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deregisterToGo", f)
	}
}

func (ptr *QmlBridge) DisconnectDeregisterToGo(objectName string) {
	defer qt.Recovering("disconnect QmlBridge::deregisterToGo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deregisterToGo")
	}
}

func (ptr *QmlBridge) DeregisterToGo(objectName string) {
	defer qt.Recovering("QmlBridge::deregisterToGo")

	if ptr.Pointer() != nil {
		C.QmlBridge_DeregisterToGo(ptr.Pointer(), C.CString(objectName))
	}
}

func NewQmlBridge(parent core.QObject_ITF) *QmlBridge {
	defer qt.Recovering("QmlBridge::QmlBridge")

	return newQmlBridgeFromPointer(C.QmlBridge_NewQmlBridge(core.PointerFromQObject(parent)))
}

func (ptr *QmlBridge) DestroyQmlBridge() {
	defer qt.Recovering("QmlBridge::~QmlBridge")

	if ptr.Pointer() != nil {
		C.QmlBridge_DestroyQmlBridge(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQmlBridge_TimerEvent
func callbackQmlBridge_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QmlBridge::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QmlBridge) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QmlBridge::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QmlBridge) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QmlBridge::timerEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QmlBridge) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QmlBridge::timerEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQmlBridge_ChildEvent
func callbackQmlBridge_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QmlBridge::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QmlBridge) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QmlBridge::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QmlBridge) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QmlBridge::childEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QmlBridge) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QmlBridge::childEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQmlBridge_ConnectNotify
func callbackQmlBridge_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QmlBridge::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QmlBridge) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QmlBridge::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QmlBridge) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QmlBridge::connectNotify")

	if ptr.Pointer() != nil {
		C.QmlBridge_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QmlBridge) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QmlBridge::connectNotify")

	if ptr.Pointer() != nil {
		C.QmlBridge_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge_CustomEvent
func callbackQmlBridge_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QmlBridge) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QmlBridge::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QmlBridge) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QmlBridge::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QmlBridge) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QmlBridge::customEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QmlBridge) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QmlBridge::customEvent")

	if ptr.Pointer() != nil {
		C.QmlBridge_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQmlBridge_DeleteLater
func callbackQmlBridge_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QmlBridge::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQmlBridgeFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QmlBridge) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QmlBridge::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QmlBridge) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QmlBridge::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QmlBridge) DeleteLater() {
	defer qt.Recovering("QmlBridge::deleteLater")

	if ptr.Pointer() != nil {
		C.QmlBridge_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QmlBridge) DeleteLaterDefault() {
	defer qt.Recovering("QmlBridge::deleteLater")

	if ptr.Pointer() != nil {
		C.QmlBridge_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQmlBridge_DisconnectNotify
func callbackQmlBridge_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQmlBridgeFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QmlBridge) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QmlBridge::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QmlBridge) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QmlBridge::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QmlBridge) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QmlBridge::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QmlBridge) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QmlBridge::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QmlBridge_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQmlBridge_Event
func callbackQmlBridge_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QmlBridge::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QmlBridge) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QmlBridge::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QmlBridge) DisconnectEvent() {
	defer qt.Recovering("disconnect QmlBridge::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QmlBridge) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QmlBridge::event")

	if ptr.Pointer() != nil {
		return C.QmlBridge_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QmlBridge) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QmlBridge::event")

	if ptr.Pointer() != nil {
		return C.QmlBridge_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQmlBridge_EventFilter
func callbackQmlBridge_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QmlBridge::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQmlBridgeFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QmlBridge) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QmlBridge::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QmlBridge) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QmlBridge::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QmlBridge) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QmlBridge::eventFilter")

	if ptr.Pointer() != nil {
		return C.QmlBridge_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QmlBridge) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QmlBridge::eventFilter")

	if ptr.Pointer() != nil {
		return C.QmlBridge_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQmlBridge_MetaObject
