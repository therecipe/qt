package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QMutableSetIterator struct {
	ptr unsafe.Pointer
}

type QMutableSetIterator_ITF interface {
	QMutableSetIterator_PTR() *QMutableSetIterator
}

func (p *QMutableSetIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableSetIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableSetIterator(ptr QMutableSetIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableSetIterator_PTR().Pointer()
	}
	return nil
}

func NewQMutableSetIteratorFromPointer(ptr unsafe.Pointer) *QMutableSetIterator {
	var n = new(QMutableSetIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableSetIterator) QMutableSetIterator_PTR() *QMutableSetIterator {
	return ptr
}
