package core

//#include "qtmetatypeprivate.h"
import "C"
import (
	"unsafe"
)

type QtMetaTypePrivate struct {
	ptr unsafe.Pointer
}

type QtMetaTypePrivateITF interface {
	QtMetaTypePrivatePTR() *QtMetaTypePrivate
}

func (p *QtMetaTypePrivate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QtMetaTypePrivate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQtMetaTypePrivate(ptr QtMetaTypePrivateITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtMetaTypePrivatePTR().Pointer()
	}
	return nil
}

func QtMetaTypePrivateFromPointer(ptr unsafe.Pointer) *QtMetaTypePrivate {
	var n = new(QtMetaTypePrivate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QtMetaTypePrivate) QtMetaTypePrivatePTR() *QtMetaTypePrivate {
	return ptr
}

//QtMetaTypePrivate::IteratorCapability
type QtMetaTypePrivate__IteratorCapability int

var (
	QtMetaTypePrivate__ForwardCapability       = QtMetaTypePrivate__IteratorCapability(1)
	QtMetaTypePrivate__BiDirectionalCapability = QtMetaTypePrivate__IteratorCapability(2)
	QtMetaTypePrivate__RandomAccessCapability  = QtMetaTypePrivate__IteratorCapability(4)
)
