package widgets

//#include "qstyleoptiongraphicsitem.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionGraphicsItem struct {
	QStyleOption
}

type QStyleOptionGraphicsItemITF interface {
	QStyleOptionITF
	QStyleOptionGraphicsItemPTR() *QStyleOptionGraphicsItem
}

func PointerFromQStyleOptionGraphicsItem(ptr QStyleOptionGraphicsItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionGraphicsItemPTR().Pointer()
	}
	return nil
}

func QStyleOptionGraphicsItemFromPointer(ptr unsafe.Pointer) *QStyleOptionGraphicsItem {
	var n = new(QStyleOptionGraphicsItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionGraphicsItem) QStyleOptionGraphicsItemPTR() *QStyleOptionGraphicsItem {
	return ptr
}

//QStyleOptionGraphicsItem::StyleOptionType
type QStyleOptionGraphicsItem__StyleOptionType int

var (
	QStyleOptionGraphicsItem__Type = QStyleOptionGraphicsItem__StyleOptionType(QStyleOption__SO_GraphicsItem)
)

//QStyleOptionGraphicsItem::StyleOptionVersion
type QStyleOptionGraphicsItem__StyleOptionVersion int

var (
	QStyleOptionGraphicsItem__Version = QStyleOptionGraphicsItem__StyleOptionVersion(1)
)

func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {
	return QStyleOptionGraphicsItemFromPointer(unsafe.Pointer(C.QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem()))
}

func NewQStyleOptionGraphicsItem2(other QStyleOptionGraphicsItemITF) *QStyleOptionGraphicsItem {
	return QStyleOptionGraphicsItemFromPointer(unsafe.Pointer(C.QStyleOptionGraphicsItem_NewQStyleOptionGraphicsItem2(C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(other)))))
}
