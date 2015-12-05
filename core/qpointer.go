package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QPointer struct {
	ptr unsafe.Pointer
}

type QPointer_ITF interface {
	QPointer_PTR() *QPointer
}

func (p *QPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPointer(ptr QPointer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPointer_PTR().Pointer()
	}
	return nil
}

func NewQPointerFromPointer(ptr unsafe.Pointer) *QPointer {
	var n = new(QPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPointer) QPointer_PTR() *QPointer {
	return ptr
}
