package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStyleOptionToolBar struct {
	QStyleOption
}

type QStyleOptionToolBar_ITF interface {
	QStyleOption_ITF
	QStyleOptionToolBar_PTR() *QStyleOptionToolBar
}

func PointerFromQStyleOptionToolBar(ptr QStyleOptionToolBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionToolBar_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionToolBarFromPointer(ptr unsafe.Pointer) *QStyleOptionToolBar {
	var n = new(QStyleOptionToolBar)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionToolBar) QStyleOptionToolBar_PTR() *QStyleOptionToolBar {
	return ptr
}

//QStyleOptionToolBar::StyleOptionType
type QStyleOptionToolBar__StyleOptionType int64

var (
	QStyleOptionToolBar__Type = QStyleOptionToolBar__StyleOptionType(QStyleOption__SO_ToolBar)
)

//QStyleOptionToolBar::StyleOptionVersion
type QStyleOptionToolBar__StyleOptionVersion int64

var (
	QStyleOptionToolBar__Version = QStyleOptionToolBar__StyleOptionVersion(1)
)

//QStyleOptionToolBar::ToolBarFeature
type QStyleOptionToolBar__ToolBarFeature int64

const (
	QStyleOptionToolBar__None    = QStyleOptionToolBar__ToolBarFeature(0x0)
	QStyleOptionToolBar__Movable = QStyleOptionToolBar__ToolBarFeature(0x1)
)

//QStyleOptionToolBar::ToolBarPosition
type QStyleOptionToolBar__ToolBarPosition int64

const (
	QStyleOptionToolBar__Beginning = QStyleOptionToolBar__ToolBarPosition(0)
	QStyleOptionToolBar__Middle    = QStyleOptionToolBar__ToolBarPosition(1)
	QStyleOptionToolBar__End       = QStyleOptionToolBar__ToolBarPosition(2)
	QStyleOptionToolBar__OnlyOne   = QStyleOptionToolBar__ToolBarPosition(3)
)

func NewQStyleOptionToolBar() *QStyleOptionToolBar {
	defer qt.Recovering("QStyleOptionToolBar::QStyleOptionToolBar")

	return NewQStyleOptionToolBarFromPointer(C.QStyleOptionToolBar_NewQStyleOptionToolBar())
}

func NewQStyleOptionToolBar2(other QStyleOptionToolBar_ITF) *QStyleOptionToolBar {
	defer qt.Recovering("QStyleOptionToolBar::QStyleOptionToolBar")

	return NewQStyleOptionToolBarFromPointer(C.QStyleOptionToolBar_NewQStyleOptionToolBar2(PointerFromQStyleOptionToolBar(other)))
}
