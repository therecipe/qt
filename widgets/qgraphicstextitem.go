package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsTextItem struct {
	QGraphicsObject
}

type QGraphicsTextItem_ITF interface {
	QGraphicsObject_ITF
	QGraphicsTextItem_PTR() *QGraphicsTextItem
}

func PointerFromQGraphicsTextItem(ptr QGraphicsTextItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsTextItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsTextItemFromPointer(ptr unsafe.Pointer) *QGraphicsTextItem {
	var n = new(QGraphicsTextItem)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsTextItem_") {
		n.SetObjectName("QGraphicsTextItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsTextItem) QGraphicsTextItem_PTR() *QGraphicsTextItem {
	return ptr
}

func (ptr *QGraphicsTextItem) OpenExternalLinks() bool {
	defer qt.Recovering("QGraphicsTextItem::openExternalLinks")

	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_OpenExternalLinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) SetOpenExternalLinks(open bool) {
	defer qt.Recovering("QGraphicsTextItem::setOpenExternalLinks")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetOpenExternalLinks(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QGraphicsTextItem) SetTextCursor(cursor gui.QTextCursor_ITF) {
	defer qt.Recovering("QGraphicsTextItem::setTextCursor")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTextCursor(ptr.Pointer(), gui.PointerFromQTextCursor(cursor))
	}
}

func (ptr *QGraphicsTextItem) AdjustSize() {
	defer qt.Recovering("QGraphicsTextItem::adjustSize")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_AdjustSize(ptr.Pointer())
	}
}

func (ptr *QGraphicsTextItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QGraphicsTextItem::contains")

	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) ConnectContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQGraphicsTextItemContextMenuEvent
func callbackQGraphicsTextItemContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*QGraphicsSceneContextMenuEvent))(NewQGraphicsSceneContextMenuEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).ContextMenuEventDefault(NewQGraphicsSceneContextMenuEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) ContextMenuEvent(event QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ContextMenuEvent(ptr.Pointer(), PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ContextMenuEventDefault(event QGraphicsSceneContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ContextMenuEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneContextMenuEvent(event))
	}
}

func (ptr *QGraphicsTextItem) DefaultTextColor() *gui.QColor {
	defer qt.Recovering("QGraphicsTextItem::defaultTextColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QGraphicsTextItem_DefaultTextColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsTextItem) Document() *gui.QTextDocument {
	defer qt.Recovering("QGraphicsTextItem::document")

	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QGraphicsTextItem_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsTextItem) ConnectDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQGraphicsTextItemDragEnterEvent
func callbackQGraphicsTextItemDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).DragEnterEventDefault(NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) DragEnterEvent(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DragEnterEvent(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsTextItem) DragEnterEventDefault(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DragEnterEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQGraphicsTextItemDragLeaveEvent
func callbackQGraphicsTextItemDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).DragLeaveEventDefault(NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) DragLeaveEvent(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DragLeaveEvent(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsTextItem) DragLeaveEventDefault(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DragLeaveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQGraphicsTextItemDragMoveEvent
func callbackQGraphicsTextItemDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).DragMoveEventDefault(NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) DragMoveEvent(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DragMoveEvent(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsTextItem) DragMoveEventDefault(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DragMoveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectDropEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQGraphicsTextItemDropEvent
func callbackQGraphicsTextItemDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).DropEventDefault(NewQGraphicsSceneDragDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) DropEvent(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DropEvent(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsTextItem) DropEventDefault(event QGraphicsSceneDragDropEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DropEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneDragDropEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQGraphicsTextItemFocusInEvent
func callbackQGraphicsTextItemFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsTextItem) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQGraphicsTextItemFocusOutEvent
func callbackQGraphicsTextItemFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsTextItem) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectHoverEnterEvent(f func(event *QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverEnterEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectHoverEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverEnterEvent")
	}
}

//export callbackQGraphicsTextItemHoverEnterEvent
func callbackQGraphicsTextItemHoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).HoverEnterEventDefault(NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) HoverEnterEvent(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_HoverEnterEvent(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsTextItem) HoverEnterEventDefault(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_HoverEnterEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverLeaveEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverLeaveEvent")
	}
}

//export callbackQGraphicsTextItemHoverLeaveEvent
func callbackQGraphicsTextItemHoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).HoverLeaveEventDefault(NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) HoverLeaveEvent(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_HoverLeaveEvent(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsTextItem) HoverLeaveEventDefault(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_HoverLeaveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverMoveEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverMoveEvent")
	}
}

//export callbackQGraphicsTextItemHoverMoveEvent
func callbackQGraphicsTextItemHoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).HoverMoveEventDefault(NewQGraphicsSceneHoverEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) HoverMoveEvent(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_HoverMoveEvent(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsTextItem) HoverMoveEventDefault(event QGraphicsSceneHoverEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_HoverMoveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneHoverEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQGraphicsTextItemInputMethodEvent
func callbackQGraphicsTextItemInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGraphicsTextItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGraphicsTextItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsTextItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsTextItem_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QGraphicsTextItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	defer qt.Recovering("QGraphicsTextItem::isObscuredBy")

	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQGraphicsTextItemKeyPressEvent
func callbackQGraphicsTextItemKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsTextItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQGraphicsTextItemKeyReleaseEvent
func callbackQGraphicsTextItemKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsTextItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectLinkActivated(f func(link string)) {
	defer qt.Recovering("connect QGraphicsTextItem::linkActivated")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectLinkActivated() {
	defer qt.Recovering("disconnect QGraphicsTextItem::linkActivated")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQGraphicsTextItemLinkActivated
func callbackQGraphicsTextItemLinkActivated(ptr unsafe.Pointer, ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QGraphicsTextItem::linkActivated")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkActivated"); signal != nil {
		signal.(func(string))(C.GoString(link))
	}

}

func (ptr *QGraphicsTextItem) LinkActivated(link string) {
	defer qt.Recovering("QGraphicsTextItem::linkActivated")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_LinkActivated(ptr.Pointer(), C.CString(link))
	}
}

func (ptr *QGraphicsTextItem) ConnectLinkHovered(f func(link string)) {
	defer qt.Recovering("connect QGraphicsTextItem::linkHovered")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ConnectLinkHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkHovered", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectLinkHovered() {
	defer qt.Recovering("disconnect QGraphicsTextItem::linkHovered")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkHovered")
	}
}

//export callbackQGraphicsTextItemLinkHovered
func callbackQGraphicsTextItemLinkHovered(ptr unsafe.Pointer, ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QGraphicsTextItem::linkHovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "linkHovered"); signal != nil {
		signal.(func(string))(C.GoString(link))
	}

}

