package core

//#include "qmultihash.h"
import "C"
import (
	"unsafe"
)

type QMultiHash struct {
	QHash
}

type QMultiHashITF interface {
	QHashITF
	QMultiHashPTR() *QMultiHash
}

func PointerFromQMultiHash(ptr QMultiHashITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMultiHashPTR().Pointer()
	}
	return nil
}

func QMultiHashFromPointer(ptr unsafe.Pointer) *QMultiHash {
	var n = new(QMultiHash)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMultiHash) QMultiHashPTR() *QMultiHash {
	return ptr
}
