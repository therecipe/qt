package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QDBusConnection struct {
	ptr unsafe.Pointer
}

type QDBusConnection_ITF interface {
	QDBusConnection_PTR() *QDBusConnection
}

func (p *QDBusConnection) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusConnection) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusConnection(ptr QDBusConnection_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnection_PTR().Pointer()
	}
	return nil
}

func NewQDBusConnectionFromPointer(ptr unsafe.Pointer) *QDBusConnection {
	var n = new(QDBusConnection)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusConnection) QDBusConnection_PTR() *QDBusConnection {
	return ptr
}

//QDBusConnection::BusType
type QDBusConnection__BusType int64

const (
	QDBusConnection__SessionBus    = QDBusConnection__BusType(0)
	QDBusConnection__SystemBus     = QDBusConnection__BusType(1)
	QDBusConnection__ActivationBus = QDBusConnection__BusType(2)
)

//QDBusConnection::ConnectionCapability
type QDBusConnection__ConnectionCapability int64

const (
	QDBusConnection__UnixFileDescriptorPassing = QDBusConnection__ConnectionCapability(0x0001)
)

//QDBusConnection::RegisterOption
type QDBusConnection__RegisterOption int64

const (
	QDBusConnection__ExportAdaptors                = QDBusConnection__RegisterOption(0x01)
	QDBusConnection__ExportScriptableSlots         = QDBusConnection__RegisterOption(0x10)
	QDBusConnection__ExportScriptableSignals       = QDBusConnection__RegisterOption(0x20)
	QDBusConnection__ExportScriptableProperties    = QDBusConnection__RegisterOption(0x40)
	QDBusConnection__ExportScriptableInvokables    = QDBusConnection__RegisterOption(0x80)
	QDBusConnection__ExportScriptableContents      = QDBusConnection__RegisterOption(0xf0)
	QDBusConnection__ExportNonScriptableSlots      = QDBusConnection__RegisterOption(0x100)
	QDBusConnection__ExportNonScriptableSignals    = QDBusConnection__RegisterOption(0x200)
	QDBusConnection__ExportNonScriptableProperties = QDBusConnection__RegisterOption(0x400)
	QDBusConnection__ExportNonScriptableInvokables = QDBusConnection__RegisterOption(0x800)
	QDBusConnection__ExportNonScriptableContents   = QDBusConnection__RegisterOption(0xf00)
	QDBusConnection__ExportAllSlots                = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableSlots | QDBusConnection__ExportNonScriptableSlots)
	QDBusConnection__ExportAllSignals              = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableSignals | QDBusConnection__ExportNonScriptableSignals)
	QDBusConnection__ExportAllProperties           = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableProperties | QDBusConnection__ExportNonScriptableProperties)
	QDBusConnection__ExportAllInvokables           = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableInvokables | QDBusConnection__ExportNonScriptableInvokables)
	QDBusConnection__ExportAllContents             = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableContents | QDBusConnection__ExportNonScriptableContents)
	QDBusConnection__ExportChildObjects            = QDBusConnection__RegisterOption(0x1000)
)

//QDBusConnection::UnregisterMode
type QDBusConnection__UnregisterMode int64

const (
	QDBusConnection__UnregisterNode = QDBusConnection__UnregisterMode(0)
	QDBusConnection__UnregisterTree = QDBusConnection__UnregisterMode(1)
)

func NewQDBusConnection2(other QDBusConnection_ITF) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::QDBusConnection")

	return NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection2(PointerFromQDBusConnection(other)))
}

func NewQDBusConnection(name string) *QDBusConnection {
	defer qt.Recovering("QDBusConnection::QDBusConnection")

	return NewQDBusConnectionFromPointer(C.QDBusConnection_NewQDBusConnection(C.CString(name)))
}

