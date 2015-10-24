package core

//#include "qdebug.h"
import "C"
import (
	"unsafe"
)

type QDebug struct {
	ptr unsafe.Pointer
}

type QDebugITF interface {
	QDebugPTR() *QDebug
}

func (p *QDebug) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDebug) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDebug(ptr QDebugITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDebugPTR().Pointer()
	}
	return nil
}

func QDebugFromPointer(ptr unsafe.Pointer) *QDebug {
	var n = new(QDebug)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDebug) QDebugPTR() *QDebug {
	return ptr
}
