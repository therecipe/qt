package core

//#include "qmutablelinkedlistiterator.h"
import "C"
import (
	"unsafe"
)

type QMutableLinkedListIterator struct {
	ptr unsafe.Pointer
}

type QMutableLinkedListIteratorITF interface {
	QMutableLinkedListIteratorPTR() *QMutableLinkedListIterator
}

func (p *QMutableLinkedListIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableLinkedListIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableLinkedListIterator(ptr QMutableLinkedListIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableLinkedListIteratorPTR().Pointer()
	}
	return nil
}

func QMutableLinkedListIteratorFromPointer(ptr unsafe.Pointer) *QMutableLinkedListIterator {
	var n = new(QMutableLinkedListIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableLinkedListIterator) QMutableLinkedListIteratorPTR() *QMutableLinkedListIterator {
	return ptr
}
