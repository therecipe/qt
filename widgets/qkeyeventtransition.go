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
func callbackQKeyEventTransitionOnTransition(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QKeyEventTransition::onTransition")

	if signal := qt.GetSignal(C.GoString(ptrName), "onTransition"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQKeyEventTransitionTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QKeyEventTransition::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQKeyEventTransitionChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QKeyEventTransition::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQKeyEventTransitionCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QKeyEventTransition::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
