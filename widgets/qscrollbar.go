package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QScrollBar struct {
	QAbstractSlider
}

type QScrollBar_ITF interface {
	QAbstractSlider_ITF
	QScrollBar_PTR() *QScrollBar
}

func PointerFromQScrollBar(ptr QScrollBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScrollBar_PTR().Pointer()
	}
	return nil
}

func NewQScrollBarFromPointer(ptr unsafe.Pointer) *QScrollBar {
	var n = new(QScrollBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QScrollBar_") {
		n.SetObjectName("QScrollBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QScrollBar) QScrollBar_PTR() *QScrollBar {
	return ptr
}

func NewQScrollBar(parent QWidget_ITF) *QScrollBar {
	defer qt.Recovering("QScrollBar::QScrollBar")

	return NewQScrollBarFromPointer(C.QScrollBar_NewQScrollBar(PointerFromQWidget(parent)))
}

func NewQScrollBar2(orientation core.Qt__Orientation, parent QWidget_ITF) *QScrollBar {
	defer qt.Recovering("QScrollBar::QScrollBar")

	return NewQScrollBarFromPointer(C.QScrollBar_NewQScrollBar2(C.int(orientation), PointerFromQWidget(parent)))
}

func (ptr *QScrollBar) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QScrollBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QScrollBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQScrollBarContextMenuEvent
func callbackQScrollBarContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QScrollBar::event")

	if ptr.Pointer() != nil {
		return C.QScrollBar_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QScrollBar) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QScrollBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QScrollBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQScrollBarHideEvent
func callbackQScrollBarHideEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QScrollBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQScrollBarMouseMoveEvent
func callbackQScrollBarMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QScrollBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQScrollBarMousePressEvent
func callbackQScrollBarMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QScrollBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQScrollBarMouseReleaseEvent
func callbackQScrollBarMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QScrollBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QScrollBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQScrollBarPaintEvent
func callbackQScrollBarPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QScrollBar) SizeHint() *core.QSize {
	defer qt.Recovering("QScrollBar::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QScrollBar_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QScrollBar) ConnectSliderChange(f func(change QAbstractSlider__SliderChange)) {
	defer qt.Recovering("connect QScrollBar::sliderChange")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sliderChange", f)
	}
}

func (ptr *QScrollBar) DisconnectSliderChange() {
	defer qt.Recovering("disconnect QScrollBar::sliderChange")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sliderChange")
	}
}

//export callbackQScrollBarSliderChange
func callbackQScrollBarSliderChange(ptrName *C.char, change C.int) bool {
	defer qt.Recovering("callback QScrollBar::sliderChange")

	if signal := qt.GetSignal(C.GoString(ptrName), "sliderChange"); signal != nil {
		signal.(func(QAbstractSlider__SliderChange))(QAbstractSlider__SliderChange(change))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QScrollBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QScrollBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQScrollBarWheelEvent
func callbackQScrollBarWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) DestroyQScrollBar() {
	defer qt.Recovering("QScrollBar::~QScrollBar")

	if ptr.Pointer() != nil {
		C.QScrollBar_DestroyQScrollBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QScrollBar) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QScrollBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QScrollBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQScrollBarChangeEvent
func callbackQScrollBarChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectKeyPressEvent(f func(ev *gui.QKeyEvent)) {
	defer qt.Recovering("connect QScrollBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QScrollBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQScrollBarKeyPressEvent
func callbackQScrollBarKeyPressEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QScrollBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QScrollBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQScrollBarTimerEvent
func callbackQScrollBarTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QScrollBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QScrollBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQScrollBarActionEvent
func callbackQScrollBarActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QScrollBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QScrollBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQScrollBarDragEnterEvent
func callbackQScrollBarDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QScrollBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QScrollBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQScrollBarDragLeaveEvent
func callbackQScrollBarDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QScrollBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QScrollBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQScrollBarDragMoveEvent
func callbackQScrollBarDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QScrollBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QScrollBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQScrollBarDropEvent
func callbackQScrollBarDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScrollBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QScrollBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQScrollBarEnterEvent
func callbackQScrollBarEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QScrollBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QScrollBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQScrollBarFocusInEvent
func callbackQScrollBarFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QScrollBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QScrollBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQScrollBarFocusOutEvent
func callbackQScrollBarFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScrollBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QScrollBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQScrollBarLeaveEvent
func callbackQScrollBarLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QScrollBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QScrollBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQScrollBarMoveEvent
func callbackQScrollBarMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QScrollBar::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QScrollBar) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QScrollBar::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQScrollBarSetVisible
func callbackQScrollBarSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QScrollBar::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QScrollBar::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QScrollBar::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQScrollBarShowEvent
func callbackQScrollBarShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QScrollBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QScrollBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQScrollBarCloseEvent
func callbackQScrollBarCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QScrollBar::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QScrollBar) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QScrollBar::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQScrollBarInitPainter
func callbackQScrollBarInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QScrollBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QScrollBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQScrollBarInputMethodEvent
func callbackQScrollBarInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QScrollBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QScrollBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQScrollBarKeyReleaseEvent
func callbackQScrollBarKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QScrollBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QScrollBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQScrollBarMouseDoubleClickEvent
func callbackQScrollBarMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QScrollBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QScrollBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQScrollBarResizeEvent
func callbackQScrollBarResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QScrollBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QScrollBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQScrollBarTabletEvent
func callbackQScrollBarTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QScrollBar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QScrollBar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQScrollBarChildEvent
func callbackQScrollBarChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QScrollBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QScrollBar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QScrollBar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QScrollBar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQScrollBarCustomEvent
func callbackQScrollBarCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QScrollBar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
