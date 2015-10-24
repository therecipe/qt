package network

//#include "qnetworkcookie.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkCookie struct {
	ptr unsafe.Pointer
}

type QNetworkCookieITF interface {
	QNetworkCookiePTR() *QNetworkCookie
}

func (p *QNetworkCookie) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkCookie) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkCookie(ptr QNetworkCookieITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookiePTR().Pointer()
	}
	return nil
}

func QNetworkCookieFromPointer(ptr unsafe.Pointer) *QNetworkCookie {
	var n = new(QNetworkCookie)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkCookie) QNetworkCookiePTR() *QNetworkCookie {
	return ptr
}

//QNetworkCookie::RawForm
type QNetworkCookie__RawForm int

var (
	QNetworkCookie__NameAndValueOnly = QNetworkCookie__RawForm(0)
	QNetworkCookie__Full             = QNetworkCookie__RawForm(1)
)

func NewQNetworkCookie(name core.QByteArrayITF, value core.QByteArrayITF) *QNetworkCookie {
	return QNetworkCookieFromPointer(unsafe.Pointer(C.QNetworkCookie_NewQNetworkCookie(C.QtObjectPtr(core.PointerFromQByteArray(name)), C.QtObjectPtr(core.PointerFromQByteArray(value)))))
}

func NewQNetworkCookie2(other QNetworkCookieITF) *QNetworkCookie {
	return QNetworkCookieFromPointer(unsafe.Pointer(C.QNetworkCookie_NewQNetworkCookie2(C.QtObjectPtr(PointerFromQNetworkCookie(other)))))
}

func (ptr *QNetworkCookie) Domain() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkCookie_Domain(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkCookie) HasSameIdentifier(other QNetworkCookieITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_HasSameIdentifier(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCookie(other))) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsHttpOnly() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsHttpOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSecure() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSecure(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSessionCookie() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSessionCookie(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkCookie) Normalize(url string) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_Normalize(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QNetworkCookie) Path() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkCookie_Path(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkCookie) SetDomain(domain string) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetDomain(C.QtObjectPtr(ptr.Pointer()), C.CString(domain))
	}
}

func (ptr *QNetworkCookie) SetExpirationDate(date core.QDateTimeITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetExpirationDate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDateTime(date)))
	}
}

func (ptr *QNetworkCookie) SetHttpOnly(enable bool) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetHttpOnly(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QNetworkCookie) SetName(cookieName core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetName(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(cookieName)))
	}
}

func (ptr *QNetworkCookie) SetPath(path string) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetPath(C.QtObjectPtr(ptr.Pointer()), C.CString(path))
	}
}

func (ptr *QNetworkCookie) SetSecure(enable bool) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetSecure(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QNetworkCookie) SetValue(value core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetValue(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(value)))
	}
}

func (ptr *QNetworkCookie) Swap(other QNetworkCookieITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkCookie(other)))
	}
}

func (ptr *QNetworkCookie) DestroyQNetworkCookie() {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_DestroyQNetworkCookie(C.QtObjectPtr(ptr.Pointer()))
	}
}
