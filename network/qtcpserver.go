package network

//#include "qtcpserver.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTcpServer struct {
	core.QObject
}

type QTcpServer_ITF interface {
	core.QObject_ITF
	QTcpServer_PTR() *QTcpServer
}

func PointerFromQTcpServer(ptr QTcpServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpServer_PTR().Pointer()
	}
	return nil
}

func NewQTcpServerFromPointer(ptr unsafe.Pointer) *QTcpServer {
	var n = new(QTcpServer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTcpServer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTcpServer) QTcpServer_PTR() *QTcpServer {
	return ptr
}

func NewQTcpServer(parent core.QObject_ITF) *QTcpServer {
	return NewQTcpServerFromPointer(C.QTcpServer_NewQTcpServer(core.PointerFromQObject(parent)))
}

func (ptr *QTcpServer) ConnectAcceptError(f func(socketError QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectAcceptError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "acceptError", f)
	}
}

func (ptr *QTcpServer) DisconnectAcceptError() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "acceptError")
	}
}

//export callbackQTcpServerAcceptError
func callbackQTcpServerAcceptError(ptrName *C.char, socketError C.int) {
	qt.GetSignal(C.GoString(ptrName), "acceptError").(func(QAbstractSocket__SocketError))(QAbstractSocket__SocketError(socketError))
}

func (ptr *QTcpServer) Close() {
	if ptr.Pointer() != nil {
		C.QTcpServer_Close(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTcpServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTcpServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(C.QTcpServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QTcpServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQTcpServerNewConnection
func callbackQTcpServerNewConnection(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "newConnection").(func())()
}

func (ptr *QTcpServer) NextPendingConnection() *QTcpSocket {
	if ptr.Pointer() != nil {
		return NewQTcpSocketFromPointer(C.QTcpServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTcpServer) PauseAccepting() {
	if ptr.Pointer() != nil {
		C.QTcpServer_PauseAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ResumeAccepting() {
	if ptr.Pointer() != nil {
		C.QTcpServer_ResumeAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ServerError() QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QTcpServer_ServerError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QTcpServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

func (ptr *QTcpServer) SetProxy(networkProxy QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QTcpServer) WaitForNewConnection(msec int, timedOut bool) bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_WaitForNewConnection(ptr.Pointer(), C.int(msec), C.int(qt.GoBoolToInt(timedOut))) != 0
	}
	return false
}

func (ptr *QTcpServer) DestroyQTcpServer() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DestroyQTcpServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
