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

type QNetworkConfigurationManagerITF interface {
	core.QObjectITF
	QNetworkConfigurationManagerPTR() *QNetworkConfigurationManager
}

func PointerFromQNetworkConfigurationManager(ptr QNetworkConfigurationManagerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkConfigurationManagerPTR().Pointer()
	}
	return nil
}

func QNetworkConfigurationManagerFromPointer(ptr unsafe.Pointer) *QNetworkConfigurationManager {
	var n = new(QNetworkConfigurationManager)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNetworkConfigurationManager_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkConfigurationManager) QNetworkConfigurationManagerPTR() *QNetworkConfigurationManager {
	return ptr
}

//QNetworkConfigurationManager::Capability
type QNetworkConfigurationManager__Capability int

var (
	QNetworkConfigurationManager__CanStartAndStopInterfaces = QNetworkConfigurationManager__Capability(0x00000001)
	QNetworkConfigurationManager__DirectConnectionRouting   = QNetworkConfigurationManager__Capability(0x00000002)
	QNetworkConfigurationManager__SystemSessionSupport      = QNetworkConfigurationManager__Capability(0x00000004)
	QNetworkConfigurationManager__ApplicationLevelRoaming   = QNetworkConfigurationManager__Capability(0x00000008)
	QNetworkConfigurationManager__ForcedRoaming             = QNetworkConfigurationManager__Capability(0x00000010)
	QNetworkConfigurationManager__DataStatistics            = QNetworkConfigurationManager__Capability(0x00000020)
	QNetworkConfigurationManager__NetworkSessionRequired    = QNetworkConfigurationManager__Capability(0x00000040)
)

func NewQNetworkConfigurationManager(parent core.QObjectITF) *QNetworkConfigurationManager {
	return QNetworkConfigurationManagerFromPointer(unsafe.Pointer(C.QNetworkConfigurationManager_NewQNetworkConfigurationManager(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QNetworkConfigurationManager) Capabilities() QNetworkConfigurationManager__Capability {
	if ptr.Pointer() != nil {
		return QNetworkConfigurationManager__Capability(C.QNetworkConfigurationManager_Capabilities(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkConfigurationManager) IsOnline() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkConfigurationManager_IsOnline(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkConfigurationManager) ConnectOnlineStateChanged(f func(isOnline bool)) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectOnlineStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "onlineStateChanged", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectOnlineStateChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectOnlineStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "onlineStateChanged")
	}
}

//export callbackQNetworkConfigurationManagerOnlineStateChanged
func callbackQNetworkConfigurationManagerOnlineStateChanged(ptrName *C.char, isOnline C.int) {
	qt.GetSignal(C.GoString(ptrName), "onlineStateChanged").(func(bool))(int(isOnline) != 0)
}

func (ptr *QNetworkConfigurationManager) ConnectUpdateCompleted(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_ConnectUpdateCompleted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "updateCompleted", f)
	}
}

func (ptr *QNetworkConfigurationManager) DisconnectUpdateCompleted() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DisconnectUpdateCompleted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "updateCompleted")
	}
}

//export callbackQNetworkConfigurationManagerUpdateCompleted
func callbackQNetworkConfigurationManagerUpdateCompleted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "updateCompleted").(func())()
}

func (ptr *QNetworkConfigurationManager) UpdateConfigurations() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_UpdateConfigurations(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkConfigurationManager) DestroyQNetworkConfigurationManager() {
	if ptr.Pointer() != nil {
		C.QNetworkConfigurationManager_DestroyQNetworkConfigurationManager(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
