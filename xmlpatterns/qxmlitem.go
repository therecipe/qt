package xmlpatterns

//#include "qxmlitem.h"
import "C"
import (
	"unsafe"
)

type QXmlItem struct {
	ptr unsafe.Pointer
}

type QXmlItemITF interface {
	QXmlItemPTR() *QXmlItem
}

func (p *QXmlItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlItem(ptr QXmlItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlItemPTR().Pointer()
	}
	return nil
}

func QXmlItemFromPointer(ptr unsafe.Pointer) *QXmlItem {
	var n = new(QXmlItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlItem) QXmlItemPTR() *QXmlItem {
	return ptr
}

func NewQXmlItem() *QXmlItem {
	return QXmlItemFromPointer(unsafe.Pointer(C.QXmlItem_NewQXmlItem()))
}

func NewQXmlItem4(atomicValue string) *QXmlItem {
	return QXmlItemFromPointer(unsafe.Pointer(C.QXmlItem_NewQXmlItem4(C.CString(atomicValue))))
}

func NewQXmlItem2(other QXmlItemITF) *QXmlItem {
	return QXmlItemFromPointer(unsafe.Pointer(C.QXmlItem_NewQXmlItem2(C.QtObjectPtr(PointerFromQXmlItem(other)))))
}

func NewQXmlItem3(node QXmlNodeModelIndexITF) *QXmlItem {
	return QXmlItemFromPointer(unsafe.Pointer(C.QXmlItem_NewQXmlItem3(C.QtObjectPtr(PointerFromQXmlNodeModelIndex(node)))))
}

func (ptr *QXmlItem) IsNode() bool {
	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlItem) DestroyQXmlItem() {
	if ptr.Pointer() != nil {
		C.QXmlItem_DestroyQXmlItem(C.QtObjectPtr(ptr.Pointer()))
	}
}
