package core

//#include "core.h"
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
	for len(n.ObjectName()) < len("QObjectCleanupHandler_") {
		n.SetObjectName("QObjectCleanupHandler_" + qt.Identifier())
	}
	return n
}

func (ptr *QObjectCleanupHandler) QObjectCleanupHandler_PTR() *QObjectCleanupHandler {
	return ptr
}

func NewQObjectCleanupHandler() *QObjectCleanupHandler {
	defer qt.Recovering("QObjectCleanupHandler::QObjectCleanupHandler")

	return NewQObjectCleanupHandlerFromPointer(C.QObjectCleanupHandler_NewQObjectCleanupHandler())
}

func (ptr *QObjectCleanupHandler) Add(object QObject_ITF) *QObject {
	defer qt.Recovering("QObjectCleanupHandler::add")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QObjectCleanupHandler_Add(ptr.Pointer(), PointerFromQObject(object)))
	}
	return nil
}

func (ptr *QObjectCleanupHandler) Clear() {
	defer qt.Recovering("QObjectCleanupHandler::clear")

	if ptr.Pointer() != nil {
		C.QObjectCleanupHandler_Clear(ptr.Pointer())
	}
}

func (ptr *QObjectCleanupHandler) IsEmpty() bool {
	defer qt.Recovering("QObjectCleanupHandler::isEmpty")

	if ptr.Pointer() != nil {
		return C.QObjectCleanupHandler_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QObjectCleanupHandler) DestroyQObjectCleanupHandler() {
	defer qt.Recovering("QObjectCleanupHandler::~QObjectCleanupHandler")

	if ptr.Pointer() != nil {
		C.QObjectCleanupHandler_DestroyQObjectCleanupHandler(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QObjectCleanupHandler) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QObjectCleanupHandler::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QObjectCleanupHandler) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QObjectCleanupHandler::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQObjectCleanupHandlerTimerEvent
func callbackQObjectCleanupHandlerTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QObjectCleanupHandler::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QObjectCleanupHandler) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QObjectCleanupHandler::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QObjectCleanupHandler) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QObjectCleanupHandler::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQObjectCleanupHandlerChildEvent
func callbackQObjectCleanupHandlerChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QObjectCleanupHandler::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QObjectCleanupHandler) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QObjectCleanupHandler::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QObjectCleanupHandler) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QObjectCleanupHandler::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQObjectCleanupHandlerCustomEvent
func callbackQObjectCleanupHandlerCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QObjectCleanupHandler::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
