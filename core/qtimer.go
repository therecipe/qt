package core

//#include "qtimer.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTimer struct {
	QObject
}

type QTimer_ITF interface {
	QObject_ITF
	QTimer_PTR() *QTimer
}

func PointerFromQTimer(ptr QTimer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimer_PTR().Pointer()
	}
	return nil
}

func NewQTimerFromPointer(ptr unsafe.Pointer) *QTimer {
	var n = new(QTimer)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QTimer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTimer) QTimer_PTR() *QTimer {
	return ptr
}

func (ptr *QTimer) RemainingTime() int {
	if ptr.Pointer() != nil {
		return int(C.QTimer_RemainingTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimer) SetInterval(msec int) {
	if ptr.Pointer() != nil {
		C.QTimer_SetInterval(ptr.Pointer(), C.int(msec))
	}
}

func NewQTimer(parent QObject_ITF) *QTimer {
	return NewQTimerFromPointer(C.QTimer_NewQTimer(PointerFromQObject(parent)))
}

func (ptr *QTimer) Interval() int {
	if ptr.Pointer() != nil {
		return int(C.QTimer_Interval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimer) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QTimer_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTimer) IsSingleShot() bool {
	if ptr.Pointer() != nil {
		return C.QTimer_IsSingleShot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTimer) SetSingleShot(singleShot bool) {
	if ptr.Pointer() != nil {
		C.QTimer_SetSingleShot(ptr.Pointer(), C.int(qt.GoBoolToInt(singleShot)))
	}
}

func (ptr *QTimer) SetTimerType(atype Qt__TimerType) {
	if ptr.Pointer() != nil {
		C.QTimer_SetTimerType(ptr.Pointer(), C.int(atype))
	}
}

func QTimer_SingleShot2(msec int, timerType Qt__TimerType, receiver QObject_ITF, member string) {
	C.QTimer_QTimer_SingleShot2(C.int(msec), C.int(timerType), PointerFromQObject(receiver), C.CString(member))
}

func QTimer_SingleShot(msec int, receiver QObject_ITF, member string) {
	C.QTimer_QTimer_SingleShot(C.int(msec), PointerFromQObject(receiver), C.CString(member))
}

func (ptr *QTimer) Start2() {
	if ptr.Pointer() != nil {
		C.QTimer_Start2(ptr.Pointer())
	}
}

func (ptr *QTimer) Start(msec int) {
	if ptr.Pointer() != nil {
		C.QTimer_Start(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QTimer) Stop() {
	if ptr.Pointer() != nil {
		C.QTimer_Stop(ptr.Pointer())
	}
}

func (ptr *QTimer) ConnectTimeout(f func()) {
	if ptr.Pointer() != nil {
		C.QTimer_ConnectTimeout(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "timeout", f)
	}
}

func (ptr *QTimer) DisconnectTimeout() {
	if ptr.Pointer() != nil {
		C.QTimer_DisconnectTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "timeout")
	}
}

//export callbackQTimerTimeout
func callbackQTimerTimeout(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "timeout").(func())()
}

func (ptr *QTimer) TimerId() int {
	if ptr.Pointer() != nil {
		return int(C.QTimer_TimerId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimer) TimerType() Qt__TimerType {
	if ptr.Pointer() != nil {
		return Qt__TimerType(C.QTimer_TimerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimer) DestroyQTimer() {
	if ptr.Pointer() != nil {
		C.QTimer_DestroyQTimer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
