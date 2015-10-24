package core

//#include "qshareddatapointer.h"
import "C"
import (
	"unsafe"
)

type QSharedDataPointer struct {
	ptr unsafe.Pointer
}

type QSharedDataPointerITF interface {
	QSharedDataPointerPTR() *QSharedDataPointer
}

func (p *QSharedDataPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSharedDataPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSharedDataPointer(ptr QSharedDataPointerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedDataPointerPTR().Pointer()
	}
	return nil
}

func QSharedDataPointerFromPointer(ptr unsafe.Pointer) *QSharedDataPointer {
	var n = new(QSharedDataPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSharedDataPointer) QSharedDataPointerPTR() *QSharedDataPointer {
	return ptr
}
