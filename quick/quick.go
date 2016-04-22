package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
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
	return n
}

func newQQuickFramebufferObjectFromPointer(ptr unsafe.Pointer) *QQuickFramebufferObject {
	var n = NewQQuickFramebufferObjectFromPointer(ptr)
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
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ClassBeginDefault()
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
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ComponentCompleteDefault()
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

func (ptr *QQuickFramebufferObject) ConnectGeometryChanged(f func(newGeometry *core.QRectF, oldGeometry *core.QRectF)) {
	defer qt.Recovering("connect QQuickFramebufferObject::geometryChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "geometryChanged", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectGeometryChanged() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::geometryChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "geometryChanged")
	}
}

//export callbackQQuickFramebufferObjectGeometryChanged
func callbackQQuickFramebufferObjectGeometryChanged(ptr unsafe.Pointer, ptrName *C.char, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::geometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	}
}

func (ptr *QQuickFramebufferObject) GeometryChanged(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::geometryChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_GeometryChanged(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

func (ptr *QQuickFramebufferObject) GeometryChangedDefault(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::geometryChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_GeometryChangedDefault(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
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

type QQuickImageProvider struct {
	qml.QQmlImageProviderBase
}

type QQuickImageProvider_ITF interface {
	qml.QQmlImageProviderBase_ITF
	QQuickImageProvider_PTR() *QQuickImageProvider
}

func PointerFromQQuickImageProvider(ptr QQuickImageProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageProvider_PTR().Pointer()
	}
	return nil
}

func NewQQuickImageProviderFromPointer(ptr unsafe.Pointer) *QQuickImageProvider {
	var n = new(QQuickImageProvider)
	n.SetPointer(ptr)
	return n
}

func newQQuickImageProviderFromPointer(ptr unsafe.Pointer) *QQuickImageProvider {
	var n = NewQQuickImageProviderFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QQuickImageProvider_") {
		n.SetObjectNameAbs("QQuickImageProvider_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickImageProvider) QQuickImageProvider_PTR() *QQuickImageProvider {
	return ptr
}

func NewQQuickImageProvider(ty qml.QQmlImageProviderBase__ImageType, flags qml.QQmlImageProviderBase__Flag) *QQuickImageProvider {
	defer qt.Recovering("QQuickImageProvider::QQuickImageProvider")

	return newQQuickImageProviderFromPointer(C.QQuickImageProvider_NewQQuickImageProvider(C.int(ty), C.int(flags)))
}

func (ptr *QQuickImageProvider) Flags() qml.QQmlImageProviderBase__Flag {
	defer qt.Recovering("QQuickImageProvider::flags")

	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__Flag(C.QQuickImageProvider_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickImageProvider) ImageType() qml.QQmlImageProviderBase__ImageType {
	defer qt.Recovering("QQuickImageProvider::imageType")

	if ptr.Pointer() != nil {
		return qml.QQmlImageProviderBase__ImageType(C.QQuickImageProvider_ImageType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickImageProvider) RequestImage(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QImage {
	defer qt.Recovering("QQuickImageProvider::requestImage")

	if ptr.Pointer() != nil {
		return gui.NewQImageFromPointer(C.QQuickImageProvider_RequestImage(ptr.Pointer(), C.CString(id), core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
	}
	return nil
}

func (ptr *QQuickImageProvider) RequestPixmap(id string, size core.QSize_ITF, requestedSize core.QSize_ITF) *gui.QPixmap {
	defer qt.Recovering("QQuickImageProvider::requestPixmap")

	if ptr.Pointer() != nil {
		return gui.NewQPixmapFromPointer(C.QQuickImageProvider_RequestPixmap(ptr.Pointer(), C.CString(id), core.PointerFromQSize(size), core.PointerFromQSize(requestedSize)))
	}
	return nil
}

func (ptr *QQuickImageProvider) DestroyQQuickImageProvider() {
	defer qt.Recovering("QQuickImageProvider::~QQuickImageProvider")

	if ptr.Pointer() != nil {
		C.QQuickImageProvider_DestroyQQuickImageProvider(ptr.Pointer())
	}
}

func (ptr *QQuickImageProvider) ObjectNameAbs() string {
	defer qt.Recovering("QQuickImageProvider::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickImageProvider_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickImageProvider) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQuickImageProvider::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQuickImageProvider_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QQuickItem struct {
	core.QObject
	qml.QQmlParserStatus
}

type QQuickItem_ITF interface {
	core.QObject_ITF
	qml.QQmlParserStatus_ITF
	QQuickItem_PTR() *QQuickItem
}

func (p *QQuickItem) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QQuickItem) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QQmlParserStatus_PTR().SetPointer(ptr)
}

func PointerFromQQuickItem(ptr QQuickItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItem_PTR().Pointer()
	}
	return nil
}

func NewQQuickItemFromPointer(ptr unsafe.Pointer) *QQuickItem {
	var n = new(QQuickItem)
	n.SetPointer(ptr)
	return n
}

func newQQuickItemFromPointer(ptr unsafe.Pointer) *QQuickItem {
	var n = NewQQuickItemFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickItem_") {
		n.SetObjectName("QQuickItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickItem) QQuickItem_PTR() *QQuickItem {
	return ptr
}

//QQuickItem::Flag
type QQuickItem__Flag int64

const (
	QQuickItem__ItemClipsChildrenToShape = QQuickItem__Flag(0x01)
	QQuickItem__ItemAcceptsInputMethod   = QQuickItem__Flag(0x02)
	QQuickItem__ItemIsFocusScope         = QQuickItem__Flag(0x04)
	QQuickItem__ItemHasContents          = QQuickItem__Flag(0x08)
	QQuickItem__ItemAcceptsDrops         = QQuickItem__Flag(0x10)
)

//QQuickItem::ItemChange
type QQuickItem__ItemChange int64

const (
	QQuickItem__ItemChildAddedChange       = QQuickItem__ItemChange(0)
	QQuickItem__ItemChildRemovedChange     = QQuickItem__ItemChange(1)
	QQuickItem__ItemSceneChange            = QQuickItem__ItemChange(2)
	QQuickItem__ItemVisibleHasChanged      = QQuickItem__ItemChange(3)
	QQuickItem__ItemParentHasChanged       = QQuickItem__ItemChange(4)
	QQuickItem__ItemOpacityHasChanged      = QQuickItem__ItemChange(5)
	QQuickItem__ItemActiveFocusHasChanged  = QQuickItem__ItemChange(6)
	QQuickItem__ItemRotationHasChanged     = QQuickItem__ItemChange(7)
	QQuickItem__ItemAntialiasingHasChanged = QQuickItem__ItemChange(8)
)

//QQuickItem::TransformOrigin
type QQuickItem__TransformOrigin int64

const (
	QQuickItem__TopLeft     = QQuickItem__TransformOrigin(0)
	QQuickItem__Top         = QQuickItem__TransformOrigin(1)
	QQuickItem__TopRight    = QQuickItem__TransformOrigin(2)
	QQuickItem__Left        = QQuickItem__TransformOrigin(3)
	QQuickItem__Center      = QQuickItem__TransformOrigin(4)
	QQuickItem__Right       = QQuickItem__TransformOrigin(5)
	QQuickItem__BottomLeft  = QQuickItem__TransformOrigin(6)
	QQuickItem__Bottom      = QQuickItem__TransformOrigin(7)
	QQuickItem__BottomRight = QQuickItem__TransformOrigin(8)
)

func NewQQuickItem(parent QQuickItem_ITF) *QQuickItem {
	defer qt.Recovering("QQuickItem::QQuickItem")

	return newQQuickItemFromPointer(C.QQuickItem_NewQQuickItem(PointerFromQQuickItem(parent)))
}

func (ptr *QQuickItem) ActiveFocusOnTab() bool {
	defer qt.Recovering("QQuickItem::activeFocusOnTab")

	if ptr.Pointer() != nil {
		return C.QQuickItem_ActiveFocusOnTab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Antialiasing() bool {
	defer qt.Recovering("QQuickItem::antialiasing")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Antialiasing(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) BaselineOffset() float64 {
	defer qt.Recovering("QQuickItem::baselineOffset")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_BaselineOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ChildrenRect() *core.QRectF {
	defer qt.Recovering("QQuickItem::childrenRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QQuickItem_ChildrenRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) Clip() bool {
	defer qt.Recovering("QQuickItem::clip")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Clip(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) HasActiveFocus() bool {
	defer qt.Recovering("QQuickItem::hasActiveFocus")

	if ptr.Pointer() != nil {
		return C.QQuickItem_HasActiveFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) HasFocus() bool {
	defer qt.Recovering("QQuickItem::hasFocus")

	if ptr.Pointer() != nil {
		return C.QQuickItem_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Height() float64 {
	defer qt.Recovering("QQuickItem::height")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ImplicitHeight() float64 {
	defer qt.Recovering("QQuickItem::implicitHeight")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_ImplicitHeight(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) IsEnabled() bool {
	defer qt.Recovering("QQuickItem::isEnabled")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsTextureProvider() bool {
	defer qt.Recovering("QQuickItem::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsVisible() bool {
	defer qt.Recovering("QQuickItem::isVisible")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Opacity() float64 {
	defer qt.Recovering("QQuickItem::opacity")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ParentItem() *QQuickItem {
	defer qt.Recovering("QQuickItem::parentItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_ParentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) ResetAntialiasing() {
	defer qt.Recovering("QQuickItem::resetAntialiasing")

	if ptr.Pointer() != nil {
		C.QQuickItem_ResetAntialiasing(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ResetHeight() {
	defer qt.Recovering("QQuickItem::resetHeight")

	if ptr.Pointer() != nil {
		C.QQuickItem_ResetHeight(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ResetWidth() {
	defer qt.Recovering("QQuickItem::resetWidth")

	if ptr.Pointer() != nil {
		C.QQuickItem_ResetWidth(ptr.Pointer())
	}
}

func (ptr *QQuickItem) Rotation() float64 {
	defer qt.Recovering("QQuickItem::rotation")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Rotation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Scale() float64 {
	defer qt.Recovering("QQuickItem::scale")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Scale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) SetActiveFocusOnTab(v bool) {
	defer qt.Recovering("QQuickItem::setActiveFocusOnTab")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetActiveFocusOnTab(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetAntialiasing(v bool) {
	defer qt.Recovering("QQuickItem::setAntialiasing")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetAntialiasing(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetBaselineOffset(v float64) {
	defer qt.Recovering("QQuickItem::setBaselineOffset")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetBaselineOffset(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetClip(v bool) {
	defer qt.Recovering("QQuickItem::setClip")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetClip(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetEnabled(v bool) {
	defer qt.Recovering("QQuickItem::setEnabled")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetFocus(v bool) {
	defer qt.Recovering("QQuickItem::setFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetFocus2(focus bool, reason core.Qt__FocusReason) {
	defer qt.Recovering("QQuickItem::setFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus2(ptr.Pointer(), C.int(qt.GoBoolToInt(focus)), C.int(reason))
	}
}

func (ptr *QQuickItem) SetHeight(v float64) {
	defer qt.Recovering("QQuickItem::setHeight")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetHeight(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetImplicitHeight(v float64) {
	defer qt.Recovering("QQuickItem::setImplicitHeight")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitHeight(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetImplicitWidth(v float64) {
	defer qt.Recovering("QQuickItem::setImplicitWidth")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitWidth(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetOpacity(v float64) {
	defer qt.Recovering("QQuickItem::setOpacity")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetOpacity(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetParentItem(parent QQuickItem_ITF) {
	defer qt.Recovering("QQuickItem::setParentItem")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetParentItem(ptr.Pointer(), PointerFromQQuickItem(parent))
	}
}

func (ptr *QQuickItem) SetRotation(v float64) {
	defer qt.Recovering("QQuickItem::setRotation")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetRotation(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetScale(v float64) {
	defer qt.Recovering("QQuickItem::setScale")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetSmooth(v bool) {
	defer qt.Recovering("QQuickItem::setSmooth")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetSmooth(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetState(v string) {
	defer qt.Recovering("QQuickItem::setState")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetState(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QQuickItem) SetTransformOrigin(v QQuickItem__TransformOrigin) {
	defer qt.Recovering("QQuickItem::setTransformOrigin")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetTransformOrigin(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QQuickItem) SetVisible(v bool) {
	defer qt.Recovering("QQuickItem::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QQuickItem) SetWidth(v float64) {
	defer qt.Recovering("QQuickItem::setWidth")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetWidth(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetX(v float64) {
	defer qt.Recovering("QQuickItem::setX")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetX(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetY(v float64) {
	defer qt.Recovering("QQuickItem::setY")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetY(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) SetZ(v float64) {
	defer qt.Recovering("QQuickItem::setZ")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetZ(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickItem) Smooth() bool {
	defer qt.Recovering("QQuickItem::smooth")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Smooth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) State() string {
	defer qt.Recovering("QQuickItem::state")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickItem_State(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickItem) TextureProvider() *QSGTextureProvider {
	defer qt.Recovering("QQuickItem::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickItem_TextureProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) TransformOrigin() QQuickItem__TransformOrigin {
	defer qt.Recovering("QQuickItem::transformOrigin")

	if ptr.Pointer() != nil {
		return QQuickItem__TransformOrigin(C.QQuickItem_TransformOrigin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Width() float64 {
	defer qt.Recovering("QQuickItem::width")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) X() float64 {
	defer qt.Recovering("QQuickItem::x")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Y() float64 {
	defer qt.Recovering("QQuickItem::y")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Y(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) Z() float64 {
	defer qt.Recovering("QQuickItem::z")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_Z(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) AcceptHoverEvents() bool {
	defer qt.Recovering("QQuickItem::acceptHoverEvents")

	if ptr.Pointer() != nil {
		return C.QQuickItem_AcceptHoverEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) AcceptedMouseButtons() core.Qt__MouseButton {
	defer qt.Recovering("QQuickItem::acceptedMouseButtons")

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QQuickItem_AcceptedMouseButtons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ChildAt(x float64, y float64) *QQuickItem {
	defer qt.Recovering("QQuickItem::childAt")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_ChildAt(ptr.Pointer(), C.double(x), C.double(y)))
	}
	return nil
}

func (ptr *QQuickItem) ChildMouseEventFilter(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItem::childMouseEventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickItem_ChildMouseEventFilter(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickItem) ConnectClassBegin(f func()) {
	defer qt.Recovering("connect QQuickItem::classBegin")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "classBegin", f)
	}
}

func (ptr *QQuickItem) DisconnectClassBegin() {
	defer qt.Recovering("disconnect QQuickItem::classBegin")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "classBegin")
	}
}

//export callbackQQuickItemClassBegin
func callbackQQuickItemClassBegin(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::classBegin")

	if signal := qt.GetSignal(C.GoString(ptrName), "classBegin"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ClassBeginDefault()
	}
}

func (ptr *QQuickItem) ClassBegin() {
	defer qt.Recovering("QQuickItem::classBegin")

	if ptr.Pointer() != nil {
		C.QQuickItem_ClassBegin(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ClassBeginDefault() {
	defer qt.Recovering("QQuickItem::classBegin")

	if ptr.Pointer() != nil {
		C.QQuickItem_ClassBeginDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ConnectComponentComplete(f func()) {
	defer qt.Recovering("connect QQuickItem::componentComplete")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "componentComplete", f)
	}
}

func (ptr *QQuickItem) DisconnectComponentComplete() {
	defer qt.Recovering("disconnect QQuickItem::componentComplete")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "componentComplete")
	}
}

//export callbackQQuickItemComponentComplete
func callbackQQuickItemComponentComplete(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::componentComplete")

	if signal := qt.GetSignal(C.GoString(ptrName), "componentComplete"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ComponentCompleteDefault()
	}
}

func (ptr *QQuickItem) ComponentComplete() {
	defer qt.Recovering("QQuickItem::componentComplete")

	if ptr.Pointer() != nil {
		C.QQuickItem_ComponentComplete(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ComponentCompleteDefault() {
	defer qt.Recovering("QQuickItem::componentComplete")

	if ptr.Pointer() != nil {
		C.QQuickItem_ComponentCompleteDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QQuickItem::contains")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickItem) Cursor() *gui.QCursor {
	defer qt.Recovering("QQuickItem::cursor")

	if ptr.Pointer() != nil {
		return gui.NewQCursorFromPointer(C.QQuickItem_Cursor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QQuickItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QQuickItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQQuickItemDragEnterEvent
func callbackQQuickItemDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QQuickItem) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QQuickItem) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QQuickItem) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QQuickItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QQuickItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQQuickItemDragLeaveEvent
func callbackQQuickItemDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QQuickItem) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QQuickItem) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QQuickItem) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QQuickItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QQuickItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQQuickItemDragMoveEvent
func callbackQQuickItemDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QQuickItem) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QQuickItem) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QQuickItem) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QQuickItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QQuickItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQQuickItemDropEvent
func callbackQQuickItemDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QQuickItem) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QQuickItem) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QQuickItem) Event(ev core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItem::event")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Event(ptr.Pointer(), core.PointerFromQEvent(ev)) != 0
	}
	return false
}

func (ptr *QQuickItem) FiltersChildMouseEvents() bool {
	defer qt.Recovering("QQuickItem::filtersChildMouseEvents")

	if ptr.Pointer() != nil {
		return C.QQuickItem_FiltersChildMouseEvents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Flags() QQuickItem__Flag {
	defer qt.Recovering("QQuickItem::flags")

	if ptr.Pointer() != nil {
		return QQuickItem__Flag(C.QQuickItem_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickItemFocusInEvent
func callbackQQuickItemFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickItem) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickItem) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickItem) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickItemFocusOutEvent
func callbackQQuickItemFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickItem) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickItem) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickItem) ForceActiveFocus() {
	defer qt.Recovering("QQuickItem::forceActiveFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ForceActiveFocus2(reason core.Qt__FocusReason) {
	defer qt.Recovering("QQuickItem::forceActiveFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_ForceActiveFocus2(ptr.Pointer(), C.int(reason))
	}
}

func (ptr *QQuickItem) ConnectGeometryChanged(f func(newGeometry *core.QRectF, oldGeometry *core.QRectF)) {
	defer qt.Recovering("connect QQuickItem::geometryChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "geometryChanged", f)
	}
}

func (ptr *QQuickItem) DisconnectGeometryChanged() {
	defer qt.Recovering("disconnect QQuickItem::geometryChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "geometryChanged")
	}
}

//export callbackQQuickItemGeometryChanged
func callbackQQuickItemGeometryChanged(ptr unsafe.Pointer, ptrName *C.char, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::geometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickItemFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	}
}

func (ptr *QQuickItem) GeometryChanged(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	defer qt.Recovering("QQuickItem::geometryChanged")

	if ptr.Pointer() != nil {
		C.QQuickItem_GeometryChanged(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

func (ptr *QQuickItem) GeometryChangedDefault(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	defer qt.Recovering("QQuickItem::geometryChanged")

	if ptr.Pointer() != nil {
		C.QQuickItem_GeometryChangedDefault(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

func (ptr *QQuickItem) GrabMouse() {
	defer qt.Recovering("QQuickItem::grabMouse")

	if ptr.Pointer() != nil {
		C.QQuickItem_GrabMouse(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ConnectHoverEnterEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverEnterEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectHoverEnterEvent() {
	defer qt.Recovering("disconnect QQuickItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverEnterEvent")
	}
}

//export callbackQQuickItemHoverEnterEvent
func callbackQQuickItemHoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickItem) HoverEnterEvent(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_HoverEnterEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickItem) HoverEnterEventDefault(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_HoverEnterEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickItem) ConnectHoverLeaveEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverLeaveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QQuickItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverLeaveEvent")
	}
}

//export callbackQQuickItemHoverLeaveEvent
func callbackQQuickItemHoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickItem) HoverLeaveEvent(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_HoverLeaveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickItem) HoverLeaveEventDefault(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_HoverLeaveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickItem) ConnectHoverMoveEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverMoveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QQuickItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverMoveEvent")
	}
}

//export callbackQQuickItemHoverMoveEvent
func callbackQQuickItemHoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickItem) HoverMoveEvent(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_HoverMoveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickItem) HoverMoveEventDefault(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_HoverMoveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickItem) ImplicitWidth() float64 {
	defer qt.Recovering("QQuickItem::implicitWidth")

	if ptr.Pointer() != nil {
		return float64(C.QQuickItem_ImplicitWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QQuickItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QQuickItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQQuickItemInputMethodEvent
func callbackQQuickItemInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickItem) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickItem_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QQuickItem) IsFocusScope() bool {
	defer qt.Recovering("QQuickItem::isFocusScope")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsFocusScope(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepMouseGrab() bool {
	defer qt.Recovering("QQuickItem::keepMouseGrab")

	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepMouseGrab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) KeepTouchGrab() bool {
	defer qt.Recovering("QQuickItem::keepTouchGrab")

	if ptr.Pointer() != nil {
		return C.QQuickItem_KeepTouchGrab(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickItemKeyPressEvent
func callbackQQuickItemKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickItem) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickItemKeyReleaseEvent
func callbackQQuickItemKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickItem) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickItem) MapFromItem(item QQuickItem_ITF, point core.QPointF_ITF) *core.QPointF {
	defer qt.Recovering("QQuickItem::mapFromItem")

	if ptr.Pointer() != nil {
		return core.NewQPointFFromPointer(C.QQuickItem_MapFromItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQPointF(point)))
	}
	return nil
}

func (ptr *QQuickItem) MapFromScene(point core.QPointF_ITF) *core.QPointF {
	defer qt.Recovering("QQuickItem::mapFromScene")

	if ptr.Pointer() != nil {
		return core.NewQPointFFromPointer(C.QQuickItem_MapFromScene(ptr.Pointer(), core.PointerFromQPointF(point)))
	}
	return nil
}

func (ptr *QQuickItem) MapRectFromItem(item QQuickItem_ITF, rect core.QRectF_ITF) *core.QRectF {
	defer qt.Recovering("QQuickItem::mapRectFromItem")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QQuickItem_MapRectFromItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQRectF(rect)))
	}
	return nil
}

func (ptr *QQuickItem) MapRectFromScene(rect core.QRectF_ITF) *core.QRectF {
	defer qt.Recovering("QQuickItem::mapRectFromScene")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QQuickItem_MapRectFromScene(ptr.Pointer(), core.PointerFromQRectF(rect)))
	}
	return nil
}

func (ptr *QQuickItem) MapRectToItem(item QQuickItem_ITF, rect core.QRectF_ITF) *core.QRectF {
	defer qt.Recovering("QQuickItem::mapRectToItem")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QQuickItem_MapRectToItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQRectF(rect)))
	}
	return nil
}

func (ptr *QQuickItem) MapRectToScene(rect core.QRectF_ITF) *core.QRectF {
	defer qt.Recovering("QQuickItem::mapRectToScene")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QQuickItem_MapRectToScene(ptr.Pointer(), core.PointerFromQRectF(rect)))
	}
	return nil
}

func (ptr *QQuickItem) MapToItem(item QQuickItem_ITF, point core.QPointF_ITF) *core.QPointF {
	defer qt.Recovering("QQuickItem::mapToItem")

	if ptr.Pointer() != nil {
		return core.NewQPointFFromPointer(C.QQuickItem_MapToItem(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQPointF(point)))
	}
	return nil
}

func (ptr *QQuickItem) MapToScene(point core.QPointF_ITF) *core.QPointF {
	defer qt.Recovering("QQuickItem::mapToScene")

	if ptr.Pointer() != nil {
		return core.NewQPointFFromPointer(C.QQuickItem_MapToScene(ptr.Pointer(), core.PointerFromQPointF(point)))
	}
	return nil
}

func (ptr *QQuickItem) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickItemMouseDoubleClickEvent
func callbackQQuickItemMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickItemMouseMoveEvent
func callbackQQuickItemMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickItemMousePressEvent
func callbackQQuickItemMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickItemMouseReleaseEvent
func callbackQQuickItemMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickItem) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickItem) ConnectMouseUngrabEvent(f func()) {
	defer qt.Recovering("connect QQuickItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseUngrabEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectMouseUngrabEvent() {
	defer qt.Recovering("disconnect QQuickItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseUngrabEvent")
	}
}

//export callbackQQuickItemMouseUngrabEvent
func callbackQQuickItemMouseUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::mouseUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).MouseUngrabEventDefault()
	}
}

func (ptr *QQuickItem) MouseUngrabEvent() {
	defer qt.Recovering("QQuickItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MouseUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickItem) MouseUngrabEventDefault() {
	defer qt.Recovering("QQuickItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_MouseUngrabEventDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem {
	defer qt.Recovering("QQuickItem::nextItemInFocusChain")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_NextItemInFocusChain(ptr.Pointer(), C.int(qt.GoBoolToInt(forward))))
	}
	return nil
}

func (ptr *QQuickItem) Polish() {
	defer qt.Recovering("QQuickItem::polish")

	if ptr.Pointer() != nil {
		C.QQuickItem_Polish(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ConnectReleaseResources(f func()) {
	defer qt.Recovering("connect QQuickItem::releaseResources")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "releaseResources", f)
	}
}

func (ptr *QQuickItem) DisconnectReleaseResources() {
	defer qt.Recovering("disconnect QQuickItem::releaseResources")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "releaseResources")
	}
}

//export callbackQQuickItemReleaseResources
func callbackQQuickItemReleaseResources(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickItem) ReleaseResources() {
	defer qt.Recovering("QQuickItem::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickItem_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ReleaseResourcesDefault() {
	defer qt.Recovering("QQuickItem::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickItem_ReleaseResourcesDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ScopedFocusItem() *QQuickItem {
	defer qt.Recovering("QQuickItem::scopedFocusItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickItem_ScopedFocusItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) SetAcceptHoverEvents(enabled bool) {
	defer qt.Recovering("QQuickItem::setAcceptHoverEvents")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptHoverEvents(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickItem) SetAcceptedMouseButtons(buttons core.Qt__MouseButton) {
	defer qt.Recovering("QQuickItem::setAcceptedMouseButtons")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetAcceptedMouseButtons(ptr.Pointer(), C.int(buttons))
	}
}

func (ptr *QQuickItem) SetCursor(cursor gui.QCursor_ITF) {
	defer qt.Recovering("QQuickItem::setCursor")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetCursor(ptr.Pointer(), gui.PointerFromQCursor(cursor))
	}
}

func (ptr *QQuickItem) SetFiltersChildMouseEvents(filter bool) {
	defer qt.Recovering("QQuickItem::setFiltersChildMouseEvents")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFiltersChildMouseEvents(ptr.Pointer(), C.int(qt.GoBoolToInt(filter)))
	}
}

func (ptr *QQuickItem) SetFlag(flag QQuickItem__Flag, enabled bool) {
	defer qt.Recovering("QQuickItem::setFlag")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlag(ptr.Pointer(), C.int(flag), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickItem) SetFlags(flags QQuickItem__Flag) {
	defer qt.Recovering("QQuickItem::setFlags")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QQuickItem) SetKeepMouseGrab(keep bool) {
	defer qt.Recovering("QQuickItem::setKeepMouseGrab")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepMouseGrab(ptr.Pointer(), C.int(qt.GoBoolToInt(keep)))
	}
}

func (ptr *QQuickItem) SetKeepTouchGrab(keep bool) {
	defer qt.Recovering("QQuickItem::setKeepTouchGrab")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetKeepTouchGrab(ptr.Pointer(), C.int(qt.GoBoolToInt(keep)))
	}
}

func (ptr *QQuickItem) StackAfter(sibling QQuickItem_ITF) {
	defer qt.Recovering("QQuickItem::stackAfter")

	if ptr.Pointer() != nil {
		C.QQuickItem_StackAfter(ptr.Pointer(), PointerFromQQuickItem(sibling))
	}
}

func (ptr *QQuickItem) StackBefore(sibling QQuickItem_ITF) {
	defer qt.Recovering("QQuickItem::stackBefore")

	if ptr.Pointer() != nil {
		C.QQuickItem_StackBefore(ptr.Pointer(), PointerFromQQuickItem(sibling))
	}
}

func (ptr *QQuickItem) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	defer qt.Recovering("connect QQuickItem::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QQuickItem::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQQuickItemTouchEvent
func callbackQQuickItemTouchEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *QQuickItem) TouchEvent(event gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickItem::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QQuickItem) TouchEventDefault(event gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickItem::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QQuickItem) ConnectTouchUngrabEvent(f func()) {
	defer qt.Recovering("connect QQuickItem::touchUngrabEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchUngrabEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectTouchUngrabEvent() {
	defer qt.Recovering("disconnect QQuickItem::touchUngrabEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchUngrabEvent")
	}
}

//export callbackQQuickItemTouchUngrabEvent
func callbackQQuickItemTouchUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::touchUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).TouchUngrabEventDefault()
	}
}

func (ptr *QQuickItem) TouchUngrabEvent() {
	defer qt.Recovering("QQuickItem::touchUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_TouchUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickItem) TouchUngrabEventDefault() {
	defer qt.Recovering("QQuickItem::touchUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_TouchUngrabEventDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UngrabMouse() {
	defer qt.Recovering("QQuickItem::ungrabMouse")

	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabMouse(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UngrabTouchPoints() {
	defer qt.Recovering("QQuickItem::ungrabTouchPoints")

	if ptr.Pointer() != nil {
		C.QQuickItem_UngrabTouchPoints(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UnsetCursor() {
	defer qt.Recovering("QQuickItem::unsetCursor")

	if ptr.Pointer() != nil {
		C.QQuickItem_UnsetCursor(ptr.Pointer())
	}
}

func (ptr *QQuickItem) Update() {
	defer qt.Recovering("QQuickItem::update")

	if ptr.Pointer() != nil {
		C.QQuickItem_Update(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ConnectUpdatePolish(f func()) {
	defer qt.Recovering("connect QQuickItem::updatePolish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updatePolish", f)
	}
}

func (ptr *QQuickItem) DisconnectUpdatePolish() {
	defer qt.Recovering("disconnect QQuickItem::updatePolish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updatePolish")
	}
}

//export callbackQQuickItemUpdatePolish
func callbackQQuickItemUpdatePolish(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::updatePolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).UpdatePolishDefault()
	}
}

func (ptr *QQuickItem) UpdatePolish() {
	defer qt.Recovering("QQuickItem::updatePolish")

	if ptr.Pointer() != nil {
		C.QQuickItem_UpdatePolish(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UpdatePolishDefault() {
	defer qt.Recovering("QQuickItem::updatePolish")

	if ptr.Pointer() != nil {
		C.QQuickItem_UpdatePolishDefault(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickItemWheelEvent
func callbackQQuickItemWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickItem) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickItem::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickItem) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickItem::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickItem) Window() *QQuickWindow {
	defer qt.Recovering("QQuickItem::window")

	if ptr.Pointer() != nil {
		return NewQQuickWindowFromPointer(C.QQuickItem_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) ConnectWindowChanged(f func(window *QQuickWindow)) {
	defer qt.Recovering("connect QQuickItem::windowChanged")

	if ptr.Pointer() != nil {
		C.QQuickItem_ConnectWindowChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "windowChanged", f)
	}
}

func (ptr *QQuickItem) DisconnectWindowChanged() {
	defer qt.Recovering("disconnect QQuickItem::windowChanged")

	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectWindowChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "windowChanged")
	}
}

//export callbackQQuickItemWindowChanged
func callbackQQuickItemWindowChanged(ptr unsafe.Pointer, ptrName *C.char, window unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::windowChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowChanged"); signal != nil {
		signal.(func(*QQuickWindow))(NewQQuickWindowFromPointer(window))
	}

}

func (ptr *QQuickItem) WindowChanged(window QQuickWindow_ITF) {
	defer qt.Recovering("QQuickItem::windowChanged")

	if ptr.Pointer() != nil {
		C.QQuickItem_WindowChanged(ptr.Pointer(), PointerFromQQuickWindow(window))
	}
}

func (ptr *QQuickItem) DestroyQQuickItem() {
	defer qt.Recovering("QQuickItem::~QQuickItem")

	if ptr.Pointer() != nil {
		C.QQuickItem_DestroyQQuickItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickItemTimerEvent
func callbackQQuickItemTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickItemChildEvent
func callbackQQuickItemChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickItem::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickItem::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickItemCustomEvent
func callbackQQuickItemCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickItem::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickItem::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQuickItemGrabResult struct {
	core.QObject
}

type QQuickItemGrabResult_ITF interface {
	core.QObject_ITF
	QQuickItemGrabResult_PTR() *QQuickItemGrabResult
}

func PointerFromQQuickItemGrabResult(ptr QQuickItemGrabResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickItemGrabResult_PTR().Pointer()
	}
	return nil
}

func NewQQuickItemGrabResultFromPointer(ptr unsafe.Pointer) *QQuickItemGrabResult {
	var n = new(QQuickItemGrabResult)
	n.SetPointer(ptr)
	return n
}

func newQQuickItemGrabResultFromPointer(ptr unsafe.Pointer) *QQuickItemGrabResult {
	var n = NewQQuickItemGrabResultFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickItemGrabResult_") {
		n.SetObjectName("QQuickItemGrabResult_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickItemGrabResult) QQuickItemGrabResult_PTR() *QQuickItemGrabResult {
	return ptr
}

func (ptr *QQuickItemGrabResult) Image() *gui.QImage {
	defer qt.Recovering("QQuickItemGrabResult::image")

	if ptr.Pointer() != nil {
		return gui.NewQImageFromPointer(C.QQuickItemGrabResult_Image(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItemGrabResult) Url() *core.QUrl {
	defer qt.Recovering("QQuickItemGrabResult::url")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQuickItemGrabResult_Url(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItemGrabResult) ConnectReady(f func()) {
	defer qt.Recovering("connect QQuickItemGrabResult::ready")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectReady(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "ready", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectReady() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::ready")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectReady(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "ready")
	}
}

//export callbackQQuickItemGrabResultReady
func callbackQQuickItemGrabResultReady(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItemGrabResult::ready")

	if signal := qt.GetSignal(C.GoString(ptrName), "ready"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickItemGrabResult) Ready() {
	defer qt.Recovering("QQuickItemGrabResult::ready")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_Ready(ptr.Pointer())
	}
}

func (ptr *QQuickItemGrabResult) SaveToFile(fileName string) bool {
	defer qt.Recovering("QQuickItemGrabResult::saveToFile")

	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_SaveToFile(ptr.Pointer(), C.CString(fileName)) != 0
	}
	return false
}

func (ptr *QQuickItemGrabResult) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickItemGrabResult::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickItemGrabResultTimerEvent
func callbackQQuickItemGrabResultTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItemGrabResult::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickItemGrabResult) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickItemGrabResult) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickItemGrabResult) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickItemGrabResult::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickItemGrabResultChildEvent
func callbackQQuickItemGrabResultChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItemGrabResult::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickItemGrabResult) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickItemGrabResult) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickItemGrabResult) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickItemGrabResult::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickItemGrabResultCustomEvent
func callbackQQuickItemGrabResultCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItemGrabResult::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickItemGrabResult) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickItemGrabResult) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQuickPaintedItem struct {
	QQuickItem
}

type QQuickPaintedItem_ITF interface {
	QQuickItem_ITF
	QQuickPaintedItem_PTR() *QQuickPaintedItem
}

func PointerFromQQuickPaintedItem(ptr QQuickPaintedItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickPaintedItem_PTR().Pointer()
	}
	return nil
}

func NewQQuickPaintedItemFromPointer(ptr unsafe.Pointer) *QQuickPaintedItem {
	var n = new(QQuickPaintedItem)
	n.SetPointer(ptr)
	return n
}

func newQQuickPaintedItemFromPointer(ptr unsafe.Pointer) *QQuickPaintedItem {
	var n = NewQQuickPaintedItemFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickPaintedItem_") {
		n.SetObjectName("QQuickPaintedItem_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickPaintedItem) QQuickPaintedItem_PTR() *QQuickPaintedItem {
	return ptr
}

//QQuickPaintedItem::PerformanceHint
type QQuickPaintedItem__PerformanceHint int64

const (
	QQuickPaintedItem__FastFBOResizing = QQuickPaintedItem__PerformanceHint(0x1)
)

//QQuickPaintedItem::RenderTarget
type QQuickPaintedItem__RenderTarget int64

const (
	QQuickPaintedItem__Image                      = QQuickPaintedItem__RenderTarget(0)
	QQuickPaintedItem__FramebufferObject          = QQuickPaintedItem__RenderTarget(1)
	QQuickPaintedItem__InvertedYFramebufferObject = QQuickPaintedItem__RenderTarget(2)
)

func (ptr *QQuickPaintedItem) ContentsScale() float64 {
	defer qt.Recovering("QQuickPaintedItem::contentsScale")

	if ptr.Pointer() != nil {
		return float64(C.QQuickPaintedItem_ContentsScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickPaintedItem) ContentsSize() *core.QSize {
	defer qt.Recovering("QQuickPaintedItem::contentsSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickPaintedItem_ContentsSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickPaintedItem) FillColor() *gui.QColor {
	defer qt.Recovering("QQuickPaintedItem::fillColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QQuickPaintedItem_FillColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickPaintedItem) RenderTarget() QQuickPaintedItem__RenderTarget {
	defer qt.Recovering("QQuickPaintedItem::renderTarget")

	if ptr.Pointer() != nil {
		return QQuickPaintedItem__RenderTarget(C.QQuickPaintedItem_RenderTarget(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickPaintedItem) SetContentsScale(v float64) {
	defer qt.Recovering("QQuickPaintedItem::setContentsScale")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetContentsScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickPaintedItem) SetContentsSize(v core.QSize_ITF) {
	defer qt.Recovering("QQuickPaintedItem::setContentsSize")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetContentsSize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QQuickPaintedItem) SetFillColor(v gui.QColor_ITF) {
	defer qt.Recovering("QQuickPaintedItem::setFillColor")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetFillColor(ptr.Pointer(), gui.PointerFromQColor(v))
	}
}

func (ptr *QQuickPaintedItem) SetRenderTarget(target QQuickPaintedItem__RenderTarget) {
	defer qt.Recovering("QQuickPaintedItem::setRenderTarget")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetRenderTarget(ptr.Pointer(), C.int(target))
	}
}

func (ptr *QQuickPaintedItem) Antialiasing() bool {
	defer qt.Recovering("QQuickPaintedItem::antialiasing")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Antialiasing(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) ContentsBoundingRect() *core.QRectF {
	defer qt.Recovering("QQuickPaintedItem::contentsBoundingRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QQuickPaintedItem_ContentsBoundingRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickPaintedItem) ConnectContentsScaleChanged(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::contentsScaleChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectContentsScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsScaleChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsScaleChanged() {
	defer qt.Recovering("disconnect QQuickPaintedItem::contentsScaleChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsScaleChanged")
	}
}

//export callbackQQuickPaintedItemContentsScaleChanged
func callbackQQuickPaintedItemContentsScaleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::contentsScaleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsScaleChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ContentsScaleChanged() {
	defer qt.Recovering("QQuickPaintedItem::contentsScaleChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ContentsScaleChanged(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ConnectContentsSizeChanged(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::contentsSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectContentsSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsSizeChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsSizeChanged() {
	defer qt.Recovering("disconnect QQuickPaintedItem::contentsSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsSizeChanged")
	}
}

//export callbackQQuickPaintedItemContentsSizeChanged
func callbackQQuickPaintedItemContentsSizeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::contentsSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsSizeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ContentsSizeChanged() {
	defer qt.Recovering("QQuickPaintedItem::contentsSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ContentsSizeChanged(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ConnectFillColorChanged(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::fillColorChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectFillColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fillColorChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectFillColorChanged() {
	defer qt.Recovering("disconnect QQuickPaintedItem::fillColorChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectFillColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fillColorChanged")
	}
}

//export callbackQQuickPaintedItemFillColorChanged
func callbackQQuickPaintedItemFillColorChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::fillColorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fillColorChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) FillColorChanged() {
	defer qt.Recovering("QQuickPaintedItem::fillColorChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FillColorChanged(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) IsTextureProvider() bool {
	defer qt.Recovering("QQuickPaintedItem::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) Mipmap() bool {
	defer qt.Recovering("QQuickPaintedItem::mipmap")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Mipmap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) OpaquePainting() bool {
	defer qt.Recovering("QQuickPaintedItem::opaquePainting")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_OpaquePainting(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) Paint(painter gui.QPainter_ITF) {
	defer qt.Recovering("QQuickPaintedItem::paint")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QQuickPaintedItem) PerformanceHints() QQuickPaintedItem__PerformanceHint {
	defer qt.Recovering("QQuickPaintedItem::performanceHints")

	if ptr.Pointer() != nil {
		return QQuickPaintedItem__PerformanceHint(C.QQuickPaintedItem_PerformanceHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickPaintedItem) ConnectReleaseResources(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::releaseResources")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "releaseResources", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectReleaseResources() {
	defer qt.Recovering("disconnect QQuickPaintedItem::releaseResources")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "releaseResources")
	}
}

//export callbackQQuickPaintedItemReleaseResources
func callbackQQuickPaintedItemReleaseResources(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickPaintedItem) ReleaseResources() {
	defer qt.Recovering("QQuickPaintedItem::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ReleaseResourcesDefault() {
	defer qt.Recovering("QQuickPaintedItem::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ReleaseResourcesDefault(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ConnectRenderTargetChanged(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::renderTargetChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectRenderTargetChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "renderTargetChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectRenderTargetChanged() {
	defer qt.Recovering("disconnect QQuickPaintedItem::renderTargetChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectRenderTargetChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "renderTargetChanged")
	}
}

//export callbackQQuickPaintedItemRenderTargetChanged
func callbackQQuickPaintedItemRenderTargetChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::renderTargetChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "renderTargetChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) RenderTargetChanged() {
	defer qt.Recovering("QQuickPaintedItem::renderTargetChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_RenderTargetChanged(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ResetContentsSize() {
	defer qt.Recovering("QQuickPaintedItem::resetContentsSize")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ResetContentsSize(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) SetAntialiasing(enable bool) {
	defer qt.Recovering("QQuickPaintedItem::setAntialiasing")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetAntialiasing(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QQuickPaintedItem) SetMipmap(enable bool) {
	defer qt.Recovering("QQuickPaintedItem::setMipmap")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetMipmap(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QQuickPaintedItem) SetOpaquePainting(opaque bool) {
	defer qt.Recovering("QQuickPaintedItem::setOpaquePainting")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetOpaquePainting(ptr.Pointer(), C.int(qt.GoBoolToInt(opaque)))
	}
}

func (ptr *QQuickPaintedItem) SetPerformanceHint(hint QQuickPaintedItem__PerformanceHint, enabled bool) {
	defer qt.Recovering("QQuickPaintedItem::setPerformanceHint")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetPerformanceHint(ptr.Pointer(), C.int(hint), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickPaintedItem) SetPerformanceHints(hints QQuickPaintedItem__PerformanceHint) {
	defer qt.Recovering("QQuickPaintedItem::setPerformanceHints")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetPerformanceHints(ptr.Pointer(), C.int(hints))
	}
}

func (ptr *QQuickPaintedItem) TextureProvider() *QSGTextureProvider {
	defer qt.Recovering("QQuickPaintedItem::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickPaintedItem_TextureProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickPaintedItem) Update(rect core.QRect_ITF) {
	defer qt.Recovering("QQuickPaintedItem::update")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_Update(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QQuickPaintedItem) DestroyQQuickPaintedItem() {
	defer qt.Recovering("QQuickPaintedItem::~QQuickPaintedItem")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DestroyQQuickPaintedItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickPaintedItem) ConnectClassBegin(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::classBegin")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "classBegin", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectClassBegin() {
	defer qt.Recovering("disconnect QQuickPaintedItem::classBegin")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "classBegin")
	}
}

//export callbackQQuickPaintedItemClassBegin
func callbackQQuickPaintedItemClassBegin(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::classBegin")

	if signal := qt.GetSignal(C.GoString(ptrName), "classBegin"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ClassBeginDefault()
	}
}

func (ptr *QQuickPaintedItem) ClassBegin() {
	defer qt.Recovering("QQuickPaintedItem::classBegin")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ClassBegin(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ClassBeginDefault() {
	defer qt.Recovering("QQuickPaintedItem::classBegin")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ClassBeginDefault(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ConnectComponentComplete(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::componentComplete")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "componentComplete", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectComponentComplete() {
	defer qt.Recovering("disconnect QQuickPaintedItem::componentComplete")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "componentComplete")
	}
}

//export callbackQQuickPaintedItemComponentComplete
func callbackQQuickPaintedItemComponentComplete(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::componentComplete")

	if signal := qt.GetSignal(C.GoString(ptrName), "componentComplete"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ComponentCompleteDefault()
	}
}

func (ptr *QQuickPaintedItem) ComponentComplete() {
	defer qt.Recovering("QQuickPaintedItem::componentComplete")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ComponentComplete(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ComponentCompleteDefault() {
	defer qt.Recovering("QQuickPaintedItem::componentComplete")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ComponentCompleteDefault(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQQuickPaintedItemDragEnterEvent
func callbackQQuickPaintedItemDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QQuickPaintedItem) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQQuickPaintedItemDragLeaveEvent
func callbackQQuickPaintedItemDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QQuickPaintedItem) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQQuickPaintedItemDragMoveEvent
func callbackQQuickPaintedItemDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QQuickPaintedItem) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQQuickPaintedItemDropEvent
func callbackQQuickPaintedItemDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QQuickPaintedItem) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickPaintedItemFocusInEvent
func callbackQQuickPaintedItemFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickPaintedItem) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickPaintedItemFocusOutEvent
func callbackQQuickPaintedItemFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickPaintedItem) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectGeometryChanged(f func(newGeometry *core.QRectF, oldGeometry *core.QRectF)) {
	defer qt.Recovering("connect QQuickPaintedItem::geometryChanged")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "geometryChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectGeometryChanged() {
	defer qt.Recovering("disconnect QQuickPaintedItem::geometryChanged")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "geometryChanged")
	}
}

//export callbackQQuickPaintedItemGeometryChanged
func callbackQQuickPaintedItemGeometryChanged(ptr unsafe.Pointer, ptrName *C.char, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::geometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	}
}

func (ptr *QQuickPaintedItem) GeometryChanged(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	defer qt.Recovering("QQuickPaintedItem::geometryChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_GeometryChanged(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

func (ptr *QQuickPaintedItem) GeometryChangedDefault(newGeometry core.QRectF_ITF, oldGeometry core.QRectF_ITF) {
	defer qt.Recovering("QQuickPaintedItem::geometryChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_GeometryChangedDefault(ptr.Pointer(), core.PointerFromQRectF(newGeometry), core.PointerFromQRectF(oldGeometry))
	}
}

func (ptr *QQuickPaintedItem) ConnectHoverEnterEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverEnterEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectHoverEnterEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::hoverEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverEnterEvent")
	}
}

//export callbackQQuickPaintedItemHoverEnterEvent
func callbackQQuickPaintedItemHoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) HoverEnterEvent(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverEnterEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickPaintedItem) HoverEnterEventDefault(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::hoverEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverEnterEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectHoverLeaveEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverLeaveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectHoverLeaveEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverLeaveEvent")
	}
}

//export callbackQQuickPaintedItemHoverLeaveEvent
func callbackQQuickPaintedItemHoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) HoverLeaveEvent(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverLeaveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickPaintedItem) HoverLeaveEventDefault(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::hoverLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverLeaveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectHoverMoveEvent(f func(event *gui.QHoverEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hoverMoveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectHoverMoveEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::hoverMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hoverMoveEvent")
	}
}

//export callbackQQuickPaintedItemHoverMoveEvent
func callbackQQuickPaintedItemHoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) HoverMoveEvent(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverMoveEvent(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickPaintedItem) HoverMoveEventDefault(event gui.QHoverEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::hoverMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_HoverMoveEventDefault(ptr.Pointer(), gui.PointerFromQHoverEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQQuickPaintedItemInputMethodEvent
func callbackQQuickPaintedItemInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickPaintedItem) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickPaintedItemKeyPressEvent
func callbackQQuickPaintedItemKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickPaintedItem) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickPaintedItemKeyReleaseEvent
func callbackQQuickPaintedItemKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickPaintedItem) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickPaintedItemMouseDoubleClickEvent
func callbackQQuickPaintedItemMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickPaintedItemMouseMoveEvent
func callbackQQuickPaintedItemMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickPaintedItemMousePressEvent
func callbackQQuickPaintedItemMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickPaintedItemMouseReleaseEvent
func callbackQQuickPaintedItemMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectMouseUngrabEvent(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseUngrabEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMouseUngrabEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseUngrabEvent")
	}
}

//export callbackQQuickPaintedItemMouseUngrabEvent
func callbackQQuickPaintedItemMouseUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::mouseUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseUngrabEventDefault()
	}
}

func (ptr *QQuickPaintedItem) MouseUngrabEvent() {
	defer qt.Recovering("QQuickPaintedItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) MouseUngrabEventDefault() {
	defer qt.Recovering("QQuickPaintedItem::mouseUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_MouseUngrabEventDefault(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ConnectTouchEvent(f func(event *gui.QTouchEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQQuickPaintedItemTouchEvent
func callbackQQuickPaintedItemTouchEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) TouchEvent(event gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QQuickPaintedItem) TouchEventDefault(event gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectTouchUngrabEvent(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::touchUngrabEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchUngrabEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTouchUngrabEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::touchUngrabEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchUngrabEvent")
	}
}

//export callbackQQuickPaintedItemTouchUngrabEvent
func callbackQQuickPaintedItemTouchUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::touchUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).TouchUngrabEventDefault()
	}
}

func (ptr *QQuickPaintedItem) TouchUngrabEvent() {
	defer qt.Recovering("QQuickPaintedItem::touchUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TouchUngrabEvent(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) TouchUngrabEventDefault() {
	defer qt.Recovering("QQuickPaintedItem::touchUngrabEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TouchUngrabEventDefault(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ConnectUpdatePolish(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::updatePolish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updatePolish", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectUpdatePolish() {
	defer qt.Recovering("disconnect QQuickPaintedItem::updatePolish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updatePolish")
	}
}

//export callbackQQuickPaintedItemUpdatePolish
func callbackQQuickPaintedItemUpdatePolish(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::updatePolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).UpdatePolishDefault()
	}
}

func (ptr *QQuickPaintedItem) UpdatePolish() {
	defer qt.Recovering("QQuickPaintedItem::updatePolish")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_UpdatePolish(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) UpdatePolishDefault() {
	defer qt.Recovering("QQuickPaintedItem::updatePolish")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_UpdatePolishDefault(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickPaintedItemWheelEvent
func callbackQQuickPaintedItemWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickPaintedItem) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickPaintedItemTimerEvent
func callbackQQuickPaintedItemTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickPaintedItem) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickPaintedItemChildEvent
func callbackQQuickPaintedItemChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickPaintedItem) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickPaintedItem::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickPaintedItem::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickPaintedItemCustomEvent
func callbackQQuickPaintedItemCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickPaintedItem) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickPaintedItem) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickPaintedItem::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQuickRenderControl struct {
	core.QObject
}

type QQuickRenderControl_ITF interface {
	core.QObject_ITF
	QQuickRenderControl_PTR() *QQuickRenderControl
}

func PointerFromQQuickRenderControl(ptr QQuickRenderControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickRenderControl_PTR().Pointer()
	}
	return nil
}

func NewQQuickRenderControlFromPointer(ptr unsafe.Pointer) *QQuickRenderControl {
	var n = new(QQuickRenderControl)
	n.SetPointer(ptr)
	return n
}

func newQQuickRenderControlFromPointer(ptr unsafe.Pointer) *QQuickRenderControl {
	var n = NewQQuickRenderControlFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickRenderControl_") {
		n.SetObjectName("QQuickRenderControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickRenderControl) QQuickRenderControl_PTR() *QQuickRenderControl {
	return ptr
}

func NewQQuickRenderControl(parent core.QObject_ITF) *QQuickRenderControl {
	defer qt.Recovering("QQuickRenderControl::QQuickRenderControl")

	return newQQuickRenderControlFromPointer(C.QQuickRenderControl_NewQQuickRenderControl(core.PointerFromQObject(parent)))
}

func (ptr *QQuickRenderControl) Grab() *gui.QImage {
	defer qt.Recovering("QQuickRenderControl::grab")

	if ptr.Pointer() != nil {
		return gui.NewQImageFromPointer(C.QQuickRenderControl_Grab(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickRenderControl) Initialize(gl gui.QOpenGLContext_ITF) {
	defer qt.Recovering("QQuickRenderControl::initialize")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Initialize(ptr.Pointer(), gui.PointerFromQOpenGLContext(gl))
	}
}

func (ptr *QQuickRenderControl) Invalidate() {
	defer qt.Recovering("QQuickRenderControl::invalidate")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Invalidate(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PolishItems() {
	defer qt.Recovering("QQuickRenderControl::polishItems")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PolishItems(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) PrepareThread(targetThread core.QThread_ITF) {
	defer qt.Recovering("QQuickRenderControl::prepareThread")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_PrepareThread(ptr.Pointer(), core.PointerFromQThread(targetThread))
	}
}

func (ptr *QQuickRenderControl) Render() {
	defer qt.Recovering("QQuickRenderControl::render")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_Render(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) ConnectRenderRequested(f func()) {
	defer qt.Recovering("connect QQuickRenderControl::renderRequested")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectRenderRequested(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "renderRequested", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderRequested() {
	defer qt.Recovering("disconnect QQuickRenderControl::renderRequested")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectRenderRequested(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "renderRequested")
	}
}

//export callbackQQuickRenderControlRenderRequested
func callbackQQuickRenderControlRenderRequested(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickRenderControl::renderRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "renderRequested"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickRenderControl) RenderRequested() {
	defer qt.Recovering("QQuickRenderControl::renderRequested")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_RenderRequested(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) RenderWindow(offset core.QPoint_ITF) *gui.QWindow {
	defer qt.Recovering("QQuickRenderControl::renderWindow")

	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindow(ptr.Pointer(), core.PointerFromQPoint(offset)))
	}
	return nil
}

func QQuickRenderControl_RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	defer qt.Recovering("QQuickRenderControl::renderWindowFor")

	return gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
}

func (ptr *QQuickRenderControl) ConnectSceneChanged(f func()) {
	defer qt.Recovering("connect QQuickRenderControl::sceneChanged")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectSceneChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneChanged", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectSceneChanged() {
	defer qt.Recovering("disconnect QQuickRenderControl::sceneChanged")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectSceneChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneChanged")
	}
}

//export callbackQQuickRenderControlSceneChanged
func callbackQQuickRenderControlSceneChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickRenderControl::sceneChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickRenderControl) SceneChanged() {
	defer qt.Recovering("QQuickRenderControl::sceneChanged")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_SceneChanged(ptr.Pointer())
	}
}

func (ptr *QQuickRenderControl) Sync() bool {
	defer qt.Recovering("QQuickRenderControl::sync")

	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_Sync(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) DestroyQQuickRenderControl() {
	defer qt.Recovering("QQuickRenderControl::~QQuickRenderControl")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DestroyQQuickRenderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickRenderControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickRenderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickRenderControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickRenderControlTimerEvent
func callbackQQuickRenderControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickRenderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickRenderControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickRenderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickRenderControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickRenderControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickRenderControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickRenderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickRenderControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickRenderControlChildEvent
func callbackQQuickRenderControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickRenderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickRenderControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickRenderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickRenderControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickRenderControl::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickRenderControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickRenderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickRenderControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickRenderControlCustomEvent
func callbackQQuickRenderControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickRenderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickRenderControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickRenderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickRenderControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickRenderControl::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQuickTextDocument struct {
	core.QObject
}

type QQuickTextDocument_ITF interface {
	core.QObject_ITF
	QQuickTextDocument_PTR() *QQuickTextDocument
}

func PointerFromQQuickTextDocument(ptr QQuickTextDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextDocument_PTR().Pointer()
	}
	return nil
}

func NewQQuickTextDocumentFromPointer(ptr unsafe.Pointer) *QQuickTextDocument {
	var n = new(QQuickTextDocument)
	n.SetPointer(ptr)
	return n
}

func newQQuickTextDocumentFromPointer(ptr unsafe.Pointer) *QQuickTextDocument {
	var n = NewQQuickTextDocumentFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickTextDocument_") {
		n.SetObjectName("QQuickTextDocument_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickTextDocument) QQuickTextDocument_PTR() *QQuickTextDocument {
	return ptr
}

func NewQQuickTextDocument(parent QQuickItem_ITF) *QQuickTextDocument {
	defer qt.Recovering("QQuickTextDocument::QQuickTextDocument")

	return newQQuickTextDocumentFromPointer(C.QQuickTextDocument_NewQQuickTextDocument(PointerFromQQuickItem(parent)))
}

func (ptr *QQuickTextDocument) TextDocument() *gui.QTextDocument {
	defer qt.Recovering("QQuickTextDocument::textDocument")

	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QQuickTextDocument_TextDocument(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextDocument) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickTextDocument::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickTextDocument::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickTextDocumentTimerEvent
func callbackQQuickTextDocumentTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextDocument) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextDocument) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickTextDocument::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickTextDocument::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickTextDocumentChildEvent
func callbackQQuickTextDocumentChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextDocument) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextDocument) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickTextDocument::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickTextDocument::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickTextDocumentCustomEvent
func callbackQQuickTextDocumentCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickTextDocument) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickTextDocument) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickTextDocument::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQuickTextureFactory struct {
	core.QObject
}

type QQuickTextureFactory_ITF interface {
	core.QObject_ITF
	QQuickTextureFactory_PTR() *QQuickTextureFactory
}

func PointerFromQQuickTextureFactory(ptr QQuickTextureFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickTextureFactory_PTR().Pointer()
	}
	return nil
}

func NewQQuickTextureFactoryFromPointer(ptr unsafe.Pointer) *QQuickTextureFactory {
	var n = new(QQuickTextureFactory)
	n.SetPointer(ptr)
	return n
}

func newQQuickTextureFactoryFromPointer(ptr unsafe.Pointer) *QQuickTextureFactory {
	var n = NewQQuickTextureFactoryFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickTextureFactory_") {
		n.SetObjectName("QQuickTextureFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickTextureFactory) QQuickTextureFactory_PTR() *QQuickTextureFactory {
	return ptr
}

func (ptr *QQuickTextureFactory) Image() *gui.QImage {
	defer qt.Recovering("QQuickTextureFactory::image")

	if ptr.Pointer() != nil {
		return gui.NewQImageFromPointer(C.QQuickTextureFactory_Image(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextureFactory) CreateTexture(window QQuickWindow_ITF) *QSGTexture {
	defer qt.Recovering("QQuickTextureFactory::createTexture")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickTextureFactory_CreateTexture(ptr.Pointer(), PointerFromQQuickWindow(window)))
	}
	return nil
}

func (ptr *QQuickTextureFactory) TextureByteCount() int {
	defer qt.Recovering("QQuickTextureFactory::textureByteCount")

	if ptr.Pointer() != nil {
		return int(C.QQuickTextureFactory_TextureByteCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickTextureFactory) TextureSize() *core.QSize {
	defer qt.Recovering("QQuickTextureFactory::textureSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickTextureFactory_TextureSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextureFactory) DestroyQQuickTextureFactory() {
	defer qt.Recovering("QQuickTextureFactory::~QQuickTextureFactory")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DestroyQQuickTextureFactory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickTextureFactory) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickTextureFactory::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickTextureFactory::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickTextureFactoryTimerEvent
func callbackQQuickTextureFactoryTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextureFactory) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickTextureFactory) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickTextureFactory::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickTextureFactory::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickTextureFactoryChildEvent
func callbackQQuickTextureFactoryChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextureFactory) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickTextureFactory) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickTextureFactory::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickTextureFactory::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickTextureFactoryCustomEvent
func callbackQQuickTextureFactoryCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickTextureFactory) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickTextureFactory) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickTextureFactory::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQuickView struct {
	QQuickWindow
}

type QQuickView_ITF interface {
	QQuickWindow_ITF
	QQuickView_PTR() *QQuickView
}

func PointerFromQQuickView(ptr QQuickView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickView_PTR().Pointer()
	}
	return nil
}

func NewQQuickViewFromPointer(ptr unsafe.Pointer) *QQuickView {
	var n = new(QQuickView)
	n.SetPointer(ptr)
	return n
}

func newQQuickViewFromPointer(ptr unsafe.Pointer) *QQuickView {
	var n = NewQQuickViewFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickView_") {
		n.SetObjectName("QQuickView_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickView) QQuickView_PTR() *QQuickView {
	return ptr
}

//QQuickView::ResizeMode
type QQuickView__ResizeMode int64

const (
	QQuickView__SizeViewToRootObject = QQuickView__ResizeMode(0)
	QQuickView__SizeRootObjectToView = QQuickView__ResizeMode(1)
)

//QQuickView::Status
type QQuickView__Status int64

const (
	QQuickView__Null    = QQuickView__Status(0)
	QQuickView__Ready   = QQuickView__Status(1)
	QQuickView__Loading = QQuickView__Status(2)
	QQuickView__Error   = QQuickView__Status(3)
)

func (ptr *QQuickView) ResizeMode() QQuickView__ResizeMode {
	defer qt.Recovering("QQuickView::resizeMode")

	if ptr.Pointer() != nil {
		return QQuickView__ResizeMode(C.QQuickView_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickView) SetResizeMode(v QQuickView__ResizeMode) {
	defer qt.Recovering("QQuickView::setResizeMode")

	if ptr.Pointer() != nil {
		C.QQuickView_SetResizeMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QQuickView) Status() QQuickView__Status {
	defer qt.Recovering("QQuickView::status")

	if ptr.Pointer() != nil {
		return QQuickView__Status(C.QQuickView_Status(ptr.Pointer()))
	}
	return 0
}

func NewQQuickView2(engine qml.QQmlEngine_ITF, parent gui.QWindow_ITF) *QQuickView {
	defer qt.Recovering("QQuickView::QQuickView")

	return newQQuickViewFromPointer(C.QQuickView_NewQQuickView2(qml.PointerFromQQmlEngine(engine), gui.PointerFromQWindow(parent)))
}

func NewQQuickView(parent gui.QWindow_ITF) *QQuickView {
	defer qt.Recovering("QQuickView::QQuickView")

	return newQQuickViewFromPointer(C.QQuickView_NewQQuickView(gui.PointerFromQWindow(parent)))
}

func NewQQuickView3(source core.QUrl_ITF, parent gui.QWindow_ITF) *QQuickView {
	defer qt.Recovering("QQuickView::QQuickView")

	return newQQuickViewFromPointer(C.QQuickView_NewQQuickView3(core.PointerFromQUrl(source), gui.PointerFromQWindow(parent)))
}

func (ptr *QQuickView) Engine() *qml.QQmlEngine {
	defer qt.Recovering("QQuickView::engine")

	if ptr.Pointer() != nil {
		return qml.NewQQmlEngineFromPointer(C.QQuickView_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) InitialSize() *core.QSize {
	defer qt.Recovering("QQuickView::initialSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickView_InitialSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickViewKeyPressEvent
func callbackQQuickViewKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickView) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickViewKeyReleaseEvent
func callbackQQuickViewKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickView) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickViewMouseMoveEvent
func callbackQQuickViewMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickView) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickViewMousePressEvent
func callbackQQuickViewMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickView) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickViewMouseReleaseEvent
func callbackQQuickViewMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickView) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) RootContext() *qml.QQmlContext {
	defer qt.Recovering("QQuickView::rootContext")

	if ptr.Pointer() != nil {
		return qml.NewQQmlContextFromPointer(C.QQuickView_RootContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) RootObject() *QQuickItem {
	defer qt.Recovering("QQuickView::rootObject")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickView_RootObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) SetSource(url core.QUrl_ITF) {
	defer qt.Recovering("QQuickView::setSource")

	if ptr.Pointer() != nil {
		C.QQuickView_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickView) Source() *core.QUrl {
	defer qt.Recovering("QQuickView::source")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQuickView_Source(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) ConnectStatusChanged(f func(status QQuickView__Status)) {
	defer qt.Recovering("connect QQuickView::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickView_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQuickView) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQuickView::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickView_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQuickViewStatusChanged
func callbackQQuickViewStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQuickView::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQuickView__Status))(QQuickView__Status(status))
	}

}

func (ptr *QQuickView) StatusChanged(status QQuickView__Status) {
	defer qt.Recovering("QQuickView::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickView_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QQuickView) DestroyQQuickView() {
	defer qt.Recovering("QQuickView::~QQuickView")

	if ptr.Pointer() != nil {
		C.QQuickView_DestroyQQuickView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickView) ConnectExposeEvent(f func(v *gui.QExposeEvent)) {
	defer qt.Recovering("connect QQuickView::exposeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "exposeEvent", f)
	}
}

func (ptr *QQuickView) DisconnectExposeEvent() {
	defer qt.Recovering("disconnect QQuickView::exposeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "exposeEvent")
	}
}

//export callbackQQuickViewExposeEvent
func callbackQQuickViewExposeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*gui.QExposeEvent))(gui.NewQExposeEventFromPointer(v))
	} else {
		NewQQuickViewFromPointer(ptr).ExposeEventDefault(gui.NewQExposeEventFromPointer(v))
	}
}

func (ptr *QQuickView) ExposeEvent(v gui.QExposeEvent_ITF) {
	defer qt.Recovering("QQuickView::exposeEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ExposeEvent(ptr.Pointer(), gui.PointerFromQExposeEvent(v))
	}
}

func (ptr *QQuickView) ExposeEventDefault(v gui.QExposeEvent_ITF) {
	defer qt.Recovering("QQuickView::exposeEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ExposeEventDefault(ptr.Pointer(), gui.PointerFromQExposeEvent(v))
	}
}

func (ptr *QQuickView) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickViewFocusInEvent
func callbackQQuickViewFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickView) FocusInEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickView) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickView) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickViewFocusOutEvent
func callbackQQuickViewFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickView) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickView) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickView) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QQuickView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QQuickView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QQuickView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQQuickViewHideEvent
func callbackQQuickViewHideEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
	} else {
		NewQQuickViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(v))
	}
}

func (ptr *QQuickView) HideEvent(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickView::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QQuickView) HideEventDefault(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickView::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QQuickView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickViewMouseDoubleClickEvent
func callbackQQuickViewMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickView) ConnectResizeEvent(f func(ev *gui.QResizeEvent)) {
	defer qt.Recovering("connect QQuickView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QQuickView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QQuickView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQQuickViewResizeEvent
func callbackQQuickViewResizeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(ev))
	}
}

func (ptr *QQuickView) ResizeEvent(ev gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

func (ptr *QQuickView) ResizeEventDefault(ev gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

func (ptr *QQuickView) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QQuickView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QQuickView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QQuickView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQQuickViewShowEvent
func callbackQQuickViewShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQQuickViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QQuickView) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickView::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QQuickView) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickView::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QQuickView) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickViewWheelEvent
func callbackQQuickViewWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickView) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickView) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickView) ConnectMoveEvent(f func(ev *gui.QMoveEvent)) {
	defer qt.Recovering("connect QQuickView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QQuickView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QQuickView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQQuickViewMoveEvent
func callbackQQuickViewMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(ev))
	}
}

func (ptr *QQuickView) MoveEvent(ev gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickView::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

func (ptr *QQuickView) MoveEventDefault(ev gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickView::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

func (ptr *QQuickView) ConnectTabletEvent(f func(ev *gui.QTabletEvent)) {
	defer qt.Recovering("connect QQuickView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QQuickView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QQuickView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQQuickViewTabletEvent
func callbackQQuickViewTabletEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(ev))
	}
}

func (ptr *QQuickView) TabletEvent(ev gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

func (ptr *QQuickView) TabletEventDefault(ev gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

func (ptr *QQuickView) ConnectTouchEvent(f func(ev *gui.QTouchEvent)) {
	defer qt.Recovering("connect QQuickView::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QQuickView) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QQuickView::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQQuickViewTouchEvent
func callbackQQuickViewTouchEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(ev))
	}
}

func (ptr *QQuickView) TouchEvent(ev gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickView::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

func (ptr *QQuickView) TouchEventDefault(ev gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickView::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

func (ptr *QQuickView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickViewTimerEvent
func callbackQQuickViewTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickView) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickView::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickView) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickView::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickViewChildEvent
func callbackQQuickViewChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickView::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickView::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickViewCustomEvent
func callbackQQuickViewCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickView::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickView::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQuickWidget struct {
	widgets.QWidget
}

type QQuickWidget_ITF interface {
	widgets.QWidget_ITF
	QQuickWidget_PTR() *QQuickWidget
}

func PointerFromQQuickWidget(ptr QQuickWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWidget_PTR().Pointer()
	}
	return nil
}

func NewQQuickWidgetFromPointer(ptr unsafe.Pointer) *QQuickWidget {
	var n = new(QQuickWidget)
	n.SetPointer(ptr)
	return n
}

func newQQuickWidgetFromPointer(ptr unsafe.Pointer) *QQuickWidget {
	var n = NewQQuickWidgetFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickWidget_") {
		n.SetObjectName("QQuickWidget_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickWidget) QQuickWidget_PTR() *QQuickWidget {
	return ptr
}

//QQuickWidget::ResizeMode
type QQuickWidget__ResizeMode int64

const (
	QQuickWidget__SizeViewToRootObject = QQuickWidget__ResizeMode(0)
	QQuickWidget__SizeRootObjectToView = QQuickWidget__ResizeMode(1)
)

//QQuickWidget::Status
type QQuickWidget__Status int64

const (
	QQuickWidget__Null    = QQuickWidget__Status(0)
	QQuickWidget__Ready   = QQuickWidget__Status(1)
	QQuickWidget__Loading = QQuickWidget__Status(2)
	QQuickWidget__Error   = QQuickWidget__Status(3)
)

func (ptr *QQuickWidget) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickWidget::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickWidgetFocusInEvent
func callbackQQuickWidgetFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWidget::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickWidget::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickWidgetFocusOutEvent
func callbackQQuickWidgetFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWidget::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QQuickWidget) ResizeMode() QQuickWidget__ResizeMode {
	defer qt.Recovering("QQuickWidget::resizeMode")

	if ptr.Pointer() != nil {
		return QQuickWidget__ResizeMode(C.QQuickWidget_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWidget) SetResizeMode(v QQuickWidget__ResizeMode) {
	defer qt.Recovering("QQuickWidget::setResizeMode")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetResizeMode(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QQuickWidget) Status() QQuickWidget__Status {
	defer qt.Recovering("QQuickWidget::status")

	if ptr.Pointer() != nil {
		return QQuickWidget__Status(C.QQuickWidget_Status(ptr.Pointer()))
	}
	return 0
}

func NewQQuickWidget2(engine qml.QQmlEngine_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	defer qt.Recovering("QQuickWidget::QQuickWidget")

	return newQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget2(qml.PointerFromQQmlEngine(engine), widgets.PointerFromQWidget(parent)))
}

func NewQQuickWidget(parent widgets.QWidget_ITF) *QQuickWidget {
	defer qt.Recovering("QQuickWidget::QQuickWidget")

	return newQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget(widgets.PointerFromQWidget(parent)))
}

func NewQQuickWidget3(source core.QUrl_ITF, parent widgets.QWidget_ITF) *QQuickWidget {
	defer qt.Recovering("QQuickWidget::QQuickWidget")

	return newQQuickWidgetFromPointer(C.QQuickWidget_NewQQuickWidget3(core.PointerFromQUrl(source), widgets.PointerFromQWidget(parent)))
}

func (ptr *QQuickWidget) ConnectDragEnterEvent(f func(e *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QQuickWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QQuickWidget::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQQuickWidgetDragEnterEvent
func callbackQQuickWidgetDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DragEnterEvent(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QQuickWidget) DragEnterEventDefault(e gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectDragLeaveEvent(f func(e *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QQuickWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQQuickWidgetDragLeaveEvent
func callbackQQuickWidgetDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DragLeaveEvent(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QQuickWidget) DragLeaveEventDefault(e gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectDragMoveEvent(f func(e *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QQuickWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQQuickWidgetDragMoveEvent
func callbackQQuickWidgetDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DragMoveEvent(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QQuickWidget) DragMoveEventDefault(e gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectDropEvent(f func(e *gui.QDropEvent)) {
	defer qt.Recovering("connect QQuickWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QQuickWidget::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQQuickWidgetDropEvent
func callbackQQuickWidgetDropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) DropEvent(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QQuickWidget) DropEventDefault(e gui.QDropEvent_ITF) {
	defer qt.Recovering("QQuickWidget::dropEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(e))
	}
}

func (ptr *QQuickWidget) Engine() *qml.QQmlEngine {
	defer qt.Recovering("QQuickWidget::engine")

	if ptr.Pointer() != nil {
		return qml.NewQQmlEngineFromPointer(C.QQuickWidget_Engine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWidget::event")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWidget) Format() *gui.QSurfaceFormat {
	defer qt.Recovering("QQuickWidget::format")

	if ptr.Pointer() != nil {
		return gui.NewQSurfaceFormatFromPointer(C.QQuickWidget_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) GrabFramebuffer() *gui.QImage {
	defer qt.Recovering("QQuickWidget::grabFramebuffer")

	if ptr.Pointer() != nil {
		return gui.NewQImageFromPointer(C.QQuickWidget_GrabFramebuffer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQQuickWidgetHideEvent
func callbackQQuickWidgetHideEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
	} else {
		NewQQuickWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(v))
	}
}

func (ptr *QQuickWidget) HideEvent(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QQuickWidget) HideEventDefault(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QQuickWidget) InitialSize() *core.QSize {
	defer qt.Recovering("QQuickWidget::initialSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWidget_InitialSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickWidget::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickWidgetKeyPressEvent
func callbackQQuickWidgetKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWidget::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickWidgetKeyReleaseEvent
func callbackQQuickWidgetKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWidget::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectMouseDoubleClickEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickWidgetMouseDoubleClickEvent
func callbackQQuickWidgetMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MouseDoubleClickEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MouseDoubleClickEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectMouseMoveEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickWidgetMouseMoveEvent
func callbackQQuickWidgetMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MouseMoveEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectMousePressEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickWidget::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickWidgetMousePressEvent
func callbackQQuickWidgetMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MousePressEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) ConnectMouseReleaseEvent(f func(e *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickWidgetMouseReleaseEvent
func callbackQQuickWidgetMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) MouseReleaseEventDefault(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickWidget) QuickWindow() *QQuickWindow {
	defer qt.Recovering("QQuickWidget::quickWindow")

	if ptr.Pointer() != nil {
		return NewQQuickWindowFromPointer(C.QQuickWidget_QuickWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) RootContext() *qml.QQmlContext {
	defer qt.Recovering("QQuickWidget::rootContext")

	if ptr.Pointer() != nil {
		return qml.NewQQmlContextFromPointer(C.QQuickWidget_RootContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) RootObject() *QQuickItem {
	defer qt.Recovering("QQuickWidget::rootObject")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWidget_RootObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	defer qt.Recovering("connect QQuickWidget::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectSceneGraphError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphError", f)
	}
}

func (ptr *QQuickWidget) DisconnectSceneGraphError() {
	defer qt.Recovering("disconnect QQuickWidget::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphError")
	}
}

//export callbackQQuickWidgetSceneGraphError
func callbackQQuickWidgetSceneGraphError(ptr unsafe.Pointer, ptrName *C.char, error C.int, message *C.char) {
	defer qt.Recovering("callback QQuickWidget::sceneGraphError")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
	}

}

func (ptr *QQuickWidget) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {
	defer qt.Recovering("QQuickWidget::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SceneGraphError(ptr.Pointer(), C.int(error), C.CString(message))
	}
}

func (ptr *QQuickWidget) SetClearColor(color gui.QColor_ITF) {
	defer qt.Recovering("QQuickWidget::setClearColor")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetClearColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QQuickWidget) SetFormat(format gui.QSurfaceFormat_ITF) {
	defer qt.Recovering("QQuickWidget::setFormat")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetFormat(ptr.Pointer(), gui.PointerFromQSurfaceFormat(format))
	}
}

func (ptr *QQuickWidget) SetSource(url core.QUrl_ITF) {
	defer qt.Recovering("QQuickWidget::setSource")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QQuickWidget) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QQuickWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QQuickWidget::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQQuickWidgetShowEvent
func callbackQQuickWidgetShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QQuickWidget) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QQuickWidget) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QQuickWidget) ConnectStatusChanged(f func(status QQuickWidget__Status)) {
	defer qt.Recovering("connect QQuickWidget::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectStatusChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "statusChanged", f)
	}
}

func (ptr *QQuickWidget) DisconnectStatusChanged() {
	defer qt.Recovering("disconnect QQuickWidget::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectStatusChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "statusChanged")
	}
}

//export callbackQQuickWidgetStatusChanged
func callbackQQuickWidgetStatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQuickWidget::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQuickWidget__Status))(QQuickWidget__Status(status))
	}

}

func (ptr *QQuickWidget) StatusChanged(status QQuickWidget__Status) {
	defer qt.Recovering("QQuickWidget::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickWidget_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QQuickWidget) Source() *core.QUrl {
	defer qt.Recovering("QQuickWidget::source")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QQuickWidget_Source(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) ConnectWheelEvent(f func(e *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickWidget::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickWidgetWheelEvent
func callbackQQuickWidgetWheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
}

func (ptr *QQuickWidget) WheelEvent(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QQuickWidget) WheelEventDefault(e gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickWidget::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(e))
	}
}

func (ptr *QQuickWidget) DestroyQQuickWidget() {
	defer qt.Recovering("QQuickWidget::~QQuickWidget")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DestroyQQuickWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWidget) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QQuickWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QQuickWidget::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQQuickWidgetActionEvent
func callbackQQuickWidgetActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QQuickWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QQuickWidget) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QQuickWidget::actionEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QQuickWidget::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQQuickWidgetEnterEvent
func callbackQQuickWidgetEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::enterEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQQuickWidgetLeaveEvent
func callbackQQuickWidgetLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::leaveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QQuickWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QQuickWidget::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQQuickWidgetMoveEvent
func callbackQQuickWidgetMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QQuickWidget) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickWidget::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QQuickWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QQuickWidget::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQQuickWidgetPaintEvent
func callbackQQuickWidgetPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QQuickWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QQuickWidget) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QQuickWidget::paintEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QQuickWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QQuickWidget::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQQuickWidgetSetVisible
func callbackQQuickWidgetSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QQuickWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QQuickWidget) SetVisible(visible bool) {
	defer qt.Recovering("QQuickWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QQuickWidget) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QQuickWidget::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QQuickWidget) ConnectChangeEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QQuickWidget::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQQuickWidgetChangeEvent
func callbackQQuickWidgetChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ChangeEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) ChangeEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::changeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QQuickWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QQuickWidget::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQQuickWidgetCloseEvent
func callbackQQuickWidgetCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QQuickWidget) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QQuickWidget::closeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QQuickWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QQuickWidget::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQQuickWidgetContextMenuEvent
func callbackQQuickWidgetContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QQuickWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QQuickWidget) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QQuickWidget::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QQuickWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QQuickWidget) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QQuickWidget::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQQuickWidgetInitPainter
func callbackQQuickWidgetInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQQuickWidgetFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QQuickWidget) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QQuickWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QQuickWidget_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QQuickWidget) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QQuickWidget::initPainter")

	if ptr.Pointer() != nil {
		C.QQuickWidget_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QQuickWidget) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QQuickWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QQuickWidget::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQQuickWidgetInputMethodEvent
func callbackQQuickWidgetInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickWidget) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QQuickWidget::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QQuickWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QQuickWidget::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQQuickWidgetResizeEvent
func callbackQQuickWidgetResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QQuickWidget) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickWidget::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QQuickWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QQuickWidget::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQQuickWidgetTabletEvent
func callbackQQuickWidgetTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QQuickWidget) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickWidget::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickWidget::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickWidgetTimerEvent
func callbackQQuickWidgetTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWidget) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickWidget::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickWidget::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickWidgetChildEvent
func callbackQQuickWidgetChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWidget) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickWidget::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWidget) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickWidget::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickWidgetCustomEvent
func callbackQQuickWidgetCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWidget) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWidget) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWidget::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QQuickWindow struct {
	gui.QWindow
}

type QQuickWindow_ITF interface {
	gui.QWindow_ITF
	QQuickWindow_PTR() *QQuickWindow
}

func PointerFromQQuickWindow(ptr QQuickWindow_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func NewQQuickWindowFromPointer(ptr unsafe.Pointer) *QQuickWindow {
	var n = new(QQuickWindow)
	n.SetPointer(ptr)
	return n
}

func newQQuickWindowFromPointer(ptr unsafe.Pointer) *QQuickWindow {
	var n = NewQQuickWindowFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickWindow_") {
		n.SetObjectName("QQuickWindow_" + qt.Identifier())
	}
	return n
}

func (ptr *QQuickWindow) QQuickWindow_PTR() *QQuickWindow {
	return ptr
}

//QQuickWindow::CreateTextureOption
type QQuickWindow__CreateTextureOption int64

const (
	QQuickWindow__TextureHasAlphaChannel = QQuickWindow__CreateTextureOption(0x0001)
	QQuickWindow__TextureHasMipmaps      = QQuickWindow__CreateTextureOption(0x0002)
	QQuickWindow__TextureOwnsGLTexture   = QQuickWindow__CreateTextureOption(0x0004)
	QQuickWindow__TextureCanUseAtlas     = QQuickWindow__CreateTextureOption(0x0008)
)

//QQuickWindow::RenderStage
type QQuickWindow__RenderStage int64

const (
	QQuickWindow__BeforeSynchronizingStage = QQuickWindow__RenderStage(0)
	QQuickWindow__AfterSynchronizingStage  = QQuickWindow__RenderStage(1)
	QQuickWindow__BeforeRenderingStage     = QQuickWindow__RenderStage(2)
	QQuickWindow__AfterRenderingStage      = QQuickWindow__RenderStage(3)
	QQuickWindow__AfterSwapStage           = QQuickWindow__RenderStage(4)
)

//QQuickWindow::SceneGraphError
type QQuickWindow__SceneGraphError int64

const (
	QQuickWindow__ContextNotAvailable = QQuickWindow__SceneGraphError(1)
)

func (ptr *QQuickWindow) ActiveFocusItem() *QQuickItem {
	defer qt.Recovering("QQuickWindow::activeFocusItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_ActiveFocusItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) Color() *gui.QColor {
	defer qt.Recovering("QQuickWindow::color")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QQuickWindow_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ContentItem() *QQuickItem {
	defer qt.Recovering("QQuickWindow::contentItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_ContentItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) SetColor(color gui.QColor_ITF) {
	defer qt.Recovering("QQuickWindow::setColor")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func NewQQuickWindow(parent gui.QWindow_ITF) *QQuickWindow {
	defer qt.Recovering("QQuickWindow::QQuickWindow")

	return newQQuickWindowFromPointer(C.QQuickWindow_NewQQuickWindow(gui.PointerFromQWindow(parent)))
}

func (ptr *QQuickWindow) AccessibleRoot() *gui.QAccessibleInterface {
	defer qt.Recovering("QQuickWindow::accessibleRoot")

	if ptr.Pointer() != nil {
		return gui.NewQAccessibleInterfaceFromPointer(C.QQuickWindow_AccessibleRoot(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectActiveFocusItemChanged(f func()) {
	defer qt.Recovering("connect QQuickWindow::activeFocusItemChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectActiveFocusItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeFocusItemChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectActiveFocusItemChanged() {
	defer qt.Recovering("disconnect QQuickWindow::activeFocusItemChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectActiveFocusItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeFocusItemChanged")
	}
}

//export callbackQQuickWindowActiveFocusItemChanged
func callbackQQuickWindowActiveFocusItemChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::activeFocusItemChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeFocusItemChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ActiveFocusItemChanged() {
	defer qt.Recovering("QQuickWindow::activeFocusItemChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ActiveFocusItemChanged(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectAfterAnimating(f func()) {
	defer qt.Recovering("connect QQuickWindow::afterAnimating")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterAnimating(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterAnimating", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterAnimating() {
	defer qt.Recovering("disconnect QQuickWindow::afterAnimating")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterAnimating(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterAnimating")
	}
}

//export callbackQQuickWindowAfterAnimating
func callbackQQuickWindowAfterAnimating(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::afterAnimating")

	if signal := qt.GetSignal(C.GoString(ptrName), "afterAnimating"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) AfterAnimating() {
	defer qt.Recovering("QQuickWindow::afterAnimating")

	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterAnimating(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectAfterRendering(f func()) {
	defer qt.Recovering("connect QQuickWindow::afterRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterRendering(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterRendering() {
	defer qt.Recovering("disconnect QQuickWindow::afterRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterRendering")
	}
}

//export callbackQQuickWindowAfterRendering
func callbackQQuickWindowAfterRendering(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::afterRendering")

	if signal := qt.GetSignal(C.GoString(ptrName), "afterRendering"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) AfterRendering() {
	defer qt.Recovering("QQuickWindow::afterRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterRendering(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectAfterSynchronizing(f func()) {
	defer qt.Recovering("connect QQuickWindow::afterSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectAfterSynchronizing(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "afterSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectAfterSynchronizing() {
	defer qt.Recovering("disconnect QQuickWindow::afterSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectAfterSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "afterSynchronizing")
	}
}

//export callbackQQuickWindowAfterSynchronizing
func callbackQQuickWindowAfterSynchronizing(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::afterSynchronizing")

	if signal := qt.GetSignal(C.GoString(ptrName), "afterSynchronizing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) AfterSynchronizing() {
	defer qt.Recovering("QQuickWindow::afterSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterSynchronizing(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectBeforeRendering(f func()) {
	defer qt.Recovering("connect QQuickWindow::beforeRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeRendering(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeRendering", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeRendering() {
	defer qt.Recovering("disconnect QQuickWindow::beforeRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeRendering(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeRendering")
	}
}

//export callbackQQuickWindowBeforeRendering
func callbackQQuickWindowBeforeRendering(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::beforeRendering")

	if signal := qt.GetSignal(C.GoString(ptrName), "beforeRendering"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) BeforeRendering() {
	defer qt.Recovering("QQuickWindow::beforeRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_BeforeRendering(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectBeforeSynchronizing(f func()) {
	defer qt.Recovering("connect QQuickWindow::beforeSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectBeforeSynchronizing(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "beforeSynchronizing", f)
	}
}

func (ptr *QQuickWindow) DisconnectBeforeSynchronizing() {
	defer qt.Recovering("disconnect QQuickWindow::beforeSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectBeforeSynchronizing(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "beforeSynchronizing")
	}
}

//export callbackQQuickWindowBeforeSynchronizing
func callbackQQuickWindowBeforeSynchronizing(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::beforeSynchronizing")

	if signal := qt.GetSignal(C.GoString(ptrName), "beforeSynchronizing"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) BeforeSynchronizing() {
	defer qt.Recovering("QQuickWindow::beforeSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_BeforeSynchronizing(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ClearBeforeRendering() bool {
	defer qt.Recovering("QQuickWindow::clearBeforeRendering")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_ClearBeforeRendering(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) ConnectColorChanged(f func(v *gui.QColor)) {
	defer qt.Recovering("connect QQuickWindow::colorChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "colorChanged", f)
	}
}

func (ptr *QQuickWindow) DisconnectColorChanged() {
	defer qt.Recovering("disconnect QQuickWindow::colorChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "colorChanged")
	}
}

//export callbackQQuickWindowColorChanged
func callbackQQuickWindowColorChanged(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::colorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(v))
	}

}

func (ptr *QQuickWindow) ColorChanged(v gui.QColor_ITF) {
	defer qt.Recovering("QQuickWindow::colorChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(v))
	}
}

func (ptr *QQuickWindow) CreateTextureFromImage2(image gui.QImage_ITF) *QSGTexture {
	defer qt.Recovering("QQuickWindow::createTextureFromImage")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage2(ptr.Pointer(), gui.PointerFromQImage(image)))
	}
	return nil
}

func (ptr *QQuickWindow) CreateTextureFromImage(image gui.QImage_ITF, options QQuickWindow__CreateTextureOption) *QSGTexture {
	defer qt.Recovering("QQuickWindow::createTextureFromImage")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickWindow_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.int(options)))
	}
	return nil
}

func (ptr *QQuickWindow) EffectiveDevicePixelRatio() float64 {
	defer qt.Recovering("QQuickWindow::effectiveDevicePixelRatio")

	if ptr.Pointer() != nil {
		return float64(C.QQuickWindow_EffectiveDevicePixelRatio(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWindow) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWindow::event")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWindow) ConnectExposeEvent(f func(v *gui.QExposeEvent)) {
	defer qt.Recovering("connect QQuickWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "exposeEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectExposeEvent() {
	defer qt.Recovering("disconnect QQuickWindow::exposeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "exposeEvent")
	}
}

//export callbackQQuickWindowExposeEvent
func callbackQQuickWindowExposeEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::exposeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "exposeEvent"); signal != nil {
		signal.(func(*gui.QExposeEvent))(gui.NewQExposeEventFromPointer(v))
	} else {
		NewQQuickWindowFromPointer(ptr).ExposeEventDefault(gui.NewQExposeEventFromPointer(v))
	}
}

func (ptr *QQuickWindow) ExposeEvent(v gui.QExposeEvent_ITF) {
	defer qt.Recovering("QQuickWindow::exposeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ExposeEvent(ptr.Pointer(), gui.PointerFromQExposeEvent(v))
	}
}

func (ptr *QQuickWindow) ExposeEventDefault(v gui.QExposeEvent_ITF) {
	defer qt.Recovering("QQuickWindow::exposeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ExposeEventDefault(ptr.Pointer(), gui.PointerFromQExposeEvent(v))
	}
}

func (ptr *QQuickWindow) ConnectFocusInEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QQuickWindow::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQQuickWindowFocusInEvent
func callbackQQuickWindowFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) FocusInEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWindow::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickWindow) FocusInEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWindow::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickWindow) ConnectFocusOutEvent(f func(ev *gui.QFocusEvent)) {
	defer qt.Recovering("connect QQuickWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QQuickWindow::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQQuickWindowFocusOutEvent
func callbackQQuickWindowFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWindow::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickWindow) FocusOutEventDefault(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWindow::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickWindow) ConnectFrameSwapped(f func()) {
	defer qt.Recovering("connect QQuickWindow::frameSwapped")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectFrameSwapped(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "frameSwapped", f)
	}
}

func (ptr *QQuickWindow) DisconnectFrameSwapped() {
	defer qt.Recovering("disconnect QQuickWindow::frameSwapped")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectFrameSwapped(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "frameSwapped")
	}
}

//export callbackQQuickWindowFrameSwapped
func callbackQQuickWindowFrameSwapped(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::frameSwapped")

	if signal := qt.GetSignal(C.GoString(ptrName), "frameSwapped"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) FrameSwapped() {
	defer qt.Recovering("QQuickWindow::frameSwapped")

	if ptr.Pointer() != nil {
		C.QQuickWindow_FrameSwapped(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) GrabWindow() *gui.QImage {
	defer qt.Recovering("QQuickWindow::grabWindow")

	if ptr.Pointer() != nil {
		return gui.NewQImageFromPointer(C.QQuickWindow_GrabWindow(ptr.Pointer()))
	}
	return nil
}

func QQuickWindow_HasDefaultAlphaBuffer() bool {
	defer qt.Recovering("QQuickWindow::hasDefaultAlphaBuffer")

	return C.QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer() != 0
}

func (ptr *QQuickWindow) ConnectHideEvent(f func(v *gui.QHideEvent)) {
	defer qt.Recovering("connect QQuickWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QQuickWindow::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQQuickWindowHideEvent
func callbackQQuickWindowHideEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(v))
	} else {
		NewQQuickWindowFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(v))
	}
}

func (ptr *QQuickWindow) HideEvent(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickWindow::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QQuickWindow) HideEventDefault(v gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickWindow::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(v))
	}
}

func (ptr *QQuickWindow) IncubationController() *qml.QQmlIncubationController {
	defer qt.Recovering("QQuickWindow::incubationController")

	if ptr.Pointer() != nil {
		return qml.NewQQmlIncubationControllerFromPointer(C.QQuickWindow_IncubationController(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) IsPersistentOpenGLContext() bool {
	defer qt.Recovering("QQuickWindow::isPersistentOpenGLContext")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentOpenGLContext(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsPersistentSceneGraph() bool {
	defer qt.Recovering("QQuickWindow::isPersistentSceneGraph")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsPersistentSceneGraph(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) IsSceneGraphInitialized() bool {
	defer qt.Recovering("QQuickWindow::isSceneGraphInitialized")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_IsSceneGraphInitialized(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) ConnectKeyPressEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QQuickWindow::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQQuickWindowKeyPressEvent
func callbackQQuickWindowKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWindowFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWindow) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWindow::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWindow) KeyPressEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWindow::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWindow) ConnectKeyReleaseEvent(f func(e *gui.QKeyEvent)) {
	defer qt.Recovering("connect QQuickWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QQuickWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQQuickWindowKeyReleaseEvent
func callbackQQuickWindowKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWindowFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
	}
}

func (ptr *QQuickWindow) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWindow) KeyReleaseEventDefault(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWindow) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QQuickWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQQuickWindowMouseDoubleClickEvent
func callbackQQuickWindowMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseGrabberItem() *QQuickItem {
	defer qt.Recovering("QQuickWindow::mouseGrabberItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_MouseGrabberItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QQuickWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQQuickWindowMouseMoveEvent
func callbackQQuickWindowMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QQuickWindow::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQQuickWindowMousePressEvent
func callbackQQuickWindowMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QQuickWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QQuickWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQQuickWindowMouseReleaseEvent
func callbackQQuickWindowMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) OpenglContext() *gui.QOpenGLContext {
	defer qt.Recovering("QQuickWindow::openglContext")

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLContextFromPointer(C.QQuickWindow_OpenglContext(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ConnectOpenglContextCreated(f func(context *gui.QOpenGLContext)) {
	defer qt.Recovering("connect QQuickWindow::openglContextCreated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectOpenglContextCreated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "openglContextCreated", f)
	}
}

func (ptr *QQuickWindow) DisconnectOpenglContextCreated() {
	defer qt.Recovering("disconnect QQuickWindow::openglContextCreated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectOpenglContextCreated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "openglContextCreated")
	}
}

//export callbackQQuickWindowOpenglContextCreated
func callbackQQuickWindowOpenglContextCreated(ptr unsafe.Pointer, ptrName *C.char, context unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::openglContextCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "openglContextCreated"); signal != nil {
		signal.(func(*gui.QOpenGLContext))(gui.NewQOpenGLContextFromPointer(context))
	}

}

func (ptr *QQuickWindow) OpenglContextCreated(context gui.QOpenGLContext_ITF) {
	defer qt.Recovering("QQuickWindow::openglContextCreated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_OpenglContextCreated(ptr.Pointer(), gui.PointerFromQOpenGLContext(context))
	}
}

func (ptr *QQuickWindow) ReleaseResources() {
	defer qt.Recovering("QQuickWindow::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RenderTarget() *gui.QOpenGLFramebufferObject {
	defer qt.Recovering("QQuickWindow::renderTarget")

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLFramebufferObjectFromPointer(C.QQuickWindow_RenderTarget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) RenderTargetSize() *core.QSize {
	defer qt.Recovering("QQuickWindow::renderTargetSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWindow_RenderTargetSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) ResetOpenGLState() {
	defer qt.Recovering("QQuickWindow::resetOpenGLState")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ResetOpenGLState(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectResizeEvent(f func(ev *gui.QResizeEvent)) {
	defer qt.Recovering("connect QQuickWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QQuickWindow::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQQuickWindowResizeEvent
func callbackQQuickWindowResizeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) ResizeEvent(ev gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickWindow::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

func (ptr *QQuickWindow) ResizeEventDefault(ev gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickWindow::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

func (ptr *QQuickWindow) ConnectSceneGraphAboutToStop(f func()) {
	defer qt.Recovering("connect QQuickWindow::sceneGraphAboutToStop")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphAboutToStop(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphAboutToStop() {
	defer qt.Recovering("disconnect QQuickWindow::sceneGraphAboutToStop")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphAboutToStop(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphAboutToStop")
	}
}

//export callbackQQuickWindowSceneGraphAboutToStop
func callbackQQuickWindowSceneGraphAboutToStop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphAboutToStop")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphAboutToStop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) SceneGraphAboutToStop() {
	defer qt.Recovering("QQuickWindow::sceneGraphAboutToStop")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphAboutToStop(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectSceneGraphError(f func(error QQuickWindow__SceneGraphError, message string)) {
	defer qt.Recovering("connect QQuickWindow::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphError", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphError() {
	defer qt.Recovering("disconnect QQuickWindow::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphError")
	}
}

//export callbackQQuickWindowSceneGraphError
func callbackQQuickWindowSceneGraphError(ptr unsafe.Pointer, ptrName *C.char, error C.int, message *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphError")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
	}

}

func (ptr *QQuickWindow) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {
	defer qt.Recovering("QQuickWindow::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphError(ptr.Pointer(), C.int(error), C.CString(message))
	}
}

func (ptr *QQuickWindow) ConnectSceneGraphInitialized(f func()) {
	defer qt.Recovering("connect QQuickWindow::sceneGraphInitialized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInitialized(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInitialized", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInitialized() {
	defer qt.Recovering("disconnect QQuickWindow::sceneGraphInitialized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInitialized(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInitialized")
	}
}

//export callbackQQuickWindowSceneGraphInitialized
func callbackQQuickWindowSceneGraphInitialized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphInitialized")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphInitialized"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) SceneGraphInitialized() {
	defer qt.Recovering("QQuickWindow::sceneGraphInitialized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphInitialized(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectSceneGraphInvalidated(f func()) {
	defer qt.Recovering("connect QQuickWindow::sceneGraphInvalidated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectSceneGraphInvalidated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphInvalidated", f)
	}
}

func (ptr *QQuickWindow) DisconnectSceneGraphInvalidated() {
	defer qt.Recovering("disconnect QQuickWindow::sceneGraphInvalidated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectSceneGraphInvalidated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphInvalidated")
	}
}

//export callbackQQuickWindowSceneGraphInvalidated
func callbackQQuickWindowSceneGraphInvalidated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphInvalidated")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphInvalidated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) SceneGraphInvalidated() {
	defer qt.Recovering("QQuickWindow::sceneGraphInvalidated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphInvalidated(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ScheduleRenderJob(job core.QRunnable_ITF, stage QQuickWindow__RenderStage) {
	defer qt.Recovering("QQuickWindow::scheduleRenderJob")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ScheduleRenderJob(ptr.Pointer(), core.PointerFromQRunnable(job), C.int(stage))
	}
}

func (ptr *QQuickWindow) SendEvent(item QQuickItem_ITF, e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWindow::sendEvent")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_SendEvent(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	defer qt.Recovering("QQuickWindow::setClearBeforeRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetClearBeforeRendering(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func QQuickWindow_SetDefaultAlphaBuffer(useAlpha bool) {
	defer qt.Recovering("QQuickWindow::setDefaultAlphaBuffer")

	C.QQuickWindow_QQuickWindow_SetDefaultAlphaBuffer(C.int(qt.GoBoolToInt(useAlpha)))
}

func (ptr *QQuickWindow) SetPersistentOpenGLContext(persistent bool) {
	defer qt.Recovering("QQuickWindow::setPersistentOpenGLContext")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentOpenGLContext(ptr.Pointer(), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetPersistentSceneGraph(persistent bool) {
	defer qt.Recovering("QQuickWindow::setPersistentSceneGraph")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetPersistentSceneGraph(ptr.Pointer(), C.int(qt.GoBoolToInt(persistent)))
	}
}

func (ptr *QQuickWindow) SetRenderTarget(fbo gui.QOpenGLFramebufferObject_ITF) {
	defer qt.Recovering("QQuickWindow::setRenderTarget")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetRenderTarget(ptr.Pointer(), gui.PointerFromQOpenGLFramebufferObject(fbo))
	}
}

func (ptr *QQuickWindow) ConnectShowEvent(f func(v *gui.QShowEvent)) {
	defer qt.Recovering("connect QQuickWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QQuickWindow::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQQuickWindowShowEvent
func callbackQQuickWindowShowEvent(ptr unsafe.Pointer, ptrName *C.char, v unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(v))
	} else {
		NewQQuickWindowFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(v))
	}
}

func (ptr *QQuickWindow) ShowEvent(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickWindow::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QQuickWindow) ShowEventDefault(v gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickWindow::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(v))
	}
}

func (ptr *QQuickWindow) Update() {
	defer qt.Recovering("QQuickWindow::update")

	if ptr.Pointer() != nil {
		C.QQuickWindow_Update(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QQuickWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QQuickWindow::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQQuickWindowWheelEvent
func callbackQQuickWindowWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickWindow::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickWindow) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickWindow::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickWindow) DestroyQQuickWindow() {
	defer qt.Recovering("QQuickWindow::~QQuickWindow")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DestroyQQuickWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWindow) ConnectMoveEvent(f func(ev *gui.QMoveEvent)) {
	defer qt.Recovering("connect QQuickWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QQuickWindow::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQQuickWindowMoveEvent
func callbackQQuickWindowMoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) MoveEvent(ev gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickWindow::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

func (ptr *QQuickWindow) MoveEventDefault(ev gui.QMoveEvent_ITF) {
	defer qt.Recovering("QQuickWindow::moveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(ev))
	}
}

func (ptr *QQuickWindow) ConnectTabletEvent(f func(ev *gui.QTabletEvent)) {
	defer qt.Recovering("connect QQuickWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QQuickWindow::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQQuickWindowTabletEvent
func callbackQQuickWindowTabletEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) TabletEvent(ev gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickWindow::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

func (ptr *QQuickWindow) TabletEventDefault(ev gui.QTabletEvent_ITF) {
	defer qt.Recovering("QQuickWindow::tabletEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(ev))
	}
}

func (ptr *QQuickWindow) ConnectTouchEvent(f func(ev *gui.QTouchEvent)) {
	defer qt.Recovering("connect QQuickWindow::touchEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "touchEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectTouchEvent() {
	defer qt.Recovering("disconnect QQuickWindow::touchEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "touchEvent")
	}
}

//export callbackQQuickWindowTouchEvent
func callbackQQuickWindowTouchEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(ev))
	}
}

func (ptr *QQuickWindow) TouchEvent(ev gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickWindow::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_TouchEvent(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

func (ptr *QQuickWindow) TouchEventDefault(ev gui.QTouchEvent_ITF) {
	defer qt.Recovering("QQuickWindow::touchEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_TouchEventDefault(ptr.Pointer(), gui.PointerFromQTouchEvent(ev))
	}
}

func (ptr *QQuickWindow) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickWindow::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQQuickWindowTimerEvent
func callbackQQuickWindowTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickWindow::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWindow) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickWindow::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickWindow) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickWindow::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQQuickWindowChildEvent
func callbackQQuickWindowChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickWindow::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWindow) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickWindow::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickWindow) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickWindow::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQQuickWindowCustomEvent
func callbackQQuickWindowCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickWindow) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWindow::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickWindow) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickWindow::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSGAbstractRenderer struct {
	core.QObject
}

type QSGAbstractRenderer_ITF interface {
	core.QObject_ITF
	QSGAbstractRenderer_PTR() *QSGAbstractRenderer
}

func PointerFromQSGAbstractRenderer(ptr QSGAbstractRenderer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGAbstractRenderer_PTR().Pointer()
	}
	return nil
}

func NewQSGAbstractRendererFromPointer(ptr unsafe.Pointer) *QSGAbstractRenderer {
	var n = new(QSGAbstractRenderer)
	n.SetPointer(ptr)
	return n
}

func newQSGAbstractRendererFromPointer(ptr unsafe.Pointer) *QSGAbstractRenderer {
	var n = NewQSGAbstractRendererFromPointer(ptr)
	for len(n.ObjectName()) < len("QSGAbstractRenderer_") {
		n.SetObjectName("QSGAbstractRenderer_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGAbstractRenderer) QSGAbstractRenderer_PTR() *QSGAbstractRenderer {
	return ptr
}

//QSGAbstractRenderer::ClearModeBit
type QSGAbstractRenderer__ClearModeBit int64

const (
	QSGAbstractRenderer__ClearColorBuffer   = QSGAbstractRenderer__ClearModeBit(0x0001)
	QSGAbstractRenderer__ClearDepthBuffer   = QSGAbstractRenderer__ClearModeBit(0x0002)
	QSGAbstractRenderer__ClearStencilBuffer = QSGAbstractRenderer__ClearModeBit(0x0004)
)

func (ptr *QSGAbstractRenderer) ClearColor() *gui.QColor {
	defer qt.Recovering("QSGAbstractRenderer::clearColor")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QSGAbstractRenderer_ClearColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGAbstractRenderer) ClearMode() QSGAbstractRenderer__ClearModeBit {
	defer qt.Recovering("QSGAbstractRenderer::clearMode")

	if ptr.Pointer() != nil {
		return QSGAbstractRenderer__ClearModeBit(C.QSGAbstractRenderer_ClearMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGAbstractRenderer) DeviceRect() *core.QRect {
	defer qt.Recovering("QSGAbstractRenderer::deviceRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QSGAbstractRenderer_DeviceRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGAbstractRenderer) ConnectSceneGraphChanged(f func()) {
	defer qt.Recovering("connect QSGAbstractRenderer::sceneGraphChanged")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ConnectSceneGraphChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "sceneGraphChanged", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectSceneGraphChanged() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::sceneGraphChanged")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectSceneGraphChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "sceneGraphChanged")
	}
}

//export callbackQSGAbstractRendererSceneGraphChanged
func callbackQSGAbstractRendererSceneGraphChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGAbstractRenderer::sceneGraphChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGAbstractRenderer) SceneGraphChanged() {
	defer qt.Recovering("QSGAbstractRenderer::sceneGraphChanged")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SceneGraphChanged(ptr.Pointer())
	}
}

func (ptr *QSGAbstractRenderer) SetClearColor(color gui.QColor_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setClearColor")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetClearColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QSGAbstractRenderer) SetClearMode(mode QSGAbstractRenderer__ClearModeBit) {
	defer qt.Recovering("QSGAbstractRenderer::setClearMode")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetClearMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSGAbstractRenderer) SetDeviceRect(rect core.QRect_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setDeviceRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetDeviceRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QSGAbstractRenderer) SetDeviceRect2(size core.QSize_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setDeviceRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetDeviceRect2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrix(matrix gui.QMatrix4x4_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setProjectionMatrix")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetProjectionMatrix(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QSGAbstractRenderer) SetProjectionMatrixToRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setProjectionMatrixToRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetProjectionMatrixToRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGAbstractRenderer) SetViewportRect(rect core.QRect_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setViewportRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetViewportRect(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QSGAbstractRenderer) SetViewportRect2(size core.QSize_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::setViewportRect")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_SetViewportRect2(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QSGAbstractRenderer) ViewportRect() *core.QRect {
	defer qt.Recovering("QSGAbstractRenderer::viewportRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QSGAbstractRenderer_ViewportRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGAbstractRenderer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSGAbstractRenderer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSGAbstractRendererTimerEvent
func callbackQSGAbstractRendererTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGAbstractRenderer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGAbstractRenderer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGAbstractRenderer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGAbstractRenderer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSGAbstractRenderer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSGAbstractRendererChildEvent
func callbackQSGAbstractRendererChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGAbstractRenderer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGAbstractRenderer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::childEvent")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGAbstractRenderer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::childEvent")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGAbstractRenderer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSGAbstractRenderer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSGAbstractRendererCustomEvent
func callbackQSGAbstractRendererCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGAbstractRenderer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGAbstractRenderer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::customEvent")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGAbstractRenderer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::customEvent")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSGBasicGeometryNode struct {
	QSGNode
}

type QSGBasicGeometryNode_ITF interface {
	QSGNode_ITF
	QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode
}

func PointerFromQSGBasicGeometryNode(ptr QSGBasicGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func NewQSGBasicGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGBasicGeometryNode {
	var n = new(QSGBasicGeometryNode)
	n.SetPointer(ptr)
	return n
}

func newQSGBasicGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGBasicGeometryNode {
	var n = NewQSGBasicGeometryNodeFromPointer(ptr)
	return n
}

func (ptr *QSGBasicGeometryNode) QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode {
	return ptr
}

func (ptr *QSGBasicGeometryNode) Geometry2() *QSGGeometry {
	defer qt.Recovering("QSGBasicGeometryNode::geometry")

	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) Geometry() *QSGGeometry {
	defer qt.Recovering("QSGBasicGeometryNode::geometry")

	if ptr.Pointer() != nil {
		return NewQSGGeometryFromPointer(C.QSGBasicGeometryNode_Geometry(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGBasicGeometryNode) SetGeometry(geometry QSGGeometry_ITF) {
	defer qt.Recovering("QSGBasicGeometryNode::setGeometry")

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_SetGeometry(ptr.Pointer(), PointerFromQSGGeometry(geometry))
	}
}

func (ptr *QSGBasicGeometryNode) DestroyQSGBasicGeometryNode() {
	defer qt.Recovering("QSGBasicGeometryNode::~QSGBasicGeometryNode")

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_DestroyQSGBasicGeometryNode(ptr.Pointer())
	}
}

func (ptr *QSGBasicGeometryNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGBasicGeometryNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGBasicGeometryNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGBasicGeometryNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGBasicGeometryNodePreprocess
func callbackQSGBasicGeometryNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGBasicGeometryNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGBasicGeometryNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGBasicGeometryNode) Preprocess() {
	defer qt.Recovering("QSGBasicGeometryNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGBasicGeometryNode) PreprocessDefault() {
	defer qt.Recovering("QSGBasicGeometryNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGBasicGeometryNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGClipNode struct {
	QSGBasicGeometryNode
}

type QSGClipNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGClipNode_PTR() *QSGClipNode
}

func PointerFromQSGClipNode(ptr QSGClipNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGClipNode_PTR().Pointer()
	}
	return nil
}

func NewQSGClipNodeFromPointer(ptr unsafe.Pointer) *QSGClipNode {
	var n = new(QSGClipNode)
	n.SetPointer(ptr)
	return n
}

func newQSGClipNodeFromPointer(ptr unsafe.Pointer) *QSGClipNode {
	var n = NewQSGClipNodeFromPointer(ptr)
	return n
}

func (ptr *QSGClipNode) QSGClipNode_PTR() *QSGClipNode {
	return ptr
}

func NewQSGClipNode() *QSGClipNode {
	defer qt.Recovering("QSGClipNode::QSGClipNode")

	return newQSGClipNodeFromPointer(C.QSGClipNode_NewQSGClipNode())
}

func (ptr *QSGClipNode) ClipRect() *core.QRectF {
	defer qt.Recovering("QSGClipNode::clipRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGClipNode_ClipRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGClipNode) IsRectangular() bool {
	defer qt.Recovering("QSGClipNode::isRectangular")

	if ptr.Pointer() != nil {
		return C.QSGClipNode_IsRectangular(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGClipNode) SetClipRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QSGClipNode::setClipRect")

	if ptr.Pointer() != nil {
		C.QSGClipNode_SetClipRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGClipNode) SetIsRectangular(rectHint bool) {
	defer qt.Recovering("QSGClipNode::setIsRectangular")

	if ptr.Pointer() != nil {
		C.QSGClipNode_SetIsRectangular(ptr.Pointer(), C.int(qt.GoBoolToInt(rectHint)))
	}
}

func (ptr *QSGClipNode) DestroyQSGClipNode() {
	defer qt.Recovering("QSGClipNode::~QSGClipNode")

	if ptr.Pointer() != nil {
		C.QSGClipNode_DestroyQSGClipNode(ptr.Pointer())
	}
}

func (ptr *QSGClipNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGClipNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGClipNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGClipNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGClipNodePreprocess
func callbackQSGClipNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGClipNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGClipNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGClipNode) Preprocess() {
	defer qt.Recovering("QSGClipNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGClipNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGClipNode) PreprocessDefault() {
	defer qt.Recovering("QSGClipNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGClipNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGDynamicTexture struct {
	QSGTexture
}

type QSGDynamicTexture_ITF interface {
	QSGTexture_ITF
	QSGDynamicTexture_PTR() *QSGDynamicTexture
}

func PointerFromQSGDynamicTexture(ptr QSGDynamicTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGDynamicTexture_PTR().Pointer()
	}
	return nil
}

func NewQSGDynamicTextureFromPointer(ptr unsafe.Pointer) *QSGDynamicTexture {
	var n = new(QSGDynamicTexture)
	n.SetPointer(ptr)
	return n
}

func newQSGDynamicTextureFromPointer(ptr unsafe.Pointer) *QSGDynamicTexture {
	var n = NewQSGDynamicTextureFromPointer(ptr)
	for len(n.ObjectName()) < len("QSGDynamicTexture_") {
		n.SetObjectName("QSGDynamicTexture_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGDynamicTexture) QSGDynamicTexture_PTR() *QSGDynamicTexture {
	return ptr
}

func (ptr *QSGDynamicTexture) UpdateTexture() bool {
	defer qt.Recovering("QSGDynamicTexture::updateTexture")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_UpdateTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSGDynamicTexture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSGDynamicTexture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSGDynamicTextureTimerEvent
func callbackQSGDynamicTextureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGDynamicTexture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGDynamicTexture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGDynamicTexture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSGDynamicTexture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSGDynamicTexture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSGDynamicTextureChildEvent
func callbackQSGDynamicTextureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGDynamicTexture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::childEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGDynamicTexture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::childEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGDynamicTexture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSGDynamicTexture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSGDynamicTexture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSGDynamicTextureCustomEvent
func callbackQSGDynamicTextureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGDynamicTexture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::customEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGDynamicTexture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSGDynamicTexture::customEvent")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSGEngine struct {
	core.QObject
}

type QSGEngine_ITF interface {
	core.QObject_ITF
	QSGEngine_PTR() *QSGEngine
}

func PointerFromQSGEngine(ptr QSGEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGEngine_PTR().Pointer()
	}
	return nil
}

func NewQSGEngineFromPointer(ptr unsafe.Pointer) *QSGEngine {
	var n = new(QSGEngine)
	n.SetPointer(ptr)
	return n
}

func newQSGEngineFromPointer(ptr unsafe.Pointer) *QSGEngine {
	var n = NewQSGEngineFromPointer(ptr)
	for len(n.ObjectName()) < len("QSGEngine_") {
		n.SetObjectName("QSGEngine_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGEngine) QSGEngine_PTR() *QSGEngine {
	return ptr
}

//QSGEngine::CreateTextureOption
type QSGEngine__CreateTextureOption int64

const (
	QSGEngine__TextureHasAlphaChannel = QSGEngine__CreateTextureOption(0x0001)
	QSGEngine__TextureOwnsGLTexture   = QSGEngine__CreateTextureOption(0x0004)
	QSGEngine__TextureCanUseAtlas     = QSGEngine__CreateTextureOption(0x0008)
)

func NewQSGEngine(parent core.QObject_ITF) *QSGEngine {
	defer qt.Recovering("QSGEngine::QSGEngine")

	return newQSGEngineFromPointer(C.QSGEngine_NewQSGEngine(core.PointerFromQObject(parent)))
}

func (ptr *QSGEngine) CreateRenderer() *QSGAbstractRenderer {
	defer qt.Recovering("QSGEngine::createRenderer")

	if ptr.Pointer() != nil {
		return NewQSGAbstractRendererFromPointer(C.QSGEngine_CreateRenderer(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGEngine) CreateTextureFromImage(image gui.QImage_ITF, options QSGEngine__CreateTextureOption) *QSGTexture {
	defer qt.Recovering("QSGEngine::createTextureFromImage")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGEngine_CreateTextureFromImage(ptr.Pointer(), gui.PointerFromQImage(image), C.int(options)))
	}
	return nil
}

func (ptr *QSGEngine) Initialize(context gui.QOpenGLContext_ITF) {
	defer qt.Recovering("QSGEngine::initialize")

	if ptr.Pointer() != nil {
		C.QSGEngine_Initialize(ptr.Pointer(), gui.PointerFromQOpenGLContext(context))
	}
}

func (ptr *QSGEngine) Invalidate() {
	defer qt.Recovering("QSGEngine::invalidate")

	if ptr.Pointer() != nil {
		C.QSGEngine_Invalidate(ptr.Pointer())
	}
}

func (ptr *QSGEngine) DestroyQSGEngine() {
	defer qt.Recovering("QSGEngine::~QSGEngine")

	if ptr.Pointer() != nil {
		C.QSGEngine_DestroyQSGEngine(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGEngine) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSGEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSGEngine) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSGEngine::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSGEngineTimerEvent
func callbackQSGEngineTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGEngine) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGEngine::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGEngine) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSGEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSGEngine) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSGEngine::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSGEngineChildEvent
func callbackQSGEngineChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGEngine) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGEngine::childEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGEngine) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSGEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSGEngine) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSGEngine::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSGEngineCustomEvent
func callbackQSGEngineCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGEngine) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSGEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGEngine) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSGEngine::customEvent")

	if ptr.Pointer() != nil {
		C.QSGEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSGFlatColorMaterial struct {
	QSGMaterial
}

type QSGFlatColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial
}

func PointerFromQSGFlatColorMaterial(ptr QSGFlatColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGFlatColorMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGFlatColorMaterialFromPointer(ptr unsafe.Pointer) *QSGFlatColorMaterial {
	var n = new(QSGFlatColorMaterial)
	n.SetPointer(ptr)
	return n
}

func newQSGFlatColorMaterialFromPointer(ptr unsafe.Pointer) *QSGFlatColorMaterial {
	var n = NewQSGFlatColorMaterialFromPointer(ptr)
	return n
}

func (ptr *QSGFlatColorMaterial) QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial {
	return ptr
}

func (ptr *QSGFlatColorMaterial) Color() *gui.QColor {
	defer qt.Recovering("QSGFlatColorMaterial::color")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QSGFlatColorMaterial_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGFlatColorMaterial) SetColor(color gui.QColor_ITF) {
	defer qt.Recovering("QSGFlatColorMaterial::setColor")

	if ptr.Pointer() != nil {
		C.QSGFlatColorMaterial_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

type QSGGeometry struct {
	ptr unsafe.Pointer
}

type QSGGeometry_ITF interface {
	QSGGeometry_PTR() *QSGGeometry
}

func (p *QSGGeometry) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGGeometry) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGGeometry(ptr QSGGeometry_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometry_PTR().Pointer()
	}
	return nil
}

func NewQSGGeometryFromPointer(ptr unsafe.Pointer) *QSGGeometry {
	var n = new(QSGGeometry)
	n.SetPointer(ptr)
	return n
}

func newQSGGeometryFromPointer(ptr unsafe.Pointer) *QSGGeometry {
	var n = NewQSGGeometryFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSGGeometry_") {
		n.SetObjectNameAbs("QSGGeometry_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGGeometry) QSGGeometry_PTR() *QSGGeometry {
	return ptr
}

//QSGGeometry::DataPattern
type QSGGeometry__DataPattern int64

const (
	QSGGeometry__AlwaysUploadPattern = QSGGeometry__DataPattern(0)
	QSGGeometry__StreamPattern       = QSGGeometry__DataPattern(1)
	QSGGeometry__DynamicPattern      = QSGGeometry__DataPattern(2)
	QSGGeometry__StaticPattern       = QSGGeometry__DataPattern(3)
)

func (ptr *QSGGeometry) Allocate(vertexCount int, indexCount int) {
	defer qt.Recovering("QSGGeometry::allocate")

	if ptr.Pointer() != nil {
		C.QSGGeometry_Allocate(ptr.Pointer(), C.int(vertexCount), C.int(indexCount))
	}
}

func (ptr *QSGGeometry) AttributeCount() int {
	defer qt.Recovering("QSGGeometry::attributeCount")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_AttributeCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexCount() int {
	defer qt.Recovering("QSGGeometry::indexCount")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_IndexCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexData() unsafe.Pointer {
	defer qt.Recovering("QSGGeometry::indexData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_IndexData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) IndexData2() unsafe.Pointer {
	defer qt.Recovering("QSGGeometry::indexData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_IndexData2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) IndexDataPattern() QSGGeometry__DataPattern {
	defer qt.Recovering("QSGGeometry::indexDataPattern")

	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_IndexDataPattern(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) IndexType() int {
	defer qt.Recovering("QSGGeometry::indexType")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_IndexType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) MarkIndexDataDirty() {
	defer qt.Recovering("QSGGeometry::markIndexDataDirty")

	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkIndexDataDirty(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) MarkVertexDataDirty() {
	defer qt.Recovering("QSGGeometry::markVertexDataDirty")

	if ptr.Pointer() != nil {
		C.QSGGeometry_MarkVertexDataDirty(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) SetIndexDataPattern(p QSGGeometry__DataPattern) {
	defer qt.Recovering("QSGGeometry::setIndexDataPattern")

	if ptr.Pointer() != nil {
		C.QSGGeometry_SetIndexDataPattern(ptr.Pointer(), C.int(p))
	}
}

func (ptr *QSGGeometry) SetVertexDataPattern(p QSGGeometry__DataPattern) {
	defer qt.Recovering("QSGGeometry::setVertexDataPattern")

	if ptr.Pointer() != nil {
		C.QSGGeometry_SetVertexDataPattern(ptr.Pointer(), C.int(p))
	}
}

func (ptr *QSGGeometry) SizeOfIndex() int {
	defer qt.Recovering("QSGGeometry::sizeOfIndex")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_SizeOfIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) SizeOfVertex() int {
	defer qt.Recovering("QSGGeometry::sizeOfVertex")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_SizeOfVertex(ptr.Pointer()))
	}
	return 0
}

func QSGGeometry_UpdateRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {
	defer qt.Recovering("QSGGeometry::updateRectGeometry")

	C.QSGGeometry_QSGGeometry_UpdateRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect))
}

func QSGGeometry_UpdateTexturedRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF, textureRect core.QRectF_ITF) {
	defer qt.Recovering("QSGGeometry::updateTexturedRectGeometry")

	C.QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect), core.PointerFromQRectF(textureRect))
}

func (ptr *QSGGeometry) VertexCount() int {
	defer qt.Recovering("QSGGeometry::vertexCount")

	if ptr.Pointer() != nil {
		return int(C.QSGGeometry_VertexCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) VertexData() unsafe.Pointer {
	defer qt.Recovering("QSGGeometry::vertexData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_VertexData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) VertexData2() unsafe.Pointer {
	defer qt.Recovering("QSGGeometry::vertexData")

	if ptr.Pointer() != nil {
		return unsafe.Pointer(C.QSGGeometry_VertexData2(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometry) VertexDataPattern() QSGGeometry__DataPattern {
	defer qt.Recovering("QSGGeometry::vertexDataPattern")

	if ptr.Pointer() != nil {
		return QSGGeometry__DataPattern(C.QSGGeometry_VertexDataPattern(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGGeometry) DestroyQSGGeometry() {
	defer qt.Recovering("QSGGeometry::~QSGGeometry")

	if ptr.Pointer() != nil {
		C.QSGGeometry_DestroyQSGGeometry(ptr.Pointer())
	}
}

func (ptr *QSGGeometry) ObjectNameAbs() string {
	defer qt.Recovering("QSGGeometry::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGGeometry_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGGeometry) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSGGeometry::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSGGeometry_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSGGeometryNode struct {
	QSGBasicGeometryNode
}

type QSGGeometryNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGGeometryNode_PTR() *QSGGeometryNode
}

func PointerFromQSGGeometryNode(ptr QSGGeometryNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func NewQSGGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGGeometryNode {
	var n = new(QSGGeometryNode)
	n.SetPointer(ptr)
	return n
}

func newQSGGeometryNodeFromPointer(ptr unsafe.Pointer) *QSGGeometryNode {
	var n = NewQSGGeometryNodeFromPointer(ptr)
	return n
}

func (ptr *QSGGeometryNode) QSGGeometryNode_PTR() *QSGGeometryNode {
	return ptr
}

func NewQSGGeometryNode() *QSGGeometryNode {
	defer qt.Recovering("QSGGeometryNode::QSGGeometryNode")

	return newQSGGeometryNodeFromPointer(C.QSGGeometryNode_NewQSGGeometryNode())
}

func (ptr *QSGGeometryNode) Material() *QSGMaterial {
	defer qt.Recovering("QSGGeometryNode::material")

	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_Material(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometryNode) OpaqueMaterial() *QSGMaterial {
	defer qt.Recovering("QSGGeometryNode::opaqueMaterial")

	if ptr.Pointer() != nil {
		return NewQSGMaterialFromPointer(C.QSGGeometryNode_OpaqueMaterial(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGGeometryNode) SetMaterial(material QSGMaterial_ITF) {
	defer qt.Recovering("QSGGeometryNode::setMaterial")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

func (ptr *QSGGeometryNode) SetOpaqueMaterial(material QSGMaterial_ITF) {
	defer qt.Recovering("QSGGeometryNode::setOpaqueMaterial")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_SetOpaqueMaterial(ptr.Pointer(), PointerFromQSGMaterial(material))
	}
}

func (ptr *QSGGeometryNode) DestroyQSGGeometryNode() {
	defer qt.Recovering("QSGGeometryNode::~QSGGeometryNode")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_DestroyQSGGeometryNode(ptr.Pointer())
	}
}

func (ptr *QSGGeometryNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGGeometryNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGGeometryNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGGeometryNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGGeometryNodePreprocess
func callbackQSGGeometryNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGGeometryNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGGeometryNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGGeometryNode) Preprocess() {
	defer qt.Recovering("QSGGeometryNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGGeometryNode) PreprocessDefault() {
	defer qt.Recovering("QSGGeometryNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGGeometryNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGMaterial struct {
	ptr unsafe.Pointer
}

type QSGMaterial_ITF interface {
	QSGMaterial_PTR() *QSGMaterial
}

func (p *QSGMaterial) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGMaterial) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGMaterial(ptr QSGMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialFromPointer(ptr unsafe.Pointer) *QSGMaterial {
	var n = new(QSGMaterial)
	n.SetPointer(ptr)
	return n
}

func newQSGMaterialFromPointer(ptr unsafe.Pointer) *QSGMaterial {
	var n = NewQSGMaterialFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSGMaterial_") {
		n.SetObjectNameAbs("QSGMaterial_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGMaterial) QSGMaterial_PTR() *QSGMaterial {
	return ptr
}

//QSGMaterial::Flag
type QSGMaterial__Flag int64

const (
	QSGMaterial__Blending                          = QSGMaterial__Flag(0x0001)
	QSGMaterial__RequiresDeterminant               = QSGMaterial__Flag(0x0002)
	QSGMaterial__RequiresFullMatrixExceptTranslate = QSGMaterial__Flag(0x0004 | QSGMaterial__RequiresDeterminant)
	QSGMaterial__RequiresFullMatrix                = QSGMaterial__Flag(0x0008 | QSGMaterial__RequiresFullMatrixExceptTranslate)
	QSGMaterial__CustomCompileStep                 = QSGMaterial__Flag(0x0010)
)

func (ptr *QSGMaterial) Compare(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

func (ptr *QSGMaterial) CreateShader() *QSGMaterialShader {
	defer qt.Recovering("QSGMaterial::createShader")

	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGMaterial) Flags() QSGMaterial__Flag {
	defer qt.Recovering("QSGMaterial::flags")

	if ptr.Pointer() != nil {
		return QSGMaterial__Flag(C.QSGMaterial_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGMaterial) SetFlag(flags QSGMaterial__Flag, on bool) {
	defer qt.Recovering("QSGMaterial::setFlag")

	if ptr.Pointer() != nil {
		C.QSGMaterial_SetFlag(ptr.Pointer(), C.int(flags), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QSGMaterial) Type() *QSGMaterialType {
	defer qt.Recovering("QSGMaterial::type")

	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGMaterial_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGMaterial) ObjectNameAbs() string {
	defer qt.Recovering("QSGMaterial::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGMaterial_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGMaterial) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSGMaterial::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSGMaterial_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSGMaterialShader struct {
	ptr unsafe.Pointer
}

type QSGMaterialShader_ITF interface {
	QSGMaterialShader_PTR() *QSGMaterialShader
}

func (p *QSGMaterialShader) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGMaterialShader) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGMaterialShader(ptr QSGMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialShader_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGMaterialShader {
	var n = new(QSGMaterialShader)
	n.SetPointer(ptr)
	return n
}

func newQSGMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGMaterialShader {
	var n = NewQSGMaterialShaderFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSGMaterialShader_") {
		n.SetObjectNameAbs("QSGMaterialShader_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGMaterialShader) QSGMaterialShader_PTR() *QSGMaterialShader {
	return ptr
}

func (ptr *QSGMaterialShader) ConnectActivate(f func()) {
	defer qt.Recovering("connect QSGMaterialShader::activate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "activate", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectActivate() {
	defer qt.Recovering("disconnect QSGMaterialShader::activate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "activate")
	}
}

//export callbackQSGMaterialShaderActivate
func callbackQSGMaterialShaderActivate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::activate")

	if signal := qt.GetSignal(C.GoString(ptrName), "activate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).ActivateDefault()
	}
}

func (ptr *QSGMaterialShader) Activate() {
	defer qt.Recovering("QSGMaterialShader::activate")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Activate(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) ActivateDefault() {
	defer qt.Recovering("QSGMaterialShader::activate")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_ActivateDefault(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) ConnectCompile(f func()) {
	defer qt.Recovering("connect QSGMaterialShader::compile")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "compile", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectCompile() {
	defer qt.Recovering("disconnect QSGMaterialShader::compile")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "compile")
	}
}

//export callbackQSGMaterialShaderCompile
func callbackQSGMaterialShaderCompile(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::compile")

	if signal := qt.GetSignal(C.GoString(ptrName), "compile"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).CompileDefault()
	}
}

func (ptr *QSGMaterialShader) Compile() {
	defer qt.Recovering("QSGMaterialShader::compile")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Compile(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) CompileDefault() {
	defer qt.Recovering("QSGMaterialShader::compile")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_CompileDefault(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) ConnectDeactivate(f func()) {
	defer qt.Recovering("connect QSGMaterialShader::deactivate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "deactivate", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectDeactivate() {
	defer qt.Recovering("disconnect QSGMaterialShader::deactivate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "deactivate")
	}
}

//export callbackQSGMaterialShaderDeactivate
func callbackQSGMaterialShaderDeactivate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::deactivate")

	if signal := qt.GetSignal(C.GoString(ptrName), "deactivate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).DeactivateDefault()
	}
}

func (ptr *QSGMaterialShader) Deactivate() {
	defer qt.Recovering("QSGMaterialShader::deactivate")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Deactivate(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) DeactivateDefault() {
	defer qt.Recovering("QSGMaterialShader::deactivate")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_DeactivateDefault(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) ConnectInitialize(f func()) {
	defer qt.Recovering("connect QSGMaterialShader::initialize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "initialize", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectInitialize() {
	defer qt.Recovering("disconnect QSGMaterialShader::initialize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "initialize")
	}
}

//export callbackQSGMaterialShaderInitialize
func callbackQSGMaterialShaderInitialize(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::initialize")

	if signal := qt.GetSignal(C.GoString(ptrName), "initialize"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).InitializeDefault()
	}
}

func (ptr *QSGMaterialShader) Initialize() {
	defer qt.Recovering("QSGMaterialShader::initialize")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_Initialize(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) InitializeDefault() {
	defer qt.Recovering("QSGMaterialShader::initialize")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_InitializeDefault(ptr.Pointer())
	}
}

func (ptr *QSGMaterialShader) Program() *gui.QOpenGLShaderProgram {
	defer qt.Recovering("QSGMaterialShader::program")

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLShaderProgramFromPointer(C.QSGMaterialShader_Program(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGMaterialShader) ObjectNameAbs() string {
	defer qt.Recovering("QSGMaterialShader::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGMaterialShader_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGMaterialShader) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSGMaterialShader::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSGMaterialType struct {
	ptr unsafe.Pointer
}

type QSGMaterialType_ITF interface {
	QSGMaterialType_PTR() *QSGMaterialType
}

func (p *QSGMaterialType) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGMaterialType) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGMaterialType(ptr QSGMaterialType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGMaterialType_PTR().Pointer()
	}
	return nil
}

func NewQSGMaterialTypeFromPointer(ptr unsafe.Pointer) *QSGMaterialType {
	var n = new(QSGMaterialType)
	n.SetPointer(ptr)
	return n
}

func newQSGMaterialTypeFromPointer(ptr unsafe.Pointer) *QSGMaterialType {
	var n = NewQSGMaterialTypeFromPointer(ptr)
	return n
}

func (ptr *QSGMaterialType) QSGMaterialType_PTR() *QSGMaterialType {
	return ptr
}

type QSGNode struct {
	ptr unsafe.Pointer
}

type QSGNode_ITF interface {
	QSGNode_PTR() *QSGNode
}

func (p *QSGNode) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSGNode) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSGNode(ptr QSGNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGNode_PTR().Pointer()
	}
	return nil
}

func NewQSGNodeFromPointer(ptr unsafe.Pointer) *QSGNode {
	var n = new(QSGNode)
	n.SetPointer(ptr)
	return n
}

func newQSGNodeFromPointer(ptr unsafe.Pointer) *QSGNode {
	var n = NewQSGNodeFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSGNode_") {
		n.SetObjectNameAbs("QSGNode_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGNode) QSGNode_PTR() *QSGNode {
	return ptr
}

//QSGNode::DirtyStateBit
type QSGNode__DirtyStateBit int64

const (
	QSGNode__DirtySubtreeBlocked = QSGNode__DirtyStateBit(0x0080)
	QSGNode__DirtyMatrix         = QSGNode__DirtyStateBit(0x0100)
	QSGNode__DirtyNodeAdded      = QSGNode__DirtyStateBit(0x0400)
	QSGNode__DirtyNodeRemoved    = QSGNode__DirtyStateBit(0x0800)
	QSGNode__DirtyGeometry       = QSGNode__DirtyStateBit(0x1000)
	QSGNode__DirtyMaterial       = QSGNode__DirtyStateBit(0x2000)
	QSGNode__DirtyOpacity        = QSGNode__DirtyStateBit(0x4000)
)

//QSGNode::Flag
type QSGNode__Flag int64

const (
	QSGNode__OwnedByParent      = QSGNode__Flag(0x0001)
	QSGNode__UsePreprocess      = QSGNode__Flag(0x0002)
	QSGNode__OwnsGeometry       = QSGNode__Flag(0x00010000)
	QSGNode__OwnsMaterial       = QSGNode__Flag(0x00020000)
	QSGNode__OwnsOpaqueMaterial = QSGNode__Flag(0x00040000)
	QSGNode__InternalReserved   = QSGNode__Flag(0x01000000)
)

//QSGNode::NodeType
type QSGNode__NodeType int64

const (
	QSGNode__BasicNodeType     = QSGNode__NodeType(0)
	QSGNode__GeometryNodeType  = QSGNode__NodeType(1)
	QSGNode__TransformNodeType = QSGNode__NodeType(2)
	QSGNode__ClipNodeType      = QSGNode__NodeType(3)
	QSGNode__OpacityNodeType   = QSGNode__NodeType(4)
)

func (ptr *QSGNode) ChildAtIndex(i int) *QSGNode {
	defer qt.Recovering("QSGNode::childAtIndex")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_ChildAtIndex(ptr.Pointer(), C.int(i)))
	}
	return nil
}

func (ptr *QSGNode) ChildCount() int {
	defer qt.Recovering("QSGNode::childCount")

	if ptr.Pointer() != nil {
		return int(C.QSGNode_ChildCount(ptr.Pointer()))
	}
	return 0
}

func NewQSGNode() *QSGNode {
	defer qt.Recovering("QSGNode::QSGNode")

	return newQSGNodeFromPointer(C.QSGNode_NewQSGNode())
}

func (ptr *QSGNode) AppendChildNode(node QSGNode_ITF) {
	defer qt.Recovering("QSGNode::appendChildNode")

	if ptr.Pointer() != nil {
		C.QSGNode_AppendChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

func (ptr *QSGNode) FirstChild() *QSGNode {
	defer qt.Recovering("QSGNode::firstChild")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_FirstChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) Flags() QSGNode__Flag {
	defer qt.Recovering("QSGNode::flags")

	if ptr.Pointer() != nil {
		return QSGNode__Flag(C.QSGNode_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGNode) InsertChildNodeAfter(node QSGNode_ITF, after QSGNode_ITF) {
	defer qt.Recovering("QSGNode::insertChildNodeAfter")

	if ptr.Pointer() != nil {
		C.QSGNode_InsertChildNodeAfter(ptr.Pointer(), PointerFromQSGNode(node), PointerFromQSGNode(after))
	}
}

func (ptr *QSGNode) InsertChildNodeBefore(node QSGNode_ITF, before QSGNode_ITF) {
	defer qt.Recovering("QSGNode::insertChildNodeBefore")

	if ptr.Pointer() != nil {
		C.QSGNode_InsertChildNodeBefore(ptr.Pointer(), PointerFromQSGNode(node), PointerFromQSGNode(before))
	}
}

func (ptr *QSGNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGNode) LastChild() *QSGNode {
	defer qt.Recovering("QSGNode::lastChild")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_LastChild(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) MarkDirty(bits QSGNode__DirtyStateBit) {
	defer qt.Recovering("QSGNode::markDirty")

	if ptr.Pointer() != nil {
		C.QSGNode_MarkDirty(ptr.Pointer(), C.int(bits))
	}
}

func (ptr *QSGNode) NextSibling() *QSGNode {
	defer qt.Recovering("QSGNode::nextSibling")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_NextSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) Parent() *QSGNode {
	defer qt.Recovering("QSGNode::parent")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_Parent(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) PrependChildNode(node QSGNode_ITF) {
	defer qt.Recovering("QSGNode::prependChildNode")

	if ptr.Pointer() != nil {
		C.QSGNode_PrependChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

func (ptr *QSGNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGNodePreprocess
func callbackQSGNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGNode) Preprocess() {
	defer qt.Recovering("QSGNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGNode) PreprocessDefault() {
	defer qt.Recovering("QSGNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGNode_PreprocessDefault(ptr.Pointer())
	}
}

func (ptr *QSGNode) PreviousSibling() *QSGNode {
	defer qt.Recovering("QSGNode::previousSibling")

	if ptr.Pointer() != nil {
		return NewQSGNodeFromPointer(C.QSGNode_PreviousSibling(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGNode) RemoveAllChildNodes() {
	defer qt.Recovering("QSGNode::removeAllChildNodes")

	if ptr.Pointer() != nil {
		C.QSGNode_RemoveAllChildNodes(ptr.Pointer())
	}
}

func (ptr *QSGNode) RemoveChildNode(node QSGNode_ITF) {
	defer qt.Recovering("QSGNode::removeChildNode")

	if ptr.Pointer() != nil {
		C.QSGNode_RemoveChildNode(ptr.Pointer(), PointerFromQSGNode(node))
	}
}

func (ptr *QSGNode) SetFlag(f QSGNode__Flag, enabled bool) {
	defer qt.Recovering("QSGNode::setFlag")

	if ptr.Pointer() != nil {
		C.QSGNode_SetFlag(ptr.Pointer(), C.int(f), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSGNode) SetFlags(f QSGNode__Flag, enabled bool) {
	defer qt.Recovering("QSGNode::setFlags")

	if ptr.Pointer() != nil {
		C.QSGNode_SetFlags(ptr.Pointer(), C.int(f), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QSGNode) Type() QSGNode__NodeType {
	defer qt.Recovering("QSGNode::type")

	if ptr.Pointer() != nil {
		return QSGNode__NodeType(C.QSGNode_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGNode) DestroyQSGNode() {
	defer qt.Recovering("QSGNode::~QSGNode")

	if ptr.Pointer() != nil {
		C.QSGNode_DestroyQSGNode(ptr.Pointer())
	}
}

func (ptr *QSGNode) ObjectNameAbs() string {
	defer qt.Recovering("QSGNode::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGNode_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGNode) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSGNode::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSGNode_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QSGOpacityNode struct {
	QSGNode
}

type QSGOpacityNode_ITF interface {
	QSGNode_ITF
	QSGOpacityNode_PTR() *QSGOpacityNode
}

func PointerFromQSGOpacityNode(ptr QSGOpacityNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpacityNode_PTR().Pointer()
	}
	return nil
}

func NewQSGOpacityNodeFromPointer(ptr unsafe.Pointer) *QSGOpacityNode {
	var n = new(QSGOpacityNode)
	n.SetPointer(ptr)
	return n
}

func newQSGOpacityNodeFromPointer(ptr unsafe.Pointer) *QSGOpacityNode {
	var n = NewQSGOpacityNodeFromPointer(ptr)
	return n
}

func (ptr *QSGOpacityNode) QSGOpacityNode_PTR() *QSGOpacityNode {
	return ptr
}

func NewQSGOpacityNode() *QSGOpacityNode {
	defer qt.Recovering("QSGOpacityNode::QSGOpacityNode")

	return newQSGOpacityNodeFromPointer(C.QSGOpacityNode_NewQSGOpacityNode())
}

func (ptr *QSGOpacityNode) Opacity() float64 {
	defer qt.Recovering("QSGOpacityNode::opacity")

	if ptr.Pointer() != nil {
		return float64(C.QSGOpacityNode_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpacityNode) SetOpacity(opacity float64) {
	defer qt.Recovering("QSGOpacityNode::setOpacity")

	if ptr.Pointer() != nil {
		C.QSGOpacityNode_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QSGOpacityNode) DestroyQSGOpacityNode() {
	defer qt.Recovering("QSGOpacityNode::~QSGOpacityNode")

	if ptr.Pointer() != nil {
		C.QSGOpacityNode_DestroyQSGOpacityNode(ptr.Pointer())
	}
}

func (ptr *QSGOpacityNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGOpacityNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGOpacityNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGOpacityNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGOpacityNodePreprocess
func callbackQSGOpacityNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGOpacityNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGOpacityNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGOpacityNode) Preprocess() {
	defer qt.Recovering("QSGOpacityNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGOpacityNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGOpacityNode) PreprocessDefault() {
	defer qt.Recovering("QSGOpacityNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGOpacityNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGOpaqueTextureMaterial struct {
	QSGMaterial
}

type QSGOpaqueTextureMaterial_ITF interface {
	QSGMaterial_ITF
	QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial
}

func PointerFromQSGOpaqueTextureMaterial(ptr QSGOpaqueTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGOpaqueTextureMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGOpaqueTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGOpaqueTextureMaterial {
	var n = new(QSGOpaqueTextureMaterial)
	n.SetPointer(ptr)
	return n
}

func newQSGOpaqueTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGOpaqueTextureMaterial {
	var n = NewQSGOpaqueTextureMaterialFromPointer(ptr)
	return n
}

func (ptr *QSGOpaqueTextureMaterial) QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial {
	return ptr
}

func (ptr *QSGOpaqueTextureMaterial) Filtering() QSGTexture__Filtering {
	defer qt.Recovering("QSGOpaqueTextureMaterial::filtering")

	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) HorizontalWrapMode() QSGTexture__WrapMode {
	defer qt.Recovering("QSGOpaqueTextureMaterial::horizontalWrapMode")

	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_HorizontalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) MipmapFiltering() QSGTexture__Filtering {
	defer qt.Recovering("QSGOpaqueTextureMaterial::mipmapFiltering")

	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGOpaqueTextureMaterial_MipmapFiltering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) SetFiltering(filtering QSGTexture__Filtering) {
	defer qt.Recovering("QSGOpaqueTextureMaterial::setFiltering")

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetFiltering(ptr.Pointer(), C.int(filtering))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetHorizontalWrapMode(mode QSGTexture__WrapMode) {
	defer qt.Recovering("QSGOpaqueTextureMaterial::setHorizontalWrapMode")

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetHorizontalWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetMipmapFiltering(filtering QSGTexture__Filtering) {
	defer qt.Recovering("QSGOpaqueTextureMaterial::setMipmapFiltering")

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetMipmapFiltering(ptr.Pointer(), C.int(filtering))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetTexture(texture QSGTexture_ITF) {
	defer qt.Recovering("QSGOpaqueTextureMaterial::setTexture")

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetTexture(ptr.Pointer(), PointerFromQSGTexture(texture))
	}
}

func (ptr *QSGOpaqueTextureMaterial) SetVerticalWrapMode(mode QSGTexture__WrapMode) {
	defer qt.Recovering("QSGOpaqueTextureMaterial::setVerticalWrapMode")

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetVerticalWrapMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSGOpaqueTextureMaterial) Texture() *QSGTexture {
	defer qt.Recovering("QSGOpaqueTextureMaterial::texture")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGOpaqueTextureMaterial_Texture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) VerticalWrapMode() QSGTexture__WrapMode {
	defer qt.Recovering("QSGOpaqueTextureMaterial::verticalWrapMode")

	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGOpaqueTextureMaterial_VerticalWrapMode(ptr.Pointer()))
	}
	return 0
}

type QSGSimpleMaterial struct {
	QSGMaterial
}

type QSGSimpleMaterial_ITF interface {
	QSGMaterial_ITF
	QSGSimpleMaterial_PTR() *QSGSimpleMaterial
}

func PointerFromQSGSimpleMaterial(ptr QSGSimpleMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleMaterialFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterial {
	var n = new(QSGSimpleMaterial)
	n.SetPointer(ptr)
	return n
}

func newQSGSimpleMaterialFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterial {
	var n = NewQSGSimpleMaterialFromPointer(ptr)
	return n
}

func (ptr *QSGSimpleMaterial) QSGSimpleMaterial_PTR() *QSGSimpleMaterial {
	return ptr
}

type QSGSimpleMaterialShader struct {
	QSGMaterialShader
}

type QSGSimpleMaterialShader_ITF interface {
	QSGMaterialShader_ITF
	QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader
}

func PointerFromQSGSimpleMaterialShader(ptr QSGSimpleMaterialShader_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleMaterialShader_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterialShader {
	var n = new(QSGSimpleMaterialShader)
	n.SetPointer(ptr)
	return n
}

func newQSGSimpleMaterialShaderFromPointer(ptr unsafe.Pointer) *QSGSimpleMaterialShader {
	var n = NewQSGSimpleMaterialShaderFromPointer(ptr)
	return n
}

func (ptr *QSGSimpleMaterialShader) QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader {
	return ptr
}

type QSGSimpleRectNode struct {
	QSGGeometryNode
}

type QSGSimpleRectNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleRectNode_PTR() *QSGSimpleRectNode
}

func PointerFromQSGSimpleRectNode(ptr QSGSimpleRectNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleRectNode_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleRectNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleRectNode {
	var n = new(QSGSimpleRectNode)
	n.SetPointer(ptr)
	return n
}

func newQSGSimpleRectNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleRectNode {
	var n = NewQSGSimpleRectNodeFromPointer(ptr)
	return n
}

func (ptr *QSGSimpleRectNode) QSGSimpleRectNode_PTR() *QSGSimpleRectNode {
	return ptr
}

func NewQSGSimpleRectNode2() *QSGSimpleRectNode {
	defer qt.Recovering("QSGSimpleRectNode::QSGSimpleRectNode")

	return newQSGSimpleRectNodeFromPointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode2())
}

func NewQSGSimpleRectNode(rect core.QRectF_ITF, color gui.QColor_ITF) *QSGSimpleRectNode {
	defer qt.Recovering("QSGSimpleRectNode::QSGSimpleRectNode")

	return newQSGSimpleRectNodeFromPointer(C.QSGSimpleRectNode_NewQSGSimpleRectNode(core.PointerFromQRectF(rect), gui.PointerFromQColor(color)))
}

func (ptr *QSGSimpleRectNode) Color() *gui.QColor {
	defer qt.Recovering("QSGSimpleRectNode::color")

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QSGSimpleRectNode_Color(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGSimpleRectNode) Rect() *core.QRectF {
	defer qt.Recovering("QSGSimpleRectNode::rect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGSimpleRectNode_Rect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGSimpleRectNode) SetColor(color gui.QColor_ITF) {
	defer qt.Recovering("QSGSimpleRectNode::setColor")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetColor(ptr.Pointer(), gui.PointerFromQColor(color))
	}
}

func (ptr *QSGSimpleRectNode) SetRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QSGSimpleRectNode::setRect")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QSGSimpleRectNode) SetRect2(x float64, y float64, w float64, h float64) {
	defer qt.Recovering("QSGSimpleRectNode::setRect")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QSGSimpleRectNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGSimpleRectNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGSimpleRectNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGSimpleRectNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGSimpleRectNodePreprocess
func callbackQSGSimpleRectNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGSimpleRectNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGSimpleRectNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGSimpleRectNode) Preprocess() {
	defer qt.Recovering("QSGSimpleRectNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGSimpleRectNode) PreprocessDefault() {
	defer qt.Recovering("QSGSimpleRectNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGSimpleRectNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGSimpleTextureNode struct {
	QSGGeometryNode
}

type QSGSimpleTextureNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode
}

func PointerFromQSGSimpleTextureNode(ptr QSGSimpleTextureNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGSimpleTextureNode_PTR().Pointer()
	}
	return nil
}

func NewQSGSimpleTextureNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleTextureNode {
	var n = new(QSGSimpleTextureNode)
	n.SetPointer(ptr)
	return n
}

func newQSGSimpleTextureNodeFromPointer(ptr unsafe.Pointer) *QSGSimpleTextureNode {
	var n = NewQSGSimpleTextureNodeFromPointer(ptr)
	return n
}

func (ptr *QSGSimpleTextureNode) QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode {
	return ptr
}

//QSGSimpleTextureNode::TextureCoordinatesTransformFlag
type QSGSimpleTextureNode__TextureCoordinatesTransformFlag int64

const (
	QSGSimpleTextureNode__NoTransform        = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x00)
	QSGSimpleTextureNode__MirrorHorizontally = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x01)
	QSGSimpleTextureNode__MirrorVertically   = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x02)
)

func NewQSGSimpleTextureNode() *QSGSimpleTextureNode {
	defer qt.Recovering("QSGSimpleTextureNode::QSGSimpleTextureNode")

	return newQSGSimpleTextureNodeFromPointer(C.QSGSimpleTextureNode_NewQSGSimpleTextureNode())
}

func (ptr *QSGSimpleTextureNode) Filtering() QSGTexture__Filtering {
	defer qt.Recovering("QSGSimpleTextureNode::filtering")

	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGSimpleTextureNode_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGSimpleTextureNode) OwnsTexture() bool {
	defer qt.Recovering("QSGSimpleTextureNode::ownsTexture")

	if ptr.Pointer() != nil {
		return C.QSGSimpleTextureNode_OwnsTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGSimpleTextureNode) Rect() *core.QRectF {
	defer qt.Recovering("QSGSimpleTextureNode::rect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGSimpleTextureNode_Rect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) SetFiltering(filtering QSGTexture__Filtering) {
	defer qt.Recovering("QSGSimpleTextureNode::setFiltering")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetFiltering(ptr.Pointer(), C.int(filtering))
	}
}

func (ptr *QSGSimpleTextureNode) SetOwnsTexture(owns bool) {
	defer qt.Recovering("QSGSimpleTextureNode::setOwnsTexture")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetOwnsTexture(ptr.Pointer(), C.int(qt.GoBoolToInt(owns)))
	}
}

func (ptr *QSGSimpleTextureNode) SetRect(r core.QRectF_ITF) {
	defer qt.Recovering("QSGSimpleTextureNode::setRect")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetRect(ptr.Pointer(), core.PointerFromQRectF(r))
	}
}

func (ptr *QSGSimpleTextureNode) SetRect2(x float64, y float64, w float64, h float64) {
	defer qt.Recovering("QSGSimpleTextureNode::setRect")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QSGSimpleTextureNode) SetSourceRect(r core.QRectF_ITF) {
	defer qt.Recovering("QSGSimpleTextureNode::setSourceRect")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetSourceRect(ptr.Pointer(), core.PointerFromQRectF(r))
	}
}

func (ptr *QSGSimpleTextureNode) SetSourceRect2(x float64, y float64, w float64, h float64) {
	defer qt.Recovering("QSGSimpleTextureNode::setSourceRect")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetSourceRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QSGSimpleTextureNode) SetTexture(texture QSGTexture_ITF) {
	defer qt.Recovering("QSGSimpleTextureNode::setTexture")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetTexture(ptr.Pointer(), PointerFromQSGTexture(texture))
	}
}

func (ptr *QSGSimpleTextureNode) SetTextureCoordinatesTransform(mode QSGSimpleTextureNode__TextureCoordinatesTransformFlag) {
	defer qt.Recovering("QSGSimpleTextureNode::setTextureCoordinatesTransform")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_SetTextureCoordinatesTransform(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QSGSimpleTextureNode) SourceRect() *core.QRectF {
	defer qt.Recovering("QSGSimpleTextureNode::sourceRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGSimpleTextureNode_SourceRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) Texture() *QSGTexture {
	defer qt.Recovering("QSGSimpleTextureNode::texture")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGSimpleTextureNode_Texture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGSimpleTextureNode) TextureCoordinatesTransform() QSGSimpleTextureNode__TextureCoordinatesTransformFlag {
	defer qt.Recovering("QSGSimpleTextureNode::textureCoordinatesTransform")

	if ptr.Pointer() != nil {
		return QSGSimpleTextureNode__TextureCoordinatesTransformFlag(C.QSGSimpleTextureNode_TextureCoordinatesTransform(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGSimpleTextureNode) DestroyQSGSimpleTextureNode() {
	defer qt.Recovering("QSGSimpleTextureNode::~QSGSimpleTextureNode")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_DestroyQSGSimpleTextureNode(ptr.Pointer())
	}
}

func (ptr *QSGSimpleTextureNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGSimpleTextureNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGSimpleTextureNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGSimpleTextureNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGSimpleTextureNodePreprocess
func callbackQSGSimpleTextureNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGSimpleTextureNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGSimpleTextureNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGSimpleTextureNode) Preprocess() {
	defer qt.Recovering("QSGSimpleTextureNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGSimpleTextureNode) PreprocessDefault() {
	defer qt.Recovering("QSGSimpleTextureNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGSimpleTextureNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGTexture struct {
	core.QObject
}

type QSGTexture_ITF interface {
	core.QObject_ITF
	QSGTexture_PTR() *QSGTexture
}

func PointerFromQSGTexture(ptr QSGTexture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTexture_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureFromPointer(ptr unsafe.Pointer) *QSGTexture {
	var n = new(QSGTexture)
	n.SetPointer(ptr)
	return n
}

func newQSGTextureFromPointer(ptr unsafe.Pointer) *QSGTexture {
	var n = NewQSGTextureFromPointer(ptr)
	for len(n.ObjectName()) < len("QSGTexture_") {
		n.SetObjectName("QSGTexture_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGTexture) QSGTexture_PTR() *QSGTexture {
	return ptr
}

//QSGTexture::Filtering
type QSGTexture__Filtering int64

const (
	QSGTexture__None    = QSGTexture__Filtering(0)
	QSGTexture__Nearest = QSGTexture__Filtering(1)
	QSGTexture__Linear  = QSGTexture__Filtering(2)
)

//QSGTexture::WrapMode
type QSGTexture__WrapMode int64

const (
	QSGTexture__Repeat      = QSGTexture__WrapMode(0)
	QSGTexture__ClampToEdge = QSGTexture__WrapMode(1)
)

func (ptr *QSGTexture) Bind() {
	defer qt.Recovering("QSGTexture::bind")

	if ptr.Pointer() != nil {
		C.QSGTexture_Bind(ptr.Pointer())
	}
}

func (ptr *QSGTexture) ConvertToNormalizedSourceRect(rect core.QRectF_ITF) *core.QRectF {
	defer qt.Recovering("QSGTexture::convertToNormalizedSourceRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGTexture_ConvertToNormalizedSourceRect(ptr.Pointer(), core.PointerFromQRectF(rect)))
	}
	return nil
}

func (ptr *QSGTexture) Filtering() QSGTexture__Filtering {
	defer qt.Recovering("QSGTexture::filtering")

	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_Filtering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) HasAlphaChannel() bool {
	defer qt.Recovering("QSGTexture::hasAlphaChannel")

	if ptr.Pointer() != nil {
		return C.QSGTexture_HasAlphaChannel(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) HasMipmaps() bool {
	defer qt.Recovering("QSGTexture::hasMipmaps")

	if ptr.Pointer() != nil {
		return C.QSGTexture_HasMipmaps(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) HorizontalWrapMode() QSGTexture__WrapMode {
	defer qt.Recovering("QSGTexture::horizontalWrapMode")

	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_HorizontalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) IsAtlasTexture() bool {
	defer qt.Recovering("QSGTexture::isAtlasTexture")

	if ptr.Pointer() != nil {
		return C.QSGTexture_IsAtlasTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) MipmapFiltering() QSGTexture__Filtering {
	defer qt.Recovering("QSGTexture::mipmapFiltering")

	if ptr.Pointer() != nil {
		return QSGTexture__Filtering(C.QSGTexture_MipmapFiltering(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) NormalizedTextureSubRect() *core.QRectF {
	defer qt.Recovering("QSGTexture::normalizedTextureSubRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGTexture_NormalizedTextureSubRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTexture) RemovedFromAtlas() *QSGTexture {
	defer qt.Recovering("QSGTexture::removedFromAtlas")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGTexture_RemovedFromAtlas(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTexture) SetFiltering(filter QSGTexture__Filtering) {
	defer qt.Recovering("QSGTexture::setFiltering")

	if ptr.Pointer() != nil {
		C.QSGTexture_SetFiltering(ptr.Pointer(), C.int(filter))
	}
}

func (ptr *QSGTexture) SetHorizontalWrapMode(hwrap QSGTexture__WrapMode) {
	defer qt.Recovering("QSGTexture::setHorizontalWrapMode")

	if ptr.Pointer() != nil {
		C.QSGTexture_SetHorizontalWrapMode(ptr.Pointer(), C.int(hwrap))
	}
}

func (ptr *QSGTexture) SetMipmapFiltering(filter QSGTexture__Filtering) {
	defer qt.Recovering("QSGTexture::setMipmapFiltering")

	if ptr.Pointer() != nil {
		C.QSGTexture_SetMipmapFiltering(ptr.Pointer(), C.int(filter))
	}
}

func (ptr *QSGTexture) SetVerticalWrapMode(vwrap QSGTexture__WrapMode) {
	defer qt.Recovering("QSGTexture::setVerticalWrapMode")

	if ptr.Pointer() != nil {
		C.QSGTexture_SetVerticalWrapMode(ptr.Pointer(), C.int(vwrap))
	}
}

func (ptr *QSGTexture) TextureId() int {
	defer qt.Recovering("QSGTexture::textureId")

	if ptr.Pointer() != nil {
		return int(C.QSGTexture_TextureId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) TextureSize() *core.QSize {
	defer qt.Recovering("QSGTexture::textureSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSGTexture_TextureSize(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTexture) UpdateBindOptions(force bool) {
	defer qt.Recovering("QSGTexture::updateBindOptions")

	if ptr.Pointer() != nil {
		C.QSGTexture_UpdateBindOptions(ptr.Pointer(), C.int(qt.GoBoolToInt(force)))
	}
}

func (ptr *QSGTexture) VerticalWrapMode() QSGTexture__WrapMode {
	defer qt.Recovering("QSGTexture::verticalWrapMode")

	if ptr.Pointer() != nil {
		return QSGTexture__WrapMode(C.QSGTexture_VerticalWrapMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSGTexture) DestroyQSGTexture() {
	defer qt.Recovering("QSGTexture::~QSGTexture")

	if ptr.Pointer() != nil {
		C.QSGTexture_DestroyQSGTexture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGTexture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSGTexture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSGTexture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSGTexture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSGTextureTimerEvent
func callbackQSGTextureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTexture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGTexture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGTexture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGTexture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGTexture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGTexture::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGTexture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGTexture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSGTexture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSGTexture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSGTexture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSGTextureChildEvent
func callbackQSGTextureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTexture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGTexture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGTexture::childEvent")

	if ptr.Pointer() != nil {
		C.QSGTexture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGTexture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGTexture::childEvent")

	if ptr.Pointer() != nil {
		C.QSGTexture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGTexture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSGTexture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSGTexture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSGTexture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSGTextureCustomEvent
func callbackQSGTextureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTexture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGTexture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSGTexture::customEvent")

	if ptr.Pointer() != nil {
		C.QSGTexture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGTexture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSGTexture::customEvent")

	if ptr.Pointer() != nil {
		C.QSGTexture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSGTextureMaterial struct {
	QSGOpaqueTextureMaterial
}

type QSGTextureMaterial_ITF interface {
	QSGOpaqueTextureMaterial_ITF
	QSGTextureMaterial_PTR() *QSGTextureMaterial
}

func PointerFromQSGTextureMaterial(ptr QSGTextureMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGTextureMaterial {
	var n = new(QSGTextureMaterial)
	n.SetPointer(ptr)
	return n
}

func newQSGTextureMaterialFromPointer(ptr unsafe.Pointer) *QSGTextureMaterial {
	var n = NewQSGTextureMaterialFromPointer(ptr)
	return n
}

func (ptr *QSGTextureMaterial) QSGTextureMaterial_PTR() *QSGTextureMaterial {
	return ptr
}

type QSGTextureProvider struct {
	core.QObject
}

type QSGTextureProvider_ITF interface {
	core.QObject_ITF
	QSGTextureProvider_PTR() *QSGTextureProvider
}

func PointerFromQSGTextureProvider(ptr QSGTextureProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTextureProvider_PTR().Pointer()
	}
	return nil
}

func NewQSGTextureProviderFromPointer(ptr unsafe.Pointer) *QSGTextureProvider {
	var n = new(QSGTextureProvider)
	n.SetPointer(ptr)
	return n
}

func newQSGTextureProviderFromPointer(ptr unsafe.Pointer) *QSGTextureProvider {
	var n = NewQSGTextureProviderFromPointer(ptr)
	for len(n.ObjectName()) < len("QSGTextureProvider_") {
		n.SetObjectName("QSGTextureProvider_" + qt.Identifier())
	}
	return n
}

func (ptr *QSGTextureProvider) QSGTextureProvider_PTR() *QSGTextureProvider {
	return ptr
}

func (ptr *QSGTextureProvider) Texture() *QSGTexture {
	defer qt.Recovering("QSGTextureProvider::texture")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGTextureProvider_Texture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureProvider) ConnectTextureChanged(f func()) {
	defer qt.Recovering("connect QSGTextureProvider::textureChanged")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectTextureChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textureChanged", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTextureChanged() {
	defer qt.Recovering("disconnect QSGTextureProvider::textureChanged")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectTextureChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textureChanged")
	}
}

//export callbackQSGTextureProviderTextureChanged
func callbackQSGTextureProviderTextureChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGTextureProvider::textureChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGTextureProvider) TextureChanged() {
	defer qt.Recovering("QSGTextureProvider::textureChanged")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TextureChanged(ptr.Pointer())
	}
}

func (ptr *QSGTextureProvider) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QSGTextureProvider::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QSGTextureProvider::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQSGTextureProviderTimerEvent
func callbackQSGTextureProviderTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGTextureProvider) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::timerEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QSGTextureProvider) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QSGTextureProvider::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QSGTextureProvider::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQSGTextureProviderChildEvent
func callbackQSGTextureProviderChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::childEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGTextureProvider) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::childEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QSGTextureProvider) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QSGTextureProvider::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QSGTextureProvider::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQSGTextureProviderCustomEvent
func callbackQSGTextureProviderCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QSGTextureProvider) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::customEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QSGTextureProvider) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QSGTextureProvider::customEvent")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

type QSGTransformNode struct {
	QSGNode
}

type QSGTransformNode_ITF interface {
	QSGNode_ITF
	QSGTransformNode_PTR() *QSGTransformNode
}

func PointerFromQSGTransformNode(ptr QSGTransformNode_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGTransformNode_PTR().Pointer()
	}
	return nil
}

func NewQSGTransformNodeFromPointer(ptr unsafe.Pointer) *QSGTransformNode {
	var n = new(QSGTransformNode)
	n.SetPointer(ptr)
	return n
}

func newQSGTransformNodeFromPointer(ptr unsafe.Pointer) *QSGTransformNode {
	var n = NewQSGTransformNodeFromPointer(ptr)
	return n
}

func (ptr *QSGTransformNode) QSGTransformNode_PTR() *QSGTransformNode {
	return ptr
}

func NewQSGTransformNode() *QSGTransformNode {
	defer qt.Recovering("QSGTransformNode::QSGTransformNode")

	return newQSGTransformNodeFromPointer(C.QSGTransformNode_NewQSGTransformNode())
}

func (ptr *QSGTransformNode) SetMatrix(matrix gui.QMatrix4x4_ITF) {
	defer qt.Recovering("QSGTransformNode::setMatrix")

	if ptr.Pointer() != nil {
		C.QSGTransformNode_SetMatrix(ptr.Pointer(), gui.PointerFromQMatrix4x4(matrix))
	}
}

func (ptr *QSGTransformNode) DestroyQSGTransformNode() {
	defer qt.Recovering("QSGTransformNode::~QSGTransformNode")

	if ptr.Pointer() != nil {
		C.QSGTransformNode_DestroyQSGTransformNode(ptr.Pointer())
	}
}

func (ptr *QSGTransformNode) ConnectPreprocess(f func()) {
	defer qt.Recovering("connect QSGTransformNode::preprocess")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "preprocess", f)
	}
}

func (ptr *QSGTransformNode) DisconnectPreprocess() {
	defer qt.Recovering("disconnect QSGTransformNode::preprocess")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "preprocess")
	}
}

//export callbackQSGTransformNodePreprocess
func callbackQSGTransformNodePreprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGTransformNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTransformNodeFromPointer(ptr).PreprocessDefault()
	}
}

func (ptr *QSGTransformNode) Preprocess() {
	defer qt.Recovering("QSGTransformNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGTransformNode_Preprocess(ptr.Pointer())
	}
}

func (ptr *QSGTransformNode) PreprocessDefault() {
	defer qt.Recovering("QSGTransformNode::preprocess")

	if ptr.Pointer() != nil {
		C.QSGTransformNode_PreprocessDefault(ptr.Pointer())
	}
}

type QSGVertexColorMaterial struct {
	QSGMaterial
}

type QSGVertexColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial
}

func PointerFromQSGVertexColorMaterial(ptr QSGVertexColorMaterial_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSGVertexColorMaterial_PTR().Pointer()
	}
	return nil
}

func NewQSGVertexColorMaterialFromPointer(ptr unsafe.Pointer) *QSGVertexColorMaterial {
	var n = new(QSGVertexColorMaterial)
	n.SetPointer(ptr)
	return n
}

func newQSGVertexColorMaterialFromPointer(ptr unsafe.Pointer) *QSGVertexColorMaterial {
	var n = NewQSGVertexColorMaterialFromPointer(ptr)
	return n
}

func (ptr *QSGVertexColorMaterial) QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial {
	return ptr
}
