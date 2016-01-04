package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMenuBar struct {
	QWidget
}

type QMenuBar_ITF interface {
	QWidget_ITF
	QMenuBar_PTR() *QMenuBar
}

func PointerFromQMenuBar(ptr QMenuBar_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenuBar_PTR().Pointer()
	}
	return nil
}

func NewQMenuBarFromPointer(ptr unsafe.Pointer) *QMenuBar {
	var n = new(QMenuBar)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMenuBar_") {
		n.SetObjectName("QMenuBar_" + qt.Identifier())
	}
	return n
}

func (ptr *QMenuBar) QMenuBar_PTR() *QMenuBar {
	return ptr
}

func (ptr *QMenuBar) IsDefaultUp() bool {
	defer qt.Recovering("QMenuBar::isDefaultUp")

	if ptr.Pointer() != nil {
		return C.QMenuBar_IsDefaultUp(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenuBar) IsNativeMenuBar() bool {
	defer qt.Recovering("QMenuBar::isNativeMenuBar")

	if ptr.Pointer() != nil {
		return C.QMenuBar_IsNativeMenuBar(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenuBar) SetCornerWidget(widget QWidget_ITF, corner core.Qt__Corner) {
	defer qt.Recovering("QMenuBar::setCornerWidget")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetCornerWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(corner))
	}
}

func (ptr *QMenuBar) SetDefaultUp(v bool) {
	defer qt.Recovering("QMenuBar::setDefaultUp")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetDefaultUp(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	defer qt.Recovering("QMenuBar::setNativeMenuBar")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetNativeMenuBar(ptr.Pointer(), C.int(qt.GoBoolToInt(nativeMenuBar)))
	}
}

func NewQMenuBar(parent QWidget_ITF) *QMenuBar {
	defer qt.Recovering("QMenuBar::QMenuBar")

	return NewQMenuBarFromPointer(C.QMenuBar_NewQMenuBar(PointerFromQWidget(parent)))
}

func (ptr *QMenuBar) ActionAt(pt core.QPoint_ITF) *QAction {
	defer qt.Recovering("QMenuBar::actionAt")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_ActionAt(ptr.Pointer(), core.PointerFromQPoint(pt)))
	}
	return nil
}

func (ptr *QMenuBar) ConnectActionEvent(f func(e *gui.QActionEvent)) {
	defer qt.Recovering("connect QMenuBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMenuBar::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMenuBarActionEvent
func callbackQMenuBarActionEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(e))
	} else {
		NewQMenuBarFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(e))
	}
}

func (ptr *QMenuBar) ActionEvent(e gui.QActionEvent_ITF) {
	defer qt.Recovering("QMenuBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(e))
	}
}

func (ptr *QMenuBar) ActionEventDefault(e gui.QActionEvent_ITF) {
	defer qt.Recovering("QMenuBar::actionEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(e))
	}
}

func (ptr *QMenuBar) ActionGeometry(act QAction_ITF) *core.QRect {
	defer qt.Recovering("QMenuBar::actionGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QMenuBar_ActionGeometry(ptr.Pointer(), PointerFromQAction(act)))
	}
	return nil
}

func (ptr *QMenuBar) ActiveAction() *QAction {
	defer qt.Recovering("QMenuBar::activeAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_ActiveAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) AddAction(text string) *QAction {
	defer qt.Recovering("QMenuBar::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddAction(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QMenuBar) AddAction2(text string, receiver core.QObject_ITF, member string) *QAction {
	defer qt.Recovering("QMenuBar::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddAction2(ptr.Pointer(), C.CString(text), core.PointerFromQObject(receiver), C.CString(member)))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu(menu QMenu_ITF) *QAction {
	defer qt.Recovering("QMenuBar::addMenu")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddMenu(ptr.Pointer(), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu3(icon gui.QIcon_ITF, title string) *QMenu {
	defer qt.Recovering("QMenuBar::addMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenuBar_AddMenu3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(title)))
	}
	return nil
}

func (ptr *QMenuBar) AddMenu2(title string) *QMenu {
	defer qt.Recovering("QMenuBar::addMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenuBar_AddMenu2(ptr.Pointer(), C.CString(title)))
	}
	return nil
}

func (ptr *QMenuBar) AddSeparator() *QAction {
	defer qt.Recovering("QMenuBar::addSeparator")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_AddSeparator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QMenuBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QMenuBar::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQMenuBarChangeEvent
func callbackQMenuBarChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQMenuBarFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QMenuBar) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QMenuBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QMenuBar) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QMenuBar::changeEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QMenuBar) Clear() {
	defer qt.Recovering("QMenuBar::clear")

	if ptr.Pointer() != nil {
		C.QMenuBar_Clear(ptr.Pointer())
	}
}

func (ptr *QMenuBar) CornerWidget(corner core.Qt__Corner) *QWidget {
	defer qt.Recovering("QMenuBar::cornerWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMenuBar_CornerWidget(ptr.Pointer(), C.int(corner)))
	}
	return nil
}

func (ptr *QMenuBar) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QMenuBar::event")

	if ptr.Pointer() != nil {
		return C.QMenuBar_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMenuBar) EventFilter(object core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QMenuBar::eventFilter")

	if ptr.Pointer() != nil {
		return C.QMenuBar_EventFilter(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QMenuBar) ConnectFocusInEvent(f func(v *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMenuBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QMenuBar::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQMenuBarFocusInEvent
func callbackQMenuBarFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(v))
	} else {
		NewQMenuBarFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(v))
	}
}

func (ptr *QMenuBar) FocusInEvent(v gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMenuBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(v))
	}
}

func (ptr *QMenuBar) FocusInEventDefault(v gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMenuBar::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(v))
	}
}

