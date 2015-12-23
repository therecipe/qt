package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothDeviceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothDeviceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent
}

func PointerFromQBluetoothDeviceDiscoveryAgent(ptr QBluetoothDeviceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothDeviceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceDiscoveryAgent {
	var n = new(QBluetoothDeviceDiscoveryAgent)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothDeviceDiscoveryAgent_") {
		n.SetObjectName("QBluetoothDeviceDiscoveryAgent_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothDeviceDiscoveryAgent) QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent {
	return ptr
}

//QBluetoothDeviceDiscoveryAgent::Error
type QBluetoothDeviceDiscoveryAgent__Error int64

const (
	QBluetoothDeviceDiscoveryAgent__NoError                      = QBluetoothDeviceDiscoveryAgent__Error(0)
	QBluetoothDeviceDiscoveryAgent__InputOutputError             = QBluetoothDeviceDiscoveryAgent__Error(1)
	QBluetoothDeviceDiscoveryAgent__PoweredOffError              = QBluetoothDeviceDiscoveryAgent__Error(2)
	QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError = QBluetoothDeviceDiscoveryAgent__Error(3)
	QBluetoothDeviceDiscoveryAgent__UnsupportedPlatformError     = QBluetoothDeviceDiscoveryAgent__Error(4)
	QBluetoothDeviceDiscoveryAgent__UnknownError                 = QBluetoothDeviceDiscoveryAgent__Error(100)
)

//QBluetoothDeviceDiscoveryAgent::InquiryType
type QBluetoothDeviceDiscoveryAgent__InquiryType int64

const (
	QBluetoothDeviceDiscoveryAgent__GeneralUnlimitedInquiry = QBluetoothDeviceDiscoveryAgent__InquiryType(0)
	QBluetoothDeviceDiscoveryAgent__LimitedInquiry          = QBluetoothDeviceDiscoveryAgent__InquiryType(1)
)

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCanceled(f func()) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCanceled() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentCanceled
func callbackQBluetoothDeviceDiscoveryAgentCanceled(ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::canceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectError2(f func(error QBluetoothDeviceDiscoveryAgent__Error)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentError2
func callbackQBluetoothDeviceDiscoveryAgentError2(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothDeviceDiscoveryAgent__Error))(QBluetoothDeviceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectFinished(f func()) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentFinished
func callbackQBluetoothDeviceDiscoveryAgentFinished(ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func NewQBluetoothDeviceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::QBluetoothDeviceDiscoveryAgent")

	return NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(core.PointerFromQObject(parent)))
}

func NewQBluetoothDeviceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::QBluetoothDeviceDiscoveryAgent")

	return NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error() QBluetoothDeviceDiscoveryAgent__Error {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__Error(C.QBluetoothDeviceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ErrorString() string {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothDeviceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceDiscoveryAgent) InquiryType() QBluetoothDeviceDiscoveryAgent__InquiryType {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::inquiryType")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__InquiryType(C.QBluetoothDeviceDiscoveryAgent_InquiryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) IsActive() bool {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::isActive")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetInquiryType(ty QBluetoothDeviceDiscoveryAgent__InquiryType) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::setInquiryType")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_SetInquiryType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::start")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Start(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Stop() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgent() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::~QBluetoothDeviceDiscoveryAgent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentTimerEvent
func callbackQBluetoothDeviceDiscoveryAgentTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentChildEvent
func callbackQBluetoothDeviceDiscoveryAgentChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentCustomEvent
func callbackQBluetoothDeviceDiscoveryAgentCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
