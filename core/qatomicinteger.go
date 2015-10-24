package core

//#include "qatomicinteger.h"
import "C"
import (
	"unsafe"
)

type QAtomicInteger struct {
	ptr unsafe.Pointer
}

type QAtomicIntegerITF interface {
	QAtomicIntegerPTR() *QAtomicInteger
}

func (p *QAtomicInteger) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAtomicInteger) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAtomicInteger(ptr QAtomicIntegerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAtomicIntegerPTR().Pointer()
	}
	return nil
}

func QAtomicIntegerFromPointer(ptr unsafe.Pointer) *QAtomicInteger {
	var n = new(QAtomicInteger)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAtomicInteger) QAtomicIntegerPTR() *QAtomicInteger {
	return ptr
}
