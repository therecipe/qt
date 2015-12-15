package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkProxy struct {
	ptr unsafe.Pointer
}

type QNetworkProxy_ITF interface {
	QNetworkProxy_PTR() *QNetworkProxy
}

func (p *QNetworkProxy) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkProxy) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkProxy(ptr QNetworkProxy_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxy_PTR().Pointer()
	}
	return nil
}

func NewQNetworkProxyFromPointer(ptr unsafe.Pointer) *QNetworkProxy {
	var n = new(QNetworkProxy)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkProxy) QNetworkProxy_PTR() *QNetworkProxy {
	return ptr
}

//QNetworkProxy::Capability
type QNetworkProxy__Capability int64

const (
	QNetworkProxy__TunnelingCapability      = QNetworkProxy__Capability(0x0001)
	QNetworkProxy__ListeningCapability      = QNetworkProxy__Capability(0x0002)
	QNetworkProxy__UdpTunnelingCapability   = QNetworkProxy__Capability(0x0004)
	QNetworkProxy__CachingCapability        = QNetworkProxy__Capability(0x0008)
	QNetworkProxy__HostNameLookupCapability = QNetworkProxy__Capability(0x0010)
)

//QNetworkProxy::ProxyType
type QNetworkProxy__ProxyType int64

const (
	QNetworkProxy__DefaultProxy     = QNetworkProxy__ProxyType(0)
	QNetworkProxy__Socks5Proxy      = QNetworkProxy__ProxyType(1)
	QNetworkProxy__NoProxy          = QNetworkProxy__ProxyType(2)
	QNetworkProxy__HttpProxy        = QNetworkProxy__ProxyType(3)
	QNetworkProxy__HttpCachingProxy = QNetworkProxy__ProxyType(4)
	QNetworkProxy__FtpCachingProxy  = QNetworkProxy__ProxyType(5)
)

func NewQNetworkProxy() *QNetworkProxy {
	defer qt.Recovering("QNetworkProxy::QNetworkProxy")

	return NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy())
}

func NewQNetworkProxy3(other QNetworkProxy_ITF) *QNetworkProxy {
	defer qt.Recovering("QNetworkProxy::QNetworkProxy")

	return NewQNetworkProxyFromPointer(C.QNetworkProxy_NewQNetworkProxy3(PointerFromQNetworkProxy(other)))
}

func (ptr *QNetworkProxy) Capabilities() QNetworkProxy__Capability {
	defer qt.Recovering("QNetworkProxy::capabilities")

	if ptr.Pointer() != nil {
		return QNetworkProxy__Capability(C.QNetworkProxy_Capabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) HasRawHeader(headerName core.QByteArray_ITF) bool {
	defer qt.Recovering("QNetworkProxy::hasRawHeader")

	if ptr.Pointer() != nil {
		return C.QNetworkProxy_HasRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Header(header QNetworkRequest__KnownHeaders) *core.QVariant {
	defer qt.Recovering("QNetworkProxy::header")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QNetworkProxy_Header(ptr.Pointer(), C.int(header)))
	}
	return nil
}

func (ptr *QNetworkProxy) HostName() string {
	defer qt.Recovering("QNetworkProxy::hostName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_HostName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) IsCachingProxy() bool {
	defer qt.Recovering("QNetworkProxy::isCachingProxy")

	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsCachingProxy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkProxy) IsTransparentProxy() bool {
	defer qt.Recovering("QNetworkProxy::isTransparentProxy")

	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsTransparentProxy(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Password() string {
	defer qt.Recovering("QNetworkProxy::password")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) RawHeader(headerName core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QNetworkProxy::rawHeader")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkProxy_RawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName)))
	}
	return nil
}

func QNetworkProxy_SetApplicationProxy(networkProxy QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkProxy::setApplicationProxy")

	C.QNetworkProxy_QNetworkProxy_SetApplicationProxy(PointerFromQNetworkProxy(networkProxy))
}

func (ptr *QNetworkProxy) SetCapabilities(capabilities QNetworkProxy__Capability) {
	defer qt.Recovering("QNetworkProxy::setCapabilities")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetCapabilities(ptr.Pointer(), C.int(capabilities))
	}
}

func (ptr *QNetworkProxy) SetHeader(header QNetworkRequest__KnownHeaders, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkProxy::setHeader")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetHeader(ptr.Pointer(), C.int(header), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkProxy) SetHostName(hostName string) {
	defer qt.Recovering("QNetworkProxy::setHostName")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetHostName(ptr.Pointer(), C.CString(hostName))
	}
}

func (ptr *QNetworkProxy) SetPassword(password string) {
	defer qt.Recovering("QNetworkProxy::setPassword")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetPassword(ptr.Pointer(), C.CString(password))
	}
}

func (ptr *QNetworkProxy) SetRawHeader(headerName core.QByteArray_ITF, headerValue core.QByteArray_ITF) {
	defer qt.Recovering("QNetworkProxy::setRawHeader")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetRawHeader(ptr.Pointer(), core.PointerFromQByteArray(headerName), core.PointerFromQByteArray(headerValue))
	}
}

func (ptr *QNetworkProxy) SetType(ty QNetworkProxy__ProxyType) {
	defer qt.Recovering("QNetworkProxy::setType")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QNetworkProxy) SetUser(user string) {
	defer qt.Recovering("QNetworkProxy::setUser")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetUser(ptr.Pointer(), C.CString(user))
	}
}

func (ptr *QNetworkProxy) Swap(other QNetworkProxy_ITF) {
	defer qt.Recovering("QNetworkProxy::swap")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_Swap(ptr.Pointer(), PointerFromQNetworkProxy(other))
	}
}

func (ptr *QNetworkProxy) Type() QNetworkProxy__ProxyType {
	defer qt.Recovering("QNetworkProxy::type")

	if ptr.Pointer() != nil {
		return QNetworkProxy__ProxyType(C.QNetworkProxy_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkProxy) User() string {
	defer qt.Recovering("QNetworkProxy::user")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_User(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkProxy) DestroyQNetworkProxy() {
	defer qt.Recovering("QNetworkProxy::~QNetworkProxy")

	if ptr.Pointer() != nil {
		C.QNetworkProxy_DestroyQNetworkProxy(ptr.Pointer())
	}
}
