package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QHoverEvent struct {
	QInputEvent
}

type QHoverEvent_ITF interface {
	QInputEvent_ITF
	QHoverEvent_PTR() *QHoverEvent
}

func PointerFromQHoverEvent(ptr QHoverEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHoverEvent_PTR().Pointer()
	}
	return nil
}

func NewQHoverEventFromPointer(ptr unsafe.Pointer) *QHoverEvent {
	var n = new(QHoverEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHoverEvent) QHoverEvent_PTR() *QHoverEvent {
	return ptr
}

func NewQHoverEvent(ty core.QEvent__Type, pos core.QPointF_ITF, oldPos core.QPointF_ITF, modifiers core.Qt__KeyboardModifier) *QHoverEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHoverEvent::QHoverEvent")
		}
	}()

	return NewQHoverEventFromPointer(C.QHoverEvent_NewQHoverEvent(C.int(ty), core.PointerFromQPointF(pos), core.PointerFromQPointF(oldPos), C.int(modifiers)))
}
