package gui

//#include "qpaintevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPaintEvent struct {
	core.QEvent
}

type QPaintEventITF interface {
	core.QEventITF
	QPaintEventPTR() *QPaintEvent
}

func PointerFromQPaintEvent(ptr QPaintEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEventPTR().Pointer()
	}
	return nil
}

func QPaintEventFromPointer(ptr unsafe.Pointer) *QPaintEvent {
	var n = new(QPaintEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPaintEvent) QPaintEventPTR() *QPaintEvent {
	return ptr
}

func NewQPaintEvent2(paintRect core.QRectITF) *QPaintEvent {
	return QPaintEventFromPointer(unsafe.Pointer(C.QPaintEvent_NewQPaintEvent2(C.QtObjectPtr(core.PointerFromQRect(paintRect)))))
}

func NewQPaintEvent(paintRegion QRegionITF) *QPaintEvent {
	return QPaintEventFromPointer(unsafe.Pointer(C.QPaintEvent_NewQPaintEvent(C.QtObjectPtr(PointerFromQRegion(paintRegion)))))
}
