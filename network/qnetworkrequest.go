package network

//#include "qnetworkrequest.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkRequest struct {
	ptr unsafe.Pointer
}

type QNetworkRequestITF interface {
	QNetworkRequestPTR() *QNetworkRequest
}

func (p *QNetworkRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkRequest(ptr QNetworkRequestITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkRequestPTR().Pointer()
	}
	return nil
}

func QNetworkRequestFromPointer(ptr unsafe.Pointer) *QNetworkRequest {
	var n = new(QNetworkRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkRequest) QNetworkRequestPTR() *QNetworkRequest {
	return ptr
}

//QNetworkRequest::Attribute
type QNetworkRequest__Attribute int

var (
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
type QNetworkRequest__CacheLoadControl int

var (
	QNetworkRequest__AlwaysNetwork = QNetworkRequest__CacheLoadControl(0)
	QNetworkRequest__PreferNetwork = QNetworkRequest__CacheLoadControl(1)
	QNetworkRequest__PreferCache   = QNetworkRequest__CacheLoadControl(2)
	QNetworkRequest__AlwaysCache   = QNetworkRequest__CacheLoadControl(3)
)

//QNetworkRequest::KnownHeaders
type QNetworkRequest__KnownHeaders int

var (
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
type QNetworkRequest__LoadControl int

var (
	QNetworkRequest__Automatic = QNetworkRequest__LoadControl(0)
	QNetworkRequest__Manual    = QNetworkRequest__LoadControl(1)
)

//QNetworkRequest::Priority
type QNetworkRequest__Priority int

var (
	QNetworkRequest__HighPriority   = QNetworkRequest__Priority(1)
	QNetworkRequest__NormalPriority = QNetworkRequest__Priority(3)
	QNetworkRequest__LowPriority    = QNetworkRequest__Priority(5)
)

func NewQNetworkRequest2(other QNetworkRequestITF) *QNetworkRequest {
	return QNetworkRequestFromPointer(unsafe.Pointer(C.QNetworkRequest_NewQNetworkRequest2(C.QtObjectPtr(PointerFromQNetworkRequest(other)))))
}

func NewQNetworkRequest(url string) *QNetworkRequest {
	return QNetworkRequestFromPointer(unsafe.Pointer(C.QNetworkRequest_NewQNetworkRequest(C.CString(url))))
}

func (ptr *QNetworkRequest) Attribute(code QNetworkRequest__Attribute, defaultValue string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkRequest_Attribute(C.QtObjectPtr(ptr.Pointer()), C.int(code), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QNetworkRequest) HasRawHeader(headerName core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkRequest_HasRawHeader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(headerName))) != 0
	}
	return false
}

func (ptr *QNetworkRequest) Header(header QNetworkRequest__KnownHeaders) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkRequest_Header(C.QtObjectPtr(ptr.Pointer()), C.int(header)))
	}
	return ""
}

func (ptr *QNetworkRequest) OriginatingObject() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QNetworkRequest_OriginatingObject(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QNetworkRequest) Priority() QNetworkRequest__Priority {
	if ptr.Pointer() != nil {
		return QNetworkRequest__Priority(C.QNetworkRequest_Priority(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkRequest) SetAttribute(code QNetworkRequest__Attribute, value string) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(code), C.CString(value))
	}
}

func (ptr *QNetworkRequest) SetHeader(header QNetworkRequest__KnownHeaders, value string) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetHeader(C.QtObjectPtr(ptr.Pointer()), C.int(header), C.CString(value))
	}
}

func (ptr *QNetworkRequest) SetOriginatingObject(object core.QObjectITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetOriginatingObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object)))
	}
}

func (ptr *QNetworkRequest) SetPriority(priority QNetworkRequest__Priority) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetPriority(C.QtObjectPtr(ptr.Pointer()), C.int(priority))
	}
}

func (ptr *QNetworkRequest) SetRawHeader(headerName core.QByteArrayITF, headerValue core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetRawHeader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(headerName)), C.QtObjectPtr(core.PointerFromQByteArray(headerValue)))
	}
}

func (ptr *QNetworkRequest) SetSslConfiguration(config QSslConfigurationITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetSslConfiguration(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslConfiguration(config)))
	}
}

func (ptr *QNetworkRequest) SetUrl(url string) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_SetUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QNetworkRequest) Swap(other QNetworkRequestITF) {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkRequest(other)))
	}
}

func (ptr *QNetworkRequest) Url() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkRequest_Url(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkRequest) DestroyQNetworkRequest() {
	if ptr.Pointer() != nil {
		C.QNetworkRequest_DestroyQNetworkRequest(C.QtObjectPtr(ptr.Pointer()))
	}
}
