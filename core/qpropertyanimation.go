package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QPropertyAnimation struct {
	QVariantAnimation
}

type QPropertyAnimation_ITF interface {
	QVariantAnimation_ITF
	QPropertyAnimation_PTR() *QPropertyAnimation
}

func PointerFromQPropertyAnimation(ptr QPropertyAnimation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPropertyAnimation_PTR().Pointer()
	}
	return nil
}

func NewQPropertyAnimationFromPointer(ptr unsafe.Pointer) *QPropertyAnimation {
	var n = new(QPropertyAnimation)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPropertyAnimation_") {
		n.SetObjectName("QPropertyAnimation_" + qt.Identifier())
	}
	return n
}

func (ptr *QPropertyAnimation) QPropertyAnimation_PTR() *QPropertyAnimation {
	return ptr
}

func (ptr *QPropertyAnimation) PropertyName() *QByteArray {
	defer qt.Recovering("QPropertyAnimation::propertyName")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QPropertyAnimation_PropertyName(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPropertyAnimation) SetPropertyName(propertyName QByteArray_ITF) {
	defer qt.Recovering("QPropertyAnimation::setPropertyName")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_SetPropertyName(ptr.Pointer(), PointerFromQByteArray(propertyName))
	}
}

func (ptr *QPropertyAnimation) SetTargetObject(target QObject_ITF) {
	defer qt.Recovering("QPropertyAnimation::setTargetObject")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_SetTargetObject(ptr.Pointer(), PointerFromQObject(target))
	}
}

func (ptr *QPropertyAnimation) TargetObject() *QObject {
	defer qt.Recovering("QPropertyAnimation::targetObject")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QPropertyAnimation_TargetObject(ptr.Pointer()))
	}
	return nil
}

func NewQPropertyAnimation(parent QObject_ITF) *QPropertyAnimation {
	defer qt.Recovering("QPropertyAnimation::QPropertyAnimation")

	return NewQPropertyAnimationFromPointer(C.QPropertyAnimation_NewQPropertyAnimation(PointerFromQObject(parent)))
}

func NewQPropertyAnimation2(target QObject_ITF, propertyName QByteArray_ITF, parent QObject_ITF) *QPropertyAnimation {
	defer qt.Recovering("QPropertyAnimation::QPropertyAnimation")

	return NewQPropertyAnimationFromPointer(C.QPropertyAnimation_NewQPropertyAnimation2(PointerFromQObject(target), PointerFromQByteArray(propertyName), PointerFromQObject(parent)))
}

func (ptr *QPropertyAnimation) Event(event QEvent_ITF) bool {
	defer qt.Recovering("QPropertyAnimation::event")

	if ptr.Pointer() != nil {
		return C.QPropertyAnimation_Event(ptr.Pointer(), PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QPropertyAnimation) ConnectUpdateCurrentValue(f func(value *QVariant)) {
	defer qt.Recovering("connect QPropertyAnimation::updateCurrentValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateCurrentValue", f)
	}
}

func (ptr *QPropertyAnimation) DisconnectUpdateCurrentValue() {
	defer qt.Recovering("disconnect QPropertyAnimation::updateCurrentValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateCurrentValue")
	}
}

//export callbackQPropertyAnimationUpdateCurrentValue
func callbackQPropertyAnimationUpdateCurrentValue(ptr unsafe.Pointer, ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QPropertyAnimation::updateCurrentValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentValue"); signal != nil {
		signal.(func(*QVariant))(NewQVariantFromPointer(value))
	} else {
		NewQPropertyAnimationFromPointer(ptr).UpdateCurrentValueDefault(NewQVariantFromPointer(value))
	}
}

func (ptr *QPropertyAnimation) UpdateCurrentValue(value QVariant_ITF) {
	defer qt.Recovering("QPropertyAnimation::updateCurrentValue")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_UpdateCurrentValue(ptr.Pointer(), PointerFromQVariant(value))
	}
}

func (ptr *QPropertyAnimation) UpdateCurrentValueDefault(value QVariant_ITF) {
	defer qt.Recovering("QPropertyAnimation::updateCurrentValue")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_UpdateCurrentValueDefault(ptr.Pointer(), PointerFromQVariant(value))
	}
}

func (ptr *QPropertyAnimation) ConnectUpdateState(f func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	defer qt.Recovering("connect QPropertyAnimation::updateState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateState", f)
	}
}

