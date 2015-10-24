package core

//#include "qflags.h"
import "C"
import (
	"unsafe"
)

type QFlags struct {
	ptr unsafe.Pointer
}

type QFlagsITF interface {
	QFlagsPTR() *QFlags
}

func (p *QFlags) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFlags) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFlags(ptr QFlagsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFlagsPTR().Pointer()
	}
	return nil
}

func QFlagsFromPointer(ptr unsafe.Pointer) *QFlags {
	var n = new(QFlags)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFlags) QFlagsPTR() *QFlags {
	return ptr
}
