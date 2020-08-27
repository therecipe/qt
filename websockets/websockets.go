// +build !minimal

package websockets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/internal"
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

func (ptr *QMaskGenerator) QMaskGenerator_PTR() *QMaskGenerator {
	return ptr
}

func (ptr *QMaskGenerator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QMaskGenerator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQMaskGenerator(ptr QMaskGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMaskGenerator_PTR().Pointer()
	}
	return nil
}

func (n *QMaskGenerator) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QMaskGenerator) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQMaskGeneratorFromPointer(ptr unsafe.Pointer) (n *QMaskGenerator) {
	n = new(QMaskGenerator)
	n.InitFromInternal(uintptr(ptr), "websockets.QMaskGenerator")
	return
}
func NewQMaskGenerator2(parent core.QObject_ITF) *QMaskGenerator {

	return internal.CallLocalFunction([]interface{}{"", "", "websockets.NewQMaskGenerator2", "", parent}).(*QMaskGenerator)
}

func (ptr *QMaskGenerator) ConnectNextMask(f func() uint) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNextMask", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMaskGenerator) DisconnectNextMask() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNextMask"})
}

func (ptr *QMaskGenerator) NextMask() uint {

	return uint(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextMask"}).(float64))
}

func (ptr *QMaskGenerator) ConnectSeed(f func() bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSeed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMaskGenerator) DisconnectSeed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSeed"})
}

func (ptr *QMaskGenerator) Seed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Seed"}).(bool)
}

func (ptr *QMaskGenerator) ConnectDestroyQMaskGenerator(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQMaskGenerator", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QMaskGenerator) DisconnectDestroyQMaskGenerator() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQMaskGenerator"})
}

func (ptr *QMaskGenerator) DestroyQMaskGenerator() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQMaskGenerator"})
}

func (ptr *QMaskGenerator) DestroyQMaskGeneratorDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQMaskGeneratorDefault"})
}

func (ptr *QMaskGenerator) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QMaskGenerator) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QMaskGenerator) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QMaskGenerator) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QMaskGenerator) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QMaskGenerator) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QMaskGenerator) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QMaskGenerator) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QMaskGenerator) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QMaskGenerator) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QMaskGenerator) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QMaskGenerator) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QMaskGenerator) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QMaskGenerator) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QMaskGenerator) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QMaskGenerator) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QMaskGenerator) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QMaskGenerator) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QMaskGenerator) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QMaskGenerator) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QMaskGenerator) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QQmlWebSocket struct {
	core.QObject
}

type QQmlWebSocket_ITF interface {
	core.QObject_ITF
	QQmlWebSocket_PTR() *QQmlWebSocket
}

func (ptr *QQmlWebSocket) QQmlWebSocket_PTR() *QQmlWebSocket {
	return ptr
}

func (ptr *QQmlWebSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlWebSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlWebSocket(ptr QQmlWebSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlWebSocket_PTR().Pointer()
	}
	return nil
}

func (n *QQmlWebSocket) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlWebSocket) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlWebSocketFromPointer(ptr unsafe.Pointer) (n *QQmlWebSocket) {
	n = new(QQmlWebSocket)
	n.InitFromInternal(uintptr(ptr), "websockets.QQmlWebSocket")
	return
}

func (ptr *QQmlWebSocket) DestroyQQmlWebSocket() {
}

type QQmlWebSocketServer struct {
	core.QObject
}

type QQmlWebSocketServer_ITF interface {
	core.QObject_ITF
	QQmlWebSocketServer_PTR() *QQmlWebSocketServer
}

func (ptr *QQmlWebSocketServer) QQmlWebSocketServer_PTR() *QQmlWebSocketServer {
	return ptr
}

func (ptr *QQmlWebSocketServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QQmlWebSocketServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQQmlWebSocketServer(ptr QQmlWebSocketServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlWebSocketServer_PTR().Pointer()
	}
	return nil
}

