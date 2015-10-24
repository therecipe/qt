package core

//#include "qatomicpointer.h"
import "C"
import (
	"unsafe"
)

type QAtomicPointer struct {
	ptr unsafe.Pointer
}

type QAtomicPointerITF interface {
	QAtomicPointerPTR() *QAtomicPointer
}

func (p *QAtomicPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAtomicPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAtomicPointer(ptr QAtomicPointerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAtomicPointerPTR().Pointer()
	}
	return nil
}

func QAtomicPointerFromPointer(ptr unsafe.Pointer) *QAtomicPointer {
	var n = new(QAtomicPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAtomicPointer) QAtomicPointerPTR() *QAtomicPointer {
	return ptr
}
