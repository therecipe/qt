package qt

//#include "qmainwindow.h"
import "C"

type qmainwindow struct {
	qwidget
}

type QMainWindow interface {
	QWidget
	AddToolBar_ToolBarArea_QToolBar(area ToolBarArea, toolbar QToolBar)
	AddToolBar_QToolBar(toolbar QToolBar)
	AddToolBar_String(title string) QToolBar
	AddToolBarBreak_ToolBarArea(area ToolBarArea)
	CentralWidget() QWidget
	Corner_Corner(corner Corner) DockWidgetArea
	CreatePopupMenu() QMenu
	DocumentMode() bool
	InsertToolBar_QToolBar_QToolBar(before QToolBar, toolbar QToolBar)
	InsertToolBarBreak_QToolBar(before QToolBar)
	IsAnimated() bool
	IsDockNestingEnabled() bool
	MenuBar() QMenuBar
	MenuWidget() QWidget
	RemoveToolBar_QToolBar(toolbar QToolBar)
	RemoveToolBarBreak_QToolBar(before QToolBar)
	SetCentralWidget_QWidget(widget QWidget)
	SetCorner_Corner_DockWidgetArea(corner Corner, area DockWidgetArea)
	SetDocumentMode_Bool(enabled bool)
	SetMenuBar_QMenuBar(menuBar QMenuBar)
	SetMenuWidget_QWidget(menuBar QWidget)
	SetStatusBar_QStatusBar(statusbar QStatusBar)
	SetToolButtonStyle_ToolButtonStyle(toolButtonStyle ToolButtonStyle)
	SetUnifiedTitleAndToolBarOnMac_Bool(set bool)
	StatusBar() QStatusBar
	TakeCentralWidget() QWidget
	ToolBarArea_QToolBar(toolbar QToolBar) ToolBarArea
	ToolBarBreak_QToolBar(toolbar QToolBar) bool
	ToolButtonStyle() ToolButtonStyle
	UnifiedTitleAndToolBarOnMac() bool
	ConnectSlotSetAnimated()
	DisconnectSlotSetAnimated()
	SlotSetAnimated_Bool(enabled bool)
	ConnectSlotSetDockNestingEnabled()
	DisconnectSlotSetDockNestingEnabled()
	SlotSetDockNestingEnabled_Bool(enabled bool)
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

func NewQMainWindow_QWidget_WindowType(parent QWidget, flags WindowType) QMainWindow {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qmainwindow = new(qmainwindow)
	qmainwindow.SetPointer(C.QMainWindow_New_QWidget_WindowType(parentPtr, C.int(flags)))
	qmainwindow.SetObjectName_String("QMainWindow_" + randomIdentifier())
	return qmainwindow
}

func (p *qmainwindow) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QMainWindow_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qmainwindow) AddToolBar_ToolBarArea_QToolBar(area ToolBarArea, toolbar QToolBar) {
	if p.Pointer() == nil {
	} else {
		var toolbarPtr C.QtObjectPtr = nil
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		C.QMainWindow_AddToolBar_ToolBarArea_QToolBar(p.Pointer(), C.int(area), toolbarPtr)
	}
}

func (p *qmainwindow) AddToolBar_QToolBar(toolbar QToolBar) {
	if p.Pointer() == nil {
	} else {
		var toolbarPtr C.QtObjectPtr = nil
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		C.QMainWindow_AddToolBar_QToolBar(p.Pointer(), toolbarPtr)
	}
}

func (p *qmainwindow) AddToolBar_String(title string) QToolBar {
	if p.Pointer() == nil {
		return nil
	} else {
		var qtoolbar = new(qtoolbar)
		qtoolbar.SetPointer(C.QMainWindow_AddToolBar_String(p.Pointer(), C.CString(title)))
		if qtoolbar.ObjectName() == "" {
			qtoolbar.SetObjectName_String("QToolBar_" + randomIdentifier())
		}
		return qtoolbar
	}
}

func (p *qmainwindow) AddToolBarBreak_ToolBarArea(area ToolBarArea) {
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
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qmainwindow) Corner_Corner(corner Corner) DockWidgetArea {
	if p.Pointer() == nil {
		return 0
	} else {
		return DockWidgetArea(C.QMainWindow_Corner_Corner(p.Pointer(), C.int(corner)))
	}
}

func (p *qmainwindow) CreatePopupMenu() QMenu {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenu = new(qmenu)
		qmenu.SetPointer(C.QMainWindow_CreatePopupMenu(p.Pointer()))
		if qmenu.ObjectName() == "" {
			qmenu.SetObjectName_String("QMenu_" + randomIdentifier())
		}
		return qmenu
	}
}

func (p *qmainwindow) DocumentMode() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QMainWindow_DocumentMode(p.Pointer()) != 0
	}
}

func (p *qmainwindow) InsertToolBar_QToolBar_QToolBar(before QToolBar, toolbar QToolBar) {
	if p.Pointer() == nil {
	} else {
		var beforePtr C.QtObjectPtr = nil
		if before != nil {
			beforePtr = before.Pointer()
		}
		var toolbarPtr C.QtObjectPtr = nil
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		C.QMainWindow_InsertToolBar_QToolBar_QToolBar(p.Pointer(), beforePtr, toolbarPtr)
	}
}

func (p *qmainwindow) InsertToolBarBreak_QToolBar(before QToolBar) {
	if p.Pointer() == nil {
	} else {
		var beforePtr C.QtObjectPtr = nil
		if before != nil {
			beforePtr = before.Pointer()
		}
		C.QMainWindow_InsertToolBarBreak_QToolBar(p.Pointer(), beforePtr)
	}
}

