package core

//#include "qstack.h"
import "C"
import (
	"unsafe"
)

type QStack struct {
	QVector
}

type QStack_ITF interface {
	QVector_ITF
	QStack_PTR() *QStack
}

func PointerFromQStack(ptr QStack_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStack_PTR().Pointer()
	}
	return nil
}

func NewQStackFromPointer(ptr unsafe.Pointer) *QStack {
	var n = new(QStack)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStack) QStack_PTR() *QStack {
	return ptr
}
