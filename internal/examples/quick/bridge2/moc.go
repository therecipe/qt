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

func (ptr *QmlBridge) QmlBridge_PTR() *QmlBridge {
	return ptr
}

func (ptr *QmlBridge) ConnectSendToQml(f func(data string)) {
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

//export callbackQmlBridgeSendToQml
func callbackQmlBridgeSendToQml(ptr unsafe.Pointer, ptrName *C.char, data *C.char) {
	defer qt.Recovering("callback QmlBridge::sendToQml")

	if signal := qt.GetSignal(C.GoString(ptrName), "sendToQml"); signal != nil {
		signal.(func(string))(C.GoString(data))
	}

}

func (ptr *QmlBridge) SendToQml(data string) {
	defer qt.Recovering("QmlBridge::sendToQml")

	if ptr.Pointer() != nil {
		C.QmlBridge_SendToQml(ptr.Pointer(), C.CString(data))
	}
}

func (ptr *QmlBridge) ConnectSendToGo(f func(data string)) {
	defer qt.Recovering("connect QmlBridge::sendToGo")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sendToGo", f)
	}
}

func (ptr *QmlBridge) DisconnectSendToGo() {
	defer qt.Recovering("disconnect QmlBridge::sendToGo")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sendToGo")
	}
}

//export callbackQmlBridgeSendToGo
func callbackQmlBridgeSendToGo(ptr unsafe.Pointer, ptrName *C.char, data *C.char) {
	defer qt.Recovering("callback QmlBridge::sendToGo")

	if signal := qt.GetSignal(C.GoString(ptrName), "sendToGo"); signal != nil {
		signal.(func(string))(C.GoString(data))
	}

}

func (ptr *QmlBridge) SendToGo(data string) {
	defer qt.Recovering("QmlBridge::sendToGo")

	if ptr.Pointer() != nil {
		C.QmlBridge_SendToGo(ptr.Pointer(), C.CString(data))
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

//export callbackQmlBridgeTimerEvent
func callbackQmlBridgeTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQmlBridgeChildEvent
func callbackQmlBridgeChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQmlBridgeCustomEvent
func callbackQmlBridgeCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QmlBridge::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQmlBridgeFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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