func (p *qmainwindow) IsAnimated() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QMainWindow_IsAnimated(p.Pointer()) != 0
	}
}

func (p *qmainwindow) IsDockNestingEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QMainWindow_IsDockNestingEnabled(p.Pointer()) != 0
	}
}

func (p *qmainwindow) MenuBar() QMenuBar {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenubar = new(qmenubar)
		qmenubar.SetPointer(C.QMainWindow_MenuBar(p.Pointer()))
		if qmenubar.ObjectName() == "" {
			qmenubar.SetObjectName_String("QMenuBar_" + randomIdentifier())
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
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qmainwindow) RemoveToolBar_QToolBar(toolbar QToolBar) {
	if p.Pointer() == nil {
	} else {
		var toolbarPtr C.QtObjectPtr = nil
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		C.QMainWindow_RemoveToolBar_QToolBar(p.Pointer(), toolbarPtr)
	}
}

func (p *qmainwindow) RemoveToolBarBreak_QToolBar(before QToolBar) {
	if p.Pointer() == nil {
	} else {
		var beforePtr C.QtObjectPtr = nil
		if before != nil {
			beforePtr = before.Pointer()
		}
		C.QMainWindow_RemoveToolBarBreak_QToolBar(p.Pointer(), beforePtr)
	}
}

func (p *qmainwindow) SetCentralWidget_QWidget(widget QWidget) {
	if p.Pointer() == nil {
	} else {
		var widgetPtr C.QtObjectPtr = nil
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QMainWindow_SetCentralWidget_QWidget(p.Pointer(), widgetPtr)
	}
}

func (p *qmainwindow) SetCorner_Corner_DockWidgetArea(corner Corner, area DockWidgetArea) {
	if p.Pointer() != nil {
		C.QMainWindow_SetCorner_Corner_DockWidgetArea(p.Pointer(), C.int(corner), C.int(area))
	}
}

func (p *qmainwindow) SetDocumentMode_Bool(enabled bool) {
	if p.Pointer() != nil {
		C.QMainWindow_SetDocumentMode_Bool(p.Pointer(), goBoolToCInt(enabled))
	}
}

func (p *qmainwindow) SetMenuBar_QMenuBar(menuBar QMenuBar) {
	if p.Pointer() == nil {
	} else {
		var menuBarPtr C.QtObjectPtr = nil
		if menuBar != nil {
			menuBarPtr = menuBar.Pointer()
		}
		C.QMainWindow_SetMenuBar_QMenuBar(p.Pointer(), menuBarPtr)
	}
}

func (p *qmainwindow) SetMenuWidget_QWidget(menuBar QWidget) {
	if p.Pointer() == nil {
	} else {
		var menuBarPtr C.QtObjectPtr = nil
		if menuBar != nil {
			menuBarPtr = menuBar.Pointer()
		}
		C.QMainWindow_SetMenuWidget_QWidget(p.Pointer(), menuBarPtr)
	}
}

func (p *qmainwindow) SetStatusBar_QStatusBar(statusbar QStatusBar) {
	if p.Pointer() == nil {
	} else {
		var statusbarPtr C.QtObjectPtr = nil
		if statusbar != nil {
			statusbarPtr = statusbar.Pointer()
		}
		C.QMainWindow_SetStatusBar_QStatusBar(p.Pointer(), statusbarPtr)
	}
}

func (p *qmainwindow) SetToolButtonStyle_ToolButtonStyle(toolButtonStyle ToolButtonStyle) {
	if p.Pointer() != nil {
		C.QMainWindow_SetToolButtonStyle_ToolButtonStyle(p.Pointer(), C.int(toolButtonStyle))
	}
}

func (p *qmainwindow) SetUnifiedTitleAndToolBarOnMac_Bool(set bool) {
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
			qstatusbar.SetObjectName_String("QStatusBar_" + randomIdentifier())
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
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qmainwindow) ToolBarArea_QToolBar(toolbar QToolBar) ToolBarArea {
	if p.Pointer() == nil {
		return 0
	} else {
		var toolbarPtr C.QtObjectPtr = nil
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		return ToolBarArea(C.QMainWindow_ToolBarArea_QToolBar(p.Pointer(), toolbarPtr))
	}
}

func (p *qmainwindow) ToolBarBreak_QToolBar(toolbar QToolBar) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var toolbarPtr C.QtObjectPtr = nil
		if toolbar != nil {
			toolbarPtr = toolbar.Pointer()
		}
		return C.QMainWindow_ToolBarBreak_QToolBar(p.Pointer(), toolbarPtr) != 0
	}
}

func (p *qmainwindow) ToolButtonStyle() ToolButtonStyle {
	if p.Pointer() == nil {
		return 0
	} else {
		return ToolButtonStyle(C.QMainWindow_ToolButtonStyle(p.Pointer()))
	}
}

func (p *qmainwindow) UnifiedTitleAndToolBarOnMac() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QMainWindow_UnifiedTitleAndToolBarOnMac(p.Pointer()) != 0
	}
}

func (p *qmainwindow) ConnectSlotSetAnimated() {
	C.QMainWindow_ConnectSlotSetAnimated(p.Pointer())
}

func (p *qmainwindow) DisconnectSlotSetAnimated() {
	C.QMainWindow_DisconnectSlotSetAnimated(p.Pointer())
}

func (p *qmainwindow) SlotSetAnimated_Bool(enabled bool) {
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

func (p *qmainwindow) SlotSetDockNestingEnabled_Bool(enabled bool) {
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