func (n *QQmlWebSocketServer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QQmlWebSocketServer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQQmlWebSocketServerFromPointer(ptr unsafe.Pointer) (n *QQmlWebSocketServer) {
	n = new(QQmlWebSocketServer)
	n.InitFromInternal(uintptr(ptr), "websockets.QQmlWebSocketServer")
	return
}

func (ptr *QQmlWebSocketServer) DestroyQQmlWebSocketServer() {
}

type QWebSocket struct {
	core.QObject
}

type QWebSocket_ITF interface {
	core.QObject_ITF
	QWebSocket_PTR() *QWebSocket
}

func (ptr *QWebSocket) QWebSocket_PTR() *QWebSocket {
	return ptr
}

func (ptr *QWebSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebSocket(ptr QWebSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocket_PTR().Pointer()
	}
	return nil
}

func (n *QWebSocket) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebSocket) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebSocketFromPointer(ptr unsafe.Pointer) (n *QWebSocket) {
	n = new(QWebSocket)
	n.InitFromInternal(uintptr(ptr), "websockets.QWebSocket")
	return
}
func NewQWebSocket2(origin string, version QWebSocketProtocol__Version, parent core.QObject_ITF) *QWebSocket {

	return internal.CallLocalFunction([]interface{}{"", "", "websockets.NewQWebSocket2", "", origin, version, parent}).(*QWebSocket)
}

func (ptr *QWebSocket) Abort() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Abort"})
}

func (ptr *QWebSocket) ConnectAboutToClose(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAboutToClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectAboutToClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAboutToClose"})
}

func (ptr *QWebSocket) AboutToClose() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AboutToClose"})
}

func (ptr *QWebSocket) ConnectBinaryFrameReceived(f func(frame *core.QByteArray, isLastFrame bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBinaryFrameReceived", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectBinaryFrameReceived() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBinaryFrameReceived"})
}

func (ptr *QWebSocket) BinaryFrameReceived(frame core.QByteArray_ITF, isLastFrame bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BinaryFrameReceived", frame, isLastFrame})
}

func (ptr *QWebSocket) ConnectBinaryMessageReceived(f func(message *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBinaryMessageReceived", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectBinaryMessageReceived() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBinaryMessageReceived"})
}

func (ptr *QWebSocket) BinaryMessageReceived(message core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BinaryMessageReceived", message})
}

func (ptr *QWebSocket) BytesToWrite() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesToWrite"}).(float64))
}

func (ptr *QWebSocket) ConnectBytesWritten(f func(bytes int64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectBytesWritten", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectBytesWritten() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectBytesWritten"})
}

func (ptr *QWebSocket) BytesWritten(bytes int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "BytesWritten", bytes})
}

func (ptr *QWebSocket) ConnectClose(f func(closeCode QWebSocketProtocol__CloseCode, reason string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClose", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectClose() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClose"})
}

func (ptr *QWebSocket) Close(closeCode QWebSocketProtocol__CloseCode, reason string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close", closeCode, reason})
}

func (ptr *QWebSocket) CloseDefault(closeCode QWebSocketProtocol__CloseCode, reason string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseDefault", closeCode, reason})
}

func (ptr *QWebSocket) CloseCode() QWebSocketProtocol__CloseCode {

	return QWebSocketProtocol__CloseCode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseCode"}).(float64))
}

func (ptr *QWebSocket) CloseReason() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CloseReason"}).(string)
}

func (ptr *QWebSocket) ConnectConnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectConnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectConnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectConnected"})
}

func (ptr *QWebSocket) Connected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Connected"})
}

func (ptr *QWebSocket) ConnectDisconnected(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDisconnected", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectDisconnected() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDisconnected"})
}

func (ptr *QWebSocket) Disconnected() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Disconnected"})
}

func (ptr *QWebSocket) Error() network.QAbstractSocket__SocketError {

	return network.QAbstractSocket__SocketError(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QWebSocket) ConnectError2(f func(error network.QAbstractSocket__SocketError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectError2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectError2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectError2"})
}

func (ptr *QWebSocket) Error2(error network.QAbstractSocket__SocketError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error2", error})
}

func (ptr *QWebSocket) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QWebSocket) Flush() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Flush"}).(bool)
}

func (ptr *QWebSocket) ConnectIgnoreSslErrors(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectIgnoreSslErrors", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectIgnoreSslErrors() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectIgnoreSslErrors"})
}