func (ptr *QMenuBar) ConnectFocusOutEvent(f func(v *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMenuBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMenuBar::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMenuBarFocusOutEvent
func callbackQMenuBarFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(v))
	} else {
		NewQMenuBarFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(v))
	}
}

func (ptr *QMenuBar) FocusOutEvent(v gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMenuBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(v))
	}
}

func (ptr *QMenuBar) FocusOutEventDefault(v gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMenuBar::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(v))
	}
}

func (ptr *QMenuBar) HeightForWidth(v int) int {
	defer qt.Recovering("QMenuBar::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QMenuBar_HeightForWidth(ptr.Pointer(), C.int(v)))
	}
	return 0
}

func (ptr *QMenuBar) ConnectHovered(f func(action *QAction)) {
	defer qt.Recovering("connect QMenuBar::hovered")

	if ptr.Pointer() != nil {
		C.QMenuBar_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QMenuBar) DisconnectHovered() {
	defer qt.Recovering("disconnect QMenuBar::hovered")

	if ptr.Pointer() != nil {
		C.QMenuBar_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQMenuBarHovered
func callbackQMenuBarHovered(ptr unsafe.Pointer, ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::hovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "hovered"); signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QMenuBar) Hovered(action QAction_ITF) {
	defer qt.Recovering("QMenuBar::hovered")

	if ptr.Pointer() != nil {
		C.QMenuBar_Hovered(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QMenuBar) InsertMenu(before QAction_ITF, menu QMenu_ITF) *QAction {
	defer qt.Recovering("QMenuBar::insertMenu")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_InsertMenu(ptr.Pointer(), PointerFromQAction(before), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenuBar) InsertSeparator(before QAction_ITF) *QAction {
	defer qt.Recovering("QMenuBar::insertSeparator")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenuBar_InsertSeparator(ptr.Pointer(), PointerFromQAction(before)))
	}
	return nil
}

func (ptr *QMenuBar) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMenuBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMenuBar::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMenuBarKeyPressEvent
func callbackQMenuBarKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQMenuBarFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QMenuBar) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMenuBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QMenuBar) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMenuBar::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QMenuBar) ConnectLeaveEvent(f func(v *core.QEvent)) {
	defer qt.Recovering("connect QMenuBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMenuBar::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMenuBarLeaveEvent
func callbackQMenuBarLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(v))
	} else {
		NewQMenuBarFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(v))
	}
}

func (ptr *QMenuBar) LeaveEvent(v core.QEvent_ITF) {
	defer qt.Recovering("QMenuBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(v))
	}
}

func (ptr *QMenuBar) LeaveEventDefault(v core.QEvent_ITF) {
	defer qt.Recovering("QMenuBar::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(v))
	}
}

func (ptr *QMenuBar) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QMenuBar::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMenuBar_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenuBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMenuBar::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMenuBarMouseMoveEvent
func callbackQMenuBarMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMenuBarFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMenuBar) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenuBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenuBar) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenuBar::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenuBar) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenuBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMenuBar::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMenuBarMousePressEvent
func callbackQMenuBarMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMenuBarFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMenuBar) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenuBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenuBar) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenuBar::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenuBar) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenuBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMenuBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMenuBarMouseReleaseEvent
func callbackQMenuBarMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMenuBarFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMenuBar) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenuBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenuBar) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenuBar::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenuBar) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMenuBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMenuBar::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMenuBarPaintEvent
func callbackQMenuBarPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQMenuBarFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QMenuBar) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMenuBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QMenuBar) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMenuBar::paintEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QMenuBar) ConnectResizeEvent(f func(v *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMenuBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMenuBar::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMenuBarResizeEvent
func callbackQMenuBarResizeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(v))
	} else {
		NewQMenuBarFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(v))
	}
}

func (ptr *QMenuBar) ResizeEvent(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMenuBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QMenuBar) ResizeEventDefault(v gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMenuBar::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(v))
	}
}

func (ptr *QMenuBar) SetActiveAction(act QAction_ITF) {
	defer qt.Recovering("QMenuBar::setActiveAction")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetActiveAction(ptr.Pointer(), PointerFromQAction(act))
	}
}

