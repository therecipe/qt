// +build !minimal

package websockets

//#include <stdint.h>
//#include <stdlib.h>
//#include "websockets.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtWebSockets_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

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
func NewQMaskGenerator(parent core.QObject_ITF) *QMaskGenerator {
	var tmpValue = NewQMaskGeneratorFromPointer(C.QMaskGenerator_NewQMaskGenerator(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQMaskGenerator_NextMask
func callbackQMaskGenerator_NextMask(ptr unsafe.Pointer) C.uint {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::nextMask"); signal != nil {
		return C.uint(uint32(signal.(func() uint)()))
	}

	return C.uint(uint32(0))
}

func (ptr *QMaskGenerator) ConnectNextMask(f func() uint) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::nextMask", f)
	}
}

func (ptr *QMaskGenerator) DisconnectNextMask() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::nextMask")
	}
}

func (ptr *QMaskGenerator) NextMask() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QMaskGenerator_NextMask(ptr.Pointer())))
	}
	return 0
}

//export callbackQMaskGenerator_Seed
func callbackQMaskGenerator_Seed(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::seed"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QMaskGenerator) ConnectSeed(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::seed", f)
	}
}

func (ptr *QMaskGenerator) DisconnectSeed() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::seed")
	}
}

func (ptr *QMaskGenerator) Seed() bool {
	if ptr.Pointer() != nil {
		return C.QMaskGenerator_Seed(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQMaskGenerator_DestroyQMaskGenerator
func callbackQMaskGenerator_DestroyQMaskGenerator(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::~QMaskGenerator"); signal != nil {
		signal.(func())()
	} else {
		NewQMaskGeneratorFromPointer(ptr).DestroyQMaskGeneratorDefault()
	}
}

func (ptr *QMaskGenerator) ConnectDestroyQMaskGenerator(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::~QMaskGenerator", f)
	}
}

func (ptr *QMaskGenerator) DisconnectDestroyQMaskGenerator() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::~QMaskGenerator")
	}
}

func (ptr *QMaskGenerator) DestroyQMaskGenerator() {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DestroyQMaskGenerator(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMaskGenerator) DestroyQMaskGeneratorDefault() {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DestroyQMaskGeneratorDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQMaskGenerator_TimerEvent
func callbackQMaskGenerator_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::timerEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::timerEvent")
	}
}

func (ptr *QMaskGenerator) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMaskGenerator) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQMaskGenerator_ChildEvent
func callbackQMaskGenerator_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::childEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::childEvent")
	}
}

func (ptr *QMaskGenerator) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMaskGenerator) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQMaskGenerator_ConnectNotify
func callbackQMaskGenerator_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMaskGeneratorFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMaskGenerator) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::connectNotify", f)
	}
}

func (ptr *QMaskGenerator) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::connectNotify")
	}
}

func (ptr *QMaskGenerator) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMaskGenerator) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMaskGenerator_CustomEvent
func callbackQMaskGenerator_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::customEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::customEvent")
	}
}

func (ptr *QMaskGenerator) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMaskGenerator) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQMaskGenerator_DeleteLater
func callbackQMaskGenerator_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQMaskGeneratorFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QMaskGenerator) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::deleteLater", f)
	}
}

func (ptr *QMaskGenerator) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::deleteLater")
	}
}

func (ptr *QMaskGenerator) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QMaskGenerator) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQMaskGenerator_DisconnectNotify
func callbackQMaskGenerator_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQMaskGeneratorFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QMaskGenerator) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::disconnectNotify", f)
	}
}

func (ptr *QMaskGenerator) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::disconnectNotify")
	}
}

func (ptr *QMaskGenerator) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QMaskGenerator) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QMaskGenerator_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQMaskGenerator_Event
func callbackQMaskGenerator_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMaskGeneratorFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QMaskGenerator) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::event", f)
	}
}

func (ptr *QMaskGenerator) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::event")
	}
}

func (ptr *QMaskGenerator) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMaskGenerator_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMaskGenerator) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMaskGenerator_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQMaskGenerator_EventFilter
func callbackQMaskGenerator_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQMaskGeneratorFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QMaskGenerator) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::eventFilter", f)
	}
}

func (ptr *QMaskGenerator) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::eventFilter")
	}
}

func (ptr *QMaskGenerator) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMaskGenerator_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMaskGenerator) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QMaskGenerator_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQMaskGenerator_MetaObject
func callbackQMaskGenerator_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QMaskGenerator::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQMaskGeneratorFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QMaskGenerator) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::metaObject", f)
	}
}

func (ptr *QMaskGenerator) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QMaskGenerator::metaObject")
	}
}

