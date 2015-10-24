package network

//#include "qsslkey.h"
import "C"
import (
	"unsafe"
)

type QSslKey struct {
	ptr unsafe.Pointer
}

type QSslKeyITF interface {
	QSslKeyPTR() *QSslKey
}

func (p *QSslKey) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslKey) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslKey(ptr QSslKeyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslKeyPTR().Pointer()
	}
	return nil
}

func QSslKeyFromPointer(ptr unsafe.Pointer) *QSslKey {
	var n = new(QSslKey)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslKey) QSslKeyPTR() *QSslKey {
	return ptr
}

func NewQSslKey() *QSslKey {
	return QSslKeyFromPointer(unsafe.Pointer(C.QSslKey_NewQSslKey()))
}

func NewQSslKey5(other QSslKeyITF) *QSslKey {
	return QSslKeyFromPointer(unsafe.Pointer(C.QSslKey_NewQSslKey5(C.QtObjectPtr(PointerFromQSslKey(other)))))
}

func (ptr *QSslKey) Clear() {
	if ptr.Pointer() != nil {
		C.QSslKey_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSslKey) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslKey_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSslKey) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QSslKey_Length(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSslKey) Swap(other QSslKeyITF) {
	if ptr.Pointer() != nil {
		C.QSslKey_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSslKey(other)))
	}
}

func (ptr *QSslKey) DestroyQSslKey() {
	if ptr.Pointer() != nil {
		C.QSslKey_DestroyQSslKey(C.QtObjectPtr(ptr.Pointer()))
	}
}
