package core

//#include "qscopedpointer.h"
import "C"
import (
	"unsafe"
)

type QScopedPointer struct {
	ptr unsafe.Pointer
}

type QScopedPointerITF interface {
	QScopedPointerPTR() *QScopedPointer
}

func (p *QScopedPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScopedPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScopedPointer(ptr QScopedPointerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScopedPointerPTR().Pointer()
	}
	return nil
}

func QScopedPointerFromPointer(ptr unsafe.Pointer) *QScopedPointer {
	var n = new(QScopedPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScopedPointer) QScopedPointerPTR() *QScopedPointer {
	return ptr
}
