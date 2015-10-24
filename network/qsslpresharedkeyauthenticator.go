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

type QSslPreSharedKeyAuthenticatorITF interface {
	QSslPreSharedKeyAuthenticatorPTR() *QSslPreSharedKeyAuthenticator
}

func (p *QSslPreSharedKeyAuthenticator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslPreSharedKeyAuthenticator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslPreSharedKeyAuthenticator(ptr QSslPreSharedKeyAuthenticatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslPreSharedKeyAuthenticatorPTR().Pointer()
	}
	return nil
}

func QSslPreSharedKeyAuthenticatorFromPointer(ptr unsafe.Pointer) *QSslPreSharedKeyAuthenticator {
	var n = new(QSslPreSharedKeyAuthenticator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslPreSharedKeyAuthenticator) QSslPreSharedKeyAuthenticatorPTR() *QSslPreSharedKeyAuthenticator {
	return ptr
}

func NewQSslPreSharedKeyAuthenticator() *QSslPreSharedKeyAuthenticator {
	return QSslPreSharedKeyAuthenticatorFromPointer(unsafe.Pointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator()))
}

func NewQSslPreSharedKeyAuthenticator2(authenticator QSslPreSharedKeyAuthenticatorITF) *QSslPreSharedKeyAuthenticator {
	return QSslPreSharedKeyAuthenticatorFromPointer(unsafe.Pointer(C.QSslPreSharedKeyAuthenticator_NewQSslPreSharedKeyAuthenticator2(C.QtObjectPtr(PointerFromQSslPreSharedKeyAuthenticator(authenticator)))))
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumIdentityLength() int {
	if ptr.Pointer() != nil {
		return int(C.QSslPreSharedKeyAuthenticator_MaximumIdentityLength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) MaximumPreSharedKeyLength() int {
	if ptr.Pointer() != nil {
		return int(C.QSslPreSharedKeyAuthenticator_MaximumPreSharedKeyLength(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslPreSharedKeyAuthenticator) SetIdentity(identity core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_SetIdentity(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(identity)))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) SetPreSharedKey(preSharedKey core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_SetPreSharedKey(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(preSharedKey)))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) Swap(authenticator QSslPreSharedKeyAuthenticatorITF) {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslPreSharedKeyAuthenticator(authenticator)))
	}
}

func (ptr *QSslPreSharedKeyAuthenticator) DestroyQSslPreSharedKeyAuthenticator() {
	if ptr.Pointer() != nil {
		C.QSslPreSharedKeyAuthenticator_DestroyQSslPreSharedKeyAuthenticator(C.QtObjectPtr(ptr.Pointer()))
	}
}
