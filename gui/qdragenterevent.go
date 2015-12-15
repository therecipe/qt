package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDragEnterEvent struct {
	QDragMoveEvent
}

type QDragEnterEvent_ITF interface {
	QDragMoveEvent_ITF
	QDragEnterEvent_PTR() *QDragEnterEvent
}

func PointerFromQDragEnterEvent(ptr QDragEnterEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDragEnterEvent_PTR().Pointer()
	}
	return nil
}

func NewQDragEnterEventFromPointer(ptr unsafe.Pointer) *QDragEnterEvent {
	var n = new(QDragEnterEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDragEnterEvent) QDragEnterEvent_PTR() *QDragEnterEvent {
	return ptr
}

func NewQDragEnterEvent(point core.QPoint_ITF, actions core.Qt__DropAction, data core.QMimeData_ITF, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QDragEnterEvent {
	defer qt.Recovering("QDragEnterEvent::QDragEnterEvent")

	return NewQDragEnterEventFromPointer(C.QDragEnterEvent_NewQDragEnterEvent(core.PointerFromQPoint(point), C.int(actions), core.PointerFromQMimeData(data), C.int(buttons), C.int(modifiers)))
}
