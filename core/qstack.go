package core

//#include "qstack.h"
import "C"
import (
	"unsafe"
)

type QStack struct {
	QVector
}

type QStackITF interface {
	QVectorITF
	QStackPTR() *QStack
}

func PointerFromQStack(ptr QStackITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStackPTR().Pointer()
	}
	return nil
}

func QStackFromPointer(ptr unsafe.Pointer) *QStack {
	var n = new(QStack)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStack) QStackPTR() *QStack {
	return ptr
}
