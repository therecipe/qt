package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QMutableListIterator struct {
	ptr unsafe.Pointer
}

type QMutableListIterator_ITF interface {
	QMutableListIterator_PTR() *QMutableListIterator
}

func (p *QMutableListIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableListIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableListIterator(ptr QMutableListIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableListIterator_PTR().Pointer()
	}
	return nil
}

func NewQMutableListIteratorFromPointer(ptr unsafe.Pointer) *QMutableListIterator {
	var n = new(QMutableListIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableListIterator) QMutableListIterator_PTR() *QMutableListIterator {
	return ptr
}
