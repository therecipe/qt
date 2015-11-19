package widgets

//#include "qwidget.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QWidget struct {
	gui.QPaintDevice
	core.QObject
}

type QWidget_ITF interface {
	gui.QPaintDevice_ITF
	core.QObject_ITF
	QWidget_PTR() *QWidget
}

func (p *QWidget) Pointer() unsafe.Pointer {
	return p.QPaintDevice_PTR().Pointer()
}

func (p *QWidget) SetPointer(ptr unsafe.Pointer) {
	p.QPaintDevice_PTR().SetPointer(ptr)
	p.QObject_PTR().SetPointer(ptr)
}

func PointerFromQWidget(ptr QWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidget_PTR().Pointer()
	}
	return nil
}

func NewQWidgetFromPointer(ptr unsafe.Pointer) *QWidget {
	var n = new(QWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWidget) QWidget_PTR() *QWidget {
	return ptr
}

//QWidget::RenderFlag
type QWidget__RenderFlag int64

const (
	QWidget__DrawWindowBackground = QWidget__RenderFlag(0x1)
	QWidget__DrawChildren         = QWidget__RenderFlag(0x2)
	QWidget__IgnoreMask           = QWidget__RenderFlag(0x4)
)

func (ptr *QWidget) AcceptDrops() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_AcceptDrops(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) AccessibleDescription() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_AccessibleDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) AccessibleName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_AccessibleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) ActivateWindow() {
	if ptr.Pointer() != nil {
		C.QWidget_ActivateWindow(ptr.Pointer())
	}
}

func (ptr *QWidget) AutoFillBackground() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_AutoFillBackground(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) ChildrenRegion() *gui.QRegion {
	if ptr.Pointer() != nil {
		return gui.NewQRegionFromPointer(C.QWidget_ChildrenRegion(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ClearMask() {
	if ptr.Pointer() != nil {
		C.QWidget_ClearMask(ptr.Pointer())
	}
}

func (ptr *QWidget) ContextMenuPolicy() core.Qt__ContextMenuPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__ContextMenuPolicy(C.QWidget_ContextMenuPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) FocusPolicy() core.Qt__FocusPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__FocusPolicy(C.QWidget_FocusPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) GrabKeyboard() {
	if ptr.Pointer() != nil {
		C.QWidget_GrabKeyboard(ptr.Pointer())
	}
}

func (ptr *QWidget) GrabMouse() {
	if ptr.Pointer() != nil {
		C.QWidget_GrabMouse(ptr.Pointer())
	}
}

func (ptr *QWidget) GrabMouse2(cursor gui.QCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_GrabMouse2(ptr.Pointer(), gui.PointerFromQCursor(cursor))
	}
}

func (ptr *QWidget) HasFocus() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) InputMethodHints() core.Qt__InputMethodHint {
	if ptr.Pointer() != nil {
		return core.Qt__InputMethodHint(C.QWidget_InputMethodHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) IsActiveWindow() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsActiveWindow(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsFullScreen() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsMaximized() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsMaximized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsMinimized() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsMinimized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsWindowModified() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsWindowModified(ptr.Pointer()) != 0
	}
	return false
}

func QWidget_KeyboardGrabber() *QWidget {
	return NewQWidgetFromPointer(C.QWidget_QWidget_KeyboardGrabber())
}

func (ptr *QWidget) LayoutDirection() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QWidget_LayoutDirection(ptr.Pointer()))
	}
	return 0
}

func QWidget_MouseGrabber() *QWidget {
	return NewQWidgetFromPointer(C.QWidget_QWidget_MouseGrabber())
}

func (ptr *QWidget) Move(v core.QPoint_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Move(ptr.Pointer(), core.PointerFromQPoint(v))
	}
}

func (ptr *QWidget) PaintEngine() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWidget_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ReleaseKeyboard() {
	if ptr.Pointer() != nil {
		C.QWidget_ReleaseKeyboard(ptr.Pointer())
	}
}

func (ptr *QWidget) ReleaseMouse() {
	if ptr.Pointer() != nil {
		C.QWidget_ReleaseMouse(ptr.Pointer())
	}
}

