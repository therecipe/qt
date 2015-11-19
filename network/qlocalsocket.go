package network

//#include "qlocalsocket.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QLocalSocket_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return C.QLocalSocket_Open(ptr.Pointer(), C.int(openMode)) != 0
	}
	return false
}

func NewQLocalSocket(parent core.QObject_ITF) *QLocalSocket {
	return NewQLocalSocketFromPointer(C.QLocalSocket_NewQLocalSocket(core.PointerFromQObject(parent)))
}

func (ptr *QLocalSocket) ConnectToServer2(name string, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer2(ptr.Pointer(), C.CString(name), C.int(openMode))
	}
}

func (ptr *QLocalSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QLocalSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQLocalSocketConnected
func callbackQLocalSocketConnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "connected").(func())()
}

func (ptr *QLocalSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QLocalSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQLocalSocketDisconnected
func callbackQLocalSocketDisconnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QLocalSocket) FullServerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_FullServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ServerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) SetServerName(name string) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_SetServerName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QLocalSocket) ConnectStateChanged(f func(socketState QLocalSocket__LocalSocketState)) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLocalSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQLocalSocketStateChanged
func callbackQLocalSocketStateChanged(ptrName *C.char, socketState C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QLocalSocket__LocalSocketState))(QLocalSocket__LocalSocketState(socketState))
}

func (ptr *QLocalSocket) DestroyQLocalSocket() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DestroyQLocalSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) Close() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Close(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) ConnectToServer(openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer(ptr.Pointer(), C.int(openMode))
	}
}

func (ptr *QLocalSocket) DisconnectFromServer() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectFromServer(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) Error() QLocalSocket__LocalSocketError {
	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketError(C.QLocalSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForConnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForDisconnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}
