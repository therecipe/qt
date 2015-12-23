package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QMainWindow struct {
	QWidget
}

type QMainWindow_ITF interface {
	QWidget_ITF
	QMainWindow_PTR() *QMainWindow
}

func PointerFromQMainWindow(ptr QMainWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindow_PTR().Pointer()
	}
	return nil
}

func NewQMainWindowFromPointer(ptr unsafe.Pointer) *QMainWindow {
	var n = new(QMainWindow)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMainWindow_") {
		n.SetObjectName("QMainWindow_" + qt.Identifier())
	}
	return n
}

func (ptr *QMainWindow) QMainWindow_PTR() *QMainWindow {
	return ptr
}

//QMainWindow::DockOption
type QMainWindow__DockOption int64

const (
	QMainWindow__AnimatedDocks    = QMainWindow__DockOption(0x01)
	QMainWindow__AllowNestedDocks = QMainWindow__DockOption(0x02)
	QMainWindow__AllowTabbedDocks = QMainWindow__DockOption(0x04)
	QMainWindow__ForceTabbedDocks = QMainWindow__DockOption(0x08)
	QMainWindow__VerticalTabs     = QMainWindow__DockOption(0x10)
)

func (ptr *QMainWindow) DockOptions() QMainWindow__DockOption {
	defer qt.Recovering("QMainWindow::dockOptions")

	if ptr.Pointer() != nil {
		return QMainWindow__DockOption(C.QMainWindow_DockOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMainWindow) DocumentMode() bool {
	defer qt.Recovering("QMainWindow::documentMode")

	if ptr.Pointer() != nil {
		return C.QMainWindow_DocumentMode(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMainWindow) IconSize() *core.QSize {
	defer qt.Recovering("QMainWindow::iconSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QMainWindow_IconSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) IsAnimated() bool {
	defer qt.Recovering("QMainWindow::isAnimated")

	if ptr.Pointer() != nil {
		return C.QMainWindow_IsAnimated(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMainWindow) IsDockNestingEnabled() bool {
	defer qt.Recovering("QMainWindow::isDockNestingEnabled")

	if ptr.Pointer() != nil {
		return C.QMainWindow_IsDockNestingEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMainWindow) SetAnimated(enabled bool) {
	defer qt.Recovering("QMainWindow::setAnimated")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetAnimated(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMainWindow) SetDockNestingEnabled(enabled bool) {
	defer qt.Recovering("QMainWindow::setDockNestingEnabled")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetDockNestingEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMainWindow) SetDockOptions(options QMainWindow__DockOption) {
	defer qt.Recovering("QMainWindow::setDockOptions")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetDockOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QMainWindow) SetDocumentMode(enabled bool) {
	defer qt.Recovering("QMainWindow::setDocumentMode")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetDocumentMode(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMainWindow) SetIconSize(iconSize core.QSize_ITF) {
	defer qt.Recovering("QMainWindow::setIconSize")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetIconSize(ptr.Pointer(), core.PointerFromQSize(iconSize))
	}
}

func (ptr *QMainWindow) SetTabShape(tabShape QTabWidget__TabShape) {
	defer qt.Recovering("QMainWindow::setTabShape")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetTabShape(ptr.Pointer(), C.int(tabShape))
	}
}

func (ptr *QMainWindow) SetToolButtonStyle(toolButtonStyle core.Qt__ToolButtonStyle) {
	defer qt.Recovering("QMainWindow::setToolButtonStyle")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetToolButtonStyle(ptr.Pointer(), C.int(toolButtonStyle))
	}
}

func (ptr *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	defer qt.Recovering("QMainWindow::setUnifiedTitleAndToolBarOnMac")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetUnifiedTitleAndToolBarOnMac(ptr.Pointer(), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QMainWindow) SplitDockWidget(first QDockWidget_ITF, second QDockWidget_ITF, orientation core.Qt__Orientation) {
	defer qt.Recovering("QMainWindow::splitDockWidget")

	if ptr.Pointer() != nil {
		C.QMainWindow_SplitDockWidget(ptr.Pointer(), PointerFromQDockWidget(first), PointerFromQDockWidget(second), C.int(orientation))
	}
}

func (ptr *QMainWindow) TabShape() QTabWidget__TabShape {
	defer qt.Recovering("QMainWindow::tabShape")

	if ptr.Pointer() != nil {
		return QTabWidget__TabShape(C.QMainWindow_TabShape(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMainWindow) TabifyDockWidget(first QDockWidget_ITF, second QDockWidget_ITF) {
	defer qt.Recovering("QMainWindow::tabifyDockWidget")

	if ptr.Pointer() != nil {
		C.QMainWindow_TabifyDockWidget(ptr.Pointer(), PointerFromQDockWidget(first), PointerFromQDockWidget(second))
	}
}

func (ptr *QMainWindow) ToolButtonStyle() core.Qt__ToolButtonStyle {
	defer qt.Recovering("QMainWindow::toolButtonStyle")

	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QMainWindow_ToolButtonStyle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMainWindow) UnifiedTitleAndToolBarOnMac() bool {
	defer qt.Recovering("QMainWindow::unifiedTitleAndToolBarOnMac")

	if ptr.Pointer() != nil {
		return C.QMainWindow_UnifiedTitleAndToolBarOnMac(ptr.Pointer()) != 0
	}
	return false
}

func NewQMainWindow(parent QWidget_ITF, flags core.Qt__WindowType) *QMainWindow {
	defer qt.Recovering("QMainWindow::QMainWindow")

	return NewQMainWindowFromPointer(C.QMainWindow_NewQMainWindow(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QMainWindow) AddDockWidget(area core.Qt__DockWidgetArea, dockwidget QDockWidget_ITF) {
	defer qt.Recovering("QMainWindow::addDockWidget")

	if ptr.Pointer() != nil {
		C.QMainWindow_AddDockWidget(ptr.Pointer(), C.int(area), PointerFromQDockWidget(dockwidget))
	}
}

func (ptr *QMainWindow) AddDockWidget2(area core.Qt__DockWidgetArea, dockwidget QDockWidget_ITF, orientation core.Qt__Orientation) {
	defer qt.Recovering("QMainWindow::addDockWidget")

	if ptr.Pointer() != nil {
		C.QMainWindow_AddDockWidget2(ptr.Pointer(), C.int(area), PointerFromQDockWidget(dockwidget), C.int(orientation))
	}
}

func (ptr *QMainWindow) AddToolBar3(title string) *QToolBar {
	defer qt.Recovering("QMainWindow::addToolBar")

	if ptr.Pointer() != nil {
		return NewQToolBarFromPointer(C.QMainWindow_AddToolBar3(ptr.Pointer(), C.CString(title)))
	}
	return nil
}

func (ptr *QMainWindow) AddToolBar2(toolbar QToolBar_ITF) {
	defer qt.Recovering("QMainWindow::addToolBar")

	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBar2(ptr.Pointer(), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) AddToolBar(area core.Qt__ToolBarArea, toolbar QToolBar_ITF) {
	defer qt.Recovering("QMainWindow::addToolBar")

	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBar(ptr.Pointer(), C.int(area), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) AddToolBarBreak(area core.Qt__ToolBarArea) {
	defer qt.Recovering("QMainWindow::addToolBarBreak")

	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBarBreak(ptr.Pointer(), C.int(area))
	}
}

func (ptr *QMainWindow) CentralWidget() *QWidget {
	defer qt.Recovering("QMainWindow::centralWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMainWindow_CentralWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QMainWindow::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QMainWindow::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQMainWindowContextMenuEvent
func callbackQMainWindowContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) Corner(corner core.Qt__Corner) core.Qt__DockWidgetArea {
	defer qt.Recovering("QMainWindow::corner")

	if ptr.Pointer() != nil {
		return core.Qt__DockWidgetArea(C.QMainWindow_Corner(ptr.Pointer(), C.int(corner)))
	}
	return 0
}

func (ptr *QMainWindow) CreatePopupMenu() *QMenu {
	defer qt.Recovering("QMainWindow::createPopupMenu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QMainWindow_CreatePopupMenu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) DockWidgetArea(dockwidget QDockWidget_ITF) core.Qt__DockWidgetArea {
	defer qt.Recovering("QMainWindow::dockWidgetArea")

	if ptr.Pointer() != nil {
		return core.Qt__DockWidgetArea(C.QMainWindow_DockWidgetArea(ptr.Pointer(), PointerFromQDockWidget(dockwidget)))
	}
	return 0
}

func (ptr *QMainWindow) ConnectIconSizeChanged(f func(iconSize *core.QSize)) {
	defer qt.Recovering("connect QMainWindow::iconSizeChanged")

	if ptr.Pointer() != nil {
		C.QMainWindow_ConnectIconSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "iconSizeChanged", f)
	}
}

func (ptr *QMainWindow) DisconnectIconSizeChanged() {
	defer qt.Recovering("disconnect QMainWindow::iconSizeChanged")

	if ptr.Pointer() != nil {
		C.QMainWindow_DisconnectIconSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "iconSizeChanged")
	}
}

//export callbackQMainWindowIconSizeChanged
func callbackQMainWindowIconSizeChanged(ptrName *C.char, iconSize unsafe.Pointer) {
	defer qt.Recovering("callback QMainWindow::iconSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "iconSizeChanged"); signal != nil {
		signal.(func(*core.QSize))(core.NewQSizeFromPointer(iconSize))
	}

}

func (ptr *QMainWindow) InsertToolBar(before QToolBar_ITF, toolbar QToolBar_ITF) {
	defer qt.Recovering("QMainWindow::insertToolBar")

	if ptr.Pointer() != nil {
		C.QMainWindow_InsertToolBar(ptr.Pointer(), PointerFromQToolBar(before), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) InsertToolBarBreak(before QToolBar_ITF) {
	defer qt.Recovering("QMainWindow::insertToolBarBreak")

	if ptr.Pointer() != nil {
		C.QMainWindow_InsertToolBarBreak(ptr.Pointer(), PointerFromQToolBar(before))
	}
}

func (ptr *QMainWindow) MenuBar() *QMenuBar {
	defer qt.Recovering("QMainWindow::menuBar")

	if ptr.Pointer() != nil {
		return NewQMenuBarFromPointer(C.QMainWindow_MenuBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) MenuWidget() *QWidget {
	defer qt.Recovering("QMainWindow::menuWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMainWindow_MenuWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) RemoveDockWidget(dockwidget QDockWidget_ITF) {
	defer qt.Recovering("QMainWindow::removeDockWidget")

	if ptr.Pointer() != nil {
		C.QMainWindow_RemoveDockWidget(ptr.Pointer(), PointerFromQDockWidget(dockwidget))
	}
}

func (ptr *QMainWindow) RemoveToolBar(toolbar QToolBar_ITF) {
	defer qt.Recovering("QMainWindow::removeToolBar")

	if ptr.Pointer() != nil {
		C.QMainWindow_RemoveToolBar(ptr.Pointer(), PointerFromQToolBar(toolbar))
	}
}

func (ptr *QMainWindow) RemoveToolBarBreak(before QToolBar_ITF) {
	defer qt.Recovering("QMainWindow::removeToolBarBreak")

	if ptr.Pointer() != nil {
		C.QMainWindow_RemoveToolBarBreak(ptr.Pointer(), PointerFromQToolBar(before))
	}
}

func (ptr *QMainWindow) RestoreDockWidget(dockwidget QDockWidget_ITF) bool {
	defer qt.Recovering("QMainWindow::restoreDockWidget")

	if ptr.Pointer() != nil {
		return C.QMainWindow_RestoreDockWidget(ptr.Pointer(), PointerFromQDockWidget(dockwidget)) != 0
	}
	return false
}

func (ptr *QMainWindow) RestoreState(state core.QByteArray_ITF, version int) bool {
	defer qt.Recovering("QMainWindow::restoreState")

	if ptr.Pointer() != nil {
		return C.QMainWindow_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state), C.int(version)) != 0
	}
	return false
}

func (ptr *QMainWindow) SaveState(version int) *core.QByteArray {
	defer qt.Recovering("QMainWindow::saveState")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QMainWindow_SaveState(ptr.Pointer(), C.int(version)))
	}
	return nil
}

func (ptr *QMainWindow) SetCentralWidget(widget QWidget_ITF) {
	defer qt.Recovering("QMainWindow::setCentralWidget")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetCentralWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QMainWindow) SetCorner(corner core.Qt__Corner, area core.Qt__DockWidgetArea) {
	defer qt.Recovering("QMainWindow::setCorner")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetCorner(ptr.Pointer(), C.int(corner), C.int(area))
	}
}

func (ptr *QMainWindow) SetMenuBar(menuBar QMenuBar_ITF) {
	defer qt.Recovering("QMainWindow::setMenuBar")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetMenuBar(ptr.Pointer(), PointerFromQMenuBar(menuBar))
	}
}

func (ptr *QMainWindow) SetMenuWidget(menuBar QWidget_ITF) {
	defer qt.Recovering("QMainWindow::setMenuWidget")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetMenuWidget(ptr.Pointer(), PointerFromQWidget(menuBar))
	}
}

func (ptr *QMainWindow) SetStatusBar(statusbar QStatusBar_ITF) {
	defer qt.Recovering("QMainWindow::setStatusBar")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetStatusBar(ptr.Pointer(), PointerFromQStatusBar(statusbar))
	}
}

func (ptr *QMainWindow) SetTabPosition(areas core.Qt__DockWidgetArea, tabPosition QTabWidget__TabPosition) {
	defer qt.Recovering("QMainWindow::setTabPosition")

	if ptr.Pointer() != nil {
		C.QMainWindow_SetTabPosition(ptr.Pointer(), C.int(areas), C.int(tabPosition))
	}
}

func (ptr *QMainWindow) StatusBar() *QStatusBar {
	defer qt.Recovering("QMainWindow::statusBar")

	if ptr.Pointer() != nil {
		return NewQStatusBarFromPointer(C.QMainWindow_StatusBar(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) TabPosition(area core.Qt__DockWidgetArea) QTabWidget__TabPosition {
	defer qt.Recovering("QMainWindow::tabPosition")

	if ptr.Pointer() != nil {
		return QTabWidget__TabPosition(C.QMainWindow_TabPosition(ptr.Pointer(), C.int(area)))
	}
	return 0
}

func (ptr *QMainWindow) TakeCentralWidget() *QWidget {
	defer qt.Recovering("QMainWindow::takeCentralWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QMainWindow_TakeCentralWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMainWindow) ToolBarArea(toolbar QToolBar_ITF) core.Qt__ToolBarArea {
	defer qt.Recovering("QMainWindow::toolBarArea")

	if ptr.Pointer() != nil {
		return core.Qt__ToolBarArea(C.QMainWindow_ToolBarArea(ptr.Pointer(), PointerFromQToolBar(toolbar)))
	}
	return 0
}

func (ptr *QMainWindow) ToolBarBreak(toolbar QToolBar_ITF) bool {
	defer qt.Recovering("QMainWindow::toolBarBreak")

	if ptr.Pointer() != nil {
		return C.QMainWindow_ToolBarBreak(ptr.Pointer(), PointerFromQToolBar(toolbar)) != 0
	}
	return false
}

func (ptr *QMainWindow) ConnectToolButtonStyleChanged(f func(toolButtonStyle core.Qt__ToolButtonStyle)) {
	defer qt.Recovering("connect QMainWindow::toolButtonStyleChanged")

	if ptr.Pointer() != nil {
		C.QMainWindow_ConnectToolButtonStyleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toolButtonStyleChanged", f)
	}
}

func (ptr *QMainWindow) DisconnectToolButtonStyleChanged() {
	defer qt.Recovering("disconnect QMainWindow::toolButtonStyleChanged")

	if ptr.Pointer() != nil {
		C.QMainWindow_DisconnectToolButtonStyleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toolButtonStyleChanged")
	}
}

//export callbackQMainWindowToolButtonStyleChanged
func callbackQMainWindowToolButtonStyleChanged(ptrName *C.char, toolButtonStyle C.int) {
	defer qt.Recovering("callback QMainWindow::toolButtonStyleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "toolButtonStyleChanged"); signal != nil {
		signal.(func(core.Qt__ToolButtonStyle))(core.Qt__ToolButtonStyle(toolButtonStyle))
	}

}

func (ptr *QMainWindow) DestroyQMainWindow() {
	defer qt.Recovering("QMainWindow::~QMainWindow")

	if ptr.Pointer() != nil {
		C.QMainWindow_DestroyQMainWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMainWindow) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QMainWindow::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QMainWindow::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQMainWindowActionEvent
func callbackQMainWindowActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QMainWindow::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QMainWindow::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQMainWindowDragEnterEvent
func callbackQMainWindowDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QMainWindow::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QMainWindow::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQMainWindowDragLeaveEvent
func callbackQMainWindowDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QMainWindow::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QMainWindow::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQMainWindowDragMoveEvent
func callbackQMainWindowDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QMainWindow::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QMainWindow::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQMainWindowDropEvent
func callbackQMainWindowDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMainWindow::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QMainWindow::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQMainWindowEnterEvent
func callbackQMainWindowEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QMainWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QMainWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQMainWindowFocusOutEvent
func callbackQMainWindowFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QMainWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QMainWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQMainWindowHideEvent
func callbackQMainWindowHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMainWindow::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QMainWindow::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQMainWindowLeaveEvent
func callbackQMainWindowLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QMainWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QMainWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQMainWindowMoveEvent
func callbackQMainWindowMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QMainWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QMainWindow::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQMainWindowPaintEvent
func callbackQMainWindowPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QMainWindow::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QMainWindow) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QMainWindow::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQMainWindowSetVisible
func callbackQMainWindowSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QMainWindow::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QMainWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QMainWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQMainWindowShowEvent
func callbackQMainWindowShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QMainWindow::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QMainWindow::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQMainWindowCloseEvent
func callbackQMainWindowCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QMainWindow::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QMainWindow) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QMainWindow::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQMainWindowInitPainter
func callbackQMainWindowInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QMainWindow::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QMainWindow::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQMainWindowInputMethodEvent
func callbackQMainWindowInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMainWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QMainWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQMainWindowKeyPressEvent
func callbackQMainWindowKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QMainWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QMainWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQMainWindowKeyReleaseEvent
func callbackQMainWindowKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMainWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QMainWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQMainWindowMouseDoubleClickEvent
func callbackQMainWindowMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMainWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QMainWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQMainWindowMouseMoveEvent
func callbackQMainWindowMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMainWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QMainWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQMainWindowMousePressEvent
func callbackQMainWindowMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QMainWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QMainWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQMainWindowMouseReleaseEvent
func callbackQMainWindowMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QMainWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QMainWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQMainWindowResizeEvent
func callbackQMainWindowResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QMainWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QMainWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQMainWindowTabletEvent
func callbackQMainWindowTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QMainWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QMainWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQMainWindowWheelEvent
func callbackQMainWindowWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMainWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMainWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMainWindowTimerEvent
func callbackQMainWindowTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMainWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMainWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMainWindowChildEvent
func callbackQMainWindowChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QMainWindow) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMainWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMainWindow) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMainWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMainWindowCustomEvent
func callbackQMainWindowCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QMainWindow::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