func (ptr *QWebSocket) IgnoreSslErrors() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrors"})
}

func (ptr *QWebSocket) IgnoreSslErrorsDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrorsDefault"})
}

func (ptr *QWebSocket) IgnoreSslErrors2(errors []*network.QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IgnoreSslErrors2", errors})
}

func (ptr *QWebSocket) IsValid() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsValid"}).(bool)
}

func (ptr *QWebSocket) LocalAddress() *network.QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalAddress"}).(*network.QHostAddress)
}

func (ptr *QWebSocket) LocalPort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocalPort"}).(float64))
}

func (ptr *QWebSocket) MaskGenerator() *QMaskGenerator {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaskGenerator"}).(*QMaskGenerator)
}

func (ptr *QWebSocket) ConnectOpen(f func(url *core.QUrl)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectOpen() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen"})
}

func (ptr *QWebSocket) Open(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open", url})
}

func (ptr *QWebSocket) OpenDefault(url core.QUrl_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OpenDefault", url})
}

func (ptr *QWebSocket) ConnectOpen2(f func(request *network.QNetworkRequest)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOpen2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectOpen2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOpen2"})
}

func (ptr *QWebSocket) Open2(request network.QNetworkRequest_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open2", request})
}

func (ptr *QWebSocket) Open2Default(request network.QNetworkRequest_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Open2Default", request})
}

func (ptr *QWebSocket) Origin() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Origin"}).(string)
}

func (ptr *QWebSocket) PauseMode() network.QAbstractSocket__PauseMode {

	return network.QAbstractSocket__PauseMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PauseMode"}).(float64))
}

func (ptr *QWebSocket) PeerAddress() *network.QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerAddress"}).(*network.QHostAddress)
}

func (ptr *QWebSocket) PeerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerName"}).(string)
}

func (ptr *QWebSocket) PeerPort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerPort"}).(float64))
}

func (ptr *QWebSocket) ConnectPing(f func(payload *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPing", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectPing() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPing"})
}

func (ptr *QWebSocket) Ping(payload core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Ping", payload})
}

func (ptr *QWebSocket) PingDefault(payload core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PingDefault", payload})
}

func (ptr *QWebSocket) ConnectPong(f func(elapsedTime uint64, payload *core.QByteArray)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPong", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectPong() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPong"})
}

func (ptr *QWebSocket) Pong(elapsedTime uint64, payload core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pong", elapsedTime, payload})
}

func (ptr *QWebSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *network.QSslPreSharedKeyAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreSharedKeyAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectPreSharedKeyAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreSharedKeyAuthenticationRequired"})
}

func (ptr *QWebSocket) PreSharedKeyAuthenticationRequired(authenticator network.QSslPreSharedKeyAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreSharedKeyAuthenticationRequired", authenticator})
}

func (ptr *QWebSocket) Proxy() *network.QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Proxy"}).(*network.QNetworkProxy)
}

func (ptr *QWebSocket) ConnectProxyAuthenticationRequired(f func(proxy *network.QNetworkProxy, authenticator *network.QAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectProxyAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectProxyAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectProxyAuthenticationRequired"})
}

func (ptr *QWebSocket) ProxyAuthenticationRequired(proxy network.QNetworkProxy_ITF, authenticator network.QAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ProxyAuthenticationRequired", proxy, authenticator})
}

func (ptr *QWebSocket) ReadBufferSize() int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadBufferSize"}).(float64))
}

func (ptr *QWebSocket) ConnectReadChannelFinished(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectReadChannelFinished", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectReadChannelFinished() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectReadChannelFinished"})
}

func (ptr *QWebSocket) ReadChannelFinished() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ReadChannelFinished"})
}

func (ptr *QWebSocket) Request() *network.QNetworkRequest {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Request"}).(*network.QNetworkRequest)
}

func (ptr *QWebSocket) RequestUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RequestUrl"}).(*core.QUrl)
}

func (ptr *QWebSocket) ResourceName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResourceName"}).(string)
}

func (ptr *QWebSocket) Resume() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Resume"})
}

