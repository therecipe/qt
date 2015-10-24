package core

//#include "qbasictimer.h"
import "C"
import (
	"unsafe"
)

type QBasicTimer struct {
	ptr unsafe.Pointer
}

type QBasicTimerITF interface {
	QBasicTimerPTR() *QBasicTimer
}

func (p *QBasicTimer) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBasicTimer) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBasicTimer(ptr QBasicTimerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBasicTimerPTR().Pointer()
	}
	return nil
}

func QBasicTimerFromPointer(ptr unsafe.Pointer) *QBasicTimer {
	var n = new(QBasicTimer)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBasicTimer) QBasicTimerPTR() *QBasicTimer {
	return ptr
}

func (ptr *QBasicTimer) Start(msec int, object QObjectITF) {
	if ptr.Pointer() != nil {
		C.QBasicTimer_Start(C.QtObjectPtr(ptr.Pointer()), C.int(msec), C.QtObjectPtr(PointerFromQObject(object)))
	}
}

func NewQBasicTimer() *QBasicTimer {
	return QBasicTimerFromPointer(unsafe.Pointer(C.QBasicTimer_NewQBasicTimer()))
}

func (ptr *QBasicTimer) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QBasicTimer_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBasicTimer) Start2(msec int, timerType Qt__TimerType, obj QObjectITF) {
	if ptr.Pointer() != nil {
		C.QBasicTimer_Start2(C.QtObjectPtr(ptr.Pointer()), C.int(msec), C.int(timerType), C.QtObjectPtr(PointerFromQObject(obj)))
	}
}

func (ptr *QBasicTimer) Stop() {
	if ptr.Pointer() != nil {
		C.QBasicTimer_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBasicTimer) TimerId() int {
	if ptr.Pointer() != nil {
		return int(C.QBasicTimer_TimerId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBasicTimer) DestroyQBasicTimer() {
	if ptr.Pointer() != nil {
		C.QBasicTimer_DestroyQBasicTimer(C.QtObjectPtr(ptr.Pointer()))
	}
}
