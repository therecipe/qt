package core

//#include "qmultihash.h"
import "C"
import (
	"unsafe"
)

type QMultiHash struct {
	QHash
}

type QMultiHash_ITF interface {
	QHash_ITF
	QMultiHash_PTR() *QMultiHash
}

func PointerFromQMultiHash(ptr QMultiHash_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMultiHash_PTR().Pointer()
	}
	return nil
}

func NewQMultiHashFromPointer(ptr unsafe.Pointer) *QMultiHash {
	var n = new(QMultiHash)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMultiHash) QMultiHash_PTR() *QMultiHash {
	return ptr
}
