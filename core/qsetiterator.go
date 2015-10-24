package core

//#include "qsetiterator.h"
import "C"
import (
	"unsafe"
)

type QSetIterator struct {
	ptr unsafe.Pointer
}

type QSetIteratorITF interface {
	QSetIteratorPTR() *QSetIterator
}

func (p *QSetIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSetIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSetIterator(ptr QSetIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSetIteratorPTR().Pointer()
	}
	return nil
}

func QSetIteratorFromPointer(ptr unsafe.Pointer) *QSetIterator {
	var n = new(QSetIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSetIterator) QSetIteratorPTR() *QSetIterator {
	return ptr
}
