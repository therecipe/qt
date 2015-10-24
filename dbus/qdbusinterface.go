package dbus

//#include "qdbusinterface.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusInterface struct {
	QDBusAbstractInterface
}

type QDBusInterfaceITF interface {
	QDBusAbstractInterfaceITF
	QDBusInterfacePTR() *QDBusInterface
}

func PointerFromQDBusInterface(ptr QDBusInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusInterfacePTR().Pointer()
	}
	return nil
}

func QDBusInterfaceFromPointer(ptr unsafe.Pointer) *QDBusInterface {
	var n = new(QDBusInterface)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDBusInterface_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusInterface) QDBusInterfacePTR() *QDBusInterface {
	return ptr
}

func NewQDBusInterface(service string, path string, interfa string, connection QDBusConnectionITF, parent core.QObjectITF) *QDBusInterface {
	return QDBusInterfaceFromPointer(unsafe.Pointer(C.QDBusInterface_NewQDBusInterface(C.CString(service), C.CString(path), C.CString(interfa), C.QtObjectPtr(PointerFromQDBusConnection(connection)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QDBusInterface) DestroyQDBusInterface() {
	if ptr.Pointer() != nil {
		C.QDBusInterface_DestroyQDBusInterface(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
