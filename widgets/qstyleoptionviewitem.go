package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStyleOptionViewItem struct {
	QStyleOption
}

type QStyleOptionViewItem_ITF interface {
	QStyleOption_ITF
	QStyleOptionViewItem_PTR() *QStyleOptionViewItem
}

func PointerFromQStyleOptionViewItem(ptr QStyleOptionViewItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionViewItem_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionViewItemFromPointer(ptr unsafe.Pointer) *QStyleOptionViewItem {
	var n = new(QStyleOptionViewItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionViewItem) QStyleOptionViewItem_PTR() *QStyleOptionViewItem {
	return ptr
}

//QStyleOptionViewItem::Position
type QStyleOptionViewItem__Position int64

const (
	QStyleOptionViewItem__Left   = QStyleOptionViewItem__Position(0)
	QStyleOptionViewItem__Right  = QStyleOptionViewItem__Position(1)
	QStyleOptionViewItem__Top    = QStyleOptionViewItem__Position(2)
	QStyleOptionViewItem__Bottom = QStyleOptionViewItem__Position(3)
)

//QStyleOptionViewItem::StyleOptionType
type QStyleOptionViewItem__StyleOptionType int64

var (
	QStyleOptionViewItem__Type = QStyleOptionViewItem__StyleOptionType(QStyleOption__SO_ViewItem)
)

//QStyleOptionViewItem::StyleOptionVersion
type QStyleOptionViewItem__StyleOptionVersion int64

var (
	QStyleOptionViewItem__Version = QStyleOptionViewItem__StyleOptionVersion(4)
)

//QStyleOptionViewItem::ViewItemFeature
type QStyleOptionViewItem__ViewItemFeature int64

const (
	QStyleOptionViewItem__None              = QStyleOptionViewItem__ViewItemFeature(0x00)
	QStyleOptionViewItem__WrapText          = QStyleOptionViewItem__ViewItemFeature(0x01)
	QStyleOptionViewItem__Alternate         = QStyleOptionViewItem__ViewItemFeature(0x02)
	QStyleOptionViewItem__HasCheckIndicator = QStyleOptionViewItem__ViewItemFeature(0x04)
	QStyleOptionViewItem__HasDisplay        = QStyleOptionViewItem__ViewItemFeature(0x08)
	QStyleOptionViewItem__HasDecoration     = QStyleOptionViewItem__ViewItemFeature(0x10)
)

//QStyleOptionViewItem::ViewItemPosition
type QStyleOptionViewItem__ViewItemPosition int64

const (
	QStyleOptionViewItem__Invalid   = QStyleOptionViewItem__ViewItemPosition(0)
	QStyleOptionViewItem__Beginning = QStyleOptionViewItem__ViewItemPosition(1)
	QStyleOptionViewItem__Middle    = QStyleOptionViewItem__ViewItemPosition(2)
	QStyleOptionViewItem__End       = QStyleOptionViewItem__ViewItemPosition(3)
	QStyleOptionViewItem__OnlyOne   = QStyleOptionViewItem__ViewItemPosition(4)
)

func NewQStyleOptionViewItem() *QStyleOptionViewItem {
	defer qt.Recovering("QStyleOptionViewItem::QStyleOptionViewItem")

	return NewQStyleOptionViewItemFromPointer(C.QStyleOptionViewItem_NewQStyleOptionViewItem())
}

func NewQStyleOptionViewItem2(other QStyleOptionViewItem_ITF) *QStyleOptionViewItem {
	defer qt.Recovering("QStyleOptionViewItem::QStyleOptionViewItem")

	return NewQStyleOptionViewItemFromPointer(C.QStyleOptionViewItem_NewQStyleOptionViewItem2(PointerFromQStyleOptionViewItem(other)))
}
