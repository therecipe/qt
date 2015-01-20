package qt

//#include "qmainwindow.h"
import "C"

type qmainwindow struct {
	qwidget
}

type QMainWindow interface {
	QWidget
	AddToolBar1(area ToolBarArea, toolbar QToolBar)
	AddToolBar2(toolbar QToolBar)
	AddToolBar3(title string) QToolBar
	AddToolBarBreak(area ToolBarArea)
	CentralWidget() QWidget
	Corner(corner Corner) DockWidgetArea
	CreatePopupMenu() QMenu
	DocumentMode() bool
	InsertToolBar(before QToolBar, toolbar QToolBar)
	InsertToolBarBreak(before QToolBar)
	IsAnimated() bool
	IsDockNestingEnabled() bool
	MenuBar() QMenuBar
	MenuWidget() QWidget
	RemoveToolBar(toolbar QToolBar)
	RemoveToolBarBreak(before QToolBar)
	SetCentralWidget(widget QWidget)
	SetCorner(corner Corner, area DockWidgetArea)
	SetDocumentMode(enabled bool)
	SetMenuBar(menuBar QMenuBar)
	SetMenuWidget(menuBar QWidget)
	SetStatusBar(statusbar QStatusBar)
	SetToolButtonStyle(toolButtonStyle ToolButtonStyle)
	SetUnifiedTitleAndToolBarOnMac(set bool)
	StatusBar() QStatusBar
	TakeCentralWidget() QWidget
	ToolBarArea(toolbar QToolBar) ToolBarArea
	ToolBarBreak(toolbar QToolBar) bool
	ToolButtonStyle() ToolButtonStyle
	UnifiedTitleAndToolBarOnMac() bool
	ConnectSlotSetAnimated()
	DisconnectSlotSetAnimated()
	SlotSetAnimated(enabled bool)
	ConnectSlotSetDockNestingEnabled()
	DisconnectSlotSetDockNestingEnabled()
	SlotSetDockNestingEnabled(enabled bool)
	ConnectSignalToolButtonStyleChanged(f func())
	DisconnectSignalToolButtonStyleChanged()
	SignalToolButtonStyleChanged() func()
}

func (p *qmainwindow) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qmainwindow) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQMainWindow(parent QWidget, flags WindowType) QMainWindow {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qmainwindow = new(qmainwindow)
	qmainwindow.SetPointer(C.QMainWindow_New_QWidget_WindowType(parentPtr, C.int(flags)))
	qmainwindow.SetObjectName("QMainWindow_" + randomIdentifier())
	return qmainwindow
}

func (p *qmainwindow) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QMainWindow_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qmainwindow) AddToolBar1(area ToolBarArea, toolbar QToolBar) {
	if p.Pointer() != nil {
		var toolbarPtr C.QtObjectPtr
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		C.QMainWindow_AddToolBar_ToolBarArea_QToolBar(p.Pointer(), C.int(area), toolbarPtr)
	}
}

func (p *qmainwindow) AddToolBar2(toolbar QToolBar) {
	if p.Pointer() != nil {
		var toolbarPtr C.QtObjectPtr
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		C.QMainWindow_AddToolBar_QToolBar(p.Pointer(), toolbarPtr)
	}
}

func (p *qmainwindow) AddToolBar3(title string) QToolBar {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtoolbar = new(qtoolbar)
		qtoolbar.SetPointer(C.QMainWindow_AddToolBar_String(p.Pointer(), C.CString(title)))
		if qtoolbar.ObjectName() == "" {
			qtoolbar.SetObjectName("QToolBar_" + randomIdentifier())
		}
		return qtoolbar
	}
}

func (p *qmainwindow) AddToolBarBreak(area ToolBarArea) {
	if p.Pointer() != nil {
		C.QMainWindow_AddToolBarBreak_ToolBarArea(p.Pointer(), C.int(area))
	}
}

func (p *qmainwindow) CentralWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QMainWindow_CentralWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qmainwindow) Corner(corner Corner) DockWidgetArea {
	if p.Pointer() == nil {
		return 0
	}
	return DockWidgetArea(C.QMainWindow_Corner_Corner(p.Pointer(), C.int(corner)))
}

