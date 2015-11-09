package core

//#include "qsetiterator.h"
import "C"
import (
	"unsafe"
)

type QSetIterator struct {
	ptr unsafe.Pointer
}

type QSetIterator_ITF interface {
	QSetIterator_PTR() *QSetIterator
}

func (p *QSetIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSetIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSetIterator(ptr QSetIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSetIterator_PTR().Pointer()
	}
	return nil
}

func NewQSetIteratorFromPointer(ptr unsafe.Pointer) *QSetIterator {
	var n = new(QSetIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSetIterator) QSetIterator_PTR() *QSetIterator {
	return ptr
}
