package dbus

//#include "qdbusserver.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusServer struct {
	core.QObject
}

type QDBusServerITF interface {
	core.QObjectITF
	QDBusServerPTR() *QDBusServer
}

func PointerFromQDBusServer(ptr QDBusServerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServerPTR().Pointer()
	}
	return nil
}

func QDBusServerFromPointer(ptr unsafe.Pointer) *QDBusServer {
	var n = new(QDBusServer)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDBusServer_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusServer) QDBusServerPTR() *QDBusServer {
	return ptr
}

func NewQDBusServer2(parent core.QObjectITF) *QDBusServer {
	return QDBusServerFromPointer(unsafe.Pointer(C.QDBusServer_NewQDBusServer2(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQDBusServer(address string, parent core.QObjectITF) *QDBusServer {
	return QDBusServerFromPointer(unsafe.Pointer(C.QDBusServer_NewQDBusServer(C.CString(address), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QDBusServer) Address() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusServer_Address(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusServer) IsAnonymousAuthenticationAllowed() bool {
	if ptr.Pointer() != nil {
		return C.QDBusServer_IsAnonymousAuthenticationAllowed(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusServer) IsConnected() bool {
	if ptr.Pointer() != nil {
		return C.QDBusServer_IsConnected(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusServer) SetAnonymousAuthenticationAllowed(value bool) {
	if ptr.Pointer() != nil {
		C.QDBusServer_SetAnonymousAuthenticationAllowed(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QDBusServer) DestroyQDBusServer() {
	if ptr.Pointer() != nil {
		C.QDBusServer_DestroyQDBusServer(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