func (ptr *QMaskGenerator) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QMaskGenerator_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMaskGenerator) MetaObjectDefault() *core.QMetaObject {
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
func NewQWebSocket(origin string, version QWebSocketProtocol__Version, parent core.QObject_ITF) *QWebSocket {
	var originC = C.CString(origin)
	defer C.free(unsafe.Pointer(originC))
	var tmpValue = NewQWebSocketFromPointer(C.QWebSocket_NewQWebSocket(originC, C.longlong(version), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QWebSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Abort(ptr.Pointer())
	}
}

//export callbackQWebSocket_AboutToClose
func callbackQWebSocket_AboutToClose(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::aboutToClose"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectAboutToClose(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectAboutToClose(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::aboutToClose", f)
	}
}

func (ptr *QWebSocket) DisconnectAboutToClose() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectAboutToClose(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::aboutToClose")
	}
}

func (ptr *QWebSocket) AboutToClose() {
	if ptr.Pointer() != nil {
		C.QWebSocket_AboutToClose(ptr.Pointer())
	}
}

//export callbackQWebSocket_BinaryFrameReceived
func callbackQWebSocket_BinaryFrameReceived(ptr unsafe.Pointer, frame unsafe.Pointer, isLastFrame C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::binaryFrameReceived"); signal != nil {
		signal.(func(*core.QByteArray, bool))(core.NewQByteArrayFromPointer(frame), int8(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectBinaryFrameReceived(f func(frame *core.QByteArray, isLastFrame bool)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryFrameReceived(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::binaryFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryFrameReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::binaryFrameReceived")
	}
}

func (ptr *QWebSocket) BinaryFrameReceived(frame core.QByteArray_ITF, isLastFrame bool) {
	if ptr.Pointer() != nil {
		C.QWebSocket_BinaryFrameReceived(ptr.Pointer(), core.PointerFromQByteArray(frame), C.char(int8(qt.GoBoolToInt(isLastFrame))))
	}
}

//export callbackQWebSocket_BinaryMessageReceived
func callbackQWebSocket_BinaryMessageReceived(ptr unsafe.Pointer, message unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::binaryMessageReceived"); signal != nil {
		signal.(func(*core.QByteArray))(core.NewQByteArrayFromPointer(message))
	}

}

func (ptr *QWebSocket) ConnectBinaryMessageReceived(f func(message *core.QByteArray)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryMessageReceived(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::binaryMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryMessageReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::binaryMessageReceived")
	}
}

func (ptr *QWebSocket) BinaryMessageReceived(message core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_BinaryMessageReceived(ptr.Pointer(), core.PointerFromQByteArray(message))
	}
}

//export callbackQWebSocket_BytesWritten
func callbackQWebSocket_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::bytesWritten"); signal != nil {
		signal.(func(int64))(int64(bytes))
	}

}

func (ptr *QWebSocket) ConnectBytesWritten(f func(bytes int64)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBytesWritten(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::bytesWritten", f)
	}
}

func (ptr *QWebSocket) DisconnectBytesWritten() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBytesWritten(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::bytesWritten")
	}
}

func (ptr *QWebSocket) BytesWritten(bytes int64) {
	if ptr.Pointer() != nil {
		C.QWebSocket_BytesWritten(ptr.Pointer(), C.longlong(bytes))
	}
}

//export callbackQWebSocket_Close
func callbackQWebSocket_Close(ptr unsafe.Pointer, closeCode C.longlong, reason C.struct_QtWebSockets_PackedString) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::close"); signal != nil {
		signal.(func(QWebSocketProtocol__CloseCode, string))(QWebSocketProtocol__CloseCode(closeCode), cGoUnpackString(reason))
	}

}

func (ptr *QWebSocket) ConnectClose(f func(closeCode QWebSocketProtocol__CloseCode, reason string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::close", f)
	}
}

func (ptr *QWebSocket) DisconnectClose(closeCode QWebSocketProtocol__CloseCode, reason string) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::close")
	}
}

func (ptr *QWebSocket) Close(closeCode QWebSocketProtocol__CloseCode, reason string) {
	if ptr.Pointer() != nil {
		var reasonC = C.CString(reason)
		defer C.free(unsafe.Pointer(reasonC))
		C.QWebSocket_Close(ptr.Pointer(), C.longlong(closeCode), reasonC)
	}
}

func (ptr *QWebSocket) CloseCode() QWebSocketProtocol__CloseCode {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__CloseCode(C.QWebSocket_CloseCode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) CloseReason() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_CloseReason(ptr.Pointer()))
	}
	return ""
}

//export callbackQWebSocket_Connected
func callbackQWebSocket_Connected(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::connected", f)
	}
}

