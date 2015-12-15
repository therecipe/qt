package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QStyleOptionTabWidgetFrame struct {
	QStyleOption
}

type QStyleOptionTabWidgetFrame_ITF interface {
	QStyleOption_ITF
	QStyleOptionTabWidgetFrame_PTR() *QStyleOptionTabWidgetFrame
}

func PointerFromQStyleOptionTabWidgetFrame(ptr QStyleOptionTabWidgetFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionTabWidgetFrame_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionTabWidgetFrameFromPointer(ptr unsafe.Pointer) *QStyleOptionTabWidgetFrame {
	var n = new(QStyleOptionTabWidgetFrame)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionTabWidgetFrame) QStyleOptionTabWidgetFrame_PTR() *QStyleOptionTabWidgetFrame {
	return ptr
}

//QStyleOptionTabWidgetFrame::StyleOptionType
type QStyleOptionTabWidgetFrame__StyleOptionType int64

var (
	QStyleOptionTabWidgetFrame__Type = QStyleOptionTabWidgetFrame__StyleOptionType(QStyleOption__SO_TabWidgetFrame)
)

//QStyleOptionTabWidgetFrame::StyleOptionVersion
type QStyleOptionTabWidgetFrame__StyleOptionVersion int64

var (
	QStyleOptionTabWidgetFrame__Version = QStyleOptionTabWidgetFrame__StyleOptionVersion(2)
)

func NewQStyleOptionTabWidgetFrame() *QStyleOptionTabWidgetFrame {
	defer qt.Recovering("QStyleOptionTabWidgetFrame::QStyleOptionTabWidgetFrame")

	return NewQStyleOptionTabWidgetFrameFromPointer(C.QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame())
}

func NewQStyleOptionTabWidgetFrame2(other QStyleOptionTabWidgetFrame_ITF) *QStyleOptionTabWidgetFrame {
	defer qt.Recovering("QStyleOptionTabWidgetFrame::QStyleOptionTabWidgetFrame")

	return NewQStyleOptionTabWidgetFrameFromPointer(C.QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame2(PointerFromQStyleOptionTabWidgetFrame(other)))
}
