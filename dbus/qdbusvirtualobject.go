package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusVirtualObject struct {
	core.QObject
}

type QDBusVirtualObject_ITF interface {
	core.QObject_ITF
	QDBusVirtualObject_PTR() *QDBusVirtualObject
}

func PointerFromQDBusVirtualObject(ptr QDBusVirtualObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVirtualObject_PTR().Pointer()
	}
	return nil
}

func NewQDBusVirtualObjectFromPointer(ptr unsafe.Pointer) *QDBusVirtualObject {
	var n = new(QDBusVirtualObject)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusVirtualObject_") {
		n.SetObjectName("QDBusVirtualObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusVirtualObject) QDBusVirtualObject_PTR() *QDBusVirtualObject {
	return ptr
}

func (ptr *QDBusVirtualObject) HandleMessage(message QDBusMessage_ITF, connection QDBusConnection_ITF) bool {
	defer qt.Recovering("QDBusVirtualObject::handleMessage")

	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_HandleMessage(ptr.Pointer(), PointerFromQDBusMessage(message), PointerFromQDBusConnection(connection)) != 0
	}
	return false
}

func (ptr *QDBusVirtualObject) Introspect(path string) string {
	defer qt.Recovering("QDBusVirtualObject::introspect")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusVirtualObject_Introspect(ptr.Pointer(), C.CString(path)))
	}
	return ""
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObject() {
	defer qt.Recovering("QDBusVirtualObject::~QDBusVirtualObject")

	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DestroyQDBusVirtualObject(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
