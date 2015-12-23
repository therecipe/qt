package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothServiceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothServiceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent
}

func PointerFromQBluetoothServiceDiscoveryAgent(ptr QBluetoothServiceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServiceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothServiceDiscoveryAgent {
	var n = new(QBluetoothServiceDiscoveryAgent)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothServiceDiscoveryAgent_") {
		n.SetObjectName("QBluetoothServiceDiscoveryAgent_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothServiceDiscoveryAgent) QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent {
	return ptr
}

//QBluetoothServiceDiscoveryAgent::DiscoveryMode
type QBluetoothServiceDiscoveryAgent__DiscoveryMode int64

const (
	QBluetoothServiceDiscoveryAgent__MinimalDiscovery = QBluetoothServiceDiscoveryAgent__DiscoveryMode(0)
	QBluetoothServiceDiscoveryAgent__FullDiscovery    = QBluetoothServiceDiscoveryAgent__DiscoveryMode(1)
)

//QBluetoothServiceDiscoveryAgent::Error
type QBluetoothServiceDiscoveryAgent__Error int64

const (
	QBluetoothServiceDiscoveryAgent__NoError                      = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__NoError)
	QBluetoothServiceDiscoveryAgent__InputOutputError             = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InputOutputError)
	QBluetoothServiceDiscoveryAgent__PoweredOffError              = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__PoweredOffError)
	QBluetoothServiceDiscoveryAgent__InvalidBluetoothAdapterError = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError)
	QBluetoothServiceDiscoveryAgent__UnknownError                 = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__UnknownError)
)

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCanceled(f func()) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCanceled() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentCanceled
func callbackQBluetoothServiceDiscoveryAgentCanceled(ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::canceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectError2(f func(error QBluetoothServiceDiscoveryAgent__Error)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentError2
func callbackQBluetoothServiceDiscoveryAgentError2(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothServiceDiscoveryAgent__Error))(QBluetoothServiceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentFinished
func callbackQBluetoothServiceDiscoveryAgentFinished(ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func NewQBluetoothServiceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::QBluetoothServiceDiscoveryAgent")

	return NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(core.PointerFromQObject(parent)))
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::QBluetoothServiceDiscoveryAgent")

	return NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothServiceDiscoveryAgent) Clear() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::clear")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		return QBluetoothServiceDiscoveryAgent__Error(C.QBluetoothServiceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::isActive")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddress_ITF) bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::setRemoteAddress")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_SetRemoteAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::setUuidFilter")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Start(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::start")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Start(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Stop() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::~QBluetoothServiceDiscoveryAgent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentTimerEvent
func callbackQBluetoothServiceDiscoveryAgentTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentChildEvent
func callbackQBluetoothServiceDiscoveryAgentChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentCustomEvent
func callbackQBluetoothServiceDiscoveryAgentCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
