package widgets

//#include "qstyleoptiontoolbutton.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionToolButton struct {
	QStyleOptionComplex
}

type QStyleOptionToolButtonITF interface {
	QStyleOptionComplexITF
	QStyleOptionToolButtonPTR() *QStyleOptionToolButton
}

func PointerFromQStyleOptionToolButton(ptr QStyleOptionToolButtonITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionToolButtonPTR().Pointer()
	}
	return nil
}

func QStyleOptionToolButtonFromPointer(ptr unsafe.Pointer) *QStyleOptionToolButton {
	var n = new(QStyleOptionToolButton)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionToolButton) QStyleOptionToolButtonPTR() *QStyleOptionToolButton {
	return ptr
}

//QStyleOptionToolButton::StyleOptionType
type QStyleOptionToolButton__StyleOptionType int

var (
	QStyleOptionToolButton__Type = QStyleOptionToolButton__StyleOptionType(QStyleOption__SO_ToolButton)
)

//QStyleOptionToolButton::StyleOptionVersion
type QStyleOptionToolButton__StyleOptionVersion int

var (
	QStyleOptionToolButton__Version = QStyleOptionToolButton__StyleOptionVersion(1)
)

//QStyleOptionToolButton::ToolButtonFeature
type QStyleOptionToolButton__ToolButtonFeature int

var (
	QStyleOptionToolButton__None            = QStyleOptionToolButton__ToolButtonFeature(0x00)
	QStyleOptionToolButton__Arrow           = QStyleOptionToolButton__ToolButtonFeature(0x01)
	QStyleOptionToolButton__Menu            = QStyleOptionToolButton__ToolButtonFeature(0x04)
	QStyleOptionToolButton__MenuButtonPopup = QStyleOptionToolButton__ToolButtonFeature(QStyleOptionToolButton__Menu)
	QStyleOptionToolButton__PopupDelay      = QStyleOptionToolButton__ToolButtonFeature(0x08)
	QStyleOptionToolButton__HasMenu         = QStyleOptionToolButton__ToolButtonFeature(0x10)
)

func NewQStyleOptionToolButton() *QStyleOptionToolButton {
	return QStyleOptionToolButtonFromPointer(unsafe.Pointer(C.QStyleOptionToolButton_NewQStyleOptionToolButton()))
}

func NewQStyleOptionToolButton2(other QStyleOptionToolButtonITF) *QStyleOptionToolButton {
	return QStyleOptionToolButtonFromPointer(unsafe.Pointer(C.QStyleOptionToolButton_NewQStyleOptionToolButton2(C.QtObjectPtr(PointerFromQStyleOptionToolButton(other)))))
}
