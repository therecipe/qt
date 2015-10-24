package widgets

//#include "qstyleoptioncombobox.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionComboBox struct {
	QStyleOptionComplex
}

type QStyleOptionComboBoxITF interface {
	QStyleOptionComplexITF
	QStyleOptionComboBoxPTR() *QStyleOptionComboBox
}

func PointerFromQStyleOptionComboBox(ptr QStyleOptionComboBoxITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionComboBoxPTR().Pointer()
	}
	return nil
}

func QStyleOptionComboBoxFromPointer(ptr unsafe.Pointer) *QStyleOptionComboBox {
	var n = new(QStyleOptionComboBox)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionComboBox) QStyleOptionComboBoxPTR() *QStyleOptionComboBox {
	return ptr
}

//QStyleOptionComboBox::StyleOptionType
type QStyleOptionComboBox__StyleOptionType int

var (
	QStyleOptionComboBox__Type = QStyleOptionComboBox__StyleOptionType(QStyleOption__SO_ComboBox)
)

//QStyleOptionComboBox::StyleOptionVersion
type QStyleOptionComboBox__StyleOptionVersion int

var (
	QStyleOptionComboBox__Version = QStyleOptionComboBox__StyleOptionVersion(1)
)

func NewQStyleOptionComboBox() *QStyleOptionComboBox {
	return QStyleOptionComboBoxFromPointer(unsafe.Pointer(C.QStyleOptionComboBox_NewQStyleOptionComboBox()))
}

func NewQStyleOptionComboBox2(other QStyleOptionComboBoxITF) *QStyleOptionComboBox {
	return QStyleOptionComboBoxFromPointer(unsafe.Pointer(C.QStyleOptionComboBox_NewQStyleOptionComboBox2(C.QtObjectPtr(PointerFromQStyleOptionComboBox(other)))))
}
