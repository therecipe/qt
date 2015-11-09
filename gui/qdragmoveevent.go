package gui

//#include "qdragmoveevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDragMoveEvent struct {
	QDropEvent
}

type QDragMoveEvent_ITF interface {
	QDropEvent_ITF
	QDragMoveEvent_PTR() *QDragMoveEvent
}

func PointerFromQDragMoveEvent(ptr QDragMoveEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDragMoveEvent_PTR().Pointer()
	}
	return nil
}

func NewQDragMoveEventFromPointer(ptr unsafe.Pointer) *QDragMoveEvent {
	var n = new(QDragMoveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDragMoveEvent) QDragMoveEvent_PTR() *QDragMoveEvent {
	return ptr
}

func NewQDragMoveEvent(pos core.QPoint_ITF, actions core.Qt__DropAction, data core.QMimeData_ITF, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, ty core.QEvent__Type) *QDragMoveEvent {
	return NewQDragMoveEventFromPointer(C.QDragMoveEvent_NewQDragMoveEvent(core.PointerFromQPoint(pos), C.int(actions), core.PointerFromQMimeData(data), C.int(buttons), C.int(modifiers), C.int(ty)))
}

func (ptr *QDragMoveEvent) Accept2() {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_Accept2(ptr.Pointer())
	}
}

func (ptr *QDragMoveEvent) Accept(rectangle core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_Accept(ptr.Pointer(), core.PointerFromQRect(rectangle))
	}
}

func (ptr *QDragMoveEvent) Ignore2() {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_Ignore2(ptr.Pointer())
	}
}

func (ptr *QDragMoveEvent) Ignore(rectangle core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_Ignore(ptr.Pointer(), core.PointerFromQRect(rectangle))
	}
}

func (ptr *QDragMoveEvent) DestroyQDragMoveEvent() {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_DestroyQDragMoveEvent(ptr.Pointer())
	}
}
