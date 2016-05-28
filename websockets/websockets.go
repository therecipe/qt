// +build !minimal

package websockets

//#include "websockets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QMaskGenerator struct {
	core.QObject
}

type QMaskGenerator_ITF interface {
	core.QObject_ITF
	QMaskGenerator_PTR() *QMaskGenerator
}

func (p *QMaskGenerator) QMaskGenerator_PTR() *QMaskGenerator {
	return p
}

func (p *QMaskGenerator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QMaskGenerator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQMaskGenerator(ptr QMaskGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMaskGenerator_PTR().Pointer()
	}
	return nil
}

func NewQMaskGeneratorFromPointer(ptr unsafe.Pointer) *QMaskGenerator {
	var n = new(QMaskGenerator)
	n.SetPointer(ptr)
	return n
}

func newQMaskGeneratorFromPointer(ptr unsafe.Pointer) *QMaskGenerator {
	var n = NewQMaskGeneratorFromPointer(ptr)
	for len(n.ObjectName()) < len("QMaskGenerator_") {
		n.SetObjectName("QMaskGenerator_" + qt.Identifier())
	}
	return n
}

//export callbackQMaskGenerator_Seed
func callbackQMaskGenerator_Seed(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QMaskGenerator::seed")

	if signal := qt.GetSignal(C.GoString(ptrName), "seed"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QMaskGenerator) ConnectSeed(f func() bool) {
	defer qt.Recovering("connect QMaskGenerator::seed")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "seed", f)
	}
}

func (ptr *QMaskGenerator) DisconnectSeed() {
	defer qt.Recovering("disconnect QMaskGenerator::seed")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "seed")
	}
}

func (ptr *QMaskGenerator) Seed() bool {
	defer qt.Recovering("QMaskGenerator::seed")

	if ptr.Pointer() != nil {
		return C.QMaskGenerator_Seed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMaskGenerator) DestroyQMaskGenerator() {
	defer qt.Recovering("QMaskGenerator::~QMaskGenerator")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QMaskGenerator_DestroyQMaskGenerator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQMaskGenerator_TimerEvent
func callbackQMaskGenerator_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QMaskGenerator) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMaskGenerator) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQMaskGenerator_ChildEvent
func callbackQMaskGenerator_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QMaskGenerator) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMaskGenerator) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMaskGenerator_ConnectNotify
func callbackQMaskGenerator_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMaskGeneratorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMaskGenerator) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QMaskGenerator::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QMaskGenerator) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QMaskGenerator::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QMaskGenerator) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMaskGenerator::connectNotify")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMaskGenerator) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMaskGenerator::connectNotify")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMaskGenerator_CustomEvent
func callbackQMaskGenerator_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QMaskGenerator) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMaskGenerator) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMaskGenerator_DeleteLater
func callbackQMaskGenerator_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMaskGenerator::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMaskGeneratorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMaskGenerator) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QMaskGenerator::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QMaskGenerator) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QMaskGenerator::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QMaskGenerator) DeleteLater() {
	defer qt.Recovering("QMaskGenerator::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QMaskGenerator_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMaskGenerator) DeleteLaterDefault() {
	defer qt.Recovering("QMaskGenerator::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QMaskGenerator_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQMaskGenerator_DisconnectNotify
func callbackQMaskGenerator_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMaskGeneratorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMaskGenerator) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QMaskGenerator::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QMaskGenerator) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QMaskGenerator::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QMaskGenerator) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMaskGenerator::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMaskGenerator) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QMaskGenerator::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMaskGenerator_Event
func callbackQMaskGenerator_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QMaskGenerator::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQMaskGeneratorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QMaskGenerator) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QMaskGenerator::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QMaskGenerator) DisconnectEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QMaskGenerator) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QMaskGenerator::event")

	if ptr.Pointer() != nil {
		return C.QMaskGenerator_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMaskGenerator) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QMaskGenerator::event")

	if ptr.Pointer() != nil {
		return C.QMaskGenerator_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQMaskGenerator_EventFilter
func callbackQMaskGenerator_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QMaskGenerator::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQMaskGeneratorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QMaskGenerator) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QMaskGenerator::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QMaskGenerator) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QMaskGenerator::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QMaskGenerator) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QMaskGenerator::eventFilter")

	if ptr.Pointer() != nil {
		return C.QMaskGenerator_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMaskGenerator) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QMaskGenerator::eventFilter")

	if ptr.Pointer() != nil {
		return C.QMaskGenerator_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQMaskGenerator_MetaObject
func callbackQMaskGenerator_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QMaskGenerator::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMaskGeneratorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMaskGenerator) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QMaskGenerator::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QMaskGenerator) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QMaskGenerator::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QMaskGenerator) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QMaskGenerator::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMaskGenerator_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMaskGenerator) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QMaskGenerator::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMaskGenerator_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebSocket struct {
	core.QObject
}

