package widgets

//#include "widgets.h"
import "C"
import (
	"log"
	"unsafe"
)

type QStyleOptionFocusRect struct {
	QStyleOption
}

type QStyleOptionFocusRect_ITF interface {
	QStyleOption_ITF
	QStyleOptionFocusRect_PTR() *QStyleOptionFocusRect
}

func PointerFromQStyleOptionFocusRect(ptr QStyleOptionFocusRect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionFocusRect_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionFocusRectFromPointer(ptr unsafe.Pointer) *QStyleOptionFocusRect {
	var n = new(QStyleOptionFocusRect)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionFocusRect) QStyleOptionFocusRect_PTR() *QStyleOptionFocusRect {
	return ptr
}

//QStyleOptionFocusRect::StyleOptionType
type QStyleOptionFocusRect__StyleOptionType int64

var (
	QStyleOptionFocusRect__Type = QStyleOptionFocusRect__StyleOptionType(QStyleOption__SO_FocusRect)
)

//QStyleOptionFocusRect::StyleOptionVersion
type QStyleOptionFocusRect__StyleOptionVersion int64

var (
	QStyleOptionFocusRect__Version = QStyleOptionFocusRect__StyleOptionVersion(1)
)

func NewQStyleOptionFocusRect() *QStyleOptionFocusRect {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOptionFocusRect::QStyleOptionFocusRect")
		}
	}()

	return NewQStyleOptionFocusRectFromPointer(C.QStyleOptionFocusRect_NewQStyleOptionFocusRect())
}

func NewQStyleOptionFocusRect2(other QStyleOptionFocusRect_ITF) *QStyleOptionFocusRect {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOptionFocusRect::QStyleOptionFocusRect")
		}
	}()

	return NewQStyleOptionFocusRectFromPointer(C.QStyleOptionFocusRect_NewQStyleOptionFocusRect2(PointerFromQStyleOptionFocusRect(other)))
}
