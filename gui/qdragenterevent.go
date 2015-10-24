package gui

//#include "qdragenterevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDragEnterEvent struct {
	QDragMoveEvent
}

type QDragEnterEventITF interface {
	QDragMoveEventITF
	QDragEnterEventPTR() *QDragEnterEvent
}

func PointerFromQDragEnterEvent(ptr QDragEnterEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDragEnterEventPTR().Pointer()
	}
	return nil
}

func QDragEnterEventFromPointer(ptr unsafe.Pointer) *QDragEnterEvent {
	var n = new(QDragEnterEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDragEnterEvent) QDragEnterEventPTR() *QDragEnterEvent {
	return ptr
}

func NewQDragEnterEvent(point core.QPointITF, actions core.Qt__DropAction, data core.QMimeDataITF, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QDragEnterEvent {
	return QDragEnterEventFromPointer(unsafe.Pointer(C.QDragEnterEvent_NewQDragEnterEvent(C.QtObjectPtr(core.PointerFromQPoint(point)), C.int(actions), C.QtObjectPtr(core.PointerFromQMimeData(data)), C.int(buttons), C.int(modifiers))))
}
