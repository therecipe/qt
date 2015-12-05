package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QtGlobalStatic struct {
	ptr unsafe.Pointer
}

type QtGlobalStatic_ITF interface {
	QtGlobalStatic_PTR() *QtGlobalStatic
}

func (p *QtGlobalStatic) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QtGlobalStatic) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQtGlobalStatic(ptr QtGlobalStatic_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtGlobalStatic_PTR().Pointer()
	}
	return nil
}

func NewQtGlobalStaticFromPointer(ptr unsafe.Pointer) *QtGlobalStatic {
	var n = new(QtGlobalStatic)
	n.SetPointer(ptr)
	return n
}

func (ptr *QtGlobalStatic) QtGlobalStatic_PTR() *QtGlobalStatic {
	return ptr
}

//QtGlobalStatic::GuardValues
type QtGlobalStatic__GuardValues int64

const (
	QtGlobalStatic__Destroyed     = QtGlobalStatic__GuardValues(-2)
	QtGlobalStatic__Initialized   = QtGlobalStatic__GuardValues(-1)
	QtGlobalStatic__Uninitialized = QtGlobalStatic__GuardValues(0)
	QtGlobalStatic__Initializing  = QtGlobalStatic__GuardValues(1)
)
