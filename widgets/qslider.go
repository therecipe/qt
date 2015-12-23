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
func callbackQSliderMouseMoveEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQSliderMousePressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQSliderMouseReleaseEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQSliderPaintEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQSliderChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQSliderKeyPressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQSliderSliderChange(ptrName *C.char, change C.int) bool {
	defer qt.Recovering("callback QSlider::sliderChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "sliderChange"); signal != nil {
		signal.(func(QAbstractSlider__SliderChange))(QAbstractSlider__SliderChange(change))
		return true
	}
	return false

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
func callbackQSliderTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

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
func callbackQSliderWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

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
func callbackQSliderActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QSlider::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQSliderShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQSliderInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQSliderCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QSlider::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
