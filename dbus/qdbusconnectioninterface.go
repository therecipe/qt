package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QDBusConnectionInterface struct {
	QDBusAbstractInterface
}

type QDBusConnectionInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusConnectionInterface_PTR() *QDBusConnectionInterface
}

func PointerFromQDBusConnectionInterface(ptr QDBusConnectionInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnectionInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusConnectionInterfaceFromPointer(ptr unsafe.Pointer) *QDBusConnectionInterface {
	var n = new(QDBusConnectionInterface)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusConnectionInterface_") {
		n.SetObjectName("QDBusConnectionInterface_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusConnectionInterface) QDBusConnectionInterface_PTR() *QDBusConnectionInterface {
	return ptr
}

//QDBusConnectionInterface::RegisterServiceReply
type QDBusConnectionInterface__RegisterServiceReply int64

const (
	QDBusConnectionInterface__ServiceNotRegistered = QDBusConnectionInterface__RegisterServiceReply(0)
	QDBusConnectionInterface__ServiceRegistered    = QDBusConnectionInterface__RegisterServiceReply(1)
	QDBusConnectionInterface__ServiceQueued        = QDBusConnectionInterface__RegisterServiceReply(2)
)

//QDBusConnectionInterface::ServiceQueueOptions
type QDBusConnectionInterface__ServiceQueueOptions int64

const (
	QDBusConnectionInterface__DontQueueService       = QDBusConnectionInterface__ServiceQueueOptions(0)
	QDBusConnectionInterface__QueueService           = QDBusConnectionInterface__ServiceQueueOptions(1)
	QDBusConnectionInterface__ReplaceExistingService = QDBusConnectionInterface__ServiceQueueOptions(2)
)

//QDBusConnectionInterface::ServiceReplacementOptions
type QDBusConnectionInterface__ServiceReplacementOptions int64

const (
	QDBusConnectionInterface__DontAllowReplacement = QDBusConnectionInterface__ServiceReplacementOptions(0)
	QDBusConnectionInterface__AllowReplacement     = QDBusConnectionInterface__ServiceReplacementOptions(1)
)

func (ptr *QDBusConnectionInterface) ConnectServiceRegistered(f func(serviceName string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusConnectionInterface::serviceRegistered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceRegistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceRegistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceRegistered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusConnectionInterface::serviceRegistered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceRegistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceRegistered")
	}
}

//export callbackQDBusConnectionInterfaceServiceRegistered
func callbackQDBusConnectionInterfaceServiceRegistered(ptrName *C.char, serviceName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusConnectionInterface::serviceRegistered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "serviceRegistered").(func(string))(C.GoString(serviceName))
}

func (ptr *QDBusConnectionInterface) ConnectServiceUnregistered(f func(serviceName string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusConnectionInterface::serviceUnregistered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_ConnectServiceUnregistered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceUnregistered", f)
	}
}

func (ptr *QDBusConnectionInterface) DisconnectServiceUnregistered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusConnectionInterface::serviceUnregistered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QDBusConnectionInterface_DisconnectServiceUnregistered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceUnregistered")
	}
}

//export callbackQDBusConnectionInterfaceServiceUnregistered
func callbackQDBusConnectionInterfaceServiceUnregistered(ptrName *C.char, serviceName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QDBusConnectionInterface::serviceUnregistered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "serviceUnregistered").(func(string))(C.GoString(serviceName))
}
