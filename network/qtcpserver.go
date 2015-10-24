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

type QTcpServerITF interface {
	core.QObjectITF
	QTcpServerPTR() *QTcpServer
}

func PointerFromQTcpServer(ptr QTcpServerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpServerPTR().Pointer()
	}
	return nil
}

func QTcpServerFromPointer(ptr unsafe.Pointer) *QTcpServer {
	var n = new(QTcpServer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTcpServer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTcpServer) QTcpServerPTR() *QTcpServer {
	return ptr
}

func NewQTcpServer(parent core.QObjectITF) *QTcpServer {
	return QTcpServerFromPointer(unsafe.Pointer(C.QTcpServer_NewQTcpServer(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QTcpServer) ConnectAcceptError(f func(socketError QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectAcceptError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "acceptError", f)
	}
}

func (ptr *QTcpServer) DisconnectAcceptError() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectAcceptError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "acceptError")
	}
}

//export callbackQTcpServerAcceptError
func callbackQTcpServerAcceptError(ptrName *C.char, socketError C.int) {
	qt.GetSignal(C.GoString(ptrName), "acceptError").(func(QAbstractSocket__SocketError))(QAbstractSocket__SocketError(socketError))
}

func (ptr *QTcpServer) Close() {
	if ptr.Pointer() != nil {
		C.QTcpServer_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTcpServer) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTcpServer_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTcpServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_HasPendingConnections(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTcpServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_IsListening(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTcpServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(C.QTcpServer_MaxPendingConnections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTcpServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNewConnection(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QTcpServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNewConnection(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQTcpServerNewConnection
func callbackQTcpServerNewConnection(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "newConnection").(func())()
}

func (ptr *QTcpServer) NextPendingConnection() *QTcpSocket {
	if ptr.Pointer() != nil {
		return QTcpSocketFromPointer(unsafe.Pointer(C.QTcpServer_NextPendingConnection(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTcpServer) PauseAccepting() {
	if ptr.Pointer() != nil {
		C.QTcpServer_PauseAccepting(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTcpServer) ResumeAccepting() {
	if ptr.Pointer() != nil {
		C.QTcpServer_ResumeAccepting(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTcpServer) ServerError() QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QTcpServer_ServerError(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTcpServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QTcpServer_SetMaxPendingConnections(C.QtObjectPtr(ptr.Pointer()), C.int(numConnections))
	}
}

func (ptr *QTcpServer) SetProxy(networkProxy QNetworkProxyITF) {
	if ptr.Pointer() != nil {
		C.QTcpServer_SetProxy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkProxy(networkProxy)))
	}
}

func (ptr *QTcpServer) WaitForNewConnection(msec int, timedOut bool) bool {
	if ptr.Pointer() != nil {
		return C.QTcpServer_WaitForNewConnection(C.QtObjectPtr(ptr.Pointer()), C.int(msec), C.int(qt.GoBoolToInt(timedOut))) != 0
	}
	return false
}

func (ptr *QTcpServer) DestroyQTcpServer() {
	if ptr.Pointer() != nil {
		C.QTcpServer_DestroyQTcpServer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
