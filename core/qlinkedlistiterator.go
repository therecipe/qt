package core

//#include "qlinkedlistiterator.h"
import "C"
import (
	"unsafe"
)

type QLinkedListIterator struct {
	ptr unsafe.Pointer
}

type QLinkedListIteratorITF interface {
	QLinkedListIteratorPTR() *QLinkedListIterator
}

func (p *QLinkedListIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLinkedListIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLinkedListIterator(ptr QLinkedListIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLinkedListIteratorPTR().Pointer()
	}
	return nil
}

func QLinkedListIteratorFromPointer(ptr unsafe.Pointer) *QLinkedListIterator {
	var n = new(QLinkedListIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLinkedListIterator) QLinkedListIteratorPTR() *QLinkedListIterator {
	return ptr
}
