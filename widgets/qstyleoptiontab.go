package widgets

//#include "qstyleoptiontab.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionTab struct {
	QStyleOption
}

type QStyleOptionTabITF interface {
	QStyleOptionITF
	QStyleOptionTabPTR() *QStyleOptionTab
}

func PointerFromQStyleOptionTab(ptr QStyleOptionTabITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionTabPTR().Pointer()
	}
	return nil
}

func QStyleOptionTabFromPointer(ptr unsafe.Pointer) *QStyleOptionTab {
	var n = new(QStyleOptionTab)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionTab) QStyleOptionTabPTR() *QStyleOptionTab {
	return ptr
}

//QStyleOptionTab::CornerWidget
type QStyleOptionTab__CornerWidget int

var (
	QStyleOptionTab__NoCornerWidgets   = QStyleOptionTab__CornerWidget(0x00)
	QStyleOptionTab__LeftCornerWidget  = QStyleOptionTab__CornerWidget(0x01)
	QStyleOptionTab__RightCornerWidget = QStyleOptionTab__CornerWidget(0x02)
)

//QStyleOptionTab::SelectedPosition
type QStyleOptionTab__SelectedPosition int

var (
	QStyleOptionTab__NotAdjacent        = QStyleOptionTab__SelectedPosition(0)
	QStyleOptionTab__NextIsSelected     = QStyleOptionTab__SelectedPosition(1)
	QStyleOptionTab__PreviousIsSelected = QStyleOptionTab__SelectedPosition(2)
)

//QStyleOptionTab::StyleOptionType
type QStyleOptionTab__StyleOptionType int

var (
	QStyleOptionTab__Type = QStyleOptionTab__StyleOptionType(QStyleOption__SO_Tab)
)

//QStyleOptionTab::StyleOptionVersion
type QStyleOptionTab__StyleOptionVersion int

var (
	QStyleOptionTab__Version = QStyleOptionTab__StyleOptionVersion(3)
)

//QStyleOptionTab::TabFeature
type QStyleOptionTab__TabFeature int

var (
	QStyleOptionTab__None     = QStyleOptionTab__TabFeature(0x00)
	QStyleOptionTab__HasFrame = QStyleOptionTab__TabFeature(0x01)
)

//QStyleOptionTab::TabPosition
type QStyleOptionTab__TabPosition int

var (
	QStyleOptionTab__Beginning  = QStyleOptionTab__TabPosition(0)
	QStyleOptionTab__Middle     = QStyleOptionTab__TabPosition(1)
	QStyleOptionTab__End        = QStyleOptionTab__TabPosition(2)
	QStyleOptionTab__OnlyOneTab = QStyleOptionTab__TabPosition(3)
)

func NewQStyleOptionTab() *QStyleOptionTab {
	return QStyleOptionTabFromPointer(unsafe.Pointer(C.QStyleOptionTab_NewQStyleOptionTab()))
}

func NewQStyleOptionTab2(other QStyleOptionTabITF) *QStyleOptionTab {
	return QStyleOptionTabFromPointer(unsafe.Pointer(C.QStyleOptionTab_NewQStyleOptionTab2(C.QtObjectPtr(PointerFromQStyleOptionTab(other)))))
}
