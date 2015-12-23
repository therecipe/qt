package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

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
func callbackQNetworkReplyClose(ptrName *C.char) bool {
	defer qt.Recovering("callback QNetworkReply::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQNetworkReplyDownloadProgress(ptrName *C.char, bytesReceived C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QNetworkReply::downloadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "downloadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesReceived), int64(bytesTotal))
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
func callbackQNetworkReplyEncrypted(ptrName *C.char) {
	defer qt.Recovering("callback QNetworkReply::encrypted")

	if signal := qt.GetSignal(C.GoString(ptrName), "encrypted"); signal != nil {
		signal.(func())()
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
func callbackQNetworkReplyError2(ptrName *C.char, code C.int) {
	defer qt.Recovering("callback QNetworkReply::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QNetworkReply__NetworkError))(QNetworkReply__NetworkError(code))
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
func callbackQNetworkReplyFinished(ptrName *C.char) {
	defer qt.Recovering("callback QNetworkReply::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
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
func callbackQNetworkReplyIgnoreSslErrors(ptrName *C.char) bool {
	defer qt.Recovering("callback QNetworkReply::ignoreSslErrors")

	if signal := qt.GetSignal(C.GoString(ptrName), "ignoreSslErrors"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQNetworkReplyMetaDataChanged(ptrName *C.char) {
	defer qt.Recovering("callback QNetworkReply::metaDataChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaDataChanged"); signal != nil {
		signal.(func())()
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
func callbackQNetworkReplyPreSharedKeyAuthenticationRequired(ptrName *C.char, authenticator unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkReply::preSharedKeyAuthenticationRequired")

	if signal := qt.GetSignal(C.GoString(ptrName), "preSharedKeyAuthenticationRequired"); signal != nil {
		signal.(func(*QSslPreSharedKeyAuthenticator))(NewQSslPreSharedKeyAuthenticatorFromPointer(authenticator))
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
func callbackQNetworkReplySetReadBufferSize(ptrName *C.char, size C.longlong) bool {
	defer qt.Recovering("callback QNetworkReply::setReadBufferSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "setReadBufferSize"); signal != nil {
		signal.(func(int64))(int64(size))
		return true
	}
	return false

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
func callbackQNetworkReplyUploadProgress(ptrName *C.char, bytesSent C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QNetworkReply::uploadProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "uploadProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesSent), int64(bytesTotal))
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
func callbackQNetworkReplyTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QNetworkReply::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQNetworkReplyChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QNetworkReply::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQNetworkReplyCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QNetworkReply::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
