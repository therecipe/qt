package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

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

	return NewQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest2(PointerFromQNetworkRequest(other)))
}

func NewQNetworkRequest(url core.QUrl_ITF) *QNetworkRequest {
	defer qt.Recovering("QNetworkRequest::QNetworkRequest")

	return NewQNetworkRequestFromPointer(C.QNetworkRequest_NewQNetworkRequest(core.PointerFromQUrl(url)))
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
