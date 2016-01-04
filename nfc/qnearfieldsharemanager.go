package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNearFieldShareManager struct {
	core.QObject
}

type QNearFieldShareManager_ITF interface {
	core.QObject_ITF
	QNearFieldShareManager_PTR() *QNearFieldShareManager
}

func PointerFromQNearFieldShareManager(ptr QNearFieldShareManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldShareManager_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldShareManagerFromPointer(ptr unsafe.Pointer) *QNearFieldShareManager {
	var n = new(QNearFieldShareManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldShareManager_") {
		n.SetObjectName("QNearFieldShareManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QNearFieldShareManager) QNearFieldShareManager_PTR() *QNearFieldShareManager {
	return ptr
}

//QNearFieldShareManager::ShareError
type QNearFieldShareManager__ShareError int64

const (
	QNearFieldShareManager__NoError                     = QNearFieldShareManager__ShareError(0)
	QNearFieldShareManager__UnknownError                = QNearFieldShareManager__ShareError(1)
	QNearFieldShareManager__InvalidShareContentError    = QNearFieldShareManager__ShareError(2)
	QNearFieldShareManager__ShareCanceledError          = QNearFieldShareManager__ShareError(3)
	QNearFieldShareManager__ShareInterruptedError       = QNearFieldShareManager__ShareError(4)
	QNearFieldShareManager__ShareRejectedError          = QNearFieldShareManager__ShareError(5)
	QNearFieldShareManager__UnsupportedShareModeError   = QNearFieldShareManager__ShareError(6)
	QNearFieldShareManager__ShareAlreadyInProgressError = QNearFieldShareManager__ShareError(7)
	QNearFieldShareManager__SharePermissionDeniedError  = QNearFieldShareManager__ShareError(8)
)

//QNearFieldShareManager::ShareMode
type QNearFieldShareManager__ShareMode int64

const (
	QNearFieldShareManager__NoShare   = QNearFieldShareManager__ShareMode(0x00)
	QNearFieldShareManager__NdefShare = QNearFieldShareManager__ShareMode(0x01)
	QNearFieldShareManager__FileShare = QNearFieldShareManager__ShareMode(0x02)
)

func NewQNearFieldShareManager(parent core.QObject_ITF) *QNearFieldShareManager {
	defer qt.Recovering("QNearFieldShareManager::QNearFieldShareManager")

	return NewQNearFieldShareManagerFromPointer(C.QNearFieldShareManager_NewQNearFieldShareManager(core.PointerFromQObject(parent)))
}

func (ptr *QNearFieldShareManager) ConnectError(f func(error QNearFieldShareManager__ShareError)) {
	defer qt.Recovering("connect QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectError() {
	defer qt.Recovering("disconnect QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQNearFieldShareManagerError
func callbackQNearFieldShareManagerError(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QNearFieldShareManager::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareError))(QNearFieldShareManager__ShareError(error))
	}

}

func (ptr *QNearFieldShareManager) Error(error QNearFieldShareManager__ShareError) {
	defer qt.Recovering("QNearFieldShareManager::error")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_Error(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QNearFieldShareManager) SetShareModes(mode QNearFieldShareManager__ShareMode) {
	defer qt.Recovering("QNearFieldShareManager::setShareModes")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_SetShareModes(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QNearFieldShareManager) ShareError() QNearFieldShareManager__ShareError {
	defer qt.Recovering("QNearFieldShareManager::shareError")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareError(C.QNearFieldShareManager_ShareError(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareManager::shareModes")

	if ptr.Pointer() != nil {
		return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_ShareModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldShareManager) ConnectShareModesChanged(f func(modes QNearFieldShareManager__ShareMode)) {
	defer qt.Recovering("connect QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectShareModesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "shareModesChanged", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectShareModesChanged() {
	defer qt.Recovering("disconnect QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectShareModesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "shareModesChanged")
	}
}

//export callbackQNearFieldShareManagerShareModesChanged
func callbackQNearFieldShareManagerShareModesChanged(ptr unsafe.Pointer, ptrName *C.char, modes C.int) {
	defer qt.Recovering("callback QNearFieldShareManager::shareModesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "shareModesChanged"); signal != nil {
		signal.(func(QNearFieldShareManager__ShareMode))(QNearFieldShareManager__ShareMode(modes))
	}

}

func (ptr *QNearFieldShareManager) ShareModesChanged(modes QNearFieldShareManager__ShareMode) {
	defer qt.Recovering("QNearFieldShareManager::shareModesChanged")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ShareModesChanged(ptr.Pointer(), C.int(modes))
	}
}

func QNearFieldShareManager_SupportedShareModes() QNearFieldShareManager__ShareMode {
	defer qt.Recovering("QNearFieldShareManager::supportedShareModes")

	return QNearFieldShareManager__ShareMode(C.QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes())
}

func (ptr *QNearFieldShareManager) ConnectTargetDetected(f func(shareTarget *QNearFieldShareTarget)) {
	defer qt.Recovering("connect QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetDetected", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTargetDetected() {
	defer qt.Recovering("disconnect QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetDetected")
	}
}

//export callbackQNearFieldShareManagerTargetDetected
func callbackQNearFieldShareManagerTargetDetected(ptr unsafe.Pointer, ptrName *C.char, shareTarget unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::targetDetected")

	if signal := qt.GetSignal(C.GoString(ptrName), "targetDetected"); signal != nil {
		signal.(func(*QNearFieldShareTarget))(NewQNearFieldShareTargetFromPointer(shareTarget))
	}

}

func (ptr *QNearFieldShareManager) TargetDetected(shareTarget QNearFieldShareTarget_ITF) {
	defer qt.Recovering("QNearFieldShareManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TargetDetected(ptr.Pointer(), PointerFromQNearFieldShareTarget(shareTarget))
	}
}

func (ptr *QNearFieldShareManager) DestroyQNearFieldShareManager() {
	defer qt.Recovering("QNearFieldShareManager::~QNearFieldShareManager")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_DestroyQNearFieldShareManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldShareManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNearFieldShareManagerTimerEvent
func callbackQNearFieldShareManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldShareManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldShareManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNearFieldShareManagerChildEvent
func callbackQNearFieldShareManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldShareManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldShareManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNearFieldShareManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNearFieldShareManagerCustomEvent
func callbackQNearFieldShareManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldShareManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldShareManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldShareManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldShareManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldShareManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldShareManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