func (ptr *QGraphicsTextItem) LinkHovered(link string) {
	defer qt.Recovering("QGraphicsTextItem::linkHovered")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_LinkHovered(ptr.Pointer(), C.CString(link))
	}
}

func (ptr *QGraphicsTextItem) ConnectMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQGraphicsTextItemMouseDoubleClickEvent
func callbackQGraphicsTextItemMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).MouseDoubleClickEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) MouseDoubleClickEvent(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_MouseDoubleClickEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsTextItem) MouseDoubleClickEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_MouseDoubleClickEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectMouseMoveEvent(f func(event *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQGraphicsTextItemMouseMoveEvent
func callbackQGraphicsTextItemMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).MouseMoveEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) MouseMoveEvent(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_MouseMoveEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsTextItem) MouseMoveEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_MouseMoveEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectMousePressEvent(f func(event *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQGraphicsTextItemMousePressEvent
func callbackQGraphicsTextItemMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).MousePressEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) MousePressEvent(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_MousePressEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsTextItem) MousePressEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_MousePressEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQGraphicsTextItemMouseReleaseEvent
func callbackQGraphicsTextItemMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).MouseReleaseEventDefault(NewQGraphicsSceneMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) MouseReleaseEvent(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_MouseReleaseEvent(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsTextItem) MouseReleaseEventDefault(event QGraphicsSceneMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_MouseReleaseEventDefault(ptr.Pointer(), PointerFromQGraphicsSceneMouseEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectPaint(f func(painter *gui.QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsTextItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectPaint() {
	defer qt.Recovering("disconnect QGraphicsTextItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
}

//export callbackQGraphicsTextItemPaint
func callbackQGraphicsTextItemPaint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
	}

}

func (ptr *QGraphicsTextItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsTextItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsTextItem) PaintDefault(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsTextItem::paint")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_PaintDefault(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsTextItem) SceneEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsTextItem::sceneEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_SceneEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) SetDefaultTextColor(col gui.QColor_ITF) {
	defer qt.Recovering("QGraphicsTextItem::setDefaultTextColor")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetDefaultTextColor(ptr.Pointer(), gui.PointerFromQColor(col))
	}
}

func (ptr *QGraphicsTextItem) SetDocument(document gui.QTextDocument_ITF) {
	defer qt.Recovering("QGraphicsTextItem::setDocument")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetDocument(ptr.Pointer(), gui.PointerFromQTextDocument(document))
	}
}

func (ptr *QGraphicsTextItem) SetFont(font gui.QFont_ITF) {
	defer qt.Recovering("QGraphicsTextItem::setFont")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QGraphicsTextItem) SetHtml(text string) {
	defer qt.Recovering("QGraphicsTextItem::setHtml")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetHtml(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QGraphicsTextItem) SetPlainText(text string) {
	defer qt.Recovering("QGraphicsTextItem::setPlainText")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QGraphicsTextItem) SetTabChangesFocus(b bool) {
	defer qt.Recovering("QGraphicsTextItem::setTabChangesFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTabChangesFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QGraphicsTextItem) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	defer qt.Recovering("QGraphicsTextItem::setTextInteractionFlags")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QGraphicsTextItem) SetTextWidth(width float64) {
	defer qt.Recovering("QGraphicsTextItem::setTextWidth")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTextWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QGraphicsTextItem) TabChangesFocus() bool {
	defer qt.Recovering("QGraphicsTextItem::tabChangesFocus")

	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_TabChangesFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) TextInteractionFlags() core.Qt__TextInteractionFlag {
	defer qt.Recovering("QGraphicsTextItem::textInteractionFlags")

	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QGraphicsTextItem_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsTextItem) TextWidth() float64 {
	defer qt.Recovering("QGraphicsTextItem::textWidth")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsTextItem_TextWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsTextItem) ToHtml() string {
	defer qt.Recovering("QGraphicsTextItem::toHtml")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsTextItem_ToHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsTextItem) ToPlainText() string {
	defer qt.Recovering("QGraphicsTextItem::toPlainText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsTextItem_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsTextItem) Type() int {
	defer qt.Recovering("QGraphicsTextItem::type")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsTextItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsTextItem) DestroyQGraphicsTextItem() {
	defer qt.Recovering("QGraphicsTextItem::~QGraphicsTextItem")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DestroyQGraphicsTextItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsTextItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsTextItemTimerEvent
func callbackQGraphicsTextItemTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsTextItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsTextItemChildEvent
func callbackQGraphicsTextItemChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsTextItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsTextItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsTextItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsTextItemCustomEvent
func callbackQGraphicsTextItemCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsTextItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsTextItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsTextItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsTextItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsTextItem::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
