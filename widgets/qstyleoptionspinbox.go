package widgets

//#include "qstyleoptionspinbox.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionSpinBox struct {
	QStyleOptionComplex
}

type QStyleOptionSpinBoxITF interface {
	QStyleOptionComplexITF
	QStyleOptionSpinBoxPTR() *QStyleOptionSpinBox
}

func PointerFromQStyleOptionSpinBox(ptr QStyleOptionSpinBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionSpinBoxPTR().Pointer()
	}
	return nil
}

func QStyleOptionSpinBoxFromPointer(ptr unsafe.Pointer) *QStyleOptionSpinBox {
	var n = new(QStyleOptionSpinBox)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionSpinBox) QStyleOptionSpinBoxPTR() *QStyleOptionSpinBox {
	return ptr
}

//QStyleOptionSpinBox::StyleOptionType
type QStyleOptionSpinBox__StyleOptionType int

var (
	QStyleOptionSpinBox__Type = QStyleOptionSpinBox__StyleOptionType(QStyleOption__SO_SpinBox)
)

//QStyleOptionSpinBox::StyleOptionVersion
type QStyleOptionSpinBox__StyleOptionVersion int

var (
	QStyleOptionSpinBox__Version = QStyleOptionSpinBox__StyleOptionVersion(1)
)

func NewQStyleOptionSpinBox() *QStyleOptionSpinBox {
	return QStyleOptionSpinBoxFromPointer(unsafe.Pointer(C.QStyleOptionSpinBox_NewQStyleOptionSpinBox()))
}

func NewQStyleOptionSpinBox2(other QStyleOptionSpinBoxITF) *QStyleOptionSpinBox {
	return QStyleOptionSpinBoxFromPointer(unsafe.Pointer(C.QStyleOptionSpinBox_NewQStyleOptionSpinBox2(C.QtObjectPtr(PointerFromQStyleOptionSpinBox(other)))))
}
