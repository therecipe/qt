package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStyleOptionSpinBox struct {
	QStyleOptionComplex
}

type QStyleOptionSpinBox_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionSpinBox_PTR() *QStyleOptionSpinBox
}

func PointerFromQStyleOptionSpinBox(ptr QStyleOptionSpinBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionSpinBox_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionSpinBoxFromPointer(ptr unsafe.Pointer) *QStyleOptionSpinBox {
	var n = new(QStyleOptionSpinBox)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionSpinBox) QStyleOptionSpinBox_PTR() *QStyleOptionSpinBox {
	return ptr
}

//QStyleOptionSpinBox::StyleOptionType
type QStyleOptionSpinBox__StyleOptionType int64

var (
	QStyleOptionSpinBox__Type = QStyleOptionSpinBox__StyleOptionType(QStyleOption__SO_SpinBox)
)

//QStyleOptionSpinBox::StyleOptionVersion
type QStyleOptionSpinBox__StyleOptionVersion int64

var (
	QStyleOptionSpinBox__Version = QStyleOptionSpinBox__StyleOptionVersion(1)
)

func NewQStyleOptionSpinBox() *QStyleOptionSpinBox {
	defer qt.Recovering("QStyleOptionSpinBox::QStyleOptionSpinBox")

	return NewQStyleOptionSpinBoxFromPointer(C.QStyleOptionSpinBox_NewQStyleOptionSpinBox())
}

func NewQStyleOptionSpinBox2(other QStyleOptionSpinBox_ITF) *QStyleOptionSpinBox {
	defer qt.Recovering("QStyleOptionSpinBox::QStyleOptionSpinBox")

	return NewQStyleOptionSpinBoxFromPointer(C.QStyleOptionSpinBox_NewQStyleOptionSpinBox2(PointerFromQStyleOptionSpinBox(other)))
}
