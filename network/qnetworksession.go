package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
		n.SetObjectName("QNetworkSession_" + qt.RandomIdentifier())
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::QNetworkSession")
		}
	}()

	return NewQNetworkSessionFromPointer(C.QNetworkSession_NewQNetworkSession(PointerFromQNetworkConfiguration(connectionConfig), core.PointerFromQObject(parent)))
}

func (ptr *QNetworkSession) Accept() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::accept")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_Accept(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) Close() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::close")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_Close(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectClosed(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::closed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectClosed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "closed", f)
	}
}

func (ptr *QNetworkSession) DisconnectClosed() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::closed")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectClosed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "closed")
	}
}

//export callbackQNetworkSessionClosed
func callbackQNetworkSessionClosed(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::closed")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "closed").(func())()
}

func (ptr *QNetworkSession) Error() QNetworkSession__SessionError {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QNetworkSession__SessionError(C.QNetworkSession_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkSession_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkSession) Ignore() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::ignore")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_Ignore(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) IsOpen() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::isOpen")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNetworkSession_IsOpen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkSession) Migrate() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::migrate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_Migrate(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectNewConfigurationActivated(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::newConfigurationActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectNewConfigurationActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConfigurationActivated", f)
	}
}

func (ptr *QNetworkSession) DisconnectNewConfigurationActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::newConfigurationActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectNewConfigurationActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConfigurationActivated")
	}
}

//export callbackQNetworkSessionNewConfigurationActivated
func callbackQNetworkSessionNewConfigurationActivated(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::newConfigurationActivated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "newConfigurationActivated").(func())()
}

func (ptr *QNetworkSession) Open() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::open")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_Open(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) ConnectOpened(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::opened")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectOpened(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opened", f)
	}
}

func (ptr *QNetworkSession) DisconnectOpened() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::opened")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectOpened(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opened")
	}
}

//export callbackQNetworkSessionOpened
func callbackQNetworkSessionOpened(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::opened")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "opened").(func())()
}

func (ptr *QNetworkSession) Reject() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::reject")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_Reject(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) SessionProperty(key string) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::sessionProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QNetworkSession_SessionProperty(ptr.Pointer(), C.CString(key)))
	}
	return nil
}

func (ptr *QNetworkSession) SetSessionProperty(key string, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::setSessionProperty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_SetSessionProperty(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QNetworkSession) ConnectStateChanged(f func(state QNetworkSession__State)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectStateChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::stateChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

//export callbackQNetworkSessionStateChanged
func callbackQNetworkSessionStateChanged(ptrName *C.char, state C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::stateChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "stateChanged").(func(QNetworkSession__State))(QNetworkSession__State(state))
}

func (ptr *QNetworkSession) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_Stop(ptr.Pointer())
	}
}

func (ptr *QNetworkSession) UsagePolicies() QNetworkSession__UsagePolicy {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::usagePolicies")
		}
	}()

	if ptr.Pointer() != nil {
		return QNetworkSession__UsagePolicy(C.QNetworkSession_UsagePolicies(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkSession) ConnectUsagePoliciesChanged(f func(usagePolicies QNetworkSession__UsagePolicy)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::usagePoliciesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_ConnectUsagePoliciesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "usagePoliciesChanged", f)
	}
}

func (ptr *QNetworkSession) DisconnectUsagePoliciesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::usagePoliciesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_DisconnectUsagePoliciesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "usagePoliciesChanged")
	}
}

//export callbackQNetworkSessionUsagePoliciesChanged
func callbackQNetworkSessionUsagePoliciesChanged(ptrName *C.char, usagePolicies C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::usagePoliciesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "usagePoliciesChanged").(func(QNetworkSession__UsagePolicy))(QNetworkSession__UsagePolicy(usagePolicies))
}

func (ptr *QNetworkSession) WaitForOpened(msecs int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::waitForOpened")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNetworkSession_WaitForOpened(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QNetworkSession) DestroyQNetworkSession() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkSession::~QNetworkSession")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkSession_DestroyQNetworkSession(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
