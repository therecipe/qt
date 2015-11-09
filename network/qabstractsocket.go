package network

//#include "qabstractsocket.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAbstractSocket_" + qt.RandomIdentifier())
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
	return NewQAbstractSocketFromPointer(C.QAbstractSocket_NewQAbstractSocket(C.int(socketType), core.PointerFromQObject(parent)))
}

func (ptr *QAbstractSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) Close() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Close(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQAbstractSocketConnected
func callbackQAbstractSocketConnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "connected").(func())()
}

func (ptr *QAbstractSocket) DisconnectFromHost() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QAbstractSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQAbstractSocketDisconnected
func callbackQAbstractSocketDisconnected(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "disconnected").(func())()
}

func (ptr *QAbstractSocket) Error() QAbstractSocket__SocketError {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QAbstractSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) ConnectHostFound(f func()) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectHostFound(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hostFound", f)
	}
}

func (ptr *QAbstractSocket) DisconnectHostFound() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectHostFound(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hostFound")
	}
}

//export callbackQAbstractSocketHostFound
func callbackQAbstractSocketHostFound(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "hostFound").(func())()
}

func (ptr *QAbstractSocket) IsSequential() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSocket) PauseMode() QAbstractSocket__PauseMode {
	if ptr.Pointer() != nil {
		return QAbstractSocket__PauseMode(C.QAbstractSocket_PauseMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QAbstractSocket) SetPauseMode(pauseMode QAbstractSocket__PauseMode) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetPauseMode(ptr.Pointer(), C.int(pauseMode))
	}
}

func (ptr *QAbstractSocket) SetProxy(networkProxy QNetworkProxy_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QAbstractSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_SetSocketOption(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QAbstractSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAbstractSocket_SocketOption(ptr.Pointer(), C.int(option)))
	}
	return nil
}

func (ptr *QAbstractSocket) SocketType() QAbstractSocket__SocketType {
	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketType(C.QAbstractSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSocket) ConnectStateChanged(f func(socketState QAbstractSocket__SocketState)) {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QAbstractSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQAbstractSocketStateChanged
func callbackQAbstractSocketStateChanged(ptrName *C.char, socketState C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QAbstractSocket__SocketState))(QAbstractSocket__SocketState(socketState))
}

func (ptr *QAbstractSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForConnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForDisconnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QAbstractSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QAbstractSocket) DestroyQAbstractSocket() {
	if ptr.Pointer() != nil {
		C.QAbstractSocket_DestroyQAbstractSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
