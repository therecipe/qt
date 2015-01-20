package qt

//#include "qaction.h"
import "C"

type qaction struct {
	qobject
}

type QAction interface {
	QObject
	AutoRepeat() bool
	IconText() string
	IsCheckable() bool
	IsChecked() bool
	IsEnabled() bool
	IsIconVisibleInMenu() bool
	IsSeparator() bool
	IsVisible() bool
	Menu() QMenu
	ParentWidget() QWidget
	SetAutoRepeat(autoRepeat bool)
	SetCheckable(checkable bool)
	SetIconText(text string)
	SetIconVisibleInMenu(visible bool)
	SetMenu(menu QMenu)
	SetSeparator(bo bool)
	SetShortcutContext(context ShortcutContext)
	SetStatusTip(statusTip string)
	SetText(text string)
	SetToolTip(tip string)
	SetWhatsThis(what string)
	ShortcutContext() ShortcutContext
	ShowStatusText(widget QWidget) bool
	StatusTip() string
	Text() string
	ToolTip() string
	WhatsThis() string
	ConnectSlotHover()
	DisconnectSlotHover()
	SlotHover()
	ConnectSlotSetChecked()
	DisconnectSlotSetChecked()
	SlotSetChecked(bo bool)
	ConnectSlotSetDisabled()
	DisconnectSlotSetDisabled()
	SlotSetDisabled(bo bool)
	ConnectSlotSetEnabled()
	DisconnectSlotSetEnabled()
	SlotSetEnabled(bo bool)
	ConnectSlotSetVisible()
	DisconnectSlotSetVisible()
	SlotSetVisible(bo bool)
	ConnectSlotToggle()
	DisconnectSlotToggle()
	SlotToggle()
	ConnectSlotTrigger()
	DisconnectSlotTrigger()
	SlotTrigger()
	ConnectSignalChanged(f func())
	DisconnectSignalChanged()
	SignalChanged() func()
	ConnectSignalHovered(f func())
	DisconnectSignalHovered()
	SignalHovered() func()
	ConnectSignalToggled(f func())
	DisconnectSignalToggled()
	SignalToggled() func()
	ConnectSignalTriggered(f func())
	DisconnectSignalTriggered()
	SignalTriggered() func()
}

func (p *qaction) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qaction) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQAction1(parent QObject) QAction {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qaction = new(qaction)
	qaction.SetPointer(C.QAction_New_QObject(parentPtr))
	qaction.SetObjectName("QAction_" + randomIdentifier())
	return qaction
}

func NewQAction2(text string, parent QObject) QAction {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qaction = new(qaction)
	qaction.SetPointer(C.QAction_New_String_QObject(C.CString(text), parentPtr))
	qaction.SetObjectName("QAction_" + randomIdentifier())
	return qaction
}

func (p *qaction) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QAction_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qaction) AutoRepeat() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAction_AutoRepeat(p.Pointer()) != 0
}

func (p *qaction) IconText() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QAction_IconText(p.Pointer()))
}

func (p *qaction) IsCheckable() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAction_IsCheckable(p.Pointer()) != 0
}

func (p *qaction) IsChecked() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAction_IsChecked(p.Pointer()) != 0
}

func (p *qaction) IsEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAction_IsEnabled(p.Pointer()) != 0
}

func (p *qaction) IsIconVisibleInMenu() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAction_IsIconVisibleInMenu(p.Pointer()) != 0
}

func (p *qaction) IsSeparator() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAction_IsSeparator(p.Pointer()) != 0
}

func (p *qaction) IsVisible() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QAction_IsVisible(p.Pointer()) != 0
}

func (p *qaction) Menu() QMenu {
	if p.Pointer() == nil {
		return nil
	} else {
		var qmenu = new(qmenu)
		qmenu.SetPointer(C.QAction_Menu(p.Pointer()))
		if qmenu.ObjectName() == "" {
			qmenu.SetObjectName("QMenu_" + randomIdentifier())
		}
		return qmenu
	}
}

