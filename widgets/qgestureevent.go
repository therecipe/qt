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

type QGestureEvent_ITF interface {
	core.QEvent_ITF
	QGestureEvent_PTR() *QGestureEvent
}

func PointerFromQGestureEvent(ptr QGestureEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGestureEvent_PTR().Pointer()
	}
	return nil
}

func NewQGestureEventFromPointer(ptr unsafe.Pointer) *QGestureEvent {
	var n = new(QGestureEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGestureEvent) QGestureEvent_PTR() *QGestureEvent {
	return ptr
}

func (ptr *QGestureEvent) Accept(gesture QGesture_ITF) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_Accept(ptr.Pointer(), PointerFromQGesture(gesture))
	}
}

func (ptr *QGestureEvent) Accept2(gestureType core.Qt__GestureType) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_Accept2(ptr.Pointer(), C.int(gestureType))
	}
}

func (ptr *QGestureEvent) Gesture(ty core.Qt__GestureType) *QGesture {
	if ptr.Pointer() != nil {
		return NewQGestureFromPointer(C.QGestureEvent_Gesture(ptr.Pointer(), C.int(ty)))
	}
	return nil
}

func (ptr *QGestureEvent) Ignore(gesture QGesture_ITF) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_Ignore(ptr.Pointer(), PointerFromQGesture(gesture))
	}
}

func (ptr *QGestureEvent) Ignore2(gestureType core.Qt__GestureType) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_Ignore2(ptr.Pointer(), C.int(gestureType))
	}
}

func (ptr *QGestureEvent) IsAccepted(gesture QGesture_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGestureEvent_IsAccepted(ptr.Pointer(), PointerFromQGesture(gesture)) != 0
	}
	return false
}

func (ptr *QGestureEvent) IsAccepted2(gestureType core.Qt__GestureType) bool {
	if ptr.Pointer() != nil {
		return C.QGestureEvent_IsAccepted2(ptr.Pointer(), C.int(gestureType)) != 0
	}
	return false
}

func (ptr *QGestureEvent) SetAccepted(gesture QGesture_ITF, value bool) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_SetAccepted(ptr.Pointer(), PointerFromQGesture(gesture), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QGestureEvent) SetAccepted2(gestureType core.Qt__GestureType, value bool) {
	if ptr.Pointer() != nil {
		C.QGestureEvent_SetAccepted2(ptr.Pointer(), C.int(gestureType), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QGestureEvent) Widget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QGestureEvent_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGestureEvent) DestroyQGestureEvent() {
	if ptr.Pointer() != nil {
		C.QGestureEvent_DestroyQGestureEvent(ptr.Pointer())
	}
}