func (ptr *QWebSocket) SendBinaryMessage(data core.QByteArray_ITF) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendBinaryMessage", data}).(float64))
}

func (ptr *QWebSocket) SendTextMessage(message string) int64 {

	return int64(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SendTextMessage", message}).(float64))
}

func (ptr *QWebSocket) SetMaskGenerator(maskGenerator QMaskGenerator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaskGenerator", maskGenerator})
}

func (ptr *QWebSocket) SetPauseMode(pauseMode network.QAbstractSocket__PauseMode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPauseMode", pauseMode})
}

func (ptr *QWebSocket) SetProxy(networkProxy network.QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProxy", networkProxy})
}

func (ptr *QWebSocket) SetReadBufferSize(size int64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetReadBufferSize", size})
}

func (ptr *QWebSocket) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSslConfiguration", sslConfiguration})
}

func (ptr *QWebSocket) SslConfiguration() *network.QSslConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslConfiguration"}).(*network.QSslConfiguration)
}

func (ptr *QWebSocket) ConnectSslErrors(f func(errors []*network.QSslError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSslErrors", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectSslErrors() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSslErrors"})
}

func (ptr *QWebSocket) SslErrors(errors []*network.QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslErrors", errors})
}

func (ptr *QWebSocket) State() network.QAbstractSocket__SocketState {

	return network.QAbstractSocket__SocketState(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QWebSocket) ConnectStateChanged(f func(state network.QAbstractSocket__SocketState)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QWebSocket) StateChanged(state network.QAbstractSocket__SocketState) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", state})
}

func (ptr *QWebSocket) ConnectTextFrameReceived(f func(frame string, isLastFrame bool)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextFrameReceived", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectTextFrameReceived() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextFrameReceived"})
}

func (ptr *QWebSocket) TextFrameReceived(frame string, isLastFrame bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextFrameReceived", frame, isLastFrame})
}

func (ptr *QWebSocket) ConnectTextMessageReceived(f func(message string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectTextMessageReceived", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectTextMessageReceived() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectTextMessageReceived"})
}

func (ptr *QWebSocket) TextMessageReceived(message string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TextMessageReceived", message})
}

func (ptr *QWebSocket) Version() QWebSocketProtocol__Version {

	return QWebSocketProtocol__Version(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Version"}).(float64))
}

func (ptr *QWebSocket) ConnectDestroyQWebSocket(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQWebSocket", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocket) DisconnectDestroyQWebSocket() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQWebSocket"})
}

func (ptr *QWebSocket) DestroyQWebSocket() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebSocket"})
}

func (ptr *QWebSocket) DestroyQWebSocketDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebSocketDefault"})
}

func (ptr *QWebSocket) __ignoreSslErrors_errors_atList2(i int) *network.QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrors_errors_atList2", i}).(*network.QSslError)
}

func (ptr *QWebSocket) __ignoreSslErrors_errors_setList2(i network.QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrors_errors_setList2", i})
}

func (ptr *QWebSocket) __ignoreSslErrors_errors_newList2() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__ignoreSslErrors_errors_newList2"}).(unsafe.Pointer)
}

func (ptr *QWebSocket) __sslErrors_errors_atList(i int) *network.QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_atList", i}).(*network.QSslError)
}

func (ptr *QWebSocket) __sslErrors_errors_setList(i network.QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_setList", i})
}

func (ptr *QWebSocket) __sslErrors_errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_newList"}).(unsafe.Pointer)
}

func (ptr *QWebSocket) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebSocket) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebSocket) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebSocket) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebSocket) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebSocket) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebSocket) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebSocket) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebSocket) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebSocket) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebSocket) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebSocket) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebSocket) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebSocket) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebSocket) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebSocket) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebSocket) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QWebSocketCorsAuthenticator struct {
	internal.Internal
}

type QWebSocketCorsAuthenticator_ITF interface {
	QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator
}

func (ptr *QWebSocketCorsAuthenticator) QWebSocketCorsAuthenticator_PTR() *QWebSocketCorsAuthenticator {
	return ptr
}

