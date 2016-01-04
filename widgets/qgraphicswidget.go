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
func callbackQGraphicsWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQGraphicsWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGraphicsWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGraphicsWidget) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWidget::event")

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
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
func callbackQGraphicsWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QGraphicsWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
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
func callbackQGraphicsWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQGraphicsWidgetGeometryChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWidget::geometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsWidget) GeometryChanged() {
	defer qt.Recovering("QGraphicsWidget::geometryChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_GeometryChanged(ptr.Pointer())
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
func callbackQGraphicsWidgetGrabKeyboardEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::grabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).GrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) GrabKeyboardEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::grabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_GrabKeyboardEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWidget) GrabKeyboardEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::grabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_GrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQGraphicsWidgetGrabMouseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::grabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).GrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) GrabMouseEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::grabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_GrabMouseEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWidget) GrabMouseEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::grabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_GrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQGraphicsWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGraphicsWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
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
func callbackQGraphicsWidgetHoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).HoverLeaveEventDefault(NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) HoverLeaveEvent(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_HoverLeaveEvent(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsWidget) HoverLeaveEventDefault(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_HoverLeaveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
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
func callbackQGraphicsWidgetHoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).HoverMoveEventDefault(NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) HoverMoveEvent(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_HoverMoveEvent(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsWidget) HoverMoveEventDefault(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_HoverMoveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
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
func callbackQGraphicsWidgetInitStyleOption(ptr unsafe.Pointer, ptrName *C.char, option unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::initStyleOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "initStyleOption"); signal != nil {
		signal.(func(*QStyleOption))(NewQStyleOptionFromPointer(option))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).InitStyleOptionDefault(NewQStyleOptionFromPointer(option))
	}
}

func (ptr *QGraphicsWidget) InitStyleOption(option QStyleOption_ITF) {
	defer qt.Recovering("QGraphicsWidget::initStyleOption")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_InitStyleOption(ptr.Pointer(), PointerFromQStyleOption(option))
	}
}

func (ptr *QGraphicsWidget) InitStyleOptionDefault(option QStyleOption_ITF) {
	defer qt.Recovering("QGraphicsWidget::initStyleOption")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_InitStyleOptionDefault(ptr.Pointer(), PointerFromQStyleOption(option))
	}
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

func (ptr *QGraphicsWidget) ItemChange(change QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QGraphicsWidget::itemChange")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsWidget_ItemChange(ptr.Pointer(), C.int(change), core.PointerFromQVariant(value)))
	}
	return nil
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
func callbackQGraphicsWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMoveEvent))(NewQGraphicsSceneMoveEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).MoveEventDefault(NewQGraphicsSceneMoveEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) MoveEvent(event QGraphicsSceneMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_MoveEvent(ptr.Pointer(), PointerFromQGraphicsSceneMoveEvent(event))
	}
}

func (ptr *QGraphicsWidget) MoveEventDefault(event QGraphicsSceneMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_MoveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMoveEvent(event))
	}
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
func callbackQGraphicsWidgetPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QGraphicsWidget) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsWidget::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsWidget) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsWidget::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
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
func callbackQGraphicsWidgetPaintWindowFrame(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::paintWindowFrame")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintWindowFrame"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).PaintWindowFrameDefault(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsWidget) PaintWindowFrame(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsWidget::paintWindowFrame")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_PaintWindowFrame(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsWidget) PaintWindowFrameDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsWidget::paintWindowFrame")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_PaintWindowFrameDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
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
func callbackQGraphicsWidgetPolishEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWidget::polishEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "polishEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWidgetFromPointer(ptr).PolishEventDefault()
	}
}

func (ptr *QGraphicsWidget) PolishEvent() {
	defer qt.Recovering("QGraphicsWidget::polishEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_PolishEvent(ptr.Pointer())
	}
}

func (ptr *QGraphicsWidget) PolishEventDefault() {
	defer qt.Recovering("QGraphicsWidget::polishEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_PolishEventDefault(ptr.Pointer())
	}
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
func callbackQGraphicsWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*QGraphicsSceneResizeEvent))(NewQGraphicsSceneResizeEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).ResizeEventDefault(NewQGraphicsSceneResizeEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) ResizeEvent(event QGraphicsSceneResizeEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ResizeEvent(ptr.Pointer(), PointerFromQGraphicsSceneResizeEvent(event))
	}
}

func (ptr *QGraphicsWidget) ResizeEventDefault(event QGraphicsSceneResizeEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ResizeEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneResizeEvent(event))
	}
}

func (ptr *QGraphicsWidget) SceneEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWidget::sceneEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_SceneEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
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
func callbackQGraphicsWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGraphicsWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
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
func callbackQGraphicsWidgetUngrabKeyboardEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::ungrabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).UngrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) UngrabKeyboardEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UngrabKeyboardEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWidget) UngrabKeyboardEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UngrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQGraphicsWidgetUngrabMouseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::ungrabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).UngrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) UngrabMouseEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::ungrabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UngrabMouseEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWidget) UngrabMouseEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::ungrabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UngrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
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
func callbackQGraphicsWidgetUpdateGeometry(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsWidget::updateGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometry"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsWidgetFromPointer(ptr).UpdateGeometryDefault()
	}
}

func (ptr *QGraphicsWidget) UpdateGeometry() {
	defer qt.Recovering("QGraphicsWidget::updateGeometry")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UpdateGeometry(ptr.Pointer())
	}
}

func (ptr *QGraphicsWidget) UpdateGeometryDefault() {
	defer qt.Recovering("QGraphicsWidget::updateGeometry")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_UpdateGeometryDefault(ptr.Pointer())
	}
}

func (ptr *QGraphicsWidget) WindowFrameEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsWidget::windowFrameEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsWidget_WindowFrameEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsWidget) WindowFrameSectionAt(pos core.QPointF_ITF) core.Qt__WindowFrameSection {
	defer qt.Recovering("QGraphicsWidget::windowFrameSectionAt")

	if ptr.Pointer() != nil {
		return core.Qt__WindowFrameSection(C.QGraphicsWidget_WindowFrameSectionAt(ptr.Pointer(), core.PointerFromQPointF(pos)))
	}
	return 0
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
func callbackQGraphicsWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQGraphicsWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQGraphicsWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
