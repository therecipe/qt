// +build !minimal

package network

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
	"unsafe"
)

type QAbstractNetworkCache struct {
	core.QObject
}

type QAbstractNetworkCache_ITF interface {
	core.QObject_ITF
	QAbstractNetworkCache_PTR() *QAbstractNetworkCache
}

func (ptr *QAbstractNetworkCache) QAbstractNetworkCache_PTR() *QAbstractNetworkCache {
	return ptr
}

func (ptr *QAbstractNetworkCache) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractNetworkCache) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractNetworkCache(ptr QAbstractNetworkCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNetworkCache_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractNetworkCache) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractNetworkCache) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQAbstractNetworkCacheFromPointer(ptr unsafe.Pointer) (n *QAbstractNetworkCache) {
	n = new(QAbstractNetworkCache)
	n.InitFromInternal(uintptr(ptr), "network.QAbstractNetworkCache")
	return
}
func NewQAbstractNetworkCache(parent core.QObject_ITF) *QAbstractNetworkCache {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQAbstractNetworkCache", "", parent}).(*QAbstractNetworkCache)
}

func (ptr *QAbstractNetworkCache) ConnectCacheSize(f func() int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCacheSize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractNetworkCache) DisconnectCacheSize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCacheSize"})
}

func (ptr *QAbstractNetworkCache) CacheSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CacheSize"}).(float64))
}

func (ptr *QAbstractNetworkCache) ConnectClear(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClear", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractNetworkCache) DisconnectClear() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClear"})
}

func (ptr *QAbstractNetworkCache) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QAbstractNetworkCache) ConnectData(f func(url *core.QUrl) *core.QIODevice) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractNetworkCache) DisconnectData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectData"})
}

func (ptr *QAbstractNetworkCache) Data(url core.QUrl_ITF) *core.QIODevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data", url}).(*core.QIODevice)
}

func (ptr *QAbstractNetworkCache) ConnectInsert(f func(device *core.QIODevice)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInsert", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractNetworkCache) DisconnectInsert() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInsert"})
}

func (ptr *QAbstractNetworkCache) Insert(device core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", device})
}

func (ptr *QAbstractNetworkCache) ConnectMetaData(f func(url *core.QUrl) *QNetworkCacheMetaData) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMetaData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractNetworkCache) DisconnectMetaData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMetaData"})
}

func (ptr *QAbstractNetworkCache) MetaData(url core.QUrl_ITF) *QNetworkCacheMetaData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaData", url}).(*QNetworkCacheMetaData)
}

func (ptr *QAbstractNetworkCache) ConnectPrepare(f func(metaData *QNetworkCacheMetaData) *core.QIODevice) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrepare", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractNetworkCache) DisconnectPrepare() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrepare"})
}

func (ptr *QAbstractNetworkCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prepare", metaData}).(*core.QIODevice)
}

func (ptr *QAbstractNetworkCache) ConnectRemove(f func(url *core.QUrl) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRemove", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractNetworkCache) DisconnectRemove() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemove"})
}

func (ptr *QAbstractNetworkCache) Remove(url core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", url}).(bool)
}

func (ptr *QAbstractNetworkCache) ConnectUpdateMetaData(f func(metaData *QNetworkCacheMetaData)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateMetaData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractNetworkCache) DisconnectUpdateMetaData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateMetaData"})
}

func (ptr *QAbstractNetworkCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMetaData", metaData})
}

func (ptr *QAbstractNetworkCache) ConnectDestroyQAbstractNetworkCache(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractNetworkCache", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractNetworkCache) DisconnectDestroyQAbstractNetworkCache() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractNetworkCache"})
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCache() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractNetworkCache"})
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCacheDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractNetworkCacheDefault"})
}

func (ptr *QAbstractNetworkCache) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QAbstractNetworkCache) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QAbstractNetworkCache) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractNetworkCache) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QAbstractNetworkCache) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QAbstractNetworkCache) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractNetworkCache) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QAbstractNetworkCache) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QAbstractNetworkCache) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractNetworkCache) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QAbstractNetworkCache) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QAbstractNetworkCache) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QAbstractNetworkCache) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QAbstractNetworkCache) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QAbstractNetworkCache) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QAbstractNetworkCache) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QAbstractNetworkCache) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QAbstractNetworkCache) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QAbstractNetworkCache) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QAbstractNetworkCache) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QAbstractNetworkCache) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QAbstractSocket struct {
	core.QIODevice
}

type QAbstractSocket_ITF interface {
	core.QIODevice_ITF
	QAbstractSocket_PTR() *QAbstractSocket
}

func (ptr *QAbstractSocket) QAbstractSocket_PTR() *QAbstractSocket {
	return ptr
}

func (ptr *QAbstractSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QAbstractSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQAbstractSocket(ptr QAbstractSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func (n *QAbstractSocket) InitFromInternal(ptr uintptr, name string) {
	n.QIODevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QAbstractSocket) ClassNameInternalF() string {
	return n.QIODevice_PTR().ClassNameInternalF()
}

func NewQAbstractSocketFromPointer(ptr unsafe.Pointer) (n *QAbstractSocket) {
	n = new(QAbstractSocket)
	n.InitFromInternal(uintptr(ptr), "network.QAbstractSocket")
	return
}

//go:generate stringer -type=QAbstractSocket__SocketType
//QAbstractSocket::SocketType
type QAbstractSocket__SocketType int64

const (
	QAbstractSocket__TcpSocket         QAbstractSocket__SocketType = QAbstractSocket__SocketType(0)
	QAbstractSocket__UdpSocket         QAbstractSocket__SocketType = QAbstractSocket__SocketType(1)
	QAbstractSocket__SctpSocket        QAbstractSocket__SocketType = QAbstractSocket__SocketType(2)
	QAbstractSocket__UnknownSocketType QAbstractSocket__SocketType = QAbstractSocket__SocketType(-1)
)

//go:generate stringer -type=QAbstractSocket__NetworkLayerProtocol
//QAbstractSocket::NetworkLayerProtocol
type QAbstractSocket__NetworkLayerProtocol int64

const (
	QAbstractSocket__IPv4Protocol                QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(0)
	QAbstractSocket__IPv6Protocol                QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(1)
	QAbstractSocket__AnyIPProtocol               QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(2)
	QAbstractSocket__UnknownNetworkLayerProtocol QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(-1)
)

//go:generate stringer -type=QAbstractSocket__SocketError
//QAbstractSocket::SocketError
type QAbstractSocket__SocketError int64

const (
	QAbstractSocket__ConnectionRefusedError           QAbstractSocket__SocketError = QAbstractSocket__SocketError(0)
	QAbstractSocket__RemoteHostClosedError            QAbstractSocket__SocketError = QAbstractSocket__SocketError(1)
	QAbstractSocket__HostNotFoundError                QAbstractSocket__SocketError = QAbstractSocket__SocketError(2)
	QAbstractSocket__SocketAccessError                QAbstractSocket__SocketError = QAbstractSocket__SocketError(3)
	QAbstractSocket__SocketResourceError              QAbstractSocket__SocketError = QAbstractSocket__SocketError(4)
	QAbstractSocket__SocketTimeoutError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(5)
	QAbstractSocket__DatagramTooLargeError            QAbstractSocket__SocketError = QAbstractSocket__SocketError(6)
	QAbstractSocket__NetworkError                     QAbstractSocket__SocketError = QAbstractSocket__SocketError(7)
	QAbstractSocket__AddressInUseError                QAbstractSocket__SocketError = QAbstractSocket__SocketError(8)
	QAbstractSocket__SocketAddressNotAvailableError   QAbstractSocket__SocketError = QAbstractSocket__SocketError(9)
	QAbstractSocket__UnsupportedSocketOperationError  QAbstractSocket__SocketError = QAbstractSocket__SocketError(10)
	QAbstractSocket__UnfinishedSocketOperationError   QAbstractSocket__SocketError = QAbstractSocket__SocketError(11)
	QAbstractSocket__ProxyAuthenticationRequiredError QAbstractSocket__SocketError = QAbstractSocket__SocketError(12)
	QAbstractSocket__SslHandshakeFailedError          QAbstractSocket__SocketError = QAbstractSocket__SocketError(13)
	QAbstractSocket__ProxyConnectionRefusedError      QAbstractSocket__SocketError = QAbstractSocket__SocketError(14)
	QAbstractSocket__ProxyConnectionClosedError       QAbstractSocket__SocketError = QAbstractSocket__SocketError(15)
	QAbstractSocket__ProxyConnectionTimeoutError      QAbstractSocket__SocketError = QAbstractSocket__SocketError(16)
	QAbstractSocket__ProxyNotFoundError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(17)
	QAbstractSocket__ProxyProtocolError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(18)
	QAbstractSocket__OperationError                   QAbstractSocket__SocketError = QAbstractSocket__SocketError(19)
	QAbstractSocket__SslInternalError                 QAbstractSocket__SocketError = QAbstractSocket__SocketError(20)
	QAbstractSocket__SslInvalidUserDataError          QAbstractSocket__SocketError = QAbstractSocket__SocketError(21)
	QAbstractSocket__TemporaryError                   QAbstractSocket__SocketError = QAbstractSocket__SocketError(22)
	QAbstractSocket__UnknownSocketError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(-1)
)

//go:generate stringer -type=QAbstractSocket__SocketState
//QAbstractSocket::SocketState
type QAbstractSocket__SocketState int64

const (
	QAbstractSocket__UnconnectedState QAbstractSocket__SocketState = QAbstractSocket__SocketState(0)
	QAbstractSocket__HostLookupState  QAbstractSocket__SocketState = QAbstractSocket__SocketState(1)
	QAbstractSocket__ConnectingState  QAbstractSocket__SocketState = QAbstractSocket__SocketState(2)
	QAbstractSocket__ConnectedState   QAbstractSocket__SocketState = QAbstractSocket__SocketState(3)
	QAbstractSocket__BoundState       QAbstractSocket__SocketState = QAbstractSocket__SocketState(4)
	QAbstractSocket__ListeningState   QAbstractSocket__SocketState = QAbstractSocket__SocketState(5)
	QAbstractSocket__ClosingState     QAbstractSocket__SocketState = QAbstractSocket__SocketState(6)
)

//go:generate stringer -type=QAbstractSocket__SocketOption
//QAbstractSocket::SocketOption
type QAbstractSocket__SocketOption int64

const (
	QAbstractSocket__LowDelayOption                QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(0)
	QAbstractSocket__KeepAliveOption               QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(1)
	QAbstractSocket__MulticastTtlOption            QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(2)
	QAbstractSocket__MulticastLoopbackOption       QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(3)
	QAbstractSocket__TypeOfServiceOption           QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(4)
	QAbstractSocket__SendBufferSizeSocketOption    QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(5)
	QAbstractSocket__ReceiveBufferSizeSocketOption QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(6)
	QAbstractSocket__PathMtuSocketOption           QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(7)
)

//go:generate stringer -type=QAbstractSocket__BindFlag
//QAbstractSocket::BindFlag
type QAbstractSocket__BindFlag int64

const (
	QAbstractSocket__DefaultForPlatform QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x0)
	QAbstractSocket__ShareAddress       QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x1)
	QAbstractSocket__DontShareAddress   QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x2)
	QAbstractSocket__ReuseAddressHint   QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x4)
)

//go:generate stringer -type=QAbstractSocket__PauseMode
//QAbstractSocket::PauseMode
type QAbstractSocket__PauseMode int64

const (
	QAbstractSocket__PauseNever       QAbstractSocket__PauseMode = QAbstractSocket__PauseMode(0x0)
	QAbstractSocket__PauseOnSslErrors QAbstractSocket__PauseMode = QAbstractSocket__PauseMode(0x1)
)

func NewQAbstractSocket(socketType QAbstractSocket__SocketType, parent core.QObject_ITF) *QAbstractSocket {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQAbstractSocket", "", socketType, parent}).(*QAbstractSocket)
}

func (ptr *QAbstractSocket) Abort() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"})
}

func (ptr *QAbstractSocket) AtEndDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtEndDefault"}).(bool)
}

func (ptr *QAbstractSocket) Bind(address QHostAddress_ITF, port uint16, mode QAbstractSocket__BindFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Bind", address, port, mode}).(bool)
}

func (ptr *QAbstractSocket) Bind2(port uint16, mode QAbstractSocket__BindFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Bind2", port, mode}).(bool)
}

func (ptr *QAbstractSocket) BytesAvailableDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesAvailableDefault"}).(float64))
}

func (ptr *QAbstractSocket) BytesToWriteDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesToWriteDefault"}).(float64))
}

func (ptr *QAbstractSocket) CanReadLineDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanReadLineDefault"}).(bool)
}

func (ptr *QAbstractSocket) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QAbstractSocket) ConnectConnectToHost(f func(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protoc QAbstractSocket__NetworkLayerProtocol)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnectToHost", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectConnectToHost() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnectToHost"})
}

func (ptr *QAbstractSocket) ConnectToHost(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protoc QAbstractSocket__NetworkLayerProtocol) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToHost", hostName, port, openMode, protoc})
}

func (ptr *QAbstractSocket) ConnectToHostDefault(hostName string, port uint16, openMode core.QIODevice__OpenModeFlag, protoc QAbstractSocket__NetworkLayerProtocol) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToHostDefault", hostName, port, openMode, protoc})
}

func (ptr *QAbstractSocket) ConnectConnectToHost2(f func(address *QHostAddress, port uint16, openMode core.QIODevice__OpenModeFlag)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnectToHost2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectConnectToHost2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnectToHost2"})
}

func (ptr *QAbstractSocket) ConnectToHost2(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToHost2", address, port, openMode})
}

func (ptr *QAbstractSocket) ConnectToHost2Default(address QHostAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToHost2Default", address, port, openMode})
}

func (ptr *QAbstractSocket) ConnectConnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectConnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnected"})
}

func (ptr *QAbstractSocket) Connected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connected"})
}

func (ptr *QAbstractSocket) ConnectDisconnectFromHost(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDisconnectFromHost", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectDisconnectFromHost() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDisconnectFromHost"})
}

func (ptr *QAbstractSocket) DisconnectFromHost() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFromHost"})
}

func (ptr *QAbstractSocket) DisconnectFromHostDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFromHostDefault"})
}

func (ptr *QAbstractSocket) ConnectDisconnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDisconnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectDisconnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDisconnected"})
}

func (ptr *QAbstractSocket) Disconnected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnected"})
}

func (ptr *QAbstractSocket) Error() QAbstractSocket__SocketError {

	return QAbstractSocket__SocketError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QAbstractSocket) ConnectError2(f func(socketError QAbstractSocket__SocketError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QAbstractSocket) Error2(socketError QAbstractSocket__SocketError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", socketError})
}

func (ptr *QAbstractSocket) Flush() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flush"}).(bool)
}

func (ptr *QAbstractSocket) ConnectHostFound(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHostFound", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectHostFound() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHostFound"})
}

func (ptr *QAbstractSocket) HostFound() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HostFound"})
}

func (ptr *QAbstractSocket) IsSequentialDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSequentialDefault"}).(bool)
}

func (ptr *QAbstractSocket) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QAbstractSocket) LocalAddress() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalAddress"}).(*QHostAddress)
}

func (ptr *QAbstractSocket) LocalPort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalPort"}).(float64))
}

func (ptr *QAbstractSocket) PauseMode() QAbstractSocket__PauseMode {

	return QAbstractSocket__PauseMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PauseMode"}).(float64))
}

func (ptr *QAbstractSocket) PeerAddress() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerAddress"}).(*QHostAddress)
}

func (ptr *QAbstractSocket) PeerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerName"}).(string)
}

func (ptr *QAbstractSocket) PeerPort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerPort"}).(float64))
}

func (ptr *QAbstractSocket) ProtocolTag() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProtocolTag"}).(string)
}

func (ptr *QAbstractSocket) Proxy() *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Proxy"}).(*QNetworkProxy)
}

func (ptr *QAbstractSocket) ConnectProxyAuthenticationRequired(f func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProxyAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectProxyAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProxyAuthenticationRequired"})
}

func (ptr *QAbstractSocket) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProxyAuthenticationRequired", proxy, authenticator})
}

func (ptr *QAbstractSocket) ReadBufferSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadBufferSize"}).(float64))
}

func (ptr *QAbstractSocket) ConnectReadData(f func(data *string, maxSize int64) int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReadData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectReadData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReadData"})
}

func (ptr *QAbstractSocket) ReadData(data *string, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadData", data, maxSize}).(float64))
}

func (ptr *QAbstractSocket) ReadDataDefault(data *string, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadDataDefault", data, maxSize}).(float64))
}

func (ptr *QAbstractSocket) ReadLineDataDefault(data []byte, maxlen int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadLineDataDefault", data, maxlen}).(float64))
}

func (ptr *QAbstractSocket) ConnectResume(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectResume", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectResume() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectResume"})
}

func (ptr *QAbstractSocket) Resume() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Resume"})
}

func (ptr *QAbstractSocket) ResumeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResumeDefault"})
}

func (ptr *QAbstractSocket) SetLocalAddress(address QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalAddress", address})
}

func (ptr *QAbstractSocket) SetLocalPort(port uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalPort", port})
}

func (ptr *QAbstractSocket) SetPauseMode(pauseMode QAbstractSocket__PauseMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPauseMode", pauseMode})
}

func (ptr *QAbstractSocket) SetPeerAddress(address QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerAddress", address})
}

func (ptr *QAbstractSocket) SetPeerName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerName", name})
}

func (ptr *QAbstractSocket) SetPeerPort(port uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerPort", port})
}

func (ptr *QAbstractSocket) SetProtocolTag(tag string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProtocolTag", tag})
}

func (ptr *QAbstractSocket) SetProxy(networkProxy QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProxy", networkProxy})
}

func (ptr *QAbstractSocket) ConnectSetReadBufferSize(f func(size int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetReadBufferSize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectSetReadBufferSize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetReadBufferSize"})
}

func (ptr *QAbstractSocket) SetReadBufferSize(size int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadBufferSize", size})
}

func (ptr *QAbstractSocket) SetReadBufferSizeDefault(size int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadBufferSizeDefault", size})
}

func (ptr *QAbstractSocket) SetSocketError(socketError QAbstractSocket__SocketError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSocketError", socketError})
}

func (ptr *QAbstractSocket) ConnectSetSocketOption(f func(option QAbstractSocket__SocketOption, value *core.QVariant)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetSocketOption", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectSetSocketOption() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetSocketOption"})
}

func (ptr *QAbstractSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSocketOption", option, value})
}

func (ptr *QAbstractSocket) SetSocketOptionDefault(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSocketOptionDefault", option, value})
}

func (ptr *QAbstractSocket) SetSocketState(state QAbstractSocket__SocketState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSocketState", state})
}

func (ptr *QAbstractSocket) ConnectSocketOption(f func(option QAbstractSocket__SocketOption) *core.QVariant) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSocketOption", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectSocketOption() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSocketOption"})
}

func (ptr *QAbstractSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SocketOption", option}).(*core.QVariant)
}

func (ptr *QAbstractSocket) SocketOptionDefault(option QAbstractSocket__SocketOption) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SocketOptionDefault", option}).(*core.QVariant)
}

func (ptr *QAbstractSocket) SocketType() QAbstractSocket__SocketType {

	return QAbstractSocket__SocketType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SocketType"}).(float64))
}

func (ptr *QAbstractSocket) State() QAbstractSocket__SocketState {

	return QAbstractSocket__SocketState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QAbstractSocket) ConnectStateChanged(f func(socketState QAbstractSocket__SocketState)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QAbstractSocket) StateChanged(socketState QAbstractSocket__SocketState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", socketState})
}

func (ptr *QAbstractSocket) WaitForBytesWrittenDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForBytesWrittenDefault", msecs}).(bool)
}

func (ptr *QAbstractSocket) ConnectWaitForConnected(f func(msecs int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWaitForConnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectWaitForConnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWaitForConnected"})
}

func (ptr *QAbstractSocket) WaitForConnected(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForConnected", msecs}).(bool)
}

func (ptr *QAbstractSocket) WaitForConnectedDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForConnectedDefault", msecs}).(bool)
}

func (ptr *QAbstractSocket) ConnectWaitForDisconnected(f func(msecs int) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWaitForDisconnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectWaitForDisconnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWaitForDisconnected"})
}

func (ptr *QAbstractSocket) WaitForDisconnected(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForDisconnected", msecs}).(bool)
}

func (ptr *QAbstractSocket) WaitForDisconnectedDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForDisconnectedDefault", msecs}).(bool)
}

func (ptr *QAbstractSocket) WaitForReadyReadDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForReadyReadDefault", msecs}).(bool)
}

func (ptr *QAbstractSocket) ConnectWriteData(f func(data []byte, size int64) int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWriteData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectWriteData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWriteData"})
}

func (ptr *QAbstractSocket) WriteData(data []byte, size int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteData", data, size}).(float64))
}

func (ptr *QAbstractSocket) WriteDataDefault(data []byte, size int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDataDefault", data, size}).(float64))
}

func (ptr *QAbstractSocket) ConnectDestroyQAbstractSocket(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQAbstractSocket", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QAbstractSocket) DisconnectDestroyQAbstractSocket() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQAbstractSocket"})
}

func (ptr *QAbstractSocket) DestroyQAbstractSocket() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractSocket"})
}

func (ptr *QAbstractSocket) DestroyQAbstractSocketDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAbstractSocketDefault"})
}

func (ptr *QAbstractSocket) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QAbstractSocket) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QAbstractSocket) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QAbstractSocket) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QAbstractSocket) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractSocket) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QAbstractSocket) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QAbstractSocket) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QAbstractSocket) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QAbstractSocket) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QAbstractSocket) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QAbstractSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault", mode}).(bool)
}

func (ptr *QAbstractSocket) PosDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PosDefault"}).(float64))
}

func (ptr *QAbstractSocket) ResetDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetDefault"}).(bool)
}

func (ptr *QAbstractSocket) SeekDefault(pos int64) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeekDefault", pos}).(bool)
}

func (ptr *QAbstractSocket) SizeDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeDefault"}).(float64))
}

func (ptr *QAbstractSocket) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QAbstractSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QAbstractSocket) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QAbstractSocket) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QAbstractSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QAbstractSocket) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QAbstractSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QAbstractSocket) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QAbstractSocket) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QAuthenticator struct {
	internal.Internal
}

type QAuthenticator_ITF interface {
	QAuthenticator_PTR() *QAuthenticator
}

func (ptr *QAuthenticator) QAuthenticator_PTR() *QAuthenticator {
	return ptr
}

func (ptr *QAuthenticator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QAuthenticator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQAuthenticator(ptr QAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAuthenticator_PTR().Pointer()
	}
	return nil
}

func (n *QAuthenticator) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQAuthenticatorFromPointer(ptr unsafe.Pointer) (n *QAuthenticator) {
	n = new(QAuthenticator)
	n.InitFromInternal(uintptr(ptr), "network.QAuthenticator")
	return
}
func NewQAuthenticator() *QAuthenticator {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQAuthenticator", ""}).(*QAuthenticator)
}

func NewQAuthenticator2(other QAuthenticator_ITF) *QAuthenticator {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQAuthenticator2", "", other}).(*QAuthenticator)
}

func (ptr *QAuthenticator) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QAuthenticator) Option(opt string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Option", opt}).(*core.QVariant)
}

func (ptr *QAuthenticator) Options() map[string]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Options"}).(map[string]*core.QVariant)
}

func (ptr *QAuthenticator) Password() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Password"}).(string)
}

func (ptr *QAuthenticator) Realm() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Realm"}).(string)
}

func (ptr *QAuthenticator) SetOption(opt string, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOption", opt, value})
}

func (ptr *QAuthenticator) SetPassword(password string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPassword", password})
}

func (ptr *QAuthenticator) SetUser(user string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUser", user})
}

func (ptr *QAuthenticator) User() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "User"}).(string)
}

func (ptr *QAuthenticator) DestroyQAuthenticator() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQAuthenticator"})
}

func (ptr *QAuthenticator) __options_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__options_atList", v, i}).(*core.QVariant)
}

func (ptr *QAuthenticator) __options_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__options_setList", key, i})
}

func (ptr *QAuthenticator) __options_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__options_newList"}).(unsafe.Pointer)
}

func (ptr *QAuthenticator) __options_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__options_keyList"}).([]string)
}

func (ptr *QAuthenticator) ____options_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____options_keyList_atList", i}).(string)
}

func (ptr *QAuthenticator) ____options_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____options_keyList_setList", i})
}

func (ptr *QAuthenticator) ____options_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____options_keyList_newList"}).(unsafe.Pointer)
}

type QDnsDomainNameRecord struct {
	internal.Internal
}

type QDnsDomainNameRecord_ITF interface {
	QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord
}

func (ptr *QDnsDomainNameRecord) QDnsDomainNameRecord_PTR() *QDnsDomainNameRecord {
	return ptr
}

func (ptr *QDnsDomainNameRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDnsDomainNameRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDnsDomainNameRecord(ptr QDnsDomainNameRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsDomainNameRecord_PTR().Pointer()
	}
	return nil
}

func (n *QDnsDomainNameRecord) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDnsDomainNameRecordFromPointer(ptr unsafe.Pointer) (n *QDnsDomainNameRecord) {
	n = new(QDnsDomainNameRecord)
	n.InitFromInternal(uintptr(ptr), "network.QDnsDomainNameRecord")
	return
}
func NewQDnsDomainNameRecord() *QDnsDomainNameRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsDomainNameRecord", ""}).(*QDnsDomainNameRecord)
}

func NewQDnsDomainNameRecord2(other QDnsDomainNameRecord_ITF) *QDnsDomainNameRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsDomainNameRecord2", "", other}).(*QDnsDomainNameRecord)
}

func (ptr *QDnsDomainNameRecord) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDnsDomainNameRecord) Swap(other QDnsDomainNameRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDnsDomainNameRecord) TimeToLive() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimeToLive"}).(float64))
}

func (ptr *QDnsDomainNameRecord) Value() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(string)
}

func (ptr *QDnsDomainNameRecord) DestroyQDnsDomainNameRecord() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDnsDomainNameRecord"})
}

type QDnsHostAddressRecord struct {
	internal.Internal
}

type QDnsHostAddressRecord_ITF interface {
	QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord
}

func (ptr *QDnsHostAddressRecord) QDnsHostAddressRecord_PTR() *QDnsHostAddressRecord {
	return ptr
}

func (ptr *QDnsHostAddressRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDnsHostAddressRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDnsHostAddressRecord(ptr QDnsHostAddressRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsHostAddressRecord_PTR().Pointer()
	}
	return nil
}

func (n *QDnsHostAddressRecord) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDnsHostAddressRecordFromPointer(ptr unsafe.Pointer) (n *QDnsHostAddressRecord) {
	n = new(QDnsHostAddressRecord)
	n.InitFromInternal(uintptr(ptr), "network.QDnsHostAddressRecord")
	return
}
func NewQDnsHostAddressRecord() *QDnsHostAddressRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsHostAddressRecord", ""}).(*QDnsHostAddressRecord)
}

func NewQDnsHostAddressRecord2(other QDnsHostAddressRecord_ITF) *QDnsHostAddressRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsHostAddressRecord2", "", other}).(*QDnsHostAddressRecord)
}

func (ptr *QDnsHostAddressRecord) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDnsHostAddressRecord) Swap(other QDnsHostAddressRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDnsHostAddressRecord) TimeToLive() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimeToLive"}).(float64))
}

func (ptr *QDnsHostAddressRecord) Value() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(*QHostAddress)
}

func (ptr *QDnsHostAddressRecord) DestroyQDnsHostAddressRecord() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDnsHostAddressRecord"})
}

type QDnsLookup struct {
	core.QObject
}

type QDnsLookup_ITF interface {
	core.QObject_ITF
	QDnsLookup_PTR() *QDnsLookup
}

func (ptr *QDnsLookup) QDnsLookup_PTR() *QDnsLookup {
	return ptr
}

func (ptr *QDnsLookup) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDnsLookup) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDnsLookup(ptr QDnsLookup_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsLookup_PTR().Pointer()
	}
	return nil
}

func (n *QDnsLookup) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDnsLookup) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDnsLookupFromPointer(ptr unsafe.Pointer) (n *QDnsLookup) {
	n = new(QDnsLookup)
	n.InitFromInternal(uintptr(ptr), "network.QDnsLookup")
	return
}

//go:generate stringer -type=QDnsLookup__Error
//QDnsLookup::Error
type QDnsLookup__Error int64

const (
	QDnsLookup__NoError                 QDnsLookup__Error = QDnsLookup__Error(0)
	QDnsLookup__ResolverError           QDnsLookup__Error = QDnsLookup__Error(1)
	QDnsLookup__OperationCancelledError QDnsLookup__Error = QDnsLookup__Error(2)
	QDnsLookup__InvalidRequestError     QDnsLookup__Error = QDnsLookup__Error(3)
	QDnsLookup__InvalidReplyError       QDnsLookup__Error = QDnsLookup__Error(4)
	QDnsLookup__ServerFailureError      QDnsLookup__Error = QDnsLookup__Error(5)
	QDnsLookup__ServerRefusedError      QDnsLookup__Error = QDnsLookup__Error(6)
	QDnsLookup__NotFoundError           QDnsLookup__Error = QDnsLookup__Error(7)
)

//go:generate stringer -type=QDnsLookup__Type
//QDnsLookup::Type
type QDnsLookup__Type int64

const (
	QDnsLookup__A     QDnsLookup__Type = QDnsLookup__Type(1)
	QDnsLookup__AAAA  QDnsLookup__Type = QDnsLookup__Type(28)
	QDnsLookup__ANY   QDnsLookup__Type = QDnsLookup__Type(255)
	QDnsLookup__CNAME QDnsLookup__Type = QDnsLookup__Type(5)
	QDnsLookup__MX    QDnsLookup__Type = QDnsLookup__Type(15)
	QDnsLookup__NS    QDnsLookup__Type = QDnsLookup__Type(2)
	QDnsLookup__PTR   QDnsLookup__Type = QDnsLookup__Type(12)
	QDnsLookup__SRV   QDnsLookup__Type = QDnsLookup__Type(33)
	QDnsLookup__TXT   QDnsLookup__Type = QDnsLookup__Type(16)
)

func NewQDnsLookup(parent core.QObject_ITF) *QDnsLookup {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsLookup", "", parent}).(*QDnsLookup)
}

func NewQDnsLookup2(ty QDnsLookup__Type, name string, parent core.QObject_ITF) *QDnsLookup {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsLookup2", "", ty, name, parent}).(*QDnsLookup)
}

func NewQDnsLookup3(ty QDnsLookup__Type, name string, nameserver QHostAddress_ITF, parent core.QObject_ITF) *QDnsLookup {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsLookup3", "", ty, name, nameserver, parent}).(*QDnsLookup)
}

func (ptr *QDnsLookup) ConnectAbort(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAbort", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDnsLookup) DisconnectAbort() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAbort"})
}

func (ptr *QDnsLookup) Abort() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"})
}

func (ptr *QDnsLookup) AbortDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AbortDefault"})
}

func (ptr *QDnsLookup) CanonicalNameRecords() []*QDnsDomainNameRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanonicalNameRecords"}).([]*QDnsDomainNameRecord)
}

func (ptr *QDnsLookup) Error() QDnsLookup__Error {

	return QDnsLookup__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QDnsLookup) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QDnsLookup) ConnectFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDnsLookup) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QDnsLookup) Finished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished"})
}

func (ptr *QDnsLookup) HostAddressRecords() []*QDnsHostAddressRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HostAddressRecords"}).([]*QDnsHostAddressRecord)
}

func (ptr *QDnsLookup) IsFinished() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFinished"}).(bool)
}

func (ptr *QDnsLookup) ConnectLookup(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLookup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDnsLookup) DisconnectLookup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLookup"})
}

func (ptr *QDnsLookup) Lookup() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Lookup"})
}

func (ptr *QDnsLookup) LookupDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LookupDefault"})
}

func (ptr *QDnsLookup) MailExchangeRecords() []*QDnsMailExchangeRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MailExchangeRecords"}).([]*QDnsMailExchangeRecord)
}

func (ptr *QDnsLookup) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDnsLookup) ConnectNameChanged(f func(name string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNameChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDnsLookup) DisconnectNameChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNameChanged"})
}

func (ptr *QDnsLookup) NameChanged(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameChanged", name})
}

func (ptr *QDnsLookup) NameServerRecords() []*QDnsDomainNameRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameServerRecords"}).([]*QDnsDomainNameRecord)
}

func (ptr *QDnsLookup) Nameserver() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Nameserver"}).(*QHostAddress)
}

func (ptr *QDnsLookup) ConnectNameserverChanged(f func(nameserver *QHostAddress)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNameserverChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDnsLookup) DisconnectNameserverChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNameserverChanged"})
}

func (ptr *QDnsLookup) NameserverChanged(nameserver QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NameserverChanged", nameserver})
}

func (ptr *QDnsLookup) PointerRecords() []*QDnsDomainNameRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PointerRecords"}).([]*QDnsDomainNameRecord)
}

