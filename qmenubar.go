package qt

//#include "qmenubar.h"
import "C"

type qmenubar struct {
	qwidget
}

type QMenuBar interface {
	QWidget
	ActiveAction() QAction
	AddAction1(text string) QAction
	AddAction2(text string, receiver QObject, member string) QAction
	AddAction3(action QAction)
	AddMenu1(menu QMenu) QAction
	AddMenu2(title string) QMenu
	AddSeparator() QAction
	Clear()
	CornerWidget(corner Corner) QWidget
	InsertMenu(before QAction, menu QMenu) QAction
	InsertSeparator(before QAction) QAction
	IsDefaultUp() bool
	IsNativeMenuBar() bool
	SetActiveAction(act QAction)
	SetCornerWidget(widget QWidget, corner Corner)
	SetDefaultUp(defaultAction bool)
	SetNativeMenuBar(nativeMenuBar bool)
	ConnectSlotSetVisible()
	DisconnectSlotSetVisible()
	SlotSetVisible(visible bool)
	ConnectSignalHovered(f func())
	DisconnectSignalHovered()
	SignalHovered() func()
	ConnectSignalTriggered(f func())
	DisconnectSignalTriggered()
	SignalTriggered() func()
}

func (p *qmenubar) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qmenubar) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQMenuBar(parent QWidget) QMenuBar {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qmenubar = new(qmenubar)
	qmenubar.SetPointer(C.QMenuBar_New_QWidget(parentPtr))
	qmenubar.SetObjectName("QMenuBar_" + randomIdentifier())
	return qmenubar
}

func (p *qmenubar) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QMenuBar_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qmenubar) ActiveAction() QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenuBar_ActiveAction(p.Pointer()))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenubar) AddAction1(text string) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenuBar_AddAction_String(p.Pointer(), C.CString(text)))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenubar) AddAction2(text string, receiver QObject, member string) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var receiverPtr C.QtObjectPtr
		if receiver != nil {
			receiverPtr = receiver.Pointer()
		}
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenuBar_AddAction_String_QObject_String(p.Pointer(), C.CString(text), receiverPtr, C.CString(member)))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenubar) AddAction3(action QAction) {
	if p.Pointer() != nil {
		var actionPtr C.QtObjectPtr
		if action != nil {
			actionPtr = action.Pointer()
		}
		C.QMenuBar_AddAction_QAction(p.Pointer(), actionPtr)
	}
}

func (p *qmenubar) AddMenu1(menu QMenu) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var menuPtr C.QtObjectPtr
		if menu != nil {
			menuPtr = menu.Pointer()
		}
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenuBar_AddMenu_QMenu(p.Pointer(), menuPtr))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenubar) AddMenu2(title string) QMenu {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenu = new(qmenu)
		qmenu.SetPointer(C.QMenuBar_AddMenu_String(p.Pointer(), C.CString(title)))
		if qmenu.ObjectName() == "" {
			qmenu.SetObjectName("QMenu_" + randomIdentifier())
		}
		return qmenu
	}
}

func (p *qmenubar) AddSeparator() QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenuBar_AddSeparator(p.Pointer()))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenubar) Clear() {
	if p.Pointer() != nil {
		C.QMenuBar_Clear(p.Pointer())
	}
}

func (p *qmenubar) CornerWidget(corner Corner) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QMenuBar_CornerWidget_Corner(p.Pointer(), C.int(corner)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qmenubar) InsertMenu(before QAction, menu QMenu) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var beforePtr C.QtObjectPtr
		if before != nil {
			beforePtr = before.Pointer()
		}
		var menuPtr C.QtObjectPtr
		if menu != nil {
			menuPtr = menu.Pointer()
		}
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenuBar_InsertMenu_QAction_QMenu(p.Pointer(), beforePtr, menuPtr))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenubar) InsertSeparator(before QAction) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var beforePtr C.QtObjectPtr
		if before != nil {
			beforePtr = before.Pointer()
		}
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenuBar_InsertSeparator_QAction(p.Pointer(), beforePtr))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenubar) IsDefaultUp() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QMenuBar_IsDefaultUp(p.Pointer()) != 0
}

func (p *qmenubar) IsNativeMenuBar() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QMenuBar_IsNativeMenuBar(p.Pointer()) != 0
}

func (p *qmenubar) SetActiveAction(act QAction) {
	if p.Pointer() != nil {
		var actPtr C.QtObjectPtr
		if act != nil {
			actPtr = act.Pointer()
		}
		C.QMenuBar_SetActiveAction_QAction(p.Pointer(), actPtr)
	}
}

func (p *qmenubar) SetCornerWidget(widget QWidget, corner Corner) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QMenuBar_SetCornerWidget_QWidget_Corner(p.Pointer(), widgetPtr, C.int(corner))
	}
}

func (p *qmenubar) SetDefaultUp(defaultAction bool) {
	if p.Pointer() != nil {
		C.QMenuBar_SetDefaultUp_Bool(p.Pointer(), goBoolToCInt(defaultAction))
	}
}

func (p *qmenubar) SetNativeMenuBar(nativeMenuBar bool) {
	if p.Pointer() != nil {
		C.QMenuBar_SetNativeMenuBar_Bool(p.Pointer(), goBoolToCInt(nativeMenuBar))
	}
}

func (p *qmenubar) ConnectSlotSetVisible() {
	C.QMenuBar_ConnectSlotSetVisible(p.Pointer())
}

func (p *qmenubar) DisconnectSlotSetVisible() {
	C.QMenuBar_DisconnectSlotSetVisible(p.Pointer())
}

func (p *qmenubar) SlotSetVisible(visible bool) {
	if p.Pointer() != nil {
		C.QMenuBar_SetVisible_Bool(p.Pointer(), goBoolToCInt(visible))
	}
}

func (p *qmenubar) ConnectSignalHovered(f func()) {
	C.QMenuBar_ConnectSignalHovered(p.Pointer())
	connectSignal(p.ObjectName(), "hovered", f)
}

func (p *qmenubar) DisconnectSignalHovered() {
	C.QMenuBar_DisconnectSignalHovered(p.Pointer())
	disconnectSignal(p.ObjectName(), "hovered")
}

func (p *qmenubar) SignalHovered() func() {
	return getSignal(p.ObjectName(), "hovered")
}

func (p *qmenubar) ConnectSignalTriggered(f func()) {
	C.QMenuBar_ConnectSignalTriggered(p.Pointer())
	connectSignal(p.ObjectName(), "triggered", f)
}

func (p *qmenubar) DisconnectSignalTriggered() {
	C.QMenuBar_DisconnectSignalTriggered(p.Pointer())
	disconnectSignal(p.ObjectName(), "triggered")
}

func (p *qmenubar) SignalTriggered() func() {
	return getSignal(p.ObjectName(), "triggered")
}

//export callbackQMenuBar
func callbackQMenuBar(ptr C.QtObjectPtr, msg *C.char) {
	var qmenubar = new(qmenubar)
	qmenubar.SetPointer(ptr)
	getSignal(qmenubar.ObjectName(), C.GoString(msg))()
}
