package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QGlobalStatic struct {
	ptr unsafe.Pointer
}

type QGlobalStatic_ITF interface {
	QGlobalStatic_PTR() *QGlobalStatic
}

func (p *QGlobalStatic) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGlobalStatic) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGlobalStatic(ptr QGlobalStatic_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGlobalStatic_PTR().Pointer()
	}
	return nil
}

func NewQGlobalStaticFromPointer(ptr unsafe.Pointer) *QGlobalStatic {
	var n = new(QGlobalStatic)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGlobalStatic) QGlobalStatic_PTR() *QGlobalStatic {
	return ptr
}
