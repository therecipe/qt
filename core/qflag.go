package core

//#include "qflag.h"
import "C"
import (
	"unsafe"
)

type QFlag struct {
	ptr unsafe.Pointer
}

type QFlagITF interface {
	QFlagPTR() *QFlag
}

func (p *QFlag) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFlag) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFlag(ptr QFlagITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFlagPTR().Pointer()
	}
	return nil
}

func QFlagFromPointer(ptr unsafe.Pointer) *QFlag {
	var n = new(QFlag)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFlag) QFlagPTR() *QFlag {
	return ptr
}

func NewQFlag(value int) *QFlag {
	return QFlagFromPointer(unsafe.Pointer(C.QFlag_NewQFlag(C.int(value))))
}
