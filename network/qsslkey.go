package network

//#include "qsslkey.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSslKey struct {
	ptr unsafe.Pointer
}

type QSslKey_ITF interface {
	QSslKey_PTR() *QSslKey
}

func (p *QSslKey) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSslKey) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSslKey(ptr QSslKey_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSslKey_PTR().Pointer()
	}
	return nil
}

func NewQSslKeyFromPointer(ptr unsafe.Pointer) *QSslKey {
	var n = new(QSslKey)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSslKey) QSslKey_PTR() *QSslKey {
	return ptr
}

func NewQSslKey() *QSslKey {
	return NewQSslKeyFromPointer(C.QSslKey_NewQSslKey())
}

func NewQSslKey5(other QSslKey_ITF) *QSslKey {
	return NewQSslKeyFromPointer(C.QSslKey_NewQSslKey5(PointerFromQSslKey(other)))
}

func (ptr *QSslKey) Clear() {
	if ptr.Pointer() != nil {
		C.QSslKey_Clear(ptr.Pointer())
	}
}

func (ptr *QSslKey) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QSslKey_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslKey) Length() int {
	if ptr.Pointer() != nil {
		return int(C.QSslKey_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslKey) Swap(other QSslKey_ITF) {
	if ptr.Pointer() != nil {
		C.QSslKey_Swap(ptr.Pointer(), PointerFromQSslKey(other))
	}
}

func (ptr *QSslKey) ToDer(passPhrase core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslKey_ToDer(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
	}
	return nil
}

func (ptr *QSslKey) ToPem(passPhrase core.QByteArray_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslKey_ToPem(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
	}
	return nil
}

func (ptr *QSslKey) DestroyQSslKey() {
	if ptr.Pointer() != nil {
		C.QSslKey_DestroyQSslKey(ptr.Pointer())
	}
}
