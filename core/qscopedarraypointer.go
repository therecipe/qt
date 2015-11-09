package core

//#include "qscopedarraypointer.h"
import "C"
import (
	"unsafe"
)

type QScopedArrayPointer struct {
	QScopedPointer
}

type QScopedArrayPointer_ITF interface {
	QScopedPointer_ITF
	QScopedArrayPointer_PTR() *QScopedArrayPointer
}

func PointerFromQScopedArrayPointer(ptr QScopedArrayPointer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScopedArrayPointer_PTR().Pointer()
	}
	return nil
}

func NewQScopedArrayPointerFromPointer(ptr unsafe.Pointer) *QScopedArrayPointer {
	var n = new(QScopedArrayPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScopedArrayPointer) QScopedArrayPointer_PTR() *QScopedArrayPointer {
	return ptr
}
