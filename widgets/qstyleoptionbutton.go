package widgets

//#include "qstyleoptionbutton.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionButton struct {
	QStyleOption
}

type QStyleOptionButtonITF interface {
	QStyleOptionITF
	QStyleOptionButtonPTR() *QStyleOptionButton
}

func PointerFromQStyleOptionButton(ptr QStyleOptionButtonITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionButtonPTR().Pointer()
	}
	return nil
}

func QStyleOptionButtonFromPointer(ptr unsafe.Pointer) *QStyleOptionButton {
	var n = new(QStyleOptionButton)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionButton) QStyleOptionButtonPTR() *QStyleOptionButton {
	return ptr
}

//QStyleOptionButton::ButtonFeature
type QStyleOptionButton__ButtonFeature int

var (
	QStyleOptionButton__None              = QStyleOptionButton__ButtonFeature(0x00)
	QStyleOptionButton__Flat              = QStyleOptionButton__ButtonFeature(0x01)
	QStyleOptionButton__HasMenu           = QStyleOptionButton__ButtonFeature(0x02)
	QStyleOptionButton__DefaultButton     = QStyleOptionButton__ButtonFeature(0x04)
	QStyleOptionButton__AutoDefaultButton = QStyleOptionButton__ButtonFeature(0x08)
	QStyleOptionButton__CommandLinkButton = QStyleOptionButton__ButtonFeature(0x10)
)

//QStyleOptionButton::StyleOptionType
type QStyleOptionButton__StyleOptionType int

var (
	QStyleOptionButton__Type = QStyleOptionButton__StyleOptionType(QStyleOption__SO_Button)
)

//QStyleOptionButton::StyleOptionVersion
type QStyleOptionButton__StyleOptionVersion int

var (
	QStyleOptionButton__Version = QStyleOptionButton__StyleOptionVersion(1)
)

func NewQStyleOptionButton() *QStyleOptionButton {
	return QStyleOptionButtonFromPointer(unsafe.Pointer(C.QStyleOptionButton_NewQStyleOptionButton()))
}

func NewQStyleOptionButton2(other QStyleOptionButtonITF) *QStyleOptionButton {
	return QStyleOptionButtonFromPointer(unsafe.Pointer(C.QStyleOptionButton_NewQStyleOptionButton2(C.QtObjectPtr(PointerFromQStyleOptionButton(other)))))
}
