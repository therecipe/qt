package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAuthenticator struct {
	ptr unsafe.Pointer
}

type QAuthenticator_ITF interface {
	QAuthenticator_PTR() *QAuthenticator
}

func (p *QAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAuthenticator(ptr QAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQAuthenticatorFromPointer(ptr unsafe.Pointer) *QAuthenticator {
	var n = new(QAuthenticator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAuthenticator) QAuthenticator_PTR() *QAuthenticator {
	return ptr
}

func NewQAuthenticator() *QAuthenticator {
	defer qt.Recovering("QAuthenticator::QAuthenticator")

	return NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator())
}

func NewQAuthenticator2(other QAuthenticator_ITF) *QAuthenticator {
	defer qt.Recovering("QAuthenticator::QAuthenticator")

	return NewQAuthenticatorFromPointer(C.QAuthenticator_NewQAuthenticator2(PointerFromQAuthenticator(other)))
}

func (ptr *QAuthenticator) IsNull() bool {
	defer qt.Recovering("QAuthenticator::isNull")

	if ptr.Pointer() != nil {
		return C.QAuthenticator_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAuthenticator) Option(opt string) *core.QVariant {
	defer qt.Recovering("QAuthenticator::option")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAuthenticator_Option(ptr.Pointer(), C.CString(opt)))
	}
	return nil
}

func (ptr *QAuthenticator) Password() string {
	defer qt.Recovering("QAuthenticator::password")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Password(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) Realm() string {
	defer qt.Recovering("QAuthenticator::realm")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_Realm(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) SetOption(opt string, value core.QVariant_ITF) {
	defer qt.Recovering("QAuthenticator::setOption")

	if ptr.Pointer() != nil {
		C.QAuthenticator_SetOption(ptr.Pointer(), C.CString(opt), core.PointerFromQVariant(value))
	}
}

func (ptr *QAuthenticator) SetPassword(password string) {
	defer qt.Recovering("QAuthenticator::setPassword")

	if ptr.Pointer() != nil {
		C.QAuthenticator_SetPassword(ptr.Pointer(), C.CString(password))
	}
}

func (ptr *QAuthenticator) SetUser(user string) {
	defer qt.Recovering("QAuthenticator::setUser")

	if ptr.Pointer() != nil {
		C.QAuthenticator_SetUser(ptr.Pointer(), C.CString(user))
	}
}

func (ptr *QAuthenticator) User() string {
	defer qt.Recovering("QAuthenticator::user")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAuthenticator_User(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAuthenticator) DestroyQAuthenticator() {
	defer qt.Recovering("QAuthenticator::~QAuthenticator")

	if ptr.Pointer() != nil {
		C.QAuthenticator_DestroyQAuthenticator(ptr.Pointer())
	}
}
