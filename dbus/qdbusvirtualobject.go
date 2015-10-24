package dbus

//#include "qdbusvirtualobject.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusVirtualObject struct {
	core.QObject
}

type QDBusVirtualObjectITF interface {
	core.QObjectITF
	QDBusVirtualObjectPTR() *QDBusVirtualObject
}

func PointerFromQDBusVirtualObject(ptr QDBusVirtualObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusVirtualObjectPTR().Pointer()
	}
	return nil
}

func QDBusVirtualObjectFromPointer(ptr unsafe.Pointer) *QDBusVirtualObject {
	var n = new(QDBusVirtualObject)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDBusVirtualObject_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusVirtualObject) QDBusVirtualObjectPTR() *QDBusVirtualObject {
	return ptr
}

func (ptr *QDBusVirtualObject) HandleMessage(message QDBusMessageITF, connection QDBusConnectionITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusVirtualObject_HandleMessage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDBusMessage(message)), C.QtObjectPtr(PointerFromQDBusConnection(connection))) != 0
	}
	return false
}

func (ptr *QDBusVirtualObject) Introspect(path string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusVirtualObject_Introspect(C.QtObjectPtr(ptr.Pointer()), C.CString(path)))
	}
	return ""
}

func (ptr *QDBusVirtualObject) DestroyQDBusVirtualObject() {
	if ptr.Pointer() != nil {
		C.QDBusVirtualObject_DestroyQDBusVirtualObject(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