func (ptr *QDnsLookup) ServiceRecords() []*QDnsServiceRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServiceRecords"}).([]*QDnsServiceRecord)
}

func (ptr *QDnsLookup) SetName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", name})
}

func (ptr *QDnsLookup) SetNameserver(nameserver QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNameserver", nameserver})
}

func (ptr *QDnsLookup) SetType(vqd QDnsLookup__Type) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetType", vqd})
}

func (ptr *QDnsLookup) TextRecords() []*QDnsTextRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextRecords"}).([]*QDnsTextRecord)
}

func (ptr *QDnsLookup) Type() QDnsLookup__Type {

	return QDnsLookup__Type(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QDnsLookup) ConnectTypeChanged(f func(ty QDnsLookup__Type)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTypeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDnsLookup) DisconnectTypeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTypeChanged"})
}

func (ptr *QDnsLookup) TypeChanged(ty QDnsLookup__Type) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TypeChanged", ty})
}

func (ptr *QDnsLookup) ConnectDestroyQDnsLookup(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQDnsLookup", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QDnsLookup) DisconnectDestroyQDnsLookup() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQDnsLookup"})
}

func (ptr *QDnsLookup) DestroyQDnsLookup() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDnsLookup"})
}

func (ptr *QDnsLookup) DestroyQDnsLookupDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDnsLookupDefault"})
}

func (ptr *QDnsLookup) __canonicalNameRecords_atList(i int) *QDnsDomainNameRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__canonicalNameRecords_atList", i}).(*QDnsDomainNameRecord)
}

func (ptr *QDnsLookup) __canonicalNameRecords_setList(i QDnsDomainNameRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__canonicalNameRecords_setList", i})
}

func (ptr *QDnsLookup) __canonicalNameRecords_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__canonicalNameRecords_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __hostAddressRecords_atList(i int) *QDnsHostAddressRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__hostAddressRecords_atList", i}).(*QDnsHostAddressRecord)
}

func (ptr *QDnsLookup) __hostAddressRecords_setList(i QDnsHostAddressRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__hostAddressRecords_setList", i})
}

func (ptr *QDnsLookup) __hostAddressRecords_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__hostAddressRecords_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __mailExchangeRecords_atList(i int) *QDnsMailExchangeRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mailExchangeRecords_atList", i}).(*QDnsMailExchangeRecord)
}

func (ptr *QDnsLookup) __mailExchangeRecords_setList(i QDnsMailExchangeRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mailExchangeRecords_setList", i})
}

func (ptr *QDnsLookup) __mailExchangeRecords_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__mailExchangeRecords_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __nameServerRecords_atList(i int) *QDnsDomainNameRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__nameServerRecords_atList", i}).(*QDnsDomainNameRecord)
}

func (ptr *QDnsLookup) __nameServerRecords_setList(i QDnsDomainNameRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__nameServerRecords_setList", i})
}

func (ptr *QDnsLookup) __nameServerRecords_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__nameServerRecords_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __pointerRecords_atList(i int) *QDnsDomainNameRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__pointerRecords_atList", i}).(*QDnsDomainNameRecord)
}

func (ptr *QDnsLookup) __pointerRecords_setList(i QDnsDomainNameRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__pointerRecords_setList", i})
}

func (ptr *QDnsLookup) __pointerRecords_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__pointerRecords_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __serviceRecords_atList(i int) *QDnsServiceRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__serviceRecords_atList", i}).(*QDnsServiceRecord)
}

func (ptr *QDnsLookup) __serviceRecords_setList(i QDnsServiceRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__serviceRecords_setList", i})
}

func (ptr *QDnsLookup) __serviceRecords_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__serviceRecords_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __textRecords_atList(i int) *QDnsTextRecord {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__textRecords_atList", i}).(*QDnsTextRecord)
}

func (ptr *QDnsLookup) __textRecords_setList(i QDnsTextRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__textRecords_setList", i})
}

func (ptr *QDnsLookup) __textRecords_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__textRecords_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QDnsLookup) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QDnsLookup) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QDnsLookup) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QDnsLookup) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QDnsLookup) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QDnsLookup) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QDnsLookup) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QDnsLookup) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QDnsLookup) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QDnsLookup) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QDnsLookup) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QDnsLookup) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QDnsLookup) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QDnsLookup) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QDnsLookup) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QDnsLookup) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QDnsLookup) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QDnsMailExchangeRecord struct {
	internal.Internal
}

type QDnsMailExchangeRecord_ITF interface {
	QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord
}

func (ptr *QDnsMailExchangeRecord) QDnsMailExchangeRecord_PTR() *QDnsMailExchangeRecord {
	return ptr
}

func (ptr *QDnsMailExchangeRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDnsMailExchangeRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDnsMailExchangeRecord(ptr QDnsMailExchangeRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsMailExchangeRecord_PTR().Pointer()
	}
	return nil
}

func (n *QDnsMailExchangeRecord) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDnsMailExchangeRecordFromPointer(ptr unsafe.Pointer) (n *QDnsMailExchangeRecord) {
	n = new(QDnsMailExchangeRecord)
	n.InitFromInternal(uintptr(ptr), "network.QDnsMailExchangeRecord")
	return
}
func NewQDnsMailExchangeRecord() *QDnsMailExchangeRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsMailExchangeRecord", ""}).(*QDnsMailExchangeRecord)
}

func NewQDnsMailExchangeRecord2(other QDnsMailExchangeRecord_ITF) *QDnsMailExchangeRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsMailExchangeRecord2", "", other}).(*QDnsMailExchangeRecord)
}

func (ptr *QDnsMailExchangeRecord) Exchange() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Exchange"}).(string)
}

func (ptr *QDnsMailExchangeRecord) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDnsMailExchangeRecord) Preference() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Preference"}).(float64))
}

func (ptr *QDnsMailExchangeRecord) Swap(other QDnsMailExchangeRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDnsMailExchangeRecord) TimeToLive() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimeToLive"}).(float64))
}

func (ptr *QDnsMailExchangeRecord) DestroyQDnsMailExchangeRecord() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDnsMailExchangeRecord"})
}

type QDnsServiceRecord struct {
	internal.Internal
}

type QDnsServiceRecord_ITF interface {
	QDnsServiceRecord_PTR() *QDnsServiceRecord
}

func (ptr *QDnsServiceRecord) QDnsServiceRecord_PTR() *QDnsServiceRecord {
	return ptr
}

func (ptr *QDnsServiceRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDnsServiceRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDnsServiceRecord(ptr QDnsServiceRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsServiceRecord_PTR().Pointer()
	}
	return nil
}

func (n *QDnsServiceRecord) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDnsServiceRecordFromPointer(ptr unsafe.Pointer) (n *QDnsServiceRecord) {
	n = new(QDnsServiceRecord)
	n.InitFromInternal(uintptr(ptr), "network.QDnsServiceRecord")
	return
}
func NewQDnsServiceRecord() *QDnsServiceRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsServiceRecord", ""}).(*QDnsServiceRecord)
}

func NewQDnsServiceRecord2(other QDnsServiceRecord_ITF) *QDnsServiceRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsServiceRecord2", "", other}).(*QDnsServiceRecord)
}

func (ptr *QDnsServiceRecord) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDnsServiceRecord) Port() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Port"}).(float64))
}

func (ptr *QDnsServiceRecord) Priority() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Priority"}).(float64))
}

func (ptr *QDnsServiceRecord) Swap(other QDnsServiceRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDnsServiceRecord) Target() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Target"}).(string)
}

func (ptr *QDnsServiceRecord) TimeToLive() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimeToLive"}).(float64))
}

func (ptr *QDnsServiceRecord) Weight() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Weight"}).(float64))
}

func (ptr *QDnsServiceRecord) DestroyQDnsServiceRecord() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDnsServiceRecord"})
}

type QDnsTextRecord struct {
	internal.Internal
}

type QDnsTextRecord_ITF interface {
	QDnsTextRecord_PTR() *QDnsTextRecord
}

func (ptr *QDnsTextRecord) QDnsTextRecord_PTR() *QDnsTextRecord {
	return ptr
}

func (ptr *QDnsTextRecord) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QDnsTextRecord) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQDnsTextRecord(ptr QDnsTextRecord_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDnsTextRecord_PTR().Pointer()
	}
	return nil
}

func (n *QDnsTextRecord) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQDnsTextRecordFromPointer(ptr unsafe.Pointer) (n *QDnsTextRecord) {
	n = new(QDnsTextRecord)
	n.InitFromInternal(uintptr(ptr), "network.QDnsTextRecord")
	return
}
func NewQDnsTextRecord() *QDnsTextRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsTextRecord", ""}).(*QDnsTextRecord)
}

func NewQDnsTextRecord2(other QDnsTextRecord_ITF) *QDnsTextRecord {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQDnsTextRecord2", "", other}).(*QDnsTextRecord)
}

func (ptr *QDnsTextRecord) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QDnsTextRecord) Swap(other QDnsTextRecord_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QDnsTextRecord) TimeToLive() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimeToLive"}).(float64))
}

func (ptr *QDnsTextRecord) Values() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Values"}).([]*core.QByteArray)
}

func (ptr *QDnsTextRecord) DestroyQDnsTextRecord() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQDnsTextRecord"})
}

func (ptr *QDnsTextRecord) __values_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__values_atList", i}).(*core.QByteArray)
}

func (ptr *QDnsTextRecord) __values_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__values_setList", i})
}

func (ptr *QDnsTextRecord) __values_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__values_newList"}).(unsafe.Pointer)
}

type QDtls struct {
	core.QObject
}

type QDtls_ITF interface {
	core.QObject_ITF
	QDtls_PTR() *QDtls
}

func (ptr *QDtls) QDtls_PTR() *QDtls {
	return ptr
}

func (ptr *QDtls) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDtls) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDtls(ptr QDtls_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDtls_PTR().Pointer()
	}
	return nil
}

func (n *QDtls) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDtls) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDtlsFromPointer(ptr unsafe.Pointer) (n *QDtls) {
	n = new(QDtls)
	n.InitFromInternal(uintptr(ptr), "network.QDtls")
	return
}

//go:generate stringer -type=QDtls__QDtlsError
//QDtls::QDtlsError
type QDtls__QDtlsError int64

const (
	QDtls__NoError                     QDtls__QDtlsError = QDtls__QDtlsError(0)
	QDtls__InvalidInputParameters      QDtls__QDtlsError = QDtls__QDtlsError(1)
	QDtls__InvalidOperation            QDtls__QDtlsError = QDtls__QDtlsError(2)
	QDtls__UnderlyingSocketError       QDtls__QDtlsError = QDtls__QDtlsError(3)
	QDtls__RemoteClosedConnectionError QDtls__QDtlsError = QDtls__QDtlsError(4)
	QDtls__PeerVerificationError       QDtls__QDtlsError = QDtls__QDtlsError(5)
	QDtls__TlsInitializationError      QDtls__QDtlsError = QDtls__QDtlsError(6)
	QDtls__TlsFatalError               QDtls__QDtlsError = QDtls__QDtlsError(7)
	QDtls__TlsNonFatalError            QDtls__QDtlsError = QDtls__QDtlsError(8)
)

//go:generate stringer -type=QDtls__HandshakeState
//QDtls::HandshakeState
type QDtls__HandshakeState int64

const (
	QDtls__HandshakeNotStarted    QDtls__HandshakeState = QDtls__HandshakeState(0)
	QDtls__HandshakeInProgress    QDtls__HandshakeState = QDtls__HandshakeState(1)
	QDtls__PeerVerificationFailed QDtls__HandshakeState = QDtls__HandshakeState(2)
	QDtls__HandshakeComplete      QDtls__HandshakeState = QDtls__HandshakeState(3)
)

type QDtlsClientVerifier struct {
	core.QObject
}

type QDtlsClientVerifier_ITF interface {
	core.QObject_ITF
	QDtlsClientVerifier_PTR() *QDtlsClientVerifier
}

func (ptr *QDtlsClientVerifier) QDtlsClientVerifier_PTR() *QDtlsClientVerifier {
	return ptr
}

func (ptr *QDtlsClientVerifier) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QDtlsClientVerifier) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQDtlsClientVerifier(ptr QDtlsClientVerifier_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDtlsClientVerifier_PTR().Pointer()
	}
	return nil
}

func (n *QDtlsClientVerifier) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QDtlsClientVerifier) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQDtlsClientVerifierFromPointer(ptr unsafe.Pointer) (n *QDtlsClientVerifier) {
	n = new(QDtlsClientVerifier)
	n.InitFromInternal(uintptr(ptr), "network.QDtlsClientVerifier")
	return
}

type QHostAddress struct {
	internal.Internal
}

type QHostAddress_ITF interface {
	QHostAddress_PTR() *QHostAddress
}

func (ptr *QHostAddress) QHostAddress_PTR() *QHostAddress {
	return ptr
}

func (ptr *QHostAddress) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHostAddress) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHostAddress(ptr QHostAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostAddress_PTR().Pointer()
	}
	return nil
}

func (n *QHostAddress) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHostAddressFromPointer(ptr unsafe.Pointer) (n *QHostAddress) {
	n = new(QHostAddress)
	n.InitFromInternal(uintptr(ptr), "network.QHostAddress")
	return
}

//go:generate stringer -type=QHostAddress__SpecialAddress
//QHostAddress::SpecialAddress
type QHostAddress__SpecialAddress int64

const (
	QHostAddress__Null          QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(0)
	QHostAddress__Broadcast     QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(1)
	QHostAddress__LocalHost     QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(2)
	QHostAddress__LocalHostIPv6 QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(3)
	QHostAddress__Any           QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(4)
	QHostAddress__AnyIPv6       QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(5)
	QHostAddress__AnyIPv4       QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(6)
)

//go:generate stringer -type=QHostAddress__ConversionModeFlag
//QHostAddress::ConversionModeFlag
type QHostAddress__ConversionModeFlag int64

const (
	QHostAddress__ConvertV4MappedToIPv4     QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(1)
	QHostAddress__ConvertV4CompatToIPv4     QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(2)
	QHostAddress__ConvertUnspecifiedAddress QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(4)
	QHostAddress__ConvertLocalHost          QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(8)
	QHostAddress__TolerantConversion        QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(0xff)
	QHostAddress__StrictConversion          QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(0)
)

func NewQHostAddress() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHostAddress", ""}).(*QHostAddress)
}

func NewQHostAddress2(ip4Addr uint) *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHostAddress2", "", ip4Addr}).(*QHostAddress)
}

func NewQHostAddress3(ip6Addr string) *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHostAddress3", "", ip6Addr}).(*QHostAddress)
}

func NewQHostAddress4(ip6Addr string) *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHostAddress4", "", ip6Addr}).(*QHostAddress)
}

func NewQHostAddress7(address string) *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHostAddress7", "", address}).(*QHostAddress)
}

func NewQHostAddress8(address QHostAddress_ITF) *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHostAddress8", "", address}).(*QHostAddress)
}

func NewQHostAddress9(address QHostAddress__SpecialAddress) *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHostAddress9", "", address}).(*QHostAddress)
}

func (ptr *QHostAddress) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QHostAddress) IsBroadcast() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBroadcast"}).(bool)
}

func (ptr *QHostAddress) IsEqual(other QHostAddress_ITF, mode QHostAddress__ConversionModeFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEqual", other, mode}).(bool)
}

func (ptr *QHostAddress) IsGlobal() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsGlobal"}).(bool)
}

func (ptr *QHostAddress) IsInSubnet(subnet QHostAddress_ITF, netmask int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsInSubnet", subnet, netmask}).(bool)
}

func (ptr *QHostAddress) IsLinkLocal() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLinkLocal"}).(bool)
}

func (ptr *QHostAddress) IsLoopback() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLoopback"}).(bool)
}

func (ptr *QHostAddress) IsMulticast() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsMulticast"}).(bool)
}

func (ptr *QHostAddress) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QHostAddress) IsSiteLocal() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSiteLocal"}).(bool)
}

func (ptr *QHostAddress) IsUniqueLocalUnicast() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsUniqueLocalUnicast"}).(bool)
}

func (ptr *QHostAddress) Protocol() QAbstractSocket__NetworkLayerProtocol {

	return QAbstractSocket__NetworkLayerProtocol(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Protocol"}).(float64))
}

func (ptr *QHostAddress) ScopeId() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ScopeId"}).(string)
}

func (ptr *QHostAddress) SetAddress(ip4Addr uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAddress", ip4Addr})
}

func (ptr *QHostAddress) SetAddress2(ip6Addr string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAddress2", ip6Addr})
}

func (ptr *QHostAddress) SetAddress3(ip6Addr string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAddress3", ip6Addr})
}

func (ptr *QHostAddress) SetAddress6(address string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAddress6", address}).(bool)
}

func (ptr *QHostAddress) SetAddress7(address QHostAddress__SpecialAddress) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAddress7", address})
}

func (ptr *QHostAddress) SetScopeId(id string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetScopeId", id})
}

func (ptr *QHostAddress) Swap(other QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QHostAddress) ToIPv4Address() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToIPv4Address"}).(float64))
}

func (ptr *QHostAddress) ToIPv4Address2(ok *bool) uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToIPv4Address2", ok}).(float64))
}

func (ptr *QHostAddress) ToString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToString"}).(string)
}

func (ptr *QHostAddress) DestroyQHostAddress() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHostAddress"})
}

type QHostInfo struct {
	internal.Internal
}

type QHostInfo_ITF interface {
	QHostInfo_PTR() *QHostInfo
}

func (ptr *QHostInfo) QHostInfo_PTR() *QHostInfo {
	return ptr
}

func (ptr *QHostInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHostInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHostInfo(ptr QHostInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostInfo_PTR().Pointer()
	}
	return nil
}

func (n *QHostInfo) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHostInfoFromPointer(ptr unsafe.Pointer) (n *QHostInfo) {
	n = new(QHostInfo)
	n.InitFromInternal(uintptr(ptr), "network.QHostInfo")
	return
}

//go:generate stringer -type=QHostInfo__HostInfoError
//QHostInfo::HostInfoError
type QHostInfo__HostInfoError int64

const (
	QHostInfo__NoError      QHostInfo__HostInfoError = QHostInfo__HostInfoError(0)
	QHostInfo__HostNotFound QHostInfo__HostInfoError = QHostInfo__HostInfoError(1)
	QHostInfo__UnknownError QHostInfo__HostInfoError = QHostInfo__HostInfoError(2)
)

func NewQHostInfo(id int) *QHostInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHostInfo", "", id}).(*QHostInfo)
}

func NewQHostInfo2(other QHostInfo_ITF) *QHostInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHostInfo2", "", other}).(*QHostInfo)
}

func QHostInfo_AbortHostLookup(id int) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_AbortHostLookup", "", id})
}

func (ptr *QHostInfo) AbortHostLookup(id int) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_AbortHostLookup", "", id})
}

func (ptr *QHostInfo) Addresses() []*QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Addresses"}).([]*QHostAddress)
}

func (ptr *QHostInfo) Error() QHostInfo__HostInfoError {

	return QHostInfo__HostInfoError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QHostInfo) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func QHostInfo_FromName(name string) *QHostInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_FromName", "", name}).(*QHostInfo)
}

func (ptr *QHostInfo) FromName(name string) *QHostInfo {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_FromName", "", name}).(*QHostInfo)
}

func (ptr *QHostInfo) HostName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HostName"}).(string)
}

func QHostInfo_LocalDomainName() string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_LocalDomainName", ""}).(string)
}

func (ptr *QHostInfo) LocalDomainName() string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_LocalDomainName", ""}).(string)
}

func QHostInfo_LocalHostName() string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_LocalHostName", ""}).(string)
}

func (ptr *QHostInfo) LocalHostName() string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_LocalHostName", ""}).(string)
}

func QHostInfo_LookupHost(name string, receiver core.QObject_ITF, member string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_LookupHost", "", name, receiver, member}).(float64))
}

func (ptr *QHostInfo) LookupHost(name string, receiver core.QObject_ITF, member string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "network.QHostInfo_LookupHost", "", name, receiver, member}).(float64))
}

func (ptr *QHostInfo) LookupId() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LookupId"}).(float64))
}

func (ptr *QHostInfo) SetAddresses(addresses []*QHostAddress) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAddresses", addresses})
}

func (ptr *QHostInfo) SetError(error QHostInfo__HostInfoError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetError", error})
}

func (ptr *QHostInfo) SetErrorString(str string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetErrorString", str})
}

func (ptr *QHostInfo) SetHostName(hostName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHostName", hostName})
}

func (ptr *QHostInfo) SetLookupId(id int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLookupId", id})
}

func (ptr *QHostInfo) Swap(other QHostInfo_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QHostInfo) DestroyQHostInfo() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHostInfo"})
}

func (ptr *QHostInfo) __addresses_atList(i int) *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addresses_atList", i}).(*QHostAddress)
}

func (ptr *QHostInfo) __addresses_setList(i QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addresses_setList", i})
}

func (ptr *QHostInfo) __addresses_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addresses_newList"}).(unsafe.Pointer)
}

func (ptr *QHostInfo) __setAddresses_addresses_atList(i int) *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAddresses_addresses_atList", i}).(*QHostAddress)
}

func (ptr *QHostInfo) __setAddresses_addresses_setList(i QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAddresses_addresses_setList", i})
}

func (ptr *QHostInfo) __setAddresses_addresses_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAddresses_addresses_newList"}).(unsafe.Pointer)
}

type QHstsPolicy struct {
	internal.Internal
}

type QHstsPolicy_ITF interface {
	QHstsPolicy_PTR() *QHstsPolicy
}

func (ptr *QHstsPolicy) QHstsPolicy_PTR() *QHstsPolicy {
	return ptr
}

func (ptr *QHstsPolicy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHstsPolicy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHstsPolicy(ptr QHstsPolicy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHstsPolicy_PTR().Pointer()
	}
	return nil
}

func (n *QHstsPolicy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHstsPolicyFromPointer(ptr unsafe.Pointer) (n *QHstsPolicy) {
	n = new(QHstsPolicy)
	n.InitFromInternal(uintptr(ptr), "network.QHstsPolicy")
	return
}

//go:generate stringer -type=QHstsPolicy__PolicyFlag
//QHstsPolicy::PolicyFlag
type QHstsPolicy__PolicyFlag int64

const (
	QHstsPolicy__IncludeSubDomains QHstsPolicy__PolicyFlag = QHstsPolicy__PolicyFlag(1)
)

func NewQHstsPolicy() *QHstsPolicy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHstsPolicy", ""}).(*QHstsPolicy)
}

func NewQHstsPolicy2(expiry core.QDateTime_ITF, flags QHstsPolicy__PolicyFlag, host string, mode core.QUrl__ParsingMode) *QHstsPolicy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHstsPolicy2", "", expiry, flags, host, mode}).(*QHstsPolicy)
}

func NewQHstsPolicy3(other QHstsPolicy_ITF) *QHstsPolicy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHstsPolicy3", "", other}).(*QHstsPolicy)
}

func (ptr *QHstsPolicy) Expiry() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Expiry"}).(*core.QDateTime)
}

func (ptr *QHstsPolicy) Host(options core.QUrl__ComponentFormattingOption) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Host", options}).(string)
}

func (ptr *QHstsPolicy) IncludesSubDomains() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncludesSubDomains"}).(bool)
}

func (ptr *QHstsPolicy) IsExpired() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsExpired"}).(bool)
}

func (ptr *QHstsPolicy) SetExpiry(expiry core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExpiry", expiry})
}

func (ptr *QHstsPolicy) SetHost(host string, mode core.QUrl__ParsingMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHost", host, mode})
}

func (ptr *QHstsPolicy) SetIncludesSubDomains(include bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIncludesSubDomains", include})
}

func (ptr *QHstsPolicy) Swap(other QHstsPolicy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QHstsPolicy) DestroyQHstsPolicy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHstsPolicy"})
}

type QHttpMultiPart struct {
	core.QObject
}

type QHttpMultiPart_ITF interface {
	core.QObject_ITF
	QHttpMultiPart_PTR() *QHttpMultiPart
}

func (ptr *QHttpMultiPart) QHttpMultiPart_PTR() *QHttpMultiPart {
	return ptr
}

func (ptr *QHttpMultiPart) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QHttpMultiPart) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQHttpMultiPart(ptr QHttpMultiPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpMultiPart_PTR().Pointer()
	}
	return nil
}

func (n *QHttpMultiPart) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QHttpMultiPart) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQHttpMultiPartFromPointer(ptr unsafe.Pointer) (n *QHttpMultiPart) {
	n = new(QHttpMultiPart)
	n.InitFromInternal(uintptr(ptr), "network.QHttpMultiPart")
	return
}

//go:generate stringer -type=QHttpMultiPart__ContentType
//QHttpMultiPart::ContentType
type QHttpMultiPart__ContentType int64

const (
	QHttpMultiPart__MixedType       QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(0)
	QHttpMultiPart__RelatedType     QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(1)
	QHttpMultiPart__FormDataType    QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(2)
	QHttpMultiPart__AlternativeType QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(3)
)

func NewQHttpMultiPart(parent core.QObject_ITF) *QHttpMultiPart {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHttpMultiPart", "", parent}).(*QHttpMultiPart)
}

func NewQHttpMultiPart2(contentType QHttpMultiPart__ContentType, parent core.QObject_ITF) *QHttpMultiPart {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHttpMultiPart2", "", contentType, parent}).(*QHttpMultiPart)
}

func (ptr *QHttpMultiPart) Append(httpPart QHttpPart_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Append", httpPart})
}

func (ptr *QHttpMultiPart) Boundary() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Boundary"}).(*core.QByteArray)
}

func (ptr *QHttpMultiPart) SetBoundary(boundary core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBoundary", boundary})
}

func (ptr *QHttpMultiPart) SetContentType(contentType QHttpMultiPart__ContentType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetContentType", contentType})
}

func (ptr *QHttpMultiPart) ConnectDestroyQHttpMultiPart(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQHttpMultiPart", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QHttpMultiPart) DisconnectDestroyQHttpMultiPart() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQHttpMultiPart"})
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPart() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHttpMultiPart"})
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPartDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHttpMultiPartDefault"})
}

func (ptr *QHttpMultiPart) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QHttpMultiPart) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QHttpMultiPart) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QHttpMultiPart) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QHttpMultiPart) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QHttpMultiPart) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QHttpMultiPart) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QHttpMultiPart) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QHttpMultiPart) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QHttpMultiPart) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QHttpMultiPart) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QHttpMultiPart) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QHttpMultiPart) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QHttpMultiPart) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QHttpMultiPart) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QHttpMultiPart) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QHttpMultiPart) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QHttpMultiPart) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QHttpMultiPart) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QHttpMultiPart) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QHttpMultiPart) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QHttpPart struct {
	internal.Internal
}

type QHttpPart_ITF interface {
	QHttpPart_PTR() *QHttpPart
}

func (ptr *QHttpPart) QHttpPart_PTR() *QHttpPart {
	return ptr
}

func (ptr *QHttpPart) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QHttpPart) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQHttpPart(ptr QHttpPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpPart_PTR().Pointer()
	}
	return nil
}

func (n *QHttpPart) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQHttpPartFromPointer(ptr unsafe.Pointer) (n *QHttpPart) {
	n = new(QHttpPart)
	n.InitFromInternal(uintptr(ptr), "network.QHttpPart")
	return
}
func NewQHttpPart() *QHttpPart {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHttpPart", ""}).(*QHttpPart)
}

func NewQHttpPart2(other QHttpPart_ITF) *QHttpPart {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQHttpPart2", "", other}).(*QHttpPart)
}

func (ptr *QHttpPart) SetBody(body core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBody", body})
}

func (ptr *QHttpPart) SetBodyDevice(device core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBodyDevice", device})
}

func (ptr *QHttpPart) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeader", header, value})
}

func (ptr *QHttpPart) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRawHeader", headerName, headerValue})
}

func (ptr *QHttpPart) Swap(other QHttpPart_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QHttpPart) DestroyQHttpPart() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQHttpPart"})
}

type QIPv6Address struct {
	internal.Internal
}

type QIPv6Address_ITF interface {
	QIPv6Address_PTR() *QIPv6Address
}

func (ptr *QIPv6Address) QIPv6Address_PTR() *QIPv6Address {
	return ptr
}

func (ptr *QIPv6Address) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QIPv6Address) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQIPv6Address(ptr QIPv6Address_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QIPv6Address_PTR().Pointer()
	}
	return nil
}

func (n *QIPv6Address) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQIPv6AddressFromPointer(ptr unsafe.Pointer) (n *QIPv6Address) {
	n = new(QIPv6Address)
	n.InitFromInternal(uintptr(ptr), "network.QIPv6Address")
	return
}

func (ptr *QIPv6Address) DestroyQIPv6Address() {
}

type QLocalServer struct {
	core.QObject
}

type QLocalServer_ITF interface {
	core.QObject_ITF
	QLocalServer_PTR() *QLocalServer
}

func (ptr *QLocalServer) QLocalServer_PTR() *QLocalServer {
	return ptr
}

func (ptr *QLocalServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QLocalServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQLocalServer(ptr QLocalServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalServer_PTR().Pointer()
	}
	return nil
}

func (n *QLocalServer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLocalServer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQLocalServerFromPointer(ptr unsafe.Pointer) (n *QLocalServer) {
	n = new(QLocalServer)
	n.InitFromInternal(uintptr(ptr), "network.QLocalServer")
	return
}

//go:generate stringer -type=QLocalServer__SocketOption
//QLocalServer::SocketOption
type QLocalServer__SocketOption int64

const (
	QLocalServer__NoOptions         QLocalServer__SocketOption = QLocalServer__SocketOption(0x0)
	QLocalServer__UserAccessOption  QLocalServer__SocketOption = QLocalServer__SocketOption(0x01)
	QLocalServer__GroupAccessOption QLocalServer__SocketOption = QLocalServer__SocketOption(0x2)
	QLocalServer__OtherAccessOption QLocalServer__SocketOption = QLocalServer__SocketOption(0x4)
	QLocalServer__WorldAccessOption QLocalServer__SocketOption = QLocalServer__SocketOption(0x7)
)

func NewQLocalServer(parent core.QObject_ITF) *QLocalServer {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQLocalServer", "", parent}).(*QLocalServer)
}

func (ptr *QLocalServer) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QLocalServer) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QLocalServer) FullServerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FullServerName"}).(string)
}

func (ptr *QLocalServer) ConnectHasPendingConnections(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasPendingConnections", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalServer) DisconnectHasPendingConnections() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasPendingConnections"})
}

func (ptr *QLocalServer) HasPendingConnections() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasPendingConnections"}).(bool)
}

func (ptr *QLocalServer) HasPendingConnectionsDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasPendingConnectionsDefault"}).(bool)
}

func (ptr *QLocalServer) ConnectIncomingConnection(f func(socketDescriptor uintptr)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIncomingConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalServer) DisconnectIncomingConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIncomingConnection"})
}

func (ptr *QLocalServer) IncomingConnection(socketDescriptor uintptr) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncomingConnection", socketDescriptor})
}

func (ptr *QLocalServer) IncomingConnectionDefault(socketDescriptor uintptr) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IncomingConnectionDefault", socketDescriptor})
}

func (ptr *QLocalServer) IsListening() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsListening"}).(bool)
}

func (ptr *QLocalServer) Listen(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Listen", name}).(bool)
}

func (ptr *QLocalServer) MaxPendingConnections() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxPendingConnections"}).(float64))
}

func (ptr *QLocalServer) ConnectNewConnection(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalServer) DisconnectNewConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewConnection"})
}

func (ptr *QLocalServer) NewConnection() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewConnection"})
}

func (ptr *QLocalServer) ConnectNextPendingConnection(f func() *QLocalSocket) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNextPendingConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalServer) DisconnectNextPendingConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNextPendingConnection"})
}

func (ptr *QLocalServer) NextPendingConnection() *QLocalSocket {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextPendingConnection"}).(*QLocalSocket)
}

func (ptr *QLocalServer) NextPendingConnectionDefault() *QLocalSocket {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextPendingConnectionDefault"}).(*QLocalSocket)
}

func QLocalServer_RemoveServer(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QLocalServer_RemoveServer", "", name}).(bool)
}

func (ptr *QLocalServer) RemoveServer(name string) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QLocalServer_RemoveServer", "", name}).(bool)
}

func (ptr *QLocalServer) ServerError() QAbstractSocket__SocketError {

	return QAbstractSocket__SocketError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerError"}).(float64))
}

func (ptr *QLocalServer) ServerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerName"}).(string)
}

func (ptr *QLocalServer) SetMaxPendingConnections(numConnections int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaxPendingConnections", numConnections})
}

func (ptr *QLocalServer) SetSocketOptions(options QLocalServer__SocketOption) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSocketOptions", options})
}

func (ptr *QLocalServer) SocketOptions() QLocalServer__SocketOption {

	return QLocalServer__SocketOption(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SocketOptions"}).(float64))
}

func (ptr *QLocalServer) WaitForNewConnection(msec int, timedOut *bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForNewConnection", msec, timedOut}).(bool)
}

func (ptr *QLocalServer) ConnectDestroyQLocalServer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLocalServer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalServer) DisconnectDestroyQLocalServer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLocalServer"})
}

func (ptr *QLocalServer) DestroyQLocalServer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLocalServer"})
}

