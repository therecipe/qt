package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QExplicitlySharedDataPointer struct {
	ptr unsafe.Pointer
}

type QExplicitlySharedDataPointer_ITF interface {
	QExplicitlySharedDataPointer_PTR() *QExplicitlySharedDataPointer
}

func (p *QExplicitlySharedDataPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QExplicitlySharedDataPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQExplicitlySharedDataPointer(ptr QExplicitlySharedDataPointer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QExplicitlySharedDataPointer_PTR().Pointer()
	}
	return nil
}

func NewQExplicitlySharedDataPointerFromPointer(ptr unsafe.Pointer) *QExplicitlySharedDataPointer {
	var n = new(QExplicitlySharedDataPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QExplicitlySharedDataPointer) QExplicitlySharedDataPointer_PTR() *QExplicitlySharedDataPointer {
	return ptr
}
