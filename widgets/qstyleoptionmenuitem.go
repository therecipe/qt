package widgets

//#include "qstyleoptionmenuitem.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionMenuItem struct {
	QStyleOption
}

type QStyleOptionMenuItemITF interface {
	QStyleOptionITF
	QStyleOptionMenuItemPTR() *QStyleOptionMenuItem
}

func PointerFromQStyleOptionMenuItem(ptr QStyleOptionMenuItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionMenuItemPTR().Pointer()
	}
	return nil
}

func QStyleOptionMenuItemFromPointer(ptr unsafe.Pointer) *QStyleOptionMenuItem {
	var n = new(QStyleOptionMenuItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionMenuItem) QStyleOptionMenuItemPTR() *QStyleOptionMenuItem {
	return ptr
}

//QStyleOptionMenuItem::CheckType
type QStyleOptionMenuItem__CheckType int

var (
	QStyleOptionMenuItem__NotCheckable = QStyleOptionMenuItem__CheckType(0)
	QStyleOptionMenuItem__Exclusive    = QStyleOptionMenuItem__CheckType(1)
	QStyleOptionMenuItem__NonExclusive = QStyleOptionMenuItem__CheckType(2)
)

//QStyleOptionMenuItem::MenuItemType
type QStyleOptionMenuItem__MenuItemType int

var (
	QStyleOptionMenuItem__Normal      = QStyleOptionMenuItem__MenuItemType(0)
	QStyleOptionMenuItem__DefaultItem = QStyleOptionMenuItem__MenuItemType(1)
	QStyleOptionMenuItem__Separator   = QStyleOptionMenuItem__MenuItemType(2)
	QStyleOptionMenuItem__SubMenu     = QStyleOptionMenuItem__MenuItemType(3)
	QStyleOptionMenuItem__Scroller    = QStyleOptionMenuItem__MenuItemType(4)
	QStyleOptionMenuItem__TearOff     = QStyleOptionMenuItem__MenuItemType(5)
	QStyleOptionMenuItem__Margin      = QStyleOptionMenuItem__MenuItemType(6)
	QStyleOptionMenuItem__EmptyArea   = QStyleOptionMenuItem__MenuItemType(7)
)

//QStyleOptionMenuItem::StyleOptionType
type QStyleOptionMenuItem__StyleOptionType int

var (
	QStyleOptionMenuItem__Type = QStyleOptionMenuItem__StyleOptionType(QStyleOption__SO_MenuItem)
)

//QStyleOptionMenuItem::StyleOptionVersion
type QStyleOptionMenuItem__StyleOptionVersion int

var (
	QStyleOptionMenuItem__Version = QStyleOptionMenuItem__StyleOptionVersion(1)
)

func NewQStyleOptionMenuItem() *QStyleOptionMenuItem {
	return QStyleOptionMenuItemFromPointer(unsafe.Pointer(C.QStyleOptionMenuItem_NewQStyleOptionMenuItem()))
}

func NewQStyleOptionMenuItem2(other QStyleOptionMenuItemITF) *QStyleOptionMenuItem {
	return QStyleOptionMenuItemFromPointer(unsafe.Pointer(C.QStyleOptionMenuItem_NewQStyleOptionMenuItem2(C.QtObjectPtr(PointerFromQStyleOptionMenuItem(other)))))
}
