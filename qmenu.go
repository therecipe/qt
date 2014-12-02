package qt

//#include "qmenu.h"
import "C"

type qmenu struct {
	qwidget
}

type QMenu interface {
	QWidget
	ActiveAction() QAction
	AddAction_String(text string) QAction
	AddAction_QAction(action QAction)
	AddMenu_QMenu(menu QMenu) QAction
	AddMenu_String(title string) QMenu
	AddSection_String(text string) QAction
	AddSeparator() QAction
	Clear()
	DefaultAction() QAction
	Exec() QAction
	HideTearOffMenu()
	InsertMenu_QAction_QMenu(before QAction, menu QMenu) QAction
	InsertSection_QAction_String(before QAction, text string) QAction
	InsertSeparator_QAction(before QAction) QAction
	IsEmpty() bool
	IsTearOffEnabled() bool
	IsTearOffMenuVisible() bool
	MenuAction() QAction
	SeparatorsCollapsible() bool
	SetActiveAction_QAction(act QAction)
	SetDefaultAction_QAction(act QAction)
	SetSeparatorsCollapsible_Bool(collapse bool)
	SetTearOffEnabled_Bool(tearOffEnabled bool)
	SetTitle_String(title string)
	SetToolTipsVisible_Bool(visible bool)
	Title() string
	ToolTipsVisible() bool
	ConnectSignalAboutToHide(f func())
	DisconnectSignalAboutToHide()
	SignalAboutToHide() func()
	ConnectSignalAboutToShow(f func())
	DisconnectSignalAboutToShow()
	SignalAboutToShow() func()
	ConnectSignalHovered(f func())
	DisconnectSignalHovered()
	SignalHovered() func()
	ConnectSignalTriggered(f func())
	DisconnectSignalTriggered()
	SignalTriggered() func()
}

func (p *qmenu) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qmenu) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQMenu_QWidget(parent QWidget) QMenu {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qmenu = new(qmenu)
	qmenu.SetPointer(C.QMenu_New_QWidget(parentPtr))
	qmenu.SetObjectName_String("QMenu_" + randomIdentifier())
	return qmenu
}

func NewQMenu_String_QWidget(title string, parent QWidget) QMenu {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qmenu = new(qmenu)
	qmenu.SetPointer(C.QMenu_New_String_QWidget(C.CString(title), parentPtr))
	qmenu.SetObjectName_String("QMenu_" + randomIdentifier())
	return qmenu
}

func (p *qmenu) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QMenu_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qmenu) ActiveAction() QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_ActiveAction(p.Pointer()))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) AddAction_String(text string) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_AddAction_String(p.Pointer(), C.CString(text)))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) AddAction_QAction(action QAction) {
	if p.Pointer() == nil {
	} else {
		var actionPtr C.QtObjectPtr = nil
		if action != nil {
			actionPtr = action.Pointer()
		}
		C.QMenu_AddAction_QAction(p.Pointer(), actionPtr)
	}
}

func (p *qmenu) AddMenu_QMenu(menu QMenu) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var menuPtr C.QtObjectPtr = nil
		if menu != nil {
			menuPtr = menu.Pointer()
		}
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_AddMenu_QMenu(p.Pointer(), menuPtr))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) AddMenu_String(title string) QMenu {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenu = new(qmenu)
		qmenu.SetPointer(C.QMenu_AddMenu_String(p.Pointer(), C.CString(title)))
		if qmenu.ObjectName() == "" {
			qmenu.SetObjectName_String("QMenu_" + randomIdentifier())
		}
		return qmenu
	}
}

func (p *qmenu) AddSection_String(text string) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_AddSection_String(p.Pointer(), C.CString(text)))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) AddSeparator() QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_AddSeparator(p.Pointer()))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) Clear() {
	if p.Pointer() != nil {
		C.QMenu_Clear(p.Pointer())
	}
}

func (p *qmenu) DefaultAction() QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_DefaultAction(p.Pointer()))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) Exec() QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_Exec(p.Pointer()))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) HideTearOffMenu() {
	if p.Pointer() != nil {
		C.QMenu_HideTearOffMenu(p.Pointer())
	}
}

func (p *qmenu) InsertMenu_QAction_QMenu(before QAction, menu QMenu) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var beforePtr C.QtObjectPtr = nil
		if before != nil {
			beforePtr = before.Pointer()
		}
		var menuPtr C.QtObjectPtr = nil
		if menu != nil {
			menuPtr = menu.Pointer()
		}
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_InsertMenu_QAction_QMenu(p.Pointer(), beforePtr, menuPtr))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) InsertSection_QAction_String(before QAction, text string) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var beforePtr C.QtObjectPtr = nil
		if before != nil {
			beforePtr = before.Pointer()
		}
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_InsertSection_QAction_String(p.Pointer(), beforePtr, C.CString(text)))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) InsertSeparator_QAction(before QAction) QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var beforePtr C.QtObjectPtr = nil
		if before != nil {
			beforePtr = before.Pointer()
		}
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_InsertSeparator_QAction(p.Pointer(), beforePtr))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) IsEmpty() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QMenu_IsEmpty(p.Pointer()) != 0
	}
}

