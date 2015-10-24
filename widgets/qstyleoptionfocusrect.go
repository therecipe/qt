package widgets

//#include "qstyleoptionfocusrect.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionFocusRect struct {
	QStyleOption
}

type QStyleOptionFocusRectITF interface {
	QStyleOptionITF
	QStyleOptionFocusRectPTR() *QStyleOptionFocusRect
}

func PointerFromQStyleOptionFocusRect(ptr QStyleOptionFocusRectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionFocusRectPTR().Pointer()
	}
	return nil
}

func QStyleOptionFocusRectFromPointer(ptr unsafe.Pointer) *QStyleOptionFocusRect {
	var n = new(QStyleOptionFocusRect)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionFocusRect) QStyleOptionFocusRectPTR() *QStyleOptionFocusRect {
	return ptr
}

//QStyleOptionFocusRect::StyleOptionType
type QStyleOptionFocusRect__StyleOptionType int

var (
	QStyleOptionFocusRect__Type = QStyleOptionFocusRect__StyleOptionType(QStyleOption__SO_FocusRect)
)

//QStyleOptionFocusRect::StyleOptionVersion
type QStyleOptionFocusRect__StyleOptionVersion int

var (
	QStyleOptionFocusRect__Version = QStyleOptionFocusRect__StyleOptionVersion(1)
)

func NewQStyleOptionFocusRect() *QStyleOptionFocusRect {
	return QStyleOptionFocusRectFromPointer(unsafe.Pointer(C.QStyleOptionFocusRect_NewQStyleOptionFocusRect()))
}

func NewQStyleOptionFocusRect2(other QStyleOptionFocusRectITF) *QStyleOptionFocusRect {
	return QStyleOptionFocusRectFromPointer(unsafe.Pointer(C.QStyleOptionFocusRect_NewQStyleOptionFocusRect2(C.QtObjectPtr(PointerFromQStyleOptionFocusRect(other)))))
}
