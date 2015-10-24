package widgets

//#include "qstyleoptionprogressbar.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionProgressBar struct {
	QStyleOption
}

type QStyleOptionProgressBarITF interface {
	QStyleOptionITF
	QStyleOptionProgressBarPTR() *QStyleOptionProgressBar
}

func PointerFromQStyleOptionProgressBar(ptr QStyleOptionProgressBarITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionProgressBarPTR().Pointer()
	}
	return nil
}

func QStyleOptionProgressBarFromPointer(ptr unsafe.Pointer) *QStyleOptionProgressBar {
	var n = new(QStyleOptionProgressBar)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionProgressBar) QStyleOptionProgressBarPTR() *QStyleOptionProgressBar {
	return ptr
}

//QStyleOptionProgressBar::StyleOptionType
type QStyleOptionProgressBar__StyleOptionType int

var (
	QStyleOptionProgressBar__Type = QStyleOptionProgressBar__StyleOptionType(QStyleOption__SO_ProgressBar)
)

//QStyleOptionProgressBar::StyleOptionVersion
type QStyleOptionProgressBar__StyleOptionVersion int

var (
	QStyleOptionProgressBar__Version = QStyleOptionProgressBar__StyleOptionVersion(2)
)

func NewQStyleOptionProgressBar() *QStyleOptionProgressBar {
	return QStyleOptionProgressBarFromPointer(unsafe.Pointer(C.QStyleOptionProgressBar_NewQStyleOptionProgressBar()))
}

func NewQStyleOptionProgressBar2(other QStyleOptionProgressBarITF) *QStyleOptionProgressBar {
	return QStyleOptionProgressBarFromPointer(unsafe.Pointer(C.QStyleOptionProgressBar_NewQStyleOptionProgressBar2(C.QtObjectPtr(PointerFromQStyleOptionProgressBar(other)))))
}
