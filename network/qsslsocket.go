package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

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

	return NewQSslSocketFromPointer(C.QSslSocket_NewQSslSocket(core.PointerFromQObject(parent)))
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
func callbackQSslSocketClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QSslSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQSslSocketEncrypted(ptrName *C.char) {
	defer qt.Recovering("callback QSslSocket::encrypted")

	if signal := qt.GetSignal(C.GoString(ptrName), "encrypted"); signal != nil {
		signal.(func())()
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
func callbackQSslSocketEncryptedBytesWritten(ptrName *C.char, written C.longlong) {
	defer qt.Recovering("callback QSslSocket::encryptedBytesWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "encryptedBytesWritten"); signal != nil {
		signal.(func(int64))(int64(written))
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
func callbackQSslSocketModeChanged(ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QSslSocket::modeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "modeChanged"); signal != nil {
		signal.(func(QSslSocket__SslMode))(QSslSocket__SslMode(mode))
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
func callbackQSslSocketPreSharedKeyAuthenticationRequired(ptrName *C.char, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QSslSocket::preSharedKeyAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
	}

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
func callbackQSslSocketResume(ptrName *C.char) bool {
	defer qt.Recovering("callback QSslSocket::resume")

	if signal := qt.GetSignal(C.GoString(ptrName), "resume"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQSslSocketSetReadBufferSize(ptrName *C.char, size C.longlong) bool {
	defer qt.Recovering("callback QSslSocket::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
		return true
	}
	return false

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
func callbackQSslSocketSetSocketOption(ptrName *C.char, option C.int, value unsafe.Pointer) bool {
	defer qt.Recovering("callback QSslSocket::setSocketOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSocketOption"); signal != nil {
		signal.(func(QAbstractSocket__SocketOption, *core.QVariant))(QAbstractSocket__SocketOption(option), core.NewQVariantFromPointer(value))
		return true
	}
	return false

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
func callbackQSslSocketDisconnectFromHost(ptrName *C.char) bool {
	defer qt.Recovering("callback QSslSocket::disconnectFromHost")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectFromHost"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQSslSocketTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSslSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSslSocketChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSslSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSslSocketCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSslSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
