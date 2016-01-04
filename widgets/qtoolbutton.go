package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QToolButton struct {
	QAbstractButton
}

type QToolButton_ITF interface {
	QAbstractButton_ITF
	QToolButton_PTR() *QToolButton
}

func PointerFromQToolButton(ptr QToolButton_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QToolButton_PTR().Pointer()
	}
	return nil
}

func NewQToolButtonFromPointer(ptr unsafe.Pointer) *QToolButton {
	var n = new(QToolButton)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QToolButton_") {
		n.SetObjectName("QToolButton_" + qt.Identifier())
	}
	return n
}

func (ptr *QToolButton) QToolButton_PTR() *QToolButton {
	return ptr
}

//QToolButton::ToolButtonPopupMode
type QToolButton__ToolButtonPopupMode int64

const (
	QToolButton__DelayedPopup    = QToolButton__ToolButtonPopupMode(0)
	QToolButton__MenuButtonPopup = QToolButton__ToolButtonPopupMode(1)
	QToolButton__InstantPopup    = QToolButton__ToolButtonPopupMode(2)
)

func (ptr *QToolButton) ArrowType() core.Qt__ArrowType {
	defer qt.Recovering("QToolButton::arrowType")

	if ptr.Pointer() != nil {
		return core.Qt__ArrowType(C.QToolButton_ArrowType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolButton) AutoRaise() bool {
	defer qt.Recovering("QToolButton::autoRaise")

	if ptr.Pointer() != nil {
		return C.QToolButton_AutoRaise(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QToolButton) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QToolButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QToolButton) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QToolButton::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQToolButtonPaintEvent
func callbackQToolButtonPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	}

}

func (ptr *QToolButton) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QToolButton::paintEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QToolButton) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QToolButton::paintEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QToolButton) PopupMode() QToolButton__ToolButtonPopupMode {
	defer qt.Recovering("QToolButton::popupMode")

	if ptr.Pointer() != nil {
		return QToolButton__ToolButtonPopupMode(C.QToolButton_PopupMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QToolButton) SetArrowType(ty core.Qt__ArrowType) {
	defer qt.Recovering("QToolButton::setArrowType")

	if ptr.Pointer() != nil {
		C.QToolButton_SetArrowType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QToolButton) SetAutoRaise(enable bool) {
	defer qt.Recovering("QToolButton::setAutoRaise")

	if ptr.Pointer() != nil {
		C.QToolButton_SetAutoRaise(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QToolButton) SetPopupMode(mode QToolButton__ToolButtonPopupMode) {
	defer qt.Recovering("QToolButton::setPopupMode")

	if ptr.Pointer() != nil {
		C.QToolButton_SetPopupMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QToolButton) SetToolButtonStyle(style core.Qt__ToolButtonStyle) {
	defer qt.Recovering("QToolButton::setToolButtonStyle")

	if ptr.Pointer() != nil {
		C.QToolButton_SetToolButtonStyle(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QToolButton) ToolButtonStyle() core.Qt__ToolButtonStyle {
	defer qt.Recovering("QToolButton::toolButtonStyle")

	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QToolButton_ToolButtonStyle(ptr.Pointer()))
	}
	return 0
}

func NewQToolButton(parent QWidget_ITF) *QToolButton {
	defer qt.Recovering("QToolButton::QToolButton")

	return NewQToolButtonFromPointer(C.QToolButton_NewQToolButton(PointerFromQWidget(parent)))
}

func (ptr *QToolButton) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QToolButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QToolButton) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QToolButton::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQToolButtonActionEvent
func callbackQToolButtonActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QToolButton) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QToolButton::actionEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QToolButton) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QToolButton::actionEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QToolButton) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QToolButton) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QToolButton::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQToolButtonChangeEvent
func callbackQToolButtonChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QToolButton) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QToolButton::changeEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QToolButton) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QToolButton::changeEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QToolButton) ConnectEnterEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QToolButton) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QToolButton::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQToolButtonEnterEvent
func callbackQToolButtonEnterEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QToolButton) EnterEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QToolButton::enterEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QToolButton) EnterEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QToolButton::enterEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QToolButton) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QToolButton::event")

	if ptr.Pointer() != nil {
		return C.QToolButton_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QToolButton) HitButton(pos core.QPoint_ITF) bool {
	defer qt.Recovering("QToolButton::hitButton")

	if ptr.Pointer() != nil {
		return C.QToolButton_HitButton(ptr.Pointer(), core.PointerFromQPoint(pos)) != 0
	}
	return false
}

func (ptr *QToolButton) ConnectLeaveEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QToolButton::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQToolButtonLeaveEvent
func callbackQToolButtonLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QToolButton) LeaveEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QToolButton::leaveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QToolButton) LeaveEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QToolButton::leaveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QToolButton) Menu() *QMenu {
	defer qt.Recovering("QToolButton::menu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QToolButton_Menu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolButton) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QToolButton::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QToolButton_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolButton) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QToolButton::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQToolButtonMousePressEvent
func callbackQToolButtonMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QToolButton) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolButton::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QToolButton) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolButton::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QToolButton) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QToolButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQToolButtonMouseReleaseEvent
func callbackQToolButtonMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QToolButton) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QToolButton) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolButton::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QToolButton) ConnectNextCheckState(f func()) {
	defer qt.Recovering("connect QToolButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nextCheckState", f)
	}
}

