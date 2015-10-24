package core

//#include "qpair.h"
import "C"
import (
	"unsafe"
)

type QPair struct {
	ptr unsafe.Pointer
}

type QPairITF interface {
	QPairPTR() *QPair
}

func (p *QPair) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPair) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPair(ptr QPairITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPairPTR().Pointer()
	}
	return nil
}

func QPairFromPointer(ptr unsafe.Pointer) *QPair {
	var n = new(QPair)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPair) QPairPTR() *QPair {
	return ptr
}
