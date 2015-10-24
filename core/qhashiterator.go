package core

//#include "qhashiterator.h"
import "C"
import (
	"unsafe"
)

type QHashIterator struct {
	ptr unsafe.Pointer
}

type QHashIteratorITF interface {
	QHashIteratorPTR() *QHashIterator
}

func (p *QHashIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHashIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHashIterator(ptr QHashIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHashIteratorPTR().Pointer()
	}
	return nil
}

func QHashIteratorFromPointer(ptr unsafe.Pointer) *QHashIterator {
	var n = new(QHashIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHashIterator) QHashIteratorPTR() *QHashIterator {
	return ptr
}
