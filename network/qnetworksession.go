package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QNetworkSession struct {
	core.QObject
}

type QNetworkSession_ITF interface {
	core.QObject_ITF
	QNetworkSession_PTR() *QNetworkSession
}

func PointerFromQNetworkSession(ptr QNetworkSession_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkSession_PTR().Pointer()
	}
	return nil
}

func NewQNetworkSessionFromPointer(ptr unsafe.Pointer) *QNetworkSession {
	var n = new(QNetworkSession)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QNetworkSession_") {
		n.SetObjectName("QNetworkSession_" + qt.Identifier())
	}
	return n
}

func (ptr *QNetworkSession) QNetworkSession_PTR() *QNetworkSession {
	return ptr
}

//QNetworkSession::SessionError
type QNetworkSession__SessionError int64

const (
	QNetworkSession__UnknownSessionError        = QNetworkSession__SessionError(0)
	QNetworkSession__SessionAbortedError        = QNetworkSession__SessionError(1)
	QNetworkSession__RoamingError               = QNetworkSession__SessionError(2)
	QNetworkSession__OperationNotSupportedError = QNetworkSession__SessionError(3)
	QNetworkSession__InvalidConfigurationError  = QNetworkSession__SessionError(4)
)

//QNetworkSession::State
type QNetworkSession__State int64

const (
	QNetworkSession__Invalid      = QNetworkSession__State(0)
	QNetworkSession__NotAvailable = QNetworkSession__State(1)
	QNetworkSession__Connecting   = QNetworkSession__State(2)
	QNetworkSession__Connected    = QNetworkSession__State(3)
	QNetworkSession__Closing      = QNetworkSession__State(4)
	QNetworkSession__Disconnected = QNetworkSession__State(5)
	QNetworkSession__Roaming      = QNetworkSession__State(6)
)

//QNetworkSession::UsagePolicy
type QNetworkSession__UsagePolicy int64

const (
	QNetworkSession__NoPolicy                  = QNetworkSession__UsagePolicy(0)
	QNetworkSession__NoBackgroundTrafficPolicy = QNetworkSession__UsagePolicy(1)
)

func NewQNetworkSession(connectionConfig QNetworkConfiguration_ITF, parent core.QObject_ITF) *QNetworkSession {
	defer qt.Recovering("QNetworkSession::QNetworkSession")

	return NewQNetworkSessionFromPointer(C.QNetworkSession_NewQNetworkSession(PointerFromQNetworkConfiguration(connectionConfig), core.PointerFromQObject(parent)))
}

func (ptr *QNetworkSession) Accept() {
	defer qt.Recovering("QNetworkSession::accept")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Accept(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Close() {
	defer qt.Recovering("QNetworkSession::close")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Close(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectClosed(f func()) {
	defer qt.Recovering("connect QNetworkSession::closed")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectClosed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "closed", f)
	}
}

func (ptr *QNetworkSession) DisconnectClosed() {
	defer qt.Recovering("disconnect QNetworkSession::closed")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "closed")
	}
}

//export callbackQNetworkSessionClosed
func callbackQNetworkSessionClosed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkSession::closed")

	if signal := qt.GetSignal(C.GoString(ptrName), "closed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) Closed() {
	defer qt.Recovering("QNetworkSession::closed")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Closed(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectError2(f func(error QNetworkSession__SessionError)) {
	defer qt.Recovering("connect QNetworkSession::error")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QNetworkSession) DisconnectError2() {
	defer qt.Recovering("disconnect QNetworkSession::error")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQNetworkSessionError2
func callbackQNetworkSessionError2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QNetworkSession::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QNetworkSession__SessionError))(QNetworkSession__SessionError(error))
	}

}

