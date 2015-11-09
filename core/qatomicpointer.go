package core

//#include "qatomicpointer.h"
import "C"
import (
	"unsafe"
)

type QAtomicPointer struct {
	ptr unsafe.Pointer
}

type QAtomicPointer_ITF interface {
	QAtomicPointer_PTR() *QAtomicPointer
}

func (p *QAtomicPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAtomicPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAtomicPointer(ptr QAtomicPointer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAtomicPointer_PTR().Pointer()
	}
	return nil
}

func NewQAtomicPointerFromPointer(ptr unsafe.Pointer) *QAtomicPointer {
	var n = new(QAtomicPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAtomicPointer) QAtomicPointer_PTR() *QAtomicPointer {
	return ptr
}
