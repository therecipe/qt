package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothTransferManager struct {
	core.QObject
}

type QBluetoothTransferManager_ITF interface {
	core.QObject_ITF
	QBluetoothTransferManager_PTR() *QBluetoothTransferManager
}

func PointerFromQBluetoothTransferManager(ptr QBluetoothTransferManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferManager_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) *QBluetoothTransferManager {
	var n = new(QBluetoothTransferManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothTransferManager_") {
		n.SetObjectName("QBluetoothTransferManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothTransferManager) QBluetoothTransferManager_PTR() *QBluetoothTransferManager {
	return ptr
}

func (ptr *QBluetoothTransferManager) Put(request QBluetoothTransferRequest_ITF, data core.QIODevice_ITF) *QBluetoothTransferReply {
	defer qt.Recovering("QBluetoothTransferManager::put")

	if ptr.Pointer() != nil {
		return NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferManager_Put(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func NewQBluetoothTransferManager(parent core.QObject_ITF) *QBluetoothTransferManager {
	defer qt.Recovering("QBluetoothTransferManager::QBluetoothTransferManager")

	return NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferManager_NewQBluetoothTransferManager(core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothTransferManager) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	defer qt.Recovering("connect QBluetoothTransferManager::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothTransferManagerFinished
func callbackQBluetoothTransferManagerFinished(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferManager) Finished(reply QBluetoothTransferReply_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_Finished(ptr.Pointer(), PointerFromQBluetoothTransferReply(reply))
	}
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManager() {
	defer qt.Recovering("QBluetoothTransferManager::~QBluetoothTransferManager")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DestroyQBluetoothTransferManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothTransferManagerTimerEvent
func callbackQBluetoothTransferManagerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothTransferManagerChildEvent
func callbackQBluetoothTransferManagerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothTransferManagerCustomEvent
func callbackQBluetoothTransferManagerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
