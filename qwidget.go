package qt

//#include "qwidget.h"
import "C"

type qwidget struct {
	qobject
}

type QWidget interface {
	QObject
	AcceptDrops() bool
	AccessibleDescription() string
	AccessibleName() string
	ActivateWindow()
	AdjustSize()
	AutoFillBackground() bool
	ChildAt(x int, y int) QWidget
	ClearFocus()
	ClearMask()
	ContextMenuPolicy() ContextMenuPolicy
	EnsurePolished()
	FocusPolicy() FocusPolicy
	FocusProxy() QWidget
	FocusWidget() QWidget
	GrabKeyboard()
	GrabMouse()
	HasFocus() bool
	HasHeightForWidth() bool
	HasMouseTracking() bool
	Height() int
	HeightForWidth(w int) int
	InputMethodHints() InputMethodHint
	IsActiveWindow() bool
	IsAncestorOf(child QWidget) bool
	IsEnabled() bool
	IsEnabledTo(ancestor QWidget) bool
	IsFullScreen() bool
	IsHidden() bool
	IsMaximized() bool
	IsMinimized() bool
	IsModal() bool
	IsVisible() bool
	IsVisibleTo(ancestor QWidget) bool
	IsWindow() bool
	IsWindowModified() bool
	Layout() QLayout
	LayoutDirection() LayoutDirection
	MaximumHeight() int
	MaximumWidth() int
	MinimumHeight() int
	MinimumWidth() int
	Move(x int, y int)
	NativeParentWidget() QWidget
	NextInFocusChain() QWidget
	OverrideWindowFlags(flags WindowType)
	ParentWidget() QWidget
	PreviousInFocusChain() QWidget
	ReleaseKeyboard()
	ReleaseMouse()
	ReleaseShortcut(id int)
	Repaint(x int, y int, w int, h int)
	Resize(w int, h int)
	Scroll(dx int, dy int)
	SetAcceptDrops(on bool)
	SetAccessibleDescription(description string)
	SetAccessibleName(name string)
	SetAttribute(attribute WidgetAttribute, on bool)
	SetAutoFillBackground(enabled bool)
	SetBaseSize(basew int, baseh int)
	SetContentsMargins(left int, top int, right int, bottom int)
	SetContextMenuPolicy(policy ContextMenuPolicy)
	SetFixedHeight(h int)
	SetFixedSize(w int, h int)
	SetFixedWidth(w int)
	SetFocus(reason FocusReason)
	SetFocusPolicy(policy FocusPolicy)
	SetFocusProxy(w QWidget)
	SetGeometry(x int, y int, w int, h int)
	SetLayout(layout QLayout)
	SetLayoutDirection(direction LayoutDirection)
	SetMaximumHeight(maxh int)
	SetMaximumSize(maxw int, maxh int)
	SetMaximumWidth(maxw int)
	SetMinimumHeight(minh int)
	SetMinimumSize(minw int, minh int)
	SetMinimumWidth(minw int)
	SetMouseTracking(enable bool)
	SetShortcutAutoRepeat(id int, enable bool)
	SetShortcutEnabled(id int, enable bool)
	SetSizeIncrement(w int, h int)
	SetStatusTip(statusTip string)
	SetToolTip(toolTip string)
	SetToolTipDuration(msec int)
	SetUpdatesEnabled(enable bool)
	SetWhatsThis(whatsThis string)
	SetWindowFilePath(filePath string)
	SetWindowFlags(typ WindowType)
	SetWindowIconText(iconText string)
	SetWindowModality(windowModality WindowModality)
	SetWindowRole(role string)
	SetWindowState(windowState WindowState)
	StackUnder(w QWidget)
	StatusTip() string
	StyleSheet() string
	TestAttribute(attribute WidgetAttribute) bool
	ToolTip() string
	ToolTipDuration() int
	UnderMouse() bool
	UngrabGesture(gesture GestureType)
	UnsetCursor()
	UnsetLayoutDirection()
	UnsetLocale()
	Update(x int, y int, w int, h int)
	UpdateGeometry()
	UpdatesEnabled() bool
	Width() int
	Window() QWidget
	WindowFilePath() string
	WindowFlags() WindowType
	WindowIconText() string
	WindowModality() WindowModality
	WindowRole() string
	WindowState() WindowState
	WindowTitle() string
	WindowType() WindowType
	X() int
	Y() int
	ConnectSlotClose()
	DisconnectSlotClose()
	SlotClose()
	ConnectSlotHide()
	DisconnectSlotHide()
	SlotHide()
	ConnectSlotLower()
	DisconnectSlotLower()
	SlotLower()
	ConnectSlotRaise()
	DisconnectSlotRaise()
	SlotRaise()
	ConnectSlotSetDisabled()
	DisconnectSlotSetDisabled()
	SlotSetDisabled(disable bool)
	ConnectSlotSetEnabled()
	DisconnectSlotSetEnabled()
	SlotSetEnabled(enabled bool)
	ConnectSlotSetHidden()
	DisconnectSlotSetHidden()
	SlotSetHidden(hidden bool)
	ConnectSlotSetStyleSheet()
	DisconnectSlotSetStyleSheet()
	SlotSetStyleSheet(styleSheet string)
	ConnectSlotSetWindowModified()
	DisconnectSlotSetWindowModified()
	SlotSetWindowModified(modified bool)
	ConnectSlotSetWindowTitle()
	DisconnectSlotSetWindowTitle()
	SlotSetWindowTitle(windowTitle string)
	ConnectSlotShow()
	DisconnectSlotShow()
	SlotShow()
	ConnectSlotShowFullScreen()
	DisconnectSlotShowFullScreen()
	SlotShowFullScreen()
	ConnectSlotShowMaximized()
	DisconnectSlotShowMaximized()
	SlotShowMaximized()
	ConnectSlotShowMinimized()
	DisconnectSlotShowMinimized()
	SlotShowMinimized()
	ConnectSlotShowNormal()
	DisconnectSlotShowNormal()
	SlotShowNormal()
	ConnectSignalWindowIconTextChanged(f func())
	DisconnectSignalWindowIconTextChanged()
	SignalWindowIconTextChanged() func()
	ConnectSignalWindowTitleChanged(f func())
	DisconnectSignalWindowTitleChanged()
	SignalWindowTitleChanged() func()
}

