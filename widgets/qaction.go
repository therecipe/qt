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

type QActionITF interface {
	core.QObjectITF
	QActionPTR() *QAction
}

func PointerFromQAction(ptr QActionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QActionPTR().Pointer()
	}
	return nil
}

func QActionFromPointer(ptr unsafe.Pointer) *QAction {
	var n = new(QAction)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QAction_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QAction) QActionPTR() *QAction {
	return ptr
}

//QAction::ActionEvent
type QAction__ActionEvent int

var (
	QAction__Trigger = QAction__ActionEvent(0)
	QAction__Hover   = QAction__ActionEvent(1)
)

//QAction::MenuRole
type QAction__MenuRole int

var (
	QAction__NoRole                  = QAction__MenuRole(0)
	QAction__TextHeuristicRole       = QAction__MenuRole(1)
	QAction__ApplicationSpecificRole = QAction__MenuRole(2)
	QAction__AboutQtRole             = QAction__MenuRole(3)
	QAction__AboutRole               = QAction__MenuRole(4)
	QAction__PreferencesRole         = QAction__MenuRole(5)
	QAction__QuitRole                = QAction__MenuRole(6)
)

//QAction::Priority
type QAction__Priority int

var (
	QAction__LowPriority    = QAction__Priority(0)
	QAction__NormalPriority = QAction__Priority(128)
	QAction__HighPriority   = QAction__Priority(256)
)

