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
func callbackQMenuAboutToHide(ptrName *C.char) {
	defer qt.Recovering("callback QMenu::aboutToHide")

	if signal := qt.GetSignal(C.GoString(ptrName), "aboutToHide"); signal != nil {
		signal.(func())()
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
func callbackQMenuAboutToShow(ptrName *C.char) {
	defer qt.Recovering("callback QMenu::aboutToShow")

	if signal := qt.GetSignal(C.GoString(ptrName), "aboutToShow"); signal != nil {
		signal.(func())()
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
func callbackQMenuActionEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(e))
		return true
	}
	return false

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
func callbackQMenuChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

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
func callbackQMenuEnterEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(v))
		return true
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
func callbackQMenuHideEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
		return true
	}
	return false

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
func callbackQMenuHovered(ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::hovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "hovered"); signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
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
func callbackQMenuKeyPressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
		return true
	}
	return false

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
func callbackQMenuLeaveEvent(ptrName *C.char, v unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(v))
		return true
	}
	return false

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
func callbackQMenuMouseMoveEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQMenuMousePressEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQMenuMouseReleaseEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
		return true
	}
	return false

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
func callbackQMenuPaintEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(e))
		return true
	}
	return false

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
func callbackQMenuTimerEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(e))
		return true
	}
	return false

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
func callbackQMenuTriggered(ptrName *C.char, action unsafe.Pointer) {
	defer qt.Recovering("callback QMenu::triggered")

	if signal := qt.GetSignal(C.GoString(ptrName), "triggered"); signal != nil {
		signal.(func(*QAction))(NewQActionFromPointer(action))
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
func callbackQMenuWheelEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
		return true
	}
	return false

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
func callbackQMenuDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMenu::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

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
func callbackQMenuShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

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
func callbackQMenuInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQMenuCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMenu::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
