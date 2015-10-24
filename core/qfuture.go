package core

//#include "qfuture.h"
import "C"
import (
	"unsafe"
)

type QFuture struct {
	ptr unsafe.Pointer
}

type QFutureITF interface {
	QFuturePTR() *QFuture
}

func (p *QFuture) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFuture) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFuture(ptr QFutureITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFuturePTR().Pointer()
	}
	return nil
}

func QFutureFromPointer(ptr unsafe.Pointer) *QFuture {
	var n = new(QFuture)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFuture) QFuturePTR() *QFuture {
	return ptr
}
