package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QVariantAnimation struct {
	QAbstractAnimation
}

type QVariantAnimation_ITF interface {
	QAbstractAnimation_ITF
	QVariantAnimation_PTR() *QVariantAnimation
}

func PointerFromQVariantAnimation(ptr QVariantAnimation_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVariantAnimation_PTR().Pointer()
	}
	return nil
}

func NewQVariantAnimationFromPointer(ptr unsafe.Pointer) *QVariantAnimation {
	var n = new(QVariantAnimation)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVariantAnimation_") {
		n.SetObjectName("QVariantAnimation_" + qt.Identifier())
	}
	return n
}

func (ptr *QVariantAnimation) QVariantAnimation_PTR() *QVariantAnimation {
	return ptr
}

func (ptr *QVariantAnimation) CurrentValue() *QVariant {
	defer qt.Recovering("QVariantAnimation::currentValue")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QVariantAnimation_CurrentValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariantAnimation) Duration() int {
	defer qt.Recovering("QVariantAnimation::duration")

	if ptr.Pointer() != nil {
		return int(C.QVariantAnimation_Duration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVariantAnimation) EasingCurve() *QEasingCurve {
	defer qt.Recovering("QVariantAnimation::easingCurve")

	if ptr.Pointer() != nil {
		return NewQEasingCurveFromPointer(C.QVariantAnimation_EasingCurve(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariantAnimation) EndValue() *QVariant {
	defer qt.Recovering("QVariantAnimation::endValue")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QVariantAnimation_EndValue(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariantAnimation) SetDuration(msecs int) {
	defer qt.Recovering("QVariantAnimation::setDuration")

	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetDuration(ptr.Pointer(), C.int(msecs))
	}
}

func (ptr *QVariantAnimation) SetEasingCurve(easing QEasingCurve_ITF) {
	defer qt.Recovering("QVariantAnimation::setEasingCurve")

	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetEasingCurve(ptr.Pointer(), PointerFromQEasingCurve(easing))
	}
}

func (ptr *QVariantAnimation) SetEndValue(value QVariant_ITF) {
	defer qt.Recovering("QVariantAnimation::setEndValue")

	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetEndValue(ptr.Pointer(), PointerFromQVariant(value))
	}
}

func (ptr *QVariantAnimation) SetStartValue(value QVariant_ITF) {
	defer qt.Recovering("QVariantAnimation::setStartValue")

	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetStartValue(ptr.Pointer(), PointerFromQVariant(value))
	}
}

func (ptr *QVariantAnimation) StartValue() *QVariant {
	defer qt.Recovering("QVariantAnimation::startValue")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QVariantAnimation_StartValue(ptr.Pointer()))
	}
	return nil
}

func NewQVariantAnimation(parent QObject_ITF) *QVariantAnimation {
	defer qt.Recovering("QVariantAnimation::QVariantAnimation")

	return NewQVariantAnimationFromPointer(C.QVariantAnimation_NewQVariantAnimation(PointerFromQObject(parent)))
}

func (ptr *QVariantAnimation) KeyValueAt(step float64) *QVariant {
	defer qt.Recovering("QVariantAnimation::keyValueAt")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QVariantAnimation_KeyValueAt(ptr.Pointer(), C.double(step)))
	}
	return nil
}

func (ptr *QVariantAnimation) SetKeyValueAt(step float64, value QVariant_ITF) {
	defer qt.Recovering("QVariantAnimation::setKeyValueAt")

	if ptr.Pointer() != nil {
		C.QVariantAnimation_SetKeyValueAt(ptr.Pointer(), C.double(step), PointerFromQVariant(value))
	}
}

func (ptr *QVariantAnimation) ConnectUpdateCurrentTime(f func(v int)) {
	defer qt.Recovering("connect QVariantAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateCurrentTime", f)
	}
}

func (ptr *QVariantAnimation) DisconnectUpdateCurrentTime() {
	defer qt.Recovering("disconnect QVariantAnimation::updateCurrentTime")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateCurrentTime")
	}
}

