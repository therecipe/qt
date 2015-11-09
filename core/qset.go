package core

//#include "qset.h"
import "C"
import (
	"unsafe"
)

type QSet struct {
	ptr unsafe.Pointer
}

type QSet_ITF interface {
	QSet_PTR() *QSet
}

func (p *QSet) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSet) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSet(ptr QSet_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSet_PTR().Pointer()
	}
	return nil
}

func NewQSetFromPointer(ptr unsafe.Pointer) *QSet {
	var n = new(QSet)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSet) QSet_PTR() *QSet {
	return ptr
}