func (ptr *QLocalServer) DestroyQLocalServerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLocalServerDefault"})
}

func (ptr *QLocalServer) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QLocalServer) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QLocalServer) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QLocalServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QLocalServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QLocalServer) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QLocalServer) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QLocalServer) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QLocalServer) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QLocalServer) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QLocalServer) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QLocalServer) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QLocalServer) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QLocalServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QLocalServer) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QLocalServer) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QLocalServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QLocalServer) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QLocalServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QLocalServer) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QLocalServer) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QLocalSocket struct {
	core.QIODevice
}

type QLocalSocket_ITF interface {
	core.QIODevice_ITF
	QLocalSocket_PTR() *QLocalSocket
}

func (ptr *QLocalSocket) QLocalSocket_PTR() *QLocalSocket {
	return ptr
}

func (ptr *QLocalSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QLocalSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQLocalSocket(ptr QLocalSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocalSocket_PTR().Pointer()
	}
	return nil
}

func (n *QLocalSocket) InitFromInternal(ptr uintptr, name string) {
	n.QIODevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QLocalSocket) ClassNameInternalF() string {
	return n.QIODevice_PTR().ClassNameInternalF()
}

func NewQLocalSocketFromPointer(ptr unsafe.Pointer) (n *QLocalSocket) {
	n = new(QLocalSocket)
	n.InitFromInternal(uintptr(ptr), "network.QLocalSocket")
	return
}

//go:generate stringer -type=QLocalSocket__LocalSocketError
//QLocalSocket::LocalSocketError
type QLocalSocket__LocalSocketError int64

const (
	QLocalSocket__ConnectionRefusedError          QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__ConnectionRefusedError)
	QLocalSocket__PeerClosedError                 QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__RemoteHostClosedError)
	QLocalSocket__ServerNotFoundError             QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__HostNotFoundError)
	QLocalSocket__SocketAccessError               QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__SocketAccessError)
	QLocalSocket__SocketResourceError             QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__SocketResourceError)
	QLocalSocket__SocketTimeoutError              QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__SocketTimeoutError)
	QLocalSocket__DatagramTooLargeError           QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__DatagramTooLargeError)
	QLocalSocket__ConnectionError                 QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__NetworkError)
	QLocalSocket__UnsupportedSocketOperationError QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__UnsupportedSocketOperationError)
	QLocalSocket__UnknownSocketError              QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__UnknownSocketError)
	QLocalSocket__OperationError                  QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__OperationError)
)

//go:generate stringer -type=QLocalSocket__LocalSocketState
//QLocalSocket::LocalSocketState
type QLocalSocket__LocalSocketState int64

const (
	QLocalSocket__UnconnectedState QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__UnconnectedState)
	QLocalSocket__ConnectingState  QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectingState)
	QLocalSocket__ConnectedState   QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectedState)
	QLocalSocket__ClosingState     QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__ClosingState)
)

func NewQLocalSocket(parent core.QObject_ITF) *QLocalSocket {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQLocalSocket", "", parent}).(*QLocalSocket)
}

func (ptr *QLocalSocket) Abort() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"})
}

func (ptr *QLocalSocket) BytesAvailableDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesAvailableDefault"}).(float64))
}

func (ptr *QLocalSocket) BytesToWriteDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesToWriteDefault"}).(float64))
}

func (ptr *QLocalSocket) CanReadLineDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanReadLineDefault"}).(bool)
}

func (ptr *QLocalSocket) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QLocalSocket) ConnectToServer(openMode core.QIODevice__OpenModeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToServer", openMode})
}

func (ptr *QLocalSocket) ConnectToServer2(name string, openMode core.QIODevice__OpenModeFlag) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToServer2", name, openMode})
}

func (ptr *QLocalSocket) ConnectConnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalSocket) DisconnectConnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnected"})
}

func (ptr *QLocalSocket) Connected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connected"})
}

func (ptr *QLocalSocket) DisconnectFromServer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFromServer"})
}

func (ptr *QLocalSocket) ConnectDisconnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDisconnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalSocket) DisconnectDisconnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDisconnected"})
}

func (ptr *QLocalSocket) Disconnected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnected"})
}

func (ptr *QLocalSocket) Error() QLocalSocket__LocalSocketError {

	return QLocalSocket__LocalSocketError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QLocalSocket) ConnectError2(f func(socketError QLocalSocket__LocalSocketError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalSocket) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QLocalSocket) Error2(socketError QLocalSocket__LocalSocketError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", socketError})
}

func (ptr *QLocalSocket) Flush() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flush"}).(bool)
}

func (ptr *QLocalSocket) FullServerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FullServerName"}).(string)
}

func (ptr *QLocalSocket) IsSequentialDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSequentialDefault"}).(bool)
}

func (ptr *QLocalSocket) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QLocalSocket) OpenDefault(openMode core.QIODevice__OpenModeFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault", openMode}).(bool)
}

func (ptr *QLocalSocket) ReadBufferSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadBufferSize"}).(float64))
}

func (ptr *QLocalSocket) ConnectReadData(f func(data *string, c int64) int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReadData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalSocket) DisconnectReadData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReadData"})
}

func (ptr *QLocalSocket) ReadData(data *string, c int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadData", data, c}).(float64))
}

func (ptr *QLocalSocket) ReadDataDefault(data *string, c int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadDataDefault", data, c}).(float64))
}

func (ptr *QLocalSocket) ServerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerName"}).(string)
}

func (ptr *QLocalSocket) SetReadBufferSize(size int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadBufferSize", size})
}

func (ptr *QLocalSocket) SetServerName(name string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServerName", name})
}

func (ptr *QLocalSocket) State() QLocalSocket__LocalSocketState {

	return QLocalSocket__LocalSocketState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QLocalSocket) ConnectStateChanged(f func(socketState QLocalSocket__LocalSocketState)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalSocket) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QLocalSocket) StateChanged(socketState QLocalSocket__LocalSocketState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", socketState})
}

func (ptr *QLocalSocket) WaitForBytesWrittenDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForBytesWrittenDefault", msecs}).(bool)
}

func (ptr *QLocalSocket) WaitForConnected(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForConnected", msecs}).(bool)
}

func (ptr *QLocalSocket) WaitForDisconnected(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForDisconnected", msecs}).(bool)
}

func (ptr *QLocalSocket) WaitForReadyReadDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForReadyReadDefault", msecs}).(bool)
}

func (ptr *QLocalSocket) ConnectWriteData(f func(data []byte, c int64) int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectWriteData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalSocket) DisconnectWriteData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectWriteData"})
}

func (ptr *QLocalSocket) WriteData(data []byte, c int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteData", data, c}).(float64))
}

func (ptr *QLocalSocket) WriteDataDefault(data []byte, c int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDataDefault", data, c}).(float64))
}

func (ptr *QLocalSocket) ConnectDestroyQLocalSocket(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQLocalSocket", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QLocalSocket) DisconnectDestroyQLocalSocket() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQLocalSocket"})
}

func (ptr *QLocalSocket) DestroyQLocalSocket() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLocalSocket"})
}

func (ptr *QLocalSocket) DestroyQLocalSocketDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQLocalSocketDefault"})
}

func (ptr *QLocalSocket) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QLocalSocket) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QLocalSocket) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QLocalSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QLocalSocket) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QLocalSocket) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QLocalSocket) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QLocalSocket) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QLocalSocket) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QLocalSocket) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QLocalSocket) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QLocalSocket) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QLocalSocket) AtEndDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtEndDefault"}).(bool)
}

func (ptr *QLocalSocket) PosDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PosDefault"}).(float64))
}

func (ptr *QLocalSocket) ReadLineDataDefault(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadLineDataDefault", data, maxSize}).(float64))
}

func (ptr *QLocalSocket) ResetDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetDefault"}).(bool)
}

func (ptr *QLocalSocket) SeekDefault(pos int64) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeekDefault", pos}).(bool)
}

func (ptr *QLocalSocket) SizeDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeDefault"}).(float64))
}

func (ptr *QLocalSocket) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QLocalSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QLocalSocket) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QLocalSocket) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QLocalSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QLocalSocket) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QLocalSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QLocalSocket) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QLocalSocket) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QNetworkAccessManager struct {
	core.QObject
}

type QNetworkAccessManager_ITF interface {
	core.QObject_ITF
	QNetworkAccessManager_PTR() *QNetworkAccessManager
}

func (ptr *QNetworkAccessManager) QNetworkAccessManager_PTR() *QNetworkAccessManager {
	return ptr
}

func (ptr *QNetworkAccessManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkAccessManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkAccessManager(ptr QNetworkAccessManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAccessManager_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkAccessManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNetworkAccessManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQNetworkAccessManagerFromPointer(ptr unsafe.Pointer) (n *QNetworkAccessManager) {
	n = new(QNetworkAccessManager)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkAccessManager")
	return
}

//go:generate stringer -type=QNetworkAccessManager__Operation
//QNetworkAccessManager::Operation
type QNetworkAccessManager__Operation int64

const (
	QNetworkAccessManager__HeadOperation    QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(1)
	QNetworkAccessManager__GetOperation     QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(2)
	QNetworkAccessManager__PutOperation     QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(3)
	QNetworkAccessManager__PostOperation    QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(4)
	QNetworkAccessManager__DeleteOperation  QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(5)
	QNetworkAccessManager__CustomOperation  QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(6)
	QNetworkAccessManager__UnknownOperation QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(0)
)

//go:generate stringer -type=QNetworkAccessManager__NetworkAccessibility
//QNetworkAccessManager::NetworkAccessibility
type QNetworkAccessManager__NetworkAccessibility int64

const (
	QNetworkAccessManager__UnknownAccessibility QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(-1)
	QNetworkAccessManager__NotAccessible        QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(0)
	QNetworkAccessManager__Accessible           QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(1)
)

func NewQNetworkAccessManager(parent core.QObject_ITF) *QNetworkAccessManager {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkAccessManager", "", parent}).(*QNetworkAccessManager)
}

func (ptr *QNetworkAccessManager) ActiveConfiguration() *QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveConfiguration"}).(*QNetworkConfiguration)
}

func (ptr *QNetworkAccessManager) AddStrictTransportSecurityHosts(knownHosts []*QHstsPolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddStrictTransportSecurityHosts", knownHosts})
}

func (ptr *QNetworkAccessManager) ConnectAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAuthenticationRequired"})
}

func (ptr *QNetworkAccessManager) AuthenticationRequired(reply QNetworkReply_ITF, authenticator QAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AuthenticationRequired", reply, authenticator})
}

func (ptr *QNetworkAccessManager) Cache() *QAbstractNetworkCache {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Cache"}).(*QAbstractNetworkCache)
}

func (ptr *QNetworkAccessManager) ClearAccessCache() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearAccessCache"})
}

func (ptr *QNetworkAccessManager) ClearConnectionCache() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearConnectionCache"})
}

func (ptr *QNetworkAccessManager) Configuration() *QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Configuration"}).(*QNetworkConfiguration)
}

func (ptr *QNetworkAccessManager) ConnectToHost(hostName string, port uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToHost", hostName, port})
}

func (ptr *QNetworkAccessManager) ConnectToHostEncrypted(hostName string, port uint16, sslConfiguration QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToHostEncrypted", hostName, port, sslConfiguration})
}

func (ptr *QNetworkAccessManager) ConnectToHostEncrypted2(hostName string, port uint16, sslConfiguration QSslConfiguration_ITF, peerName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToHostEncrypted2", hostName, port, sslConfiguration, peerName})
}

func (ptr *QNetworkAccessManager) CookieJar() *QNetworkCookieJar {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CookieJar"}).(*QNetworkCookieJar)
}

func (ptr *QNetworkAccessManager) ConnectCreateRequest(f func(op QNetworkAccessManager__Operation, originalReq *QNetworkRequest, outgoingData *core.QIODevice) *QNetworkReply) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateRequest", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectCreateRequest() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateRequest"})
}

func (ptr *QNetworkAccessManager) CreateRequest(op QNetworkAccessManager__Operation, originalReq QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateRequest", op, originalReq, outgoingData}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) CreateRequestDefault(op QNetworkAccessManager__Operation, originalReq QNetworkRequest_ITF, outgoingData core.QIODevice_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateRequestDefault", op, originalReq, outgoingData}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) DeleteResource(request QNetworkRequest_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteResource", request}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) EnableStrictTransportSecurityStore(enabled bool, storeDir string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EnableStrictTransportSecurityStore", enabled, storeDir})
}

func (ptr *QNetworkAccessManager) ConnectEncrypted(f func(reply *QNetworkReply)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEncrypted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectEncrypted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEncrypted"})
}

func (ptr *QNetworkAccessManager) Encrypted(reply QNetworkReply_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Encrypted", reply})
}

func (ptr *QNetworkAccessManager) ConnectFinished(f func(reply *QNetworkReply)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QNetworkAccessManager) Finished(reply QNetworkReply_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished", reply})
}

func (ptr *QNetworkAccessManager) Get(request QNetworkRequest_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Get", request}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) Head(request QNetworkRequest_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Head", request}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) IsStrictTransportSecurityEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsStrictTransportSecurityEnabled"}).(bool)
}

func (ptr *QNetworkAccessManager) IsStrictTransportSecurityStoreEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsStrictTransportSecurityStoreEnabled"}).(bool)
}

func (ptr *QNetworkAccessManager) NetworkAccessible() QNetworkAccessManager__NetworkAccessibility {

	return QNetworkAccessManager__NetworkAccessibility(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NetworkAccessible"}).(float64))
}

func (ptr *QNetworkAccessManager) ConnectNetworkAccessibleChanged(f func(accessible QNetworkAccessManager__NetworkAccessibility)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNetworkAccessibleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectNetworkAccessibleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNetworkAccessibleChanged"})
}

func (ptr *QNetworkAccessManager) NetworkAccessibleChanged(accessible QNetworkAccessManager__NetworkAccessibility) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NetworkAccessibleChanged", accessible})
}

func (ptr *QNetworkAccessManager) Post(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Post", request, data}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) Post2(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Post2", request, data}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) Post3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Post3", request, multiPart}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) ConnectPreSharedKeyAuthenticationRequired(f func(reply *QNetworkReply, authenticator *QSslPreSharedKeyAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreSharedKeyAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectPreSharedKeyAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreSharedKeyAuthenticationRequired"})
}

func (ptr *QNetworkAccessManager) PreSharedKeyAuthenticationRequired(reply QNetworkReply_ITF, authenticator QSslPreSharedKeyAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreSharedKeyAuthenticationRequired", reply, authenticator})
}

func (ptr *QNetworkAccessManager) Proxy() *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Proxy"}).(*QNetworkProxy)
}

func (ptr *QNetworkAccessManager) ConnectProxyAuthenticationRequired(f func(proxy *QNetworkProxy, authenticator *QAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProxyAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectProxyAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProxyAuthenticationRequired"})
}

func (ptr *QNetworkAccessManager) ProxyAuthenticationRequired(proxy QNetworkProxy_ITF, authenticator QAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProxyAuthenticationRequired", proxy, authenticator})
}

func (ptr *QNetworkAccessManager) ProxyFactory() *QNetworkProxyFactory {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProxyFactory"}).(*QNetworkProxyFactory)
}

func (ptr *QNetworkAccessManager) Put(request QNetworkRequest_ITF, data core.QIODevice_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Put", request, data}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) Put2(request QNetworkRequest_ITF, data core.QByteArray_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Put2", request, data}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) Put3(request QNetworkRequest_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Put3", request, multiPart}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) RedirectPolicy() QNetworkRequest__RedirectPolicy {

	return QNetworkRequest__RedirectPolicy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RedirectPolicy"}).(float64))
}

func (ptr *QNetworkAccessManager) SendCustomRequest(request QNetworkRequest_ITF, verb core.QByteArray_ITF, data core.QIODevice_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendCustomRequest", request, verb, data}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) SendCustomRequest2(request QNetworkRequest_ITF, verb core.QByteArray_ITF, data core.QByteArray_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendCustomRequest2", request, verb, data}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) SendCustomRequest3(request QNetworkRequest_ITF, verb core.QByteArray_ITF, multiPart QHttpMultiPart_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendCustomRequest3", request, verb, multiPart}).(*QNetworkReply)
}

func (ptr *QNetworkAccessManager) SetCache(cache QAbstractNetworkCache_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCache", cache})
}

func (ptr *QNetworkAccessManager) SetConfiguration(config QNetworkConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetConfiguration", config})
}

func (ptr *QNetworkAccessManager) SetCookieJar(cookieJar QNetworkCookieJar_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCookieJar", cookieJar})
}

func (ptr *QNetworkAccessManager) SetNetworkAccessible(accessible QNetworkAccessManager__NetworkAccessibility) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNetworkAccessible", accessible})
}

func (ptr *QNetworkAccessManager) SetProxy(proxy QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProxy", proxy})
}

func (ptr *QNetworkAccessManager) SetProxyFactory(factory QNetworkProxyFactory_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProxyFactory", factory})
}

func (ptr *QNetworkAccessManager) SetRedirectPolicy(policy QNetworkRequest__RedirectPolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRedirectPolicy", policy})
}

func (ptr *QNetworkAccessManager) SetStrictTransportSecurityEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetStrictTransportSecurityEnabled", enabled})
}

func (ptr *QNetworkAccessManager) ConnectSslErrors(f func(reply *QNetworkReply, errors []*QSslError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSslErrors", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectSslErrors() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSslErrors"})
}

func (ptr *QNetworkAccessManager) SslErrors(reply QNetworkReply_ITF, errors []*QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslErrors", reply, errors})
}

func (ptr *QNetworkAccessManager) StrictTransportSecurityHosts() []*QHstsPolicy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StrictTransportSecurityHosts"}).([]*QHstsPolicy)
}

func (ptr *QNetworkAccessManager) SupportedSchemes() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedSchemes"}).([]string)
}

func (ptr *QNetworkAccessManager) ConnectSupportedSchemesImplementation(f func() []string) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSupportedSchemesImplementation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectSupportedSchemesImplementation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSupportedSchemesImplementation"})
}

func (ptr *QNetworkAccessManager) SupportedSchemesImplementation() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedSchemesImplementation"}).([]string)
}

func (ptr *QNetworkAccessManager) SupportedSchemesImplementationDefault() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedSchemesImplementationDefault"}).([]string)
}

func (ptr *QNetworkAccessManager) ConnectDestroyQNetworkAccessManager(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNetworkAccessManager", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkAccessManager) DisconnectDestroyQNetworkAccessManager() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNetworkAccessManager"})
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManager() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkAccessManager"})
}

func (ptr *QNetworkAccessManager) DestroyQNetworkAccessManagerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkAccessManagerDefault"})
}

func (ptr *QNetworkAccessManager) __addStrictTransportSecurityHosts_knownHosts_atList(i int) *QHstsPolicy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addStrictTransportSecurityHosts_knownHosts_atList", i}).(*QHstsPolicy)
}

func (ptr *QNetworkAccessManager) __addStrictTransportSecurityHosts_knownHosts_setList(i QHstsPolicy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addStrictTransportSecurityHosts_knownHosts_setList", i})
}

func (ptr *QNetworkAccessManager) __addStrictTransportSecurityHosts_knownHosts_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addStrictTransportSecurityHosts_knownHosts_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkAccessManager) __sslErrors_errors_atList(i int) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_atList", i}).(*QSslError)
}

func (ptr *QNetworkAccessManager) __sslErrors_errors_setList(i QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_setList", i})
}

func (ptr *QNetworkAccessManager) __sslErrors_errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkAccessManager) __strictTransportSecurityHosts_atList(i int) *QHstsPolicy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__strictTransportSecurityHosts_atList", i}).(*QHstsPolicy)
}

func (ptr *QNetworkAccessManager) __strictTransportSecurityHosts_setList(i QHstsPolicy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__strictTransportSecurityHosts_setList", i})
}

func (ptr *QNetworkAccessManager) __strictTransportSecurityHosts_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__strictTransportSecurityHosts_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkAccessManager) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QNetworkAccessManager) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNetworkAccessManager) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkAccessManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QNetworkAccessManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QNetworkAccessManager) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkAccessManager) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QNetworkAccessManager) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QNetworkAccessManager) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkAccessManager) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QNetworkAccessManager) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QNetworkAccessManager) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QNetworkAccessManager) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QNetworkAccessManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QNetworkAccessManager) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QNetworkAccessManager) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QNetworkAccessManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QNetworkAccessManager) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QNetworkAccessManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QNetworkAccessManager) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QNetworkAccessManager) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QNetworkAddressEntry struct {
	internal.Internal
}

type QNetworkAddressEntry_ITF interface {
	QNetworkAddressEntry_PTR() *QNetworkAddressEntry
}

func (ptr *QNetworkAddressEntry) QNetworkAddressEntry_PTR() *QNetworkAddressEntry {
	return ptr
}

func (ptr *QNetworkAddressEntry) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkAddressEntry) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkAddressEntry(ptr QNetworkAddressEntry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkAddressEntry_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkAddressEntry) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkAddressEntryFromPointer(ptr unsafe.Pointer) (n *QNetworkAddressEntry) {
	n = new(QNetworkAddressEntry)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkAddressEntry")
	return
}

//go:generate stringer -type=QNetworkAddressEntry__DnsEligibilityStatus
//QNetworkAddressEntry::DnsEligibilityStatus
type QNetworkAddressEntry__DnsEligibilityStatus int64

const (
	QNetworkAddressEntry__DnsEligibilityUnknown QNetworkAddressEntry__DnsEligibilityStatus = QNetworkAddressEntry__DnsEligibilityStatus(-1)
	QNetworkAddressEntry__DnsIneligible         QNetworkAddressEntry__DnsEligibilityStatus = QNetworkAddressEntry__DnsEligibilityStatus(0)
	QNetworkAddressEntry__DnsEligible           QNetworkAddressEntry__DnsEligibilityStatus = QNetworkAddressEntry__DnsEligibilityStatus(1)
)

func NewQNetworkAddressEntry() *QNetworkAddressEntry {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkAddressEntry", ""}).(*QNetworkAddressEntry)
}

func NewQNetworkAddressEntry2(other QNetworkAddressEntry_ITF) *QNetworkAddressEntry {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkAddressEntry2", "", other}).(*QNetworkAddressEntry)
}

func (ptr *QNetworkAddressEntry) Broadcast() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Broadcast"}).(*QHostAddress)
}

func (ptr *QNetworkAddressEntry) ClearAddressLifetime() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearAddressLifetime"})
}

func (ptr *QNetworkAddressEntry) DnsEligibility() QNetworkAddressEntry__DnsEligibilityStatus {

	return QNetworkAddressEntry__DnsEligibilityStatus(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DnsEligibility"}).(float64))
}

func (ptr *QNetworkAddressEntry) Ip() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Ip"}).(*QHostAddress)
}

func (ptr *QNetworkAddressEntry) IsLifetimeKnown() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsLifetimeKnown"}).(bool)
}

func (ptr *QNetworkAddressEntry) IsPermanent() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsPermanent"}).(bool)
}

func (ptr *QNetworkAddressEntry) IsTemporary() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsTemporary"}).(bool)
}

func (ptr *QNetworkAddressEntry) Netmask() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Netmask"}).(*QHostAddress)
}

func (ptr *QNetworkAddressEntry) PrefixLength() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrefixLength"}).(float64))
}

func (ptr *QNetworkAddressEntry) SetAddressLifetime(preferred core.QDeadlineTimer_ITF, validity core.QDeadlineTimer_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAddressLifetime", preferred, validity})
}

func (ptr *QNetworkAddressEntry) SetBroadcast(newBroadcast QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBroadcast", newBroadcast})
}

func (ptr *QNetworkAddressEntry) SetDnsEligibility(status QNetworkAddressEntry__DnsEligibilityStatus) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDnsEligibility", status})
}

func (ptr *QNetworkAddressEntry) SetIp(newIp QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIp", newIp})
}

func (ptr *QNetworkAddressEntry) SetNetmask(newNetmask QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetNetmask", newNetmask})
}

func (ptr *QNetworkAddressEntry) SetPrefixLength(length int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrefixLength", length})
}

func (ptr *QNetworkAddressEntry) Swap(other QNetworkAddressEntry_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QNetworkAddressEntry) DestroyQNetworkAddressEntry() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkAddressEntry"})
}

type QNetworkCacheMetaData struct {
	internal.Internal
}

type QNetworkCacheMetaData_ITF interface {
	QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData
}

func (ptr *QNetworkCacheMetaData) QNetworkCacheMetaData_PTR() *QNetworkCacheMetaData {
	return ptr
}

func (ptr *QNetworkCacheMetaData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkCacheMetaData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkCacheMetaData(ptr QNetworkCacheMetaData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCacheMetaData_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkCacheMetaData) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkCacheMetaDataFromPointer(ptr unsafe.Pointer) (n *QNetworkCacheMetaData) {
	n = new(QNetworkCacheMetaData)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkCacheMetaData")
	return
}
func NewQNetworkCacheMetaData() *QNetworkCacheMetaData {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkCacheMetaData", ""}).(*QNetworkCacheMetaData)
}

func NewQNetworkCacheMetaData2(other QNetworkCacheMetaData_ITF) *QNetworkCacheMetaData {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkCacheMetaData2", "", other}).(*QNetworkCacheMetaData)
}

func (ptr *QNetworkCacheMetaData) ExpirationDate() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExpirationDate"}).(*core.QDateTime)
}

func (ptr *QNetworkCacheMetaData) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QNetworkCacheMetaData) LastModified() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LastModified"}).(*core.QDateTime)
}

func (ptr *QNetworkCacheMetaData) RawHeaders() []*QNetworkCacheMetaData_RawHeader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RawHeaders"}).([]*QNetworkCacheMetaData_RawHeader)
}

func (ptr *QNetworkCacheMetaData) SaveToDisk() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SaveToDisk"}).(bool)
}

func (ptr *QNetworkCacheMetaData) SetExpirationDate(dateTime core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExpirationDate", dateTime})
}

func (ptr *QNetworkCacheMetaData) SetLastModified(dateTime core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLastModified", dateTime})
}

func (ptr *QNetworkCacheMetaData) SetRawHeaders(list []*QNetworkCacheMetaData_RawHeader) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRawHeaders", list})
}

func (ptr *QNetworkCacheMetaData) SetSaveToDisk(allow bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSaveToDisk", allow})
}

func (ptr *QNetworkCacheMetaData) SetUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrl", url})
}

func (ptr *QNetworkCacheMetaData) Swap(other QNetworkCacheMetaData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QNetworkCacheMetaData) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QNetworkCacheMetaData) DestroyQNetworkCacheMetaData() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkCacheMetaData"})
}

func (ptr *QNetworkCacheMetaData) __rawHeaders_atList(i int) *QNetworkCacheMetaData_RawHeader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaders_atList", i}).(*QNetworkCacheMetaData_RawHeader)
}

func (ptr *QNetworkCacheMetaData) __rawHeaders_setList(i QNetworkCacheMetaData_RawHeader_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaders_setList", i})
}

func (ptr *QNetworkCacheMetaData) __rawHeaders_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaders_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkCacheMetaData) __setRawHeaders_list_atList(i int) *QNetworkCacheMetaData_RawHeader {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setRawHeaders_list_atList", i}).(*QNetworkCacheMetaData_RawHeader)
}

func (ptr *QNetworkCacheMetaData) __setRawHeaders_list_setList(i QNetworkCacheMetaData_RawHeader_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setRawHeaders_list_setList", i})
}

func (ptr *QNetworkCacheMetaData) __setRawHeaders_list_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setRawHeaders_list_newList"}).(unsafe.Pointer)
}

type QNetworkConfiguration struct {
	internal.Internal
}

type QNetworkConfiguration_ITF interface {
	QNetworkConfiguration_PTR() *QNetworkConfiguration
}

func (ptr *QNetworkConfiguration) QNetworkConfiguration_PTR() *QNetworkConfiguration {
	return ptr
}

func (ptr *QNetworkConfiguration) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkConfiguration) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkConfiguration(ptr QNetworkConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfiguration_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkConfiguration) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkConfigurationFromPointer(ptr unsafe.Pointer) (n *QNetworkConfiguration) {
	n = new(QNetworkConfiguration)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkConfiguration")
	return
}

//go:generate stringer -type=QNetworkConfiguration__Type
//QNetworkConfiguration::Type
type QNetworkConfiguration__Type int64

const (
	QNetworkConfiguration__InternetAccessPoint QNetworkConfiguration__Type = QNetworkConfiguration__Type(0)
	QNetworkConfiguration__ServiceNetwork      QNetworkConfiguration__Type = QNetworkConfiguration__Type(1)
	QNetworkConfiguration__UserChoice          QNetworkConfiguration__Type = QNetworkConfiguration__Type(2)
	QNetworkConfiguration__Invalid             QNetworkConfiguration__Type = QNetworkConfiguration__Type(3)
)

//go:generate stringer -type=QNetworkConfiguration__Purpose
//QNetworkConfiguration::Purpose
type QNetworkConfiguration__Purpose int64

const (
	QNetworkConfiguration__UnknownPurpose         QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(0)
	QNetworkConfiguration__PublicPurpose          QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(1)
	QNetworkConfiguration__PrivatePurpose         QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(2)
	QNetworkConfiguration__ServiceSpecificPurpose QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(3)
)

//go:generate stringer -type=QNetworkConfiguration__StateFlag
//QNetworkConfiguration::StateFlag
type QNetworkConfiguration__StateFlag int64

const (
	QNetworkConfiguration__Undefined  QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x0000001)
	QNetworkConfiguration__Defined    QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x0000002)
	QNetworkConfiguration__Discovered QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x0000006)
	QNetworkConfiguration__Active     QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x000000e)
)

//go:generate stringer -type=QNetworkConfiguration__BearerType
//QNetworkConfiguration::BearerType
type QNetworkConfiguration__BearerType int64

const (
	QNetworkConfiguration__BearerUnknown   QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(0)
	QNetworkConfiguration__BearerEthernet  QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(1)
	QNetworkConfiguration__BearerWLAN      QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(2)
	QNetworkConfiguration__Bearer2G        QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(3)
	QNetworkConfiguration__BearerCDMA2000  QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(4)
	QNetworkConfiguration__BearerWCDMA     QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(5)
	QNetworkConfiguration__BearerHSPA      QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(6)
	QNetworkConfiguration__BearerBluetooth QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(7)
	QNetworkConfiguration__BearerWiMAX     QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(8)
	QNetworkConfiguration__BearerEVDO      QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(9)
	QNetworkConfiguration__BearerLTE       QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(10)
	QNetworkConfiguration__Bearer3G        QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(11)
	QNetworkConfiguration__Bearer4G        QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(12)
)

func NewQNetworkConfiguration() *QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkConfiguration", ""}).(*QNetworkConfiguration)
}

func NewQNetworkConfiguration2(other QNetworkConfiguration_ITF) *QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkConfiguration2", "", other}).(*QNetworkConfiguration)
}

func (ptr *QNetworkConfiguration) BearerType() QNetworkConfiguration__BearerType {

	return QNetworkConfiguration__BearerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BearerType"}).(float64))
}

func (ptr *QNetworkConfiguration) BearerTypeFamily() QNetworkConfiguration__BearerType {

	return QNetworkConfiguration__BearerType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BearerTypeFamily"}).(float64))
}

func (ptr *QNetworkConfiguration) BearerTypeName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BearerTypeName"}).(string)
}

func (ptr *QNetworkConfiguration) Children() []*QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Children"}).([]*QNetworkConfiguration)
}

func (ptr *QNetworkConfiguration) ConnectTimeout() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTimeout"}).(float64))
}

func (ptr *QNetworkConfiguration) Identifier() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Identifier"}).(string)
}

func (ptr *QNetworkConfiguration) IsRoamingAvailable() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRoamingAvailable"}).(bool)
}

func (ptr *QNetworkConfiguration) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QNetworkConfiguration) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QNetworkConfiguration) Purpose() QNetworkConfiguration__Purpose {

	return QNetworkConfiguration__Purpose(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Purpose"}).(float64))
}

func (ptr *QNetworkConfiguration) SetConnectTimeout(timeout int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetConnectTimeout", timeout}).(bool)
}

func (ptr *QNetworkConfiguration) State() QNetworkConfiguration__StateFlag {

	return QNetworkConfiguration__StateFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QNetworkConfiguration) Swap(other QNetworkConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QNetworkConfiguration) Type() QNetworkConfiguration__Type {

	return QNetworkConfiguration__Type(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QNetworkConfiguration) DestroyQNetworkConfiguration() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkConfiguration"})
}

func (ptr *QNetworkConfiguration) __children_atList(i int) *QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*QNetworkConfiguration)
}

func (ptr *QNetworkConfiguration) __children_setList(i QNetworkConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNetworkConfiguration) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

type QNetworkConfigurationManager struct {
	core.QObject
}

type QNetworkConfigurationManager_ITF interface {
	core.QObject_ITF
	QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager
}

func (ptr *QNetworkConfigurationManager) QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager {
	return ptr
}

