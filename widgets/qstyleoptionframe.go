package widgets

//#include "qstyleoptionframe.h"
import "C"
import (
	"unsafe"
)

type QStyleOptionFrame struct {
	QStyleOption
}

type QStyleOptionFrameITF interface {
	QStyleOptionITF
	QStyleOptionFramePTR() *QStyleOptionFrame
}

func PointerFromQStyleOptionFrame(ptr QStyleOptionFrameITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionFramePTR().Pointer()
	}
	return nil
}

func QStyleOptionFrameFromPointer(ptr unsafe.Pointer) *QStyleOptionFrame {
	var n = new(QStyleOptionFrame)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionFrame) QStyleOptionFramePTR() *QStyleOptionFrame {
	return ptr
}

//QStyleOptionFrame::FrameFeature
type QStyleOptionFrame__FrameFeature int

var (
	QStyleOptionFrame__None    = QStyleOptionFrame__FrameFeature(0x00)
	QStyleOptionFrame__Flat    = QStyleOptionFrame__FrameFeature(0x01)
	QStyleOptionFrame__Rounded = QStyleOptionFrame__FrameFeature(0x02)
)

//QStyleOptionFrame::StyleOptionType
type QStyleOptionFrame__StyleOptionType int

var (
	QStyleOptionFrame__Type = QStyleOptionFrame__StyleOptionType(QStyleOption__SO_Frame)
)

//QStyleOptionFrame::StyleOptionVersion
type QStyleOptionFrame__StyleOptionVersion int

var (
	QStyleOptionFrame__Version = QStyleOptionFrame__StyleOptionVersion(3)
)

func NewQStyleOptionFrame() *QStyleOptionFrame {
	return QStyleOptionFrameFromPointer(unsafe.Pointer(C.QStyleOptionFrame_NewQStyleOptionFrame()))
}

func NewQStyleOptionFrame2(other QStyleOptionFrameITF) *QStyleOptionFrame {
	return QStyleOptionFrameFromPointer(unsafe.Pointer(C.QStyleOptionFrame_NewQStyleOptionFrame2(C.QtObjectPtr(PointerFromQStyleOptionFrame(other)))))
}
