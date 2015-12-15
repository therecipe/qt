package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QSslKey::QSslKey")

	return NewQSslKeyFromPointer(C.QSslKey_NewQSslKey())
}

func NewQSslKey5(other QSslKey_ITF) *QSslKey {
	defer qt.Recovering("QSslKey::QSslKey")

	return NewQSslKeyFromPointer(C.QSslKey_NewQSslKey5(PointerFromQSslKey(other)))
}

func (ptr *QSslKey) Clear() {
	defer qt.Recovering("QSslKey::clear")

	if ptr.Pointer() != nil {
		C.QSslKey_Clear(ptr.Pointer())
	}
}

func (ptr *QSslKey) IsNull() bool {
	defer qt.Recovering("QSslKey::isNull")

	if ptr.Pointer() != nil {
		return C.QSslKey_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSslKey) Length() int {
	defer qt.Recovering("QSslKey::length")

	if ptr.Pointer() != nil {
		return int(C.QSslKey_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSslKey) Swap(other QSslKey_ITF) {
	defer qt.Recovering("QSslKey::swap")

	if ptr.Pointer() != nil {
		C.QSslKey_Swap(ptr.Pointer(), PointerFromQSslKey(other))
	}
}

func (ptr *QSslKey) ToDer(passPhrase core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QSslKey::toDer")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslKey_ToDer(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
	}
	return nil
}

func (ptr *QSslKey) ToPem(passPhrase core.QByteArray_ITF) *core.QByteArray {
	defer qt.Recovering("QSslKey::toPem")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QSslKey_ToPem(ptr.Pointer(), core.PointerFromQByteArray(passPhrase)))
	}
	return nil
}

func (ptr *QSslKey) DestroyQSslKey() {
	defer qt.Recovering("QSslKey::~QSslKey")

	if ptr.Pointer() != nil {
		C.QSslKey_DestroyQSslKey(ptr.Pointer())
	}
}