func (p *qwidget) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qwidget) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQWidget(parent QWidget, f WindowType) QWidget {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qwidget = new(qwidget)
	qwidget.SetPointer(C.QWidget_New_QWidget_WindowType(parentPtr, C.int(f)))
	qwidget.SetObjectName("QWidget_" + randomIdentifier())
	return qwidget
}

func (p *qwidget) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QWidget_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qwidget) AcceptDrops() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_AcceptDrops(p.Pointer()) != 0
}

func (p *qwidget) AccessibleDescription() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QWidget_AccessibleDescription(p.Pointer()))
}

func (p *qwidget) AccessibleName() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QWidget_AccessibleName(p.Pointer()))
}

func (p *qwidget) ActivateWindow() {
	if p.Pointer() != nil {
		C.QWidget_ActivateWindow(p.Pointer())
	}
}

func (p *qwidget) AdjustSize() {
	if p.Pointer() != nil {
		C.QWidget_AdjustSize(p.Pointer())
	}
}

func (p *qwidget) AutoFillBackground() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_AutoFillBackground(p.Pointer()) != 0
}

func (p *qwidget) ChildAt(x int, y int) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QWidget_ChildAt_Int_Int(p.Pointer(), C.int(x), C.int(y)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qwidget) ClearFocus() {
	if p.Pointer() != nil {
		C.QWidget_ClearFocus(p.Pointer())
	}
}

func (p *qwidget) ClearMask() {
	if p.Pointer() != nil {
		C.QWidget_ClearMask(p.Pointer())
	}
}

func (p *qwidget) ContextMenuPolicy() ContextMenuPolicy {
	if p.Pointer() == nil {
		return 0
	}
	return ContextMenuPolicy(C.QWidget_ContextMenuPolicy(p.Pointer()))
}

func (p *qwidget) EnsurePolished() {
	if p.Pointer() != nil {
		C.QWidget_EnsurePolished(p.Pointer())
	}
}

func (p *qwidget) FocusPolicy() FocusPolicy {
	if p.Pointer() == nil {
		return 0
	}
	return FocusPolicy(C.QWidget_FocusPolicy(p.Pointer()))
}

func (p *qwidget) FocusProxy() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QWidget_FocusProxy(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qwidget) FocusWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QWidget_FocusWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qwidget) GrabKeyboard() {
	if p.Pointer() != nil {
		C.QWidget_GrabKeyboard(p.Pointer())
	}
}

func (p *qwidget) GrabMouse() {
	if p.Pointer() != nil {
		C.QWidget_GrabMouse(p.Pointer())
	}
}

