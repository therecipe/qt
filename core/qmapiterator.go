package core

//#include "qmapiterator.h"
import "C"
import (
	"unsafe"
)

type QMapIterator struct {
	ptr unsafe.Pointer
}

type QMapIteratorITF interface {
	QMapIteratorPTR() *QMapIterator
}

func (p *QMapIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMapIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMapIterator(ptr QMapIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMapIteratorPTR().Pointer()
	}
	return nil
}

func QMapIteratorFromPointer(ptr unsafe.Pointer) *QMapIterator {
	var n = new(QMapIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMapIterator) QMapIteratorPTR() *QMapIterator {
	return ptr
}
