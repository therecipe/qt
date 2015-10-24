package xmlpatterns

//#include "qxmlresultitems.h"
import "C"
import (
	"unsafe"
)

type QXmlResultItems struct {
	ptr unsafe.Pointer
}

type QXmlResultItemsITF interface {
	QXmlResultItemsPTR() *QXmlResultItems
}

func (p *QXmlResultItems) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlResultItems) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlResultItems(ptr QXmlResultItemsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlResultItemsPTR().Pointer()
	}
	return nil
}

func QXmlResultItemsFromPointer(ptr unsafe.Pointer) *QXmlResultItems {
	var n = new(QXmlResultItems)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlResultItems) QXmlResultItemsPTR() *QXmlResultItems {
	return ptr
}

func NewQXmlResultItems() *QXmlResultItems {
	return QXmlResultItemsFromPointer(unsafe.Pointer(C.QXmlResultItems_NewQXmlResultItems()))
}

func (ptr *QXmlResultItems) HasError() bool {
	if ptr.Pointer() != nil {
		return C.QXmlResultItems_HasError(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlResultItems) DestroyQXmlResultItems() {
	if ptr.Pointer() != nil {
		C.QXmlResultItems_DestroyQXmlResultItems(C.QtObjectPtr(ptr.Pointer()))
	}
}