func (ptr *QWebSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::connected")
	}
}

func (ptr *QWebSocket) Connected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Connected(ptr.Pointer())
	}
}

//export callbackQWebSocket_Disconnected
func callbackQWebSocket_Disconnected(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::disconnected", f)
	}
}

func (ptr *QWebSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::disconnected")
	}
}

func (ptr *QWebSocket) Disconnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Disconnected(ptr.Pointer())
	}
}

//export callbackQWebSocket_Error2
func callbackQWebSocket_Error2(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::error2"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(error))
	}

}

func (ptr *QWebSocket) ConnectError2(f func(error network.QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::error2", f)
	}
}

func (ptr *QWebSocket) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::error2")
	}
}

func (ptr *QWebSocket) Error2(error network.QAbstractSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Error2(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QWebSocket) Error() network.QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketError(C.QWebSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQWebSocket_IgnoreSslErrors
func callbackQWebSocket_IgnoreSslErrors(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::ignoreSslErrors"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectIgnoreSslErrors(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::ignoreSslErrors", f)
	}
}

func (ptr *QWebSocket) DisconnectIgnoreSslErrors() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::ignoreSslErrors")
	}
}

func (ptr *QWebSocket) IgnoreSslErrors() {
	if ptr.Pointer() != nil {
		C.QWebSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QWebSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocket) LocalAddress() *network.QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQHostAddressFromPointer(C.QWebSocket_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) LocalPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QWebSocket_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) MaskGenerator() *QMaskGenerator {
	if ptr.Pointer() != nil {
		var tmpValue = NewQMaskGeneratorFromPointer(C.QWebSocket_MaskGenerator(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebSocket_Open2
func callbackQWebSocket_Open2(ptr unsafe.Pointer, request unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::open2"); signal != nil {
		signal.(func(*network.QNetworkRequest))(network.NewQNetworkRequestFromPointer(request))
	}

}

func (ptr *QWebSocket) ConnectOpen2(f func(request *network.QNetworkRequest)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::open2", f)
	}
}

func (ptr *QWebSocket) DisconnectOpen2(request network.QNetworkRequest_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::open2")
	}
}

func (ptr *QWebSocket) Open2(request network.QNetworkRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Open2(ptr.Pointer(), network.PointerFromQNetworkRequest(request))
	}
}

//export callbackQWebSocket_Open
func callbackQWebSocket_Open(ptr unsafe.Pointer, url unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::open"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QWebSocket) ConnectOpen(f func(url *core.QUrl)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::open", f)
	}
}

func (ptr *QWebSocket) DisconnectOpen(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::open")
	}
}

func (ptr *QWebSocket) Open(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Open(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebSocket) Origin() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) PauseMode() network.QAbstractSocket__PauseMode {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__PauseMode(C.QWebSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) PeerAddress() *network.QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQHostAddressFromPointer(C.QWebSocket_PeerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) PeerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QWebSocket_PeerPort(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_Ping
func callbackQWebSocket_Ping(ptr unsafe.Pointer, payload unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::ping"); signal != nil {
		signal.(func(*core.QByteArray))(core.NewQByteArrayFromPointer(payload))
	}

}

func (ptr *QWebSocket) ConnectPing(f func(payload *core.QByteArray)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::ping", f)
	}
}

func (ptr *QWebSocket) DisconnectPing(payload core.QByteArray_ITF) {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::ping")
	}
}

func (ptr *QWebSocket) Ping(payload core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Ping(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

//export callbackQWebSocket_Pong
func callbackQWebSocket_Pong(ptr unsafe.Pointer, elapsedTime C.ulonglong, payload unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::pong"); signal != nil {
		signal.(func(uint64, *core.QByteArray))(uint64(elapsedTime), core.NewQByteArrayFromPointer(payload))
	}

}

func (ptr *QWebSocket) ConnectPong(f func(elapsedTime uint64, payload *core.QByteArray)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectPong(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::pong", f)
	}
}

func (ptr *QWebSocket) DisconnectPong() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectPong(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::pong")
	}
}

func (ptr *QWebSocket) Pong(elapsedTime uint64, payload core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Pong(ptr.Pointer(), C.ulonglong(elapsedTime), core.PointerFromQByteArray(payload))
	}
}

func (ptr *QWebSocket) Proxy() *network.QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkProxyFromPointer(C.QWebSocket_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

//export callbackQWebSocket_ProxyAuthenticationRequired
func callbackQWebSocket_ProxyAuthenticationRequired(ptr unsafe.Pointer, proxy unsafe.Pointer, authenticator unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::proxyAuthenticationRequired"); signal != nil {
		signal.(func(*network.QNetworkProxy, *network.QAuthenticator))(network.NewQNetworkProxyFromPointer(proxy), network.NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocket) ConnectProxyAuthenticationRequired(f func(proxy *network.QNetworkProxy, authenticator *network.QAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectProxyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::proxyAuthenticationRequired", f)
	}
}

func (ptr *QWebSocket) DisconnectProxyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectProxyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::proxyAuthenticationRequired")
	}
}