func (p *qwidget) HasFocus() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_HasFocus(p.Pointer()) != 0
}

func (p *qwidget) HasHeightForWidth() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_HasHeightForWidth(p.Pointer()) != 0
}

func (p *qwidget) HasMouseTracking() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_HasMouseTracking(p.Pointer()) != 0
}

func (p *qwidget) Height() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_Height(p.Pointer()))
}

func (p *qwidget) HeightForWidth(w int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_HeightForWidth_Int(p.Pointer(), C.int(w)))
}

func (p *qwidget) InputMethodHints() InputMethodHint {
	if p.Pointer() == nil {
		return 0
	}
	return InputMethodHint(C.QWidget_InputMethodHints(p.Pointer()))
}

func (p *qwidget) IsActiveWindow() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsActiveWindow(p.Pointer()) != 0
}

func (p *qwidget) IsAncestorOf(child QWidget) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var childPtr C.QtObjectPtr
		if child != nil {
			childPtr = child.Pointer()
		}
		return C.QWidget_IsAncestorOf_QWidget(p.Pointer(), childPtr) != 0
	}
}

func (p *qwidget) IsEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsEnabled(p.Pointer()) != 0
}

func (p *qwidget) IsEnabledTo(ancestor QWidget) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var ancestorPtr C.QtObjectPtr
		if ancestor != nil {
			ancestorPtr = ancestor.Pointer()
		}
		return C.QWidget_IsEnabledTo_QWidget(p.Pointer(), ancestorPtr) != 0
	}
}

func (p *qwidget) IsFullScreen() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsFullScreen(p.Pointer()) != 0
}

func (p *qwidget) IsHidden() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsHidden(p.Pointer()) != 0
}

func (p *qwidget) IsMaximized() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsMaximized(p.Pointer()) != 0
}

func (p *qwidget) IsMinimized() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsMinimized(p.Pointer()) != 0
}

func (p *qwidget) IsModal() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsModal(p.Pointer()) != 0
}

func (p *qwidget) IsVisible() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsVisible(p.Pointer()) != 0
}

func (p *qwidget) IsVisibleTo(ancestor QWidget) bool {
	if p.Pointer() == nil {
		return false
	} else {
		var ancestorPtr C.QtObjectPtr
		if ancestor != nil {
			ancestorPtr = ancestor.Pointer()
		}
		return C.QWidget_IsVisibleTo_QWidget(p.Pointer(), ancestorPtr) != 0
	}
}

func (p *qwidget) IsWindow() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsWindow(p.Pointer()) != 0
}

func (p *qwidget) IsWindowModified() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_IsWindowModified(p.Pointer()) != 0
}

func (p *qwidget) Layout() QLayout {
	if p.Pointer() == nil {
		return nil
	} else {
		var qlayout = new(qlayout)
		qlayout.SetPointer(C.QWidget_Layout(p.Pointer()))
		if qlayout.ObjectName() == "" {
			qlayout.SetObjectName("QLayout_" + randomIdentifier())
		}
		return qlayout
	}
}

func (p *qwidget) LayoutDirection() LayoutDirection {
	if p.Pointer() == nil {
		return 0
	}
	return LayoutDirection(C.QWidget_LayoutDirection(p.Pointer()))
}

func (p *qwidget) MaximumHeight() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_MaximumHeight(p.Pointer()))
}

func (p *qwidget) MaximumWidth() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_MaximumWidth(p.Pointer()))
}

func (p *qwidget) MinimumHeight() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_MinimumHeight(p.Pointer()))
}

func (p *qwidget) MinimumWidth() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_MinimumWidth(p.Pointer()))
}

func (p *qwidget) Move(x int, y int) {
	if p.Pointer() != nil {
		C.QWidget_Move_Int_Int(p.Pointer(), C.int(x), C.int(y))
	}
}

func (p *qwidget) NativeParentWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QWidget_NativeParentWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qwidget) NextInFocusChain() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QWidget_NextInFocusChain(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qwidget) OverrideWindowFlags(flags WindowType) {
	if p.Pointer() != nil {
		C.QWidget_OverrideWindowFlags_WindowType(p.Pointer(), C.int(flags))
	}
}

func (p *qwidget) ParentWidget() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QWidget_ParentWidget(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qwidget) PreviousInFocusChain() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QWidget_PreviousInFocusChain(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qwidget) ReleaseKeyboard() {
	if p.Pointer() != nil {
		C.QWidget_ReleaseKeyboard(p.Pointer())
	}
}

