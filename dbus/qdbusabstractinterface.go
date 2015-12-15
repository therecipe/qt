package dbus

//#include "dbus.h"
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
	for len(n.ObjectName()) < len("QDBusAbstractInterface_") {
		n.SetObjectName("QDBusAbstractInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusAbstractInterface) QDBusAbstractInterface_PTR() *QDBusAbstractInterface {
	return ptr
}

func (ptr *QDBusAbstractInterface) Interface() string {
	defer qt.Recovering("QDBusAbstractInterface::interface")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Interface(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) IsValid() bool {
	defer qt.Recovering("QDBusAbstractInterface::isValid")

	if ptr.Pointer() != nil {
		return C.QDBusAbstractInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusAbstractInterface) Path() string {
	defer qt.Recovering("QDBusAbstractInterface::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) Service() string {
	defer qt.Recovering("QDBusAbstractInterface::service")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusAbstractInterface_Service(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusAbstractInterface) SetTimeout(timeout int) {
	defer qt.Recovering("QDBusAbstractInterface::setTimeout")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_SetTimeout(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QDBusAbstractInterface) Timeout() int {
	defer qt.Recovering("QDBusAbstractInterface::timeout")

	if ptr.Pointer() != nil {
		return int(C.QDBusAbstractInterface_Timeout(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDBusAbstractInterface) DestroyQDBusAbstractInterface() {
	defer qt.Recovering("QDBusAbstractInterface::~QDBusAbstractInterface")

	if ptr.Pointer() != nil {
		C.QDBusAbstractInterface_DestroyQDBusAbstractInterface(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
