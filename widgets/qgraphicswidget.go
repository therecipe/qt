package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsWidget struct {
	QGraphicsObject
	QGraphicsLayoutItem
}

type QGraphicsWidget_ITF interface {
	QGraphicsObject_ITF
	QGraphicsLayoutItem_ITF
	QGraphicsWidget_PTR() *QGraphicsWidget
}

func (p *QGraphicsWidget) Pointer() unsafe.Pointer {
	return p.QGraphicsObject_PTR().Pointer()
}

func (p *QGraphicsWidget) SetPointer(ptr unsafe.Pointer) {
	p.QGraphicsObject_PTR().SetPointer(ptr)
	p.QGraphicsLayoutItem_PTR().SetPointer(ptr)
}

func PointerFromQGraphicsWidget(ptr QGraphicsWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsWidget_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsWidgetFromPointer(ptr unsafe.Pointer) *QGraphicsWidget {
	var n = new(QGraphicsWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsWidget_") {
		n.SetObjectName("QGraphicsWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsWidget) QGraphicsWidget_PTR() *QGraphicsWidget {
	return ptr
}

func (ptr *QGraphicsWidget) AutoFillBackground() bool {
	defer qt.Recovering("QGraphicsWidget::autoFillBackground")

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_AutoFillBackground(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) FocusPolicy() core.Qt__FocusPolicy {
	defer qt.Recovering("QGraphicsWidget::focusPolicy")

	if ptr.Pointer() != nil {
		return core.Qt__FocusPolicy(C.QGraphicsWidget_FocusPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) LayoutDirection() core.Qt__LayoutDirection {
	defer qt.Recovering("QGraphicsWidget::layoutDirection")

	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QGraphicsWidget_LayoutDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) Resize(size core.QSizeF_ITF) {
	defer qt.Recovering("QGraphicsWidget::resize")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Resize(ptr.Pointer(), core.PointerFromQSizeF(size))
	}
}

func (ptr *QGraphicsWidget) SetAutoFillBackground(enabled bool) {
	defer qt.Recovering("QGraphicsWidget::setAutoFillBackground")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetAutoFillBackground(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWidget) SetFocusPolicy(policy core.Qt__FocusPolicy) {
	defer qt.Recovering("QGraphicsWidget::setFocusPolicy")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetFocusPolicy(ptr.Pointer(), C.int(policy))
	}
}

func (ptr *QGraphicsWidget) SetFont(font gui.QFont_ITF) {
	defer qt.Recovering("QGraphicsWidget::setFont")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QGraphicsWidget) SetLayout(layout QGraphicsLayout_ITF) {
	defer qt.Recovering("QGraphicsWidget::setLayout")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetLayout(ptr.Pointer(), PointerFromQGraphicsLayout(layout))
	}
}

func (ptr *QGraphicsWidget) SetLayoutDirection(direction core.Qt__LayoutDirection) {
	defer qt.Recovering("QGraphicsWidget::setLayoutDirection")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetLayoutDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QGraphicsWidget) SetPalette(palette gui.QPalette_ITF) {
	defer qt.Recovering("QGraphicsWidget::setPalette")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetPalette(ptr.Pointer(), gui.PointerFromQPalette(palette))
	}
}

func (ptr *QGraphicsWidget) SetWindowFlags(wFlags core.Qt__WindowType) {
	defer qt.Recovering("QGraphicsWidget::setWindowFlags")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetWindowFlags(ptr.Pointer(), C.int(wFlags))
	}
}

func (ptr *QGraphicsWidget) SetWindowTitle(title string) {
	defer qt.Recovering("QGraphicsWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetWindowTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QGraphicsWidget) UnsetLayoutDirection() {
	defer qt.Recovering("QGraphicsWidget::unsetLayoutDirection")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UnsetLayoutDirection(ptr.Pointer())
	}
}

func (ptr *QGraphicsWidget) WindowFlags() core.Qt__WindowType {
	defer qt.Recovering("QGraphicsWidget::windowFlags")

	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QGraphicsWidget_WindowFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) WindowTitle() string {
	defer qt.Recovering("QGraphicsWidget::windowTitle")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsWidget_WindowTitle(ptr.Pointer()))
	}
	return ""
}

func NewQGraphicsWidget(parent QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QGraphicsWidget {
	defer qt.Recovering("QGraphicsWidget::QGraphicsWidget")

	return NewQGraphicsWidgetFromPointer(C.QGraphicsWidget_NewQGraphicsWidget(PointerFromQGraphicsItem(parent), C.int(wFlags)))
}

func (ptr *QGraphicsWidget) AddAction(action QAction_ITF) {
	defer qt.Recovering("QGraphicsWidget::addAction")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_AddAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QGraphicsWidget) AdjustSize() {
	defer qt.Recovering("QGraphicsWidget::adjustSize")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_AdjustSize(ptr.Pointer())
	}
}

