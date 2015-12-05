package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QFutureIterator struct {
	ptr unsafe.Pointer
}

type QFutureIterator_ITF interface {
	QFutureIterator_PTR() *QFutureIterator
}

func (p *QFutureIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFutureIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFutureIterator(ptr QFutureIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFutureIterator_PTR().Pointer()
	}
	return nil
}

func NewQFutureIteratorFromPointer(ptr unsafe.Pointer) *QFutureIterator {
	var n = new(QFutureIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFutureIterator) QFutureIterator_PTR() *QFutureIterator {
	return ptr
}
