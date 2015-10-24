package core

//#include "qobjectcleanuphandler.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QObjectCleanupHandler struct {
	QObject
}

type QObjectCleanupHandlerITF interface {
	QObjectITF
	QObjectCleanupHandlerPTR() *QObjectCleanupHandler
}

func PointerFromQObjectCleanupHandler(ptr QObjectCleanupHandlerITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QObjectCleanupHandlerPTR().Pointer()
	}
	return nil
}

func QObjectCleanupHandlerFromPointer(ptr unsafe.Pointer) *QObjectCleanupHandler {
	var n = new(QObjectCleanupHandler)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QObjectCleanupHandler_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QObjectCleanupHandler) QObjectCleanupHandlerPTR() *QObjectCleanupHandler {
	return ptr
}

func NewQObjectCleanupHandler() *QObjectCleanupHandler {
	return QObjectCleanupHandlerFromPointer(unsafe.Pointer(C.QObjectCleanupHandler_NewQObjectCleanupHandler()))
}

func (ptr *QObjectCleanupHandler) Add(object QObjectITF) *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QObjectCleanupHandler_Add(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object)))))
	}
	return nil
}

func (ptr *QObjectCleanupHandler) Clear() {
	if ptr.Pointer() != nil {
		C.QObjectCleanupHandler_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QObjectCleanupHandler) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QObjectCleanupHandler_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QObjectCleanupHandler) DestroyQObjectCleanupHandler() {
	if ptr.Pointer() != nil {
		C.QObjectCleanupHandler_DestroyQObjectCleanupHandler(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
