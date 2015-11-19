package bluetooth

//#include "qbluetoothsocket.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QBluetoothSocket_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQBluetoothSocketConnected
func callbackQBluetoothSocketConnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "connected").(func())()
}

func (ptr *QBluetoothSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQBluetoothSocketDisconnected
func callbackQBluetoothSocketDisconnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QBluetoothSocket) ConnectStateChanged(f func(state QBluetoothSocket__SocketState)) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQBluetoothSocketStateChanged
func callbackQBluetoothSocketStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QBluetoothSocket__SocketState))(QBluetoothSocket__SocketState(state))
}

func NewQBluetoothSocket(socketType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothSocket {
	return NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket(C.int(socketType), core.PointerFromQObject(parent)))
}

func NewQBluetoothSocket2(parent core.QObject_ITF) *QBluetoothSocket {
	return NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket2(core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) Close() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) ConnectToService2(address QBluetoothAddress_ITF, uuid QBluetoothUuid_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService2(ptr.Pointer(), PointerFromQBluetoothAddress(address), PointerFromQBluetoothUuid(uuid), C.int(openMode))
	}
}

func (ptr *QBluetoothSocket) ConnectToService(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.int(openMode))
	}
}

func (ptr *QBluetoothSocket) DisconnectFromService() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectFromService(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) Error() QBluetoothSocket__SocketError {
	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketError(C.QBluetoothSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) LocalName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) SetSocketDescriptor(socketDescriptor int, socketType QBluetoothServiceInfo__Protocol, socketState QBluetoothSocket__SocketState, openMode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_SetSocketDescriptor(ptr.Pointer(), C.int(socketDescriptor), C.int(socketType), C.int(socketState), C.int(openMode)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) SocketDescriptor() int {
	if ptr.Pointer() != nil {
		return int(C.QBluetoothSocket_SocketDescriptor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) SocketType() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocket() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DestroyQBluetoothSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
