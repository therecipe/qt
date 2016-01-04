package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QUdpSocket struct {
	QAbstractSocket
}

type QUdpSocket_ITF interface {
	QAbstractSocket_ITF
	QUdpSocket_PTR() *QUdpSocket
}

func PointerFromQUdpSocket(ptr QUdpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUdpSocket_PTR().Pointer()
	}
	return nil
}

func NewQUdpSocketFromPointer(ptr unsafe.Pointer) *QUdpSocket {
	var n = new(QUdpSocket)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QUdpSocket_") {
		n.SetObjectName("QUdpSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QUdpSocket) QUdpSocket_PTR() *QUdpSocket {
	return ptr
}

func NewQUdpSocket(parent core.QObject_ITF) *QUdpSocket {
	defer qt.Recovering("QUdpSocket::QUdpSocket")

	return NewQUdpSocketFromPointer(C.QUdpSocket_NewQUdpSocket(core.PointerFromQObject(parent)))
}

func (ptr *QUdpSocket) HasPendingDatagrams() bool {
	defer qt.Recovering("QUdpSocket::hasPendingDatagrams")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_HasPendingDatagrams(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddress_ITF) bool {
	defer qt.Recovering("QUdpSocket::joinMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	defer qt.Recovering("QUdpSocket::joinMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddress_ITF) bool {
	defer qt.Recovering("QUdpSocket::leaveMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	defer qt.Recovering("QUdpSocket::leaveMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) PendingDatagramSize() int64 {
	defer qt.Recovering("QUdpSocket::pendingDatagramSize")

	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_PendingDatagramSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) SetMulticastInterface(iface QNetworkInterface_ITF) {
	defer qt.Recovering("QUdpSocket::setMulticastInterface")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetMulticastInterface(ptr.Pointer(), PointerFromQNetworkInterface(iface))
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocket() {
	defer qt.Recovering("QUdpSocket::~QUdpSocket")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DestroyQUdpSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QUdpSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QUdpSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QUdpSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QUdpSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQUdpSocketClose
func callbackQUdpSocketClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QUdpSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QUdpSocket) Close() {
	defer qt.Recovering("QUdpSocket::close")

	if ptr.Pointer() != nil {
		C.QUdpSocket_Close(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) CloseDefault() {
	defer qt.Recovering("QUdpSocket::close")

	if ptr.Pointer() != nil {
		C.QUdpSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectFromHost", f)
	}
}

func (ptr *QUdpSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectFromHost")
	}
}

//export callbackQUdpSocketDisconnectFromHost
func callbackQUdpSocketDisconnectFromHost(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QUdpSocket::disconnectFromHost")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QUdpSocket) DisconnectFromHost() {
	defer qt.Recovering("QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QUdpSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resume", f)
	}
}

func (ptr *QUdpSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QUdpSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resume")
	}
}

//export callbackQUdpSocketResume
func callbackQUdpSocketResume(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QUdpSocket::resume")

	if signal := qt.GetSignal(C.GoString(ptrName), "resume"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QUdpSocket) Resume() {
	defer qt.Recovering("QUdpSocket::resume")

	if ptr.Pointer() != nil {
		C.QUdpSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ResumeDefault() {
	defer qt.Recovering("QUdpSocket::resume")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setReadBufferSize", f)
	}
}

func (ptr *QUdpSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setReadBufferSize")
	}
}

//export callbackQUdpSocketSetReadBufferSize
func callbackQUdpSocketSetReadBufferSize(ptr unsafe.Pointer, ptrName *C.char, size C.longlong) {
	defer qt.Recovering("callback QUdpSocket::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQUdpSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QUdpSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QUdpSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QUdpSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSocketOption", f)
	}
}

func (ptr *QUdpSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSocketOption")
	}
}

//export callbackQUdpSocketSetSocketOption
func callbackQUdpSocketSetSocketOption(ptr unsafe.Pointer, ptrName *C.char, option C.int, value unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::setSocketOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQUdpSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QUdpSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetSocketOption(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QUdpSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetSocketOptionDefault(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QUdpSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQUdpSocketTimerEvent
func callbackQUdpSocketTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QUdpSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QUdpSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QUdpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QUdpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQUdpSocketChildEvent
func callbackQUdpSocketChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUdpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUdpSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUdpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUdpSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QUdpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QUdpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQUdpSocketCustomEvent
func callbackQUdpSocketCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QUdpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUdpSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QUdpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
