package gui

//#include "qkeyevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QKeyEvent struct {
	QInputEvent
}

type QKeyEvent_ITF interface {
	QInputEvent_ITF
	QKeyEvent_PTR() *QKeyEvent
}

func PointerFromQKeyEvent(ptr QKeyEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeyEvent_PTR().Pointer()
	}
	return nil
}

func NewQKeyEventFromPointer(ptr unsafe.Pointer) *QKeyEvent {
	var n = new(QKeyEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QKeyEvent) QKeyEvent_PTR() *QKeyEvent {
	return ptr
}

func (ptr *QKeyEvent) Matches(key QKeySequence__StandardKey) bool {
	if ptr.Pointer() != nil {
		return C.QKeyEvent_Matches(ptr.Pointer(), C.int(key)) != 0
	}
	return false
}

func (ptr *QKeyEvent) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QKeyEvent_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEvent) IsAutoRepeat() bool {
	if ptr.Pointer() != nil {
		return C.QKeyEvent_IsAutoRepeat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QKeyEvent) Key() int {
	if ptr.Pointer() != nil {
		return int(C.QKeyEvent_Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEvent) Modifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QKeyEvent_Modifiers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEvent) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QKeyEvent_Text(ptr.Pointer()))
	}
	return ""
}