//export callbackQVariantAnimationUpdateCurrentTime
func callbackQVariantAnimationUpdateCurrentTime(ptrName *C.char, v C.int) bool {
	defer qt.Recovering("callback QVariantAnimation::updateCurrentTime")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentTime"); signal != nil {
		signal.(func(int))(int(v))
		return true
	}
	return false

}

func (ptr *QVariantAnimation) ConnectUpdateCurrentValue(f func(value *QVariant)) {
	defer qt.Recovering("connect QVariantAnimation::updateCurrentValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateCurrentValue", f)
	}
}

func (ptr *QVariantAnimation) DisconnectUpdateCurrentValue() {
	defer qt.Recovering("disconnect QVariantAnimation::updateCurrentValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateCurrentValue")
	}
}

//export callbackQVariantAnimationUpdateCurrentValue
func callbackQVariantAnimationUpdateCurrentValue(ptrName *C.char, value unsafe.Pointer) bool {
	defer qt.Recovering("callback QVariantAnimation::updateCurrentValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateCurrentValue"); signal != nil {
		signal.(func(*QVariant))(NewQVariantFromPointer(value))
		return true
	}
	return false

}

func (ptr *QVariantAnimation) ConnectUpdateState(f func(newState QAbstractAnimation__State, oldState QAbstractAnimation__State)) {
	defer qt.Recovering("connect QVariantAnimation::updateState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateState", f)
	}
}

func (ptr *QVariantAnimation) DisconnectUpdateState() {
	defer qt.Recovering("disconnect QVariantAnimation::updateState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateState")
	}
}

//export callbackQVariantAnimationUpdateState
func callbackQVariantAnimationUpdateState(ptrName *C.char, newState C.int, oldState C.int) bool {
	defer qt.Recovering("callback QVariantAnimation::updateState")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateState"); signal != nil {
		signal.(func(QAbstractAnimation__State, QAbstractAnimation__State))(QAbstractAnimation__State(newState), QAbstractAnimation__State(oldState))
		return true
	}
	return false

}

func (ptr *QVariantAnimation) ConnectValueChanged(f func(value *QVariant)) {
	defer qt.Recovering("connect QVariantAnimation::valueChanged")

	if ptr.Pointer() != nil {
		C.QVariantAnimation_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QVariantAnimation) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QVariantAnimation::valueChanged")

	if ptr.Pointer() != nil {
		C.QVariantAnimation_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQVariantAnimationValueChanged
func callbackQVariantAnimationValueChanged(ptrName *C.char, value unsafe.Pointer) {
	defer qt.Recovering("callback QVariantAnimation::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged"); signal != nil {
		signal.(func(*QVariant))(NewQVariantFromPointer(value))
	}

}

func (ptr *QVariantAnimation) DestroyQVariantAnimation() {
	defer qt.Recovering("QVariantAnimation::~QVariantAnimation")

	if ptr.Pointer() != nil {
		C.QVariantAnimation_DestroyQVariantAnimation(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVariantAnimation) ConnectUpdateDirection(f func(direction QAbstractAnimation__Direction)) {
	defer qt.Recovering("connect QVariantAnimation::updateDirection")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateDirection", f)
	}
}

func (ptr *QVariantAnimation) DisconnectUpdateDirection() {
	defer qt.Recovering("disconnect QVariantAnimation::updateDirection")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateDirection")
	}
}

//export callbackQVariantAnimationUpdateDirection
func callbackQVariantAnimationUpdateDirection(ptrName *C.char, direction C.int) bool {
	defer qt.Recovering("callback QVariantAnimation::updateDirection")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateDirection"); signal != nil {
		signal.(func(QAbstractAnimation__Direction))(QAbstractAnimation__Direction(direction))
		return true
	}
	return false

}

func (ptr *QVariantAnimation) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QVariantAnimation::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVariantAnimation) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVariantAnimation::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVariantAnimationTimerEvent
func callbackQVariantAnimationTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVariantAnimation::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QVariantAnimation) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QVariantAnimation::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVariantAnimation) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVariantAnimation::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVariantAnimationChildEvent
func callbackQVariantAnimationChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVariantAnimation::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QVariantAnimation) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QVariantAnimation::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVariantAnimation) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVariantAnimation::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVariantAnimationCustomEvent
func callbackQVariantAnimationCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QVariantAnimation::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
