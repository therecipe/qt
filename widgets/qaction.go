package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QAction struct {
	core.QObject
}

type QAction_ITF interface {
	core.QObject_ITF
	QAction_PTR() *QAction
}

func PointerFromQAction(ptr QAction_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAction_PTR().Pointer()
	}
	return nil
}

func NewQActionFromPointer(ptr unsafe.Pointer) *QAction {
	var n = new(QAction)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAction_") {
		n.SetObjectName("QAction_" + qt.Identifier())
	}
	return n
}

func (ptr *QAction) QAction_PTR() *QAction {
	return ptr
}

//QAction::ActionEvent
type QAction__ActionEvent int64

const (
	QAction__Trigger = QAction__ActionEvent(0)
	QAction__Hover   = QAction__ActionEvent(1)
)

//QAction::MenuRole
type QAction__MenuRole int64

const (
	QAction__NoRole                  = QAction__MenuRole(0)
	QAction__TextHeuristicRole       = QAction__MenuRole(1)
	QAction__ApplicationSpecificRole = QAction__MenuRole(2)
	QAction__AboutQtRole             = QAction__MenuRole(3)
	QAction__AboutRole               = QAction__MenuRole(4)
	QAction__PreferencesRole         = QAction__MenuRole(5)
	QAction__QuitRole                = QAction__MenuRole(6)
)

//QAction::Priority
type QAction__Priority int64

const (
	QAction__LowPriority    = QAction__Priority(0)
	QAction__NormalPriority = QAction__Priority(128)
	QAction__HighPriority   = QAction__Priority(256)
)

