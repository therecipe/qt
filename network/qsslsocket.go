package network

//#include "qsslsocket.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QSslSocket_" + qt.RandomIdentifier())
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
	return NewQSslSocketFromPointer(C.QSslSocket_NewQSslSocket(core.PointerFromQObject(parent)))
}

func (ptr *QSslSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QSslSocket) AddCaCertificate(certificate QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_AddCaCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(PointerFromQSslCertificate(certificate))
}

func (ptr *QSslSocket) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) Close() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Close(ptr.Pointer())
	}
}

func (ptr *QSslSocket) ConnectEncrypted(f func()) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QSslSocket) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQSslSocketEncrypted
func callbackQSslSocketEncrypted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "encrypted").(func())()
}

func (ptr *QSslSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) IgnoreSslErrors() {
	if ptr.Pointer() != nil {
		C.QSslSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QSslSocket) IsEncrypted() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_IsEncrypted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) Mode() QSslSocket__SslMode {
	if ptr.Pointer() != nil {
		return QSslSocket__SslMode(C.QSslSocket_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) ConnectModeChanged(f func(mode QSslSocket__SslMode)) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modeChanged", f)
	}
}

func (ptr *QSslSocket) DisconnectModeChanged() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modeChanged")
	}
}

//export callbackQSslSocketModeChanged
func callbackQSslSocketModeChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "modeChanged").(func(QSslSocket__SslMode))(QSslSocket__SslMode(mode))
}

func (ptr *QSslSocket) PeerVerifyDepth() int {
	if ptr.Pointer() != nil {
		return int(C.QSslSocket_PeerVerifyDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslSocket_PeerVerifyMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslSocket_PeerVerifyName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QSslSocket) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQSslSocketPreSharedKeyAuthenticationRequired
func callbackQSslSocketPreSharedKeyAuthenticationRequired(ptrName *C.char, authenticator unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired").(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
}

func (ptr *QSslSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QSslSocket) SetLocalCertificate(certificate QSslCertificate_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslSocket) SetPeerVerifyDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QSslSocket) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSslSocket) SetPeerVerifyName(hostName string) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyName(ptr.Pointer(), C.CString(hostName))
	}
}

func (ptr *QSslSocket) SetPrivateKey(key QSslKey_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

func (ptr *QSslSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetSocketOption(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QSslSocket) SetSslConfiguration(configuration QSslConfiguration_ITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

func (ptr *QSslSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSslSocket_SocketOption(ptr.Pointer(), C.int(option)))
	}
	return nil
}

func QSslSocket_SslLibraryBuildVersionString() string {
	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryBuildVersionString())
}

func QSslSocket_SslLibraryVersionString() string {
	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryVersionString())
}

func (ptr *QSslSocket) StartClientEncryption() {
	if ptr.Pointer() != nil {
		C.QSslSocket_StartClientEncryption(ptr.Pointer())
	}
}

func (ptr *QSslSocket) StartServerEncryption() {
	if ptr.Pointer() != nil {
		C.QSslSocket_StartServerEncryption(ptr.Pointer())
	}
}

func QSslSocket_SupportsSsl() bool {
	return C.QSslSocket_QSslSocket_SupportsSsl() != 0
}

func (ptr *QSslSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForConnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForDisconnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForEncrypted(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForEncrypted(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) DestroyQSslSocket() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DestroyQSslSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
