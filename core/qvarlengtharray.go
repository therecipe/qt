package core

//#include "qvarlengtharray.h"
import "C"
import (
	"unsafe"
)

type QVarLengthArray struct {
	ptr unsafe.Pointer
}

type QVarLengthArray_ITF interface {
	QVarLengthArray_PTR() *QVarLengthArray
}

func (p *QVarLengthArray) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVarLengthArray) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVarLengthArray(ptr QVarLengthArray_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVarLengthArray_PTR().Pointer()
	}
	return nil
}

func NewQVarLengthArrayFromPointer(ptr unsafe.Pointer) *QVarLengthArray {
	var n = new(QVarLengthArray)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVarLengthArray) QVarLengthArray_PTR() *QVarLengthArray {
	return ptr
}
