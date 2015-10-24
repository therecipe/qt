package network

//#include "qnetworkreply.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkReply struct {
	core.QIODevice
}

type QNetworkReplyITF interface {
	core.QIODeviceITF
	QNetworkReplyPTR() *QNetworkReply
}

func PointerFromQNetworkReply(ptr QNetworkReplyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkReplyPTR().Pointer()
	}
	return nil
}

func QNetworkReplyFromPointer(ptr unsafe.Pointer) *QNetworkReply {
	var n = new(QNetworkReply)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNetworkReply_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkReply) QNetworkReplyPTR() *QNetworkReply {
	return ptr
}

//QNetworkReply::NetworkError
type QNetworkReply__NetworkError int

var (
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
	if ptr.Pointer() != nil {
		C.QNetworkReply_Abort(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkReply) Attribute(code QNetworkRequest__Attribute) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkReply_Attribute(C.QtObjectPtr(ptr.Pointer()), C.int(code)))
	}
	return ""
}

func (ptr *QNetworkReply) Close() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkReply) ConnectEncrypted(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectEncrypted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "encrypted", f)
	}
}

func (ptr *QNetworkReply) DisconnectEncrypted() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectEncrypted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "encrypted")
	}
}

//export callbackQNetworkReplyEncrypted
func callbackQNetworkReplyEncrypted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "encrypted").(func())()
}

func (ptr *QNetworkReply) Error() QNetworkReply__NetworkError {
	if ptr.Pointer() != nil {
		return QNetworkReply__NetworkError(C.QNetworkReply_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkReply) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QNetworkReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQNetworkReplyFinished
func callbackQNetworkReplyFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func (ptr *QNetworkReply) HasRawHeader(headerName core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_HasRawHeader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(headerName))) != 0
	}
	return false
}

func (ptr *QNetworkReply) Header(header QNetworkRequest__KnownHeaders) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkReply_Header(C.QtObjectPtr(ptr.Pointer()), C.int(header)))
	}
	return ""
}

func (ptr *QNetworkReply) IgnoreSslErrors() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_IgnoreSslErrors(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsFinished(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkReply) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkReply_IsRunning(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkReply) Manager() *QNetworkAccessManager {
	if ptr.Pointer() != nil {
		return QNetworkAccessManagerFromPointer(unsafe.Pointer(C.QNetworkReply_Manager(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QNetworkReply) ConnectMetaDataChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectMetaDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "metaDataChanged", f)
	}
}

func (ptr *QNetworkReply) DisconnectMetaDataChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectMetaDataChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "metaDataChanged")
	}
}

//export callbackQNetworkReplyMetaDataChanged
func callbackQNetworkReplyMetaDataChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "metaDataChanged").(func())()
}

func (ptr *QNetworkReply) Operation() QNetworkAccessManager__Operation {
	if ptr.Pointer() != nil {
		return QNetworkAccessManager__Operation(C.QNetworkReply_Operation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkReply) ConnectPreSharedKeyAuthenticationRequired(f func(authenticator QSslPreSharedKeyAuthenticatorITF)) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_ConnectPreSharedKeyAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired", f)
	}
}

func (ptr *QNetworkReply) DisconnectPreSharedKeyAuthenticationRequired() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DisconnectPreSharedKeyAuthenticationRequired(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "preSharedKeyAuthenticationRequired")
	}
}

//export callbackQNetworkReplyPreSharedKeyAuthenticationRequired
func callbackQNetworkReplyPreSharedKeyAuthenticationRequired(ptrName *C.char, authenticator unsafe.Pointer) {
	qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired").(func(*QSslPreSharedKeyAuthenticator))(QSslPreSharedKeyAuthenticatorFromPointer(authenticator))
}

func (ptr *QNetworkReply) SetSslConfiguration(config QSslConfigurationITF) {
	if ptr.Pointer() != nil {
		C.QNetworkReply_SetSslConfiguration(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslConfiguration(config)))
	}
}

func (ptr *QNetworkReply) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkReply_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkReply) DestroyQNetworkReply() {
	if ptr.Pointer() != nil {
		C.QNetworkReply_DestroyQNetworkReply(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
