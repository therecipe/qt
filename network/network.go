package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QAbstractNetworkCache struct {
	core.QObject
}

type QAbstractNetworkCache_ITF interface {
	core.QObject_ITF
	QAbstractNetworkCache_PTR() *QAbstractNetworkCache
}

func PointerFromQAbstractNetworkCache(ptr QAbstractNetworkCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNetworkCache_PTR().Pointer()
	}
	return nil
}

func NewQAbstractNetworkCacheFromPointer(ptr unsafe.Pointer) *QAbstractNetworkCache {
	var n = new(QAbstractNetworkCache)
	n.SetPointer(ptr)
	return n
}

func newQAbstractNetworkCacheFromPointer(ptr unsafe.Pointer) *QAbstractNetworkCache {
	var n = NewQAbstractNetworkCacheFromPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractNetworkCache_") {
		n.SetObjectName("QAbstractNetworkCache_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractNetworkCache) QAbstractNetworkCache_PTR() *QAbstractNetworkCache {
	return ptr
}

func (ptr *QAbstractNetworkCache) CacheSize() int64 {
	defer qt.Recovering("QAbstractNetworkCache::cacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractNetworkCache_CacheSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractNetworkCache) Clear() {
	defer qt.Recovering("QAbstractNetworkCache::clear")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Clear(ptr.Pointer())
	}
}

func (ptr *QAbstractNetworkCache) Data(url core.QUrl_ITF) *core.QIODevice {
	defer qt.Recovering("QAbstractNetworkCache::data")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	defer qt.Recovering("QAbstractNetworkCache::prepare")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::updateMetaData")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCache() {
	defer qt.Recovering("QAbstractNetworkCache::~QAbstractNetworkCache")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCache(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractNetworkCache) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractNetworkCacheTimerEvent
func callbackQAbstractNetworkCacheTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractNetworkCacheChildEvent
func callbackQAbstractNetworkCacheChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractNetworkCacheCustomEvent
func callbackQAbstractNetworkCacheCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

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
	return n
}

func newQAbstractSocketFromPointer(ptr unsafe.Pointer) *QAbstractSocket {
	var n = NewQAbstractSocketFromPointer(ptr)
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

	return newQAbstractSocketFromPointer(C.QAbstractSocket_NewQAbstractSocket(C.int(socketType), core.PointerFromQObject(parent)))
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

type QAuthenticator struct {
	ptr unsafe.Pointer
}

type QAuthenticator_ITF interface {
	QAuthenticator_PTR() *QAuthenticator
}

func (p *QAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAuthenticator(ptr QAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQAuthenticatorFromPointer(ptr unsafe.Pointer) *QAuthenticator {
	var n = new(QAuthenticator)
	n.SetPointer(ptr)
	return n
}

func newQAuthenticatorFromPointer(ptr unsafe.Pointer) *QAuthenticator {
	var n = NewQAuthenticatorFromPointer(ptr)
	return n
}

func (ptr *QAuthenticator) QAuthenticator_PTR() *QAuthenticator {
	return ptr
}

func NewQAuthenticator() *QAuthenticator {
	defer qt.Recovering("QAuthenticator::QAuthenticator")

	return newQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator())
}

func NewQAuthenticator2(other QAuthenticator_ITF) *QAuthenticator {
	defer qt.Recovering("QAuthenticator::QAuthenticator")

	return newQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator2(PointerFromQAuthenticator(other)))
}

func (ptr *QAuthenticator) IsNull() bool {
	defer qt.Recovering("QAuthenticator::isNull")

	if ptr.Pointer() != nil {
		return C.QAuthenticator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAuthenticator) Option(opt string) *core.QVariant {
	defer qt.Recovering("QAuthenticator::option")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAuthenticator_Option(ptr.Pointer(), C.CString(opt)))
	}
	return nil
}

func (ptr *QAuthenticator) Password() string {
	defer qt.Recovering("QAuthenticator::password")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) Realm() string {
	defer qt.Recovering("QAuthenticator::realm")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Realm(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) SetOption(opt string, value core.QVariant_ITF) {
	defer qt.Recovering("QAuthenticator::setOption")

	if ptr.Pointer() != nil {
		C.QAuthenticator_SetOption(ptr.Pointer(), C.CString(opt), core.PointerFromQVariant(value))
	}
}

func (ptr *QAuthenticator) SetPassword(password string) {
	defer qt.Recovering("QAuthenticator::setPassword")

	if ptr.Pointer() != nil {
		C.QAuthenticator_SetPassword(ptr.Pointer(), C.CString(password))
	}
}

func (ptr *QAuthenticator) SetUser(user string) {
	defer qt.Recovering("QAuthenticator::setUser")

	if ptr.Pointer() != nil {
		C.QAuthenticator_SetUser(ptr.Pointer(), C.CString(user))
	}
}

func (ptr *QAuthenticator) User() string {
	defer qt.Recovering("QAuthenticator::user")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_User(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) DestroyQAuthenticator() {
	defer qt.Recovering("QAuthenticator::~QAuthenticator")

	if ptr.Pointer() != nil {
		C.QAuthenticator_DestroyQAuthenticator(ptr.Pointer())
	}
}

type QDnsDomainNameRecord struct {
	ptr unsafe.Pointer
}

type QDnsDomainNameRecord_ITF interface {
	QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord
}

func (p *QDnsDomainNameRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsDomainNameRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsDomainNameRecord(ptr QDnsDomainNameRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsDomainNameRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsDomainNameRecordFromPointer(ptr unsafe.Pointer) *QDnsDomainNameRecord {
	var n = new(QDnsDomainNameRecord)
	n.SetPointer(ptr)
	return n
}

func newQDnsDomainNameRecordFromPointer(ptr unsafe.Pointer) *QDnsDomainNameRecord {
	var n = NewQDnsDomainNameRecordFromPointer(ptr)
	return n
}

func (ptr *QDnsDomainNameRecord) QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord {
	return ptr
}

func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {
	defer qt.Recovering("QDnsDomainNameRecord::QDnsDomainNameRecord")

	return newQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord())
}

func NewQDnsDomainNameRecord2(other QDnsDomainNameRecord_ITF) *QDnsDomainNameRecord {
	defer qt.Recovering("QDnsDomainNameRecord::QDnsDomainNameRecord")

	return newQDnsDomainNameRecordFromPointer(C.QDnsDomainNameRecord_NewQDnsDomainNameRecord2(PointerFromQDnsDomainNameRecord(other)))
}

func (ptr *QDnsDomainNameRecord) Name() string {
	defer qt.Recovering("QDnsDomainNameRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsDomainNameRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) Swap(other QDnsDomainNameRecord_ITF) {
	defer qt.Recovering("QDnsDomainNameRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_Swap(ptr.Pointer(), PointerFromQDnsDomainNameRecord(other))
	}
}

func (ptr *QDnsDomainNameRecord) Value() string {
	defer qt.Recovering("QDnsDomainNameRecord::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsDomainNameRecord_Value(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsDomainNameRecord) DestroyQDnsDomainNameRecord() {
	defer qt.Recovering("QDnsDomainNameRecord::~QDnsDomainNameRecord")

	if ptr.Pointer() != nil {
		C.QDnsDomainNameRecord_DestroyQDnsDomainNameRecord(ptr.Pointer())
	}
}

type QDnsHostAddressRecord struct {
	ptr unsafe.Pointer
}

type QDnsHostAddressRecord_ITF interface {
	QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord
}

func (p *QDnsHostAddressRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsHostAddressRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsHostAddressRecord(ptr QDnsHostAddressRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsHostAddressRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsHostAddressRecordFromPointer(ptr unsafe.Pointer) *QDnsHostAddressRecord {
	var n = new(QDnsHostAddressRecord)
	n.SetPointer(ptr)
	return n
}

func newQDnsHostAddressRecordFromPointer(ptr unsafe.Pointer) *QDnsHostAddressRecord {
	var n = NewQDnsHostAddressRecordFromPointer(ptr)
	return n
}

func (ptr *QDnsHostAddressRecord) QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord {
	return ptr
}

func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {
	defer qt.Recovering("QDnsHostAddressRecord::QDnsHostAddressRecord")

	return newQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord())
}

func NewQDnsHostAddressRecord2(other QDnsHostAddressRecord_ITF) *QDnsHostAddressRecord {
	defer qt.Recovering("QDnsHostAddressRecord::QDnsHostAddressRecord")

	return newQDnsHostAddressRecordFromPointer(C.QDnsHostAddressRecord_NewQDnsHostAddressRecord2(PointerFromQDnsHostAddressRecord(other)))
}

func (ptr *QDnsHostAddressRecord) Name() string {
	defer qt.Recovering("QDnsHostAddressRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsHostAddressRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsHostAddressRecord) Swap(other QDnsHostAddressRecord_ITF) {
	defer qt.Recovering("QDnsHostAddressRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_Swap(ptr.Pointer(), PointerFromQDnsHostAddressRecord(other))
	}
}

func (ptr *QDnsHostAddressRecord) DestroyQDnsHostAddressRecord() {
	defer qt.Recovering("QDnsHostAddressRecord::~QDnsHostAddressRecord")

	if ptr.Pointer() != nil {
		C.QDnsHostAddressRecord_DestroyQDnsHostAddressRecord(ptr.Pointer())
	}
}

type QDnsLookup struct {
	core.QObject
}

type QDnsLookup_ITF interface {
	core.QObject_ITF
	QDnsLookup_PTR() *QDnsLookup
}

func PointerFromQDnsLookup(ptr QDnsLookup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsLookup_PTR().Pointer()
	}
	return nil
}

func NewQDnsLookupFromPointer(ptr unsafe.Pointer) *QDnsLookup {
	var n = new(QDnsLookup)
	n.SetPointer(ptr)
	return n
}

func newQDnsLookupFromPointer(ptr unsafe.Pointer) *QDnsLookup {
	var n = NewQDnsLookupFromPointer(ptr)
	for len(n.ObjectName()) < len("QDnsLookup_") {
		n.SetObjectName("QDnsLookup_" + qt.Identifier())
	}
	return n
}

func (ptr *QDnsLookup) QDnsLookup_PTR() *QDnsLookup {
	return ptr
}

//QDnsLookup::Error
type QDnsLookup__Error int64

const (
	QDnsLookup__NoError                 = QDnsLookup__Error(0)
	QDnsLookup__ResolverError           = QDnsLookup__Error(1)
	QDnsLookup__OperationCancelledError = QDnsLookup__Error(2)
	QDnsLookup__InvalidRequestError     = QDnsLookup__Error(3)
	QDnsLookup__InvalidReplyError       = QDnsLookup__Error(4)
	QDnsLookup__ServerFailureError      = QDnsLookup__Error(5)
	QDnsLookup__ServerRefusedError      = QDnsLookup__Error(6)
	QDnsLookup__NotFoundError           = QDnsLookup__Error(7)
)

//QDnsLookup::Type
type QDnsLookup__Type int64

const (
	QDnsLookup__A     = QDnsLookup__Type(1)
	QDnsLookup__AAAA  = QDnsLookup__Type(28)
	QDnsLookup__ANY   = QDnsLookup__Type(255)
	QDnsLookup__CNAME = QDnsLookup__Type(5)
	QDnsLookup__MX    = QDnsLookup__Type(15)
	QDnsLookup__NS    = QDnsLookup__Type(2)
	QDnsLookup__PTR   = QDnsLookup__Type(12)
	QDnsLookup__SRV   = QDnsLookup__Type(33)
	QDnsLookup__TXT   = QDnsLookup__Type(16)
)

func NewQDnsLookup3(ty QDnsLookup__Type, name string, nameserver QHostAddress_ITF, parent core.QObject_ITF) *QDnsLookup {
	defer qt.Recovering("QDnsLookup::QDnsLookup")

	return newQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup3(C.int(ty), C.CString(name), PointerFromQHostAddress(nameserver), core.PointerFromQObject(parent)))
}

func (ptr *QDnsLookup) Error() QDnsLookup__Error {
	defer qt.Recovering("QDnsLookup::error")

	if ptr.Pointer() != nil {
		return QDnsLookup__Error(C.QDnsLookup_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDnsLookup) ErrorString() string {
	defer qt.Recovering("QDnsLookup::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) Name() string {
	defer qt.Recovering("QDnsLookup::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsLookup_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsLookup) SetName(name string) {
	defer qt.Recovering("QDnsLookup::setName")

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDnsLookup) SetNameserver(nameserver QHostAddress_ITF) {
	defer qt.Recovering("QDnsLookup::setNameserver")

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetNameserver(ptr.Pointer(), PointerFromQHostAddress(nameserver))
	}
}

func (ptr *QDnsLookup) SetType(v QDnsLookup__Type) {
	defer qt.Recovering("QDnsLookup::setType")

	if ptr.Pointer() != nil {
		C.QDnsLookup_SetType(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QDnsLookup) Type() QDnsLookup__Type {
	defer qt.Recovering("QDnsLookup::type")

	if ptr.Pointer() != nil {
		return QDnsLookup__Type(C.QDnsLookup_Type(ptr.Pointer()))
	}
	return 0
}

func NewQDnsLookup(parent core.QObject_ITF) *QDnsLookup {
	defer qt.Recovering("QDnsLookup::QDnsLookup")

	return newQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup(core.PointerFromQObject(parent)))
}

func NewQDnsLookup2(ty QDnsLookup__Type, name string, parent core.QObject_ITF) *QDnsLookup {
	defer qt.Recovering("QDnsLookup::QDnsLookup")

	return newQDnsLookupFromPointer(C.QDnsLookup_NewQDnsLookup2(C.int(ty), C.CString(name), core.PointerFromQObject(parent)))
}

func (ptr *QDnsLookup) Abort() {
	defer qt.Recovering("QDnsLookup::abort")

	if ptr.Pointer() != nil {
		C.QDnsLookup_Abort(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) ConnectFinished(f func()) {
	defer qt.Recovering("connect QDnsLookup::finished")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QDnsLookup) DisconnectFinished() {
	defer qt.Recovering("disconnect QDnsLookup::finished")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQDnsLookupFinished
func callbackQDnsLookupFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QDnsLookup::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QDnsLookup) Finished() {
	defer qt.Recovering("QDnsLookup::finished")

	if ptr.Pointer() != nil {
		C.QDnsLookup_Finished(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) IsFinished() bool {
	defer qt.Recovering("QDnsLookup::isFinished")

	if ptr.Pointer() != nil {
		return C.QDnsLookup_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDnsLookup) Lookup() {
	defer qt.Recovering("QDnsLookup::lookup")

	if ptr.Pointer() != nil {
		C.QDnsLookup_Lookup(ptr.Pointer())
	}
}

func (ptr *QDnsLookup) ConnectNameChanged(f func(name string)) {
	defer qt.Recovering("connect QDnsLookup::nameChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "nameChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectNameChanged() {
	defer qt.Recovering("disconnect QDnsLookup::nameChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "nameChanged")
	}
}

//export callbackQDnsLookupNameChanged
func callbackQDnsLookupNameChanged(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QDnsLookup::nameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "nameChanged"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QDnsLookup) NameChanged(name string) {
	defer qt.Recovering("QDnsLookup::nameChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_NameChanged(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QDnsLookup) ConnectTypeChanged(f func(ty QDnsLookup__Type)) {
	defer qt.Recovering("connect QDnsLookup::typeChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ConnectTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "typeChanged", f)
	}
}

func (ptr *QDnsLookup) DisconnectTypeChanged() {
	defer qt.Recovering("disconnect QDnsLookup::typeChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DisconnectTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "typeChanged")
	}
}

//export callbackQDnsLookupTypeChanged
func callbackQDnsLookupTypeChanged(ptr unsafe.Pointer, ptrName *C.char, ty C.int) {
	defer qt.Recovering("callback QDnsLookup::typeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "typeChanged"); signal != nil {
		signal.(func(QDnsLookup__Type))(QDnsLookup__Type(ty))
	}

}

func (ptr *QDnsLookup) TypeChanged(ty QDnsLookup__Type) {
	defer qt.Recovering("QDnsLookup::typeChanged")

	if ptr.Pointer() != nil {
		C.QDnsLookup_TypeChanged(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QDnsLookup) DestroyQDnsLookup() {
	defer qt.Recovering("QDnsLookup::~QDnsLookup")

	if ptr.Pointer() != nil {
		C.QDnsLookup_DestroyQDnsLookup(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDnsLookup) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDnsLookupTimerEvent
func callbackQDnsLookupTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDnsLookup) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDnsLookup::timerEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDnsLookup) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDnsLookup::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDnsLookup::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDnsLookupChildEvent
func callbackQDnsLookupChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDnsLookup::childEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDnsLookup) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDnsLookup::childEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDnsLookup) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDnsLookup::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDnsLookup) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDnsLookup::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDnsLookupCustomEvent
func callbackQDnsLookupCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDnsLookup::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDnsLookupFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDnsLookup) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDnsLookup::customEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDnsLookup) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDnsLookup::customEvent")

	if ptr.Pointer() != nil {
		C.QDnsLookup_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QDnsMailExchangeRecord struct {
	ptr unsafe.Pointer
}

type QDnsMailExchangeRecord_ITF interface {
	QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord
}

func (p *QDnsMailExchangeRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsMailExchangeRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsMailExchangeRecord(ptr QDnsMailExchangeRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsMailExchangeRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsMailExchangeRecordFromPointer(ptr unsafe.Pointer) *QDnsMailExchangeRecord {
	var n = new(QDnsMailExchangeRecord)
	n.SetPointer(ptr)
	return n
}

func newQDnsMailExchangeRecordFromPointer(ptr unsafe.Pointer) *QDnsMailExchangeRecord {
	var n = NewQDnsMailExchangeRecordFromPointer(ptr)
	return n
}

func (ptr *QDnsMailExchangeRecord) QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord {
	return ptr
}

func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {
	defer qt.Recovering("QDnsMailExchangeRecord::QDnsMailExchangeRecord")

	return newQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord())
}

func NewQDnsMailExchangeRecord2(other QDnsMailExchangeRecord_ITF) *QDnsMailExchangeRecord {
	defer qt.Recovering("QDnsMailExchangeRecord::QDnsMailExchangeRecord")

	return newQDnsMailExchangeRecordFromPointer(C.QDnsMailExchangeRecord_NewQDnsMailExchangeRecord2(PointerFromQDnsMailExchangeRecord(other)))
}

func (ptr *QDnsMailExchangeRecord) Exchange() string {
	defer qt.Recovering("QDnsMailExchangeRecord::exchange")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsMailExchangeRecord_Exchange(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Name() string {
	defer qt.Recovering("QDnsMailExchangeRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsMailExchangeRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsMailExchangeRecord) Swap(other QDnsMailExchangeRecord_ITF) {
	defer qt.Recovering("QDnsMailExchangeRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_Swap(ptr.Pointer(), PointerFromQDnsMailExchangeRecord(other))
	}
}

func (ptr *QDnsMailExchangeRecord) DestroyQDnsMailExchangeRecord() {
	defer qt.Recovering("QDnsMailExchangeRecord::~QDnsMailExchangeRecord")

	if ptr.Pointer() != nil {
		C.QDnsMailExchangeRecord_DestroyQDnsMailExchangeRecord(ptr.Pointer())
	}
}

type QDnsServiceRecord struct {
	ptr unsafe.Pointer
}

type QDnsServiceRecord_ITF interface {
	QDnsServiceRecord_PTR() *QDnsServiceRecord
}

func (p *QDnsServiceRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsServiceRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsServiceRecord(ptr QDnsServiceRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsServiceRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsServiceRecordFromPointer(ptr unsafe.Pointer) *QDnsServiceRecord {
	var n = new(QDnsServiceRecord)
	n.SetPointer(ptr)
	return n
}

func newQDnsServiceRecordFromPointer(ptr unsafe.Pointer) *QDnsServiceRecord {
	var n = NewQDnsServiceRecordFromPointer(ptr)
	return n
}

func (ptr *QDnsServiceRecord) QDnsServiceRecord_PTR() *QDnsServiceRecord {
	return ptr
}

func NewQDnsServiceRecord() *QDnsServiceRecord {
	defer qt.Recovering("QDnsServiceRecord::QDnsServiceRecord")

	return newQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord())
}

func NewQDnsServiceRecord2(other QDnsServiceRecord_ITF) *QDnsServiceRecord {
	defer qt.Recovering("QDnsServiceRecord::QDnsServiceRecord")

	return newQDnsServiceRecordFromPointer(C.QDnsServiceRecord_NewQDnsServiceRecord2(PointerFromQDnsServiceRecord(other)))
}

func (ptr *QDnsServiceRecord) Name() string {
	defer qt.Recovering("QDnsServiceRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsServiceRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsServiceRecord) Swap(other QDnsServiceRecord_ITF) {
	defer qt.Recovering("QDnsServiceRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_Swap(ptr.Pointer(), PointerFromQDnsServiceRecord(other))
	}
}

func (ptr *QDnsServiceRecord) Target() string {
	defer qt.Recovering("QDnsServiceRecord::target")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsServiceRecord_Target(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsServiceRecord) DestroyQDnsServiceRecord() {
	defer qt.Recovering("QDnsServiceRecord::~QDnsServiceRecord")

	if ptr.Pointer() != nil {
		C.QDnsServiceRecord_DestroyQDnsServiceRecord(ptr.Pointer())
	}
}

type QDnsTextRecord struct {
	ptr unsafe.Pointer
}

type QDnsTextRecord_ITF interface {
	QDnsTextRecord_PTR() *QDnsTextRecord
}

func (p *QDnsTextRecord) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDnsTextRecord) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDnsTextRecord(ptr QDnsTextRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsTextRecord_PTR().Pointer()
	}
	return nil
}

func NewQDnsTextRecordFromPointer(ptr unsafe.Pointer) *QDnsTextRecord {
	var n = new(QDnsTextRecord)
	n.SetPointer(ptr)
	return n
}

func newQDnsTextRecordFromPointer(ptr unsafe.Pointer) *QDnsTextRecord {
	var n = NewQDnsTextRecordFromPointer(ptr)
	return n
}

func (ptr *QDnsTextRecord) QDnsTextRecord_PTR() *QDnsTextRecord {
	return ptr
}

func NewQDnsTextRecord() *QDnsTextRecord {
	defer qt.Recovering("QDnsTextRecord::QDnsTextRecord")

	return newQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord())
}

func NewQDnsTextRecord2(other QDnsTextRecord_ITF) *QDnsTextRecord {
	defer qt.Recovering("QDnsTextRecord::QDnsTextRecord")

	return newQDnsTextRecordFromPointer(C.QDnsTextRecord_NewQDnsTextRecord2(PointerFromQDnsTextRecord(other)))
}

func (ptr *QDnsTextRecord) Name() string {
	defer qt.Recovering("QDnsTextRecord::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDnsTextRecord_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDnsTextRecord) Swap(other QDnsTextRecord_ITF) {
	defer qt.Recovering("QDnsTextRecord::swap")

	if ptr.Pointer() != nil {
		C.QDnsTextRecord_Swap(ptr.Pointer(), PointerFromQDnsTextRecord(other))
	}
}

func (ptr *QDnsTextRecord) DestroyQDnsTextRecord() {
	defer qt.Recovering("QDnsTextRecord::~QDnsTextRecord")

	if ptr.Pointer() != nil {
		C.QDnsTextRecord_DestroyQDnsTextRecord(ptr.Pointer())
	}
}

type QHostAddress struct {
	ptr unsafe.Pointer
}

type QHostAddress_ITF interface {
	QHostAddress_PTR() *QHostAddress
}

func (p *QHostAddress) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHostAddress) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHostAddress(ptr QHostAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostAddress_PTR().Pointer()
	}
	return nil
}

func NewQHostAddressFromPointer(ptr unsafe.Pointer) *QHostAddress {
	var n = new(QHostAddress)
	n.SetPointer(ptr)
	return n
}

func newQHostAddressFromPointer(ptr unsafe.Pointer) *QHostAddress {
	var n = NewQHostAddressFromPointer(ptr)
	return n
}

func (ptr *QHostAddress) QHostAddress_PTR() *QHostAddress {
	return ptr
}

//QHostAddress::SpecialAddress
type QHostAddress__SpecialAddress int64

const (
	QHostAddress__Null          = QHostAddress__SpecialAddress(0)
	QHostAddress__Broadcast     = QHostAddress__SpecialAddress(1)
	QHostAddress__LocalHost     = QHostAddress__SpecialAddress(2)
	QHostAddress__LocalHostIPv6 = QHostAddress__SpecialAddress(3)
	QHostAddress__Any           = QHostAddress__SpecialAddress(4)
	QHostAddress__AnyIPv6       = QHostAddress__SpecialAddress(5)
	QHostAddress__AnyIPv4       = QHostAddress__SpecialAddress(6)
)

func NewQHostAddress() *QHostAddress {
	defer qt.Recovering("QHostAddress::QHostAddress")

	return newQHostAddressFromPointer(C.QHostAddress_NewQHostAddress())
}

func NewQHostAddress9(address QHostAddress__SpecialAddress) *QHostAddress {
	defer qt.Recovering("QHostAddress::QHostAddress")

	return newQHostAddressFromPointer(C.QHostAddress_NewQHostAddress9(C.int(address)))
}

func NewQHostAddress8(address QHostAddress_ITF) *QHostAddress {
	defer qt.Recovering("QHostAddress::QHostAddress")

	return newQHostAddressFromPointer(C.QHostAddress_NewQHostAddress8(PointerFromQHostAddress(address)))
}

func NewQHostAddress7(address string) *QHostAddress {
	defer qt.Recovering("QHostAddress::QHostAddress")

	return newQHostAddressFromPointer(C.QHostAddress_NewQHostAddress7(C.CString(address)))
}

func (ptr *QHostAddress) Clear() {
	defer qt.Recovering("QHostAddress::clear")

	if ptr.Pointer() != nil {
		C.QHostAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QHostAddress) IsInSubnet(subnet QHostAddress_ITF, netmask int) bool {
	defer qt.Recovering("QHostAddress::isInSubnet")

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsInSubnet(ptr.Pointer(), PointerFromQHostAddress(subnet), C.int(netmask)) != 0
	}
	return false
}

func (ptr *QHostAddress) IsLoopback() bool {
	defer qt.Recovering("QHostAddress::isLoopback")

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsLoopback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) IsNull() bool {
	defer qt.Recovering("QHostAddress::isNull")

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) Protocol() QAbstractSocket__NetworkLayerProtocol {
	defer qt.Recovering("QHostAddress::protocol")

	if ptr.Pointer() != nil {
		return QAbstractSocket__NetworkLayerProtocol(C.QHostAddress_Protocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostAddress) ScopeId() string {
	defer qt.Recovering("QHostAddress::scopeId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostAddress_ScopeId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostAddress) SetAddress5(address string) bool {
	defer qt.Recovering("QHostAddress::setAddress")

	if ptr.Pointer() != nil {
		return C.QHostAddress_SetAddress5(ptr.Pointer(), C.CString(address)) != 0
	}
	return false
}

func (ptr *QHostAddress) SetScopeId(id string) {
	defer qt.Recovering("QHostAddress::setScopeId")

	if ptr.Pointer() != nil {
		C.QHostAddress_SetScopeId(ptr.Pointer(), C.CString(id))
	}
}

func (ptr *QHostAddress) ToString() string {
	defer qt.Recovering("QHostAddress::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostAddress) DestroyQHostAddress() {
	defer qt.Recovering("QHostAddress::~QHostAddress")

	if ptr.Pointer() != nil {
		C.QHostAddress_DestroyQHostAddress(ptr.Pointer())
	}
}

type QHostInfo struct {
	ptr unsafe.Pointer
}

type QHostInfo_ITF interface {
	QHostInfo_PTR() *QHostInfo
}

func (p *QHostInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHostInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHostInfo(ptr QHostInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostInfo_PTR().Pointer()
	}
	return nil
}

func NewQHostInfoFromPointer(ptr unsafe.Pointer) *QHostInfo {
	var n = new(QHostInfo)
	n.SetPointer(ptr)
	return n
}

func newQHostInfoFromPointer(ptr unsafe.Pointer) *QHostInfo {
	var n = NewQHostInfoFromPointer(ptr)
	return n
}

func (ptr *QHostInfo) QHostInfo_PTR() *QHostInfo {
	return ptr
}

//QHostInfo::HostInfoError
type QHostInfo__HostInfoError int64

const (
	QHostInfo__NoError      = QHostInfo__HostInfoError(0)
	QHostInfo__HostNotFound = QHostInfo__HostInfoError(1)
	QHostInfo__UnknownError = QHostInfo__HostInfoError(2)
)

func NewQHostInfo2(other QHostInfo_ITF) *QHostInfo {
	defer qt.Recovering("QHostInfo::QHostInfo")

	return newQHostInfoFromPointer(C.QHostInfo_NewQHostInfo2(PointerFromQHostInfo(other)))
}

func NewQHostInfo(id int) *QHostInfo {
	defer qt.Recovering("QHostInfo::QHostInfo")

	return newQHostInfoFromPointer(C.QHostInfo_NewQHostInfo(C.int(id)))
}

func QHostInfo_AbortHostLookup(id int) {
	defer qt.Recovering("QHostInfo::abortHostLookup")

	C.QHostInfo_QHostInfo_AbortHostLookup(C.int(id))
}

func (ptr *QHostInfo) Error() QHostInfo__HostInfoError {
	defer qt.Recovering("QHostInfo::error")

	if ptr.Pointer() != nil {
		return QHostInfo__HostInfoError(C.QHostInfo_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostInfo) ErrorString() string {
	defer qt.Recovering("QHostInfo::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostInfo) HostName() string {
	defer qt.Recovering("QHostInfo::hostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostInfo_HostName(ptr.Pointer()))
	}
	return ""
}

func QHostInfo_LookupHost(name string, receiver core.QObject_ITF, member string) int {
	defer qt.Recovering("QHostInfo::lookupHost")

	return int(C.QHostInfo_QHostInfo_LookupHost(C.CString(name), core.PointerFromQObject(receiver), C.CString(member)))
}

func (ptr *QHostInfo) LookupId() int {
	defer qt.Recovering("QHostInfo::lookupId")

	if ptr.Pointer() != nil {
		return int(C.QHostInfo_LookupId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostInfo) SetError(error QHostInfo__HostInfoError) {
	defer qt.Recovering("QHostInfo::setError")

	if ptr.Pointer() != nil {
		C.QHostInfo_SetError(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QHostInfo) SetErrorString(str string) {
	defer qt.Recovering("QHostInfo::setErrorString")

	if ptr.Pointer() != nil {
		C.QHostInfo_SetErrorString(ptr.Pointer(), C.CString(str))
	}
}

func (ptr *QHostInfo) SetHostName(hostName string) {
	defer qt.Recovering("QHostInfo::setHostName")

	if ptr.Pointer() != nil {
		C.QHostInfo_SetHostName(ptr.Pointer(), C.CString(hostName))
	}
}

func (ptr *QHostInfo) SetLookupId(id int) {
	defer qt.Recovering("QHostInfo::setLookupId")

	if ptr.Pointer() != nil {
		C.QHostInfo_SetLookupId(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QHostInfo) DestroyQHostInfo() {
	defer qt.Recovering("QHostInfo::~QHostInfo")

	if ptr.Pointer() != nil {
		C.QHostInfo_DestroyQHostInfo(ptr.Pointer())
	}
}

func QHostInfo_LocalHostName() string {
	defer qt.Recovering("QHostInfo::localHostName")

	return C.GoString(C.QHostInfo_QHostInfo_LocalHostName())
}

func QHostInfo_LocalDomainName() string {
	defer qt.Recovering("QHostInfo::localDomainName")

	return C.GoString(C.QHostInfo_QHostInfo_LocalDomainName())
}

type QHttpMultiPart struct {
	core.QObject
}

type QHttpMultiPart_ITF interface {
	core.QObject_ITF
	QHttpMultiPart_PTR() *QHttpMultiPart
}

func PointerFromQHttpMultiPart(ptr QHttpMultiPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpMultiPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpMultiPartFromPointer(ptr unsafe.Pointer) *QHttpMultiPart {
	var n = new(QHttpMultiPart)
	n.SetPointer(ptr)
	return n
}

func newQHttpMultiPartFromPointer(ptr unsafe.Pointer) *QHttpMultiPart {
	var n = NewQHttpMultiPartFromPointer(ptr)
	for len(n.ObjectName()) < len("QHttpMultiPart_") {
		n.SetObjectName("QHttpMultiPart_" + qt.Identifier())
	}
	return n
}

func (ptr *QHttpMultiPart) QHttpMultiPart_PTR() *QHttpMultiPart {
	return ptr
}

//QHttpMultiPart::ContentType
type QHttpMultiPart__ContentType int64

const (
	QHttpMultiPart__MixedType       = QHttpMultiPart__ContentType(0)
	QHttpMultiPart__RelatedType     = QHttpMultiPart__ContentType(1)
	QHttpMultiPart__FormDataType    = QHttpMultiPart__ContentType(2)
	QHttpMultiPart__AlternativeType = QHttpMultiPart__ContentType(3)
)

func NewQHttpMultiPart2(contentType QHttpMultiPart__ContentType, parent core.QObject_ITF) *QHttpMultiPart {
	defer qt.Recovering("QHttpMultiPart::QHttpMultiPart")

	return newQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart2(C.int(contentType), core.PointerFromQObject(parent)))
}

func NewQHttpMultiPart(parent core.QObject_ITF) *QHttpMultiPart {
	defer qt.Recovering("QHttpMultiPart::QHttpMultiPart")

	return newQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart(core.PointerFromQObject(parent)))
}

func (ptr *QHttpMultiPart) Append(httpPart QHttpPart_ITF) {
	defer qt.Recovering("QHttpMultiPart::append")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_Append(ptr.Pointer(), PointerFromQHttpPart(httpPart))
	}
}

func (ptr *QHttpMultiPart) Boundary() *core.QByteArray {
	defer qt.Recovering("QHttpMultiPart::boundary")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QHttpMultiPart_Boundary(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHttpMultiPart) SetBoundary(boundary core.QByteArray_ITF) {
	defer qt.Recovering("QHttpMultiPart::setBoundary")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetBoundary(ptr.Pointer(), core.PointerFromQByteArray(boundary))
	}
}

func (ptr *QHttpMultiPart) SetContentType(contentType QHttpMultiPart__ContentType) {
	defer qt.Recovering("QHttpMultiPart::setContentType")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetContentType(ptr.Pointer(), C.int(contentType))
	}
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPart() {
	defer qt.Recovering("QHttpMultiPart::~QHttpMultiPart")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DestroyQHttpMultiPart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHttpMultiPart) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHttpMultiPartTimerEvent
func callbackQHttpMultiPartTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHttpMultiPart) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::timerEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QHttpMultiPart) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHttpMultiPartChildEvent
func callbackQHttpMultiPartChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHttpMultiPart) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::childEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QHttpMultiPart) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHttpMultiPart) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHttpMultiPartCustomEvent
func callbackQHttpMultiPartCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QHttpMultiPart::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQHttpMultiPartFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QHttpMultiPart) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QHttpMultiPart) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QHttpMultiPart::customEvent")

	if ptr.Pointer() != nil {
		C.QHttpMultiPart_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QHttpPart struct {
	ptr unsafe.Pointer
}

type QHttpPart_ITF interface {
	QHttpPart_PTR() *QHttpPart
}

func (p *QHttpPart) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHttpPart) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHttpPart(ptr QHttpPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpPartFromPointer(ptr unsafe.Pointer) *QHttpPart {
	var n = new(QHttpPart)
	n.SetPointer(ptr)
	return n
}

func newQHttpPartFromPointer(ptr unsafe.Pointer) *QHttpPart {
	var n = NewQHttpPartFromPointer(ptr)
	return n
}

func (ptr *QHttpPart) QHttpPart_PTR() *QHttpPart {
	return ptr
}

func NewQHttpPart() *QHttpPart {
	defer qt.Recovering("QHttpPart::QHttpPart")

	return newQHttpPartFromPointer(C.QHttpPart_NewQHttpPart())
}

func NewQHttpPart2(other QHttpPart_ITF) *QHttpPart {
	defer qt.Recovering("QHttpPart::QHttpPart")

	return newQHttpPartFromPointer(C.QHttpPart_NewQHttpPart2(PointerFromQHttpPart(other)))
}

func (ptr *QHttpPart) SetBody(body core.QByteArray_ITF) {
	defer qt.Recovering("QHttpPart::setBody")

	if ptr.Pointer() != nil {
		C.QHttpPart_SetBody(ptr.Pointer(), core.PointerFromQByteArray(body))
	}
}

func (ptr *QHttpPart) SetBodyDevice(device core.QIODevice_ITF) {
	defer qt.Recovering("QHttpPart::setBodyDevice")

	if ptr.Pointer() != nil {
		C.QHttpPart_SetBodyDevice(ptr.Pointer(), core.PointerFromQIODevice(device))
	}
}

func (ptr *QHttpPart) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	defer qt.Recovering("QHttpPart::setHeader")

	if ptr.Pointer() != nil {
		C.QHttpPart_SetHeader(ptr.Pointer(), C.int(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QHttpPart) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {
	defer qt.Recovering("QHttpPart::setRawHeader")

	if ptr.Pointer() != nil {
		C.QHttpPart_SetRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(headerValue))
	}
}

func (ptr *QHttpPart) Swap(other QHttpPart_ITF) {
	defer qt.Recovering("QHttpPart::swap")

	if ptr.Pointer() != nil {
		C.QHttpPart_Swap(ptr.Pointer(), PointerFromQHttpPart(other))
	}
}

func (ptr *QHttpPart) DestroyQHttpPart() {
	defer qt.Recovering("QHttpPart::~QHttpPart")

	if ptr.Pointer() != nil {
		C.QHttpPart_DestroyQHttpPart(ptr.Pointer())
	}
}

type QLocalServer struct {
	core.QObject
}

type QLocalServer_ITF interface {
	core.QObject_ITF
	QLocalServer_PTR() *QLocalServer
}

func PointerFromQLocalServer(ptr QLocalServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalServer_PTR().Pointer()
	}
	return nil
}

func NewQLocalServerFromPointer(ptr unsafe.Pointer) *QLocalServer {
	var n = new(QLocalServer)
	n.SetPointer(ptr)
	return n
}

func newQLocalServerFromPointer(ptr unsafe.Pointer) *QLocalServer {
	var n = NewQLocalServerFromPointer(ptr)
	for len(n.ObjectName()) < len("QLocalServer_") {
		n.SetObjectName("QLocalServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QLocalServer) QLocalServer_PTR() *QLocalServer {
	return ptr
}

//QLocalServer::SocketOption
type QLocalServer__SocketOption int64

const (
	QLocalServer__NoOptions         = QLocalServer__SocketOption(0x0)
	QLocalServer__UserAccessOption  = QLocalServer__SocketOption(0x01)
	QLocalServer__GroupAccessOption = QLocalServer__SocketOption(0x2)
	QLocalServer__OtherAccessOption = QLocalServer__SocketOption(0x4)
	QLocalServer__WorldAccessOption = QLocalServer__SocketOption(0x7)
)

func (ptr *QLocalServer) SetSocketOptions(options QLocalServer__SocketOption) {
	defer qt.Recovering("QLocalServer::setSocketOptions")

	if ptr.Pointer() != nil {
		C.QLocalServer_SetSocketOptions(ptr.Pointer(), C.int(options))
	}
}

func NewQLocalServer(parent core.QObject_ITF) *QLocalServer {
	defer qt.Recovering("QLocalServer::QLocalServer")

	return newQLocalServerFromPointer(C.QLocalServer_NewQLocalServer(core.PointerFromQObject(parent)))
}

func (ptr *QLocalServer) Close() {
	defer qt.Recovering("QLocalServer::close")

	if ptr.Pointer() != nil {
		C.QLocalServer_Close(ptr.Pointer())
	}
}

func (ptr *QLocalServer) ErrorString() string {
	defer qt.Recovering("QLocalServer::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) FullServerName() string {
	defer qt.Recovering("QLocalServer::fullServerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_FullServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) HasPendingConnections() bool {
	defer qt.Recovering("QLocalServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QLocalServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalServer) IsListening() bool {
	defer qt.Recovering("QLocalServer::isListening")

	if ptr.Pointer() != nil {
		return C.QLocalServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalServer) Listen(name string) bool {
	defer qt.Recovering("QLocalServer::listen")

	if ptr.Pointer() != nil {
		return C.QLocalServer_Listen(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QLocalServer) MaxPendingConnections() int {
	defer qt.Recovering("QLocalServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(C.QLocalServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QLocalServer::newConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QLocalServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QLocalServer::newConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQLocalServerNewConnection
func callbackQLocalServerNewConnection(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLocalServer::newConnection")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalServer) NewConnection() {
	defer qt.Recovering("QLocalServer::newConnection")

	if ptr.Pointer() != nil {
		C.QLocalServer_NewConnection(ptr.Pointer())
	}
}

func (ptr *QLocalServer) NextPendingConnection() *QLocalSocket {
	defer qt.Recovering("QLocalServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQLocalSocketFromPointer(C.QLocalServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func QLocalServer_RemoveServer(name string) bool {
	defer qt.Recovering("QLocalServer::removeServer")

	return C.QLocalServer_QLocalServer_RemoveServer(C.CString(name)) != 0
}

func (ptr *QLocalServer) ServerError() QAbstractSocket__SocketError {
	defer qt.Recovering("QLocalServer::serverError")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QLocalServer_ServerError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) ServerName() string {
	defer qt.Recovering("QLocalServer::serverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalServer_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QLocalServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QLocalServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

func (ptr *QLocalServer) SocketOptions() QLocalServer__SocketOption {
	defer qt.Recovering("QLocalServer::socketOptions")

	if ptr.Pointer() != nil {
		return QLocalServer__SocketOption(C.QLocalServer_SocketOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalServer) WaitForNewConnection(msec int, timedOut bool) bool {
	defer qt.Recovering("QLocalServer::waitForNewConnection")

	if ptr.Pointer() != nil {
		return C.QLocalServer_WaitForNewConnection(ptr.Pointer(), C.int(msec), C.int(qt.GoBoolToInt(timedOut))) != 0
	}
	return false
}

func (ptr *QLocalServer) DestroyQLocalServer() {
	defer qt.Recovering("QLocalServer::~QLocalServer")

	if ptr.Pointer() != nil {
		C.QLocalServer_DestroyQLocalServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLocalServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLocalServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLocalServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLocalServerTimerEvent
func callbackQLocalServerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLocalServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLocalServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLocalServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLocalServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLocalServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLocalServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLocalServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLocalServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLocalServerChildEvent
func callbackQLocalServerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLocalServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLocalServer::childEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLocalServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLocalServer::childEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLocalServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLocalServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLocalServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLocalServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLocalServerCustomEvent
func callbackQLocalServerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLocalServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLocalServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLocalServer::customEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLocalServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLocalServer::customEvent")

	if ptr.Pointer() != nil {
		C.QLocalServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QLocalSocket struct {
	core.QIODevice
}

type QLocalSocket_ITF interface {
	core.QIODevice_ITF
	QLocalSocket_PTR() *QLocalSocket
}

func PointerFromQLocalSocket(ptr QLocalSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalSocket_PTR().Pointer()
	}
	return nil
}

func NewQLocalSocketFromPointer(ptr unsafe.Pointer) *QLocalSocket {
	var n = new(QLocalSocket)
	n.SetPointer(ptr)
	return n
}

func newQLocalSocketFromPointer(ptr unsafe.Pointer) *QLocalSocket {
	var n = NewQLocalSocketFromPointer(ptr)
	for len(n.ObjectName()) < len("QLocalSocket_") {
		n.SetObjectName("QLocalSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QLocalSocket) QLocalSocket_PTR() *QLocalSocket {
	return ptr
}

//QLocalSocket::LocalSocketError
type QLocalSocket__LocalSocketError int64

const (
	QLocalSocket__ConnectionRefusedError          = QLocalSocket__LocalSocketError(QAbstractSocket__ConnectionRefusedError)
	QLocalSocket__PeerClosedError                 = QLocalSocket__LocalSocketError(QAbstractSocket__RemoteHostClosedError)
	QLocalSocket__ServerNotFoundError             = QLocalSocket__LocalSocketError(QAbstractSocket__HostNotFoundError)
	QLocalSocket__SocketAccessError               = QLocalSocket__LocalSocketError(QAbstractSocket__SocketAccessError)
	QLocalSocket__SocketResourceError             = QLocalSocket__LocalSocketError(QAbstractSocket__SocketResourceError)
	QLocalSocket__SocketTimeoutError              = QLocalSocket__LocalSocketError(QAbstractSocket__SocketTimeoutError)
	QLocalSocket__DatagramTooLargeError           = QLocalSocket__LocalSocketError(QAbstractSocket__DatagramTooLargeError)
	QLocalSocket__ConnectionError                 = QLocalSocket__LocalSocketError(QAbstractSocket__NetworkError)
	QLocalSocket__UnsupportedSocketOperationError = QLocalSocket__LocalSocketError(QAbstractSocket__UnsupportedSocketOperationError)
	QLocalSocket__UnknownSocketError              = QLocalSocket__LocalSocketError(QAbstractSocket__UnknownSocketError)
	QLocalSocket__OperationError                  = QLocalSocket__LocalSocketError(QAbstractSocket__OperationError)
)

//QLocalSocket::LocalSocketState
type QLocalSocket__LocalSocketState int64

const (
	QLocalSocket__UnconnectedState = QLocalSocket__LocalSocketState(QAbstractSocket__UnconnectedState)
	QLocalSocket__ConnectingState  = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectingState)
	QLocalSocket__ConnectedState   = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectedState)
	QLocalSocket__ClosingState     = QLocalSocket__LocalSocketState(QAbstractSocket__ClosingState)
)

func (ptr *QLocalSocket) Open(openMode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QLocalSocket::open")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Open(ptr.Pointer(), C.int(openMode)) != 0
	}
	return false
}

func NewQLocalSocket(parent core.QObject_ITF) *QLocalSocket {
	defer qt.Recovering("QLocalSocket::QLocalSocket")

	return newQLocalSocketFromPointer(C.QLocalSocket_NewQLocalSocket(core.PointerFromQObject(parent)))
}

func (ptr *QLocalSocket) ConnectToServer2(name string, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QLocalSocket::connectToServer")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer2(ptr.Pointer(), C.CString(name), C.int(openMode))
	}
}

func (ptr *QLocalSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QLocalSocket::connected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QLocalSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QLocalSocket::connected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

//export callbackQLocalSocketConnected
func callbackQLocalSocketConnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLocalSocket::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalSocket) Connected() {
	defer qt.Recovering("QLocalSocket::connected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Connected(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QLocalSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QLocalSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QLocalSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

//export callbackQLocalSocketDisconnected
func callbackQLocalSocketDisconnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLocalSocket::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLocalSocket) Disconnected() {
	defer qt.Recovering("QLocalSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Disconnected(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) ConnectError2(f func(socketError QLocalSocket__LocalSocketError)) {
	defer qt.Recovering("connect QLocalSocket::error")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QLocalSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QLocalSocket::error")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQLocalSocketError2
func callbackQLocalSocketError2(ptr unsafe.Pointer, ptrName *C.char, socketError C.int) {
	defer qt.Recovering("callback QLocalSocket::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QLocalSocket__LocalSocketError))(QLocalSocket__LocalSocketError(socketError))
	}

}

func (ptr *QLocalSocket) Error2(socketError QLocalSocket__LocalSocketError) {
	defer qt.Recovering("QLocalSocket::error")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Error2(ptr.Pointer(), C.int(socketError))
	}
}

func (ptr *QLocalSocket) FullServerName() string {
	defer qt.Recovering("QLocalSocket::fullServerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_FullServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) IsSequential() bool {
	defer qt.Recovering("QLocalSocket::isSequential")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ServerName() string {
	defer qt.Recovering("QLocalSocket::serverName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocalSocket_ServerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocalSocket) SetServerName(name string) {
	defer qt.Recovering("QLocalSocket::setServerName")

	if ptr.Pointer() != nil {
		C.QLocalSocket_SetServerName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QLocalSocket) State() QLocalSocket__LocalSocketState {
	defer qt.Recovering("QLocalSocket::state")

	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketState(C.QLocalSocket_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) ConnectStateChanged(f func(socketState QLocalSocket__LocalSocketState)) {
	defer qt.Recovering("connect QLocalSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLocalSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QLocalSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQLocalSocketStateChanged
func callbackQLocalSocketStateChanged(ptr unsafe.Pointer, ptrName *C.char, socketState C.int) {
	defer qt.Recovering("callback QLocalSocket::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QLocalSocket__LocalSocketState))(QLocalSocket__LocalSocketState(socketState))
	}

}

func (ptr *QLocalSocket) StateChanged(socketState QLocalSocket__LocalSocketState) {
	defer qt.Recovering("QLocalSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QLocalSocket_StateChanged(ptr.Pointer(), C.int(socketState))
	}
}

func (ptr *QLocalSocket) DestroyQLocalSocket() {
	defer qt.Recovering("QLocalSocket::~QLocalSocket")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DestroyQLocalSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLocalSocket) Abort() {
	defer qt.Recovering("QLocalSocket::abort")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) BytesAvailable() int64 {
	defer qt.Recovering("QLocalSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) BytesToWrite() int64 {
	defer qt.Recovering("QLocalSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) CanReadLine() bool {
	defer qt.Recovering("QLocalSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QLocalSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QLocalSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QLocalSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQLocalSocketClose
func callbackQLocalSocketClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLocalSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQLocalSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QLocalSocket) Close() {
	defer qt.Recovering("QLocalSocket::close")

	if ptr.Pointer() != nil {
		C.QLocalSocket_Close(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) CloseDefault() {
	defer qt.Recovering("QLocalSocket::close")

	if ptr.Pointer() != nil {
		C.QLocalSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) ConnectToServer(openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QLocalSocket::connectToServer")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ConnectToServer(ptr.Pointer(), C.int(openMode))
	}
}

func (ptr *QLocalSocket) DisconnectFromServer() {
	defer qt.Recovering("QLocalSocket::disconnectFromServer")

	if ptr.Pointer() != nil {
		C.QLocalSocket_DisconnectFromServer(ptr.Pointer())
	}
}

func (ptr *QLocalSocket) Error() QLocalSocket__LocalSocketError {
	defer qt.Recovering("QLocalSocket::error")

	if ptr.Pointer() != nil {
		return QLocalSocket__LocalSocketError(C.QLocalSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) Flush() bool {
	defer qt.Recovering("QLocalSocket::flush")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) IsValid() bool {
	defer qt.Recovering("QLocalSocket::isValid")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLocalSocket) ReadBufferSize() int64 {
	defer qt.Recovering("QLocalSocket::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocalSocket) ReadData(data string, c int64) int64 {
	defer qt.Recovering("QLocalSocket::readData")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_ReadData(ptr.Pointer(), C.CString(data), C.longlong(c)))
	}
	return 0
}

func (ptr *QLocalSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QLocalSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QLocalSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QLocalSocket) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForConnected(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForConnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForDisconnected(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForDisconnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QLocalSocket::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QLocalSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QLocalSocket) WriteData(data string, c int64) int64 {
	defer qt.Recovering("QLocalSocket::writeData")

	if ptr.Pointer() != nil {
		return int64(C.QLocalSocket_WriteData(ptr.Pointer(), C.CString(data), C.longlong(c)))
	}
	return 0
}

func (ptr *QLocalSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLocalSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLocalSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLocalSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQLocalSocketTimerEvent
func callbackQLocalSocketTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLocalSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLocalSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLocalSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLocalSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLocalSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLocalSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLocalSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQLocalSocketChildEvent
func callbackQLocalSocketChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLocalSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLocalSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLocalSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLocalSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLocalSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLocalSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLocalSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQLocalSocketCustomEvent
func callbackQLocalSocketCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLocalSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLocalSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLocalSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLocalSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLocalSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLocalSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QLocalSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QNetworkAccessManager struct {
	core.QObject
}

type QNetworkAccessManager_ITF interface {
	core.QObject_ITF
	QNetworkAccessManager_PTR() *QNetworkAccessManager
}

func PointerFromQNetworkAccessManager(ptr QNetworkAccessManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAccessManager_PTR().Pointer()
	}
	return nil
}

func NewQNetworkAccessManagerFromPointer(ptr unsafe.Pointer) *QNetworkAccessManager {
	var n = new(QNetworkAccessManager)
	n.SetPointer(ptr)
	return n
}

func newQNetworkAccessManagerFromPointer(ptr unsafe.Pointer) *QNetworkAccessManager {
	var n = NewQNetworkAccessManagerFromPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkAccessManager_") {
		n.SetObjectName("QNetworkAccessManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkAccessManager) QNetworkAccessManager_PTR() *QNetworkAccessManager {
	return ptr
}

//QNetworkAccessManager::NetworkAccessibility
type QNetworkAccessManager__NetworkAccessibility int64

const (
	QNetworkAccessManager__UnknownAccessibility = QNetworkAccessManager__NetworkAccessibility(-1)
	QNetworkAccessManager__NotAccessible        = QNetworkAccessManager__NetworkAccessibility(0)
	QNetworkAccessManager__Accessible           = QNetworkAccessManager__NetworkAccessibility(1)
)

//QNetworkAccessManager::Operation
type QNetworkAccessManager__Operation int64

const (
	QNetworkAccessManager__HeadOperation    = QNetworkAccessManager__Operation(1)
	QNetworkAccessManager__GetOperation     = QNetworkAccessManager__Operation(2)
	QNetworkAccessManager__PutOperation     = QNetworkAccessManager__Operation(3)
	QNetworkAccessManager__PostOperation    = QNetworkAccessManager__Operation(4)
	QNetworkAccessManager__DeleteOperation  = QNetworkAccessManager__Operation(5)
	QNetworkAccessManager__CustomOperation  = QNetworkAccessManager__Operation(6)
	QNetworkAccessManager__UnknownOperation = QNetworkAccessManager__Operation(0)
)

func (ptr *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory {
	defer qt.Recovering("QNetworkAccessManager::proxyFactory")

	if ptr.Pointer() != nil {
		return NewQNetworkProxyFactoryFromPointer(C.QNetworkAccessManager_ProxyFactory(ptr.Pointer()))
	}
	return nil
}

func NewQNetworkAccessManager(parent core.QObject_ITF) *QNetworkAccessManager {
	defer qt.Recovering("QNetworkAccessManager::QNetworkAccessManager")

	return newQNetworkAccessManagerFromPointer(C.QNetworkAccessManager_NewQNetworkAccessManager(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkAccessManager) ConnectAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QAuthenticator)) {
	defer qt.Recovering("connect QNetworkAccessManager::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "authenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectAuthenticationRequired() {
	defer qt.Recovering("disconnect QNetworkAccessManager::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "authenticationRequired")
	}
}

//export callbackQNetworkAccessManagerAuthenticationRequired
func callbackQNetworkAccessManagerAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::authenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "authenticationRequired"); signal != nil {
		signal.(func(*QNetworkReply, *QAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) AuthenticationRequired(reply QNetworkReply_ITF, authenticator QAuthenticator_ITF) {
	defer qt.Recovering("QNetworkAccessManager::authenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_AuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Cache() *QAbstractNetworkCache {
	defer qt.Recovering("QNetworkAccessManager::cache")

	if ptr.Pointer() != nil {
		return NewQAbstractNetworkCacheFromPointer(C.QNetworkAccessManager_Cache(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ClearAccessCache() {
	defer qt.Recovering("QNetworkAccessManager::clearAccessCache")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ClearAccessCache(ptr.Pointer())
	}
}

func (ptr *QNetworkAccessManager) CookieJar() *QNetworkCookieJar {
	defer qt.Recovering("QNetworkAccessManager::cookieJar")

	if ptr.Pointer() != nil {
		return NewQNetworkCookieJarFromPointer(C.QNetworkAccessManager_CookieJar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkAccessManager) CreateRequest(op QNetworkAccessManager__Operation, req QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::createRequest")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_CreateRequest(ptr.Pointer(), C.int(op), PointerFromQNetworkRequest(req), core.PointerFromQIODevice(outgoingData)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) DeleteResource(request QNetworkRequest_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::deleteResource")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_DeleteResource(ptr.Pointer(), PointerFromQNetworkRequest(request)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectEncrypted(f func(reply *QNetworkReply)) {
	defer qt.Recovering("connect QNetworkAccessManager::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectEncrypted() {
	defer qt.Recovering("disconnect QNetworkAccessManager::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQNetworkAccessManagerEncrypted
func callbackQNetworkAccessManagerEncrypted(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::encrypted")

	if signal := qt.GetSignal(C.GoString(ptrName), "encrypted"); signal != nil {
		signal.(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) Encrypted(reply QNetworkReply_ITF) {
	defer qt.Recovering("QNetworkAccessManager::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Encrypted(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

func (ptr *QNetworkAccessManager) ConnectFinished(f func(reply *QNetworkReply)) {
	defer qt.Recovering("connect QNetworkAccessManager::finished")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectFinished() {
	defer qt.Recovering("disconnect QNetworkAccessManager::finished")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQNetworkAccessManagerFinished
func callbackQNetworkAccessManagerFinished(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QNetworkReply))(NewQNetworkReplyFromPointer(reply))
	}

}

func (ptr *QNetworkAccessManager) Finished(reply QNetworkReply_ITF) {
	defer qt.Recovering("QNetworkAccessManager::finished")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_Finished(ptr.Pointer(), PointerFromQNetworkReply(reply))
	}
}

func (ptr *QNetworkAccessManager) Get(request QNetworkRequest_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::get")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Get(ptr.Pointer(), PointerFromQNetworkRequest(request)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Head(request QNetworkRequest_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::head")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Head(ptr.Pointer(), PointerFromQNetworkRequest(request)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) NetworkAccessible() QNetworkAccessManager__NetworkAccessibility {
	defer qt.Recovering("QNetworkAccessManager::networkAccessible")

	if ptr.Pointer() != nil {
		return QNetworkAccessManager__NetworkAccessibility(C.QNetworkAccessManager_NetworkAccessible(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkAccessManager) ConnectNetworkAccessibleChanged(f func(accessible QNetworkAccessManager__NetworkAccessibility)) {
	defer qt.Recovering("connect QNetworkAccessManager::networkAccessibleChanged")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectNetworkAccessibleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "networkAccessibleChanged", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectNetworkAccessibleChanged() {
	defer qt.Recovering("disconnect QNetworkAccessManager::networkAccessibleChanged")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectNetworkAccessibleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "networkAccessibleChanged")
	}
}

//export callbackQNetworkAccessManagerNetworkAccessibleChanged
func callbackQNetworkAccessManagerNetworkAccessibleChanged(ptr unsafe.Pointer, ptrName *C.char, accessible C.int) {
	defer qt.Recovering("callback QNetworkAccessManager::networkAccessibleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "networkAccessibleChanged"); signal != nil {
		signal.(func(QNetworkAccessManager__NetworkAccessibility))(QNetworkAccessManager__NetworkAccessibility(accessible))
	}

}

func (ptr *QNetworkAccessManager) NetworkAccessibleChanged(accessible QNetworkAccessManager__NetworkAccessibility) {
	defer qt.Recovering("QNetworkAccessManager::networkAccessibleChanged")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_NetworkAccessibleChanged(ptr.Pointer(), C.int(accessible))
	}
}

func (ptr *QNetworkAccessManager) Post3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::post")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post3(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::post")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Post2(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::post")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Post2(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) ConnectPreSharedKeyAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator)) {
	defer qt.Recovering("connect QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectPreSharedKeyAuthenticationRequired() {
	defer qt.Recovering("disconnect QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired
func callbackQNetworkAccessManagerPreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QNetworkReply, *QSslPreSharedKeyAuthenticator))(NewQNetworkReplyFromPointer(reply), NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply QNetworkReply_ITF, authenticator QSslPreSharedKeyAuthenticator_ITF) {
	defer qt.Recovering("QNetworkAccessManager::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQNetworkReply(reply), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QNetworkAccessManager) Put2(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::put")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put2(ptr.Pointer(), PointerFromQNetworkRequest(request), PointerFromQHttpMultiPart(multiPart)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::put")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) Put3(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::put")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_Put3(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) SendCustomRequest(request QNetworkRequest_ITF, verb core.QByteArray_ITF, data core.QIODevice_ITF) *QNetworkReply {
	defer qt.Recovering("QNetworkAccessManager::sendCustomRequest")

	if ptr.Pointer() != nil {
		return NewQNetworkReplyFromPointer(C.QNetworkAccessManager_SendCustomRequest(ptr.Pointer(), PointerFromQNetworkRequest(request), core.PointerFromQByteArray(verb), core.PointerFromQIODevice(data)))
	}
	return nil
}

func (ptr *QNetworkAccessManager) SetCache(cache QAbstractNetworkCache_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setCache")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCache(ptr.Pointer(), PointerFromQAbstractNetworkCache(cache))
	}
}

func (ptr *QNetworkAccessManager) SetConfiguration(config QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetConfiguration(ptr.Pointer(), PointerFromQNetworkConfiguration(config))
	}
}

func (ptr *QNetworkAccessManager) SetCookieJar(cookieJar QNetworkCookieJar_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setCookieJar")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetCookieJar(ptr.Pointer(), PointerFromQNetworkCookieJar(cookieJar))
	}
}

func (ptr *QNetworkAccessManager) SetNetworkAccessible(accessible QNetworkAccessManager__NetworkAccessibility) {
	defer qt.Recovering("QNetworkAccessManager::setNetworkAccessible")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetNetworkAccessible(ptr.Pointer(), C.int(accessible))
	}
}

func (ptr *QNetworkAccessManager) SetProxy(proxy QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setProxy")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(proxy))
	}
}

func (ptr *QNetworkAccessManager) SetProxyFactory(factory QNetworkProxyFactory_ITF) {
	defer qt.Recovering("QNetworkAccessManager::setProxyFactory")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_SetProxyFactory(ptr.Pointer(), PointerFromQNetworkProxyFactory(factory))
	}
}

func (ptr *QNetworkAccessManager) SupportedSchemes() []string {
	defer qt.Recovering("QNetworkAccessManager::supportedSchemes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QNetworkAccessManager_SupportedSchemes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManager() {
	defer qt.Recovering("QNetworkAccessManager::~QNetworkAccessManager")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_DestroyQNetworkAccessManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkAccessManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkAccessManagerTimerEvent
func callbackQNetworkAccessManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkAccessManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkAccessManagerChildEvent
func callbackQNetworkAccessManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkAccessManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkAccessManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkAccessManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkAccessManagerCustomEvent
func callbackQNetworkAccessManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkAccessManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkAccessManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkAccessManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkAccessManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkAccessManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkAccessManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QNetworkAddressEntry struct {
	ptr unsafe.Pointer
}

type QNetworkAddressEntry_ITF interface {
	QNetworkAddressEntry_PTR() *QNetworkAddressEntry
}

func (p *QNetworkAddressEntry) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkAddressEntry) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkAddressEntry(ptr QNetworkAddressEntry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAddressEntry_PTR().Pointer()
	}
	return nil
}

func NewQNetworkAddressEntryFromPointer(ptr unsafe.Pointer) *QNetworkAddressEntry {
	var n = new(QNetworkAddressEntry)
	n.SetPointer(ptr)
	return n
}

func newQNetworkAddressEntryFromPointer(ptr unsafe.Pointer) *QNetworkAddressEntry {
	var n = NewQNetworkAddressEntryFromPointer(ptr)
	return n
}

func (ptr *QNetworkAddressEntry) QNetworkAddressEntry_PTR() *QNetworkAddressEntry {
	return ptr
}

func NewQNetworkAddressEntry() *QNetworkAddressEntry {
	defer qt.Recovering("QNetworkAddressEntry::QNetworkAddressEntry")

	return newQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry())
}

func NewQNetworkAddressEntry2(other QNetworkAddressEntry_ITF) *QNetworkAddressEntry {
	defer qt.Recovering("QNetworkAddressEntry::QNetworkAddressEntry")

	return newQNetworkAddressEntryFromPointer(C.QNetworkAddressEntry_NewQNetworkAddressEntry2(PointerFromQNetworkAddressEntry(other)))
}

func (ptr *QNetworkAddressEntry) PrefixLength() int {
	defer qt.Recovering("QNetworkAddressEntry::prefixLength")

	if ptr.Pointer() != nil {
		return int(C.QNetworkAddressEntry_PrefixLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkAddressEntry) SetBroadcast(newBroadcast QHostAddress_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::setBroadcast")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetBroadcast(ptr.Pointer(), PointerFromQHostAddress(newBroadcast))
	}
}

func (ptr *QNetworkAddressEntry) SetIp(newIp QHostAddress_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::setIp")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetIp(ptr.Pointer(), PointerFromQHostAddress(newIp))
	}
}

func (ptr *QNetworkAddressEntry) SetNetmask(newNetmask QHostAddress_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::setNetmask")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetNetmask(ptr.Pointer(), PointerFromQHostAddress(newNetmask))
	}
}

func (ptr *QNetworkAddressEntry) SetPrefixLength(length int) {
	defer qt.Recovering("QNetworkAddressEntry::setPrefixLength")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_SetPrefixLength(ptr.Pointer(), C.int(length))
	}
}

func (ptr *QNetworkAddressEntry) Swap(other QNetworkAddressEntry_ITF) {
	defer qt.Recovering("QNetworkAddressEntry::swap")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_Swap(ptr.Pointer(), PointerFromQNetworkAddressEntry(other))
	}
}

func (ptr *QNetworkAddressEntry) DestroyQNetworkAddressEntry() {
	defer qt.Recovering("QNetworkAddressEntry::~QNetworkAddressEntry")

	if ptr.Pointer() != nil {
		C.QNetworkAddressEntry_DestroyQNetworkAddressEntry(ptr.Pointer())
	}
}

type QNetworkCacheMetaData struct {
	ptr unsafe.Pointer
}

type QNetworkCacheMetaData_ITF interface {
	QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData
}

func (p *QNetworkCacheMetaData) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkCacheMetaData) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkCacheMetaData(ptr QNetworkCacheMetaData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCacheMetaData_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCacheMetaDataFromPointer(ptr unsafe.Pointer) *QNetworkCacheMetaData {
	var n = new(QNetworkCacheMetaData)
	n.SetPointer(ptr)
	return n
}

func newQNetworkCacheMetaDataFromPointer(ptr unsafe.Pointer) *QNetworkCacheMetaData {
	var n = NewQNetworkCacheMetaDataFromPointer(ptr)
	return n
}

func (ptr *QNetworkCacheMetaData) QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData {
	return ptr
}

func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {
	defer qt.Recovering("QNetworkCacheMetaData::QNetworkCacheMetaData")

	return newQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData())
}

func NewQNetworkCacheMetaData2(other QNetworkCacheMetaData_ITF) *QNetworkCacheMetaData {
	defer qt.Recovering("QNetworkCacheMetaData::QNetworkCacheMetaData")

	return newQNetworkCacheMetaDataFromPointer(C.QNetworkCacheMetaData_NewQNetworkCacheMetaData2(PointerFromQNetworkCacheMetaData(other)))
}

func (ptr *QNetworkCacheMetaData) ExpirationDate() *core.QDateTime {
	defer qt.Recovering("QNetworkCacheMetaData::expirationDate")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_ExpirationDate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) IsValid() bool {
	defer qt.Recovering("QNetworkCacheMetaData::isValid")

	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) LastModified() *core.QDateTime {
	defer qt.Recovering("QNetworkCacheMetaData::lastModified")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QNetworkCacheMetaData_LastModified(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) SaveToDisk() bool {
	defer qt.Recovering("QNetworkCacheMetaData::saveToDisk")

	if ptr.Pointer() != nil {
		return C.QNetworkCacheMetaData_SaveToDisk(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCacheMetaData) SetExpirationDate(dateTime core.QDateTime_ITF) {
	defer qt.Recovering("QNetworkCacheMetaData::setExpirationDate")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetExpirationDate(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetLastModified(dateTime core.QDateTime_ITF) {
	defer qt.Recovering("QNetworkCacheMetaData::setLastModified")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetLastModified(ptr.Pointer(), core.PointerFromQDateTime(dateTime))
	}
}

func (ptr *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {
	defer qt.Recovering("QNetworkCacheMetaData::setSaveToDisk")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetSaveToDisk(ptr.Pointer(), C.int(qt.GoBoolToInt(allow)))
	}
}

func (ptr *QNetworkCacheMetaData) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkCacheMetaData::setUrl")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkCacheMetaData) Swap(other QNetworkCacheMetaData_ITF) {
	defer qt.Recovering("QNetworkCacheMetaData::swap")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_Swap(ptr.Pointer(), PointerFromQNetworkCacheMetaData(other))
	}
}

func (ptr *QNetworkCacheMetaData) Url() *core.QUrl {
	defer qt.Recovering("QNetworkCacheMetaData::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QNetworkCacheMetaData_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) DestroyQNetworkCacheMetaData() {
	defer qt.Recovering("QNetworkCacheMetaData::~QNetworkCacheMetaData")

	if ptr.Pointer() != nil {
		C.QNetworkCacheMetaData_DestroyQNetworkCacheMetaData(ptr.Pointer())
	}
}

type QNetworkConfiguration struct {
	ptr unsafe.Pointer
}

type QNetworkConfiguration_ITF interface {
	QNetworkConfiguration_PTR() *QNetworkConfiguration
}

func (p *QNetworkConfiguration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkConfiguration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkConfiguration(ptr QNetworkConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfiguration_PTR().Pointer()
	}
	return nil
}

func NewQNetworkConfigurationFromPointer(ptr unsafe.Pointer) *QNetworkConfiguration {
	var n = new(QNetworkConfiguration)
	n.SetPointer(ptr)
	return n
}

func newQNetworkConfigurationFromPointer(ptr unsafe.Pointer) *QNetworkConfiguration {
	var n = NewQNetworkConfigurationFromPointer(ptr)
	return n
}

func (ptr *QNetworkConfiguration) QNetworkConfiguration_PTR() *QNetworkConfiguration {
	return ptr
}

//QNetworkConfiguration::BearerType
type QNetworkConfiguration__BearerType int64

const (
	QNetworkConfiguration__BearerUnknown   = QNetworkConfiguration__BearerType(0)
	QNetworkConfiguration__BearerEthernet  = QNetworkConfiguration__BearerType(1)
	QNetworkConfiguration__BearerWLAN      = QNetworkConfiguration__BearerType(2)
	QNetworkConfiguration__Bearer2G        = QNetworkConfiguration__BearerType(3)
	QNetworkConfiguration__BearerCDMA2000  = QNetworkConfiguration__BearerType(4)
	QNetworkConfiguration__BearerWCDMA     = QNetworkConfiguration__BearerType(5)
	QNetworkConfiguration__BearerHSPA      = QNetworkConfiguration__BearerType(6)
	QNetworkConfiguration__BearerBluetooth = QNetworkConfiguration__BearerType(7)
	QNetworkConfiguration__BearerWiMAX     = QNetworkConfiguration__BearerType(8)
	QNetworkConfiguration__BearerEVDO      = QNetworkConfiguration__BearerType(9)
	QNetworkConfiguration__BearerLTE       = QNetworkConfiguration__BearerType(10)
	QNetworkConfiguration__Bearer3G        = QNetworkConfiguration__BearerType(11)
	QNetworkConfiguration__Bearer4G        = QNetworkConfiguration__BearerType(12)
)

//QNetworkConfiguration::Purpose
type QNetworkConfiguration__Purpose int64

const (
	QNetworkConfiguration__UnknownPurpose         = QNetworkConfiguration__Purpose(0)
	QNetworkConfiguration__PublicPurpose          = QNetworkConfiguration__Purpose(1)
	QNetworkConfiguration__PrivatePurpose         = QNetworkConfiguration__Purpose(2)
	QNetworkConfiguration__ServiceSpecificPurpose = QNetworkConfiguration__Purpose(3)
)

//QNetworkConfiguration::StateFlag
type QNetworkConfiguration__StateFlag int64

const (
	QNetworkConfiguration__Undefined  = QNetworkConfiguration__StateFlag(0x0000001)
	QNetworkConfiguration__Defined    = QNetworkConfiguration__StateFlag(0x0000002)
	QNetworkConfiguration__Discovered = QNetworkConfiguration__StateFlag(0x0000006)
	QNetworkConfiguration__Active     = QNetworkConfiguration__StateFlag(0x000000e)
)

//QNetworkConfiguration::Type
type QNetworkConfiguration__Type int64

const (
	QNetworkConfiguration__InternetAccessPoint = QNetworkConfiguration__Type(0)
	QNetworkConfiguration__ServiceNetwork      = QNetworkConfiguration__Type(1)
	QNetworkConfiguration__UserChoice          = QNetworkConfiguration__Type(2)
	QNetworkConfiguration__Invalid             = QNetworkConfiguration__Type(3)
)

func NewQNetworkConfiguration() *QNetworkConfiguration {
	defer qt.Recovering("QNetworkConfiguration::QNetworkConfiguration")

	return newQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration())
}

func NewQNetworkConfiguration2(other QNetworkConfiguration_ITF) *QNetworkConfiguration {
	defer qt.Recovering("QNetworkConfiguration::QNetworkConfiguration")

	return newQNetworkConfigurationFromPointer(C.QNetworkConfiguration_NewQNetworkConfiguration2(PointerFromQNetworkConfiguration(other)))
}

func (ptr *QNetworkConfiguration) BearerType() QNetworkConfiguration__BearerType {
	defer qt.Recovering("QNetworkConfiguration::bearerType")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeFamily() QNetworkConfiguration__BearerType {
	defer qt.Recovering("QNetworkConfiguration::bearerTypeFamily")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__BearerType(C.QNetworkConfiguration_BearerTypeFamily(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) BearerTypeName() string {
	defer qt.Recovering("QNetworkConfiguration::bearerTypeName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_BearerTypeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Identifier() string {
	defer qt.Recovering("QNetworkConfiguration::identifier")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_Identifier(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) IsRoamingAvailable() bool {
	defer qt.Recovering("QNetworkConfiguration::isRoamingAvailable")

	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsRoamingAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) IsValid() bool {
	defer qt.Recovering("QNetworkConfiguration::isValid")

	if ptr.Pointer() != nil {
		return C.QNetworkConfiguration_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfiguration) Name() string {
	defer qt.Recovering("QNetworkConfiguration::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkConfiguration_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkConfiguration) Purpose() QNetworkConfiguration__Purpose {
	defer qt.Recovering("QNetworkConfiguration::purpose")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Purpose(C.QNetworkConfiguration_Purpose(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) State() QNetworkConfiguration__StateFlag {
	defer qt.Recovering("QNetworkConfiguration::state")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__StateFlag(C.QNetworkConfiguration_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) Swap(other QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkConfiguration::swap")

	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_Swap(ptr.Pointer(), PointerFromQNetworkConfiguration(other))
	}
}

func (ptr *QNetworkConfiguration) Type() QNetworkConfiguration__Type {
	defer qt.Recovering("QNetworkConfiguration::type")

	if ptr.Pointer() != nil {
		return QNetworkConfiguration__Type(C.QNetworkConfiguration_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfiguration) DestroyQNetworkConfiguration() {
	defer qt.Recovering("QNetworkConfiguration::~QNetworkConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkConfiguration_DestroyQNetworkConfiguration(ptr.Pointer())
	}
}

type QNetworkConfigurationManager struct {
	core.QObject
}

type QNetworkConfigurationManager_ITF interface {
	core.QObject_ITF
	QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager
}

func PointerFromQNetworkConfigurationManager(ptr QNetworkConfigurationManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfigurationManager_PTR().Pointer()
	}
	return nil
}

func NewQNetworkConfigurationManagerFromPointer(ptr unsafe.Pointer) *QNetworkConfigurationManager {
	var n = new(QNetworkConfigurationManager)
	n.SetPointer(ptr)
	return n
}

func newQNetworkConfigurationManagerFromPointer(ptr unsafe.Pointer) *QNetworkConfigurationManager {
	var n = NewQNetworkConfigurationManagerFromPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkConfigurationManager_") {
		n.SetObjectName("QNetworkConfigurationManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkConfigurationManager) QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager {
	return ptr
}

//QNetworkConfigurationManager::Capability
type QNetworkConfigurationManager__Capability int64

const (
	QNetworkConfigurationManager__CanStartAndStopInterfaces = QNetworkConfigurationManager__Capability(0x00000001)
	QNetworkConfigurationManager__DirectConnectionRouting   = QNetworkConfigurationManager__Capability(0x00000002)
	QNetworkConfigurationManager__SystemSessionSupport      = QNetworkConfigurationManager__Capability(0x00000004)
	QNetworkConfigurationManager__ApplicationLevelRoaming   = QNetworkConfigurationManager__Capability(0x00000008)
	QNetworkConfigurationManager__ForcedRoaming             = QNetworkConfigurationManager__Capability(0x00000010)
	QNetworkConfigurationManager__DataStatistics            = QNetworkConfigurationManager__Capability(0x00000020)
	QNetworkConfigurationManager__NetworkSessionRequired    = QNetworkConfigurationManager__Capability(0x00000040)
)

func NewQNetworkConfigurationManager(parent core.QObject_ITF) *QNetworkConfigurationManager {
	defer qt.Recovering("QNetworkConfigurationManager::QNetworkConfigurationManager")

	return newQNetworkConfigurationManagerFromPointer(C.QNetworkConfigurationManager_NewQNetworkConfigurationManager(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkConfigurationManager) Capabilities() QNetworkConfigurationManager__Capability {
	defer qt.Recovering("QNetworkConfigurationManager::capabilities")

	if ptr.Pointer() != nil {
		return QNetworkConfigurationManager__Capability(C.QNetworkConfigurationManager_Capabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfigurationManager) IsOnline() bool {
	defer qt.Recovering("QNetworkConfigurationManager::isOnline")

	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_IsOnline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfigurationManager) ConnectOnlineStateChanged(f func(isOnline bool)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::onlineStateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectOnlineStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "onlineStateChanged", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectOnlineStateChanged() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::onlineStateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectOnlineStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "onlineStateChanged")
	}
}

//export callbackQNetworkConfigurationManagerOnlineStateChanged
func callbackQNetworkConfigurationManagerOnlineStateChanged(ptr unsafe.Pointer, ptrName *C.char, isOnline C.int) {
	defer qt.Recovering("callback QNetworkConfigurationManager::onlineStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "onlineStateChanged"); signal != nil {
		signal.(func(bool))(int(isOnline) != 0)
	}

}

func (ptr *QNetworkConfigurationManager) OnlineStateChanged(isOnline bool) {
	defer qt.Recovering("QNetworkConfigurationManager::onlineStateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_OnlineStateChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(isOnline)))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectUpdateCompleted(f func()) {
	defer qt.Recovering("connect QNetworkConfigurationManager::updateCompleted")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectUpdateCompleted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updateCompleted", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateCompleted() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::updateCompleted")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectUpdateCompleted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updateCompleted")
	}
}

//export callbackQNetworkConfigurationManagerUpdateCompleted
func callbackQNetworkConfigurationManagerUpdateCompleted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkConfigurationManager::updateCompleted")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCompleted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkConfigurationManager) UpdateCompleted() {
	defer qt.Recovering("QNetworkConfigurationManager::updateCompleted")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateCompleted(ptr.Pointer())
	}
}

func (ptr *QNetworkConfigurationManager) UpdateConfigurations() {
	defer qt.Recovering("QNetworkConfigurationManager::updateConfigurations")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateConfigurations(ptr.Pointer())
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManager() {
	defer qt.Recovering("QNetworkConfigurationManager::~QNetworkConfigurationManager")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfigurationManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkConfigurationManagerTimerEvent
func callbackQNetworkConfigurationManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkConfigurationManagerChildEvent
func callbackQNetworkConfigurationManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkConfigurationManagerCustomEvent
func callbackQNetworkConfigurationManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkConfigurationManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkConfigurationManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkConfigurationManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkConfigurationManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QNetworkCookie struct {
	ptr unsafe.Pointer
}

type QNetworkCookie_ITF interface {
	QNetworkCookie_PTR() *QNetworkCookie
}

func (p *QNetworkCookie) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkCookie) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkCookie(ptr QNetworkCookie_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookie_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCookieFromPointer(ptr unsafe.Pointer) *QNetworkCookie {
	var n = new(QNetworkCookie)
	n.SetPointer(ptr)
	return n
}

func newQNetworkCookieFromPointer(ptr unsafe.Pointer) *QNetworkCookie {
	var n = NewQNetworkCookieFromPointer(ptr)
	return n
}

func (ptr *QNetworkCookie) QNetworkCookie_PTR() *QNetworkCookie {
	return ptr
}

//QNetworkCookie::RawForm
type QNetworkCookie__RawForm int64

const (
	QNetworkCookie__NameAndValueOnly = QNetworkCookie__RawForm(0)
	QNetworkCookie__Full             = QNetworkCookie__RawForm(1)
)

func NewQNetworkCookie(name core.QByteArray_ITF, value core.QByteArray_ITF) *QNetworkCookie {
	defer qt.Recovering("QNetworkCookie::QNetworkCookie")

	return newQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie(core.PointerFromQByteArray(name), core.PointerFromQByteArray(value)))
}

func NewQNetworkCookie2(other QNetworkCookie_ITF) *QNetworkCookie {
	defer qt.Recovering("QNetworkCookie::QNetworkCookie")

	return newQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie2(PointerFromQNetworkCookie(other)))
}

func (ptr *QNetworkCookie) Domain() string {
	defer qt.Recovering("QNetworkCookie::domain")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkCookie_Domain(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkCookie) ExpirationDate() *core.QDateTime {
	defer qt.Recovering("QNetworkCookie::expirationDate")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QNetworkCookie_ExpirationDate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCookie) HasSameIdentifier(other QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookie::hasSameIdentifier")

	if ptr.Pointer() != nil {
		return C.QNetworkCookie_HasSameIdentifier(ptr.Pointer(), PointerFromQNetworkCookie(other)) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsHttpOnly() bool {
	defer qt.Recovering("QNetworkCookie::isHttpOnly")

	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsHttpOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSecure() bool {
	defer qt.Recovering("QNetworkCookie::isSecure")

	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSecure(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSessionCookie() bool {
	defer qt.Recovering("QNetworkCookie::isSessionCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSessionCookie(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) Name() *core.QByteArray {
	defer qt.Recovering("QNetworkCookie::name")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkCookie_Name(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCookie) Normalize(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkCookie::normalize")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_Normalize(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkCookie) Path() string {
	defer qt.Recovering("QNetworkCookie::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkCookie_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkCookie) SetDomain(domain string) {
	defer qt.Recovering("QNetworkCookie::setDomain")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetDomain(ptr.Pointer(), C.CString(domain))
	}
}

func (ptr *QNetworkCookie) SetExpirationDate(date core.QDateTime_ITF) {
	defer qt.Recovering("QNetworkCookie::setExpirationDate")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetExpirationDate(ptr.Pointer(), core.PointerFromQDateTime(date))
	}
}

func (ptr *QNetworkCookie) SetHttpOnly(enable bool) {
	defer qt.Recovering("QNetworkCookie::setHttpOnly")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetHttpOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QNetworkCookie) SetName(cookieName core.QByteArray_ITF) {
	defer qt.Recovering("QNetworkCookie::setName")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetName(ptr.Pointer(), core.PointerFromQByteArray(cookieName))
	}
}

func (ptr *QNetworkCookie) SetPath(path string) {
	defer qt.Recovering("QNetworkCookie::setPath")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QNetworkCookie) SetSecure(enable bool) {
	defer qt.Recovering("QNetworkCookie::setSecure")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetSecure(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QNetworkCookie) SetValue(value core.QByteArray_ITF) {
	defer qt.Recovering("QNetworkCookie::setValue")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetValue(ptr.Pointer(), core.PointerFromQByteArray(value))
	}
}

func (ptr *QNetworkCookie) Swap(other QNetworkCookie_ITF) {
	defer qt.Recovering("QNetworkCookie::swap")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_Swap(ptr.Pointer(), PointerFromQNetworkCookie(other))
	}
}

func (ptr *QNetworkCookie) ToRawForm(form QNetworkCookie__RawForm) *core.QByteArray {
	defer qt.Recovering("QNetworkCookie::toRawForm")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkCookie_ToRawForm(ptr.Pointer(), C.int(form)))
	}
	return nil
}

func (ptr *QNetworkCookie) Value() *core.QByteArray {
	defer qt.Recovering("QNetworkCookie::value")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkCookie_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCookie) DestroyQNetworkCookie() {
	defer qt.Recovering("QNetworkCookie::~QNetworkCookie")

	if ptr.Pointer() != nil {
		C.QNetworkCookie_DestroyQNetworkCookie(ptr.Pointer())
	}
}

type QNetworkCookieJar struct {
	core.QObject
}

type QNetworkCookieJar_ITF interface {
	core.QObject_ITF
	QNetworkCookieJar_PTR() *QNetworkCookieJar
}

func PointerFromQNetworkCookieJar(ptr QNetworkCookieJar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookieJar_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCookieJarFromPointer(ptr unsafe.Pointer) *QNetworkCookieJar {
	var n = new(QNetworkCookieJar)
	n.SetPointer(ptr)
	return n
}

func newQNetworkCookieJarFromPointer(ptr unsafe.Pointer) *QNetworkCookieJar {
	var n = NewQNetworkCookieJarFromPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkCookieJar_") {
		n.SetObjectName("QNetworkCookieJar_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkCookieJar) QNetworkCookieJar_PTR() *QNetworkCookieJar {
	return ptr
}

func NewQNetworkCookieJar(parent core.QObject_ITF) *QNetworkCookieJar {
	defer qt.Recovering("QNetworkCookieJar::QNetworkCookieJar")

	return newQNetworkCookieJarFromPointer(C.QNetworkCookieJar_NewQNetworkCookieJar(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkCookieJar) DeleteCookie(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::deleteCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_DeleteCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) InsertCookie(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::insertCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_InsertCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) UpdateCookie(cookie QNetworkCookie_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::updateCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_UpdateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) ValidateCookie(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {
	defer qt.Recovering("QNetworkCookieJar::validateCookie")

	if ptr.Pointer() != nil {
		return C.QNetworkCookieJar_ValidateCookie(ptr.Pointer(), PointerFromQNetworkCookie(cookie), core.PointerFromQUrl(url)) != 0
	}
	return false
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJar() {
	defer qt.Recovering("QNetworkCookieJar::~QNetworkCookieJar")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_DestroyQNetworkCookieJar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkCookieJar) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkCookieJarTimerEvent
func callbackQNetworkCookieJarTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkCookieJar) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkCookieJarChildEvent
func callbackQNetworkCookieJarChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkCookieJar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkCookieJar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkCookieJar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkCookieJarCustomEvent
func callbackQNetworkCookieJarCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkCookieJar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkCookieJarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkCookieJar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkCookieJar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkCookieJar::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkCookieJar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QNetworkDiskCache struct {
	QAbstractNetworkCache
}

type QNetworkDiskCache_ITF interface {
	QAbstractNetworkCache_ITF
	QNetworkDiskCache_PTR() *QNetworkDiskCache
}

func PointerFromQNetworkDiskCache(ptr QNetworkDiskCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkDiskCache_PTR().Pointer()
	}
	return nil
}

func NewQNetworkDiskCacheFromPointer(ptr unsafe.Pointer) *QNetworkDiskCache {
	var n = new(QNetworkDiskCache)
	n.SetPointer(ptr)
	return n
}

func newQNetworkDiskCacheFromPointer(ptr unsafe.Pointer) *QNetworkDiskCache {
	var n = NewQNetworkDiskCacheFromPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkDiskCache_") {
		n.SetObjectName("QNetworkDiskCache_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkDiskCache) QNetworkDiskCache_PTR() *QNetworkDiskCache {
	return ptr
}

func NewQNetworkDiskCache(parent core.QObject_ITF) *QNetworkDiskCache {
	defer qt.Recovering("QNetworkDiskCache::QNetworkDiskCache")

	return newQNetworkDiskCacheFromPointer(C.QNetworkDiskCache_NewQNetworkDiskCache(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkDiskCache) CacheDirectory() string {
	defer qt.Recovering("QNetworkDiskCache::cacheDirectory")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkDiskCache_CacheDirectory(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkDiskCache) CacheSize() int64 {
	defer qt.Recovering("QNetworkDiskCache::cacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_CacheSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) ConnectClear(f func()) {
	defer qt.Recovering("connect QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectClear() {
	defer qt.Recovering("disconnect QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

//export callbackQNetworkDiskCacheClear
func callbackQNetworkDiskCacheClear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkDiskCache::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkDiskCache) Clear() {
	defer qt.Recovering("QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_Clear(ptr.Pointer())
	}
}

func (ptr *QNetworkDiskCache) ClearDefault() {
	defer qt.Recovering("QNetworkDiskCache::clear")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ClearDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkDiskCache) Data(url core.QUrl_ITF) *core.QIODevice {
	defer qt.Recovering("QNetworkDiskCache::data")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QNetworkDiskCache) Expire() int64 {
	defer qt.Recovering("QNetworkDiskCache::expire")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_Expire(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) MaximumCacheSize() int64 {
	defer qt.Recovering("QNetworkDiskCache::maximumCacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkDiskCache_MaximumCacheSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	defer qt.Recovering("QNetworkDiskCache::prepare")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QNetworkDiskCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
	}
	return nil
}

func (ptr *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	defer qt.Recovering("QNetworkDiskCache::setCacheDirectory")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetCacheDirectory(ptr.Pointer(), C.CString(cacheDir))
	}
}

func (ptr *QNetworkDiskCache) SetMaximumCacheSize(size int64) {
	defer qt.Recovering("QNetworkDiskCache::setMaximumCacheSize")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_SetMaximumCacheSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCache() {
	defer qt.Recovering("QNetworkDiskCache::~QNetworkDiskCache")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_DestroyQNetworkDiskCache(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkDiskCache) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkDiskCacheTimerEvent
func callbackQNetworkDiskCacheTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkDiskCache) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkDiskCacheChildEvent
func callbackQNetworkDiskCacheChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkDiskCache) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkDiskCache) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkDiskCache) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkDiskCacheCustomEvent
func callbackQNetworkDiskCacheCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkDiskCache::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkDiskCacheFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkDiskCache) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkDiskCache) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkDiskCache::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkDiskCache_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QNetworkInterface struct {
	ptr unsafe.Pointer
}

type QNetworkInterface_ITF interface {
	QNetworkInterface_PTR() *QNetworkInterface
}

func (p *QNetworkInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkInterface(ptr QNetworkInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkInterface_PTR().Pointer()
	}
	return nil
}

func NewQNetworkInterfaceFromPointer(ptr unsafe.Pointer) *QNetworkInterface {
	var n = new(QNetworkInterface)
	n.SetPointer(ptr)
	return n
}

func newQNetworkInterfaceFromPointer(ptr unsafe.Pointer) *QNetworkInterface {
	var n = NewQNetworkInterfaceFromPointer(ptr)
	return n
}

func (ptr *QNetworkInterface) QNetworkInterface_PTR() *QNetworkInterface {
	return ptr
}

//QNetworkInterface::InterfaceFlag
type QNetworkInterface__InterfaceFlag int64

const (
	QNetworkInterface__IsUp           = QNetworkInterface__InterfaceFlag(0x1)
	QNetworkInterface__IsRunning      = QNetworkInterface__InterfaceFlag(0x2)
	QNetworkInterface__CanBroadcast   = QNetworkInterface__InterfaceFlag(0x4)
	QNetworkInterface__IsLoopBack     = QNetworkInterface__InterfaceFlag(0x8)
	QNetworkInterface__IsPointToPoint = QNetworkInterface__InterfaceFlag(0x10)
	QNetworkInterface__CanMulticast   = QNetworkInterface__InterfaceFlag(0x20)
)

func NewQNetworkInterface() *QNetworkInterface {
	defer qt.Recovering("QNetworkInterface::QNetworkInterface")

	return newQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface())
}

func NewQNetworkInterface2(other QNetworkInterface_ITF) *QNetworkInterface {
	defer qt.Recovering("QNetworkInterface::QNetworkInterface")

	return newQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface2(PointerFromQNetworkInterface(other)))
}

func (ptr *QNetworkInterface) Flags() QNetworkInterface__InterfaceFlag {
	defer qt.Recovering("QNetworkInterface::flags")

	if ptr.Pointer() != nil {
		return QNetworkInterface__InterfaceFlag(C.QNetworkInterface_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkInterface) HardwareAddress() string {
	defer qt.Recovering("QNetworkInterface::hardwareAddress")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HardwareAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) HumanReadableName() string {
	defer qt.Recovering("QNetworkInterface::humanReadableName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HumanReadableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Index() int {
	defer qt.Recovering("QNetworkInterface::index")

	if ptr.Pointer() != nil {
		return int(C.QNetworkInterface_Index(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkInterface) IsValid() bool {
	defer qt.Recovering("QNetworkInterface::isValid")

	if ptr.Pointer() != nil {
		return C.QNetworkInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkInterface) Name() string {
	defer qt.Recovering("QNetworkInterface::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Swap(other QNetworkInterface_ITF) {
	defer qt.Recovering("QNetworkInterface::swap")

	if ptr.Pointer() != nil {
		C.QNetworkInterface_Swap(ptr.Pointer(), PointerFromQNetworkInterface(other))
	}
}

func (ptr *QNetworkInterface) DestroyQNetworkInterface() {
	defer qt.Recovering("QNetworkInterface::~QNetworkInterface")

	if ptr.Pointer() != nil {
		C.QNetworkInterface_DestroyQNetworkInterface(ptr.Pointer())
	}
}

type QNetworkProxy struct {
	ptr unsafe.Pointer
}

type QNetworkProxy_ITF interface {
	QNetworkProxy_PTR() *QNetworkProxy
}

func (p *QNetworkProxy) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkProxy) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkProxy(ptr QNetworkProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxy_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyFromPointer(ptr unsafe.Pointer) *QNetworkProxy {
	var n = new(QNetworkProxy)
	n.SetPointer(ptr)
	return n
}

func newQNetworkProxyFromPointer(ptr unsafe.Pointer) *QNetworkProxy {
	var n = NewQNetworkProxyFromPointer(ptr)
	return n
}

func (ptr *QNetworkProxy) QNetworkProxy_PTR() *QNetworkProxy {
	return ptr
}

//QNetworkProxy::Capability
type QNetworkProxy__Capability int64

const (
	QNetworkProxy__TunnelingCapability      = QNetworkProxy__Capability(0x0001)
	QNetworkProxy__ListeningCapability      = QNetworkProxy__Capability(0x0002)
	QNetworkProxy__UdpTunnelingCapability   = QNetworkProxy__Capability(0x0004)
	QNetworkProxy__CachingCapability        = QNetworkProxy__Capability(0x0008)
	QNetworkProxy__HostNameLookupCapability = QNetworkProxy__Capability(0x0010)
)

//QNetworkProxy::ProxyType
type QNetworkProxy__ProxyType int64

const (
	QNetworkProxy__DefaultProxy     = QNetworkProxy__ProxyType(0)
	QNetworkProxy__Socks5Proxy      = QNetworkProxy__ProxyType(1)
	QNetworkProxy__NoProxy          = QNetworkProxy__ProxyType(2)
	QNetworkProxy__HttpProxy        = QNetworkProxy__ProxyType(3)
	QNetworkProxy__HttpCachingProxy = QNetworkProxy__ProxyType(4)
	QNetworkProxy__FtpCachingProxy  = QNetworkProxy__ProxyType(5)
)

func NewQNetworkProxy() *QNetworkProxy {
	defer qt.Recovering("QNetworkProxy::QNetworkProxy")

	return newQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy())
}

func NewQNetworkProxy3(other QNetworkProxy_ITF) *QNetworkProxy {
	defer qt.Recovering("QNetworkProxy::QNetworkProxy")

	return newQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy3(PointerFromQNetworkProxy(other)))
}

func (ptr *QNetworkProxy) Capabilities() QNetworkProxy__Capability {
	defer qt.Recovering("QNetworkProxy::capabilities")

	if ptr.Pointer() != nil {
		return QNetworkProxy__Capability(C.QNetworkProxy_Capabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) HasRawHeader(headerName core.QByteArray_ITF) bool {
	defer qt.Recovering("QNetworkProxy::hasRawHeader")

	if ptr.Pointer() != nil {
		return C.QNetworkProxy_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	defer qt.Recovering("QNetworkProxy::header")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QNetworkProxy_Header(ptr.Pointer(), C.int(header)))
	}
	return nil
}

func (ptr *QNetworkProxy) HostName() string {
	defer qt.Recovering("QNetworkProxy::hostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_HostName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) IsCachingProxy() bool {
	defer qt.Recovering("QNetworkProxy::isCachingProxy")

	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsCachingProxy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkProxy) IsTransparentProxy() bool {
	defer qt.Recovering("QNetworkProxy::isTransparentProxy")

	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsTransparentProxy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Password() string {
	defer qt.Recovering("QNetworkProxy::password")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QNetworkProxy::rawHeader")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkProxy_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
	}
	return nil
}

func QNetworkProxy_SetApplicationProxy(networkProxy QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkProxy::setApplicationProxy")

	C.QNetworkProxy_QNetworkProxy_SetApplicationProxy(PointerFromQNetworkProxy(networkProxy))
}

func (ptr *QNetworkProxy) SetCapabilities(capabilities QNetworkProxy__Capability) {
	defer qt.Recovering("QNetworkProxy::setCapabilities")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetCapabilities(ptr.Pointer(), C.int(capabilities))
	}
}

func (ptr *QNetworkProxy) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkProxy::setHeader")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetHeader(ptr.Pointer(), C.int(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkProxy) SetHostName(hostName string) {
	defer qt.Recovering("QNetworkProxy::setHostName")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetHostName(ptr.Pointer(), C.CString(hostName))
	}
}

func (ptr *QNetworkProxy) SetPassword(password string) {
	defer qt.Recovering("QNetworkProxy::setPassword")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetPassword(ptr.Pointer(), C.CString(password))
	}
}

func (ptr *QNetworkProxy) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {
	defer qt.Recovering("QNetworkProxy::setRawHeader")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(headerValue))
	}
}

func (ptr *QNetworkProxy) SetType(ty QNetworkProxy__ProxyType) {
	defer qt.Recovering("QNetworkProxy::setType")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QNetworkProxy) SetUser(user string) {
	defer qt.Recovering("QNetworkProxy::setUser")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetUser(ptr.Pointer(), C.CString(user))
	}
}

func (ptr *QNetworkProxy) Swap(other QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkProxy::swap")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_Swap(ptr.Pointer(), PointerFromQNetworkProxy(other))
	}
}

func (ptr *QNetworkProxy) Type() QNetworkProxy__ProxyType {
	defer qt.Recovering("QNetworkProxy::type")

	if ptr.Pointer() != nil {
		return QNetworkProxy__ProxyType(C.QNetworkProxy_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) User() string {
	defer qt.Recovering("QNetworkProxy::user")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_User(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) DestroyQNetworkProxy() {
	defer qt.Recovering("QNetworkProxy::~QNetworkProxy")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_DestroyQNetworkProxy(ptr.Pointer())
	}
}

type QNetworkProxyFactory struct {
	ptr unsafe.Pointer
}

type QNetworkProxyFactory_ITF interface {
	QNetworkProxyFactory_PTR() *QNetworkProxyFactory
}

func (p *QNetworkProxyFactory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkProxyFactory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkProxyFactory(ptr QNetworkProxyFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyFactory_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyFactoryFromPointer(ptr unsafe.Pointer) *QNetworkProxyFactory {
	var n = new(QNetworkProxyFactory)
	n.SetPointer(ptr)
	return n
}

func newQNetworkProxyFactoryFromPointer(ptr unsafe.Pointer) *QNetworkProxyFactory {
	var n = NewQNetworkProxyFactoryFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QNetworkProxyFactory_") {
		n.SetObjectNameAbs("QNetworkProxyFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkProxyFactory) QNetworkProxyFactory_PTR() *QNetworkProxyFactory {
	return ptr
}

func QNetworkProxyFactory_SetApplicationProxyFactory(factory QNetworkProxyFactory_ITF) {
	defer qt.Recovering("QNetworkProxyFactory::setApplicationProxyFactory")

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetApplicationProxyFactory(PointerFromQNetworkProxyFactory(factory))
}

func QNetworkProxyFactory_SetUseSystemConfiguration(enable bool) {
	defer qt.Recovering("QNetworkProxyFactory::setUseSystemConfiguration")

	C.QNetworkProxyFactory_QNetworkProxyFactory_SetUseSystemConfiguration(C.int(qt.GoBoolToInt(enable)))
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactory() {
	defer qt.Recovering("QNetworkProxyFactory::~QNetworkProxyFactory")

	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_DestroyQNetworkProxyFactory(ptr.Pointer())
	}
}

func (ptr *QNetworkProxyFactory) ObjectNameAbs() string {
	defer qt.Recovering("QNetworkProxyFactory::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyFactory_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyFactory) SetObjectNameAbs(name string) {
	defer qt.Recovering("QNetworkProxyFactory::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QNetworkProxyFactory_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QNetworkProxyQuery struct {
	ptr unsafe.Pointer
}

type QNetworkProxyQuery_ITF interface {
	QNetworkProxyQuery_PTR() *QNetworkProxyQuery
}

func (p *QNetworkProxyQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkProxyQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkProxyQuery(ptr QNetworkProxyQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyQuery_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyQueryFromPointer(ptr unsafe.Pointer) *QNetworkProxyQuery {
	var n = new(QNetworkProxyQuery)
	n.SetPointer(ptr)
	return n
}

func newQNetworkProxyQueryFromPointer(ptr unsafe.Pointer) *QNetworkProxyQuery {
	var n = NewQNetworkProxyQueryFromPointer(ptr)
	return n
}

func (ptr *QNetworkProxyQuery) QNetworkProxyQuery_PTR() *QNetworkProxyQuery {
	return ptr
}

//QNetworkProxyQuery::QueryType
type QNetworkProxyQuery__QueryType int64

const (
	QNetworkProxyQuery__TcpSocket  = QNetworkProxyQuery__QueryType(0)
	QNetworkProxyQuery__UdpSocket  = QNetworkProxyQuery__QueryType(1)
	QNetworkProxyQuery__TcpServer  = QNetworkProxyQuery__QueryType(100)
	QNetworkProxyQuery__UrlRequest = QNetworkProxyQuery__QueryType(101)
)

func NewQNetworkProxyQuery() *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return newQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery())
}

func NewQNetworkProxyQuery7(networkConfiguration QNetworkConfiguration_ITF, hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return newQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery7(PointerFromQNetworkConfiguration(networkConfiguration), C.CString(hostname), C.int(port), C.CString(protocolTag), C.int(queryType)))
}

func NewQNetworkProxyQuery6(networkConfiguration QNetworkConfiguration_ITF, requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return newQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery6(PointerFromQNetworkConfiguration(networkConfiguration), core.PointerFromQUrl(requestUrl), C.int(queryType)))
}

func NewQNetworkProxyQuery5(other QNetworkProxyQuery_ITF) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return newQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery5(PointerFromQNetworkProxyQuery(other)))
}

func NewQNetworkProxyQuery3(hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return newQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery3(C.CString(hostname), C.int(port), C.CString(protocolTag), C.int(queryType)))
}

func NewQNetworkProxyQuery2(requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {
	defer qt.Recovering("QNetworkProxyQuery::QNetworkProxyQuery")

	return newQNetworkProxyQueryFromPointer(C.QNetworkProxyQuery_NewQNetworkProxyQuery2(core.PointerFromQUrl(requestUrl), C.int(queryType)))
}

func (ptr *QNetworkProxyQuery) LocalPort() int {
	defer qt.Recovering("QNetworkProxyQuery::localPort")

	if ptr.Pointer() != nil {
		return int(C.QNetworkProxyQuery_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) PeerHostName() string {
	defer qt.Recovering("QNetworkProxyQuery::peerHostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyQuery_PeerHostName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) PeerPort() int {
	defer qt.Recovering("QNetworkProxyQuery::peerPort")

	if ptr.Pointer() != nil {
		return int(C.QNetworkProxyQuery_PeerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) ProtocolTag() string {
	defer qt.Recovering("QNetworkProxyQuery::protocolTag")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxyQuery_ProtocolTag(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxyQuery) QueryType() QNetworkProxyQuery__QueryType {
	defer qt.Recovering("QNetworkProxyQuery::queryType")

	if ptr.Pointer() != nil {
		return QNetworkProxyQuery__QueryType(C.QNetworkProxyQuery_QueryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxyQuery) SetLocalPort(port int) {
	defer qt.Recovering("QNetworkProxyQuery::setLocalPort")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetLocalPort(ptr.Pointer(), C.int(port))
	}
}

func (ptr *QNetworkProxyQuery) SetNetworkConfiguration(networkConfiguration QNetworkConfiguration_ITF) {
	defer qt.Recovering("QNetworkProxyQuery::setNetworkConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetNetworkConfiguration(ptr.Pointer(), PointerFromQNetworkConfiguration(networkConfiguration))
	}
}

func (ptr *QNetworkProxyQuery) SetPeerHostName(hostname string) {
	defer qt.Recovering("QNetworkProxyQuery::setPeerHostName")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetPeerHostName(ptr.Pointer(), C.CString(hostname))
	}
}

func (ptr *QNetworkProxyQuery) SetPeerPort(port int) {
	defer qt.Recovering("QNetworkProxyQuery::setPeerPort")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetPeerPort(ptr.Pointer(), C.int(port))
	}
}

func (ptr *QNetworkProxyQuery) SetProtocolTag(protocolTag string) {
	defer qt.Recovering("QNetworkProxyQuery::setProtocolTag")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetProtocolTag(ptr.Pointer(), C.CString(protocolTag))
	}
}

func (ptr *QNetworkProxyQuery) SetQueryType(ty QNetworkProxyQuery__QueryType) {
	defer qt.Recovering("QNetworkProxyQuery::setQueryType")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetQueryType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QNetworkProxyQuery) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkProxyQuery::setUrl")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkProxyQuery) Swap(other QNetworkProxyQuery_ITF) {
	defer qt.Recovering("QNetworkProxyQuery::swap")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_Swap(ptr.Pointer(), PointerFromQNetworkProxyQuery(other))
	}
}

func (ptr *QNetworkProxyQuery) Url() *core.QUrl {
	defer qt.Recovering("QNetworkProxyQuery::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QNetworkProxyQuery_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkProxyQuery) DestroyQNetworkProxyQuery() {
	defer qt.Recovering("QNetworkProxyQuery::~QNetworkProxyQuery")

	if ptr.Pointer() != nil {
		C.QNetworkProxyQuery_DestroyQNetworkProxyQuery(ptr.Pointer())
	}
}

type QNetworkReply struct {
	core.QIODevice
}

type QNetworkReply_ITF interface {
	core.QIODevice_ITF
	QNetworkReply_PTR() *QNetworkReply
}

func PointerFromQNetworkReply(ptr QNetworkReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkReply_PTR().Pointer()
	}
	return nil
}

func NewQNetworkReplyFromPointer(ptr unsafe.Pointer) *QNetworkReply {
	var n = new(QNetworkReply)
	n.SetPointer(ptr)
	return n
}

func newQNetworkReplyFromPointer(ptr unsafe.Pointer) *QNetworkReply {
	var n = NewQNetworkReplyFromPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkReply_") {
		n.SetObjectName("QNetworkReply_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkReply) QNetworkReply_PTR() *QNetworkReply {
	return ptr
}

//QNetworkReply::NetworkError
type QNetworkReply__NetworkError int64

const (
	QNetworkReply__NoError                           = QNetworkReply__NetworkError(0)
	QNetworkReply__ConnectionRefusedError            = QNetworkReply__NetworkError(1)
	QNetworkReply__RemoteHostClosedError             = QNetworkReply__NetworkError(2)
	QNetworkReply__HostNotFoundError                 = QNetworkReply__NetworkError(3)
	QNetworkReply__TimeoutError                      = QNetworkReply__NetworkError(4)
	QNetworkReply__OperationCanceledError            = QNetworkReply__NetworkError(5)
	QNetworkReply__SslHandshakeFailedError           = QNetworkReply__NetworkError(6)
	QNetworkReply__TemporaryNetworkFailureError      = QNetworkReply__NetworkError(7)
	QNetworkReply__NetworkSessionFailedError         = QNetworkReply__NetworkError(8)
	QNetworkReply__BackgroundRequestNotAllowedError  = QNetworkReply__NetworkError(9)
	QNetworkReply__UnknownNetworkError               = QNetworkReply__NetworkError(99)
	QNetworkReply__ProxyConnectionRefusedError       = QNetworkReply__NetworkError(101)
	QNetworkReply__ProxyConnectionClosedError        = QNetworkReply__NetworkError(102)
	QNetworkReply__ProxyNotFoundError                = QNetworkReply__NetworkError(103)
	QNetworkReply__ProxyTimeoutError                 = QNetworkReply__NetworkError(104)
	QNetworkReply__ProxyAuthenticationRequiredError  = QNetworkReply__NetworkError(105)
	QNetworkReply__UnknownProxyError                 = QNetworkReply__NetworkError(199)
	QNetworkReply__ContentAccessDenied               = QNetworkReply__NetworkError(201)
	QNetworkReply__ContentOperationNotPermittedError = QNetworkReply__NetworkError(202)
	QNetworkReply__ContentNotFoundError              = QNetworkReply__NetworkError(203)
	QNetworkReply__AuthenticationRequiredError       = QNetworkReply__NetworkError(204)
	QNetworkReply__ContentReSendError                = QNetworkReply__NetworkError(205)
	QNetworkReply__ContentConflictError              = QNetworkReply__NetworkError(206)
	QNetworkReply__ContentGoneError                  = QNetworkReply__NetworkError(207)
	QNetworkReply__UnknownContentError               = QNetworkReply__NetworkError(299)
	QNetworkReply__ProtocolUnknownError              = QNetworkReply__NetworkError(301)
	QNetworkReply__ProtocolInvalidOperationError     = QNetworkReply__NetworkError(302)
	QNetworkReply__ProtocolFailure                   = QNetworkReply__NetworkError(399)
	QNetworkReply__InternalServerError               = QNetworkReply__NetworkError(401)
	QNetworkReply__OperationNotImplementedError      = QNetworkReply__NetworkError(402)
	QNetworkReply__ServiceUnavailableError           = QNetworkReply__NetworkError(403)
	QNetworkReply__UnknownServerError                = QNetworkReply__NetworkError(499)
)

func (ptr *QNetworkReply) Abort() {
	defer qt.Recovering("QNetworkReply::abort")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Abort(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) Attribute(code QNetworkRequest__Attribute) *core.QVariant {
	defer qt.Recovering("QNetworkReply::attribute")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QNetworkReply_Attribute(ptr.Pointer(), C.int(code)))
	}
	return nil
}

func (ptr *QNetworkReply) ConnectClose(f func()) {
	defer qt.Recovering("connect QNetworkReply::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QNetworkReply) DisconnectClose() {
	defer qt.Recovering("disconnect QNetworkReply::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQNetworkReplyClose
func callbackQNetworkReplyClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkReply::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQNetworkReplyFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QNetworkReply) Close() {
	defer qt.Recovering("QNetworkReply::close")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Close(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) CloseDefault() {
	defer qt.Recovering("QNetworkReply::close")

	if ptr.Pointer() != nil {
		C.QNetworkReply_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) ConnectDownloadProgress(f func(bytesReceived int64, bytesTotal int64)) {
	defer qt.Recovering("connect QNetworkReply::downloadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectDownloadProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "downloadProgress", f)
	}
}

func (ptr *QNetworkReply) DisconnectDownloadProgress() {
	defer qt.Recovering("disconnect QNetworkReply::downloadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectDownloadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "downloadProgress")
	}
}

//export callbackQNetworkReplyDownloadProgress
func callbackQNetworkReplyDownloadProgress(ptr unsafe.Pointer, ptrName *C.char, bytesReceived C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QNetworkReply::downloadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "downloadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesReceived), int64(bytesTotal))
	}

}

func (ptr *QNetworkReply) DownloadProgress(bytesReceived int64, bytesTotal int64) {
	defer qt.Recovering("QNetworkReply::downloadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DownloadProgress(ptr.Pointer(), C.longlong(bytesReceived), C.longlong(bytesTotal))
	}
}

func (ptr *QNetworkReply) ConnectEncrypted(f func()) {
	defer qt.Recovering("connect QNetworkReply::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QNetworkReply) DisconnectEncrypted() {
	defer qt.Recovering("disconnect QNetworkReply::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQNetworkReplyEncrypted
func callbackQNetworkReplyEncrypted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkReply::encrypted")

	if signal := qt.GetSignal(C.GoString(ptrName), "encrypted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) Encrypted() {
	defer qt.Recovering("QNetworkReply::encrypted")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Encrypted(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) ConnectError2(f func(code QNetworkReply__NetworkError)) {
	defer qt.Recovering("connect QNetworkReply::error")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QNetworkReply) DisconnectError2() {
	defer qt.Recovering("disconnect QNetworkReply::error")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQNetworkReplyError2
func callbackQNetworkReplyError2(ptr unsafe.Pointer, ptrName *C.char, code C.int) {
	defer qt.Recovering("callback QNetworkReply::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QNetworkReply__NetworkError))(QNetworkReply__NetworkError(code))
	}

}

func (ptr *QNetworkReply) Error2(code QNetworkReply__NetworkError) {
	defer qt.Recovering("QNetworkReply::error")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Error2(ptr.Pointer(), C.int(code))
	}
}

func (ptr *QNetworkReply) Error() QNetworkReply__NetworkError {
	defer qt.Recovering("QNetworkReply::error")

	if ptr.Pointer() != nil {
		return QNetworkReply__NetworkError(C.QNetworkReply_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) ConnectFinished(f func()) {
	defer qt.Recovering("connect QNetworkReply::finished")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QNetworkReply) DisconnectFinished() {
	defer qt.Recovering("disconnect QNetworkReply::finished")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQNetworkReplyFinished
func callbackQNetworkReplyFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkReply::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) Finished() {
	defer qt.Recovering("QNetworkReply::finished")

	if ptr.Pointer() != nil {
		C.QNetworkReply_Finished(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) HasRawHeader(headerName core.QByteArray_ITF) bool {
	defer qt.Recovering("QNetworkReply::hasRawHeader")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)) != 0
	}
	return false
}

func (ptr *QNetworkReply) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	defer qt.Recovering("QNetworkReply::header")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QNetworkReply_Header(ptr.Pointer(), C.int(header)))
	}
	return nil
}

func (ptr *QNetworkReply) ConnectIgnoreSslErrors(f func()) {
	defer qt.Recovering("connect QNetworkReply::ignoreSslErrors")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "ignoreSslErrors", f)
	}
}

func (ptr *QNetworkReply) DisconnectIgnoreSslErrors() {
	defer qt.Recovering("disconnect QNetworkReply::ignoreSslErrors")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "ignoreSslErrors")
	}
}

//export callbackQNetworkReplyIgnoreSslErrors
func callbackQNetworkReplyIgnoreSslErrors(ptr unsafe.Pointer, ptrName *C.char) bool {
	defer qt.Recovering("callback QNetworkReply::ignoreSslErrors")

	if signal := qt.GetSignal(C.GoString(ptrName), "ignoreSslErrors"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QNetworkReply) IgnoreSslErrors() {
	defer qt.Recovering("QNetworkReply::ignoreSslErrors")

	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) IgnoreSslErrorsDefault() {
	defer qt.Recovering("QNetworkReply::ignoreSslErrors")

	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrorsDefault(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) IsFinished() bool {
	defer qt.Recovering("QNetworkReply::isFinished")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) IsRunning() bool {
	defer qt.Recovering("QNetworkReply::isRunning")

	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkReply) Manager() *QNetworkAccessManager {
	defer qt.Recovering("QNetworkReply::manager")

	if ptr.Pointer() != nil {
		return NewQNetworkAccessManagerFromPointer(C.QNetworkReply_Manager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkReply) ConnectMetaDataChanged(f func()) {
	defer qt.Recovering("connect QNetworkReply::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectMetaDataChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QNetworkReply) DisconnectMetaDataChanged() {
	defer qt.Recovering("disconnect QNetworkReply::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectMetaDataChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQNetworkReplyMetaDataChanged
func callbackQNetworkReplyMetaDataChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkReply::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkReply) MetaDataChanged() {
	defer qt.Recovering("QNetworkReply::metaDataChanged")

	if ptr.Pointer() != nil {
		C.QNetworkReply_MetaDataChanged(ptr.Pointer())
	}
}

func (ptr *QNetworkReply) Operation() QNetworkAccessManager__Operation {
	defer qt.Recovering("QNetworkReply::operation")

	if ptr.Pointer() != nil {
		return QNetworkAccessManager__Operation(C.QNetworkReply_Operation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	defer qt.Recovering("connect QNetworkReply::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkReply) DisconnectPreSharedKeyAuthenticationRequired() {
	defer qt.Recovering("disconnect QNetworkReply::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQNetworkReplyPreSharedKeyAuthenticationRequired
func callbackQNetworkReplyPreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::preSharedKeyAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QNetworkReply) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	defer qt.Recovering("QNetworkReply::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QNetworkReply_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QNetworkReply) RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QNetworkReply::rawHeader")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkReply_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
	}
	return nil
}

func (ptr *QNetworkReply) ReadBufferSize() int64 {
	defer qt.Recovering("QNetworkReply::readBufferSize")

	if ptr.Pointer() != nil {
		return int64(C.QNetworkReply_ReadBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkReply) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QNetworkReply::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setReadBufferSize", f)
	}
}

func (ptr *QNetworkReply) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QNetworkReply::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setReadBufferSize")
	}
}

//export callbackQNetworkReplySetReadBufferSize
func callbackQNetworkReplySetReadBufferSize(ptr unsafe.Pointer, ptrName *C.char, size C.longlong) {
	defer qt.Recovering("callback QNetworkReply::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQNetworkReplyFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QNetworkReply) SetReadBufferSize(size int64) {
	defer qt.Recovering("QNetworkReply::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QNetworkReply) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QNetworkReply::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QNetworkReply) SetSslConfiguration(config QSslConfiguration_ITF) {
	defer qt.Recovering("QNetworkReply::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkReply_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(config))
	}
}

func (ptr *QNetworkReply) ConnectUploadProgress(f func(bytesSent int64, bytesTotal int64)) {
	defer qt.Recovering("connect QNetworkReply::uploadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectUploadProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "uploadProgress", f)
	}
}

func (ptr *QNetworkReply) DisconnectUploadProgress() {
	defer qt.Recovering("disconnect QNetworkReply::uploadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectUploadProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "uploadProgress")
	}
}

//export callbackQNetworkReplyUploadProgress
func callbackQNetworkReplyUploadProgress(ptr unsafe.Pointer, ptrName *C.char, bytesSent C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QNetworkReply::uploadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "uploadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesSent), int64(bytesTotal))
	}

}

func (ptr *QNetworkReply) UploadProgress(bytesSent int64, bytesTotal int64) {
	defer qt.Recovering("QNetworkReply::uploadProgress")

	if ptr.Pointer() != nil {
		C.QNetworkReply_UploadProgress(ptr.Pointer(), C.longlong(bytesSent), C.longlong(bytesTotal))
	}
}

func (ptr *QNetworkReply) Url() *core.QUrl {
	defer qt.Recovering("QNetworkReply::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QNetworkReply_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkReply) DestroyQNetworkReply() {
	defer qt.Recovering("QNetworkReply::~QNetworkReply")

	if ptr.Pointer() != nil {
		C.QNetworkReply_DestroyQNetworkReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkReply) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkReplyTimerEvent
func callbackQNetworkReplyTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkReply::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkReply) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkReply::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkReplyChildEvent
func callbackQNetworkReplyChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkReply::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkReply) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkReply::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkReply::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkReply) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkReply::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkReplyCustomEvent
func callbackQNetworkReplyCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkReply) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkReply::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkReply) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkReply::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QNetworkRequest struct {
	ptr unsafe.Pointer
}

type QNetworkRequest_ITF interface {
	QNetworkRequest_PTR() *QNetworkRequest
}

func (p *QNetworkRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkRequest(ptr QNetworkRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkRequest_PTR().Pointer()
	}
	return nil
}

func NewQNetworkRequestFromPointer(ptr unsafe.Pointer) *QNetworkRequest {
	var n = new(QNetworkRequest)
	n.SetPointer(ptr)
	return n
}

func newQNetworkRequestFromPointer(ptr unsafe.Pointer) *QNetworkRequest {
	var n = NewQNetworkRequestFromPointer(ptr)
	return n
}

func (ptr *QNetworkRequest) QNetworkRequest_PTR() *QNetworkRequest {
	return ptr
}

//QNetworkRequest::Attribute
type QNetworkRequest__Attribute int64

const (
	QNetworkRequest__HttpStatusCodeAttribute               = QNetworkRequest__Attribute(0)
	QNetworkRequest__HttpReasonPhraseAttribute             = QNetworkRequest__Attribute(1)
	QNetworkRequest__RedirectionTargetAttribute            = QNetworkRequest__Attribute(2)
	QNetworkRequest__ConnectionEncryptedAttribute          = QNetworkRequest__Attribute(3)
	QNetworkRequest__CacheLoadControlAttribute             = QNetworkRequest__Attribute(4)
	QNetworkRequest__CacheSaveControlAttribute             = QNetworkRequest__Attribute(5)
	QNetworkRequest__SourceIsFromCacheAttribute            = QNetworkRequest__Attribute(6)
	QNetworkRequest__DoNotBufferUploadDataAttribute        = QNetworkRequest__Attribute(7)
	QNetworkRequest__HttpPipeliningAllowedAttribute        = QNetworkRequest__Attribute(8)
	QNetworkRequest__HttpPipeliningWasUsedAttribute        = QNetworkRequest__Attribute(9)
	QNetworkRequest__CustomVerbAttribute                   = QNetworkRequest__Attribute(10)
	QNetworkRequest__CookieLoadControlAttribute            = QNetworkRequest__Attribute(11)
	QNetworkRequest__AuthenticationReuseAttribute          = QNetworkRequest__Attribute(12)
	QNetworkRequest__CookieSaveControlAttribute            = QNetworkRequest__Attribute(13)
	QNetworkRequest__MaximumDownloadBufferSizeAttribute    = QNetworkRequest__Attribute(14)
	QNetworkRequest__DownloadBufferAttribute               = QNetworkRequest__Attribute(15)
	QNetworkRequest__SynchronousRequestAttribute           = QNetworkRequest__Attribute(16)
	QNetworkRequest__BackgroundRequestAttribute            = QNetworkRequest__Attribute(17)
	QNetworkRequest__SpdyAllowedAttribute                  = QNetworkRequest__Attribute(18)
	QNetworkRequest__SpdyWasUsedAttribute                  = QNetworkRequest__Attribute(19)
	QNetworkRequest__EmitAllUploadProgressSignalsAttribute = QNetworkRequest__Attribute(20)
	QNetworkRequest__User                                  = QNetworkRequest__Attribute(1000)
	QNetworkRequest__UserMax                               = QNetworkRequest__Attribute(32767)
)

//QNetworkRequest::CacheLoadControl
type QNetworkRequest__CacheLoadControl int64

const (
	QNetworkRequest__AlwaysNetwork = QNetworkRequest__CacheLoadControl(0)
	QNetworkRequest__PreferNetwork = QNetworkRequest__CacheLoadControl(1)
	QNetworkRequest__PreferCache   = QNetworkRequest__CacheLoadControl(2)
	QNetworkRequest__AlwaysCache   = QNetworkRequest__CacheLoadControl(3)
)

//QNetworkRequest::KnownHeaders
type QNetworkRequest__KnownHeaders int64

const (
	QNetworkRequest__ContentTypeHeader        = QNetworkRequest__KnownHeaders(0)
	QNetworkRequest__ContentLengthHeader      = QNetworkRequest__KnownHeaders(1)
	QNetworkRequest__LocationHeader           = QNetworkRequest__KnownHeaders(2)
	QNetworkRequest__LastModifiedHeader       = QNetworkRequest__KnownHeaders(3)
	QNetworkRequest__CookieHeader             = QNetworkRequest__KnownHeaders(4)
	QNetworkRequest__SetCookieHeader          = QNetworkRequest__KnownHeaders(5)
	QNetworkRequest__ContentDispositionHeader = QNetworkRequest__KnownHeaders(6)
	QNetworkRequest__UserAgentHeader          = QNetworkRequest__KnownHeaders(7)
	QNetworkRequest__ServerHeader             = QNetworkRequest__KnownHeaders(8)
)

//QNetworkRequest::LoadControl
type QNetworkRequest__LoadControl int64

const (
	QNetworkRequest__Automatic = QNetworkRequest__LoadControl(0)
	QNetworkRequest__Manual    = QNetworkRequest__LoadControl(1)
)

//QNetworkRequest::Priority
type QNetworkRequest__Priority int64

const (
	QNetworkRequest__HighPriority   = QNetworkRequest__Priority(1)
	QNetworkRequest__NormalPriority = QNetworkRequest__Priority(3)
	QNetworkRequest__LowPriority    = QNetworkRequest__Priority(5)
)

func NewQNetworkRequest2(other QNetworkRequest_ITF) *QNetworkRequest {
	defer qt.Recovering("QNetworkRequest::QNetworkRequest")

	return newQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest2(PointerFromQNetworkRequest(other)))
}

func NewQNetworkRequest(url core.QUrl_ITF) *QNetworkRequest {
	defer qt.Recovering("QNetworkRequest::QNetworkRequest")

	return newQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest(core.PointerFromQUrl(url)))
}

func (ptr *QNetworkRequest) Attribute(code QNetworkRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QNetworkRequest::attribute")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QNetworkRequest_Attribute(ptr.Pointer(), C.int(code), core.PointerFromQVariant(defaultValue)))
	}
	return nil
}

func (ptr *QNetworkRequest) HasRawHeader(headerName core.QByteArray_ITF) bool {
	defer qt.Recovering("QNetworkRequest::hasRawHeader")

	if ptr.Pointer() != nil {
		return C.QNetworkRequest_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)) != 0
	}
	return false
}

func (ptr *QNetworkRequest) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	defer qt.Recovering("QNetworkRequest::header")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QNetworkRequest_Header(ptr.Pointer(), C.int(header)))
	}
	return nil
}

func (ptr *QNetworkRequest) OriginatingObject() *core.QObject {
	defer qt.Recovering("QNetworkRequest::originatingObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QNetworkRequest_OriginatingObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkRequest) Priority() QNetworkRequest__Priority {
	defer qt.Recovering("QNetworkRequest::priority")

	if ptr.Pointer() != nil {
		return QNetworkRequest__Priority(C.QNetworkRequest_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkRequest) RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QNetworkRequest::rawHeader")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkRequest_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
	}
	return nil
}

func (ptr *QNetworkRequest) SetAttribute(code QNetworkRequest__Attribute, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkRequest::setAttribute")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetAttribute(ptr.Pointer(), C.int(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkRequest) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkRequest::setHeader")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetHeader(ptr.Pointer(), C.int(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkRequest) SetOriginatingObject(object core.QObject_ITF) {
	defer qt.Recovering("QNetworkRequest::setOriginatingObject")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetOriginatingObject(ptr.Pointer(), core.PointerFromQObject(object))
	}
}

func (ptr *QNetworkRequest) SetPriority(priority QNetworkRequest__Priority) {
	defer qt.Recovering("QNetworkRequest::setPriority")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetPriority(ptr.Pointer(), C.int(priority))
	}
}

func (ptr *QNetworkRequest) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {
	defer qt.Recovering("QNetworkRequest::setRawHeader")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(headerValue))
	}
}

func (ptr *QNetworkRequest) SetSslConfiguration(config QSslConfiguration_ITF) {
	defer qt.Recovering("QNetworkRequest::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(config))
	}
}

func (ptr *QNetworkRequest) SetUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QNetworkRequest::setUrl")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkRequest) Swap(other QNetworkRequest_ITF) {
	defer qt.Recovering("QNetworkRequest::swap")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_Swap(ptr.Pointer(), PointerFromQNetworkRequest(other))
	}
}

func (ptr *QNetworkRequest) Url() *core.QUrl {
	defer qt.Recovering("QNetworkRequest::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QNetworkRequest_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkRequest) DestroyQNetworkRequest() {
	defer qt.Recovering("QNetworkRequest::~QNetworkRequest")

	if ptr.Pointer() != nil {
		C.QNetworkRequest_DestroyQNetworkRequest(ptr.Pointer())
	}
}

type QNetworkSession struct {
	core.QObject
}

type QNetworkSession_ITF interface {
	core.QObject_ITF
	QNetworkSession_PTR() *QNetworkSession
}

func PointerFromQNetworkSession(ptr QNetworkSession_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkSession_PTR().Pointer()
	}
	return nil
}

func NewQNetworkSessionFromPointer(ptr unsafe.Pointer) *QNetworkSession {
	var n = new(QNetworkSession)
	n.SetPointer(ptr)
	return n
}

func newQNetworkSessionFromPointer(ptr unsafe.Pointer) *QNetworkSession {
	var n = NewQNetworkSessionFromPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkSession_") {
		n.SetObjectName("QNetworkSession_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkSession) QNetworkSession_PTR() *QNetworkSession {
	return ptr
}

//QNetworkSession::SessionError
type QNetworkSession__SessionError int64

const (
	QNetworkSession__UnknownSessionError        = QNetworkSession__SessionError(0)
	QNetworkSession__SessionAbortedError        = QNetworkSession__SessionError(1)
	QNetworkSession__RoamingError               = QNetworkSession__SessionError(2)
	QNetworkSession__OperationNotSupportedError = QNetworkSession__SessionError(3)
	QNetworkSession__InvalidConfigurationError  = QNetworkSession__SessionError(4)
)

//QNetworkSession::State
type QNetworkSession__State int64

const (
	QNetworkSession__Invalid      = QNetworkSession__State(0)
	QNetworkSession__NotAvailable = QNetworkSession__State(1)
	QNetworkSession__Connecting   = QNetworkSession__State(2)
	QNetworkSession__Connected    = QNetworkSession__State(3)
	QNetworkSession__Closing      = QNetworkSession__State(4)
	QNetworkSession__Disconnected = QNetworkSession__State(5)
	QNetworkSession__Roaming      = QNetworkSession__State(6)
)

//QNetworkSession::UsagePolicy
type QNetworkSession__UsagePolicy int64

const (
	QNetworkSession__NoPolicy                  = QNetworkSession__UsagePolicy(0)
	QNetworkSession__NoBackgroundTrafficPolicy = QNetworkSession__UsagePolicy(1)
)

func NewQNetworkSession(connectionConfig QNetworkConfiguration_ITF, parent core.QObject_ITF) *QNetworkSession {
	defer qt.Recovering("QNetworkSession::QNetworkSession")

	return newQNetworkSessionFromPointer(C.QNetworkSession_NewQNetworkSession(PointerFromQNetworkConfiguration(connectionConfig), core.PointerFromQObject(parent)))
}

func (ptr *QNetworkSession) Accept() {
	defer qt.Recovering("QNetworkSession::accept")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Accept(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Close() {
	defer qt.Recovering("QNetworkSession::close")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Close(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectClosed(f func()) {
	defer qt.Recovering("connect QNetworkSession::closed")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectClosed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "closed", f)
	}
}

func (ptr *QNetworkSession) DisconnectClosed() {
	defer qt.Recovering("disconnect QNetworkSession::closed")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "closed")
	}
}

//export callbackQNetworkSessionClosed
func callbackQNetworkSessionClosed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkSession::closed")

	if signal := qt.GetSignal(C.GoString(ptrName), "closed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) Closed() {
	defer qt.Recovering("QNetworkSession::closed")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Closed(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectError2(f func(error QNetworkSession__SessionError)) {
	defer qt.Recovering("connect QNetworkSession::error")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QNetworkSession) DisconnectError2() {
	defer qt.Recovering("disconnect QNetworkSession::error")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQNetworkSessionError2
func callbackQNetworkSessionError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QNetworkSession::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QNetworkSession__SessionError))(QNetworkSession__SessionError(error))
	}

}

func (ptr *QNetworkSession) Error2(error QNetworkSession__SessionError) {
	defer qt.Recovering("QNetworkSession::error")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QNetworkSession) Error() QNetworkSession__SessionError {
	defer qt.Recovering("QNetworkSession::error")

	if ptr.Pointer() != nil {
		return QNetworkSession__SessionError(C.QNetworkSession_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ErrorString() string {
	defer qt.Recovering("QNetworkSession::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkSession_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkSession) Ignore() {
	defer qt.Recovering("QNetworkSession::ignore")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Ignore(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) IsOpen() bool {
	defer qt.Recovering("QNetworkSession::isOpen")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkSession) Migrate() {
	defer qt.Recovering("QNetworkSession::migrate")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Migrate(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectNewConfigurationActivated(f func()) {
	defer qt.Recovering("connect QNetworkSession::newConfigurationActivated")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNewConfigurationActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConfigurationActivated", f)
	}
}

func (ptr *QNetworkSession) DisconnectNewConfigurationActivated() {
	defer qt.Recovering("disconnect QNetworkSession::newConfigurationActivated")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNewConfigurationActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConfigurationActivated")
	}
}

//export callbackQNetworkSessionNewConfigurationActivated
func callbackQNetworkSessionNewConfigurationActivated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkSession::newConfigurationActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConfigurationActivated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) NewConfigurationActivated() {
	defer qt.Recovering("QNetworkSession::newConfigurationActivated")

	if ptr.Pointer() != nil {
		C.QNetworkSession_NewConfigurationActivated(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Open() {
	defer qt.Recovering("QNetworkSession::open")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Open(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectOpened(f func()) {
	defer qt.Recovering("connect QNetworkSession::opened")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectOpened(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opened", f)
	}
}

func (ptr *QNetworkSession) DisconnectOpened() {
	defer qt.Recovering("disconnect QNetworkSession::opened")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectOpened(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opened")
	}
}

//export callbackQNetworkSessionOpened
func callbackQNetworkSessionOpened(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkSession::opened")

	if signal := qt.GetSignal(C.GoString(ptrName), "opened"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) Opened() {
	defer qt.Recovering("QNetworkSession::opened")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Opened(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Reject() {
	defer qt.Recovering("QNetworkSession::reject")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Reject(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) SessionProperty(key string) *core.QVariant {
	defer qt.Recovering("QNetworkSession::sessionProperty")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QNetworkSession_SessionProperty(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QNetworkSession) SetSessionProperty(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkSession::setSessionProperty")

	if ptr.Pointer() != nil {
		C.QNetworkSession_SetSessionProperty(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkSession) State() QNetworkSession__State {
	defer qt.Recovering("QNetworkSession::state")

	if ptr.Pointer() != nil {
		return QNetworkSession__State(C.QNetworkSession_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ConnectStateChanged(f func(state QNetworkSession__State)) {
	defer qt.Recovering("connect QNetworkSession::stateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QNetworkSession::stateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQNetworkSessionStateChanged
func callbackQNetworkSessionStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QNetworkSession::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QNetworkSession__State))(QNetworkSession__State(state))
	}

}

func (ptr *QNetworkSession) StateChanged(state QNetworkSession__State) {
	defer qt.Recovering("QNetworkSession::stateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QNetworkSession) Stop() {
	defer qt.Recovering("QNetworkSession::stop")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Stop(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) UsagePolicies() QNetworkSession__UsagePolicy {
	defer qt.Recovering("QNetworkSession::usagePolicies")

	if ptr.Pointer() != nil {
		return QNetworkSession__UsagePolicy(C.QNetworkSession_UsagePolicies(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ConnectUsagePoliciesChanged(f func(usagePolicies QNetworkSession__UsagePolicy)) {
	defer qt.Recovering("connect QNetworkSession::usagePoliciesChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectUsagePoliciesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "usagePoliciesChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectUsagePoliciesChanged() {
	defer qt.Recovering("disconnect QNetworkSession::usagePoliciesChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectUsagePoliciesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "usagePoliciesChanged")
	}
}

//export callbackQNetworkSessionUsagePoliciesChanged
func callbackQNetworkSessionUsagePoliciesChanged(ptr unsafe.Pointer, ptrName *C.char, usagePolicies C.int) {
	defer qt.Recovering("callback QNetworkSession::usagePoliciesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "usagePoliciesChanged"); signal != nil {
		signal.(func(QNetworkSession__UsagePolicy))(QNetworkSession__UsagePolicy(usagePolicies))
	}

}

func (ptr *QNetworkSession) UsagePoliciesChanged(usagePolicies QNetworkSession__UsagePolicy) {
	defer qt.Recovering("QNetworkSession::usagePoliciesChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_UsagePoliciesChanged(ptr.Pointer(), C.int(usagePolicies))
	}
}

func (ptr *QNetworkSession) WaitForOpened(msecs int) bool {
	defer qt.Recovering("QNetworkSession::waitForOpened")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_WaitForOpened(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QNetworkSession) DestroyQNetworkSession() {
	defer qt.Recovering("QNetworkSession::~QNetworkSession")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DestroyQNetworkSession(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkSession) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkSessionTimerEvent
func callbackQNetworkSessionTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkSession) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkSession) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkSession::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkSession::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkSessionChildEvent
func callbackQNetworkSessionChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkSession::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkSession) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkSession::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkSession) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkSession::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkSession::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkSessionCustomEvent
func callbackQNetworkSessionCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkSession::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkSession) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkSession::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSslCertificate struct {
	ptr unsafe.Pointer
}

type QSslCertificate_ITF interface {
	QSslCertificate_PTR() *QSslCertificate
}

func (p *QSslCertificate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslCertificate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslCertificate(ptr QSslCertificate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificate_PTR().Pointer()
	}
	return nil
}

func NewQSslCertificateFromPointer(ptr unsafe.Pointer) *QSslCertificate {
	var n = new(QSslCertificate)
	n.SetPointer(ptr)
	return n
}

func newQSslCertificateFromPointer(ptr unsafe.Pointer) *QSslCertificate {
	var n = NewQSslCertificateFromPointer(ptr)
	return n
}

func (ptr *QSslCertificate) QSslCertificate_PTR() *QSslCertificate {
	return ptr
}

//QSslCertificate::SubjectInfo
type QSslCertificate__SubjectInfo int64

const (
	QSslCertificate__Organization               = QSslCertificate__SubjectInfo(0)
	QSslCertificate__CommonName                 = QSslCertificate__SubjectInfo(1)
	QSslCertificate__LocalityName               = QSslCertificate__SubjectInfo(2)
	QSslCertificate__OrganizationalUnitName     = QSslCertificate__SubjectInfo(3)
	QSslCertificate__CountryName                = QSslCertificate__SubjectInfo(4)
	QSslCertificate__StateOrProvinceName        = QSslCertificate__SubjectInfo(5)
	QSslCertificate__DistinguishedNameQualifier = QSslCertificate__SubjectInfo(6)
	QSslCertificate__SerialNumber               = QSslCertificate__SubjectInfo(7)
	QSslCertificate__EmailAddress               = QSslCertificate__SubjectInfo(8)
)

func NewQSslCertificate3(other QSslCertificate_ITF) *QSslCertificate {
	defer qt.Recovering("QSslCertificate::QSslCertificate")

	return newQSslCertificateFromPointer(C.QSslCertificate_NewQSslCertificate3(PointerFromQSslCertificate(other)))
}

func (ptr *QSslCertificate) Clear() {
	defer qt.Recovering("QSslCertificate::clear")

	if ptr.Pointer() != nil {
		C.QSslCertificate_Clear(ptr.Pointer())
	}
}

func (ptr *QSslCertificate) Digest(algorithm core.QCryptographicHash__Algorithm) *core.QByteArray {
	defer qt.Recovering("QSslCertificate::digest")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslCertificate_Digest(ptr.Pointer(), C.int(algorithm)))
	}
	return nil
}

func (ptr *QSslCertificate) IsBlacklisted() bool {
	defer qt.Recovering("QSslCertificate::isBlacklisted")

	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsBlacklisted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificate) Swap(other QSslCertificate_ITF) {
	defer qt.Recovering("QSslCertificate::swap")

	if ptr.Pointer() != nil {
		C.QSslCertificate_Swap(ptr.Pointer(), PointerFromQSslCertificate(other))
	}
}

func (ptr *QSslCertificate) DestroyQSslCertificate() {
	defer qt.Recovering("QSslCertificate::~QSslCertificate")

	if ptr.Pointer() != nil {
		C.QSslCertificate_DestroyQSslCertificate(ptr.Pointer())
	}
}

func (ptr *QSslCertificate) EffectiveDate() *core.QDateTime {
	defer qt.Recovering("QSslCertificate::effectiveDate")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QSslCertificate_EffectiveDate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificate) ExpiryDate() *core.QDateTime {
	defer qt.Recovering("QSslCertificate::expiryDate")

	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QSslCertificate_ExpiryDate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificate) IsNull() bool {
	defer qt.Recovering("QSslCertificate::isNull")

	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificate) IsSelfSigned() bool {
	defer qt.Recovering("QSslCertificate::isSelfSigned")

	if ptr.Pointer() != nil {
		return C.QSslCertificate_IsSelfSigned(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificate) IssuerInfo(subject QSslCertificate__SubjectInfo) []string {
	defer qt.Recovering("QSslCertificate::issuerInfo")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_IssuerInfo(ptr.Pointer(), C.int(subject))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) IssuerInfo2(attribute core.QByteArray_ITF) []string {
	defer qt.Recovering("QSslCertificate::issuerInfo")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_IssuerInfo2(ptr.Pointer(), core.PointerFromQByteArray(attribute))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) SerialNumber() *core.QByteArray {
	defer qt.Recovering("QSslCertificate::serialNumber")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslCertificate_SerialNumber(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificate) SubjectInfo(subject QSslCertificate__SubjectInfo) []string {
	defer qt.Recovering("QSslCertificate::subjectInfo")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_SubjectInfo(ptr.Pointer(), C.int(subject))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) SubjectInfo2(attribute core.QByteArray_ITF) []string {
	defer qt.Recovering("QSslCertificate::subjectInfo")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSslCertificate_SubjectInfo2(ptr.Pointer(), core.PointerFromQByteArray(attribute))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSslCertificate) ToDer() *core.QByteArray {
	defer qt.Recovering("QSslCertificate::toDer")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslCertificate_ToDer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificate) ToPem() *core.QByteArray {
	defer qt.Recovering("QSslCertificate::toPem")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslCertificate_ToPem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificate) ToText() string {
	defer qt.Recovering("QSslCertificate::toText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificate_ToText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificate) Version() *core.QByteArray {
	defer qt.Recovering("QSslCertificate::version")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslCertificate_Version(ptr.Pointer()))
	}
	return nil
}

type QSslCertificateExtension struct {
	ptr unsafe.Pointer
}

type QSslCertificateExtension_ITF interface {
	QSslCertificateExtension_PTR() *QSslCertificateExtension
}

func (p *QSslCertificateExtension) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslCertificateExtension) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslCertificateExtension(ptr QSslCertificateExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificateExtension_PTR().Pointer()
	}
	return nil
}

func NewQSslCertificateExtensionFromPointer(ptr unsafe.Pointer) *QSslCertificateExtension {
	var n = new(QSslCertificateExtension)
	n.SetPointer(ptr)
	return n
}

func newQSslCertificateExtensionFromPointer(ptr unsafe.Pointer) *QSslCertificateExtension {
	var n = NewQSslCertificateExtensionFromPointer(ptr)
	return n
}

func (ptr *QSslCertificateExtension) QSslCertificateExtension_PTR() *QSslCertificateExtension {
	return ptr
}

func NewQSslCertificateExtension() *QSslCertificateExtension {
	defer qt.Recovering("QSslCertificateExtension::QSslCertificateExtension")

	return newQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension())
}

func NewQSslCertificateExtension2(other QSslCertificateExtension_ITF) *QSslCertificateExtension {
	defer qt.Recovering("QSslCertificateExtension::QSslCertificateExtension")

	return newQSslCertificateExtensionFromPointer(C.QSslCertificateExtension_NewQSslCertificateExtension2(PointerFromQSslCertificateExtension(other)))
}

func (ptr *QSslCertificateExtension) IsCritical() bool {
	defer qt.Recovering("QSslCertificateExtension::isCritical")

	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsCritical(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) IsSupported() bool {
	defer qt.Recovering("QSslCertificateExtension::isSupported")

	if ptr.Pointer() != nil {
		return C.QSslCertificateExtension_IsSupported(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCertificateExtension) Name() string {
	defer qt.Recovering("QSslCertificateExtension::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Oid() string {
	defer qt.Recovering("QSslCertificateExtension::oid")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCertificateExtension_Oid(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCertificateExtension) Swap(other QSslCertificateExtension_ITF) {
	defer qt.Recovering("QSslCertificateExtension::swap")

	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_Swap(ptr.Pointer(), PointerFromQSslCertificateExtension(other))
	}
}

func (ptr *QSslCertificateExtension) Value() *core.QVariant {
	defer qt.Recovering("QSslCertificateExtension::value")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSslCertificateExtension_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslCertificateExtension) DestroyQSslCertificateExtension() {
	defer qt.Recovering("QSslCertificateExtension::~QSslCertificateExtension")

	if ptr.Pointer() != nil {
		C.QSslCertificateExtension_DestroyQSslCertificateExtension(ptr.Pointer())
	}
}

type QSslCipher struct {
	ptr unsafe.Pointer
}

type QSslCipher_ITF interface {
	QSslCipher_PTR() *QSslCipher
}

func (p *QSslCipher) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslCipher) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslCipher(ptr QSslCipher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCipher_PTR().Pointer()
	}
	return nil
}

func NewQSslCipherFromPointer(ptr unsafe.Pointer) *QSslCipher {
	var n = new(QSslCipher)
	n.SetPointer(ptr)
	return n
}

func newQSslCipherFromPointer(ptr unsafe.Pointer) *QSslCipher {
	var n = NewQSslCipherFromPointer(ptr)
	return n
}

func (ptr *QSslCipher) QSslCipher_PTR() *QSslCipher {
	return ptr
}

func NewQSslCipher() *QSslCipher {
	defer qt.Recovering("QSslCipher::QSslCipher")

	return newQSslCipherFromPointer(C.QSslCipher_NewQSslCipher())
}

func NewQSslCipher4(other QSslCipher_ITF) *QSslCipher {
	defer qt.Recovering("QSslCipher::QSslCipher")

	return newQSslCipherFromPointer(C.QSslCipher_NewQSslCipher4(PointerFromQSslCipher(other)))
}

func NewQSslCipher2(name string) *QSslCipher {
	defer qt.Recovering("QSslCipher::QSslCipher")

	return newQSslCipherFromPointer(C.QSslCipher_NewQSslCipher2(C.CString(name)))
}

func (ptr *QSslCipher) AuthenticationMethod() string {
	defer qt.Recovering("QSslCipher::authenticationMethod")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_AuthenticationMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) EncryptionMethod() string {
	defer qt.Recovering("QSslCipher::encryptionMethod")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_EncryptionMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) IsNull() bool {
	defer qt.Recovering("QSslCipher::isNull")

	if ptr.Pointer() != nil {
		return C.QSslCipher_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslCipher) KeyExchangeMethod() string {
	defer qt.Recovering("QSslCipher::keyExchangeMethod")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_KeyExchangeMethod(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) Name() string {
	defer qt.Recovering("QSslCipher::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) ProtocolString() string {
	defer qt.Recovering("QSslCipher::protocolString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_ProtocolString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslCipher) SupportedBits() int {
	defer qt.Recovering("QSslCipher::supportedBits")

	if ptr.Pointer() != nil {
		return int(C.QSslCipher_SupportedBits(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslCipher) Swap(other QSslCipher_ITF) {
	defer qt.Recovering("QSslCipher::swap")

	if ptr.Pointer() != nil {
		C.QSslCipher_Swap(ptr.Pointer(), PointerFromQSslCipher(other))
	}
}

func (ptr *QSslCipher) UsedBits() int {
	defer qt.Recovering("QSslCipher::usedBits")

	if ptr.Pointer() != nil {
		return int(C.QSslCipher_UsedBits(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslCipher) DestroyQSslCipher() {
	defer qt.Recovering("QSslCipher::~QSslCipher")

	if ptr.Pointer() != nil {
		C.QSslCipher_DestroyQSslCipher(ptr.Pointer())
	}
}

type QSslConfiguration struct {
	ptr unsafe.Pointer
}

type QSslConfiguration_ITF interface {
	QSslConfiguration_PTR() *QSslConfiguration
}

func (p *QSslConfiguration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslConfiguration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslConfiguration(ptr QSslConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslConfiguration_PTR().Pointer()
	}
	return nil
}

func NewQSslConfigurationFromPointer(ptr unsafe.Pointer) *QSslConfiguration {
	var n = new(QSslConfiguration)
	n.SetPointer(ptr)
	return n
}

func newQSslConfigurationFromPointer(ptr unsafe.Pointer) *QSslConfiguration {
	var n = NewQSslConfigurationFromPointer(ptr)
	return n
}

func (ptr *QSslConfiguration) QSslConfiguration_PTR() *QSslConfiguration {
	return ptr
}

//QSslConfiguration::NextProtocolNegotiationStatus
type QSslConfiguration__NextProtocolNegotiationStatus int64

const (
	QSslConfiguration__NextProtocolNegotiationNone        = QSslConfiguration__NextProtocolNegotiationStatus(0)
	QSslConfiguration__NextProtocolNegotiationNegotiated  = QSslConfiguration__NextProtocolNegotiationStatus(1)
	QSslConfiguration__NextProtocolNegotiationUnsupported = QSslConfiguration__NextProtocolNegotiationStatus(2)
)

func NewQSslConfiguration() *QSslConfiguration {
	defer qt.Recovering("QSslConfiguration::QSslConfiguration")

	return newQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration())
}

func NewQSslConfiguration2(other QSslConfiguration_ITF) *QSslConfiguration {
	defer qt.Recovering("QSslConfiguration::QSslConfiguration")

	return newQSslConfigurationFromPointer(C.QSslConfiguration_NewQSslConfiguration2(PointerFromQSslConfiguration(other)))
}

func (ptr *QSslConfiguration) IsNull() bool {
	defer qt.Recovering("QSslConfiguration::isNull")

	if ptr.Pointer() != nil {
		return C.QSslConfiguration_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslConfiguration) NextNegotiatedProtocol() *core.QByteArray {
	defer qt.Recovering("QSslConfiguration::nextNegotiatedProtocol")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslConfiguration_NextNegotiatedProtocol(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslConfiguration) NextProtocolNegotiationStatus() QSslConfiguration__NextProtocolNegotiationStatus {
	defer qt.Recovering("QSslConfiguration::nextProtocolNegotiationStatus")

	if ptr.Pointer() != nil {
		return QSslConfiguration__NextProtocolNegotiationStatus(C.QSslConfiguration_NextProtocolNegotiationStatus(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerVerifyDepth() int {
	defer qt.Recovering("QSslConfiguration::peerVerifyDepth")

	if ptr.Pointer() != nil {
		return int(C.QSslConfiguration_PeerVerifyDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	defer qt.Recovering("QSslConfiguration::peerVerifyMode")

	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslConfiguration_PeerVerifyMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslConfiguration) SessionTicket() *core.QByteArray {
	defer qt.Recovering("QSslConfiguration::sessionTicket")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslConfiguration_SessionTicket(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslConfiguration) SessionTicketLifeTimeHint() int {
	defer qt.Recovering("QSslConfiguration::sessionTicketLifeTimeHint")

	if ptr.Pointer() != nil {
		return int(C.QSslConfiguration_SessionTicketLifeTimeHint(ptr.Pointer()))
	}
	return 0
}

func QSslConfiguration_SetDefaultConfiguration(configuration QSslConfiguration_ITF) {
	defer qt.Recovering("QSslConfiguration::setDefaultConfiguration")

	C.QSslConfiguration_QSslConfiguration_SetDefaultConfiguration(PointerFromQSslConfiguration(configuration))
}

func (ptr *QSslConfiguration) SetLocalCertificate(certificate QSslCertificate_ITF) {
	defer qt.Recovering("QSslConfiguration::setLocalCertificate")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyDepth(depth int) {
	defer qt.Recovering("QSslConfiguration::setPeerVerifyDepth")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QSslConfiguration) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	defer qt.Recovering("QSslConfiguration::setPeerVerifyMode")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPeerVerifyMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSslConfiguration) SetPrivateKey(key QSslKey_ITF) {
	defer qt.Recovering("QSslConfiguration::setPrivateKey")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

func (ptr *QSslConfiguration) SetSessionTicket(sessionTicket core.QByteArray_ITF) {
	defer qt.Recovering("QSslConfiguration::setSessionTicket")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_SetSessionTicket(ptr.Pointer(), core.PointerFromQByteArray(sessionTicket))
	}
}

func (ptr *QSslConfiguration) Swap(other QSslConfiguration_ITF) {
	defer qt.Recovering("QSslConfiguration::swap")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_Swap(ptr.Pointer(), PointerFromQSslConfiguration(other))
	}
}

func (ptr *QSslConfiguration) DestroyQSslConfiguration() {
	defer qt.Recovering("QSslConfiguration::~QSslConfiguration")

	if ptr.Pointer() != nil {
		C.QSslConfiguration_DestroyQSslConfiguration(ptr.Pointer())
	}
}

type QSslEllipticCurve struct {
	ptr unsafe.Pointer
}

type QSslEllipticCurve_ITF interface {
	QSslEllipticCurve_PTR() *QSslEllipticCurve
}

func (p *QSslEllipticCurve) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslEllipticCurve) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslEllipticCurve(ptr QSslEllipticCurve_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslEllipticCurve_PTR().Pointer()
	}
	return nil
}

func NewQSslEllipticCurveFromPointer(ptr unsafe.Pointer) *QSslEllipticCurve {
	var n = new(QSslEllipticCurve)
	n.SetPointer(ptr)
	return n
}

func newQSslEllipticCurveFromPointer(ptr unsafe.Pointer) *QSslEllipticCurve {
	var n = NewQSslEllipticCurveFromPointer(ptr)
	return n
}

func (ptr *QSslEllipticCurve) QSslEllipticCurve_PTR() *QSslEllipticCurve {
	return ptr
}

func NewQSslEllipticCurve() *QSslEllipticCurve {
	defer qt.Recovering("QSslEllipticCurve::QSslEllipticCurve")

	return newQSslEllipticCurveFromPointer(C.QSslEllipticCurve_NewQSslEllipticCurve())
}

func (ptr *QSslEllipticCurve) IsValid() bool {
	defer qt.Recovering("QSslEllipticCurve::isValid")

	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) IsTlsNamedCurve() bool {
	defer qt.Recovering("QSslEllipticCurve::isTlsNamedCurve")

	if ptr.Pointer() != nil {
		return C.QSslEllipticCurve_IsTlsNamedCurve(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslEllipticCurve) LongName() string {
	defer qt.Recovering("QSslEllipticCurve::longName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslEllipticCurve_LongName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslEllipticCurve) ShortName() string {
	defer qt.Recovering("QSslEllipticCurve::shortName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslEllipticCurve_ShortName(ptr.Pointer()))
	}
	return ""
}

type QSslError struct {
	ptr unsafe.Pointer
}

type QSslError_ITF interface {
	QSslError_PTR() *QSslError
}

func (p *QSslError) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslError) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslError(ptr QSslError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslError_PTR().Pointer()
	}
	return nil
}

func NewQSslErrorFromPointer(ptr unsafe.Pointer) *QSslError {
	var n = new(QSslError)
	n.SetPointer(ptr)
	return n
}

func newQSslErrorFromPointer(ptr unsafe.Pointer) *QSslError {
	var n = NewQSslErrorFromPointer(ptr)
	return n
}

func (ptr *QSslError) QSslError_PTR() *QSslError {
	return ptr
}

//QSslError::SslError
type QSslError__SslError int64

const (
	QSslError__NoError                             = QSslError__SslError(0)
	QSslError__UnableToGetIssuerCertificate        = QSslError__SslError(1)
	QSslError__UnableToDecryptCertificateSignature = QSslError__SslError(2)
	QSslError__UnableToDecodeIssuerPublicKey       = QSslError__SslError(3)
	QSslError__CertificateSignatureFailed          = QSslError__SslError(4)
	QSslError__CertificateNotYetValid              = QSslError__SslError(5)
	QSslError__CertificateExpired                  = QSslError__SslError(6)
	QSslError__InvalidNotBeforeField               = QSslError__SslError(7)
	QSslError__InvalidNotAfterField                = QSslError__SslError(8)
	QSslError__SelfSignedCertificate               = QSslError__SslError(9)
	QSslError__SelfSignedCertificateInChain        = QSslError__SslError(10)
	QSslError__UnableToGetLocalIssuerCertificate   = QSslError__SslError(11)
	QSslError__UnableToVerifyFirstCertificate      = QSslError__SslError(12)
	QSslError__CertificateRevoked                  = QSslError__SslError(13)
	QSslError__InvalidCaCertificate                = QSslError__SslError(14)
	QSslError__PathLengthExceeded                  = QSslError__SslError(15)
	QSslError__InvalidPurpose                      = QSslError__SslError(16)
	QSslError__CertificateUntrusted                = QSslError__SslError(17)
	QSslError__CertificateRejected                 = QSslError__SslError(18)
	QSslError__SubjectIssuerMismatch               = QSslError__SslError(19)
	QSslError__AuthorityIssuerSerialNumberMismatch = QSslError__SslError(20)
	QSslError__NoPeerCertificate                   = QSslError__SslError(21)
	QSslError__HostNameMismatch                    = QSslError__SslError(22)
	QSslError__NoSslSupport                        = QSslError__SslError(23)
	QSslError__CertificateBlacklisted              = QSslError__SslError(24)
	QSslError__UnspecifiedError                    = QSslError__SslError(-1)
)

func NewQSslError() *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	return newQSslErrorFromPointer(C.QSslError_NewQSslError())
}

func NewQSslError2(error QSslError__SslError) *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	return newQSslErrorFromPointer(C.QSslError_NewQSslError2(C.int(error)))
}

func NewQSslError3(error QSslError__SslError, certificate QSslCertificate_ITF) *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	return newQSslErrorFromPointer(C.QSslError_NewQSslError3(C.int(error), PointerFromQSslCertificate(certificate)))
}

func NewQSslError4(other QSslError_ITF) *QSslError {
	defer qt.Recovering("QSslError::QSslError")

	return newQSslErrorFromPointer(C.QSslError_NewQSslError4(PointerFromQSslError(other)))
}

func (ptr *QSslError) Error() QSslError__SslError {
	defer qt.Recovering("QSslError::error")

	if ptr.Pointer() != nil {
		return QSslError__SslError(C.QSslError_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslError) ErrorString() string {
	defer qt.Recovering("QSslError::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslError_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslError) Swap(other QSslError_ITF) {
	defer qt.Recovering("QSslError::swap")

	if ptr.Pointer() != nil {
		C.QSslError_Swap(ptr.Pointer(), PointerFromQSslError(other))
	}
}

func (ptr *QSslError) DestroyQSslError() {
	defer qt.Recovering("QSslError::~QSslError")

	if ptr.Pointer() != nil {
		C.QSslError_DestroyQSslError(ptr.Pointer())
	}
}

type QSslKey struct {
	ptr unsafe.Pointer
}

type QSslKey_ITF interface {
	QSslKey_PTR() *QSslKey
}

func (p *QSslKey) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslKey) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslKey(ptr QSslKey_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslKey_PTR().Pointer()
	}
	return nil
}

func NewQSslKeyFromPointer(ptr unsafe.Pointer) *QSslKey {
	var n = new(QSslKey)
	n.SetPointer(ptr)
	return n
}

func newQSslKeyFromPointer(ptr unsafe.Pointer) *QSslKey {
	var n = NewQSslKeyFromPointer(ptr)
	return n
}

func (ptr *QSslKey) QSslKey_PTR() *QSslKey {
	return ptr
}

func NewQSslKey() *QSslKey {
	defer qt.Recovering("QSslKey::QSslKey")

	return newQSslKeyFromPointer(C.QSslKey_NewQSslKey())
}

func NewQSslKey5(other QSslKey_ITF) *QSslKey {
	defer qt.Recovering("QSslKey::QSslKey")

	return newQSslKeyFromPointer(C.QSslKey_NewQSslKey5(PointerFromQSslKey(other)))
}

func (ptr *QSslKey) Clear() {
	defer qt.Recovering("QSslKey::clear")

	if ptr.Pointer() != nil {
		C.QSslKey_Clear(ptr.Pointer())
	}
}

func (ptr *QSslKey) IsNull() bool {
	defer qt.Recovering("QSslKey::isNull")

	if ptr.Pointer() != nil {
		return C.QSslKey_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslKey) Length() int {
	defer qt.Recovering("QSslKey::length")

	if ptr.Pointer() != nil {
		return int(C.QSslKey_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslKey) Swap(other QSslKey_ITF) {
	defer qt.Recovering("QSslKey::swap")

	if ptr.Pointer() != nil {
		C.QSslKey_Swap(ptr.Pointer(), PointerFromQSslKey(other))
	}
}

func (ptr *QSslKey) ToDer(passPhrase core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QSslKey::toDer")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslKey_ToDer(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
	}
	return nil
}

func (ptr *QSslKey) ToPem(passPhrase core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QSslKey::toPem")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslKey_ToPem(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
	}
	return nil
}

func (ptr *QSslKey) DestroyQSslKey() {
	defer qt.Recovering("QSslKey::~QSslKey")

	if ptr.Pointer() != nil {
		C.QSslKey_DestroyQSslKey(ptr.Pointer())
	}
}

type QSslPreSharedKeyAuthenticator struct {
	ptr unsafe.Pointer
}

type QSslPreSharedKeyAuthenticator_ITF interface {
	QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator
}

func (p *QSslPreSharedKeyAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslPreSharedKeyAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslPreSharedKeyAuthenticator(ptr QSslPreSharedKeyAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslPreSharedKeyAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQSslPreSharedKeyAuthenticatorFromPointer(ptr unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	var n = new(QSslPreSharedKeyAuthenticator)
	n.SetPointer(ptr)
	return n
}

func newQSslPreSharedKeyAuthenticatorFromPointer(ptr unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	var n = NewQSslPreSharedKeyAuthenticatorFromPointer(ptr)
	return n
}

func (ptr *QSslPreSharedKeyAuthenticator) QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator {
	return ptr
}

func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::QSslPreSharedKeyAuthenticator")

	return newQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator())
}

func NewQSslPreSharedKeyAuthenticator2(authenticator QSslPreSharedKeyAuthenticator_ITF) *QSslPreSharedKeyAuthenticator {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::QSslPreSharedKeyAuthenticator")

	return newQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(PointerFromQSslPreSharedKeyAuthenticator(authenticator)))
}

func (ptr *QSslPreSharedKeyAuthenticator) Identity() *core.QByteArray {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::identity")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_Identity(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) IdentityHint() *core.QByteArray {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::identityHint")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_IdentityHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::maximumIdentityLength")

	if ptr.Pointer() != nil {
		return int(C.QSslPreSharedKeyAuthenticator_MaximumIdentityLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::maximumPreSharedKeyLength")

	if ptr.Pointer() != nil {
		return int(C.QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) PreSharedKey() *core.QByteArray {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::preSharedKey")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_PreSharedKey(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) SetIdentity(identity core.QByteArray_ITF) {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::setIdentity")

	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_SetIdentity(ptr.Pointer(), core.PointerFromQByteArray(identity))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey core.QByteArray_ITF) {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::setPreSharedKey")

	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_SetPreSharedKey(ptr.Pointer(), core.PointerFromQByteArray(preSharedKey))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) Swap(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::swap")

	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_Swap(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) DestroyQSslPreSharedKeyAuthenticator() {
	defer qt.Recovering("QSslPreSharedKeyAuthenticator::~QSslPreSharedKeyAuthenticator")

	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(ptr.Pointer())
	}
}

type QSslSocket struct {
	QTcpSocket
}

type QSslSocket_ITF interface {
	QTcpSocket_ITF
	QSslSocket_PTR() *QSslSocket
}

func PointerFromQSslSocket(ptr QSslSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslSocket_PTR().Pointer()
	}
	return nil
}

func NewQSslSocketFromPointer(ptr unsafe.Pointer) *QSslSocket {
	var n = new(QSslSocket)
	n.SetPointer(ptr)
	return n
}

func newQSslSocketFromPointer(ptr unsafe.Pointer) *QSslSocket {
	var n = NewQSslSocketFromPointer(ptr)
	for len(n.ObjectName()) < len("QSslSocket_") {
		n.SetObjectName("QSslSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QSslSocket) QSslSocket_PTR() *QSslSocket {
	return ptr
}

//QSslSocket::PeerVerifyMode
type QSslSocket__PeerVerifyMode int64

const (
	QSslSocket__VerifyNone     = QSslSocket__PeerVerifyMode(0)
	QSslSocket__QueryPeer      = QSslSocket__PeerVerifyMode(1)
	QSslSocket__VerifyPeer     = QSslSocket__PeerVerifyMode(2)
	QSslSocket__AutoVerifyPeer = QSslSocket__PeerVerifyMode(3)
)

//QSslSocket::SslMode
type QSslSocket__SslMode int64

const (
	QSslSocket__UnencryptedMode = QSslSocket__SslMode(0)
	QSslSocket__SslClientMode   = QSslSocket__SslMode(1)
	QSslSocket__SslServerMode   = QSslSocket__SslMode(2)
)

func NewQSslSocket(parent core.QObject_ITF) *QSslSocket {
	defer qt.Recovering("QSslSocket::QSslSocket")

	return newQSslSocketFromPointer(C.QSslSocket_NewQSslSocket(core.PointerFromQObject(parent)))
}

func (ptr *QSslSocket) Abort() {
	defer qt.Recovering("QSslSocket::abort")

	if ptr.Pointer() != nil {
		C.QSslSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QSslSocket) AddCaCertificate(certificate QSslCertificate_ITF) {
	defer qt.Recovering("QSslSocket::addCaCertificate")

	if ptr.Pointer() != nil {
		C.QSslSocket_AddCaCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	defer qt.Recovering("QSslSocket::addDefaultCaCertificate")

	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(PointerFromQSslCertificate(certificate))
}

func (ptr *QSslSocket) AtEnd() bool {
	defer qt.Recovering("QSslSocket::atEnd")

	if ptr.Pointer() != nil {
		return C.QSslSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) BytesAvailable() int64 {
	defer qt.Recovering("QSslSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) BytesToWrite() int64 {
	defer qt.Recovering("QSslSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) CanReadLine() bool {
	defer qt.Recovering("QSslSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QSslSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QSslSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QSslSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QSslSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQSslSocketClose
func callbackQSslSocketClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSslSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QSslSocket) Close() {
	defer qt.Recovering("QSslSocket::close")

	if ptr.Pointer() != nil {
		C.QSslSocket_Close(ptr.Pointer())
	}
}

func (ptr *QSslSocket) CloseDefault() {
	defer qt.Recovering("QSslSocket::close")

	if ptr.Pointer() != nil {
		C.QSslSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QSslSocket) ConnectEncrypted(f func()) {
	defer qt.Recovering("connect QSslSocket::encrypted")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QSslSocket) DisconnectEncrypted() {
	defer qt.Recovering("disconnect QSslSocket::encrypted")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQSslSocketEncrypted
func callbackQSslSocketEncrypted(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSslSocket::encrypted")

	if signal := qt.GetSignal(C.GoString(ptrName), "encrypted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSslSocket) Encrypted() {
	defer qt.Recovering("QSslSocket::encrypted")

	if ptr.Pointer() != nil {
		C.QSslSocket_Encrypted(ptr.Pointer())
	}
}

func (ptr *QSslSocket) EncryptedBytesAvailable() int64 {
	defer qt.Recovering("QSslSocket::encryptedBytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_EncryptedBytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) EncryptedBytesToWrite() int64 {
	defer qt.Recovering("QSslSocket::encryptedBytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_EncryptedBytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) ConnectEncryptedBytesWritten(f func(written int64)) {
	defer qt.Recovering("connect QSslSocket::encryptedBytesWritten")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectEncryptedBytesWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "encryptedBytesWritten", f)
	}
}

func (ptr *QSslSocket) DisconnectEncryptedBytesWritten() {
	defer qt.Recovering("disconnect QSslSocket::encryptedBytesWritten")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncryptedBytesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "encryptedBytesWritten")
	}
}

//export callbackQSslSocketEncryptedBytesWritten
func callbackQSslSocketEncryptedBytesWritten(ptr unsafe.Pointer, ptrName *C.char, written C.longlong) {
	defer qt.Recovering("callback QSslSocket::encryptedBytesWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "encryptedBytesWritten"); signal != nil {
		signal.(func(int64))(int64(written))
	}

}

func (ptr *QSslSocket) EncryptedBytesWritten(written int64) {
	defer qt.Recovering("QSslSocket::encryptedBytesWritten")

	if ptr.Pointer() != nil {
		C.QSslSocket_EncryptedBytesWritten(ptr.Pointer(), C.longlong(written))
	}
}

func (ptr *QSslSocket) Flush() bool {
	defer qt.Recovering("QSslSocket::flush")

	if ptr.Pointer() != nil {
		return C.QSslSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) IgnoreSslErrors() {
	defer qt.Recovering("QSslSocket::ignoreSslErrors")

	if ptr.Pointer() != nil {
		C.QSslSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QSslSocket) IsEncrypted() bool {
	defer qt.Recovering("QSslSocket::isEncrypted")

	if ptr.Pointer() != nil {
		return C.QSslSocket_IsEncrypted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) Mode() QSslSocket__SslMode {
	defer qt.Recovering("QSslSocket::mode")

	if ptr.Pointer() != nil {
		return QSslSocket__SslMode(C.QSslSocket_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) ConnectModeChanged(f func(mode QSslSocket__SslMode)) {
	defer qt.Recovering("connect QSslSocket::modeChanged")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modeChanged", f)
	}
}

func (ptr *QSslSocket) DisconnectModeChanged() {
	defer qt.Recovering("disconnect QSslSocket::modeChanged")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modeChanged")
	}
}

//export callbackQSslSocketModeChanged
func callbackQSslSocketModeChanged(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QSslSocket::modeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "modeChanged"); signal != nil {
		signal.(func(QSslSocket__SslMode))(QSslSocket__SslMode(mode))
	}

}

func (ptr *QSslSocket) ModeChanged(mode QSslSocket__SslMode) {
	defer qt.Recovering("QSslSocket::modeChanged")

	if ptr.Pointer() != nil {
		C.QSslSocket_ModeChanged(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSslSocket) PeerVerifyDepth() int {
	defer qt.Recovering("QSslSocket::peerVerifyDepth")

	if ptr.Pointer() != nil {
		return int(C.QSslSocket_PeerVerifyDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	defer qt.Recovering("QSslSocket::peerVerifyMode")

	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslSocket_PeerVerifyMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyName() string {
	defer qt.Recovering("QSslSocket::peerVerifyName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslSocket_PeerVerifyName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	defer qt.Recovering("connect QSslSocket::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QSslSocket) DisconnectPreSharedKeyAuthenticationRequired() {
	defer qt.Recovering("disconnect QSslSocket::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQSslSocketPreSharedKeyAuthenticationRequired
func callbackQSslSocketPreSharedKeyAuthenticationRequired(ptr unsafe.Pointer, ptrName *C.char, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::preSharedKeyAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

}

func (ptr *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	defer qt.Recovering("QSslSocket::preSharedKeyAuthenticationRequired")

	if ptr.Pointer() != nil {
		C.QSslSocket_PreSharedKeyAuthenticationRequired(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QSslSocket) ReadData(data string, maxlen int64) int64 {
	defer qt.Recovering("QSslSocket::readData")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_ReadData(ptr.Pointer(), C.CString(data), C.longlong(maxlen)))
	}
	return 0
}

func (ptr *QSslSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QSslSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resume", f)
	}
}

func (ptr *QSslSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QSslSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resume")
	}
}

//export callbackQSslSocketResume
func callbackQSslSocketResume(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSslSocket::resume")

	if signal := qt.GetSignal(C.GoString(ptrName), "resume"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QSslSocket) Resume() {
	defer qt.Recovering("QSslSocket::resume")

	if ptr.Pointer() != nil {
		C.QSslSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QSslSocket) ResumeDefault() {
	defer qt.Recovering("QSslSocket::resume")

	if ptr.Pointer() != nil {
		C.QSslSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QSslSocket) SetLocalCertificate(certificate QSslCertificate_ITF) {
	defer qt.Recovering("QSslSocket::setLocalCertificate")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslSocket) SetPeerVerifyDepth(depth int) {
	defer qt.Recovering("QSslSocket::setPeerVerifyDepth")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QSslSocket) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	defer qt.Recovering("QSslSocket::setPeerVerifyMode")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSslSocket) SetPeerVerifyName(hostName string) {
	defer qt.Recovering("QSslSocket::setPeerVerifyName")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyName(ptr.Pointer(), C.CString(hostName))
	}
}

func (ptr *QSslSocket) SetPrivateKey(key QSslKey_ITF) {
	defer qt.Recovering("QSslSocket::setPrivateKey")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

func (ptr *QSslSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QSslSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setReadBufferSize", f)
	}
}

func (ptr *QSslSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QSslSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setReadBufferSize")
	}
}

//export callbackQSslSocketSetReadBufferSize
func callbackQSslSocketSetReadBufferSize(ptr unsafe.Pointer, ptrName *C.char, size C.longlong) {
	defer qt.Recovering("callback QSslSocket::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQSslSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QSslSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QSslSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QSslSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QSslSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QSslSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QSslSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSocketOption", f)
	}
}

func (ptr *QSslSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QSslSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSocketOption")
	}
}

//export callbackQSslSocketSetSocketOption
func callbackQSslSocketSetSocketOption(ptr unsafe.Pointer, ptrName *C.char, option C.int, value unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::setSocketOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQSslSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QSslSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QSslSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetSocketOption(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QSslSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QSslSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetSocketOptionDefault(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QSslSocket) SetSslConfiguration(configuration QSslConfiguration_ITF) {
	defer qt.Recovering("QSslSocket::setSslConfiguration")

	if ptr.Pointer() != nil {
		C.QSslSocket_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

func (ptr *QSslSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	defer qt.Recovering("QSslSocket::socketOption")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSslSocket_SocketOption(ptr.Pointer(), C.int(option)))
	}
	return nil
}

func QSslSocket_SslLibraryBuildVersionString() string {
	defer qt.Recovering("QSslSocket::sslLibraryBuildVersionString")

	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryBuildVersionString())
}

func QSslSocket_SslLibraryVersionString() string {
	defer qt.Recovering("QSslSocket::sslLibraryVersionString")

	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryVersionString())
}

func (ptr *QSslSocket) StartClientEncryption() {
	defer qt.Recovering("QSslSocket::startClientEncryption")

	if ptr.Pointer() != nil {
		C.QSslSocket_StartClientEncryption(ptr.Pointer())
	}
}

func (ptr *QSslSocket) StartServerEncryption() {
	defer qt.Recovering("QSslSocket::startServerEncryption")

	if ptr.Pointer() != nil {
		C.QSslSocket_StartServerEncryption(ptr.Pointer())
	}
}

func QSslSocket_SupportsSsl() bool {
	defer qt.Recovering("QSslSocket::supportsSsl")

	return C.QSslSocket_QSslSocket_SupportsSsl() != 0
}

func (ptr *QSslSocket) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForConnected(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForConnected")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForConnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForDisconnected(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForDisconnected")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForDisconnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForEncrypted(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForEncrypted")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForEncrypted(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QSslSocket::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WriteData(data string, len int64) int64 {
	defer qt.Recovering("QSslSocket::writeData")

	if ptr.Pointer() != nil {
		return int64(C.QSslSocket_WriteData(ptr.Pointer(), C.CString(data), C.longlong(len)))
	}
	return 0
}

func (ptr *QSslSocket) DestroyQSslSocket() {
	defer qt.Recovering("QSslSocket::~QSslSocket")

	if ptr.Pointer() != nil {
		C.QSslSocket_DestroyQSslSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSslSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QSslSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectFromHost", f)
	}
}

func (ptr *QSslSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QSslSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectFromHost")
	}
}

//export callbackQSslSocketDisconnectFromHost
func callbackQSslSocketDisconnectFromHost(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSslSocket::disconnectFromHost")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQSslSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QSslSocket) DisconnectFromHost() {
	defer qt.Recovering("QSslSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QSslSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QSslSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

func (ptr *QSslSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSslSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSslSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSslSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSslSocketTimerEvent
func callbackQSslSocketTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSslSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSslSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSslSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSslSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSslSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSslSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSslSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSslSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSslSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSslSocketChildEvent
func callbackQSslSocketChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSslSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSslSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSslSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSslSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSslSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSslSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSslSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSslSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSslSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSslSocketCustomEvent
func callbackQSslSocketCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSslSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSslSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSslSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSslSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSslSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QSslSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QTcpServer struct {
	core.QObject
}

type QTcpServer_ITF interface {
	core.QObject_ITF
	QTcpServer_PTR() *QTcpServer
}

func PointerFromQTcpServer(ptr QTcpServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpServer_PTR().Pointer()
	}
	return nil
}

func NewQTcpServerFromPointer(ptr unsafe.Pointer) *QTcpServer {
	var n = new(QTcpServer)
	n.SetPointer(ptr)
	return n
}

func newQTcpServerFromPointer(ptr unsafe.Pointer) *QTcpServer {
	var n = NewQTcpServerFromPointer(ptr)
	for len(n.ObjectName()) < len("QTcpServer_") {
		n.SetObjectName("QTcpServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QTcpServer) QTcpServer_PTR() *QTcpServer {
	return ptr
}

func NewQTcpServer(parent core.QObject_ITF) *QTcpServer {
	defer qt.Recovering("QTcpServer::QTcpServer")

	return newQTcpServerFromPointer(C.QTcpServer_NewQTcpServer(core.PointerFromQObject(parent)))
}

func (ptr *QTcpServer) ConnectAcceptError(f func(socketError QAbstractSocket__SocketError)) {
	defer qt.Recovering("connect QTcpServer::acceptError")

	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectAcceptError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "acceptError", f)
	}
}

func (ptr *QTcpServer) DisconnectAcceptError() {
	defer qt.Recovering("disconnect QTcpServer::acceptError")

	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectAcceptError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "acceptError")
	}
}

//export callbackQTcpServerAcceptError
func callbackQTcpServerAcceptError(ptr unsafe.Pointer, ptrName *C.char, socketError C.int) {
	defer qt.Recovering("callback QTcpServer::acceptError")

	if signal := qt.GetSignal(C.GoString(ptrName), "acceptError"); signal != nil {
		signal.(func(QAbstractSocket__SocketError))(QAbstractSocket__SocketError(socketError))
	}

}

func (ptr *QTcpServer) AcceptError(socketError QAbstractSocket__SocketError) {
	defer qt.Recovering("QTcpServer::acceptError")

	if ptr.Pointer() != nil {
		C.QTcpServer_AcceptError(ptr.Pointer(), C.int(socketError))
	}
}

func (ptr *QTcpServer) Close() {
	defer qt.Recovering("QTcpServer::close")

	if ptr.Pointer() != nil {
		C.QTcpServer_Close(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ErrorString() string {
	defer qt.Recovering("QTcpServer::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTcpServer_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTcpServer) HasPendingConnections() bool {
	defer qt.Recovering("QTcpServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QTcpServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) IsListening() bool {
	defer qt.Recovering("QTcpServer::isListening")

	if ptr.Pointer() != nil {
		return C.QTcpServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTcpServer) MaxPendingConnections() int {
	defer qt.Recovering("QTcpServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(C.QTcpServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QTcpServer::newConnection")

	if ptr.Pointer() != nil {
		C.QTcpServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QTcpServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QTcpServer::newConnection")

	if ptr.Pointer() != nil {
		C.QTcpServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

//export callbackQTcpServerNewConnection
func callbackQTcpServerNewConnection(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTcpServer::newConnection")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTcpServer) NewConnection() {
	defer qt.Recovering("QTcpServer::newConnection")

	if ptr.Pointer() != nil {
		C.QTcpServer_NewConnection(ptr.Pointer())
	}
}

func (ptr *QTcpServer) NextPendingConnection() *QTcpSocket {
	defer qt.Recovering("QTcpServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQTcpSocketFromPointer(C.QTcpServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTcpServer) PauseAccepting() {
	defer qt.Recovering("QTcpServer::pauseAccepting")

	if ptr.Pointer() != nil {
		C.QTcpServer_PauseAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ResumeAccepting() {
	defer qt.Recovering("QTcpServer::resumeAccepting")

	if ptr.Pointer() != nil {
		C.QTcpServer_ResumeAccepting(ptr.Pointer())
	}
}

func (ptr *QTcpServer) ServerError() QAbstractSocket__SocketError {
	defer qt.Recovering("QTcpServer::serverError")

	if ptr.Pointer() != nil {
		return QAbstractSocket__SocketError(C.QTcpServer_ServerError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTcpServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QTcpServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QTcpServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

func (ptr *QTcpServer) SetProxy(networkProxy QNetworkProxy_ITF) {
	defer qt.Recovering("QTcpServer::setProxy")

	if ptr.Pointer() != nil {
		C.QTcpServer_SetProxy(ptr.Pointer(), PointerFromQNetworkProxy(networkProxy))
	}
}

func (ptr *QTcpServer) WaitForNewConnection(msec int, timedOut bool) bool {
	defer qt.Recovering("QTcpServer::waitForNewConnection")

	if ptr.Pointer() != nil {
		return C.QTcpServer_WaitForNewConnection(ptr.Pointer(), C.int(msec), C.int(qt.GoBoolToInt(timedOut))) != 0
	}
	return false
}

func (ptr *QTcpServer) DestroyQTcpServer() {
	defer qt.Recovering("QTcpServer::~QTcpServer")

	if ptr.Pointer() != nil {
		C.QTcpServer_DestroyQTcpServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTcpServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTcpServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTcpServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTcpServerTimerEvent
func callbackQTcpServerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTcpServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTcpServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTcpServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTcpServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTcpServerChildEvent
func callbackQTcpServerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTcpServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpServer::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpServer::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTcpServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTcpServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTcpServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTcpServerCustomEvent
func callbackQTcpServerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTcpServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTcpServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpServer::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTcpServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpServer::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QTcpSocket struct {
	QAbstractSocket
}

type QTcpSocket_ITF interface {
	QAbstractSocket_ITF
	QTcpSocket_PTR() *QTcpSocket
}

func PointerFromQTcpSocket(ptr QTcpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func NewQTcpSocketFromPointer(ptr unsafe.Pointer) *QTcpSocket {
	var n = new(QTcpSocket)
	n.SetPointer(ptr)
	return n
}

func newQTcpSocketFromPointer(ptr unsafe.Pointer) *QTcpSocket {
	var n = NewQTcpSocketFromPointer(ptr)
	for len(n.ObjectName()) < len("QTcpSocket_") {
		n.SetObjectName("QTcpSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QTcpSocket) QTcpSocket_PTR() *QTcpSocket {
	return ptr
}

func NewQTcpSocket(parent core.QObject_ITF) *QTcpSocket {
	defer qt.Recovering("QTcpSocket::QTcpSocket")

	return newQTcpSocketFromPointer(C.QTcpSocket_NewQTcpSocket(core.PointerFromQObject(parent)))
}

func (ptr *QTcpSocket) DestroyQTcpSocket() {
	defer qt.Recovering("QTcpSocket::~QTcpSocket")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DestroyQTcpSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QTcpSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QTcpSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QTcpSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QTcpSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQTcpSocketClose
func callbackQTcpSocketClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTcpSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QTcpSocket) Close() {
	defer qt.Recovering("QTcpSocket::close")

	if ptr.Pointer() != nil {
		C.QTcpSocket_Close(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) CloseDefault() {
	defer qt.Recovering("QTcpSocket::close")

	if ptr.Pointer() != nil {
		C.QTcpSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectFromHost", f)
	}
}

func (ptr *QTcpSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectFromHost")
	}
}

//export callbackQTcpSocketDisconnectFromHost
func callbackQTcpSocketDisconnectFromHost(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTcpSocket::disconnectFromHost")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QTcpSocket) DisconnectFromHost() {
	defer qt.Recovering("QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QTcpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QTcpSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QTcpSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resume", f)
	}
}

func (ptr *QTcpSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QTcpSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resume")
	}
}

//export callbackQTcpSocketResume
func callbackQTcpSocketResume(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QTcpSocket::resume")

	if signal := qt.GetSignal(C.GoString(ptrName), "resume"); signal != nil {
		signal.(func())()
	} else {
		NewQTcpSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QTcpSocket) Resume() {
	defer qt.Recovering("QTcpSocket::resume")

	if ptr.Pointer() != nil {
		C.QTcpSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ResumeDefault() {
	defer qt.Recovering("QTcpSocket::resume")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QTcpSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setReadBufferSize", f)
	}
}

func (ptr *QTcpSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setReadBufferSize")
	}
}

//export callbackQTcpSocketSetReadBufferSize
func callbackQTcpSocketSetReadBufferSize(ptr unsafe.Pointer, ptrName *C.char, size C.longlong) {
	defer qt.Recovering("callback QTcpSocket::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQTcpSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QTcpSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QTcpSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QTcpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QTcpSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSocketOption", f)
	}
}

func (ptr *QTcpSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSocketOption")
	}
}

//export callbackQTcpSocketSetSocketOption
func callbackQTcpSocketSetSocketOption(ptr unsafe.Pointer, ptrName *C.char, option C.int, value unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::setSocketOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQTcpSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QTcpSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetSocketOption(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QTcpSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QTcpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QTcpSocket_SetSocketOptionDefault(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QTcpSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTcpSocketTimerEvent
func callbackQTcpSocketTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QTcpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QTcpSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QTcpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QTcpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQTcpSocketChildEvent
func callbackQTcpSocketChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QTcpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QTcpSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QTcpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QTcpSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QTcpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQTcpSocketCustomEvent
func callbackQTcpSocketCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QTcpSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTcpSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTcpSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QTcpSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QTcpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QTcpSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QUdpSocket struct {
	QAbstractSocket
}

type QUdpSocket_ITF interface {
	QAbstractSocket_ITF
	QUdpSocket_PTR() *QUdpSocket
}

func PointerFromQUdpSocket(ptr QUdpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUdpSocket_PTR().Pointer()
	}
	return nil
}

func NewQUdpSocketFromPointer(ptr unsafe.Pointer) *QUdpSocket {
	var n = new(QUdpSocket)
	n.SetPointer(ptr)
	return n
}

func newQUdpSocketFromPointer(ptr unsafe.Pointer) *QUdpSocket {
	var n = NewQUdpSocketFromPointer(ptr)
	for len(n.ObjectName()) < len("QUdpSocket_") {
		n.SetObjectName("QUdpSocket_" + qt.Identifier())
	}
	return n
}

func (ptr *QUdpSocket) QUdpSocket_PTR() *QUdpSocket {
	return ptr
}

func NewQUdpSocket(parent core.QObject_ITF) *QUdpSocket {
	defer qt.Recovering("QUdpSocket::QUdpSocket")

	return newQUdpSocketFromPointer(C.QUdpSocket_NewQUdpSocket(core.PointerFromQObject(parent)))
}

func (ptr *QUdpSocket) HasPendingDatagrams() bool {
	defer qt.Recovering("QUdpSocket::hasPendingDatagrams")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_HasPendingDatagrams(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddress_ITF) bool {
	defer qt.Recovering("QUdpSocket::joinMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) JoinMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	defer qt.Recovering("QUdpSocket::joinMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_JoinMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddress_ITF) bool {
	defer qt.Recovering("QUdpSocket::leaveMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup(ptr.Pointer(), PointerFromQHostAddress(groupAddress)) != 0
	}
	return false
}

func (ptr *QUdpSocket) LeaveMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {
	defer qt.Recovering("QUdpSocket::leaveMulticastGroup")

	if ptr.Pointer() != nil {
		return C.QUdpSocket_LeaveMulticastGroup2(ptr.Pointer(), PointerFromQHostAddress(groupAddress), PointerFromQNetworkInterface(iface)) != 0
	}
	return false
}

func (ptr *QUdpSocket) PendingDatagramSize() int64 {
	defer qt.Recovering("QUdpSocket::pendingDatagramSize")

	if ptr.Pointer() != nil {
		return int64(C.QUdpSocket_PendingDatagramSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QUdpSocket) SetMulticastInterface(iface QNetworkInterface_ITF) {
	defer qt.Recovering("QUdpSocket::setMulticastInterface")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetMulticastInterface(ptr.Pointer(), PointerFromQNetworkInterface(iface))
	}
}

func (ptr *QUdpSocket) DestroyQUdpSocket() {
	defer qt.Recovering("QUdpSocket::~QUdpSocket")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DestroyQUdpSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QUdpSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QUdpSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QUdpSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QUdpSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQUdpSocketClose
func callbackQUdpSocketClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QUdpSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QUdpSocket) Close() {
	defer qt.Recovering("QUdpSocket::close")

	if ptr.Pointer() != nil {
		C.QUdpSocket_Close(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) CloseDefault() {
	defer qt.Recovering("QUdpSocket::close")

	if ptr.Pointer() != nil {
		C.QUdpSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ConnectDisconnectFromHost(f func()) {
	defer qt.Recovering("connect QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectFromHost", f)
	}
}

func (ptr *QUdpSocket) DisconnectDisconnectFromHost() {
	defer qt.Recovering("disconnect QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectFromHost")
	}
}

//export callbackQUdpSocketDisconnectFromHost
func callbackQUdpSocketDisconnectFromHost(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QUdpSocket::disconnectFromHost")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectFromHost"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).DisconnectFromHostDefault()
	}
}

func (ptr *QUdpSocket) DisconnectFromHost() {
	defer qt.Recovering("QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectFromHost(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) DisconnectFromHostDefault() {
	defer qt.Recovering("QUdpSocket::disconnectFromHost")

	if ptr.Pointer() != nil {
		C.QUdpSocket_DisconnectFromHostDefault(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ConnectResume(f func()) {
	defer qt.Recovering("connect QUdpSocket::resume")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resume", f)
	}
}

func (ptr *QUdpSocket) DisconnectResume() {
	defer qt.Recovering("disconnect QUdpSocket::resume")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resume")
	}
}

//export callbackQUdpSocketResume
func callbackQUdpSocketResume(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QUdpSocket::resume")

	if signal := qt.GetSignal(C.GoString(ptrName), "resume"); signal != nil {
		signal.(func())()
	} else {
		NewQUdpSocketFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QUdpSocket) Resume() {
	defer qt.Recovering("QUdpSocket::resume")

	if ptr.Pointer() != nil {
		C.QUdpSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ResumeDefault() {
	defer qt.Recovering("QUdpSocket::resume")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ResumeDefault(ptr.Pointer())
	}
}

func (ptr *QUdpSocket) ConnectSetReadBufferSize(f func(size int64)) {
	defer qt.Recovering("connect QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setReadBufferSize", f)
	}
}

func (ptr *QUdpSocket) DisconnectSetReadBufferSize() {
	defer qt.Recovering("disconnect QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setReadBufferSize")
	}
}

//export callbackQUdpSocketSetReadBufferSize
func callbackQUdpSocketSetReadBufferSize(ptr unsafe.Pointer, ptrName *C.char, size C.longlong) {
	defer qt.Recovering("callback QUdpSocket::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
	} else {
		NewQUdpSocketFromPointer(ptr).SetReadBufferSizeDefault(int64(size))
	}
}

func (ptr *QUdpSocket) SetReadBufferSize(size int64) {
	defer qt.Recovering("QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetReadBufferSize(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QUdpSocket) SetReadBufferSizeDefault(size int64) {
	defer qt.Recovering("QUdpSocket::setReadBufferSize")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetReadBufferSizeDefault(ptr.Pointer(), C.longlong(size))
	}
}

func (ptr *QUdpSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {
	defer qt.Recovering("connect QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSocketOption", f)
	}
}

func (ptr *QUdpSocket) DisconnectSetSocketOption() {
	defer qt.Recovering("disconnect QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSocketOption")
	}
}

//export callbackQUdpSocketSetSocketOption
func callbackQUdpSocketSetSocketOption(ptr unsafe.Pointer, ptrName *C.char, option C.int, value unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::setSocketOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	} else {
		NewQUdpSocketFromPointer(ptr).SetSocketOptionDefault(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QUdpSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetSocketOption(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QUdpSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer qt.Recovering("QUdpSocket::setSocketOption")

	if ptr.Pointer() != nil {
		C.QUdpSocket_SetSocketOptionDefault(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QUdpSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQUdpSocketTimerEvent
func callbackQUdpSocketTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QUdpSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QUdpSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QUdpSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QUdpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QUdpSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQUdpSocketChildEvent
func callbackQUdpSocketChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUdpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUdpSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QUdpSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QUdpSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QUdpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QUdpSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QUdpSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQUdpSocketCustomEvent
func callbackQUdpSocketCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QUdpSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQUdpSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QUdpSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QUdpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QUdpSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QUdpSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QUdpSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
