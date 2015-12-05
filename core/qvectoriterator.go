package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QVectorIterator struct {
	ptr unsafe.Pointer
}

type QVectorIterator_ITF interface {
	QVectorIterator_PTR() *QVectorIterator
}

func (p *QVectorIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVectorIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVectorIterator(ptr QVectorIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVectorIterator_PTR().Pointer()
	}
	return nil
}

func NewQVectorIteratorFromPointer(ptr unsafe.Pointer) *QVectorIterator {
	var n = new(QVectorIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVectorIterator) QVectorIterator_PTR() *QVectorIterator {
	return ptr
}