func (ptr *QWidget) Resize(v core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Resize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetAcceptDrops(on bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAcceptDrops(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWidget) SetAccessibleDescription(description string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAccessibleDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QWidget) SetAccessibleName(name string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAccessibleName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QWidget) SetAutoFillBackground(enabled bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAutoFillBackground(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QWidget) SetContextMenuPolicy(policy core.Qt__ContextMenuPolicy) {
	if ptr.Pointer() != nil {
		C.QWidget_SetContextMenuPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QWidget) SetCursor(v gui.QCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetCursor(ptr.Pointer(), gui.PointerFromQCursor(v))
	}
}

func (ptr *QWidget) SetEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QWidget) SetFixedSize2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedSize2(ptr.Pointer(), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) SetFocusPolicy(policy core.Qt__FocusPolicy) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocusPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QWidget) SetFont(v gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFont(ptr.Pointer(), gui.PointerFromQFont(v))
	}
}

func (ptr *QWidget) SetGeometry(v core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetGeometry(ptr.Pointer(), core.PointerFromQRect(v))
	}
}

func (ptr *QWidget) SetInputMethodHints(hints core.Qt__InputMethodHint) {
	if ptr.Pointer() != nil {
		C.QWidget_SetInputMethodHints(ptr.Pointer(), C.int(hints))
	}
}

func (ptr *QWidget) SetLayout(layout QLayout_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetLayout(ptr.Pointer(), PointerFromQLayout(layout))
	}
}

func (ptr *QWidget) SetLayoutDirection(direction core.Qt__LayoutDirection) {
	if ptr.Pointer() != nil {
		C.QWidget_SetLayoutDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QWidget) SetLocale(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QWidget) SetMask(bitmap gui.QBitmap_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMask(ptr.Pointer(), gui.PointerFromQBitmap(bitmap))
	}
}

func (ptr *QWidget) SetMask2(region gui.QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMask2(ptr.Pointer(), gui.PointerFromQRegion(region))
	}
}

func (ptr *QWidget) SetMaximumHeight(maxh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumHeight(ptr.Pointer(), C.int(maxh))
	}
}

func (ptr *QWidget) SetMaximumWidth(maxw int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumWidth(ptr.Pointer(), C.int(maxw))
	}
}

func (ptr *QWidget) SetMinimumHeight(minh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumHeight(ptr.Pointer(), C.int(minh))
	}
}

func (ptr *QWidget) SetMinimumWidth(minw int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumWidth(ptr.Pointer(), C.int(minw))
	}
}

func (ptr *QWidget) SetPalette(v gui.QPalette_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetPalette(ptr.Pointer(), gui.PointerFromQPalette(v))
	}
}

func (ptr *QWidget) SetSizePolicy(v QSizePolicy_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizePolicy(ptr.Pointer(), PointerFromQSizePolicy(v))
	}
}

func (ptr *QWidget) SetStatusTip(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetStatusTip(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) SetStyleSheet(styleSheet string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetStyleSheet(ptr.Pointer(), C.CString(styleSheet))
	}
}

func (ptr *QWidget) SetToolTip(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetToolTip(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) SetToolTipDuration(msec int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetToolTipDuration(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QWidget) SetUpdatesEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetUpdatesEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWidget) SetWhatsThis(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWhatsThis(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) SetWindowFilePath(filePath string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowFilePath(ptr.Pointer(), C.CString(filePath))
	}
}

func (ptr *QWidget) SetWindowFlags(ty core.Qt__WindowType) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowFlags(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QWidget) SetWindowIcon(icon gui.QIcon_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QWidget) SetWindowIconText(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowIconText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) SetWindowModality(windowModality core.Qt__WindowModality) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowModality(ptr.Pointer(), C.int(windowModality))
	}
}

func (ptr *QWidget) SetWindowModified(v bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QWidget) SetWindowOpacity(level float64) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowOpacity(ptr.Pointer(), C.double(level))
	}
}

func (ptr *QWidget) SetWindowState(windowState core.Qt__WindowState) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowState(ptr.Pointer(), C.int(windowState))
	}
}

func (ptr *QWidget) SetWindowTitle(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowTitle(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) StatusTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) StyleSheet() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_StyleSheet(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) ToolTipDuration() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_ToolTipDuration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) UnsetCursor() {
	if ptr.Pointer() != nil {
		C.QWidget_UnsetCursor(ptr.Pointer())
	}
}

func (ptr *QWidget) UnsetLayoutDirection() {
	if ptr.Pointer() != nil {
		C.QWidget_UnsetLayoutDirection(ptr.Pointer())
	}
}

func (ptr *QWidget) UnsetLocale() {
	if ptr.Pointer() != nil {
		C.QWidget_UnsetLocale(ptr.Pointer())
	}
}

