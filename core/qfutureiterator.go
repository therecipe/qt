package core

//#include "qfutureiterator.h"
import "C"
import (
	"unsafe"
)

type QFutureIterator struct {
	ptr unsafe.Pointer
}

type QFutureIteratorITF interface {
	QFutureIteratorPTR() *QFutureIterator
}

func (p *QFutureIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFutureIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFutureIterator(ptr QFutureIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFutureIteratorPTR().Pointer()
	}
	return nil
}

func QFutureIteratorFromPointer(ptr unsafe.Pointer) *QFutureIterator {
	var n = new(QFutureIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFutureIterator) QFutureIteratorPTR() *QFutureIterator {
	return ptr
}
