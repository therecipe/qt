package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraLocksControl struct {
	QMediaControl
}

type QCameraLocksControl_ITF interface {
	QMediaControl_ITF
	QCameraLocksControl_PTR() *QCameraLocksControl
}

func PointerFromQCameraLocksControl(ptr QCameraLocksControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraLocksControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraLocksControlFromPointer(ptr unsafe.Pointer) *QCameraLocksControl {
	var n = new(QCameraLocksControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraLocksControl_") {
		n.SetObjectName("QCameraLocksControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraLocksControl) QCameraLocksControl_PTR() *QCameraLocksControl {
	return ptr
}

func (ptr *QCameraLocksControl) LockStatus(lock QCamera__LockType) QCamera__LockStatus {
	defer qt.Recovering("QCameraLocksControl::lockStatus")

	if ptr.Pointer() != nil {
		return QCamera__LockStatus(C.QCameraLocksControl_LockStatus(ptr.Pointer(), C.int(lock)))
	}
	return 0
}

func (ptr *QCameraLocksControl) ConnectLockStatusChanged(f func(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason)) {
	defer qt.Recovering("connect QCameraLocksControl::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_ConnectLockStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "lockStatusChanged", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectLockStatusChanged() {
	defer qt.Recovering("disconnect QCameraLocksControl::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DisconnectLockStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "lockStatusChanged")
	}
}

//export callbackQCameraLocksControlLockStatusChanged
func callbackQCameraLocksControlLockStatusChanged(ptr unsafe.Pointer, ptrName *C.char, lock C.int, status C.int, reason C.int) {
	defer qt.Recovering("callback QCameraLocksControl::lockStatusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "lockStatusChanged"); signal != nil {
		signal.(func(QCamera__LockType, QCamera__LockStatus, QCamera__LockChangeReason))(QCamera__LockType(lock), QCamera__LockStatus(status), QCamera__LockChangeReason(reason))
	}

}

func (ptr *QCameraLocksControl) LockStatusChanged(lock QCamera__LockType, status QCamera__LockStatus, reason QCamera__LockChangeReason) {
	defer qt.Recovering("QCameraLocksControl::lockStatusChanged")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_LockStatusChanged(ptr.Pointer(), C.int(lock), C.int(status), C.int(reason))
	}
}

func (ptr *QCameraLocksControl) SearchAndLock(locks QCamera__LockType) {
	defer qt.Recovering("QCameraLocksControl::searchAndLock")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_SearchAndLock(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) SupportedLocks() QCamera__LockType {
	defer qt.Recovering("QCameraLocksControl::supportedLocks")

	if ptr.Pointer() != nil {
		return QCamera__LockType(C.QCameraLocksControl_SupportedLocks(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraLocksControl) Unlock(locks QCamera__LockType) {
	defer qt.Recovering("QCameraLocksControl::unlock")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_Unlock(ptr.Pointer(), C.int(locks))
	}
}

func (ptr *QCameraLocksControl) DestroyQCameraLocksControl() {
	defer qt.Recovering("QCameraLocksControl::~QCameraLocksControl")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_DestroyQCameraLocksControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraLocksControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraLocksControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraLocksControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraLocksControlTimerEvent
func callbackQCameraLocksControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraLocksControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraLocksControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraLocksControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraLocksControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraLocksControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraLocksControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraLocksControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraLocksControlChildEvent
func callbackQCameraLocksControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraLocksControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraLocksControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraLocksControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraLocksControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraLocksControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraLocksControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraLocksControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraLocksControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraLocksControlCustomEvent
func callbackQCameraLocksControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraLocksControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraLocksControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraLocksControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraLocksControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraLocksControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraLocksControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
