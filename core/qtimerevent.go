package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTimerEvent struct {
	QEvent
}

type QTimerEvent_ITF interface {
	QEvent_ITF
	QTimerEvent_PTR() *QTimerEvent
}

func PointerFromQTimerEvent(ptr QTimerEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimerEvent_PTR().Pointer()
	}
	return nil
}

func NewQTimerEventFromPointer(ptr unsafe.Pointer) *QTimerEvent {
	var n = new(QTimerEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTimerEvent) QTimerEvent_PTR() *QTimerEvent {
	return ptr
}

func NewQTimerEvent(timerId int) *QTimerEvent {
	defer qt.Recovering("QTimerEvent::QTimerEvent")

	return NewQTimerEventFromPointer(C.QTimerEvent_NewQTimerEvent(C.int(timerId)))
}

func (ptr *QTimerEvent) TimerId() int {
	defer qt.Recovering("QTimerEvent::timerId")

	if ptr.Pointer() != nil {
		return int(C.QTimerEvent_TimerId(ptr.Pointer()))
	}
	return 0
}
