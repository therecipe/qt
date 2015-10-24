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

type QWidgetITF interface {
	gui.QPaintDeviceITF
	core.QObjectITF
	QWidgetPTR() *QWidget
}

func (p *QWidget) Pointer() unsafe.Pointer {
	return p.QPaintDevicePTR().Pointer()
}

func (p *QWidget) SetPointer(ptr unsafe.Pointer) {
	p.QPaintDevicePTR().SetPointer(ptr)
	p.QObjectPTR().SetPointer(ptr)
}

func PointerFromQWidget(ptr QWidgetITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWidgetPTR().Pointer()
	}
	return nil
}

func QWidgetFromPointer(ptr unsafe.Pointer) *QWidget {
	var n = new(QWidget)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QWidget) QWidgetPTR() *QWidget {
	return ptr
}

//QWidget::RenderFlag
type QWidget__RenderFlag int

var (
	QWidget__DrawWindowBackground = QWidget__RenderFlag(0x1)
	QWidget__DrawChildren         = QWidget__RenderFlag(0x2)
	QWidget__IgnoreMask           = QWidget__RenderFlag(0x4)
)

func (ptr *QWidget) AcceptDrops() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_AcceptDrops(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) AccessibleDescription() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_AccessibleDescription(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) AccessibleName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_AccessibleName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) ActivateWindow() {
	if ptr.Pointer() != nil {
		C.QWidget_ActivateWindow(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) AutoFillBackground() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_AutoFillBackground(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) ClearMask() {
	if ptr.Pointer() != nil {
		C.QWidget_ClearMask(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) ContextMenuPolicy() core.Qt__ContextMenuPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__ContextMenuPolicy(C.QWidget_ContextMenuPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) FocusPolicy() core.Qt__FocusPolicy {
	if ptr.Pointer() != nil {
		return core.Qt__FocusPolicy(C.QWidget_FocusPolicy(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) GrabKeyboard() {
	if ptr.Pointer() != nil {
		C.QWidget_GrabKeyboard(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) GrabMouse() {
	if ptr.Pointer() != nil {
		C.QWidget_GrabMouse(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) GrabMouse2(cursor gui.QCursorITF) {
	if ptr.Pointer() != nil {
		C.QWidget_GrabMouse2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQCursor(cursor)))
	}
}

func (ptr *QWidget) HasFocus() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_HasFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) InputMethodHints() core.Qt__InputMethodHint {
	if ptr.Pointer() != nil {
		return core.Qt__InputMethodHint(C.QWidget_InputMethodHints(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) IsActiveWindow() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsActiveWindow(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) IsFullScreen() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsFullScreen(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) IsMaximized() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsMaximized(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) IsMinimized() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsMinimized(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) IsWindowModified() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsWindowModified(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func QWidget_KeyboardGrabber() *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QWidget_QWidget_KeyboardGrabber()))
}

func (ptr *QWidget) LayoutDirection() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QWidget_LayoutDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QWidget_MouseGrabber() *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QWidget_QWidget_MouseGrabber()))
}

func (ptr *QWidget) Move(v core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Move(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(v)))
	}
}

func (ptr *QWidget) PaintEngine() *gui.QPaintEngine {
	if ptr.Pointer() != nil {
		return gui.QPaintEngineFromPointer(unsafe.Pointer(C.QWidget_PaintEngine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) ReleaseKeyboard() {
	if ptr.Pointer() != nil {
		C.QWidget_ReleaseKeyboard(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) ReleaseMouse() {
	if ptr.Pointer() != nil {
		C.QWidget_ReleaseMouse(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) Resize(v core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Resize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(v)))
	}
}

func (ptr *QWidget) SetAcceptDrops(on bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAcceptDrops(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWidget) SetAccessibleDescription(description string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAccessibleDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(description))
	}
}

func (ptr *QWidget) SetAccessibleName(name string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAccessibleName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QWidget) SetAutoFillBackground(enabled bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAutoFillBackground(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QWidget) SetContextMenuPolicy(policy core.Qt__ContextMenuPolicy) {
	if ptr.Pointer() != nil {
		C.QWidget_SetContextMenuPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QWidget) SetCursor(v gui.QCursorITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetCursor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQCursor(v)))
	}
}

func (ptr *QWidget) SetEnabled(v bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QWidget) SetFixedSize2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedSize2(C.QtObjectPtr(ptr.Pointer()), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) SetFocusPolicy(policy core.Qt__FocusPolicy) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocusPolicy(C.QtObjectPtr(ptr.Pointer()), C.int(policy))
	}
}

func (ptr *QWidget) SetFont(v gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(v)))
	}
}

func (ptr *QWidget) SetGeometry(v core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(v)))
	}
}

