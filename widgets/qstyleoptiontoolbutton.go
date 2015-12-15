package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStyleOptionToolButton struct {
	QStyleOptionComplex
}

type QStyleOptionToolButton_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionToolButton_PTR() *QStyleOptionToolButton
}

func PointerFromQStyleOptionToolButton(ptr QStyleOptionToolButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionToolButton_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionToolButtonFromPointer(ptr unsafe.Pointer) *QStyleOptionToolButton {
	var n = new(QStyleOptionToolButton)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionToolButton) QStyleOptionToolButton_PTR() *QStyleOptionToolButton {
	return ptr
}

//QStyleOptionToolButton::StyleOptionType
type QStyleOptionToolButton__StyleOptionType int64

var (
	QStyleOptionToolButton__Type = QStyleOptionToolButton__StyleOptionType(QStyleOption__SO_ToolButton)
)

//QStyleOptionToolButton::StyleOptionVersion
type QStyleOptionToolButton__StyleOptionVersion int64

var (
	QStyleOptionToolButton__Version = QStyleOptionToolButton__StyleOptionVersion(1)
)

//QStyleOptionToolButton::ToolButtonFeature
type QStyleOptionToolButton__ToolButtonFeature int64

const (
	QStyleOptionToolButton__None            = QStyleOptionToolButton__ToolButtonFeature(0x00)
	QStyleOptionToolButton__Arrow           = QStyleOptionToolButton__ToolButtonFeature(0x01)
	QStyleOptionToolButton__Menu            = QStyleOptionToolButton__ToolButtonFeature(0x04)
	QStyleOptionToolButton__MenuButtonPopup = QStyleOptionToolButton__ToolButtonFeature(QStyleOptionToolButton__Menu)
	QStyleOptionToolButton__PopupDelay      = QStyleOptionToolButton__ToolButtonFeature(0x08)
	QStyleOptionToolButton__HasMenu         = QStyleOptionToolButton__ToolButtonFeature(0x10)
)

func NewQStyleOptionToolButton() *QStyleOptionToolButton {
	defer qt.Recovering("QStyleOptionToolButton::QStyleOptionToolButton")

	return NewQStyleOptionToolButtonFromPointer(C.QStyleOptionToolButton_NewQStyleOptionToolButton())
}

func NewQStyleOptionToolButton2(other QStyleOptionToolButton_ITF) *QStyleOptionToolButton {
	defer qt.Recovering("QStyleOptionToolButton::QStyleOptionToolButton")

	return NewQStyleOptionToolButtonFromPointer(C.QStyleOptionToolButton_NewQStyleOptionToolButton2(PointerFromQStyleOptionToolButton(other)))
}
