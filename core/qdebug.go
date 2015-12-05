package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QDebug struct {
	ptr unsafe.Pointer
}

type QDebug_ITF interface {
	QDebug_PTR() *QDebug
}

func (p *QDebug) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDebug) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDebug(ptr QDebug_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDebug_PTR().Pointer()
	}
	return nil
}

func NewQDebugFromPointer(ptr unsafe.Pointer) *QDebug {
	var n = new(QDebug)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDebug) QDebug_PTR() *QDebug {
	return ptr
}
