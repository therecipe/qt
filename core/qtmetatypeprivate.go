package core

//#include "qtmetatypeprivate.h"
import "C"
import (
	"unsafe"
)

type QtMetaTypePrivate struct {
	ptr unsafe.Pointer
}

type QtMetaTypePrivate_ITF interface {
	QtMetaTypePrivate_PTR() *QtMetaTypePrivate
}

func (p *QtMetaTypePrivate) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QtMetaTypePrivate) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQtMetaTypePrivate(ptr QtMetaTypePrivate_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtMetaTypePrivate_PTR().Pointer()
	}
	return nil
}

func NewQtMetaTypePrivateFromPointer(ptr unsafe.Pointer) *QtMetaTypePrivate {
	var n = new(QtMetaTypePrivate)
	n.SetPointer(ptr)
	return n
}

func (ptr *QtMetaTypePrivate) QtMetaTypePrivate_PTR() *QtMetaTypePrivate {
	return ptr
}

//QtMetaTypePrivate::IteratorCapability
type QtMetaTypePrivate__IteratorCapability int64

const (
	QtMetaTypePrivate__ForwardCapability       = QtMetaTypePrivate__IteratorCapability(1)
	QtMetaTypePrivate__BiDirectionalCapability = QtMetaTypePrivate__IteratorCapability(2)
	QtMetaTypePrivate__RandomAccessCapability  = QtMetaTypePrivate__IteratorCapability(4)
)
