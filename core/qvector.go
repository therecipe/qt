package core

//#include "qvector.h"
import "C"
import (
	"unsafe"
)

type QVector struct {
	ptr unsafe.Pointer
}

type QVectorITF interface {
	QVectorPTR() *QVector
}

func (p *QVector) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVector) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVector(ptr QVectorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVectorPTR().Pointer()
	}
	return nil
}

func QVectorFromPointer(ptr unsafe.Pointer) *QVector {
	var n = new(QVector)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVector) QVectorPTR() *QVector {
	return ptr
}
