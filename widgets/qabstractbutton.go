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

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "checkStateSet")
	if signal != nil {
		defer signal.(func())()
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

	var signal = qt.GetSignal(C.GoString(ptrName), "clicked")
	if signal != nil {
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

	var signal = qt.GetSignal(C.GoString(ptrName), "focusInEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "focusOutEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "nextCheckState")
	if signal != nil {
		defer signal.(func())()
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

	var signal = qt.GetSignal(C.GoString(ptrName), "pressed")
	if signal != nil {
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

	var signal = qt.GetSignal(C.GoString(ptrName), "released")
	if signal != nil {
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

	var signal = qt.GetSignal(C.GoString(ptrName), "timerEvent")
	if signal != nil {
		defer signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
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

	var signal = qt.GetSignal(C.GoString(ptrName), "toggled")
	if signal != nil {
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
