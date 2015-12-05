package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QScopedValueRollback struct {
	ptr unsafe.Pointer
}

type QScopedValueRollback_ITF interface {
	QScopedValueRollback_PTR() *QScopedValueRollback
}

func (p *QScopedValueRollback) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScopedValueRollback) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScopedValueRollback(ptr QScopedValueRollback_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScopedValueRollback_PTR().Pointer()
	}
	return nil
}

func NewQScopedValueRollbackFromPointer(ptr unsafe.Pointer) *QScopedValueRollback {
	var n = new(QScopedValueRollback)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScopedValueRollback) QScopedValueRollback_PTR() *QScopedValueRollback {
	return ptr
}
