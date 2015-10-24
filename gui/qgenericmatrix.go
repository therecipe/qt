package gui

//#include "qgenericmatrix.h"
import "C"
import (
	"unsafe"
)

type QGenericMatrix struct {
	ptr unsafe.Pointer
}

type QGenericMatrixITF interface {
	QGenericMatrixPTR() *QGenericMatrix
}

func (p *QGenericMatrix) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGenericMatrix) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGenericMatrix(ptr QGenericMatrixITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGenericMatrixPTR().Pointer()
	}
	return nil
}

func QGenericMatrixFromPointer(ptr unsafe.Pointer) *QGenericMatrix {
	var n = new(QGenericMatrix)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGenericMatrix) QGenericMatrixPTR() *QGenericMatrix {
	return ptr
}