func (ptr *QNetworkSession) Error2(error QNetworkSession__SessionError) {
	defer qt.Recovering("QNetworkSession::error")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Error2(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QNetworkSession) Error() QNetworkSession__SessionError {
	defer qt.Recovering("QNetworkSession::error")

	if ptr.Pointer() != nil {
		return QNetworkSession__SessionError(C.QNetworkSession_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ErrorString() string {
	defer qt.Recovering("QNetworkSession::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkSession_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkSession) Ignore() {
	defer qt.Recovering("QNetworkSession::ignore")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Ignore(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) IsOpen() bool {
	defer qt.Recovering("QNetworkSession::isOpen")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkSession) Migrate() {
	defer qt.Recovering("QNetworkSession::migrate")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Migrate(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectNewConfigurationActivated(f func()) {
	defer qt.Recovering("connect QNetworkSession::newConfigurationActivated")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNewConfigurationActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConfigurationActivated", f)
	}
}

func (ptr *QNetworkSession) DisconnectNewConfigurationActivated() {
	defer qt.Recovering("disconnect QNetworkSession::newConfigurationActivated")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNewConfigurationActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConfigurationActivated")
	}
}

//export callbackQNetworkSessionNewConfigurationActivated
func callbackQNetworkSessionNewConfigurationActivated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkSession::newConfigurationActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConfigurationActivated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) NewConfigurationActivated() {
	defer qt.Recovering("QNetworkSession::newConfigurationActivated")

	if ptr.Pointer() != nil {
		C.QNetworkSession_NewConfigurationActivated(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Open() {
	defer qt.Recovering("QNetworkSession::open")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Open(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectOpened(f func()) {
	defer qt.Recovering("connect QNetworkSession::opened")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectOpened(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opened", f)
	}
}

func (ptr *QNetworkSession) DisconnectOpened() {
	defer qt.Recovering("disconnect QNetworkSession::opened")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectOpened(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opened")
	}
}

//export callbackQNetworkSessionOpened
func callbackQNetworkSessionOpened(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QNetworkSession::opened")

	if signal := qt.GetSignal(C.GoString(ptrName), "opened"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QNetworkSession) Opened() {
	defer qt.Recovering("QNetworkSession::opened")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Opened(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Reject() {
	defer qt.Recovering("QNetworkSession::reject")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Reject(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) SessionProperty(key string) *core.QVariant {
	defer qt.Recovering("QNetworkSession::sessionProperty")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QNetworkSession_SessionProperty(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QNetworkSession) SetSessionProperty(key string, value core.QVariant_ITF) {
	defer qt.Recovering("QNetworkSession::setSessionProperty")

	if ptr.Pointer() != nil {
		C.QNetworkSession_SetSessionProperty(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkSession) State() QNetworkSession__State {
	defer qt.Recovering("QNetworkSession::state")

	if ptr.Pointer() != nil {
		return QNetworkSession__State(C.QNetworkSession_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ConnectStateChanged(f func(state QNetworkSession__State)) {
	defer qt.Recovering("connect QNetworkSession::stateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QNetworkSession::stateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQNetworkSessionStateChanged
func callbackQNetworkSessionStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QNetworkSession::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QNetworkSession__State))(QNetworkSession__State(state))
	}

}

func (ptr *QNetworkSession) StateChanged(state QNetworkSession__State) {
	defer qt.Recovering("QNetworkSession::stateChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QNetworkSession) Stop() {
	defer qt.Recovering("QNetworkSession::stop")

	if ptr.Pointer() != nil {
		C.QNetworkSession_Stop(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) UsagePolicies() QNetworkSession__UsagePolicy {
	defer qt.Recovering("QNetworkSession::usagePolicies")

	if ptr.Pointer() != nil {
		return QNetworkSession__UsagePolicy(C.QNetworkSession_UsagePolicies(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ConnectUsagePoliciesChanged(f func(usagePolicies QNetworkSession__UsagePolicy)) {
	defer qt.Recovering("connect QNetworkSession::usagePoliciesChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectUsagePoliciesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "usagePoliciesChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectUsagePoliciesChanged() {
	defer qt.Recovering("disconnect QNetworkSession::usagePoliciesChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectUsagePoliciesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "usagePoliciesChanged")
	}
}

//export callbackQNetworkSessionUsagePoliciesChanged
func callbackQNetworkSessionUsagePoliciesChanged(ptr unsafe.Pointer, ptrName *C.char, usagePolicies C.int) {
	defer qt.Recovering("callback QNetworkSession::usagePoliciesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "usagePoliciesChanged"); signal != nil {
		signal.(func(QNetworkSession__UsagePolicy))(QNetworkSession__UsagePolicy(usagePolicies))
	}

}

func (ptr *QNetworkSession) UsagePoliciesChanged(usagePolicies QNetworkSession__UsagePolicy) {
	defer qt.Recovering("QNetworkSession::usagePoliciesChanged")

	if ptr.Pointer() != nil {
		C.QNetworkSession_UsagePoliciesChanged(ptr.Pointer(), C.int(usagePolicies))
	}
}

func (ptr *QNetworkSession) WaitForOpened(msecs int) bool {
	defer qt.Recovering("QNetworkSession::waitForOpened")

	if ptr.Pointer() != nil {
		return C.QNetworkSession_WaitForOpened(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QNetworkSession) DestroyQNetworkSession() {
	defer qt.Recovering("QNetworkSession::~QNetworkSession")

	if ptr.Pointer() != nil {
		C.QNetworkSession_DestroyQNetworkSession(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QNetworkSession) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQNetworkSessionTimerEvent
func callbackQNetworkSessionTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkSession) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QNetworkSession::timerEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QNetworkSession) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QNetworkSession::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QNetworkSession::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQNetworkSessionChildEvent
func callbackQNetworkSessionChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkSession::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkSession) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QNetworkSession::childEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QNetworkSession) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QNetworkSession::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QNetworkSession) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QNetworkSession::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQNetworkSessionCustomEvent
func callbackQNetworkSessionCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QNetworkSession::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQNetworkSessionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QNetworkSession) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkSession::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QNetworkSession) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QNetworkSession::customEvent")

	if ptr.Pointer() != nil {
		C.QNetworkSession_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
