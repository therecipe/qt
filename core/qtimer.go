package core

//#include "core.h"
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
	for len(n.ObjectName()) < len("QTimer_") {
		n.SetObjectName("QTimer_" + qt.Identifier())
	}
	return n
}

func (ptr *QTimer) QTimer_PTR() *QTimer {
	return ptr
}

func (ptr *QTimer) RemainingTime() int {
	defer qt.Recovering("QTimer::remainingTime")

	if ptr.Pointer() != nil {
		return int(C.QTimer_RemainingTime(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimer) SetInterval(msec int) {
	defer qt.Recovering("QTimer::setInterval")

	if ptr.Pointer() != nil {
		C.QTimer_SetInterval(ptr.Pointer(), C.int(msec))
	}
}

func NewQTimer(parent QObject_ITF) *QTimer {
	defer qt.Recovering("QTimer::QTimer")

	return NewQTimerFromPointer(C.QTimer_NewQTimer(PointerFromQObject(parent)))
}

func (ptr *QTimer) Interval() int {
	defer qt.Recovering("QTimer::interval")

	if ptr.Pointer() != nil {
		return int(C.QTimer_Interval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimer) IsActive() bool {
	defer qt.Recovering("QTimer::isActive")

	if ptr.Pointer() != nil {
		return C.QTimer_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTimer) IsSingleShot() bool {
	defer qt.Recovering("QTimer::isSingleShot")

	if ptr.Pointer() != nil {
		return C.QTimer_IsSingleShot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTimer) SetSingleShot(singleShot bool) {
	defer qt.Recovering("QTimer::setSingleShot")

	if ptr.Pointer() != nil {
		C.QTimer_SetSingleShot(ptr.Pointer(), C.int(qt.GoBoolToInt(singleShot)))
	}
}

func (ptr *QTimer) SetTimerType(atype Qt__TimerType) {
	defer qt.Recovering("QTimer::setTimerType")

	if ptr.Pointer() != nil {
		C.QTimer_SetTimerType(ptr.Pointer(), C.int(atype))
	}
}

func QTimer_SingleShot2(msec int, timerType Qt__TimerType, receiver QObject_ITF, member string) {
	defer qt.Recovering("QTimer::singleShot")

	C.QTimer_QTimer_SingleShot2(C.int(msec), C.int(timerType), PointerFromQObject(receiver), C.CString(member))
}

func QTimer_SingleShot(msec int, receiver QObject_ITF, member string) {
	defer qt.Recovering("QTimer::singleShot")

	C.QTimer_QTimer_SingleShot(C.int(msec), PointerFromQObject(receiver), C.CString(member))
}

func (ptr *QTimer) Start2() {
	defer qt.Recovering("QTimer::start")

	if ptr.Pointer() != nil {
		C.QTimer_Start2(ptr.Pointer())
	}
}

func (ptr *QTimer) Start(msec int) {
	defer qt.Recovering("QTimer::start")

	if ptr.Pointer() != nil {
		C.QTimer_Start(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QTimer) Stop() {
	defer qt.Recovering("QTimer::stop")

	if ptr.Pointer() != nil {
		C.QTimer_Stop(ptr.Pointer())
	}
}

func (ptr *QTimer) ConnectTimeout(f func()) {
	defer qt.Recovering("connect QTimer::timeout")

	if ptr.Pointer() != nil {
		C.QTimer_ConnectTimeout(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "timeout", f)
	}
}

func (ptr *QTimer) DisconnectTimeout() {
	defer qt.Recovering("disconnect QTimer::timeout")

	if ptr.Pointer() != nil {
		C.QTimer_DisconnectTimeout(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "timeout")
	}
}

//export callbackQTimerTimeout
func callbackQTimerTimeout(ptrName *C.char) {
	defer qt.Recovering("callback QTimer::timeout")

	var signal = qt.GetSignal(C.GoString(ptrName), "timeout")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QTimer) ConnectTimerEvent(f func(e *QTimerEvent)) {
	defer qt.Recovering("connect QTimer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QTimer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QTimer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQTimerTimerEvent
func callbackQTimerTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QTimer::timerEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QTimer) TimerId() int {
	defer qt.Recovering("QTimer::timerId")

	if ptr.Pointer() != nil {
		return int(C.QTimer_TimerId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimer) TimerType() Qt__TimerType {
	defer qt.Recovering("QTimer::timerType")

	if ptr.Pointer() != nil {
		return Qt__TimerType(C.QTimer_TimerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTimer) DestroyQTimer() {
	defer qt.Recovering("QTimer::~QTimer")

	if ptr.Pointer() != nil {
		C.QTimer_DestroyQTimer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
