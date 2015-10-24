package core

//#include "qvectoriterator.h"
import "C"
import (
	"unsafe"
)

type QVectorIterator struct {
	ptr unsafe.Pointer
}

type QVectorIteratorITF interface {
	QVectorIteratorPTR() *QVectorIterator
}

func (p *QVectorIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVectorIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVectorIterator(ptr QVectorIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVectorIteratorPTR().Pointer()
	}
	return nil
}

func QVectorIteratorFromPointer(ptr unsafe.Pointer) *QVectorIterator {
	var n = new(QVectorIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVectorIterator) QVectorIteratorPTR() *QVectorIterator {
	return ptr
}