func (ptr *QMenuBar) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMenuBar::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMenuBar) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMenuBar::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMenuBarSetVisible
func callbackQMenuBarSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMenuBar::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMenuBar) SetVisible(visible bool) {
	defer qt.Recovering("QMenuBar::setVisible")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenuBar) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QMenuBar::setVisible")

	if ptr.Pointer() != nil {
		C.QMenuBar_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenuBar) SizeHint() *core.QSize {
	defer qt.Recovering("QMenuBar::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMenuBar_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenuBar) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QMenuBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMenuBar::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMenuBarTimerEvent
func callbackQMenuBarTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQMenuBarFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QMenuBar) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QMenuBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QMenuBar) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QMenuBar::timerEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QMenuBar) ConnectTriggered(f func(action *QAction)) {
	defer qt.Recovering("connect QMenuBar::triggered")

	if ptr.Pointer() != nil {
		C.QMenuBar_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QMenuBar) DisconnectTriggered() {
	defer qt.Recovering("disconnect QMenuBar::triggered")

	if ptr.Pointer() != nil {
		C.QMenuBar_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQMenuBarTriggered
func callbackQMenuBarTriggered(ptr unsafe.Pointer, ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::triggered")

	if signal := qt.GetSignal(C.GoString(ptrName), "triggered"); signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QMenuBar) Triggered(action QAction_ITF) {
	defer qt.Recovering("QMenuBar::triggered")

	if ptr.Pointer() != nil {
		C.QMenuBar_Triggered(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QMenuBar) DestroyQMenuBar() {
	defer qt.Recovering("QMenuBar::~QMenuBar")

	if ptr.Pointer() != nil {
		C.QMenuBar_DestroyQMenuBar(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMenuBar) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QMenuBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QMenuBar::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQMenuBarDragEnterEvent
func callbackQMenuBarDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QMenuBar) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMenuBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QMenuBar) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMenuBar::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QMenuBar) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QMenuBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QMenuBar::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQMenuBarDragLeaveEvent
func callbackQMenuBarDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QMenuBar) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMenuBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QMenuBar) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMenuBar::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QMenuBar) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QMenuBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QMenuBar::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQMenuBarDragMoveEvent
func callbackQMenuBarDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QMenuBar) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMenuBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QMenuBar) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMenuBar::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QMenuBar) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QMenuBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QMenuBar::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQMenuBarDropEvent
func callbackQMenuBarDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QMenuBar) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMenuBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QMenuBar) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMenuBar::dropEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QMenuBar) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMenuBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QMenuBar::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQMenuBarEnterEvent
func callbackQMenuBarEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMenuBar) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMenuBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMenuBar) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMenuBar::enterEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMenuBar) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QMenuBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QMenuBar::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQMenuBarHideEvent
func callbackQMenuBarHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QMenuBar) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QMenuBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QMenuBar) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QMenuBar::hideEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QMenuBar) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QMenuBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QMenuBar::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQMenuBarMoveEvent
func callbackQMenuBarMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QMenuBar) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMenuBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QMenuBar) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMenuBar::moveEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QMenuBar) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QMenuBar::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QMenuBar::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQMenuBarShowEvent
func callbackQMenuBarShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QMenuBar) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QMenuBar::showEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QMenuBar) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QMenuBar::showEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QMenuBar) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QMenuBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QMenuBar::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQMenuBarCloseEvent
func callbackQMenuBarCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QMenuBar) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMenuBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QMenuBar) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMenuBar::closeEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QMenuBar) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QMenuBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QMenuBar::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQMenuBarContextMenuEvent
func callbackQMenuBarContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QMenuBar) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMenuBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QMenuBar) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMenuBar::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QMenuBar) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QMenuBar::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QMenuBar) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QMenuBar::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQMenuBarInitPainter
func callbackQMenuBarInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQMenuBarFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QMenuBar) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMenuBar::initPainter")

	if ptr.Pointer() != nil {
		C.QMenuBar_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QMenuBar) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMenuBar::initPainter")

	if ptr.Pointer() != nil {
		C.QMenuBar_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QMenuBar) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QMenuBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QMenuBar::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQMenuBarInputMethodEvent
func callbackQMenuBarInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QMenuBar) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMenuBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QMenuBar) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMenuBar::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QMenuBar) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMenuBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QMenuBar::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQMenuBarKeyReleaseEvent
func callbackQMenuBarKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QMenuBar) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMenuBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMenuBar) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMenuBar::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMenuBar) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenuBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QMenuBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQMenuBarMouseDoubleClickEvent
func callbackQMenuBarMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMenuBar) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenuBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMenuBar) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenuBar::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMenuBar) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QMenuBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QMenuBar::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQMenuBarTabletEvent
func callbackQMenuBarTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QMenuBar) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMenuBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QMenuBar) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMenuBar::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QMenuBar) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QMenuBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QMenuBar::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQMenuBarWheelEvent
func callbackQMenuBarWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QMenuBar) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMenuBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QMenuBar) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMenuBar::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QMenuBar) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMenuBar::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMenuBar::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMenuBarChildEvent
func callbackQMenuBarChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMenuBar) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMenuBar::childEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMenuBar) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMenuBar::childEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMenuBar) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMenuBar::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMenuBar) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMenuBar::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMenuBarCustomEvent
func callbackQMenuBarCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenuBar::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMenuBarFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMenuBar) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMenuBar::customEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMenuBar) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMenuBar::customEvent")

	if ptr.Pointer() != nil {
		C.QMenuBar_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
