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

type QNetworkCookie_ITF interface {
	QNetworkCookie_PTR() *QNetworkCookie
}

func (p *QNetworkCookie) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkCookie) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkCookie(ptr QNetworkCookie_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkCookie_PTR().Pointer()
	}
	return nil
}

func NewQNetworkCookieFromPointer(ptr unsafe.Pointer) *QNetworkCookie {
	var n = new(QNetworkCookie)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkCookie) QNetworkCookie_PTR() *QNetworkCookie {
	return ptr
}

//QNetworkCookie::RawForm
type QNetworkCookie__RawForm int64

const (
	QNetworkCookie__NameAndValueOnly = QNetworkCookie__RawForm(0)
	QNetworkCookie__Full             = QNetworkCookie__RawForm(1)
)

func NewQNetworkCookie(name core.QByteArray_ITF, value core.QByteArray_ITF) *QNetworkCookie {
	return NewQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie(core.PointerFromQByteArray(name), core.PointerFromQByteArray(value)))
}

func NewQNetworkCookie2(other QNetworkCookie_ITF) *QNetworkCookie {
	return NewQNetworkCookieFromPointer(C.QNetworkCookie_NewQNetworkCookie2(PointerFromQNetworkCookie(other)))
}

func (ptr *QNetworkCookie) Domain() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkCookie_Domain(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkCookie) ExpirationDate() *core.QDateTime {
	if ptr.Pointer() != nil {
		return core.NewQDateTimeFromPointer(C.QNetworkCookie_ExpirationDate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCookie) HasSameIdentifier(other QNetworkCookie_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_HasSameIdentifier(ptr.Pointer(), PointerFromQNetworkCookie(other)) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsHttpOnly() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsHttpOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSecure() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSecure(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) IsSessionCookie() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkCookie_IsSessionCookie(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkCookie) Name() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkCookie_Name(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCookie) Normalize(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_Normalize(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QNetworkCookie) Path() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkCookie_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkCookie) SetDomain(domain string) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetDomain(ptr.Pointer(), C.CString(domain))
	}
}

func (ptr *QNetworkCookie) SetExpirationDate(date core.QDateTime_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetExpirationDate(ptr.Pointer(), core.PointerFromQDateTime(date))
	}
}

func (ptr *QNetworkCookie) SetHttpOnly(enable bool) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetHttpOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QNetworkCookie) SetName(cookieName core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetName(ptr.Pointer(), core.PointerFromQByteArray(cookieName))
	}
}

func (ptr *QNetworkCookie) SetPath(path string) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetPath(ptr.Pointer(), C.CString(path))
	}
}

func (ptr *QNetworkCookie) SetSecure(enable bool) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetSecure(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QNetworkCookie) SetValue(value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_SetValue(ptr.Pointer(), core.PointerFromQByteArray(value))
	}
}

func (ptr *QNetworkCookie) Swap(other QNetworkCookie_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_Swap(ptr.Pointer(), PointerFromQNetworkCookie(other))
	}
}

func (ptr *QNetworkCookie) ToRawForm(form QNetworkCookie__RawForm) *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkCookie_ToRawForm(ptr.Pointer(), C.int(form)))
	}
	return nil
}

func (ptr *QNetworkCookie) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QNetworkCookie_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QNetworkCookie) DestroyQNetworkCookie() {
	if ptr.Pointer() != nil {
		C.QNetworkCookie_DestroyQNetworkCookie(ptr.Pointer())
	}
}
