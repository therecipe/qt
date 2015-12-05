package widgets

//#include "widgets.h"
import "C"
import (
	"log"
	"unsafe"
)

type QStyleOptionFrame struct {
	QStyleOption
}

type QStyleOptionFrame_ITF interface {
	QStyleOption_ITF
	QStyleOptionFrame_PTR() *QStyleOptionFrame
}

func PointerFromQStyleOptionFrame(ptr QStyleOptionFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyleOptionFrame_PTR().Pointer()
	}
	return nil
}

func NewQStyleOptionFrameFromPointer(ptr unsafe.Pointer) *QStyleOptionFrame {
	var n = new(QStyleOptionFrame)
	n.SetPointer(ptr)
	return n
}

func (ptr *QStyleOptionFrame) QStyleOptionFrame_PTR() *QStyleOptionFrame {
	return ptr
}

//QStyleOptionFrame::FrameFeature
type QStyleOptionFrame__FrameFeature int64

const (
	QStyleOptionFrame__None    = QStyleOptionFrame__FrameFeature(0x00)
	QStyleOptionFrame__Flat    = QStyleOptionFrame__FrameFeature(0x01)
	QStyleOptionFrame__Rounded = QStyleOptionFrame__FrameFeature(0x02)
)

//QStyleOptionFrame::StyleOptionType
type QStyleOptionFrame__StyleOptionType int64

var (
	QStyleOptionFrame__Type = QStyleOptionFrame__StyleOptionType(QStyleOption__SO_Frame)
)

//QStyleOptionFrame::StyleOptionVersion
type QStyleOptionFrame__StyleOptionVersion int64

var (
	QStyleOptionFrame__Version = QStyleOptionFrame__StyleOptionVersion(3)
)

func NewQStyleOptionFrame() *QStyleOptionFrame {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOptionFrame::QStyleOptionFrame")
		}
	}()

	return NewQStyleOptionFrameFromPointer(C.QStyleOptionFrame_NewQStyleOptionFrame())
}

func NewQStyleOptionFrame2(other QStyleOptionFrame_ITF) *QStyleOptionFrame {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QStyleOptionFrame::QStyleOptionFrame")
		}
	}()

	return NewQStyleOptionFrameFromPointer(C.QStyleOptionFrame_NewQStyleOptionFrame2(PointerFromQStyleOptionFrame(other)))
}