func (p *qmainwindow) CreatePopupMenu() QMenu {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenu = new(qmenu)
		qmenu.SetPointer(C.QMainWindow_CreatePopupMenu(p.Pointer()))
		if qmenu.ObjectName() == "" {
			qmenu.SetObjectName("QMenu_" + randomIdentifier())
		}
		return qmenu
	}
}

func (p *qmainwindow) DocumentMode() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QMainWindow_DocumentMode(p.Pointer()) != 0
}

func (p *qmainwindow) InsertToolBar(before QToolBar, toolbar QToolBar) {
	if p.Pointer() != nil {
		var beforePtr C.QtObjectPtr
		if before != nil {
			beforePtr = before.Pointer()
		}
		var toolbarPtr C.QtObjectPtr
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		C.QMainWindow_InsertToolBar_QToolBar_QToolBar(p.Pointer(), beforePtr, toolbarPtr)
	}
}

func (p *qmainwindow) InsertToolBarBreak(before QToolBar) {
	if p.Pointer() != nil {
		var beforePtr C.QtObjectPtr
		if before != nil {
			beforePtr = before.Pointer()
		}
		C.QMainWindow_InsertToolBarBreak_QToolBar(p.Pointer(), beforePtr)
	}
}

func (p *qmainwindow) IsAnimated() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QMainWindow_IsAnimated(p.Pointer()) != 0
}

func (p *qmainwindow) IsDockNestingEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QMainWindow_IsDockNestingEnabled(p.Pointer()) != 0
}

func (p *qmainwindow) MenuBar() QMenuBar {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenubar = new(qmenubar)
		qmenubar.SetPointer(C.QMainWindow_MenuBar(p.Pointer()))
		if qmenubar.ObjectName() == "" {
			qmenubar.SetObjectName("QMenuBar_" + randomIdentifier())
		}
		return qmenubar
	}
}

func (p *qmainwindow) MenuWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QMainWindow_MenuWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qmainwindow) RemoveToolBar(toolbar QToolBar) {
	if p.Pointer() != nil {
		var toolbarPtr C.QtObjectPtr
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		C.QMainWindow_RemoveToolBar_QToolBar(p.Pointer(), toolbarPtr)
	}
}

func (p *qmainwindow) RemoveToolBarBreak(before QToolBar) {
	if p.Pointer() != nil {
		var beforePtr C.QtObjectPtr
		if before != nil {
			beforePtr = before.Pointer()
		}
		C.QMainWindow_RemoveToolBarBreak_QToolBar(p.Pointer(), beforePtr)
	}
}

func (p *qmainwindow) SetCentralWidget(widget QWidget) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QMainWindow_SetCentralWidget_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qmainwindow) SetCorner(corner Corner, area DockWidgetArea) {
	if p.Pointer() != nil {
		C.QMainWindow_SetCorner_Corner_DockWidgetArea(p.Pointer(), C.int(corner), C.int(area))
	}
}

func (p *qmainwindow) SetDocumentMode(enabled bool) {
	if p.Pointer() != nil {
		C.QMainWindow_SetDocumentMode_Bool(p.Pointer(), goBoolToCInt(enabled))
	}
}

func (p *qmainwindow) SetMenuBar(menuBar QMenuBar) {
	if p.Pointer() != nil {
		var menuBarPtr C.QtObjectPtr
		if menuBar != nil {
			menuBarPtr = menuBar.Pointer()
		}
		C.QMainWindow_SetMenuBar_QMenuBar(p.Pointer(), menuBarPtr)
	}
}

func (p *qmainwindow) SetMenuWidget(menuBar QWidget) {
	if p.Pointer() != nil {
		var menuBarPtr C.QtObjectPtr
		if menuBar != nil {
			menuBarPtr = menuBar.Pointer()
		}
		C.QMainWindow_SetMenuWidget_QWidget(p.Pointer(), menuBarPtr)
	}
}

func (p *qmainwindow) SetStatusBar(statusbar QStatusBar) {
	if p.Pointer() != nil {
		var statusbarPtr C.QtObjectPtr
		if statusbar != nil {
			statusbarPtr = statusbar.Pointer()
		}
		C.QMainWindow_SetStatusBar_QStatusBar(p.Pointer(), statusbarPtr)
	}
}

