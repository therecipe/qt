package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::open")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Open(ptr.Pointer(), C.int(openMode)) != 0
	}
	return false
}

func NewQLocalSocket(parent core.QObject_ITF) *QLocalSocket {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::QLocalSocket")
		}
	}()

	return NewQLocalSocketFromPointer(C.QLocalSocket_NewQLocalSocket(core.PointerFromQObject(parent)))
}

func (ptr *QLocalSocket) ConnectToServer2(name string, openMode core.QIODevice__OpenModeFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::connectToServer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer2(ptr.Pointer(), C.CString(name), C.int(openMode))
	}
}

func (ptr *QLocalSocket) ConnectConnected(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::connected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QLocalSocket) DisconnectConnected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::connected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQLocalSocketConnected
func callbackQLocalSocketConnected(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::connected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "connected").(func())()
}

func (ptr *QLocalSocket) ConnectDisconnected(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::disconnected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QLocalSocket) DisconnectDisconnected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::disconnected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQLocalSocketDisconnected
func callbackQLocalSocketDisconnected(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::disconnected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QLocalSocket) FullServerName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::fullServerName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_FullServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) IsSequential() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::isSequential")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ServerName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::serverName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) SetServerName(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::setServerName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_SetServerName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QLocalSocket) ConnectStateChanged(f func(socketState QLocalSocket__LocalSocketState)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLocalSocket) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQLocalSocketStateChanged
func callbackQLocalSocketStateChanged(ptrName *C.char, socketState C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QLocalSocket__LocalSocketState))(QLocalSocket__LocalSocketState(socketState))
}

func (ptr *QLocalSocket) DestroyQLocalSocket() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::~QLocalSocket")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_DestroyQLocalSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) Abort() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::abort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) CanReadLine() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::canReadLine")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLocalSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) Close() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::close")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_Close(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) ConnectToServer(openMode core.QIODevice__OpenModeFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::connectToServer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer(ptr.Pointer(), C.int(openMode))
	}
}

func (ptr *QLocalSocket) DisconnectFromServer() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::disconnectFromServer")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectFromServer(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) Error() QLocalSocket__LocalSocketError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketError(C.QLocalSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) Flush() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::flush")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::waitForBytesWritten")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForConnected(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::waitForConnected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForConnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForDisconnected(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::waitForDisconnected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForDisconnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForReadyRead(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocalSocket::waitForReadyRead")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}
