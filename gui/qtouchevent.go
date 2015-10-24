package gui

//#include "qtouchevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTouchEvent struct {
	QInputEvent
}

type QTouchEventITF interface {
	QInputEventITF
	QTouchEventPTR() *QTouchEvent
}

func PointerFromQTouchEvent(ptr QTouchEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTouchEventPTR().Pointer()
	}
	return nil
}

func QTouchEventFromPointer(ptr unsafe.Pointer) *QTouchEvent {
	var n = new(QTouchEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTouchEvent) QTouchEventPTR() *QTouchEvent {
	return ptr
}

func (ptr *QTouchEvent) Device() *QTouchDevice {
	if ptr.Pointer() != nil {
		return QTouchDeviceFromPointer(unsafe.Pointer(C.QTouchEvent_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTouchEvent) Target() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QTouchEvent_Target(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTouchEvent) TouchPointStates() core.Qt__TouchPointState {
	if ptr.Pointer() != nil {
		return core.Qt__TouchPointState(C.QTouchEvent_TouchPointStates(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTouchEvent) Window() *QWindow {
	if ptr.Pointer() != nil {
		return QWindowFromPointer(unsafe.Pointer(C.QTouchEvent_Window(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QTouchEvent) DestroyQTouchEvent() {
	if ptr.Pointer() != nil {
		C.QTouchEvent_DestroyQTouchEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
