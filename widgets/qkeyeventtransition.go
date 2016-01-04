package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QKeyEventTransition struct {
	core.QEventTransition
}

type QKeyEventTransition_ITF interface {
	core.QEventTransition_ITF
	QKeyEventTransition_PTR() *QKeyEventTransition
}

func PointerFromQKeyEventTransition(ptr QKeyEventTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QKeyEventTransition_PTR().Pointer()
	}
	return nil
}

func NewQKeyEventTransitionFromPointer(ptr unsafe.Pointer) *QKeyEventTransition {
	var n = new(QKeyEventTransition)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QKeyEventTransition_") {
		n.SetObjectName("QKeyEventTransition_" + qt.Identifier())
	}
	return n
}

func (ptr *QKeyEventTransition) QKeyEventTransition_PTR() *QKeyEventTransition {
	return ptr
}

func NewQKeyEventTransition2(object core.QObject_ITF, ty core.QEvent__Type, key int, sourceState core.QState_ITF) *QKeyEventTransition {
	defer qt.Recovering("QKeyEventTransition::QKeyEventTransition")

	return NewQKeyEventTransitionFromPointer(C.QKeyEventTransition_NewQKeyEventTransition2(core.PointerFromQObject(object), C.int(ty), C.int(key), core.PointerFromQState(sourceState)))
}

func NewQKeyEventTransition(sourceState core.QState_ITF) *QKeyEventTransition {
	defer qt.Recovering("QKeyEventTransition::QKeyEventTransition")

	return NewQKeyEventTransitionFromPointer(C.QKeyEventTransition_NewQKeyEventTransition(core.PointerFromQState(sourceState)))
}

func (ptr *QKeyEventTransition) EventTest(event core.QEvent_ITF) bool {
	defer qt.Recovering("QKeyEventTransition::eventTest")

	if ptr.Pointer() != nil {
		return C.QKeyEventTransition_EventTest(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QKeyEventTransition) Key() int {
	defer qt.Recovering("QKeyEventTransition::key")

	if ptr.Pointer() != nil {
		return int(C.QKeyEventTransition_Key(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEventTransition) ModifierMask() core.Qt__KeyboardModifier {
	defer qt.Recovering("QKeyEventTransition::modifierMask")

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QKeyEventTransition_ModifierMask(ptr.Pointer()))
	}
	return 0
}

func (ptr *QKeyEventTransition) ConnectOnTransition(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QKeyEventTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onTransition", f)
	}
}

func (ptr *QKeyEventTransition) DisconnectOnTransition() {
	defer qt.Recovering("disconnect QKeyEventTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onTransition")
	}
}

//export callbackQKeyEventTransitionOnTransition
func callbackQKeyEventTransitionOnTransition(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeyEventTransition::onTransition")

	if signal := qt.GetSignal(C.GoString(ptrName), "onTransition"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	}

}

func (ptr *QKeyEventTransition) OnTransition(event core.QEvent_ITF) {
	defer qt.Recovering("QKeyEventTransition::onTransition")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_OnTransition(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeyEventTransition) OnTransitionDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QKeyEventTransition::onTransition")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_OnTransitionDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeyEventTransition) SetKey(key int) {
	defer qt.Recovering("QKeyEventTransition::setKey")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_SetKey(ptr.Pointer(), C.int(key))
	}
}

func (ptr *QKeyEventTransition) SetModifierMask(modifierMask core.Qt__KeyboardModifier) {
	defer qt.Recovering("QKeyEventTransition::setModifierMask")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_SetModifierMask(ptr.Pointer(), C.int(modifierMask))
	}
}

func (ptr *QKeyEventTransition) DestroyQKeyEventTransition() {
	defer qt.Recovering("QKeyEventTransition::~QKeyEventTransition")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_DestroyQKeyEventTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QKeyEventTransition) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QKeyEventTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QKeyEventTransition) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QKeyEventTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQKeyEventTransitionTimerEvent
func callbackQKeyEventTransitionTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeyEventTransition::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQKeyEventTransitionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QKeyEventTransition) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QKeyEventTransition::timerEvent")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QKeyEventTransition) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QKeyEventTransition::timerEvent")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QKeyEventTransition) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QKeyEventTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QKeyEventTransition) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QKeyEventTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQKeyEventTransitionChildEvent
func callbackQKeyEventTransitionChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeyEventTransition::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQKeyEventTransitionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QKeyEventTransition) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QKeyEventTransition::childEvent")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QKeyEventTransition) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QKeyEventTransition::childEvent")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QKeyEventTransition) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QKeyEventTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QKeyEventTransition) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QKeyEventTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQKeyEventTransitionCustomEvent
func callbackQKeyEventTransitionCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QKeyEventTransition::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQKeyEventTransitionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QKeyEventTransition) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QKeyEventTransition::customEvent")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QKeyEventTransition) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QKeyEventTransition::customEvent")

	if ptr.Pointer() != nil {
		C.QKeyEventTransition_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
