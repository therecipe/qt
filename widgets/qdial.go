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
func callbackQDialMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQDialMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQDialMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQDialPaintEvent(ptrName *C.char, pe unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(pe))
		return true
	}
	return false

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
func callbackQDialResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

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
func callbackQDialSliderChange(ptrName *C.char, change C.int) bool {
	defer qt.Recovering("callback QDial::sliderChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "sliderChange"); signal != nil {
		signal.(func(QAbstractSlider__SliderChange))(QAbstractSlider__SliderChange(change))
		return true
	}
	return false

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
func callbackQDialChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQDialKeyPressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
		return true
	}
	return false

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
func callbackQDialTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

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
func callbackQDialWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

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
func callbackQDialActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QDial::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQDialShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQDialInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQDialCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QDial::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
