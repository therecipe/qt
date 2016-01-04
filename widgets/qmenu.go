package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMenu struct {
	QWidget
}

type QMenu_ITF interface {
	QWidget_ITF
	QMenu_PTR() *QMenu
}

func PointerFromQMenu(ptr QMenu_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMenu_PTR().Pointer()
	}
	return nil
}

func NewQMenuFromPointer(ptr unsafe.Pointer) *QMenu {
	var n = new(QMenu)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMenu_") {
		n.SetObjectName("QMenu_" + qt.Identifier())
	}
	return n
}

func (ptr *QMenu) QMenu_PTR() *QMenu {
	return ptr
}

func (ptr *QMenu) Icon() *gui.QIcon {
	defer qt.Recovering("QMenu::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QMenu_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) IsTearOffEnabled() bool {
	defer qt.Recovering("QMenu::isTearOffEnabled")

	if ptr.Pointer() != nil {
		return C.QMenu_IsTearOffEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) SeparatorsCollapsible() bool {
	defer qt.Recovering("QMenu::separatorsCollapsible")

	if ptr.Pointer() != nil {
		return C.QMenu_SeparatorsCollapsible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) SetIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QMenu::setIcon")

	if ptr.Pointer() != nil {
		C.QMenu_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QMenu) SetSeparatorsCollapsible(collapse bool) {
	defer qt.Recovering("QMenu::setSeparatorsCollapsible")

	if ptr.Pointer() != nil {
		C.QMenu_SetSeparatorsCollapsible(ptr.Pointer(), C.int(qt.GoBoolToInt(collapse)))
	}
}

func (ptr *QMenu) SetTearOffEnabled(v bool) {
	defer qt.Recovering("QMenu::setTearOffEnabled")

	if ptr.Pointer() != nil {
		C.QMenu_SetTearOffEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QMenu) SetTitle(title string) {
	defer qt.Recovering("QMenu::setTitle")

	if ptr.Pointer() != nil {
		C.QMenu_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QMenu) SetToolTipsVisible(visible bool) {
	defer qt.Recovering("QMenu::setToolTipsVisible")

	if ptr.Pointer() != nil {
		C.QMenu_SetToolTipsVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenu) Title() string {
	defer qt.Recovering("QMenu::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMenu_Title(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMenu) ToolTipsVisible() bool {
	defer qt.Recovering("QMenu::toolTipsVisible")

	if ptr.Pointer() != nil {
		return C.QMenu_ToolTipsVisible(ptr.Pointer()) != 0
	}
	return false
}

func NewQMenu(parent QWidget_ITF) *QMenu {
	defer qt.Recovering("QMenu::QMenu")

	return NewQMenuFromPointer(C.QMenu_NewQMenu(PointerFromQWidget(parent)))
}

func NewQMenu2(title string, parent QWidget_ITF) *QMenu {
	defer qt.Recovering("QMenu::QMenu")

	return NewQMenuFromPointer(C.QMenu_NewQMenu2(C.CString(title), PointerFromQWidget(parent)))
}

func (ptr *QMenu) ConnectAboutToHide(f func()) {
	defer qt.Recovering("connect QMenu::aboutToHide")

	if ptr.Pointer() != nil {
		C.QMenu_ConnectAboutToHide(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToHide", f)
	}
}

func (ptr *QMenu) DisconnectAboutToHide() {
	defer qt.Recovering("disconnect QMenu::aboutToHide")

	if ptr.Pointer() != nil {
		C.QMenu_DisconnectAboutToHide(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToHide")
	}
}

//export callbackQMenuAboutToHide
func callbackQMenuAboutToHide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMenu::aboutToHide")

	if signal := qt.GetSignal(C.GoString(ptrName), "aboutToHide"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMenu) AboutToHide() {
	defer qt.Recovering("QMenu::aboutToHide")

	if ptr.Pointer() != nil {
		C.QMenu_AboutToHide(ptr.Pointer())
	}
}

func (ptr *QMenu) ConnectAboutToShow(f func()) {
	defer qt.Recovering("connect QMenu::aboutToShow")

	if ptr.Pointer() != nil {
		C.QMenu_ConnectAboutToShow(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToShow", f)
	}
}

func (ptr *QMenu) DisconnectAboutToShow() {
	defer qt.Recovering("disconnect QMenu::aboutToShow")

	if ptr.Pointer() != nil {
		C.QMenu_DisconnectAboutToShow(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToShow")
	}
}

//export callbackQMenuAboutToShow
func callbackQMenuAboutToShow(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QMenu::aboutToShow")

	if signal := qt.GetSignal(C.GoString(ptrName), "aboutToShow"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QMenu) AboutToShow() {
	defer qt.Recovering("QMenu::aboutToShow")

	if ptr.Pointer() != nil {
		C.QMenu_AboutToShow(ptr.Pointer())
	}
}

func (ptr *QMenu) ActionAt(pt core.QPoint_ITF) *QAction {
	defer qt.Recovering("QMenu::actionAt")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_ActionAt(ptr.Pointer(), core.PointerFromQPoint(pt)))
	}
	return nil
}

func (ptr *QMenu) ConnectActionEvent(f func(e *gui.QActionEvent)) {
	defer qt.Recovering("connect QMenu::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMenu) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMenu::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMenuActionEvent
func callbackQMenuActionEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(e))
	} else {
		NewQMenuFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(e))
	}
}

func (ptr *QMenu) ActionEvent(e gui.QActionEvent_ITF) {
	defer qt.Recovering("QMenu::actionEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(e))
	}
}

func (ptr *QMenu) ActionEventDefault(e gui.QActionEvent_ITF) {
	defer qt.Recovering("QMenu::actionEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(e))
	}
}

func (ptr *QMenu) ActionGeometry(act QAction_ITF) *core.QRect {
	defer qt.Recovering("QMenu::actionGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QMenu_ActionGeometry(ptr.Pointer(), PointerFromQAction(act)))
	}
	return nil
}

func (ptr *QMenu) ActiveAction() *QAction {
	defer qt.Recovering("QMenu::activeAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_ActiveAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) AddAction2(icon gui.QIcon_ITF, text string) *QAction {
	defer qt.Recovering("QMenu::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddAction4(icon gui.QIcon_ITF, text string, receiver core.QObject_ITF, member string, shortcut gui.QKeySequence_ITF) *QAction {
	defer qt.Recovering("QMenu::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction4(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQObject(receiver), C.CString(member), gui.PointerFromQKeySequence(shortcut)))
	}
	return nil
}

func (ptr *QMenu) AddAction(text string) *QAction {
	defer qt.Recovering("QMenu::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddAction3(text string, receiver core.QObject_ITF, member string, shortcut gui.QKeySequence_ITF) *QAction {
	defer qt.Recovering("QMenu::addAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddAction3(ptr.Pointer(), C.CString(text), core.PointerFromQObject(receiver), C.CString(member), gui.PointerFromQKeySequence(shortcut)))
	}
	return nil
}

func (ptr *QMenu) AddMenu(menu QMenu_ITF) *QAction {
	defer qt.Recovering("QMenu::addMenu")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddMenu(ptr.Pointer(), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenu) AddMenu3(icon gui.QIcon_ITF, title string) *QMenu {
	defer qt.Recovering("QMenu::addMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenu_AddMenu3(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(title)))
	}
	return nil
}

func (ptr *QMenu) AddMenu2(title string) *QMenu {
	defer qt.Recovering("QMenu::addMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMenu_AddMenu2(ptr.Pointer(), C.CString(title)))
	}
	return nil
}

func (ptr *QMenu) AddSection2(icon gui.QIcon_ITF, text string) *QAction {
	defer qt.Recovering("QMenu::addSection")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddSection2(ptr.Pointer(), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddSection(text string) *QAction {
	defer qt.Recovering("QMenu::addSection")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddSection(ptr.Pointer(), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) AddSeparator() *QAction {
	defer qt.Recovering("QMenu::addSeparator")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_AddSeparator(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QMenu::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QMenu) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QMenu::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQMenuChangeEvent
func callbackQMenuChangeEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
	} else {
		NewQMenuFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(e))
	}
}

func (ptr *QMenu) ChangeEvent(e core.QEvent_ITF) {
	defer qt.Recovering("QMenu::changeEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QMenu) ChangeEventDefault(e core.QEvent_ITF) {
	defer qt.Recovering("QMenu::changeEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(e))
	}
}

func (ptr *QMenu) Clear() {
	defer qt.Recovering("QMenu::clear")

	if ptr.Pointer() != nil {
		C.QMenu_Clear(ptr.Pointer())
	}
}

func (ptr *QMenu) ConnectEnterEvent(f func(v *core.QEvent)) {
	defer qt.Recovering("connect QMenu::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QMenu) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QMenu::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQMenuEnterEvent
func callbackQMenuEnterEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(v))
	} else {
		NewQMenuFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(v))
	}
}

func (ptr *QMenu) EnterEvent(v core.QEvent_ITF) {
	defer qt.Recovering("QMenu::enterEvent")

	if ptr.Pointer() != nil {
		C.QMenu_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(v))
	}
}

func (ptr *QMenu) EnterEventDefault(v core.QEvent_ITF) {
	defer qt.Recovering("QMenu::enterEvent")

	if ptr.Pointer() != nil {
		C.QMenu_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(v))
	}
}

func (ptr *QMenu) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QMenu::event")

	if ptr.Pointer() != nil {
		return C.QMenu_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QMenu) Exec() *QAction {
	defer qt.Recovering("QMenu::exec")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_Exec(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) Exec2(p core.QPoint_ITF, action QAction_ITF) *QAction {
	defer qt.Recovering("QMenu::exec")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_Exec2(ptr.Pointer(), core.PointerFromQPoint(p), PointerFromQAction(action)))
	}
	return nil
}

func (ptr *QMenu) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QMenu::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QMenu_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QMenu) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QMenu::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QMenu) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QMenu::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQMenuHideEvent
func callbackQMenuHideEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
	} else {
		NewQMenuFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(v))
	}
}

func (ptr *QMenu) HideEvent(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QMenu::hideEvent")

	if ptr.Pointer() != nil {
		C.QMenu_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QMenu) HideEventDefault(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QMenu::hideEvent")

	if ptr.Pointer() != nil {
		C.QMenu_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QMenu) HideTearOffMenu() {
	defer qt.Recovering("QMenu::hideTearOffMenu")

	if ptr.Pointer() != nil {
		C.QMenu_HideTearOffMenu(ptr.Pointer())
	}
}

func (ptr *QMenu) ConnectHovered(f func(action *QAction)) {
	defer qt.Recovering("connect QMenu::hovered")

	if ptr.Pointer() != nil {
		C.QMenu_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QMenu) DisconnectHovered() {
	defer qt.Recovering("disconnect QMenu::hovered")

	if ptr.Pointer() != nil {
		C.QMenu_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQMenuHovered
func callbackQMenuHovered(ptr unsafe.Pointer, ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::hovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "hovered"); signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QMenu) Hovered(action QAction_ITF) {
	defer qt.Recovering("QMenu::hovered")

	if ptr.Pointer() != nil {
		C.QMenu_Hovered(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QMenu) InsertMenu(before QAction_ITF, menu QMenu_ITF) *QAction {
	defer qt.Recovering("QMenu::insertMenu")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertMenu(ptr.Pointer(), PointerFromQAction(before), PointerFromQMenu(menu)))
	}
	return nil
}

func (ptr *QMenu) InsertSection2(before QAction_ITF, icon gui.QIcon_ITF, text string) *QAction {
	defer qt.Recovering("QMenu::insertSection")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertSection2(ptr.Pointer(), PointerFromQAction(before), gui.PointerFromQIcon(icon), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) InsertSection(before QAction_ITF, text string) *QAction {
	defer qt.Recovering("QMenu::insertSection")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertSection(ptr.Pointer(), PointerFromQAction(before), C.CString(text)))
	}
	return nil
}

func (ptr *QMenu) InsertSeparator(before QAction_ITF) *QAction {
	defer qt.Recovering("QMenu::insertSeparator")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_InsertSeparator(ptr.Pointer(), PointerFromQAction(before)))
	}
	return nil
}

func (ptr *QMenu) IsEmpty() bool {
	defer qt.Recovering("QMenu::isEmpty")

	if ptr.Pointer() != nil {
		return C.QMenu_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) IsTearOffMenuVisible() bool {
	defer qt.Recovering("QMenu::isTearOffMenuVisible")

	if ptr.Pointer() != nil {
		return C.QMenu_IsTearOffMenuVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMenu) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMenu::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMenu) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMenu::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMenuKeyPressEvent
func callbackQMenuKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQMenuFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QMenu) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMenu::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMenu_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QMenu) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMenu::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QMenu_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QMenu) ConnectLeaveEvent(f func(v *core.QEvent)) {
	defer qt.Recovering("connect QMenu::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMenu) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMenu::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMenuLeaveEvent
func callbackQMenuLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(v))
	} else {
		NewQMenuFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(v))
	}
}

func (ptr *QMenu) LeaveEvent(v core.QEvent_ITF) {
	defer qt.Recovering("QMenu::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(v))
	}
}

func (ptr *QMenu) LeaveEventDefault(v core.QEvent_ITF) {
	defer qt.Recovering("QMenu::leaveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(v))
	}
}

func (ptr *QMenu) MenuAction() *QAction {
	defer qt.Recovering("QMenu::menuAction")

	if ptr.Pointer() != nil {
		return NewQActionFromPointer(C.QMenu_MenuAction(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenu::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMenu) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMenu::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMenuMouseMoveEvent
func callbackQMenuMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMenuFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMenu) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenu::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenu) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenu::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenu) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenu::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMenu) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMenu::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMenuMousePressEvent
func callbackQMenuMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMenuFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMenu) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenu::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenu) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenu::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenu) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenu::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMenu) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMenu::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMenuMouseReleaseEvent
func callbackQMenuMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQMenuFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QMenu) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenu::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenu) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenu::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QMenu) ConnectPaintEvent(f func(e *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMenu::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMenu) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMenu::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMenuPaintEvent
func callbackQMenuPaintEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
	} else {
		NewQMenuFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(e))
	}
}