func (ptr *QWebSocketCorsAuthenticator) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebSocketCorsAuthenticator) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebSocketCorsAuthenticator(ptr QWebSocketCorsAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketCorsAuthenticator_PTR().Pointer()
	}
	return nil
}

func (n *QWebSocketCorsAuthenticator) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebSocketCorsAuthenticatorFromPointer(ptr unsafe.Pointer) (n *QWebSocketCorsAuthenticator) {
	n = new(QWebSocketCorsAuthenticator)
	n.InitFromInternal(uintptr(ptr), "websockets.QWebSocketCorsAuthenticator")
	return
}
func NewQWebSocketCorsAuthenticator(origin string) *QWebSocketCorsAuthenticator {

	return internal.CallLocalFunction([]interface{}{"", "", "websockets.NewQWebSocketCorsAuthenticator", "", origin}).(*QWebSocketCorsAuthenticator)
}

func NewQWebSocketCorsAuthenticator2(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {

	return internal.CallLocalFunction([]interface{}{"", "", "websockets.NewQWebSocketCorsAuthenticator2", "", other}).(*QWebSocketCorsAuthenticator)
}

func NewQWebSocketCorsAuthenticator3(other QWebSocketCorsAuthenticator_ITF) *QWebSocketCorsAuthenticator {

	return internal.CallLocalFunction([]interface{}{"", "", "websockets.NewQWebSocketCorsAuthenticator3", "", other}).(*QWebSocketCorsAuthenticator)
}

func (ptr *QWebSocketCorsAuthenticator) Allowed() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Allowed"}).(bool)
}

func (ptr *QWebSocketCorsAuthenticator) Origin() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Origin"}).(string)
}

func (ptr *QWebSocketCorsAuthenticator) SetAllowed(allowed bool) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetAllowed", allowed})
}

func (ptr *QWebSocketCorsAuthenticator) Swap(other QWebSocketCorsAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Swap", other})
}

func (ptr *QWebSocketCorsAuthenticator) DestroyQWebSocketCorsAuthenticator() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebSocketCorsAuthenticator"})
}

type QWebSocketProtocol struct {
	internal.Internal
}

type QWebSocketProtocol_ITF interface {
	QWebSocketProtocol_PTR() *QWebSocketProtocol
}

func (ptr *QWebSocketProtocol) QWebSocketProtocol_PTR() *QWebSocketProtocol {
	return ptr
}

func (ptr *QWebSocketProtocol) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QWebSocketProtocol) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQWebSocketProtocol(ptr QWebSocketProtocol_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketProtocol_PTR().Pointer()
	}
	return nil
}

func (n *QWebSocketProtocol) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQWebSocketProtocolFromPointer(ptr unsafe.Pointer) (n *QWebSocketProtocol) {
	n = new(QWebSocketProtocol)
	n.InitFromInternal(uintptr(ptr), "websockets.QWebSocketProtocol")
	return
}

func (ptr *QWebSocketProtocol) DestroyQWebSocketProtocol() {
}

//go:generate stringer -type=QWebSocketProtocol__CloseCode
//QWebSocketProtocol::CloseCode
type QWebSocketProtocol__CloseCode int64

const (
	QWebSocketProtocol__CloseCodeNormal                QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1000)
	QWebSocketProtocol__CloseCodeGoingAway             QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1001)
	QWebSocketProtocol__CloseCodeProtocolError         QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1002)
	QWebSocketProtocol__CloseCodeDatatypeNotSupported  QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1003)
	QWebSocketProtocol__CloseCodeReserved1004          QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1004)
	QWebSocketProtocol__CloseCodeMissingStatusCode     QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1005)
	QWebSocketProtocol__CloseCodeAbnormalDisconnection QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1006)
	QWebSocketProtocol__CloseCodeWrongDatatype         QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1007)
	QWebSocketProtocol__CloseCodePolicyViolated        QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1008)
	QWebSocketProtocol__CloseCodeTooMuchData           QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1009)
	QWebSocketProtocol__CloseCodeMissingExtension      QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1010)
	QWebSocketProtocol__CloseCodeBadOperation          QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1011)
	QWebSocketProtocol__CloseCodeTlsHandshakeFailed    QWebSocketProtocol__CloseCode = QWebSocketProtocol__CloseCode(1015)
)