func (ptr *QGraphicsWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQGraphicsWidgetChangeEvent
func callbackQGraphicsWidgetChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) Close() bool {
	defer qt.Recovering("QGraphicsWidget::close")

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQGraphicsWidgetCloseEvent
func callbackQGraphicsWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQGraphicsWidgetFocusInEvent
func callbackQGraphicsWidgetFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQGraphicsWidgetFocusOutEvent
func callbackQGraphicsWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) FocusWidget() *QGraphicsWidget {
	defer qt.Recovering("QGraphicsWidget::focusWidget")

	if ptr.Pointer() != nil {
		return NewQGraphicsWidgetFromPointer(C.QGraphicsWidget_FocusWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWidget) ConnectGeometryChanged(f func()) {
	defer qt.Recovering("connect QGraphicsWidget::geometryChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ConnectGeometryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "geometryChanged", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectGeometryChanged() {
	defer qt.Recovering("disconnect QGraphicsWidget::geometryChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_DisconnectGeometryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "geometryChanged")
	}
}

//export callbackQGraphicsWidgetGeometryChanged
func callbackQGraphicsWidgetGeometryChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWidget::geometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWidget) ConnectGrabKeyboardEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::grabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "grabKeyboardEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectGrabKeyboardEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::grabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "grabKeyboardEvent")
	}
}

//export callbackQGraphicsWidgetGrabKeyboardEvent
func callbackQGraphicsWidgetGrabKeyboardEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::grabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectGrabMouseEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::grabMouseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "grabMouseEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectGrabMouseEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::grabMouseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "grabMouseEvent")
	}
}

//export callbackQGraphicsWidgetGrabMouseEvent
func callbackQGraphicsWidgetGrabMouseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::grabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) GrabShortcut(sequence gui.QKeySequence_ITF, context core.Qt__ShortcutContext) int {
	defer qt.Recovering("QGraphicsWidget::grabShortcut")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsWidget_GrabShortcut(ptr.Pointer(), gui.PointerFromQKeySequence(sequence), C.int(context)))
	}
	return 0
}

func (ptr *QGraphicsWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQGraphicsWidgetHideEvent
func callbackQGraphicsWidgetHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverLeaveEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverLeaveEvent")
	}
}

//export callbackQGraphicsWidgetHoverLeaveEvent
func callbackQGraphicsWidgetHoverLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverMoveEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverMoveEvent")
	}
}

//export callbackQGraphicsWidgetHoverMoveEvent
func callbackQGraphicsWidgetHoverMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectInitStyleOption(f func(option *QStyleOption)) {
	defer qt.Recovering("connect QGraphicsWidget::initStyleOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initStyleOption", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectInitStyleOption() {
	defer qt.Recovering("disconnect QGraphicsWidget::initStyleOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initStyleOption")
	}
}

//export callbackQGraphicsWidgetInitStyleOption
func callbackQGraphicsWidgetInitStyleOption(ptrName *C.char, option unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::initStyleOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "initStyleOption"); signal != nil {
		signal.(func(*QStyleOption))(NewQStyleOptionFromPointer(option))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) InsertAction(before QAction_ITF, action QAction_ITF) {
	defer qt.Recovering("QGraphicsWidget::insertAction")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_InsertAction(ptr.Pointer(), PointerFromQAction(before), PointerFromQAction(action))
	}
}