func (p *qmenu) IsTearOffEnabled() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QMenu_IsTearOffEnabled(p.Pointer()) != 0
	}
}

func (p *qmenu) IsTearOffMenuVisible() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QMenu_IsTearOffMenuVisible(p.Pointer()) != 0
	}
}

func (p *qmenu) MenuAction() QAction {
	if p.Pointer() == nil {
		return nil
	} else {
		var qaction = new(qaction)
		qaction.SetPointer(C.QMenu_MenuAction(p.Pointer()))
		if qaction.ObjectName() == "" {
			qaction.SetObjectName_String("QAction_" + randomIdentifier())
		}
		return qaction
	}
}

func (p *qmenu) SeparatorsCollapsible() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QMenu_SeparatorsCollapsible(p.Pointer()) != 0
	}
}

func (p *qmenu) SetActiveAction_QAction(act QAction) {
	if p.Pointer() == nil {
	} else {
		var actPtr C.QtObjectPtr = nil
		if act != nil {
			actPtr = act.Pointer()
		}
		C.QMenu_SetActiveAction_QAction(p.Pointer(), actPtr)
	}
}

func (p *qmenu) SetDefaultAction_QAction(act QAction) {
	if p.Pointer() == nil {
	} else {
		var actPtr C.QtObjectPtr = nil
		if act != nil {
			actPtr = act.Pointer()
		}
		C.QMenu_SetDefaultAction_QAction(p.Pointer(), actPtr)
	}
}

func (p *qmenu) SetSeparatorsCollapsible_Bool(collapse bool) {
	if p.Pointer() != nil {
		C.QMenu_SetSeparatorsCollapsible_Bool(p.Pointer(), goBoolToCInt(collapse))
	}
}

func (p *qmenu) SetTearOffEnabled_Bool(tearOffEnabled bool) {
	if p.Pointer() != nil {
		C.QMenu_SetTearOffEnabled_Bool(p.Pointer(), goBoolToCInt(tearOffEnabled))
	}
}

func (p *qmenu) SetTitle_String(title string) {
	if p.Pointer() != nil {
		C.QMenu_SetTitle_String(p.Pointer(), C.CString(title))
	}
}

func (p *qmenu) SetToolTipsVisible_Bool(visible bool) {
	if p.Pointer() != nil {
		C.QMenu_SetToolTipsVisible_Bool(p.Pointer(), goBoolToCInt(visible))
	}
}

func (p *qmenu) Title() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QMenu_Title(p.Pointer()))
	}
}

func (p *qmenu) ToolTipsVisible() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QMenu_ToolTipsVisible(p.Pointer()) != 0
	}
}

func (p *qmenu) ConnectSignalAboutToHide(f func()) {
	C.QMenu_ConnectSignalAboutToHide(p.Pointer())
	connectSignal(p.ObjectName(), "aboutToHide", f)
}

func (p *qmenu) DisconnectSignalAboutToHide() {
	C.QMenu_DisconnectSignalAboutToHide(p.Pointer())
	disconnectSignal(p.ObjectName(), "aboutToHide")
}

func (p *qmenu) SignalAboutToHide() func() {
	return getSignal(p.ObjectName(), "aboutToHide")
}

func (p *qmenu) ConnectSignalAboutToShow(f func()) {
	C.QMenu_ConnectSignalAboutToShow(p.Pointer())
	connectSignal(p.ObjectName(), "aboutToShow", f)
}

func (p *qmenu) DisconnectSignalAboutToShow() {
	C.QMenu_DisconnectSignalAboutToShow(p.Pointer())
	disconnectSignal(p.ObjectName(), "aboutToShow")
}

func (p *qmenu) SignalAboutToShow() func() {
	return getSignal(p.ObjectName(), "aboutToShow")
}

func (p *qmenu) ConnectSignalHovered(f func()) {
	C.QMenu_ConnectSignalHovered(p.Pointer())
	connectSignal(p.ObjectName(), "hovered", f)
}

func (p *qmenu) DisconnectSignalHovered() {
	C.QMenu_DisconnectSignalHovered(p.Pointer())
	disconnectSignal(p.ObjectName(), "hovered")
}

func (p *qmenu) SignalHovered() func() {
	return getSignal(p.ObjectName(), "hovered")
}

func (p *qmenu) ConnectSignalTriggered(f func()) {
	C.QMenu_ConnectSignalTriggered(p.Pointer())
	connectSignal(p.ObjectName(), "triggered", f)
}

func (p *qmenu) DisconnectSignalTriggered() {
	C.QMenu_DisconnectSignalTriggered(p.Pointer())
	disconnectSignal(p.ObjectName(), "triggered")
}

func (p *qmenu) SignalTriggered() func() {
	return getSignal(p.ObjectName(), "triggered")
}

//export callbackQMenu
func callbackQMenu(ptr C.QtObjectPtr, msg *C.char) {
	var qmenu = new(qmenu)
	qmenu.SetPointer(ptr)
	getSignal(qmenu.ObjectName(), C.GoString(msg))()
}