func (p *qwidget) ReleaseMouse() {
	if p.Pointer() != nil {
		C.QWidget_ReleaseMouse(p.Pointer())
	}
}

func (p *qwidget) ReleaseShortcut(id int) {
	if p.Pointer() != nil {
		C.QWidget_ReleaseShortcut_Int(p.Pointer(), C.int(id))
	}
}

func (p *qwidget) Repaint(x int, y int, w int, h int) {
	if p.Pointer() != nil {
		C.QWidget_Repaint_Int_Int_Int_Int(p.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (p *qwidget) Resize(w int, h int) {
	if p.Pointer() != nil {
		C.QWidget_Resize_Int_Int(p.Pointer(), C.int(w), C.int(h))
	}
}

func (p *qwidget) Scroll(dx int, dy int) {
	if p.Pointer() != nil {
		C.QWidget_Scroll_Int_Int(p.Pointer(), C.int(dx), C.int(dy))
	}
}

func (p *qwidget) SetAcceptDrops(on bool) {
	if p.Pointer() != nil {
		C.QWidget_SetAcceptDrops_Bool(p.Pointer(), goBoolToCInt(on))
	}
}

func (p *qwidget) SetAccessibleDescription(description string) {
	if p.Pointer() != nil {
		C.QWidget_SetAccessibleDescription_String(p.Pointer(), C.CString(description))
	}
}

func (p *qwidget) SetAccessibleName(name string) {
	if p.Pointer() != nil {
		C.QWidget_SetAccessibleName_String(p.Pointer(), C.CString(name))
	}
}

func (p *qwidget) SetAttribute(attribute WidgetAttribute, on bool) {
	if p.Pointer() != nil {
		C.QWidget_SetAttribute_WidgetAttribute_Bool(p.Pointer(), C.int(attribute), goBoolToCInt(on))
	}
}

func (p *qwidget) SetAutoFillBackground(enabled bool) {
	if p.Pointer() != nil {
		C.QWidget_SetAutoFillBackground_Bool(p.Pointer(), goBoolToCInt(enabled))
	}
}

func (p *qwidget) SetBaseSize(basew int, baseh int) {
	if p.Pointer() != nil {
		C.QWidget_SetBaseSize_Int_Int(p.Pointer(), C.int(basew), C.int(baseh))
	}
}

func (p *qwidget) SetContentsMargins(left int, top int, right int, bottom int) {
	if p.Pointer() != nil {
		C.QWidget_SetContentsMargins_Int_Int_Int_Int(p.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (p *qwidget) SetContextMenuPolicy(policy ContextMenuPolicy) {
	if p.Pointer() != nil {
		C.QWidget_SetContextMenuPolicy_ContextMenuPolicy(p.Pointer(), C.int(policy))
	}
}

func (p *qwidget) SetFixedHeight(h int) {
	if p.Pointer() != nil {
		C.QWidget_SetFixedHeight_Int(p.Pointer(), C.int(h))
	}
}

func (p *qwidget) SetFixedSize(w int, h int) {
	if p.Pointer() != nil {
		C.QWidget_SetFixedSize_Int_Int(p.Pointer(), C.int(w), C.int(h))
	}
}

func (p *qwidget) SetFixedWidth(w int) {
	if p.Pointer() != nil {
		C.QWidget_SetFixedWidth_Int(p.Pointer(), C.int(w))
	}
}

func (p *qwidget) SetFocus(reason FocusReason) {
	if p.Pointer() != nil {
		C.QWidget_SetFocus_FocusReason(p.Pointer(), C.int(reason))
	}
}

func (p *qwidget) SetFocusPolicy(policy FocusPolicy) {
	if p.Pointer() != nil {
		C.QWidget_SetFocusPolicy_FocusPolicy(p.Pointer(), C.int(policy))
	}
}

func (p *qwidget) SetFocusProxy(w QWidget) {
	if p.Pointer() != nil {
		var wPtr C.QtObjectPtr
		if w != nil {
			wPtr = w.Pointer()
		}
		C.QWidget_SetFocusProxy_QWidget(p.Pointer(), wPtr)
	}
}

func (p *qwidget) SetGeometry(x int, y int, w int, h int) {
	if p.Pointer() != nil {
		C.QWidget_SetGeometry_Int_Int_Int_Int(p.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (p *qwidget) SetLayout(layout QLayout) {
	if p.Pointer() != nil {
		var layoutPtr C.QtObjectPtr
		if layout != nil {
			layoutPtr = layout.Pointer()
		}
		C.QWidget_SetLayout_QLayout(p.Pointer(), layoutPtr)
	}
}

func (p *qwidget) SetLayoutDirection(direction LayoutDirection) {
	if p.Pointer() != nil {
		C.QWidget_SetLayoutDirection_LayoutDirection(p.Pointer(), C.int(direction))
	}
}

func (p *qwidget) SetMaximumHeight(maxh int) {
	if p.Pointer() != nil {
		C.QWidget_SetMaximumHeight_Int(p.Pointer(), C.int(maxh))
	}
}

func (p *qwidget) SetMaximumSize(maxw int, maxh int) {
	if p.Pointer() != nil {
		C.QWidget_SetMaximumSize_Int_Int(p.Pointer(), C.int(maxw), C.int(maxh))
	}
}

func (p *qwidget) SetMaximumWidth(maxw int) {
	if p.Pointer() != nil {
		C.QWidget_SetMaximumWidth_Int(p.Pointer(), C.int(maxw))
	}
}

func (p *qwidget) SetMinimumHeight(minh int) {
	if p.Pointer() != nil {
		C.QWidget_SetMinimumHeight_Int(p.Pointer(), C.int(minh))
	}
}

func (p *qwidget) SetMinimumSize(minw int, minh int) {
	if p.Pointer() != nil {
		C.QWidget_SetMinimumSize_Int_Int(p.Pointer(), C.int(minw), C.int(minh))
	}
}

func (p *qwidget) SetMinimumWidth(minw int) {
	if p.Pointer() != nil {
		C.QWidget_SetMinimumWidth_Int(p.Pointer(), C.int(minw))
	}
}

func (p *qwidget) SetMouseTracking(enable bool) {
	if p.Pointer() != nil {
		C.QWidget_SetMouseTracking_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qwidget) SetShortcutAutoRepeat(id int, enable bool) {
	if p.Pointer() != nil {
		C.QWidget_SetShortcutAutoRepeat_Int_Bool(p.Pointer(), C.int(id), goBoolToCInt(enable))
	}
}

func (p *qwidget) SetShortcutEnabled(id int, enable bool) {
	if p.Pointer() != nil {
		C.QWidget_SetShortcutEnabled_Int_Bool(p.Pointer(), C.int(id), goBoolToCInt(enable))
	}
}

func (p *qwidget) SetSizeIncrement(w int, h int) {
	if p.Pointer() != nil {
		C.QWidget_SetSizeIncrement_Int_Int(p.Pointer(), C.int(w), C.int(h))
	}
}

func (p *qwidget) SetStatusTip(statusTip string) {
	if p.Pointer() != nil {
		C.QWidget_SetStatusTip_String(p.Pointer(), C.CString(statusTip))
	}
}

func (p *qwidget) SetToolTip(toolTip string) {
	if p.Pointer() != nil {
		C.QWidget_SetToolTip_String(p.Pointer(), C.CString(toolTip))
	}
}

func (p *qwidget) SetToolTipDuration(msec int) {
	if p.Pointer() != nil {
		C.QWidget_SetToolTipDuration_Int(p.Pointer(), C.int(msec))
	}
}

func (p *qwidget) SetUpdatesEnabled(enable bool) {
	if p.Pointer() != nil {
		C.QWidget_SetUpdatesEnabled_Bool(p.Pointer(), goBoolToCInt(enable))
	}
}

func (p *qwidget) SetWhatsThis(whatsThis string) {
	if p.Pointer() != nil {
		C.QWidget_SetWhatsThis_String(p.Pointer(), C.CString(whatsThis))
	}
}

func (p *qwidget) SetWindowFilePath(filePath string) {
	if p.Pointer() != nil {
		C.QWidget_SetWindowFilePath_String(p.Pointer(), C.CString(filePath))
	}
}

func (p *qwidget) SetWindowFlags(typ WindowType) {
	if p.Pointer() != nil {
		C.QWidget_SetWindowFlags_WindowType(p.Pointer(), C.int(typ))
	}
}

func (p *qwidget) SetWindowIconText(iconText string) {
	if p.Pointer() != nil {
		C.QWidget_SetWindowIconText_String(p.Pointer(), C.CString(iconText))
	}
}

func (p *qwidget) SetWindowModality(windowModality WindowModality) {
	if p.Pointer() != nil {
		C.QWidget_SetWindowModality_WindowModality(p.Pointer(), C.int(windowModality))
	}
}

func (p *qwidget) SetWindowRole(role string) {
	if p.Pointer() != nil {
		C.QWidget_SetWindowRole_String(p.Pointer(), C.CString(role))
	}
}

func (p *qwidget) SetWindowState(windowState WindowState) {
	if p.Pointer() != nil {
		C.QWidget_SetWindowState_WindowState(p.Pointer(), C.int(windowState))
	}
}

func (p *qwidget) StackUnder(w QWidget) {
	if p.Pointer() != nil {
		var wPtr C.QtObjectPtr
		if w != nil {
			wPtr = w.Pointer()
		}
		C.QWidget_StackUnder_QWidget(p.Pointer(), wPtr)
	}
}

func (p *qwidget) StatusTip() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QWidget_StatusTip(p.Pointer()))
}

func (p *qwidget) StyleSheet() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QWidget_StyleSheet(p.Pointer()))
}

func (p *qwidget) TestAttribute(attribute WidgetAttribute) bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_TestAttribute_WidgetAttribute(p.Pointer(), C.int(attribute)) != 0
}

func (p *qwidget) ToolTip() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QWidget_ToolTip(p.Pointer()))
}

func (p *qwidget) ToolTipDuration() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_ToolTipDuration(p.Pointer()))
}

func (p *qwidget) UnderMouse() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_UnderMouse(p.Pointer()) != 0
}