func (ptr *QNetworkConfigurationManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkConfigurationManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkConfigurationManager(ptr QNetworkConfigurationManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfigurationManager_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkConfigurationManager) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNetworkConfigurationManager) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQNetworkConfigurationManagerFromPointer(ptr unsafe.Pointer) (n *QNetworkConfigurationManager) {
	n = new(QNetworkConfigurationManager)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkConfigurationManager")
	return
}

//go:generate stringer -type=QNetworkConfigurationManager__Capability
//QNetworkConfigurationManager::Capability
type QNetworkConfigurationManager__Capability int64

const (
	QNetworkConfigurationManager__CanStartAndStopInterfaces QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000001)
	QNetworkConfigurationManager__DirectConnectionRouting   QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000002)
	QNetworkConfigurationManager__SystemSessionSupport      QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000004)
	QNetworkConfigurationManager__ApplicationLevelRoaming   QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000008)
	QNetworkConfigurationManager__ForcedRoaming             QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000010)
	QNetworkConfigurationManager__DataStatistics            QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000020)
	QNetworkConfigurationManager__NetworkSessionRequired    QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000040)
)

func NewQNetworkConfigurationManager(parent core.QObject_ITF) *QNetworkConfigurationManager {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkConfigurationManager", "", parent}).(*QNetworkConfigurationManager)
}

func (ptr *QNetworkConfigurationManager) AllConfigurations(filter QNetworkConfiguration__StateFlag) []*QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AllConfigurations", filter}).([]*QNetworkConfiguration)
}

func (ptr *QNetworkConfigurationManager) Capabilities() QNetworkConfigurationManager__Capability {

	return QNetworkConfigurationManager__Capability(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Capabilities"}).(float64))
}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationAdded(f func(config *QNetworkConfiguration)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConfigurationAdded", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationAdded() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConfigurationAdded"})
}

func (ptr *QNetworkConfigurationManager) ConfigurationAdded(config QNetworkConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConfigurationAdded", config})
}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationChanged(f func(config *QNetworkConfiguration)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConfigurationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConfigurationChanged"})
}

func (ptr *QNetworkConfigurationManager) ConfigurationChanged(config QNetworkConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConfigurationChanged", config})
}

func (ptr *QNetworkConfigurationManager) ConfigurationFromIdentifier(identifier string) *QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConfigurationFromIdentifier", identifier}).(*QNetworkConfiguration)
}

func (ptr *QNetworkConfigurationManager) ConnectConfigurationRemoved(f func(config *QNetworkConfiguration)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConfigurationRemoved", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkConfigurationManager) DisconnectConfigurationRemoved() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConfigurationRemoved"})
}

func (ptr *QNetworkConfigurationManager) ConfigurationRemoved(config QNetworkConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConfigurationRemoved", config})
}

func (ptr *QNetworkConfigurationManager) DefaultConfiguration() *QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DefaultConfiguration"}).(*QNetworkConfiguration)
}

func (ptr *QNetworkConfigurationManager) IsOnline() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOnline"}).(bool)
}

func (ptr *QNetworkConfigurationManager) ConnectOnlineStateChanged(f func(isOnline bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOnlineStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkConfigurationManager) DisconnectOnlineStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOnlineStateChanged"})
}

func (ptr *QNetworkConfigurationManager) OnlineStateChanged(isOnline bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OnlineStateChanged", isOnline})
}

func (ptr *QNetworkConfigurationManager) ConnectUpdateCompleted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateCompleted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateCompleted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateCompleted"})
}

func (ptr *QNetworkConfigurationManager) UpdateCompleted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateCompleted"})
}

func (ptr *QNetworkConfigurationManager) ConnectUpdateConfigurations(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateConfigurations", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateConfigurations() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateConfigurations"})
}

func (ptr *QNetworkConfigurationManager) UpdateConfigurations() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateConfigurations"})
}

func (ptr *QNetworkConfigurationManager) UpdateConfigurationsDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateConfigurationsDefault"})
}

func (ptr *QNetworkConfigurationManager) ConnectDestroyQNetworkConfigurationManager(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNetworkConfigurationManager", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkConfigurationManager) DisconnectDestroyQNetworkConfigurationManager() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNetworkConfigurationManager"})
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManager() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkConfigurationManager"})
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManagerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkConfigurationManagerDefault"})
}

func (ptr *QNetworkConfigurationManager) __allConfigurations_atList(i int) *QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allConfigurations_atList", i}).(*QNetworkConfiguration)
}

func (ptr *QNetworkConfigurationManager) __allConfigurations_setList(i QNetworkConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allConfigurations_setList", i})
}

func (ptr *QNetworkConfigurationManager) __allConfigurations_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allConfigurations_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkConfigurationManager) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QNetworkConfigurationManager) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNetworkConfigurationManager) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkConfigurationManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QNetworkConfigurationManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QNetworkConfigurationManager) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkConfigurationManager) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QNetworkConfigurationManager) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QNetworkConfigurationManager) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkConfigurationManager) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QNetworkConfigurationManager) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QNetworkConfigurationManager) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QNetworkConfigurationManager) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QNetworkConfigurationManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QNetworkConfigurationManager) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QNetworkConfigurationManager) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QNetworkConfigurationManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QNetworkConfigurationManager) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QNetworkConfigurationManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QNetworkConfigurationManager) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QNetworkConfigurationManager) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QNetworkCookie struct {
	internal.Internal
}

type QNetworkCookie_ITF interface {
	QNetworkCookie_PTR() *QNetworkCookie
}

func (ptr *QNetworkCookie) QNetworkCookie_PTR() *QNetworkCookie {
	return ptr
}

func (ptr *QNetworkCookie) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkCookie) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkCookie(ptr QNetworkCookie_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookie_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkCookie) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkCookieFromPointer(ptr unsafe.Pointer) (n *QNetworkCookie) {
	n = new(QNetworkCookie)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkCookie")
	return
}

//go:generate stringer -type=QNetworkCookie__RawForm
//QNetworkCookie::RawForm
type QNetworkCookie__RawForm int64

const (
	QNetworkCookie__NameAndValueOnly QNetworkCookie__RawForm = QNetworkCookie__RawForm(0)
	QNetworkCookie__Full             QNetworkCookie__RawForm = QNetworkCookie__RawForm(1)
)

func NewQNetworkCookie(name core.QByteArray_ITF, value core.QByteArray_ITF) *QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkCookie", "", name, value}).(*QNetworkCookie)
}

func NewQNetworkCookie2(other QNetworkCookie_ITF) *QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkCookie2", "", other}).(*QNetworkCookie)
}

func (ptr *QNetworkCookie) Domain() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Domain"}).(string)
}

func (ptr *QNetworkCookie) ExpirationDate() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExpirationDate"}).(*core.QDateTime)
}

func (ptr *QNetworkCookie) HasSameIdentifier(other QNetworkCookie_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasSameIdentifier", other}).(bool)
}

func (ptr *QNetworkCookie) IsHttpOnly() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsHttpOnly"}).(bool)
}

func (ptr *QNetworkCookie) IsSecure() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSecure"}).(bool)
}

func (ptr *QNetworkCookie) IsSessionCookie() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSessionCookie"}).(bool)
}

func (ptr *QNetworkCookie) Name() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(*core.QByteArray)
}

func (ptr *QNetworkCookie) Normalize(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Normalize", url})
}

func QNetworkCookie_ParseCookies(cookieString core.QByteArray_ITF) []*QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkCookie_ParseCookies", "", cookieString}).([]*QNetworkCookie)
}

func (ptr *QNetworkCookie) ParseCookies(cookieString core.QByteArray_ITF) []*QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkCookie_ParseCookies", "", cookieString}).([]*QNetworkCookie)
}

func (ptr *QNetworkCookie) Path() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Path"}).(string)
}

func (ptr *QNetworkCookie) SetDomain(domain string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDomain", domain})
}

func (ptr *QNetworkCookie) SetExpirationDate(date core.QDateTime_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetExpirationDate", date})
}

func (ptr *QNetworkCookie) SetHttpOnly(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHttpOnly", enable})
}

func (ptr *QNetworkCookie) SetName(cookieName core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetName", cookieName})
}

func (ptr *QNetworkCookie) SetPath(path string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPath", path})
}

func (ptr *QNetworkCookie) SetSecure(enable bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSecure", enable})
}

func (ptr *QNetworkCookie) SetValue(value core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetValue", value})
}

func (ptr *QNetworkCookie) Swap(other QNetworkCookie_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QNetworkCookie) ToRawForm(form QNetworkCookie__RawForm) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToRawForm", form}).(*core.QByteArray)
}

func (ptr *QNetworkCookie) Value() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(*core.QByteArray)
}

func (ptr *QNetworkCookie) DestroyQNetworkCookie() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkCookie"})
}

func (ptr *QNetworkCookie) __parseCookies_atList(i int) *QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parseCookies_atList", i}).(*QNetworkCookie)
}

func (ptr *QNetworkCookie) __parseCookies_setList(i QNetworkCookie_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parseCookies_setList", i})
}

func (ptr *QNetworkCookie) __parseCookies_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__parseCookies_newList"}).(unsafe.Pointer)
}

type QNetworkCookieJar struct {
	core.QObject
}

type QNetworkCookieJar_ITF interface {
	core.QObject_ITF
	QNetworkCookieJar_PTR() *QNetworkCookieJar
}

func (ptr *QNetworkCookieJar) QNetworkCookieJar_PTR() *QNetworkCookieJar {
	return ptr
}

func (ptr *QNetworkCookieJar) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkCookieJar) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkCookieJar(ptr QNetworkCookieJar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookieJar_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkCookieJar) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNetworkCookieJar) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQNetworkCookieJarFromPointer(ptr unsafe.Pointer) (n *QNetworkCookieJar) {
	n = new(QNetworkCookieJar)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkCookieJar")
	return
}
func NewQNetworkCookieJar(parent core.QObject_ITF) *QNetworkCookieJar {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkCookieJar", "", parent}).(*QNetworkCookieJar)
}

func (ptr *QNetworkCookieJar) AllCookies() []*QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AllCookies"}).([]*QNetworkCookie)
}

func (ptr *QNetworkCookieJar) ConnectCookiesForUrl(f func(url *core.QUrl) []*QNetworkCookie) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCookiesForUrl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkCookieJar) DisconnectCookiesForUrl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCookiesForUrl"})
}

func (ptr *QNetworkCookieJar) CookiesForUrl(url core.QUrl_ITF) []*QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CookiesForUrl", url}).([]*QNetworkCookie)
}

func (ptr *QNetworkCookieJar) CookiesForUrlDefault(url core.QUrl_ITF) []*QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CookiesForUrlDefault", url}).([]*QNetworkCookie)
}

func (ptr *QNetworkCookieJar) ConnectDeleteCookie(f func(cookie *QNetworkCookie) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDeleteCookie", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkCookieJar) DisconnectDeleteCookie() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDeleteCookie"})
}

func (ptr *QNetworkCookieJar) DeleteCookie(cookie QNetworkCookie_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteCookie", cookie}).(bool)
}

func (ptr *QNetworkCookieJar) DeleteCookieDefault(cookie QNetworkCookie_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteCookieDefault", cookie}).(bool)
}

func (ptr *QNetworkCookieJar) ConnectInsertCookie(f func(cookie *QNetworkCookie) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInsertCookie", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkCookieJar) DisconnectInsertCookie() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInsertCookie"})
}

func (ptr *QNetworkCookieJar) InsertCookie(cookie QNetworkCookie_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertCookie", cookie}).(bool)
}

func (ptr *QNetworkCookieJar) InsertCookieDefault(cookie QNetworkCookie_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertCookieDefault", cookie}).(bool)
}

func (ptr *QNetworkCookieJar) SetAllCookies(cookieList []*QNetworkCookie) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAllCookies", cookieList})
}

func (ptr *QNetworkCookieJar) ConnectSetCookiesFromUrl(f func(cookieList []*QNetworkCookie, url *core.QUrl) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetCookiesFromUrl", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkCookieJar) DisconnectSetCookiesFromUrl() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetCookiesFromUrl"})
}

func (ptr *QNetworkCookieJar) SetCookiesFromUrl(cookieList []*QNetworkCookie, url core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCookiesFromUrl", cookieList, url}).(bool)
}

func (ptr *QNetworkCookieJar) SetCookiesFromUrlDefault(cookieList []*QNetworkCookie, url core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCookiesFromUrlDefault", cookieList, url}).(bool)
}

func (ptr *QNetworkCookieJar) ConnectUpdateCookie(f func(cookie *QNetworkCookie) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateCookie", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkCookieJar) DisconnectUpdateCookie() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateCookie"})
}

func (ptr *QNetworkCookieJar) UpdateCookie(cookie QNetworkCookie_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateCookie", cookie}).(bool)
}

func (ptr *QNetworkCookieJar) UpdateCookieDefault(cookie QNetworkCookie_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateCookieDefault", cookie}).(bool)
}

func (ptr *QNetworkCookieJar) ConnectValidateCookie(f func(cookie *QNetworkCookie, url *core.QUrl) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectValidateCookie", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkCookieJar) DisconnectValidateCookie() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectValidateCookie"})
}

func (ptr *QNetworkCookieJar) ValidateCookie(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValidateCookie", cookie, url}).(bool)
}

func (ptr *QNetworkCookieJar) ValidateCookieDefault(cookie QNetworkCookie_ITF, url core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ValidateCookieDefault", cookie, url}).(bool)
}

func (ptr *QNetworkCookieJar) ConnectDestroyQNetworkCookieJar(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNetworkCookieJar", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkCookieJar) DisconnectDestroyQNetworkCookieJar() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNetworkCookieJar"})
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJar() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkCookieJar"})
}

func (ptr *QNetworkCookieJar) DestroyQNetworkCookieJarDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkCookieJarDefault"})
}

func (ptr *QNetworkCookieJar) __allCookies_atList(i int) *QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allCookies_atList", i}).(*QNetworkCookie)
}

func (ptr *QNetworkCookieJar) __allCookies_setList(i QNetworkCookie_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allCookies_setList", i})
}

func (ptr *QNetworkCookieJar) __allCookies_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allCookies_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkCookieJar) __cookiesForUrl_atList(i int) *QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__cookiesForUrl_atList", i}).(*QNetworkCookie)
}

func (ptr *QNetworkCookieJar) __cookiesForUrl_setList(i QNetworkCookie_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__cookiesForUrl_setList", i})
}

func (ptr *QNetworkCookieJar) __cookiesForUrl_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__cookiesForUrl_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkCookieJar) __setAllCookies_cookieList_atList(i int) *QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAllCookies_cookieList_atList", i}).(*QNetworkCookie)
}

func (ptr *QNetworkCookieJar) __setAllCookies_cookieList_setList(i QNetworkCookie_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAllCookies_cookieList_setList", i})
}

func (ptr *QNetworkCookieJar) __setAllCookies_cookieList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAllCookies_cookieList_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkCookieJar) __setCookiesFromUrl_cookieList_atList(i int) *QNetworkCookie {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCookiesFromUrl_cookieList_atList", i}).(*QNetworkCookie)
}

func (ptr *QNetworkCookieJar) __setCookiesFromUrl_cookieList_setList(i QNetworkCookie_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCookiesFromUrl_cookieList_setList", i})
}

func (ptr *QNetworkCookieJar) __setCookiesFromUrl_cookieList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCookiesFromUrl_cookieList_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkCookieJar) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QNetworkCookieJar) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNetworkCookieJar) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkCookieJar) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QNetworkCookieJar) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QNetworkCookieJar) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkCookieJar) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QNetworkCookieJar) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QNetworkCookieJar) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkCookieJar) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QNetworkCookieJar) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QNetworkCookieJar) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QNetworkCookieJar) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QNetworkCookieJar) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QNetworkCookieJar) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QNetworkCookieJar) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QNetworkCookieJar) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QNetworkCookieJar) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QNetworkCookieJar) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QNetworkCookieJar) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QNetworkCookieJar) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QNetworkDatagram struct {
	internal.Internal
}

type QNetworkDatagram_ITF interface {
	QNetworkDatagram_PTR() *QNetworkDatagram
}

func (ptr *QNetworkDatagram) QNetworkDatagram_PTR() *QNetworkDatagram {
	return ptr
}

func (ptr *QNetworkDatagram) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkDatagram) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkDatagram(ptr QNetworkDatagram_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkDatagram_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkDatagram) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkDatagramFromPointer(ptr unsafe.Pointer) (n *QNetworkDatagram) {
	n = new(QNetworkDatagram)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkDatagram")
	return
}

func (ptr *QNetworkDatagram) DestroyQNetworkDatagram() {
}

func NewQNetworkDatagram() *QNetworkDatagram {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkDatagram", ""}).(*QNetworkDatagram)
}

func NewQNetworkDatagram2(data core.QByteArray_ITF, destinationAddress QHostAddress_ITF, port uint16) *QNetworkDatagram {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkDatagram2", "", data, destinationAddress, port}).(*QNetworkDatagram)
}

func NewQNetworkDatagram3(other QNetworkDatagram_ITF) *QNetworkDatagram {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkDatagram3", "", other}).(*QNetworkDatagram)
}

func (ptr *QNetworkDatagram) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QNetworkDatagram) Data() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data"}).(*core.QByteArray)
}

func (ptr *QNetworkDatagram) DestinationAddress() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestinationAddress"}).(*QHostAddress)
}

func (ptr *QNetworkDatagram) DestinationPort() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestinationPort"}).(float64))
}

func (ptr *QNetworkDatagram) HopLimit() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HopLimit"}).(float64))
}

func (ptr *QNetworkDatagram) InterfaceIndex() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InterfaceIndex"}).(float64))
}

func (ptr *QNetworkDatagram) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QNetworkDatagram) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QNetworkDatagram) MakeReply(payload core.QByteArray_ITF) *QNetworkDatagram {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MakeReply", payload}).(*QNetworkDatagram)
}

func (ptr *QNetworkDatagram) MakeReply2(payload core.QByteArray_ITF) *QNetworkDatagram {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MakeReply2", payload}).(*QNetworkDatagram)
}

func (ptr *QNetworkDatagram) SenderAddress() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SenderAddress"}).(*QHostAddress)
}

func (ptr *QNetworkDatagram) SenderPort() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SenderPort"}).(float64))
}

func (ptr *QNetworkDatagram) SetData(data core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetData", data})
}

func (ptr *QNetworkDatagram) SetDestination(address QHostAddress_ITF, port uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDestination", address, port})
}

func (ptr *QNetworkDatagram) SetHopLimit(count int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHopLimit", count})
}

func (ptr *QNetworkDatagram) SetInterfaceIndex(index uint) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetInterfaceIndex", index})
}

func (ptr *QNetworkDatagram) SetSender(address QHostAddress_ITF, port uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSender", address, port})
}

func (ptr *QNetworkDatagram) Swap(other QNetworkDatagram_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

type QNetworkDiskCache struct {
	QAbstractNetworkCache
}

type QNetworkDiskCache_ITF interface {
	QAbstractNetworkCache_ITF
	QNetworkDiskCache_PTR() *QNetworkDiskCache
}

func (ptr *QNetworkDiskCache) QNetworkDiskCache_PTR() *QNetworkDiskCache {
	return ptr
}

func (ptr *QNetworkDiskCache) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNetworkCache_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkDiskCache) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractNetworkCache_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkDiskCache(ptr QNetworkDiskCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkDiskCache_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkDiskCache) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractNetworkCache_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNetworkDiskCache) ClassNameInternalF() string {
	return n.QAbstractNetworkCache_PTR().ClassNameInternalF()
}

func NewQNetworkDiskCacheFromPointer(ptr unsafe.Pointer) (n *QNetworkDiskCache) {
	n = new(QNetworkDiskCache)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkDiskCache")
	return
}
func NewQNetworkDiskCache(parent core.QObject_ITF) *QNetworkDiskCache {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkDiskCache", "", parent}).(*QNetworkDiskCache)
}

func (ptr *QNetworkDiskCache) CacheDirectory() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CacheDirectory"}).(string)
}

func (ptr *QNetworkDiskCache) ConnectCacheSize(f func() int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCacheSize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectCacheSize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCacheSize"})
}

func (ptr *QNetworkDiskCache) CacheSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CacheSize"}).(float64))
}

func (ptr *QNetworkDiskCache) CacheSizeDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CacheSizeDefault"}).(float64))
}

func (ptr *QNetworkDiskCache) ConnectClear(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClear", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectClear() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClear"})
}

func (ptr *QNetworkDiskCache) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QNetworkDiskCache) ClearDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ClearDefault"})
}

func (ptr *QNetworkDiskCache) ConnectData(f func(url *core.QUrl) *core.QIODevice) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectData"})
}

func (ptr *QNetworkDiskCache) Data(url core.QUrl_ITF) *core.QIODevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Data", url}).(*core.QIODevice)
}

func (ptr *QNetworkDiskCache) DataDefault(url core.QUrl_ITF) *core.QIODevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DataDefault", url}).(*core.QIODevice)
}

func (ptr *QNetworkDiskCache) ConnectExpire(f func() int64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectExpire", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectExpire() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectExpire"})
}

func (ptr *QNetworkDiskCache) Expire() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Expire"}).(float64))
}

func (ptr *QNetworkDiskCache) ExpireDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExpireDefault"}).(float64))
}

func (ptr *QNetworkDiskCache) FileMetaData(fileName string) *QNetworkCacheMetaData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "FileMetaData", fileName}).(*QNetworkCacheMetaData)
}

func (ptr *QNetworkDiskCache) ConnectInsert(f func(device *core.QIODevice)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectInsert", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectInsert() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectInsert"})
}

func (ptr *QNetworkDiskCache) Insert(device core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Insert", device})
}

func (ptr *QNetworkDiskCache) InsertDefault(device core.QIODevice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "InsertDefault", device})
}

func (ptr *QNetworkDiskCache) MaximumCacheSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumCacheSize"}).(float64))
}

func (ptr *QNetworkDiskCache) ConnectMetaData(f func(url *core.QUrl) *QNetworkCacheMetaData) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMetaData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectMetaData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMetaData"})
}

func (ptr *QNetworkDiskCache) MetaData(url core.QUrl_ITF) *QNetworkCacheMetaData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaData", url}).(*QNetworkCacheMetaData)
}

func (ptr *QNetworkDiskCache) MetaDataDefault(url core.QUrl_ITF) *QNetworkCacheMetaData {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaDataDefault", url}).(*QNetworkCacheMetaData)
}

func (ptr *QNetworkDiskCache) ConnectPrepare(f func(metaData *QNetworkCacheMetaData) *core.QIODevice) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPrepare", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectPrepare() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPrepare"})
}

func (ptr *QNetworkDiskCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Prepare", metaData}).(*core.QIODevice)
}

func (ptr *QNetworkDiskCache) PrepareDefault(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrepareDefault", metaData}).(*core.QIODevice)
}

func (ptr *QNetworkDiskCache) ConnectRemove(f func(url *core.QUrl) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRemove", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectRemove() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRemove"})
}

func (ptr *QNetworkDiskCache) Remove(url core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Remove", url}).(bool)
}

func (ptr *QNetworkDiskCache) RemoveDefault(url core.QUrl_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RemoveDefault", url}).(bool)
}

func (ptr *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCacheDirectory", cacheDir})
}

func (ptr *QNetworkDiskCache) SetMaximumCacheSize(size int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaximumCacheSize", size})
}

func (ptr *QNetworkDiskCache) ConnectUpdateMetaData(f func(metaData *QNetworkCacheMetaData)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUpdateMetaData", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectUpdateMetaData() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUpdateMetaData"})
}

func (ptr *QNetworkDiskCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMetaData", metaData})
}

func (ptr *QNetworkDiskCache) UpdateMetaDataDefault(metaData QNetworkCacheMetaData_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UpdateMetaDataDefault", metaData})
}

func (ptr *QNetworkDiskCache) ConnectDestroyQNetworkDiskCache(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNetworkDiskCache", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkDiskCache) DisconnectDestroyQNetworkDiskCache() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNetworkDiskCache"})
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCache() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkDiskCache"})
}

func (ptr *QNetworkDiskCache) DestroyQNetworkDiskCacheDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkDiskCacheDefault"})
}

type QNetworkInterface struct {
	internal.Internal
}

type QNetworkInterface_ITF interface {
	QNetworkInterface_PTR() *QNetworkInterface
}

func (ptr *QNetworkInterface) QNetworkInterface_PTR() *QNetworkInterface {
	return ptr
}

func (ptr *QNetworkInterface) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkInterface) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkInterface(ptr QNetworkInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkInterface_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkInterface) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkInterfaceFromPointer(ptr unsafe.Pointer) (n *QNetworkInterface) {
	n = new(QNetworkInterface)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkInterface")
	return
}

//go:generate stringer -type=QNetworkInterface__InterfaceType
//QNetworkInterface::InterfaceType
type QNetworkInterface__InterfaceType int64

const (
	QNetworkInterface__Loopback   QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(1)
	QNetworkInterface__Virtual    QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(2)
	QNetworkInterface__Ethernet   QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(3)
	QNetworkInterface__Slip       QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(4)
	QNetworkInterface__CanBus     QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(5)
	QNetworkInterface__Ppp        QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(6)
	QNetworkInterface__Fddi       QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(7)
	QNetworkInterface__Wifi       QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(8)
	QNetworkInterface__Ieee80211  QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(QNetworkInterface__Wifi)
	QNetworkInterface__Phonet     QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(9)
	QNetworkInterface__Ieee802154 QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(10)
	QNetworkInterface__SixLoWPAN  QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(11)
	QNetworkInterface__Ieee80216  QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(12)
	QNetworkInterface__Ieee1394   QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(13)
	QNetworkInterface__Unknown    QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(0)
)

//go:generate stringer -type=QNetworkInterface__InterfaceFlag
//QNetworkInterface::InterfaceFlag
type QNetworkInterface__InterfaceFlag int64

const (
	QNetworkInterface__IsUp           QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x1)
	QNetworkInterface__IsRunning      QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x2)
	QNetworkInterface__CanBroadcast   QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x4)
	QNetworkInterface__IsLoopBack     QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x8)
	QNetworkInterface__IsPointToPoint QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x10)
	QNetworkInterface__CanMulticast   QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x20)
)

func NewQNetworkInterface() *QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkInterface", ""}).(*QNetworkInterface)
}

func NewQNetworkInterface2(other QNetworkInterface_ITF) *QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkInterface2", "", other}).(*QNetworkInterface)
}

func (ptr *QNetworkInterface) AddressEntries() []*QNetworkAddressEntry {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddressEntries"}).([]*QNetworkAddressEntry)
}

func QNetworkInterface_AllAddresses() []*QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_AllAddresses", ""}).([]*QHostAddress)
}

func (ptr *QNetworkInterface) AllAddresses() []*QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_AllAddresses", ""}).([]*QHostAddress)
}

func QNetworkInterface_AllInterfaces() []*QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_AllInterfaces", ""}).([]*QNetworkInterface)
}

func (ptr *QNetworkInterface) AllInterfaces() []*QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_AllInterfaces", ""}).([]*QNetworkInterface)
}

func (ptr *QNetworkInterface) Flags() QNetworkInterface__InterfaceFlag {

	return QNetworkInterface__InterfaceFlag(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flags"}).(float64))
}

func (ptr *QNetworkInterface) HardwareAddress() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HardwareAddress"}).(string)
}

func (ptr *QNetworkInterface) HumanReadableName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HumanReadableName"}).(string)
}

func (ptr *QNetworkInterface) Index() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Index"}).(float64))
}

func QNetworkInterface_InterfaceFromIndex(index int) *QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_InterfaceFromIndex", "", index}).(*QNetworkInterface)
}

func (ptr *QNetworkInterface) InterfaceFromIndex(index int) *QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_InterfaceFromIndex", "", index}).(*QNetworkInterface)
}

func QNetworkInterface_InterfaceFromName(name string) *QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_InterfaceFromName", "", name}).(*QNetworkInterface)
}

func (ptr *QNetworkInterface) InterfaceFromName(name string) *QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_InterfaceFromName", "", name}).(*QNetworkInterface)
}

func QNetworkInterface_InterfaceIndexFromName(name string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_InterfaceIndexFromName", "", name}).(float64))
}

func (ptr *QNetworkInterface) InterfaceIndexFromName(name string) int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_InterfaceIndexFromName", "", name}).(float64))
}

func QNetworkInterface_InterfaceNameFromIndex(index int) string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_InterfaceNameFromIndex", "", index}).(string)
}

func (ptr *QNetworkInterface) InterfaceNameFromIndex(index int) string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkInterface_InterfaceNameFromIndex", "", index}).(string)
}

func (ptr *QNetworkInterface) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QNetworkInterface) MaximumTransmissionUnit() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumTransmissionUnit"}).(float64))
}

func (ptr *QNetworkInterface) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QNetworkInterface) Swap(other QNetworkInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QNetworkInterface) Type() QNetworkInterface__InterfaceType {

	return QNetworkInterface__InterfaceType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QNetworkInterface) DestroyQNetworkInterface() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkInterface"})
}

func (ptr *QNetworkInterface) __addressEntries_atList(i int) *QNetworkAddressEntry {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addressEntries_atList", i}).(*QNetworkAddressEntry)
}

func (ptr *QNetworkInterface) __addressEntries_setList(i QNetworkAddressEntry_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addressEntries_setList", i})
}

func (ptr *QNetworkInterface) __addressEntries_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addressEntries_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkInterface) __allAddresses_atList(i int) *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allAddresses_atList", i}).(*QHostAddress)
}

func (ptr *QNetworkInterface) __allAddresses_setList(i QHostAddress_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allAddresses_setList", i})
}

func (ptr *QNetworkInterface) __allAddresses_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allAddresses_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkInterface) __allInterfaces_atList(i int) *QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allInterfaces_atList", i}).(*QNetworkInterface)
}

func (ptr *QNetworkInterface) __allInterfaces_setList(i QNetworkInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allInterfaces_setList", i})
}

func (ptr *QNetworkInterface) __allInterfaces_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allInterfaces_newList"}).(unsafe.Pointer)
}

type QNetworkProxy struct {
	internal.Internal
}

type QNetworkProxy_ITF interface {
	QNetworkProxy_PTR() *QNetworkProxy
}

func (ptr *QNetworkProxy) QNetworkProxy_PTR() *QNetworkProxy {
	return ptr
}

func (ptr *QNetworkProxy) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkProxy) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkProxy(ptr QNetworkProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxy_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkProxy) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkProxyFromPointer(ptr unsafe.Pointer) (n *QNetworkProxy) {
	n = new(QNetworkProxy)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkProxy")
	return
}

//go:generate stringer -type=QNetworkProxy__ProxyType
//QNetworkProxy::ProxyType
type QNetworkProxy__ProxyType int64

const (
	QNetworkProxy__DefaultProxy     QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(0)
	QNetworkProxy__Socks5Proxy      QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(1)
	QNetworkProxy__NoProxy          QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(2)
	QNetworkProxy__HttpProxy        QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(3)
	QNetworkProxy__HttpCachingProxy QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(4)
	QNetworkProxy__FtpCachingProxy  QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(5)
)

//go:generate stringer -type=QNetworkProxy__Capability
//QNetworkProxy::Capability
type QNetworkProxy__Capability int64

const (
	QNetworkProxy__TunnelingCapability      QNetworkProxy__Capability = QNetworkProxy__Capability(0x0001)
	QNetworkProxy__ListeningCapability      QNetworkProxy__Capability = QNetworkProxy__Capability(0x0002)
	QNetworkProxy__UdpTunnelingCapability   QNetworkProxy__Capability = QNetworkProxy__Capability(0x0004)
	QNetworkProxy__CachingCapability        QNetworkProxy__Capability = QNetworkProxy__Capability(0x0008)
	QNetworkProxy__HostNameLookupCapability QNetworkProxy__Capability = QNetworkProxy__Capability(0x0010)
	QNetworkProxy__SctpTunnelingCapability  QNetworkProxy__Capability = QNetworkProxy__Capability(0x00020)
	QNetworkProxy__SctpListeningCapability  QNetworkProxy__Capability = QNetworkProxy__Capability(0x00040)
)

func NewQNetworkProxy() *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkProxy", ""}).(*QNetworkProxy)
}

func NewQNetworkProxy2(ty QNetworkProxy__ProxyType, hostName string, port uint16, user string, password string) *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkProxy2", "", ty, hostName, port, user, password}).(*QNetworkProxy)
}

func NewQNetworkProxy3(other QNetworkProxy_ITF) *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkProxy3", "", other}).(*QNetworkProxy)
}

func QNetworkProxy_ApplicationProxy() *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxy_ApplicationProxy", ""}).(*QNetworkProxy)
}

func (ptr *QNetworkProxy) ApplicationProxy() *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxy_ApplicationProxy", ""}).(*QNetworkProxy)
}

func (ptr *QNetworkProxy) Capabilities() QNetworkProxy__Capability {

	return QNetworkProxy__Capability(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Capabilities"}).(float64))
}

func (ptr *QNetworkProxy) HasRawHeader(headerName core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasRawHeader", headerName}).(bool)
}

func (ptr *QNetworkProxy) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Header", header}).(*core.QVariant)
}

