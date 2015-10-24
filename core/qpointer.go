package core

//#include "qpointer.h"
import "C"
import (
	"unsafe"
)

type QPointer struct {
	ptr unsafe.Pointer
}

type QPointerITF interface {
	QPointerPTR() *QPointer
}

func (p *QPointer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPointer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPointer(ptr QPointerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPointerPTR().Pointer()
	}
	return nil
}

func QPointerFromPointer(ptr unsafe.Pointer) *QPointer {
	var n = new(QPointer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPointer) QPointerPTR() *QPointer {
	return ptr
}
