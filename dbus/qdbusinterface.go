package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusInterface struct {
	QDBusAbstractInterface
}

type QDBusInterface_ITF interface {
	QDBusAbstractInterface_ITF
	QDBusInterface_PTR() *QDBusInterface
}

func PointerFromQDBusInterface(ptr QDBusInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusInterfaceFromPointer(ptr unsafe.Pointer) *QDBusInterface {
	var n = new(QDBusInterface)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusInterface_") {
		n.SetObjectName("QDBusInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusInterface) QDBusInterface_PTR() *QDBusInterface {
	return ptr
}

func NewQDBusInterface(service string, path string, interfa string, connection QDBusConnection_ITF, parent core.QObject_ITF) *QDBusInterface {
	defer qt.Recovering("QDBusInterface::QDBusInterface")

	return NewQDBusInterfaceFromPointer(C.QDBusInterface_NewQDBusInterface(C.CString(service), C.CString(path), C.CString(interfa), PointerFromQDBusConnection(connection), core.PointerFromQObject(parent)))
}

func (ptr *QDBusInterface) DestroyQDBusInterface() {
	defer qt.Recovering("QDBusInterface::~QDBusInterface")

	if ptr.Pointer() != nil {
		C.QDBusInterface_DestroyQDBusInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
