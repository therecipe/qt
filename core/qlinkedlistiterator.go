package core

//#include "qlinkedlistiterator.h"
import "C"
import (
	"unsafe"
)

type QLinkedListIterator struct {
	ptr unsafe.Pointer
}

type QLinkedListIterator_ITF interface {
	QLinkedListIterator_PTR() *QLinkedListIterator
}

func (p *QLinkedListIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLinkedListIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLinkedListIterator(ptr QLinkedListIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLinkedListIterator_PTR().Pointer()
	}
	return nil
}

func NewQLinkedListIteratorFromPointer(ptr unsafe.Pointer) *QLinkedListIterator {
	var n = new(QLinkedListIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLinkedListIterator) QLinkedListIterator_PTR() *QLinkedListIterator {
	return ptr
}
