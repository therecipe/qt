package xmlpatterns

//#include "qxmlnamepool.h"
import "C"
import (
	"unsafe"
)

type QXmlNamePool struct {
	ptr unsafe.Pointer
}

type QXmlNamePool_ITF interface {
	QXmlNamePool_PTR() *QXmlNamePool
}

func (p *QXmlNamePool) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlNamePool) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlNamePool(ptr QXmlNamePool_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamePool_PTR().Pointer()
	}
	return nil
}

func NewQXmlNamePoolFromPointer(ptr unsafe.Pointer) *QXmlNamePool {
	var n = new(QXmlNamePool)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlNamePool) QXmlNamePool_PTR() *QXmlNamePool {
	return ptr
}

func NewQXmlNamePool() *QXmlNamePool {
	return NewQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool())
}

func NewQXmlNamePool2(other QXmlNamePool_ITF) *QXmlNamePool {
	return NewQXmlNamePoolFromPointer(C.QXmlNamePool_NewQXmlNamePool2(PointerFromQXmlNamePool(other)))
}

func (ptr *QXmlNamePool) DestroyQXmlNamePool() {
	if ptr.Pointer() != nil {
		C.QXmlNamePool_DestroyQXmlNamePool(ptr.Pointer())
	}
}
