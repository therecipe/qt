package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QSharedPointer struct {
	ptr unsafe.Pointer
}

type QSharedPointer_ITF interface {
	QSharedPointer_PTR() *QSharedPointer
}

func (p *QSharedPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSharedPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSharedPointer(ptr QSharedPointer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedPointer_PTR().Pointer()
	}
	return nil
}

func NewQSharedPointerFromPointer(ptr unsafe.Pointer) *QSharedPointer {
	var n = new(QSharedPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSharedPointer) QSharedPointer_PTR() *QSharedPointer {
	return ptr
}
