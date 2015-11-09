package core

//#include "qhashiterator.h"
import "C"
import (
	"unsafe"
)

type QHashIterator struct {
	ptr unsafe.Pointer
}

type QHashIterator_ITF interface {
	QHashIterator_PTR() *QHashIterator
}

func (p *QHashIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHashIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHashIterator(ptr QHashIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHashIterator_PTR().Pointer()
	}
	return nil
}

func NewQHashIteratorFromPointer(ptr unsafe.Pointer) *QHashIterator {
	var n = new(QHashIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHashIterator) QHashIterator_PTR() *QHashIterator {
	return ptr
}
