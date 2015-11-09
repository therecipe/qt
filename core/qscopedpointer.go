package core

//#include "qscopedpointer.h"
import "C"
import (
	"unsafe"
)

type QScopedPointer struct {
	ptr unsafe.Pointer
}

type QScopedPointer_ITF interface {
	QScopedPointer_PTR() *QScopedPointer
}

func (p *QScopedPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScopedPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScopedPointer(ptr QScopedPointer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScopedPointer_PTR().Pointer()
	}
	return nil
}

func NewQScopedPointerFromPointer(ptr unsafe.Pointer) *QScopedPointer {
	var n = new(QScopedPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScopedPointer) QScopedPointer_PTR() *QScopedPointer {
	return ptr
}
