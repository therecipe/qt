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
func callbackQWebSocketServerAcceptError(ptr unsafe.Pointer, ptrName *C.char, socketError C.int) {
	defer qt.Recovering("callback QWebSocketServer::acceptError")

	if signal := qt.GetSignal(C.GoString(ptrName), "acceptError"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QWebSocketServer) AcceptError(socketError network.QAbstractSocket__SocketError) {
	defer qt.Recovering("QWebSocketServer::acceptError")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_AcceptError(ptr.Pointer(), C.int(socketError))
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
func callbackQWebSocketServerClosed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocketServer::closed")

	if signal := qt.GetSignal(C.GoString(ptrName), "closed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) Closed() {
	defer qt.Recovering("QWebSocketServer::closed")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_Closed(ptr.Pointer())
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
func callbackQWebSocketServerNewConnection(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocketServer::newConnection")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocketServer) NewConnection() {
	defer qt.Recovering("QWebSocketServer::newConnection")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_NewConnection(ptr.Pointer())
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
func callbackQWebSocketServerOriginAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::originAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "originAuthenticationRequired"); signal != nil {
		signal.(func(*QWebSocketCorsAuthenticator))(NewQWebSocketCorsAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QWebSocketServer) OriginAuthenticationRequired(authenticator QWebSocketCorsAuthenticator_ITF) {
	defer qt.Recovering("QWebSocketServer::originAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_OriginAuthenticationRequired(ptr.Pointer(), PointerFromQWebSocketCorsAuthenticator(authenticator))
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

func (ptr *QWebSocketServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebSocketServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebSocketServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebSocketServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQWebSocketServerTimerEvent
func callbackQWebSocketServerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebSocketServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebSocketServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebSocketServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebSocketServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebSocketServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQWebSocketServerChildEvent
func callbackQWebSocketServerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::childEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebSocketServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::childEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebSocketServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebSocketServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebSocketServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebSocketServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQWebSocketServerCustomEvent
func callbackQWebSocketServerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocketServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebSocketServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebSocketServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::customEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebSocketServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebSocketServer::customEvent")

	if ptr.Pointer() != nil {
		C.QWebSocketServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
