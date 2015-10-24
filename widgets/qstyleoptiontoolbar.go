package widgets

//#include "qstyleoptiontoolbar.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionToolBar struct {
	QStyleOption
}

type QStyleOptionToolBarITF interface {
	QStyleOptionITF
	QStyleOptionToolBarPTR() *QStyleOptionToolBar
}

func PointerFromQStyleOptionToolBar(ptr QStyleOptionToolBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionToolBarPTR().Pointer()
	}
	return nil
}

func QStyleOptionToolBarFromPointer(ptr unsafe.Pointer) *QStyleOptionToolBar {
	var n = new(QStyleOptionToolBar)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionToolBar) QStyleOptionToolBarPTR() *QStyleOptionToolBar {
	return ptr
}

//QStyleOptionToolBar::StyleOptionType
type QStyleOptionToolBar__StyleOptionType int

var (
	QStyleOptionToolBar__Type = QStyleOptionToolBar__StyleOptionType(QStyleOption__SO_ToolBar)
)

//QStyleOptionToolBar::StyleOptionVersion
type QStyleOptionToolBar__StyleOptionVersion int

var (
	QStyleOptionToolBar__Version = QStyleOptionToolBar__StyleOptionVersion(1)
)

//QStyleOptionToolBar::ToolBarFeature
type QStyleOptionToolBar__ToolBarFeature int

var (
	QStyleOptionToolBar__None    = QStyleOptionToolBar__ToolBarFeature(0x0)
	QStyleOptionToolBar__Movable = QStyleOptionToolBar__ToolBarFeature(0x1)
)

//QStyleOptionToolBar::ToolBarPosition
type QStyleOptionToolBar__ToolBarPosition int

var (
	QStyleOptionToolBar__Beginning = QStyleOptionToolBar__ToolBarPosition(0)
	QStyleOptionToolBar__Middle    = QStyleOptionToolBar__ToolBarPosition(1)
	QStyleOptionToolBar__End       = QStyleOptionToolBar__ToolBarPosition(2)
	QStyleOptionToolBar__OnlyOne   = QStyleOptionToolBar__ToolBarPosition(3)
)

func NewQStyleOptionToolBar() *QStyleOptionToolBar {
	return QStyleOptionToolBarFromPointer(unsafe.Pointer(C.QStyleOptionToolBar_NewQStyleOptionToolBar()))
}

func NewQStyleOptionToolBar2(other QStyleOptionToolBarITF) *QStyleOptionToolBar {
	return QStyleOptionToolBarFromPointer(unsafe.Pointer(C.QStyleOptionToolBar_NewQStyleOptionToolBar2(C.QtObjectPtr(PointerFromQStyleOptionToolBar(other)))))
}
