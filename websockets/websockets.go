package websockets

//#include "websockets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"unsafe"
)

type QMaskGenerator struct {
	core.QObject
}

type QMaskGenerator_ITF interface {
	core.QObject_ITF
	QMaskGenerator_PTR() *QMaskGenerator
}

func PointerFromQMaskGenerator(ptr QMaskGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMaskGenerator_PTR().Pointer()
	}
	return nil
}

func NewQMaskGeneratorFromPointer(ptr unsafe.Pointer) *QMaskGenerator {
	var n = new(QMaskGenerator)
	n.SetPointer(ptr)
	return n
}

func newQMaskGeneratorFromPointer(ptr unsafe.Pointer) *QMaskGenerator {
	var n = NewQMaskGeneratorFromPointer(ptr)
	for len(n.ObjectName()) < len("QMaskGenerator_") {
		n.SetObjectName("QMaskGenerator_" + qt.Identifier())
	}
	return n
}

func (ptr *QMaskGenerator) QMaskGenerator_PTR() *QMaskGenerator {
	return ptr
}

func (ptr *QMaskGenerator) Seed() bool {
	defer qt.Recovering("QMaskGenerator::seed")

	if ptr.Pointer() != nil {
		return C.QMaskGenerator_Seed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMaskGenerator) DestroyQMaskGenerator() {
	defer qt.Recovering("QMaskGenerator::~QMaskGenerator")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_DestroyQMaskGenerator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMaskGenerator) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMaskGeneratorTimerEvent
func callbackQMaskGeneratorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMaskGenerator) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMaskGenerator) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMaskGeneratorChildEvent
func callbackQMaskGeneratorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMaskGenerator) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMaskGenerator) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMaskGeneratorCustomEvent
func callbackQMaskGeneratorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMaskGenerator) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

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
	return n
}

