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
func callbackQUdpSocketClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QUdpSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQUdpSocketDisconnectFromHost(ptrName *C.char) bool {
	defer qt.Recovering("callback QUdpSocket::disconnectFromHost")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectFromHost"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQUdpSocketResume(ptrName *C.char) bool {
	defer qt.Recovering("callback QUdpSocket::resume")

	if signal := qt.GetSignal(C.GoString(ptrName), "resume"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQUdpSocketSetReadBufferSize(ptrName *C.char, size C.longlong) bool {
	defer qt.Recovering("callback QUdpSocket::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
		return true
	}
	return false

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
func callbackQUdpSocketSetSocketOption(ptrName *C.char, option C.int, value unsafe.Pointer) bool {
	defer qt.Recovering("callback QUdpSocket::setSocketOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
		return true
	}
	return false

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
func callbackQUdpSocketTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QUdpSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQUdpSocketChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QUdpSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQUdpSocketCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QUdpSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
