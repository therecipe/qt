package xmlpatterns

//#include "qxmlnamepool.h"
import "C"
import (
	"unsafe"
)

type QXmlNamePool struct {
	ptr unsafe.Pointer
}

type QXmlNamePoolITF interface {
	QXmlNamePoolPTR() *QXmlNamePool
}

func (p *QXmlNamePool) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlNamePool) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlNamePool(ptr QXmlNamePoolITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlNamePoolPTR().Pointer()
	}
	return nil
}

func QXmlNamePoolFromPointer(ptr unsafe.Pointer) *QXmlNamePool {
	var n = new(QXmlNamePool)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlNamePool) QXmlNamePoolPTR() *QXmlNamePool {
	return ptr
}

func NewQXmlNamePool() *QXmlNamePool {
	return QXmlNamePoolFromPointer(unsafe.Pointer(C.QXmlNamePool_NewQXmlNamePool()))
}

func NewQXmlNamePool2(other QXmlNamePoolITF) *QXmlNamePool {
	return QXmlNamePoolFromPointer(unsafe.Pointer(C.QXmlNamePool_NewQXmlNamePool2(C.QtObjectPtr(PointerFromQXmlNamePool(other)))))
}

func (ptr *QXmlNamePool) DestroyQXmlNamePool() {
	if ptr.Pointer() != nil {
		C.QXmlNamePool_DestroyQXmlNamePool(C.QtObjectPtr(ptr.Pointer()))
	}
}
