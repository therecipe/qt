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
func callbackQAbstractSliderActionTriggered(ptr unsafe.Pointer, ptrName *C.char, action C.int) {
	defer qt.Recovering("callback QAbstractSlider::actionTriggered")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionTriggered"); signal != nil {
		signal.(func(int))(int(action))
	}

}

func (ptr *QAbstractSlider) ActionTriggered(action int) {
	defer qt.Recovering("QAbstractSlider::actionTriggered")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ActionTriggered(ptr.Pointer(), C.int(action))
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
func callbackQAbstractSliderChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQAbstractSliderFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QAbstractSlider) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QAbstractSlider) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QAbstractSlider) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractSlider::event")

	if ptr.Pointer() != nil {
		return C.QAbstractSlider_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
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
func callbackQAbstractSliderKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQAbstractSliderFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QAbstractSlider) KeyPressEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QAbstractSlider) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
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
func callbackQAbstractSliderRangeChanged(ptr unsafe.Pointer, ptrName *C.char, min C.int, max C.int) {
	defer qt.Recovering("callback QAbstractSlider::rangeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "rangeChanged"); signal != nil {
		signal.(func(int, int))(int(min), int(max))
	}

}

func (ptr *QAbstractSlider) RangeChanged(min int, max int) {
	defer qt.Recovering("QAbstractSlider::rangeChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_RangeChanged(ptr.Pointer(), C.int(min), C.int(max))
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
func callbackQAbstractSliderSliderChange(ptr unsafe.Pointer, ptrName *C.char, change C.int) {
	defer qt.Recovering("callback QAbstractSlider::sliderChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "sliderChange"); signal != nil {
		signal.(func(QAbstractSlider__SliderChange))(QAbstractSlider__SliderChange(change))
	} else {
		NewQAbstractSliderFromPointer(ptr).SliderChangeDefault(QAbstractSlider__SliderChange(change))
	}
}

func (ptr *QAbstractSlider) SliderChange(change QAbstractSlider__SliderChange) {
	defer qt.Recovering("QAbstractSlider::sliderChange")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SliderChange(ptr.Pointer(), C.int(change))
	}
}

func (ptr *QAbstractSlider) SliderChangeDefault(change QAbstractSlider__SliderChange) {
	defer qt.Recovering("QAbstractSlider::sliderChange")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SliderChangeDefault(ptr.Pointer(), C.int(change))
	}
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
func callbackQAbstractSliderSliderMoved(ptr unsafe.Pointer, ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QAbstractSlider::sliderMoved")

	if signal := qt.GetSignal(C.GoString(ptrName), "sliderMoved"); signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QAbstractSlider) SliderMoved(value int) {
	defer qt.Recovering("QAbstractSlider::sliderMoved")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SliderMoved(ptr.Pointer(), C.int(value))
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
func callbackQAbstractSliderSliderPressed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSlider::sliderPressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "sliderPressed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSlider) SliderPressed() {
	defer qt.Recovering("QAbstractSlider::sliderPressed")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SliderPressed(ptr.Pointer())
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
func callbackQAbstractSliderSliderReleased(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractSlider::sliderReleased")

	if signal := qt.GetSignal(C.GoString(ptrName), "sliderReleased"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractSlider) SliderReleased() {
	defer qt.Recovering("QAbstractSlider::sliderReleased")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SliderReleased(ptr.Pointer())
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
func callbackQAbstractSliderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQAbstractSliderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QAbstractSlider) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QAbstractSlider) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
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
func callbackQAbstractSliderValueChanged(ptr unsafe.Pointer, ptrName *C.char, value C.int) {
	defer qt.Recovering("callback QAbstractSlider::valueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "valueChanged"); signal != nil {
		signal.(func(int))(int(value))
	}

}

func (ptr *QAbstractSlider) ValueChanged(value int) {
	defer qt.Recovering("QAbstractSlider::valueChanged")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ValueChanged(ptr.Pointer(), C.int(value))
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
func callbackQAbstractSliderWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQAbstractSliderFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QAbstractSlider) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QAbstractSlider) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QAbstractSlider) DestroyQAbstractSlider() {
	defer qt.Recovering("QAbstractSlider::~QAbstractSlider")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DestroyQAbstractSlider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractSlider) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QAbstractSlider::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQAbstractSliderActionEvent
func callbackQAbstractSliderActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QAbstractSlider) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QAbstractSlider::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQAbstractSliderDragEnterEvent
func callbackQAbstractSliderDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QAbstractSlider) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QAbstractSlider::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQAbstractSliderDragLeaveEvent
func callbackQAbstractSliderDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QAbstractSlider) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QAbstractSlider::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQAbstractSliderDragMoveEvent
func callbackQAbstractSliderDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QAbstractSlider) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QAbstractSlider::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQAbstractSliderDropEvent
func callbackQAbstractSliderDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QAbstractSlider) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSlider::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQAbstractSliderEnterEvent
func callbackQAbstractSliderEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSlider) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractSlider::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQAbstractSliderFocusInEvent
func callbackQAbstractSliderFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractSlider) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractSlider::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQAbstractSliderFocusOutEvent
func callbackQAbstractSliderFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractSlider) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QAbstractSlider::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQAbstractSliderHideEvent
func callbackQAbstractSliderHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QAbstractSlider) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSlider::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQAbstractSliderLeaveEvent
func callbackQAbstractSliderLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSlider) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QAbstractSlider::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQAbstractSliderMoveEvent
func callbackQAbstractSliderMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QAbstractSlider) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QAbstractSlider::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQAbstractSliderPaintEvent
func callbackQAbstractSliderPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::paintEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QAbstractSlider) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::paintEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QAbstractSlider::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QAbstractSlider) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QAbstractSlider::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQAbstractSliderSetVisible
func callbackQAbstractSliderSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QAbstractSlider::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QAbstractSlider) SetVisible(visible bool) {
	defer qt.Recovering("QAbstractSlider::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QAbstractSlider) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QAbstractSlider::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QAbstractSlider) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QAbstractSlider::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQAbstractSliderShowEvent
func callbackQAbstractSliderShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QAbstractSlider) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QAbstractSlider::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQAbstractSliderCloseEvent
func callbackQAbstractSliderCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QAbstractSlider) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QAbstractSlider::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQAbstractSliderContextMenuEvent
func callbackQAbstractSliderContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QAbstractSlider) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QAbstractSlider::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QAbstractSlider) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QAbstractSlider::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQAbstractSliderInitPainter
func callbackQAbstractSliderInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQAbstractSliderFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QAbstractSlider) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QAbstractSlider::initPainter")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QAbstractSlider) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QAbstractSlider::initPainter")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QAbstractSlider) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QAbstractSlider::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQAbstractSliderInputMethodEvent
func callbackQAbstractSliderInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QAbstractSlider) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractSlider::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQAbstractSliderKeyReleaseEvent
func callbackQAbstractSliderKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QAbstractSlider) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractSlider::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQAbstractSliderMouseDoubleClickEvent
func callbackQAbstractSliderMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSlider) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQAbstractSliderMouseMoveEvent
func callbackQAbstractSliderMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSlider) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractSlider::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQAbstractSliderMousePressEvent
func callbackQAbstractSliderMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSlider) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQAbstractSliderMouseReleaseEvent
func callbackQAbstractSliderMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSlider) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QAbstractSlider::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQAbstractSliderResizeEvent
func callbackQAbstractSliderResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QAbstractSlider) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QAbstractSlider::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQAbstractSliderTabletEvent
func callbackQAbstractSliderTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QAbstractSlider) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractSlider::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractSliderChildEvent
func callbackQAbstractSliderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractSlider) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractSlider) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractSlider::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractSlider) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractSlider::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractSliderCustomEvent
func callbackQAbstractSliderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractSlider::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractSliderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractSlider) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractSlider) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractSlider::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractSlider_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
