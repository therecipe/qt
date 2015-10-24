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

type QTimerITF interface {
	QObjectITF
	QTimerPTR() *QTimer
}

func PointerFromQTimer(ptr QTimerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTimerPTR().Pointer()
	}
	return nil
}

func QTimerFromPointer(ptr unsafe.Pointer) *QTimer {
	var n = new(QTimer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QTimer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QTimer) QTimerPTR() *QTimer {
	return ptr
}

func (ptr *QTimer) RemainingTime() int {
	if ptr.Pointer() != nil {
		return int(C.QTimer_RemainingTime(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTimer) SetInterval(msec int) {
	if ptr.Pointer() != nil {
		C.QTimer_SetInterval(C.QtObjectPtr(ptr.Pointer()), C.int(msec))
	}
}

func NewQTimer(parent QObjectITF) *QTimer {
	return QTimerFromPointer(unsafe.Pointer(C.QTimer_NewQTimer(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QTimer) Interval() int {
	if ptr.Pointer() != nil {
		return int(C.QTimer_Interval(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTimer) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QTimer_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTimer) IsSingleShot() bool {
	if ptr.Pointer() != nil {
		return C.QTimer_IsSingleShot(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QTimer) SetSingleShot(singleShot bool) {
	if ptr.Pointer() != nil {
		C.QTimer_SetSingleShot(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(singleShot)))
	}
}

func (ptr *QTimer) SetTimerType(atype Qt__TimerType) {
	if ptr.Pointer() != nil {
		C.QTimer_SetTimerType(C.QtObjectPtr(ptr.Pointer()), C.int(atype))
	}
}

func QTimer_SingleShot2(msec int, timerType Qt__TimerType, receiver QObjectITF, member string) {
	C.QTimer_QTimer_SingleShot2(C.int(msec), C.int(timerType), C.QtObjectPtr(PointerFromQObject(receiver)), C.CString(member))
}

func QTimer_SingleShot(msec int, receiver QObjectITF, member string) {
	C.QTimer_QTimer_SingleShot(C.int(msec), C.QtObjectPtr(PointerFromQObject(receiver)), C.CString(member))
}

func (ptr *QTimer) Start2() {
	if ptr.Pointer() != nil {
		C.QTimer_Start2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTimer) Start(msec int) {
	if ptr.Pointer() != nil {
		C.QTimer_Start(C.QtObjectPtr(ptr.Pointer()), C.int(msec))
	}
}

func (ptr *QTimer) Stop() {
	if ptr.Pointer() != nil {
		C.QTimer_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QTimer) ConnectTimeout(f func()) {
	if ptr.Pointer() != nil {
		C.QTimer_ConnectTimeout(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "timeout", f)
	}
}

func (ptr *QTimer) DisconnectTimeout() {
	if ptr.Pointer() != nil {
		C.QTimer_DisconnectTimeout(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "timeout")
	}
}

//export callbackQTimerTimeout
func callbackQTimerTimeout(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "timeout").(func())()
}

func (ptr *QTimer) TimerId() int {
	if ptr.Pointer() != nil {
		return int(C.QTimer_TimerId(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTimer) TimerType() Qt__TimerType {
	if ptr.Pointer() != nil {
		return Qt__TimerType(C.QTimer_TimerType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTimer) DestroyQTimer() {
	if ptr.Pointer() != nil {
		C.QTimer_DestroyQTimer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
