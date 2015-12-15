package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusServer struct {
	core.QObject
}

type QDBusServer_ITF interface {
	core.QObject_ITF
	QDBusServer_PTR() *QDBusServer
}

func PointerFromQDBusServer(ptr QDBusServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServer_PTR().Pointer()
	}
	return nil
}

func NewQDBusServerFromPointer(ptr unsafe.Pointer) *QDBusServer {
	var n = new(QDBusServer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusServer_") {
		n.SetObjectName("QDBusServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusServer) QDBusServer_PTR() *QDBusServer {
	return ptr
}

func NewQDBusServer2(parent core.QObject_ITF) *QDBusServer {
	defer qt.Recovering("QDBusServer::QDBusServer")

	return NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer2(core.PointerFromQObject(parent)))
}

func NewQDBusServer(address string, parent core.QObject_ITF) *QDBusServer {
	defer qt.Recovering("QDBusServer::QDBusServer")

	return NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer(C.CString(address), core.PointerFromQObject(parent)))
}

func (ptr *QDBusServer) Address() string {
	defer qt.Recovering("QDBusServer::address")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusServer_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusServer) IsAnonymousAuthenticationAllowed() bool {
	defer qt.Recovering("QDBusServer::isAnonymousAuthenticationAllowed")

	if ptr.Pointer() != nil {
		return C.QDBusServer_IsAnonymousAuthenticationAllowed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) IsConnected() bool {
	defer qt.Recovering("QDBusServer::isConnected")

	if ptr.Pointer() != nil {
		return C.QDBusServer_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) SetAnonymousAuthenticationAllowed(value bool) {
	defer qt.Recovering("QDBusServer::setAnonymousAuthenticationAllowed")

	if ptr.Pointer() != nil {
		C.QDBusServer_SetAnonymousAuthenticationAllowed(ptr.Pointer(), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QDBusServer) DestroyQDBusServer() {
	defer qt.Recovering("QDBusServer::~QDBusServer")

	if ptr.Pointer() != nil {
		C.QDBusServer_DestroyQDBusServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
