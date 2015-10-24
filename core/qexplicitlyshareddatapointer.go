package core

//#include "qexplicitlyshareddatapointer.h"
import "C"
import (
	"unsafe"
)

type QExplicitlySharedDataPointer struct {
	ptr unsafe.Pointer
}

type QExplicitlySharedDataPointerITF interface {
	QExplicitlySharedDataPointerPTR() *QExplicitlySharedDataPointer
}

func (p *QExplicitlySharedDataPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QExplicitlySharedDataPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQExplicitlySharedDataPointer(ptr QExplicitlySharedDataPointerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QExplicitlySharedDataPointerPTR().Pointer()
	}
	return nil
}

func QExplicitlySharedDataPointerFromPointer(ptr unsafe.Pointer) *QExplicitlySharedDataPointer {
	var n = new(QExplicitlySharedDataPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QExplicitlySharedDataPointer) QExplicitlySharedDataPointerPTR() *QExplicitlySharedDataPointer {
	return ptr
}
