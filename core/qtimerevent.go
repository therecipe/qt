package core

//#include "qtimerevent.h"
import "C"
import (
	"unsafe"
)

type QTimerEvent struct {
	QEvent
}

type QTimerEventITF interface {
	QEventITF
	QTimerEventPTR() *QTimerEvent
}

func PointerFromQTimerEvent(ptr QTimerEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimerEventPTR().Pointer()
	}
	return nil
}

func QTimerEventFromPointer(ptr unsafe.Pointer) *QTimerEvent {
	var n = new(QTimerEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTimerEvent) QTimerEventPTR() *QTimerEvent {
	return ptr
}

func NewQTimerEvent(timerId int) *QTimerEvent {
	return QTimerEventFromPointer(unsafe.Pointer(C.QTimerEvent_NewQTimerEvent(C.int(timerId))))
}

func (ptr *QTimerEvent) TimerId() int {
	if ptr.Pointer() != nil {
		return int(C.QTimerEvent_TimerId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
