package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QBluetoothSocket struct {
	core.QIODevice
}

type QBluetoothSocket_ITF interface {
	core.QIODevice_ITF
	QBluetoothSocket_PTR() *QBluetoothSocket
}

func PointerFromQBluetoothSocket(ptr QBluetoothSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothSocket_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothSocketFromPointer(ptr unsafe.Pointer) *QBluetoothSocket {
	var n = new(QBluetoothSocket)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothSocket_") {
		n.SetObjectName("QBluetoothSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothSocket) QBluetoothSocket_PTR() *QBluetoothSocket {
	return ptr
}

//QBluetoothSocket::SocketError
type QBluetoothSocket__SocketError int64

const (
	QBluetoothSocket__NoSocketError            = QBluetoothSocket__SocketError(-2)
	QBluetoothSocket__UnknownSocketError       = QBluetoothSocket__SocketError(network.QAbstractSocket__UnknownSocketError)
	QBluetoothSocket__HostNotFoundError        = QBluetoothSocket__SocketError(network.QAbstractSocket__HostNotFoundError)
	QBluetoothSocket__ServiceNotFoundError     = QBluetoothSocket__SocketError(network.QAbstractSocket__SocketAddressNotAvailableError)
	QBluetoothSocket__NetworkError             = QBluetoothSocket__SocketError(network.QAbstractSocket__NetworkError)
	QBluetoothSocket__UnsupportedProtocolError = QBluetoothSocket__SocketError(8)
	QBluetoothSocket__OperationError           = QBluetoothSocket__SocketError(network.QAbstractSocket__OperationError)
)

//QBluetoothSocket::SocketState
type QBluetoothSocket__SocketState int64

const (
	QBluetoothSocket__UnconnectedState   = QBluetoothSocket__SocketState(network.QAbstractSocket__UnconnectedState)
	QBluetoothSocket__ServiceLookupState = QBluetoothSocket__SocketState(network.QAbstractSocket__HostLookupState)
	QBluetoothSocket__ConnectingState    = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectingState)
	QBluetoothSocket__ConnectedState     = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectedState)
	QBluetoothSocket__BoundState         = QBluetoothSocket__SocketState(network.QAbstractSocket__BoundState)
	QBluetoothSocket__ClosingState       = QBluetoothSocket__SocketState(network.QAbstractSocket__ClosingState)
	QBluetoothSocket__ListeningState     = QBluetoothSocket__SocketState(network.QAbstractSocket__ListeningState)
)

func (ptr *QBluetoothSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::connected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QBluetoothSocket::connected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQBluetoothSocketConnected
func callbackQBluetoothSocketConnected(ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothSocket::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QBluetoothSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQBluetoothSocketDisconnected
func callbackQBluetoothSocketDisconnected(ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothSocket::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) ConnectError2(f func(error QBluetoothSocket__SocketError)) {
	defer qt.Recovering("connect QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothSocketError2
func callbackQBluetoothSocketError2(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothSocket::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothSocket__SocketError))(QBluetoothSocket__SocketError(error))
	}

}

func (ptr *QBluetoothSocket) ConnectStateChanged(f func(state QBluetoothSocket__SocketState)) {
	defer qt.Recovering("connect QBluetoothSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QBluetoothSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQBluetoothSocketStateChanged
func callbackQBluetoothSocketStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QBluetoothSocket::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QBluetoothSocket__SocketState))(QBluetoothSocket__SocketState(state))
	}

}

func NewQBluetoothSocket(socketType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothSocket {
	defer qt.Recovering("QBluetoothSocket::QBluetoothSocket")

	return NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket(C.int(socketType), core.PointerFromQObject(parent)))
}

func NewQBluetoothSocket2(parent core.QObject_ITF) *QBluetoothSocket {
	defer qt.Recovering("QBluetoothSocket::QBluetoothSocket")

	return NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket2(core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothSocket) Abort() {
	defer qt.Recovering("QBluetoothSocket::abort")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) BytesAvailable() int64 {
	defer qt.Recovering("QBluetoothSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) BytesToWrite() int64 {
	defer qt.Recovering("QBluetoothSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) CanReadLine() bool {
	defer qt.Recovering("QBluetoothSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QBluetoothSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQBluetoothSocketClose
func callbackQBluetoothSocketClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QBluetoothSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QBluetoothSocket) ConnectToService2(address QBluetoothAddress_ITF, uuid QBluetoothUuid_ITF, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QBluetoothSocket::connectToService")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService2(ptr.Pointer(), PointerFromQBluetoothAddress(address), PointerFromQBluetoothUuid(uuid), C.int(openMode))
	}
}

func (ptr *QBluetoothSocket) ConnectToService(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QBluetoothSocket::connectToService")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.int(openMode))
	}
}

func (ptr *QBluetoothSocket) DisconnectFromService() {
	defer qt.Recovering("QBluetoothSocket::disconnectFromService")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectFromService(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) Error() QBluetoothSocket__SocketError {
	defer qt.Recovering("QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketError(C.QBluetoothSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) ErrorString() string {
	defer qt.Recovering("QBluetoothSocket::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) IsSequential() bool {
	defer qt.Recovering("QBluetoothSocket::isSequential")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) LocalName() string {
	defer qt.Recovering("QBluetoothSocket::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) PeerName() string {
	defer qt.Recovering("QBluetoothSocket::peerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) SetSocketDescriptor(socketDescriptor int, socketType QBluetoothServiceInfo__Protocol, socketState QBluetoothSocket__SocketState, openMode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QBluetoothSocket::setSocketDescriptor")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_SetSocketDescriptor(ptr.Pointer(), C.int(socketDescriptor), C.int(socketType), C.int(socketState), C.int(openMode)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) SocketDescriptor() int {
	defer qt.Recovering("QBluetoothSocket::socketDescriptor")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothSocket_SocketDescriptor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) SocketType() QBluetoothServiceInfo__Protocol {
	defer qt.Recovering("QBluetoothSocket::socketType")

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) State() QBluetoothSocket__SocketState {
	defer qt.Recovering("QBluetoothSocket::state")

	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketState(C.QBluetoothSocket_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocket() {
	defer qt.Recovering("QBluetoothSocket::~QBluetoothSocket")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DestroyQBluetoothSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
