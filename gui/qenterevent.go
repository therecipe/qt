package gui

//#include "qenterevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QEnterEvent struct {
	core.QEvent
}

type QEnterEventITF interface {
	core.QEventITF
	QEnterEventPTR() *QEnterEvent
}

func PointerFromQEnterEvent(ptr QEnterEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEnterEventPTR().Pointer()
	}
	return nil
}

func QEnterEventFromPointer(ptr unsafe.Pointer) *QEnterEvent {
	var n = new(QEnterEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QEnterEvent) QEnterEventPTR() *QEnterEvent {
	return ptr
}

func NewQEnterEvent(localPos core.QPointFITF, windowPos core.QPointFITF, screenPos core.QPointFITF) *QEnterEvent {
	return QEnterEventFromPointer(unsafe.Pointer(C.QEnterEvent_NewQEnterEvent(C.QtObjectPtr(core.PointerFromQPointF(localPos)), C.QtObjectPtr(core.PointerFromQPointF(windowPos)), C.QtObjectPtr(core.PointerFromQPointF(screenPos)))))
}

func (ptr *QEnterEvent) GlobalX() int {
	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_GlobalX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QEnterEvent) GlobalY() int {
	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_GlobalY(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QEnterEvent) X() int {
	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QEnterEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