func (ptr *QAction) AutoRepeat() bool {
	if ptr.Pointer() != nil {
		return C.QAction_AutoRepeat(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAction) IconText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_IconText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAction) IsCheckable() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsCheckable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAction) IsChecked() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsChecked(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAction) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAction) IsIconVisibleInMenu() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsIconVisibleInMenu(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAction) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAction) MenuRole() QAction__MenuRole {
	if ptr.Pointer() != nil {
		return QAction__MenuRole(C.QAction_MenuRole(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAction) Priority() QAction__Priority {
	if ptr.Pointer() != nil {
		return QAction__Priority(C.QAction_Priority(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAction) SetAutoRepeat(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetAutoRepeat(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetCheckable(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetCheckable(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetChecked(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetChecked(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetData(userData string) {
	if ptr.Pointer() != nil {
		C.QAction_SetData(C.QtObjectPtr(ptr.Pointer()), C.CString(userData))
	}
}

func (ptr *QAction) SetEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetFont(font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QAction) SetIcon(icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QAction) SetIconText(text string) {
	if ptr.Pointer() != nil {
		C.QAction_SetIconText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QAction) SetIconVisibleInMenu(visible bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetIconVisibleInMenu(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QAction) SetMenuRole(menuRole QAction__MenuRole) {
	if ptr.Pointer() != nil {
		C.QAction_SetMenuRole(C.QtObjectPtr(ptr.Pointer()), C.int(menuRole))
	}
}

func (ptr *QAction) SetPriority(priority QAction__Priority) {
	if ptr.Pointer() != nil {
		C.QAction_SetPriority(C.QtObjectPtr(ptr.Pointer()), C.int(priority))
	}
}

func (ptr *QAction) SetShortcut(shortcut gui.QKeySequenceITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetShortcut(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQKeySequence(shortcut)))
	}
}

func (ptr *QAction) SetShortcutContext(context core.Qt__ShortcutContext) {
	if ptr.Pointer() != nil {
		C.QAction_SetShortcutContext(C.QtObjectPtr(ptr.Pointer()), C.int(context))
	}
}

func (ptr *QAction) SetStatusTip(statusTip string) {
	if ptr.Pointer() != nil {
		C.QAction_SetStatusTip(C.QtObjectPtr(ptr.Pointer()), C.CString(statusTip))
	}
}

func (ptr *QAction) SetText(text string) {
	if ptr.Pointer() != nil {
		C.QAction_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QAction) SetToolTip(tip string) {
	if ptr.Pointer() != nil {
		C.QAction_SetToolTip(C.QtObjectPtr(ptr.Pointer()), C.CString(tip))
	}
}

func (ptr *QAction) SetVisible(v bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QAction) SetWhatsThis(what string) {
	if ptr.Pointer() != nil {
		C.QAction_SetWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.CString(what))
	}
}

func (ptr *QAction) ShortcutContext() core.Qt__ShortcutContext {
	if ptr.Pointer() != nil {
		return core.Qt__ShortcutContext(C.QAction_ShortcutContext(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAction) StatusTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_StatusTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAction) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAction) Toggle() {
	if ptr.Pointer() != nil {
		C.QAction_Toggle(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAction) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_ToolTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAction) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_WhatsThis(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func NewQAction(parent core.QObjectITF) *QAction {
	return QActionFromPointer(unsafe.Pointer(C.QAction_NewQAction(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQAction3(icon gui.QIconITF, text string, parent core.QObjectITF) *QAction {
	return QActionFromPointer(unsafe.Pointer(C.QAction_NewQAction3(C.QtObjectPtr(gui.PointerFromQIcon(icon)), C.CString(text), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQAction2(text string, parent core.QObjectITF) *QAction {
	return QActionFromPointer(unsafe.Pointer(C.QAction_NewQAction2(C.CString(text), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QAction) ActionGroup() *QActionGroup {
	if ptr.Pointer() != nil {
		return QActionGroupFromPointer(unsafe.Pointer(C.QAction_ActionGroup(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAction) Activate(event QAction__ActionEvent) {
	if ptr.Pointer() != nil {
		C.QAction_Activate(C.QtObjectPtr(ptr.Pointer()), C.int(event))
	}
}

func (ptr *QAction) ConnectChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QAction_ConnectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "changed", f)
	}
}

func (ptr *QAction) DisconnectChanged() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "changed")
	}
}

//export callbackQActionChanged
func callbackQActionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "changed").(func())()
}

func (ptr *QAction) Data() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAction_Data(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAction) Hover() {
	if ptr.Pointer() != nil {
		C.QAction_Hover(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAction) ConnectHovered(f func()) {
	if ptr.Pointer() != nil {
		C.QAction_ConnectHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "hovered", f)
	}
}

func (ptr *QAction) DisconnectHovered() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "hovered")
	}
}

//export callbackQActionHovered
func callbackQActionHovered(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "hovered").(func())()
}

func (ptr *QAction) IsSeparator() bool {
	if ptr.Pointer() != nil {
		return C.QAction_IsSeparator(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QAction) Menu() *QMenu {
	if ptr.Pointer() != nil {
		return QMenuFromPointer(unsafe.Pointer(C.QAction_Menu(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAction) ParentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QAction_ParentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QAction) SetActionGroup(group QActionGroupITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetActionGroup(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQActionGroup(group)))
	}
}

func (ptr *QAction) SetDisabled(b bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetDisabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QAction) SetMenu(menu QMenuITF) {
	if ptr.Pointer() != nil {
		C.QAction_SetMenu(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQMenu(menu)))
	}
}

func (ptr *QAction) SetSeparator(b bool) {
	if ptr.Pointer() != nil {
		C.QAction_SetSeparator(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QAction) SetShortcuts2(key gui.QKeySequence__StandardKey) {
	if ptr.Pointer() != nil {
		C.QAction_SetShortcuts2(C.QtObjectPtr(ptr.Pointer()), C.int(key))
	}
}

func (ptr *QAction) ShowStatusText(widget QWidgetITF) bool {
	if ptr.Pointer() != nil {
		return C.QAction_ShowStatusText(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget))) != 0
	}
	return false
}

func (ptr *QAction) ConnectToggled(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QAction_ConnectToggled(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QAction) DisconnectToggled() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectToggled(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQActionToggled
func callbackQActionToggled(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "toggled").(func(bool))(int(checked) != 0)
}

func (ptr *QAction) Trigger() {
	if ptr.Pointer() != nil {
		C.QAction_Trigger(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QAction) ConnectTriggered(f func(checked bool)) {
	if ptr.Pointer() != nil {
		C.QAction_ConnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "triggered", f)
	}
}

func (ptr *QAction) DisconnectTriggered() {
	if ptr.Pointer() != nil {
		C.QAction_DisconnectTriggered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "triggered")
	}
}

//export callbackQActionTriggered
func callbackQActionTriggered(ptrName *C.char, checked C.int) {
	qt.GetSignal(C.GoString(ptrName), "triggered").(func(bool))(int(checked) != 0)
}

func (ptr *QAction) DestroyQAction() {
	if ptr.Pointer() != nil {
		C.QAction_DestroyQAction(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
