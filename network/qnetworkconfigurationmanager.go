package network

//#include "qnetworkconfigurationmanager.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QNetworkConfigurationManager_" + qt.RandomIdentifier())
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
	return NewQNetworkConfigurationManagerFromPointer(C.QNetworkConfigurationManager_NewQNetworkConfigurationManager(core.PointerFromQObject(parent)))
}

func (ptr *QNetworkConfigurationManager) Capabilities() QNetworkConfigurationManager__Capability {
	if ptr.Pointer() != nil {
		return QNetworkConfigurationManager__Capability(C.QNetworkConfigurationManager_Capabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkConfigurationManager) IsOnline() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_IsOnline(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkConfigurationManager) ConnectOnlineStateChanged(f func(isOnline bool)) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectOnlineStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "onlineStateChanged", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectOnlineStateChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectOnlineStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "onlineStateChanged")
	}
}

//export callbackQNetworkConfigurationManagerOnlineStateChanged
func callbackQNetworkConfigurationManagerOnlineStateChanged(ptrName *C.char, isOnline C.int) {
	qt.GetSignal(C.GoString(ptrName), "onlineStateChanged").(func(bool))(int(isOnline) != 0)
}

func (ptr *QNetworkConfigurationManager) ConnectUpdateCompleted(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectUpdateCompleted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "updateCompleted", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateCompleted() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectUpdateCompleted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "updateCompleted")
	}
}

//export callbackQNetworkConfigurationManagerUpdateCompleted
func callbackQNetworkConfigurationManagerUpdateCompleted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "updateCompleted").(func())()
}

func (ptr *QNetworkConfigurationManager) UpdateConfigurations() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateConfigurations(ptr.Pointer())
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManager() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
