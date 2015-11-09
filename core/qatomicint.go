package core

//#include "qatomicint.h"
import "C"
import (
	"unsafe"
)

type QAtomicInt struct {
	QAtomicInteger
}

type QAtomicInt_ITF interface {
	QAtomicInteger_ITF
	QAtomicInt_PTR() *QAtomicInt
}

func PointerFromQAtomicInt(ptr QAtomicInt_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAtomicInt_PTR().Pointer()
	}
	return nil
}

func NewQAtomicIntFromPointer(ptr unsafe.Pointer) *QAtomicInt {
	var n = new(QAtomicInt)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAtomicInt) QAtomicInt_PTR() *QAtomicInt {
	return ptr
}
