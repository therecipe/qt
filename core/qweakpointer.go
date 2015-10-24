package core

//#include "qweakpointer.h"
import "C"
import (
	"unsafe"
)

type QWeakPointer struct {
	ptr unsafe.Pointer
}

type QWeakPointerITF interface {
	QWeakPointerPTR() *QWeakPointer
}

func (p *QWeakPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QWeakPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQWeakPointer(ptr QWeakPointerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWeakPointerPTR().Pointer()
	}
	return nil
}

func QWeakPointerFromPointer(ptr unsafe.Pointer) *QWeakPointer {
	var n = new(QWeakPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWeakPointer) QWeakPointerPTR() *QWeakPointer {
	return ptr
}
