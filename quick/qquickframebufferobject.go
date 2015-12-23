package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QQuickFramebufferObject struct {
	QQuickItem
}

type QQuickFramebufferObject_ITF interface {
	QQuickItem_ITF
	QQuickFramebufferObject_PTR() *QQuickFramebufferObject
}

func PointerFromQQuickFramebufferObject(ptr QQuickFramebufferObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickFramebufferObject_PTR().Pointer()
	}
	return nil
}

func NewQQuickFramebufferObjectFromPointer(ptr unsafe.Pointer) *QQuickFramebufferObject {
	var n = new(QQuickFramebufferObject)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQuickFramebufferObject_") {
		n.SetObjectName("QQuickFramebufferObject_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickFramebufferObject) QQuickFramebufferObject_PTR() *QQuickFramebufferObject {
	return ptr
}

func (ptr *QQuickFramebufferObject) SetTextureFollowsItemSize(follows bool) {
	defer qt.Recovering("QQuickFramebufferObject::setTextureFollowsItemSize")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_SetTextureFollowsItemSize(ptr.Pointer(), C.int(qt.GoBoolToInt(follows)))
	}
}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSize() bool {
	defer qt.Recovering("QQuickFramebufferObject::textureFollowsItemSize")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_TextureFollowsItemSize(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) IsTextureProvider() bool {
	defer qt.Recovering("QQuickFramebufferObject::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) ConnectReleaseResources(f func()) {
	defer qt.Recovering("connect QQuickFramebufferObject::releaseResources")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "releaseResources", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectReleaseResources() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::releaseResources")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "releaseResources")
	}
}

//export callbackQQuickFramebufferObjectReleaseResources
func callbackQQuickFramebufferObjectReleaseResources(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectTextureFollowsItemSizeChanged(f func(v bool)) {
	defer qt.Recovering("connect QQuickFramebufferObject::textureFollowsItemSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectTextureFollowsItemSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textureFollowsItemSizeChanged", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTextureFollowsItemSizeChanged() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::textureFollowsItemSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectTextureFollowsItemSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textureFollowsItemSizeChanged")
	}
}

//export callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged
func callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged(ptrName *C.char, v C.int) {
	defer qt.Recovering("callback QQuickFramebufferObject::textureFollowsItemSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureFollowsItemSizeChanged"); signal != nil {
		signal.(func(bool))(int(v) != 0)
	}

}

func (ptr *QQuickFramebufferObject) TextureProvider() *QSGTextureProvider {
	defer qt.Recovering("QQuickFramebufferObject::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickFramebufferObject_TextureProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickFramebufferObject) ConnectClassBegin(f func()) {
	defer qt.Recovering("connect QQuickFramebufferObject::classBegin")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "classBegin", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectClassBegin() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::classBegin")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "classBegin")
	}
}

//export callbackQQuickFramebufferObjectClassBegin
func callbackQQuickFramebufferObjectClassBegin(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::classBegin")

	if signal := qt.GetSignal(C.GoString(ptrName), "classBegin"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectComponentComplete(f func()) {
	defer qt.Recovering("connect QQuickFramebufferObject::componentComplete")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "componentComplete", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectComponentComplete() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::componentComplete")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "componentComplete")
	}
}

//export callbackQQuickFramebufferObjectComponentComplete
func callbackQQuickFramebufferObjectComponentComplete(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::componentComplete")

	if signal := qt.GetSignal(C.GoString(ptrName), "componentComplete"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQQuickFramebufferObjectDragEnterEvent
func callbackQQuickFramebufferObjectDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQQuickFramebufferObjectDragLeaveEvent
func callbackQQuickFramebufferObjectDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQQuickFramebufferObjectDragMoveEvent
func callbackQQuickFramebufferObjectDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQQuickFramebufferObjectDropEvent
func callbackQQuickFramebufferObjectDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickFramebufferObjectFocusInEvent
func callbackQQuickFramebufferObjectFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickFramebufferObjectFocusOutEvent
func callbackQQuickFramebufferObjectFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectHoverEnterEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverEnterEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectHoverEnterEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverEnterEvent")
	}
}

//export callbackQQuickFramebufferObjectHoverEnterEvent
func callbackQQuickFramebufferObjectHoverEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectHoverLeaveEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverLeaveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverLeaveEvent")
	}
}

//export callbackQQuickFramebufferObjectHoverLeaveEvent
func callbackQQuickFramebufferObjectHoverLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectHoverMoveEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverMoveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverMoveEvent")
	}
}

//export callbackQQuickFramebufferObjectHoverMoveEvent
func callbackQQuickFramebufferObjectHoverMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQQuickFramebufferObjectInputMethodEvent
func callbackQQuickFramebufferObjectInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickFramebufferObjectKeyPressEvent
func callbackQQuickFramebufferObjectKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickFramebufferObjectKeyReleaseEvent
func callbackQQuickFramebufferObjectKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickFramebufferObjectMouseDoubleClickEvent
func callbackQQuickFramebufferObjectMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickFramebufferObjectMouseMoveEvent
func callbackQQuickFramebufferObjectMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickFramebufferObjectMousePressEvent
func callbackQQuickFramebufferObjectMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickFramebufferObjectMouseReleaseEvent
func callbackQQuickFramebufferObjectMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectMouseUngrabEvent(f func()) {
	defer qt.Recovering("connect QQuickFramebufferObject::mouseUngrabEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseUngrabEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMouseUngrabEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::mouseUngrabEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseUngrabEvent")
	}
}

//export callbackQQuickFramebufferObjectMouseUngrabEvent
func callbackQQuickFramebufferObjectMouseUngrabEvent(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseUngrabEvent"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQQuickFramebufferObjectTouchEvent
func callbackQQuickFramebufferObjectTouchEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectTouchUngrabEvent(f func()) {
	defer qt.Recovering("connect QQuickFramebufferObject::touchUngrabEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchUngrabEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTouchUngrabEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::touchUngrabEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchUngrabEvent")
	}
}

//export callbackQQuickFramebufferObjectTouchUngrabEvent
func callbackQQuickFramebufferObjectTouchUngrabEvent(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::touchUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchUngrabEvent"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectUpdatePolish(f func()) {
	defer qt.Recovering("connect QQuickFramebufferObject::updatePolish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updatePolish", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectUpdatePolish() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::updatePolish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updatePolish")
	}
}

//export callbackQQuickFramebufferObjectUpdatePolish
func callbackQQuickFramebufferObjectUpdatePolish(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::updatePolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePolish"); signal != nil {
		signal.(func())()
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickFramebufferObjectWheelEvent
func callbackQQuickFramebufferObjectWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickFramebufferObjectTimerEvent
func callbackQQuickFramebufferObjectTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickFramebufferObjectChildEvent
func callbackQQuickFramebufferObjectChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QQuickFramebufferObject) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickFramebufferObject::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickFramebufferObjectCustomEvent
func callbackQQuickFramebufferObjectCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickFramebufferObject::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
