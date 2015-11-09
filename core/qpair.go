package core

//#include "qpair.h"
import "C"
import (
	"unsafe"
)

type QPair struct {
	ptr unsafe.Pointer
}

type QPair_ITF interface {
	QPair_PTR() *QPair
}

func (p *QPair) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPair) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPair(ptr QPair_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPair_PTR().Pointer()
	}
	return nil
}

func NewQPairFromPointer(ptr unsafe.Pointer) *QPair {
	var n = new(QPair)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPair) QPair_PTR() *QPair {
	return ptr
}
