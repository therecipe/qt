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

type QDBusAbstractInterfaceITF interface {
	core.QObjectITF
	QDBusAbstractInterfacePTR() *QDBusAbstractInterface
}

func PointerFromQDBusAbstractInterface(ptr QDBusAbstractInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractInterfacePTR().Pointer()
	}
	return nil
}

func QDBusAbstractInterfaceFromPointer(ptr unsafe.Pointer) *QDBusAbstractInterface {
	var n = new(QDBusAbstractInterface)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDBusAbstractInterface_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusAbstractInterface) QDBusAbstractInterfacePTR() *QDBusAbstractInterface {
	return ptr
}

func (ptr *QDBusAbstractInterface) Interface() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Interface(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) Path() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Path(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) Service() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Service(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) SetTimeout(timeout int) {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_SetTimeout(C.QtObjectPtr(ptr.Pointer()), C.int(timeout))
	}
}

func (ptr *QDBusAbstractInterface) Timeout() int {
	if ptr.Pointer() != nil {
		return int(C.QDBusAbstractInterface_Timeout(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterface() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterface(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
