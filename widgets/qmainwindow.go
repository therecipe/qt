package widgets

//#include "qmainwindow.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMainWindow struct {
	QWidget
}

type QMainWindowITF interface {
	QWidgetITF
	QMainWindowPTR() *QMainWindow
}

func PointerFromQMainWindow(ptr QMainWindowITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMainWindowPTR().Pointer()
	}
	return nil
}

func QMainWindowFromPointer(ptr unsafe.Pointer) *QMainWindow {
	var n = new(QMainWindow)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMainWindow_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMainWindow) QMainWindowPTR() *QMainWindow {
	return ptr
}

//QMainWindow::DockOption
type QMainWindow__DockOption int

var (
	QMainWindow__AnimatedDocks    = QMainWindow__DockOption(0x01)
	QMainWindow__AllowNestedDocks = QMainWindow__DockOption(0x02)
	QMainWindow__AllowTabbedDocks = QMainWindow__DockOption(0x04)
	QMainWindow__ForceTabbedDocks = QMainWindow__DockOption(0x08)
	QMainWindow__VerticalTabs     = QMainWindow__DockOption(0x10)
)

func (ptr *QMainWindow) DockOptions() QMainWindow__DockOption {
	if ptr.Pointer() != nil {
		return QMainWindow__DockOption(C.QMainWindow_DockOptions(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMainWindow) DocumentMode() bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_DocumentMode(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMainWindow) IsAnimated() bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_IsAnimated(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMainWindow) IsDockNestingEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_IsDockNestingEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMainWindow) SetAnimated(enabled bool) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetAnimated(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMainWindow) SetDockNestingEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetDockNestingEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMainWindow) SetDockOptions(options QMainWindow__DockOption) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetDockOptions(C.QtObjectPtr(ptr.Pointer()), C.int(options))
	}
}

func (ptr *QMainWindow) SetDocumentMode(enabled bool) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetDocumentMode(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QMainWindow) SetIconSize(iconSize core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetIconSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(iconSize)))
	}
}

func (ptr *QMainWindow) SetTabShape(tabShape QTabWidget__TabShape) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetTabShape(C.QtObjectPtr(ptr.Pointer()), C.int(tabShape))
	}
}

func (ptr *QMainWindow) SetToolButtonStyle(toolButtonStyle core.Qt__ToolButtonStyle) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetToolButtonStyle(C.QtObjectPtr(ptr.Pointer()), C.int(toolButtonStyle))
	}
}

func (ptr *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetUnifiedTitleAndToolBarOnMac(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(set)))
	}
}

func (ptr *QMainWindow) SplitDockWidget(first QDockWidgetITF, second QDockWidgetITF, orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SplitDockWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDockWidget(first)), C.QtObjectPtr(PointerFromQDockWidget(second)), C.int(orientation))
	}
}

func (ptr *QMainWindow) TabShape() QTabWidget__TabShape {
	if ptr.Pointer() != nil {
		return QTabWidget__TabShape(C.QMainWindow_TabShape(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMainWindow) TabifyDockWidget(first QDockWidgetITF, second QDockWidgetITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_TabifyDockWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDockWidget(first)), C.QtObjectPtr(PointerFromQDockWidget(second)))
	}
}

func (ptr *QMainWindow) ToolButtonStyle() core.Qt__ToolButtonStyle {
	if ptr.Pointer() != nil {
		return core.Qt__ToolButtonStyle(C.QMainWindow_ToolButtonStyle(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMainWindow) UnifiedTitleAndToolBarOnMac() bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_UnifiedTitleAndToolBarOnMac(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQMainWindow(parent QWidgetITF, flags core.Qt__WindowType) *QMainWindow {
	return QMainWindowFromPointer(unsafe.Pointer(C.QMainWindow_NewQMainWindow(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(flags))))
}

func (ptr *QMainWindow) AddDockWidget(area core.Qt__DockWidgetArea, dockwidget QDockWidgetITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddDockWidget(C.QtObjectPtr(ptr.Pointer()), C.int(area), C.QtObjectPtr(PointerFromQDockWidget(dockwidget)))
	}
}

func (ptr *QMainWindow) AddDockWidget2(area core.Qt__DockWidgetArea, dockwidget QDockWidgetITF, orientation core.Qt__Orientation) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddDockWidget2(C.QtObjectPtr(ptr.Pointer()), C.int(area), C.QtObjectPtr(PointerFromQDockWidget(dockwidget)), C.int(orientation))
	}
}

func (ptr *QMainWindow) AddToolBar3(title string) *QToolBar {
	if ptr.Pointer() != nil {
		return QToolBarFromPointer(unsafe.Pointer(C.QMainWindow_AddToolBar3(C.QtObjectPtr(ptr.Pointer()), C.CString(title))))
	}
	return nil
}

func (ptr *QMainWindow) AddToolBar2(toolbar QToolBarITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBar2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQToolBar(toolbar)))
	}
}

func (ptr *QMainWindow) AddToolBar(area core.Qt__ToolBarArea, toolbar QToolBarITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBar(C.QtObjectPtr(ptr.Pointer()), C.int(area), C.QtObjectPtr(PointerFromQToolBar(toolbar)))
	}
}

func (ptr *QMainWindow) AddToolBarBreak(area core.Qt__ToolBarArea) {
	if ptr.Pointer() != nil {
		C.QMainWindow_AddToolBarBreak(C.QtObjectPtr(ptr.Pointer()), C.int(area))
	}
}

