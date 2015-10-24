package core

//#include "qmutablevectoriterator.h"
import "C"
import (
	"unsafe"
)

type QMutableVectorIterator struct {
	ptr unsafe.Pointer
}

type QMutableVectorIteratorITF interface {
	QMutableVectorIteratorPTR() *QMutableVectorIterator
}

func (p *QMutableVectorIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableVectorIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableVectorIterator(ptr QMutableVectorIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableVectorIteratorPTR().Pointer()
	}
	return nil
}

func QMutableVectorIteratorFromPointer(ptr unsafe.Pointer) *QMutableVectorIterator {
	var n = new(QMutableVectorIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableVectorIterator) QMutableVectorIteratorPTR() *QMutableVectorIterator {
	return ptr
}
