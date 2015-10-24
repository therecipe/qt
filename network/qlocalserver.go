package network

//#include "qlocalserver.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QLocalServer struct {
	core.QObject
}

type QLocalServerITF interface {
	core.QObjectITF
	QLocalServerPTR() *QLocalServer
}

func PointerFromQLocalServer(ptr QLocalServerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalServerPTR().Pointer()
	}
	return nil
}

func QLocalServerFromPointer(ptr unsafe.Pointer) *QLocalServer {
	var n = new(QLocalServer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLocalServer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLocalServer) QLocalServerPTR() *QLocalServer {
	return ptr
}

//QLocalServer::SocketOption
type QLocalServer__SocketOption int

var (
	QLocalServer__NoOptions         = QLocalServer__SocketOption(0x0)
	QLocalServer__UserAccessOption  = QLocalServer__SocketOption(0x01)
	QLocalServer__GroupAccessOption = QLocalServer__SocketOption(0x2)
	QLocalServer__OtherAccessOption = QLocalServer__SocketOption(0x4)
	QLocalServer__WorldAccessOption = QLocalServer__SocketOption(0x7)
)

func (ptr *QLocalServer) SetSocketOptions(options QLocalServer__SocketOption) {
	if ptr.Pointer() != nil {
		C.QLocalServer_SetSocketOptions(C.QtObjectPtr(ptr.Pointer()), C.int(options))
	}
}

func NewQLocalServer(parent core.QObjectITF) *QLocalServer {
	return QLocalServerFromPointer(unsafe.Pointer(C.QLocalServer_NewQLocalServer(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QLocalServer) Close() {
	if ptr.Pointer() != nil {
		C.QLocalServer_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLocalServer) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLocalServer) FullServerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_FullServerName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLocalServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_HasPendingConnections(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_IsListening(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLocalServer) Listen(name string) bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_Listen(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QLocalServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(C.QLocalServer_MaxPendingConnections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLocalServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNewConnection(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QLocalServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNewConnection(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQLocalServerNewConnection
func callbackQLocalServerNewConnection(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "newConnection").(func())()
}

func (ptr *QLocalServer) NextPendingConnection() *QLocalSocket {
	if ptr.Pointer() != nil {
		return QLocalSocketFromPointer(unsafe.Pointer(C.QLocalServer_NextPendingConnection(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func QLocalServer_RemoveServer(name string) bool {
	return C.QLocalServer_QLocalServer_RemoveServer(C.CString(name)) != 0
}

func (ptr *QLocalServer) ServerError() QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QLocalServer_ServerError(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLocalServer) ServerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_ServerName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLocalServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QLocalServer_SetMaxPendingConnections(C.QtObjectPtr(ptr.Pointer()), C.int(numConnections))
	}
}

func (ptr *QLocalServer) SocketOptions() QLocalServer__SocketOption {
	if ptr.Pointer() != nil {
		return QLocalServer__SocketOption(C.QLocalServer_SocketOptions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLocalServer) WaitForNewConnection(msec int, timedOut bool) bool {
	if ptr.Pointer() != nil {
		return C.QLocalServer_WaitForNewConnection(C.QtObjectPtr(ptr.Pointer()), C.int(msec), C.int(qt.GoBoolToInt(timedOut))) != 0
	}
	return false
}

func (ptr *QLocalServer) DestroyQLocalServer() {
	if ptr.Pointer() != nil {
		C.QLocalServer_DestroyQLocalServer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
