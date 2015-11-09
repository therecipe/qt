package core

//#include "qfuture.h"
import "C"
import (
	"unsafe"
)

type QFuture struct {
	ptr unsafe.Pointer
}

type QFuture_ITF interface {
	QFuture_PTR() *QFuture
}

func (p *QFuture) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFuture) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFuture(ptr QFuture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFuture_PTR().Pointer()
	}
	return nil
}

func NewQFutureFromPointer(ptr unsafe.Pointer) *QFuture {
	var n = new(QFuture)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFuture) QFuture_PTR() *QFuture {
	return ptr
}
