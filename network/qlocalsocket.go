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

type QLocalSocketITF interface {
	core.QIODeviceITF
	QLocalSocketPTR() *QLocalSocket
}

func PointerFromQLocalSocket(ptr QLocalSocketITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalSocketPTR().Pointer()
	}
	return nil
}

func QLocalSocketFromPointer(ptr unsafe.Pointer) *QLocalSocket {
	var n = new(QLocalSocket)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLocalSocket_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLocalSocket) QLocalSocketPTR() *QLocalSocket {
	return ptr
}

//QLocalSocket::LocalSocketError
type QLocalSocket__LocalSocketError int

var (
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
type QLocalSocket__LocalSocketState int

var (
	QLocalSocket__UnconnectedState = QLocalSocket__LocalSocketState(QAbstractSocket__UnconnectedState)
	QLocalSocket__ConnectingState  = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectingState)
	QLocalSocket__ConnectedState   = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectedState)
	QLocalSocket__ClosingState     = QLocalSocket__LocalSocketState(QAbstractSocket__ClosingState)
)

func (ptr *QLocalSocket) Open(openMode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_Open(C.QtObjectPtr(ptr.Pointer()), C.int(openMode)) != 0
	}
	return false
}

func NewQLocalSocket(parent core.QObjectITF) *QLocalSocket {
	return QLocalSocketFromPointer(unsafe.Pointer(C.QLocalSocket_NewQLocalSocket(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QLocalSocket) ConnectToServer2(name string, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer2(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(openMode))
	}
}

func (ptr *QLocalSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectConnected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QLocalSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectConnected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQLocalSocketConnected
func callbackQLocalSocketConnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "connected").(func())()
}

func (ptr *QLocalSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectDisconnected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QLocalSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectDisconnected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQLocalSocketDisconnected
func callbackQLocalSocketDisconnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QLocalSocket) FullServerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_FullServerName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLocalSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsSequential(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalSocket) ServerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_ServerName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLocalSocket) SetServerName(name string) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_SetServerName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QLocalSocket) ConnectStateChanged(f func(socketState QLocalSocket__LocalSocketState)) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLocalSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQLocalSocketStateChanged
func callbackQLocalSocketStateChanged(ptrName *C.char, socketState C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QLocalSocket__LocalSocketState))(QLocalSocket__LocalSocketState(socketState))
}

func (ptr *QLocalSocket) DestroyQLocalSocket() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DestroyQLocalSocket(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Abort(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLocalSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_CanReadLine(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalSocket) Close() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLocalSocket) ConnectToServer(openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer(C.QtObjectPtr(ptr.Pointer()), C.int(openMode))
	}
}

func (ptr *QLocalSocket) DisconnectFromServer() {
	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectFromServer(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLocalSocket) Error() QLocalSocket__LocalSocketError {
	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketError(C.QLocalSocket_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLocalSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_Flush(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForBytesWritten(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForConnected(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForDisconnected(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForReadyRead(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}
