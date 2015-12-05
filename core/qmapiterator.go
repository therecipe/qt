package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QMapIterator struct {
	ptr unsafe.Pointer
}

type QMapIterator_ITF interface {
	QMapIterator_PTR() *QMapIterator
}

func (p *QMapIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMapIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMapIterator(ptr QMapIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMapIterator_PTR().Pointer()
	}
	return nil
}

func NewQMapIteratorFromPointer(ptr unsafe.Pointer) *QMapIterator {
	var n = new(QMapIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMapIterator) QMapIterator_PTR() *QMapIterator {
	return ptr
}
