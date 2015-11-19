package bluetooth

//#include "qbluetoothservicediscoveryagent.h"
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
	if n.ObjectName() == "" {
		n.SetObjectName("QBluetoothServiceDiscoveryAgent_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCanceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentCanceled
func callbackQBluetoothServiceDiscoveryAgentCanceled(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "canceled").(func())()
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentFinished
func callbackQBluetoothServiceDiscoveryAgentFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func NewQBluetoothServiceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	return NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(core.PointerFromQObject(parent)))
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	return NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothServiceDiscoveryAgent) Clear() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {
	if ptr.Pointer() != nil {
		return QBluetoothServiceDiscoveryAgent__Error(C.QBluetoothServiceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_SetRemoteAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Start(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Start(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Stop() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
