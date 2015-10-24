package widgets

//#include "qstyleoptiontitlebar.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionTitleBar struct {
	QStyleOptionComplex
}

type QStyleOptionTitleBarITF interface {
	QStyleOptionComplexITF
	QStyleOptionTitleBarPTR() *QStyleOptionTitleBar
}

func PointerFromQStyleOptionTitleBar(ptr QStyleOptionTitleBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionTitleBarPTR().Pointer()
	}
	return nil
}

func QStyleOptionTitleBarFromPointer(ptr unsafe.Pointer) *QStyleOptionTitleBar {
	var n = new(QStyleOptionTitleBar)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionTitleBar) QStyleOptionTitleBarPTR() *QStyleOptionTitleBar {
	return ptr
}

//QStyleOptionTitleBar::StyleOptionType
type QStyleOptionTitleBar__StyleOptionType int

var (
	QStyleOptionTitleBar__Type = QStyleOptionTitleBar__StyleOptionType(QStyleOption__SO_TitleBar)
)

//QStyleOptionTitleBar::StyleOptionVersion
type QStyleOptionTitleBar__StyleOptionVersion int

var (
	QStyleOptionTitleBar__Version = QStyleOptionTitleBar__StyleOptionVersion(1)
)

func NewQStyleOptionTitleBar() *QStyleOptionTitleBar {
	return QStyleOptionTitleBarFromPointer(unsafe.Pointer(C.QStyleOptionTitleBar_NewQStyleOptionTitleBar()))
}

func NewQStyleOptionTitleBar2(other QStyleOptionTitleBarITF) *QStyleOptionTitleBar {
	return QStyleOptionTitleBarFromPointer(unsafe.Pointer(C.QStyleOptionTitleBar_NewQStyleOptionTitleBar2(C.QtObjectPtr(PointerFromQStyleOptionTitleBar(other)))))
}