func (ptr *QWidget) SetInputMethodHints(hints core.Qt__InputMethodHint) {
	if ptr.Pointer() != nil {
		C.QWidget_SetInputMethodHints(C.QtObjectPtr(ptr.Pointer()), C.int(hints))
	}
}

func (ptr *QWidget) SetLayout(layout QLayoutITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetLayout(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLayout(layout)))
	}
}

func (ptr *QWidget) SetLayoutDirection(direction core.Qt__LayoutDirection) {
	if ptr.Pointer() != nil {
		C.QWidget_SetLayoutDirection(C.QtObjectPtr(ptr.Pointer()), C.int(direction))
	}
}

func (ptr *QWidget) SetLocale(locale core.QLocaleITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetLocale(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLocale(locale)))
	}
}

func (ptr *QWidget) SetMask(bitmap gui.QBitmapITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMask(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBitmap(bitmap)))
	}
}

func (ptr *QWidget) SetMask2(region gui.QRegionITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMask2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQRegion(region)))
	}
}

func (ptr *QWidget) SetMaximumHeight(maxh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumHeight(C.QtObjectPtr(ptr.Pointer()), C.int(maxh))
	}
}

func (ptr *QWidget) SetMaximumWidth(maxw int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumWidth(C.QtObjectPtr(ptr.Pointer()), C.int(maxw))
	}
}

func (ptr *QWidget) SetMinimumHeight(minh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumHeight(C.QtObjectPtr(ptr.Pointer()), C.int(minh))
	}
}

func (ptr *QWidget) SetMinimumWidth(minw int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumWidth(C.QtObjectPtr(ptr.Pointer()), C.int(minw))
	}
}

func (ptr *QWidget) SetPalette(v gui.QPaletteITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetPalette(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPalette(v)))
	}
}

func (ptr *QWidget) SetSizePolicy(v QSizePolicyITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizePolicy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQSizePolicy(v)))
	}
}

func (ptr *QWidget) SetStatusTip(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetStatusTip(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QWidget) SetStyleSheet(styleSheet string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetStyleSheet(C.QtObjectPtr(ptr.Pointer()), C.CString(styleSheet))
	}
}

func (ptr *QWidget) SetToolTip(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetToolTip(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QWidget) SetToolTipDuration(msec int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetToolTipDuration(C.QtObjectPtr(ptr.Pointer()), C.int(msec))
	}
}

func (ptr *QWidget) SetUpdatesEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetUpdatesEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QWidget) SetWhatsThis(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWhatsThis(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QWidget) SetWindowFilePath(filePath string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowFilePath(C.QtObjectPtr(ptr.Pointer()), C.CString(filePath))
	}
}

func (ptr *QWidget) SetWindowFlags(ty core.Qt__WindowType) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowFlags(C.QtObjectPtr(ptr.Pointer()), C.int(ty))
	}
}

func (ptr *QWidget) SetWindowIcon(icon gui.QIconITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowIcon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQIcon(icon)))
	}
}

func (ptr *QWidget) SetWindowIconText(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowIconText(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QWidget) SetWindowModality(windowModality core.Qt__WindowModality) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowModality(C.QtObjectPtr(ptr.Pointer()), C.int(windowModality))
	}
}

func (ptr *QWidget) SetWindowModified(v bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowModified(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QWidget) SetWindowState(windowState core.Qt__WindowState) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowState(C.QtObjectPtr(ptr.Pointer()), C.int(windowState))
	}
}

