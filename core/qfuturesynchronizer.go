package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QFutureSynchronizer struct {
	ptr unsafe.Pointer
}

type QFutureSynchronizer_ITF interface {
	QFutureSynchronizer_PTR() *QFutureSynchronizer
}

func (p *QFutureSynchronizer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFutureSynchronizer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFutureSynchronizer(ptr QFutureSynchronizer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFutureSynchronizer_PTR().Pointer()
	}
	return nil
}

func NewQFutureSynchronizerFromPointer(ptr unsafe.Pointer) *QFutureSynchronizer {
	var n = new(QFutureSynchronizer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFutureSynchronizer) QFutureSynchronizer_PTR() *QFutureSynchronizer {
	return ptr
}
