package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLocalSocket struct {
	core.QIODevice
}

type QLocalSocket_ITF interface {
	core.QIODevice_ITF
	QLocalSocket_PTR() *QLocalSocket
}

func PointerFromQLocalSocket(ptr QLocalSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalSocket_PTR().Pointer()
	}
	return nil
}

func NewQLocalSocketFromPointer(ptr unsafe.Pointer) *QLocalSocket {
	var n = new(QLocalSocket)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLocalSocket_") {
		n.SetObjectName("QLocalSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QLocalSocket) QLocalSocket_PTR() *QLocalSocket {
	return ptr
}

//QLocalSocket::LocalSocketError
type QLocalSocket__LocalSocketError int64

const (
	QLocalSocket__ConnectionRefusedError          = QLocalSocket__LocalSocketError(QAbstractSocket__ConnectionRefusedError)
	QLocalSocket__PeerClosedError                 = QLocalSocket__LocalSocketError(QAbstractSocket__RemoteHostClosedError)
	QLocalSocket__ServerNotFoundError             = QLocalSocket__LocalSocketError(QAbstractSocket__HostNotFoundError)
	QLocalSocket__SocketAccessError               = QLocalSocket__LocalSocketError(QAbstractSocket__SocketAccessError)
	QLocalSocket__SocketResourceError             = QLocalSocket__LocalSocketError(QAbstractSocket__SocketResourceError)
	QLocalSocket__SocketTimeoutError              = QLocalSocket__LocalSocketError(QAbstractSocket__SocketTimeoutError)
	QLocalSocket__DatagramTooLargeError           = QLocalSocket__LocalSocketError(QAbstractSocket__DatagramTooLargeError)
	QLocalSocket__ConnectionError                 = QLocalSocket__LocalSocketError(QAbstractSocket__NetworkError)
	QLocalSocket__UnsupportedSocketOperationError = QLocalSocket__LocalSocketError(QAbstractSocket__UnsupportedSocketOperationError)
	QLocalSocket__UnknownSocketError              = QLocalSocket__LocalSocketError(QAbstractSocket__UnknownSocketError)
	QLocalSocket__OperationError                  = QLocalSocket__LocalSocketError(QAbstractSocket__OperationError)
)

//QLocalSocket::LocalSocketState
type QLocalSocket__LocalSocketState int64

const (
	QLocalSocket__UnconnectedState = QLocalSocket__LocalSocketState(QAbstractSocket__UnconnectedState)
	QLocalSocket__ConnectingState  = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectingState)
	QLocalSocket__ConnectedState   = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectedState)
	QLocalSocket__ClosingState     = QLocalSocket__LocalSocketState(QAbstractSocket__ClosingState)
)

func (ptr *QLocalSocket) Open(openMode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QLocalSocket::open")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Open(ptr.Pointer(), C.int(openMode)) != 0
	}
	return false
}

func NewQLocalSocket(parent core.QObject_ITF) *QLocalSocket {
	defer qt.Recovering("QLocalSocket::QLocalSocket")

	return NewQLocalSocketFromPointer(C.QLocalSocket_NewQLocalSocket(core.PointerFromQObject(parent)))
}

func (ptr *QLocalSocket) ConnectToServer2(name string, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QLocalSocket::connectToServer")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer2(ptr.Pointer(), C.CString(name), C.int(openMode))
	}
}

func (ptr *QLocalSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QLocalSocket::connected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QLocalSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QLocalSocket::connected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQLocalSocketConnected
func callbackQLocalSocketConnected(ptrName *C.char) {
	defer qt.Recovering("callback QLocalSocket::connected")

	var signal = qt.GetSignal(C.GoString(ptrName), "connected")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QLocalSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QLocalSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QLocalSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQLocalSocketDisconnected
func callbackQLocalSocketDisconnected(ptrName *C.char) {
	defer qt.Recovering("callback QLocalSocket::disconnected")

	var signal = qt.GetSignal(C.GoString(ptrName), "disconnected")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalSocket) ConnectError2(f func(socketError QLocalSocket__LocalSocketError)) {
	defer qt.Recovering("connect QLocalSocket::error")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QLocalSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QLocalSocket::error")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQLocalSocketError2
func callbackQLocalSocketError2(ptrName *C.char, socketError C.int) {
	defer qt.Recovering("callback QLocalSocket::error")

	var signal = qt.GetSignal(C.GoString(ptrName), "error")
	if signal != nil {
		signal.(func(QLocalSocket__LocalSocketError))(QLocalSocket__LocalSocketError(socketError))
	}

}

func (ptr *QLocalSocket) FullServerName() string {
	defer qt.Recovering("QLocalSocket::fullServerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_FullServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) IsSequential() bool {
	defer qt.Recovering("QLocalSocket::isSequential")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ServerName() string {
	defer qt.Recovering("QLocalSocket::serverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) SetServerName(name string) {
	defer qt.Recovering("QLocalSocket::setServerName")

	if ptr.Pointer() != nil {
		C.QLocalSocket_SetServerName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QLocalSocket) State() QLocalSocket__LocalSocketState {
	defer qt.Recovering("QLocalSocket::state")

	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketState(C.QLocalSocket_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) ConnectStateChanged(f func(socketState QLocalSocket__LocalSocketState)) {
	defer qt.Recovering("connect QLocalSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLocalSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QLocalSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQLocalSocketStateChanged
func callbackQLocalSocketStateChanged(ptrName *C.char, socketState C.int) {
	defer qt.Recovering("callback QLocalSocket::stateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stateChanged")
	if signal != nil {
		signal.(func(QLocalSocket__LocalSocketState))(QLocalSocket__LocalSocketState(socketState))
	}

}

func (ptr *QLocalSocket) DestroyQLocalSocket() {
	defer qt.Recovering("QLocalSocket::~QLocalSocket")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DestroyQLocalSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) Abort() {
	defer qt.Recovering("QLocalSocket::abort")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) BytesAvailable() int64 {
	defer qt.Recovering("QLocalSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) BytesToWrite() int64 {
	defer qt.Recovering("QLocalSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) CanReadLine() bool {
	defer qt.Recovering("QLocalSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QLocalSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QLocalSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QLocalSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQLocalSocketClose
func callbackQLocalSocketClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QLocalSocket::close")

	var signal = qt.GetSignal(C.GoString(ptrName), "close")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QLocalSocket) ConnectToServer(openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QLocalSocket::connectToServer")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer(ptr.Pointer(), C.int(openMode))
	}
}

func (ptr *QLocalSocket) DisconnectFromServer() {
	defer qt.Recovering("QLocalSocket::disconnectFromServer")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectFromServer(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) Error() QLocalSocket__LocalSocketError {
	defer qt.Recovering("QLocalSocket::error")

	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketError(C.QLocalSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) Flush() bool {
	defer qt.Recovering("QLocalSocket::flush")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsValid() bool {
	defer qt.Recovering("QLocalSocket::isValid")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ReadBufferSize() int64 {
	defer qt.Recovering("QLocalSocket::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QLocalSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QLocalSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForConnected(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForConnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForDisconnected(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForDisconnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}
