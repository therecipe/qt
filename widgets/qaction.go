package widgets

//#include "qaction.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QAction_" + qt.RandomIdentifier())
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
	if ptr.Pointer() != nil {
		return C.QAction_AutoRepeat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) IconText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_IconText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) IsCheckable() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) IsChecked() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsChecked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) IsIconVisibleInMenu() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsIconVisibleInMenu(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) MenuRole() QAction__MenuRole {
	if ptr.Pointer() != nil {
		return QAction__MenuRole(C.QAction_MenuRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAction) Priority() QAction__Priority {
	if ptr.Pointer() != nil {
		return QAction__Priority(C.QAction_Priority(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAction) SetAutoRepeat(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetAutoRepeat(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetCheckable(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetChecked(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetChecked(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetData(userData core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetData(ptr.Pointer(), core.PointerFromQVariant(userData))
	}
}

func (ptr *QAction) SetEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QAction) SetIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QAction) SetIconText(text string) {
	if ptr.Pointer() != nil {
		C.QAction_SetIconText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QAction) SetIconVisibleInMenu(visible bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetIconVisibleInMenu(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QAction) SetMenuRole(menuRole QAction__MenuRole) {
	if ptr.Pointer() != nil {
		C.QAction_SetMenuRole(ptr.Pointer(), C.int(menuRole))
	}
}

func (ptr *QAction) SetPriority(priority QAction__Priority) {
	if ptr.Pointer() != nil {
		C.QAction_SetPriority(ptr.Pointer(), C.int(priority))
	}
}

func (ptr *QAction) SetShortcut(shortcut gui.QKeySequence_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(shortcut))
	}
}

func (ptr *QAction) SetShortcutContext(context core.Qt__ShortcutContext) {
	if ptr.Pointer() != nil {
		C.QAction_SetShortcutContext(ptr.Pointer(), C.int(context))
	}
}

func (ptr *QAction) SetStatusTip(statusTip string) {
	if ptr.Pointer() != nil {
		C.QAction_SetStatusTip(ptr.Pointer(), C.CString(statusTip))
	}
}

func (ptr *QAction) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QAction_SetText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QAction) SetToolTip(tip string) {
	if ptr.Pointer() != nil {
		C.QAction_SetToolTip(ptr.Pointer(), C.CString(tip))
	}
}

func (ptr *QAction) SetVisible(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetWhatsThis(what string) {
	if ptr.Pointer() != nil {
		C.QAction_SetWhatsThis(ptr.Pointer(), C.CString(what))
	}
}

func (ptr *QAction) ShortcutContext() core.Qt__ShortcutContext {
	if ptr.Pointer() != nil {
		return core.Qt__ShortcutContext(C.QAction_ShortcutContext(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAction) StatusTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) Toggle() {
	if ptr.Pointer() != nil {
		C.QAction_Toggle(ptr.Pointer())
	}
}

func (ptr *QAction) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAction) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func NewQAction(parent core.QObject_ITF) *QAction {
	return NewQActionFromPointer(C.QAction_NewQAction(core.PointerFromQObject(parent)))
}

func NewQAction3(icon gui.QIcon_ITF, text string, parent core.QObject_ITF) *QAction {
	return NewQActionFromPointer(C.QAction_NewQAction3(gui.PointerFromQIcon(icon), C.CString(text), core.PointerFromQObject(parent)))
}

func NewQAction2(text string, parent core.QObject_ITF) *QAction {
	return NewQActionFromPointer(C.QAction_NewQAction2(C.CString(text), core.PointerFromQObject(parent)))
}

func (ptr *QAction) ActionGroup() *QActionGroup {
	if ptr.Pointer() != nil {
		return NewQActionGroupFromPointer(C.QAction_ActionGroup(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAction) Activate(event QAction__ActionEvent) {
	if ptr.Pointer() != nil {
		C.QAction_Activate(ptr.Pointer(), C.int(event))
	}
}

func (ptr *QAction) ConnectChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAction_ConnectChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "changed", f)
	}
}

func (ptr *QAction) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "changed")
	}
}

//export callbackQActionChanged
func callbackQActionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "changed").(func())()
}

func (ptr *QAction) Data() *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QAction_Data(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAction) Hover() {
	if ptr.Pointer() != nil {
		C.QAction_Hover(ptr.Pointer())
	}
}

func (ptr *QAction) ConnectHovered(f func()) {
	if ptr.Pointer() != nil {
		C.QAction_ConnectHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QAction) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQActionHovered
func callbackQActionHovered(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "hovered").(func())()
}

func (ptr *QAction) IsSeparator() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsSeparator(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAction) Menu() *QMenu {
	if ptr.Pointer() != nil {
		return NewQMenuFromPointer(C.QAction_Menu(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAction) ParentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QAction_ParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAction) SetActionGroup(group QActionGroup_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetActionGroup(ptr.Pointer(), PointerFromQActionGroup(group))
	}
}

func (ptr *QAction) SetDisabled(b bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QAction) SetMenu(menu QMenu_ITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetMenu(ptr.Pointer(), PointerFromQMenu(menu))
	}
}

func (ptr *QAction) SetSeparator(b bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetSeparator(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QAction) SetShortcuts2(key gui.QKeySequence__StandardKey) {
	if ptr.Pointer() != nil {
		C.QAction_SetShortcuts2(ptr.Pointer(), C.int(key))
	}
}

func (ptr *QAction) ShowStatusText(widget QWidget_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QAction_ShowStatusText(ptr.Pointer(), PointerFromQWidget(widget)) != 0
	}
	return false
}

func (ptr *QAction) ConnectToggled(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QAction_ConnectToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QAction) DisconnectToggled() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQActionToggled
func callbackQActionToggled(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "toggled").(func(bool))(int(checked) != 0)
}

func (ptr *QAction) Trigger() {
	if ptr.Pointer() != nil {
		C.QAction_Trigger(ptr.Pointer())
	}
}

func (ptr *QAction) ConnectTriggered(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QAction_ConnectTriggered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QAction) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectTriggered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQActionTriggered
func callbackQActionTriggered(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func(bool))(int(checked) != 0)
}

func (ptr *QAction) DestroyQAction() {
	if ptr.Pointer() != nil {
		C.QAction_DestroyQAction(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
