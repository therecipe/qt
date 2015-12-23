package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QWidget_") {
		n.SetObjectName("QWidget_" + qt.Identifier())
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
	defer qt.Recovering("QWidget::acceptDrops")

	if ptr.Pointer() != nil {
		return C.QWidget_AcceptDrops(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) AccessibleDescription() string {
	defer qt.Recovering("QWidget::accessibleDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_AccessibleDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) AccessibleName() string {
	defer qt.Recovering("QWidget::accessibleName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_AccessibleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQWidgetActionEvent
func callbackQWidgetActionEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ActivateWindow() {
	defer qt.Recovering("QWidget::activateWindow")

	if ptr.Pointer() != nil {
		C.QWidget_ActivateWindow(ptr.Pointer())
	}
}

func (ptr *QWidget) AutoFillBackground() bool {
	defer qt.Recovering("QWidget::autoFillBackground")

	if ptr.Pointer() != nil {
		return C.QWidget_AutoFillBackground(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) BaseSize() *core.QSize {
	defer qt.Recovering("QWidget::baseSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWidget_BaseSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ChildrenRect() *core.QRect {
	defer qt.Recovering("QWidget::childrenRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWidget_ChildrenRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ChildrenRegion() *gui.QRegion {
	defer qt.Recovering("QWidget::childrenRegion")

	if ptr.Pointer() != nil {
		return gui.NewQRegionFromPointer(C.QWidget_ChildrenRegion(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ClearMask() {
	defer qt.Recovering("QWidget::clearMask")

	if ptr.Pointer() != nil {
		C.QWidget_ClearMask(ptr.Pointer())
	}
}

func (ptr *QWidget) ContextMenuPolicy() core.Qt__ContextMenuPolicy {
	defer qt.Recovering("QWidget::contextMenuPolicy")

	if ptr.Pointer() != nil {
		return core.Qt__ContextMenuPolicy(C.QWidget_ContextMenuPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQWidgetDragEnterEvent
func callbackQWidgetDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQWidgetDragLeaveEvent
func callbackQWidgetDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQWidgetDragMoveEvent
func callbackQWidgetDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQWidgetDropEvent
func callbackQWidgetDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQWidgetEnterEvent
func callbackQWidgetEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQWidgetFocusOutEvent
func callbackQWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) FocusPolicy() core.Qt__FocusPolicy {
	defer qt.Recovering("QWidget::focusPolicy")

	if ptr.Pointer() != nil {
		return core.Qt__FocusPolicy(C.QWidget_FocusPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) FrameGeometry() *core.QRect {
	defer qt.Recovering("QWidget::frameGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWidget_FrameGeometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) FrameSize() *core.QSize {
	defer qt.Recovering("QWidget::frameSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWidget_FrameSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) GrabKeyboard() {
	defer qt.Recovering("QWidget::grabKeyboard")

	if ptr.Pointer() != nil {
		C.QWidget_GrabKeyboard(ptr.Pointer())
	}
}

func (ptr *QWidget) GrabMouse() {
	defer qt.Recovering("QWidget::grabMouse")

	if ptr.Pointer() != nil {
		C.QWidget_GrabMouse(ptr.Pointer())
	}
}

func (ptr *QWidget) GrabMouse2(cursor gui.QCursor_ITF) {
	defer qt.Recovering("QWidget::grabMouse")

	if ptr.Pointer() != nil {
		C.QWidget_GrabMouse2(ptr.Pointer(), gui.PointerFromQCursor(cursor))
	}
}

func (ptr *QWidget) HasFocus() bool {
	defer qt.Recovering("QWidget::hasFocus")

	if ptr.Pointer() != nil {
		return C.QWidget_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQWidgetHideEvent
func callbackQWidgetHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) InputMethodHints() core.Qt__InputMethodHint {
	defer qt.Recovering("QWidget::inputMethodHints")

	if ptr.Pointer() != nil {
		return core.Qt__InputMethodHint(C.QWidget_InputMethodHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) IsActiveWindow() bool {
	defer qt.Recovering("QWidget::isActiveWindow")

	if ptr.Pointer() != nil {
		return C.QWidget_IsActiveWindow(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsFullScreen() bool {
	defer qt.Recovering("QWidget::isFullScreen")

	if ptr.Pointer() != nil {
		return C.QWidget_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsMaximized() bool {
	defer qt.Recovering("QWidget::isMaximized")

	if ptr.Pointer() != nil {
		return C.QWidget_IsMaximized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsMinimized() bool {
	defer qt.Recovering("QWidget::isMinimized")

	if ptr.Pointer() != nil {
		return C.QWidget_IsMinimized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsWindowModified() bool {
	defer qt.Recovering("QWidget::isWindowModified")

	if ptr.Pointer() != nil {
		return C.QWidget_IsWindowModified(ptr.Pointer()) != 0
	}
	return false
}

func QWidget_KeyboardGrabber() *QWidget {
	defer qt.Recovering("QWidget::keyboardGrabber")

	return NewQWidgetFromPointer(C.QWidget_QWidget_KeyboardGrabber())
}

func (ptr *QWidget) LayoutDirection() core.Qt__LayoutDirection {
	defer qt.Recovering("QWidget::layoutDirection")

	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QWidget_LayoutDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQWidgetLeaveEvent
func callbackQWidgetLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) MapFromGlobal(pos core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QWidget::mapFromGlobal")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWidget_MapFromGlobal(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QWidget) MapToGlobal(pos core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QWidget::mapToGlobal")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWidget_MapToGlobal(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QWidget) MaximumSize() *core.QSize {
	defer qt.Recovering("QWidget::maximumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWidget_MaximumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) MinimumSize() *core.QSize {
	defer qt.Recovering("QWidget::minimumSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWidget_MinimumSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWidget_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func QWidget_MouseGrabber() *QWidget {
	defer qt.Recovering("QWidget::mouseGrabber")

	return NewQWidgetFromPointer(C.QWidget_QWidget_MouseGrabber())
}

func (ptr *QWidget) Move(v core.QPoint_ITF) {
	defer qt.Recovering("QWidget::move")

	if ptr.Pointer() != nil {
		C.QWidget_Move(ptr.Pointer(), core.PointerFromQPoint(v))
	}
}

func (ptr *QWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQWidgetMoveEvent
func callbackQWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) NormalGeometry() *core.QRect {
	defer qt.Recovering("QWidget::normalGeometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWidget_NormalGeometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) PaintEngine() *gui.QPaintEngine {
	defer qt.Recovering("QWidget::paintEngine")

	if ptr.Pointer() != nil {
		return gui.NewQPaintEngineFromPointer(C.QWidget_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQWidgetPaintEvent
func callbackQWidgetPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) Pos() *core.QPoint {
	defer qt.Recovering("QWidget::pos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWidget_Pos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ReleaseKeyboard() {
	defer qt.Recovering("QWidget::releaseKeyboard")

	if ptr.Pointer() != nil {
		C.QWidget_ReleaseKeyboard(ptr.Pointer())
	}
}

func (ptr *QWidget) ReleaseMouse() {
	defer qt.Recovering("QWidget::releaseMouse")

	if ptr.Pointer() != nil {
		C.QWidget_ReleaseMouse(ptr.Pointer())
	}
}

func (ptr *QWidget) Resize(v core.QSize_ITF) {
	defer qt.Recovering("QWidget::resize")

	if ptr.Pointer() != nil {
		C.QWidget_Resize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetAcceptDrops(on bool) {
	defer qt.Recovering("QWidget::setAcceptDrops")

	if ptr.Pointer() != nil {
		C.QWidget_SetAcceptDrops(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWidget) SetAccessibleDescription(description string) {
	defer qt.Recovering("QWidget::setAccessibleDescription")

	if ptr.Pointer() != nil {
		C.QWidget_SetAccessibleDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QWidget) SetAccessibleName(name string) {
	defer qt.Recovering("QWidget::setAccessibleName")

	if ptr.Pointer() != nil {
		C.QWidget_SetAccessibleName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QWidget) SetAutoFillBackground(enabled bool) {
	defer qt.Recovering("QWidget::setAutoFillBackground")

	if ptr.Pointer() != nil {
		C.QWidget_SetAutoFillBackground(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QWidget) SetContextMenuPolicy(policy core.Qt__ContextMenuPolicy) {
	defer qt.Recovering("QWidget::setContextMenuPolicy")

	if ptr.Pointer() != nil {
		C.QWidget_SetContextMenuPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QWidget) SetCursor(v gui.QCursor_ITF) {
	defer qt.Recovering("QWidget::setCursor")

	if ptr.Pointer() != nil {
		C.QWidget_SetCursor(ptr.Pointer(), gui.PointerFromQCursor(v))
	}
}

func (ptr *QWidget) SetEnabled(v bool) {
	defer qt.Recovering("QWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QWidget_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QWidget) SetFixedSize2(w int, h int) {
	defer qt.Recovering("QWidget::setFixedSize")

	if ptr.Pointer() != nil {
		C.QWidget_SetFixedSize2(ptr.Pointer(), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) SetFocusPolicy(policy core.Qt__FocusPolicy) {
	defer qt.Recovering("QWidget::setFocusPolicy")

	if ptr.Pointer() != nil {
		C.QWidget_SetFocusPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QWidget) SetFont(v gui.QFont_ITF) {
	defer qt.Recovering("QWidget::setFont")

	if ptr.Pointer() != nil {
		C.QWidget_SetFont(ptr.Pointer(), gui.PointerFromQFont(v))
	}
}

func (ptr *QWidget) SetGeometry(v core.QRect_ITF) {
	defer qt.Recovering("QWidget::setGeometry")

	if ptr.Pointer() != nil {
		C.QWidget_SetGeometry(ptr.Pointer(), core.PointerFromQRect(v))
	}
}

func (ptr *QWidget) SetInputMethodHints(hints core.Qt__InputMethodHint) {
	defer qt.Recovering("QWidget::setInputMethodHints")

	if ptr.Pointer() != nil {
		C.QWidget_SetInputMethodHints(ptr.Pointer(), C.int(hints))
	}
}

func (ptr *QWidget) SetLayout(layout QLayout_ITF) {
	defer qt.Recovering("QWidget::setLayout")

	if ptr.Pointer() != nil {
		C.QWidget_SetLayout(ptr.Pointer(), PointerFromQLayout(layout))
	}
}

func (ptr *QWidget) SetLayoutDirection(direction core.Qt__LayoutDirection) {
	defer qt.Recovering("QWidget::setLayoutDirection")

	if ptr.Pointer() != nil {
		C.QWidget_SetLayoutDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QWidget) SetLocale(locale core.QLocale_ITF) {
	defer qt.Recovering("QWidget::setLocale")

	if ptr.Pointer() != nil {
		C.QWidget_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QWidget) SetMask(bitmap gui.QBitmap_ITF) {
	defer qt.Recovering("QWidget::setMask")

	if ptr.Pointer() != nil {
		C.QWidget_SetMask(ptr.Pointer(), gui.PointerFromQBitmap(bitmap))
	}
}

func (ptr *QWidget) SetMask2(region gui.QRegion_ITF) {
	defer qt.Recovering("QWidget::setMask")

	if ptr.Pointer() != nil {
		C.QWidget_SetMask2(ptr.Pointer(), gui.PointerFromQRegion(region))
	}
}

func (ptr *QWidget) SetMaximumHeight(maxh int) {
	defer qt.Recovering("QWidget::setMaximumHeight")

	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumHeight(ptr.Pointer(), C.int(maxh))
	}
}

func (ptr *QWidget) SetMaximumWidth(maxw int) {
	defer qt.Recovering("QWidget::setMaximumWidth")

	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumWidth(ptr.Pointer(), C.int(maxw))
	}
}

func (ptr *QWidget) SetMinimumHeight(minh int) {
	defer qt.Recovering("QWidget::setMinimumHeight")

	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumHeight(ptr.Pointer(), C.int(minh))
	}
}

func (ptr *QWidget) SetMinimumWidth(minw int) {
	defer qt.Recovering("QWidget::setMinimumWidth")

	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumWidth(ptr.Pointer(), C.int(minw))
	}
}

func (ptr *QWidget) SetPalette(v gui.QPalette_ITF) {
	defer qt.Recovering("QWidget::setPalette")

	if ptr.Pointer() != nil {
		C.QWidget_SetPalette(ptr.Pointer(), gui.PointerFromQPalette(v))
	}
}

func (ptr *QWidget) SetSizePolicy(v QSizePolicy_ITF) {
	defer qt.Recovering("QWidget::setSizePolicy")

	if ptr.Pointer() != nil {
		C.QWidget_SetSizePolicy(ptr.Pointer(), PointerFromQSizePolicy(v))
	}
}

func (ptr *QWidget) SetStatusTip(v string) {
	defer qt.Recovering("QWidget::setStatusTip")

	if ptr.Pointer() != nil {
		C.QWidget_SetStatusTip(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QWidget_SetStyleSheet(ptr.Pointer(), C.CString(styleSheet))
	}
}

func (ptr *QWidget) SetToolTip(v string) {
	defer qt.Recovering("QWidget::setToolTip")

	if ptr.Pointer() != nil {
		C.QWidget_SetToolTip(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) SetToolTipDuration(msec int) {
	defer qt.Recovering("QWidget::setToolTipDuration")

	if ptr.Pointer() != nil {
		C.QWidget_SetToolTipDuration(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QWidget) SetUpdatesEnabled(enable bool) {
	defer qt.Recovering("QWidget::setUpdatesEnabled")

	if ptr.Pointer() != nil {
		C.QWidget_SetUpdatesEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQWidgetSetVisible
func callbackQWidgetSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QWidget) SetWhatsThis(v string) {
	defer qt.Recovering("QWidget::setWhatsThis")

	if ptr.Pointer() != nil {
		C.QWidget_SetWhatsThis(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) SetWindowFilePath(filePath string) {
	defer qt.Recovering("QWidget::setWindowFilePath")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowFilePath(ptr.Pointer(), C.CString(filePath))
	}
}

func (ptr *QWidget) SetWindowFlags(ty core.Qt__WindowType) {
	defer qt.Recovering("QWidget::setWindowFlags")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowFlags(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QWidget) SetWindowIcon(icon gui.QIcon_ITF) {
	defer qt.Recovering("QWidget::setWindowIcon")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowIcon(ptr.Pointer(), gui.PointerFromQIcon(icon))
	}
}

func (ptr *QWidget) SetWindowIconText(v string) {
	defer qt.Recovering("QWidget::setWindowIconText")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowIconText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) SetWindowModality(windowModality core.Qt__WindowModality) {
	defer qt.Recovering("QWidget::setWindowModality")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowModality(ptr.Pointer(), C.int(windowModality))
	}
}

func (ptr *QWidget) SetWindowModified(v bool) {
	defer qt.Recovering("QWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QWidget) SetWindowOpacity(level float64) {
	defer qt.Recovering("QWidget::setWindowOpacity")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowOpacity(ptr.Pointer(), C.double(level))
	}
}

func (ptr *QWidget) SetWindowState(windowState core.Qt__WindowState) {
	defer qt.Recovering("QWidget::setWindowState")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowState(ptr.Pointer(), C.int(windowState))
	}
}

func (ptr *QWidget) SetWindowTitle(v string) {
	defer qt.Recovering("QWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowTitle(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQWidgetShowEvent
func callbackQWidgetShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) SizeIncrement() *core.QSize {
	defer qt.Recovering("QWidget::sizeIncrement")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWidget_SizeIncrement(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) StatusTip() string {
	defer qt.Recovering("QWidget::statusTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_StatusTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) StyleSheet() string {
	defer qt.Recovering("QWidget::styleSheet")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_StyleSheet(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) ToolTip() string {
	defer qt.Recovering("QWidget::toolTip")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_ToolTip(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) ToolTipDuration() int {
	defer qt.Recovering("QWidget::toolTipDuration")

	if ptr.Pointer() != nil {
		return int(C.QWidget_ToolTipDuration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) UnsetCursor() {
	defer qt.Recovering("QWidget::unsetCursor")

	if ptr.Pointer() != nil {
		C.QWidget_UnsetCursor(ptr.Pointer())
	}
}

func (ptr *QWidget) UnsetLayoutDirection() {
	defer qt.Recovering("QWidget::unsetLayoutDirection")

	if ptr.Pointer() != nil {
		C.QWidget_UnsetLayoutDirection(ptr.Pointer())
	}
}

func (ptr *QWidget) UnsetLocale() {
	defer qt.Recovering("QWidget::unsetLocale")

	if ptr.Pointer() != nil {
		C.QWidget_UnsetLocale(ptr.Pointer())
	}
}

func (ptr *QWidget) WhatsThis() string {
	defer qt.Recovering("QWidget::whatsThis")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WhatsThis(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) WindowFilePath() string {
	defer qt.Recovering("QWidget::windowFilePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowFilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) WindowIcon() *gui.QIcon {
	defer qt.Recovering("QWidget::windowIcon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QWidget_WindowIcon(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) WindowIconText() string {
	defer qt.Recovering("QWidget::windowIconText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowIconText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) WindowModality() core.Qt__WindowModality {
	defer qt.Recovering("QWidget::windowModality")

	if ptr.Pointer() != nil {
		return core.Qt__WindowModality(C.QWidget_WindowModality(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) WindowOpacity() float64 {
	defer qt.Recovering("QWidget::windowOpacity")

	if ptr.Pointer() != nil {
		return float64(C.QWidget_WindowOpacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) WindowTitle() string {
	defer qt.Recovering("QWidget::windowTitle")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowTitle(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) X() int {
	defer qt.Recovering("QWidget::x")

	if ptr.Pointer() != nil {
		return int(C.QWidget_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) Y() int {
	defer qt.Recovering("QWidget::y")

	if ptr.Pointer() != nil {
		return int(C.QWidget_Y(ptr.Pointer()))
	}
	return 0
}

func NewQWidget(parent QWidget_ITF, f core.Qt__WindowType) *QWidget {
	defer qt.Recovering("QWidget::QWidget")

	return NewQWidgetFromPointer(C.QWidget_NewQWidget(PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QWidget) AddAction(action QAction_ITF) {
	defer qt.Recovering("QWidget::addAction")

	if ptr.Pointer() != nil {
		C.QWidget_AddAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QWidget) AdjustSize() {
	defer qt.Recovering("QWidget::adjustSize")

	if ptr.Pointer() != nil {
		C.QWidget_AdjustSize(ptr.Pointer())
	}
}

func (ptr *QWidget) BackgroundRole() gui.QPalette__ColorRole {
	defer qt.Recovering("QWidget::backgroundRole")

	if ptr.Pointer() != nil {
		return gui.QPalette__ColorRole(C.QWidget_BackgroundRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) BackingStore() *gui.QBackingStore {
	defer qt.Recovering("QWidget::backingStore")

	if ptr.Pointer() != nil {
		return gui.NewQBackingStoreFromPointer(C.QWidget_BackingStore(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ChildAt2(p core.QPoint_ITF) *QWidget {
	defer qt.Recovering("QWidget::childAt")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_ChildAt2(ptr.Pointer(), core.PointerFromQPoint(p)))
	}
	return nil
}

func (ptr *QWidget) ChildAt(x int, y int) *QWidget {
	defer qt.Recovering("QWidget::childAt")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_ChildAt(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QWidget) ClearFocus() {
	defer qt.Recovering("QWidget::clearFocus")

	if ptr.Pointer() != nil {
		C.QWidget_ClearFocus(ptr.Pointer())
	}
}

func (ptr *QWidget) Close() bool {
	defer qt.Recovering("QWidget::close")

	if ptr.Pointer() != nil {
		return C.QWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQWidgetCloseEvent
func callbackQWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ContentsRect() *core.QRect {
	defer qt.Recovering("QWidget::contentsRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWidget_ContentsRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQWidgetContextMenuEvent
func callbackQWidgetContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectCustomContextMenuRequested(f func(pos *core.QPoint)) {
	defer qt.Recovering("connect QWidget::customContextMenuRequested")

	if ptr.Pointer() != nil {
		C.QWidget_ConnectCustomContextMenuRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "customContextMenuRequested", f)
	}
}

func (ptr *QWidget) DisconnectCustomContextMenuRequested() {
	defer qt.Recovering("disconnect QWidget::customContextMenuRequested")

	if ptr.Pointer() != nil {
		C.QWidget_DisconnectCustomContextMenuRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "customContextMenuRequested")
	}
}

//export callbackQWidgetCustomContextMenuRequested
func callbackQWidgetCustomContextMenuRequested(ptrName *C.char, pos unsafe.Pointer) {
	defer qt.Recovering("callback QWidget::customContextMenuRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "customContextMenuRequested"); signal != nil {
		signal.(func(*core.QPoint))(core.NewQPointFromPointer(pos))
	}

}

func (ptr *QWidget) EnsurePolished() {
	defer qt.Recovering("QWidget::ensurePolished")

	if ptr.Pointer() != nil {
		C.QWidget_EnsurePolished(ptr.Pointer())
	}
}

func (ptr *QWidget) FocusProxy() *QWidget {
	defer qt.Recovering("QWidget::focusProxy")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_FocusProxy(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) FocusWidget() *QWidget {
	defer qt.Recovering("QWidget::focusWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_FocusWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ForegroundRole() gui.QPalette__ColorRole {
	defer qt.Recovering("QWidget::foregroundRole")

	if ptr.Pointer() != nil {
		return gui.QPalette__ColorRole(C.QWidget_ForegroundRole(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) Geometry() *core.QRect {
	defer qt.Recovering("QWidget::geometry")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWidget_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) GetContentsMargins(left int, top int, right int, bottom int) {
	defer qt.Recovering("QWidget::getContentsMargins")

	if ptr.Pointer() != nil {
		C.QWidget_GetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QWidget) GrabGesture(gesture core.Qt__GestureType, flags core.Qt__GestureFlag) {
	defer qt.Recovering("QWidget::grabGesture")

	if ptr.Pointer() != nil {
		C.QWidget_GrabGesture(ptr.Pointer(), C.int(gesture), C.int(flags))
	}
}

func (ptr *QWidget) GrabShortcut(key gui.QKeySequence_ITF, context core.Qt__ShortcutContext) int {
	defer qt.Recovering("QWidget::grabShortcut")

	if ptr.Pointer() != nil {
		return int(C.QWidget_GrabShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(key), C.int(context)))
	}
	return 0
}

func (ptr *QWidget) GraphicsEffect() *QGraphicsEffect {
	defer qt.Recovering("QWidget::graphicsEffect")

	if ptr.Pointer() != nil {
		return NewQGraphicsEffectFromPointer(C.QWidget_GraphicsEffect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) GraphicsProxyWidget() *QGraphicsProxyWidget {
	defer qt.Recovering("QWidget::graphicsProxyWidget")

	if ptr.Pointer() != nil {
		return NewQGraphicsProxyWidgetFromPointer(C.QWidget_GraphicsProxyWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) HasMouseTracking() bool {
	defer qt.Recovering("QWidget::hasMouseTracking")

	if ptr.Pointer() != nil {
		return C.QWidget_HasMouseTracking(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) Height() int {
	defer qt.Recovering("QWidget::height")

	if ptr.Pointer() != nil {
		return int(C.QWidget_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) HeightForWidth(w int) int {
	defer qt.Recovering("QWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QWidget_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QWidget) Hide() {
	defer qt.Recovering("QWidget::hide")

	if ptr.Pointer() != nil {
		C.QWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQWidgetInitPainter
func callbackQWidgetInitPainter(ptrName *C.char, painter unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQWidgetInputMethodEvent
func callbackQWidgetInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QWidget_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QWidget) InsertAction(before QAction_ITF, action QAction_ITF) {
	defer qt.Recovering("QWidget::insertAction")

	if ptr.Pointer() != nil {
		C.QWidget_InsertAction(ptr.Pointer(), PointerFromQAction(before), PointerFromQAction(action))
	}
}

func (ptr *QWidget) IsAncestorOf(child QWidget_ITF) bool {
	defer qt.Recovering("QWidget::isAncestorOf")

	if ptr.Pointer() != nil {
		return C.QWidget_IsAncestorOf(ptr.Pointer(), PointerFromQWidget(child)) != 0
	}
	return false
}

func (ptr *QWidget) IsEnabled() bool {
	defer qt.Recovering("QWidget::isEnabled")

	if ptr.Pointer() != nil {
		return C.QWidget_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsEnabledTo(ancestor QWidget_ITF) bool {
	defer qt.Recovering("QWidget::isEnabledTo")

	if ptr.Pointer() != nil {
		return C.QWidget_IsEnabledTo(ptr.Pointer(), PointerFromQWidget(ancestor)) != 0
	}
	return false
}

func (ptr *QWidget) IsHidden() bool {
	defer qt.Recovering("QWidget::isHidden")

	if ptr.Pointer() != nil {
		return C.QWidget_IsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsModal() bool {
	defer qt.Recovering("QWidget::isModal")

	if ptr.Pointer() != nil {
		return C.QWidget_IsModal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsVisible() bool {
	defer qt.Recovering("QWidget::isVisible")

	if ptr.Pointer() != nil {
		return C.QWidget_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) IsVisibleTo(ancestor QWidget_ITF) bool {
	defer qt.Recovering("QWidget::isVisibleTo")

	if ptr.Pointer() != nil {
		return C.QWidget_IsVisibleTo(ptr.Pointer(), PointerFromQWidget(ancestor)) != 0
	}
	return false
}

func (ptr *QWidget) IsWindow() bool {
	defer qt.Recovering("QWidget::isWindow")

	if ptr.Pointer() != nil {
		return C.QWidget_IsWindow(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQWidgetKeyPressEvent
func callbackQWidgetKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQWidgetKeyReleaseEvent
func callbackQWidgetKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) Layout() *QLayout {
	defer qt.Recovering("QWidget::layout")

	if ptr.Pointer() != nil {
		return NewQLayoutFromPointer(C.QWidget_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) Lower() {
	defer qt.Recovering("QWidget::lower")

	if ptr.Pointer() != nil {
		C.QWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QWidget) MapFrom(parent QWidget_ITF, pos core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QWidget::mapFrom")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWidget_MapFrom(ptr.Pointer(), PointerFromQWidget(parent), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QWidget) MapFromParent(pos core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QWidget::mapFromParent")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWidget_MapFromParent(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QWidget) MapTo(parent QWidget_ITF, pos core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QWidget::mapTo")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWidget_MapTo(ptr.Pointer(), PointerFromQWidget(parent), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QWidget) MapToParent(pos core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QWidget::mapToParent")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWidget_MapToParent(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QWidget) Mask() *gui.QRegion {
	defer qt.Recovering("QWidget::mask")

	if ptr.Pointer() != nil {
		return gui.NewQRegionFromPointer(C.QWidget_Mask(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) MaximumHeight() int {
	defer qt.Recovering("QWidget::maximumHeight")

	if ptr.Pointer() != nil {
		return int(C.QWidget_MaximumHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) MaximumWidth() int {
	defer qt.Recovering("QWidget::maximumWidth")

	if ptr.Pointer() != nil {
		return int(C.QWidget_MaximumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) MinimumHeight() int {
	defer qt.Recovering("QWidget::minimumHeight")

	if ptr.Pointer() != nil {
		return int(C.QWidget_MinimumHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) MinimumWidth() int {
	defer qt.Recovering("QWidget::minimumWidth")

	if ptr.Pointer() != nil {
		return int(C.QWidget_MinimumWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQWidgetMouseDoubleClickEvent
func callbackQWidgetMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQWidgetMouseMoveEvent
func callbackQWidgetMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQWidgetMousePressEvent
func callbackQWidgetMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQWidgetMouseReleaseEvent
func callbackQWidgetMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) Move2(x int, y int) {
	defer qt.Recovering("QWidget::move")

	if ptr.Pointer() != nil {
		C.QWidget_Move2(ptr.Pointer(), C.int(x), C.int(y))
	}
}

func (ptr *QWidget) NativeParentWidget() *QWidget {
	defer qt.Recovering("QWidget::nativeParentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_NativeParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) NextInFocusChain() *QWidget {
	defer qt.Recovering("QWidget::nextInFocusChain")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_NextInFocusChain(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) OverrideWindowFlags(flags core.Qt__WindowType) {
	defer qt.Recovering("QWidget::overrideWindowFlags")

	if ptr.Pointer() != nil {
		C.QWidget_OverrideWindowFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QWidget) ParentWidget() *QWidget {
	defer qt.Recovering("QWidget::parentWidget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_ParentWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) PreviousInFocusChain() *QWidget {
	defer qt.Recovering("QWidget::previousInFocusChain")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_PreviousInFocusChain(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) Raise() {
	defer qt.Recovering("QWidget::raise")

	if ptr.Pointer() != nil {
		C.QWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QWidget) Rect() *core.QRect {
	defer qt.Recovering("QWidget::rect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QWidget_Rect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ReleaseShortcut(id int) {
	defer qt.Recovering("QWidget::releaseShortcut")

	if ptr.Pointer() != nil {
		C.QWidget_ReleaseShortcut(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QWidget) RemoveAction(action QAction_ITF) {
	defer qt.Recovering("QWidget::removeAction")

	if ptr.Pointer() != nil {
		C.QWidget_RemoveAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QWidget) Render(target gui.QPaintDevice_ITF, targetOffset core.QPoint_ITF, sourceRegion gui.QRegion_ITF, renderFlags QWidget__RenderFlag) {
	defer qt.Recovering("QWidget::render")

	if ptr.Pointer() != nil {
		C.QWidget_Render(ptr.Pointer(), gui.PointerFromQPaintDevice(target), core.PointerFromQPoint(targetOffset), gui.PointerFromQRegion(sourceRegion), C.int(renderFlags))
	}
}

func (ptr *QWidget) Render2(painter gui.QPainter_ITF, targetOffset core.QPoint_ITF, sourceRegion gui.QRegion_ITF, renderFlags QWidget__RenderFlag) {
	defer qt.Recovering("QWidget::render")

	if ptr.Pointer() != nil {
		C.QWidget_Render2(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQPoint(targetOffset), gui.PointerFromQRegion(sourceRegion), C.int(renderFlags))
	}
}

func (ptr *QWidget) Repaint() {
	defer qt.Recovering("QWidget::repaint")

	if ptr.Pointer() != nil {
		C.QWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QWidget) Repaint3(rect core.QRect_ITF) {
	defer qt.Recovering("QWidget::repaint")

	if ptr.Pointer() != nil {
		C.QWidget_Repaint3(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWidget) Repaint4(rgn gui.QRegion_ITF) {
	defer qt.Recovering("QWidget::repaint")

	if ptr.Pointer() != nil {
		C.QWidget_Repaint4(ptr.Pointer(), gui.PointerFromQRegion(rgn))
	}
}

func (ptr *QWidget) Repaint2(x int, y int, w int, h int) {
	defer qt.Recovering("QWidget::repaint")

	if ptr.Pointer() != nil {
		C.QWidget_Repaint2(ptr.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) Resize2(w int, h int) {
	defer qt.Recovering("QWidget::resize")

	if ptr.Pointer() != nil {
		C.QWidget_Resize2(ptr.Pointer(), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQWidgetResizeEvent
func callbackQWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) RestoreGeometry(geometry core.QByteArray_ITF) bool {
	defer qt.Recovering("QWidget::restoreGeometry")

	if ptr.Pointer() != nil {
		return C.QWidget_RestoreGeometry(ptr.Pointer(), core.PointerFromQByteArray(geometry)) != 0
	}
	return false
}

func (ptr *QWidget) SaveGeometry() *core.QByteArray {
	defer qt.Recovering("QWidget::saveGeometry")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QWidget_SaveGeometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) Scroll(dx int, dy int) {
	defer qt.Recovering("QWidget::scroll")

	if ptr.Pointer() != nil {
		C.QWidget_Scroll(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QWidget) Scroll2(dx int, dy int, r core.QRect_ITF) {
	defer qt.Recovering("QWidget::scroll")

	if ptr.Pointer() != nil {
		C.QWidget_Scroll2(ptr.Pointer(), C.int(dx), C.int(dy), core.PointerFromQRect(r))
	}
}

func (ptr *QWidget) SetAttribute(attribute core.Qt__WidgetAttribute, on bool) {
	defer qt.Recovering("QWidget::setAttribute")

	if ptr.Pointer() != nil {
		C.QWidget_SetAttribute(ptr.Pointer(), C.int(attribute), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QWidget) SetBackgroundRole(role gui.QPalette__ColorRole) {
	defer qt.Recovering("QWidget::setBackgroundRole")

	if ptr.Pointer() != nil {
		C.QWidget_SetBackgroundRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QWidget) SetBaseSize(v core.QSize_ITF) {
	defer qt.Recovering("QWidget::setBaseSize")

	if ptr.Pointer() != nil {
		C.QWidget_SetBaseSize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetBaseSize2(basew int, baseh int) {
	defer qt.Recovering("QWidget::setBaseSize")

	if ptr.Pointer() != nil {
		C.QWidget_SetBaseSize2(ptr.Pointer(), C.int(basew), C.int(baseh))
	}
}

func (ptr *QWidget) SetContentsMargins2(margins core.QMargins_ITF) {
	defer qt.Recovering("QWidget::setContentsMargins")

	if ptr.Pointer() != nil {
		C.QWidget_SetContentsMargins2(ptr.Pointer(), core.PointerFromQMargins(margins))
	}
}

func (ptr *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	defer qt.Recovering("QWidget::setContentsMargins")

	if ptr.Pointer() != nil {
		C.QWidget_SetContentsMargins(ptr.Pointer(), C.int(left), C.int(top), C.int(right), C.int(bottom))
	}
}

func (ptr *QWidget) SetDisabled(disable bool) {
	defer qt.Recovering("QWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QWidget_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *QWidget) SetFixedHeight(h int) {
	defer qt.Recovering("QWidget::setFixedHeight")

	if ptr.Pointer() != nil {
		C.QWidget_SetFixedHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QWidget) SetFixedSize(s core.QSize_ITF) {
	defer qt.Recovering("QWidget::setFixedSize")

	if ptr.Pointer() != nil {
		C.QWidget_SetFixedSize(ptr.Pointer(), core.PointerFromQSize(s))
	}
}

func (ptr *QWidget) SetFixedWidth(w int) {
	defer qt.Recovering("QWidget::setFixedWidth")

	if ptr.Pointer() != nil {
		C.QWidget_SetFixedWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QWidget) SetFocus2() {
	defer qt.Recovering("QWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QWidget) SetFocus(reason core.Qt__FocusReason) {
	defer qt.Recovering("QWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QWidget_SetFocus(ptr.Pointer(), C.int(reason))
	}
}

func (ptr *QWidget) SetFocusProxy(w QWidget_ITF) {
	defer qt.Recovering("QWidget::setFocusProxy")

	if ptr.Pointer() != nil {
		C.QWidget_SetFocusProxy(ptr.Pointer(), PointerFromQWidget(w))
	}
}

func (ptr *QWidget) SetForegroundRole(role gui.QPalette__ColorRole) {
	defer qt.Recovering("QWidget::setForegroundRole")

	if ptr.Pointer() != nil {
		C.QWidget_SetForegroundRole(ptr.Pointer(), C.int(role))
	}
}

func (ptr *QWidget) SetGeometry2(x int, y int, w int, h int) {
	defer qt.Recovering("QWidget::setGeometry")

	if ptr.Pointer() != nil {
		C.QWidget_SetGeometry2(ptr.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) SetGraphicsEffect(effect QGraphicsEffect_ITF) {
	defer qt.Recovering("QWidget::setGraphicsEffect")

	if ptr.Pointer() != nil {
		C.QWidget_SetGraphicsEffect(ptr.Pointer(), PointerFromQGraphicsEffect(effect))
	}
}

func (ptr *QWidget) SetHidden(hidden bool) {
	defer qt.Recovering("QWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QWidget_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *QWidget) SetMaximumSize(v core.QSize_ITF) {
	defer qt.Recovering("QWidget::setMaximumSize")

	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumSize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetMaximumSize2(maxw int, maxh int) {
	defer qt.Recovering("QWidget::setMaximumSize")

	if ptr.Pointer() != nil {
		C.QWidget_SetMaximumSize2(ptr.Pointer(), C.int(maxw), C.int(maxh))
	}
}

func (ptr *QWidget) SetMinimumSize(v core.QSize_ITF) {
	defer qt.Recovering("QWidget::setMinimumSize")

	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetMinimumSize2(minw int, minh int) {
	defer qt.Recovering("QWidget::setMinimumSize")

	if ptr.Pointer() != nil {
		C.QWidget_SetMinimumSize2(ptr.Pointer(), C.int(minw), C.int(minh))
	}
}

func (ptr *QWidget) SetMouseTracking(enable bool) {
	defer qt.Recovering("QWidget::setMouseTracking")

	if ptr.Pointer() != nil {
		C.QWidget_SetMouseTracking(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetParent(parent QWidget_ITF) {
	defer qt.Recovering("QWidget::setParent")

	if ptr.Pointer() != nil {
		C.QWidget_SetParent(ptr.Pointer(), PointerFromQWidget(parent))
	}
}

func (ptr *QWidget) SetParent2(parent QWidget_ITF, f core.Qt__WindowType) {
	defer qt.Recovering("QWidget::setParent")

	if ptr.Pointer() != nil {
		C.QWidget_SetParent2(ptr.Pointer(), PointerFromQWidget(parent), C.int(f))
	}
}

func (ptr *QWidget) SetShortcutAutoRepeat(id int, enable bool) {
	defer qt.Recovering("QWidget::setShortcutAutoRepeat")

	if ptr.Pointer() != nil {
		C.QWidget_SetShortcutAutoRepeat(ptr.Pointer(), C.int(id), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetShortcutEnabled(id int, enable bool) {
	defer qt.Recovering("QWidget::setShortcutEnabled")

	if ptr.Pointer() != nil {
		C.QWidget_SetShortcutEnabled(ptr.Pointer(), C.int(id), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QWidget) SetSizeIncrement(v core.QSize_ITF) {
	defer qt.Recovering("QWidget::setSizeIncrement")

	if ptr.Pointer() != nil {
		C.QWidget_SetSizeIncrement(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QWidget) SetSizeIncrement2(w int, h int) {
	defer qt.Recovering("QWidget::setSizeIncrement")

	if ptr.Pointer() != nil {
		C.QWidget_SetSizeIncrement2(ptr.Pointer(), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) SetSizePolicy2(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy) {
	defer qt.Recovering("QWidget::setSizePolicy")

	if ptr.Pointer() != nil {
		C.QWidget_SetSizePolicy2(ptr.Pointer(), C.int(horizontal), C.int(vertical))
	}
}

func (ptr *QWidget) SetStyle(style QStyle_ITF) {
	defer qt.Recovering("QWidget::setStyle")

	if ptr.Pointer() != nil {
		C.QWidget_SetStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

func QWidget_SetTabOrder(first QWidget_ITF, second QWidget_ITF) {
	defer qt.Recovering("QWidget::setTabOrder")

	C.QWidget_QWidget_SetTabOrder(PointerFromQWidget(first), PointerFromQWidget(second))
}

func (ptr *QWidget) SetWindowRole(role string) {
	defer qt.Recovering("QWidget::setWindowRole")

	if ptr.Pointer() != nil {
		C.QWidget_SetWindowRole(ptr.Pointer(), C.CString(role))
	}
}

func (ptr *QWidget) Show() {
	defer qt.Recovering("QWidget::show")

	if ptr.Pointer() != nil {
		C.QWidget_Show(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowFullScreen() {
	defer qt.Recovering("QWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowMaximized() {
	defer qt.Recovering("QWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowMinimized() {
	defer qt.Recovering("QWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QWidget) ShowNormal() {
	defer qt.Recovering("QWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QWidget) Size() *core.QSize {
	defer qt.Recovering("QWidget::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QWidget_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) StackUnder(w QWidget_ITF) {
	defer qt.Recovering("QWidget::stackUnder")

	if ptr.Pointer() != nil {
		C.QWidget_StackUnder(ptr.Pointer(), PointerFromQWidget(w))
	}
}

func (ptr *QWidget) Style() *QStyle {
	defer qt.Recovering("QWidget::style")

	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QWidget_Style(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQWidgetTabletEvent
func callbackQWidgetTabletEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) TestAttribute(attribute core.Qt__WidgetAttribute) bool {
	defer qt.Recovering("QWidget::testAttribute")

	if ptr.Pointer() != nil {
		return C.QWidget_TestAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QWidget) UnderMouse() bool {
	defer qt.Recovering("QWidget::underMouse")

	if ptr.Pointer() != nil {
		return C.QWidget_UnderMouse(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) UngrabGesture(gesture core.Qt__GestureType) {
	defer qt.Recovering("QWidget::ungrabGesture")

	if ptr.Pointer() != nil {
		C.QWidget_UngrabGesture(ptr.Pointer(), C.int(gesture))
	}
}

func (ptr *QWidget) Update() {
	defer qt.Recovering("QWidget::update")

	if ptr.Pointer() != nil {
		C.QWidget_Update(ptr.Pointer())
	}
}

func (ptr *QWidget) Update3(rect core.QRect_ITF) {
	defer qt.Recovering("QWidget::update")

	if ptr.Pointer() != nil {
		C.QWidget_Update3(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QWidget) Update4(rgn gui.QRegion_ITF) {
	defer qt.Recovering("QWidget::update")

	if ptr.Pointer() != nil {
		C.QWidget_Update4(ptr.Pointer(), gui.PointerFromQRegion(rgn))
	}
}

func (ptr *QWidget) Update2(x int, y int, w int, h int) {
	defer qt.Recovering("QWidget::update")

	if ptr.Pointer() != nil {
		C.QWidget_Update2(ptr.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h))
	}
}

func (ptr *QWidget) UpdateGeometry() {
	defer qt.Recovering("QWidget::updateGeometry")

	if ptr.Pointer() != nil {
		C.QWidget_UpdateGeometry(ptr.Pointer())
	}
}

func (ptr *QWidget) UpdatesEnabled() bool {
	defer qt.Recovering("QWidget::updatesEnabled")

	if ptr.Pointer() != nil {
		return C.QWidget_UpdatesEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QWidget) VisibleRegion() *gui.QRegion {
	defer qt.Recovering("QWidget::visibleRegion")

	if ptr.Pointer() != nil {
		return gui.NewQRegionFromPointer(C.QWidget_VisibleRegion(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQWidgetWheelEvent
func callbackQWidgetWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) Width() int {
	defer qt.Recovering("QWidget::width")

	if ptr.Pointer() != nil {
		return int(C.QWidget_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) Window() *QWidget {
	defer qt.Recovering("QWidget::window")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QWidget_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) WindowFlags() core.Qt__WindowType {
	defer qt.Recovering("QWidget::windowFlags")

	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWidget_WindowFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) WindowHandle() *gui.QWindow {
	defer qt.Recovering("QWidget::windowHandle")

	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QWidget_WindowHandle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWidget) ConnectWindowIconChanged(f func(icon *gui.QIcon)) {
	defer qt.Recovering("connect QWidget::windowIconChanged")

	if ptr.Pointer() != nil {
		C.QWidget_ConnectWindowIconChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowIconChanged", f)
	}
}

func (ptr *QWidget) DisconnectWindowIconChanged() {
	defer qt.Recovering("disconnect QWidget::windowIconChanged")

	if ptr.Pointer() != nil {
		C.QWidget_DisconnectWindowIconChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowIconChanged")
	}
}

//export callbackQWidgetWindowIconChanged
func callbackQWidgetWindowIconChanged(ptrName *C.char, icon unsafe.Pointer) {
	defer qt.Recovering("callback QWidget::windowIconChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowIconChanged"); signal != nil {
		signal.(func(*gui.QIcon))(gui.NewQIconFromPointer(icon))
	}

}

func (ptr *QWidget) ConnectWindowIconTextChanged(f func(iconText string)) {
	defer qt.Recovering("connect QWidget::windowIconTextChanged")

	if ptr.Pointer() != nil {
		C.QWidget_ConnectWindowIconTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowIconTextChanged", f)
	}
}

func (ptr *QWidget) DisconnectWindowIconTextChanged() {
	defer qt.Recovering("disconnect QWidget::windowIconTextChanged")

	if ptr.Pointer() != nil {
		C.QWidget_DisconnectWindowIconTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowIconTextChanged")
	}
}

//export callbackQWidgetWindowIconTextChanged
func callbackQWidgetWindowIconTextChanged(ptrName *C.char, iconText *C.char) {
	defer qt.Recovering("callback QWidget::windowIconTextChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowIconTextChanged"); signal != nil {
		signal.(func(string))(C.GoString(iconText))
	}

}

func (ptr *QWidget) WindowRole() string {
	defer qt.Recovering("QWidget::windowRole")

	if ptr.Pointer() != nil {
		return C.GoString(C.QWidget_WindowRole(ptr.Pointer()))
	}
	return ""
}

func (ptr *QWidget) WindowState() core.Qt__WindowState {
	defer qt.Recovering("QWidget::windowState")

	if ptr.Pointer() != nil {
		return core.Qt__WindowState(C.QWidget_WindowState(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) ConnectWindowTitleChanged(f func(title string)) {
	defer qt.Recovering("connect QWidget::windowTitleChanged")

	if ptr.Pointer() != nil {
		C.QWidget_ConnectWindowTitleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowTitleChanged", f)
	}
}

func (ptr *QWidget) DisconnectWindowTitleChanged() {
	defer qt.Recovering("disconnect QWidget::windowTitleChanged")

	if ptr.Pointer() != nil {
		C.QWidget_DisconnectWindowTitleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowTitleChanged")
	}
}

//export callbackQWidgetWindowTitleChanged
func callbackQWidgetWindowTitleChanged(ptrName *C.char, title *C.char) {
	defer qt.Recovering("callback QWidget::windowTitleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowTitleChanged"); signal != nil {
		signal.(func(string))(C.GoString(title))
	}

}

func (ptr *QWidget) WindowType() core.Qt__WindowType {
	defer qt.Recovering("QWidget::windowType")

	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QWidget_WindowType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWidget) DestroyQWidget() {
	defer qt.Recovering("QWidget::~QWidget")

	if ptr.Pointer() != nil {
		C.QWidget_DestroyQWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func QWidget_CreateWindowContainer(window gui.QWindow_ITF, parent QWidget_ITF, flags core.Qt__WindowType) *QWidget {
	defer qt.Recovering("QWidget::createWindowContainer")

	return NewQWidgetFromPointer(C.QWidget_QWidget_CreateWindowContainer(gui.PointerFromQWindow(window), PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQWidgetTimerEvent
func callbackQWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQWidgetChildEvent
func callbackQWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQWidgetCustomEvent
func callbackQWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
