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
func callbackQAbstractButtonChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQAbstractButtonFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QAbstractButton) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QAbstractButton::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QAbstractButton) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QAbstractButton::changeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
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
func callbackQAbstractButtonCheckStateSet(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractButton::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractButtonFromPointer(ptr).CheckStateSetDefault()
	}
}

func (ptr *QAbstractButton) CheckStateSet() {
	defer qt.Recovering("QAbstractButton::checkStateSet")

	if ptr.Pointer() != nil {
		C.QAbstractButton_CheckStateSet(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) CheckStateSetDefault() {
	defer qt.Recovering("QAbstractButton::checkStateSet")

	if ptr.Pointer() != nil {
		C.QAbstractButton_CheckStateSetDefault(ptr.Pointer())
	}
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
func callbackQAbstractButtonClicked(ptr unsafe.Pointer, ptrName *C.char, checked C.int) {
	defer qt.Recovering("callback QAbstractButton::clicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "clicked"); signal != nil {
		signal.(func(bool))(int(checked) != 0)
	}

}

func (ptr *QAbstractButton) Clicked(checked bool) {
	defer qt.Recovering("QAbstractButton::clicked")

	if ptr.Pointer() != nil {
		C.QAbstractButton_Clicked(ptr.Pointer(), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QAbstractButton) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAbstractButton::event")

	if ptr.Pointer() != nil {
		return C.QAbstractButton_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
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
func callbackQAbstractButtonFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQAbstractButtonFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QAbstractButton) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractButton::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QAbstractButton) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractButton::focusInEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
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
func callbackQAbstractButtonFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQAbstractButtonFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QAbstractButton) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractButton::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QAbstractButton) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QAbstractButton::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QAbstractButton) Group() *QButtonGroup {
	defer qt.Recovering("QAbstractButton::group")

	if ptr.Pointer() != nil {
		return NewQButtonGroupFromPointer(C.QAbstractButton_Group(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAbstractButton) HitButton(pos core.QPoint_ITF) bool {
	defer qt.Recovering("QAbstractButton::hitButton")

	if ptr.Pointer() != nil {
		return C.QAbstractButton_HitButton(ptr.Pointer(), core.PointerFromQPoint(pos)) != 0
	}
	return false
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
func callbackQAbstractButtonKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQAbstractButtonFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QAbstractButton) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractButton::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QAbstractButton) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractButton::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQAbstractButtonKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQAbstractButtonFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QAbstractButton) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractButton::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QAbstractButton) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QAbstractButton::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
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
func callbackQAbstractButtonMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQAbstractButtonFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QAbstractButton) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractButton) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQAbstractButtonMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQAbstractButtonFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QAbstractButton) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractButton) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQAbstractButtonMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQAbstractButtonFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QAbstractButton) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QAbstractButton) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
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
func callbackQAbstractButtonNextCheckState(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractButton::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
	} else {
		NewQAbstractButtonFromPointer(ptr).NextCheckStateDefault()
	}
}

func (ptr *QAbstractButton) NextCheckState() {
	defer qt.Recovering("QAbstractButton::nextCheckState")

	if ptr.Pointer() != nil {
		C.QAbstractButton_NextCheckState(ptr.Pointer())
	}
}

func (ptr *QAbstractButton) NextCheckStateDefault() {
	defer qt.Recovering("QAbstractButton::nextCheckState")

	if ptr.Pointer() != nil {
		C.QAbstractButton_NextCheckStateDefault(ptr.Pointer())
	}
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
func callbackQAbstractButtonPressed(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractButton::pressed")

	if signal := qt.GetSignal(C.GoString(ptrName), "pressed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractButton) Pressed() {
	defer qt.Recovering("QAbstractButton::pressed")

	if ptr.Pointer() != nil {
		C.QAbstractButton_Pressed(ptr.Pointer())
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
func callbackQAbstractButtonReleased(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAbstractButton::released")

	if signal := qt.GetSignal(C.GoString(ptrName), "released"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAbstractButton) Released() {
	defer qt.Recovering("QAbstractButton::released")

	if ptr.Pointer() != nil {
		C.QAbstractButton_Released(ptr.Pointer())
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
func callbackQAbstractButtonTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQAbstractButtonFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QAbstractButton) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractButton::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QAbstractButton) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractButton::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
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
func callbackQAbstractButtonToggled(ptr unsafe.Pointer, ptrName *C.char, checked C.int) {
	defer qt.Recovering("callback QAbstractButton::toggled")

	if signal := qt.GetSignal(C.GoString(ptrName), "toggled"); signal != nil {
		signal.(func(bool))(int(checked) != 0)
	}

}

func (ptr *QAbstractButton) Toggled(checked bool) {
	defer qt.Recovering("QAbstractButton::toggled")

	if ptr.Pointer() != nil {
		C.QAbstractButton_Toggled(ptr.Pointer(), C.int(qt.GoBoolToInt(checked)))
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
func callbackQAbstractButtonActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractButton::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QAbstractButton) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QAbstractButton::actionEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
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
func callbackQAbstractButtonDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractButton::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QAbstractButton) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QAbstractButton::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQAbstractButtonDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractButton::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QAbstractButton) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QAbstractButton::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQAbstractButtonDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractButton::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QAbstractButton) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QAbstractButton::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQAbstractButtonDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractButton::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QAbstractButton) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QAbstractButton::dropEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQAbstractButtonEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractButton::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractButton) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractButton::enterEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQAbstractButtonHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractButton::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QAbstractButton) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QAbstractButton::hideEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQAbstractButtonLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractButton::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractButton) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractButton::leaveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQAbstractButtonMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractButton::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QAbstractButton) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QAbstractButton::moveEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
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
func callbackQAbstractButtonSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QAbstractButton::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QAbstractButton) SetVisible(visible bool) {
	defer qt.Recovering("QAbstractButton::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QAbstractButton) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QAbstractButton::setVisible")

	if ptr.Pointer() != nil {
		C.QAbstractButton_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
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
func callbackQAbstractButtonShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractButton::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QAbstractButton) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QAbstractButton::showEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
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
func callbackQAbstractButtonCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QAbstractButton) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::closeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
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
func callbackQAbstractButtonContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractButton::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QAbstractButton) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QAbstractButton::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
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
func callbackQAbstractButtonInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQAbstractButtonFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QAbstractButton) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QAbstractButton::initPainter")

	if ptr.Pointer() != nil {
		C.QAbstractButton_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QAbstractButton) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QAbstractButton::initPainter")

	if ptr.Pointer() != nil {
		C.QAbstractButton_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
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
func callbackQAbstractButtonInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractButton::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QAbstractButton) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QAbstractButton::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
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
func callbackQAbstractButtonMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QAbstractButton) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QAbstractButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQAbstractButtonResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractButton::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QAbstractButton) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QAbstractButton::resizeEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
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
func callbackQAbstractButtonTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractButton::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QAbstractButton) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QAbstractButton::tabletEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
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
func callbackQAbstractButtonWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractButton::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QAbstractButton) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QAbstractButton::wheelEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
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
func callbackQAbstractButtonChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractButton::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractButton) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractButton::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQAbstractButtonCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractButton::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractButtonFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractButton) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractButton::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractButton) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractButton::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractButton_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