func (ptr *QWebSocket) ProxyAuthenticationRequired(proxy network.QNetworkProxy_ITF, authenticator network.QAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ProxyAuthenticationRequired(ptr.Pointer(), network.PointerFromQNetworkProxy(proxy), network.PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QWebSocket) ReadBufferSize() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_ReadChannelFinished
func callbackQWebSocket_ReadChannelFinished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::readChannelFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectReadChannelFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectReadChannelFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::readChannelFinished", f)
	}
}

func (ptr *QWebSocket) DisconnectReadChannelFinished() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectReadChannelFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::readChannelFinished")
	}
}

func (ptr *QWebSocket) ReadChannelFinished() {
	if ptr.Pointer() != nil {
		C.QWebSocket_ReadChannelFinished(ptr.Pointer())
	}
}

func (ptr *QWebSocket) Request() *network.QNetworkRequest {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkRequestFromPointer(C.QWebSocket_Request(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QNetworkRequest).DestroyQNetworkRequest)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) RequestUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebSocket_RequestUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) ResourceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocket_ResourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QWebSocket) SendBinaryMessage(data core.QByteArray_ITF) int64 {
	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_SendBinaryMessage(ptr.Pointer(), core.PointerFromQByteArray(data)))
	}
	return 0
}

func (ptr *QWebSocket) SendTextMessage(message string) int64 {
	if ptr.Pointer() != nil {
		var messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
		return int64(C.QWebSocket_SendTextMessage(ptr.Pointer(), messageC))
	}
	return 0
}

func (ptr *QWebSocket) SetMaskGenerator(maskGenerator QMaskGenerator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetMaskGenerator(ptr.Pointer(), PointerFromQMaskGenerator(maskGenerator))
	}
}

func (ptr *QWebSocket) SetPauseMode(pauseMode network.QAbstractSocket__PauseMode) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetPauseMode(ptr.Pointer(), C.longlong(pauseMode))
	}
}

func (ptr *QWebSocket) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocket) SetReadBufferSize(size int64) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QWebSocket) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocket) SslConfiguration() *network.QSslConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQSslConfigurationFromPointer(C.QWebSocket_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocket) State() network.QAbstractSocket__SocketState {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketState(C.QWebSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_StateChanged
func callbackQWebSocket_StateChanged(ptr unsafe.Pointer, state C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::stateChanged"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketState))(network.QAbstractSocket__SocketState(state))
	}

}

func (ptr *QWebSocket) ConnectStateChanged(f func(state network.QAbstractSocket__SocketState)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::stateChanged", f)
	}
}

func (ptr *QWebSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::stateChanged")
	}
}

func (ptr *QWebSocket) StateChanged(state network.QAbstractSocket__SocketState) {
	if ptr.Pointer() != nil {
		C.QWebSocket_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQWebSocket_TextFrameReceived
func callbackQWebSocket_TextFrameReceived(ptr unsafe.Pointer, frame C.struct_QtWebSockets_PackedString, isLastFrame C.char) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::textFrameReceived"); signal != nil {
		signal.(func(string, bool))(cGoUnpackString(frame), int8(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectTextFrameReceived(f func(frame string, isLastFrame bool)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextFrameReceived(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::textFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextFrameReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::textFrameReceived")
	}
}

func (ptr *QWebSocket) TextFrameReceived(frame string, isLastFrame bool) {
	if ptr.Pointer() != nil {
		var frameC = C.CString(frame)
		defer C.free(unsafe.Pointer(frameC))
		C.QWebSocket_TextFrameReceived(ptr.Pointer(), frameC, C.char(int8(qt.GoBoolToInt(isLastFrame))))
	}
}

//export callbackQWebSocket_TextMessageReceived
func callbackQWebSocket_TextMessageReceived(ptr unsafe.Pointer, message C.struct_QtWebSockets_PackedString) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::textMessageReceived"); signal != nil {
		signal.(func(string))(cGoUnpackString(message))
	}

}

func (ptr *QWebSocket) ConnectTextMessageReceived(f func(message string)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextMessageReceived(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::textMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextMessageReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::textMessageReceived")
	}
}

func (ptr *QWebSocket) TextMessageReceived(message string) {
	if ptr.Pointer() != nil {
		var messageC = C.CString(message)
		defer C.free(unsafe.Pointer(messageC))
		C.QWebSocket_TextMessageReceived(ptr.Pointer(), messageC)
	}
}

func (ptr *QWebSocket) Version() QWebSocketProtocol__Version {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__Version(C.QWebSocket_Version(ptr.Pointer()))
	}
	return 0
}

//export callbackQWebSocket_DestroyQWebSocket
func callbackQWebSocket_DestroyQWebSocket(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::~QWebSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQWebSocketFromPointer(ptr).DestroyQWebSocketDefault()
	}
}

