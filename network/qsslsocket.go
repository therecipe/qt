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

type QSslSocketITF interface {
	QTcpSocketITF
	QSslSocketPTR() *QSslSocket
}

func PointerFromQSslSocket(ptr QSslSocketITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslSocketPTR().Pointer()
	}
	return nil
}

func QSslSocketFromPointer(ptr unsafe.Pointer) *QSslSocket {
	var n = new(QSslSocket)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSslSocket_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSslSocket) QSslSocketPTR() *QSslSocket {
	return ptr
}

//QSslSocket::PeerVerifyMode
type QSslSocket__PeerVerifyMode int

var (
	QSslSocket__VerifyNone     = QSslSocket__PeerVerifyMode(0)
	QSslSocket__QueryPeer      = QSslSocket__PeerVerifyMode(1)
	QSslSocket__VerifyPeer     = QSslSocket__PeerVerifyMode(2)
	QSslSocket__AutoVerifyPeer = QSslSocket__PeerVerifyMode(3)
)

//QSslSocket::SslMode
type QSslSocket__SslMode int

var (
	QSslSocket__UnencryptedMode = QSslSocket__SslMode(0)
	QSslSocket__SslClientMode   = QSslSocket__SslMode(1)
	QSslSocket__SslServerMode   = QSslSocket__SslMode(2)
)

func NewQSslSocket(parent core.QObjectITF) *QSslSocket {
	return QSslSocketFromPointer(unsafe.Pointer(C.QSslSocket_NewQSslSocket(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QSslSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Abort(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSslSocket) AddCaCertificate(certificate QSslCertificateITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_AddCaCertificate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslCertificate(certificate)))
	}
}

func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificateITF) {
	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(C.QtObjectPtr(PointerFromQSslCertificate(certificate)))
}

func (ptr *QSslSocket) AtEnd() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_AtEnd(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslSocket) CanReadLine() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_CanReadLine(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslSocket) Close() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSslSocket) ConnectEncrypted(f func()) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectEncrypted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QSslSocket) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncrypted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQSslSocketEncrypted
func callbackQSslSocketEncrypted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "encrypted").(func())()
}

func (ptr *QSslSocket) Flush() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_Flush(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslSocket) IgnoreSslErrors() {
	if ptr.Pointer() != nil {
		C.QSslSocket_IgnoreSslErrors(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSslSocket) IsEncrypted() bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_IsEncrypted(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslSocket) Mode() QSslSocket__SslMode {
	if ptr.Pointer() != nil {
		return QSslSocket__SslMode(C.QSslSocket_Mode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslSocket) ConnectModeChanged(f func(mode QSslSocket__SslMode)) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "modeChanged", f)
	}
}

func (ptr *QSslSocket) DisconnectModeChanged() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectModeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "modeChanged")
	}
}

//export callbackQSslSocketModeChanged
func callbackQSslSocketModeChanged(ptrName *C.char, mode C.int) {
	qt.GetSignal(C.GoString(ptrName), "modeChanged").(func(QSslSocket__SslMode))(QSslSocket__SslMode(mode))
}

func (ptr *QSslSocket) PeerVerifyDepth() int {
	if ptr.Pointer() != nil {
		return int(C.QSslSocket_PeerVerifyDepth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslSocket_PeerVerifyMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslSocket_PeerVerifyName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator QSslPreSharedKeyAuthenticatorITF)) {
	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectPreSharedKeyAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QSslSocket) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQSslSocketPreSharedKeyAuthenticationRequired
func callbackQSslSocketPreSharedKeyAuthenticationRequired(ptrName *C.char, authenticator unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired").(func(*QSslPreSharedKeyAuthenticator))(QSslPreSharedKeyAuthenticatorFromPointer(authenticator))
}

func (ptr *QSslSocket) Resume() {
	if ptr.Pointer() != nil {
		C.QSslSocket_Resume(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSslSocket) SetLocalCertificate(certificate QSslCertificateITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetLocalCertificate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslCertificate(certificate)))
	}
}

func (ptr *QSslSocket) SetPeerVerifyDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyDepth(C.QtObjectPtr(ptr.Pointer()), C.int(depth))
	}
}

func (ptr *QSslSocket) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QSslSocket) SetPeerVerifyName(hostName string) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyName(C.QtObjectPtr(ptr.Pointer()), C.CString(hostName))
	}
}

func (ptr *QSslSocket) SetPrivateKey(key QSslKeyITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetPrivateKey(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslKey(key)))
	}
}

func (ptr *QSslSocket) SetSocketOption(option QAbstractSocket__SocketOption, value string) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetSocketOption(C.QtObjectPtr(ptr.Pointer()), C.int(option), C.CString(value))
	}
}

func (ptr *QSslSocket) SetSslConfiguration(configuration QSslConfigurationITF) {
	if ptr.Pointer() != nil {
		C.QSslSocket_SetSslConfiguration(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslConfiguration(configuration)))
	}
}

func (ptr *QSslSocket) SocketOption(option QAbstractSocket__SocketOption) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslSocket_SocketOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)))
	}
	return ""
}

func QSslSocket_SslLibraryBuildVersionString() string {
	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryBuildVersionString())
}

func QSslSocket_SslLibraryVersionString() string {
	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryVersionString())
}

func (ptr *QSslSocket) StartClientEncryption() {
	if ptr.Pointer() != nil {
		C.QSslSocket_StartClientEncryption(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSslSocket) StartServerEncryption() {
	if ptr.Pointer() != nil {
		C.QSslSocket_StartServerEncryption(C.QtObjectPtr(ptr.Pointer()))
	}
}

func QSslSocket_SupportsSsl() bool {
	return C.QSslSocket_QSslSocket_SupportsSsl() != 0
}

func (ptr *QSslSocket) WaitForBytesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForBytesWritten(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForConnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForConnected(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForDisconnected(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForDisconnected(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForEncrypted(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForEncrypted(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForReadyRead(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForReadyRead(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) DestroyQSslSocket() {
	if ptr.Pointer() != nil {
		C.QSslSocket_DestroyQSslSocket(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