func (ptr *QMenu) PaintEvent(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMenu::paintEvent")

	if ptr.Pointer() != nil {
		C.QMenu_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QMenu) PaintEventDefault(e gui.QPaintEvent_ITF) {
	defer qt.Recovering("QMenu::paintEvent")

	if ptr.Pointer() != nil {
		C.QMenu_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(e))
	}
}

func (ptr *QMenu) Popup(p core.QPoint_ITF, atAction QAction_ITF) {
	defer qt.Recovering("QMenu::popup")

	if ptr.Pointer() != nil {
		C.QMenu_Popup(ptr.Pointer(), core.PointerFromQPoint(p), PointerFromQAction(atAction))
	}
}

func (ptr *QMenu) SetActiveAction(act QAction_ITF) {
	defer qt.Recovering("QMenu::setActiveAction")

	if ptr.Pointer() != nil {
		C.QMenu_SetActiveAction(ptr.Pointer(), PointerFromQAction(act))
	}
}

func (ptr *QMenu) SizeHint() *core.QSize {
	defer qt.Recovering("QMenu::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMenu_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMenu) ConnectTimerEvent(f func(e *core.QTimerEvent)) {
	defer qt.Recovering("connect QMenu::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMenu) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMenu::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMenuTimerEvent
func callbackQMenuTimerEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
	} else {
		NewQMenuFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(e))
	}
}