func (p *qwidget) UngrabGesture(gesture GestureType) {
	if p.Pointer() != nil {
		C.QWidget_UngrabGesture_GestureType(p.Pointer(), C.int(gesture))
	}
}

func (p *qwidget) UnsetCursor() {
	if p.Pointer() != nil {
		C.QWidget_UnsetCursor(p.Pointer())
	}
}

func (p *qwidget) UnsetLayoutDirection() {
	if p.Pointer() != nil {
		C.QWidget_UnsetLayoutDirection(p.Pointer())
	}
}

func (p *qwidget) UnsetLocale() {
	if p.Pointer() != nil {
		C.QWidget_UnsetLocale(p.Pointer())
	}
}

func (p *qwidget) Update(x int, y int, w int, h int) {
	if p.Pointer() != nil {
		C.QWidget_Update_Int_Int_Int_Int(p.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (p *qwidget) UpdateGeometry() {
	if p.Pointer() != nil {
		C.QWidget_UpdateGeometry(p.Pointer())
	}
}

func (p *qwidget) UpdatesEnabled() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QWidget_UpdatesEnabled(p.Pointer()) != 0
}

func (p *qwidget) Width() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_Width(p.Pointer()))
}

func (p *qwidget) Window() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QWidget_Window(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qwidget) WindowFilePath() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QWidget_WindowFilePath(p.Pointer()))
}

