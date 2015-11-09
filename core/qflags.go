package core

//#include "qflags.h"
import "C"
import (
	"unsafe"
)

type QFlags struct {
	ptr unsafe.Pointer
}

type QFlags_ITF interface {
	QFlags_PTR() *QFlags
}

func (p *QFlags) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFlags) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFlags(ptr QFlags_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFlags_PTR().Pointer()
	}
	return nil
}

func NewQFlagsFromPointer(ptr unsafe.Pointer) *QFlags {
	var n = new(QFlags)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFlags) QFlags_PTR() *QFlags {
	return ptr
}
