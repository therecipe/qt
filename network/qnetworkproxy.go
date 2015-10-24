package network

//#include "qnetworkproxy.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkProxy struct {
	ptr unsafe.Pointer
}

type QNetworkProxyITF interface {
	QNetworkProxyPTR() *QNetworkProxy
}

func (p *QNetworkProxy) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkProxy) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkProxy(ptr QNetworkProxyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkProxyPTR().Pointer()
	}
	return nil
}

func QNetworkProxyFromPointer(ptr unsafe.Pointer) *QNetworkProxy {
	var n = new(QNetworkProxy)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkProxy) QNetworkProxyPTR() *QNetworkProxy {
	return ptr
}

//QNetworkProxy::Capability
type QNetworkProxy__Capability int

var (
	QNetworkProxy__TunnelingCapability      = QNetworkProxy__Capability(0x0001)
	QNetworkProxy__ListeningCapability      = QNetworkProxy__Capability(0x0002)
	QNetworkProxy__UdpTunnelingCapability   = QNetworkProxy__Capability(0x0004)
	QNetworkProxy__CachingCapability        = QNetworkProxy__Capability(0x0008)
	QNetworkProxy__HostNameLookupCapability = QNetworkProxy__Capability(0x0010)
)

//QNetworkProxy::ProxyType
type QNetworkProxy__ProxyType int

var (
	QNetworkProxy__DefaultProxy     = QNetworkProxy__ProxyType(0)
	QNetworkProxy__Socks5Proxy      = QNetworkProxy__ProxyType(1)
	QNetworkProxy__NoProxy          = QNetworkProxy__ProxyType(2)
	QNetworkProxy__HttpProxy        = QNetworkProxy__ProxyType(3)
	QNetworkProxy__HttpCachingProxy = QNetworkProxy__ProxyType(4)
	QNetworkProxy__FtpCachingProxy  = QNetworkProxy__ProxyType(5)
)

func NewQNetworkProxy() *QNetworkProxy {
	return QNetworkProxyFromPointer(unsafe.Pointer(C.QNetworkProxy_NewQNetworkProxy()))
}

func NewQNetworkProxy3(other QNetworkProxyITF) *QNetworkProxy {
	return QNetworkProxyFromPointer(unsafe.Pointer(C.QNetworkProxy_NewQNetworkProxy3(C.QtObjectPtr(PointerFromQNetworkProxy(other)))))
}

func (ptr *QNetworkProxy) Capabilities() QNetworkProxy__Capability {
	if ptr.Pointer() != nil {
		return QNetworkProxy__Capability(C.QNetworkProxy_Capabilities(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkProxy) HasRawHeader(headerName core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkProxy_HasRawHeader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(headerName))) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Header(header QNetworkRequest__KnownHeaders) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_Header(C.QtObjectPtr(ptr.Pointer()), C.int(header)))
	}
	return ""
}

func (ptr *QNetworkProxy) HostName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_HostName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkProxy) IsCachingProxy() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsCachingProxy(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkProxy) IsTransparentProxy() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkProxy_IsTransparentProxy(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkProxy) Password() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_Password(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func QNetworkProxy_SetApplicationProxy(networkProxy QNetworkProxyITF) {
	C.QNetworkProxy_QNetworkProxy_SetApplicationProxy(C.QtObjectPtr(PointerFromQNetworkProxy(networkProxy)))
}

func (ptr *QNetworkProxy) SetCapabilities(capabilities QNetworkProxy__Capability) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetCapabilities(C.QtObjectPtr(ptr.Pointer()), C.int(capabilities))
	}
}

func (ptr *QNetworkProxy) SetHeader(header QNetworkRequest__KnownHeaders, value string) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetHeader(C.QtObjectPtr(ptr.Pointer()), C.int(header), C.CString(value))
	}
}

func (ptr *QNetworkProxy) SetHostName(hostName string) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetHostName(C.QtObjectPtr(ptr.Pointer()), C.CString(hostName))
	}
}

func (ptr *QNetworkProxy) SetPassword(password string) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetPassword(C.QtObjectPtr(ptr.Pointer()), C.CString(password))
	}
}

func (ptr *QNetworkProxy) SetRawHeader(headerName core.QByteArrayITF, headerValue core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetRawHeader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(headerName)), C.QtObjectPtr(core.PointerFromQByteArray(headerValue)))
	}
}

func (ptr *QNetworkProxy) SetType(ty QNetworkProxy__ProxyType) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetType(C.QtObjectPtr(ptr.Pointer()), C.int(ty))
	}
}

func (ptr *QNetworkProxy) SetUser(user string) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_SetUser(C.QtObjectPtr(ptr.Pointer()), C.CString(user))
	}
}

func (ptr *QNetworkProxy) Swap(other QNetworkProxyITF) {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkProxy(other)))
	}
}

func (ptr *QNetworkProxy) Type() QNetworkProxy__ProxyType {
	if ptr.Pointer() != nil {
		return QNetworkProxy__ProxyType(C.QNetworkProxy_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkProxy) User() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkProxy_User(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkProxy) DestroyQNetworkProxy() {
	if ptr.Pointer() != nil {
		C.QNetworkProxy_DestroyQNetworkProxy(C.QtObjectPtr(ptr.Pointer()))
	}
}
