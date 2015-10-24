package widgets

//#include "qstyleoptiontabwidgetframe.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionTabWidgetFrame struct {
	QStyleOption
}

type QStyleOptionTabWidgetFrameITF interface {
	QStyleOptionITF
	QStyleOptionTabWidgetFramePTR() *QStyleOptionTabWidgetFrame
}

func PointerFromQStyleOptionTabWidgetFrame(ptr QStyleOptionTabWidgetFrameITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionTabWidgetFramePTR().Pointer()
	}
	return nil
}

func QStyleOptionTabWidgetFrameFromPointer(ptr unsafe.Pointer) *QStyleOptionTabWidgetFrame {
	var n = new(QStyleOptionTabWidgetFrame)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionTabWidgetFrame) QStyleOptionTabWidgetFramePTR() *QStyleOptionTabWidgetFrame {
	return ptr
}

//QStyleOptionTabWidgetFrame::StyleOptionType
type QStyleOptionTabWidgetFrame__StyleOptionType int

var (
	QStyleOptionTabWidgetFrame__Type = QStyleOptionTabWidgetFrame__StyleOptionType(QStyleOption__SO_TabWidgetFrame)
)

//QStyleOptionTabWidgetFrame::StyleOptionVersion
type QStyleOptionTabWidgetFrame__StyleOptionVersion int

var (
	QStyleOptionTabWidgetFrame__Version = QStyleOptionTabWidgetFrame__StyleOptionVersion(2)
)

func NewQStyleOptionTabWidgetFrame() *QStyleOptionTabWidgetFrame {
	return QStyleOptionTabWidgetFrameFromPointer(unsafe.Pointer(C.QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame()))
}

func NewQStyleOptionTabWidgetFrame2(other QStyleOptionTabWidgetFrameITF) *QStyleOptionTabWidgetFrame {
	return QStyleOptionTabWidgetFrameFromPointer(unsafe.Pointer(C.QStyleOptionTabWidgetFrame_NewQStyleOptionTabWidgetFrame2(C.QtObjectPtr(PointerFromQStyleOptionTabWidgetFrame(other)))))
}
