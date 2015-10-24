package core

//#include "qvarlengtharray.h"
import "C"
import (
	"unsafe"
)

type QVarLengthArray struct {
	ptr unsafe.Pointer
}

type QVarLengthArrayITF interface {
	QVarLengthArrayPTR() *QVarLengthArray
}

func (p *QVarLengthArray) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVarLengthArray) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVarLengthArray(ptr QVarLengthArrayITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVarLengthArrayPTR().Pointer()
	}
	return nil
}

func QVarLengthArrayFromPointer(ptr unsafe.Pointer) *QVarLengthArray {
	var n = new(QVarLengthArray)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVarLengthArray) QVarLengthArrayPTR() *QVarLengthArray {
	return ptr
}
