package nfc

//#include "nfc.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNearFieldManager struct {
	core.QObject
}

type QNearFieldManager_ITF interface {
	core.QObject_ITF
	QNearFieldManager_PTR() *QNearFieldManager
}

func PointerFromQNearFieldManager(ptr QNearFieldManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNearFieldManager_PTR().Pointer()
	}
	return nil
}

func NewQNearFieldManagerFromPointer(ptr unsafe.Pointer) *QNearFieldManager {
	var n = new(QNearFieldManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNearFieldManager_") {
		n.SetObjectName("QNearFieldManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QNearFieldManager) QNearFieldManager_PTR() *QNearFieldManager {
	return ptr
}

//QNearFieldManager::TargetAccessMode
type QNearFieldManager__TargetAccessMode int64

const (
	QNearFieldManager__NoTargetAccess              = QNearFieldManager__TargetAccessMode(0x00)
	QNearFieldManager__NdefReadTargetAccess        = QNearFieldManager__TargetAccessMode(0x01)
	QNearFieldManager__NdefWriteTargetAccess       = QNearFieldManager__TargetAccessMode(0x02)
	QNearFieldManager__TagTypeSpecificTargetAccess = QNearFieldManager__TargetAccessMode(0x04)
)

func (ptr *QNearFieldManager) RegisterNdefMessageHandler(object core.QObject_ITF, method string) int {
	defer qt.Recovering("QNearFieldManager::registerNdefMessageHandler")

	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler(ptr.Pointer(), core.PointerFromQObject(object), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) StartTargetDetection() bool {
	defer qt.Recovering("QNearFieldManager::startTargetDetection")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_StartTargetDetection(ptr.Pointer()) != 0
	}
	return false
}

func NewQNearFieldManager(parent core.QObject_ITF) *QNearFieldManager {
	defer qt.Recovering("QNearFieldManager::QNearFieldManager")

	return NewQNearFieldManagerFromPointer(C.QNearFieldManager_NewQNearFieldManager(core.PointerFromQObject(parent)))
}

func (ptr *QNearFieldManager) IsAvailable() bool {
	defer qt.Recovering("QNearFieldManager::isAvailable")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_IsAvailable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler2(typeNameFormat QNdefRecord__TypeNameFormat, ty core.QByteArray_ITF, object core.QObject_ITF, method string) int {
	defer qt.Recovering("QNearFieldManager::registerNdefMessageHandler")

	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler2(ptr.Pointer(), C.int(typeNameFormat), core.PointerFromQByteArray(ty), core.PointerFromQObject(object), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) RegisterNdefMessageHandler3(filter QNdefFilter_ITF, object core.QObject_ITF, method string) int {
	defer qt.Recovering("QNearFieldManager::registerNdefMessageHandler")

	if ptr.Pointer() != nil {
		return int(C.QNearFieldManager_RegisterNdefMessageHandler3(ptr.Pointer(), PointerFromQNdefFilter(filter), core.PointerFromQObject(object), C.CString(method)))
	}
	return 0
}

func (ptr *QNearFieldManager) SetTargetAccessModes(accessModes QNearFieldManager__TargetAccessMode) {
	defer qt.Recovering("QNearFieldManager::setTargetAccessModes")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_SetTargetAccessModes(ptr.Pointer(), C.int(accessModes))
	}
}

func (ptr *QNearFieldManager) StopTargetDetection() {
	defer qt.Recovering("QNearFieldManager::stopTargetDetection")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_StopTargetDetection(ptr.Pointer())
	}
}

func (ptr *QNearFieldManager) TargetAccessModes() QNearFieldManager__TargetAccessMode {
	defer qt.Recovering("QNearFieldManager::targetAccessModes")

	if ptr.Pointer() != nil {
		return QNearFieldManager__TargetAccessMode(C.QNearFieldManager_TargetAccessModes(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNearFieldManager) ConnectTargetDetected(f func(target *QNearFieldTarget)) {
	defer qt.Recovering("connect QNearFieldManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetDetected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetDetected", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetDetected() {
	defer qt.Recovering("disconnect QNearFieldManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetDetected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetDetected")
	}
}

//export callbackQNearFieldManagerTargetDetected
func callbackQNearFieldManagerTargetDetected(ptr unsafe.Pointer, ptrName *C.char, target unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::targetDetected")

	if signal := qt.GetSignal(C.GoString(ptrName), "targetDetected"); signal != nil {
		signal.(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
	}

}

func (ptr *QNearFieldManager) TargetDetected(target QNearFieldTarget_ITF) {
	defer qt.Recovering("QNearFieldManager::targetDetected")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_TargetDetected(ptr.Pointer(), PointerFromQNearFieldTarget(target))
	}
}

func (ptr *QNearFieldManager) ConnectTargetLost(f func(target *QNearFieldTarget)) {
	defer qt.Recovering("connect QNearFieldManager::targetLost")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ConnectTargetLost(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "targetLost", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTargetLost() {
	defer qt.Recovering("disconnect QNearFieldManager::targetLost")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DisconnectTargetLost(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "targetLost")
	}
}

//export callbackQNearFieldManagerTargetLost
func callbackQNearFieldManagerTargetLost(ptr unsafe.Pointer, ptrName *C.char, target unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::targetLost")

	if signal := qt.GetSignal(C.GoString(ptrName), "targetLost"); signal != nil {
		signal.(func(*QNearFieldTarget))(NewQNearFieldTargetFromPointer(target))
	}

}

func (ptr *QNearFieldManager) TargetLost(target QNearFieldTarget_ITF) {
	defer qt.Recovering("QNearFieldManager::targetLost")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_TargetLost(ptr.Pointer(), PointerFromQNearFieldTarget(target))
	}
}

func (ptr *QNearFieldManager) UnregisterNdefMessageHandler(handlerId int) bool {
	defer qt.Recovering("QNearFieldManager::unregisterNdefMessageHandler")

	if ptr.Pointer() != nil {
		return C.QNearFieldManager_UnregisterNdefMessageHandler(ptr.Pointer(), C.int(handlerId)) != 0
	}
	return false
}

func (ptr *QNearFieldManager) DestroyQNearFieldManager() {
	defer qt.Recovering("QNearFieldManager::~QNearFieldManager")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_DestroyQNearFieldManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNearFieldManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNearFieldManagerTimerEvent
func callbackQNearFieldManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNearFieldManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNearFieldManagerChildEvent
func callbackQNearFieldManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::childEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNearFieldManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNearFieldManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNearFieldManagerCustomEvent
func callbackQNearFieldManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNearFieldManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNearFieldManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNearFieldManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNearFieldManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNearFieldManager::customEvent")

	if ptr.Pointer() != nil {
		C.QNearFieldManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