func newQWebSocketFromPointer(ptr unsafe.Pointer) *QWebSocket {
	var n = NewQWebSocketFromPointer(ptr)
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
func callbackQWebSocketAboutToClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::aboutToClose")

	if signal := qt.GetSignal(C.GoString(ptrName), "aboutToClose"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) AboutToClose() {
	defer qt.Recovering("QWebSocket::aboutToClose")

	if ptr.Pointer() != nil {
		C.QWebSocket_AboutToClose(ptr.Pointer())
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
func callbackQWebSocketBinaryFrameReceived(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer, isLastFrame C.int) {
	defer qt.Recovering("callback QWebSocket::binaryFrameReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "binaryFrameReceived"); signal != nil {
		signal.(func(*core.QByteArray, bool))(core.NewQByteArrayFromPointer(frame), int(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) BinaryFrameReceived(frame core.QByteArray_ITF, isLastFrame bool) {
	defer qt.Recovering("QWebSocket::binaryFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_BinaryFrameReceived(ptr.Pointer(), core.PointerFromQByteArray(frame), C.int(qt.GoBoolToInt(isLastFrame)))
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
func callbackQWebSocketBinaryMessageReceived(ptr unsafe.Pointer, ptrName *C.char, message unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::binaryMessageReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "binaryMessageReceived"); signal != nil {
		signal.(func(*core.QByteArray))(core.NewQByteArrayFromPointer(message))
	}

}

func (ptr *QWebSocket) BinaryMessageReceived(message core.QByteArray_ITF) {
	defer qt.Recovering("QWebSocket::binaryMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_BinaryMessageReceived(ptr.Pointer(), core.PointerFromQByteArray(message))
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
func callbackQWebSocketBytesWritten(ptr unsafe.Pointer, ptrName *C.char, bytes C.longlong) {
	defer qt.Recovering("callback QWebSocket::bytesWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "bytesWritten"); signal != nil {
		signal.(func(int64))(int64(bytes))
	}

}

func (ptr *QWebSocket) BytesWritten(bytes int64) {
	defer qt.Recovering("QWebSocket::bytesWritten")

	if ptr.Pointer() != nil {
		C.QWebSocket_BytesWritten(ptr.Pointer(), C.longlong(bytes))
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
func callbackQWebSocketConnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) Connected() {
	defer qt.Recovering("QWebSocket::connected")

	if ptr.Pointer() != nil {
		C.QWebSocket_Connected(ptr.Pointer())
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
func callbackQWebSocketDisconnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) Disconnected() {
	defer qt.Recovering("QWebSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QWebSocket_Disconnected(ptr.Pointer())
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
func callbackQWebSocketError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QWebSocket::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketError))(network.QAbstractSocket__SocketError(error))
	}

}

func (ptr *QWebSocket) Error2(error network.QAbstractSocket__SocketError) {
	defer qt.Recovering("QWebSocket::error")

	if ptr.Pointer() != nil {
		C.QWebSocket_Error2(ptr.Pointer(), C.int(error))
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
func callbackQWebSocketReadChannelFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QWebSocket::readChannelFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "readChannelFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QWebSocket) ReadChannelFinished() {
	defer qt.Recovering("QWebSocket::readChannelFinished")

	if ptr.Pointer() != nil {
		C.QWebSocket_ReadChannelFinished(ptr.Pointer())
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
func callbackQWebSocketStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QWebSocket::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(network.QAbstractSocket__SocketState))(network.QAbstractSocket__SocketState(state))
	}

}

func (ptr *QWebSocket) StateChanged(state network.QAbstractSocket__SocketState) {
	defer qt.Recovering("QWebSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QWebSocket_StateChanged(ptr.Pointer(), C.int(state))
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
func callbackQWebSocketTextFrameReceived(ptr unsafe.Pointer, ptrName *C.char, frame *C.char, isLastFrame C.int) {
	defer qt.Recovering("callback QWebSocket::textFrameReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "textFrameReceived"); signal != nil {
		signal.(func(string, bool))(C.GoString(frame), int(isLastFrame) != 0)
	}

}

func (ptr *QWebSocket) TextFrameReceived(frame string, isLastFrame bool) {
	defer qt.Recovering("QWebSocket::textFrameReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_TextFrameReceived(ptr.Pointer(), C.CString(frame), C.int(qt.GoBoolToInt(isLastFrame)))
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
func callbackQWebSocketTextMessageReceived(ptr unsafe.Pointer, ptrName *C.char, message *C.char) {
	defer qt.Recovering("callback QWebSocket::textMessageReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "textMessageReceived"); signal != nil {
		signal.(func(string))(C.GoString(message))
	}

}

func (ptr *QWebSocket) TextMessageReceived(message string) {
	defer qt.Recovering("QWebSocket::textMessageReceived")

	if ptr.Pointer() != nil {
		C.QWebSocket_TextMessageReceived(ptr.Pointer(), C.CString(message))
	}
}

func (ptr *QWebSocket) DestroyQWebSocket() {
	defer qt.Recovering("QWebSocket::~QWebSocket")

	if ptr.Pointer() != nil {
		C.QWebSocket_DestroyQWebSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQWebSocketTimerEvent
func callbackQWebSocketTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQWebSocketChildEvent
func callbackQWebSocketChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQWebSocketCustomEvent
func callbackQWebSocketCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QWebSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QWebSocketCorsAuthenticator struct {
	ptr unsafe.Pointer
}

type QWebSocketCorsAuthenticator_ITF interface {
	QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator
}

func (p *QWebSocketCorsAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWebSocketCorsAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWebSocketCorsAuthenticator(ptr QWebSocketCorsAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketCorsAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQWebSocketCorsAuthenticatorFromPointer(ptr unsafe.Pointer) *QWebSocketCorsAuthenticator {
	var n = new(QWebSocketCorsAuthenticator)
	n.SetPointer(ptr)
	return n
}

func newQWebSocketCorsAuthenticatorFromPointer(ptr unsafe.Pointer) *QWebSocketCorsAuthenticator {
	var n = NewQWebSocketCorsAuthenticatorFromPointer(ptr)
	return n
}

func (ptr *QWebSocketCorsAuthenticator) QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator {
	return ptr
}

func NewQWebSocketCorsAuthenticator3(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	defer qt.Recovering("QWebSocketCorsAuthenticator::QWebSocketCorsAuthenticator")

	return newQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(PointerFromQWebSocketCorsAuthenticator(other)))
}

func NewQWebSocketCorsAuthenticator(origin string) *QWebSocketCorsAuthenticator {
	defer qt.Recovering("QWebSocketCorsAuthenticator::QWebSocketCorsAuthenticator")

	return newQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(C.CString(origin)))
}

func NewQWebSocketCorsAuthenticator2(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {
	defer qt.Recovering("QWebSocketCorsAuthenticator::QWebSocketCorsAuthenticator")

	return newQWebSocketCorsAuthenticatorFromPointer(C.QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(PointerFromQWebSocketCorsAuthenticator(other)))
}

func (ptr *QWebSocketCorsAuthenticator) Allowed() bool {
	defer qt.Recovering("QWebSocketCorsAuthenticator::allowed")

	if ptr.Pointer() != nil {
		return C.QWebSocketCorsAuthenticator_Allowed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocketCorsAuthenticator) Origin() string {
	defer qt.Recovering("QWebSocketCorsAuthenticator::origin")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocketCorsAuthenticator_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocketCorsAuthenticator) SetAllowed(allowed bool) {
	defer qt.Recovering("QWebSocketCorsAuthenticator::setAllowed")

	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_SetAllowed(ptr.Pointer(), C.int(qt.GoBoolToInt(allowed)))
	}
}

func (ptr *QWebSocketCorsAuthenticator) Swap(other QWebSocketCorsAuthenticator_ITF) {
	defer qt.Recovering("QWebSocketCorsAuthenticator::swap")

	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_Swap(ptr.Pointer(), PointerFromQWebSocketCorsAuthenticator(other))
	}
}

func (ptr *QWebSocketCorsAuthenticator) DestroyQWebSocketCorsAuthenticator() {
	defer qt.Recovering("QWebSocketCorsAuthenticator::~QWebSocketCorsAuthenticator")

	if ptr.Pointer() != nil {
		C.QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(ptr.Pointer())
	}
}

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
	return n
}

func newQWebSocketServerFromPointer(ptr unsafe.Pointer) *QWebSocketServer {
	var n = NewQWebSocketServerFromPointer(ptr)
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

	return newQWebSocketServerFromPointer(C.QWebSocketServer_NewQWebSocketServer(C.CString(serverName), C.int(secureMode), core.PointerFromQObject(parent)))
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
