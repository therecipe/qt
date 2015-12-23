package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAbstractButton struct {
	QWidget
}

type QAbstractButton_ITF interface {
	QWidget_ITF
	QAbstractButton_PTR() *QAbstractButton
}

func PointerFromQAbstractButton(ptr QAbstractButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractButton_PTR().Pointer()
	}
	return nil
}

func NewQAbstractButtonFromPointer(ptr unsafe.Pointer) *QAbstractButton {
	var n = new(QAbstractButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractButton_") {
		n.SetObjectName("QAbstractButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractButton) QAbstractButton_PTR() *QAbstractButton {
	return ptr
}

func (ptr *QAbstractButton) AutoExclusive() bool {
	defer qt.Recovering("QAbstractButton::autoExclusive")

	if ptr.Pointer() != nil {
		return C.QAbstractButton_AutoExclusive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) AutoRepeat() bool {
	defer qt.Recovering("QAbstractButton::autoRepeat")

	if ptr.Pointer() != nil {
		return C.QAbstractButton_AutoRepeat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) AutoRepeatDelay() int {
	defer qt.Recovering("QAbstractButton::autoRepeatDelay")

	if ptr.Pointer() != nil {
		return int(C.QAbstractButton_AutoRepeatDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractButton) AutoRepeatInterval() int {
	defer qt.Recovering("QAbstractButton::autoRepeatInterval")

	if ptr.Pointer() != nil {
		return int(C.QAbstractButton_AutoRepeatInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractButton) Icon() *gui.QIcon {
	defer qt.Recovering("QAbstractButton::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QAbstractButton_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractButton) IconSize() *core.QSize {
	defer qt.Recovering("QAbstractButton::iconSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QAbstractButton_IconSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractButton) IsCheckable() bool {
	defer qt.Recovering("QAbstractButton::isCheckable")

	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) IsChecked() bool {
	defer qt.Recovering("QAbstractButton::isChecked")

	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsChecked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) IsDown() bool {
	defer qt.Recovering("QAbstractButton::isDown")

	if ptr.Pointer() != nil {
		return C.QAbstractButton_IsDown(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAbstractButton) SetAutoExclusive(v bool) {
	defer qt.Recovering("QAbstractButton::setAutoExclusive")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoExclusive(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetAutoRepeat(v bool) {
	defer qt.Recovering("QAbstractButton::setAutoRepeat")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeat(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetAutoRepeatDelay(v int) {
	defer qt.Recovering("QAbstractButton::setAutoRepeatDelay")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatDelay(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractButton) SetAutoRepeatInterval(v int) {
	defer qt.Recovering("QAbstractButton::setAutoRepeatInterval")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetAutoRepeatInterval(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QAbstractButton) SetCheckable(v bool) {
	defer qt.Recovering("QAbstractButton::setCheckable")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetChecked(v bool) {
	defer qt.Recovering("QAbstractButton::setChecked")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetChecked(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetDown(v bool) {
	defer qt.Recovering("QAbstractButton::setDown")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetDown(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAbstractButton) SetIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QAbstractButton::setIcon")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QAbstractButton) SetIconSize(size core.QSize_ITF) {
	defer qt.Recovering("QAbstractButton::setIconSize")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetIconSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QAbstractButton) SetShortcut(key gui.QKeySequence_ITF) {
	defer qt.Recovering("QAbstractButton::setShortcut")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(key))
	}
}

func (ptr *QAbstractButton) SetText(text string) {
	defer qt.Recovering("QAbstractButton::setText")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QAbstractButton) Text() string {
	defer qt.Recovering("QAbstractButton::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAbstractButton_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAbstractButton) Toggle() {
	defer qt.Recovering("QAbstractButton::toggle")

	if ptr.Pointer() != nil {
		C.QAbstractButton_Toggle(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) AnimateClick(msec int) {
	defer qt.Recovering("QAbstractButton::animateClick")

	if ptr.Pointer() != nil {
		C.QAbstractButton_AnimateClick(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QAbstractButton) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QAbstractButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QAbstractButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQAbstractButtonChangeEvent
func callbackQAbstractButtonChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QAbstractButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QAbstractButton) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QAbstractButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQAbstractButtonCheckStateSet
func callbackQAbstractButtonCheckStateSet(ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractButton::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractButton) Click() {
	defer qt.Recovering("QAbstractButton::click")

	if ptr.Pointer() != nil {
		C.QAbstractButton_Click(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) ConnectClicked(f func(checked bool)) {
	defer qt.Recovering("connect QAbstractButton::clicked")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QAbstractButton) DisconnectClicked() {
	defer qt.Recovering("disconnect QAbstractButton::clicked")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQAbstractButtonClicked
func callbackQAbstractButtonClicked(ptrName *C.char, checked C.int) {
	defer qt.Recovering("callback QAbstractButton::clicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "clicked"); signal != nil {
		signal.(func(bool))(int(checked) != 0)
	}

}

func (ptr *QAbstractButton) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QAbstractButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQAbstractButtonFocusInEvent
func callbackQAbstractButtonFocusInEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QAbstractButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QAbstractButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQAbstractButtonFocusOutEvent
func callbackQAbstractButtonFocusOutEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractButton) Group() *QButtonGroup {
	defer qt.Recovering("QAbstractButton::group")

	if ptr.Pointer() != nil {
		return NewQButtonGroupFromPointer(C.QAbstractButton_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractButton) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QAbstractButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQAbstractButtonKeyPressEvent
func callbackQAbstractButtonKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QAbstractButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQAbstractButtonKeyReleaseEvent
func callbackQAbstractButtonKeyReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QAbstractButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQAbstractButtonMouseMoveEvent
func callbackQAbstractButtonMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QAbstractButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQAbstractButtonMousePressEvent
func callbackQAbstractButtonMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QAbstractButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQAbstractButtonMouseReleaseEvent
func callbackQAbstractButtonMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QAbstractButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QAbstractButton) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QAbstractButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQAbstractButtonNextCheckState
func callbackQAbstractButtonNextCheckState(ptrName *C.char) bool {
	defer qt.Recovering("callback QAbstractButton::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectPressed(f func()) {
	defer qt.Recovering("connect QAbstractButton::pressed")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectPressed(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pressed", f)
	}
}

func (ptr *QAbstractButton) DisconnectPressed() {
	defer qt.Recovering("disconnect QAbstractButton::pressed")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectPressed(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pressed")
	}
}

//export callbackQAbstractButtonPressed
func callbackQAbstractButtonPressed(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractButton::pressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "pressed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractButton) ConnectReleased(f func()) {
	defer qt.Recovering("connect QAbstractButton::released")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectReleased(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "released", f)
	}
}

func (ptr *QAbstractButton) DisconnectReleased() {
	defer qt.Recovering("disconnect QAbstractButton::released")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectReleased(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "released")
	}
}

//export callbackQAbstractButtonReleased
func callbackQAbstractButtonReleased(ptrName *C.char) {
	defer qt.Recovering("callback QAbstractButton::released")

	if signal := qt.GetSignal(C.GoString(ptrName), "released"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractButton) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractButtonTimerEvent
func callbackQAbstractButtonTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectToggled(f func(checked bool)) {
	defer qt.Recovering("connect QAbstractButton::toggled")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ConnectToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QAbstractButton) DisconnectToggled() {
	defer qt.Recovering("disconnect QAbstractButton::toggled")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DisconnectToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQAbstractButtonToggled
func callbackQAbstractButtonToggled(ptrName *C.char, checked C.int) {
	defer qt.Recovering("callback QAbstractButton::toggled")

	if signal := qt.GetSignal(C.GoString(ptrName), "toggled"); signal != nil {
		signal.(func(bool))(int(checked) != 0)
	}

}

func (ptr *QAbstractButton) DestroyQAbstractButton() {
	defer qt.Recovering("QAbstractButton::~QAbstractButton")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DestroyQAbstractButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractButton) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QAbstractButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QAbstractButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQAbstractButtonActionEvent
func callbackQAbstractButtonActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QAbstractButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QAbstractButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQAbstractButtonDragEnterEvent
func callbackQAbstractButtonDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QAbstractButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQAbstractButtonDragLeaveEvent
func callbackQAbstractButtonDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QAbstractButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QAbstractButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQAbstractButtonDragMoveEvent
func callbackQAbstractButtonDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QAbstractButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QAbstractButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQAbstractButtonDropEvent
func callbackQAbstractButtonDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QAbstractButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQAbstractButtonEnterEvent
func callbackQAbstractButtonEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QAbstractButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QAbstractButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQAbstractButtonHideEvent
func callbackQAbstractButtonHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QAbstractButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQAbstractButtonLeaveEvent
func callbackQAbstractButtonLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QAbstractButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QAbstractButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQAbstractButtonMoveEvent
func callbackQAbstractButtonMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QAbstractButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QAbstractButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQAbstractButtonPaintEvent
func callbackQAbstractButtonPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QAbstractButton::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QAbstractButton) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QAbstractButton::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQAbstractButtonSetVisible
func callbackQAbstractButtonSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QAbstractButton::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QAbstractButton::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QAbstractButton::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQAbstractButtonShowEvent
func callbackQAbstractButtonShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QAbstractButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QAbstractButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQAbstractButtonCloseEvent
func callbackQAbstractButtonCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QAbstractButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QAbstractButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQAbstractButtonContextMenuEvent
func callbackQAbstractButtonContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QAbstractButton::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QAbstractButton) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QAbstractButton::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQAbstractButtonInitPainter
func callbackQAbstractButtonInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QAbstractButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QAbstractButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQAbstractButtonInputMethodEvent
func callbackQAbstractButtonInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QAbstractButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QAbstractButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQAbstractButtonMouseDoubleClickEvent
func callbackQAbstractButtonMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QAbstractButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QAbstractButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQAbstractButtonResizeEvent
func callbackQAbstractButtonResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QAbstractButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QAbstractButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQAbstractButtonTabletEvent
func callbackQAbstractButtonTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QAbstractButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QAbstractButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQAbstractButtonWheelEvent
func callbackQAbstractButtonWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractButton::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractButton::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractButtonChildEvent
func callbackQAbstractButtonChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QAbstractButton) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractButton::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractButton) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractButton::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractButtonCustomEvent
func callbackQAbstractButtonCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QAbstractButton::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
