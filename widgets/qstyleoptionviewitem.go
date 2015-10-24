package widgets

//#include "qstyleoptionviewitem.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionViewItem struct {
	QStyleOption
}

type QStyleOptionViewItemITF interface {
	QStyleOptionITF
	QStyleOptionViewItemPTR() *QStyleOptionViewItem
}

func PointerFromQStyleOptionViewItem(ptr QStyleOptionViewItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionViewItemPTR().Pointer()
	}
	return nil
}

func QStyleOptionViewItemFromPointer(ptr unsafe.Pointer) *QStyleOptionViewItem {
	var n = new(QStyleOptionViewItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionViewItem) QStyleOptionViewItemPTR() *QStyleOptionViewItem {
	return ptr
}

//QStyleOptionViewItem::Position
type QStyleOptionViewItem__Position int

var (
	QStyleOptionViewItem__Left   = QStyleOptionViewItem__Position(0)
	QStyleOptionViewItem__Right  = QStyleOptionViewItem__Position(1)
	QStyleOptionViewItem__Top    = QStyleOptionViewItem__Position(2)
	QStyleOptionViewItem__Bottom = QStyleOptionViewItem__Position(3)
)

//QStyleOptionViewItem::StyleOptionType
type QStyleOptionViewItem__StyleOptionType int

var (
	QStyleOptionViewItem__Type = QStyleOptionViewItem__StyleOptionType(QStyleOption__SO_ViewItem)
)

//QStyleOptionViewItem::StyleOptionVersion
type QStyleOptionViewItem__StyleOptionVersion int

var (
	QStyleOptionViewItem__Version = QStyleOptionViewItem__StyleOptionVersion(4)
)

//QStyleOptionViewItem::ViewItemFeature
type QStyleOptionViewItem__ViewItemFeature int

var (
	QStyleOptionViewItem__None              = QStyleOptionViewItem__ViewItemFeature(0x00)
	QStyleOptionViewItem__WrapText          = QStyleOptionViewItem__ViewItemFeature(0x01)
	QStyleOptionViewItem__Alternate         = QStyleOptionViewItem__ViewItemFeature(0x02)
	QStyleOptionViewItem__HasCheckIndicator = QStyleOptionViewItem__ViewItemFeature(0x04)
	QStyleOptionViewItem__HasDisplay        = QStyleOptionViewItem__ViewItemFeature(0x08)
	QStyleOptionViewItem__HasDecoration     = QStyleOptionViewItem__ViewItemFeature(0x10)
)

//QStyleOptionViewItem::ViewItemPosition
type QStyleOptionViewItem__ViewItemPosition int

var (
	QStyleOptionViewItem__Invalid   = QStyleOptionViewItem__ViewItemPosition(0)
	QStyleOptionViewItem__Beginning = QStyleOptionViewItem__ViewItemPosition(1)
	QStyleOptionViewItem__Middle    = QStyleOptionViewItem__ViewItemPosition(2)
	QStyleOptionViewItem__End       = QStyleOptionViewItem__ViewItemPosition(3)
	QStyleOptionViewItem__OnlyOne   = QStyleOptionViewItem__ViewItemPosition(4)
)

func NewQStyleOptionViewItem() *QStyleOptionViewItem {
	return QStyleOptionViewItemFromPointer(unsafe.Pointer(C.QStyleOptionViewItem_NewQStyleOptionViewItem()))
}

func NewQStyleOptionViewItem2(other QStyleOptionViewItemITF) *QStyleOptionViewItem {
	return QStyleOptionViewItemFromPointer(unsafe.Pointer(C.QStyleOptionViewItem_NewQStyleOptionViewItem2(C.QtObjectPtr(PointerFromQStyleOptionViewItem(other)))))
}
