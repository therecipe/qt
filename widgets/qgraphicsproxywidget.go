package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsProxyWidget struct {
	QGraphicsWidget
}

type QGraphicsProxyWidget_ITF interface {
	QGraphicsWidget_ITF
	QGraphicsProxyWidget_PTR() *QGraphicsProxyWidget
}

func PointerFromQGraphicsProxyWidget(ptr QGraphicsProxyWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsProxyWidget_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsProxyWidgetFromPointer(ptr unsafe.Pointer) *QGraphicsProxyWidget {
	var n = new(QGraphicsProxyWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsProxyWidget_") {
		n.SetObjectName("QGraphicsProxyWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsProxyWidget) QGraphicsProxyWidget_PTR() *QGraphicsProxyWidget {
	return ptr
}

func NewQGraphicsProxyWidget(parent QGraphicsItem_ITF, wFlags core.Qt__WindowType) *QGraphicsProxyWidget {
	defer qt.Recovering("QGraphicsProxyWidget::QGraphicsProxyWidget")

	return NewQGraphicsProxyWidgetFromPointer(C.QGraphicsProxyWidget_NewQGraphicsProxyWidget(PointerFromQGraphicsItem(parent), C.int(wFlags)))
}

func (ptr *QGraphicsProxyWidget) ConnectContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQGraphicsProxyWidgetContextMenuEvent
func callbackQGraphicsProxyWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*QGraphicsSceneContextMenuEvent))(NewQGraphicsSceneContextMenuEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).ContextMenuEventDefault(NewQGraphicsSceneContextMenuEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) ContextMenuEvent(event QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ContextMenuEvent(ptr.Pointer(), PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ContextMenuEventDefault(event QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ContextMenuEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) CreateProxyForChildWidget(child QWidget_ITF) *QGraphicsProxyWidget {
	defer qt.Recovering("QGraphicsProxyWidget::createProxyForChildWidget")

	if ptr.Pointer() != nil {
		return NewQGraphicsProxyWidgetFromPointer(C.QGraphicsProxyWidget_CreateProxyForChildWidget(ptr.Pointer(), PointerFromQWidget(child)))
	}
	return nil
}

func (ptr *QGraphicsProxyWidget) ConnectDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQGraphicsProxyWidgetDragEnterEvent
func callbackQGraphicsProxyWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).DragEnterEventDefault(NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) DragEnterEvent(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DragEnterEvent(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) DragEnterEventDefault(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DragEnterEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQGraphicsProxyWidgetDragLeaveEvent
func callbackQGraphicsProxyWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).DragLeaveEventDefault(NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) DragLeaveEvent(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DragLeaveEvent(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) DragLeaveEventDefault(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DragLeaveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQGraphicsProxyWidgetDragMoveEvent
func callbackQGraphicsProxyWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).DragMoveEventDefault(NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) DragMoveEvent(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DragMoveEvent(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) DragMoveEventDefault(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DragMoveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectDropEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQGraphicsProxyWidgetDropEvent
func callbackQGraphicsProxyWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).DropEventDefault(NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) DropEvent(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DropEvent(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) DropEventDefault(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DropEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsProxyWidget::event")

	if ptr.Pointer() != nil {
		return C.QGraphicsProxyWidget_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsProxyWidget) EventFilter(object core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsProxyWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QGraphicsProxyWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(object), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsProxyWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQGraphicsProxyWidgetFocusInEvent
func callbackQGraphicsProxyWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QGraphicsProxyWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QGraphicsProxyWidget_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QGraphicsProxyWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQGraphicsProxyWidgetFocusOutEvent
func callbackQGraphicsProxyWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectGrabMouseEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::grabMouseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "grabMouseEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectGrabMouseEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::grabMouseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "grabMouseEvent")
	}
}

//export callbackQGraphicsProxyWidgetGrabMouseEvent
func callbackQGraphicsProxyWidgetGrabMouseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::grabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).GrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) GrabMouseEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::grabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_GrabMouseEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) GrabMouseEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::grabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_GrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQGraphicsProxyWidgetHideEvent
func callbackQGraphicsProxyWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectHoverEnterEvent(f func(event *QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverEnterEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectHoverEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverEnterEvent")
	}
}

//export callbackQGraphicsProxyWidgetHoverEnterEvent
func callbackQGraphicsProxyWidgetHoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).HoverEnterEventDefault(NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) HoverEnterEvent(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_HoverEnterEvent(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) HoverEnterEventDefault(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_HoverEnterEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverLeaveEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverLeaveEvent")
	}
}

//export callbackQGraphicsProxyWidgetHoverLeaveEvent
func callbackQGraphicsProxyWidgetHoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).HoverLeaveEventDefault(NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) HoverLeaveEvent(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_HoverLeaveEvent(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) HoverLeaveEventDefault(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_HoverLeaveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverMoveEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverMoveEvent")
	}
}

//export callbackQGraphicsProxyWidgetHoverMoveEvent
func callbackQGraphicsProxyWidgetHoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).HoverMoveEventDefault(NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) HoverMoveEvent(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_HoverMoveEvent(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) HoverMoveEventDefault(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_HoverMoveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQGraphicsProxyWidgetInputMethodEvent
func callbackQGraphicsProxyWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsProxyWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsProxyWidget_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QGraphicsProxyWidget) ItemChange(change QGraphicsItem__GraphicsItemChange, value core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QGraphicsProxyWidget::itemChange")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsProxyWidget_ItemChange(ptr.Pointer(), C.int(change), core.PointerFromQVariant(value)))
	}
	return nil
}

