package xmlpatterns

//#include "qabstractmessagehandler.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractMessageHandler struct {
	core.QObject
}

type QAbstractMessageHandlerITF interface {
	core.QObjectITF
	QAbstractMessageHandlerPTR() *QAbstractMessageHandler
}

func PointerFromQAbstractMessageHandler(ptr QAbstractMessageHandlerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractMessageHandlerPTR().Pointer()
	}
	return nil
}

func QAbstractMessageHandlerFromPointer(ptr unsafe.Pointer) *QAbstractMessageHandler {
	var n = new(QAbstractMessageHandler)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAbstractMessageHandler_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAbstractMessageHandler) QAbstractMessageHandlerPTR() *QAbstractMessageHandler {
	return ptr
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandler() {
	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandler(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
