package core

//#include "qmutablemapiterator.h"
import "C"
import (
	"unsafe"
)

type QMutableMapIterator struct {
	ptr unsafe.Pointer
}

type QMutableMapIteratorITF interface {
	QMutableMapIteratorPTR() *QMutableMapIterator
}

func (p *QMutableMapIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableMapIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableMapIterator(ptr QMutableMapIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableMapIteratorPTR().Pointer()
	}
	return nil
}

func QMutableMapIteratorFromPointer(ptr unsafe.Pointer) *QMutableMapIterator {
	var n = new(QMutableMapIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableMapIterator) QMutableMapIteratorPTR() *QMutableMapIterator {
	return ptr
}
