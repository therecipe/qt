package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QInputEvent struct {
	core.QEvent
}

type QInputEvent_ITF interface {
	core.QEvent_ITF
	QInputEvent_PTR() *QInputEvent
}

func PointerFromQInputEvent(ptr QInputEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputEvent_PTR().Pointer()
	}
	return nil
}

func NewQInputEventFromPointer(ptr unsafe.Pointer) *QInputEvent {
	var n = new(QInputEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QInputEvent) QInputEvent_PTR() *QInputEvent {
	return ptr
}

func (ptr *QInputEvent) Modifiers() core.Qt__KeyboardModifier {
	defer qt.Recovering("QInputEvent::modifiers")

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QInputEvent_Modifiers(ptr.Pointer()))
	}
	return 0
}