func (ptr *QGraphicsWidget) IsActiveWindow() bool {
	defer qt.Recovering("QGraphicsWidget::isActiveWindow")

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_IsActiveWindow(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) Layout() *QGraphicsLayout {
	defer qt.Recovering("QGraphicsWidget::layout")

	if ptr.Pointer() != nil {
		return NewQGraphicsLayoutFromPointer(C.QGraphicsWidget_Layout(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWidget) ConnectMoveEvent(f func(event *QGraphicsSceneMoveEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQGraphicsWidgetMoveEvent
func callbackQGraphicsWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMoveEvent))(NewQGraphicsSceneMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsWidget::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsWidget::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
}

//export callbackQGraphicsWidgetPaint
func callbackQGraphicsWidgetPaint(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectPaintWindowFrame(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsWidget::paintWindowFrame")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintWindowFrame", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectPaintWindowFrame() {
	defer qt.Recovering("disconnect QGraphicsWidget::paintWindowFrame")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintWindowFrame")
	}
}

//export callbackQGraphicsWidgetPaintWindowFrame
func callbackQGraphicsWidgetPaintWindowFrame(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::paintWindowFrame")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintWindowFrame"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectPolishEvent(f func()) {
	defer qt.Recovering("connect QGraphicsWidget::polishEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "polishEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectPolishEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::polishEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "polishEvent")
	}
}

//export callbackQGraphicsWidgetPolishEvent
func callbackQGraphicsWidgetPolishEvent(ptrName *C.char) bool {
	defer qt.Recovering("callback QGraphicsWidget::polishEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "polishEvent"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ReleaseShortcut(id int) {
	defer qt.Recovering("QGraphicsWidget::releaseShortcut")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ReleaseShortcut(ptr.Pointer(), C.int(id))
	}
}

func (ptr *QGraphicsWidget) RemoveAction(action QAction_ITF) {
	defer qt.Recovering("QGraphicsWidget::removeAction")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_RemoveAction(ptr.Pointer(), PointerFromQAction(action))
	}
}

func (ptr *QGraphicsWidget) Resize2(w float64, h float64) {
	defer qt.Recovering("QGraphicsWidget::resize")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Resize2(ptr.Pointer(), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsWidget) ConnectResizeEvent(f func(event *QGraphicsSceneResizeEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQGraphicsWidgetResizeEvent
func callbackQGraphicsWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*QGraphicsSceneResizeEvent))(NewQGraphicsSceneResizeEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) SetAttribute(attribute core.Qt__WidgetAttribute, on bool) {
	defer qt.Recovering("QGraphicsWidget::setAttribute")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetAttribute(ptr.Pointer(), C.int(attribute), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QGraphicsWidget) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	defer qt.Recovering("QGraphicsWidget::setContentsMargins")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetContentsMargins(ptr.Pointer(), C.double(left), C.double(top), C.double(right), C.double(bottom))
	}
}

func (ptr *QGraphicsWidget) SetGeometry2(x float64, y float64, w float64, h float64) {
	defer qt.Recovering("QGraphicsWidget::setGeometry")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetGeometry2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsWidget) SetShortcutAutoRepeat(id int, enabled bool) {
	defer qt.Recovering("QGraphicsWidget::setShortcutAutoRepeat")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetShortcutAutoRepeat(ptr.Pointer(), C.int(id), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWidget) SetShortcutEnabled(id int, enabled bool) {
	defer qt.Recovering("QGraphicsWidget::setShortcutEnabled")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetShortcutEnabled(ptr.Pointer(), C.int(id), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsWidget) SetStyle(style QStyle_ITF) {
	defer qt.Recovering("QGraphicsWidget::setStyle")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

func QGraphicsWidget_SetTabOrder(first QGraphicsWidget_ITF, second QGraphicsWidget_ITF) {
	defer qt.Recovering("QGraphicsWidget::setTabOrder")

	C.QGraphicsWidget_QGraphicsWidget_SetTabOrder(PointerFromQGraphicsWidget(first), PointerFromQGraphicsWidget(second))
}

func (ptr *QGraphicsWidget) SetWindowFrameMargins(left float64, top float64, right float64, bottom float64) {
	defer qt.Recovering("QGraphicsWidget::setWindowFrameMargins")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_SetWindowFrameMargins(ptr.Pointer(), C.double(left), C.double(top), C.double(right), C.double(bottom))
	}
}

func (ptr *QGraphicsWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQGraphicsWidgetShowEvent
func callbackQGraphicsWidgetShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) Style() *QStyle {
	defer qt.Recovering("QGraphicsWidget::style")

	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QGraphicsWidget_Style(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsWidget) TestAttribute(attribute core.Qt__WidgetAttribute) bool {
	defer qt.Recovering("QGraphicsWidget::testAttribute")

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_TestAttribute(ptr.Pointer(), C.int(attribute)) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) Type() int {
	defer qt.Recovering("QGraphicsWidget::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsWidget_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) ConnectUngrabKeyboardEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "ungrabKeyboardEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectUngrabKeyboardEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "ungrabKeyboardEvent")
	}
}

//export callbackQGraphicsWidgetUngrabKeyboardEvent
func callbackQGraphicsWidgetUngrabKeyboardEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::ungrabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectUngrabMouseEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::ungrabMouseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "ungrabMouseEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectUngrabMouseEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::ungrabMouseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "ungrabMouseEvent")
	}
}

//export callbackQGraphicsWidgetUngrabMouseEvent
func callbackQGraphicsWidgetUngrabMouseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::ungrabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) UnsetWindowFrameMargins() {
	defer qt.Recovering("QGraphicsWidget::unsetWindowFrameMargins")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UnsetWindowFrameMargins(ptr.Pointer())
	}
}

func (ptr *QGraphicsWidget) ConnectUpdateGeometry(f func()) {
	defer qt.Recovering("connect QGraphicsWidget::updateGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometry", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectUpdateGeometry() {
	defer qt.Recovering("disconnect QGraphicsWidget::updateGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometry")
	}
}

//export callbackQGraphicsWidgetUpdateGeometry
func callbackQGraphicsWidgetUpdateGeometry(ptrName *C.char) bool {
	defer qt.Recovering("callback QGraphicsWidget::updateGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometry"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) WindowType() core.Qt__WindowType {
	defer qt.Recovering("QGraphicsWidget::windowType")

	if ptr.Pointer() != nil {
		return core.Qt__WindowType(C.QGraphicsWidget_WindowType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsWidget) DestroyQGraphicsWidget() {
	defer qt.Recovering("QGraphicsWidget::~QGraphicsWidget")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_DestroyQGraphicsWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsWidgetTimerEvent
func callbackQGraphicsWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsWidgetChildEvent
func callbackQGraphicsWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsWidgetCustomEvent
func callbackQGraphicsWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
