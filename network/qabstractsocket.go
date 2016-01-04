package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractSocket struct {
	core.QIODevice
}

type QAbstractSocket_ITF interface {
	core.QIODevice_ITF
	QAbstractSocket_PTR() *QAbstractSocket
}

func PointerFromQAbstractSocket(ptr QAbstractSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSocketFromPointer(ptr unsafe.Pointer) *QAbstractSocket {
	var n = new(QAbstractSocket)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractSocket_") {
		n.SetObjectName("QAbstractSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractSocket) QAbstractSocket_PTR() *QAbstractSocket {
	return ptr
}

//QAbstractSocket::BindFlag
type QAbstractSocket__BindFlag int64

const (
	QAbstractSocket__DefaultForPlatform = QAbstractSocket__BindFlag(0x0)
	QAbstractSocket__ShareAddress       = QAbstractSocket__BindFlag(0x1)
	QAbstractSocket__DontShareAddress   = QAbstractSocket__BindFlag(0x2)
	QAbstractSocket__ReuseAddressHint   = QAbstractSocket__BindFlag(0x4)
)

//QAbstractSocket::NetworkLayerProtocol
type QAbstractSocket__NetworkLayerProtocol int64

const (
	QAbstractSocket__IPv4Protocol                = QAbstractSocket__NetworkLayerProtocol(0)
	QAbstractSocket__IPv6Protocol                = QAbstractSocket__NetworkLayerProtocol(1)
	QAbstractSocket__AnyIPProtocol               = QAbstractSocket__NetworkLayerProtocol(2)
	QAbstractSocket__UnknownNetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(-1)
)

//QAbstractSocket::PauseMode
type QAbstractSocket__PauseMode int64

const (
	QAbstractSocket__PauseNever       = QAbstractSocket__PauseMode(0x0)
	QAbstractSocket__PauseOnSslErrors = QAbstractSocket__PauseMode(0x1)
)

//QAbstractSocket::SocketError
type QAbstractSocket__SocketError int64

const (
	QAbstractSocket__ConnectionRefusedError           = QAbstractSocket__SocketError(0)
	QAbstractSocket__RemoteHostClosedError            = QAbstractSocket__SocketError(1)
	QAbstractSocket__HostNotFoundError                = QAbstractSocket__SocketError(2)
	QAbstractSocket__SocketAccessError                = QAbstractSocket__SocketError(3)
	QAbstractSocket__SocketResourceError              = QAbstractSocket__SocketError(4)
	QAbstractSocket__SocketTimeoutError               = QAbstractSocket__SocketError(5)
	QAbstractSocket__DatagramTooLargeError            = QAbstractSocket__SocketError(6)
	QAbstractSocket__NetworkError                     = QAbstractSocket__SocketError(7)
	QAbstractSocket__AddressInUseError                = QAbstractSocket__SocketError(8)
	QAbstractSocket__SocketAddressNotAvailableError   = QAbstractSocket__SocketError(9)
	QAbstractSocket__UnsupportedSocketOperationError  = QAbstractSocket__SocketError(10)
	QAbstractSocket__UnfinishedSocketOperationError   = QAbstractSocket__SocketError(11)
	QAbstractSocket__ProxyAuthenticationRequiredError = QAbstractSocket__SocketError(12)
	QAbstractSocket__SslHandshakeFailedError          = QAbstractSocket__SocketError(13)
	QAbstractSocket__ProxyConnectionRefusedError      = QAbstractSocket__SocketError(14)
	QAbstractSocket__ProxyConnectionClosedError       = QAbstractSocket__SocketError(15)
	QAbstractSocket__ProxyConnectionTimeoutError      = QAbstractSocket__SocketError(16)
	QAbstractSocket__ProxyNotFoundError               = QAbstractSocket__SocketError(17)
	QAbstractSocket__ProxyProtocolError               = QAbstractSocket__SocketError(18)
	QAbstractSocket__OperationError                   = QAbstractSocket__SocketError(19)
	QAbstractSocket__SslInternalError                 = QAbstractSocket__SocketError(20)
	QAbstractSocket__SslInvalidUserDataError          = QAbstractSocket__SocketError(21)
	QAbstractSocket__TemporaryError                   = QAbstractSocket__SocketError(22)
	QAbstractSocket__UnknownSocketError               = QAbstractSocket__SocketError(-1)
)

//QAbstractSocket::SocketOption
type QAbstractSocket__SocketOption int64

const (
	QAbstractSocket__LowDelayOption                = QAbstractSocket__SocketOption(0)
	QAbstractSocket__KeepAliveOption               = QAbstractSocket__SocketOption(1)
	QAbstractSocket__MulticastTtlOption            = QAbstractSocket__SocketOption(2)
	QAbstractSocket__MulticastLoopbackOption       = QAbstractSocket__SocketOption(3)
	QAbstractSocket__TypeOfServiceOption           = QAbstractSocket__SocketOption(4)
	QAbstractSocket__SendBufferSizeSocketOption    = QAbstractSocket__SocketOption(5)
	QAbstractSocket__ReceiveBufferSizeSocketOption = QAbstractSocket__SocketOption(6)
)

//QAbstractSocket::SocketState
type QAbstractSocket__SocketState int64

const (
	QAbstractSocket__UnconnectedState = QAbstractSocket__SocketState(0)
	QAbstractSocket__HostLookupState  = QAbstractSocket__SocketState(1)
	QAbstractSocket__ConnectingState  = QAbstractSocket__SocketState(2)
	QAbstractSocket__ConnectedState   = QAbstractSocket__SocketState(3)
	QAbstractSocket__BoundState       = QAbstractSocket__SocketState(4)
	QAbstractSocket__ListeningState   = QAbstractSocket__SocketState(5)
	QAbstractSocket__ClosingState     = QAbstractSocket__SocketState(6)
)

//QAbstractSocket::SocketType
type QAbstractSocket__SocketType int64

const (
	QAbstractSocket__TcpSocket         = QAbstractSocket__SocketType(0)
	QAbstractSocket__UdpSocket         = QAbstractSocket__SocketType(1)
	QAbstractSocket__UnknownSocketType = QAbstractSocket__SocketType(-1)
)

func NewQAbstractSocket(socketType QAbstractSocket__SocketType, parent core.QObject_ITF) *QAbstractSocket {
	defer qt.Recovering("QAbstractSocket::QAbstractSocket")

	return NewQAbstractSocketFromPointer(C.QAbstractSocket_NewQAbstractSocket(C.int(socketType), core.PointerFromQObject(parent)))
}

func (ptr *QAbstractSocket) Abort() {
	defer qt.Recovering("QAbstractSocket::abort")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) AtEnd() bool {
	defer qt.Recovering("QAbstractSocket::atEnd")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) BytesAvailable() int64 {
	defer qt.Recovering("QAbstractSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) BytesToWrite() int64 {
	defer qt.Recovering("QAbstractSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) CanReadLine() bool {
	defer qt.Recovering("QAbstractSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QAbstractSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QAbstractSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QAbstractSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQAbstractSocketClose
func callbackQAbstractSocketClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QAbstractSocket) Close() {
	defer qt.Recovering("QAbstractSocket::close")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Close(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) CloseDefault() {
	defer qt.Recovering("QAbstractSocket::close")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QAbstractSocket::connected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QAbstractSocket::connected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQAbstractSocketConnected
func callbackQAbstractSocketConnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSocket::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSocket) Connected() {
	defer qt.Recovering("QAbstractSocket::connected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Connected(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QAbstractSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectFromHost", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QAbstractSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectFromHost")
	}
}

//export callbackQAbstractSocketDisconnectFromHost
func callbackQAbstractSocketDisconnectFromHost(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSocket::disconnectFromHost")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QAbstractSocket) DisconnectFromHost() {
	defer qt.Recovering("QAbstractSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QAbstractSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QAbstractSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QAbstractSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQAbstractSocketDisconnected
func callbackQAbstractSocketDisconnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSocket::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSocket) Disconnected() {
	defer qt.Recovering("QAbstractSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Disconnected(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) ConnectError2(f func(socketError QAbstractSocket__SocketError)) {
	defer qt.Recovering("connect QAbstractSocket::error")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QAbstractSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QAbstractSocket::error")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQAbstractSocketError2
func callbackQAbstractSocketError2(ptr unsafe.Pointer, ptrName *C.char, socketError C.int) {
	defer qt.Recovering("callback QAbstractSocket::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QAbstractSocket__SocketError))(QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QAbstractSocket) Error2(socketError QAbstractSocket__SocketError) {
	defer qt.Recovering("QAbstractSocket::error")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Error2(ptr.Pointer(), C.int(socketError))
	}
}

func (ptr *QAbstractSocket) Error() QAbstractSocket__SocketError {
	defer qt.Recovering("QAbstractSocket::error")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QAbstractSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) Flush() bool {
	defer qt.Recovering("QAbstractSocket::flush")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) ConnectHostFound(f func()) {
	defer qt.Recovering("connect QAbstractSocket::hostFound")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectHostFound(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hostFound", f)
	}
}

func (ptr *QAbstractSocket) DisconnectHostFound() {
	defer qt.Recovering("disconnect QAbstractSocket::hostFound")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectHostFound(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hostFound")
	}
}

//export callbackQAbstractSocketHostFound
func callbackQAbstractSocketHostFound(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSocket::hostFound")

	if signal := qt.GetSignal(C.GoString(ptrName), "hostFound"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSocket) HostFound() {
	defer qt.Recovering("QAbstractSocket::hostFound")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_HostFound(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) IsSequential() bool {
	defer qt.Recovering("QAbstractSocket::isSequential")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) IsValid() bool {
	defer qt.Recovering("QAbstractSocket::isValid")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) PauseMode() QAbstractSocket__PauseMode {
	defer qt.Recovering("QAbstractSocket::pauseMode")

	if ptr.Pointer() != nil {
		return QAbstractSocket__PauseMode(C.QAbstractSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) PeerName() string {
	defer qt.Recovering("QAbstractSocket::peerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSocket) ReadBufferSize() int64 {
	defer qt.Recovering("QAbstractSocket::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) ReadData(data string, maxSize int64) int64 {
	defer qt.Recovering("QAbstractSocket::readData")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_ReadData(ptr.Pointer(), C.CString(data), C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QAbstractSocket) ReadLineData(data string, maxlen int64) int64 {
	defer qt.Recovering("QAbstractSocket::readLineData")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_ReadLineData(ptr.Pointer(), C.CString(data), C.longlong(maxlen)))
	}
	return 0
}

func (ptr *QAbstractSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QAbstractSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resume", f)
	}
}

func (ptr *QAbstractSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QAbstractSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resume")
	}
}

//export callbackQAbstractSocketResume
func callbackQAbstractSocketResume(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSocket::resume")

	if signal := qt.GetSignal(C.GoString(ptrName), "resume"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QAbstractSocket) Resume() {
	defer qt.Recovering("QAbstractSocket::resume")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) ResumeDefault() {
	defer qt.Recovering("QAbstractSocket::resume")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) SetPauseMode(pauseMode QAbstractSocket__PauseMode) {
	defer qt.Recovering("QAbstractSocket::setPauseMode")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetPauseMode(ptr.Pointer(), C.int(pauseMode))
	}
}

func (ptr *QAbstractSocket) SetProxy(networkProxy QNetworkProxy_ITF) {
	defer qt.Recovering("QAbstractSocket::setProxy")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QAbstractSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QAbstractSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setReadBufferSize", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QAbstractSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setReadBufferSize")
	}
}

//export callbackQAbstractSocketSetReadBufferSize
func callbackQAbstractSocketSetReadBufferSize(ptr unsafe.Pointer, ptrName *C.char, size C.longlong) {
	defer qt.Recovering("callback QAbstractSocket::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQAbstractSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QAbstractSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QAbstractSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QAbstractSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QAbstractSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QAbstractSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QAbstractSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSocketOption", f)
	}
}

func (ptr *QAbstractSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QAbstractSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSocketOption")
	}
}

//export callbackQAbstractSocketSetSocketOption
func callbackQAbstractSocketSetSocketOption(ptr unsafe.Pointer, ptrName *C.char, option C.int, value unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::setSocketOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQAbstractSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QAbstractSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QAbstractSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketOption(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAbstractSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QAbstractSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketOptionDefault(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAbstractSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QAbstractSocket::socketOption")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractSocket_SocketOption(ptr.Pointer(), C.int(option)))
	}
	return nil
}

func (ptr *QAbstractSocket) SocketType() QAbstractSocket__SocketType {
	defer qt.Recovering("QAbstractSocket::socketType")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketType(C.QAbstractSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) State() QAbstractSocket__SocketState {
	defer qt.Recovering("QAbstractSocket::state")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketState(C.QAbstractSocket_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) ConnectStateChanged(f func(socketState QAbstractSocket__SocketState)) {
	defer qt.Recovering("connect QAbstractSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAbstractSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QAbstractSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAbstractSocketStateChanged
func callbackQAbstractSocketStateChanged(ptr unsafe.Pointer, ptrName *C.char, socketState C.int) {
	defer qt.Recovering("callback QAbstractSocket::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QAbstractSocket__SocketState))(QAbstractSocket__SocketState(socketState))
	}

}

func (ptr *QAbstractSocket) StateChanged(socketState QAbstractSocket__SocketState) {
	defer qt.Recovering("QAbstractSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_StateChanged(ptr.Pointer(), C.int(socketState))
	}
}

func (ptr *QAbstractSocket) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForConnected(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForConnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForDisconnected(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForDisconnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QAbstractSocket::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WriteData(data string, size int64) int64 {
	defer qt.Recovering("QAbstractSocket::writeData")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractSocket_WriteData(ptr.Pointer(), C.CString(data), C.longlong(size)))
	}
	return 0
}

func (ptr *QAbstractSocket) DestroyQAbstractSocket() {
	defer qt.Recovering("QAbstractSocket::~QAbstractSocket")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_DestroyQAbstractSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractSocketTimerEvent
func callbackQAbstractSocketTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractSocketChildEvent
func callbackQAbstractSocketChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractSocketCustomEvent
func callbackQAbstractSocketCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
