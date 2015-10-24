package core

//#include "qscopedvaluerollback.h"
import "C"
import (
	"unsafe"
)

type QScopedValueRollback struct {
	ptr unsafe.Pointer
}

type QScopedValueRollbackITF interface {
	QScopedValueRollbackPTR() *QScopedValueRollback
}

func (p *QScopedValueRollback) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScopedValueRollback) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScopedValueRollback(ptr QScopedValueRollbackITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScopedValueRollbackPTR().Pointer()
	}
	return nil
}

func QScopedValueRollbackFromPointer(ptr unsafe.Pointer) *QScopedValueRollback {
	var n = new(QScopedValueRollback)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScopedValueRollback) QScopedValueRollbackPTR() *QScopedValueRollback {
	return ptr
}
