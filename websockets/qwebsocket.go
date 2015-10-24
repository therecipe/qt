package websockets

//#include "qwebsocket.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QWebSocket struct {
	core.QObject
}

type QWebSocketITF interface {
	core.QObjectITF
	QWebSocketPTR() *QWebSocket
}

func PointerFromQWebSocket(ptr QWebSocketITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketPTR().Pointer()
	}
	return nil
}

func QWebSocketFromPointer(ptr unsafe.Pointer) *QWebSocket {
	var n = new(QWebSocket)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWebSocket_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWebSocket) QWebSocketPTR() *QWebSocket {
	return ptr
}

func (ptr *QWebSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Abort(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWebSocket) ConnectAboutToClose(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectAboutToClose(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "aboutToClose", f)
	}
}

func (ptr *QWebSocket) DisconnectAboutToClose() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectAboutToClose(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToClose")
	}
}

//export callbackQWebSocketAboutToClose
func callbackQWebSocketAboutToClose(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToClose").(func())()
}

func (ptr *QWebSocket) CloseReason() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_CloseReason(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectConnected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QWebSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectConnected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQWebSocketConnected
func callbackQWebSocketConnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "connected").(func())()
}

func (ptr *QWebSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectDisconnected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QWebSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectDisconnected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQWebSocketDisconnected
func callbackQWebSocketDisconnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QWebSocket) Error() network.QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketError(C.QWebSocket_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebSocket) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_Flush(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebSocket) IgnoreSslErrors() {
	if ptr.Pointer() != nil {
		C.QWebSocket_IgnoreSslErrors(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWebSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWebSocket) MaskGenerator() *QMaskGenerator {
	if ptr.Pointer() != nil {
		return QMaskGeneratorFromPointer(unsafe.Pointer(C.QWebSocket_MaskGenerator(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWebSocket) Open(url string) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Open(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QWebSocket) Origin() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_Origin(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocket) PauseMode() network.QAbstractSocket__PauseMode {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__PauseMode(C.QWebSocket_PauseMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWebSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_PeerName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocket) Ping(payload core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Ping(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(payload)))
	}
}

func (ptr *QWebSocket) ConnectReadChannelFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectReadChannelFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "readChannelFinished", f)
	}
}

func (ptr *QWebSocket) DisconnectReadChannelFinished() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectReadChannelFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "readChannelFinished")
	}
}

//export callbackQWebSocketReadChannelFinished
func callbackQWebSocketReadChannelFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readChannelFinished").(func())()
}

func (ptr *QWebSocket) RequestUrl() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_RequestUrl(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocket) ResourceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ResourceName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWebSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Resume(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWebSocket) SetMaskGenerator(maskGenerator QMaskGeneratorITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetMaskGenerator(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMaskGenerator(maskGenerator)))
	}
}

func (ptr *QWebSocket) SetPauseMode(pauseMode network.QAbstractSocket__PauseMode) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetPauseMode(C.QtObjectPtr(ptr.Pointer()), C.int(pauseMode))
	}
}

func (ptr *QWebSocket) SetProxy(networkProxy network.QNetworkProxyITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetProxy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(network.PointerFromQNetworkProxy(networkProxy)))
	}
}

func (ptr *QWebSocket) SetSslConfiguration(sslConfiguration network.QSslConfigurationITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetSslConfiguration(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(network.PointerFromQSslConfiguration(sslConfiguration)))
	}
}

func (ptr *QWebSocket) ConnectStateChanged(f func(state network.QAbstractSocket__SocketState)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QWebSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQWebSocketStateChanged
func callbackQWebSocketStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(network.QAbstractSocket__SocketState))(network.QAbstractSocket__SocketState(state))
}

func (ptr *QWebSocket) ConnectTextFrameReceived(f func(frame string, isLastFrame bool)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextFrameReceived(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextFrameReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextFrameReceived(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textFrameReceived")
	}
}

//export callbackQWebSocketTextFrameReceived
func callbackQWebSocketTextFrameReceived(ptrName *C.char, frame *C.char, isLastFrame C.int) {
	qt.GetSignal(C.GoString(ptrName), "textFrameReceived").(func(string, bool))(C.GoString(frame), int(isLastFrame) != 0)
}

func (ptr *QWebSocket) ConnectTextMessageReceived(f func(message string)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextMessageReceived(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "textMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextMessageReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextMessageReceived(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "textMessageReceived")
	}
}

//export callbackQWebSocketTextMessageReceived
func callbackQWebSocketTextMessageReceived(ptrName *C.char, message *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textMessageReceived").(func(string))(C.GoString(message))
}

func (ptr *QWebSocket) DestroyQWebSocket() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DestroyQWebSocket(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
