package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTcpSocket struct {
	QAbstractSocket
}

type QTcpSocket_ITF interface {
	QAbstractSocket_ITF
	QTcpSocket_PTR() *QTcpSocket
}

func PointerFromQTcpSocket(ptr QTcpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func NewQTcpSocketFromPointer(ptr unsafe.Pointer) *QTcpSocket {
	var n = new(QTcpSocket)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QTcpSocket_") {
		n.SetObjectName("QTcpSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QTcpSocket) QTcpSocket_PTR() *QTcpSocket {
	return ptr
}

func NewQTcpSocket(parent core.QObject_ITF) *QTcpSocket {
	defer qt.Recovering("QTcpSocket::QTcpSocket")

	return NewQTcpSocketFromPointer(C.QTcpSocket_NewQTcpSocket(core.PointerFromQObject(parent)))
}

func (ptr *QTcpSocket) DestroyQTcpSocket() {
	defer qt.Recovering("QTcpSocket::~QTcpSocket")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DestroyQTcpSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QTcpSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QTcpSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QTcpSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQTcpSocketClose
func callbackQTcpSocketClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTcpSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QTcpSocket) Close() {
	defer qt.Recovering("QTcpSocket::close")

	if ptr.Pointer() != nil {
		C.QTcpSocket_Close(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) CloseDefault() {
	defer qt.Recovering("QTcpSocket::close")

	if ptr.Pointer() != nil {
		C.QTcpSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectFromHost", f)
	}
}

func (ptr *QTcpSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectFromHost")
	}
}

//export callbackQTcpSocketDisconnectFromHost
func callbackQTcpSocketDisconnectFromHost(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTcpSocket::disconnectFromHost")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QTcpSocket) DisconnectFromHost() {
	defer qt.Recovering("QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QTcpSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resume", f)
	}
}

func (ptr *QTcpSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QTcpSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resume")
	}
}

//export callbackQTcpSocketResume
func callbackQTcpSocketResume(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTcpSocket::resume")

	if signal := qt.GetSignal(C.GoString(ptrName), "resume"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QTcpSocket) Resume() {
	defer qt.Recovering("QTcpSocket::resume")

	if ptr.Pointer() != nil {
		C.QTcpSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ResumeDefault() {
	defer qt.Recovering("QTcpSocket::resume")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setReadBufferSize", f)
	}
}

func (ptr *QTcpSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setReadBufferSize")
	}
}

//export callbackQTcpSocketSetReadBufferSize
func callbackQTcpSocketSetReadBufferSize(ptr unsafe.Pointer, ptrName *C.char, size C.longlong) {
	defer qt.Recovering("callback QTcpSocket::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQTcpSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QTcpSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QTcpSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QTcpSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSocketOption", f)
	}
}

func (ptr *QTcpSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSocketOption")
	}
}

//export callbackQTcpSocketSetSocketOption
func callbackQTcpSocketSetSocketOption(ptr unsafe.Pointer, ptrName *C.char, option C.int, value unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::setSocketOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQTcpSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QTcpSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetSocketOption(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QTcpSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetSocketOptionDefault(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QTcpSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTcpSocketTimerEvent
func callbackQTcpSocketTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTcpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTcpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTcpSocketChildEvent
func callbackQTcpSocketChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTcpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTcpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTcpSocketCustomEvent
func callbackQTcpSocketCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTcpSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
