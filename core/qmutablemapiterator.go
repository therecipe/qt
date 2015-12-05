package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QMutableMapIterator struct {
	ptr unsafe.Pointer
}

type QMutableMapIterator_ITF interface {
	QMutableMapIterator_PTR() *QMutableMapIterator
}

func (p *QMutableMapIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableMapIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableMapIterator(ptr QMutableMapIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableMapIterator_PTR().Pointer()
	}
	return nil
}

func NewQMutableMapIteratorFromPointer(ptr unsafe.Pointer) *QMutableMapIterator {
	var n = new(QMutableMapIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableMapIterator) QMutableMapIterator_PTR() *QMutableMapIterator {
	return ptr
}
