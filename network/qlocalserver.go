package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLocalServer struct {
	core.QObject
}

type QLocalServer_ITF interface {
	core.QObject_ITF
	QLocalServer_PTR() *QLocalServer
}

func PointerFromQLocalServer(ptr QLocalServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalServer_PTR().Pointer()
	}
	return nil
}

func NewQLocalServerFromPointer(ptr unsafe.Pointer) *QLocalServer {
	var n = new(QLocalServer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QLocalServer_") {
		n.SetObjectName("QLocalServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QLocalServer) QLocalServer_PTR() *QLocalServer {
	return ptr
}

//QLocalServer::SocketOption
type QLocalServer__SocketOption int64

const (
	QLocalServer__NoOptions         = QLocalServer__SocketOption(0x0)
	QLocalServer__UserAccessOption  = QLocalServer__SocketOption(0x01)
	QLocalServer__GroupAccessOption = QLocalServer__SocketOption(0x2)
	QLocalServer__OtherAccessOption = QLocalServer__SocketOption(0x4)
	QLocalServer__WorldAccessOption = QLocalServer__SocketOption(0x7)
)

func (ptr *QLocalServer) SetSocketOptions(options QLocalServer__SocketOption) {
	defer qt.Recovering("QLocalServer::setSocketOptions")

	if ptr.Pointer() != nil {
		C.QLocalServer_SetSocketOptions(ptr.Pointer(), C.int(options))
	}
}

func NewQLocalServer(parent core.QObject_ITF) *QLocalServer {
	defer qt.Recovering("QLocalServer::QLocalServer")

	return NewQLocalServerFromPointer(C.QLocalServer_NewQLocalServer(core.PointerFromQObject(parent)))
}

func (ptr *QLocalServer) Close() {
	defer qt.Recovering("QLocalServer::close")

	if ptr.Pointer() != nil {
		C.QLocalServer_Close(ptr.Pointer())
	}
}

func (ptr *QLocalServer) ErrorString() string {
	defer qt.Recovering("QLocalServer::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) FullServerName() string {
	defer qt.Recovering("QLocalServer::fullServerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_FullServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) HasPendingConnections() bool {
	defer qt.Recovering("QLocalServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QLocalServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalServer) IsListening() bool {
	defer qt.Recovering("QLocalServer::isListening")

	if ptr.Pointer() != nil {
		return C.QLocalServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalServer) Listen(name string) bool {
	defer qt.Recovering("QLocalServer::listen")

	if ptr.Pointer() != nil {
		return C.QLocalServer_Listen(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QLocalServer) MaxPendingConnections() int {
	defer qt.Recovering("QLocalServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(C.QLocalServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QLocalServer::newConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QLocalServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QLocalServer::newConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQLocalServerNewConnection
func callbackQLocalServerNewConnection(ptrName *C.char) {
	defer qt.Recovering("callback QLocalServer::newConnection")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalServer) NextPendingConnection() *QLocalSocket {
	defer qt.Recovering("QLocalServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQLocalSocketFromPointer(C.QLocalServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func QLocalServer_RemoveServer(name string) bool {
	defer qt.Recovering("QLocalServer::removeServer")

	return C.QLocalServer_QLocalServer_RemoveServer(C.CString(name)) != 0
}

func (ptr *QLocalServer) ServerError() QAbstractSocket__SocketError {
	defer qt.Recovering("QLocalServer::serverError")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QLocalServer_ServerError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) ServerName() string {
	defer qt.Recovering("QLocalServer::serverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QLocalServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QLocalServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

func (ptr *QLocalServer) SocketOptions() QLocalServer__SocketOption {
	defer qt.Recovering("QLocalServer::socketOptions")

	if ptr.Pointer() != nil {
		return QLocalServer__SocketOption(C.QLocalServer_SocketOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) WaitForNewConnection(msec int, timedOut bool) bool {
	defer qt.Recovering("QLocalServer::waitForNewConnection")

	if ptr.Pointer() != nil {
		return C.QLocalServer_WaitForNewConnection(ptr.Pointer(), C.int(msec), C.int(qt.GoBoolToInt(timedOut))) != 0
	}
	return false
}

func (ptr *QLocalServer) DestroyQLocalServer() {
	defer qt.Recovering("QLocalServer::~QLocalServer")

	if ptr.Pointer() != nil {
		C.QLocalServer_DestroyQLocalServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
