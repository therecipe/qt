package network

//#include "qsslpresharedkeyauthenticator.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSslPreSharedKeyAuthenticator struct {
	ptr unsafe.Pointer
}

type QSslPreSharedKeyAuthenticator_ITF interface {
	QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator
}

func (p *QSslPreSharedKeyAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslPreSharedKeyAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslPreSharedKeyAuthenticator(ptr QSslPreSharedKeyAuthenticator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslPreSharedKeyAuthenticator_PTR().Pointer()
	}
	return nil
}

func NewQSslPreSharedKeyAuthenticatorFromPointer(ptr unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	var n = new(QSslPreSharedKeyAuthenticator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslPreSharedKeyAuthenticator) QSslPreSharedKeyAuthenticator_PTR() *QSslPreSharedKeyAuthenticator {
	return ptr
}

func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	return NewQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator())
}

func NewQSslPreSharedKeyAuthenticator2(authenticator QSslPreSharedKeyAuthenticator_ITF) *QSslPreSharedKeyAuthenticator {
	return NewQSslPreSharedKeyAuthenticatorFromPointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(PointerFromQSslPreSharedKeyAuthenticator(authenticator)))
}

func (ptr *QSslPreSharedKeyAuthenticator) Identity() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_Identity(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) IdentityHint() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_IdentityHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	if ptr.Pointer() != nil {
		return int(C.QSslPreSharedKeyAuthenticator_MaximumIdentityLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {
	if ptr.Pointer() != nil {
		return int(C.QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) PreSharedKey() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslPreSharedKeyAuthenticator_PreSharedKey(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSslPreSharedKeyAuthenticator) SetIdentity(identity core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_SetIdentity(ptr.Pointer(), core.PointerFromQByteArray(identity))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_SetPreSharedKey(ptr.Pointer(), core.PointerFromQByteArray(preSharedKey))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) Swap(authenticator QSslPreSharedKeyAuthenticator_ITF) {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_Swap(ptr.Pointer(), PointerFromQSslPreSharedKeyAuthenticator(authenticator))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) DestroyQSslPreSharedKeyAuthenticator() {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(ptr.Pointer())
	}
}
