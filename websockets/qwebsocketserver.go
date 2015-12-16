package websockets

//#include "websockets.h"
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

type QWebSocketServer_ITF interface {
	core.QObject_ITF
	QWebSocketServer_PTR() *QWebSocketServer
}

func PointerFromQWebSocketServer(ptr QWebSocketServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketServer_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketServerFromPointer(ptr unsafe.Pointer) *QWebSocketServer {
	var n = new(QWebSocketServer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWebSocketServer_") {
		n.SetObjectName("QWebSocketServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebSocketServer) QWebSocketServer_PTR() *QWebSocketServer {
	return ptr
}

//QWebSocketServer::SslMode
type QWebSocketServer__SslMode int64

const (
	QWebSocketServer__SecureMode    = QWebSocketServer__SslMode(0)
	QWebSocketServer__NonSecureMode = QWebSocketServer__SslMode(1)
)

func NewQWebSocketServer(serverName string, secureMode QWebSocketServer__SslMode, parent core.QObject_ITF) *QWebSocketServer {
	defer qt.Recovering("QWebSocketServer::QWebSocketServer")

	return NewQWebSocketServerFromPointer(C.QWebSocketServer_NewQWebSocketServer(C.CString(serverName), C.int(secureMode), core.PointerFromQObject(parent)))
}

func (ptr *QWebSocketServer) ConnectAcceptError(f func(socketError network.QAbstractSocket__SocketError)) {
	defer qt.Recovering("connect QWebSocketServer::acceptError")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectAcceptError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "acceptError", f)
	}
}

func (ptr *QWebSocketServer) DisconnectAcceptError() {
	defer qt.Recovering("disconnect QWebSocketServer::acceptError")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "acceptError")
	}
}

//export callbackQWebSocketServerAcceptError
func callbackQWebSocketServerAcceptError(ptrName *C.char, socketError C.int) {
	defer qt.Recovering("callback QWebSocketServer::acceptError")

	var signal = qt.GetSignal(C.GoString(ptrName), "acceptError")
	if signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QWebSocketServer) Close() {
	defer qt.Recovering("QWebSocketServer::close")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_Close(ptr.Pointer())
	}
}

func (ptr *QWebSocketServer) ConnectClosed(f func()) {
	defer qt.Recovering("connect QWebSocketServer::closed")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectClosed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "closed", f)
	}
}

func (ptr *QWebSocketServer) DisconnectClosed() {
	defer qt.Recovering("disconnect QWebSocketServer::closed")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "closed")
	}
}

//export callbackQWebSocketServerClosed
func callbackQWebSocketServerClosed(ptrName *C.char) {
	defer qt.Recovering("callback QWebSocketServer::closed")

	var signal = qt.GetSignal(C.GoString(ptrName), "closed")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) ErrorString() string {
	defer qt.Recovering("QWebSocketServer::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) HasPendingConnections() bool {
	defer qt.Recovering("QWebSocketServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketServer) IsListening() bool {
	defer qt.Recovering("QWebSocketServer::isListening")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketServer) MaxPendingConnections() int {
	defer qt.Recovering("QWebSocketServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(C.QWebSocketServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QWebSocketServer::newConnection")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QWebSocketServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QWebSocketServer::newConnection")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQWebSocketServerNewConnection
func callbackQWebSocketServerNewConnection(ptrName *C.char) {
	defer qt.Recovering("callback QWebSocketServer::newConnection")

	var signal = qt.GetSignal(C.GoString(ptrName), "newConnection")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) NextPendingConnection() *QWebSocket {
	defer qt.Recovering("QWebSocketServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQWebSocketFromPointer(C.QWebSocketServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocketServer) ConnectOriginAuthenticationRequired(f func(authenticator *QWebSocketCorsAuthenticator)) {
	defer qt.Recovering("connect QWebSocketServer::originAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ConnectOriginAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "originAuthenticationRequired", f)
	}
}

func (ptr *QWebSocketServer) DisconnectOriginAuthenticationRequired() {
	defer qt.Recovering("disconnect QWebSocketServer::originAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DisconnectOriginAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "originAuthenticationRequired")
	}
}

//export callbackQWebSocketServerOriginAuthenticationRequired
func callbackQWebSocketServerOriginAuthenticationRequired(ptrName *C.char, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::originAuthenticationRequired")

	var signal = qt.GetSignal(C.GoString(ptrName), "originAuthenticationRequired")
	if signal != nil {
		signal.(func(*QWebSocketCorsAuthenticator))(NewQWebSocketCorsAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocketServer) PauseAccepting() {
	defer qt.Recovering("QWebSocketServer::pauseAccepting")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_PauseAccepting(ptr.Pointer())
	}
}

func (ptr *QWebSocketServer) ResumeAccepting() {
	defer qt.Recovering("QWebSocketServer::resumeAccepting")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ResumeAccepting(ptr.Pointer())
	}
}

func (ptr *QWebSocketServer) SecureMode() QWebSocketServer__SslMode {
	defer qt.Recovering("QWebSocketServer::secureMode")

	if ptr.Pointer() != nil {
		return QWebSocketServer__SslMode(C.QWebSocketServer_SecureMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) ServerName() string {
	defer qt.Recovering("QWebSocketServer::serverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketServer_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketServer) ServerUrl() *core.QUrl {
	defer qt.Recovering("QWebSocketServer::serverUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebSocketServer_ServerUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocketServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QWebSocketServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

func (ptr *QWebSocketServer) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	defer qt.Recovering("QWebSocketServer::setProxy")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocketServer) SetServerName(serverName string) {
	defer qt.Recovering("QWebSocketServer::setServerName")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetServerName(ptr.Pointer(), C.CString(serverName))
	}
}

func (ptr *QWebSocketServer) SetSocketDescriptor(socketDescriptor int) bool {
	defer qt.Recovering("QWebSocketServer::setSocketDescriptor")

	if ptr.Pointer() != nil {
		return C.QWebSocketServer_SetSocketDescriptor(ptr.Pointer(), C.int(socketDescriptor)) != 0
	}
	return false
}

func (ptr *QWebSocketServer) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	defer qt.Recovering("QWebSocketServer::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocketServer) SocketDescriptor() int {
	defer qt.Recovering("QWebSocketServer::socketDescriptor")

	if ptr.Pointer() != nil {
		return int(C.QWebSocketServer_SocketDescriptor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocketServer) DestroyQWebSocketServer() {
	defer qt.Recovering("QWebSocketServer::~QWebSocketServer")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_DestroyQWebSocketServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
