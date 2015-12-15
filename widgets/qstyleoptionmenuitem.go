package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStyleOptionMenuItem struct {
	QStyleOption
}

type QStyleOptionMenuItem_ITF interface {
	QStyleOption_ITF
	QStyleOptionMenuItem_PTR() *QStyleOptionMenuItem
}

func PointerFromQStyleOptionMenuItem(ptr QStyleOptionMenuItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionMenuItem_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionMenuItemFromPointer(ptr unsafe.Pointer) *QStyleOptionMenuItem {
	var n = new(QStyleOptionMenuItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionMenuItem) QStyleOptionMenuItem_PTR() *QStyleOptionMenuItem {
	return ptr
}

//QStyleOptionMenuItem::CheckType
type QStyleOptionMenuItem__CheckType int64

const (
	QStyleOptionMenuItem__NotCheckable = QStyleOptionMenuItem__CheckType(0)
	QStyleOptionMenuItem__Exclusive    = QStyleOptionMenuItem__CheckType(1)
	QStyleOptionMenuItem__NonExclusive = QStyleOptionMenuItem__CheckType(2)
)

//QStyleOptionMenuItem::MenuItemType
type QStyleOptionMenuItem__MenuItemType int64

const (
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
type QStyleOptionMenuItem__StyleOptionType int64

var (
	QStyleOptionMenuItem__Type = QStyleOptionMenuItem__StyleOptionType(QStyleOption__SO_MenuItem)
)

//QStyleOptionMenuItem::StyleOptionVersion
type QStyleOptionMenuItem__StyleOptionVersion int64

var (
	QStyleOptionMenuItem__Version = QStyleOptionMenuItem__StyleOptionVersion(1)
)

func NewQStyleOptionMenuItem() *QStyleOptionMenuItem {
	defer qt.Recovering("QStyleOptionMenuItem::QStyleOptionMenuItem")

	return NewQStyleOptionMenuItemFromPointer(C.QStyleOptionMenuItem_NewQStyleOptionMenuItem())
}

func NewQStyleOptionMenuItem2(other QStyleOptionMenuItem_ITF) *QStyleOptionMenuItem {
	defer qt.Recovering("QStyleOptionMenuItem::QStyleOptionMenuItem")

	return NewQStyleOptionMenuItemFromPointer(C.QStyleOptionMenuItem_NewQStyleOptionMenuItem2(PointerFromQStyleOptionMenuItem(other)))
}
