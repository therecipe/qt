package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::canceled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCanceled() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::canceled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentCanceled
func callbackQBluetoothServiceDiscoveryAgentCanceled(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::canceled")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "canceled").(func())()
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectFinished() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::finished")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentFinished
func callbackQBluetoothServiceDiscoveryAgentFinished(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::finished")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func NewQBluetoothServiceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::QBluetoothServiceDiscoveryAgent")
		}
	}()

	return NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(core.PointerFromQObject(parent)))
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::QBluetoothServiceDiscoveryAgent")
		}
	}()

	return NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothServiceDiscoveryAgent) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::error")
		}
	}()

	if ptr.Pointer() != nil {
		return QBluetoothServiceDiscoveryAgent__Error(C.QBluetoothServiceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::isActive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddress_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::setRemoteAddress")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_SetRemoteAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuid_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::setUuidFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Start(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::start")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Start(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Stop() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::stop")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothServiceDiscoveryAgent::~QBluetoothServiceDiscoveryAgent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
