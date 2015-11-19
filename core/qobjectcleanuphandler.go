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

type QObjectCleanupHandler_ITF interface {
	QObject_ITF
	QObjectCleanupHandler_PTR() *QObjectCleanupHandler
}

func PointerFromQObjectCleanupHandler(ptr QObjectCleanupHandler_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QObjectCleanupHandler_PTR().Pointer()
	}
	return nil
}

func NewQObjectCleanupHandlerFromPointer(ptr unsafe.Pointer) *QObjectCleanupHandler {
	var n = new(QObjectCleanupHandler)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QObjectCleanupHandler_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QObjectCleanupHandler) QObjectCleanupHandler_PTR() *QObjectCleanupHandler {
	return ptr
}

func NewQObjectCleanupHandler() *QObjectCleanupHandler {
	return NewQObjectCleanupHandlerFromPointer(C.QObjectCleanupHandler_NewQObjectCleanupHandler())
}

func (ptr *QObjectCleanupHandler) Add(object QObject_ITF) *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QObjectCleanupHandler_Add(ptr.Pointer(), PointerFromQObject(object)))
	}
	return nil
}

func (ptr *QObjectCleanupHandler) Clear() {
	if ptr.Pointer() != nil {
		C.QObjectCleanupHandler_Clear(ptr.Pointer())
	}
}

func (ptr *QObjectCleanupHandler) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QObjectCleanupHandler_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QObjectCleanupHandler) DestroyQObjectCleanupHandler() {
	if ptr.Pointer() != nil {
		C.QObjectCleanupHandler_DestroyQObjectCleanupHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
