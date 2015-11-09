package core

//#include "qshareddatapointer.h"
import "C"
import (
	"unsafe"
)

type QSharedDataPointer struct {
	ptr unsafe.Pointer
}

type QSharedDataPointer_ITF interface {
	QSharedDataPointer_PTR() *QSharedDataPointer
}

func (p *QSharedDataPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSharedDataPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSharedDataPointer(ptr QSharedDataPointer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedDataPointer_PTR().Pointer()
	}
	return nil
}

func NewQSharedDataPointerFromPointer(ptr unsafe.Pointer) *QSharedDataPointer {
	var n = new(QSharedDataPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSharedDataPointer) QSharedDataPointer_PTR() *QSharedDataPointer {
	return ptr
}