func (ptr *QWebSocket) ConnectDestroyQWebSocket(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::~QWebSocket", f)
	}
}

func (ptr *QWebSocket) DisconnectDestroyQWebSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::~QWebSocket")
	}
}

func (ptr *QWebSocket) DestroyQWebSocket() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DestroyQWebSocket(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocket) DestroyQWebSocketDefault() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DestroyQWebSocketDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebSocket_TimerEvent
func callbackQWebSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::timerEvent", f)
	}
}

func (ptr *QWebSocket) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::timerEvent")
	}
}

func (ptr *QWebSocket) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebSocket_ChildEvent
func callbackQWebSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::childEvent", f)
	}
}

func (ptr *QWebSocket) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::childEvent")
	}
}

func (ptr *QWebSocket) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebSocket_ConnectNotify
func callbackQWebSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::connectNotify", f)
	}
}

func (ptr *QWebSocket) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::connectNotify")
	}
}

func (ptr *QWebSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocket_CustomEvent
func callbackQWebSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::customEvent", f)
	}
}

func (ptr *QWebSocket) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::customEvent")
	}
}

func (ptr *QWebSocket) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebSocket_DeleteLater
func callbackQWebSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebSocket) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::deleteLater", f)
	}
}

func (ptr *QWebSocket) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::deleteLater")
	}
}

func (ptr *QWebSocket) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebSocket_DisconnectNotify
func callbackQWebSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::disconnectNotify", f)
	}
}

func (ptr *QWebSocket) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::disconnectNotify")
	}
}

func (ptr *QWebSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocket_Event
func callbackQWebSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::event", f)
	}
}

func (ptr *QWebSocket) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::event")
	}
}

func (ptr *QWebSocket) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebSocket_EventFilter
func callbackQWebSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::eventFilter", f)
	}
}

func (ptr *QWebSocket) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::eventFilter")
	}
}

func (ptr *QWebSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebSocket_MetaObject
func callbackQWebSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocket::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::metaObject", f)
	}
}

func (ptr *QWebSocket) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocket::metaObject")
	}
}

func (ptr *QWebSocket) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) MetaObjectDefault() *core.QMetaObject {
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
func NewQWebSocketCorsAuthenticator3(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	var tmpValue = NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(PointerFromQWebSocketCorsAuthenticator(other)))
	runtime.SetFinalizer(tmpValue, (*QWebSocketCorsAuthenticator).DestroyQWebSocketCorsAuthenticator)
	return tmpValue
}

func NewQWebSocketCorsAuthenticator(origin string) *QWebSocketCorsAuthenticator {
	var originC = C.CString(origin)
	defer C.free(unsafe.Pointer(originC))
	var tmpValue = NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(originC))
	runtime.SetFinalizer(tmpValue, (*QWebSocketCorsAuthenticator).DestroyQWebSocketCorsAuthenticator)
	return tmpValue
}

func NewQWebSocketCorsAuthenticator2(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	var tmpValue = NewQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(PointerFromQWebSocketCorsAuthenticator(other)))
	runtime.SetFinalizer(tmpValue, (*QWebSocketCorsAuthenticator).DestroyQWebSocketCorsAuthenticator)
	return tmpValue
}

func (ptr *QWebSocketCorsAuthenticator) Allowed() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketCorsAuthenticator_Allowed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketCorsAuthenticator) Origin() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocketCorsAuthenticator_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketCorsAuthenticator) SetAllowed(allowed bool) {
	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_SetAllowed(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(allowed))))
	}
}

func (ptr *QWebSocketCorsAuthenticator) Swap(other QWebSocketCorsAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_Swap(ptr.Pointer(), PointerFromQWebSocketCorsAuthenticator(other))
	}
}