func (ptr *QMainWindow) CentralWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QMainWindow_CentralWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMainWindow) Corner(corner core.Qt__Corner) core.Qt__DockWidgetArea {
	if ptr.Pointer() != nil {
		return core.Qt__DockWidgetArea(C.QMainWindow_Corner(C.QtObjectPtr(ptr.Pointer()), C.int(corner)))
	}
	return 0
}

func (ptr *QMainWindow) CreatePopupMenu() *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QMainWindow_CreatePopupMenu(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMainWindow) DockWidgetArea(dockwidget QDockWidgetITF) core.Qt__DockWidgetArea {
	if ptr.Pointer() != nil {
		return core.Qt__DockWidgetArea(C.QMainWindow_DockWidgetArea(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDockWidget(dockwidget))))
	}
	return 0
}

func (ptr *QMainWindow) InsertToolBar(before QToolBarITF, toolbar QToolBarITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_InsertToolBar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQToolBar(before)), C.QtObjectPtr(PointerFromQToolBar(toolbar)))
	}
}

func (ptr *QMainWindow) InsertToolBarBreak(before QToolBarITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_InsertToolBarBreak(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQToolBar(before)))
	}
}

func (ptr *QMainWindow) MenuBar() *QMenuBar {
	if ptr.Pointer() != nil {
		return QMenuBarFromPointer(unsafe.Pointer(C.QMainWindow_MenuBar(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMainWindow) MenuWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QMainWindow_MenuWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMainWindow) RemoveDockWidget(dockwidget QDockWidgetITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_RemoveDockWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDockWidget(dockwidget)))
	}
}

func (ptr *QMainWindow) RemoveToolBar(toolbar QToolBarITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_RemoveToolBar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQToolBar(toolbar)))
	}
}

func (ptr *QMainWindow) RemoveToolBarBreak(before QToolBarITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_RemoveToolBarBreak(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQToolBar(before)))
	}
}

func (ptr *QMainWindow) RestoreDockWidget(dockwidget QDockWidgetITF) bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_RestoreDockWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDockWidget(dockwidget))) != 0
	}
	return false
}

func (ptr *QMainWindow) RestoreState(state core.QByteArrayITF, version int) bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_RestoreState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(state)), C.int(version)) != 0
	}
	return false
}

func (ptr *QMainWindow) SetCentralWidget(widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetCentralWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QMainWindow) SetCorner(corner core.Qt__Corner, area core.Qt__DockWidgetArea) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetCorner(C.QtObjectPtr(ptr.Pointer()), C.int(corner), C.int(area))
	}
}

func (ptr *QMainWindow) SetMenuBar(menuBar QMenuBarITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetMenuBar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMenuBar(menuBar)))
	}
}

func (ptr *QMainWindow) SetMenuWidget(menuBar QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetMenuWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(menuBar)))
	}
}

func (ptr *QMainWindow) SetStatusBar(statusbar QStatusBarITF) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetStatusBar(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStatusBar(statusbar)))
	}
}

func (ptr *QMainWindow) SetTabPosition(areas core.Qt__DockWidgetArea, tabPosition QTabWidget__TabPosition) {
	if ptr.Pointer() != nil {
		C.QMainWindow_SetTabPosition(C.QtObjectPtr(ptr.Pointer()), C.int(areas), C.int(tabPosition))
	}
}

func (ptr *QMainWindow) StatusBar() *QStatusBar {
	if ptr.Pointer() != nil {
		return QStatusBarFromPointer(unsafe.Pointer(C.QMainWindow_StatusBar(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMainWindow) TabPosition(area core.Qt__DockWidgetArea) QTabWidget__TabPosition {
	if ptr.Pointer() != nil {
		return QTabWidget__TabPosition(C.QMainWindow_TabPosition(C.QtObjectPtr(ptr.Pointer()), C.int(area)))
	}
	return 0
}

func (ptr *QMainWindow) TakeCentralWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QMainWindow_TakeCentralWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QMainWindow) ToolBarArea(toolbar QToolBarITF) core.Qt__ToolBarArea {
	if ptr.Pointer() != nil {
		return core.Qt__ToolBarArea(C.QMainWindow_ToolBarArea(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQToolBar(toolbar))))
	}
	return 0
}

func (ptr *QMainWindow) ToolBarBreak(toolbar QToolBarITF) bool {
	if ptr.Pointer() != nil {
		return C.QMainWindow_ToolBarBreak(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQToolBar(toolbar))) != 0
	}
	return false
}

func (ptr *QMainWindow) ConnectToolButtonStyleChanged(f func(toolButtonStyle core.Qt__ToolButtonStyle)) {
	if ptr.Pointer() != nil {
		C.QMainWindow_ConnectToolButtonStyleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "toolButtonStyleChanged", f)
	}
}

func (ptr *QMainWindow) DisconnectToolButtonStyleChanged() {
	if ptr.Pointer() != nil {
		C.QMainWindow_DisconnectToolButtonStyleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "toolButtonStyleChanged")
	}
}

//export callbackQMainWindowToolButtonStyleChanged
func callbackQMainWindowToolButtonStyleChanged(ptrName *C.char, toolButtonStyle C.int) {
	qt.GetSignal(C.GoString(ptrName), "toolButtonStyleChanged").(func(core.Qt__ToolButtonStyle))(core.Qt__ToolButtonStyle(toolButtonStyle))
}

func (ptr *QMainWindow) DestroyQMainWindow() {
	if ptr.Pointer() != nil {
		C.QMainWindow_DestroyQMainWindow(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