//go:generate stringer -type=QWebSocketProtocol__Version
//QWebSocketProtocol::Version
type QWebSocketProtocol__Version int64

const (
	QWebSocketProtocol__VersionUnknown QWebSocketProtocol__Version = QWebSocketProtocol__Version(-1)
	QWebSocketProtocol__Version0       QWebSocketProtocol__Version = QWebSocketProtocol__Version(0)
	QWebSocketProtocol__Version4       QWebSocketProtocol__Version = QWebSocketProtocol__Version(4)
	QWebSocketProtocol__Version5       QWebSocketProtocol__Version = QWebSocketProtocol__Version(5)
	QWebSocketProtocol__Version6       QWebSocketProtocol__Version = QWebSocketProtocol__Version(6)
	QWebSocketProtocol__Version7       QWebSocketProtocol__Version = QWebSocketProtocol__Version(7)
	QWebSocketProtocol__Version8       QWebSocketProtocol__Version = QWebSocketProtocol__Version(8)
	QWebSocketProtocol__Version13      QWebSocketProtocol__Version = QWebSocketProtocol__Version(13)
	QWebSocketProtocol__VersionLatest  QWebSocketProtocol__Version = QWebSocketProtocol__Version(QWebSocketProtocol__Version13)
)

type QWebSocketServer struct {
	core.QObject
}

type QWebSocketServer_ITF interface {
	core.QObject_ITF
	QWebSocketServer_PTR() *QWebSocketServer
}

func (ptr *QWebSocketServer) QWebSocketServer_PTR() *QWebSocketServer {
	return ptr
}

func (ptr *QWebSocketServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QWebSocketServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQWebSocketServer(ptr QWebSocketServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebSocketServer_PTR().Pointer()
	}
	return nil
}

func (n *QWebSocketServer) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QWebSocketServer) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQWebSocketServerFromPointer(ptr unsafe.Pointer) (n *QWebSocketServer) {
	n = new(QWebSocketServer)
	n.InitFromInternal(uintptr(ptr), "websockets.QWebSocketServer")
	return
}

//go:generate stringer -type=QWebSocketServer__SslMode
//QWebSocketServer::SslMode
type QWebSocketServer__SslMode int64

const (
	QWebSocketServer__SecureMode    QWebSocketServer__SslMode = QWebSocketServer__SslMode(0)
	QWebSocketServer__NonSecureMode QWebSocketServer__SslMode = QWebSocketServer__SslMode(1)
)

func NewQWebSocketServer2(serverName string, secureMode QWebSocketServer__SslMode, parent core.QObject_ITF) *QWebSocketServer {

	return internal.CallLocalFunction([]interface{}{"", "", "websockets.NewQWebSocketServer2", "", serverName, secureMode, parent}).(*QWebSocketServer)
}

func (ptr *QWebSocketServer) ConnectAcceptError(f func(socketError network.QAbstractSocket__SocketError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAcceptError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectAcceptError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAcceptError"})
}

func (ptr *QWebSocketServer) AcceptError(socketError network.QAbstractSocket__SocketError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AcceptError", socketError})
}

func (ptr *QWebSocketServer) Close() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Close"})
}

func (ptr *QWebSocketServer) ConnectClosed(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectClosed", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectClosed() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectClosed"})
}

func (ptr *QWebSocketServer) Closed() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Closed"})
}

func (ptr *QWebSocketServer) Error() QWebSocketProtocol__CloseCode {

	return QWebSocketProtocol__CloseCode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Error"}).(float64))
}

func (ptr *QWebSocketServer) ErrorString() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ErrorString"}).(string)
}

func (ptr *QWebSocketServer) HandleConnection(socket network.QTcpSocket_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HandleConnection", socket})
}

func (ptr *QWebSocketServer) HasPendingConnections() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "HasPendingConnections"}).(bool)
}

func (ptr *QWebSocketServer) IsListening() bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "IsListening"}).(bool)
}

