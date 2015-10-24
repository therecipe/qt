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

type QDragMoveEventITF interface {
	QDropEventITF
	QDragMoveEventPTR() *QDragMoveEvent
}

func PointerFromQDragMoveEvent(ptr QDragMoveEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDragMoveEventPTR().Pointer()
	}
	return nil
}

func QDragMoveEventFromPointer(ptr unsafe.Pointer) *QDragMoveEvent {
	var n = new(QDragMoveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDragMoveEvent) QDragMoveEventPTR() *QDragMoveEvent {
	return ptr
}

func NewQDragMoveEvent(pos core.QPointITF, actions core.Qt__DropAction, data core.QMimeDataITF, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, ty core.QEvent__Type) *QDragMoveEvent {
	return QDragMoveEventFromPointer(unsafe.Pointer(C.QDragMoveEvent_NewQDragMoveEvent(C.QtObjectPtr(core.PointerFromQPoint(pos)), C.int(actions), C.QtObjectPtr(core.PointerFromQMimeData(data)), C.int(buttons), C.int(modifiers), C.int(ty))))
}

func (ptr *QDragMoveEvent) Accept2() {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_Accept2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDragMoveEvent) Accept(rectangle core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_Accept(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)))
	}
}

func (ptr *QDragMoveEvent) Ignore2() {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_Ignore2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QDragMoveEvent) Ignore(rectangle core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_Ignore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)))
	}
}

func (ptr *QDragMoveEvent) DestroyQDragMoveEvent() {
	if ptr.Pointer() != nil {
		C.QDragMoveEvent_DestroyQDragMoveEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