func (p *qwidget) WindowFlags() WindowType {
	if p.Pointer() == nil {
		return 0
	}
	return WindowType(C.QWidget_WindowFlags(p.Pointer()))
}

func (p *qwidget) WindowIconText() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QWidget_WindowIconText(p.Pointer()))
}

func (p *qwidget) WindowModality() WindowModality {
	if p.Pointer() == nil {
		return 0
	}
	return WindowModality(C.QWidget_WindowModality(p.Pointer()))
}

func (p *qwidget) WindowRole() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QWidget_WindowRole(p.Pointer()))
}

func (p *qwidget) WindowState() WindowState {
	if p.Pointer() == nil {
		return 0
	}
	return WindowState(C.QWidget_WindowState(p.Pointer()))
}

func (p *qwidget) WindowTitle() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QWidget_WindowTitle(p.Pointer()))
}

func (p *qwidget) WindowType() WindowType {
	if p.Pointer() == nil {
		return 0
	}
	return WindowType(C.QWidget_WindowType(p.Pointer()))
}

func (p *qwidget) X() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_X(p.Pointer()))
}

func (p *qwidget) Y() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QWidget_Y(p.Pointer()))
}

func (p *qwidget) ConnectSlotClose() {
	C.QWidget_ConnectSlotClose(p.Pointer())
}

func (p *qwidget) DisconnectSlotClose() {
	C.QWidget_DisconnectSlotClose(p.Pointer())
}

func (p *qwidget) SlotClose() {
	if p.Pointer() != nil {
		C.QWidget_Close(p.Pointer())
	}
}