func (ptr *QWidget) SetWindowTitle(v string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowTitle(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QWidget) StatusTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_StatusTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) StyleSheet() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_StyleSheet(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) ToolTip() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_ToolTip(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) ToolTipDuration() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_ToolTipDuration(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) UnsetCursor() {
	if ptr.Pointer() != nil {
		C.QWidget_UnsetCursor(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) UnsetLayoutDirection() {
	if ptr.Pointer() != nil {
		C.QWidget_UnsetLayoutDirection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) UnsetLocale() {
	if ptr.Pointer() != nil {
		C.QWidget_UnsetLocale(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) WhatsThis() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WhatsThis(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) WindowFilePath() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowFilePath(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) WindowIconText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowIconText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) WindowModality() core.Qt__WindowModality {
	if ptr.Pointer() != nil {
		return core.Qt__WindowModality(C.QWidget_WindowModality(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) WindowTitle() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowTitle(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) X() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQWidget(parent QWidgetITF, f core.Qt__WindowType) *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QWidget_NewQWidget(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(f))))
}

func (ptr *QWidget) AddAction(action QActionITF) {
	if ptr.Pointer() != nil {
		C.QWidget_AddAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(action)))
	}
}

func (ptr *QWidget) AdjustSize() {
	if ptr.Pointer() != nil {
		C.QWidget_AdjustSize(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) BackgroundRole() gui.QPalette__ColorRole {
	if ptr.Pointer() != nil {
		return gui.QPalette__ColorRole(C.QWidget_BackgroundRole(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) BackingStore() *gui.QBackingStore {
	if ptr.Pointer() != nil {
		return gui.QBackingStoreFromPointer(unsafe.Pointer(C.QWidget_BackingStore(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) ChildAt2(p core.QPointITF) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidget_ChildAt2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(p)))))
	}
	return nil
}

func (ptr *QWidget) ChildAt(x int, y int) *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidget_ChildAt(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))))
	}
	return nil
}

func (ptr *QWidget) ClearFocus() {
	if ptr.Pointer() != nil {
		C.QWidget_ClearFocus(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) Close() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_Close(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) EnsurePolished() {
	if ptr.Pointer() != nil {
		C.QWidget_EnsurePolished(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) FocusProxy() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidget_FocusProxy(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) FocusWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidget_FocusWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) ForegroundRole() gui.QPalette__ColorRole {
	if ptr.Pointer() != nil {
		return gui.QPalette__ColorRole(C.QWidget_ForegroundRole(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) GetContentsMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QWidget_GetContentsMargins(C.QtObjectPtr(ptr.Pointer()), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QWidget) GrabGesture(gesture core.Qt__GestureType, flags core.Qt__GestureFlag) {
	if ptr.Pointer() != nil {
		C.QWidget_GrabGesture(C.QtObjectPtr(ptr.Pointer()), C.int(gesture), C.int(flags))
	}
}

func (ptr *QWidget) GrabShortcut(key gui.QKeySequenceITF, context core.Qt__ShortcutContext) int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_GrabShortcut(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQKeySequence(key)), C.int(context)))
	}
	return 0
}

func (ptr *QWidget) GraphicsEffect() *QGraphicsEffect {
	if ptr.Pointer() != nil {
		return QGraphicsEffectFromPointer(unsafe.Pointer(C.QWidget_GraphicsEffect(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) GraphicsProxyWidget() *QGraphicsProxyWidget {
	if ptr.Pointer() != nil {
		return QGraphicsProxyWidgetFromPointer(unsafe.Pointer(C.QWidget_GraphicsProxyWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) HasHeightForWidth() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_HasHeightForWidth(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) HasMouseTracking() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_HasMouseTracking(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) Height() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_Height(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w)))
	}
	return 0
}

func (ptr *QWidget) Hide() {
	if ptr.Pointer() != nil {
		C.QWidget_Hide(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) InputMethodQuery(query core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(query)))
	}
	return ""
}

func (ptr *QWidget) InsertAction(before QActionITF, action QActionITF) {
	if ptr.Pointer() != nil {
		C.QWidget_InsertAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(before)), C.QtObjectPtr(PointerFromQAction(action)))
	}
}

func (ptr *QWidget) IsAncestorOf(child QWidgetITF) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsAncestorOf(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(child))) != 0
	}
	return false
}