func (ptr *QNetworkProxy) HostName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HostName"}).(string)
}

func (ptr *QNetworkProxy) IsCachingProxy() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCachingProxy"}).(bool)
}

func (ptr *QNetworkProxy) IsTransparentProxy() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsTransparentProxy"}).(bool)
}

func (ptr *QNetworkProxy) Password() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Password"}).(string)
}

func (ptr *QNetworkProxy) Port() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Port"}).(float64))
}

func (ptr *QNetworkProxy) QNetworkCacheMetaData_RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QNetworkCacheMetaData_RawHeader", headerName}).(*core.QByteArray)
}

func (ptr *QNetworkProxy) RawHeaderList() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RawHeaderList"}).([]*core.QByteArray)
}

func QNetworkProxy_SetApplicationProxy(networkProxy QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxy_SetApplicationProxy", "", networkProxy})
}

func (ptr *QNetworkProxy) SetApplicationProxy(networkProxy QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxy_SetApplicationProxy", "", networkProxy})
}

func (ptr *QNetworkProxy) SetCapabilities(capabilities QNetworkProxy__Capability) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCapabilities", capabilities})
}

func (ptr *QNetworkProxy) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeader", header, value})
}

func (ptr *QNetworkProxy) SetHostName(hostName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHostName", hostName})
}

func (ptr *QNetworkProxy) SetPassword(password string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPassword", password})
}

func (ptr *QNetworkProxy) SetPort(port uint16) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPort", port})
}

func (ptr *QNetworkProxy) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRawHeader", headerName, headerValue})
}

func (ptr *QNetworkProxy) SetType(ty QNetworkProxy__ProxyType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetType", ty})
}

func (ptr *QNetworkProxy) SetUser(user string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUser", user})
}

func (ptr *QNetworkProxy) Swap(other QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QNetworkProxy) Type() QNetworkProxy__ProxyType {

	return QNetworkProxy__ProxyType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QNetworkProxy) User() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "User"}).(string)
}

func (ptr *QNetworkProxy) DestroyQNetworkProxy() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkProxy"})
}

func (ptr *QNetworkProxy) __rawHeaderList_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaderList_atList", i}).(*core.QByteArray)
}

func (ptr *QNetworkProxy) __rawHeaderList_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaderList_setList", i})
}

func (ptr *QNetworkProxy) __rawHeaderList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaderList_newList"}).(unsafe.Pointer)
}

type QNetworkProxyFactory struct {
	internal.Internal
}

type QNetworkProxyFactory_ITF interface {
	QNetworkProxyFactory_PTR() *QNetworkProxyFactory
}

func (ptr *QNetworkProxyFactory) QNetworkProxyFactory_PTR() *QNetworkProxyFactory {
	return ptr
}

func (ptr *QNetworkProxyFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkProxyFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkProxyFactory(ptr QNetworkProxyFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyFactory_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkProxyFactory) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkProxyFactoryFromPointer(ptr unsafe.Pointer) (n *QNetworkProxyFactory) {
	n = new(QNetworkProxyFactory)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkProxyFactory")
	return
}
func NewQNetworkProxyFactory() *QNetworkProxyFactory {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkProxyFactory", ""}).(*QNetworkProxyFactory)
}

func QNetworkProxyFactory_ProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_ProxyForQuery", "", query}).([]*QNetworkProxy)
}

func (ptr *QNetworkProxyFactory) ProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_ProxyForQuery", "", query}).([]*QNetworkProxy)
}

func (ptr *QNetworkProxyFactory) ConnectQueryProxy(f func(query *QNetworkProxyQuery) []*QNetworkProxy) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectQueryProxy", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkProxyFactory) DisconnectQueryProxy() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectQueryProxy"})
}

func (ptr *QNetworkProxyFactory) QueryProxy(query QNetworkProxyQuery_ITF) []*QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QueryProxy", query}).([]*QNetworkProxy)
}

func QNetworkProxyFactory_SetApplicationProxyFactory(factory QNetworkProxyFactory_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_SetApplicationProxyFactory", "", factory})
}

func (ptr *QNetworkProxyFactory) SetApplicationProxyFactory(factory QNetworkProxyFactory_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_SetApplicationProxyFactory", "", factory})
}

func QNetworkProxyFactory_SetUseSystemConfiguration(enable bool) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_SetUseSystemConfiguration", "", enable})
}

func (ptr *QNetworkProxyFactory) SetUseSystemConfiguration(enable bool) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_SetUseSystemConfiguration", "", enable})
}

func QNetworkProxyFactory_SystemProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_SystemProxyForQuery", "", query}).([]*QNetworkProxy)
}

func (ptr *QNetworkProxyFactory) SystemProxyForQuery(query QNetworkProxyQuery_ITF) []*QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_SystemProxyForQuery", "", query}).([]*QNetworkProxy)
}

func QNetworkProxyFactory_UsesSystemConfiguration() bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_UsesSystemConfiguration", ""}).(bool)
}

func (ptr *QNetworkProxyFactory) UsesSystemConfiguration() bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QNetworkProxyFactory_UsesSystemConfiguration", ""}).(bool)
}

func (ptr *QNetworkProxyFactory) ConnectDestroyQNetworkProxyFactory(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNetworkProxyFactory", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkProxyFactory) DisconnectDestroyQNetworkProxyFactory() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNetworkProxyFactory"})
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactory() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkProxyFactory"})
}

func (ptr *QNetworkProxyFactory) DestroyQNetworkProxyFactoryDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkProxyFactoryDefault"})
}

func (ptr *QNetworkProxyFactory) __proxyForQuery_atList(i int) *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__proxyForQuery_atList", i}).(*QNetworkProxy)
}

func (ptr *QNetworkProxyFactory) __proxyForQuery_setList(i QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__proxyForQuery_setList", i})
}

func (ptr *QNetworkProxyFactory) __proxyForQuery_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__proxyForQuery_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkProxyFactory) __queryProxy_atList(i int) *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__queryProxy_atList", i}).(*QNetworkProxy)
}

func (ptr *QNetworkProxyFactory) __queryProxy_setList(i QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__queryProxy_setList", i})
}

func (ptr *QNetworkProxyFactory) __queryProxy_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__queryProxy_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkProxyFactory) __systemProxyForQuery_atList(i int) *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__systemProxyForQuery_atList", i}).(*QNetworkProxy)
}

func (ptr *QNetworkProxyFactory) __systemProxyForQuery_setList(i QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__systemProxyForQuery_setList", i})
}

func (ptr *QNetworkProxyFactory) __systemProxyForQuery_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__systemProxyForQuery_newList"}).(unsafe.Pointer)
}

type QNetworkProxyQuery struct {
	internal.Internal
}

type QNetworkProxyQuery_ITF interface {
	QNetworkProxyQuery_PTR() *QNetworkProxyQuery
}

func (ptr *QNetworkProxyQuery) QNetworkProxyQuery_PTR() *QNetworkProxyQuery {
	return ptr
}

func (ptr *QNetworkProxyQuery) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkProxyQuery) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkProxyQuery(ptr QNetworkProxyQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyQuery_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkProxyQuery) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkProxyQueryFromPointer(ptr unsafe.Pointer) (n *QNetworkProxyQuery) {
	n = new(QNetworkProxyQuery)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkProxyQuery")
	return
}

//go:generate stringer -type=QNetworkProxyQuery__QueryType
//QNetworkProxyQuery::QueryType
type QNetworkProxyQuery__QueryType int64

const (
	QNetworkProxyQuery__TcpSocket  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(0)
	QNetworkProxyQuery__UdpSocket  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(1)
	QNetworkProxyQuery__SctpSocket QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(2)
	QNetworkProxyQuery__TcpServer  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(100)
	QNetworkProxyQuery__UrlRequest QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(101)
	QNetworkProxyQuery__SctpServer QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(102)
)

func NewQNetworkProxyQuery() *QNetworkProxyQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkProxyQuery", ""}).(*QNetworkProxyQuery)
}

func NewQNetworkProxyQuery2(requestUrl core.QUrl_ITF, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkProxyQuery2", "", requestUrl, queryType}).(*QNetworkProxyQuery)
}

func NewQNetworkProxyQuery3(hostname string, port int, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkProxyQuery3", "", hostname, port, protocolTag, queryType}).(*QNetworkProxyQuery)
}

func NewQNetworkProxyQuery4(bindPort uint16, protocolTag string, queryType QNetworkProxyQuery__QueryType) *QNetworkProxyQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkProxyQuery4", "", bindPort, protocolTag, queryType}).(*QNetworkProxyQuery)
}

func NewQNetworkProxyQuery8(other QNetworkProxyQuery_ITF) *QNetworkProxyQuery {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkProxyQuery8", "", other}).(*QNetworkProxyQuery)
}

func (ptr *QNetworkProxyQuery) LocalPort() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalPort"}).(float64))
}

func (ptr *QNetworkProxyQuery) PeerHostName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerHostName"}).(string)
}

func (ptr *QNetworkProxyQuery) PeerPort() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerPort"}).(float64))
}

func (ptr *QNetworkProxyQuery) ProtocolTag() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProtocolTag"}).(string)
}

func (ptr *QNetworkProxyQuery) QueryType() QNetworkProxyQuery__QueryType {

	return QNetworkProxyQuery__QueryType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QueryType"}).(float64))
}

func (ptr *QNetworkProxyQuery) SetLocalPort(port int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalPort", port})
}

func (ptr *QNetworkProxyQuery) SetPeerHostName(hostname string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerHostName", hostname})
}

func (ptr *QNetworkProxyQuery) SetPeerPort(port int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerPort", port})
}

func (ptr *QNetworkProxyQuery) SetProtocolTag(protocolTag string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProtocolTag", protocolTag})
}

func (ptr *QNetworkProxyQuery) SetQueryType(ty QNetworkProxyQuery__QueryType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetQueryType", ty})
}

func (ptr *QNetworkProxyQuery) SetUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrl", url})
}

func (ptr *QNetworkProxyQuery) Swap(other QNetworkProxyQuery_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QNetworkProxyQuery) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QNetworkProxyQuery) DestroyQNetworkProxyQuery() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkProxyQuery"})
}

type QNetworkReply struct {
	core.QIODevice
}

type QNetworkReply_ITF interface {
	core.QIODevice_ITF
	QNetworkReply_PTR() *QNetworkReply
}

func (ptr *QNetworkReply) QNetworkReply_PTR() *QNetworkReply {
	return ptr
}

func (ptr *QNetworkReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkReply(ptr QNetworkReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkReply_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkReply) InitFromInternal(ptr uintptr, name string) {
	n.QIODevice_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNetworkReply) ClassNameInternalF() string {
	return n.QIODevice_PTR().ClassNameInternalF()
}

func NewQNetworkReplyFromPointer(ptr unsafe.Pointer) (n *QNetworkReply) {
	n = new(QNetworkReply)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkReply")
	return
}

//go:generate stringer -type=QNetworkReply__NetworkError
//QNetworkReply::NetworkError
type QNetworkReply__NetworkError int64

const (
	QNetworkReply__NoError                           QNetworkReply__NetworkError = QNetworkReply__NetworkError(0)
	QNetworkReply__ConnectionRefusedError            QNetworkReply__NetworkError = QNetworkReply__NetworkError(1)
	QNetworkReply__RemoteHostClosedError             QNetworkReply__NetworkError = QNetworkReply__NetworkError(2)
	QNetworkReply__HostNotFoundError                 QNetworkReply__NetworkError = QNetworkReply__NetworkError(3)
	QNetworkReply__TimeoutError                      QNetworkReply__NetworkError = QNetworkReply__NetworkError(4)
	QNetworkReply__OperationCanceledError            QNetworkReply__NetworkError = QNetworkReply__NetworkError(5)
	QNetworkReply__SslHandshakeFailedError           QNetworkReply__NetworkError = QNetworkReply__NetworkError(6)
	QNetworkReply__TemporaryNetworkFailureError      QNetworkReply__NetworkError = QNetworkReply__NetworkError(7)
	QNetworkReply__NetworkSessionFailedError         QNetworkReply__NetworkError = QNetworkReply__NetworkError(8)
	QNetworkReply__BackgroundRequestNotAllowedError  QNetworkReply__NetworkError = QNetworkReply__NetworkError(9)
	QNetworkReply__TooManyRedirectsError             QNetworkReply__NetworkError = QNetworkReply__NetworkError(10)
	QNetworkReply__InsecureRedirectError             QNetworkReply__NetworkError = QNetworkReply__NetworkError(11)
	QNetworkReply__UnknownNetworkError               QNetworkReply__NetworkError = QNetworkReply__NetworkError(99)
	QNetworkReply__ProxyConnectionRefusedError       QNetworkReply__NetworkError = QNetworkReply__NetworkError(101)
	QNetworkReply__ProxyConnectionClosedError        QNetworkReply__NetworkError = QNetworkReply__NetworkError(102)
	QNetworkReply__ProxyNotFoundError                QNetworkReply__NetworkError = QNetworkReply__NetworkError(103)
	QNetworkReply__ProxyTimeoutError                 QNetworkReply__NetworkError = QNetworkReply__NetworkError(104)
	QNetworkReply__ProxyAuthenticationRequiredError  QNetworkReply__NetworkError = QNetworkReply__NetworkError(105)
	QNetworkReply__UnknownProxyError                 QNetworkReply__NetworkError = QNetworkReply__NetworkError(199)
	QNetworkReply__ContentAccessDenied               QNetworkReply__NetworkError = QNetworkReply__NetworkError(201)
	QNetworkReply__ContentOperationNotPermittedError QNetworkReply__NetworkError = QNetworkReply__NetworkError(202)
	QNetworkReply__ContentNotFoundError              QNetworkReply__NetworkError = QNetworkReply__NetworkError(203)
	QNetworkReply__AuthenticationRequiredError       QNetworkReply__NetworkError = QNetworkReply__NetworkError(204)
	QNetworkReply__ContentReSendError                QNetworkReply__NetworkError = QNetworkReply__NetworkError(205)
	QNetworkReply__ContentConflictError              QNetworkReply__NetworkError = QNetworkReply__NetworkError(206)
	QNetworkReply__ContentGoneError                  QNetworkReply__NetworkError = QNetworkReply__NetworkError(207)
	QNetworkReply__UnknownContentError               QNetworkReply__NetworkError = QNetworkReply__NetworkError(299)
	QNetworkReply__ProtocolUnknownError              QNetworkReply__NetworkError = QNetworkReply__NetworkError(301)
	QNetworkReply__ProtocolInvalidOperationError     QNetworkReply__NetworkError = QNetworkReply__NetworkError(302)
	QNetworkReply__ProtocolFailure                   QNetworkReply__NetworkError = QNetworkReply__NetworkError(399)
	QNetworkReply__InternalServerError               QNetworkReply__NetworkError = QNetworkReply__NetworkError(401)
	QNetworkReply__OperationNotImplementedError      QNetworkReply__NetworkError = QNetworkReply__NetworkError(402)
	QNetworkReply__ServiceUnavailableError           QNetworkReply__NetworkError = QNetworkReply__NetworkError(403)
	QNetworkReply__UnknownServerError                QNetworkReply__NetworkError = QNetworkReply__NetworkError(499)
)

func NewQNetworkReply(parent core.QObject_ITF) *QNetworkReply {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkReply", "", parent}).(*QNetworkReply)
}

func (ptr *QNetworkReply) ConnectAbort(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAbort", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectAbort() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAbort"})
}

func (ptr *QNetworkReply) Abort() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"})
}

func (ptr *QNetworkReply) Attribute(code QNetworkRequest__Attribute) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Attribute", code}).(*core.QVariant)
}

func (ptr *QNetworkReply) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QNetworkReply) ConnectDownloadProgress(f func(bytesReceived int64, bytesTotal int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDownloadProgress", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectDownloadProgress() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDownloadProgress"})
}

func (ptr *QNetworkReply) DownloadProgress(bytesReceived int64, bytesTotal int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DownloadProgress", bytesReceived, bytesTotal})
}

func (ptr *QNetworkReply) ConnectEncrypted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEncrypted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectEncrypted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEncrypted"})
}

func (ptr *QNetworkReply) Encrypted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Encrypted"})
}

func (ptr *QNetworkReply) Error() QNetworkReply__NetworkError {

	return QNetworkReply__NetworkError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QNetworkReply) ConnectError2(f func(code QNetworkReply__NetworkError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QNetworkReply) Error2(code QNetworkReply__NetworkError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", code})
}

func (ptr *QNetworkReply) ConnectFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectFinished"})
}

func (ptr *QNetworkReply) Finished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Finished"})
}

func (ptr *QNetworkReply) HasRawHeader(headerName core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasRawHeader", headerName}).(bool)
}

func (ptr *QNetworkReply) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Header", header}).(*core.QVariant)
}

func (ptr *QNetworkReply) ConnectIgnoreSslErrors(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIgnoreSslErrors", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectIgnoreSslErrors() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIgnoreSslErrors"})
}

func (ptr *QNetworkReply) IgnoreSslErrors() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrors"})
}

func (ptr *QNetworkReply) IgnoreSslErrorsDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrorsDefault"})
}

func (ptr *QNetworkReply) IgnoreSslErrors2(errors []*QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrors2", errors})
}

func (ptr *QNetworkReply) ConnectIgnoreSslErrorsImplementation(f func(errors []*QSslError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIgnoreSslErrorsImplementation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectIgnoreSslErrorsImplementation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIgnoreSslErrorsImplementation"})
}

func (ptr *QNetworkReply) IgnoreSslErrorsImplementation(errors []*QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrorsImplementation", errors})
}

func (ptr *QNetworkReply) IgnoreSslErrorsImplementationDefault(errors []*QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrorsImplementationDefault", errors})
}

func (ptr *QNetworkReply) IsFinished() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsFinished"}).(bool)
}

func (ptr *QNetworkReply) IsRunning() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsRunning"}).(bool)
}

func (ptr *QNetworkReply) Manager() *QNetworkAccessManager {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Manager"}).(*QNetworkAccessManager)
}

func (ptr *QNetworkReply) ConnectMetaDataChanged(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMetaDataChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectMetaDataChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMetaDataChanged"})
}

func (ptr *QNetworkReply) MetaDataChanged() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaDataChanged"})
}

func (ptr *QNetworkReply) Operation() QNetworkAccessManager__Operation {

	return QNetworkAccessManager__Operation(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Operation"}).(float64))
}

func (ptr *QNetworkReply) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreSharedKeyAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectPreSharedKeyAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreSharedKeyAuthenticationRequired"})
}

func (ptr *QNetworkReply) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreSharedKeyAuthenticationRequired", authenticator})
}

func (ptr *QNetworkReply) QNetworkCacheMetaData_RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QNetworkCacheMetaData_RawHeader", headerName}).(*core.QByteArray)
}

func (ptr *QNetworkReply) RawHeaderList() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RawHeaderList"}).([]*core.QByteArray)
}

func (ptr *QNetworkReply) ReadBufferSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadBufferSize"}).(float64))
}

func (ptr *QNetworkReply) ConnectRedirectAllowed(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRedirectAllowed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectRedirectAllowed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRedirectAllowed"})
}

func (ptr *QNetworkReply) RedirectAllowed() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RedirectAllowed"})
}

func (ptr *QNetworkReply) ConnectRedirected(f func(url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRedirected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectRedirected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRedirected"})
}

func (ptr *QNetworkReply) Redirected(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Redirected", url})
}

func (ptr *QNetworkReply) Request() *QNetworkRequest {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Request"}).(*QNetworkRequest)
}

func (ptr *QNetworkReply) SetAttribute(code QNetworkRequest__Attribute, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute", code, value})
}

func (ptr *QNetworkReply) SetError(errorCode QNetworkReply__NetworkError, errorString string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetError", errorCode, errorString})
}

func (ptr *QNetworkReply) SetFinished(finished bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFinished", finished})
}

func (ptr *QNetworkReply) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeader", header, value})
}

func (ptr *QNetworkReply) SetOperation(operation QNetworkAccessManager__Operation) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOperation", operation})
}

func (ptr *QNetworkReply) SetRawHeader(headerName core.QByteArray_ITF, value core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRawHeader", headerName, value})
}

func (ptr *QNetworkReply) ConnectSetReadBufferSize(f func(size int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetReadBufferSize", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectSetReadBufferSize() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetReadBufferSize"})
}

func (ptr *QNetworkReply) SetReadBufferSize(size int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadBufferSize", size})
}

func (ptr *QNetworkReply) SetReadBufferSizeDefault(size int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadBufferSizeDefault", size})
}

func (ptr *QNetworkReply) SetRequest(request QNetworkRequest_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRequest", request})
}

func (ptr *QNetworkReply) SetSslConfiguration(config QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSslConfiguration", config})
}

func (ptr *QNetworkReply) ConnectSetSslConfigurationImplementation(f func(configuration *QSslConfiguration)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetSslConfigurationImplementation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectSetSslConfigurationImplementation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetSslConfigurationImplementation"})
}

func (ptr *QNetworkReply) SetSslConfigurationImplementation(configuration QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSslConfigurationImplementation", configuration})
}

func (ptr *QNetworkReply) SetSslConfigurationImplementationDefault(configuration QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSslConfigurationImplementationDefault", configuration})
}

func (ptr *QNetworkReply) SetUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrl", url})
}

func (ptr *QNetworkReply) SslConfiguration() *QSslConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslConfiguration"}).(*QSslConfiguration)
}

func (ptr *QNetworkReply) ConnectSslConfigurationImplementation(f func(configuration *QSslConfiguration)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSslConfigurationImplementation", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectSslConfigurationImplementation() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSslConfigurationImplementation"})
}

func (ptr *QNetworkReply) SslConfigurationImplementation(configuration QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslConfigurationImplementation", configuration})
}

func (ptr *QNetworkReply) SslConfigurationImplementationDefault(configuration QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslConfigurationImplementationDefault", configuration})
}

func (ptr *QNetworkReply) ConnectSslErrors(f func(errors []*QSslError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSslErrors", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectSslErrors() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSslErrors"})
}

func (ptr *QNetworkReply) SslErrors(errors []*QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslErrors", errors})
}

func (ptr *QNetworkReply) ConnectUploadProgress(f func(bytesSent int64, bytesTotal int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUploadProgress", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectUploadProgress() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUploadProgress"})
}

func (ptr *QNetworkReply) UploadProgress(bytesSent int64, bytesTotal int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UploadProgress", bytesSent, bytesTotal})
}

func (ptr *QNetworkReply) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QNetworkReply) ConnectDestroyQNetworkReply(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNetworkReply", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkReply) DisconnectDestroyQNetworkReply() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNetworkReply"})
}

func (ptr *QNetworkReply) DestroyQNetworkReply() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkReply"})
}

func (ptr *QNetworkReply) DestroyQNetworkReplyDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkReplyDefault"})
}

func (ptr *QNetworkReply) __ignoreSslErrors_errors_atList2(i int) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrors_errors_atList2", i}).(*QSslError)
}

func (ptr *QNetworkReply) __ignoreSslErrors_errors_setList2(i QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrors_errors_setList2", i})
}

func (ptr *QNetworkReply) __ignoreSslErrors_errors_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrors_errors_newList2"}).(unsafe.Pointer)
}

func (ptr *QNetworkReply) __ignoreSslErrorsImplementation_errors_atList(i int) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrorsImplementation_errors_atList", i}).(*QSslError)
}

func (ptr *QNetworkReply) __ignoreSslErrorsImplementation_errors_setList(i QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrorsImplementation_errors_setList", i})
}

func (ptr *QNetworkReply) __ignoreSslErrorsImplementation_errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrorsImplementation_errors_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkReply) __rawHeaderList_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaderList_atList", i}).(*core.QByteArray)
}

func (ptr *QNetworkReply) __rawHeaderList_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaderList_setList", i})
}

func (ptr *QNetworkReply) __rawHeaderList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaderList_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkReply) __sslErrors_errors_atList(i int) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_atList", i}).(*QSslError)
}

func (ptr *QNetworkReply) __sslErrors_errors_setList(i QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_setList", i})
}

func (ptr *QNetworkReply) __sslErrors_errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkReply) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QNetworkReply) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNetworkReply) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkReply) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QNetworkReply) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QNetworkReply) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkReply) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QNetworkReply) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QNetworkReply) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkReply) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QNetworkReply) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QNetworkReply) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QNetworkReply) AtEndDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AtEndDefault"}).(bool)
}

func (ptr *QNetworkReply) BytesAvailableDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesAvailableDefault"}).(float64))
}

func (ptr *QNetworkReply) BytesToWriteDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesToWriteDefault"}).(float64))
}

func (ptr *QNetworkReply) CanReadLineDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CanReadLineDefault"}).(bool)
}

func (ptr *QNetworkReply) IsSequentialDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSequentialDefault"}).(bool)
}

func (ptr *QNetworkReply) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault", mode}).(bool)
}

func (ptr *QNetworkReply) PosDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PosDefault"}).(float64))
}

func (ptr *QNetworkReply) ReadData(data *string, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadData", data, maxSize}).(float64))
}

func (ptr *QNetworkReply) ReadDataDefault(data *string, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadDataDefault", data, maxSize}).(float64))
}

func (ptr *QNetworkReply) ReadLineDataDefault(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadLineDataDefault", data, maxSize}).(float64))
}

func (ptr *QNetworkReply) ResetDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResetDefault"}).(bool)
}

func (ptr *QNetworkReply) SeekDefault(pos int64) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SeekDefault", pos}).(bool)
}

func (ptr *QNetworkReply) SizeDefault() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SizeDefault"}).(float64))
}

func (ptr *QNetworkReply) WaitForBytesWrittenDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForBytesWrittenDefault", msecs}).(bool)
}

func (ptr *QNetworkReply) WaitForReadyReadDefault(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForReadyReadDefault", msecs}).(bool)
}

func (ptr *QNetworkReply) WriteData(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteData", data, maxSize}).(float64))
}

func (ptr *QNetworkReply) WriteDataDefault(data []byte, maxSize int64) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDataDefault", data, maxSize}).(float64))
}

func (ptr *QNetworkReply) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QNetworkReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QNetworkReply) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QNetworkReply) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QNetworkReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QNetworkReply) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QNetworkReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QNetworkReply) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QNetworkReply) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QNetworkRequest struct {
	internal.Internal
}

type QNetworkRequest_ITF interface {
	QNetworkRequest_PTR() *QNetworkRequest
}

func (ptr *QNetworkRequest) QNetworkRequest_PTR() *QNetworkRequest {
	return ptr
}

func (ptr *QNetworkRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkRequest(ptr QNetworkRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkRequest_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkRequest) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkRequestFromPointer(ptr unsafe.Pointer) (n *QNetworkRequest) {
	n = new(QNetworkRequest)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkRequest")
	return
}

//go:generate stringer -type=QNetworkRequest__Attribute
//QNetworkRequest::Attribute
type QNetworkRequest__Attribute int64

const (
	QNetworkRequest__HttpStatusCodeAttribute               QNetworkRequest__Attribute = QNetworkRequest__Attribute(0)
	QNetworkRequest__HttpReasonPhraseAttribute             QNetworkRequest__Attribute = QNetworkRequest__Attribute(1)
	QNetworkRequest__RedirectionTargetAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(2)
	QNetworkRequest__ConnectionEncryptedAttribute          QNetworkRequest__Attribute = QNetworkRequest__Attribute(3)
	QNetworkRequest__CacheLoadControlAttribute             QNetworkRequest__Attribute = QNetworkRequest__Attribute(4)
	QNetworkRequest__CacheSaveControlAttribute             QNetworkRequest__Attribute = QNetworkRequest__Attribute(5)
	QNetworkRequest__SourceIsFromCacheAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(6)
	QNetworkRequest__DoNotBufferUploadDataAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(7)
	QNetworkRequest__HttpPipeliningAllowedAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(8)
	QNetworkRequest__HttpPipeliningWasUsedAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(9)
	QNetworkRequest__CustomVerbAttribute                   QNetworkRequest__Attribute = QNetworkRequest__Attribute(10)
	QNetworkRequest__CookieLoadControlAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(11)
	QNetworkRequest__AuthenticationReuseAttribute          QNetworkRequest__Attribute = QNetworkRequest__Attribute(12)
	QNetworkRequest__CookieSaveControlAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(13)
	QNetworkRequest__MaximumDownloadBufferSizeAttribute    QNetworkRequest__Attribute = QNetworkRequest__Attribute(14)
	QNetworkRequest__DownloadBufferAttribute               QNetworkRequest__Attribute = QNetworkRequest__Attribute(15)
	QNetworkRequest__SynchronousRequestAttribute           QNetworkRequest__Attribute = QNetworkRequest__Attribute(16)
	QNetworkRequest__BackgroundRequestAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(17)
	QNetworkRequest__SpdyAllowedAttribute                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(18)
	QNetworkRequest__SpdyWasUsedAttribute                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(19)
	QNetworkRequest__EmitAllUploadProgressSignalsAttribute QNetworkRequest__Attribute = QNetworkRequest__Attribute(20)
	QNetworkRequest__FollowRedirectsAttribute              QNetworkRequest__Attribute = QNetworkRequest__Attribute(21)
	QNetworkRequest__HTTP2AllowedAttribute                 QNetworkRequest__Attribute = QNetworkRequest__Attribute(22)
	QNetworkRequest__HTTP2WasUsedAttribute                 QNetworkRequest__Attribute = QNetworkRequest__Attribute(23)
	QNetworkRequest__OriginalContentLengthAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(24)
	QNetworkRequest__RedirectPolicyAttribute               QNetworkRequest__Attribute = QNetworkRequest__Attribute(25)
	QNetworkRequest__Http2DirectAttribute                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(26)
	QNetworkRequest__ResourceTypeAttribute                 QNetworkRequest__Attribute = QNetworkRequest__Attribute(27)
	QNetworkRequest__User                                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(1000)
	QNetworkRequest__UserMax                               QNetworkRequest__Attribute = QNetworkRequest__Attribute(32767)
)

//go:generate stringer -type=QNetworkRequest__KnownHeaders
//QNetworkRequest::KnownHeaders
type QNetworkRequest__KnownHeaders int64

const (
	QNetworkRequest__ContentTypeHeader        QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(0)
	QNetworkRequest__ContentLengthHeader      QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(1)
	QNetworkRequest__LocationHeader           QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(2)
	QNetworkRequest__LastModifiedHeader       QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(3)
	QNetworkRequest__CookieHeader             QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(4)
	QNetworkRequest__SetCookieHeader          QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(5)
	QNetworkRequest__ContentDispositionHeader QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(6)
	QNetworkRequest__UserAgentHeader          QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(7)
	QNetworkRequest__ServerHeader             QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(8)
	QNetworkRequest__IfModifiedSinceHeader    QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(9)
	QNetworkRequest__ETagHeader               QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(10)
	QNetworkRequest__IfMatchHeader            QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(11)
	QNetworkRequest__IfNoneMatchHeader        QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(12)
)

//go:generate stringer -type=QNetworkRequest__CacheLoadControl
//QNetworkRequest::CacheLoadControl
type QNetworkRequest__CacheLoadControl int64

const (
	QNetworkRequest__AlwaysNetwork QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(0)
	QNetworkRequest__PreferNetwork QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(1)
	QNetworkRequest__PreferCache   QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(2)
	QNetworkRequest__AlwaysCache   QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(3)
)

//go:generate stringer -type=QNetworkRequest__LoadControl
//QNetworkRequest::LoadControl
type QNetworkRequest__LoadControl int64

const (
	QNetworkRequest__Automatic QNetworkRequest__LoadControl = QNetworkRequest__LoadControl(0)
	QNetworkRequest__Manual    QNetworkRequest__LoadControl = QNetworkRequest__LoadControl(1)
)

//go:generate stringer -type=QNetworkRequest__Priority
//QNetworkRequest::Priority
type QNetworkRequest__Priority int64

const (
	QNetworkRequest__HighPriority   QNetworkRequest__Priority = QNetworkRequest__Priority(1)
	QNetworkRequest__NormalPriority QNetworkRequest__Priority = QNetworkRequest__Priority(3)
	QNetworkRequest__LowPriority    QNetworkRequest__Priority = QNetworkRequest__Priority(5)
)

//go:generate stringer -type=QNetworkRequest__RedirectPolicy
//QNetworkRequest::RedirectPolicy
type QNetworkRequest__RedirectPolicy int64

const (
	QNetworkRequest__ManualRedirectPolicy       QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(0)
	QNetworkRequest__NoLessSafeRedirectPolicy   QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(1)
	QNetworkRequest__SameOriginRedirectPolicy   QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(2)
	QNetworkRequest__UserVerifiedRedirectPolicy QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(3)
)

func NewQNetworkRequest(url core.QUrl_ITF) *QNetworkRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkRequest", "", url}).(*QNetworkRequest)
}

func NewQNetworkRequest2(other QNetworkRequest_ITF) *QNetworkRequest {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkRequest2", "", other}).(*QNetworkRequest)
}