type QWebSocket_ITF interface {
	core.QObject_ITF
	QWebSocket_PTR() *QWebSocket
}

func (p *QWebSocket) QWebSocket_PTR() *QWebSocket {
	return p
}

func (p *QWebSocket) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebSocket) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQWebSocket(ptr QWebSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocket_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketFromPointer(ptr unsafe.Pointer) *QWebSocket {
	var n = new(QWebSocket)
	n.SetPointer(ptr)
	return n
}

func newQWebSocketFromPointer(ptr unsafe.Pointer) *QWebSocket {
	var n = NewQWebSocketFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebSocket_") {
		n.SetObjectName("QWebSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebSocket) Abort() {
	defer qt.Recovering("QWebSocket::abort")

	if ptr.Pointer() != nil {
		C.QWebSocket_Abort(ptr.Pointer())
	}
}

//export callbackQWebSocket_AboutToClose
func callbackQWebSocket_AboutToClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::aboutToClose")

	if signal := qt.GetSignal(C.GoString(ptrName), "aboutToClose"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectAboutToClose(f func()) {
	defer qt.Recovering("connect QWebSocket::aboutToClose")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectAboutToClose(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToClose", f)
	}
}

func (ptr *QWebSocket) DisconnectAboutToClose() {
	defer qt.Recovering("disconnect QWebSocket::aboutToClose")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectAboutToClose(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToClose")
	}
}

func (ptr *QWebSocket) AboutToClose() {
	defer qt.Recovering("QWebSocket::aboutToClose")

	if ptr.Pointer() != nil {
		C.QWebSocket_AboutToClose(ptr.Pointer())
	}
}

//export callbackQWebSocket_BinaryFrameReceived
func callbackQWebSocket_BinaryFrameReceived(ptr unsafe.Pointer, ptrName *C.char, frame *C.char, isLastFrame C.int) {
	defer qt.Recovering("callback QWebSocket::binaryFrameReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "binaryFrameReceived"); signal != nil {
		signal.(func(string, bool))(C.GoString(frame), int(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectBinaryFrameReceived(f func(frame string, isLastFrame bool)) {
	defer qt.Recovering("connect QWebSocket::binaryFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryFrameReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "binaryFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryFrameReceived() {
	defer qt.Recovering("disconnect QWebSocket::binaryFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "binaryFrameReceived")
	}
}

func (ptr *QWebSocket) BinaryFrameReceived(frame string, isLastFrame bool) {
	defer qt.Recovering("QWebSocket::binaryFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_BinaryFrameReceived(ptr.Pointer(), C.CString(frame), C.int(qt.GoBoolToInt(isLastFrame)))
	}
}

//export callbackQWebSocket_BinaryMessageReceived
func callbackQWebSocket_BinaryMessageReceived(ptr unsafe.Pointer, ptrName *C.char, message *C.char) {
	defer qt.Recovering("callback QWebSocket::binaryMessageReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "binaryMessageReceived"); signal != nil {
		signal.(func(string))(C.GoString(message))
	}

}

func (ptr *QWebSocket) ConnectBinaryMessageReceived(f func(message string)) {
	defer qt.Recovering("connect QWebSocket::binaryMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "binaryMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryMessageReceived() {
	defer qt.Recovering("disconnect QWebSocket::binaryMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "binaryMessageReceived")
	}
}

func (ptr *QWebSocket) BinaryMessageReceived(message string) {
	defer qt.Recovering("QWebSocket::binaryMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_BinaryMessageReceived(ptr.Pointer(), C.CString(message))
	}
}

//export callbackQWebSocket_BytesWritten
func callbackQWebSocket_BytesWritten(ptr unsafe.Pointer, ptrName *C.char, bytes C.longlong) {
	defer qt.Recovering("callback QWebSocket::bytesWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "bytesWritten"); signal != nil {
		signal.(func(int64))(int64(bytes))
	}

}

func (ptr *QWebSocket) ConnectBytesWritten(f func(bytes int64)) {
	defer qt.Recovering("connect QWebSocket::bytesWritten")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBytesWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bytesWritten", f)
	}
}

func (ptr *QWebSocket) DisconnectBytesWritten() {
	defer qt.Recovering("disconnect QWebSocket::bytesWritten")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBytesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bytesWritten")
	}
}

func (ptr *QWebSocket) BytesWritten(bytes int64) {
	defer qt.Recovering("QWebSocket::bytesWritten")

	if ptr.Pointer() != nil {
		C.QWebSocket_BytesWritten(ptr.Pointer(), C.longlong(bytes))
	}
}

func (ptr *QWebSocket) CloseReason() string {
	defer qt.Recovering("QWebSocket::closeReason")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_CloseReason(ptr.Pointer()))
	}
	return ""
}

//export callbackQWebSocket_Connected
func callbackQWebSocket_Connected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QWebSocket::connected")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QWebSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QWebSocket::connected")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

func (ptr *QWebSocket) Connected() {
	defer qt.Recovering("QWebSocket::connected")

	if ptr.Pointer() != nil {
		C.QWebSocket_Connected(ptr.Pointer())
	}
}

//export callbackQWebSocket_Disconnected
func callbackQWebSocket_Disconnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QWebSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QWebSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QWebSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

func (ptr *QWebSocket) Disconnected() {
	defer qt.Recovering("QWebSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QWebSocket_Disconnected(ptr.Pointer())
	}
}

//export callbackQWebSocket_Error2
func callbackQWebSocket_Error2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QWebSocket::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(error))
	}

}

func (ptr *QWebSocket) ConnectError2(f func(error network.QAbstractSocket__SocketError)) {
	defer qt.Recovering("connect QWebSocket::error")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QWebSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QWebSocket::error")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

func (ptr *QWebSocket) Error2(error network.QAbstractSocket__SocketError) {
	defer qt.Recovering("QWebSocket::error")

	if ptr.Pointer() != nil {
		C.QWebSocket_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QWebSocket) Error() network.QAbstractSocket__SocketError {
	defer qt.Recovering("QWebSocket::error")

	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketError(C.QWebSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) ErrorString() string {
	defer qt.Recovering("QWebSocket::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Flush() bool {
	defer qt.Recovering("QWebSocket::flush")

	if ptr.Pointer() != nil {
		return C.QWebSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebSocket_IgnoreSslErrors
func callbackQWebSocket_IgnoreSslErrors(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::ignoreSslErrors")

	if signal := qt.GetSignal(C.GoString(ptrName), "ignoreSslErrors"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectIgnoreSslErrors(f func()) {
	defer qt.Recovering("connect QWebSocket::ignoreSslErrors")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "ignoreSslErrors", f)
	}
}

func (ptr *QWebSocket) DisconnectIgnoreSslErrors() {
	defer qt.Recovering("disconnect QWebSocket::ignoreSslErrors")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "ignoreSslErrors")
	}
}

func (ptr *QWebSocket) IgnoreSslErrors() {
	defer qt.Recovering("QWebSocket::ignoreSslErrors")

	if ptr.Pointer() != nil {
		C.QWebSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QWebSocket) IsValid() bool {
	defer qt.Recovering("QWebSocket::isValid")

	if ptr.Pointer() != nil {
		return C.QWebSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocket) LocalAddress() *network.QHostAddress {
	defer qt.Recovering("QWebSocket::localAddress")

	if ptr.Pointer() != nil {
		return network.NewQHostAddressFromPointer(C.QWebSocket_LocalAddress(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) MaskGenerator() *QMaskGenerator {
	defer qt.Recovering("QWebSocket::maskGenerator")

	if ptr.Pointer() != nil {
		return NewQMaskGeneratorFromPointer(C.QWebSocket_MaskGenerator(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebSocket_Open2
func callbackQWebSocket_Open2(ptr unsafe.Pointer, ptrName *C.char, request unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open2"); signal != nil {
		signal.(func(*network.QNetworkRequest))(network.NewQNetworkRequestFromPointer(request))
	}

}

func (ptr *QWebSocket) ConnectOpen2(f func(request *network.QNetworkRequest)) {
	defer qt.Recovering("connect QWebSocket::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open2", f)
	}
}

func (ptr *QWebSocket) DisconnectOpen2(request network.QNetworkRequest_ITF) {
	defer qt.Recovering("disconnect QWebSocket::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open2")
	}
}

func (ptr *QWebSocket) Open2(request network.QNetworkRequest_ITF) {
	defer qt.Recovering("QWebSocket::open")

	if ptr.Pointer() != nil {
		C.QWebSocket_Open2(ptr.Pointer(), network.PointerFromQNetworkRequest(request))
	}
}

//export callbackQWebSocket_Open
func callbackQWebSocket_Open(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebSocket) ConnectOpen(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QWebSocket::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QWebSocket) DisconnectOpen(url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QWebSocket::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QWebSocket) Open(url core.QUrl_ITF) {
	defer qt.Recovering("QWebSocket::open")

	if ptr.Pointer() != nil {
		C.QWebSocket_Open(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebSocket) Origin() string {
	defer qt.Recovering("QWebSocket::origin")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) PauseMode() network.QAbstractSocket__PauseMode {
	defer qt.Recovering("QWebSocket::pauseMode")

	if ptr.Pointer() != nil {
		return network.QAbstractSocket__PauseMode(C.QWebSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) PeerAddress() *network.QHostAddress {
	defer qt.Recovering("QWebSocket::peerAddress")

	if ptr.Pointer() != nil {
		return network.NewQHostAddressFromPointer(C.QWebSocket_PeerAddress(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) PeerName() string {
	defer qt.Recovering("QWebSocket::peerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

//export callbackQWebSocket_Ping
func callbackQWebSocket_Ping(ptr unsafe.Pointer, ptrName *C.char, payload *C.char) {
	defer qt.Recovering("callback QWebSocket::ping")

	if signal := qt.GetSignal(C.GoString(ptrName), "ping"); signal != nil {
		signal.(func(string))(C.GoString(payload))
	}

}

func (ptr *QWebSocket) ConnectPing(f func(payload string)) {
	defer qt.Recovering("connect QWebSocket::ping")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "ping", f)
	}
}

func (ptr *QWebSocket) DisconnectPing(payload string) {
	defer qt.Recovering("disconnect QWebSocket::ping")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "ping")
	}
}

func (ptr *QWebSocket) Ping(payload string) {
	defer qt.Recovering("QWebSocket::ping")

	if ptr.Pointer() != nil {
		C.QWebSocket_Ping(ptr.Pointer(), C.CString(payload))
	}
}

func (ptr *QWebSocket) Proxy() *network.QNetworkProxy {
	defer qt.Recovering("QWebSocket::proxy")

	if ptr.Pointer() != nil {
		return network.NewQNetworkProxyFromPointer(C.QWebSocket_Proxy(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebSocket_ProxyAuthenticationRequired
func callbackQWebSocket_ProxyAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, proxy unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::proxyAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "proxyAuthenticationRequired"); signal != nil {
		signal.(func(*network.QNetworkProxy, *network.QAuthenticator))(network.NewQNetworkProxyFromPointer(proxy), network.NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocket) ConnectProxyAuthenticationRequired(f func(proxy *network.QNetworkProxy, authenticator *network.QAuthenticator)) {
	defer qt.Recovering("connect QWebSocket::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectProxyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "proxyAuthenticationRequired", f)
	}
}

func (ptr *QWebSocket) DisconnectProxyAuthenticationRequired() {
	defer qt.Recovering("disconnect QWebSocket::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "proxyAuthenticationRequired")
	}
}

func (ptr *QWebSocket) ProxyAuthenticationRequired(proxy network.QNetworkProxy_ITF, authenticator network.QAuthenticator_ITF) {
	defer qt.Recovering("QWebSocket::proxyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebSocket_ProxyAuthenticationRequired(ptr.Pointer(), network.PointerFromQNetworkProxy(proxy), network.PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QWebSocket) ReadBufferSize() int64 {
	defer qt.Recovering("QWebSocket::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_ReadChannelFinished
func callbackQWebSocket_ReadChannelFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::readChannelFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "readChannelFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectReadChannelFinished(f func()) {
	defer qt.Recovering("connect QWebSocket::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectReadChannelFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readChannelFinished", f)
	}
}

func (ptr *QWebSocket) DisconnectReadChannelFinished() {
	defer qt.Recovering("disconnect QWebSocket::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectReadChannelFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readChannelFinished")
	}
}

func (ptr *QWebSocket) ReadChannelFinished() {
	defer qt.Recovering("QWebSocket::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QWebSocket_ReadChannelFinished(ptr.Pointer())
	}
}

func (ptr *QWebSocket) Request() *network.QNetworkRequest {
	defer qt.Recovering("QWebSocket::request")

	if ptr.Pointer() != nil {
		return network.NewQNetworkRequestFromPointer(C.QWebSocket_Request(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) RequestUrl() *core.QUrl {
	defer qt.Recovering("QWebSocket::requestUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebSocket_RequestUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) ResourceName() string {
	defer qt.Recovering("QWebSocket::resourceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ResourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Resume() {
	defer qt.Recovering("QWebSocket::resume")

	if ptr.Pointer() != nil {
		C.QWebSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QWebSocket) SendBinaryMessage(data string) int64 {
	defer qt.Recovering("QWebSocket::sendBinaryMessage")

	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_SendBinaryMessage(ptr.Pointer(), C.CString(data)))
	}
	return 0
}

func (ptr *QWebSocket) SendTextMessage(message string) int64 {
	defer qt.Recovering("QWebSocket::sendTextMessage")

	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_SendTextMessage(ptr.Pointer(), C.CString(message)))
	}
	return 0
}

func (ptr *QWebSocket) SetMaskGenerator(maskGenerator QMaskGenerator_ITF) {
	defer qt.Recovering("QWebSocket::setMaskGenerator")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetMaskGenerator(ptr.Pointer(), PointerFromQMaskGenerator(maskGenerator))
	}
}

func (ptr *QWebSocket) SetPauseMode(pauseMode network.QAbstractSocket__PauseMode) {
	defer qt.Recovering("QWebSocket::setPauseMode")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetPauseMode(ptr.Pointer(), C.int(pauseMode))
	}
}

func (ptr *QWebSocket) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	defer qt.Recovering("QWebSocket::setProxy")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QWebSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QWebSocket) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	defer qt.Recovering("QWebSocket::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocket) SslConfiguration() *network.QSslConfiguration {
	defer qt.Recovering("QWebSocket::sslConfiguration")

	if ptr.Pointer() != nil {
		return network.NewQSslConfigurationFromPointer(C.QWebSocket_SslConfiguration(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) State() network.QAbstractSocket__SocketState {
	defer qt.Recovering("QWebSocket::state")

	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketState(C.QWebSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_StateChanged
func callbackQWebSocket_StateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QWebSocket::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketState))(network.QAbstractSocket__SocketState(state))
	}

}

func (ptr *QWebSocket) ConnectStateChanged(f func(state network.QAbstractSocket__SocketState)) {
	defer qt.Recovering("connect QWebSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QWebSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QWebSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

func (ptr *QWebSocket) StateChanged(state network.QAbstractSocket__SocketState) {
	defer qt.Recovering("QWebSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QWebSocket_StateChanged(ptr.Pointer(), C.int(state))
	}
}

//export callbackQWebSocket_TextFrameReceived
func callbackQWebSocket_TextFrameReceived(ptr unsafe.Pointer, ptrName *C.char, frame *C.char, isLastFrame C.int) {
	defer qt.Recovering("callback QWebSocket::textFrameReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "textFrameReceived"); signal != nil {
		signal.(func(string, bool))(C.GoString(frame), int(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectTextFrameReceived(f func(frame string, isLastFrame bool)) {
	defer qt.Recovering("connect QWebSocket::textFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextFrameReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextFrameReceived() {
	defer qt.Recovering("disconnect QWebSocket::textFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textFrameReceived")
	}
}

func (ptr *QWebSocket) TextFrameReceived(frame string, isLastFrame bool) {
	defer qt.Recovering("QWebSocket::textFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_TextFrameReceived(ptr.Pointer(), C.CString(frame), C.int(qt.GoBoolToInt(isLastFrame)))
	}
}

//export callbackQWebSocket_TextMessageReceived
func callbackQWebSocket_TextMessageReceived(ptr unsafe.Pointer, ptrName *C.char, message *C.char) {
	defer qt.Recovering("callback QWebSocket::textMessageReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "textMessageReceived"); signal != nil {
		signal.(func(string))(C.GoString(message))
	}

}

func (ptr *QWebSocket) ConnectTextMessageReceived(f func(message string)) {
	defer qt.Recovering("connect QWebSocket::textMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextMessageReceived() {
	defer qt.Recovering("disconnect QWebSocket::textMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textMessageReceived")
	}
}

func (ptr *QWebSocket) TextMessageReceived(message string) {
	defer qt.Recovering("QWebSocket::textMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_TextMessageReceived(ptr.Pointer(), C.CString(message))
	}
}

func (ptr *QWebSocket) DestroyQWebSocket() {
	defer qt.Recovering("QWebSocket::~QWebSocket")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebSocket_DestroyQWebSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebSocket_TimerEvent
func callbackQWebSocket_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebSocket_ChildEvent
func callbackQWebSocket_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebSocket_ConnectNotify
func callbackQWebSocket_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebSocket) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocket_CustomEvent
func callbackQWebSocket_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebSocket_DeleteLater
func callbackQWebSocket_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebSocket) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebSocket) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebSocket) DeleteLater() {
	defer qt.Recovering("QWebSocket::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebSocket_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocket) DeleteLaterDefault() {
	defer qt.Recovering("QWebSocket::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebSocket_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebSocket_DisconnectNotify
func callbackQWebSocket_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebSocket) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocket_Event
func callbackQWebSocket_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebSocket::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebSocket::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebSocket) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebSocket::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebSocket) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebSocket::event")

	if ptr.Pointer() != nil {
		return C.QWebSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebSocket) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebSocket::event")

	if ptr.Pointer() != nil {
		return C.QWebSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebSocket_EventFilter
func callbackQWebSocket_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebSocket::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebSocket) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebSocket_MetaObject
func callbackQWebSocket_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebSocket::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebSocket) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebSocket) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QWebSocketCorsAuthenticator struct {
	ptr unsafe.Pointer
}

type QWebSocketCorsAuthenticator_ITF interface {
	QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator
}

func (p *QWebSocketCorsAuthenticator) QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator {
	return p
}

func (p *QWebSocketCorsAuthenticator) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebSocketCorsAuthenticator) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebSocketCorsAuthenticator(ptr QWebSocketCorsAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketCorsAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketCorsAuthenticatorFromPointer(ptr unsafe.Pointer) *QWebSocketCorsAuthenticator {
	var n = new(QWebSocketCorsAuthenticator)
	n.SetPointer(ptr)
	return n
}

func newQWebSocketCorsAuthenticatorFromPointer(ptr unsafe.Pointer) *QWebSocketCorsAuthenticator {
	var n = NewQWebSocketCorsAuthenticatorFromPointer(ptr)
	return n
}

func NewQWebSocketCorsAuthenticator3(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	defer qt.Recovering("QWebSocketCorsAuthenticator::QWebSocketCorsAuthenticator")

	return newQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(PointerFromQWebSocketCorsAuthenticator(other)))
}

func NewQWebSocketCorsAuthenticator(origin string) *QWebSocketCorsAuthenticator {
	defer qt.Recovering("QWebSocketCorsAuthenticator::QWebSocketCorsAuthenticator")

	return newQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(C.CString(origin)))
}

func NewQWebSocketCorsAuthenticator2(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	defer qt.Recovering("QWebSocketCorsAuthenticator::QWebSocketCorsAuthenticator")

	return newQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(PointerFromQWebSocketCorsAuthenticator(other)))
}

func (ptr *QWebSocketCorsAuthenticator) Allowed() bool {
	defer qt.Recovering("QWebSocketCorsAuthenticator::allowed")

	if ptr.Pointer() != nil {
		return C.QWebSocketCorsAuthenticator_Allowed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketCorsAuthenticator) Origin() string {
	defer qt.Recovering("QWebSocketCorsAuthenticator::origin")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketCorsAuthenticator_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketCorsAuthenticator) SetAllowed(allowed bool) {
	defer qt.Recovering("QWebSocketCorsAuthenticator::setAllowed")

	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_SetAllowed(ptr.Pointer(), C.int(qt.GoBoolToInt(allowed)))
	}
}

func (ptr *QWebSocketCorsAuthenticator) Swap(other QWebSocketCorsAuthenticator_ITF) {
	defer qt.Recovering("QWebSocketCorsAuthenticator::swap")

	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_Swap(ptr.Pointer(), PointerFromQWebSocketCorsAuthenticator(other))
	}
}

func (ptr *QWebSocketCorsAuthenticator) DestroyQWebSocketCorsAuthenticator() {
	defer qt.Recovering("QWebSocketCorsAuthenticator::~QWebSocketCorsAuthenticator")

	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QWebSocketServer::SslMode
type QWebSocketServer__SslMode int64

const (
	QWebSocketServer__SecureMode    = QWebSocketServer__SslMode(0)
	QWebSocketServer__NonSecureMode = QWebSocketServer__SslMode(1)
)

type QWebSocketServer struct {
	core.QObject
}

type QWebSocketServer_ITF interface {
	core.QObject_ITF
	QWebSocketServer_PTR() *QWebSocketServer
}

func (p *QWebSocketServer) QWebSocketServer_PTR() *QWebSocketServer {
	return p
}

func (p *QWebSocketServer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QWebSocketServer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQWebSocketServer(ptr QWebSocketServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketServer_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketServerFromPointer(ptr unsafe.Pointer) *QWebSocketServer {
	var n = new(QWebSocketServer)
	n.SetPointer(ptr)
	return n
}

func newQWebSocketServerFromPointer(ptr unsafe.Pointer) *QWebSocketServer {
	var n = NewQWebSocketServerFromPointer(ptr)
	for len(n.ObjectName()) < len("QWebSocketServer_") {
		n.SetObjectName("QWebSocketServer_" + qt.Identifier())
	}
	return n
}

func NewQWebSocketServer(serverName string, secureMode QWebSocketServer__SslMode, parent core.QObject_ITF) *QWebSocketServer {
	defer qt.Recovering("QWebSocketServer::QWebSocketServer")

	return newQWebSocketServerFromPointer(C.QWebSocketServer_NewQWebSocketServer(C.CString(serverName), C.int(secureMode), core.PointerFromQObject(parent)))
}

//export callbackQWebSocketServer_AcceptError
func callbackQWebSocketServer_AcceptError(ptr unsafe.Pointer, ptrName *C.char, socketError C.int) {
	defer qt.Recovering("callback QWebSocketServer::acceptError")

	if signal := qt.GetSignal(C.GoString(ptrName), "acceptError"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QWebSocketServer) ConnectAcceptError(f func(socketError network.QAbstractSocket__SocketError)) {
	defer qt.Recovering("connect QWebSocketServer::acceptError")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectAcceptError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "acceptError", f)
	}
}

func (ptr *QWebSocketServer) DisconnectAcceptError() {
	defer qt.Recovering("disconnect QWebSocketServer::acceptError")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "acceptError")
	}
}

func (ptr *QWebSocketServer) AcceptError(socketError network.QAbstractSocket__SocketError) {
	defer qt.Recovering("QWebSocketServer::acceptError")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_AcceptError(ptr.Pointer(), C.int(socketError))
	}
}

func (ptr *QWebSocketServer) Close() {
	defer qt.Recovering("QWebSocketServer::close")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_Close(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_Closed
func callbackQWebSocketServer_Closed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocketServer::closed")

	if signal := qt.GetSignal(C.GoString(ptrName), "closed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) ConnectClosed(f func()) {
	defer qt.Recovering("connect QWebSocketServer::closed")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectClosed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "closed", f)
	}
}

func (ptr *QWebSocketServer) DisconnectClosed() {
	defer qt.Recovering("disconnect QWebSocketServer::closed")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "closed")
	}
}

func (ptr *QWebSocketServer) Closed() {
	defer qt.Recovering("QWebSocketServer::closed")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_Closed(ptr.Pointer())
	}
}

func (ptr *QWebSocketServer) ErrorString() string {
	defer qt.Recovering("QWebSocketServer::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) HasPendingConnections() bool {
	defer qt.Recovering("QWebSocketServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketServer) IsListening() bool {
	defer qt.Recovering("QWebSocketServer::isListening")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketServer) MaxPendingConnections() int {
	defer qt.Recovering("QWebSocketServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(C.QWebSocketServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocketServer_NewConnection
func callbackQWebSocketServer_NewConnection(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocketServer::newConnection")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QWebSocketServer::newConnection")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QWebSocketServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QWebSocketServer::newConnection")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

func (ptr *QWebSocketServer) NewConnection() {
	defer qt.Recovering("QWebSocketServer::newConnection")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_NewConnection(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_NextPendingConnection
func callbackQWebSocketServer_NextPendingConnection(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebSocketServer::nextPendingConnection")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextPendingConnection"); signal != nil {
		return PointerFromQWebSocket(signal.(func() *QWebSocket)())
	}

	return PointerFromQWebSocket(NewQWebSocketServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QWebSocketServer) ConnectNextPendingConnection(f func() *QWebSocket) {
	defer qt.Recovering("connect QWebSocketServer::nextPendingConnection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextPendingConnection", f)
	}
}

func (ptr *QWebSocketServer) DisconnectNextPendingConnection() {
	defer qt.Recovering("disconnect QWebSocketServer::nextPendingConnection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextPendingConnection")
	}
}

func (ptr *QWebSocketServer) NextPendingConnection() *QWebSocket {
	defer qt.Recovering("QWebSocketServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQWebSocketFromPointer(C.QWebSocketServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocketServer) NextPendingConnectionDefault() *QWebSocket {
	defer qt.Recovering("QWebSocketServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQWebSocketFromPointer(C.QWebSocketServer_NextPendingConnectionDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQWebSocketServer_OriginAuthenticationRequired
func callbackQWebSocketServer_OriginAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::originAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "originAuthenticationRequired"); signal != nil {
		signal.(func(*QWebSocketCorsAuthenticator))(NewQWebSocketCorsAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocketServer) ConnectOriginAuthenticationRequired(f func(authenticator *QWebSocketCorsAuthenticator)) {
	defer qt.Recovering("connect QWebSocketServer::originAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectOriginAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "originAuthenticationRequired", f)
	}
}

func (ptr *QWebSocketServer) DisconnectOriginAuthenticationRequired() {
	defer qt.Recovering("disconnect QWebSocketServer::originAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectOriginAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "originAuthenticationRequired")
	}
}

func (ptr *QWebSocketServer) OriginAuthenticationRequired(authenticator QWebSocketCorsAuthenticator_ITF) {
	defer qt.Recovering("QWebSocketServer::originAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_OriginAuthenticationRequired(ptr.Pointer(), PointerFromQWebSocketCorsAuthenticator(authenticator))
	}
}

func (ptr *QWebSocketServer) PauseAccepting() {
	defer qt.Recovering("QWebSocketServer::pauseAccepting")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_PauseAccepting(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_PeerVerifyError
func callbackQWebSocketServer_PeerVerifyError(ptr unsafe.Pointer, ptrName *C.char, error unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::peerVerifyError")

	if signal := qt.GetSignal(C.GoString(ptrName), "peerVerifyError"); signal != nil {
		signal.(func(*network.QSslError))(network.NewQSslErrorFromPointer(error))
	}

}

func (ptr *QWebSocketServer) ConnectPeerVerifyError(f func(error *network.QSslError)) {
	defer qt.Recovering("connect QWebSocketServer::peerVerifyError")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectPeerVerifyError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "peerVerifyError", f)
	}
}

func (ptr *QWebSocketServer) DisconnectPeerVerifyError() {
	defer qt.Recovering("disconnect QWebSocketServer::peerVerifyError")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectPeerVerifyError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "peerVerifyError")
	}
}

func (ptr *QWebSocketServer) PeerVerifyError(error network.QSslError_ITF) {
	defer qt.Recovering("QWebSocketServer::peerVerifyError")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_PeerVerifyError(ptr.Pointer(), network.PointerFromQSslError(error))
	}
}

func (ptr *QWebSocketServer) Proxy() *network.QNetworkProxy {
	defer qt.Recovering("QWebSocketServer::proxy")

	if ptr.Pointer() != nil {
		return network.NewQNetworkProxyFromPointer(C.QWebSocketServer_Proxy(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocketServer) ResumeAccepting() {
	defer qt.Recovering("QWebSocketServer::resumeAccepting")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ResumeAccepting(ptr.Pointer())
	}
}

func (ptr *QWebSocketServer) SecureMode() QWebSocketServer__SslMode {
	defer qt.Recovering("QWebSocketServer::secureMode")

	if ptr.Pointer() != nil {
		return QWebSocketServer__SslMode(C.QWebSocketServer_SecureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) ServerAddress() *network.QHostAddress {
	defer qt.Recovering("QWebSocketServer::serverAddress")

	if ptr.Pointer() != nil {
		return network.NewQHostAddressFromPointer(C.QWebSocketServer_ServerAddress(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocketServer) ServerName() string {
	defer qt.Recovering("QWebSocketServer::serverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketServer_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) ServerUrl() *core.QUrl {
	defer qt.Recovering("QWebSocketServer::serverUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebSocketServer_ServerUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocketServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QWebSocketServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

func (ptr *QWebSocketServer) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	defer qt.Recovering("QWebSocketServer::setProxy")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocketServer) SetServerName(serverName string) {
	defer qt.Recovering("QWebSocketServer::setServerName")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetServerName(ptr.Pointer(), C.CString(serverName))
	}
}

func (ptr *QWebSocketServer) SetSocketDescriptor(socketDescriptor int) bool {
	defer qt.Recovering("QWebSocketServer::setSocketDescriptor")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_SetSocketDescriptor(ptr.Pointer(), C.int(socketDescriptor)) != 0
	}
	return false
}

func (ptr *QWebSocketServer) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	defer qt.Recovering("QWebSocketServer::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocketServer) SocketDescriptor() int {
	defer qt.Recovering("QWebSocketServer::socketDescriptor")

	if ptr.Pointer() != nil {
		return int(C.QWebSocketServer_SocketDescriptor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) SslConfiguration() *network.QSslConfiguration {
	defer qt.Recovering("QWebSocketServer::sslConfiguration")

	if ptr.Pointer() != nil {
		return network.NewQSslConfigurationFromPointer(C.QWebSocketServer_SslConfiguration(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocketServer) DestroyQWebSocketServer() {
	defer qt.Recovering("QWebSocketServer::~QWebSocketServer")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebSocketServer_DestroyQWebSocketServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebSocketServer_TimerEvent
func callbackQWebSocketServer_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebSocketServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebSocketServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebSocketServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QWebSocketServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebSocketServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebSocketServer_ChildEvent
func callbackQWebSocketServer_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebSocketServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebSocketServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebSocketServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QWebSocketServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::childEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebSocketServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::childEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebSocketServer_ConnectNotify
func callbackQWebSocketServer_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocketServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebSocketServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QWebSocketServer) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QWebSocketServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QWebSocketServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebSocketServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebSocketServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebSocketServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocketServer_CustomEvent
func callbackQWebSocketServer_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebSocketServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebSocketServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebSocketServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QWebSocketServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::customEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebSocketServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::customEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebSocketServer_DeleteLater
func callbackQWebSocketServer_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocketServer::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebSocketServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebSocketServer) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QWebSocketServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QWebSocketServer) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QWebSocketServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QWebSocketServer) DeleteLater() {
	defer qt.Recovering("QWebSocketServer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebSocketServer_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocketServer) DeleteLaterDefault() {
	defer qt.Recovering("QWebSocketServer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QWebSocketServer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQWebSocketServer_DisconnectNotify
func callbackQWebSocketServer_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocketServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QWebSocketServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QWebSocketServer) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QWebSocketServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QWebSocketServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebSocketServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebSocketServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QWebSocketServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocketServer_Event
func callbackQWebSocketServer_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebSocketServer::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQWebSocketServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QWebSocketServer) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QWebSocketServer::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QWebSocketServer) DisconnectEvent() {
	defer qt.Recovering("disconnect QWebSocketServer::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QWebSocketServer) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebSocketServer::event")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebSocketServer) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QWebSocketServer::event")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebSocketServer_EventFilter
func callbackQWebSocketServer_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QWebSocketServer::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQWebSocketServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QWebSocketServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QWebSocketServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QWebSocketServer) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QWebSocketServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QWebSocketServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebSocketServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebSocketServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QWebSocketServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebSocketServer_MetaObject
func callbackQWebSocketServer_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QWebSocketServer::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebSocketServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebSocketServer) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QWebSocketServer::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QWebSocketServer) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QWebSocketServer::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QWebSocketServer) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QWebSocketServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocketServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocketServer) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QWebSocketServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocketServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