func (ptr *QWidget) IsEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) IsEnabledTo(ancestor QWidgetITF) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsEnabledTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(ancestor))) != 0
	}
	return false
}

func (ptr *QWidget) IsHidden() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsHidden(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) IsModal() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsModal(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) IsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) IsVisibleTo(ancestor QWidgetITF) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsVisibleTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(ancestor))) != 0
	}
	return false
}

func (ptr *QWidget) IsWindow() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_IsWindow(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) Layout() *QLayout {
	if ptr.Pointer() != nil {
		return QLayoutFromPointer(unsafe.Pointer(C.QWidget_Layout(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) Lower() {
	if ptr.Pointer() != nil {
		C.QWidget_Lower(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) MaximumHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_MaximumHeight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) MaximumWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_MaximumWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) MinimumHeight() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_MinimumHeight(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) MinimumWidth() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_MinimumWidth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) Move2(x int, y int) {
	if ptr.Pointer() != nil {
		C.QWidget_Move2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))
	}
}

func (ptr *QWidget) NativeParentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidget_NativeParentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) NextInFocusChain() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidget_NextInFocusChain(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) OverrideWindowFlags(flags core.Qt__WindowType) {
	if ptr.Pointer() != nil {
		C.QWidget_OverrideWindowFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QWidget) ParentWidget() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidget_ParentWidget(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) PreviousInFocusChain() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidget_PreviousInFocusChain(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) Raise() {
	if ptr.Pointer() != nil {
		C.QWidget_Raise(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) ReleaseShortcut(id int) {
	if ptr.Pointer() != nil {
		C.QWidget_ReleaseShortcut(C.QtObjectPtr(ptr.Pointer()), C.int(id))
	}
}

func (ptr *QWidget) RemoveAction(action QActionITF) {
	if ptr.Pointer() != nil {
		C.QWidget_RemoveAction(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAction(action)))
	}
}

func (ptr *QWidget) Render(target gui.QPaintDeviceITF, targetOffset core.QPointITF, sourceRegion gui.QRegionITF, renderFlags QWidget__RenderFlag) {
	if ptr.Pointer() != nil {
		C.QWidget_Render(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPaintDevice(target)), C.QtObjectPtr(core.PointerFromQPoint(targetOffset)), C.QtObjectPtr(gui.PointerFromQRegion(sourceRegion)), C.int(renderFlags))
	}
}

func (ptr *QWidget) Render2(painter gui.QPainterITF, targetOffset core.QPointITF, sourceRegion gui.QRegionITF, renderFlags QWidget__RenderFlag) {
	if ptr.Pointer() != nil {
		C.QWidget_Render2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQPoint(targetOffset)), C.QtObjectPtr(gui.PointerFromQRegion(sourceRegion)), C.int(renderFlags))
	}
}

func (ptr *QWidget) Repaint() {
	if ptr.Pointer() != nil {
		C.QWidget_Repaint(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) Repaint3(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Repaint3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QWidget) Repaint4(rgn gui.QRegionITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Repaint4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQRegion(rgn)))
	}
}

func (ptr *QWidget) Repaint2(x int, y int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_Repaint2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) Resize2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_Resize2(C.QtObjectPtr(ptr.Pointer()), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) RestoreGeometry(geometry core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_RestoreGeometry(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(geometry))) != 0
	}
	return false
}

func (ptr *QWidget) Scroll(dx int, dy int) {
	if ptr.Pointer() != nil {
		C.QWidget_Scroll(C.QtObjectPtr(ptr.Pointer()), C.int(dx), C.int(dy))
	}
}

func (ptr *QWidget) Scroll2(dx int, dy int, r core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Scroll2(C.QtObjectPtr(ptr.Pointer()), C.int(dx), C.int(dy), C.QtObjectPtr(core.PointerFromQRect(r)))
	}
}

func (ptr *QWidget) SetAttribute(attribute core.Qt__WidgetAttribute, on bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(attribute), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWidget) SetBackgroundRole(role gui.QPalette__ColorRole) {
	if ptr.Pointer() != nil {
		C.QWidget_SetBackgroundRole(C.QtObjectPtr(ptr.Pointer()), C.int(role))
	}
}

