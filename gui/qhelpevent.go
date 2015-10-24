package gui

//#include "qhelpevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHelpEvent struct {
	core.QEvent
}

type QHelpEventITF interface {
	core.QEventITF
	QHelpEventPTR() *QHelpEvent
}

func PointerFromQHelpEvent(ptr QHelpEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEventPTR().Pointer()
	}
	return nil
}

func QHelpEventFromPointer(ptr unsafe.Pointer) *QHelpEvent {
	var n = new(QHelpEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpEvent) QHelpEventPTR() *QHelpEvent {
	return ptr
}

func NewQHelpEvent(ty core.QEvent__Type, pos core.QPointITF, globalPos core.QPointITF) *QHelpEvent {
	return QHelpEventFromPointer(unsafe.Pointer(C.QHelpEvent_NewQHelpEvent(C.int(ty), C.QtObjectPtr(core.PointerFromQPoint(pos)), C.QtObjectPtr(core.PointerFromQPoint(globalPos)))))
}

func (ptr *QHelpEvent) GlobalX() int {
	if ptr.Pointer() != nil {
		return int(C.QHelpEvent_GlobalX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpEvent) GlobalY() int {
	if ptr.Pointer() != nil {
		return int(C.QHelpEvent_GlobalY(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpEvent) X() int {
	if ptr.Pointer() != nil {
		return int(C.QHelpEvent_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QHelpEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QHelpEvent_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
