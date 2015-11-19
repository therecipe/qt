package dbus

//#include "qdbusabstractinterface.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusAbstractInterface struct {
	core.QObject
}

type QDBusAbstractInterface_ITF interface {
	core.QObject_ITF
	QDBusAbstractInterface_PTR() *QDBusAbstractInterface
}

func PointerFromQDBusAbstractInterface(ptr QDBusAbstractInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterface_PTR().Pointer()
	}
	return nil
}

func NewQDBusAbstractInterfaceFromPointer(ptr unsafe.Pointer) *QDBusAbstractInterface {
	var n = new(QDBusAbstractInterface)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDBusAbstractInterface_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusAbstractInterface) QDBusAbstractInterface_PTR() *QDBusAbstractInterface {
	return ptr
}

func (ptr *QDBusAbstractInterface) Interface() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) Path() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) Service() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) SetTimeout(timeout int) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_SetTimeout(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QDBusAbstractInterface) Timeout() int {
	if ptr.Pointer() != nil {
		return int(C.QDBusAbstractInterface_Timeout(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterface() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
