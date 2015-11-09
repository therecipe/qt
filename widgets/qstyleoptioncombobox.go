package widgets

//#include "qstyleoptioncombobox.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionComboBox struct {
	QStyleOptionComplex
}

type QStyleOptionComboBox_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionComboBox_PTR() *QStyleOptionComboBox
}

func PointerFromQStyleOptionComboBox(ptr QStyleOptionComboBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionComboBox_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionComboBoxFromPointer(ptr unsafe.Pointer) *QStyleOptionComboBox {
	var n = new(QStyleOptionComboBox)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionComboBox) QStyleOptionComboBox_PTR() *QStyleOptionComboBox {
	return ptr
}

//QStyleOptionComboBox::StyleOptionType
type QStyleOptionComboBox__StyleOptionType int64

var (
	QStyleOptionComboBox__Type = QStyleOptionComboBox__StyleOptionType(QStyleOption__SO_ComboBox)
)

//QStyleOptionComboBox::StyleOptionVersion
type QStyleOptionComboBox__StyleOptionVersion int64

var (
	QStyleOptionComboBox__Version = QStyleOptionComboBox__StyleOptionVersion(1)
)

func NewQStyleOptionComboBox() *QStyleOptionComboBox {
	return NewQStyleOptionComboBoxFromPointer(C.QStyleOptionComboBox_NewQStyleOptionComboBox())
}

func NewQStyleOptionComboBox2(other QStyleOptionComboBox_ITF) *QStyleOptionComboBox {
	return NewQStyleOptionComboBoxFromPointer(C.QStyleOptionComboBox_NewQStyleOptionComboBox2(PointerFromQStyleOptionComboBox(other)))
}