func (ptr *QGraphicsProxyWidget) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQGraphicsProxyWidgetKeyPressEvent
func callbackQGraphicsProxyWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQGraphicsProxyWidgetKeyReleaseEvent
func callbackQGraphicsProxyWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQGraphicsProxyWidgetMouseDoubleClickEvent
func callbackQGraphicsProxyWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).MouseDoubleClickEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) MouseDoubleClickEvent(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MouseDoubleClickEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) MouseDoubleClickEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MouseDoubleClickEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectMouseMoveEvent(f func(event *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQGraphicsProxyWidgetMouseMoveEvent
func callbackQGraphicsProxyWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).MouseMoveEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) MouseMoveEvent(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MouseMoveEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) MouseMoveEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MouseMoveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectMousePressEvent(f func(event *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQGraphicsProxyWidgetMousePressEvent
func callbackQGraphicsProxyWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).MousePressEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) MousePressEvent(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MousePressEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) MousePressEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MousePressEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQGraphicsProxyWidgetMouseReleaseEvent
func callbackQGraphicsProxyWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).MouseReleaseEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) MouseReleaseEvent(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MouseReleaseEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) MouseReleaseEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MouseReleaseEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
}

//export callbackQGraphicsProxyWidgetPaint
func callbackQGraphicsProxyWidgetPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QGraphicsProxyWidget) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsProxyWidget) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectResizeEvent(f func(event *QGraphicsSceneResizeEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQGraphicsProxyWidgetResizeEvent
func callbackQGraphicsProxyWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*QGraphicsSceneResizeEvent))(NewQGraphicsSceneResizeEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).ResizeEventDefault(NewQGraphicsSceneResizeEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) ResizeEvent(event QGraphicsSceneResizeEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ResizeEvent(ptr.Pointer(), PointerFromQGraphicsSceneResizeEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ResizeEventDefault(event QGraphicsSceneResizeEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ResizeEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneResizeEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) SetWidget(widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::setWidget")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_SetWidget(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQGraphicsProxyWidgetShowEvent
func callbackQGraphicsProxyWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) Type() int {
	defer qt.Recovering("QGraphicsProxyWidget::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsProxyWidget_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsProxyWidget) ConnectUngrabMouseEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::ungrabMouseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "ungrabMouseEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectUngrabMouseEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::ungrabMouseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "ungrabMouseEvent")
	}
}

//export callbackQGraphicsProxyWidgetUngrabMouseEvent
func callbackQGraphicsProxyWidgetUngrabMouseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::ungrabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).UngrabMouseEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) UngrabMouseEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::ungrabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_UngrabMouseEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) UngrabMouseEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::ungrabMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_UngrabMouseEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectWheelEvent(f func(event *QGraphicsSceneWheelEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQGraphicsProxyWidgetWheelEvent
func callbackQGraphicsProxyWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*QGraphicsSceneWheelEvent))(NewQGraphicsSceneWheelEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).WheelEventDefault(NewQGraphicsSceneWheelEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) WheelEvent(event QGraphicsSceneWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_WheelEvent(ptr.Pointer(), PointerFromQGraphicsSceneWheelEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) WheelEventDefault(event QGraphicsSceneWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_WheelEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneWheelEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) Widget() *QWidget {
	defer qt.Recovering("QGraphicsProxyWidget::widget")

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QGraphicsProxyWidget_Widget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsProxyWidget) DestroyQGraphicsProxyWidget() {
	defer qt.Recovering("QGraphicsProxyWidget::~QGraphicsProxyWidget")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_DestroyQGraphicsProxyWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsProxyWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQGraphicsProxyWidgetChangeEvent
func callbackQGraphicsProxyWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQGraphicsProxyWidgetCloseEvent
func callbackQGraphicsProxyWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectGrabKeyboardEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::grabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "grabKeyboardEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectGrabKeyboardEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::grabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "grabKeyboardEvent")
	}
}

//export callbackQGraphicsProxyWidgetGrabKeyboardEvent
func callbackQGraphicsProxyWidgetGrabKeyboardEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::grabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).GrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) GrabKeyboardEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::grabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_GrabKeyboardEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) GrabKeyboardEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::grabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_GrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectInitStyleOption(f func(option *QStyleOption)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::initStyleOption")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initStyleOption", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectInitStyleOption() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::initStyleOption")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initStyleOption")
	}
}

