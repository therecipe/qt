package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QXmlItem::QXmlItem")

	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem())
}

func NewQXmlItem4(atomicValue core.QVariant_ITF) *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem4(core.PointerFromQVariant(atomicValue)))
}

func NewQXmlItem2(other QXmlItem_ITF) *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem2(PointerFromQXmlItem(other)))
}

func NewQXmlItem3(node QXmlNodeModelIndex_ITF) *QXmlItem {
	defer qt.Recovering("QXmlItem::QXmlItem")

	return NewQXmlItemFromPointer(C.QXmlItem_NewQXmlItem3(PointerFromQXmlNodeModelIndex(node)))
}

func (ptr *QXmlItem) IsNode() bool {
	defer qt.Recovering("QXmlItem::isNode")

	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) IsNull() bool {
	defer qt.Recovering("QXmlItem::isNull")

	if ptr.Pointer() != nil {
		return C.QXmlItem_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlItem) DestroyQXmlItem() {
	defer qt.Recovering("QXmlItem::~QXmlItem")

	if ptr.Pointer() != nil {
		C.QXmlItem_DestroyQXmlItem(ptr.Pointer())
	}
}
