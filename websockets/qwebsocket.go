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

type QWebSocket_ITF interface {
	core.QObject_ITF
	QWebSocket_PTR() *QWebSocket
}

func PointerFromQWebSocket(ptr QWebSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocket_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketFromPointer(ptr unsafe.Pointer) *QWebSocket {
	var n = new(QWebSocket)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWebSocket_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWebSocket) QWebSocket_PTR() *QWebSocket {
	return ptr
}

func (ptr *QWebSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QWebSocket) ConnectAboutToClose(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectAboutToClose(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToClose", f)
	}
}

func (ptr *QWebSocket) DisconnectAboutToClose() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectAboutToClose(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToClose")
	}
}

//export callbackQWebSocketAboutToClose
func callbackQWebSocketAboutToClose(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToClose").(func())()
}

func (ptr *QWebSocket) ConnectBinaryFrameReceived(f func(frame *core.QByteArray, isLastFrame bool)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryFrameReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "binaryFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryFrameReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "binaryFrameReceived")
	}
}

//export callbackQWebSocketBinaryFrameReceived
func callbackQWebSocketBinaryFrameReceived(ptrName *C.char, frame unsafe.Pointer, isLastFrame C.int) {
	qt.GetSignal(C.GoString(ptrName), "binaryFrameReceived").(func(*core.QByteArray, bool))(core.NewQByteArrayFromPointer(frame), int(isLastFrame) != 0)
}

func (ptr *QWebSocket) ConnectBinaryMessageReceived(f func(message *core.QByteArray)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "binaryMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryMessageReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "binaryMessageReceived")
	}
}

//export callbackQWebSocketBinaryMessageReceived
func callbackQWebSocketBinaryMessageReceived(ptrName *C.char, message unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "binaryMessageReceived").(func(*core.QByteArray))(core.NewQByteArrayFromPointer(message))
}

func (ptr *QWebSocket) CloseReason() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_CloseReason(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QWebSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQWebSocketConnected
func callbackQWebSocketConnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "connected").(func())()
}

func (ptr *QWebSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QWebSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQWebSocketDisconnected
func callbackQWebSocketDisconnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QWebSocket) Error() network.QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketError(C.QWebSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocket) IgnoreSslErrors() {
	if ptr.Pointer() != nil {
		C.QWebSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QWebSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QWebSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocket) MaskGenerator() *QMaskGenerator {
	if ptr.Pointer() != nil {
		return NewQMaskGeneratorFromPointer(C.QWebSocket_MaskGenerator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) Open(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Open(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebSocket) Origin() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) PauseMode() network.QAbstractSocket__PauseMode {
	if ptr.Pointer() != nil {
		return network.QAbstractSocket__PauseMode(C.QWebSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Ping(payload core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_Ping(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

func (ptr *QWebSocket) ConnectReadChannelFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectReadChannelFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readChannelFinished", f)
	}
}

func (ptr *QWebSocket) DisconnectReadChannelFinished() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectReadChannelFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readChannelFinished")
	}
}

//export callbackQWebSocketReadChannelFinished
func callbackQWebSocketReadChannelFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readChannelFinished").(func())()
}

func (ptr *QWebSocket) ResourceName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ResourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QWebSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QWebSocket) SetMaskGenerator(maskGenerator QMaskGenerator_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetMaskGenerator(ptr.Pointer(), PointerFromQMaskGenerator(maskGenerator))
	}
}

func (ptr *QWebSocket) SetPauseMode(pauseMode network.QAbstractSocket__PauseMode) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetPauseMode(ptr.Pointer(), C.int(pauseMode))
	}
}

func (ptr *QWebSocket) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocket) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QWebSocket_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocket) ConnectStateChanged(f func(state network.QAbstractSocket__SocketState)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QWebSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQWebSocketStateChanged
func callbackQWebSocketStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(network.QAbstractSocket__SocketState))(network.QAbstractSocket__SocketState(state))
}

func (ptr *QWebSocket) ConnectTextFrameReceived(f func(frame string, isLastFrame bool)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextFrameReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextFrameReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textFrameReceived")
	}
}

//export callbackQWebSocketTextFrameReceived
func callbackQWebSocketTextFrameReceived(ptrName *C.char, frame *C.char, isLastFrame C.int) {
	qt.GetSignal(C.GoString(ptrName), "textFrameReceived").(func(string, bool))(C.GoString(frame), int(isLastFrame) != 0)
}

func (ptr *QWebSocket) ConnectTextMessageReceived(f func(message string)) {
	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextMessageReceived() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textMessageReceived")
	}
}

//export callbackQWebSocketTextMessageReceived
func callbackQWebSocketTextMessageReceived(ptrName *C.char, message *C.char) {
	qt.GetSignal(C.GoString(ptrName), "textMessageReceived").(func(string))(C.GoString(message))
}

func (ptr *QWebSocket) DestroyQWebSocket() {
	if ptr.Pointer() != nil {
		C.QWebSocket_DestroyQWebSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
