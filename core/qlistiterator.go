package core

//#include "qlistiterator.h"
import "C"
import (
	"unsafe"
)

type QListIterator struct {
	ptr unsafe.Pointer
}

type QListIterator_ITF interface {
	QListIterator_PTR() *QListIterator
}

func (p *QListIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QListIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQListIterator(ptr QListIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QListIterator_PTR().Pointer()
	}
	return nil
}

func NewQListIteratorFromPointer(ptr unsafe.Pointer) *QListIterator {
	var n = new(QListIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QListIterator) QListIterator_PTR() *QListIterator {
	return ptr
}
