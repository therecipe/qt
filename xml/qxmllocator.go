package xml

//#include "qxmllocator.h"
import "C"
import (
	"unsafe"
)

type QXmlLocator struct {
	ptr unsafe.Pointer
}

type QXmlLocatorITF interface {
	QXmlLocatorPTR() *QXmlLocator
}

func (p *QXmlLocator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlLocator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlLocator(ptr QXmlLocatorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlLocatorPTR().Pointer()
	}
	return nil
}

func QXmlLocatorFromPointer(ptr unsafe.Pointer) *QXmlLocator {
	var n = new(QXmlLocator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlLocator) QXmlLocatorPTR() *QXmlLocator {
	return ptr
}

func (ptr *QXmlLocator) ColumnNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_ColumnNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlLocator) LineNumber() int {
	if ptr.Pointer() != nil {
		return int(C.QXmlLocator_LineNumber(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QXmlLocator) DestroyQXmlLocator() {
	if ptr.Pointer() != nil {
		C.QXmlLocator_DestroyQXmlLocator(C.QtObjectPtr(ptr.Pointer()))
	}
}
