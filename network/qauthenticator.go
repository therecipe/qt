package network

//#include "qauthenticator.h"
import "C"
import (
	"unsafe"
)

type QAuthenticator struct {
	ptr unsafe.Pointer
}

type QAuthenticatorITF interface {
	QAuthenticatorPTR() *QAuthenticator
}

func (p *QAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAuthenticator(ptr QAuthenticatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAuthenticatorPTR().Pointer()
	}
	return nil
}

func QAuthenticatorFromPointer(ptr unsafe.Pointer) *QAuthenticator {
	var n = new(QAuthenticator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAuthenticator) QAuthenticatorPTR() *QAuthenticator {
	return ptr
}

func NewQAuthenticator() *QAuthenticator {
	return QAuthenticatorFromPointer(unsafe.Pointer(C.QAuthenticator_NewQAuthenticator()))
}

func NewQAuthenticator2(other QAuthenticatorITF) *QAuthenticator {
	return QAuthenticatorFromPointer(unsafe.Pointer(C.QAuthenticator_NewQAuthenticator2(C.QtObjectPtr(PointerFromQAuthenticator(other)))))
}

func (ptr *QAuthenticator) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QAuthenticator_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAuthenticator) Option(opt string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Option(C.QtObjectPtr(ptr.Pointer()), C.CString(opt)))
	}
	return ""
}

func (ptr *QAuthenticator) Password() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Password(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAuthenticator) Realm() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Realm(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAuthenticator) SetOption(opt string, value string) {
	if ptr.Pointer() != nil {
		C.QAuthenticator_SetOption(C.QtObjectPtr(ptr.Pointer()), C.CString(opt), C.CString(value))
	}
}

func (ptr *QAuthenticator) SetPassword(password string) {
	if ptr.Pointer() != nil {
		C.QAuthenticator_SetPassword(C.QtObjectPtr(ptr.Pointer()), C.CString(password))
	}
}

func (ptr *QAuthenticator) SetUser(user string) {
	if ptr.Pointer() != nil {
		C.QAuthenticator_SetUser(C.QtObjectPtr(ptr.Pointer()), C.CString(user))
	}
}

func (ptr *QAuthenticator) User() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_User(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAuthenticator) DestroyQAuthenticator() {
	if ptr.Pointer() != nil {
		C.QAuthenticator_DestroyQAuthenticator(C.QtObjectPtr(ptr.Pointer()))
	}
}