func (ptr *QMenu) TimerEvent(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QMenu::timerEvent")

	if ptr.Pointer() != nil {
		C.QMenu_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QMenu) TimerEventDefault(e core.QTimerEvent_ITF) {
	defer qt.Recovering("QMenu::timerEvent")

	if ptr.Pointer() != nil {
		C.QMenu_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(e))
	}
}

func (ptr *QMenu) ConnectTriggered(f func(action *QAction)) {
	defer qt.Recovering("connect QMenu::triggered")

	if ptr.Pointer() != nil {
		C.QMenu_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QMenu) DisconnectTriggered() {
	defer qt.Recovering("disconnect QMenu::triggered")

	if ptr.Pointer() != nil {
		C.QMenu_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQMenuTriggered
func callbackQMenuTriggered(ptr unsafe.Pointer, ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::triggered")

	if signal := qt.GetSignal(C.GoString(ptrName), "triggered"); signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
	}

}

func (ptr *QMenu) Triggered(action QAction_ITF) {
	defer qt.Recovering("QMenu::triggered")

	if ptr.Pointer() != nil {
		C.QMenu_Triggered(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QMenu) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QMenu::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QMenu) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QMenu::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQMenuWheelEvent
func callbackQMenuWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQMenuFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QMenu) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMenu::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMenu_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QMenu) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QMenu::wheelEvent")

	if ptr.Pointer() != nil {
		C.QMenu_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QMenu) DestroyQMenu() {
	defer qt.Recovering("QMenu::~QMenu")

	if ptr.Pointer() != nil {
		C.QMenu_DestroyQMenu(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMenu) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QMenu::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QMenu) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QMenu::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQMenuDragEnterEvent
func callbackQMenuDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QMenu) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMenu::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMenu_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QMenu) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QMenu::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QMenu_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QMenu) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QMenu::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QMenu) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QMenu::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQMenuDragLeaveEvent
func callbackQMenuDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QMenu) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMenu::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QMenu) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QMenu::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QMenu) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QMenu::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QMenu) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QMenu::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQMenuDragMoveEvent
func callbackQMenuDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QMenu) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMenu::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QMenu) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QMenu::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QMenu) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QMenu::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QMenu) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QMenu::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQMenuDropEvent
func callbackQMenuDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QMenu) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMenu::dropEvent")

	if ptr.Pointer() != nil {
		C.QMenu_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QMenu) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QMenu::dropEvent")

	if ptr.Pointer() != nil {
		C.QMenu_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QMenu) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMenu::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QMenu) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QMenu::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQMenuFocusInEvent
func callbackQMenuFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QMenu) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMenu::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMenu_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMenu) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMenu::focusInEvent")

	if ptr.Pointer() != nil {
		C.QMenu_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMenu) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMenu::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMenu) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMenu::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMenuFocusOutEvent
func callbackQMenuFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QMenu) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMenu::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMenu_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMenu) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QMenu::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QMenu_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QMenu) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QMenu::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QMenu) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QMenu::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQMenuMoveEvent
func callbackQMenuMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QMenu) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMenu::moveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QMenu) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QMenu::moveEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QMenu) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMenu::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMenu) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMenu::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMenuSetVisible
func callbackQMenuSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMenu::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMenu) SetVisible(visible bool) {
	defer qt.Recovering("QMenu::setVisible")

	if ptr.Pointer() != nil {
		C.QMenu_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenu) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QMenu::setVisible")

	if ptr.Pointer() != nil {
		C.QMenu_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QMenu) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QMenu::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QMenu) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QMenu::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQMenuShowEvent
func callbackQMenuShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QMenu) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QMenu::showEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QMenu) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QMenu::showEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QMenu) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QMenu::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QMenu) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QMenu::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQMenuCloseEvent
func callbackQMenuCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QMenu) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMenu::closeEvent")

	if ptr.Pointer() != nil {
		C.QMenu_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QMenu) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QMenu::closeEvent")

	if ptr.Pointer() != nil {
		C.QMenu_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QMenu) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QMenu::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QMenu) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QMenu::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQMenuContextMenuEvent
func callbackQMenuContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QMenu) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMenu::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QMenu) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QMenu::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QMenu) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QMenu::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QMenu) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QMenu::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQMenuInitPainter
func callbackQMenuInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQMenuFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QMenu) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMenu::initPainter")

	if ptr.Pointer() != nil {
		C.QMenu_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QMenu) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QMenu::initPainter")

	if ptr.Pointer() != nil {
		C.QMenu_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QMenu) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QMenu::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QMenu) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QMenu::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQMenuInputMethodEvent
func callbackQMenuInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QMenu) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMenu::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMenu_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QMenu) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QMenu::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QMenu_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QMenu) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMenu::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QMenu) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QMenu::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQMenuKeyReleaseEvent
func callbackQMenuKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QMenu) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMenu::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMenu_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMenu) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QMenu::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QMenu_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QMenu) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMenu::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QMenu) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QMenu::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQMenuMouseDoubleClickEvent
func callbackQMenuMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QMenu) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenu::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMenu) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QMenu::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QMenu_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QMenu) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMenu::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMenu) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMenu::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMenuResizeEvent
func callbackQMenuResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QMenu) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMenu::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QMenu) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QMenu::resizeEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QMenu) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QMenu::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QMenu) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QMenu::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQMenuTabletEvent
func callbackQMenuTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QMenu) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMenu::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMenu_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QMenu) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QMenu::tabletEvent")

	if ptr.Pointer() != nil {
		C.QMenu_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QMenu) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMenu::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMenu) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMenu::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMenuChildEvent
func callbackQMenuChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMenu) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMenu::childEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMenu) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMenu::childEvent")

	if ptr.Pointer() != nil {
		C.QMenu_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMenu) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMenu::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMenu) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMenu::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMenuCustomEvent
func callbackQMenuCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMenuFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMenu) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMenu::customEvent")

	if ptr.Pointer() != nil {
		C.QMenu_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMenu) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMenu::customEvent")

	if ptr.Pointer() != nil {
		C.QMenu_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