func (p *qwidget) ConnectSlotHide() {
	C.QWidget_ConnectSlotHide(p.Pointer())
}

func (p *qwidget) DisconnectSlotHide() {
	C.QWidget_DisconnectSlotHide(p.Pointer())
}

func (p *qwidget) SlotHide() {
	if p.Pointer() != nil {
		C.QWidget_Hide(p.Pointer())
	}
}

func (p *qwidget) ConnectSlotLower() {
	C.QWidget_ConnectSlotLower(p.Pointer())
}

func (p *qwidget) DisconnectSlotLower() {
	C.QWidget_DisconnectSlotLower(p.Pointer())
}

func (p *qwidget) SlotLower() {
	if p.Pointer() != nil {
		C.QWidget_Lower(p.Pointer())
	}
}

func (p *qwidget) ConnectSlotRaise() {
	C.QWidget_ConnectSlotRaise(p.Pointer())
}

func (p *qwidget) DisconnectSlotRaise() {
	C.QWidget_DisconnectSlotRaise(p.Pointer())
}

func (p *qwidget) SlotRaise() {
	if p.Pointer() != nil {
		C.QWidget_Raise(p.Pointer())
	}
}

func (p *qwidget) ConnectSlotSetDisabled() {
	C.QWidget_ConnectSlotSetDisabled(p.Pointer())
}

func (p *qwidget) DisconnectSlotSetDisabled() {
	C.QWidget_DisconnectSlotSetDisabled(p.Pointer())
}

func (p *qwidget) SlotSetDisabled(disable bool) {
	if p.Pointer() != nil {
		C.QWidget_SetDisabled_Bool(p.Pointer(), goBoolToCInt(disable))
	}
}

func (p *qwidget) ConnectSlotSetEnabled() {
	C.QWidget_ConnectSlotSetEnabled(p.Pointer())
}

func (p *qwidget) DisconnectSlotSetEnabled() {
	C.QWidget_DisconnectSlotSetEnabled(p.Pointer())
}

func (p *qwidget) SlotSetEnabled(enabled bool) {
	if p.Pointer() != nil {
		C.QWidget_SetEnabled_Bool(p.Pointer(), goBoolToCInt(enabled))
	}
}

func (p *qwidget) ConnectSlotSetHidden() {
	C.QWidget_ConnectSlotSetHidden(p.Pointer())
}

func (p *qwidget) DisconnectSlotSetHidden() {
	C.QWidget_DisconnectSlotSetHidden(p.Pointer())
}

func (p *qwidget) SlotSetHidden(hidden bool) {
	if p.Pointer() != nil {
		C.QWidget_SetHidden_Bool(p.Pointer(), goBoolToCInt(hidden))
	}
}

func (p *qwidget) ConnectSlotSetStyleSheet() {
	C.QWidget_ConnectSlotSetStyleSheet(p.Pointer())
}

func (p *qwidget) DisconnectSlotSetStyleSheet() {
	C.QWidget_DisconnectSlotSetStyleSheet(p.Pointer())
}

func (p *qwidget) SlotSetStyleSheet(styleSheet string) {
	if p.Pointer() != nil {
		C.QWidget_SetStyleSheet_String(p.Pointer(), C.CString(styleSheet))
	}
}

func (p *qwidget) ConnectSlotSetWindowModified() {
	C.QWidget_ConnectSlotSetWindowModified(p.Pointer())
}

func (p *qwidget) DisconnectSlotSetWindowModified() {
	C.QWidget_DisconnectSlotSetWindowModified(p.Pointer())
}

func (p *qwidget) SlotSetWindowModified(modified bool) {
	if p.Pointer() != nil {
		C.QWidget_SetWindowModified_Bool(p.Pointer(), goBoolToCInt(modified))
	}
}

func (p *qwidget) ConnectSlotSetWindowTitle() {
	C.QWidget_ConnectSlotSetWindowTitle(p.Pointer())
}

func (p *qwidget) DisconnectSlotSetWindowTitle() {
	C.QWidget_DisconnectSlotSetWindowTitle(p.Pointer())
}

func (p *qwidget) SlotSetWindowTitle(windowTitle string) {
	if p.Pointer() != nil {
		C.QWidget_SetWindowTitle_String(p.Pointer(), C.CString(windowTitle))
	}
}

func (p *qwidget) ConnectSlotShow() {
	C.QWidget_ConnectSlotShow(p.Pointer())
}

