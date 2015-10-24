package core

//#include "qhash.h"
import "C"
import (
	"unsafe"
)

type QHash struct {
	ptr unsafe.Pointer
}

type QHashITF interface {
	QHashPTR() *QHash
}

func (p *QHash) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHash) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHash(ptr QHashITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHashPTR().Pointer()
	}
	return nil
}

func QHashFromPointer(ptr unsafe.Pointer) *QHash {
	var n = new(QHash)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHash) QHashPTR() *QHash {
	return ptr
}
