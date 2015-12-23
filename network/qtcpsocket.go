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
func callbackQTcpSocketClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QTcpSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQTcpSocketDisconnectFromHost(ptrName *C.char) bool {
	defer qt.Recovering("callback QTcpSocket::disconnectFromHost")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectFromHost"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQTcpSocketResume(ptrName *C.char) bool {
	defer qt.Recovering("callback QTcpSocket::resume")

	if signal := qt.GetSignal(C.GoString(ptrName), "resume"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQTcpSocketSetReadBufferSize(ptrName *C.char, size C.longlong) bool {
	defer qt.Recovering("callback QTcpSocket::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
		return true
	}
	return false

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
func callbackQTcpSocketSetSocketOption(ptrName *C.char, option C.int, value unsafe.Pointer) bool {
	defer qt.Recovering("callback QTcpSocket::setSocketOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
		return true
	}
	return false

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
func callbackQTcpSocketTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTcpSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTcpSocketChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTcpSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQTcpSocketCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QTcpSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