func (p *qmainwindow) SetToolButtonStyle(toolButtonStyle ToolButtonStyle) {
	if p.Pointer() != nil {
		C.QMainWindow_SetToolButtonStyle_ToolButtonStyle(p.Pointer(), C.int(toolButtonStyle))
	}
}

func (p *qmainwindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	if p.Pointer() != nil {
		C.QMainWindow_SetUnifiedTitleAndToolBarOnMac_Bool(p.Pointer(), goBoolToCInt(set))
	}
}

func (p *qmainwindow) StatusBar() QStatusBar {
	if p.Pointer() == nil {
		return nil
	} else {
		var qstatusbar = new(qstatusbar)
		qstatusbar.SetPointer(C.QMainWindow_StatusBar(p.Pointer()))
		if qstatusbar.ObjectName() == "" {
			qstatusbar.SetObjectName("QStatusBar_" + randomIdentifier())
		}
		return qstatusbar
	}
}

func (p *qmainwindow) TakeCentralWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QMainWindow_TakeCentralWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qmainwindow) ToolBarArea(toolbar QToolBar) ToolBarArea {
	if p.Pointer() == nil {
		return 0
	} else {
		var toolbarPtr C.QtObjectPtr
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		return ToolBarArea(C.QMainWindow_ToolBarArea_QToolBar(p.Pointer(), toolbarPtr))
	}
}

func (p *qmainwindow) ToolBarBreak(toolbar QToolBar) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var toolbarPtr C.QtObjectPtr
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		return C.QMainWindow_ToolBarBreak_QToolBar(p.Pointer(), toolbarPtr) != 0
	}
}

func (p *qmainwindow) ToolButtonStyle() ToolButtonStyle {
	if p.Pointer() == nil {
		return 0
	}
	return ToolButtonStyle(C.QMainWindow_ToolButtonStyle(p.Pointer()))
}

func (p *qmainwindow) UnifiedTitleAndToolBarOnMac() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QMainWindow_UnifiedTitleAndToolBarOnMac(p.Pointer()) != 0
}

func (p *qmainwindow) ConnectSlotSetAnimated() {
	C.QMainWindow_ConnectSlotSetAnimated(p.Pointer())
}

func (p *qmainwindow) DisconnectSlotSetAnimated() {
	C.QMainWindow_DisconnectSlotSetAnimated(p.Pointer())
}

func (p *qmainwindow) SlotSetAnimated(enabled bool) {
	if p.Pointer() != nil {
		C.QMainWindow_SetAnimated_Bool(p.Pointer(), goBoolToCInt(enabled))
	}
}

func (p *qmainwindow) ConnectSlotSetDockNestingEnabled() {
	C.QMainWindow_ConnectSlotSetDockNestingEnabled(p.Pointer())
}

func (p *qmainwindow) DisconnectSlotSetDockNestingEnabled() {
	C.QMainWindow_DisconnectSlotSetDockNestingEnabled(p.Pointer())
}

func (p *qmainwindow) SlotSetDockNestingEnabled(enabled bool) {
	if p.Pointer() != nil {
		C.QMainWindow_SetDockNestingEnabled_Bool(p.Pointer(), goBoolToCInt(enabled))
	}
}

func (p *qmainwindow) ConnectSignalToolButtonStyleChanged(f func()) {
	C.QMainWindow_ConnectSignalToolButtonStyleChanged(p.Pointer())
	connectSignal(p.ObjectName(), "toolButtonStyleChanged", f)
}

func (p *qmainwindow) DisconnectSignalToolButtonStyleChanged() {
	C.QMainWindow_DisconnectSignalToolButtonStyleChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "toolButtonStyleChanged")
}

func (p *qmainwindow) SignalToolButtonStyleChanged() func() {
	return getSignal(p.ObjectName(), "toolButtonStyleChanged")
}

//export callbackQMainWindow
func callbackQMainWindow(ptr C.QtObjectPtr, msg *C.char) {
	var qmainwindow = new(qmainwindow)
	qmainwindow.SetPointer(ptr)
	getSignal(qmainwindow.ObjectName(), C.GoString(msg))()
}
