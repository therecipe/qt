package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QMutableHashIterator struct {
	ptr unsafe.Pointer
}

type QMutableHashIterator_ITF interface {
	QMutableHashIterator_PTR() *QMutableHashIterator
}

func (p *QMutableHashIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableHashIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableHashIterator(ptr QMutableHashIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableHashIterator_PTR().Pointer()
	}
	return nil
}

func NewQMutableHashIteratorFromPointer(ptr unsafe.Pointer) *QMutableHashIterator {
	var n = new(QMutableHashIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableHashIterator) QMutableHashIterator_PTR() *QMutableHashIterator {
	return ptr
}
