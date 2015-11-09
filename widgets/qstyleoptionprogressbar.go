package widgets

//#include "qstyleoptionprogressbar.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionProgressBar struct {
	QStyleOption
}

type QStyleOptionProgressBar_ITF interface {
	QStyleOption_ITF
	QStyleOptionProgressBar_PTR() *QStyleOptionProgressBar
}

func PointerFromQStyleOptionProgressBar(ptr QStyleOptionProgressBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionProgressBar_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionProgressBarFromPointer(ptr unsafe.Pointer) *QStyleOptionProgressBar {
	var n = new(QStyleOptionProgressBar)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionProgressBar) QStyleOptionProgressBar_PTR() *QStyleOptionProgressBar {
	return ptr
}

//QStyleOptionProgressBar::StyleOptionType
type QStyleOptionProgressBar__StyleOptionType int64

var (
	QStyleOptionProgressBar__Type = QStyleOptionProgressBar__StyleOptionType(QStyleOption__SO_ProgressBar)
)

//QStyleOptionProgressBar::StyleOptionVersion
type QStyleOptionProgressBar__StyleOptionVersion int64

var (
	QStyleOptionProgressBar__Version = QStyleOptionProgressBar__StyleOptionVersion(2)
)

func NewQStyleOptionProgressBar() *QStyleOptionProgressBar {
	return NewQStyleOptionProgressBarFromPointer(C.QStyleOptionProgressBar_NewQStyleOptionProgressBar())
}

func NewQStyleOptionProgressBar2(other QStyleOptionProgressBar_ITF) *QStyleOptionProgressBar {
	return NewQStyleOptionProgressBarFromPointer(C.QStyleOptionProgressBar_NewQStyleOptionProgressBar2(PointerFromQStyleOptionProgressBar(other)))
}
