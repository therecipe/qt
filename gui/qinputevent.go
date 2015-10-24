package gui

//#include "qinputevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QInputEvent struct {
	core.QEvent
}

type QInputEventITF interface {
	core.QEventITF
	QInputEventPTR() *QInputEvent
}

func PointerFromQInputEvent(ptr QInputEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QInputEventPTR().Pointer()
	}
	return nil
}

func QInputEventFromPointer(ptr unsafe.Pointer) *QInputEvent {
	var n = new(QInputEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QInputEvent) QInputEventPTR() *QInputEvent {
	return ptr
}

func (ptr *QInputEvent) Modifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QInputEvent_Modifiers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
