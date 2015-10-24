package core

//#include "qtglobalstatic.h"
import "C"
import (
	"unsafe"
)

type QtGlobalStatic struct {
	ptr unsafe.Pointer
}

type QtGlobalStaticITF interface {
	QtGlobalStaticPTR() *QtGlobalStatic
}

func (p *QtGlobalStatic) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QtGlobalStatic) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQtGlobalStatic(ptr QtGlobalStaticITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtGlobalStaticPTR().Pointer()
	}
	return nil
}

func QtGlobalStaticFromPointer(ptr unsafe.Pointer) *QtGlobalStatic {
	var n = new(QtGlobalStatic)
	n.SetPointer(ptr)
	return n
}

func (ptr *QtGlobalStatic) QtGlobalStaticPTR() *QtGlobalStatic {
	return ptr
}

//QtGlobalStatic::GuardValues
type QtGlobalStatic__GuardValues int

var (
	QtGlobalStatic__Destroyed     = QtGlobalStatic__GuardValues(-2)
	QtGlobalStatic__Initialized   = QtGlobalStatic__GuardValues(-1)
	QtGlobalStatic__Uninitialized = QtGlobalStatic__GuardValues(0)
	QtGlobalStatic__Initializing  = QtGlobalStatic__GuardValues(1)
)
