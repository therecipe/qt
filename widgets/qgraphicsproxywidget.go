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
func callbackQGraphicsProxyWidgetContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*QGraphicsSceneContextMenuEvent))(NewQGraphicsSceneContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
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
func callbackQGraphicsProxyWidgetFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
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
func callbackQGraphicsProxyWidgetFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetGrabMouseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::grabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetHideEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetHoverEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetHoverLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetHoverMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetPaint(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetResizeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*QGraphicsSceneResizeEvent))(NewQGraphicsSceneResizeEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetShowEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetUngrabMouseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::ungrabMouseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabMouseEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*QGraphicsSceneWheelEvent))(NewQGraphicsSceneWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetChangeEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetCloseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetGrabKeyboardEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::grabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "grabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetInitStyleOption(ptrName *C.char, option unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::initStyleOption")

	if signal := qt.GetSignal(C.GoString(ptrName), "initStyleOption"); signal != nil {
		signal.(func(*QStyleOption))(NewQStyleOptionFromPointer(option))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMoveEvent))(NewQGraphicsSceneMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetPaintWindowFrame(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::paintWindowFrame")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintWindowFrame"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetPolishEvent(ptrName *C.char) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::polishEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "polishEvent"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetUngrabKeyboardEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::ungrabKeyboardEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "ungrabKeyboardEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetUpdateGeometry(ptrName *C.char) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::updateGeometry")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateGeometry"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsProxyWidgetCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsProxyWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
