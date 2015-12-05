package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QGenericMatrix struct {
	ptr unsafe.Pointer
}

type QGenericMatrix_ITF interface {
	QGenericMatrix_PTR() *QGenericMatrix
}

func (p *QGenericMatrix) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGenericMatrix) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGenericMatrix(ptr QGenericMatrix_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGenericMatrix_PTR().Pointer()
	}
	return nil
}

func NewQGenericMatrixFromPointer(ptr unsafe.Pointer) *QGenericMatrix {
	var n = new(QGenericMatrix)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGenericMatrix) QGenericMatrix_PTR() *QGenericMatrix {
	return ptr
}
