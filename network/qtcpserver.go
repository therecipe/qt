package network

//#include "network.h"
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
	for len(n.ObjectName()) < len("QTcpServer_") {
		n.SetObjectName("QTcpServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QTcpServer) QTcpServer_PTR() *QTcpServer {
	return ptr
}

func NewQTcpServer(parent core.QObject_ITF) *QTcpServer {
	defer qt.Recovering("QTcpServer::QTcpServer")

	return NewQTcpServerFromPointer(C.QTcpServer_NewQTcpServer(core.PointerFromQObject(parent)))
}

func (ptr *QTcpServer) ConnectAcceptError(f func(socketError QAbstractSocket__SocketError)) {
	defer qt.Recovering("connect QTcpServer::acceptError")

	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectAcceptError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "acceptError", f)
	}
}

func (ptr *QTcpServer) DisconnectAcceptError() {
	defer qt.Recovering("disconnect QTcpServer::acceptError")

	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "acceptError")
	}
}

//export callbackQTcpServerAcceptError
func callbackQTcpServerAcceptError(ptrName *C.char, socketError C.int) {
	defer qt.Recovering("callback QTcpServer::acceptError")

	var signal = qt.GetSignal(C.GoString(ptrName), "acceptError")
	if signal != nil {
		signal.(func(QAbstractSocket__SocketError))(QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QTcpServer) Close() {
	defer qt.Recovering("QTcpServer::close")

	if ptr.Pointer() != nil {
		C.QTcpServer_Close(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ErrorString() string {
	defer qt.Recovering("QTcpServer::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTcpServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTcpServer) HasPendingConnections() bool {
	defer qt.Recovering("QTcpServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QTcpServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) IsListening() bool {
	defer qt.Recovering("QTcpServer::isListening")

	if ptr.Pointer() != nil {
		return C.QTcpServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) MaxPendingConnections() int {
	defer qt.Recovering("QTcpServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(C.QTcpServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QTcpServer::newConnection")

	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QTcpServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QTcpServer::newConnection")

	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQTcpServerNewConnection
func callbackQTcpServerNewConnection(ptrName *C.char) {
	defer qt.Recovering("callback QTcpServer::newConnection")

	var signal = qt.GetSignal(C.GoString(ptrName), "newConnection")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QTcpServer) NextPendingConnection() *QTcpSocket {
	defer qt.Recovering("QTcpServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQTcpSocketFromPointer(C.QTcpServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTcpServer) PauseAccepting() {
	defer qt.Recovering("QTcpServer::pauseAccepting")

	if ptr.Pointer() != nil {
		C.QTcpServer_PauseAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ResumeAccepting() {
	defer qt.Recovering("QTcpServer::resumeAccepting")

	if ptr.Pointer() != nil {
		C.QTcpServer_ResumeAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ServerError() QAbstractSocket__SocketError {
	defer qt.Recovering("QTcpServer::serverError")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QTcpServer_ServerError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QTcpServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QTcpServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

func (ptr *QTcpServer) SetProxy(networkProxy QNetworkProxy_ITF) {
	defer qt.Recovering("QTcpServer::setProxy")

	if ptr.Pointer() != nil {
		C.QTcpServer_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QTcpServer) WaitForNewConnection(msec int, timedOut bool) bool {
	defer qt.Recovering("QTcpServer::waitForNewConnection")

	if ptr.Pointer() != nil {
		return C.QTcpServer_WaitForNewConnection(ptr.Pointer(), C.int(msec), C.int(qt.GoBoolToInt(timedOut))) != 0
	}
	return false
}

func (ptr *QTcpServer) DestroyQTcpServer() {
	defer qt.Recovering("QTcpServer::~QTcpServer")

	if ptr.Pointer() != nil {
		C.QTcpServer_DestroyQTcpServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
