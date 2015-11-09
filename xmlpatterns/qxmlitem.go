package xmlpatterns

//#include "qxmlitem.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QXmlItem struct {
	ptr unsafe.Pointer
}

type QXmlItem_ITF interface {
	QXmlItem_PTR() *QXmlItem
}

func (p *QXmlItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlItem(ptr QXmlItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlItem_PTR().Pointer()
	}
	return nil
}

func NewQXmlItemFromPointer(ptr unsafe.Pointer) *QXmlItem {
	var n = new(QXmlItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlItem) QXmlItem_PTR() *QXmlItem {
	return ptr
}

func NewQXmlItem() *QXmlItem {
	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem())
}

func NewQXmlItem4(atomicValue core.QVariant_ITF) *QXmlItem {
	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem4(core.PointerFromQVariant(atomicValue)))
}

func NewQXmlItem2(other QXmlItem_ITF) *QXmlItem {
	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem2(PointerFromQXmlItem(other)))
}

func NewQXmlItem3(node QXmlNodeModelIndex_ITF) *QXmlItem {
	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem3(PointerFromQXmlNodeModelIndex(node)))
}

func (ptr *QXmlItem) IsNode() bool {
	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) DestroyQXmlItem() {
	if ptr.Pointer() != nil {
		C.QXmlItem_DestroyQXmlItem(ptr.Pointer())
	}
}