func (p *qaction) ParentWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QAction_ParentWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qaction) SetAutoRepeat(autoRepeat bool) {
	if p.Pointer() != nil {
		C.QAction_SetAutoRepeat_Bool(p.Pointer(), goBoolToCInt(autoRepeat))
	}
}

func (p *qaction) SetCheckable(checkable bool) {
	if p.Pointer() != nil {
		C.QAction_SetCheckable_Bool(p.Pointer(), goBoolToCInt(checkable))
	}
}

func (p *qaction) SetIconText(text string) {
	if p.Pointer() != nil {
		C.QAction_SetIconText_String(p.Pointer(), C.CString(text))
	}
}

func (p *qaction) SetIconVisibleInMenu(visible bool) {
	if p.Pointer() != nil {
		C.QAction_SetIconVisibleInMenu_Bool(p.Pointer(), goBoolToCInt(visible))
	}
}

func (p *qaction) SetMenu(menu QMenu) {
	if p.Pointer() != nil {
		var menuPtr C.QtObjectPtr
		if menu != nil {
			menuPtr = menu.Pointer()
		}
		C.QAction_SetMenu_QMenu(p.Pointer(), menuPtr)
	}
}

func (p *qaction) SetSeparator(bo bool) {
	if p.Pointer() != nil {
		C.QAction_SetSeparator_Bool(p.Pointer(), goBoolToCInt(bo))
	}
}

func (p *qaction) SetShortcutContext(context ShortcutContext) {
	if p.Pointer() != nil {
		C.QAction_SetShortcutContext_ShortcutContext(p.Pointer(), C.int(context))
	}
}

func (p *qaction) SetStatusTip(statusTip string) {
	if p.Pointer() != nil {
		C.QAction_SetStatusTip_String(p.Pointer(), C.CString(statusTip))
	}
}

func (p *qaction) SetText(text string) {
	if p.Pointer() != nil {
		C.QAction_SetText_String(p.Pointer(), C.CString(text))
	}
}

func (p *qaction) SetToolTip(tip string) {
	if p.Pointer() != nil {
		C.QAction_SetToolTip_String(p.Pointer(), C.CString(tip))
	}
}

func (p *qaction) SetWhatsThis(what string) {
	if p.Pointer() != nil {
		C.QAction_SetWhatsThis_String(p.Pointer(), C.CString(what))
	}
}

func (p *qaction) ShortcutContext() ShortcutContext {
	if p.Pointer() == nil {
		return 0
	}
	return ShortcutContext(C.QAction_ShortcutContext(p.Pointer()))
}

func (p *qaction) ShowStatusText(widget QWidget) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		return C.QAction_ShowStatusText_QWidget(p.Pointer(), widgetPtr) != 0
	}
}

func (p *qaction) StatusTip() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QAction_StatusTip(p.Pointer()))
}

func (p *qaction) Text() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QAction_Text(p.Pointer()))
}

func (p *qaction) ToolTip() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QAction_ToolTip(p.Pointer()))
}

func (p *qaction) WhatsThis() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QAction_WhatsThis(p.Pointer()))
}

func (p *qaction) ConnectSlotHover() {
	C.QAction_ConnectSlotHover(p.Pointer())
}

func (p *qaction) DisconnectSlotHover() {
	C.QAction_DisconnectSlotHover(p.Pointer())
}

func (p *qaction) SlotHover() {
	if p.Pointer() != nil {
		C.QAction_Hover(p.Pointer())
	}
}

func (p *qaction) ConnectSlotSetChecked() {
	C.QAction_ConnectSlotSetChecked(p.Pointer())
}

func (p *qaction) DisconnectSlotSetChecked() {
	C.QAction_DisconnectSlotSetChecked(p.Pointer())
}

func (p *qaction) SlotSetChecked(bo bool) {
	if p.Pointer() != nil {
		C.QAction_SetChecked_Bool(p.Pointer(), goBoolToCInt(bo))
	}
}

func (p *qaction) ConnectSlotSetDisabled() {
	C.QAction_ConnectSlotSetDisabled(p.Pointer())
}

func (p *qaction) DisconnectSlotSetDisabled() {
	C.QAction_DisconnectSlotSetDisabled(p.Pointer())
}

