package core

//#include "qglobalstatic.h"
import "C"
import (
	"unsafe"
)

type QGlobalStatic struct {
	ptr unsafe.Pointer
}

type QGlobalStaticITF interface {
	QGlobalStaticPTR() *QGlobalStatic
}

func (p *QGlobalStatic) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGlobalStatic) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGlobalStatic(ptr QGlobalStaticITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGlobalStaticPTR().Pointer()
	}
	return nil
}

func QGlobalStaticFromPointer(ptr unsafe.Pointer) *QGlobalStatic {
	var n = new(QGlobalStatic)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGlobalStatic) QGlobalStaticPTR() *QGlobalStatic {
	return ptr
}
