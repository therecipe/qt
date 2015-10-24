package core

//#include "qmutablehashiterator.h"
import "C"
import (
	"unsafe"
)

type QMutableHashIterator struct {
	ptr unsafe.Pointer
}

type QMutableHashIteratorITF interface {
	QMutableHashIteratorPTR() *QMutableHashIterator
}

func (p *QMutableHashIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableHashIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableHashIterator(ptr QMutableHashIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableHashIteratorPTR().Pointer()
	}
	return nil
}

func QMutableHashIteratorFromPointer(ptr unsafe.Pointer) *QMutableHashIterator {
	var n = new(QMutableHashIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableHashIterator) QMutableHashIteratorPTR() *QMutableHashIterator {
	return ptr
}