func (ptr *QWebSocketCorsAuthenticator) DestroyQWebSocketCorsAuthenticator() {
	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QWebSocketProtocol::CloseCode
type QWebSocketProtocol__CloseCode int64

const (
	QWebSocketProtocol__CloseCodeNormal                = QWebSocketProtocol__CloseCode(1000)
	QWebSocketProtocol__CloseCodeGoingAway             = QWebSocketProtocol__CloseCode(1001)
	QWebSocketProtocol__CloseCodeProtocolError         = QWebSocketProtocol__CloseCode(1002)
	QWebSocketProtocol__CloseCodeDatatypeNotSupported  = QWebSocketProtocol__CloseCode(1003)
	QWebSocketProtocol__CloseCodeReserved1004          = QWebSocketProtocol__CloseCode(1004)
	QWebSocketProtocol__CloseCodeMissingStatusCode     = QWebSocketProtocol__CloseCode(1005)
	QWebSocketProtocol__CloseCodeAbnormalDisconnection = QWebSocketProtocol__CloseCode(1006)
	QWebSocketProtocol__CloseCodeWrongDatatype         = QWebSocketProtocol__CloseCode(1007)
	QWebSocketProtocol__CloseCodePolicyViolated        = QWebSocketProtocol__CloseCode(1008)
	QWebSocketProtocol__CloseCodeTooMuchData           = QWebSocketProtocol__CloseCode(1009)
	QWebSocketProtocol__CloseCodeMissingExtension      = QWebSocketProtocol__CloseCode(1010)
	QWebSocketProtocol__CloseCodeBadOperation          = QWebSocketProtocol__CloseCode(1011)
	QWebSocketProtocol__CloseCodeTlsHandshakeFailed    = QWebSocketProtocol__CloseCode(1015)
)

//QWebSocketProtocol::Version
type QWebSocketProtocol__Version int64

const (
	QWebSocketProtocol__VersionUnknown = QWebSocketProtocol__Version(-1)
	QWebSocketProtocol__Version0       = QWebSocketProtocol__Version(0)
	QWebSocketProtocol__Version4       = QWebSocketProtocol__Version(4)
	QWebSocketProtocol__Version5       = QWebSocketProtocol__Version(5)
	QWebSocketProtocol__Version6       = QWebSocketProtocol__Version(6)
	QWebSocketProtocol__Version7       = QWebSocketProtocol__Version(7)
	QWebSocketProtocol__Version8       = QWebSocketProtocol__Version(8)
	QWebSocketProtocol__Version13      = QWebSocketProtocol__Version(13)
	QWebSocketProtocol__VersionLatest  = QWebSocketProtocol__Version(QWebSocketProtocol__Version13)
)

type QWebSocketProtocol struct {
	ptr unsafe.Pointer
}

type QWebSocketProtocol_ITF interface {
	QWebSocketProtocol_PTR() *QWebSocketProtocol
}

func (p *QWebSocketProtocol) QWebSocketProtocol_PTR() *QWebSocketProtocol {
	return p
}

func (p *QWebSocketProtocol) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QWebSocketProtocol) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQWebSocketProtocol(ptr QWebSocketProtocol_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketProtocol_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketProtocolFromPointer(ptr unsafe.Pointer) *QWebSocketProtocol {
	var n = new(QWebSocketProtocol)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWebSocketProtocol) DestroyQWebSocketProtocol() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
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
func NewQWebSocketServer(serverName string, secureMode QWebSocketServer__SslMode, parent core.QObject_ITF) *QWebSocketServer {
	var serverNameC = C.CString(serverName)
	defer C.free(unsafe.Pointer(serverNameC))
	var tmpValue = NewQWebSocketServerFromPointer(C.QWebSocketServer_NewQWebSocketServer(serverNameC, C.longlong(secureMode), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQWebSocketServer_AcceptError
func callbackQWebSocketServer_AcceptError(ptr unsafe.Pointer, socketError C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::acceptError"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QWebSocketServer) ConnectAcceptError(f func(socketError network.QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectAcceptError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::acceptError", f)
	}
}

func (ptr *QWebSocketServer) DisconnectAcceptError() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::acceptError")
	}
}

func (ptr *QWebSocketServer) AcceptError(socketError network.QAbstractSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_AcceptError(ptr.Pointer(), C.longlong(socketError))
	}
}

func (ptr *QWebSocketServer) Close() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_Close(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_Closed
func callbackQWebSocketServer_Closed(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::closed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) ConnectClosed(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectClosed(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::closed", f)
	}
}

func (ptr *QWebSocketServer) DisconnectClosed() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::closed")
	}
}

func (ptr *QWebSocketServer) Closed() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_Closed(ptr.Pointer())
	}
}

func (ptr *QWebSocketServer) Error() QWebSocketProtocol__CloseCode {
	if ptr.Pointer() != nil {
		return QWebSocketProtocol__CloseCode(C.QWebSocketServer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocketServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketServer) Listen(address network.QHostAddress_ITF, port uint16) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_Listen(ptr.Pointer(), network.PointerFromQHostAddress(address), C.ushort(port)) != 0
	}
	return false
}

