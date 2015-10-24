package widgets

//#include "qstyleoptiontoolbox.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionToolBox struct {
	QStyleOption
}

type QStyleOptionToolBoxITF interface {
	QStyleOptionITF
	QStyleOptionToolBoxPTR() *QStyleOptionToolBox
}

func PointerFromQStyleOptionToolBox(ptr QStyleOptionToolBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionToolBoxPTR().Pointer()
	}
	return nil
}

func QStyleOptionToolBoxFromPointer(ptr unsafe.Pointer) *QStyleOptionToolBox {
	var n = new(QStyleOptionToolBox)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionToolBox) QStyleOptionToolBoxPTR() *QStyleOptionToolBox {
	return ptr
}

//QStyleOptionToolBox::SelectedPosition
type QStyleOptionToolBox__SelectedPosition int

var (
	QStyleOptionToolBox__NotAdjacent        = QStyleOptionToolBox__SelectedPosition(0)
	QStyleOptionToolBox__NextIsSelected     = QStyleOptionToolBox__SelectedPosition(1)
	QStyleOptionToolBox__PreviousIsSelected = QStyleOptionToolBox__SelectedPosition(2)
)

//QStyleOptionToolBox::StyleOptionType
type QStyleOptionToolBox__StyleOptionType int

var (
	QStyleOptionToolBox__Type = QStyleOptionToolBox__StyleOptionType(QStyleOption__SO_ToolBox)
)

//QStyleOptionToolBox::StyleOptionVersion
type QStyleOptionToolBox__StyleOptionVersion int

var (
	QStyleOptionToolBox__Version = QStyleOptionToolBox__StyleOptionVersion(2)
)

//QStyleOptionToolBox::TabPosition
type QStyleOptionToolBox__TabPosition int

var (
	QStyleOptionToolBox__Beginning  = QStyleOptionToolBox__TabPosition(0)
	QStyleOptionToolBox__Middle     = QStyleOptionToolBox__TabPosition(1)
	QStyleOptionToolBox__End        = QStyleOptionToolBox__TabPosition(2)
	QStyleOptionToolBox__OnlyOneTab = QStyleOptionToolBox__TabPosition(3)
)

func NewQStyleOptionToolBox() *QStyleOptionToolBox {
	return QStyleOptionToolBoxFromPointer(unsafe.Pointer(C.QStyleOptionToolBox_NewQStyleOptionToolBox()))
}

func NewQStyleOptionToolBox2(other QStyleOptionToolBoxITF) *QStyleOptionToolBox {
	return QStyleOptionToolBoxFromPointer(unsafe.Pointer(C.QStyleOptionToolBox_NewQStyleOptionToolBox2(C.QtObjectPtr(PointerFromQStyleOptionToolBox(other)))))
}
