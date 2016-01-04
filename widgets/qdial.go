package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QDial struct {
	QAbstractSlider
}

type QDial_ITF interface {
	QAbstractSlider_ITF
	QDial_PTR() *QDial
}

func PointerFromQDial(ptr QDial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDial_PTR().Pointer()
	}
	return nil
}

func NewQDialFromPointer(ptr unsafe.Pointer) *QDial {
	var n = new(QDial)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDial_") {
		n.SetObjectName("QDial_" + qt.Identifier())
	}
	return n
}

func (ptr *QDial) QDial_PTR() *QDial {
	return ptr
}

func (ptr *QDial) NotchSize() int {
	defer qt.Recovering("QDial::notchSize")

	if ptr.Pointer() != nil {
		return int(C.QDial_NotchSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDial) NotchTarget() float64 {
	defer qt.Recovering("QDial::notchTarget")

	if ptr.Pointer() != nil {
		return float64(C.QDial_NotchTarget(ptr.Pointer()))
	}
	return 0
}

func (ptr *QDial) NotchesVisible() bool {
	defer qt.Recovering("QDial::notchesVisible")

	if ptr.Pointer() != nil {
		return C.QDial_NotchesVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDial) SetNotchesVisible(visible bool) {
	defer qt.Recovering("QDial::setNotchesVisible")

	if ptr.Pointer() != nil {
		C.QDial_SetNotchesVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDial) SetWrapping(on bool) {
	defer qt.Recovering("QDial::setWrapping")

	if ptr.Pointer() != nil {
		C.QDial_SetWrapping(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QDial) Wrapping() bool {
	defer qt.Recovering("QDial::wrapping")

	if ptr.Pointer() != nil {
		return C.QDial_Wrapping(ptr.Pointer()) != 0
	}
	return false
}

func NewQDial(parent QWidget_ITF) *QDial {
	defer qt.Recovering("QDial::QDial")

	return NewQDialFromPointer(C.QDial_NewQDial(PointerFromQWidget(parent)))
}

func (ptr *QDial) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QDial::event")

	if ptr.Pointer() != nil {
		return C.QDial_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QDial) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QDial::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QDial_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDial) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDial::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QDial) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QDial::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQDialMouseMoveEvent
func callbackQDialMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QDial::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQDialFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QDial) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDial::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDial_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QDial) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDial::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QDial_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QDial) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDial::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QDial) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QDial::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQDialMousePressEvent
func callbackQDialMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QDial::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQDialFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QDial) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDial::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDial_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QDial) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDial::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QDial_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QDial) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDial::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QDial) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QDial::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQDialMouseReleaseEvent
func callbackQDialMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QDial::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQDialFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QDial) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDial::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDial_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QDial) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDial::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDial_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QDial) ConnectPaintEvent(f func(pe *gui.QPaintEvent)) {
	defer qt.Recovering("connect QDial::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QDial) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QDial::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQDialPaintEvent
func callbackQDialPaintEvent(ptr unsafe.Pointer, ptrName *C.char, pe unsafe.Pointer) {
	defer qt.Recovering("callback QDial::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(pe))
	} else {
		NewQDialFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(pe))
	}
}

func (ptr *QDial) PaintEvent(pe gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDial::paintEvent")

	if ptr.Pointer() != nil {
		C.QDial_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(pe))
	}
}

func (ptr *QDial) PaintEventDefault(pe gui.QPaintEvent_ITF) {
	defer qt.Recovering("QDial::paintEvent")

	if ptr.Pointer() != nil {
		C.QDial_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(pe))
	}
}

func (ptr *QDial) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QDial::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QDial) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QDial::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQDialResizeEvent
func callbackQDialResizeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QDial::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
	} else {
		NewQDialFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(e))
	}
}

func (ptr *QDial) ResizeEvent(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDial::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDial_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QDial) ResizeEventDefault(e gui.QResizeEvent_ITF) {
	defer qt.Recovering("QDial::resizeEvent")

	if ptr.Pointer() != nil {
		C.QDial_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(e))
	}
}

