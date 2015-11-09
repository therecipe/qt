package widgets

//#include "qstyleoptiontoolbox.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionToolBox struct {
	QStyleOption
}

type QStyleOptionToolBox_ITF interface {
	QStyleOption_ITF
	QStyleOptionToolBox_PTR() *QStyleOptionToolBox
}

func PointerFromQStyleOptionToolBox(ptr QStyleOptionToolBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionToolBox_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionToolBoxFromPointer(ptr unsafe.Pointer) *QStyleOptionToolBox {
	var n = new(QStyleOptionToolBox)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionToolBox) QStyleOptionToolBox_PTR() *QStyleOptionToolBox {
	return ptr
}

//QStyleOptionToolBox::SelectedPosition
type QStyleOptionToolBox__SelectedPosition int64

const (
	QStyleOptionToolBox__NotAdjacent        = QStyleOptionToolBox__SelectedPosition(0)
	QStyleOptionToolBox__NextIsSelected     = QStyleOptionToolBox__SelectedPosition(1)
	QStyleOptionToolBox__PreviousIsSelected = QStyleOptionToolBox__SelectedPosition(2)
)

//QStyleOptionToolBox::StyleOptionType
type QStyleOptionToolBox__StyleOptionType int64

var (
	QStyleOptionToolBox__Type = QStyleOptionToolBox__StyleOptionType(QStyleOption__SO_ToolBox)
)

//QStyleOptionToolBox::StyleOptionVersion
type QStyleOptionToolBox__StyleOptionVersion int64

var (
	QStyleOptionToolBox__Version = QStyleOptionToolBox__StyleOptionVersion(2)
)

//QStyleOptionToolBox::TabPosition
type QStyleOptionToolBox__TabPosition int64

const (
	QStyleOptionToolBox__Beginning  = QStyleOptionToolBox__TabPosition(0)
	QStyleOptionToolBox__Middle     = QStyleOptionToolBox__TabPosition(1)
	QStyleOptionToolBox__End        = QStyleOptionToolBox__TabPosition(2)
	QStyleOptionToolBox__OnlyOneTab = QStyleOptionToolBox__TabPosition(3)
)

func NewQStyleOptionToolBox() *QStyleOptionToolBox {
	return NewQStyleOptionToolBoxFromPointer(C.QStyleOptionToolBox_NewQStyleOptionToolBox())
}

func NewQStyleOptionToolBox2(other QStyleOptionToolBox_ITF) *QStyleOptionToolBox {
	return NewQStyleOptionToolBoxFromPointer(C.QStyleOptionToolBox_NewQStyleOptionToolBox2(PointerFromQStyleOptionToolBox(other)))
}
