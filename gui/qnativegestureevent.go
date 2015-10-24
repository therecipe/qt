package gui

//#include "qnativegestureevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNativeGestureEvent struct {
	QInputEvent
}

type QNativeGestureEventITF interface {
	QInputEventITF
	QNativeGestureEventPTR() *QNativeGestureEvent
}

func PointerFromQNativeGestureEvent(ptr QNativeGestureEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNativeGestureEventPTR().Pointer()
	}
	return nil
}

func QNativeGestureEventFromPointer(ptr unsafe.Pointer) *QNativeGestureEvent {
	var n = new(QNativeGestureEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNativeGestureEvent) QNativeGestureEventPTR() *QNativeGestureEvent {
	return ptr
}

func (ptr *QNativeGestureEvent) GestureType() core.Qt__NativeGestureType {
	if ptr.Pointer() != nil {
		return core.Qt__NativeGestureType(C.QNativeGestureEvent_GestureType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
