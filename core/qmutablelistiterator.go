package core

//#include "qmutablelistiterator.h"
import "C"
import (
	"unsafe"
)

type QMutableListIterator struct {
	ptr unsafe.Pointer
}

type QMutableListIteratorITF interface {
	QMutableListIteratorPTR() *QMutableListIterator
}

func (p *QMutableListIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMutableListIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMutableListIterator(ptr QMutableListIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMutableListIteratorPTR().Pointer()
	}
	return nil
}

func QMutableListIteratorFromPointer(ptr unsafe.Pointer) *QMutableListIterator {
	var n = new(QMutableListIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMutableListIterator) QMutableListIteratorPTR() *QMutableListIterator {
	return ptr
}
