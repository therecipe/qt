package gui

//#include "qresizeevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QResizeEvent struct {
	core.QEvent
}

type QResizeEvent_ITF interface {
	core.QEvent_ITF
	QResizeEvent_PTR() *QResizeEvent
}

func PointerFromQResizeEvent(ptr QResizeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QResizeEvent_PTR().Pointer()
	}
	return nil
}

func NewQResizeEventFromPointer(ptr unsafe.Pointer) *QResizeEvent {
	var n = new(QResizeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QResizeEvent) QResizeEvent_PTR() *QResizeEvent {
	return ptr
}

func NewQResizeEvent(size core.QSize_ITF, oldSize core.QSize_ITF) *QResizeEvent {
	return NewQResizeEventFromPointer(C.QResizeEvent_NewQResizeEvent(core.PointerFromQSize(size), core.PointerFromQSize(oldSize)))
}
