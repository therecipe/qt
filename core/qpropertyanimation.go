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
func callbackQPropertyAnimationUpdateCurrentValue(ptrName *C.char, value unsafe.Pointer) bool {
	defer qt.Recovering("callback QPropertyAnimation::updateCurrentValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentValue"); signal != nil {
		signal.(func(*QVariant))(NewQVariantFromPointer(value))
		return true
	}
	return false

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
func callbackQPropertyAnimationUpdateState(ptrName *C.char, newState C.int, oldState C.int) bool {
	defer qt.Recovering("callback QPropertyAnimation::updateState")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateState"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
		return true
	}
	return false

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
func callbackQPropertyAnimationUpdateCurrentTime(ptrName *C.char, v C.int) bool {
	defer qt.Recovering("callback QPropertyAnimation::updateCurrentTime")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentTime"); signal != nil {
		signal.(func(int))(int(v))
		return true
	}
	return false

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
func callbackQPropertyAnimationUpdateDirection(ptrName *C.char, direction C.int) bool {
	defer qt.Recovering("callback QPropertyAnimation::updateDirection")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateDirection"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
		return true
	}
	return false

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
func callbackQPropertyAnimationTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPropertyAnimation::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQPropertyAnimationChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPropertyAnimation::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQPropertyAnimationCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QPropertyAnimation::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