func (ptr *QDial) SizeHint() *core.QSize {
	defer qt.Recovering("QDial::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QDial_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QDial) ConnectSliderChange(f func(change QAbstractSlider__SliderChange)) {
	defer qt.Recovering("connect QDial::sliderChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sliderChange", f)
	}
}

func (ptr *QDial) DisconnectSliderChange() {
	defer qt.Recovering("disconnect QDial::sliderChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sliderChange")
	}
}

//export callbackQDialSliderChange
func callbackQDialSliderChange(ptr unsafe.Pointer, ptrName *C.char, change C.int) {
	defer qt.Recovering("callback QDial::sliderChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "sliderChange"); signal != nil {
		signal.(func(QAbstractSlider__SliderChange))(QAbstractSlider__SliderChange(change))
	} else {
		NewQDialFromPointer(ptr).SliderChangeDefault(QAbstractSlider__SliderChange(change))
	}
}

func (ptr *QDial) SliderChange(change QAbstractSlider__SliderChange) {
	defer qt.Recovering("QDial::sliderChange")

	if ptr.Pointer() != nil {
		C.QDial_SliderChange(ptr.Pointer(), C.int(change))
	}
}

func (ptr *QDial) SliderChangeDefault(change QAbstractSlider__SliderChange) {
	defer qt.Recovering("QDial::sliderChange")

	if ptr.Pointer() != nil {
		C.QDial_SliderChangeDefault(ptr.Pointer(), C.int(change))
	}
}

func (ptr *QDial) DestroyQDial() {
	defer qt.Recovering("QDial::~QDial")

	if ptr.Pointer() != nil {
		C.QDial_DestroyQDial(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDial) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QDial::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QDial) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QDial::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQDialChangeEvent
func callbackQDialChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QDial::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQDialFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QDial) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QDial::changeEvent")

	if ptr.Pointer() != nil {
		C.QDial_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QDial) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QDial::changeEvent")

	if ptr.Pointer() != nil {
		C.QDial_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QDial) ConnectKeyPressEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDial::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QDial) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QDial::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQDialKeyPressEvent
func callbackQDialKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QDial::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQDialFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QDial) KeyPressEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDial::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDial_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QDial) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDial::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QDial_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QDial) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QDial::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDial) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDial::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDialTimerEvent
func callbackQDialTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QDial::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQDialFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QDial) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QDial::timerEvent")

	if ptr.Pointer() != nil {
		C.QDial_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QDial) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QDial::timerEvent")

	if ptr.Pointer() != nil {
		C.QDial_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QDial) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QDial::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QDial) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QDial::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQDialWheelEvent
func callbackQDialWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QDial::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQDialFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QDial) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDial::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDial_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QDial) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QDial::wheelEvent")

	if ptr.Pointer() != nil {
		C.QDial_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QDial) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QDial::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QDial) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QDial::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQDialActionEvent
func callbackQDialActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QDial) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDial::actionEvent")

	if ptr.Pointer() != nil {
		C.QDial_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDial) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QDial::actionEvent")

	if ptr.Pointer() != nil {
		C.QDial_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QDial) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QDial::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QDial) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QDial::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQDialDragEnterEvent
func callbackQDialDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QDial) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDial::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDial_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDial) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QDial::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QDial_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QDial) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QDial::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QDial) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QDial::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQDialDragLeaveEvent
func callbackQDialDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QDial) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDial::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDial_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDial) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QDial::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QDial_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QDial) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QDial::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QDial) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QDial::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQDialDragMoveEvent
func callbackQDialDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QDial) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDial::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDial_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDial) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QDial::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QDial_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QDial) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QDial::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QDial) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QDial::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQDialDropEvent
func callbackQDialDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QDial) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDial::dropEvent")

	if ptr.Pointer() != nil {
		C.QDial_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDial) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QDial::dropEvent")

	if ptr.Pointer() != nil {
		C.QDial_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QDial) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDial::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QDial) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QDial::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQDialEnterEvent
func callbackQDialEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDial) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDial::enterEvent")

	if ptr.Pointer() != nil {
		C.QDial_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDial) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDial::enterEvent")

	if ptr.Pointer() != nil {
		C.QDial_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDial) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDial::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QDial) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QDial::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQDialFocusInEvent
func callbackQDialFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDial) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDial::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDial_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDial) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDial::focusInEvent")

	if ptr.Pointer() != nil {
		C.QDial_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDial) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QDial::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QDial) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QDial::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQDialFocusOutEvent
func callbackQDialFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QDial) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDial::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDial_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDial) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QDial::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QDial_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QDial) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QDial::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QDial) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QDial::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQDialHideEvent
func callbackQDialHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QDial) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDial::hideEvent")

	if ptr.Pointer() != nil {
		C.QDial_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDial) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QDial::hideEvent")

	if ptr.Pointer() != nil {
		C.QDial_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QDial) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDial::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QDial) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QDial::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQDialLeaveEvent
func callbackQDialLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDial) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDial::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDial_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDial) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDial::leaveEvent")

	if ptr.Pointer() != nil {
		C.QDial_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDial) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QDial::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QDial) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QDial::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQDialMoveEvent
func callbackQDialMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QDial) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDial::moveEvent")

	if ptr.Pointer() != nil {
		C.QDial_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDial) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QDial::moveEvent")

	if ptr.Pointer() != nil {
		C.QDial_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QDial) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QDial::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QDial) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QDial::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQDialSetVisible
func callbackQDialSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDial::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QDial) SetVisible(visible bool) {
	defer qt.Recovering("QDial::setVisible")

	if ptr.Pointer() != nil {
		C.QDial_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDial) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QDial::setVisible")

	if ptr.Pointer() != nil {
		C.QDial_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QDial) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QDial::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QDial) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QDial::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQDialShowEvent
func callbackQDialShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QDial) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDial::showEvent")

	if ptr.Pointer() != nil {
		C.QDial_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDial) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QDial::showEvent")

	if ptr.Pointer() != nil {
		C.QDial_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QDial) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QDial::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QDial) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QDial::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQDialCloseEvent
func callbackQDialCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QDial) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDial::closeEvent")

	if ptr.Pointer() != nil {
		C.QDial_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDial) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QDial::closeEvent")

	if ptr.Pointer() != nil {
		C.QDial_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QDial) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QDial::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QDial) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QDial::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQDialContextMenuEvent
func callbackQDialContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QDial) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDial::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDial_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDial) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QDial::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QDial_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QDial) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QDial::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QDial) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QDial::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQDialInitPainter
func callbackQDialInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QDial::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQDialFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QDial) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDial::initPainter")

	if ptr.Pointer() != nil {
		C.QDial_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDial) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QDial::initPainter")

	if ptr.Pointer() != nil {
		C.QDial_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QDial) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QDial::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QDial) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QDial::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQDialInputMethodEvent
func callbackQDialInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QDial) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDial::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDial_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDial) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QDial::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QDial_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QDial) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QDial::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QDial) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QDial::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQDialKeyReleaseEvent
func callbackQDialKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QDial) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDial::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDial_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDial) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QDial::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QDial_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QDial) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QDial::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QDial) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QDial::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQDialMouseDoubleClickEvent
func callbackQDialMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QDial) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDial::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDial_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDial) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QDial::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QDial_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QDial) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QDial::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QDial) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QDial::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQDialTabletEvent
func callbackQDialTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QDial) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDial::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDial_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDial) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QDial::tabletEvent")

	if ptr.Pointer() != nil {
		C.QDial_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QDial) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDial::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDial) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDial::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDialChildEvent
func callbackQDialChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDial) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDial::childEvent")

	if ptr.Pointer() != nil {
		C.QDial_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDial) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDial::childEvent")

	if ptr.Pointer() != nil {
		C.QDial_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDial) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDial::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDial) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDial::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDialCustomEvent
func callbackQDialCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDial::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDialFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDial) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDial::customEvent")

	if ptr.Pointer() != nil {
		C.QDial_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDial) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDial::customEvent")

	if ptr.Pointer() != nil {
		C.QDial_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
