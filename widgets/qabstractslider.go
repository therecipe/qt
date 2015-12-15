package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAbstractSlider struct {
	QWidget
}

type QAbstractSlider_ITF interface {
	QWidget_ITF
	QAbstractSlider_PTR() *QAbstractSlider
}

func PointerFromQAbstractSlider(ptr QAbstractSlider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractSlider_PTR().Pointer()
	}
	return nil
}

func NewQAbstractSliderFromPointer(ptr unsafe.Pointer) *QAbstractSlider {
	var n = new(QAbstractSlider)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractSlider_") {
		n.SetObjectName("QAbstractSlider_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractSlider) QAbstractSlider_PTR() *QAbstractSlider {
	return ptr
}

//QAbstractSlider::SliderAction
type QAbstractSlider__SliderAction int64

const (
	QAbstractSlider__SliderNoAction      = QAbstractSlider__SliderAction(0)
	QAbstractSlider__SliderSingleStepAdd = QAbstractSlider__SliderAction(1)
	QAbstractSlider__SliderSingleStepSub = QAbstractSlider__SliderAction(2)
	QAbstractSlider__SliderPageStepAdd   = QAbstractSlider__SliderAction(3)
	QAbstractSlider__SliderPageStepSub   = QAbstractSlider__SliderAction(4)
	QAbstractSlider__SliderToMinimum     = QAbstractSlider__SliderAction(5)
	QAbstractSlider__SliderToMaximum     = QAbstractSlider__SliderAction(6)
	QAbstractSlider__SliderMove          = QAbstractSlider__SliderAction(7)
)

//QAbstractSlider::SliderChange
type QAbstractSlider__SliderChange int64

const (
	QAbstractSlider__SliderRangeChange       = QAbstractSlider__SliderChange(0)
	QAbstractSlider__SliderOrientationChange = QAbstractSlider__SliderChange(1)
	QAbstractSlider__SliderStepsChange       = QAbstractSlider__SliderChange(2)
	QAbstractSlider__SliderValueChange       = QAbstractSlider__SliderChange(3)
)

func (ptr *QAbstractSlider) HasTracking() bool {
	defer qt.Recovering("QAbstractSlider::hasTracking")

	if ptr.Pointer() != nil {
		return C.QAbstractSlider_HasTracking(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) InvertedAppearance() bool {
	defer qt.Recovering("QAbstractSlider::invertedAppearance")

	if ptr.Pointer() != nil {
		return C.QAbstractSlider_InvertedAppearance(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) InvertedControls() bool {
	defer qt.Recovering("QAbstractSlider::invertedControls")

	if ptr.Pointer() != nil {
		return C.QAbstractSlider_InvertedControls(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) IsSliderDown() bool {
	defer qt.Recovering("QAbstractSlider::isSliderDown")

	if ptr.Pointer() != nil {
		return C.QAbstractSlider_IsSliderDown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractSlider) Maximum() int {
	defer qt.Recovering("QAbstractSlider::maximum")

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Maximum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) Minimum() int {
	defer qt.Recovering("QAbstractSlider::minimum")

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Minimum(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QAbstractSlider::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QAbstractSlider_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) PageStep() int {
	defer qt.Recovering("QAbstractSlider::pageStep")

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_PageStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) SetInvertedAppearance(v bool) {
	defer qt.Recovering("QAbstractSlider::setInvertedAppearance")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetInvertedAppearance(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetInvertedControls(v bool) {
	defer qt.Recovering("QAbstractSlider::setInvertedControls")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetInvertedControls(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetMaximum(v int) {
	defer qt.Recovering("QAbstractSlider::setMaximum")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMaximum(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetMinimum(v int) {
	defer qt.Recovering("QAbstractSlider::setMinimum")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetMinimum(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetOrientation(v core.Qt__Orientation) {
	defer qt.Recovering("QAbstractSlider::setOrientation")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetOrientation(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetPageStep(v int) {
	defer qt.Recovering("QAbstractSlider::setPageStep")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetPageStep(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetSingleStep(v int) {
	defer qt.Recovering("QAbstractSlider::setSingleStep")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSingleStep(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetSliderDown(v bool) {
	defer qt.Recovering("QAbstractSlider::setSliderDown")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSliderDown(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractSlider) SetSliderPosition(v int) {
	defer qt.Recovering("QAbstractSlider::setSliderPosition")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetSliderPosition(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SetTracking(enable bool) {
	defer qt.Recovering("QAbstractSlider::setTracking")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetTracking(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QAbstractSlider) SetValue(v int) {
	defer qt.Recovering("QAbstractSlider::setValue")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetValue(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractSlider) SingleStep() int {
	defer qt.Recovering("QAbstractSlider::singleStep")

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_SingleStep(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) SliderPosition() int {
	defer qt.Recovering("QAbstractSlider::sliderPosition")

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_SliderPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractSlider) Value() int {
	defer qt.Recovering("QAbstractSlider::value")

	if ptr.Pointer() != nil {
		return int(C.QAbstractSlider_Value(ptr.Pointer()))
	}
	return 0
}

func NewQAbstractSlider(parent QWidget_ITF) *QAbstractSlider {
	defer qt.Recovering("QAbstractSlider::QAbstractSlider")

	return NewQAbstractSliderFromPointer(C.QAbstractSlider_NewQAbstractSlider(PointerFromQWidget(parent)))
}

func (ptr *QAbstractSlider) ConnectActionTriggered(f func(action int)) {
	defer qt.Recovering("connect QAbstractSlider::actionTriggered")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectActionTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "actionTriggered", f)
	}
}

func (ptr *QAbstractSlider) DisconnectActionTriggered() {
	defer qt.Recovering("disconnect QAbstractSlider::actionTriggered")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectActionTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "actionTriggered")
	}
}

//export callbackQAbstractSliderActionTriggered
func callbackQAbstractSliderActionTriggered(ptrName *C.char, action C.int) {
	defer qt.Recovering("callback QAbstractSlider::actionTriggered")

	var signal = qt.GetSignal(C.GoString(ptrName), "actionTriggered")
	if signal != nil {
		signal.(func(int))(int(action))
	}

}

func (ptr *QAbstractSlider) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSlider::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQAbstractSliderChangeEvent
func callbackQAbstractSliderChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSlider::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QAbstractSlider) ConnectKeyPressEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractSlider::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQAbstractSliderKeyPressEvent
func callbackQAbstractSliderKeyPressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSlider::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QAbstractSlider) ConnectRangeChanged(f func(min int, max int)) {
	defer qt.Recovering("connect QAbstractSlider::rangeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectRangeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "rangeChanged", f)
	}
}

func (ptr *QAbstractSlider) DisconnectRangeChanged() {
	defer qt.Recovering("disconnect QAbstractSlider::rangeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectRangeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "rangeChanged")
	}
}

//export callbackQAbstractSliderRangeChanged
func callbackQAbstractSliderRangeChanged(ptrName *C.char, min C.int, max C.int) {
	defer qt.Recovering("callback QAbstractSlider::rangeChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "rangeChanged")
	if signal != nil {
		signal.(func(int, int))(int(min), int(max))
	}

}

func (ptr *QAbstractSlider) SetRange(min int, max int) {
	defer qt.Recovering("QAbstractSlider::setRange")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetRange(ptr.Pointer(), C.int(min), C.int(max))
	}
}

func (ptr *QAbstractSlider) ConnectSliderChange(f func(change QAbstractSlider__SliderChange)) {
	defer qt.Recovering("connect QAbstractSlider::sliderChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sliderChange", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderChange() {
	defer qt.Recovering("disconnect QAbstractSlider::sliderChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sliderChange")
	}
}

//export callbackQAbstractSliderSliderChange
func callbackQAbstractSliderSliderChange(ptrName *C.char, change C.int) bool {
	defer qt.Recovering("callback QAbstractSlider::sliderChange")

	var signal = qt.GetSignal(C.GoString(ptrName), "sliderChange")
	if signal != nil {
		defer signal.(func(QAbstractSlider__SliderChange))(QAbstractSlider__SliderChange(change))
		return true
	}
	return false

}

func (ptr *QAbstractSlider) ConnectSliderMoved(f func(value int)) {
	defer qt.Recovering("connect QAbstractSlider::sliderMoved")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderMoved(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sliderMoved", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderMoved() {
	defer qt.Recovering("disconnect QAbstractSlider::sliderMoved")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderMoved(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sliderMoved")
	}
}

//export callbackQAbstractSliderSliderMoved
func callbackQAbstractSliderSliderMoved(ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QAbstractSlider::sliderMoved")

	var signal = qt.GetSignal(C.GoString(ptrName), "sliderMoved")
	if signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QAbstractSlider) ConnectSliderPressed(f func()) {
	defer qt.Recovering("connect QAbstractSlider::sliderPressed")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sliderPressed", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderPressed() {
	defer qt.Recovering("disconnect QAbstractSlider::sliderPressed")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sliderPressed")
	}
}

//export callbackQAbstractSliderSliderPressed
func callbackQAbstractSliderSliderPressed(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSlider::sliderPressed")

	var signal = qt.GetSignal(C.GoString(ptrName), "sliderPressed")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSlider) ConnectSliderReleased(f func()) {
	defer qt.Recovering("connect QAbstractSlider::sliderReleased")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectSliderReleased(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sliderReleased", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSliderReleased() {
	defer qt.Recovering("disconnect QAbstractSlider::sliderReleased")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectSliderReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sliderReleased")
	}
}

//export callbackQAbstractSliderSliderReleased
func callbackQAbstractSliderSliderReleased(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSlider::sliderReleased")

	var signal = qt.GetSignal(C.GoString(ptrName), "sliderReleased")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSlider) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractSlider::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractSliderTimerEvent
func callbackQAbstractSliderTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSlider::timerEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractSlider) TriggerAction(action QAbstractSlider__SliderAction) {
	defer qt.Recovering("QAbstractSlider::triggerAction")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_TriggerAction(ptr.Pointer(), C.int(action))
	}
}

func (ptr *QAbstractSlider) ConnectValueChanged(f func(value int)) {
	defer qt.Recovering("connect QAbstractSlider::valueChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ConnectValueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "valueChanged", f)
	}
}

func (ptr *QAbstractSlider) DisconnectValueChanged() {
	defer qt.Recovering("disconnect QAbstractSlider::valueChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DisconnectValueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "valueChanged")
	}
}

//export callbackQAbstractSliderValueChanged
func callbackQAbstractSliderValueChanged(ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QAbstractSlider::valueChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "valueChanged")
	if signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QAbstractSlider) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QAbstractSlider::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQAbstractSliderWheelEvent
func callbackQAbstractSliderWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractSlider::wheelEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "wheelEvent")
	if signal != nil {
		defer signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractSlider) DestroyQAbstractSlider() {
	defer qt.Recovering("QAbstractSlider::~QAbstractSlider")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DestroyQAbstractSlider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
