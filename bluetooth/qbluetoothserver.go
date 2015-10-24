package bluetooth

//#include "qbluetoothserver.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothServer struct {
	core.QObject
}

type QBluetoothServerITF interface {
	core.QObjectITF
	QBluetoothServerPTR() *QBluetoothServer
}

func PointerFromQBluetoothServer(ptr QBluetoothServerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServerPTR().Pointer()
	}
	return nil
}

func QBluetoothServerFromPointer(ptr unsafe.Pointer) *QBluetoothServer {
	var n = new(QBluetoothServer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QBluetoothServer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothServer) QBluetoothServerPTR() *QBluetoothServer {
	return ptr
}

//QBluetoothServer::Error
type QBluetoothServer__Error int

var (
	QBluetoothServer__NoError                       = QBluetoothServer__Error(0)
	QBluetoothServer__UnknownError                  = QBluetoothServer__Error(1)
	QBluetoothServer__PoweredOffError               = QBluetoothServer__Error(2)
	QBluetoothServer__InputOutputError              = QBluetoothServer__Error(3)
	QBluetoothServer__ServiceAlreadyRegisteredError = QBluetoothServer__Error(4)
	QBluetoothServer__UnsupportedProtocolError      = QBluetoothServer__Error(5)
)

func NewQBluetoothServer(serverType QBluetoothServiceInfo__Protocol, parent core.QObjectITF) *QBluetoothServer {
	return QBluetoothServerFromPointer(unsafe.Pointer(C.QBluetoothServer_NewQBluetoothServer(C.int(serverType), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QBluetoothServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNewConnection(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QBluetoothServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNewConnection(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQBluetoothServerNewConnection
func callbackQBluetoothServerNewConnection(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "newConnection").(func())()
}

func (ptr *QBluetoothServer) Error() QBluetoothServer__Error {
	if ptr.Pointer() != nil {
		return QBluetoothServer__Error(C.QBluetoothServer_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_IsListening(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(C.QBluetoothServer_MaxPendingConnections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerType() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServer_ServerType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServer) DestroyQBluetoothServer() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DestroyQBluetoothServer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) Close() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBluetoothServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_HasPendingConnections(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServer) NextPendingConnection() *QBluetoothSocket {
	if ptr.Pointer() != nil {
		return QBluetoothSocketFromPointer(unsafe.Pointer(C.QBluetoothServer_NextPendingConnection(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QBluetoothServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_SetMaxPendingConnections(C.QtObjectPtr(ptr.Pointer()), C.int(numConnections))
	}
}
