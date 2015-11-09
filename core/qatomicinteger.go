package core

//#include "qatomicinteger.h"
import "C"
import (
	"unsafe"
)

type QAtomicInteger struct {
	ptr unsafe.Pointer
}

type QAtomicInteger_ITF interface {
	QAtomicInteger_PTR() *QAtomicInteger
}

func (p *QAtomicInteger) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAtomicInteger) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAtomicInteger(ptr QAtomicInteger_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAtomicInteger_PTR().Pointer()
	}
	return nil
}

func NewQAtomicIntegerFromPointer(ptr unsafe.Pointer) *QAtomicInteger {
	var n = new(QAtomicInteger)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAtomicInteger) QAtomicInteger_PTR() *QAtomicInteger {
	return ptr
}
