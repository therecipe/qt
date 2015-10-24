package gui

//#include "qhoverevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHoverEvent struct {
	QInputEvent
}

type QHoverEventITF interface {
	QInputEventITF
	QHoverEventPTR() *QHoverEvent
}

func PointerFromQHoverEvent(ptr QHoverEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHoverEventPTR().Pointer()
	}
	return nil
}

func QHoverEventFromPointer(ptr unsafe.Pointer) *QHoverEvent {
	var n = new(QHoverEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHoverEvent) QHoverEventPTR() *QHoverEvent {
	return ptr
}

func NewQHoverEvent(ty core.QEvent__Type, pos core.QPointFITF, oldPos core.QPointFITF, modifiers core.Qt__KeyboardModifier) *QHoverEvent {
	return QHoverEventFromPointer(unsafe.Pointer(C.QHoverEvent_NewQHoverEvent(C.int(ty), C.QtObjectPtr(core.PointerFromQPointF(pos)), C.QtObjectPtr(core.PointerFromQPointF(oldPos)), C.int(modifiers))))
}