func (ptr *QNetworkRequest) Attribute(code QNetworkRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Attribute", code, defaultValue}).(*core.QVariant)
}

func (ptr *QNetworkRequest) HasRawHeader(headerName core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasRawHeader", headerName}).(bool)
}

func (ptr *QNetworkRequest) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Header", header}).(*core.QVariant)
}

func (ptr *QNetworkRequest) MaximumRedirectsAllowed() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumRedirectsAllowed"}).(float64))
}

func (ptr *QNetworkRequest) OriginatingObject() *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OriginatingObject"}).(*core.QObject)
}

func (ptr *QNetworkRequest) PeerVerifyName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerVerifyName"}).(string)
}

func (ptr *QNetworkRequest) Priority() QNetworkRequest__Priority {

	return QNetworkRequest__Priority(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Priority"}).(float64))
}

func (ptr *QNetworkRequest) QNetworkCacheMetaData_RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "QNetworkCacheMetaData_RawHeader", headerName}).(*core.QByteArray)
}

func (ptr *QNetworkRequest) RawHeaderList() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RawHeaderList"}).([]*core.QByteArray)
}

func (ptr *QNetworkRequest) SetAttribute(code QNetworkRequest__Attribute, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAttribute", code, value})
}

func (ptr *QNetworkRequest) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetHeader", header, value})
}

func (ptr *QNetworkRequest) SetMaximumRedirectsAllowed(maxRedirectsAllowed int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaximumRedirectsAllowed", maxRedirectsAllowed})
}

func (ptr *QNetworkRequest) SetOriginatingObject(object core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOriginatingObject", object})
}

func (ptr *QNetworkRequest) SetPeerVerifyName(peerName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerVerifyName", peerName})
}

func (ptr *QNetworkRequest) SetPriority(priority QNetworkRequest__Priority) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPriority", priority})
}

func (ptr *QNetworkRequest) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRawHeader", headerName, headerValue})
}

func (ptr *QNetworkRequest) SetSslConfiguration(config QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSslConfiguration", config})
}

func (ptr *QNetworkRequest) SetUrl(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetUrl", url})
}

func (ptr *QNetworkRequest) SslConfiguration() *QSslConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslConfiguration"}).(*QSslConfiguration)
}

func (ptr *QNetworkRequest) Swap(other QNetworkRequest_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QNetworkRequest) Url() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Url"}).(*core.QUrl)
}

func (ptr *QNetworkRequest) DestroyQNetworkRequest() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkRequest"})
}

func (ptr *QNetworkRequest) __rawHeaderList_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaderList_atList", i}).(*core.QByteArray)
}

func (ptr *QNetworkRequest) __rawHeaderList_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaderList_setList", i})
}

func (ptr *QNetworkRequest) __rawHeaderList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__rawHeaderList_newList"}).(unsafe.Pointer)
}

type QNetworkSession struct {
	core.QObject
}

type QNetworkSession_ITF interface {
	core.QObject_ITF
	QNetworkSession_PTR() *QNetworkSession
}

func (ptr *QNetworkSession) QNetworkSession_PTR() *QNetworkSession {
	return ptr
}

func (ptr *QNetworkSession) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QNetworkSession) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQNetworkSession(ptr QNetworkSession_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkSession_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkSession) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QNetworkSession) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQNetworkSessionFromPointer(ptr unsafe.Pointer) (n *QNetworkSession) {
	n = new(QNetworkSession)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkSession")
	return
}

//go:generate stringer -type=QNetworkSession__State
//QNetworkSession::State
type QNetworkSession__State int64

const (
	QNetworkSession__Invalid      QNetworkSession__State = QNetworkSession__State(0)
	QNetworkSession__NotAvailable QNetworkSession__State = QNetworkSession__State(1)
	QNetworkSession__Connecting   QNetworkSession__State = QNetworkSession__State(2)
	QNetworkSession__Connected    QNetworkSession__State = QNetworkSession__State(3)
	QNetworkSession__Closing      QNetworkSession__State = QNetworkSession__State(4)
	QNetworkSession__Disconnected QNetworkSession__State = QNetworkSession__State(5)
	QNetworkSession__Roaming      QNetworkSession__State = QNetworkSession__State(6)
)

//go:generate stringer -type=QNetworkSession__SessionError
//QNetworkSession::SessionError
type QNetworkSession__SessionError int64

const (
	QNetworkSession__UnknownSessionError        QNetworkSession__SessionError = QNetworkSession__SessionError(0)
	QNetworkSession__SessionAbortedError        QNetworkSession__SessionError = QNetworkSession__SessionError(1)
	QNetworkSession__RoamingError               QNetworkSession__SessionError = QNetworkSession__SessionError(2)
	QNetworkSession__OperationNotSupportedError QNetworkSession__SessionError = QNetworkSession__SessionError(3)
	QNetworkSession__InvalidConfigurationError  QNetworkSession__SessionError = QNetworkSession__SessionError(4)
)

//go:generate stringer -type=QNetworkSession__UsagePolicy
//QNetworkSession::UsagePolicy
type QNetworkSession__UsagePolicy int64

const (
	QNetworkSession__NoPolicy                  QNetworkSession__UsagePolicy = QNetworkSession__UsagePolicy(0)
	QNetworkSession__NoBackgroundTrafficPolicy QNetworkSession__UsagePolicy = QNetworkSession__UsagePolicy(1)
)

func NewQNetworkSession(connectionConfig QNetworkConfiguration_ITF, parent core.QObject_ITF) *QNetworkSession {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkSession", "", connectionConfig, parent}).(*QNetworkSession)
}

func (ptr *QNetworkSession) ConnectAccept(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAccept", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectAccept() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAccept"})
}

func (ptr *QNetworkSession) Accept() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Accept"})
}

func (ptr *QNetworkSession) AcceptDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcceptDefault"})
}

func (ptr *QNetworkSession) ActiveTime() uint64 {

	return uint64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ActiveTime"}).(float64))
}

func (ptr *QNetworkSession) BytesReceived() uint64 {

	return uint64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesReceived"}).(float64))
}

func (ptr *QNetworkSession) BytesWritten() uint64 {

	return uint64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesWritten"}).(float64))
}

func (ptr *QNetworkSession) ConnectClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QNetworkSession) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QNetworkSession) CloseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault"})
}

func (ptr *QNetworkSession) ConnectClosed(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClosed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectClosed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClosed"})
}

func (ptr *QNetworkSession) Closed() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Closed"})
}

func (ptr *QNetworkSession) Configuration() *QNetworkConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Configuration"}).(*QNetworkConfiguration)
}

func (ptr *QNetworkSession) Error() QNetworkSession__SessionError {

	return QNetworkSession__SessionError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QNetworkSession) ConnectError2(f func(error QNetworkSession__SessionError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QNetworkSession) Error2(error QNetworkSession__SessionError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", error})
}

func (ptr *QNetworkSession) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QNetworkSession) ConnectIgnore(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIgnore", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectIgnore() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIgnore"})
}

func (ptr *QNetworkSession) Ignore() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Ignore"})
}

func (ptr *QNetworkSession) IgnoreDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreDefault"})
}

func (ptr *QNetworkSession) Interface() *QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Interface"}).(*QNetworkInterface)
}

func (ptr *QNetworkSession) IsOpen() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsOpen"}).(bool)
}

func (ptr *QNetworkSession) ConnectMigrate(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectMigrate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectMigrate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectMigrate"})
}

func (ptr *QNetworkSession) Migrate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Migrate"})
}

func (ptr *QNetworkSession) MigrateDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MigrateDefault"})
}

func (ptr *QNetworkSession) ConnectNewConfigurationActivated(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewConfigurationActivated", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectNewConfigurationActivated() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewConfigurationActivated"})
}

func (ptr *QNetworkSession) NewConfigurationActivated() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewConfigurationActivated"})
}

func (ptr *QNetworkSession) ConnectOpen(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen"})
}

func (ptr *QNetworkSession) Open() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open"})
}

func (ptr *QNetworkSession) OpenDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault"})
}

func (ptr *QNetworkSession) ConnectOpened(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpened", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectOpened() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpened"})
}

func (ptr *QNetworkSession) Opened() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Opened"})
}

func (ptr *QNetworkSession) ConnectPreferredConfigurationChanged(f func(config *QNetworkConfiguration, isSeamless bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreferredConfigurationChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectPreferredConfigurationChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreferredConfigurationChanged"})
}

func (ptr *QNetworkSession) PreferredConfigurationChanged(config QNetworkConfiguration_ITF, isSeamless bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreferredConfigurationChanged", config, isSeamless})
}

func (ptr *QNetworkSession) ConnectReject(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReject", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectReject() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReject"})
}

func (ptr *QNetworkSession) Reject() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Reject"})
}

func (ptr *QNetworkSession) RejectDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RejectDefault"})
}

func (ptr *QNetworkSession) SessionProperty(key string) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SessionProperty", key}).(*core.QVariant)
}

func (ptr *QNetworkSession) SetSessionProperty(key string, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSessionProperty", key, value})
}

func (ptr *QNetworkSession) State() QNetworkSession__State {

	return QNetworkSession__State(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QNetworkSession) ConnectStateChanged(f func(state QNetworkSession__State)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QNetworkSession) StateChanged(state QNetworkSession__State) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", state})
}

func (ptr *QNetworkSession) ConnectStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QNetworkSession) Stop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"})
}

func (ptr *QNetworkSession) StopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopDefault"})
}

func (ptr *QNetworkSession) UsagePolicies() QNetworkSession__UsagePolicy {

	return QNetworkSession__UsagePolicy(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UsagePolicies"}).(float64))
}

func (ptr *QNetworkSession) ConnectUsagePoliciesChanged(f func(usagePolicies QNetworkSession__UsagePolicy)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectUsagePoliciesChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectUsagePoliciesChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectUsagePoliciesChanged"})
}

func (ptr *QNetworkSession) UsagePoliciesChanged(usagePolicies QNetworkSession__UsagePolicy) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UsagePoliciesChanged", usagePolicies})
}

func (ptr *QNetworkSession) WaitForOpened(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForOpened", msecs}).(bool)
}

func (ptr *QNetworkSession) ConnectDestroyQNetworkSession(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQNetworkSession", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QNetworkSession) DisconnectDestroyQNetworkSession() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQNetworkSession"})
}

func (ptr *QNetworkSession) DestroyQNetworkSession() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkSession"})
}

func (ptr *QNetworkSession) DestroyQNetworkSessionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQNetworkSessionDefault"})
}

func (ptr *QNetworkSession) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QNetworkSession) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QNetworkSession) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkSession) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QNetworkSession) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QNetworkSession) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkSession) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QNetworkSession) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QNetworkSession) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QNetworkSession) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QNetworkSession) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QNetworkSession) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QNetworkSession) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QNetworkSession) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QNetworkSession) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QNetworkSession) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QNetworkSession) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QNetworkSession) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QNetworkSession) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QNetworkSession) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QNetworkSession) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QOcspResponse struct {
	internal.Internal
}

type QOcspResponse_ITF interface {
	QOcspResponse_PTR() *QOcspResponse
}

func (ptr *QOcspResponse) QOcspResponse_PTR() *QOcspResponse {
	return ptr
}

func (ptr *QOcspResponse) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QOcspResponse) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQOcspResponse(ptr QOcspResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOcspResponse_PTR().Pointer()
	}
	return nil
}

func (n *QOcspResponse) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQOcspResponseFromPointer(ptr unsafe.Pointer) (n *QOcspResponse) {
	n = new(QOcspResponse)
	n.InitFromInternal(uintptr(ptr), "network.QOcspResponse")
	return
}

//go:generate stringer -type=QOcspResponse__QOcspCertificateStatus
//QOcspResponse::QOcspCertificateStatus
type QOcspResponse__QOcspCertificateStatus int64

const (
	QOcspResponse__Good    QOcspResponse__QOcspCertificateStatus = QOcspResponse__QOcspCertificateStatus(0)
	QOcspResponse__Revoked QOcspResponse__QOcspCertificateStatus = QOcspResponse__QOcspCertificateStatus(1)
	QOcspResponse__Unknown QOcspResponse__QOcspCertificateStatus = QOcspResponse__QOcspCertificateStatus(2)
)

//go:generate stringer -type=QOcspResponse__QOcspRevocationReason
//QOcspResponse::QOcspRevocationReason
type QOcspResponse__QOcspRevocationReason int64

const (
	QOcspResponse__None                 QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(-1)
	QOcspResponse__Unspecified          QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(0)
	QOcspResponse__KeyCompromise        QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(1)
	QOcspResponse__CACompromise         QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(2)
	QOcspResponse__AffiliationChanged   QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(3)
	QOcspResponse__Superseded           QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(4)
	QOcspResponse__CessationOfOperation QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(5)
	QOcspResponse__CertificateHold      QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(6)
	QOcspResponse__RemoveFromCRL        QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(7)
)

func NewQOcspResponse() *QOcspResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQOcspResponse", ""}).(*QOcspResponse)
}

func NewQOcspResponse2(other QOcspResponse_ITF) *QOcspResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQOcspResponse2", "", other}).(*QOcspResponse)
}

func NewQOcspResponse3(other QOcspResponse_ITF) *QOcspResponse {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQOcspResponse3", "", other}).(*QOcspResponse)
}

func (ptr *QOcspResponse) Subject() *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Subject"}).(*QSslCertificate)
}

func (ptr *QOcspResponse) Swap(other QOcspResponse_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QOcspResponse) DestroyQOcspResponse() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQOcspResponse"})
}

type QPasswordDigestor struct {
	internal.Internal
}

type QPasswordDigestor_ITF interface {
	QPasswordDigestor_PTR() *QPasswordDigestor
}

func (ptr *QPasswordDigestor) QPasswordDigestor_PTR() *QPasswordDigestor {
	return ptr
}

func (ptr *QPasswordDigestor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QPasswordDigestor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQPasswordDigestor(ptr QPasswordDigestor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPasswordDigestor_PTR().Pointer()
	}
	return nil
}

func (n *QPasswordDigestor) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQPasswordDigestorFromPointer(ptr unsafe.Pointer) (n *QPasswordDigestor) {
	n = new(QPasswordDigestor)
	n.InitFromInternal(uintptr(ptr), "network.QPasswordDigestor")
	return
}

func (ptr *QPasswordDigestor) DestroyQPasswordDigestor() {
}

type QSctpServer struct {
	QTcpServer
}

type QSctpServer_ITF interface {
	QTcpServer_ITF
	QSctpServer_PTR() *QSctpServer
}

func (ptr *QSctpServer) QSctpServer_PTR() *QSctpServer {
	return ptr
}

func (ptr *QSctpServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpServer_PTR().Pointer()
	}
	return nil
}

func (ptr *QSctpServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTcpServer_PTR().SetPointer(p)
	}
}

func PointerFromQSctpServer(ptr QSctpServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSctpServer_PTR().Pointer()
	}
	return nil
}

func (n *QSctpServer) InitFromInternal(ptr uintptr, name string) {
	n.QTcpServer_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSctpServer) ClassNameInternalF() string {
	return n.QTcpServer_PTR().ClassNameInternalF()
}

func NewQSctpServerFromPointer(ptr unsafe.Pointer) (n *QSctpServer) {
	n = new(QSctpServer)
	n.InitFromInternal(uintptr(ptr), "network.QSctpServer")
	return
}

type QSctpSocket struct {
	QTcpSocket
}

type QSctpSocket_ITF interface {
	QTcpSocket_ITF
	QSctpSocket_PTR() *QSctpSocket
}

func (ptr *QSctpSocket) QSctpSocket_PTR() *QSctpSocket {
	return ptr
}

func (ptr *QSctpSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func (ptr *QSctpSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTcpSocket_PTR().SetPointer(p)
	}
}

func PointerFromQSctpSocket(ptr QSctpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSctpSocket_PTR().Pointer()
	}
	return nil
}

func (n *QSctpSocket) InitFromInternal(ptr uintptr, name string) {
	n.QTcpSocket_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSctpSocket) ClassNameInternalF() string {
	return n.QTcpSocket_PTR().ClassNameInternalF()
}

func NewQSctpSocketFromPointer(ptr unsafe.Pointer) (n *QSctpSocket) {
	n = new(QSctpSocket)
	n.InitFromInternal(uintptr(ptr), "network.QSctpSocket")
	return
}

type QSsl struct {
	internal.Internal
}

type QSsl_ITF interface {
	QSsl_PTR() *QSsl
}

func (ptr *QSsl) QSsl_PTR() *QSsl {
	return ptr
}

func (ptr *QSsl) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSsl) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSsl(ptr QSsl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSsl_PTR().Pointer()
	}
	return nil
}

func (n *QSsl) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslFromPointer(ptr unsafe.Pointer) (n *QSsl) {
	n = new(QSsl)
	n.InitFromInternal(uintptr(ptr), "network.QSsl")
	return
}

func (ptr *QSsl) DestroyQSsl() {
}

//go:generate stringer -type=QSsl__KeyType
//QSsl::KeyType
type QSsl__KeyType int64

const (
	QSsl__PrivateKey QSsl__KeyType = QSsl__KeyType(0)
	QSsl__PublicKey  QSsl__KeyType = QSsl__KeyType(1)
)

//go:generate stringer -type=QSsl__EncodingFormat
//QSsl::EncodingFormat
type QSsl__EncodingFormat int64

const (
	QSsl__Pem QSsl__EncodingFormat = QSsl__EncodingFormat(0)
	QSsl__Der QSsl__EncodingFormat = QSsl__EncodingFormat(1)
)

//go:generate stringer -type=QSsl__KeyAlgorithm
//QSsl::KeyAlgorithm
type QSsl__KeyAlgorithm int64

const (
	QSsl__Opaque QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(0)
	QSsl__Rsa    QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(1)
	QSsl__Dsa    QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(2)
	QSsl__Ec     QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(3)
	QSsl__Dh     QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(4)
)

//go:generate stringer -type=QSsl__AlternativeNameEntryType
//QSsl::AlternativeNameEntryType
type QSsl__AlternativeNameEntryType int64

const (
	QSsl__EmailEntry     QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(0)
	QSsl__DnsEntry       QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(1)
	QSsl__IpAddressEntry QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(2)
)

//go:generate stringer -type=QSsl__SslProtocol
//QSsl::SslProtocol
type QSsl__SslProtocol int64

const (
	QSsl__SslV3           QSsl__SslProtocol = QSsl__SslProtocol(0)
	QSsl__SslV2           QSsl__SslProtocol = QSsl__SslProtocol(1)
	QSsl__TlsV1_0         QSsl__SslProtocol = QSsl__SslProtocol(2)
	QSsl__TlsV1           QSsl__SslProtocol = QSsl__SslProtocol(QSsl__TlsV1_0)
	QSsl__TlsV1_1         QSsl__SslProtocol = QSsl__SslProtocol(3)
	QSsl__TlsV1_2         QSsl__SslProtocol = QSsl__SslProtocol(4)
	QSsl__AnyProtocol     QSsl__SslProtocol = QSsl__SslProtocol(5)
	QSsl__TlsV1SslV3      QSsl__SslProtocol = QSsl__SslProtocol(6)
	QSsl__SecureProtocols QSsl__SslProtocol = QSsl__SslProtocol(7)
	QSsl__TlsV1_0OrLater  QSsl__SslProtocol = QSsl__SslProtocol(8)
	QSsl__TlsV1_1OrLater  QSsl__SslProtocol = QSsl__SslProtocol(9)
	QSsl__TlsV1_2OrLater  QSsl__SslProtocol = QSsl__SslProtocol(10)
	QSsl__DtlsV1_0        QSsl__SslProtocol = QSsl__SslProtocol(11)
	QSsl__DtlsV1_0OrLater QSsl__SslProtocol = QSsl__SslProtocol(12)
	QSsl__DtlsV1_2        QSsl__SslProtocol = QSsl__SslProtocol(13)
	QSsl__DtlsV1_2OrLater QSsl__SslProtocol = QSsl__SslProtocol(14)
	QSsl__TlsV1_3         QSsl__SslProtocol = QSsl__SslProtocol(15)
	QSsl__TlsV1_3OrLater  QSsl__SslProtocol = QSsl__SslProtocol(16)
	QSsl__UnknownProtocol QSsl__SslProtocol = QSsl__SslProtocol(-1)
)

//go:generate stringer -type=QSsl__SslOption
//QSsl::SslOption
type QSsl__SslOption int64

const (
	QSsl__SslOptionDisableEmptyFragments         QSsl__SslOption = QSsl__SslOption(0x01)
	QSsl__SslOptionDisableSessionTickets         QSsl__SslOption = QSsl__SslOption(0x02)
	QSsl__SslOptionDisableCompression            QSsl__SslOption = QSsl__SslOption(0x04)
	QSsl__SslOptionDisableServerNameIndication   QSsl__SslOption = QSsl__SslOption(0x08)
	QSsl__SslOptionDisableLegacyRenegotiation    QSsl__SslOption = QSsl__SslOption(0x10)
	QSsl__SslOptionDisableSessionSharing         QSsl__SslOption = QSsl__SslOption(0x20)
	QSsl__SslOptionDisableSessionPersistence     QSsl__SslOption = QSsl__SslOption(0x40)
	QSsl__SslOptionDisableServerCipherPreference QSsl__SslOption = QSsl__SslOption(0x80)
)

type QSslCertificate struct {
	internal.Internal
}

type QSslCertificate_ITF interface {
	QSslCertificate_PTR() *QSslCertificate
}

func (ptr *QSslCertificate) QSslCertificate_PTR() *QSslCertificate {
	return ptr
}

func (ptr *QSslCertificate) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSslCertificate) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSslCertificate(ptr QSslCertificate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificate_PTR().Pointer()
	}
	return nil
}

func (n *QSslCertificate) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslCertificateFromPointer(ptr unsafe.Pointer) (n *QSslCertificate) {
	n = new(QSslCertificate)
	n.InitFromInternal(uintptr(ptr), "network.QSslCertificate")
	return
}

//go:generate stringer -type=QSslCertificate__SubjectInfo
//QSslCertificate::SubjectInfo
type QSslCertificate__SubjectInfo int64

const (
	QSslCertificate__Organization               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(0)
	QSslCertificate__CommonName                 QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(1)
	QSslCertificate__LocalityName               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(2)
	QSslCertificate__OrganizationalUnitName     QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(3)
	QSslCertificate__CountryName                QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(4)
	QSslCertificate__StateOrProvinceName        QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(5)
	QSslCertificate__DistinguishedNameQualifier QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(6)
	QSslCertificate__SerialNumber               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(7)
	QSslCertificate__EmailAddress               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(8)
)

func NewQSslCertificate(device core.QIODevice_ITF, format QSsl__EncodingFormat) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslCertificate", "", device, format}).(*QSslCertificate)
}

func NewQSslCertificate2(data core.QByteArray_ITF, format QSsl__EncodingFormat) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslCertificate2", "", data, format}).(*QSslCertificate)
}

func NewQSslCertificate3(other QSslCertificate_ITF) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslCertificate3", "", other}).(*QSslCertificate)
}

func (ptr *QSslCertificate) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QSslCertificate) Digest(algorithm core.QCryptographicHash__Algorithm) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Digest", algorithm}).(*core.QByteArray)
}

func (ptr *QSslCertificate) EffectiveDate() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EffectiveDate"}).(*core.QDateTime)
}

func (ptr *QSslCertificate) ExpiryDate() *core.QDateTime {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ExpiryDate"}).(*core.QDateTime)
}

func (ptr *QSslCertificate) Extensions() []*QSslCertificateExtension {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Extensions"}).([]*QSslCertificateExtension)
}

func QSslCertificate_FromData(data core.QByteArray_ITF, format QSsl__EncodingFormat) []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_FromData", "", data, format}).([]*QSslCertificate)
}

func (ptr *QSslCertificate) FromData(data core.QByteArray_ITF, format QSsl__EncodingFormat) []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_FromData", "", data, format}).([]*QSslCertificate)
}

func QSslCertificate_FromDevice(device core.QIODevice_ITF, format QSsl__EncodingFormat) []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_FromDevice", "", device, format}).([]*QSslCertificate)
}

func (ptr *QSslCertificate) FromDevice(device core.QIODevice_ITF, format QSsl__EncodingFormat) []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_FromDevice", "", device, format}).([]*QSslCertificate)
}

func QSslCertificate_FromPath(path string, format QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_FromPath", "", path, format, syntax}).([]*QSslCertificate)
}

func (ptr *QSslCertificate) FromPath(path string, format QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_FromPath", "", path, format, syntax}).([]*QSslCertificate)
}

func QSslCertificate_ImportPkcs12(device core.QIODevice_ITF, key QSslKey_ITF, certificate QSslCertificate_ITF, caCertificates []*QSslCertificate, passPhrase core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_ImportPkcs12", "", device, key, certificate, caCertificates, passPhrase}).(bool)
}

func (ptr *QSslCertificate) ImportPkcs12(device core.QIODevice_ITF, key QSslKey_ITF, certificate QSslCertificate_ITF, caCertificates []*QSslCertificate, passPhrase core.QByteArray_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_ImportPkcs12", "", device, key, certificate, caCertificates, passPhrase}).(bool)
}

func (ptr *QSslCertificate) IsBlacklisted() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsBlacklisted"}).(bool)
}

func (ptr *QSslCertificate) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QSslCertificate) IsSelfSigned() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSelfSigned"}).(bool)
}

func (ptr *QSslCertificate) IssuerDisplayName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IssuerDisplayName"}).(string)
}

func (ptr *QSslCertificate) IssuerInfo(subject QSslCertificate__SubjectInfo) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IssuerInfo", subject}).([]string)
}

func (ptr *QSslCertificate) IssuerInfo2(attribute core.QByteArray_ITF) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IssuerInfo2", attribute}).([]string)
}

func (ptr *QSslCertificate) IssuerInfoAttributes() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IssuerInfoAttributes"}).([]*core.QByteArray)
}

func (ptr *QSslCertificate) PublicKey() *QSslKey {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PublicKey"}).(*QSslKey)
}

func (ptr *QSslCertificate) SerialNumber() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SerialNumber"}).(*core.QByteArray)
}

func (ptr *QSslCertificate) SubjectAlternativeNames() map[QSsl__AlternativeNameEntryType]string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubjectAlternativeNames"}).(map[QSsl__AlternativeNameEntryType]string)
}

func (ptr *QSslCertificate) SubjectDisplayName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubjectDisplayName"}).(string)
}

func (ptr *QSslCertificate) SubjectInfo(subject QSslCertificate__SubjectInfo) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubjectInfo", subject}).([]string)
}

func (ptr *QSslCertificate) SubjectInfo2(attribute core.QByteArray_ITF) []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubjectInfo2", attribute}).([]string)
}

func (ptr *QSslCertificate) SubjectInfoAttributes() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SubjectInfoAttributes"}).([]*core.QByteArray)
}

func (ptr *QSslCertificate) Swap(other QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QSslCertificate) ToDer() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToDer"}).(*core.QByteArray)
}

func (ptr *QSslCertificate) ToPem() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToPem"}).(*core.QByteArray)
}

func (ptr *QSslCertificate) ToText() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToText"}).(string)
}

func QSslCertificate_Verify(certificateChain []*QSslCertificate, hostName string) []*QSslError {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_Verify", "", certificateChain, hostName}).([]*QSslError)
}

func (ptr *QSslCertificate) Verify(certificateChain []*QSslCertificate, hostName string) []*QSslError {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslCertificate_Verify", "", certificateChain, hostName}).([]*QSslError)
}

func (ptr *QSslCertificate) Version() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Version"}).(*core.QByteArray)
}

func (ptr *QSslCertificate) DestroyQSslCertificate() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslCertificate"})
}

func (ptr *QSslCertificate) __extensions_atList(i int) *QSslCertificateExtension {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extensions_atList", i}).(*QSslCertificateExtension)
}

func (ptr *QSslCertificate) __extensions_setList(i QSslCertificateExtension_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extensions_setList", i})
}

func (ptr *QSslCertificate) __extensions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__extensions_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) __fromData_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__fromData_atList", i}).(*QSslCertificate)
}

func (ptr *QSslCertificate) __fromData_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__fromData_setList", i})
}

func (ptr *QSslCertificate) __fromData_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__fromData_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) __fromDevice_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__fromDevice_atList", i}).(*QSslCertificate)
}

func (ptr *QSslCertificate) __fromDevice_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__fromDevice_setList", i})
}

func (ptr *QSslCertificate) __fromDevice_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__fromDevice_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) __fromPath_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__fromPath_atList", i}).(*QSslCertificate)
}

func (ptr *QSslCertificate) __fromPath_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__fromPath_setList", i})
}

func (ptr *QSslCertificate) __fromPath_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__fromPath_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) __importPkcs12_caCertificates_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__importPkcs12_caCertificates_atList", i}).(*QSslCertificate)
}

func (ptr *QSslCertificate) __importPkcs12_caCertificates_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__importPkcs12_caCertificates_setList", i})
}

func (ptr *QSslCertificate) __importPkcs12_caCertificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__importPkcs12_caCertificates_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) __issuerInfoAttributes_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__issuerInfoAttributes_atList", i}).(*core.QByteArray)
}

func (ptr *QSslCertificate) __issuerInfoAttributes_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__issuerInfoAttributes_setList", i})
}

func (ptr *QSslCertificate) __issuerInfoAttributes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__issuerInfoAttributes_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) __subjectAlternativeNames_atList(v QSsl__AlternativeNameEntryType, i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subjectAlternativeNames_atList", v, i}).(string)
}

func (ptr *QSslCertificate) __subjectAlternativeNames_setList(key QSsl__AlternativeNameEntryType, i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subjectAlternativeNames_setList", key, i})
}

func (ptr *QSslCertificate) __subjectAlternativeNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subjectAlternativeNames_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) __subjectAlternativeNames_keyList() []QSsl__AlternativeNameEntryType {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subjectAlternativeNames_keyList"}).([]QSsl__AlternativeNameEntryType)
}

func (ptr *QSslCertificate) __subjectInfoAttributes_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subjectInfoAttributes_atList", i}).(*core.QByteArray)
}

func (ptr *QSslCertificate) __subjectInfoAttributes_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subjectInfoAttributes_setList", i})
}

func (ptr *QSslCertificate) __subjectInfoAttributes_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__subjectInfoAttributes_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) __verify_atList(i int) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__verify_atList", i}).(*QSslError)
}

func (ptr *QSslCertificate) __verify_setList(i QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__verify_setList", i})
}

func (ptr *QSslCertificate) __verify_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__verify_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) __verify_certificateChain_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__verify_certificateChain_atList", i}).(*QSslCertificate)
}

func (ptr *QSslCertificate) __verify_certificateChain_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__verify_certificateChain_setList", i})
}

func (ptr *QSslCertificate) __verify_certificateChain_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__verify_certificateChain_newList"}).(unsafe.Pointer)
}

func (ptr *QSslCertificate) ____subjectAlternativeNames_keyList_atList(i int) QSsl__AlternativeNameEntryType {

	return QSsl__AlternativeNameEntryType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____subjectAlternativeNames_keyList_atList", i}).(float64))
}

func (ptr *QSslCertificate) ____subjectAlternativeNames_keyList_setList(i QSsl__AlternativeNameEntryType) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____subjectAlternativeNames_keyList_setList", i})
}

func (ptr *QSslCertificate) ____subjectAlternativeNames_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____subjectAlternativeNames_keyList_newList"}).(unsafe.Pointer)
}

type QSslCertificateExtension struct {
	internal.Internal
}

type QSslCertificateExtension_ITF interface {
	QSslCertificateExtension_PTR() *QSslCertificateExtension
}

func (ptr *QSslCertificateExtension) QSslCertificateExtension_PTR() *QSslCertificateExtension {
	return ptr
}

func (ptr *QSslCertificateExtension) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSslCertificateExtension) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSslCertificateExtension(ptr QSslCertificateExtension_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCertificateExtension_PTR().Pointer()
	}
	return nil
}

func (n *QSslCertificateExtension) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslCertificateExtensionFromPointer(ptr unsafe.Pointer) (n *QSslCertificateExtension) {
	n = new(QSslCertificateExtension)
	n.InitFromInternal(uintptr(ptr), "network.QSslCertificateExtension")
	return
}
func NewQSslCertificateExtension() *QSslCertificateExtension {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslCertificateExtension", ""}).(*QSslCertificateExtension)
}

func NewQSslCertificateExtension2(other QSslCertificateExtension_ITF) *QSslCertificateExtension {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslCertificateExtension2", "", other}).(*QSslCertificateExtension)
}

func (ptr *QSslCertificateExtension) IsCritical() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsCritical"}).(bool)
}

func (ptr *QSslCertificateExtension) IsSupported() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsSupported"}).(bool)
}

func (ptr *QSslCertificateExtension) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QSslCertificateExtension) Oid() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Oid"}).(string)
}

func (ptr *QSslCertificateExtension) Swap(other QSslCertificateExtension_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QSslCertificateExtension) Value() *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Value"}).(*core.QVariant)
}

func (ptr *QSslCertificateExtension) DestroyQSslCertificateExtension() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslCertificateExtension"})
}

type QSslCipher struct {
	internal.Internal
}

