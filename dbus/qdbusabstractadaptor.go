package dbus

//#include "qdbusabstractadaptor.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusAbstractAdaptor struct {
	core.QObject
}

type QDBusAbstractAdaptorITF interface {
	core.QObjectITF
	QDBusAbstractAdaptorPTR() *QDBusAbstractAdaptor
}

func PointerFromQDBusAbstractAdaptor(ptr QDBusAbstractAdaptorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusAbstractAdaptorPTR().Pointer()
	}
	return nil
}

func QDBusAbstractAdaptorFromPointer(ptr unsafe.Pointer) *QDBusAbstractAdaptor {
	var n = new(QDBusAbstractAdaptor)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QDBusAbstractAdaptor_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QDBusAbstractAdaptor) QDBusAbstractAdaptorPTR() *QDBusAbstractAdaptor {
	return ptr
}

func (ptr *QDBusAbstractAdaptor) DestroyQDBusAbstractAdaptor() {
	if ptr.Pointer() != nil {
		C.QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
