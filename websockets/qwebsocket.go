package websockets

//#include "websockets.h"
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
	for len(n.ObjectName()) < len("QWebSocket_") {
		n.SetObjectName("QWebSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebSocket) QWebSocket_PTR() *QWebSocket {
	return ptr
}

func (ptr *QWebSocket) Abort() {
	defer qt.Recovering("QWebSocket::abort")

	if ptr.Pointer() != nil {
		C.QWebSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QWebSocket) ConnectAboutToClose(f func()) {
	defer qt.Recovering("connect QWebSocket::aboutToClose")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectAboutToClose(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToClose", f)
	}
}

func (ptr *QWebSocket) DisconnectAboutToClose() {
	defer qt.Recovering("disconnect QWebSocket::aboutToClose")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectAboutToClose(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToClose")
	}
}

//export callbackQWebSocketAboutToClose
func callbackQWebSocketAboutToClose(ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::aboutToClose")

	var signal = qt.GetSignal(C.GoString(ptrName), "aboutToClose")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectBinaryFrameReceived(f func(frame *core.QByteArray, isLastFrame bool)) {
	defer qt.Recovering("connect QWebSocket::binaryFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryFrameReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "binaryFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryFrameReceived() {
	defer qt.Recovering("disconnect QWebSocket::binaryFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "binaryFrameReceived")
	}
}

//export callbackQWebSocketBinaryFrameReceived
func callbackQWebSocketBinaryFrameReceived(ptrName *C.char, frame unsafe.Pointer, isLastFrame C.int) {
	defer qt.Recovering("callback QWebSocket::binaryFrameReceived")

	var signal = qt.GetSignal(C.GoString(ptrName), "binaryFrameReceived")
	if signal != nil {
		signal.(func(*core.QByteArray, bool))(core.NewQByteArrayFromPointer(frame), int(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectBinaryMessageReceived(f func(message *core.QByteArray)) {
	defer qt.Recovering("connect QWebSocket::binaryMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "binaryMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryMessageReceived() {
	defer qt.Recovering("disconnect QWebSocket::binaryMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "binaryMessageReceived")
	}
}

//export callbackQWebSocketBinaryMessageReceived
func callbackQWebSocketBinaryMessageReceived(ptrName *C.char, message unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::binaryMessageReceived")

	var signal = qt.GetSignal(C.GoString(ptrName), "binaryMessageReceived")
	if signal != nil {
		signal.(func(*core.QByteArray))(core.NewQByteArrayFromPointer(message))
	}

}

func (ptr *QWebSocket) ConnectBytesWritten(f func(bytes int64)) {
	defer qt.Recovering("connect QWebSocket::bytesWritten")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBytesWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "bytesWritten", f)
	}
}

func (ptr *QWebSocket) DisconnectBytesWritten() {
	defer qt.Recovering("disconnect QWebSocket::bytesWritten")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBytesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "bytesWritten")
	}
}

//export callbackQWebSocketBytesWritten
func callbackQWebSocketBytesWritten(ptrName *C.char, bytes C.longlong) {
	defer qt.Recovering("callback QWebSocket::bytesWritten")

	var signal = qt.GetSignal(C.GoString(ptrName), "bytesWritten")
	if signal != nil {
		signal.(func(int64))(int64(bytes))
	}

}

func (ptr *QWebSocket) CloseReason() string {
	defer qt.Recovering("QWebSocket::closeReason")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_CloseReason(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QWebSocket::connected")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QWebSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QWebSocket::connected")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQWebSocketConnected
func callbackQWebSocketConnected(ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::connected")

	var signal = qt.GetSignal(C.GoString(ptrName), "connected")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QWebSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QWebSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QWebSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQWebSocketDisconnected
func callbackQWebSocketDisconnected(ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::disconnected")

	var signal = qt.GetSignal(C.GoString(ptrName), "disconnected")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ConnectError2(f func(error network.QAbstractSocket__SocketError)) {
	defer qt.Recovering("connect QWebSocket::error")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QWebSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QWebSocket::error")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQWebSocketError2
func callbackQWebSocketError2(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QWebSocket::error")

	var signal = qt.GetSignal(C.GoString(ptrName), "error2")
	if signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(error))
	}

}

func (ptr *QWebSocket) Error() network.QAbstractSocket__SocketError {
	defer qt.Recovering("QWebSocket::error")

	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketError(C.QWebSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) ErrorString() string {
	defer qt.Recovering("QWebSocket::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Flush() bool {
	defer qt.Recovering("QWebSocket::flush")

	if ptr.Pointer() != nil {
		return C.QWebSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocket) IgnoreSslErrors() {
	defer qt.Recovering("QWebSocket::ignoreSslErrors")

	if ptr.Pointer() != nil {
		C.QWebSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QWebSocket) IsValid() bool {
	defer qt.Recovering("QWebSocket::isValid")

	if ptr.Pointer() != nil {
		return C.QWebSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocket) MaskGenerator() *QMaskGenerator {
	defer qt.Recovering("QWebSocket::maskGenerator")

	if ptr.Pointer() != nil {
		return NewQMaskGeneratorFromPointer(C.QWebSocket_MaskGenerator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) Open(url core.QUrl_ITF) {
	defer qt.Recovering("QWebSocket::open")

	if ptr.Pointer() != nil {
		C.QWebSocket_Open(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebSocket) Origin() string {
	defer qt.Recovering("QWebSocket::origin")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) PauseMode() network.QAbstractSocket__PauseMode {
	defer qt.Recovering("QWebSocket::pauseMode")

	if ptr.Pointer() != nil {
		return network.QAbstractSocket__PauseMode(C.QWebSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) PeerName() string {
	defer qt.Recovering("QWebSocket::peerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Ping(payload core.QByteArray_ITF) {
	defer qt.Recovering("QWebSocket::ping")

	if ptr.Pointer() != nil {
		C.QWebSocket_Ping(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

func (ptr *QWebSocket) ReadBufferSize() int64 {
	defer qt.Recovering("QWebSocket::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) ConnectReadChannelFinished(f func()) {
	defer qt.Recovering("connect QWebSocket::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectReadChannelFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readChannelFinished", f)
	}
}

func (ptr *QWebSocket) DisconnectReadChannelFinished() {
	defer qt.Recovering("disconnect QWebSocket::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectReadChannelFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readChannelFinished")
	}
}

//export callbackQWebSocketReadChannelFinished
func callbackQWebSocketReadChannelFinished(ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::readChannelFinished")

	var signal = qt.GetSignal(C.GoString(ptrName), "readChannelFinished")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) RequestUrl() *core.QUrl {
	defer qt.Recovering("QWebSocket::requestUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QWebSocket_RequestUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) ResourceName() string {
	defer qt.Recovering("QWebSocket::resourceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ResourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Resume() {
	defer qt.Recovering("QWebSocket::resume")

	if ptr.Pointer() != nil {
		C.QWebSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QWebSocket) SendBinaryMessage(data core.QByteArray_ITF) int64 {
	defer qt.Recovering("QWebSocket::sendBinaryMessage")

	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_SendBinaryMessage(ptr.Pointer(), core.PointerFromQByteArray(data)))
	}
	return 0
}

func (ptr *QWebSocket) SendTextMessage(message string) int64 {
	defer qt.Recovering("QWebSocket::sendTextMessage")

	if ptr.Pointer() != nil {
		return int64(C.QWebSocket_SendTextMessage(ptr.Pointer(), C.CString(message)))
	}
	return 0
}

func (ptr *QWebSocket) SetMaskGenerator(maskGenerator QMaskGenerator_ITF) {
	defer qt.Recovering("QWebSocket::setMaskGenerator")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetMaskGenerator(ptr.Pointer(), PointerFromQMaskGenerator(maskGenerator))
	}
}

func (ptr *QWebSocket) SetPauseMode(pauseMode network.QAbstractSocket__PauseMode) {
	defer qt.Recovering("QWebSocket::setPauseMode")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetPauseMode(ptr.Pointer(), C.int(pauseMode))
	}
}

func (ptr *QWebSocket) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	defer qt.Recovering("QWebSocket::setProxy")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QWebSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QWebSocket) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	defer qt.Recovering("QWebSocket::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QWebSocket_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocket) State() network.QAbstractSocket__SocketState {
	defer qt.Recovering("QWebSocket::state")

	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketState(C.QWebSocket_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) ConnectStateChanged(f func(state network.QAbstractSocket__SocketState)) {
	defer qt.Recovering("connect QWebSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QWebSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QWebSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQWebSocketStateChanged
func callbackQWebSocketStateChanged(ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QWebSocket::stateChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "stateChanged")
	if signal != nil {
		signal.(func(network.QAbstractSocket__SocketState))(network.QAbstractSocket__SocketState(state))
	}

}

func (ptr *QWebSocket) ConnectTextFrameReceived(f func(frame string, isLastFrame bool)) {
	defer qt.Recovering("connect QWebSocket::textFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextFrameReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextFrameReceived() {
	defer qt.Recovering("disconnect QWebSocket::textFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textFrameReceived")
	}
}

//export callbackQWebSocketTextFrameReceived
func callbackQWebSocketTextFrameReceived(ptrName *C.char, frame *C.char, isLastFrame C.int) {
	defer qt.Recovering("callback QWebSocket::textFrameReceived")

	var signal = qt.GetSignal(C.GoString(ptrName), "textFrameReceived")
	if signal != nil {
		signal.(func(string, bool))(C.GoString(frame), int(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) ConnectTextMessageReceived(f func(message string)) {
	defer qt.Recovering("connect QWebSocket::textMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextMessageReceived() {
	defer qt.Recovering("disconnect QWebSocket::textMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textMessageReceived")
	}
}

//export callbackQWebSocketTextMessageReceived
func callbackQWebSocketTextMessageReceived(ptrName *C.char, message *C.char) {
	defer qt.Recovering("callback QWebSocket::textMessageReceived")

	var signal = qt.GetSignal(C.GoString(ptrName), "textMessageReceived")
	if signal != nil {
		signal.(func(string))(C.GoString(message))
	}

}

func (ptr *QWebSocket) DestroyQWebSocket() {
	defer qt.Recovering("QWebSocket::~QWebSocket")

	if ptr.Pointer() != nil {
		C.QWebSocket_DestroyQWebSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