func (ptr *QWebSocketServer) Listen(address network.QHostAddress_ITF, port uint16) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Listen", address, port}).(bool)
}

func (ptr *QWebSocketServer) MaxPendingConnections() int {

	return int(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MaxPendingConnections"}).(float64))
}

func (ptr *QWebSocketServer) ConnectNewConnection(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNewConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectNewConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNewConnection"})
}

func (ptr *QWebSocketServer) NewConnection() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NewConnection"})
}

func (ptr *QWebSocketServer) ConnectNextPendingConnection(f func() *QWebSocket) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNextPendingConnection", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectNextPendingConnection() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNextPendingConnection"})
}

func (ptr *QWebSocketServer) NextPendingConnection() *QWebSocket {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextPendingConnection"}).(*QWebSocket)
}

func (ptr *QWebSocketServer) NextPendingConnectionDefault() *QWebSocket {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "NextPendingConnectionDefault"}).(*QWebSocket)
}

func (ptr *QWebSocketServer) ConnectOriginAuthenticationRequired(f func(authenticator *QWebSocketCorsAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectOriginAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectOriginAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectOriginAuthenticationRequired"})
}

func (ptr *QWebSocketServer) OriginAuthenticationRequired(authenticator QWebSocketCorsAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "OriginAuthenticationRequired", authenticator})
}

func (ptr *QWebSocketServer) PauseAccepting() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PauseAccepting"})
}

func (ptr *QWebSocketServer) ConnectPeerVerifyError(f func(error *network.QSslError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPeerVerifyError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectPeerVerifyError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPeerVerifyError"})
}

func (ptr *QWebSocketServer) PeerVerifyError(error network.QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PeerVerifyError", error})
}

func (ptr *QWebSocketServer) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *network.QSslPreSharedKeyAuthenticator)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPreSharedKeyAuthenticationRequired", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectPreSharedKeyAuthenticationRequired() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPreSharedKeyAuthenticationRequired"})
}

func (ptr *QWebSocketServer) PreSharedKeyAuthenticationRequired(authenticator network.QSslPreSharedKeyAuthenticator_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PreSharedKeyAuthenticationRequired", authenticator})
}

func (ptr *QWebSocketServer) Proxy() *network.QNetworkProxy {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Proxy"}).(*network.QNetworkProxy)
}

func (ptr *QWebSocketServer) ResumeAccepting() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResumeAccepting"})
}

func (ptr *QWebSocketServer) SecureMode() QWebSocketServer__SslMode {

	return QWebSocketServer__SslMode(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SecureMode"}).(float64))
}

func (ptr *QWebSocketServer) ServerAddress() *network.QHostAddress {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerAddress"}).(*network.QHostAddress)
}

func (ptr *QWebSocketServer) ConnectServerError(f func(closeCode QWebSocketProtocol__CloseCode)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectServerError", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectServerError() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectServerError"})
}

func (ptr *QWebSocketServer) ServerError(closeCode QWebSocketProtocol__CloseCode) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerError", closeCode})
}

func (ptr *QWebSocketServer) ServerName() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerName"}).(string)
}

func (ptr *QWebSocketServer) ServerPort() uint16 {

	return uint16(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerPort"}).(float64))
}

func (ptr *QWebSocketServer) ServerUrl() *core.QUrl {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ServerUrl"}).(*core.QUrl)
}

func (ptr *QWebSocketServer) SetMaxPendingConnections(numConnections int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetMaxPendingConnections", numConnections})
}

func (ptr *QWebSocketServer) SetProxy(networkProxy network.QNetworkProxy_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetProxy", networkProxy})
}

func (ptr *QWebSocketServer) SetServerName(serverName string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetServerName", serverName})
}

func (ptr *QWebSocketServer) SetSslConfiguration(sslConfiguration network.QSslConfiguration_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetSslConfiguration", sslConfiguration})
}

func (ptr *QWebSocketServer) SslConfiguration() *network.QSslConfiguration {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslConfiguration"}).(*network.QSslConfiguration)
}

func (ptr *QWebSocketServer) ConnectSslErrors(f func(errors []*network.QSslError)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSslErrors", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectSslErrors() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSslErrors"})
}

