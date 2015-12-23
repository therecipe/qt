package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QEventTransition struct {
	QAbstractTransition
}

type QEventTransition_ITF interface {
	QAbstractTransition_ITF
	QEventTransition_PTR() *QEventTransition
}

func PointerFromQEventTransition(ptr QEventTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEventTransition_PTR().Pointer()
	}
	return nil
}

func NewQEventTransitionFromPointer(ptr unsafe.Pointer) *QEventTransition {
	var n = new(QEventTransition)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QEventTransition_") {
		n.SetObjectName("QEventTransition_" + qt.Identifier())
	}
	return n
}

func (ptr *QEventTransition) QEventTransition_PTR() *QEventTransition {
	return ptr
}

func NewQEventTransition2(object QObject_ITF, ty QEvent__Type, sourceState QState_ITF) *QEventTransition {
	defer qt.Recovering("QEventTransition::QEventTransition")

	return NewQEventTransitionFromPointer(C.QEventTransition_NewQEventTransition2(PointerFromQObject(object), C.int(ty), PointerFromQState(sourceState)))
}

func NewQEventTransition(sourceState QState_ITF) *QEventTransition {
	defer qt.Recovering("QEventTransition::QEventTransition")

	return NewQEventTransitionFromPointer(C.QEventTransition_NewQEventTransition(PointerFromQState(sourceState)))
}

func (ptr *QEventTransition) EventSource() *QObject {
	defer qt.Recovering("QEventTransition::eventSource")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QEventTransition_EventSource(ptr.Pointer()))
	}
	return nil
}

func (ptr *QEventTransition) EventType() QEvent__Type {
	defer qt.Recovering("QEventTransition::eventType")

	if ptr.Pointer() != nil {
		return QEvent__Type(C.QEventTransition_EventType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEventTransition) ConnectOnTransition(f func(event *QEvent)) {
	defer qt.Recovering("connect QEventTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onTransition", f)
	}
}

func (ptr *QEventTransition) DisconnectOnTransition() {
	defer qt.Recovering("disconnect QEventTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onTransition")
	}
}

//export callbackQEventTransitionOnTransition
func callbackQEventTransitionOnTransition(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QEventTransition::onTransition")

	if signal := qt.GetSignal(C.GoString(ptrName), "onTransition"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QEventTransition) SetEventSource(object QObject_ITF) {
	defer qt.Recovering("QEventTransition::setEventSource")

	if ptr.Pointer() != nil {
		C.QEventTransition_SetEventSource(ptr.Pointer(), PointerFromQObject(object))
	}
}

func (ptr *QEventTransition) SetEventType(ty QEvent__Type) {
	defer qt.Recovering("QEventTransition::setEventType")

	if ptr.Pointer() != nil {
		C.QEventTransition_SetEventType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QEventTransition) DestroyQEventTransition() {
	defer qt.Recovering("QEventTransition::~QEventTransition")

	if ptr.Pointer() != nil {
		C.QEventTransition_DestroyQEventTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QEventTransition) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QEventTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QEventTransition) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QEventTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQEventTransitionTimerEvent
func callbackQEventTransitionTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QEventTransition::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QEventTransition) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QEventTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QEventTransition) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QEventTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQEventTransitionChildEvent
func callbackQEventTransitionChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QEventTransition::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QEventTransition) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QEventTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QEventTransition) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QEventTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQEventTransitionCustomEvent
func callbackQEventTransitionCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QEventTransition::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
