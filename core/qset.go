package core

//#include "qset.h"
import "C"
import (
	"unsafe"
)

type QSet struct {
	ptr unsafe.Pointer
}

type QSetITF interface {
	QSetPTR() *QSet
}

func (p *QSet) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSet) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSet(ptr QSetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSetPTR().Pointer()
	}
	return nil
}

func QSetFromPointer(ptr unsafe.Pointer) *QSet {
	var n = new(QSet)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSet) QSetPTR() *QSet {
	return ptr
}