func (ptr *QToolButton) DisconnectNextCheckState() {
	defer qt.Recovering("disconnect QToolButton::nextCheckState")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nextCheckState")
	}
}

//export callbackQToolButtonNextCheckState
func callbackQToolButtonNextCheckState(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QToolButton::nextCheckState")

	if signal := qt.GetSignal(C.GoString(ptrName), "nextCheckState"); signal != nil {
		signal.(func())()
	} else {
		NewQToolButtonFromPointer(ptr).NextCheckStateDefault()
	}
}

func (ptr *QToolButton) NextCheckState() {
	defer qt.Recovering("QToolButton::nextCheckState")

	if ptr.Pointer() != nil {
		C.QToolButton_NextCheckState(ptr.Pointer())
	}
}

func (ptr *QToolButton) NextCheckStateDefault() {
	defer qt.Recovering("QToolButton::nextCheckState")

	if ptr.Pointer() != nil {
		C.QToolButton_NextCheckStateDefault(ptr.Pointer())
	}
}

func (ptr *QToolButton) SetMenu(menu QMenu_ITF) {
	defer qt.Recovering("QToolButton::setMenu")

	if ptr.Pointer() != nil {
		C.QToolButton_SetMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QToolButton) ShowMenu() {
	defer qt.Recovering("QToolButton::showMenu")

	if ptr.Pointer() != nil {
		C.QToolButton_ShowMenu(ptr.Pointer())
	}
}

func (ptr *QToolButton) SizeHint() *core.QSize {
	defer qt.Recovering("QToolButton::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QToolButton_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QToolButton) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QToolButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QToolButton) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QToolButton::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQToolButtonTimerEvent
func callbackQToolButtonTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QToolButton) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QToolButton::timerEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QToolButton) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QToolButton::timerEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QToolButton) ConnectTriggered(f func(action *QAction)) {
	defer qt.Recovering("connect QToolButton::triggered")

	if ptr.Pointer() != nil {
		C.QToolButton_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QToolButton) DisconnectTriggered() {
	defer qt.Recovering("disconnect QToolButton::triggered")

	if ptr.Pointer() != nil {
		C.QToolButton_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQToolButtonTriggered
func callbackQToolButtonTriggered(ptr unsafe.Pointer, ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::triggered")

	if signal := qt.GetSignal(C.GoString(ptrName), "triggered"); signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QToolButton) Triggered(action QAction_ITF) {
	defer qt.Recovering("QToolButton::triggered")

	if ptr.Pointer() != nil {
		C.QToolButton_Triggered(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QToolButton) DestroyQToolButton() {
	defer qt.Recovering("QToolButton::~QToolButton")

	if ptr.Pointer() != nil {
		C.QToolButton_DestroyQToolButton(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QToolButton) ConnectCheckStateSet(f func()) {
	defer qt.Recovering("connect QToolButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "checkStateSet", f)
	}
}

func (ptr *QToolButton) DisconnectCheckStateSet() {
	defer qt.Recovering("disconnect QToolButton::checkStateSet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "checkStateSet")
	}
}

//export callbackQToolButtonCheckStateSet
func callbackQToolButtonCheckStateSet(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QToolButton::checkStateSet")

	if signal := qt.GetSignal(C.GoString(ptrName), "checkStateSet"); signal != nil {
		signal.(func())()
	} else {
		NewQToolButtonFromPointer(ptr).CheckStateSetDefault()
	}
}

func (ptr *QToolButton) CheckStateSet() {
	defer qt.Recovering("QToolButton::checkStateSet")

	if ptr.Pointer() != nil {
		C.QToolButton_CheckStateSet(ptr.Pointer())
	}
}

func (ptr *QToolButton) CheckStateSetDefault() {
	defer qt.Recovering("QToolButton::checkStateSet")

	if ptr.Pointer() != nil {
		C.QToolButton_CheckStateSetDefault(ptr.Pointer())
	}
}

func (ptr *QToolButton) ConnectFocusInEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QToolButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QToolButton) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QToolButton::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQToolButtonFocusInEvent
func callbackQToolButtonFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QToolButton) FocusInEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolButton::focusInEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QToolButton) FocusInEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolButton::focusInEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QToolButton) ConnectFocusOutEvent(f func(e *gui.QFocusEvent)) {
	defer qt.Recovering("connect QToolButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QToolButton) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QToolButton::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQToolButtonFocusOutEvent
func callbackQToolButtonFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(e))
	}
}

func (ptr *QToolButton) FocusOutEvent(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolButton::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QToolButton) FocusOutEventDefault(e gui.QFocusEvent_ITF) {
	defer qt.Recovering("QToolButton::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(e))
	}
}

func (ptr *QToolButton) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QToolButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QToolButton) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QToolButton::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQToolButtonKeyPressEvent
func callbackQToolButtonKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QToolButton) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolButton::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QToolButton) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolButton::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QToolButton) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QToolButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QToolButton) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QToolButton::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQToolButtonKeyReleaseEvent
func callbackQToolButtonKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QToolButton) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolButton::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QToolButton) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QToolButton::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QToolButton) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QToolButton::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQToolButtonMouseMoveEvent
func callbackQToolButtonMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQToolButtonFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QToolButton) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolButton::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QToolButton) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolButton::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QToolButton) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QToolButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QToolButton) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QToolButton::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQToolButtonDragEnterEvent
func callbackQToolButtonDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QToolButton) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QToolButton::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QToolButton) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QToolButton::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QToolButton) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QToolButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QToolButton::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQToolButtonDragLeaveEvent
func callbackQToolButtonDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QToolButton) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QToolButton::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QToolButton) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QToolButton::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QToolButton) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QToolButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QToolButton::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQToolButtonDragMoveEvent
func callbackQToolButtonDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QToolButton) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QToolButton::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QToolButton) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QToolButton::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QToolButton) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QToolButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QToolButton) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QToolButton::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQToolButtonDropEvent
func callbackQToolButtonDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QToolButton) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QToolButton::dropEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QToolButton) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QToolButton::dropEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QToolButton) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QToolButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QToolButton) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QToolButton::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQToolButtonHideEvent
func callbackQToolButtonHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QToolButton) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QToolButton::hideEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QToolButton) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QToolButton::hideEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QToolButton) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QToolButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QToolButton::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQToolButtonMoveEvent
func callbackQToolButtonMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QToolButton) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QToolButton::moveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QToolButton) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QToolButton::moveEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QToolButton) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QToolButton::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QToolButton) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QToolButton::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQToolButtonSetVisible
func callbackQToolButtonSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QToolButton::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QToolButton) SetVisible(visible bool) {
	defer qt.Recovering("QToolButton::setVisible")

	if ptr.Pointer() != nil {
		C.QToolButton_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QToolButton) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QToolButton::setVisible")

	if ptr.Pointer() != nil {
		C.QToolButton_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QToolButton) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QToolButton::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QToolButton) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QToolButton::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQToolButtonShowEvent
func callbackQToolButtonShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QToolButton) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QToolButton::showEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QToolButton) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QToolButton::showEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QToolButton) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QToolButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QToolButton) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QToolButton::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQToolButtonCloseEvent
func callbackQToolButtonCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QToolButton) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QToolButton::closeEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QToolButton) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QToolButton::closeEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QToolButton) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QToolButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QToolButton) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QToolButton::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQToolButtonContextMenuEvent
func callbackQToolButtonContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QToolButton) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QToolButton::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QToolButton) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QToolButton::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QToolButton) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QToolButton::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QToolButton) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QToolButton::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQToolButtonInitPainter
func callbackQToolButtonInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQToolButtonFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QToolButton) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QToolButton::initPainter")

	if ptr.Pointer() != nil {
		C.QToolButton_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QToolButton) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QToolButton::initPainter")

	if ptr.Pointer() != nil {
		C.QToolButton_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QToolButton) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QToolButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QToolButton) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QToolButton::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQToolButtonInputMethodEvent
func callbackQToolButtonInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QToolButton) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QToolButton::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QToolButton) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QToolButton::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QToolButton) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QToolButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QToolButton) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QToolButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQToolButtonMouseDoubleClickEvent
func callbackQToolButtonMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QToolButton) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolButton) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QToolButton::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QToolButton) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QToolButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QToolButton) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QToolButton::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQToolButtonResizeEvent
func callbackQToolButtonResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QToolButton) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QToolButton::resizeEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QToolButton) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QToolButton::resizeEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QToolButton) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QToolButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QToolButton) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QToolButton::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQToolButtonTabletEvent
func callbackQToolButtonTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QToolButton) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QToolButton::tabletEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QToolButton) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QToolButton::tabletEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QToolButton) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QToolButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QToolButton) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QToolButton::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQToolButtonWheelEvent
func callbackQToolButtonWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QToolButton) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QToolButton::wheelEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QToolButton) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QToolButton::wheelEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QToolButton) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QToolButton::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QToolButton) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QToolButton::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQToolButtonChildEvent
func callbackQToolButtonChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QToolButton) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QToolButton::childEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QToolButton) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QToolButton::childEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QToolButton) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QToolButton::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QToolButton) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QToolButton::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQToolButtonCustomEvent
func callbackQToolButtonCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QToolButton::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQToolButtonFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QToolButton) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QToolButton::customEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QToolButton) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QToolButton::customEvent")

	if ptr.Pointer() != nil {
		C.QToolButton_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