func (ptr *QWidget) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) WindowFilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowFilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) WindowIconText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowIconText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) WindowModality() core.Qt__WindowModality {
	if ptr.Pointer() != nil {
		return core.Qt__WindowModality(C.QWidget_WindowModality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) WindowOpacity() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QWidget_WindowOpacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) WindowTitle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) X() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_Y(ptr.Pointer()))
	}
	return 0
}

func NewQWidget(parent QWidget_ITF, f core.Qt__WindowType) *QWidget {
	return NewQWidgetFromPointer(C.QWidget_NewQWidget(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QWidget) AddAction(action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_AddAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QWidget) AdjustSize() {
	if ptr.Pointer() != nil {
		C.QWidget_AdjustSize(ptr.Pointer())
	}
}

func (ptr *QWidget) BackgroundRole() gui.QPalette__ColorRole {
	if ptr.Pointer() != nil {
		return gui.QPalette__ColorRole(C.QWidget_BackgroundRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) BackingStore() *gui.QBackingStore {
	if ptr.Pointer() != nil {
		return gui.NewQBackingStoreFromPointer(C.QWidget_BackingStore(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ChildAt2(p core.QPoint_ITF) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_ChildAt2(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QWidget) ChildAt(x int, y int) *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_ChildAt(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QWidget) ClearFocus() {
	if ptr.Pointer() != nil {
		C.QWidget_ClearFocus(ptr.Pointer())
	}
}

func (ptr *QWidget) Close() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) EnsurePolished() {
	if ptr.Pointer() != nil {
		C.QWidget_EnsurePolished(ptr.Pointer())
	}
}

func (ptr *QWidget) FocusProxy() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_FocusProxy(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) FocusWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_FocusWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ForegroundRole() gui.QPalette__ColorRole {
	if ptr.Pointer() != nil {
		return gui.QPalette__ColorRole(C.QWidget_ForegroundRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) GetContentsMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QWidget_GetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QWidget) GrabGesture(gesture core.Qt__GestureType, flags core.Qt__GestureFlag) {
	if ptr.Pointer() != nil {
		C.QWidget_GrabGesture(ptr.Pointer(), C.int(gesture), C.int(flags))
	}
}

func (ptr *QWidget) GrabShortcut(key gui.QKeySequence_ITF, context core.Qt__ShortcutContext) int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_GrabShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(key), C.int(context)))
	}
	return 0
}

func (ptr *QWidget) GraphicsEffect() *QGraphicsEffect {
	if ptr.Pointer() != nil {
		return NewQGraphicsEffectFromPointer(C.QWidget_GraphicsEffect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) GraphicsProxyWidget() *QGraphicsProxyWidget {
	if ptr.Pointer() != nil {
		return NewQGraphicsProxyWidgetFromPointer(C.QWidget_GraphicsProxyWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) HasMouseTracking() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_HasMouseTracking(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QWidget) Hide() {
	if ptr.Pointer() != nil {
		C.QWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWidget_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QWidget) InsertAction(before QAction_ITF, action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_InsertAction(ptr.Pointer(), PointerFromQAction(before), PointerFromQAction(action))
	}
}

func (ptr *QWidget) IsAncestorOf(child QWidget_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsAncestorOf(ptr.Pointer(), PointerFromQWidget(child)) != 0
	}
	return false
}

func (ptr *QWidget) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsEnabledTo(ancestor QWidget_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsEnabledTo(ptr.Pointer(), PointerFromQWidget(ancestor)) != 0
	}
	return false
}

func (ptr *QWidget) IsHidden() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsModal() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsModal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsVisibleTo(ancestor QWidget_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsVisibleTo(ptr.Pointer(), PointerFromQWidget(ancestor)) != 0
	}
	return false
}

func (ptr *QWidget) IsWindow() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsWindow(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) Layout() *QLayout {
	if ptr.Pointer() != nil {
		return NewQLayoutFromPointer(C.QWidget_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) Lower() {
	if ptr.Pointer() != nil {
		C.QWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QWidget) Mask() *gui.QRegion {
	if ptr.Pointer() != nil {
		return gui.NewQRegionFromPointer(C.QWidget_Mask(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) MaximumHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_MaximumHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) MaximumWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_MaximumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) MinimumHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_MinimumHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) MinimumWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_MinimumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) Move2(x int, y int) {
	if ptr.Pointer() != nil {
		C.QWidget_Move2(ptr.Pointer(), C.int(x), C.int(y))
	}
}

func (ptr *QWidget) NativeParentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_NativeParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) NextInFocusChain() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_NextInFocusChain(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) OverrideWindowFlags(flags core.Qt__WindowType) {
	if ptr.Pointer() != nil {
		C.QWidget_OverrideWindowFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QWidget) ParentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_ParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) PreviousInFocusChain() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_PreviousInFocusChain(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) Raise() {
	if ptr.Pointer() != nil {
		C.QWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QWidget) ReleaseShortcut(id int) {
	if ptr.Pointer() != nil {
		C.QWidget_ReleaseShortcut(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWidget) RemoveAction(action QAction_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_RemoveAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QWidget) Render(target gui.QPaintDevice_ITF, targetOffset core.QPoint_ITF, sourceRegion gui.QRegion_ITF, renderFlags QWidget__RenderFlag) {
	if ptr.Pointer() != nil {
		C.QWidget_Render(ptr.Pointer(), gui.PointerFromQPaintDevice(target), core.PointerFromQPoint(targetOffset), gui.PointerFromQRegion(sourceRegion), C.int(renderFlags))
	}
}

func (ptr *QWidget) Render2(painter gui.QPainter_ITF, targetOffset core.QPoint_ITF, sourceRegion gui.QRegion_ITF, renderFlags QWidget__RenderFlag) {
	if ptr.Pointer() != nil {
		C.QWidget_Render2(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQPoint(targetOffset), gui.PointerFromQRegion(sourceRegion), C.int(renderFlags))
	}
}

func (ptr *QWidget) Repaint() {
	if ptr.Pointer() != nil {
		C.QWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QWidget) Repaint3(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Repaint3(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWidget) Repaint4(rgn gui.QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Repaint4(ptr.Pointer(), gui.PointerFromQRegion(rgn))
	}
}

func (ptr *QWidget) Repaint2(x int, y int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_Repaint2(ptr.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) Resize2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_Resize2(ptr.Pointer(), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) RestoreGeometry(geometry core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_RestoreGeometry(ptr.Pointer(), core.PointerFromQByteArray(geometry)) != 0
	}
	return false
}

func (ptr *QWidget) SaveGeometry() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QWidget_SaveGeometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) Scroll(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QWidget_Scroll(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QWidget) Scroll2(dx int, dy int, r core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Scroll2(ptr.Pointer(), C.int(dx), C.int(dy), core.PointerFromQRect(r))
	}
}

func (ptr *QWidget) SetAttribute(attribute core.Qt__WidgetAttribute, on bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAttribute(ptr.Pointer(), C.int(attribute), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWidget) SetBackgroundRole(role gui.QPalette__ColorRole) {
	if ptr.Pointer() != nil {
		C.QWidget_SetBackgroundRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QWidget) SetBaseSize(v core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetBaseSize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetBaseSize2(basew int, baseh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetBaseSize2(ptr.Pointer(), C.int(basew), C.int(baseh))
	}
}

func (ptr *QWidget) SetContentsMargins2(margins core.QMargins_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetContentsMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QWidget) SetDisabled(disable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *QWidget) SetFixedHeight(h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QWidget) SetFixedSize(s core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedSize(ptr.Pointer(), core.PointerFromQSize(s))
	}
}

func (ptr *QWidget) SetFixedWidth(w int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QWidget) SetFocus2() {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QWidget) SetFocus(reason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocus(ptr.Pointer(), C.int(reason))
	}
}

func (ptr *QWidget) SetFocusProxy(w QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocusProxy(ptr.Pointer(), PointerFromQWidget(w))
	}
}

func (ptr *QWidget) SetForegroundRole(role gui.QPalette__ColorRole) {
	if ptr.Pointer() != nil {
		C.QWidget_SetForegroundRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QWidget) SetGeometry2(x int, y int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetGeometry2(ptr.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) SetGraphicsEffect(effect QGraphicsEffect_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetGraphicsEffect(ptr.Pointer(), PointerFromQGraphicsEffect(effect))
	}
}

func (ptr *QWidget) SetHidden(hidden bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *QWidget) SetMaximumSize(v core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumSize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetMaximumSize2(maxw int, maxh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumSize2(ptr.Pointer(), C.int(maxw), C.int(maxh))
	}
}

func (ptr *QWidget) SetMinimumSize(v core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetMinimumSize2(minw int, minh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize2(ptr.Pointer(), C.int(minw), C.int(minh))
	}
}

func (ptr *QWidget) SetMouseTracking(enable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMouseTracking(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetParent(parent QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetParent(ptr.Pointer(), PointerFromQWidget(parent))
	}
}

func (ptr *QWidget) SetParent2(parent QWidget_ITF, f core.Qt__WindowType) {
	if ptr.Pointer() != nil {
		C.QWidget_SetParent2(ptr.Pointer(), PointerFromQWidget(parent), C.int(f))
	}
}

func (ptr *QWidget) SetShortcutAutoRepeat(id int, enable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetShortcutAutoRepeat(ptr.Pointer(), C.int(id), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetShortcutEnabled(id int, enable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetShortcutEnabled(ptr.Pointer(), C.int(id), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetSizeIncrement(v core.QSize_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizeIncrement(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetSizeIncrement2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizeIncrement2(ptr.Pointer(), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) SetSizePolicy2(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizePolicy2(ptr.Pointer(), C.int(horizontal), C.int(vertical))
	}
}

func (ptr *QWidget) SetStyle(style QStyle_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

func QWidget_SetTabOrder(first QWidget_ITF, second QWidget_ITF) {
	C.QWidget_QWidget_SetTabOrder(PointerFromQWidget(first), PointerFromQWidget(second))
}

func (ptr *QWidget) SetWindowRole(role string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowRole(ptr.Pointer(), C.CString(role))
	}
}

func (ptr *QWidget) Show() {
	if ptr.Pointer() != nil {
		C.QWidget_Show(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QWidget) StackUnder(w QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_StackUnder(ptr.Pointer(), PointerFromQWidget(w))
	}
}

func (ptr *QWidget) Style() *QStyle {
	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QWidget_Style(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) TestAttribute(attribute core.Qt__WidgetAttribute) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_TestAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QWidget) UnderMouse() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_UnderMouse(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) UngrabGesture(gesture core.Qt__GestureType) {
	if ptr.Pointer() != nil {
		C.QWidget_UngrabGesture(ptr.Pointer(), C.int(gesture))
	}
}

func (ptr *QWidget) Update() {
	if ptr.Pointer() != nil {
		C.QWidget_Update(ptr.Pointer())
	}
}

func (ptr *QWidget) Update3(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Update3(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWidget) Update4(rgn gui.QRegion_ITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Update4(ptr.Pointer(), gui.PointerFromQRegion(rgn))
	}
}

func (ptr *QWidget) Update2(x int, y int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_Update2(ptr.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) UpdateGeometry() {
	if ptr.Pointer() != nil {
		C.QWidget_UpdateGeometry(ptr.Pointer())
	}
}

func (ptr *QWidget) UpdatesEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_UpdatesEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) VisibleRegion() *gui.QRegion {
	if ptr.Pointer() != nil {
		return gui.NewQRegionFromPointer(C.QWidget_VisibleRegion(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) Window() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) WindowFlags() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWidget_WindowFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) WindowHandle() *gui.QWindow {
	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QWidget_WindowHandle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ConnectWindowIconTextChanged(f func(iconText string)) {
	if ptr.Pointer() != nil {
		C.QWidget_ConnectWindowIconTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowIconTextChanged", f)
	}
}

func (ptr *QWidget) DisconnectWindowIconTextChanged() {
	if ptr.Pointer() != nil {
		C.QWidget_DisconnectWindowIconTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowIconTextChanged")
	}
}

//export callbackQWidgetWindowIconTextChanged
func callbackQWidgetWindowIconTextChanged(ptrName *C.char, iconText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "windowIconTextChanged").(func(string))(C.GoString(iconText))
}

func (ptr *QWidget) WindowRole() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) WindowState() core.Qt__WindowState {
	if ptr.Pointer() != nil {
		return core.Qt__WindowState(C.QWidget_WindowState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) ConnectWindowTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {
		C.QWidget_ConnectWindowTitleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowTitleChanged", f)
	}
}

func (ptr *QWidget) DisconnectWindowTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWidget_DisconnectWindowTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowTitleChanged")
	}
}

//export callbackQWidgetWindowTitleChanged
func callbackQWidgetWindowTitleChanged(ptrName *C.char, title *C.char) {
	qt.GetSignal(C.GoString(ptrName), "windowTitleChanged").(func(string))(C.GoString(title))
}

func (ptr *QWidget) WindowType() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWidget_WindowType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) DestroyQWidget() {
	if ptr.Pointer() != nil {
		C.QWidget_DestroyQWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func QWidget_CreateWindowContainer(window gui.QWindow_ITF, parent QWidget_ITF, flags core.Qt__WindowType) *QWidget {
	return NewQWidgetFromPointer(C.QWidget_QWidget_CreateWindowContainer(gui.PointerFromQWindow(window), PointerFromQWidget(parent), C.int(flags)))
}
