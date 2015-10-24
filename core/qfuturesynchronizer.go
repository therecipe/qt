package core

//#include "qfuturesynchronizer.h"
import "C"
import (
	"unsafe"
)

type QFutureSynchronizer struct {
	ptr unsafe.Pointer
}

type QFutureSynchronizerITF interface {
	QFutureSynchronizerPTR() *QFutureSynchronizer
}

func (p *QFutureSynchronizer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFutureSynchronizer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFutureSynchronizer(ptr QFutureSynchronizerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFutureSynchronizerPTR().Pointer()
	}
	return nil
}

func QFutureSynchronizerFromPointer(ptr unsafe.Pointer) *QFutureSynchronizer {
	var n = new(QFutureSynchronizer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFutureSynchronizer) QFutureSynchronizerPTR() *QFutureSynchronizer {
	return ptr
}
