package widgets

//#include "qgestureevent.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGestureEvent struct {
	core.QEvent
}

type QGestureEventITF interface {
	core.QEventITF
	QGestureEventPTR() *QGestureEvent
}

func PointerFromQGestureEvent(ptr QGestureEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGestureEventPTR().Pointer()
	}
	return nil
}

func QGestureEventFromPointer(ptr unsafe.Pointer) *QGestureEvent {
	var n = new(QGestureEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGestureEvent) QGestureEventPTR() *QGestureEvent {
	return ptr
}

func (ptr *QGestureEvent) Accept(gesture QGestureITF) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_Accept(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGesture(gesture)))
	}
}

func (ptr *QGestureEvent) Accept2(gestureType core.Qt__GestureType) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_Accept2(C.QtObjectPtr(ptr.Pointer()), C.int(gestureType))
	}
}

func (ptr *QGestureEvent) Gesture(ty core.Qt__GestureType) *QGesture {
	if ptr.Pointer() != nil {
		return QGestureFromPointer(unsafe.Pointer(C.QGestureEvent_Gesture(C.QtObjectPtr(ptr.Pointer()), C.int(ty))))
	}
	return nil
}

func (ptr *QGestureEvent) Ignore(gesture QGestureITF) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_Ignore(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGesture(gesture)))
	}
}

func (ptr *QGestureEvent) Ignore2(gestureType core.Qt__GestureType) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_Ignore2(C.QtObjectPtr(ptr.Pointer()), C.int(gestureType))
	}
}

func (ptr *QGestureEvent) IsAccepted(gesture QGestureITF) bool {
	if ptr.Pointer() != nil {
		return C.QGestureEvent_IsAccepted(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGesture(gesture))) != 0
	}
	return false
}

func (ptr *QGestureEvent) IsAccepted2(gestureType core.Qt__GestureType) bool {
	if ptr.Pointer() != nil {
		return C.QGestureEvent_IsAccepted2(C.QtObjectPtr(ptr.Pointer()), C.int(gestureType)) != 0
	}
	return false
}

func (ptr *QGestureEvent) SetAccepted(gesture QGestureITF, value bool) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_SetAccepted(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGesture(gesture)), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QGestureEvent) SetAccepted2(gestureType core.Qt__GestureType, value bool) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_SetAccepted2(C.QtObjectPtr(ptr.Pointer()), C.int(gestureType), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QGestureEvent) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QGestureEvent_Widget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGestureEvent) DestroyQGestureEvent() {
	if ptr.Pointer() != nil {
		C.QGestureEvent_DestroyQGestureEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
