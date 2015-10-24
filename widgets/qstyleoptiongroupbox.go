package widgets

//#include "qstyleoptiongroupbox.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionGroupBox struct {
	QStyleOptionComplex
}

type QStyleOptionGroupBoxITF interface {
	QStyleOptionComplexITF
	QStyleOptionGroupBoxPTR() *QStyleOptionGroupBox
}

func PointerFromQStyleOptionGroupBox(ptr QStyleOptionGroupBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionGroupBoxPTR().Pointer()
	}
	return nil
}

func QStyleOptionGroupBoxFromPointer(ptr unsafe.Pointer) *QStyleOptionGroupBox {
	var n = new(QStyleOptionGroupBox)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionGroupBox) QStyleOptionGroupBoxPTR() *QStyleOptionGroupBox {
	return ptr
}

//QStyleOptionGroupBox::StyleOptionType
type QStyleOptionGroupBox__StyleOptionType int

var (
	QStyleOptionGroupBox__Type = QStyleOptionGroupBox__StyleOptionType(QStyleOption__SO_GroupBox)
)

//QStyleOptionGroupBox::StyleOptionVersion
type QStyleOptionGroupBox__StyleOptionVersion int

var (
	QStyleOptionGroupBox__Version = QStyleOptionGroupBox__StyleOptionVersion(1)
)

func NewQStyleOptionGroupBox() *QStyleOptionGroupBox {
	return QStyleOptionGroupBoxFromPointer(unsafe.Pointer(C.QStyleOptionGroupBox_NewQStyleOptionGroupBox()))
}

func NewQStyleOptionGroupBox2(other QStyleOptionGroupBoxITF) *QStyleOptionGroupBox {
	return QStyleOptionGroupBoxFromPointer(unsafe.Pointer(C.QStyleOptionGroupBox_NewQStyleOptionGroupBox2(C.QtObjectPtr(PointerFromQStyleOptionGroupBox(other)))))
}
