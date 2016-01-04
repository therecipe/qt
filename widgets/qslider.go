package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QSlider struct {
	QAbstractSlider
}

type QSlider_ITF interface {
	QAbstractSlider_ITF
	QSlider_PTR() *QSlider
}

func PointerFromQSlider(ptr QSlider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSlider_PTR().Pointer()
	}
	return nil
}

func NewQSliderFromPointer(ptr unsafe.Pointer) *QSlider {
	var n = new(QSlider)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QSlider_") {
		n.SetObjectName("QSlider_" + qt.Identifier())
	}
	return n
}

func (ptr *QSlider) QSlider_PTR() *QSlider {
	return ptr
}

//QSlider::TickPosition
type QSlider__TickPosition int64

const (
	QSlider__NoTicks        = QSlider__TickPosition(0)
	QSlider__TicksAbove     = QSlider__TickPosition(1)
	QSlider__TicksLeft      = QSlider__TickPosition(QSlider__TicksAbove)
	QSlider__TicksBelow     = QSlider__TickPosition(2)
	QSlider__TicksRight     = QSlider__TickPosition(QSlider__TicksBelow)
	QSlider__TicksBothSides = QSlider__TickPosition(3)
)

func (ptr *QSlider) SetTickInterval(ti int) {
	defer qt.Recovering("QSlider::setTickInterval")

	if ptr.Pointer() != nil {
		C.QSlider_SetTickInterval(ptr.Pointer(), C.int(ti))
	}
}

func (ptr *QSlider) SetTickPosition(position QSlider__TickPosition) {
	defer qt.Recovering("QSlider::setTickPosition")

	if ptr.Pointer() != nil {
		C.QSlider_SetTickPosition(ptr.Pointer(), C.int(position))
	}
}

func (ptr *QSlider) TickInterval() int {
	defer qt.Recovering("QSlider::tickInterval")

	if ptr.Pointer() != nil {
		return int(C.QSlider_TickInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSlider) TickPosition() QSlider__TickPosition {
	defer qt.Recovering("QSlider::tickPosition")

	if ptr.Pointer() != nil {
		return QSlider__TickPosition(C.QSlider_TickPosition(ptr.Pointer()))
	}
	return 0
}

func NewQSlider(parent QWidget_ITF) *QSlider {
	defer qt.Recovering("QSlider::QSlider")

	return NewQSliderFromPointer(C.QSlider_NewQSlider(PointerFromQWidget(parent)))
}

func NewQSlider2(orientation core.Qt__Orientation, parent QWidget_ITF) *QSlider {
	defer qt.Recovering("QSlider::QSlider")

	return NewQSliderFromPointer(C.QSlider_NewQSlider2(C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QSlider) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QSlider::event")

	if ptr.Pointer() != nil {
		return C.QSlider_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSlider) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QSlider::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSlider_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSlider) ConnectMouseMoveEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QSlider) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQSliderMouseMoveEvent
func callbackQSliderMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQSliderFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QSlider) MouseMoveEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QSlider) MouseMoveEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSlider::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QSlider) ConnectMousePressEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSlider::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QSlider) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QSlider::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQSliderMousePressEvent
func callbackQSliderMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQSliderFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QSlider) MousePressEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSlider::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QSlider) MousePressEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSlider::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QSlider) ConnectMouseReleaseEvent(f func(ev *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QSlider) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQSliderMouseReleaseEvent
func callbackQSliderMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
	} else {
		NewQSliderFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(ev))
	}
}

func (ptr *QSlider) MouseReleaseEvent(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QSlider) MouseReleaseEventDefault(ev gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSlider::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(ev))
	}
}

func (ptr *QSlider) ConnectPaintEvent(f func(ev *gui.QPaintEvent)) {
	defer qt.Recovering("connect QSlider::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QSlider) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QSlider::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQSliderPaintEvent
func callbackQSliderPaintEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(ev))
	} else {
		NewQSliderFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(ev))
	}
}

func (ptr *QSlider) PaintEvent(ev gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSlider::paintEvent")

	if ptr.Pointer() != nil {
		C.QSlider_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(ev))
	}
}

func (ptr *QSlider) PaintEventDefault(ev gui.QPaintEvent_ITF) {
	defer qt.Recovering("QSlider::paintEvent")

	if ptr.Pointer() != nil {
		C.QSlider_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(ev))
	}
}