func (ptr *QWidget) SetBaseSize(v core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetBaseSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(v)))
	}
}

func (ptr *QWidget) SetBaseSize2(basew int, baseh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetBaseSize2(C.QtObjectPtr(ptr.Pointer()), C.int(basew), C.int(baseh))
	}
}

func (ptr *QWidget) SetContentsMargins2(margins core.QMarginsITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetContentsMargins2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQMargins(margins)))
	}
}

func (ptr *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetContentsMargins(C.QtObjectPtr(ptr.Pointer()), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QWidget) SetDisabled(disable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetDisabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *QWidget) SetFixedHeight(h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedHeight(C.QtObjectPtr(ptr.Pointer()), C.int(h))
	}
}

func (ptr *QWidget) SetFixedSize(s core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(s)))
	}
}

func (ptr *QWidget) SetFixedWidth(w int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFixedWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w))
	}
}

func (ptr *QWidget) SetFocus2() {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocus2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) SetFocus(reason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocus(C.QtObjectPtr(ptr.Pointer()), C.int(reason))
	}
}

func (ptr *QWidget) SetFocusProxy(w QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetFocusProxy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(w)))
	}
}

func (ptr *QWidget) SetForegroundRole(role gui.QPalette__ColorRole) {
	if ptr.Pointer() != nil {
		C.QWidget_SetForegroundRole(C.QtObjectPtr(ptr.Pointer()), C.int(role))
	}
}

func (ptr *QWidget) SetGeometry2(x int, y int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetGeometry2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) SetGraphicsEffect(effect QGraphicsEffectITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetGraphicsEffect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsEffect(effect)))
	}
}

func (ptr *QWidget) SetHidden(hidden bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetHidden(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *QWidget) SetMaximumSize(v core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(v)))
	}
}

func (ptr *QWidget) SetMaximumSize2(maxw int, maxh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumSize2(C.QtObjectPtr(ptr.Pointer()), C.int(maxw), C.int(maxh))
	}
}

func (ptr *QWidget) SetMinimumSize(v core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(v)))
	}
}

func (ptr *QWidget) SetMinimumSize2(minw int, minh int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize2(C.QtObjectPtr(ptr.Pointer()), C.int(minw), C.int(minh))
	}
}

func (ptr *QWidget) SetMouseTracking(enable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetMouseTracking(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetParent(parent QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetParent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(parent)))
	}
}

func (ptr *QWidget) SetParent2(parent QWidgetITF, f core.Qt__WindowType) {
	if ptr.Pointer() != nil {
		C.QWidget_SetParent2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(parent)), C.int(f))
	}
}

