package widgets

//#include "qstyleoptiontitlebar.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionTitleBar struct {
	QStyleOptionComplex
}

type QStyleOptionTitleBar_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionTitleBar_PTR() *QStyleOptionTitleBar
}

func PointerFromQStyleOptionTitleBar(ptr QStyleOptionTitleBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionTitleBar_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionTitleBarFromPointer(ptr unsafe.Pointer) *QStyleOptionTitleBar {
	var n = new(QStyleOptionTitleBar)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionTitleBar) QStyleOptionTitleBar_PTR() *QStyleOptionTitleBar {
	return ptr
}

//QStyleOptionTitleBar::StyleOptionType
type QStyleOptionTitleBar__StyleOptionType int64

var (
	QStyleOptionTitleBar__Type = QStyleOptionTitleBar__StyleOptionType(QStyleOption__SO_TitleBar)
)

//QStyleOptionTitleBar::StyleOptionVersion
type QStyleOptionTitleBar__StyleOptionVersion int64

var (
	QStyleOptionTitleBar__Version = QStyleOptionTitleBar__StyleOptionVersion(1)
)

func NewQStyleOptionTitleBar() *QStyleOptionTitleBar {
	return NewQStyleOptionTitleBarFromPointer(C.QStyleOptionTitleBar_NewQStyleOptionTitleBar())
}

func NewQStyleOptionTitleBar2(other QStyleOptionTitleBar_ITF) *QStyleOptionTitleBar {
	return NewQStyleOptionTitleBarFromPointer(C.QStyleOptionTitleBar_NewQStyleOptionTitleBar2(PointerFromQStyleOptionTitleBar(other)))
}
