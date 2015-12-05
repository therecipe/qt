package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QCollatorSortKey struct {
	ptr unsafe.Pointer
}

type QCollatorSortKey_ITF interface {
	QCollatorSortKey_PTR() *QCollatorSortKey
}

func (p *QCollatorSortKey) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCollatorSortKey) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCollatorSortKey(ptr QCollatorSortKey_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCollatorSortKey_PTR().Pointer()
	}
	return nil
}

func NewQCollatorSortKeyFromPointer(ptr unsafe.Pointer) *QCollatorSortKey {
	var n = new(QCollatorSortKey)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCollatorSortKey) QCollatorSortKey_PTR() *QCollatorSortKey {
	return ptr
}

func NewQCollatorSortKey(other QCollatorSortKey_ITF) *QCollatorSortKey {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCollatorSortKey::QCollatorSortKey")
		}
	}()

	return NewQCollatorSortKeyFromPointer(C.QCollatorSortKey_NewQCollatorSortKey(PointerFromQCollatorSortKey(other)))
}

func (ptr *QCollatorSortKey) Swap(other QCollatorSortKey_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCollatorSortKey::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCollatorSortKey_Swap(ptr.Pointer(), PointerFromQCollatorSortKey(other))
	}
}

func (ptr *QCollatorSortKey) DestroyQCollatorSortKey() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCollatorSortKey::~QCollatorSortKey")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCollatorSortKey_DestroyQCollatorSortKey(ptr.Pointer())
	}
}

func (ptr *QCollatorSortKey) Compare(otherKey QCollatorSortKey_ITF) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCollatorSortKey::compare")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QCollatorSortKey_Compare(ptr.Pointer(), PointerFromQCollatorSortKey(otherKey)))
	}
	return 0
}