func (ptr *QWidget) SetShortcutAutoRepeat(id int, enable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetShortcutAutoRepeat(C.QtObjectPtr(ptr.Pointer()), C.int(id), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetShortcutEnabled(id int, enable bool) {
	if ptr.Pointer() != nil {
		C.QWidget_SetShortcutEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(id), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetSizeIncrement(v core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizeIncrement(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(v)))
	}
}

func (ptr *QWidget) SetSizeIncrement2(w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizeIncrement2(C.QtObjectPtr(ptr.Pointer()), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) SetSizePolicy2(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy) {
	if ptr.Pointer() != nil {
		C.QWidget_SetSizePolicy2(C.QtObjectPtr(ptr.Pointer()), C.int(horizontal), C.int(vertical))
	}
}

func (ptr *QWidget) SetStyle(style QStyleITF) {
	if ptr.Pointer() != nil {
		C.QWidget_SetStyle(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStyle(style)))
	}
}

func QWidget_SetTabOrder(first QWidgetITF, second QWidgetITF) {
	C.QWidget_QWidget_SetTabOrder(C.QtObjectPtr(PointerFromQWidget(first)), C.QtObjectPtr(PointerFromQWidget(second)))
}

func (ptr *QWidget) SetWindowRole(role string) {
	if ptr.Pointer() != nil {
		C.QWidget_SetWindowRole(C.QtObjectPtr(ptr.Pointer()), C.CString(role))
	}
}

func (ptr *QWidget) Show() {
	if ptr.Pointer() != nil {
		C.QWidget_Show(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) ShowFullScreen() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowFullScreen(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) ShowMaximized() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowMaximized(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) ShowMinimized() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowMinimized(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) ShowNormal() {
	if ptr.Pointer() != nil {
		C.QWidget_ShowNormal(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) StackUnder(w QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QWidget_StackUnder(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(w)))
	}
}

func (ptr *QWidget) Style() *QStyle {
	if ptr.Pointer() != nil {
		return QStyleFromPointer(unsafe.Pointer(C.QWidget_Style(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) TestAttribute(attribute core.Qt__WidgetAttribute) bool {
	if ptr.Pointer() != nil {
		return C.QWidget_TestAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QWidget) UnderMouse() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_UnderMouse(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) UngrabGesture(gesture core.Qt__GestureType) {
	if ptr.Pointer() != nil {
		C.QWidget_UngrabGesture(C.QtObjectPtr(ptr.Pointer()), C.int(gesture))
	}
}

func (ptr *QWidget) Update() {
	if ptr.Pointer() != nil {
		C.QWidget_Update(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) Update3(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Update3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QWidget) Update4(rgn gui.QRegionITF) {
	if ptr.Pointer() != nil {
		C.QWidget_Update4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQRegion(rgn)))
	}
}

func (ptr *QWidget) Update2(x int, y int, w int, h int) {
	if ptr.Pointer() != nil {
		C.QWidget_Update2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) UpdateGeometry() {
	if ptr.Pointer() != nil {
		C.QWidget_UpdateGeometry(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QWidget) UpdatesEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QWidget_UpdatesEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QWidget) Width() int {
	if ptr.Pointer() != nil {
		return int(C.QWidget_Width(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) Window() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QWidget_Window(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) WindowFlags() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWidget_WindowFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) WindowHandle() *gui.QWindow {
	if ptr.Pointer() != nil {
		return gui.QWindowFromPointer(unsafe.Pointer(C.QWidget_WindowHandle(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QWidget) ConnectWindowIconTextChanged(f func(iconText string)) {
	if ptr.Pointer() != nil {
		C.QWidget_ConnectWindowIconTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "windowIconTextChanged", f)
	}
}

func (ptr *QWidget) DisconnectWindowIconTextChanged() {
	if ptr.Pointer() != nil {
		C.QWidget_DisconnectWindowIconTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "windowIconTextChanged")
	}
}

//export callbackQWidgetWindowIconTextChanged
func callbackQWidgetWindowIconTextChanged(ptrName *C.char, iconText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "windowIconTextChanged").(func(string))(C.GoString(iconText))
}

func (ptr *QWidget) WindowRole() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowRole(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QWidget) WindowState() core.Qt__WindowState {
	if ptr.Pointer() != nil {
		return core.Qt__WindowState(C.QWidget_WindowState(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) ConnectWindowTitleChanged(f func(title string)) {
	if ptr.Pointer() != nil {
		C.QWidget_ConnectWindowTitleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "windowTitleChanged", f)
	}
}

func (ptr *QWidget) DisconnectWindowTitleChanged() {
	if ptr.Pointer() != nil {
		C.QWidget_DisconnectWindowTitleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "windowTitleChanged")
	}
}

//export callbackQWidgetWindowTitleChanged
func callbackQWidgetWindowTitleChanged(ptrName *C.char, title *C.char) {
	qt.GetSignal(C.GoString(ptrName), "windowTitleChanged").(func(string))(C.GoString(title))
}

func (ptr *QWidget) WindowType() core.Qt__WindowType {
	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWidget_WindowType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWidget) DestroyQWidget() {
	if ptr.Pointer() != nil {
		C.QWidget_DestroyQWidget(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func QWidget_CreateWindowContainer(window gui.QWindowITF, parent QWidgetITF, flags core.Qt__WindowType) *QWidget {
	return QWidgetFromPointer(unsafe.Pointer(C.QWidget_QWidget_CreateWindowContainer(C.QtObjectPtr(gui.PointerFromQWindow(window)), C.QtObjectPtr(PointerFromQWidget(parent)), C.int(flags))))
}
