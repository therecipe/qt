package core

//#include "qflag.h"
import "C"
import (
	"unsafe"
)

type QFlag struct {
	ptr unsafe.Pointer
}

type QFlag_ITF interface {
	QFlag_PTR() *QFlag
}

func (p *QFlag) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFlag) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFlag(ptr QFlag_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFlag_PTR().Pointer()
	}
	return nil
}

func NewQFlagFromPointer(ptr unsafe.Pointer) *QFlag {
	var n = new(QFlag)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFlag) QFlag_PTR() *QFlag {
	return ptr
}

func NewQFlag(value int) *QFlag {
	return NewQFlagFromPointer(C.QFlag_NewQFlag(C.int(value)))
}