func (ptr *QWebSocketServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebSocketServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

//export callbackQWebSocketServer_NewConnection
func callbackQWebSocketServer_NewConnection(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::newConnection", f)
	}
}

func (ptr *QWebSocketServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::newConnection")
	}
}

func (ptr *QWebSocketServer) NewConnection() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_NewConnection(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_NextPendingConnection
func callbackQWebSocketServer_NextPendingConnection(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::nextPendingConnection"); signal != nil {
		return PointerFromQWebSocket(signal.(func() *QWebSocket)())
	}

	return PointerFromQWebSocket(NewQWebSocketServerFromPointer(ptr).NextPendingConnectionDefault())
}

func (ptr *QWebSocketServer) ConnectNextPendingConnection(f func() *QWebSocket) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::nextPendingConnection", f)
	}
}

func (ptr *QWebSocketServer) DisconnectNextPendingConnection() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::nextPendingConnection")
	}
}

func (ptr *QWebSocketServer) NextPendingConnection() *QWebSocket {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebSocketFromPointer(C.QWebSocketServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) NextPendingConnectionDefault() *QWebSocket {
	if ptr.Pointer() != nil {
		var tmpValue = NewQWebSocketFromPointer(C.QWebSocketServer_NextPendingConnectionDefault(ptr.Pointer()))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQWebSocketServer_OriginAuthenticationRequired
func callbackQWebSocketServer_OriginAuthenticationRequired(ptr unsafe.Pointer, authenticator unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::originAuthenticationRequired"); signal != nil {
		signal.(func(*QWebSocketCorsAuthenticator))(NewQWebSocketCorsAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocketServer) ConnectOriginAuthenticationRequired(f func(authenticator *QWebSocketCorsAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectOriginAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::originAuthenticationRequired", f)
	}
}

func (ptr *QWebSocketServer) DisconnectOriginAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectOriginAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::originAuthenticationRequired")
	}
}

func (ptr *QWebSocketServer) OriginAuthenticationRequired(authenticator QWebSocketCorsAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_OriginAuthenticationRequired(ptr.Pointer(), PointerFromQWebSocketCorsAuthenticator(authenticator))
	}
}

func (ptr *QWebSocketServer) PauseAccepting() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_PauseAccepting(ptr.Pointer())
	}
}

//export callbackQWebSocketServer_PeerVerifyError
func callbackQWebSocketServer_PeerVerifyError(ptr unsafe.Pointer, error unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::peerVerifyError"); signal != nil {
		signal.(func(*network.QSslError))(network.NewQSslErrorFromPointer(error))
	}

}

func (ptr *QWebSocketServer) ConnectPeerVerifyError(f func(error *network.QSslError)) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectPeerVerifyError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::peerVerifyError", f)
	}
}

func (ptr *QWebSocketServer) DisconnectPeerVerifyError() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectPeerVerifyError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::peerVerifyError")
	}
}

func (ptr *QWebSocketServer) PeerVerifyError(error network.QSslError_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_PeerVerifyError(ptr.Pointer(), network.PointerFromQSslError(error))
	}
}

func (ptr *QWebSocketServer) Proxy() *network.QNetworkProxy {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQNetworkProxyFromPointer(C.QWebSocketServer_Proxy(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QNetworkProxy).DestroyQNetworkProxy)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) ResumeAccepting() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ResumeAccepting(ptr.Pointer())
	}
}

func (ptr *QWebSocketServer) SecureMode() QWebSocketServer__SslMode {
	if ptr.Pointer() != nil {
		return QWebSocketServer__SslMode(C.QWebSocketServer_SecureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) ServerAddress() *network.QHostAddress {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQHostAddressFromPointer(C.QWebSocketServer_ServerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QHostAddress).DestroyQHostAddress)
		return tmpValue
	}
	return nil
}

//export callbackQWebSocketServer_ServerError
func callbackQWebSocketServer_ServerError(ptr unsafe.Pointer, closeCode C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::serverError"); signal != nil {
		signal.(func(QWebSocketProtocol__CloseCode))(QWebSocketProtocol__CloseCode(closeCode))
	}

}

func (ptr *QWebSocketServer) ConnectServerError(f func(closeCode QWebSocketProtocol__CloseCode)) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectServerError(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::serverError", f)
	}
}

func (ptr *QWebSocketServer) DisconnectServerError() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectServerError(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::serverError")
	}
}

func (ptr *QWebSocketServer) ServerError(closeCode QWebSocketProtocol__CloseCode) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ServerError(ptr.Pointer(), C.longlong(closeCode))
	}
}

