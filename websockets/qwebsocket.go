package websockets

//#include "websockets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"log"
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
		n.SetObjectName("QWebSocket_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWebSocket) QWebSocket_PTR() *QWebSocket {
	return ptr
}

func (ptr *QWebSocket) Abort() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::abort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QWebSocket) ConnectAboutToClose(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::aboutToClose")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectAboutToClose(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToClose", f)
	}
}

func (ptr *QWebSocket) DisconnectAboutToClose() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::aboutToClose")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectAboutToClose(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToClose")
	}
}

//export callbackQWebSocketAboutToClose
func callbackQWebSocketAboutToClose(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::aboutToClose")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "aboutToClose").(func())()
}

func (ptr *QWebSocket) ConnectBinaryFrameReceived(f func(frame *core.QByteArray, isLastFrame bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::binaryFrameReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryFrameReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "binaryFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryFrameReceived() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::binaryFrameReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "binaryFrameReceived")
	}
}

//export callbackQWebSocketBinaryFrameReceived
func callbackQWebSocketBinaryFrameReceived(ptrName *C.char, frame unsafe.Pointer, isLastFrame C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::binaryFrameReceived")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "binaryFrameReceived").(func(*core.QByteArray, bool))(core.NewQByteArrayFromPointer(frame), int(isLastFrame) != 0)
}

func (ptr *QWebSocket) ConnectBinaryMessageReceived(f func(message *core.QByteArray)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::binaryMessageReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectBinaryMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "binaryMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectBinaryMessageReceived() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::binaryMessageReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectBinaryMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "binaryMessageReceived")
	}
}

//export callbackQWebSocketBinaryMessageReceived
func callbackQWebSocketBinaryMessageReceived(ptrName *C.char, message unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::binaryMessageReceived")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "binaryMessageReceived").(func(*core.QByteArray))(core.NewQByteArrayFromPointer(message))
}

func (ptr *QWebSocket) CloseReason() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::closeReason")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_CloseReason(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) ConnectConnected(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::connected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QWebSocket) DisconnectConnected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::connected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQWebSocketConnected
func callbackQWebSocketConnected(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::connected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "connected").(func())()
}

func (ptr *QWebSocket) ConnectDisconnected(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::disconnected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QWebSocket) DisconnectDisconnected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::disconnected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQWebSocketDisconnected
func callbackQWebSocketDisconnected(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::disconnected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QWebSocket) Error() network.QAbstractSocket__SocketError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::error")
		}
	}()

	if ptr.Pointer() != nil {
		return network.QAbstractSocket__SocketError(C.QWebSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Flush() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::flush")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWebSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocket) IgnoreSslErrors() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::ignoreSslErrors")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QWebSocket) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QWebSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWebSocket) MaskGenerator() *QMaskGenerator {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::maskGenerator")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQMaskGeneratorFromPointer(C.QWebSocket_MaskGenerator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWebSocket) Open(url core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::open")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_Open(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QWebSocket) Origin() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::origin")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_Origin(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) PauseMode() network.QAbstractSocket__PauseMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::pauseMode")
		}
	}()

	if ptr.Pointer() != nil {
		return network.QAbstractSocket__PauseMode(C.QWebSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWebSocket) PeerName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::peerName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Ping(payload core.QByteArray_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::ping")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_Ping(ptr.Pointer(), core.PointerFromQByteArray(payload))
	}
}

func (ptr *QWebSocket) ConnectReadChannelFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::readChannelFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectReadChannelFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readChannelFinished", f)
	}
}

func (ptr *QWebSocket) DisconnectReadChannelFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::readChannelFinished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectReadChannelFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readChannelFinished")
	}
}

//export callbackQWebSocketReadChannelFinished
func callbackQWebSocketReadChannelFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::readChannelFinished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "readChannelFinished").(func())()
}

func (ptr *QWebSocket) ResourceName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::resourceName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QWebSocket_ResourceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWebSocket) Resume() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::resume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QWebSocket) SetMaskGenerator(maskGenerator QMaskGenerator_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::setMaskGenerator")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_SetMaskGenerator(ptr.Pointer(), PointerFromQMaskGenerator(maskGenerator))
	}
}

func (ptr *QWebSocket) SetPauseMode(pauseMode network.QAbstractSocket__PauseMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::setPauseMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_SetPauseMode(ptr.Pointer(), C.int(pauseMode))
	}
}

func (ptr *QWebSocket) SetProxy(networkProxy network.QNetworkProxy_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::setProxy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_SetProxy(ptr.Pointer(), network.PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QWebSocket) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::setSslConfiguration")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_SetSslConfiguration(ptr.Pointer(), network.PointerFromQSslConfiguration(sslConfiguration))
	}
}

func (ptr *QWebSocket) ConnectStateChanged(f func(state network.QAbstractSocket__SocketState)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QWebSocket) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQWebSocketStateChanged
func callbackQWebSocketStateChanged(ptrName *C.char, state C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(network.QAbstractSocket__SocketState))(network.QAbstractSocket__SocketState(state))
}

func (ptr *QWebSocket) ConnectTextFrameReceived(f func(frame string, isLastFrame bool)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::textFrameReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextFrameReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textFrameReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextFrameReceived() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::textFrameReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextFrameReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textFrameReceived")
	}
}

//export callbackQWebSocketTextFrameReceived
func callbackQWebSocketTextFrameReceived(ptrName *C.char, frame *C.char, isLastFrame C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::textFrameReceived")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "textFrameReceived").(func(string, bool))(C.GoString(frame), int(isLastFrame) != 0)
}

func (ptr *QWebSocket) ConnectTextMessageReceived(f func(message string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::textMessageReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_ConnectTextMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textMessageReceived", f)
	}
}

func (ptr *QWebSocket) DisconnectTextMessageReceived() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::textMessageReceived")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DisconnectTextMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textMessageReceived")
	}
}

//export callbackQWebSocketTextMessageReceived
func callbackQWebSocketTextMessageReceived(ptrName *C.char, message *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::textMessageReceived")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "textMessageReceived").(func(string))(C.GoString(message))
}

func (ptr *QWebSocket) DestroyQWebSocket() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QWebSocket::~QWebSocket")
		}
	}()

	if ptr.Pointer() != nil {
		C.QWebSocket_DestroyQWebSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
