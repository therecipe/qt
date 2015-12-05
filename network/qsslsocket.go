package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::QSslSocket")
		}
	}()

	return NewQSslSocketFromPointer(C.QSslSocket_NewQSslSocket(core.PointerFromQObject(parent)))
}

func (ptr *QSslSocket) Abort() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::abort")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_Abort(ptr.Pointer())
	}
}

func (ptr *QSslSocket) AddCaCertificate(certificate QSslCertificate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::addCaCertificate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_AddCaCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func QSslSocket_AddDefaultCaCertificate(certificate QSslCertificate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::addDefaultCaCertificate")
		}
	}()

	C.QSslSocket_QSslSocket_AddDefaultCaCertificate(PointerFromQSslCertificate(certificate))
}

func (ptr *QSslSocket) AtEnd() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::atEnd")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) CanReadLine() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::canReadLine")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) Close() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::close")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_Close(ptr.Pointer())
	}
}

func (ptr *QSslSocket) ConnectEncrypted(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::encrypted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectEncrypted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QSslSocket) DisconnectEncrypted() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::encrypted")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectEncrypted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQSslSocketEncrypted
func callbackQSslSocketEncrypted(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::encrypted")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "encrypted").(func())()
}

func (ptr *QSslSocket) Flush() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::flush")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslSocket_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) IgnoreSslErrors() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::ignoreSslErrors")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_IgnoreSslErrors(ptr.Pointer())
	}
}

func (ptr *QSslSocket) IsEncrypted() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::isEncrypted")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslSocket_IsEncrypted(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslSocket) Mode() QSslSocket__SslMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::mode")
		}
	}()

	if ptr.Pointer() != nil {
		return QSslSocket__SslMode(C.QSslSocket_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) ConnectModeChanged(f func(mode QSslSocket__SslMode)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::modeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectModeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "modeChanged", f)
	}
}

func (ptr *QSslSocket) DisconnectModeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::modeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectModeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "modeChanged")
	}
}

//export callbackQSslSocketModeChanged
func callbackQSslSocketModeChanged(ptrName *C.char, mode C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::modeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "modeChanged").(func(QSslSocket__SslMode))(QSslSocket__SslMode(mode))
}

func (ptr *QSslSocket) PeerVerifyDepth() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::peerVerifyDepth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QSslSocket_PeerVerifyDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyMode() QSslSocket__PeerVerifyMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::peerVerifyMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QSslSocket__PeerVerifyMode(C.QSslSocket_PeerVerifyMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslSocket) PeerVerifyName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::peerVerifyName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QSslSocket_PeerVerifyName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSslSocket) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator *QSslPreSharedKeyAuthenticator)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::preSharedKeyAuthenticationRequired")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_ConnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QSslSocket) DisconnectPreSharedKeyAuthenticationRequired() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::preSharedKeyAuthenticationRequired")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_DisconnectPreSharedKeyAuthenticationRequired(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQSslSocketPreSharedKeyAuthenticationRequired
func callbackQSslSocketPreSharedKeyAuthenticationRequired(ptrName *C.char, authenticator unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::preSharedKeyAuthenticationRequired")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired").(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
}

func (ptr *QSslSocket) Resume() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::resume")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_Resume(ptr.Pointer())
	}
}

func (ptr *QSslSocket) SetLocalCertificate(certificate QSslCertificate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::setLocalCertificate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_SetLocalCertificate(ptr.Pointer(), PointerFromQSslCertificate(certificate))
	}
}

func (ptr *QSslSocket) SetPeerVerifyDepth(depth int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::setPeerVerifyDepth")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QSslSocket) SetPeerVerifyMode(mode QSslSocket__PeerVerifyMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::setPeerVerifyMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSslSocket) SetPeerVerifyName(hostName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::setPeerVerifyName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPeerVerifyName(ptr.Pointer(), C.CString(hostName))
	}
}

func (ptr *QSslSocket) SetPrivateKey(key QSslKey_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::setPrivateKey")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_SetPrivateKey(ptr.Pointer(), PointerFromQSslKey(key))
	}
}

func (ptr *QSslSocket) SetSocketOption(option QAbstractSocket__SocketOption, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::setSocketOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_SetSocketOption(ptr.Pointer(), C.int(option), core.PointerFromQVariant(value))
	}
}

func (ptr *QSslSocket) SetSslConfiguration(configuration QSslConfiguration_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::setSslConfiguration")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_SetSslConfiguration(ptr.Pointer(), PointerFromQSslConfiguration(configuration))
	}
}

func (ptr *QSslSocket) SocketOption(option QAbstractSocket__SocketOption) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::socketOption")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSslSocket_SocketOption(ptr.Pointer(), C.int(option)))
	}
	return nil
}

func QSslSocket_SslLibraryBuildVersionString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::sslLibraryBuildVersionString")
		}
	}()

	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryBuildVersionString())
}

func QSslSocket_SslLibraryVersionString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::sslLibraryVersionString")
		}
	}()

	return C.GoString(C.QSslSocket_QSslSocket_SslLibraryVersionString())
}

func (ptr *QSslSocket) StartClientEncryption() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::startClientEncryption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_StartClientEncryption(ptr.Pointer())
	}
}

func (ptr *QSslSocket) StartServerEncryption() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::startServerEncryption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_StartServerEncryption(ptr.Pointer())
	}
}

func QSslSocket_SupportsSsl() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::supportsSsl")
		}
	}()

	return C.QSslSocket_QSslSocket_SupportsSsl() != 0
}

func (ptr *QSslSocket) WaitForBytesWritten(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::waitForBytesWritten")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForConnected(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::waitForConnected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForConnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForDisconnected(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::waitForDisconnected")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForDisconnected(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForEncrypted(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::waitForEncrypted")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForEncrypted(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) WaitForReadyRead(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::waitForReadyRead")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QSslSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QSslSocket) DestroyQSslSocket() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QSslSocket::~QSslSocket")
		}
	}()

	if ptr.Pointer() != nil {
		C.QSslSocket_DestroyQSslSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