func (ptr *QWebSocketServer) ServerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QWebSocketServer_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) ServerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QWebSocketServer_ServerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) ServerUrl() *core.QUrl {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQUrlFromPointer(C.QWebSocketServer_ServerUrl(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QUrl).DestroyQUrl)
		return tmpValue
	}
	return nil
}

func (ptr *QWebSocketServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetMaxPendingConnections(ptr.Pointer(), C.int(int32(numConnections)))
	}
}

func (ptr *QWebSocketServer) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocketServer) SetServerName(serverName string) {
	if ptr.Pointer() != nil {
		var serverNameC = C.CString(serverName)
		defer C.free(unsafe.Pointer(serverNameC))
		C.QWebSocketServer_SetServerName(ptr.Pointer(), serverNameC)
	}
}

func (ptr *QWebSocketServer) SetSocketDescriptor(socketDescriptor int) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_SetSocketDescriptor(ptr.Pointer(), C.int(int32(socketDescriptor))) != 0
	}
	return false
}

func (ptr *QWebSocketServer) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocketServer) SocketDescriptor() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QWebSocketServer_SocketDescriptor(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebSocketServer) SslConfiguration() *network.QSslConfiguration {
	if ptr.Pointer() != nil {
		var tmpValue = network.NewQSslConfigurationFromPointer(C.QWebSocketServer_SslConfiguration(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*network.QSslConfiguration).DestroyQSslConfiguration)
		return tmpValue
	}
	return nil
}

//export callbackQWebSocketServer_DestroyQWebSocketServer
func callbackQWebSocketServer_DestroyQWebSocketServer(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::~QWebSocketServer"); signal != nil {
		signal.(func())()
	} else {
		NewQWebSocketServerFromPointer(ptr).DestroyQWebSocketServerDefault()
	}
}

func (ptr *QWebSocketServer) ConnectDestroyQWebSocketServer(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::~QWebSocketServer", f)
	}
}

func (ptr *QWebSocketServer) DisconnectDestroyQWebSocketServer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::~QWebSocketServer")
	}
}

func (ptr *QWebSocketServer) DestroyQWebSocketServer() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DestroyQWebSocketServer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocketServer) DestroyQWebSocketServerDefault() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DestroyQWebSocketServerDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebSocketServer_TimerEvent
func callbackQWebSocketServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::timerEvent", f)
	}
}

func (ptr *QWebSocketServer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::timerEvent")
	}
}

func (ptr *QWebSocketServer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebSocketServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQWebSocketServer_ChildEvent
func callbackQWebSocketServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::childEvent", f)
	}
}

func (ptr *QWebSocketServer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::childEvent")
	}
}

func (ptr *QWebSocketServer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebSocketServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQWebSocketServer_ConnectNotify
func callbackQWebSocketServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocketServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::connectNotify", f)
	}
}

func (ptr *QWebSocketServer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::connectNotify")
	}
}

func (ptr *QWebSocketServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebSocketServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocketServer_CustomEvent
func callbackQWebSocketServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::customEvent", f)
	}
}

func (ptr *QWebSocketServer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::customEvent")
	}
}

func (ptr *QWebSocketServer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebSocketServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQWebSocketServer_DeleteLater
func callbackQWebSocketServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQWebSocketServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QWebSocketServer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::deleteLater", f)
	}
}

func (ptr *QWebSocketServer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::deleteLater")
	}
}

func (ptr *QWebSocketServer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocketServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQWebSocketServer_DisconnectNotify
func callbackQWebSocketServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQWebSocketServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QWebSocketServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::disconnectNotify", f)
	}
}

func (ptr *QWebSocketServer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::disconnectNotify")
	}
}

func (ptr *QWebSocketServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QWebSocketServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQWebSocketServer_Event
func callbackQWebSocketServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QWebSocketServer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::event", f)
	}
}

func (ptr *QWebSocketServer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::event")
	}
}

func (ptr *QWebSocketServer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QWebSocketServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQWebSocketServer_EventFilter
func callbackQWebSocketServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQWebSocketServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QWebSocketServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::eventFilter", f)
	}
}

func (ptr *QWebSocketServer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::eventFilter")
	}
}

func (ptr *QWebSocketServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QWebSocketServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQWebSocketServer_MetaObject
func callbackQWebSocketServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QWebSocketServer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQWebSocketServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QWebSocketServer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::metaObject", f)
	}
}

func (ptr *QWebSocketServer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QWebSocketServer::metaObject")
	}
}

func (ptr *QWebSocketServer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocketServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocketServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QWebSocketServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
