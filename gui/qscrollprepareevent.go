package gui

//#include "qscrollprepareevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QScrollPrepareEvent struct {
	core.QEvent
}

type QScrollPrepareEventITF interface {
	core.QEventITF
	QScrollPrepareEventPTR() *QScrollPrepareEvent
}

func PointerFromQScrollPrepareEvent(ptr QScrollPrepareEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollPrepareEventPTR().Pointer()
	}
	return nil
}

func QScrollPrepareEventFromPointer(ptr unsafe.Pointer) *QScrollPrepareEvent {
	var n = new(QScrollPrepareEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScrollPrepareEvent) QScrollPrepareEventPTR() *QScrollPrepareEvent {
	return ptr
}

func NewQScrollPrepareEvent(startPos core.QPointFITF) *QScrollPrepareEvent {
	return QScrollPrepareEventFromPointer(unsafe.Pointer(C.QScrollPrepareEvent_NewQScrollPrepareEvent(C.QtObjectPtr(core.PointerFromQPointF(startPos)))))
}

func (ptr *QScrollPrepareEvent) SetContentPos(pos core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QScrollPrepareEvent_SetContentPos(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(pos)))
	}
}

func (ptr *QScrollPrepareEvent) SetContentPosRange(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QScrollPrepareEvent_SetContentPosRange(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QScrollPrepareEvent) SetViewportSize(size core.QSizeFITF) {
	if ptr.Pointer() != nil {
		C.QScrollPrepareEvent_SetViewportSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSizeF(size)))
	}
}

func (ptr *QScrollPrepareEvent) DestroyQScrollPrepareEvent() {
	if ptr.Pointer() != nil {
		C.QScrollPrepareEvent_DestroyQScrollPrepareEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
