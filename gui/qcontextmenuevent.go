package gui

//#include "qcontextmenuevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QContextMenuEvent struct {
	QInputEvent
}

type QContextMenuEventITF interface {
	QInputEventITF
	QContextMenuEventPTR() *QContextMenuEvent
}

func PointerFromQContextMenuEvent(ptr QContextMenuEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QContextMenuEventPTR().Pointer()
	}
	return nil
}

func QContextMenuEventFromPointer(ptr unsafe.Pointer) *QContextMenuEvent {
	var n = new(QContextMenuEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QContextMenuEvent) QContextMenuEventPTR() *QContextMenuEvent {
	return ptr
}

//QContextMenuEvent::Reason
type QContextMenuEvent__Reason int

var (
	QContextMenuEvent__Mouse    = QContextMenuEvent__Reason(0)
	QContextMenuEvent__Keyboard = QContextMenuEvent__Reason(1)
	QContextMenuEvent__Other    = QContextMenuEvent__Reason(2)
)

func NewQContextMenuEvent3(reason QContextMenuEvent__Reason, pos core.QPointITF) *QContextMenuEvent {
	return QContextMenuEventFromPointer(unsafe.Pointer(C.QContextMenuEvent_NewQContextMenuEvent3(C.int(reason), C.QtObjectPtr(core.PointerFromQPoint(pos)))))
}

func NewQContextMenuEvent2(reason QContextMenuEvent__Reason, pos core.QPointITF, globalPos core.QPointITF) *QContextMenuEvent {
	return QContextMenuEventFromPointer(unsafe.Pointer(C.QContextMenuEvent_NewQContextMenuEvent2(C.int(reason), C.QtObjectPtr(core.PointerFromQPoint(pos)), C.QtObjectPtr(core.PointerFromQPoint(globalPos)))))
}

func NewQContextMenuEvent(reason QContextMenuEvent__Reason, pos core.QPointITF, globalPos core.QPointITF, modifiers core.Qt__KeyboardModifier) *QContextMenuEvent {
	return QContextMenuEventFromPointer(unsafe.Pointer(C.QContextMenuEvent_NewQContextMenuEvent(C.int(reason), C.QtObjectPtr(core.PointerFromQPoint(pos)), C.QtObjectPtr(core.PointerFromQPoint(globalPos)), C.int(modifiers))))
}

func (ptr *QContextMenuEvent) GlobalX() int {
	if ptr.Pointer() != nil {
		return int(C.QContextMenuEvent_GlobalX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QContextMenuEvent) GlobalY() int {
	if ptr.Pointer() != nil {
		return int(C.QContextMenuEvent_GlobalY(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QContextMenuEvent) Reason() QContextMenuEvent__Reason {
	if ptr.Pointer() != nil {
		return QContextMenuEvent__Reason(C.QContextMenuEvent_Reason(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QContextMenuEvent) X() int {
	if ptr.Pointer() != nil {
		return int(C.QContextMenuEvent_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QContextMenuEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QContextMenuEvent_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
