package network

//#include "qsslcipher.h"
import "C"
import (
	"unsafe"
)

type QSslCipher struct {
	ptr unsafe.Pointer
}

type QSslCipherITF interface {
	QSslCipherPTR() *QSslCipher
}

func (p *QSslCipher) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslCipher) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslCipher(ptr QSslCipherITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslCipherPTR().Pointer()
	}
	return nil
}

func QSslCipherFromPointer(ptr unsafe.Pointer) *QSslCipher {
	var n = new(QSslCipher)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslCipher) QSslCipherPTR() *QSslCipher {
	return ptr
}

func NewQSslCipher() *QSslCipher {
	return QSslCipherFromPointer(unsafe.Pointer(C.QSslCipher_NewQSslCipher()))
}

func NewQSslCipher4(other QSslCipherITF) *QSslCipher {
	return QSslCipherFromPointer(unsafe.Pointer(C.QSslCipher_NewQSslCipher4(C.QtObjectPtr(PointerFromQSslCipher(other)))))
}

func NewQSslCipher2(name string) *QSslCipher {
	return QSslCipherFromPointer(unsafe.Pointer(C.QSslCipher_NewQSslCipher2(C.CString(name))))
}

func (ptr *QSslCipher) AuthenticationMethod() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_AuthenticationMethod(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCipher) EncryptionMethod() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_EncryptionMethod(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCipher) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslCipher_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslCipher) KeyExchangeMethod() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_KeyExchangeMethod(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCipher) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCipher) ProtocolString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSslCipher_ProtocolString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSslCipher) SupportedBits() int {
	if ptr.Pointer() != nil {
		return int(C.QSslCipher_SupportedBits(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslCipher) Swap(other QSslCipherITF) {
	if ptr.Pointer() != nil {
		C.QSslCipher_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslCipher(other)))
	}
}

func (ptr *QSslCipher) UsedBits() int {
	if ptr.Pointer() != nil {
		return int(C.QSslCipher_UsedBits(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslCipher) DestroyQSslCipher() {
	if ptr.Pointer() != nil {
		C.QSslCipher_DestroyQSslCipher(C.QtObjectPtr(ptr.Pointer()))
	}
}