func (ptr *QSlider) SizeHint() *core.QSize {
	defer qt.Recovering("QSlider::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSlider_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSlider) DestroyQSlider() {
	defer qt.Recovering("QSlider::~QSlider")

	if ptr.Pointer() != nil {
		C.QSlider_DestroyQSlider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSlider) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QSlider::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QSlider) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QSlider::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQSliderChangeEvent
func callbackQSliderChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQSliderFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QSlider) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QSlider::changeEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QSlider) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QSlider::changeEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QSlider) ConnectKeyPressEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSlider::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QSlider) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QSlider::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQSliderKeyPressEvent
func callbackQSliderKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
	} else {
		NewQSliderFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(ev))
	}
}

func (ptr *QSlider) KeyPressEvent(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSlider::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSlider_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QSlider) KeyPressEventDefault(ev gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSlider::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QSlider_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(ev))
	}
}

func (ptr *QSlider) ConnectSliderChange(f func(change QAbstractSlider__SliderChange)) {
	defer qt.Recovering("connect QSlider::sliderChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sliderChange", f)
	}
}

func (ptr *QSlider) DisconnectSliderChange() {
	defer qt.Recovering("disconnect QSlider::sliderChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sliderChange")
	}
}

//export callbackQSliderSliderChange
func callbackQSliderSliderChange(ptr unsafe.Pointer, ptrName *C.char, change C.int) {
	defer qt.Recovering("callback QSlider::sliderChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "sliderChange"); signal != nil {
		signal.(func(QAbstractSlider__SliderChange))(QAbstractSlider__SliderChange(change))
	} else {
		NewQSliderFromPointer(ptr).SliderChangeDefault(QAbstractSlider__SliderChange(change))
	}
}

func (ptr *QSlider) SliderChange(change QAbstractSlider__SliderChange) {
	defer qt.Recovering("QSlider::sliderChange")

	if ptr.Pointer() != nil {
		C.QSlider_SliderChange(ptr.Pointer(), C.int(change))
	}
}

func (ptr *QSlider) SliderChangeDefault(change QAbstractSlider__SliderChange) {
	defer qt.Recovering("QSlider::sliderChange")

	if ptr.Pointer() != nil {
		C.QSlider_SliderChangeDefault(ptr.Pointer(), C.int(change))
	}
}

func (ptr *QSlider) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QSlider::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSlider) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSlider::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSliderTimerEvent
func callbackQSliderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQSliderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QSlider) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QSlider::timerEvent")

	if ptr.Pointer() != nil {
		C.QSlider_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QSlider) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QSlider::timerEvent")

	if ptr.Pointer() != nil {
		C.QSlider_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QSlider) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QSlider::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QSlider) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QSlider::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQSliderWheelEvent
func callbackQSliderWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQSliderFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QSlider) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSlider::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSlider_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QSlider) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QSlider::wheelEvent")

	if ptr.Pointer() != nil {
		C.QSlider_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QSlider) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QSlider::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QSlider) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QSlider::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQSliderActionEvent
func callbackQSliderActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QSlider) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSlider::actionEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSlider) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QSlider::actionEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QSlider) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QSlider::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QSlider) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QSlider::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQSliderDragEnterEvent
func callbackQSliderDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QSlider) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSlider::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSlider_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSlider) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QSlider::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QSlider_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QSlider) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QSlider::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QSlider) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QSlider::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQSliderDragLeaveEvent
func callbackQSliderDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QSlider) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSlider::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSlider) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QSlider::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QSlider) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QSlider::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QSlider) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QSlider::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQSliderDragMoveEvent
func callbackQSliderDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QSlider) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSlider::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSlider) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QSlider::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QSlider) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QSlider::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QSlider) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QSlider::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQSliderDropEvent
func callbackQSliderDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QSlider) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSlider::dropEvent")

	if ptr.Pointer() != nil {
		C.QSlider_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSlider) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QSlider::dropEvent")

	if ptr.Pointer() != nil {
		C.QSlider_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QSlider) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSlider::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QSlider) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QSlider::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQSliderEnterEvent
func callbackQSliderEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSlider) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSlider::enterEvent")

	if ptr.Pointer() != nil {
		C.QSlider_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSlider) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSlider::enterEvent")

	if ptr.Pointer() != nil {
		C.QSlider_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSlider) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSlider::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QSlider) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QSlider::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQSliderFocusInEvent
func callbackQSliderFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSlider) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSlider::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSlider_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSlider) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSlider::focusInEvent")

	if ptr.Pointer() != nil {
		C.QSlider_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSlider) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QSlider::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QSlider) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QSlider::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQSliderFocusOutEvent
func callbackQSliderFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QSlider) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSlider::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSlider_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSlider) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QSlider::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QSlider_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QSlider) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QSlider::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QSlider) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QSlider::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQSliderHideEvent
func callbackQSliderHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QSlider) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSlider::hideEvent")

	if ptr.Pointer() != nil {
		C.QSlider_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSlider) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QSlider::hideEvent")

	if ptr.Pointer() != nil {
		C.QSlider_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QSlider) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSlider::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QSlider) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QSlider::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQSliderLeaveEvent
func callbackQSliderLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSlider) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSlider::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSlider) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSlider::leaveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSlider) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QSlider::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QSlider) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QSlider::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQSliderMoveEvent
func callbackQSliderMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QSlider) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSlider::moveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSlider) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QSlider::moveEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QSlider) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QSlider::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QSlider) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QSlider::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQSliderSetVisible
func callbackQSliderSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSlider::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QSlider) SetVisible(visible bool) {
	defer qt.Recovering("QSlider::setVisible")

	if ptr.Pointer() != nil {
		C.QSlider_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSlider) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QSlider::setVisible")

	if ptr.Pointer() != nil {
		C.QSlider_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QSlider) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QSlider::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QSlider) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QSlider::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQSliderShowEvent
func callbackQSliderShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QSlider) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSlider::showEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSlider) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QSlider::showEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QSlider) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QSlider::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QSlider) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QSlider::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQSliderCloseEvent
func callbackQSliderCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QSlider) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSlider::closeEvent")

	if ptr.Pointer() != nil {
		C.QSlider_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSlider) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QSlider::closeEvent")

	if ptr.Pointer() != nil {
		C.QSlider_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QSlider) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QSlider::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QSlider) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QSlider::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQSliderContextMenuEvent
func callbackQSliderContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QSlider) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSlider::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSlider) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QSlider::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QSlider) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QSlider::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QSlider) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QSlider::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQSliderInitPainter
func callbackQSliderInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQSliderFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QSlider) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSlider::initPainter")

	if ptr.Pointer() != nil {
		C.QSlider_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSlider) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QSlider::initPainter")

	if ptr.Pointer() != nil {
		C.QSlider_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QSlider) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QSlider::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QSlider) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QSlider::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQSliderInputMethodEvent
func callbackQSliderInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QSlider) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSlider::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSlider_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSlider) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QSlider::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QSlider_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QSlider) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QSlider::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QSlider) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QSlider::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQSliderKeyReleaseEvent
func callbackQSliderKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QSlider) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSlider::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSlider_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSlider) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QSlider::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QSlider_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QSlider) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QSlider::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QSlider) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QSlider::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQSliderMouseDoubleClickEvent
func callbackQSliderMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QSlider) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSlider::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSlider) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QSlider::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QSlider_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QSlider) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QSlider::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QSlider) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QSlider::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQSliderResizeEvent
func callbackQSliderResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QSlider) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSlider::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSlider) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QSlider::resizeEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QSlider) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QSlider::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QSlider) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QSlider::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQSliderTabletEvent
func callbackQSliderTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QSlider) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSlider::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSlider_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSlider) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QSlider::tabletEvent")

	if ptr.Pointer() != nil {
		C.QSlider_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QSlider) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSlider::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSlider) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSlider::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSliderChildEvent
func callbackQSliderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSlider) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSlider::childEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSlider) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSlider::childEvent")

	if ptr.Pointer() != nil {
		C.QSlider_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSlider) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSlider::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSlider) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSlider::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSliderCustomEvent
func callbackQSliderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSlider::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSliderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSlider) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSlider::customEvent")

	if ptr.Pointer() != nil {
		C.QSlider_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSlider) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSlider::customEvent")

	if ptr.Pointer() != nil {
		C.QSlider_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
