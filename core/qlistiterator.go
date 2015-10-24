package core

//#include "qlistiterator.h"
import "C"
import (
	"unsafe"
)

type QListIterator struct {
	ptr unsafe.Pointer
}

type QListIteratorITF interface {
	QListIteratorPTR() *QListIterator
}

func (p *QListIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QListIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQListIterator(ptr QListIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QListIteratorPTR().Pointer()
	}
	return nil
}

func QListIteratorFromPointer(ptr unsafe.Pointer) *QListIterator {
	var n = new(QListIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QListIterator) QListIteratorPTR() *QListIterator {
	return ptr
}
