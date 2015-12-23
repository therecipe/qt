package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QRadioButton struct {
	QAbstractButton
}

type QRadioButton_ITF interface {
	QAbstractButton_ITF
	QRadioButton_PTR() *QRadioButton
}

func PointerFromQRadioButton(ptr QRadioButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioButton_PTR().Pointer()
	}
	return nil
}

func NewQRadioButtonFromPointer(ptr unsafe.Pointer) *QRadioButton {
	var n = new(QRadioButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRadioButton_") {
		n.SetObjectName("QRadioButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QRadioButton) QRadioButton_PTR() *QRadioButton {
	return ptr
}

func NewQRadioButton(parent QWidget_ITF) *QRadioButton {
	defer qt.Recovering("QRadioButton::QRadioButton")

	return NewQRadioButtonFromPointer(C.QRadioButton_NewQRadioButton(PointerFromQWidget(parent)))
}

func NewQRadioButton2(text string, parent QWidget_ITF) *QRadioButton {
	defer qt.Recovering("QRadioButton::QRadioButton")

	return NewQRadioButtonFromPointer(C.QRadioButton_NewQRadioButton2(C.CString(text), PointerFromQWidget(parent)))
}

func (ptr *QRadioButton) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QRadioButton::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QRadioButton_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioButton) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRadioButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QRadioButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQRadioButtonMouseMoveEvent
func callbackQRadioButtonMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectPaintEvent(f func(v *gui.QPaintEvent)) {
	defer qt.Recovering("connect QRadioButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QRadioButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQRadioButtonPaintEvent
func callbackQRadioButtonPaintEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(v))
		return true
	}
	return false

}

func (ptr *QRadioButton) SizeHint() *core.QSize {
	defer qt.Recovering("QRadioButton::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QRadioButton_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioButton) DestroyQRadioButton() {
	defer qt.Recovering("QRadioButton::~QRadioButton")

	if ptr.Pointer() != nil {
		C.QRadioButton_DestroyQRadioButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QRadioButton) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QRadioButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QRadioButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQRadioButtonChangeEvent
func callbackQRadioButtonChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QRadioButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QRadioButton) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QRadioButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQRadioButtonCheckStateSet
func callbackQRadioButtonCheckStateSet(ptrName *C.char) bool {
	defer qt.Recovering("callback QRadioButton::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QRadioButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QRadioButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQRadioButtonFocusInEvent
func callbackQRadioButtonFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QRadioButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QRadioButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQRadioButtonFocusOutEvent
func callbackQRadioButtonFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QRadioButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QRadioButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQRadioButtonKeyPressEvent
func callbackQRadioButtonKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QRadioButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QRadioButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQRadioButtonKeyReleaseEvent
func callbackQRadioButtonKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRadioButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QRadioButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQRadioButtonMousePressEvent
func callbackQRadioButtonMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRadioButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QRadioButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQRadioButtonMouseReleaseEvent
func callbackQRadioButtonMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QRadioButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QRadioButton) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QRadioButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQRadioButtonNextCheckState
func callbackQRadioButtonNextCheckState(ptrName *C.char) bool {
	defer qt.Recovering("callback QRadioButton::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QRadioButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QRadioButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQRadioButtonTimerEvent
func callbackQRadioButtonTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QRadioButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QRadioButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQRadioButtonActionEvent
func callbackQRadioButtonActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QRadioButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QRadioButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQRadioButtonDragEnterEvent
func callbackQRadioButtonDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QRadioButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QRadioButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQRadioButtonDragLeaveEvent
func callbackQRadioButtonDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QRadioButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QRadioButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQRadioButtonDragMoveEvent
func callbackQRadioButtonDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QRadioButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QRadioButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQRadioButtonDropEvent
func callbackQRadioButtonDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QRadioButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQRadioButtonEnterEvent
func callbackQRadioButtonEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QRadioButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QRadioButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQRadioButtonHideEvent
func callbackQRadioButtonHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QRadioButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQRadioButtonLeaveEvent
func callbackQRadioButtonLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QRadioButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QRadioButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQRadioButtonMoveEvent
func callbackQRadioButtonMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QRadioButton::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QRadioButton) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QRadioButton::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQRadioButtonSetVisible
func callbackQRadioButtonSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QRadioButton::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QRadioButton::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QRadioButton::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQRadioButtonShowEvent
func callbackQRadioButtonShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QRadioButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QRadioButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQRadioButtonCloseEvent
func callbackQRadioButtonCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QRadioButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QRadioButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQRadioButtonContextMenuEvent
func callbackQRadioButtonContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QRadioButton::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QRadioButton) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QRadioButton::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQRadioButtonInitPainter
func callbackQRadioButtonInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QRadioButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QRadioButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQRadioButtonInputMethodEvent
func callbackQRadioButtonInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QRadioButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QRadioButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQRadioButtonMouseDoubleClickEvent
func callbackQRadioButtonMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QRadioButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QRadioButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQRadioButtonResizeEvent
func callbackQRadioButtonResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QRadioButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QRadioButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQRadioButtonTabletEvent
func callbackQRadioButtonTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QRadioButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QRadioButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQRadioButtonWheelEvent
func callbackQRadioButtonWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QRadioButton::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QRadioButton::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQRadioButtonChildEvent
func callbackQRadioButtonChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QRadioButton) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QRadioButton::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QRadioButton) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QRadioButton::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQRadioButtonCustomEvent
func callbackQRadioButtonCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QRadioButton::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
