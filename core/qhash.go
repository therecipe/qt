package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QHash struct {
	ptr unsafe.Pointer
}

type QHash_ITF interface {
	QHash_PTR() *QHash
}

func (p *QHash) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHash) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHash(ptr QHash_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHash_PTR().Pointer()
	}
	return nil
}

func NewQHashFromPointer(ptr unsafe.Pointer) *QHash {
	var n = new(QHash)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHash) QHash_PTR() *QHash {
	return ptr
}