func (ptr *QPropertyAnimation) DisconnectUpdateState() {
	defer qt.Recovering("disconnect QPropertyAnimation::updateState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateState")
	}
}

//export callbackQPropertyAnimationUpdateState
func callbackQPropertyAnimationUpdateState(ptr unsafe.Pointer, ptrName *C.char, newState C.int, oldState C.int) {
	defer qt.Recovering("callback QPropertyAnimation::updateState")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateState"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	} else {
		NewQPropertyAnimationFromPointer(ptr).UpdateStateDefault(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
	}
}

func (ptr *QPropertyAnimation) UpdateState(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QPropertyAnimation::updateState")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_UpdateState(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QPropertyAnimation) UpdateStateDefault(newState QAbstractAnimation__State, oldState QAbstractAnimation__State) {
	defer qt.Recovering("QPropertyAnimation::updateState")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_UpdateStateDefault(ptr.Pointer(), C.int(newState), C.int(oldState))
	}
}

func (ptr *QPropertyAnimation) DestroyQPropertyAnimation() {
	defer qt.Recovering("QPropertyAnimation::~QPropertyAnimation")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_DestroyQPropertyAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPropertyAnimation) ConnectUpdateCurrentTime(f func(v int)) {
	defer qt.Recovering("connect QPropertyAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateCurrentTime", f)
	}
}

func (ptr *QPropertyAnimation) DisconnectUpdateCurrentTime() {
	defer qt.Recovering("disconnect QPropertyAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateCurrentTime")
	}
}

//export callbackQPropertyAnimationUpdateCurrentTime
func callbackQPropertyAnimationUpdateCurrentTime(ptr unsafe.Pointer, ptrName *C.char, v C.int) {
	defer qt.Recovering("callback QPropertyAnimation::updateCurrentTime")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentTime"); signal != nil {
		signal.(func(int))(int(v))
	}

}

func (ptr *QPropertyAnimation) UpdateCurrentTime(v int) {
	defer qt.Recovering("QPropertyAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_UpdateCurrentTime(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QPropertyAnimation) UpdateCurrentTimeDefault(v int) {
	defer qt.Recovering("QPropertyAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_UpdateCurrentTimeDefault(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QPropertyAnimation) ConnectUpdateDirection(f func(direction QAbstractAnimation__Direction)) {
	defer qt.Recovering("connect QPropertyAnimation::updateDirection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateDirection", f)
	}
}

func (ptr *QPropertyAnimation) DisconnectUpdateDirection() {
	defer qt.Recovering("disconnect QPropertyAnimation::updateDirection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateDirection")
	}
}

//export callbackQPropertyAnimationUpdateDirection
func callbackQPropertyAnimationUpdateDirection(ptr unsafe.Pointer, ptrName *C.char, direction C.int) {
	defer qt.Recovering("callback QPropertyAnimation::updateDirection")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateDirection"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
	} else {
		NewQPropertyAnimationFromPointer(ptr).UpdateDirectionDefault(QAbstractAnimation__Direction(direction))
	}
}

func (ptr *QPropertyAnimation) UpdateDirection(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QPropertyAnimation::updateDirection")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_UpdateDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QPropertyAnimation) UpdateDirectionDefault(direction QAbstractAnimation__Direction) {
	defer qt.Recovering("QPropertyAnimation::updateDirection")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_UpdateDirectionDefault(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QPropertyAnimation) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QPropertyAnimation::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPropertyAnimation) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPropertyAnimation::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPropertyAnimationTimerEvent
func callbackQPropertyAnimationTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPropertyAnimation::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQPropertyAnimationFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPropertyAnimation) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QPropertyAnimation::timerEvent")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QPropertyAnimation) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QPropertyAnimation::timerEvent")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QPropertyAnimation) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QPropertyAnimation::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPropertyAnimation) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPropertyAnimation::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPropertyAnimationChildEvent
func callbackQPropertyAnimationChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPropertyAnimation::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQPropertyAnimationFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QPropertyAnimation) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QPropertyAnimation::childEvent")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QPropertyAnimation) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QPropertyAnimation::childEvent")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QPropertyAnimation) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QPropertyAnimation::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPropertyAnimation) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPropertyAnimation::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPropertyAnimationCustomEvent
func callbackQPropertyAnimationCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPropertyAnimation::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQPropertyAnimationFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QPropertyAnimation) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QPropertyAnimation::customEvent")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QPropertyAnimation) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QPropertyAnimation::customEvent")

	if ptr.Pointer() != nil {
		C.QPropertyAnimation_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
