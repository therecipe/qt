package core

//#include "qmutablesetiterator.h"
import "C"
import (
	"unsafe"
)

type QMutableSetIterator struct {
	ptr unsafe.Pointer
}

type QMutableSetIteratorITF interface {
	QMutableSetIteratorPTR() *QMutableSetIterator
}

func (p *QMutableSetIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableSetIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableSetIterator(ptr QMutableSetIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableSetIteratorPTR().Pointer()
	}
	return nil
}

func QMutableSetIteratorFromPointer(ptr unsafe.Pointer) *QMutableSetIterator {
	var n = new(QMutableSetIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableSetIterator) QMutableSetIteratorPTR() *QMutableSetIterator {
	return ptr
}
