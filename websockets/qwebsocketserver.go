package websockets

//#include "qwebsocketserver.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QWebSocketServer struct {
	core.QObject
}

type QWebSocketServerITF interface {
	core.QObjectITF
	QWebSocketServerPTR() *QWebSocketServer
}

func PointerFromQWebSocketServer(ptr QWebSocketServerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketServerPTR().Pointer()
	}
	return nil
}

func QWebSocketServerFromPointer(ptr unsafe.Pointer) *QWebSocketServer {
	var n = new(QWebSocketServer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWebSocketServer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWebSocketServer) QWebSocketServerPTR() *QWebSocketServer {
	return ptr
}

//QWebSocketServer::SslMode
type QWebSocketServer__SslMode int

var (
	QWebSocketServer__SecureMode    = QWebSocketServer__SslMode(0)
	QWebSocketServer__NonSecureMode = QWebSocketServer__SslMode(1)
)

func NewQWebSocketServer(serverName string, secureMode QWebSocketServer__SslMode, parent core.QObjectITF) *QWebSocketServer {
	return QWebSocketServerFromPointer(unsafe.Pointer(C.QWebSocketServer_NewQWebSocketServer(C.CString(serverName), C.int(secureMode), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QWebSocketServer) ConnectAcceptError(f func(socketError network.QAbstractSocket__SocketError)) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectAcceptError(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "acceptError", f)
	}
}

func (ptr *QWebSocketServer) DisconnectAcceptError() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectAcceptError(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "acceptError")
	}
}

//export callbackQWebSocketServerAcceptError
func callbackQWebSocketServerAcceptError(ptrName *C.char, socketError C.int) {
	qt.GetSignal(C.GoString(ptrName), "acceptError").(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(socketError))
}

func (ptr *QWebSocketServer) Close() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWebSocketServer) ConnectClosed(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectClosed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "closed", f)
	}
}

func (ptr *QWebSocketServer) DisconnectClosed() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectClosed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "closed")
	}
}

//export callbackQWebSocketServerClosed
func callbackQWebSocketServerClosed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "closed").(func())()
}

func (ptr *QWebSocketServer) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketServer_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocketServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_HasPendingConnections(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebSocketServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_IsListening(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebSocketServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(C.QWebSocketServer_MaxPendingConnections(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebSocketServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectNewConnection(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QWebSocketServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNewConnection(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQWebSocketServerNewConnection
func callbackQWebSocketServerNewConnection(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "newConnection").(func())()
}

func (ptr *QWebSocketServer) NextPendingConnection() *QWebSocket {
	if ptr.Pointer() != nil {
		return QWebSocketFromPointer(unsafe.Pointer(C.QWebSocketServer_NextPendingConnection(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWebSocketServer) ConnectOriginAuthenticationRequired(f func(authenticator QWebSocketCorsAuthenticatorITF)) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectOriginAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "originAuthenticationRequired", f)
	}
}

func (ptr *QWebSocketServer) DisconnectOriginAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectOriginAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "originAuthenticationRequired")
	}
}

//export callbackQWebSocketServerOriginAuthenticationRequired
func callbackQWebSocketServerOriginAuthenticationRequired(ptrName *C.char, authenticator unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "originAuthenticationRequired").(func(*QWebSocketCorsAuthenticator))(QWebSocketCorsAuthenticatorFromPointer(authenticator))
}

func (ptr *QWebSocketServer) PauseAccepting() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_PauseAccepting(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWebSocketServer) ResumeAccepting() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_ResumeAccepting(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWebSocketServer) SecureMode() QWebSocketServer__SslMode {
	if ptr.Pointer() != nil {
		return QWebSocketServer__SslMode(C.QWebSocketServer_SecureMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebSocketServer) ServerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketServer_ServerName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocketServer) ServerUrl() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketServer_ServerUrl(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocketServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetMaxPendingConnections(C.QtObjectPtr(ptr.Pointer()), C.int(numConnections))
	}
}

func (ptr *QWebSocketServer) SetProxy(networkProxy network.QNetworkProxyITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetProxy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(network.PointerFromQNetworkProxy(networkProxy)))
	}
}

func (ptr *QWebSocketServer) SetServerName(serverName string) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetServerName(C.QtObjectPtr(ptr.Pointer()), C.CString(serverName))
	}
}

func (ptr *QWebSocketServer) SetSocketDescriptor(socketDescriptor int) bool {
	if ptr.Pointer() != nil {
		return C.QWebSocketServer_SetSocketDescriptor(C.QtObjectPtr(ptr.Pointer()), C.int(socketDescriptor)) != 0
	}
	return false
}

func (ptr *QWebSocketServer) SetSslConfiguration(sslConfiguration network.QSslConfigurationITF) {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetSslConfiguration(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(network.PointerFromQSslConfiguration(sslConfiguration)))
	}
}

func (ptr *QWebSocketServer) SocketDescriptor() int {
	if ptr.Pointer() != nil {
		return int(C.QWebSocketServer_SocketDescriptor(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebSocketServer) DestroyQWebSocketServer() {
	if ptr.Pointer() != nil {
		C.QWebSocketServer_DestroyQWebSocketServer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