type QSslCipher_ITF interface {
	QSslCipher_PTR() *QSslCipher
}

func (ptr *QSslCipher) QSslCipher_PTR() *QSslCipher {
	return ptr
}

func (ptr *QSslCipher) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSslCipher) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSslCipher(ptr QSslCipher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCipher_PTR().Pointer()
	}
	return nil
}

func (n *QSslCipher) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslCipherFromPointer(ptr unsafe.Pointer) (n *QSslCipher) {
	n = new(QSslCipher)
	n.InitFromInternal(uintptr(ptr), "network.QSslCipher")
	return
}
func NewQSslCipher() *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslCipher", ""}).(*QSslCipher)
}

func NewQSslCipher2(name string) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslCipher2", "", name}).(*QSslCipher)
}

func NewQSslCipher3(name string, protoc QSsl__SslProtocol) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslCipher3", "", name, protoc}).(*QSslCipher)
}

func NewQSslCipher4(other QSslCipher_ITF) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslCipher4", "", other}).(*QSslCipher)
}

func (ptr *QSslCipher) AuthenticationMethod() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AuthenticationMethod"}).(string)
}

func (ptr *QSslCipher) EncryptionMethod() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EncryptionMethod"}).(string)
}

func (ptr *QSslCipher) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QSslCipher) KeyExchangeMethod() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "KeyExchangeMethod"}).(string)
}

func (ptr *QSslCipher) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func (ptr *QSslCipher) Protocol() QSsl__SslProtocol {

	return QSsl__SslProtocol(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Protocol"}).(float64))
}

func (ptr *QSslCipher) ProtocolString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProtocolString"}).(string)
}

func (ptr *QSslCipher) SupportedBits() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedBits"}).(float64))
}

func (ptr *QSslCipher) Swap(other QSslCipher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QSslCipher) UsedBits() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "UsedBits"}).(float64))
}

func (ptr *QSslCipher) DestroyQSslCipher() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslCipher"})
}

type QSslConfiguration struct {
	internal.Internal
}

type QSslConfiguration_ITF interface {
	QSslConfiguration_PTR() *QSslConfiguration
}

func (ptr *QSslConfiguration) QSslConfiguration_PTR() *QSslConfiguration {
	return ptr
}

func (ptr *QSslConfiguration) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSslConfiguration) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSslConfiguration(ptr QSslConfiguration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslConfiguration_PTR().Pointer()
	}
	return nil
}

func (n *QSslConfiguration) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslConfigurationFromPointer(ptr unsafe.Pointer) (n *QSslConfiguration) {
	n = new(QSslConfiguration)
	n.InitFromInternal(uintptr(ptr), "network.QSslConfiguration")
	return
}

//go:generate stringer -type=QSslConfiguration__NextProtocolNegotiationStatus
//QSslConfiguration::NextProtocolNegotiationStatus
type QSslConfiguration__NextProtocolNegotiationStatus int64

const (
	QSslConfiguration__NextProtocolNegotiationNone        QSslConfiguration__NextProtocolNegotiationStatus = QSslConfiguration__NextProtocolNegotiationStatus(0)
	QSslConfiguration__NextProtocolNegotiationNegotiated  QSslConfiguration__NextProtocolNegotiationStatus = QSslConfiguration__NextProtocolNegotiationStatus(1)
	QSslConfiguration__NextProtocolNegotiationUnsupported QSslConfiguration__NextProtocolNegotiationStatus = QSslConfiguration__NextProtocolNegotiationStatus(2)
)

func NewQSslConfiguration() *QSslConfiguration {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslConfiguration", ""}).(*QSslConfiguration)
}

func NewQSslConfiguration2(other QSslConfiguration_ITF) *QSslConfiguration {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslConfiguration2", "", other}).(*QSslConfiguration)
}

func (ptr *QSslConfiguration) AllowedNextProtocols() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AllowedNextProtocols"}).([]*core.QByteArray)
}

func (ptr *QSslConfiguration) BackendConfiguration() map[*core.QByteArray]*core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BackendConfiguration"}).(map[*core.QByteArray]*core.QVariant)
}

func (ptr *QSslConfiguration) CaCertificates() []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CaCertificates"}).([]*QSslCertificate)
}

func (ptr *QSslConfiguration) Ciphers() []*QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Ciphers"}).([]*QSslCipher)
}

func QSslConfiguration_DefaultConfiguration() *QSslConfiguration {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslConfiguration_DefaultConfiguration", ""}).(*QSslConfiguration)
}

func (ptr *QSslConfiguration) DefaultConfiguration() *QSslConfiguration {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslConfiguration_DefaultConfiguration", ""}).(*QSslConfiguration)
}

func (ptr *QSslConfiguration) DiffieHellmanParameters() *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DiffieHellmanParameters"}).(*QSslDiffieHellmanParameters)
}

func (ptr *QSslConfiguration) EphemeralServerKey() *QSslKey {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EphemeralServerKey"}).(*QSslKey)
}

func (ptr *QSslConfiguration) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QSslConfiguration) LocalCertificate() *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalCertificate"}).(*QSslCertificate)
}

func (ptr *QSslConfiguration) LocalCertificateChain() []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalCertificateChain"}).([]*QSslCertificate)
}

func (ptr *QSslConfiguration) NextNegotiatedProtocol() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextNegotiatedProtocol"}).(*core.QByteArray)
}

func (ptr *QSslConfiguration) NextProtocolNegotiationStatus() QSslConfiguration__NextProtocolNegotiationStatus {

	return QSslConfiguration__NextProtocolNegotiationStatus(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextProtocolNegotiationStatus"}).(float64))
}

func (ptr *QSslConfiguration) OcspStaplingEnabled() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OcspStaplingEnabled"}).(bool)
}

func (ptr *QSslConfiguration) PeerCertificate() *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerCertificate"}).(*QSslCertificate)
}

func (ptr *QSslConfiguration) PeerCertificateChain() []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerCertificateChain"}).([]*QSslCertificate)
}

func (ptr *QSslConfiguration) PeerVerifyDepth() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerVerifyDepth"}).(float64))
}

func (ptr *QSslConfiguration) PeerVerifyMode() QSslSocket__PeerVerifyMode {

	return QSslSocket__PeerVerifyMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerVerifyMode"}).(float64))
}

func (ptr *QSslConfiguration) PreSharedKeyIdentityHint() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreSharedKeyIdentityHint"}).(*core.QByteArray)
}

func (ptr *QSslConfiguration) PrivateKey() *QSslKey {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrivateKey"}).(*QSslKey)
}

func (ptr *QSslConfiguration) Protocol() QSsl__SslProtocol {

	return QSsl__SslProtocol(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Protocol"}).(float64))
}

func (ptr *QSslConfiguration) SessionCipher() *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SessionCipher"}).(*QSslCipher)
}

func (ptr *QSslConfiguration) SessionProtocol() QSsl__SslProtocol {

	return QSsl__SslProtocol(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SessionProtocol"}).(float64))
}

func (ptr *QSslConfiguration) SessionTicket() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SessionTicket"}).(*core.QByteArray)
}

func (ptr *QSslConfiguration) SessionTicketLifeTimeHint() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SessionTicketLifeTimeHint"}).(float64))
}

func (ptr *QSslConfiguration) SetAllowedNextProtocols(protocols []*core.QByteArray) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAllowedNextProtocols", protocols})
}

func (ptr *QSslConfiguration) SetBackendConfiguration(backendConfiguration map[*core.QByteArray]*core.QVariant) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackendConfiguration", backendConfiguration})
}

func (ptr *QSslConfiguration) SetBackendConfigurationOption(name core.QByteArray_ITF, value core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetBackendConfigurationOption", name, value})
}

func (ptr *QSslConfiguration) SetCaCertificates(certificates []*QSslCertificate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCaCertificates", certificates})
}

func (ptr *QSslConfiguration) SetCiphers(ciphers []*QSslCipher) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetCiphers", ciphers})
}

func QSslConfiguration_SetDefaultConfiguration(configuration QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QSslConfiguration_SetDefaultConfiguration", "", configuration})
}

func (ptr *QSslConfiguration) SetDefaultConfiguration(configuration QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QSslConfiguration_SetDefaultConfiguration", "", configuration})
}

func (ptr *QSslConfiguration) SetDiffieHellmanParameters(dhparams QSslDiffieHellmanParameters_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetDiffieHellmanParameters", dhparams})
}

func (ptr *QSslConfiguration) SetEllipticCurves(curves []*QSslEllipticCurve) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetEllipticCurves", curves})
}

func (ptr *QSslConfiguration) SetLocalCertificate(certificate QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalCertificate", certificate})
}

func (ptr *QSslConfiguration) SetLocalCertificateChain(localChain []*QSslCertificate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalCertificateChain", localChain})
}

func (ptr *QSslConfiguration) SetOcspStaplingEnabled(enabled bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetOcspStaplingEnabled", enabled})
}

func (ptr *QSslConfiguration) SetPeerVerifyDepth(depth int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerVerifyDepth", depth})
}

func (ptr *QSslConfiguration) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerVerifyMode", mode})
}

func (ptr *QSslConfiguration) SetPreSharedKeyIdentityHint(hint core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPreSharedKeyIdentityHint", hint})
}

func (ptr *QSslConfiguration) SetPrivateKey(key QSslKey_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrivateKey", key})
}

func (ptr *QSslConfiguration) SetProtocol(protoc QSsl__SslProtocol) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProtocol", protoc})
}

func (ptr *QSslConfiguration) SetSessionTicket(sessionTicket core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSessionTicket", sessionTicket})
}

func (ptr *QSslConfiguration) SetSslOption(option QSsl__SslOption, on bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSslOption", option, on})
}

func QSslConfiguration_SupportedCiphers() []*QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslConfiguration_SupportedCiphers", ""}).([]*QSslCipher)
}

func (ptr *QSslConfiguration) SupportedCiphers() []*QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslConfiguration_SupportedCiphers", ""}).([]*QSslCipher)
}

func (ptr *QSslConfiguration) Swap(other QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func QSslConfiguration_SystemCaCertificates() []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslConfiguration_SystemCaCertificates", ""}).([]*QSslCertificate)
}

func (ptr *QSslConfiguration) SystemCaCertificates() []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslConfiguration_SystemCaCertificates", ""}).([]*QSslCertificate)
}

func (ptr *QSslConfiguration) TestSslOption(option QSsl__SslOption) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TestSslOption", option}).(bool)
}

func (ptr *QSslConfiguration) DestroyQSslConfiguration() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslConfiguration"})
}

func (ptr *QSslConfiguration) __allowedNextProtocols_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allowedNextProtocols_atList", i}).(*core.QByteArray)
}

func (ptr *QSslConfiguration) __allowedNextProtocols_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allowedNextProtocols_setList", i})
}

func (ptr *QSslConfiguration) __allowedNextProtocols_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__allowedNextProtocols_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __backendConfiguration_atList(v core.QByteArray_ITF, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__backendConfiguration_atList", v, i}).(*core.QVariant)
}

func (ptr *QSslConfiguration) __backendConfiguration_setList(key core.QByteArray_ITF, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__backendConfiguration_setList", key, i})
}

func (ptr *QSslConfiguration) __backendConfiguration_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__backendConfiguration_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __backendConfiguration_keyList() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__backendConfiguration_keyList"}).([]*core.QByteArray)
}

func (ptr *QSslConfiguration) __caCertificates_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__caCertificates_atList", i}).(*QSslCertificate)
}

func (ptr *QSslConfiguration) __caCertificates_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__caCertificates_setList", i})
}

func (ptr *QSslConfiguration) __caCertificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__caCertificates_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __ciphers_atList(i int) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ciphers_atList", i}).(*QSslCipher)
}

func (ptr *QSslConfiguration) __ciphers_setList(i QSslCipher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ciphers_setList", i})
}

func (ptr *QSslConfiguration) __ciphers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ciphers_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __ellipticCurves_atList(i int) *QSslEllipticCurve {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ellipticCurves_atList", i}).(*QSslEllipticCurve)
}

func (ptr *QSslConfiguration) __ellipticCurves_setList(i QSslEllipticCurve_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ellipticCurves_setList", i})
}

func (ptr *QSslConfiguration) __ellipticCurves_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ellipticCurves_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __localCertificateChain_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__localCertificateChain_atList", i}).(*QSslCertificate)
}

func (ptr *QSslConfiguration) __localCertificateChain_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__localCertificateChain_setList", i})
}

func (ptr *QSslConfiguration) __localCertificateChain_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__localCertificateChain_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __peerCertificateChain_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__peerCertificateChain_atList", i}).(*QSslCertificate)
}

func (ptr *QSslConfiguration) __peerCertificateChain_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__peerCertificateChain_setList", i})
}

func (ptr *QSslConfiguration) __peerCertificateChain_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__peerCertificateChain_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __setAllowedNextProtocols_protocols_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAllowedNextProtocols_protocols_atList", i}).(*core.QByteArray)
}

func (ptr *QSslConfiguration) __setAllowedNextProtocols_protocols_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAllowedNextProtocols_protocols_setList", i})
}

func (ptr *QSslConfiguration) __setAllowedNextProtocols_protocols_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setAllowedNextProtocols_protocols_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __setBackendConfiguration_backendConfiguration_atList(v core.QByteArray_ITF, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBackendConfiguration_backendConfiguration_atList", v, i}).(*core.QVariant)
}

func (ptr *QSslConfiguration) __setBackendConfiguration_backendConfiguration_setList(key core.QByteArray_ITF, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBackendConfiguration_backendConfiguration_setList", key, i})
}

func (ptr *QSslConfiguration) __setBackendConfiguration_backendConfiguration_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBackendConfiguration_backendConfiguration_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __setBackendConfiguration_backendConfiguration_keyList() []*core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setBackendConfiguration_backendConfiguration_keyList"}).([]*core.QByteArray)
}

func (ptr *QSslConfiguration) __setCaCertificates_certificates_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCaCertificates_certificates_atList", i}).(*QSslCertificate)
}

func (ptr *QSslConfiguration) __setCaCertificates_certificates_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCaCertificates_certificates_setList", i})
}

func (ptr *QSslConfiguration) __setCaCertificates_certificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCaCertificates_certificates_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __setCiphers_ciphers_atList(i int) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCiphers_ciphers_atList", i}).(*QSslCipher)
}

func (ptr *QSslConfiguration) __setCiphers_ciphers_setList(i QSslCipher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCiphers_ciphers_setList", i})
}

func (ptr *QSslConfiguration) __setCiphers_ciphers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCiphers_ciphers_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __setEllipticCurves_curves_atList(i int) *QSslEllipticCurve {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setEllipticCurves_curves_atList", i}).(*QSslEllipticCurve)
}

func (ptr *QSslConfiguration) __setEllipticCurves_curves_setList(i QSslEllipticCurve_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setEllipticCurves_curves_setList", i})
}

func (ptr *QSslConfiguration) __setEllipticCurves_curves_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setEllipticCurves_curves_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __setLocalCertificateChain_localChain_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setLocalCertificateChain_localChain_atList", i}).(*QSslCertificate)
}

func (ptr *QSslConfiguration) __setLocalCertificateChain_localChain_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setLocalCertificateChain_localChain_setList", i})
}

func (ptr *QSslConfiguration) __setLocalCertificateChain_localChain_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setLocalCertificateChain_localChain_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __supportedCiphers_atList(i int) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedCiphers_atList", i}).(*QSslCipher)
}

func (ptr *QSslConfiguration) __supportedCiphers_setList(i QSslCipher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedCiphers_setList", i})
}

func (ptr *QSslConfiguration) __supportedCiphers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedCiphers_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __supportedEllipticCurves_atList(i int) *QSslEllipticCurve {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedEllipticCurves_atList", i}).(*QSslEllipticCurve)
}

func (ptr *QSslConfiguration) __supportedEllipticCurves_setList(i QSslEllipticCurve_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedEllipticCurves_setList", i})
}

func (ptr *QSslConfiguration) __supportedEllipticCurves_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedEllipticCurves_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) __systemCaCertificates_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__systemCaCertificates_atList", i}).(*QSslCertificate)
}

func (ptr *QSslConfiguration) __systemCaCertificates_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__systemCaCertificates_setList", i})
}

func (ptr *QSslConfiguration) __systemCaCertificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__systemCaCertificates_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) ____backendConfiguration_keyList_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____backendConfiguration_keyList_atList", i}).(*core.QByteArray)
}

func (ptr *QSslConfiguration) ____backendConfiguration_keyList_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____backendConfiguration_keyList_setList", i})
}

func (ptr *QSslConfiguration) ____backendConfiguration_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____backendConfiguration_keyList_newList"}).(unsafe.Pointer)
}

func (ptr *QSslConfiguration) ____setBackendConfiguration_backendConfiguration_keyList_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setBackendConfiguration_backendConfiguration_keyList_atList", i}).(*core.QByteArray)
}

func (ptr *QSslConfiguration) ____setBackendConfiguration_backendConfiguration_keyList_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setBackendConfiguration_backendConfiguration_keyList_setList", i})
}

func (ptr *QSslConfiguration) ____setBackendConfiguration_backendConfiguration_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____setBackendConfiguration_backendConfiguration_keyList_newList"}).(unsafe.Pointer)
}

type QSslDiffieHellmanParameters struct {
	internal.Internal
}

type QSslDiffieHellmanParameters_ITF interface {
	QSslDiffieHellmanParameters_PTR() *QSslDiffieHellmanParameters
}

func (ptr *QSslDiffieHellmanParameters) QSslDiffieHellmanParameters_PTR() *QSslDiffieHellmanParameters {
	return ptr
}

func (ptr *QSslDiffieHellmanParameters) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSslDiffieHellmanParameters) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSslDiffieHellmanParameters(ptr QSslDiffieHellmanParameters_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslDiffieHellmanParameters_PTR().Pointer()
	}
	return nil
}

func (n *QSslDiffieHellmanParameters) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslDiffieHellmanParametersFromPointer(ptr unsafe.Pointer) (n *QSslDiffieHellmanParameters) {
	n = new(QSslDiffieHellmanParameters)
	n.InitFromInternal(uintptr(ptr), "network.QSslDiffieHellmanParameters")
	return
}

//go:generate stringer -type=QSslDiffieHellmanParameters__Error
//QSslDiffieHellmanParameters::Error
type QSslDiffieHellmanParameters__Error int64

const (
	QSslDiffieHellmanParameters__NoError               QSslDiffieHellmanParameters__Error = QSslDiffieHellmanParameters__Error(0)
	QSslDiffieHellmanParameters__InvalidInputDataError QSslDiffieHellmanParameters__Error = QSslDiffieHellmanParameters__Error(1)
	QSslDiffieHellmanParameters__UnsafeParametersError QSslDiffieHellmanParameters__Error = QSslDiffieHellmanParameters__Error(2)
)

func NewQSslDiffieHellmanParameters() *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslDiffieHellmanParameters", ""}).(*QSslDiffieHellmanParameters)
}

func NewQSslDiffieHellmanParameters2(other QSslDiffieHellmanParameters_ITF) *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslDiffieHellmanParameters2", "", other}).(*QSslDiffieHellmanParameters)
}

func NewQSslDiffieHellmanParameters3(other QSslDiffieHellmanParameters_ITF) *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslDiffieHellmanParameters3", "", other}).(*QSslDiffieHellmanParameters)
}

func QSslDiffieHellmanParameters_DefaultParameters() *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslDiffieHellmanParameters_DefaultParameters", ""}).(*QSslDiffieHellmanParameters)
}

func (ptr *QSslDiffieHellmanParameters) DefaultParameters() *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslDiffieHellmanParameters_DefaultParameters", ""}).(*QSslDiffieHellmanParameters)
}

func (ptr *QSslDiffieHellmanParameters) Error() QSslDiffieHellmanParameters__Error {

	return QSslDiffieHellmanParameters__Error(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QSslDiffieHellmanParameters) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func QSslDiffieHellmanParameters_FromEncoded(encoded core.QByteArray_ITF, encoding QSsl__EncodingFormat) *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslDiffieHellmanParameters_FromEncoded", "", encoded, encoding}).(*QSslDiffieHellmanParameters)
}

func (ptr *QSslDiffieHellmanParameters) FromEncoded(encoded core.QByteArray_ITF, encoding QSsl__EncodingFormat) *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslDiffieHellmanParameters_FromEncoded", "", encoded, encoding}).(*QSslDiffieHellmanParameters)
}

func QSslDiffieHellmanParameters_FromEncoded2(device core.QIODevice_ITF, encoding QSsl__EncodingFormat) *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslDiffieHellmanParameters_FromEncoded2", "", device, encoding}).(*QSslDiffieHellmanParameters)
}

func (ptr *QSslDiffieHellmanParameters) FromEncoded2(device core.QIODevice_ITF, encoding QSsl__EncodingFormat) *QSslDiffieHellmanParameters {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslDiffieHellmanParameters_FromEncoded2", "", device, encoding}).(*QSslDiffieHellmanParameters)
}

func (ptr *QSslDiffieHellmanParameters) IsEmpty() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEmpty"}).(bool)
}

func (ptr *QSslDiffieHellmanParameters) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSslDiffieHellmanParameters) Swap(other QSslDiffieHellmanParameters_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QSslDiffieHellmanParameters) DestroyQSslDiffieHellmanParameters() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslDiffieHellmanParameters"})
}

type QSslEllipticCurve struct {
	internal.Internal
}

type QSslEllipticCurve_ITF interface {
	QSslEllipticCurve_PTR() *QSslEllipticCurve
}

func (ptr *QSslEllipticCurve) QSslEllipticCurve_PTR() *QSslEllipticCurve {
	return ptr
}

func (ptr *QSslEllipticCurve) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSslEllipticCurve) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSslEllipticCurve(ptr QSslEllipticCurve_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslEllipticCurve_PTR().Pointer()
	}
	return nil
}

func (n *QSslEllipticCurve) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslEllipticCurveFromPointer(ptr unsafe.Pointer) (n *QSslEllipticCurve) {
	n = new(QSslEllipticCurve)
	n.InitFromInternal(uintptr(ptr), "network.QSslEllipticCurve")
	return
}

func (ptr *QSslEllipticCurve) DestroyQSslEllipticCurve() {
}

func NewQSslEllipticCurve() *QSslEllipticCurve {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslEllipticCurve", ""}).(*QSslEllipticCurve)
}

func QSslEllipticCurve_FromLongName(name string) *QSslEllipticCurve {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslEllipticCurve_FromLongName", "", name}).(*QSslEllipticCurve)
}

func (ptr *QSslEllipticCurve) FromLongName(name string) *QSslEllipticCurve {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslEllipticCurve_FromLongName", "", name}).(*QSslEllipticCurve)
}

func QSslEllipticCurve_FromShortName(name string) *QSslEllipticCurve {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslEllipticCurve_FromShortName", "", name}).(*QSslEllipticCurve)
}

func (ptr *QSslEllipticCurve) FromShortName(name string) *QSslEllipticCurve {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslEllipticCurve_FromShortName", "", name}).(*QSslEllipticCurve)
}

func (ptr *QSslEllipticCurve) IsTlsNamedCurve() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsTlsNamedCurve"}).(bool)
}

func (ptr *QSslEllipticCurve) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QSslEllipticCurve) LongName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LongName"}).(string)
}

func (ptr *QSslEllipticCurve) ShortName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ShortName"}).(string)
}

type QSslError struct {
	internal.Internal
}

type QSslError_ITF interface {
	QSslError_PTR() *QSslError
}

func (ptr *QSslError) QSslError_PTR() *QSslError {
	return ptr
}

func (ptr *QSslError) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSslError) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSslError(ptr QSslError_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslError_PTR().Pointer()
	}
	return nil
}

func (n *QSslError) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslErrorFromPointer(ptr unsafe.Pointer) (n *QSslError) {
	n = new(QSslError)
	n.InitFromInternal(uintptr(ptr), "network.QSslError")
	return
}

//go:generate stringer -type=QSslError__SslError
//QSslError::SslError
type QSslError__SslError int64

const (
	QSslError__NoError                             QSslError__SslError = QSslError__SslError(0)
	QSslError__UnableToGetIssuerCertificate        QSslError__SslError = QSslError__SslError(1)
	QSslError__UnableToDecryptCertificateSignature QSslError__SslError = QSslError__SslError(2)
	QSslError__UnableToDecodeIssuerPublicKey       QSslError__SslError = QSslError__SslError(3)
	QSslError__CertificateSignatureFailed          QSslError__SslError = QSslError__SslError(4)
	QSslError__CertificateNotYetValid              QSslError__SslError = QSslError__SslError(5)
	QSslError__CertificateExpired                  QSslError__SslError = QSslError__SslError(6)
	QSslError__InvalidNotBeforeField               QSslError__SslError = QSslError__SslError(7)
	QSslError__InvalidNotAfterField                QSslError__SslError = QSslError__SslError(8)
	QSslError__SelfSignedCertificate               QSslError__SslError = QSslError__SslError(9)
	QSslError__SelfSignedCertificateInChain        QSslError__SslError = QSslError__SslError(10)
	QSslError__UnableToGetLocalIssuerCertificate   QSslError__SslError = QSslError__SslError(11)
	QSslError__UnableToVerifyFirstCertificate      QSslError__SslError = QSslError__SslError(12)
	QSslError__CertificateRevoked                  QSslError__SslError = QSslError__SslError(13)
	QSslError__InvalidCaCertificate                QSslError__SslError = QSslError__SslError(14)
	QSslError__PathLengthExceeded                  QSslError__SslError = QSslError__SslError(15)
	QSslError__InvalidPurpose                      QSslError__SslError = QSslError__SslError(16)
	QSslError__CertificateUntrusted                QSslError__SslError = QSslError__SslError(17)
	QSslError__CertificateRejected                 QSslError__SslError = QSslError__SslError(18)
	QSslError__SubjectIssuerMismatch               QSslError__SslError = QSslError__SslError(19)
	QSslError__AuthorityIssuerSerialNumberMismatch QSslError__SslError = QSslError__SslError(20)
	QSslError__NoPeerCertificate                   QSslError__SslError = QSslError__SslError(21)
	QSslError__HostNameMismatch                    QSslError__SslError = QSslError__SslError(22)
	QSslError__NoSslSupport                        QSslError__SslError = QSslError__SslError(23)
	QSslError__CertificateBlacklisted              QSslError__SslError = QSslError__SslError(24)
	QSslError__CertificateStatusUnknown            QSslError__SslError = QSslError__SslError(25)
	QSslError__OcspNoResponseFound                 QSslError__SslError = QSslError__SslError(26)
	QSslError__OcspMalformedRequest                QSslError__SslError = QSslError__SslError(27)
	QSslError__OcspMalformedResponse               QSslError__SslError = QSslError__SslError(28)
	QSslError__OcspInternalError                   QSslError__SslError = QSslError__SslError(29)
	QSslError__OcspTryLater                        QSslError__SslError = QSslError__SslError(30)
	QSslError__OcspSigRequred                      QSslError__SslError = QSslError__SslError(31)
	QSslError__OcspUnauthorized                    QSslError__SslError = QSslError__SslError(32)
	QSslError__OcspResponseCannotBeTrusted         QSslError__SslError = QSslError__SslError(33)
	QSslError__OcspResponseCertIdUnknown           QSslError__SslError = QSslError__SslError(34)
	QSslError__OcspResponseExpired                 QSslError__SslError = QSslError__SslError(35)
	QSslError__OcspStatusUnknown                   QSslError__SslError = QSslError__SslError(36)
	QSslError__UnspecifiedError                    QSslError__SslError = QSslError__SslError(-1)
)

func NewQSslError() *QSslError {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslError", ""}).(*QSslError)
}

func NewQSslError2(error QSslError__SslError) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslError2", "", error}).(*QSslError)
}

func NewQSslError3(error QSslError__SslError, certificate QSslCertificate_ITF) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslError3", "", error, certificate}).(*QSslError)
}

func NewQSslError4(other QSslError_ITF) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslError4", "", other}).(*QSslError)
}

func (ptr *QSslError) Certificate() *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Certificate"}).(*QSslCertificate)
}

func (ptr *QSslError) Error() QSslError__SslError {

	return QSslError__SslError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QSslError) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QSslError) Swap(other QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QSslError) DestroyQSslError() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslError"})
}

type QSslKey struct {
	internal.Internal
}

type QSslKey_ITF interface {
	QSslKey_PTR() *QSslKey
}

func (ptr *QSslKey) QSslKey_PTR() *QSslKey {
	return ptr
}

func (ptr *QSslKey) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSslKey) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSslKey(ptr QSslKey_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslKey_PTR().Pointer()
	}
	return nil
}

func (n *QSslKey) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslKeyFromPointer(ptr unsafe.Pointer) (n *QSslKey) {
	n = new(QSslKey)
	n.InitFromInternal(uintptr(ptr), "network.QSslKey")
	return
}
func NewQSslKey() *QSslKey {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslKey", ""}).(*QSslKey)
}

func NewQSslKey2(encoded core.QByteArray_ITF, algorithm QSsl__KeyAlgorithm, encoding QSsl__EncodingFormat, ty QSsl__KeyType, passPhrase core.QByteArray_ITF) *QSslKey {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslKey2", "", encoded, algorithm, encoding, ty, passPhrase}).(*QSslKey)
}

func NewQSslKey3(device core.QIODevice_ITF, algorithm QSsl__KeyAlgorithm, encoding QSsl__EncodingFormat, ty QSsl__KeyType, passPhrase core.QByteArray_ITF) *QSslKey {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslKey3", "", device, algorithm, encoding, ty, passPhrase}).(*QSslKey)
}

func NewQSslKey5(other QSslKey_ITF) *QSslKey {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslKey5", "", other}).(*QSslKey)
}

func (ptr *QSslKey) Algorithm() QSsl__KeyAlgorithm {

	return QSsl__KeyAlgorithm(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Algorithm"}).(float64))
}

func (ptr *QSslKey) Clear() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Clear"})
}

func (ptr *QSslKey) IsNull() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsNull"}).(bool)
}

func (ptr *QSslKey) Length() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Length"}).(float64))
}

func (ptr *QSslKey) Swap(other QSslKey_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QSslKey) ToDer(passPhrase core.QByteArray_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToDer", passPhrase}).(*core.QByteArray)
}

func (ptr *QSslKey) ToPem(passPhrase core.QByteArray_ITF) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ToPem", passPhrase}).(*core.QByteArray)
}

func (ptr *QSslKey) Type() QSsl__KeyType {

	return QSsl__KeyType(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Type"}).(float64))
}

func (ptr *QSslKey) DestroyQSslKey() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslKey"})
}

type QSslPreSharedKeyAuthenticator struct {
	internal.Internal
}

type QSslPreSharedKeyAuthenticator_ITF interface {
	QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator
}

func (ptr *QSslPreSharedKeyAuthenticator) QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator {
	return ptr
}

func (ptr *QSslPreSharedKeyAuthenticator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQSslPreSharedKeyAuthenticator(ptr QSslPreSharedKeyAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslPreSharedKeyAuthenticator_PTR().Pointer()
	}
	return nil
}

func (n *QSslPreSharedKeyAuthenticator) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQSslPreSharedKeyAuthenticatorFromPointer(ptr unsafe.Pointer) (n *QSslPreSharedKeyAuthenticator) {
	n = new(QSslPreSharedKeyAuthenticator)
	n.InitFromInternal(uintptr(ptr), "network.QSslPreSharedKeyAuthenticator")
	return
}
func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslPreSharedKeyAuthenticator", ""}).(*QSslPreSharedKeyAuthenticator)
}

func NewQSslPreSharedKeyAuthenticator2(authenticator QSslPreSharedKeyAuthenticator_ITF) *QSslPreSharedKeyAuthenticator {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslPreSharedKeyAuthenticator2", "", authenticator}).(*QSslPreSharedKeyAuthenticator)
}

func (ptr *QSslPreSharedKeyAuthenticator) Identity() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Identity"}).(*core.QByteArray)
}

func (ptr *QSslPreSharedKeyAuthenticator) IdentityHint() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IdentityHint"}).(*core.QByteArray)
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumIdentityLength"}).(float64))
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaximumPreSharedKeyLength"}).(float64))
}

func (ptr *QSslPreSharedKeyAuthenticator) PreSharedKey() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreSharedKey"}).(*core.QByteArray)
}

func (ptr *QSslPreSharedKeyAuthenticator) SetIdentity(identity core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetIdentity", identity})
}

func (ptr *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPreSharedKey", preSharedKey})
}

func (ptr *QSslPreSharedKeyAuthenticator) Swap(authenticator QSslPreSharedKeyAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", authenticator})
}

func (ptr *QSslPreSharedKeyAuthenticator) DestroyQSslPreSharedKeyAuthenticator() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslPreSharedKeyAuthenticator"})
}

type QSslSocket struct {
	QTcpSocket
}

type QSslSocket_ITF interface {
	QTcpSocket_ITF
	QSslSocket_PTR() *QSslSocket
}

func (ptr *QSslSocket) QSslSocket_PTR() *QSslSocket {
	return ptr
}

func (ptr *QSslSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func (ptr *QSslSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTcpSocket_PTR().SetPointer(p)
	}
}

func PointerFromQSslSocket(ptr QSslSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslSocket_PTR().Pointer()
	}
	return nil
}