func (ptr *QAction) AutoRepeat() bool {
	defer qt.Recovering("QAction::autoRepeat")

	if ptr.Pointer() != nil {
		return C.QAction_AutoRepeat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) Icon() *gui.QIcon {
	defer qt.Recovering("QAction::icon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QAction_Icon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAction) IconText() string {
	defer qt.Recovering("QAction::iconText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_IconText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) IsCheckable() bool {
	defer qt.Recovering("QAction::isCheckable")

	if ptr.Pointer() != nil {
		return C.QAction_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) IsChecked() bool {
	defer qt.Recovering("QAction::isChecked")

	if ptr.Pointer() != nil {
		return C.QAction_IsChecked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) IsEnabled() bool {
	defer qt.Recovering("QAction::isEnabled")

	if ptr.Pointer() != nil {
		return C.QAction_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) IsIconVisibleInMenu() bool {
	defer qt.Recovering("QAction::isIconVisibleInMenu")

	if ptr.Pointer() != nil {
		return C.QAction_IsIconVisibleInMenu(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) IsVisible() bool {
	defer qt.Recovering("QAction::isVisible")

	if ptr.Pointer() != nil {
		return C.QAction_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) MenuRole() QAction__MenuRole {
	defer qt.Recovering("QAction::menuRole")

	if ptr.Pointer() != nil {
		return QAction__MenuRole(C.QAction_MenuRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAction) Priority() QAction__Priority {
	defer qt.Recovering("QAction::priority")

	if ptr.Pointer() != nil {
		return QAction__Priority(C.QAction_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAction) SetAutoRepeat(v bool) {
	defer qt.Recovering("QAction::setAutoRepeat")

	if ptr.Pointer() != nil {
		C.QAction_SetAutoRepeat(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetCheckable(v bool) {
	defer qt.Recovering("QAction::setCheckable")

	if ptr.Pointer() != nil {
		C.QAction_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetChecked(v bool) {
	defer qt.Recovering("QAction::setChecked")

	if ptr.Pointer() != nil {
		C.QAction_SetChecked(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetData(userData core.QVariant_ITF) {
	defer qt.Recovering("QAction::setData")

	if ptr.Pointer() != nil {
		C.QAction_SetData(ptr.Pointer(), core.PointerFromQVariant(userData))
	}
}

func (ptr *QAction) SetEnabled(v bool) {
	defer qt.Recovering("QAction::setEnabled")

	if ptr.Pointer() != nil {
		C.QAction_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetFont(font gui.QFont_ITF) {
	defer qt.Recovering("QAction::setFont")

	if ptr.Pointer() != nil {
		C.QAction_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QAction) SetIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QAction::setIcon")

	if ptr.Pointer() != nil {
		C.QAction_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QAction) SetIconText(text string) {
	defer qt.Recovering("QAction::setIconText")

	if ptr.Pointer() != nil {
		C.QAction_SetIconText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QAction) SetIconVisibleInMenu(visible bool) {
	defer qt.Recovering("QAction::setIconVisibleInMenu")

	if ptr.Pointer() != nil {
		C.QAction_SetIconVisibleInMenu(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QAction) SetMenuRole(menuRole QAction__MenuRole) {
	defer qt.Recovering("QAction::setMenuRole")

	if ptr.Pointer() != nil {
		C.QAction_SetMenuRole(ptr.Pointer(), C.int(menuRole))
	}
}

func (ptr *QAction) SetPriority(priority QAction__Priority) {
	defer qt.Recovering("QAction::setPriority")

	if ptr.Pointer() != nil {
		C.QAction_SetPriority(ptr.Pointer(), C.int(priority))
	}
}

func (ptr *QAction) SetShortcut(shortcut gui.QKeySequence_ITF) {
	defer qt.Recovering("QAction::setShortcut")

	if ptr.Pointer() != nil {
		C.QAction_SetShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(shortcut))
	}
}

func (ptr *QAction) SetShortcutContext(context core.Qt__ShortcutContext) {
	defer qt.Recovering("QAction::setShortcutContext")

	if ptr.Pointer() != nil {
		C.QAction_SetShortcutContext(ptr.Pointer(), C.int(context))
	}
}

func (ptr *QAction) SetStatusTip(statusTip string) {
	defer qt.Recovering("QAction::setStatusTip")

	if ptr.Pointer() != nil {
		C.QAction_SetStatusTip(ptr.Pointer(), C.CString(statusTip))
	}
}

func (ptr *QAction) SetText(text string) {
	defer qt.Recovering("QAction::setText")

	if ptr.Pointer() != nil {
		C.QAction_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QAction) SetToolTip(tip string) {
	defer qt.Recovering("QAction::setToolTip")

	if ptr.Pointer() != nil {
		C.QAction_SetToolTip(ptr.Pointer(), C.CString(tip))
	}
}

func (ptr *QAction) SetVisible(v bool) {
	defer qt.Recovering("QAction::setVisible")

	if ptr.Pointer() != nil {
		C.QAction_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetWhatsThis(what string) {
	defer qt.Recovering("QAction::setWhatsThis")

	if ptr.Pointer() != nil {
		C.QAction_SetWhatsThis(ptr.Pointer(), C.CString(what))
	}
}

func (ptr *QAction) ShortcutContext() core.Qt__ShortcutContext {
	defer qt.Recovering("QAction::shortcutContext")

	if ptr.Pointer() != nil {
		return core.Qt__ShortcutContext(C.QAction_ShortcutContext(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAction) StatusTip() string {
	defer qt.Recovering("QAction::statusTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) Text() string {
	defer qt.Recovering("QAction::text")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) Toggle() {
	defer qt.Recovering("QAction::toggle")

	if ptr.Pointer() != nil {
		C.QAction_Toggle(ptr.Pointer())
	}
}

func (ptr *QAction) ToolTip() string {
	defer qt.Recovering("QAction::toolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) WhatsThis() string {
	defer qt.Recovering("QAction::whatsThis")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func NewQAction(parent core.QObject_ITF) *QAction {
	defer qt.Recovering("QAction::QAction")

	return NewQActionFromPointer(C.QAction_NewQAction(core.PointerFromQObject(parent)))
}

func NewQAction3(icon gui.QIcon_ITF, text string, parent core.QObject_ITF) *QAction {
	defer qt.Recovering("QAction::QAction")

	return NewQActionFromPointer(C.QAction_NewQAction3(gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQObject(parent)))
}

func NewQAction2(text string, parent core.QObject_ITF) *QAction {
	defer qt.Recovering("QAction::QAction")

	return NewQActionFromPointer(C.QAction_NewQAction2(C.CString(text), core.PointerFromQObject(parent)))
}

func (ptr *QAction) ActionGroup() *QActionGroup {
	defer qt.Recovering("QAction::actionGroup")

	if ptr.Pointer() != nil {
		return NewQActionGroupFromPointer(C.QAction_ActionGroup(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAction) Activate(event QAction__ActionEvent) {
	defer qt.Recovering("QAction::activate")

	if ptr.Pointer() != nil {
		C.QAction_Activate(ptr.Pointer(), C.int(event))
	}
}

func (ptr *QAction) ConnectChanged(f func()) {
	defer qt.Recovering("connect QAction::changed")

	if ptr.Pointer() != nil {
		C.QAction_ConnectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "changed", f)
	}
}

func (ptr *QAction) DisconnectChanged() {
	defer qt.Recovering("disconnect QAction::changed")

	if ptr.Pointer() != nil {
		C.QAction_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "changed")
	}
}

//export callbackQActionChanged
func callbackQActionChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAction::changed")

	if signal := qt.GetSignal(C.GoString(ptrName), "changed"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAction) Changed() {
	defer qt.Recovering("QAction::changed")

	if ptr.Pointer() != nil {
		C.QAction_Changed(ptr.Pointer())
	}
}

func (ptr *QAction) Data() *core.QVariant {
	defer qt.Recovering("QAction::data")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAction_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAction) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QAction::event")

	if ptr.Pointer() != nil {
		return C.QAction_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QAction) Hover() {
	defer qt.Recovering("QAction::hover")

	if ptr.Pointer() != nil {
		C.QAction_Hover(ptr.Pointer())
	}
}

func (ptr *QAction) ConnectHovered(f func()) {
	defer qt.Recovering("connect QAction::hovered")

	if ptr.Pointer() != nil {
		C.QAction_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QAction) DisconnectHovered() {
	defer qt.Recovering("disconnect QAction::hovered")

	if ptr.Pointer() != nil {
		C.QAction_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQActionHovered
func callbackQActionHovered(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAction::hovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "hovered"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAction) Hovered() {
	defer qt.Recovering("QAction::hovered")

	if ptr.Pointer() != nil {
		C.QAction_Hovered(ptr.Pointer())
	}
}

func (ptr *QAction) IsSeparator() bool {
	defer qt.Recovering("QAction::isSeparator")

	if ptr.Pointer() != nil {
		return C.QAction_IsSeparator(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) Menu() *QMenu {
	defer qt.Recovering("QAction::menu")

	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QAction_Menu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAction) ParentWidget() *QWidget {
	defer qt.Recovering("QAction::parentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAction_ParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAction) SetActionGroup(group QActionGroup_ITF) {
	defer qt.Recovering("QAction::setActionGroup")

	if ptr.Pointer() != nil {
		C.QAction_SetActionGroup(ptr.Pointer(), PointerFromQActionGroup(group))
	}
}

func (ptr *QAction) SetDisabled(b bool) {
	defer qt.Recovering("QAction::setDisabled")

	if ptr.Pointer() != nil {
		C.QAction_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QAction) SetMenu(menu QMenu_ITF) {
	defer qt.Recovering("QAction::setMenu")

	if ptr.Pointer() != nil {
		C.QAction_SetMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QAction) SetSeparator(b bool) {
	defer qt.Recovering("QAction::setSeparator")

	if ptr.Pointer() != nil {
		C.QAction_SetSeparator(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QAction) SetShortcuts2(key gui.QKeySequence__StandardKey) {
	defer qt.Recovering("QAction::setShortcuts")

	if ptr.Pointer() != nil {
		C.QAction_SetShortcuts2(ptr.Pointer(), C.int(key))
	}
}

func (ptr *QAction) ShowStatusText(widget QWidget_ITF) bool {
	defer qt.Recovering("QAction::showStatusText")

	if ptr.Pointer() != nil {
		return C.QAction_ShowStatusText(ptr.Pointer(), PointerFromQWidget(widget)) != 0
	}
	return false
}

func (ptr *QAction) ConnectToggled(f func(checked bool)) {
	defer qt.Recovering("connect QAction::toggled")

	if ptr.Pointer() != nil {
		C.QAction_ConnectToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QAction) DisconnectToggled() {
	defer qt.Recovering("disconnect QAction::toggled")

	if ptr.Pointer() != nil {
		C.QAction_DisconnectToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQActionToggled
func callbackQActionToggled(ptr unsafe.Pointer, ptrName *C.char, checked C.int) {
	defer qt.Recovering("callback QAction::toggled")

	if signal := qt.GetSignal(C.GoString(ptrName), "toggled"); signal != nil {
		signal.(func(bool))(int(checked) != 0)
	}

}

func (ptr *QAction) Toggled(checked bool) {
	defer qt.Recovering("QAction::toggled")

	if ptr.Pointer() != nil {
		C.QAction_Toggled(ptr.Pointer(), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QAction) Trigger() {
	defer qt.Recovering("QAction::trigger")

	if ptr.Pointer() != nil {
		C.QAction_Trigger(ptr.Pointer())
	}
}

func (ptr *QAction) ConnectTriggered(f func(checked bool)) {
	defer qt.Recovering("connect QAction::triggered")

	if ptr.Pointer() != nil {
		C.QAction_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QAction) DisconnectTriggered() {
	defer qt.Recovering("disconnect QAction::triggered")

	if ptr.Pointer() != nil {
		C.QAction_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQActionTriggered
func callbackQActionTriggered(ptr unsafe.Pointer, ptrName *C.char, checked C.int) {
	defer qt.Recovering("callback QAction::triggered")

	if signal := qt.GetSignal(C.GoString(ptrName), "triggered"); signal != nil {
		signal.(func(bool))(int(checked) != 0)
	}

}

func (ptr *QAction) Triggered(checked bool) {
	defer qt.Recovering("QAction::triggered")

	if ptr.Pointer() != nil {
		C.QAction_Triggered(ptr.Pointer(), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QAction) DestroyQAction() {
	defer qt.Recovering("QAction::~QAction")

	if ptr.Pointer() != nil {
		C.QAction_DestroyQAction(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAction) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAction::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAction) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAction::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQActionTimerEvent
func callbackQActionTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAction::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQActionFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAction) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAction::timerEvent")

	if ptr.Pointer() != nil {
		C.QAction_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAction) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAction::timerEvent")

	if ptr.Pointer() != nil {
		C.QAction_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAction) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAction::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAction) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAction::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQActionChildEvent
func callbackQActionChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAction::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQActionFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAction) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAction::childEvent")

	if ptr.Pointer() != nil {
		C.QAction_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAction) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAction::childEvent")

	if ptr.Pointer() != nil {
		C.QAction_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAction) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAction::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAction) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAction::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQActionCustomEvent
func callbackQActionCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAction::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQActionFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAction) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAction::customEvent")

	if ptr.Pointer() != nil {
		C.QAction_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAction) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAction::customEvent")

	if ptr.Pointer() != nil {
		C.QAction_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
