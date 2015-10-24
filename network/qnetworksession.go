package network

//#include "qnetworksession.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkSession struct {
	core.QObject
}

type QNetworkSessionITF interface {
	core.QObjectITF
	QNetworkSessionPTR() *QNetworkSession
}

func PointerFromQNetworkSession(ptr QNetworkSessionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkSessionPTR().Pointer()
	}
	return nil
}

func QNetworkSessionFromPointer(ptr unsafe.Pointer) *QNetworkSession {
	var n = new(QNetworkSession)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QNetworkSession_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QNetworkSession) QNetworkSessionPTR() *QNetworkSession {
	return ptr
}

//QNetworkSession::SessionError
type QNetworkSession__SessionError int

var (
	QNetworkSession__UnknownSessionError        = QNetworkSession__SessionError(0)
	QNetworkSession__SessionAbortedError        = QNetworkSession__SessionError(1)
	QNetworkSession__RoamingError               = QNetworkSession__SessionError(2)
	QNetworkSession__OperationNotSupportedError = QNetworkSession__SessionError(3)
	QNetworkSession__InvalidConfigurationError  = QNetworkSession__SessionError(4)
)

//QNetworkSession::State
type QNetworkSession__State int

var (
	QNetworkSession__Invalid      = QNetworkSession__State(0)
	QNetworkSession__NotAvailable = QNetworkSession__State(1)
	QNetworkSession__Connecting   = QNetworkSession__State(2)
	QNetworkSession__Connected    = QNetworkSession__State(3)
	QNetworkSession__Closing      = QNetworkSession__State(4)
	QNetworkSession__Disconnected = QNetworkSession__State(5)
	QNetworkSession__Roaming      = QNetworkSession__State(6)
)

//QNetworkSession::UsagePolicy
type QNetworkSession__UsagePolicy int

var (
	QNetworkSession__NoPolicy                  = QNetworkSession__UsagePolicy(0)
	QNetworkSession__NoBackgroundTrafficPolicy = QNetworkSession__UsagePolicy(1)
)

func NewQNetworkSession(connectionConfig QNetworkConfigurationITF, parent core.QObjectITF) *QNetworkSession {
	return QNetworkSessionFromPointer(unsafe.Pointer(C.QNetworkSession_NewQNetworkSession(C.QtObjectPtr(PointerFromQNetworkConfiguration(connectionConfig)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QNetworkSession) Accept() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Accept(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkSession) Close() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Close(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkSession) ConnectClosed(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectClosed(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "closed", f)
	}
}

func (ptr *QNetworkSession) DisconnectClosed() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectClosed(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "closed")
	}
}

//export callbackQNetworkSessionClosed
func callbackQNetworkSessionClosed(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "closed").(func())()
}

func (ptr *QNetworkSession) Error() QNetworkSession__SessionError {
	if ptr.Pointer() != nil {
		return QNetworkSession__SessionError(C.QNetworkSession_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkSession) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkSession_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkSession) Ignore() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Ignore(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkSession) IsOpen() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkSession_IsOpen(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkSession) Migrate() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Migrate(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkSession) ConnectNewConfigurationActivated(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNewConfigurationActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "newConfigurationActivated", f)
	}
}

func (ptr *QNetworkSession) DisconnectNewConfigurationActivated() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNewConfigurationActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "newConfigurationActivated")
	}
}

//export callbackQNetworkSessionNewConfigurationActivated
func callbackQNetworkSessionNewConfigurationActivated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "newConfigurationActivated").(func())()
}

func (ptr *QNetworkSession) Open() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Open(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkSession) ConnectOpened(f func()) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectOpened(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "opened", f)
	}
}

func (ptr *QNetworkSession) DisconnectOpened() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectOpened(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "opened")
	}
}

//export callbackQNetworkSessionOpened
func callbackQNetworkSessionOpened(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "opened").(func())()
}

func (ptr *QNetworkSession) Reject() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Reject(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkSession) SessionProperty(key string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkSession_SessionProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(key)))
	}
	return ""
}

func (ptr *QNetworkSession) SetSessionProperty(key string, value string) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_SetSessionProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(value))
	}
}

func (ptr *QNetworkSession) ConnectStateChanged(f func(state QNetworkSession__State)) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectStateChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQNetworkSessionStateChanged
func callbackQNetworkSessionStateChanged(ptrName *C.char, state C.int) {
	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QNetworkSession__State))(QNetworkSession__State(state))
}

func (ptr *QNetworkSession) Stop() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QNetworkSession) UsagePolicies() QNetworkSession__UsagePolicy {
	if ptr.Pointer() != nil {
		return QNetworkSession__UsagePolicy(C.QNetworkSession_UsagePolicies(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkSession) ConnectUsagePoliciesChanged(f func(usagePolicies QNetworkSession__UsagePolicy)) {
	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectUsagePoliciesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "usagePoliciesChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectUsagePoliciesChanged() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectUsagePoliciesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "usagePoliciesChanged")
	}
}

//export callbackQNetworkSessionUsagePoliciesChanged
func callbackQNetworkSessionUsagePoliciesChanged(ptrName *C.char, usagePolicies C.int) {
	qt.GetSignal(C.GoString(ptrName), "usagePoliciesChanged").(func(QNetworkSession__UsagePolicy))(QNetworkSession__UsagePolicy(usagePolicies))
}

func (ptr *QNetworkSession) WaitForOpened(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QNetworkSession_WaitForOpened(C.QtObjectPtr(ptr.Pointer()), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QNetworkSession) DestroyQNetworkSession() {
	if ptr.Pointer() != nil {
		C.QNetworkSession_DestroyQNetworkSession(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