func (n *QSslSocket) InitFromInternal(ptr uintptr, name string) {
	n.QTcpSocket_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QSslSocket) ClassNameInternalF() string {
	return n.QTcpSocket_PTR().ClassNameInternalF()
}

func NewQSslSocketFromPointer(ptr unsafe.Pointer) (n *QSslSocket) {
	n = new(QSslSocket)
	n.InitFromInternal(uintptr(ptr), "network.QSslSocket")
	return
}

//go:generate stringer -type=QSslSocket__SslMode
//QSslSocket::SslMode
type QSslSocket__SslMode int64

const (
	QSslSocket__UnencryptedMode QSslSocket__SslMode = QSslSocket__SslMode(0)
	QSslSocket__SslClientMode   QSslSocket__SslMode = QSslSocket__SslMode(1)
	QSslSocket__SslServerMode   QSslSocket__SslMode = QSslSocket__SslMode(2)
)

//go:generate stringer -type=QSslSocket__PeerVerifyMode
//QSslSocket::PeerVerifyMode
type QSslSocket__PeerVerifyMode int64

const (
	QSslSocket__VerifyNone     QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(0)
	QSslSocket__QueryPeer      QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(1)
	QSslSocket__VerifyPeer     QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(2)
	QSslSocket__AutoVerifyPeer QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(3)
)

func NewQSslSocket(parent core.QObject_ITF) *QSslSocket {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQSslSocket", "", parent}).(*QSslSocket)
}

func (ptr *QSslSocket) AddCaCertificate(certificate QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddCaCertificate", certificate})
}

func (ptr *QSslSocket) AddCaCertificates(path string, format QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddCaCertificates", path, format, syntax}).(bool)
}

func (ptr *QSslSocket) AddCaCertificates2(certificates []*QSslCertificate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddCaCertificates2", certificates})
}

func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_AddDefaultCaCertificate", "", certificate})
}

func (ptr *QSslSocket) AddDefaultCaCertificate(certificate QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_AddDefaultCaCertificate", "", certificate})
}

func QSslSocket_AddDefaultCaCertificates(path string, encoding QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_AddDefaultCaCertificates", "", path, encoding, syntax}).(bool)
}

func (ptr *QSslSocket) AddDefaultCaCertificates(path string, encoding QSsl__EncodingFormat, syntax core.QRegExp__PatternSyntax) bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_AddDefaultCaCertificates", "", path, encoding, syntax}).(bool)
}

func QSslSocket_AddDefaultCaCertificates2(certificates []*QSslCertificate) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_AddDefaultCaCertificates2", "", certificates})
}

func (ptr *QSslSocket) AddDefaultCaCertificates2(certificates []*QSslCertificate) {

	internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_AddDefaultCaCertificates2", "", certificates})
}

func (ptr *QSslSocket) ConnectToHostEncrypted(hostName string, port uint16, mode core.QIODevice__OpenModeFlag, protoc QAbstractSocket__NetworkLayerProtocol) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToHostEncrypted", hostName, port, mode, protoc})
}

func (ptr *QSslSocket) ConnectToHostEncrypted2(hostName string, port uint16, sslPeerName string, mode core.QIODevice__OpenModeFlag, protoc QAbstractSocket__NetworkLayerProtocol) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectToHostEncrypted2", hostName, port, sslPeerName, mode, protoc})
}

func (ptr *QSslSocket) ConnectEncrypted(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEncrypted", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectEncrypted() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEncrypted"})
}

func (ptr *QSslSocket) Encrypted() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Encrypted"})
}

func (ptr *QSslSocket) EncryptedBytesAvailable() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EncryptedBytesAvailable"}).(float64))
}

func (ptr *QSslSocket) EncryptedBytesToWrite() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EncryptedBytesToWrite"}).(float64))
}

func (ptr *QSslSocket) ConnectEncryptedBytesWritten(f func(written int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectEncryptedBytesWritten", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectEncryptedBytesWritten() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectEncryptedBytesWritten"})
}

func (ptr *QSslSocket) EncryptedBytesWritten(written int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EncryptedBytesWritten", written})
}

func (ptr *QSslSocket) ConnectIgnoreSslErrors(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIgnoreSslErrors", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectIgnoreSslErrors() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIgnoreSslErrors"})
}

func (ptr *QSslSocket) IgnoreSslErrors() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrors"})
}

func (ptr *QSslSocket) IgnoreSslErrorsDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrorsDefault"})
}

func (ptr *QSslSocket) IgnoreSslErrors2(errors []*QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrors2", errors})
}

func (ptr *QSslSocket) IsEncrypted() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsEncrypted"}).(bool)
}

func (ptr *QSslSocket) LocalCertificate() *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalCertificate"}).(*QSslCertificate)
}

func (ptr *QSslSocket) LocalCertificateChain() []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalCertificateChain"}).([]*QSslCertificate)
}

func (ptr *QSslSocket) Mode() QSslSocket__SslMode {

	return QSslSocket__SslMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Mode"}).(float64))
}

func (ptr *QSslSocket) ConnectModeChanged(f func(mode QSslSocket__SslMode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectModeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectModeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectModeChanged"})
}

func (ptr *QSslSocket) ModeChanged(mode QSslSocket__SslMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ModeChanged", mode})
}

func (ptr *QSslSocket) OcspResponses() []*QOcspResponse {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OcspResponses"}).([]*QOcspResponse)
}

func (ptr *QSslSocket) PeerCertificate() *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerCertificate"}).(*QSslCertificate)
}

func (ptr *QSslSocket) PeerCertificateChain() []*QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerCertificateChain"}).([]*QSslCertificate)
}

func (ptr *QSslSocket) PeerVerifyDepth() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerVerifyDepth"}).(float64))
}

func (ptr *QSslSocket) ConnectPeerVerifyError(f func(error *QSslError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPeerVerifyError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectPeerVerifyError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPeerVerifyError"})
}

func (ptr *QSslSocket) PeerVerifyError(error QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerVerifyError", error})
}

func (ptr *QSslSocket) PeerVerifyMode() QSslSocket__PeerVerifyMode {

	return QSslSocket__PeerVerifyMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerVerifyMode"}).(float64))
}

func (ptr *QSslSocket) PeerVerifyName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerVerifyName"}).(string)
}

func (ptr *QSslSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreSharedKeyAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectPreSharedKeyAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreSharedKeyAuthenticationRequired"})
}

func (ptr *QSslSocket) PreSharedKeyAuthenticationRequired(authenticator QSslPreSharedKeyAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreSharedKeyAuthenticationRequired", authenticator})
}

func (ptr *QSslSocket) PrivateKey() *QSslKey {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PrivateKey"}).(*QSslKey)
}

func (ptr *QSslSocket) Protocol() QSsl__SslProtocol {

	return QSsl__SslProtocol(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Protocol"}).(float64))
}

func (ptr *QSslSocket) SessionCipher() *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SessionCipher"}).(*QSslCipher)
}

func (ptr *QSslSocket) SessionProtocol() QSsl__SslProtocol {

	return QSsl__SslProtocol(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SessionProtocol"}).(float64))
}

func (ptr *QSslSocket) SetLocalCertificate(certificate QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalCertificate", certificate})
}

func (ptr *QSslSocket) SetLocalCertificate2(path string, format QSsl__EncodingFormat) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalCertificate2", path, format})
}

func (ptr *QSslSocket) SetLocalCertificateChain(localChain []*QSslCertificate) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocalCertificateChain", localChain})
}

func (ptr *QSslSocket) SetPeerVerifyDepth(depth int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerVerifyDepth", depth})
}

func (ptr *QSslSocket) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerVerifyMode", mode})
}

func (ptr *QSslSocket) SetPeerVerifyName(hostName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPeerVerifyName", hostName})
}

func (ptr *QSslSocket) SetPrivateKey(key QSslKey_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrivateKey", key})
}

func (ptr *QSslSocket) SetPrivateKey2(fileName string, algorithm QSsl__KeyAlgorithm, format QSsl__EncodingFormat, passPhrase core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPrivateKey2", fileName, algorithm, format, passPhrase})
}

func (ptr *QSslSocket) SetProtocol(protoc QSsl__SslProtocol) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProtocol", protoc})
}

func (ptr *QSslSocket) SetSslConfiguration(configuration QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSslConfiguration", configuration})
}

func (ptr *QSslSocket) SslConfiguration() *QSslConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslConfiguration"}).(*QSslConfiguration)
}

func (ptr *QSslSocket) SslErrors() []*QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslErrors"}).([]*QSslError)
}

func (ptr *QSslSocket) ConnectSslErrors2(f func(errors []*QSslError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSslErrors2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectSslErrors2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSslErrors2"})
}

func (ptr *QSslSocket) SslErrors2(errors []*QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslErrors2", errors})
}

func QSslSocket_SslLibraryBuildVersionNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SslLibraryBuildVersionNumber", ""}).(float64))
}

func (ptr *QSslSocket) SslLibraryBuildVersionNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SslLibraryBuildVersionNumber", ""}).(float64))
}

func QSslSocket_SslLibraryBuildVersionString() string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SslLibraryBuildVersionString", ""}).(string)
}

func (ptr *QSslSocket) SslLibraryBuildVersionString() string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SslLibraryBuildVersionString", ""}).(string)
}

func QSslSocket_SslLibraryVersionNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SslLibraryVersionNumber", ""}).(float64))
}

func (ptr *QSslSocket) SslLibraryVersionNumber() int {

	return int(internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SslLibraryVersionNumber", ""}).(float64))
}

func QSslSocket_SslLibraryVersionString() string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SslLibraryVersionString", ""}).(string)
}

func (ptr *QSslSocket) SslLibraryVersionString() string {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SslLibraryVersionString", ""}).(string)
}

func (ptr *QSslSocket) ConnectStartClientEncryption(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartClientEncryption", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectStartClientEncryption() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartClientEncryption"})
}

func (ptr *QSslSocket) StartClientEncryption() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartClientEncryption"})
}

func (ptr *QSslSocket) StartClientEncryptionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartClientEncryptionDefault"})
}

func (ptr *QSslSocket) ConnectStartServerEncryption(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStartServerEncryption", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectStartServerEncryption() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStartServerEncryption"})
}

func (ptr *QSslSocket) StartServerEncryption() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartServerEncryption"})
}

func (ptr *QSslSocket) StartServerEncryptionDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StartServerEncryptionDefault"})
}

func QSslSocket_SupportsSsl() bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SupportsSsl", ""}).(bool)
}

func (ptr *QSslSocket) SupportsSsl() bool {

	return internal.CallLocalFunction([]interface{}{"", "", "network.QSslSocket_SupportsSsl", ""}).(bool)
}

func (ptr *QSslSocket) WaitForEncrypted(msecs int) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForEncrypted", msecs}).(bool)
}

func (ptr *QSslSocket) ConnectDestroyQSslSocket(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQSslSocket", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QSslSocket) DisconnectDestroyQSslSocket() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQSslSocket"})
}

func (ptr *QSslSocket) DestroyQSslSocket() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslSocket"})
}

func (ptr *QSslSocket) DestroyQSslSocketDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQSslSocketDefault"})
}

func (ptr *QSslSocket) __addCaCertificates_certificates_atList2(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addCaCertificates_certificates_atList2", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __addCaCertificates_certificates_setList2(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addCaCertificates_certificates_setList2", i})
}

func (ptr *QSslSocket) __addCaCertificates_certificates_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addCaCertificates_certificates_newList2"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __addDefaultCaCertificates_certificates_atList2(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addDefaultCaCertificates_certificates_atList2", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __addDefaultCaCertificates_certificates_setList2(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addDefaultCaCertificates_certificates_setList2", i})
}

func (ptr *QSslSocket) __addDefaultCaCertificates_certificates_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__addDefaultCaCertificates_certificates_newList2"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __caCertificates_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__caCertificates_atList", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __caCertificates_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__caCertificates_setList", i})
}

func (ptr *QSslSocket) __caCertificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__caCertificates_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __ciphers_atList(i int) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ciphers_atList", i}).(*QSslCipher)
}

func (ptr *QSslSocket) __ciphers_setList(i QSslCipher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ciphers_setList", i})
}

func (ptr *QSslSocket) __ciphers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ciphers_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __defaultCaCertificates_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__defaultCaCertificates_atList", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __defaultCaCertificates_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__defaultCaCertificates_setList", i})
}

func (ptr *QSslSocket) __defaultCaCertificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__defaultCaCertificates_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __defaultCiphers_atList(i int) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__defaultCiphers_atList", i}).(*QSslCipher)
}

func (ptr *QSslSocket) __defaultCiphers_setList(i QSslCipher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__defaultCiphers_setList", i})
}

func (ptr *QSslSocket) __defaultCiphers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__defaultCiphers_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __ignoreSslErrors_errors_atList2(i int) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrors_errors_atList2", i}).(*QSslError)
}

func (ptr *QSslSocket) __ignoreSslErrors_errors_setList2(i QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrors_errors_setList2", i})
}

func (ptr *QSslSocket) __ignoreSslErrors_errors_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrors_errors_newList2"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __localCertificateChain_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__localCertificateChain_atList", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __localCertificateChain_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__localCertificateChain_setList", i})
}

func (ptr *QSslSocket) __localCertificateChain_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__localCertificateChain_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __ocspResponses_atList(i int) *QOcspResponse {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ocspResponses_atList", i}).(*QOcspResponse)
}

func (ptr *QSslSocket) __ocspResponses_setList(i QOcspResponse_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ocspResponses_setList", i})
}

func (ptr *QSslSocket) __ocspResponses_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ocspResponses_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __peerCertificateChain_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__peerCertificateChain_atList", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __peerCertificateChain_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__peerCertificateChain_setList", i})
}

func (ptr *QSslSocket) __peerCertificateChain_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__peerCertificateChain_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __setCaCertificates_certificates_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCaCertificates_certificates_atList", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __setCaCertificates_certificates_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCaCertificates_certificates_setList", i})
}

func (ptr *QSslSocket) __setCaCertificates_certificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCaCertificates_certificates_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __setCiphers_ciphers_atList(i int) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCiphers_ciphers_atList", i}).(*QSslCipher)
}

func (ptr *QSslSocket) __setCiphers_ciphers_setList(i QSslCipher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCiphers_ciphers_setList", i})
}

func (ptr *QSslSocket) __setCiphers_ciphers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setCiphers_ciphers_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __setDefaultCaCertificates_certificates_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setDefaultCaCertificates_certificates_atList", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __setDefaultCaCertificates_certificates_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setDefaultCaCertificates_certificates_setList", i})
}

func (ptr *QSslSocket) __setDefaultCaCertificates_certificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setDefaultCaCertificates_certificates_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __setDefaultCiphers_ciphers_atList(i int) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setDefaultCiphers_ciphers_atList", i}).(*QSslCipher)
}

func (ptr *QSslSocket) __setDefaultCiphers_ciphers_setList(i QSslCipher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setDefaultCiphers_ciphers_setList", i})
}

func (ptr *QSslSocket) __setDefaultCiphers_ciphers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setDefaultCiphers_ciphers_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __setLocalCertificateChain_localChain_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setLocalCertificateChain_localChain_atList", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __setLocalCertificateChain_localChain_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setLocalCertificateChain_localChain_setList", i})
}

func (ptr *QSslSocket) __setLocalCertificateChain_localChain_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__setLocalCertificateChain_localChain_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __sslErrors_atList(i int) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_atList", i}).(*QSslError)
}

func (ptr *QSslSocket) __sslErrors_setList(i QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_setList", i})
}

func (ptr *QSslSocket) __sslErrors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __sslErrors_errors_atList2(i int) *QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_atList2", i}).(*QSslError)
}

func (ptr *QSslSocket) __sslErrors_errors_setList2(i QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_setList2", i})
}

func (ptr *QSslSocket) __sslErrors_errors_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_newList2"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __supportedCiphers_atList(i int) *QSslCipher {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedCiphers_atList", i}).(*QSslCipher)
}

func (ptr *QSslSocket) __supportedCiphers_setList(i QSslCipher_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedCiphers_setList", i})
}

func (ptr *QSslSocket) __supportedCiphers_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedCiphers_newList"}).(unsafe.Pointer)
}

func (ptr *QSslSocket) __systemCaCertificates_atList(i int) *QSslCertificate {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__systemCaCertificates_atList", i}).(*QSslCertificate)
}

func (ptr *QSslSocket) __systemCaCertificates_setList(i QSslCertificate_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__systemCaCertificates_setList", i})
}

func (ptr *QSslSocket) __systemCaCertificates_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__systemCaCertificates_newList"}).(unsafe.Pointer)
}

type QTcpServer struct {
	core.QObject
}

type QTcpServer_ITF interface {
	core.QObject_ITF
	QTcpServer_PTR() *QTcpServer
}

func (ptr *QTcpServer) QTcpServer_PTR() *QTcpServer {
	return ptr
}

func (ptr *QTcpServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTcpServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQTcpServer(ptr QTcpServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpServer_PTR().Pointer()
	}
	return nil
}

func (n *QTcpServer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTcpServer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTcpServerFromPointer(ptr unsafe.Pointer) (n *QTcpServer) {
	n = new(QTcpServer)
	n.InitFromInternal(uintptr(ptr), "network.QTcpServer")
	return
}
func NewQTcpServer(parent core.QObject_ITF) *QTcpServer {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQTcpServer", "", parent}).(*QTcpServer)
}

func (ptr *QTcpServer) ConnectAcceptError(f func(socketError QAbstractSocket__SocketError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAcceptError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTcpServer) DisconnectAcceptError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAcceptError"})
}

func (ptr *QTcpServer) AcceptError(socketError QAbstractSocket__SocketError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcceptError", socketError})
}

func (ptr *QTcpServer) AddPendingConnection(socket QTcpSocket_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AddPendingConnection", socket})
}

func (ptr *QTcpServer) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QTcpServer) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QTcpServer) ConnectHasPendingConnections(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectHasPendingConnections", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTcpServer) DisconnectHasPendingConnections() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectHasPendingConnections"})
}

func (ptr *QTcpServer) HasPendingConnections() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasPendingConnections"}).(bool)
}

func (ptr *QTcpServer) HasPendingConnectionsDefault() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasPendingConnectionsDefault"}).(bool)
}

func (ptr *QTcpServer) IsListening() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsListening"}).(bool)
}

func (ptr *QTcpServer) Listen(address QHostAddress_ITF, port uint16) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Listen", address, port}).(bool)
}

func (ptr *QTcpServer) MaxPendingConnections() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxPendingConnections"}).(float64))
}

func (ptr *QTcpServer) ConnectNewConnection(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTcpServer) DisconnectNewConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewConnection"})
}

func (ptr *QTcpServer) NewConnection() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewConnection"})
}

func (ptr *QTcpServer) ConnectNextPendingConnection(f func() *QTcpSocket) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNextPendingConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTcpServer) DisconnectNextPendingConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNextPendingConnection"})
}

func (ptr *QTcpServer) NextPendingConnection() *QTcpSocket {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextPendingConnection"}).(*QTcpSocket)
}

func (ptr *QTcpServer) NextPendingConnectionDefault() *QTcpSocket {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextPendingConnectionDefault"}).(*QTcpSocket)
}

func (ptr *QTcpServer) PauseAccepting() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PauseAccepting"})
}

func (ptr *QTcpServer) Proxy() *QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Proxy"}).(*QNetworkProxy)
}

func (ptr *QTcpServer) ResumeAccepting() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResumeAccepting"})
}

func (ptr *QTcpServer) ServerAddress() *QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerAddress"}).(*QHostAddress)
}

func (ptr *QTcpServer) ServerError() QAbstractSocket__SocketError {

	return QAbstractSocket__SocketError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerError"}).(float64))
}

func (ptr *QTcpServer) ServerPort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerPort"}).(float64))
}

func (ptr *QTcpServer) SetMaxPendingConnections(numConnections int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaxPendingConnections", numConnections})
}

func (ptr *QTcpServer) SetProxy(networkProxy QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProxy", networkProxy})
}

func (ptr *QTcpServer) WaitForNewConnection(msec int, timedOut *bool) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WaitForNewConnection", msec, timedOut}).(bool)
}

func (ptr *QTcpServer) ConnectDestroyQTcpServer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQTcpServer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTcpServer) DisconnectDestroyQTcpServer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQTcpServer"})
}

func (ptr *QTcpServer) DestroyQTcpServer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTcpServer"})
}

func (ptr *QTcpServer) DestroyQTcpServerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTcpServerDefault"})
}

func (ptr *QTcpServer) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QTcpServer) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QTcpServer) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QTcpServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QTcpServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QTcpServer) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QTcpServer) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QTcpServer) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QTcpServer) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QTcpServer) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QTcpServer) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QTcpServer) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QTcpServer) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QTcpServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QTcpServer) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QTcpServer) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QTcpServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QTcpServer) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QTcpServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QTcpServer) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QTcpServer) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QTcpSocket struct {
	QAbstractSocket
}

type QTcpSocket_ITF interface {
	QAbstractSocket_ITF
	QTcpSocket_PTR() *QTcpSocket
}

func (ptr *QTcpSocket) QTcpSocket_PTR() *QTcpSocket {
	return ptr
}

func (ptr *QTcpSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func (ptr *QTcpSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSocket_PTR().SetPointer(p)
	}
}

func PointerFromQTcpSocket(ptr QTcpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTcpSocket_PTR().Pointer()
	}
	return nil
}

func (n *QTcpSocket) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractSocket_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTcpSocket) ClassNameInternalF() string {
	return n.QAbstractSocket_PTR().ClassNameInternalF()
}

func NewQTcpSocketFromPointer(ptr unsafe.Pointer) (n *QTcpSocket) {
	n = new(QTcpSocket)
	n.InitFromInternal(uintptr(ptr), "network.QTcpSocket")
	return
}
func NewQTcpSocket(parent core.QObject_ITF) *QTcpSocket {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQTcpSocket", "", parent}).(*QTcpSocket)
}

func (ptr *QTcpSocket) ConnectDestroyQTcpSocket(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQTcpSocket", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTcpSocket) DisconnectDestroyQTcpSocket() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQTcpSocket"})
}

func (ptr *QTcpSocket) DestroyQTcpSocket() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTcpSocket"})
}

func (ptr *QTcpSocket) DestroyQTcpSocketDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQTcpSocketDefault"})
}

type QUdpSocket struct {
	QAbstractSocket
}

type QUdpSocket_ITF interface {
	QAbstractSocket_ITF
	QUdpSocket_PTR() *QUdpSocket
}

func (ptr *QUdpSocket) QUdpSocket_PTR() *QUdpSocket {
	return ptr
}

func (ptr *QUdpSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSocket_PTR().Pointer()
	}
	return nil
}

func (ptr *QUdpSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QAbstractSocket_PTR().SetPointer(p)
	}
}

func PointerFromQUdpSocket(ptr QUdpSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUdpSocket_PTR().Pointer()
	}
	return nil
}

func (n *QUdpSocket) InitFromInternal(ptr uintptr, name string) {
	n.QAbstractSocket_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QUdpSocket) ClassNameInternalF() string {
	return n.QAbstractSocket_PTR().ClassNameInternalF()
}

func NewQUdpSocketFromPointer(ptr unsafe.Pointer) (n *QUdpSocket) {
	n = new(QUdpSocket)
	n.InitFromInternal(uintptr(ptr), "network.QUdpSocket")
	return
}
func NewQUdpSocket(parent core.QObject_ITF) *QUdpSocket {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQUdpSocket", "", parent}).(*QUdpSocket)
}

func (ptr *QUdpSocket) HasPendingDatagrams() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasPendingDatagrams"}).(bool)
}

func (ptr *QUdpSocket) JoinMulticastGroup(groupAddress QHostAddress_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JoinMulticastGroup", groupAddress}).(bool)
}

func (ptr *QUdpSocket) JoinMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "JoinMulticastGroup2", groupAddress, iface}).(bool)
}

func (ptr *QUdpSocket) LeaveMulticastGroup(groupAddress QHostAddress_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveMulticastGroup", groupAddress}).(bool)
}

func (ptr *QUdpSocket) LeaveMulticastGroup2(groupAddress QHostAddress_ITF, iface QNetworkInterface_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LeaveMulticastGroup2", groupAddress, iface}).(bool)
}

func (ptr *QUdpSocket) MulticastInterface() *QNetworkInterface {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MulticastInterface"}).(*QNetworkInterface)
}

func (ptr *QUdpSocket) PendingDatagramSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PendingDatagramSize"}).(float64))
}

func (ptr *QUdpSocket) ReadDatagram(data []byte, maxSize int64, address QHostAddress_ITF, port uint16) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadDatagram", data, maxSize, address, port}).(float64))
}

func (ptr *QUdpSocket) ReceiveDatagram(maxSize int64) *QNetworkDatagram {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReceiveDatagram", maxSize}).(*QNetworkDatagram)
}

func (ptr *QUdpSocket) SetMulticastInterface(iface QNetworkInterface_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMulticastInterface", iface})
}

func (ptr *QUdpSocket) WriteDatagram(data []byte, size int64, address QHostAddress_ITF, port uint16) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDatagram", data, size, address, port}).(float64))
}

func (ptr *QUdpSocket) WriteDatagram2(datagram QNetworkDatagram_ITF) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDatagram2", datagram}).(float64))
}

func (ptr *QUdpSocket) WriteDatagram3(datagram core.QByteArray_ITF, host QHostAddress_ITF, port uint16) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "WriteDatagram3", datagram, host, port}).(float64))
}

func (ptr *QUdpSocket) ConnectDestroyQUdpSocket(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQUdpSocket", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QUdpSocket) DisconnectDestroyQUdpSocket() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQUdpSocket"})
}

func (ptr *QUdpSocket) DestroyQUdpSocket() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQUdpSocket"})
}

func (ptr *QUdpSocket) DestroyQUdpSocketDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQUdpSocketDefault"})
}

type QNetworkCacheMetaData_RawHeader struct {
	internal.Internal
}

type QNetworkCacheMetaData_RawHeader_ITF interface {
	QNetworkCacheMetaData_RawHeader_PTR() *QNetworkCacheMetaData_RawHeader
}

func (ptr *QNetworkCacheMetaData_RawHeader) QNetworkCacheMetaData_RawHeader_PTR() *QNetworkCacheMetaData_RawHeader {
	return ptr
}

func (ptr *QNetworkCacheMetaData_RawHeader) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QNetworkCacheMetaData_RawHeader) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQNetworkCacheMetaData_RawHeader(ptr QNetworkCacheMetaData_RawHeader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCacheMetaData_RawHeader_PTR().Pointer()
	}
	return nil
}

func (n *QNetworkCacheMetaData_RawHeader) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQNetworkCacheMetaData_RawHeaderFromPointer(ptr unsafe.Pointer) (n *QNetworkCacheMetaData_RawHeader) {
	n = new(QNetworkCacheMetaData_RawHeader)
	n.InitFromInternal(uintptr(ptr), "network.QNetworkCacheMetaData_RawHeader")
	return
}

func (ptr *QNetworkCacheMetaData_RawHeader) DestroyQNetworkCacheMetaData_RawHeader() {
}

func NewQNetworkCacheMetaData_RawHeader() *QNetworkCacheMetaData_RawHeader {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkCacheMetaData_RawHeader", ""}).(*QNetworkCacheMetaData_RawHeader)
}

func NewQNetworkCacheMetaData_RawHeader2(first core.QByteArray_ITF, second core.QByteArray_ITF) *QNetworkCacheMetaData_RawHeader {

	return internal.CallLocalFunction([]interface{}{"", "", "network.NewQNetworkCacheMetaData_RawHeader2", "", first, second}).(*QNetworkCacheMetaData_RawHeader)
}

func (ptr *QNetworkCacheMetaData_RawHeader) First() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "First"}).(*core.QByteArray)
}

func (ptr *QNetworkCacheMetaData_RawHeader) SetFirst(vqb core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetFirst", vqb})
}

func (ptr *QNetworkCacheMetaData_RawHeader) Second() *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Second"}).(*core.QByteArray)
}

func (ptr *QNetworkCacheMetaData_RawHeader) SetSecond(vqb core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSecond", vqb})
}

func init() {
	internal.ConstructorTable["network.QAbstractNetworkCache"] = NewQAbstractNetworkCacheFromPointer
	internal.ConstructorTable["network.QAbstractSocket"] = NewQAbstractSocketFromPointer
	internal.ConstructorTable["network.QAuthenticator"] = NewQAuthenticatorFromPointer
	internal.ConstructorTable["network.QDnsDomainNameRecord"] = NewQDnsDomainNameRecordFromPointer
	internal.ConstructorTable["network.QDnsHostAddressRecord"] = NewQDnsHostAddressRecordFromPointer
	internal.ConstructorTable["network.QDnsLookup"] = NewQDnsLookupFromPointer
	internal.ConstructorTable["network.QDnsMailExchangeRecord"] = NewQDnsMailExchangeRecordFromPointer
	internal.ConstructorTable["network.QDnsServiceRecord"] = NewQDnsServiceRecordFromPointer
	internal.ConstructorTable["network.QDnsTextRecord"] = NewQDnsTextRecordFromPointer
	internal.ConstructorTable["network.QHostAddress"] = NewQHostAddressFromPointer
	internal.ConstructorTable["network.QHostInfo"] = NewQHostInfoFromPointer
	internal.ConstructorTable["network.QHstsPolicy"] = NewQHstsPolicyFromPointer
	internal.ConstructorTable["network.QHttpMultiPart"] = NewQHttpMultiPartFromPointer
	internal.ConstructorTable["network.QHttpPart"] = NewQHttpPartFromPointer
	internal.ConstructorTable["network.QIPv6Address"] = NewQIPv6AddressFromPointer
	internal.ConstructorTable["network.QLocalServer"] = NewQLocalServerFromPointer
	internal.ConstructorTable["network.QLocalSocket"] = NewQLocalSocketFromPointer
	internal.ConstructorTable["network.QNetworkAccessManager"] = NewQNetworkAccessManagerFromPointer
	internal.ConstructorTable["network.QNetworkAddressEntry"] = NewQNetworkAddressEntryFromPointer
	internal.ConstructorTable["network.QNetworkCacheMetaData"] = NewQNetworkCacheMetaDataFromPointer
	internal.ConstructorTable["network.QNetworkConfiguration"] = NewQNetworkConfigurationFromPointer
	internal.ConstructorTable["network.QNetworkConfigurationManager"] = NewQNetworkConfigurationManagerFromPointer
	internal.ConstructorTable["network.QNetworkCookie"] = NewQNetworkCookieFromPointer
	internal.ConstructorTable["network.QNetworkCookieJar"] = NewQNetworkCookieJarFromPointer
	internal.ConstructorTable["network.QNetworkDatagram"] = NewQNetworkDatagramFromPointer
	internal.ConstructorTable["network.QNetworkDiskCache"] = NewQNetworkDiskCacheFromPointer
	internal.ConstructorTable["network.QNetworkInterface"] = NewQNetworkInterfaceFromPointer
	internal.ConstructorTable["network.QNetworkProxy"] = NewQNetworkProxyFromPointer
	internal.ConstructorTable["network.QNetworkProxyFactory"] = NewQNetworkProxyFactoryFromPointer
	internal.ConstructorTable["network.QNetworkProxyQuery"] = NewQNetworkProxyQueryFromPointer
	internal.ConstructorTable["network.QNetworkReply"] = NewQNetworkReplyFromPointer
	internal.ConstructorTable["network.QNetworkRequest"] = NewQNetworkRequestFromPointer
	internal.ConstructorTable["network.QNetworkSession"] = NewQNetworkSessionFromPointer
	internal.ConstructorTable["network.QOcspResponse"] = NewQOcspResponseFromPointer
	internal.ConstructorTable["network.QSsl"] = NewQSslFromPointer
	internal.ConstructorTable["network.QSslCertificate"] = NewQSslCertificateFromPointer
	internal.ConstructorTable["network.QSslCertificateExtension"] = NewQSslCertificateExtensionFromPointer
	internal.ConstructorTable["network.QSslCipher"] = NewQSslCipherFromPointer
	internal.ConstructorTable["network.QSslConfiguration"] = NewQSslConfigurationFromPointer
	internal.ConstructorTable["network.QSslDiffieHellmanParameters"] = NewQSslDiffieHellmanParametersFromPointer
	internal.ConstructorTable["network.QSslEllipticCurve"] = NewQSslEllipticCurveFromPointer
	internal.ConstructorTable["network.QSslError"] = NewQSslErrorFromPointer
	internal.ConstructorTable["network.QSslKey"] = NewQSslKeyFromPointer
	internal.ConstructorTable["network.QSslPreSharedKeyAuthenticator"] = NewQSslPreSharedKeyAuthenticatorFromPointer
	internal.ConstructorTable["network.QSslSocket"] = NewQSslSocketFromPointer
	internal.ConstructorTable["network.QTcpServer"] = NewQTcpServerFromPointer
	internal.ConstructorTable["network.QTcpSocket"] = NewQTcpSocketFromPointer
	internal.ConstructorTable["network.QUdpSocket"] = NewQUdpSocketFromPointer
	internal.ConstructorTable["network.QNetworkCacheMetaData_RawHeader"] = NewQNetworkCacheMetaData_RawHeaderFromPointer
}
