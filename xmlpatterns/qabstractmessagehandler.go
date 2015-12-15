package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractMessageHandler struct {
	core.QObject
}

type QAbstractMessageHandler_ITF interface {
	core.QObject_ITF
	QAbstractMessageHandler_PTR() *QAbstractMessageHandler
}

func PointerFromQAbstractMessageHandler(ptr QAbstractMessageHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractMessageHandler_PTR().Pointer()
	}
	return nil
}

func NewQAbstractMessageHandlerFromPointer(ptr unsafe.Pointer) *QAbstractMessageHandler {
	var n = new(QAbstractMessageHandler)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractMessageHandler_") {
		n.SetObjectName("QAbstractMessageHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractMessageHandler) QAbstractMessageHandler_PTR() *QAbstractMessageHandler {
	return ptr
}

func (ptr *QAbstractMessageHandler) DestroyQAbstractMessageHandler() {
	defer qt.Recovering("QAbstractMessageHandler::~QAbstractMessageHandler")

	if ptr.Pointer() != nil {
		C.QAbstractMessageHandler_DestroyQAbstractMessageHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
