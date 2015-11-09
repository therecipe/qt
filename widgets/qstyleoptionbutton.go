package widgets

//#include "qstyleoptionbutton.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionButton struct {
	QStyleOption
}

type QStyleOptionButton_ITF interface {
	QStyleOption_ITF
	QStyleOptionButton_PTR() *QStyleOptionButton
}

func PointerFromQStyleOptionButton(ptr QStyleOptionButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionButton_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionButtonFromPointer(ptr unsafe.Pointer) *QStyleOptionButton {
	var n = new(QStyleOptionButton)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionButton) QStyleOptionButton_PTR() *QStyleOptionButton {
	return ptr
}

//QStyleOptionButton::ButtonFeature
type QStyleOptionButton__ButtonFeature int64

const (
	QStyleOptionButton__None              = QStyleOptionButton__ButtonFeature(0x00)
	QStyleOptionButton__Flat              = QStyleOptionButton__ButtonFeature(0x01)
	QStyleOptionButton__HasMenu           = QStyleOptionButton__ButtonFeature(0x02)
	QStyleOptionButton__DefaultButton     = QStyleOptionButton__ButtonFeature(0x04)
	QStyleOptionButton__AutoDefaultButton = QStyleOptionButton__ButtonFeature(0x08)
	QStyleOptionButton__CommandLinkButton = QStyleOptionButton__ButtonFeature(0x10)
)

//QStyleOptionButton::StyleOptionType
type QStyleOptionButton__StyleOptionType int64

var (
	QStyleOptionButton__Type = QStyleOptionButton__StyleOptionType(QStyleOption__SO_Button)
)

//QStyleOptionButton::StyleOptionVersion
type QStyleOptionButton__StyleOptionVersion int64

var (
	QStyleOptionButton__Version = QStyleOptionButton__StyleOptionVersion(1)
)

func NewQStyleOptionButton() *QStyleOptionButton {
	return NewQStyleOptionButtonFromPointer(C.QStyleOptionButton_NewQStyleOptionButton())
}

func NewQStyleOptionButton2(other QStyleOptionButton_ITF) *QStyleOptionButton {
	return NewQStyleOptionButtonFromPointer(C.QStyleOptionButton_NewQStyleOptionButton2(PointerFromQStyleOptionButton(other)))
}
