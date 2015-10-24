package core

//#include "qcollatorsortkey.h"
import "C"
import (
	"unsafe"
)

type QCollatorSortKey struct {
	ptr unsafe.Pointer
}

type QCollatorSortKeyITF interface {
	QCollatorSortKeyPTR() *QCollatorSortKey
}

func (p *QCollatorSortKey) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCollatorSortKey) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCollatorSortKey(ptr QCollatorSortKeyITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCollatorSortKeyPTR().Pointer()
	}
	return nil
}

func QCollatorSortKeyFromPointer(ptr unsafe.Pointer) *QCollatorSortKey {
	var n = new(QCollatorSortKey)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCollatorSortKey) QCollatorSortKeyPTR() *QCollatorSortKey {
	return ptr
}

func NewQCollatorSortKey(other QCollatorSortKeyITF) *QCollatorSortKey {
	return QCollatorSortKeyFromPointer(unsafe.Pointer(C.QCollatorSortKey_NewQCollatorSortKey(C.QtObjectPtr(PointerFromQCollatorSortKey(other)))))
}

func (ptr *QCollatorSortKey) Swap(other QCollatorSortKeyITF) {
	if ptr.Pointer() != nil {
		C.QCollatorSortKey_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCollatorSortKey(other)))
	}
}

func (ptr *QCollatorSortKey) DestroyQCollatorSortKey() {
	if ptr.Pointer() != nil {
		C.QCollatorSortKey_DestroyQCollatorSortKey(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCollatorSortKey) Compare(otherKey QCollatorSortKeyITF) int {
	if ptr.Pointer() != nil {
		return int(C.QCollatorSortKey_Compare(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCollatorSortKey(otherKey))))
	}
	return 0
}