func (p *qaction) SlotSetDisabled(bo bool) {
	if p.Pointer() != nil {
		C.QAction_SetDisabled_Bool(p.Pointer(), goBoolToCInt(bo))
	}
}

func (p *qaction) ConnectSlotSetEnabled() {
	C.QAction_ConnectSlotSetEnabled(p.Pointer())
}

func (p *qaction) DisconnectSlotSetEnabled() {
	C.QAction_DisconnectSlotSetEnabled(p.Pointer())
}

func (p *qaction) SlotSetEnabled(bo bool) {
	if p.Pointer() != nil {
		C.QAction_SetEnabled_Bool(p.Pointer(), goBoolToCInt(bo))
	}
}

func (p *qaction) ConnectSlotSetVisible() {
	C.QAction_ConnectSlotSetVisible(p.Pointer())
}

func (p *qaction) DisconnectSlotSetVisible() {
	C.QAction_DisconnectSlotSetVisible(p.Pointer())
}

func (p *qaction) SlotSetVisible(bo bool) {
	if p.Pointer() != nil {
		C.QAction_SetVisible_Bool(p.Pointer(), goBoolToCInt(bo))
	}
}

func (p *qaction) ConnectSlotToggle() {
	C.QAction_ConnectSlotToggle(p.Pointer())
}

func (p *qaction) DisconnectSlotToggle() {
	C.QAction_DisconnectSlotToggle(p.Pointer())
}

func (p *qaction) SlotToggle() {
	if p.Pointer() != nil {
		C.QAction_Toggle(p.Pointer())
	}
}

func (p *qaction) ConnectSlotTrigger() {
	C.QAction_ConnectSlotTrigger(p.Pointer())
}

func (p *qaction) DisconnectSlotTrigger() {
	C.QAction_DisconnectSlotTrigger(p.Pointer())
}

func (p *qaction) SlotTrigger() {
	if p.Pointer() != nil {
		C.QAction_Trigger(p.Pointer())
	}
}

func (p *qaction) ConnectSignalChanged(f func()) {
	C.QAction_ConnectSignalChanged(p.Pointer())
	connectSignal(p.ObjectName(), "changed", f)
}

func (p *qaction) DisconnectSignalChanged() {
	C.QAction_DisconnectSignalChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "changed")
}

func (p *qaction) SignalChanged() func() {
	return getSignal(p.ObjectName(), "changed")
}

func (p *qaction) ConnectSignalHovered(f func()) {
	C.QAction_ConnectSignalHovered(p.Pointer())
	connectSignal(p.ObjectName(), "hovered", f)
}

func (p *qaction) DisconnectSignalHovered() {
	C.QAction_DisconnectSignalHovered(p.Pointer())
	disconnectSignal(p.ObjectName(), "hovered")
}

func (p *qaction) SignalHovered() func() {
	return getSignal(p.ObjectName(), "hovered")
}

func (p *qaction) ConnectSignalToggled(f func()) {
	C.QAction_ConnectSignalToggled(p.Pointer())
	connectSignal(p.ObjectName(), "toggled", f)
}

func (p *qaction) DisconnectSignalToggled() {
	C.QAction_DisconnectSignalToggled(p.Pointer())
	disconnectSignal(p.ObjectName(), "toggled")
}

func (p *qaction) SignalToggled() func() {
	return getSignal(p.ObjectName(), "toggled")
}

func (p *qaction) ConnectSignalTriggered(f func()) {
	C.QAction_ConnectSignalTriggered(p.Pointer())
	connectSignal(p.ObjectName(), "triggered", f)
}

func (p *qaction) DisconnectSignalTriggered() {
	C.QAction_DisconnectSignalTriggered(p.Pointer())
	disconnectSignal(p.ObjectName(), "triggered")
}

func (p *qaction) SignalTriggered() func() {
	return getSignal(p.ObjectName(), "triggered")
}

//export callbackQAction
func callbackQAction(ptr C.QtObjectPtr, msg *C.char) {
	var qaction = new(qaction)
	qaction.SetPointer(ptr)
	getSignal(qaction.ObjectName(), C.GoString(msg))()
}
