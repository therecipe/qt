package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QMutableVectorIterator struct {
	ptr unsafe.Pointer
}

type QMutableVectorIterator_ITF interface {
	QMutableVectorIterator_PTR() *QMutableVectorIterator
}

func (p *QMutableVectorIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableVectorIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableVectorIterator(ptr QMutableVectorIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableVectorIterator_PTR().Pointer()
	}
	return nil
}

func NewQMutableVectorIteratorFromPointer(ptr unsafe.Pointer) *QMutableVectorIterator {
	var n = new(QMutableVectorIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableVectorIterator) QMutableVectorIterator_PTR() *QMutableVectorIterator {
	return ptr
}