func (p *qwidget) DisconnectSlotShow() {
	C.QWidget_DisconnectSlotShow(p.Pointer())
}

func (p *qwidget) SlotShow() {
	if p.Pointer() != nil {
		C.QWidget_Show(p.Pointer())
	}
}

func (p *qwidget) ConnectSlotShowFullScreen() {
	C.QWidget_ConnectSlotShowFullScreen(p.Pointer())
}

func (p *qwidget) DisconnectSlotShowFullScreen() {
	C.QWidget_DisconnectSlotShowFullScreen(p.Pointer())
}

func (p *qwidget) SlotShowFullScreen() {
	if p.Pointer() != nil {
		C.QWidget_ShowFullScreen(p.Pointer())
	}
}

func (p *qwidget) ConnectSlotShowMaximized() {
	C.QWidget_ConnectSlotShowMaximized(p.Pointer())
}

func (p *qwidget) DisconnectSlotShowMaximized() {
	C.QWidget_DisconnectSlotShowMaximized(p.Pointer())
}

func (p *qwidget) SlotShowMaximized() {
	if p.Pointer() != nil {
		C.QWidget_ShowMaximized(p.Pointer())
	}
}

func (p *qwidget) ConnectSlotShowMinimized() {
	C.QWidget_ConnectSlotShowMinimized(p.Pointer())
}

func (p *qwidget) DisconnectSlotShowMinimized() {
	C.QWidget_DisconnectSlotShowMinimized(p.Pointer())
}

func (p *qwidget) SlotShowMinimized() {
	if p.Pointer() != nil {
		C.QWidget_ShowMinimized(p.Pointer())
	}
}

func (p *qwidget) ConnectSlotShowNormal() {
	C.QWidget_ConnectSlotShowNormal(p.Pointer())
}

func (p *qwidget) DisconnectSlotShowNormal() {
	C.QWidget_DisconnectSlotShowNormal(p.Pointer())
}

func (p *qwidget) SlotShowNormal() {
	if p.Pointer() != nil {
		C.QWidget_ShowNormal(p.Pointer())
	}
}

func (p *qwidget) ConnectSignalWindowIconTextChanged(f func()) {
	C.QWidget_ConnectSignalWindowIconTextChanged(p.Pointer())
	connectSignal(p.ObjectName(), "windowIconTextChanged", f)
}

func (p *qwidget) DisconnectSignalWindowIconTextChanged() {
	C.QWidget_DisconnectSignalWindowIconTextChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "windowIconTextChanged")
}

func (p *qwidget) SignalWindowIconTextChanged() func() {
	return getSignal(p.ObjectName(), "windowIconTextChanged")
}

func (p *qwidget) ConnectSignalWindowTitleChanged(f func()) {
	C.QWidget_ConnectSignalWindowTitleChanged(p.Pointer())
	connectSignal(p.ObjectName(), "windowTitleChanged", f)
}

func (p *qwidget) DisconnectSignalWindowTitleChanged() {
	C.QWidget_DisconnectSignalWindowTitleChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "windowTitleChanged")
}

func (p *qwidget) SignalWindowTitleChanged() func() {
	return getSignal(p.ObjectName(), "windowTitleChanged")
}

func QWidget_KeyboardGrabber() QWidget {
	var qwidget = new(qwidget)
	qwidget.SetPointer(C.QWidget_KeyboardGrabber())
	if qwidget.ObjectName() == "" {
		qwidget.SetObjectName("QWidget_" + randomIdentifier())
	}
	return qwidget
}

func QWidget_MouseGrabber() QWidget {
	var qwidget = new(qwidget)
	qwidget.SetPointer(C.QWidget_MouseGrabber())
	if qwidget.ObjectName() == "" {
		qwidget.SetObjectName("QWidget_" + randomIdentifier())
	}
	return qwidget
}

func QWidget_SetTabOrder(first QWidget, second QWidget) {
	var firstPtr C.QtObjectPtr
	if first != nil {
		firstPtr = first.Pointer()
	}
	var secondPtr C.QtObjectPtr
	if second != nil {
		secondPtr = second.Pointer()
	}
	C.QWidget_SetTabOrder_QWidget_QWidget(firstPtr, secondPtr)
}

//export callbackQWidget
func callbackQWidget(ptr C.QtObjectPtr, msg *C.char) {
	var qwidget = new(qwidget)
	qwidget.SetPointer(ptr)
	getSignal(qwidget.ObjectName(), C.GoString(msg))()
}