//export callbackQGraphicsProxyWidgetInitStyleOption
func callbackQGraphicsProxyWidgetInitStyleOption(ptr unsafe.Pointer, ptrName *C.char, option unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::initStyleOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "initStyleOption"); signal != nil {
		signal.(func(*QStyleOption))(NewQStyleOptionFromPointer(option))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).InitStyleOptionDefault(NewQStyleOptionFromPointer(option))
	}
}

func (ptr *QGraphicsProxyWidget) InitStyleOption(option QStyleOption_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::initStyleOption")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_InitStyleOption(ptr.Pointer(), PointerFromQStyleOption(option))
	}
}

func (ptr *QGraphicsProxyWidget) InitStyleOptionDefault(option QStyleOption_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::initStyleOption")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_InitStyleOptionDefault(ptr.Pointer(), PointerFromQStyleOption(option))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectMoveEvent(f func(event *QGraphicsSceneMoveEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQGraphicsProxyWidgetMoveEvent
func callbackQGraphicsProxyWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMoveEvent))(NewQGraphicsSceneMoveEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).MoveEventDefault(NewQGraphicsSceneMoveEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) MoveEvent(event QGraphicsSceneMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MoveEvent(ptr.Pointer(), PointerFromQGraphicsSceneMoveEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) MoveEventDefault(event QGraphicsSceneMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_MoveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMoveEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectPaintWindowFrame(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::paintWindowFrame")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintWindowFrame", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectPaintWindowFrame() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::paintWindowFrame")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintWindowFrame")
	}
}

//export callbackQGraphicsProxyWidgetPaintWindowFrame
func callbackQGraphicsProxyWidgetPaintWindowFrame(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::paintWindowFrame")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintWindowFrame"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).PaintWindowFrameDefault(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}
}

func (ptr *QGraphicsProxyWidget) PaintWindowFrame(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::paintWindowFrame")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_PaintWindowFrame(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsProxyWidget) PaintWindowFrameDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::paintWindowFrame")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_PaintWindowFrameDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectPolishEvent(f func()) {
	defer qt.Recovering("connect QGraphicsProxyWidget::polishEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "polishEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectPolishEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::polishEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "polishEvent")
	}
}

//export callbackQGraphicsProxyWidgetPolishEvent
func callbackQGraphicsProxyWidgetPolishEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsProxyWidget::polishEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "polishEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).PolishEventDefault()
	}
}

func (ptr *QGraphicsProxyWidget) PolishEvent() {
	defer qt.Recovering("QGraphicsProxyWidget::polishEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_PolishEvent(ptr.Pointer())
	}
}

func (ptr *QGraphicsProxyWidget) PolishEventDefault() {
	defer qt.Recovering("QGraphicsProxyWidget::polishEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_PolishEventDefault(ptr.Pointer())
	}
}

func (ptr *QGraphicsProxyWidget) ConnectUngrabKeyboardEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "ungrabKeyboardEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectUngrabKeyboardEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "ungrabKeyboardEvent")
	}
}

//export callbackQGraphicsProxyWidgetUngrabKeyboardEvent
func callbackQGraphicsProxyWidgetUngrabKeyboardEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::ungrabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).UngrabKeyboardEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) UngrabKeyboardEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_UngrabKeyboardEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) UngrabKeyboardEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::ungrabKeyboardEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_UngrabKeyboardEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectUpdateGeometry(f func()) {
	defer qt.Recovering("connect QGraphicsProxyWidget::updateGeometry")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateGeometry", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectUpdateGeometry() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::updateGeometry")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateGeometry")
	}
}

//export callbackQGraphicsProxyWidgetUpdateGeometry
func callbackQGraphicsProxyWidgetUpdateGeometry(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsProxyWidget::updateGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometry"); signal != nil {
		signal.(func())()
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).UpdateGeometryDefault()
	}
}

func (ptr *QGraphicsProxyWidget) UpdateGeometry() {
	defer qt.Recovering("QGraphicsProxyWidget::updateGeometry")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_UpdateGeometry(ptr.Pointer())
	}
}

func (ptr *QGraphicsProxyWidget) UpdateGeometryDefault() {
	defer qt.Recovering("QGraphicsProxyWidget::updateGeometry")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_UpdateGeometryDefault(ptr.Pointer())
	}
}

func (ptr *QGraphicsProxyWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsProxyWidgetTimerEvent
func callbackQGraphicsProxyWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsProxyWidgetChildEvent
func callbackQGraphicsProxyWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsProxyWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsProxyWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsProxyWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsProxyWidgetCustomEvent
func callbackQGraphicsProxyWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsProxyWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsProxyWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsProxyWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsProxyWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsProxyWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsProxyWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
