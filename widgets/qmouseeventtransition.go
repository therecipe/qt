package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMouseEventTransition struct {
	core.QEventTransition
}

type QMouseEventTransition_ITF interface {
	core.QEventTransition_ITF
	QMouseEventTransition_PTR() *QMouseEventTransition
}

func PointerFromQMouseEventTransition(ptr QMouseEventTransition_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMouseEventTransition_PTR().Pointer()
	}
	return nil
}

func NewQMouseEventTransitionFromPointer(ptr unsafe.Pointer) *QMouseEventTransition {
	var n = new(QMouseEventTransition)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMouseEventTransition_") {
		n.SetObjectName("QMouseEventTransition_" + qt.Identifier())
	}
	return n
}

func (ptr *QMouseEventTransition) QMouseEventTransition_PTR() *QMouseEventTransition {
	return ptr
}

func NewQMouseEventTransition2(object core.QObject_ITF, ty core.QEvent__Type, button core.Qt__MouseButton, sourceState core.QState_ITF) *QMouseEventTransition {
	defer qt.Recovering("QMouseEventTransition::QMouseEventTransition")

	return NewQMouseEventTransitionFromPointer(C.QMouseEventTransition_NewQMouseEventTransition2(core.PointerFromQObject(object), C.int(ty), C.int(button), core.PointerFromQState(sourceState)))
}

func NewQMouseEventTransition(sourceState core.QState_ITF) *QMouseEventTransition {
	defer qt.Recovering("QMouseEventTransition::QMouseEventTransition")

	return NewQMouseEventTransitionFromPointer(C.QMouseEventTransition_NewQMouseEventTransition(core.PointerFromQState(sourceState)))
}

func (ptr *QMouseEventTransition) Button() core.Qt__MouseButton {
	defer qt.Recovering("QMouseEventTransition::button")

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEventTransition_Button(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEventTransition) EventTest(event core.QEvent_ITF) bool {
	defer qt.Recovering("QMouseEventTransition::eventTest")

	if ptr.Pointer() != nil {
		return C.QMouseEventTransition_EventTest(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMouseEventTransition) ModifierMask() core.Qt__KeyboardModifier {
	defer qt.Recovering("QMouseEventTransition::modifierMask")

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QMouseEventTransition_ModifierMask(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMouseEventTransition) ConnectOnTransition(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMouseEventTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "onTransition", f)
	}
}

func (ptr *QMouseEventTransition) DisconnectOnTransition() {
	defer qt.Recovering("disconnect QMouseEventTransition::onTransition")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "onTransition")
	}
}

//export callbackQMouseEventTransitionOnTransition
func callbackQMouseEventTransitionOnTransition(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMouseEventTransition::onTransition")

	if signal := qt.GetSignal(C.GoString(ptrName), "onTransition"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMouseEventTransitionFromPointer(ptr).OnTransitionDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMouseEventTransition) OnTransition(event core.QEvent_ITF) {
	defer qt.Recovering("QMouseEventTransition::onTransition")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_OnTransition(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMouseEventTransition) OnTransitionDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMouseEventTransition::onTransition")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_OnTransitionDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMouseEventTransition) SetButton(button core.Qt__MouseButton) {
	defer qt.Recovering("QMouseEventTransition::setButton")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetButton(ptr.Pointer(), C.int(button))
	}
}

func (ptr *QMouseEventTransition) SetHitTestPath(path gui.QPainterPath_ITF) {
	defer qt.Recovering("QMouseEventTransition::setHitTestPath")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetHitTestPath(ptr.Pointer(), gui.PointerFromQPainterPath(path))
	}
}

func (ptr *QMouseEventTransition) SetModifierMask(modifierMask core.Qt__KeyboardModifier) {
	defer qt.Recovering("QMouseEventTransition::setModifierMask")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_SetModifierMask(ptr.Pointer(), C.int(modifierMask))
	}
}

func (ptr *QMouseEventTransition) DestroyQMouseEventTransition() {
	defer qt.Recovering("QMouseEventTransition::~QMouseEventTransition")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_DestroyQMouseEventTransition(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMouseEventTransition) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMouseEventTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMouseEventTransition) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMouseEventTransition::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMouseEventTransitionTimerEvent
func callbackQMouseEventTransitionTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMouseEventTransition::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMouseEventTransitionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMouseEventTransition) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMouseEventTransition::timerEvent")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMouseEventTransition) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMouseEventTransition::timerEvent")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMouseEventTransition) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMouseEventTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMouseEventTransition) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMouseEventTransition::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMouseEventTransitionChildEvent
func callbackQMouseEventTransitionChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMouseEventTransition::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMouseEventTransitionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMouseEventTransition) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMouseEventTransition::childEvent")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMouseEventTransition) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMouseEventTransition::childEvent")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMouseEventTransition) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMouseEventTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMouseEventTransition) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMouseEventTransition::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMouseEventTransitionCustomEvent
func callbackQMouseEventTransitionCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMouseEventTransition::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMouseEventTransitionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMouseEventTransition) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMouseEventTransition::customEvent")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMouseEventTransition) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMouseEventTransition::customEvent")

	if ptr.Pointer() != nil {
		C.QMouseEventTransition_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
