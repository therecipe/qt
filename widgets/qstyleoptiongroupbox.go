package widgets

//#include "qstyleoptiongroupbox.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionGroupBox struct {
	QStyleOptionComplex
}

type QStyleOptionGroupBox_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionGroupBox_PTR() *QStyleOptionGroupBox
}

func PointerFromQStyleOptionGroupBox(ptr QStyleOptionGroupBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionGroupBox_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionGroupBoxFromPointer(ptr unsafe.Pointer) *QStyleOptionGroupBox {
	var n = new(QStyleOptionGroupBox)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionGroupBox) QStyleOptionGroupBox_PTR() *QStyleOptionGroupBox {
	return ptr
}

//QStyleOptionGroupBox::StyleOptionType
type QStyleOptionGroupBox__StyleOptionType int64

var (
	QStyleOptionGroupBox__Type = QStyleOptionGroupBox__StyleOptionType(QStyleOption__SO_GroupBox)
)

//QStyleOptionGroupBox::StyleOptionVersion
type QStyleOptionGroupBox__StyleOptionVersion int64

var (
	QStyleOptionGroupBox__Version = QStyleOptionGroupBox__StyleOptionVersion(1)
)

func NewQStyleOptionGroupBox() *QStyleOptionGroupBox {
	return NewQStyleOptionGroupBoxFromPointer(C.QStyleOptionGroupBox_NewQStyleOptionGroupBox())
}

func NewQStyleOptionGroupBox2(other QStyleOptionGroupBox_ITF) *QStyleOptionGroupBox {
	return NewQStyleOptionGroupBoxFromPointer(C.QStyleOptionGroupBox_NewQStyleOptionGroupBox2(PointerFromQStyleOptionGroupBox(other)))
}
