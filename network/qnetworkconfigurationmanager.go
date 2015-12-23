package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkConfigurationManager struct {
	core.QObject
}

type QNetworkConfigurationManager_ITF interface {
	core.QObject_ITF
	QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager
}

func PointerFromQNetworkConfigurationManager(ptr QNetworkConfigurationManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfigurationManager_PTR().Pointer()
	}
	return nil
}

func NewQNetworkConfigurationManagerFromPointer(ptr unsafe.Pointer) *QNetworkConfigurationManager {
	var n = new(QNetworkConfigurationManager)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkConfigurationManager_") {
		n.SetObjectName("QNetworkConfigurationManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkConfigurationManager) QNetworkConfigurationManager_PTR() *QNetworkConfigurationManager {
	return ptr
}

//QNetworkConfigurationManager::Capability
type QNetworkConfigurationManager__Capability int64

const (
	QNetworkConfigurationManager__CanStartAndStopInterfaces = QNetworkConfigurationManager__Capability(0x00000001)
	QNetworkConfigurationManager__DirectConnectionRouting   = QNetworkConfigurationManager__Capability(0x00000002)
	QNetworkConfigurationManager__SystemSessionSupport      = QNetworkConfigurationManager__Capability(0x00000004)
	QNetworkConfigurationManager__ApplicationLevelRoaming   = QNetworkConfigurationManager__Capability(0x00000008)
	QNetworkConfigurationManager__ForcedRoaming             = QNetworkConfigurationManager__Capability(0x00000010)
	QNetworkConfigurationManager__DataStatistics            = QNetworkConfigurationManager__Capability(0x00000020)
	QNetworkConfigurationManager__NetworkSessionRequired    = QNetworkConfigurationManager__Capability(0x00000040)
)

func NewQNetworkConfigurationManager(parent core.QObject_ITF) *QNetworkConfigurationManager {
	defer qt.Recovering("QNetworkConfigurationManager::QNetworkConfigurationManager")

	return NewQNetworkConfigurationManagerFromPointer(C.QNetworkConfigurationManager_NewQNetworkConfigurationManager(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkConfigurationManager) Capabilities() QNetworkConfigurationManager__Capability {
	defer qt.Recovering("QNetworkConfigurationManager::capabilities")

	if ptr.Pointer() != nil {
		return QNetworkConfigurationManager__Capability(C.QNetworkConfigurationManager_Capabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfigurationManager) IsOnline() bool {
	defer qt.Recovering("QNetworkConfigurationManager::isOnline")

	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_IsOnline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfigurationManager) ConnectOnlineStateChanged(f func(isOnline bool)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::onlineStateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectOnlineStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "onlineStateChanged", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectOnlineStateChanged() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::onlineStateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectOnlineStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "onlineStateChanged")
	}
}

//export callbackQNetworkConfigurationManagerOnlineStateChanged
func callbackQNetworkConfigurationManagerOnlineStateChanged(ptrName *C.char, isOnline C.int) {
	defer qt.Recovering("callback QNetworkConfigurationManager::onlineStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "onlineStateChanged"); signal != nil {
		signal.(func(bool))(int(isOnline) != 0)
	}

}

func (ptr *QNetworkConfigurationManager) ConnectUpdateCompleted(f func()) {
	defer qt.Recovering("connect QNetworkConfigurationManager::updateCompleted")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectUpdateCompleted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updateCompleted", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateCompleted() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::updateCompleted")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectUpdateCompleted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updateCompleted")
	}
}

//export callbackQNetworkConfigurationManagerUpdateCompleted
func callbackQNetworkConfigurationManagerUpdateCompleted(ptrName *C.char) {
	defer qt.Recovering("callback QNetworkConfigurationManager::updateCompleted")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCompleted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkConfigurationManager) UpdateConfigurations() {
	defer qt.Recovering("QNetworkConfigurationManager::updateConfigurations")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateConfigurations(ptr.Pointer())
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManager() {
	defer qt.Recovering("QNetworkConfigurationManager::~QNetworkConfigurationManager")

	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkConfigurationManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkConfigurationManagerTimerEvent
func callbackQNetworkConfigurationManagerTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QNetworkConfigurationManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QNetworkConfigurationManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkConfigurationManagerChildEvent
func callbackQNetworkConfigurationManagerChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QNetworkConfigurationManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QNetworkConfigurationManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkConfigurationManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkConfigurationManagerCustomEvent
func callbackQNetworkConfigurationManagerCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QNetworkConfigurationManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
