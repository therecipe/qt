package core

//#include "qmutablelinkedlistiterator.h"
import "C"
import (
	"unsafe"
)

type QMutableLinkedListIterator struct {
	ptr unsafe.Pointer
}

type QMutableLinkedListIterator_ITF interface {
	QMutableLinkedListIterator_PTR() *QMutableLinkedListIterator
}

func (p *QMutableLinkedListIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableLinkedListIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableLinkedListIterator(ptr QMutableLinkedListIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableLinkedListIterator_PTR().Pointer()
	}
	return nil
}

func NewQMutableLinkedListIteratorFromPointer(ptr unsafe.Pointer) *QMutableLinkedListIterator {
	var n = new(QMutableLinkedListIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableLinkedListIterator) QMutableLinkedListIterator_PTR() *QMutableLinkedListIterator {
	return ptr
}
