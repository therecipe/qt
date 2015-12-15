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
func callbackQGraphicsTextItemContextMenuEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::contextMenuEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "contextMenuEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneContextMenuEvent))(NewQGraphicsSceneContextMenuEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::dragEnterEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "dragEnterEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::dragLeaveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::dragMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "dragMoveEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::dropEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "dropEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::focusInEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusInEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::focusOutEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "focusOutEvent")
	if signal != nil {
		defer signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemHoverEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::hoverEnterEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemHoverLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::hoverLeaveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemHoverMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::hoverMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneHoverEvent))(NewQGraphicsSceneHoverEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::inputMethodEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "inputMethodEvent")
	if signal != nil {
		defer signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::keyPressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyPressEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::keyReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent")
	if signal != nil {
		defer signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemLinkActivated(ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QGraphicsTextItem::linkActivated")

	var signal = qt.GetSignal(C.GoString(ptrName), "linkActivated")
	if signal != nil {
		signal.(func(string))(C.GoString(link))
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
func callbackQGraphicsTextItemLinkHovered(ptrName *C.char, link *C.char) {
	defer qt.Recovering("callback QGraphicsTextItem::linkHovered")

	var signal = qt.GetSignal(C.GoString(ptrName), "linkHovered")
	if signal != nil {
		signal.(func(string))(C.GoString(link))
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
func callbackQGraphicsTextItemMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::mouseDoubleClickEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::mouseMoveEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::mousePressEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mousePressEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::mouseReleaseEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent")
	if signal != nil {
		defer signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQGraphicsTextItemPaint(ptrName *C.char, painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsTextItem::paint")

	var signal = qt.GetSignal(C.GoString(ptrName), "paint")
	if signal != nil {
		defer signal.(func(*gui.QPainter, *QStyleOptionGraphicsItem, *QWidget))(gui.NewQPainterFromPointer(painter), NewQStyleOptionGraphicsItemFromPointer(option), NewQWidgetFromPointer(widget))
		return true
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