func (ptr *QDBusConnection) BaseService() string {
	defer qt.Recovering("QDBusConnection::baseService")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusConnection_BaseService(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) CallWithCallback(message QDBusMessage_ITF, receiver core.QObject_ITF, returnMethod string, errorMethod string, timeout int) bool {
	defer qt.Recovering("QDBusConnection::callWithCallback")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_CallWithCallback(ptr.Pointer(), PointerFromQDBusMessage(message), core.PointerFromQObject(receiver), C.CString(returnMethod), C.CString(errorMethod), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect(service string, path string, interfa string, name string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::connect")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_Connect(ptr.Pointer(), C.CString(service), C.CString(path), C.CString(interfa), C.CString(name), core.PointerFromQObject(receiver), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect2(service string, path string, interfa string, name string, signature string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::connect")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_Connect2(ptr.Pointer(), C.CString(service), C.CString(path), C.CString(interfa), C.CString(name), C.CString(signature), core.PointerFromQObject(receiver), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect3(service string, path string, interfa string, name string, argumentMatch []string, signature string, receiver core.QObject_ITF, slot string) bool {
	defer qt.Recovering("QDBusConnection::connect")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_Connect3(ptr.Pointer(), C.CString(service), C.CString(path), C.CString(interfa), C.CString(name), C.CString(strings.Join(argumentMatch, ",,,")), C.CString(signature), core.PointerFromQObject(receiver), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QDBusConnection) ConnectionCapabilities() QDBusConnection__ConnectionCapability {
	defer qt.Recovering("QDBusConnection::connectionCapabilities")

	if ptr.Pointer() != nil {
		return QDBusConnection__ConnectionCapability(C.QDBusConnection_ConnectionCapabilities(ptr.Pointer()))
	}
	return 0
}

func QDBusConnection_DisconnectFromBus(name string) {
	defer qt.Recovering("QDBusConnection::disconnectFromBus")

	C.QDBusConnection_QDBusConnection_DisconnectFromBus(C.CString(name))
}

func QDBusConnection_DisconnectFromPeer(name string) {
	defer qt.Recovering("QDBusConnection::disconnectFromPeer")

	C.QDBusConnection_QDBusConnection_DisconnectFromPeer(C.CString(name))
}

func (ptr *QDBusConnection) Interface() *QDBusConnectionInterface {
	defer qt.Recovering("QDBusConnection::interface")

	if ptr.Pointer() != nil {
		return NewQDBusConnectionInterfaceFromPointer(C.QDBusConnection_Interface(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDBusConnection) IsConnected() bool {
	defer qt.Recovering("QDBusConnection::isConnected")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func QDBusConnection_LocalMachineId() *core.QByteArray {
	defer qt.Recovering("QDBusConnection::localMachineId")

	return core.NewQByteArrayFromPointer(C.QDBusConnection_QDBusConnection_LocalMachineId())
}

func (ptr *QDBusConnection) Name() string {
	defer qt.Recovering("QDBusConnection::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusConnection_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusConnection) ObjectRegisteredAt(path string) *core.QObject {
	defer qt.Recovering("QDBusConnection::objectRegisteredAt")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QDBusConnection_ObjectRegisteredAt(ptr.Pointer(), C.CString(path)))
	}
	return nil
}

func (ptr *QDBusConnection) RegisterObject(path string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
	defer qt.Recovering("QDBusConnection::registerObject")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_RegisterObject(ptr.Pointer(), C.CString(path), core.PointerFromQObject(object), C.int(options)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterObject2(path string, interfa string, object core.QObject_ITF, options QDBusConnection__RegisterOption) bool {
	defer qt.Recovering("QDBusConnection::registerObject")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_RegisterObject2(ptr.Pointer(), C.CString(path), C.CString(interfa), core.PointerFromQObject(object), C.int(options)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterService(serviceName string) bool {
	defer qt.Recovering("QDBusConnection::registerService")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_RegisterService(ptr.Pointer(), C.CString(serviceName)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Send(message QDBusMessage_ITF) bool {
	defer qt.Recovering("QDBusConnection::send")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_Send(ptr.Pointer(), PointerFromQDBusMessage(message)) != 0
	}
	return false
}

func (ptr *QDBusConnection) UnregisterObject(path string, mode QDBusConnection__UnregisterMode) {
	defer qt.Recovering("QDBusConnection::unregisterObject")

	if ptr.Pointer() != nil {
		C.QDBusConnection_UnregisterObject(ptr.Pointer(), C.CString(path), C.int(mode))
	}
}

func (ptr *QDBusConnection) UnregisterService(serviceName string) bool {
	defer qt.Recovering("QDBusConnection::unregisterService")

	if ptr.Pointer() != nil {
		return C.QDBusConnection_UnregisterService(ptr.Pointer(), C.CString(serviceName)) != 0
	}
	return false
}

func (ptr *QDBusConnection) DestroyQDBusConnection() {
	defer qt.Recovering("QDBusConnection::~QDBusConnection")

	if ptr.Pointer() != nil {
		C.QDBusConnection_DestroyQDBusConnection(ptr.Pointer())
	}
}
