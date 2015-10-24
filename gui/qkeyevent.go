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

type QKeyEventITF interface {
	QInputEventITF
	QKeyEventPTR() *QKeyEvent
}

func PointerFromQKeyEvent(ptr QKeyEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeyEventPTR().Pointer()
	}
	return nil
}

func QKeyEventFromPointer(ptr unsafe.Pointer) *QKeyEvent {
	var n = new(QKeyEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QKeyEvent) QKeyEventPTR() *QKeyEvent {
	return ptr
}

func (ptr *QKeyEvent) Matches(key QKeySequence__StandardKey) bool {
	if ptr.Pointer() != nil {
		return C.QKeyEvent_Matches(C.QtObjectPtr(ptr.Pointer()), C.int(key)) != 0
	}
	return false
}

func (ptr *QKeyEvent) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QKeyEvent_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QKeyEvent) IsAutoRepeat() bool {
	if ptr.Pointer() != nil {
		return C.QKeyEvent_IsAutoRepeat(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QKeyEvent) Key() int {
	if ptr.Pointer() != nil {
		return int(C.QKeyEvent_Key(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QKeyEvent) Modifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QKeyEvent_Modifiers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QKeyEvent) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QKeyEvent_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