func (ptr *QWebSocketServer) SslErrors(errors []*network.QSslError) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SslErrors", errors})
}

func (ptr *QWebSocketServer) SupportedVersions() []QWebSocketProtocol__Version {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SupportedVersions"}).([]QWebSocketProtocol__Version)
}

func (ptr *QWebSocketServer) ConnectDestroyQWebSocketServer(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectDestroyQWebSocketServer", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QWebSocketServer) DisconnectDestroyQWebSocketServer() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectDestroyQWebSocketServer"})
}

func (ptr *QWebSocketServer) DestroyQWebSocketServer() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebSocketServer"})
}

func (ptr *QWebSocketServer) DestroyQWebSocketServerDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DestroyQWebSocketServerDefault"})
}

func (ptr *QWebSocketServer) __sslErrors_errors_atList(i int) *network.QSslError {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_atList", i}).(*network.QSslError)
}

func (ptr *QWebSocketServer) __sslErrors_errors_setList(i network.QSslError_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_setList", i})
}

func (ptr *QWebSocketServer) __sslErrors_errors_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__sslErrors_errors_newList"}).(unsafe.Pointer)
}

func (ptr *QWebSocketServer) __supportedVersions_atList(i int) QWebSocketProtocol__Version {

	return QWebSocketProtocol__Version(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedVersions_atList", i}).(float64))
}

func (ptr *QWebSocketServer) __supportedVersions_setList(i QWebSocketProtocol__Version) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedVersions_setList", i})
}

func (ptr *QWebSocketServer) __supportedVersions_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__supportedVersions_newList"}).(unsafe.Pointer)
}

func (ptr *QWebSocketServer) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QWebSocketServer) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QWebSocketServer) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QWebSocketServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QWebSocketServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QWebSocketServer) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QWebSocketServer) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QWebSocketServer) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QWebSocketServer) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QWebSocketServer) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QWebSocketServer) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QWebSocketServer) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QWebSocketServer) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QWebSocketServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QWebSocketServer) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QWebSocketServer) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QWebSocketServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QWebSocketServer) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QWebSocketServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QWebSocketServer) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QWebSocketServer) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QtWebSocketsDeclarativeModule struct {
	internal.Internal
}

type QtWebSocketsDeclarativeModule_ITF interface {
	QtWebSocketsDeclarativeModule_PTR() *QtWebSocketsDeclarativeModule
}

func (ptr *QtWebSocketsDeclarativeModule) QtWebSocketsDeclarativeModule_PTR() *QtWebSocketsDeclarativeModule {
	return ptr
}

func (ptr *QtWebSocketsDeclarativeModule) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QtWebSocketsDeclarativeModule) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQtWebSocketsDeclarativeModule(ptr QtWebSocketsDeclarativeModule_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtWebSocketsDeclarativeModule_PTR().Pointer()
	}
	return nil
}

func (n *QtWebSocketsDeclarativeModule) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQtWebSocketsDeclarativeModuleFromPointer(ptr unsafe.Pointer) (n *QtWebSocketsDeclarativeModule) {
	n = new(QtWebSocketsDeclarativeModule)
	n.InitFromInternal(uintptr(ptr), "websockets.QtWebSocketsDeclarativeModule")
	return
}

func (ptr *QtWebSocketsDeclarativeModule) DestroyQtWebSocketsDeclarativeModule() {
}

func init() {
	internal.ConstructorTable["websockets.QMaskGenerator"] = NewQMaskGeneratorFromPointer
	internal.ConstructorTable["websockets.QWebSocket"] = NewQWebSocketFromPointer
	internal.ConstructorTable["websockets.QWebSocketCorsAuthenticator"] = NewQWebSocketCorsAuthenticatorFromPointer
	internal.ConstructorTable["websockets.QWebSocketProtocol"] = NewQWebSocketProtocolFromPointer
	internal.ConstructorTable["websockets.QWebSocketServer"] = NewQWebSocketServerFromPointer
	internal.ConstructorTable["websockets.QtWebSocketsDeclarativeModule"] = NewQtWebSocketsDeclarativeModuleFromPointer
}
