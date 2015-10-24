package core

//#include "qscopedarraypointer.h"
import "C"
import (
	"unsafe"
)

type QScopedArrayPointer struct {
	QScopedPointer
}

type QScopedArrayPointerITF interface {
	QScopedPointerITF
	QScopedArrayPointerPTR() *QScopedArrayPointer
}

func PointerFromQScopedArrayPointer(ptr QScopedArrayPointerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScopedArrayPointerPTR().Pointer()
	}
	return nil
}

func QScopedArrayPointerFromPointer(ptr unsafe.Pointer) *QScopedArrayPointer {
	var n = new(QScopedArrayPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScopedArrayPointer) QScopedArrayPointerPTR() *QScopedArrayPointer {
	return ptr
}
