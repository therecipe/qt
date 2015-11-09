package bluetooth

//#include "qbluetoothdevicediscoveryagent.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QBluetoothDeviceDiscoveryAgent_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCanceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentCanceled
func callbackQBluetoothDeviceDiscoveryAgentCanceled(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "canceled").(func())()
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothDeviceDiscoveryAgentFinished
func callbackQBluetoothDeviceDiscoveryAgentFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func NewQBluetoothDeviceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	return NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(core.PointerFromQObject(parent)))
}

func NewQBluetoothDeviceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	return NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error() QBluetoothDeviceDiscoveryAgent__Error {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__Error(C.QBluetoothDeviceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothDeviceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceDiscoveryAgent) InquiryType() QBluetoothDeviceDiscoveryAgent__InquiryType {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__InquiryType(C.QBluetoothDeviceDiscoveryAgent_InquiryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetInquiryType(ty QBluetoothDeviceDiscoveryAgent__InquiryType) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_SetInquiryType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Start(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Stop() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgent() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
