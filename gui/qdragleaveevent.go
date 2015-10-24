package gui

//#include "qdragleaveevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDragLeaveEvent struct {
	core.QEvent
}

type QDragLeaveEventITF interface {
	core.QEventITF
	QDragLeaveEventPTR() *QDragLeaveEvent
}

func PointerFromQDragLeaveEvent(ptr QDragLeaveEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDragLeaveEventPTR().Pointer()
	}
	return nil
}

func QDragLeaveEventFromPointer(ptr unsafe.Pointer) *QDragLeaveEvent {
	var n = new(QDragLeaveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDragLeaveEvent) QDragLeaveEventPTR() *QDragLeaveEvent {
	return ptr
}

func NewQDragLeaveEvent() *QDragLeaveEvent {
	return QDragLeaveEventFromPointer(unsafe.Pointer(C.QDragLeaveEvent_NewQDragLeaveEvent()))
}
