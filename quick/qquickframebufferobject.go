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
func callbackQQuickFramebufferObjectReleaseResources(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickFramebufferObject) ReleaseResources() {
	defer qt.Recovering("QQuickFramebufferObject::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) ReleaseResourcesDefault() {
	defer qt.Recovering("QQuickFramebufferObject::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ReleaseResourcesDefault(ptr.Pointer())
	}
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
func callbackQQuickFramebufferObjectTextureFollowsItemSizeChanged(ptr unsafe.Pointer, ptrName *C.char, v C.int) {
	defer qt.Recovering("callback QQuickFramebufferObject::textureFollowsItemSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureFollowsItemSizeChanged"); signal != nil {
		signal.(func(bool))(int(v) != 0)
	}

}

func (ptr *QQuickFramebufferObject) TextureFollowsItemSizeChanged(v bool) {
	defer qt.Recovering("QQuickFramebufferObject::textureFollowsItemSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TextureFollowsItemSizeChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
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
func callbackQQuickFramebufferObjectClassBegin(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::classBegin")

	if signal := qt.GetSignal(C.GoString(ptrName), "classBegin"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickFramebufferObject) ClassBegin() {
	defer qt.Recovering("QQuickFramebufferObject::classBegin")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ClassBegin(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) ClassBeginDefault() {
	defer qt.Recovering("QQuickFramebufferObject::classBegin")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ClassBeginDefault(ptr.Pointer())
	}
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
func callbackQQuickFramebufferObjectComponentComplete(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::componentComplete")

	if signal := qt.GetSignal(C.GoString(ptrName), "componentComplete"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickFramebufferObject) ComponentComplete() {
	defer qt.Recovering("QQuickFramebufferObject::componentComplete")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ComponentComplete(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) ComponentCompleteDefault() {
	defer qt.Recovering("QQuickFramebufferObject::componentComplete")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ComponentCompleteDefault(ptr.Pointer())
	}
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
func callbackQQuickFramebufferObjectDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
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
func callbackQQuickFramebufferObjectDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
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
func callbackQQuickFramebufferObjectDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
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
func callbackQQuickFramebufferObjectDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
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
func callbackQQuickFramebufferObjectFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQQuickFramebufferObjectFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
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
func callbackQQuickFramebufferObjectHoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) HoverEnterEvent(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverEnterEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) HoverEnterEventDefault(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverEnterEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
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
func callbackQQuickFramebufferObjectHoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) HoverLeaveEvent(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverLeaveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) HoverLeaveEventDefault(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverLeaveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
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
func callbackQQuickFramebufferObjectHoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) HoverMoveEvent(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverMoveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) HoverMoveEventDefault(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_HoverMoveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
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
func callbackQQuickFramebufferObjectInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
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
func callbackQQuickFramebufferObjectKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQQuickFramebufferObjectKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
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
func callbackQQuickFramebufferObjectMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQQuickFramebufferObjectMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQQuickFramebufferObjectMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQQuickFramebufferObjectMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
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
func callbackQQuickFramebufferObjectMouseUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseUngrabEventDefault()
	}
}

func (ptr *QQuickFramebufferObject) MouseUngrabEvent() {
	defer qt.Recovering("QQuickFramebufferObject::mouseUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) MouseUngrabEventDefault() {
	defer qt.Recovering("QQuickFramebufferObject::mouseUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MouseUngrabEventDefault(ptr.Pointer())
	}
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
func callbackQQuickFramebufferObjectTouchEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) TouchEvent(event gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) TouchEventDefault(event gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
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
func callbackQQuickFramebufferObjectTouchUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::touchUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).TouchUngrabEventDefault()
	}
}

func (ptr *QQuickFramebufferObject) TouchUngrabEvent() {
	defer qt.Recovering("QQuickFramebufferObject::touchUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TouchUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) TouchUngrabEventDefault() {
	defer qt.Recovering("QQuickFramebufferObject::touchUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TouchUngrabEventDefault(ptr.Pointer())
	}
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
func callbackQQuickFramebufferObjectUpdatePolish(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::updatePolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).UpdatePolishDefault()
	}
}

func (ptr *QQuickFramebufferObject) UpdatePolish() {
	defer qt.Recovering("QQuickFramebufferObject::updatePolish")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_UpdatePolish(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) UpdatePolishDefault() {
	defer qt.Recovering("QQuickFramebufferObject::updatePolish")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_UpdatePolishDefault(ptr.Pointer())
	}
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
func callbackQQuickFramebufferObjectWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
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
func callbackQQuickFramebufferObjectTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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
func callbackQQuickFramebufferObjectChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
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
func callbackQQuickFramebufferObjectCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickFramebufferObject) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickFramebufferObject) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
