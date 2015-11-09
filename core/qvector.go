package core

//#include "qvector.h"
import "C"
import (
	"unsafe"
)

type QVector struct {
	ptr unsafe.Pointer
}

type QVector_ITF interface {
	QVector_PTR() *QVector
}

func (p *QVector) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVector) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVector(ptr QVector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVector_PTR().Pointer()
	}
	return nil
}

func NewQVectorFromPointer(ptr unsafe.Pointer) *QVector {
	var n = new(QVector)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVector) QVector_PTR() *QVector {
	return ptr
}
