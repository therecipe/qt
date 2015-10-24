package core

//#include "qatomicint.h"
import "C"
import (
	"unsafe"
)

type QAtomicInt struct {
	QAtomicInteger
}

type QAtomicIntITF interface {
	QAtomicIntegerITF
	QAtomicIntPTR() *QAtomicInt
}

func PointerFromQAtomicInt(ptr QAtomicIntITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAtomicIntPTR().Pointer()
	}
	return nil
}

func QAtomicIntFromPointer(ptr unsafe.Pointer) *QAtomicInt {
	var n = new(QAtomicInt)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAtomicInt) QAtomicIntPTR() *QAtomicInt {
	return ptr
}
