package core

//#include "qweakpointer.h"
import "C"
import (
	"unsafe"
)

type QWeakPointer struct {
	ptr unsafe.Pointer
}

type QWeakPointer_ITF interface {
	QWeakPointer_PTR() *QWeakPointer
}

func (p *QWeakPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWeakPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWeakPointer(ptr QWeakPointer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWeakPointer_PTR().Pointer()
	}
	return nil
}

func NewQWeakPointerFromPointer(ptr unsafe.Pointer) *QWeakPointer {
	var n = new(QWeakPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWeakPointer) QWeakPointer_PTR() *QWeakPointer {
	return ptr
}
