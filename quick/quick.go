// +build !minimal

package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
	"strings"
	"unsafe"
)

type QQuickAsyncImageProvider struct {
	QQuickImageProvider
}

type QQuickAsyncImageProvider_ITF interface {
	QQuickImageProvider_ITF
	QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider
}

func (p *QQuickAsyncImageProvider) QQuickAsyncImageProvider_PTR() *QQuickAsyncImageProvider {
	return p
}

func (p *QQuickAsyncImageProvider) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQuickImageProvider_PTR().Pointer()
	}
	return nil
}

func (p *QQuickAsyncImageProvider) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQuickImageProvider_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickAsyncImageProvider(ptr QQuickAsyncImageProvider_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickAsyncImageProvider_PTR().Pointer()
	}
	return nil
}

func NewQQuickAsyncImageProviderFromPointer(ptr unsafe.Pointer) *QQuickAsyncImageProvider {
	var n = new(QQuickAsyncImageProvider)
	n.SetPointer(ptr)
	return n
}

func newQQuickAsyncImageProviderFromPointer(ptr unsafe.Pointer) *QQuickAsyncImageProvider {
	var n = NewQQuickAsyncImageProviderFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QQuickAsyncImageProvider_") {
		n.SetObjectNameAbs("QQuickAsyncImageProvider_" + qt.Identifier())
	}
	return n
}

func NewQQuickAsyncImageProvider() *QQuickAsyncImageProvider {
	defer qt.Recovering("QQuickAsyncImageProvider::QQuickAsyncImageProvider")

	return newQQuickAsyncImageProviderFromPointer(C.QQuickAsyncImageProvider_NewQQuickAsyncImageProvider())
}

func (ptr *QQuickAsyncImageProvider) DestroyQQuickAsyncImageProvider() {
	defer qt.Recovering("QQuickAsyncImageProvider::~QQuickAsyncImageProvider")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QQuickAsyncImageProvider_DestroyQQuickAsyncImageProvider(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickAsyncImageProvider_RequestImageResponse
func callbackQQuickAsyncImageProvider_RequestImageResponse(ptr unsafe.Pointer, ptrName *C.char, id *C.char, requestedSize unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQuickAsyncImageProvider::requestImageResponse")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestImageResponse"); signal != nil {
		return PointerFromQQuickImageResponse(signal.(func(string, *core.QSize) *QQuickImageResponse)(C.GoString(id), core.NewQSizeFromPointer(requestedSize)))
	}

	return PointerFromQQuickImageResponse(nil)
}

func (ptr *QQuickAsyncImageProvider) ConnectRequestImageResponse(f func(id string, requestedSize *core.QSize) *QQuickImageResponse) {
	defer qt.Recovering("connect QQuickAsyncImageProvider::requestImageResponse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "requestImageResponse", f)
	}
}

func (ptr *QQuickAsyncImageProvider) DisconnectRequestImageResponse(id string, requestedSize core.QSize_ITF) {
	defer qt.Recovering("disconnect QQuickAsyncImageProvider::requestImageResponse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "requestImageResponse")
	}
}

func (ptr *QQuickAsyncImageProvider) RequestImageResponse(id string, requestedSize core.QSize_ITF) *QQuickImageResponse {
	defer qt.Recovering("QQuickAsyncImageProvider::requestImageResponse")

	if ptr.Pointer() != nil {
		return NewQQuickImageResponseFromPointer(C.QQuickAsyncImageProvider_RequestImageResponse(ptr.Pointer(), C.CString(id), core.PointerFromQSize(requestedSize)))
	}
	return nil
}

func (ptr *QQuickAsyncImageProvider) ObjectNameAbs() string {
	defer qt.Recovering("QQuickAsyncImageProvider::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickAsyncImageProvider_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickAsyncImageProvider) SetObjectNameAbs(name string) {
	defer qt.Recovering("QQuickAsyncImageProvider::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QQuickAsyncImageProvider_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}

type QQuickFramebufferObject struct {
	QQuickItem
}

type QQuickFramebufferObject_ITF interface {
	QQuickItem_ITF
	QQuickFramebufferObject_PTR() *QQuickFramebufferObject
}

func (p *QQuickFramebufferObject) QQuickFramebufferObject_PTR() *QQuickFramebufferObject {
	return p
}

func (p *QQuickFramebufferObject) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (p *QQuickFramebufferObject) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQuickItem_PTR().SetPointer(ptr)
	}
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

func (ptr *QQuickFramebufferObject) MirrorVertically() bool {
	defer qt.Recovering("QQuickFramebufferObject::mirrorVertically")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_MirrorVertically(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) SetMirrorVertically(enable bool) {
	defer qt.Recovering("QQuickFramebufferObject::setMirrorVertically")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_SetMirrorVertically(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
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

//export callbackQQuickFramebufferObject_IsTextureProvider
func callbackQQuickFramebufferObject_IsTextureProvider(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickFramebufferObject::isTextureProvider")

	if signal := qt.GetSignal(C.GoString(ptrName), "isTextureProvider"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQQuickFramebufferObjectFromPointer(ptr).IsTextureProviderDefault()))
}

func (ptr *QQuickFramebufferObject) ConnectIsTextureProvider(f func() bool) {
	defer qt.Recovering("connect QQuickFramebufferObject::isTextureProvider")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isTextureProvider", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectIsTextureProvider() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::isTextureProvider")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isTextureProvider")
	}
}

func (ptr *QQuickFramebufferObject) IsTextureProvider() bool {
	defer qt.Recovering("QQuickFramebufferObject::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) IsTextureProviderDefault() bool {
	defer qt.Recovering("QQuickFramebufferObject::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_IsTextureProviderDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickFramebufferObject_MirrorVerticallyChanged
func callbackQQuickFramebufferObject_MirrorVerticallyChanged(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QQuickFramebufferObject::mirrorVerticallyChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "mirrorVerticallyChanged"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	}

}

func (ptr *QQuickFramebufferObject) ConnectMirrorVerticallyChanged(f func(vbo bool)) {
	defer qt.Recovering("connect QQuickFramebufferObject::mirrorVerticallyChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectMirrorVerticallyChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "mirrorVerticallyChanged", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMirrorVerticallyChanged() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::mirrorVerticallyChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectMirrorVerticallyChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "mirrorVerticallyChanged")
	}
}

func (ptr *QQuickFramebufferObject) MirrorVerticallyChanged(vbo bool) {
	defer qt.Recovering("QQuickFramebufferObject::mirrorVerticallyChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_MirrorVerticallyChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQQuickFramebufferObject_ReleaseResources
func callbackQQuickFramebufferObject_ReleaseResources(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ReleaseResourcesDefault()
	}
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

//export callbackQQuickFramebufferObject_TextureFollowsItemSizeChanged
func callbackQQuickFramebufferObject_TextureFollowsItemSizeChanged(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QQuickFramebufferObject::textureFollowsItemSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureFollowsItemSizeChanged"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	}

}

func (ptr *QQuickFramebufferObject) ConnectTextureFollowsItemSizeChanged(f func(vbo bool)) {
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

func (ptr *QQuickFramebufferObject) TextureFollowsItemSizeChanged(vbo bool) {
	defer qt.Recovering("QQuickFramebufferObject::textureFollowsItemSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_TextureFollowsItemSizeChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQQuickFramebufferObject_TextureProvider
func callbackQQuickFramebufferObject_TextureProvider(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickFramebufferObject::textureProvider")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureProvider"); signal != nil {
		return PointerFromQSGTextureProvider(signal.(func() *QSGTextureProvider)())
	}

	return PointerFromQSGTextureProvider(NewQQuickFramebufferObjectFromPointer(ptr).TextureProviderDefault())
}

func (ptr *QQuickFramebufferObject) ConnectTextureProvider(f func() *QSGTextureProvider) {
	defer qt.Recovering("connect QQuickFramebufferObject::textureProvider")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureProvider", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectTextureProvider() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::textureProvider")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureProvider")
	}
}

func (ptr *QQuickFramebufferObject) TextureProvider() *QSGTextureProvider {
	defer qt.Recovering("QQuickFramebufferObject::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickFramebufferObject_TextureProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickFramebufferObject) TextureProviderDefault() *QSGTextureProvider {
	defer qt.Recovering("QQuickFramebufferObject::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickFramebufferObject_TextureProviderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickFramebufferObject_ChildMouseEventFilter
func callbackQQuickFramebufferObject_ChildMouseEventFilter(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickFramebufferObject::childMouseEventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "childMouseEventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QQuickItem, *core.QEvent) bool)(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickFramebufferObjectFromPointer(ptr).ChildMouseEventFilterDefault(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickFramebufferObject) ConnectChildMouseEventFilter(f func(item *QQuickItem, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickFramebufferObject::childMouseEventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childMouseEventFilter", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectChildMouseEventFilter() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::childMouseEventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childMouseEventFilter")
	}
}

func (ptr *QQuickFramebufferObject) ChildMouseEventFilter(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickFramebufferObject::childMouseEventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_ChildMouseEventFilter(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) ChildMouseEventFilterDefault(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickFramebufferObject::childMouseEventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_ChildMouseEventFilterDefault(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickFramebufferObject_Contains
func callbackQQuickFramebufferObject_Contains(ptr unsafe.Pointer, ptrName *C.char, point unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickFramebufferObject::contains")

	if signal := qt.GetSignal(C.GoString(ptrName), "contains"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickFramebufferObjectFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point))))
}

func (ptr *QQuickFramebufferObject) ConnectContains(f func(point *core.QPointF) bool) {
	defer qt.Recovering("connect QQuickFramebufferObject::contains")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contains", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectContains() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::contains")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contains")
	}
}

func (ptr *QQuickFramebufferObject) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QQuickFramebufferObject::contains")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) ContainsDefault(point core.QPointF_ITF) bool {
	defer qt.Recovering("QQuickFramebufferObject::contains")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

//export callbackQQuickFramebufferObject_DragEnterEvent
func callbackQQuickFramebufferObject_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_DragLeaveEvent
func callbackQQuickFramebufferObject_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_DragMoveEvent
func callbackQQuickFramebufferObject_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_DropEvent
func callbackQQuickFramebufferObject_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_FocusInEvent
func callbackQQuickFramebufferObject_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_FocusOutEvent
func callbackQQuickFramebufferObject_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_GeometryChanged
func callbackQQuickFramebufferObject_GeometryChanged(ptr unsafe.Pointer, ptrName *C.char, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::geometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
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

//export callbackQQuickFramebufferObject_HoverEnterEvent
func callbackQQuickFramebufferObject_HoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_HoverLeaveEvent
func callbackQQuickFramebufferObject_HoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_HoverMoveEvent
func callbackQQuickFramebufferObject_HoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_InputMethodEvent
func callbackQQuickFramebufferObject_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_InputMethodQuery
func callbackQQuickFramebufferObject_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QQuickFramebufferObject::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickFramebufferObjectFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickFramebufferObject) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QQuickFramebufferObject::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QQuickFramebufferObject) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickFramebufferObject::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickFramebufferObject_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QQuickFramebufferObject) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickFramebufferObject::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickFramebufferObject_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

//export callbackQQuickFramebufferObject_KeyPressEvent
func callbackQQuickFramebufferObject_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_KeyReleaseEvent
func callbackQQuickFramebufferObject_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_MouseDoubleClickEvent
func callbackQQuickFramebufferObject_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_MouseMoveEvent
func callbackQQuickFramebufferObject_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_MousePressEvent
func callbackQQuickFramebufferObject_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_MouseReleaseEvent
func callbackQQuickFramebufferObject_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_MouseUngrabEvent
func callbackQQuickFramebufferObject_MouseUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::mouseUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).MouseUngrabEventDefault()
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

//export callbackQQuickFramebufferObject_TouchEvent
func callbackQQuickFramebufferObject_TouchEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_TouchUngrabEvent
func callbackQQuickFramebufferObject_TouchUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::touchUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).TouchUngrabEventDefault()
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

//export callbackQQuickFramebufferObject_Update
func callbackQQuickFramebufferObject_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QQuickFramebufferObject::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectUpdate() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QQuickFramebufferObject) Update() {
	defer qt.Recovering("QQuickFramebufferObject::update")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_Update(ptr.Pointer())
	}
}

func (ptr *QQuickFramebufferObject) UpdateDefault() {
	defer qt.Recovering("QQuickFramebufferObject::update")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickFramebufferObject_UpdatePolish
func callbackQQuickFramebufferObject_UpdatePolish(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::updatePolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).UpdatePolishDefault()
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

//export callbackQQuickFramebufferObject_WheelEvent
func callbackQQuickFramebufferObject_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_TimerEvent
func callbackQQuickFramebufferObject_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_ChildEvent
func callbackQQuickFramebufferObject_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_ConnectNotify
func callbackQQuickFramebufferObject_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickFramebufferObject) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickFramebufferObject::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickFramebufferObject) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickFramebufferObject) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickFramebufferObject_CustomEvent
func callbackQQuickFramebufferObject_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickFramebufferObject_DeleteLater
func callbackQQuickFramebufferObject_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickFramebufferObject::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickFramebufferObject) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickFramebufferObject::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickFramebufferObject) DeleteLater() {
	defer qt.Recovering("QQuickFramebufferObject::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickFramebufferObject_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickFramebufferObject) DeleteLaterDefault() {
	defer qt.Recovering("QQuickFramebufferObject::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickFramebufferObject_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickFramebufferObject_DisconnectNotify
func callbackQQuickFramebufferObject_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickFramebufferObject::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickFramebufferObjectFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickFramebufferObject) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickFramebufferObject::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickFramebufferObject) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickFramebufferObject) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickFramebufferObject::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickFramebufferObject_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickFramebufferObject_EventFilter
func callbackQQuickFramebufferObject_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickFramebufferObject::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickFramebufferObjectFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickFramebufferObject) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickFramebufferObject::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickFramebufferObject) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickFramebufferObject::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickFramebufferObject) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickFramebufferObject::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickFramebufferObject_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickFramebufferObject_MetaObject
func callbackQQuickFramebufferObject_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickFramebufferObject::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickFramebufferObjectFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickFramebufferObject) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickFramebufferObject::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickFramebufferObject) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickFramebufferObject::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickFramebufferObject) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickFramebufferObject::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickFramebufferObject_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickFramebufferObject) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickFramebufferObject::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickFramebufferObject_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQuickImageProvider struct {
	qml.QQmlImageProviderBase
}

type QQuickImageProvider_ITF interface {
	qml.QQmlImageProviderBase_ITF
	QQuickImageProvider_PTR() *QQuickImageProvider
}

func (p *QQuickImageProvider) QQuickImageProvider_PTR() *QQuickImageProvider {
	return p
}

func (p *QQuickImageProvider) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQmlImageProviderBase_PTR().Pointer()
	}
	return nil
}

func (p *QQuickImageProvider) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQmlImageProviderBase_PTR().SetPointer(ptr)
	}
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

func (ptr *QQuickImageProvider) DestroyQQuickImageProvider() {
	defer qt.Recovering("QQuickImageProvider::~QQuickImageProvider")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QQuickImageProvider_DestroyQQuickImageProvider(ptr.Pointer())
		ptr.SetPointer(nil)
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

type QQuickImageResponse struct {
	core.QObject
}

type QQuickImageResponse_ITF interface {
	core.QObject_ITF
	QQuickImageResponse_PTR() *QQuickImageResponse
}

func (p *QQuickImageResponse) QQuickImageResponse_PTR() *QQuickImageResponse {
	return p
}

func (p *QQuickImageResponse) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickImageResponse) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQQuickImageResponse(ptr QQuickImageResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickImageResponse_PTR().Pointer()
	}
	return nil
}

func NewQQuickImageResponseFromPointer(ptr unsafe.Pointer) *QQuickImageResponse {
	var n = new(QQuickImageResponse)
	n.SetPointer(ptr)
	return n
}

func newQQuickImageResponseFromPointer(ptr unsafe.Pointer) *QQuickImageResponse {
	var n = NewQQuickImageResponseFromPointer(ptr)
	for len(n.ObjectName()) < len("QQuickImageResponse_") {
		n.SetObjectName("QQuickImageResponse_" + qt.Identifier())
	}
	return n
}

func NewQQuickImageResponse() *QQuickImageResponse {
	defer qt.Recovering("QQuickImageResponse::QQuickImageResponse")

	return newQQuickImageResponseFromPointer(C.QQuickImageResponse_NewQQuickImageResponse())
}

//export callbackQQuickImageResponse_Cancel
func callbackQQuickImageResponse_Cancel(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickImageResponse::cancel")

	if signal := qt.GetSignal(C.GoString(ptrName), "cancel"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageResponseFromPointer(ptr).CancelDefault()
	}
}

func (ptr *QQuickImageResponse) ConnectCancel(f func()) {
	defer qt.Recovering("connect QQuickImageResponse::cancel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "cancel", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectCancel() {
	defer qt.Recovering("disconnect QQuickImageResponse::cancel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "cancel")
	}
}

func (ptr *QQuickImageResponse) Cancel() {
	defer qt.Recovering("QQuickImageResponse::cancel")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_Cancel(ptr.Pointer())
	}
}

func (ptr *QQuickImageResponse) CancelDefault() {
	defer qt.Recovering("QQuickImageResponse::cancel")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_CancelDefault(ptr.Pointer())
	}
}

//export callbackQQuickImageResponse_ErrorString
func callbackQQuickImageResponse_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QQuickImageResponse::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQQuickImageResponseFromPointer(ptr).ErrorStringDefault())
}

func (ptr *QQuickImageResponse) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QQuickImageResponse::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "errorString", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectErrorString() {
	defer qt.Recovering("disconnect QQuickImageResponse::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "errorString")
	}
}

func (ptr *QQuickImageResponse) ErrorString() string {
	defer qt.Recovering("QQuickImageResponse::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickImageResponse_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QQuickImageResponse) ErrorStringDefault() string {
	defer qt.Recovering("QQuickImageResponse::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QQuickImageResponse_ErrorStringDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQQuickImageResponse_Finished
func callbackQQuickImageResponse_Finished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickImageResponse::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickImageResponse) ConnectFinished(f func()) {
	defer qt.Recovering("connect QQuickImageResponse::finished")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectFinished() {
	defer qt.Recovering("disconnect QQuickImageResponse::finished")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

func (ptr *QQuickImageResponse) Finished() {
	defer qt.Recovering("QQuickImageResponse::finished")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_Finished(ptr.Pointer())
	}
}

//export callbackQQuickImageResponse_TextureFactory
func callbackQQuickImageResponse_TextureFactory(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickImageResponse::textureFactory")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureFactory"); signal != nil {
		return PointerFromQQuickTextureFactory(signal.(func() *QQuickTextureFactory)())
	}

	return PointerFromQQuickTextureFactory(nil)
}

func (ptr *QQuickImageResponse) ConnectTextureFactory(f func() *QQuickTextureFactory) {
	defer qt.Recovering("connect QQuickImageResponse::textureFactory")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureFactory", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectTextureFactory() {
	defer qt.Recovering("disconnect QQuickImageResponse::textureFactory")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureFactory")
	}
}

func (ptr *QQuickImageResponse) TextureFactory() *QQuickTextureFactory {
	defer qt.Recovering("QQuickImageResponse::textureFactory")

	if ptr.Pointer() != nil {
		return NewQQuickTextureFactoryFromPointer(C.QQuickImageResponse_TextureFactory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickImageResponse) DestroyQQuickImageResponse() {
	defer qt.Recovering("QQuickImageResponse::~QQuickImageResponse")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickImageResponse_DestroyQQuickImageResponse(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickImageResponse_TimerEvent
func callbackQQuickImageResponse_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickImageResponse::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickImageResponseFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QQuickImageResponse) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QQuickImageResponse::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QQuickImageResponse::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QQuickImageResponse) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickImageResponse::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QQuickImageResponse) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QQuickImageResponse::timerEvent")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQQuickImageResponse_ChildEvent
func callbackQQuickImageResponse_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickImageResponse::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickImageResponseFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QQuickImageResponse) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QQuickImageResponse::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QQuickImageResponse::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QQuickImageResponse) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickImageResponse::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QQuickImageResponse) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QQuickImageResponse::childEvent")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQQuickImageResponse_ConnectNotify
func callbackQQuickImageResponse_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickImageResponse::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickImageResponseFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickImageResponse) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickImageResponse::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickImageResponse::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickImageResponse) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickImageResponse::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickImageResponse) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickImageResponse::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickImageResponse_CustomEvent
func callbackQQuickImageResponse_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickImageResponse::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickImageResponseFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QQuickImageResponse) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QQuickImageResponse::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QQuickImageResponse::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QQuickImageResponse) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickImageResponse::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QQuickImageResponse) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QQuickImageResponse::customEvent")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQQuickImageResponse_DeleteLater
func callbackQQuickImageResponse_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickImageResponse::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickImageResponseFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickImageResponse) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickImageResponse::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickImageResponse::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickImageResponse) DeleteLater() {
	defer qt.Recovering("QQuickImageResponse::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickImageResponse_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickImageResponse) DeleteLaterDefault() {
	defer qt.Recovering("QQuickImageResponse::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickImageResponse_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickImageResponse_DisconnectNotify
func callbackQQuickImageResponse_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickImageResponse::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickImageResponseFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickImageResponse) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickImageResponse::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickImageResponse::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickImageResponse) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickImageResponse::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickImageResponse) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickImageResponse::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickImageResponse_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickImageResponse_Event
func callbackQQuickImageResponse_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickImageResponse::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickImageResponseFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QQuickImageResponse) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickImageResponse::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectEvent() {
	defer qt.Recovering("disconnect QQuickImageResponse::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QQuickImageResponse) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickImageResponse::event")

	if ptr.Pointer() != nil {
		return C.QQuickImageResponse_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickImageResponse) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickImageResponse::event")

	if ptr.Pointer() != nil {
		return C.QQuickImageResponse_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickImageResponse_EventFilter
func callbackQQuickImageResponse_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickImageResponse::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickImageResponseFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickImageResponse) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickImageResponse::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickImageResponse::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickImageResponse) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickImageResponse::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickImageResponse_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickImageResponse) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickImageResponse::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickImageResponse_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickImageResponse_MetaObject
func callbackQQuickImageResponse_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickImageResponse::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickImageResponseFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickImageResponse) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickImageResponse::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickImageResponse) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickImageResponse::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickImageResponse) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickImageResponse::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickImageResponse_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickImageResponse) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickImageResponse::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickImageResponse_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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
	QQuickItem__ItemChildAddedChange           = QQuickItem__ItemChange(0)
	QQuickItem__ItemChildRemovedChange         = QQuickItem__ItemChange(1)
	QQuickItem__ItemSceneChange                = QQuickItem__ItemChange(2)
	QQuickItem__ItemVisibleHasChanged          = QQuickItem__ItemChange(3)
	QQuickItem__ItemParentHasChanged           = QQuickItem__ItemChange(4)
	QQuickItem__ItemOpacityHasChanged          = QQuickItem__ItemChange(5)
	QQuickItem__ItemActiveFocusHasChanged      = QQuickItem__ItemChange(6)
	QQuickItem__ItemRotationHasChanged         = QQuickItem__ItemChange(7)
	QQuickItem__ItemAntialiasingHasChanged     = QQuickItem__ItemChange(8)
	QQuickItem__ItemDevicePixelRatioHasChanged = QQuickItem__ItemChange(9)
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

type QQuickItem struct {
	core.QObject
	qml.QQmlParserStatus
}

type QQuickItem_ITF interface {
	core.QObject_ITF
	qml.QQmlParserStatus_ITF
	QQuickItem_PTR() *QQuickItem
}

func (p *QQuickItem) QQuickItem_PTR() *QQuickItem {
	return p
}

func (p *QQuickItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
		p.QQmlParserStatus_PTR().SetPointer(ptr)
	}
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

//export callbackQQuickItem_IsTextureProvider
func callbackQQuickItem_IsTextureProvider(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickItem::isTextureProvider")

	if signal := qt.GetSignal(C.GoString(ptrName), "isTextureProvider"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).IsTextureProviderDefault()))
}

func (ptr *QQuickItem) ConnectIsTextureProvider(f func() bool) {
	defer qt.Recovering("connect QQuickItem::isTextureProvider")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isTextureProvider", f)
	}
}

func (ptr *QQuickItem) DisconnectIsTextureProvider() {
	defer qt.Recovering("disconnect QQuickItem::isTextureProvider")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isTextureProvider")
	}
}

func (ptr *QQuickItem) IsTextureProvider() bool {
	defer qt.Recovering("QQuickItem::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) IsTextureProviderDefault() bool {
	defer qt.Recovering("QQuickItem::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsTextureProviderDefault(ptr.Pointer()) != 0
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

func (ptr *QQuickItem) SetActiveFocusOnTab(vbo bool) {
	defer qt.Recovering("QQuickItem::setActiveFocusOnTab")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetActiveFocusOnTab(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QQuickItem) SetAntialiasing(vbo bool) {
	defer qt.Recovering("QQuickItem::setAntialiasing")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetAntialiasing(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QQuickItem) SetBaselineOffset(vqr float64) {
	defer qt.Recovering("QQuickItem::setBaselineOffset")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetBaselineOffset(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetClip(vbo bool) {
	defer qt.Recovering("QQuickItem::setClip")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetClip(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QQuickItem) SetEnabled(vbo bool) {
	defer qt.Recovering("QQuickItem::setEnabled")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QQuickItem) SetFocus(vbo bool) {
	defer qt.Recovering("QQuickItem::setFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QQuickItem) SetFocus2(focus bool, reason core.Qt__FocusReason) {
	defer qt.Recovering("QQuickItem::setFocus")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetFocus2(ptr.Pointer(), C.int(qt.GoBoolToInt(focus)), C.int(reason))
	}
}

func (ptr *QQuickItem) SetHeight(vqr float64) {
	defer qt.Recovering("QQuickItem::setHeight")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetHeight(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetImplicitHeight(vqr float64) {
	defer qt.Recovering("QQuickItem::setImplicitHeight")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitHeight(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetImplicitWidth(vqr float64) {
	defer qt.Recovering("QQuickItem::setImplicitWidth")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetImplicitWidth(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetOpacity(vqr float64) {
	defer qt.Recovering("QQuickItem::setOpacity")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetOpacity(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetParentItem(parent QQuickItem_ITF) {
	defer qt.Recovering("QQuickItem::setParentItem")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetParentItem(ptr.Pointer(), PointerFromQQuickItem(parent))
	}
}

func (ptr *QQuickItem) SetRotation(vqr float64) {
	defer qt.Recovering("QQuickItem::setRotation")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetRotation(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetScale(vqr float64) {
	defer qt.Recovering("QQuickItem::setScale")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetScale(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetSmooth(vbo bool) {
	defer qt.Recovering("QQuickItem::setSmooth")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetSmooth(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QQuickItem) SetState(vqs string) {
	defer qt.Recovering("QQuickItem::setState")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetState(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QQuickItem) SetTransformOrigin(vtr QQuickItem__TransformOrigin) {
	defer qt.Recovering("QQuickItem::setTransformOrigin")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetTransformOrigin(ptr.Pointer(), C.int(vtr))
	}
}

func (ptr *QQuickItem) SetVisible(vbo bool) {
	defer qt.Recovering("QQuickItem::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QQuickItem) SetWidth(vqr float64) {
	defer qt.Recovering("QQuickItem::setWidth")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetWidth(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetX(vqr float64) {
	defer qt.Recovering("QQuickItem::setX")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetX(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetY(vqr float64) {
	defer qt.Recovering("QQuickItem::setY")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetY(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickItem) SetZ(vqr float64) {
	defer qt.Recovering("QQuickItem::setZ")

	if ptr.Pointer() != nil {
		C.QQuickItem_SetZ(ptr.Pointer(), C.double(vqr))
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

//export callbackQQuickItem_TextureProvider
func callbackQQuickItem_TextureProvider(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickItem::textureProvider")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureProvider"); signal != nil {
		return PointerFromQSGTextureProvider(signal.(func() *QSGTextureProvider)())
	}

	return PointerFromQSGTextureProvider(NewQQuickItemFromPointer(ptr).TextureProviderDefault())
}

func (ptr *QQuickItem) ConnectTextureProvider(f func() *QSGTextureProvider) {
	defer qt.Recovering("connect QQuickItem::textureProvider")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureProvider", f)
	}
}

func (ptr *QQuickItem) DisconnectTextureProvider() {
	defer qt.Recovering("disconnect QQuickItem::textureProvider")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureProvider")
	}
}

func (ptr *QQuickItem) TextureProvider() *QSGTextureProvider {
	defer qt.Recovering("QQuickItem::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickItem_TextureProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) TextureProviderDefault() *QSGTextureProvider {
	defer qt.Recovering("QQuickItem::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickItem_TextureProviderDefault(ptr.Pointer()))
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

//export callbackQQuickItem_ChildMouseEventFilter
func callbackQQuickItem_ChildMouseEventFilter(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickItem::childMouseEventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "childMouseEventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QQuickItem, *core.QEvent) bool)(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).ChildMouseEventFilterDefault(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickItem) ConnectChildMouseEventFilter(f func(item *QQuickItem, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickItem::childMouseEventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childMouseEventFilter", f)
	}
}

func (ptr *QQuickItem) DisconnectChildMouseEventFilter() {
	defer qt.Recovering("disconnect QQuickItem::childMouseEventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childMouseEventFilter")
	}
}

func (ptr *QQuickItem) ChildMouseEventFilter(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItem::childMouseEventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickItem_ChildMouseEventFilter(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickItem) ChildMouseEventFilterDefault(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItem::childMouseEventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickItem_ChildMouseEventFilterDefault(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickItem) ClassBegin() {
	defer qt.Recovering("QQuickItem::classBegin")

	if ptr.Pointer() != nil {
		C.QQuickItem_ClassBegin(ptr.Pointer())
	}
}

func (ptr *QQuickItem) ComponentComplete() {
	defer qt.Recovering("QQuickItem::componentComplete")

	if ptr.Pointer() != nil {
		C.QQuickItem_ComponentComplete(ptr.Pointer())
	}
}

//export callbackQQuickItem_Contains
func callbackQQuickItem_Contains(ptr unsafe.Pointer, ptrName *C.char, point unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickItem::contains")

	if signal := qt.GetSignal(C.GoString(ptrName), "contains"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point))))
}

func (ptr *QQuickItem) ConnectContains(f func(point *core.QPointF) bool) {
	defer qt.Recovering("connect QQuickItem::contains")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contains", f)
	}
}

func (ptr *QQuickItem) DisconnectContains() {
	defer qt.Recovering("disconnect QQuickItem::contains")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contains")
	}
}

func (ptr *QQuickItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QQuickItem::contains")

	if ptr.Pointer() != nil {
		return C.QQuickItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickItem) ContainsDefault(point core.QPointF_ITF) bool {
	defer qt.Recovering("QQuickItem::contains")

	if ptr.Pointer() != nil {
		return C.QQuickItem_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
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

//export callbackQQuickItem_DragEnterEvent
func callbackQQuickItem_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
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

//export callbackQQuickItem_DragLeaveEvent
func callbackQQuickItem_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
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

//export callbackQQuickItem_DragMoveEvent
func callbackQQuickItem_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
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

//export callbackQQuickItem_DropEvent
func callbackQQuickItem_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
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

//export callbackQQuickItem_FocusInEvent
func callbackQQuickItem_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
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

//export callbackQQuickItem_FocusOutEvent
func callbackQQuickItem_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
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

//export callbackQQuickItem_GeometryChanged
func callbackQQuickItem_GeometryChanged(ptr unsafe.Pointer, ptrName *C.char, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::geometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickItemFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
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

func (ptr *QQuickItem) HeightValid() bool {
	defer qt.Recovering("QQuickItem::heightValid")

	if ptr.Pointer() != nil {
		return C.QQuickItem_HeightValid(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickItem_HoverEnterEvent
func callbackQQuickItem_HoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
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

//export callbackQQuickItem_HoverLeaveEvent
func callbackQQuickItem_HoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
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

//export callbackQQuickItem_HoverMoveEvent
func callbackQQuickItem_HoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
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

//export callbackQQuickItem_InputMethodEvent
func callbackQQuickItem_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
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

//export callbackQQuickItem_InputMethodQuery
func callbackQQuickItem_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QQuickItem::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickItemFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickItem) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QQuickItem::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QQuickItem) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QQuickItem::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QQuickItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickItem_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QQuickItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickItem_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QQuickItem) IsComponentComplete() bool {
	defer qt.Recovering("QQuickItem::isComponentComplete")

	if ptr.Pointer() != nil {
		return C.QQuickItem_IsComponentComplete(ptr.Pointer()) != 0
	}
	return false
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

//export callbackQQuickItem_KeyPressEvent
func callbackQQuickItem_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
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

//export callbackQQuickItem_KeyReleaseEvent
func callbackQQuickItem_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
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

//export callbackQQuickItem_MouseDoubleClickEvent
func callbackQQuickItem_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
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

//export callbackQQuickItem_MouseMoveEvent
func callbackQQuickItem_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickItem_MousePressEvent
func callbackQQuickItem_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickItem_MouseReleaseEvent
func callbackQQuickItem_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickItem_MouseUngrabEvent
func callbackQQuickItem_MouseUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::mouseUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).MouseUngrabEventDefault()
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

//export callbackQQuickItem_ReleaseResources
func callbackQQuickItem_ReleaseResources(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).ReleaseResourcesDefault()
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

//export callbackQQuickItem_TouchEvent
func callbackQQuickItem_TouchEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
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

//export callbackQQuickItem_TouchUngrabEvent
func callbackQQuickItem_TouchUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::touchUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).TouchUngrabEventDefault()
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

//export callbackQQuickItem_Update
func callbackQQuickItem_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickItem) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QQuickItem::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QQuickItem) DisconnectUpdate() {
	defer qt.Recovering("disconnect QQuickItem::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QQuickItem) Update() {
	defer qt.Recovering("QQuickItem::update")

	if ptr.Pointer() != nil {
		C.QQuickItem_Update(ptr.Pointer())
	}
}

func (ptr *QQuickItem) UpdateInputMethod(queries core.Qt__InputMethodQuery) {
	defer qt.Recovering("QQuickItem::updateInputMethod")

	if ptr.Pointer() != nil {
		C.QQuickItem_UpdateInputMethod(ptr.Pointer(), C.int(queries))
	}
}

//export callbackQQuickItem_UpdatePolish
func callbackQQuickItem_UpdatePolish(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::updatePolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).UpdatePolishDefault()
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

//export callbackQQuickItem_WheelEvent
func callbackQQuickItem_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
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

func (ptr *QQuickItem) WidthValid() bool {
	defer qt.Recovering("QQuickItem::widthValid")

	if ptr.Pointer() != nil {
		return C.QQuickItem_WidthValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickItem) Window() *QQuickWindow {
	defer qt.Recovering("QQuickItem::window")

	if ptr.Pointer() != nil {
		return NewQQuickWindowFromPointer(C.QQuickItem_Window(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickItem_WindowChanged
func callbackQQuickItem_WindowChanged(ptr unsafe.Pointer, ptrName *C.char, window unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::windowChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "windowChanged"); signal != nil {
		signal.(func(*QQuickWindow))(NewQQuickWindowFromPointer(window))
	}

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

func (ptr *QQuickItem) WindowChanged(window QQuickWindow_ITF) {
	defer qt.Recovering("QQuickItem::windowChanged")

	if ptr.Pointer() != nil {
		C.QQuickItem_WindowChanged(ptr.Pointer(), PointerFromQQuickWindow(window))
	}
}

func (ptr *QQuickItem) DestroyQQuickItem() {
	defer qt.Recovering("QQuickItem::~QQuickItem")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickItem_DestroyQQuickItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickItem_TimerEvent
func callbackQQuickItem_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQQuickItem_ChildEvent
func callbackQQuickItem_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickItem_ConnectNotify
func callbackQQuickItem_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickItem) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItem_CustomEvent
func callbackQQuickItem_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickItem_DeleteLater
func callbackQQuickItem_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItem::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickItem) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickItem) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickItem) DeleteLater() {
	defer qt.Recovering("QQuickItem::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickItem_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickItem) DeleteLaterDefault() {
	defer qt.Recovering("QQuickItem::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickItem_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickItem_DisconnectNotify
func callbackQQuickItem_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItem::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickItem) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItem_EventFilter
func callbackQQuickItem_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickItem::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickItem) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickItem_MetaObject
func callbackQQuickItem_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickItem::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickItem) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickItem::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickItem) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickItem::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickItem) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItem) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQuickItemGrabResult struct {
	core.QObject
}

type QQuickItemGrabResult_ITF interface {
	core.QObject_ITF
	QQuickItemGrabResult_PTR() *QQuickItemGrabResult
}

func (p *QQuickItemGrabResult) QQuickItemGrabResult_PTR() *QQuickItemGrabResult {
	return p
}

func (p *QQuickItemGrabResult) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickItemGrabResult) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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

//export callbackQQuickItemGrabResult_Ready
func callbackQQuickItemGrabResult_Ready(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItemGrabResult::ready")

	if signal := qt.GetSignal(C.GoString(ptrName), "ready"); signal != nil {
		signal.(func())()
	}

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

//export callbackQQuickItemGrabResult_TimerEvent
func callbackQQuickItemGrabResult_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItemGrabResult::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
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

//export callbackQQuickItemGrabResult_ChildEvent
func callbackQQuickItemGrabResult_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItemGrabResult::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickItemGrabResult_ConnectNotify
func callbackQQuickItemGrabResult_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItemGrabResult::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItemGrabResult) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickItemGrabResult::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickItemGrabResult) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickItemGrabResult) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItemGrabResult_CustomEvent
func callbackQQuickItemGrabResult_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItemGrabResult::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickItemGrabResult_DeleteLater
func callbackQQuickItemGrabResult_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickItemGrabResult::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickItemGrabResult) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickItemGrabResult::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickItemGrabResult) DeleteLater() {
	defer qt.Recovering("QQuickItemGrabResult::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickItemGrabResult_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickItemGrabResult) DeleteLaterDefault() {
	defer qt.Recovering("QQuickItemGrabResult::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickItemGrabResult_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickItemGrabResult_DisconnectNotify
func callbackQQuickItemGrabResult_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickItemGrabResult::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickItemGrabResultFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickItemGrabResult) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickItemGrabResult::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickItemGrabResult) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickItemGrabResult) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickItemGrabResult::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickItemGrabResult_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickItemGrabResult_Event
func callbackQQuickItemGrabResult_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickItemGrabResult::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickItemGrabResultFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QQuickItemGrabResult) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickItemGrabResult::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectEvent() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QQuickItemGrabResult) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItemGrabResult::event")

	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickItemGrabResult) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItemGrabResult::event")

	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickItemGrabResult_EventFilter
func callbackQQuickItemGrabResult_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickItemGrabResult::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickItemGrabResultFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickItemGrabResult) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickItemGrabResult::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickItemGrabResult) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItemGrabResult::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickItemGrabResult) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickItemGrabResult::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickItemGrabResult_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickItemGrabResult_MetaObject
func callbackQQuickItemGrabResult_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickItemGrabResult::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickItemGrabResultFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickItemGrabResult) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickItemGrabResult::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickItemGrabResult) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickItemGrabResult::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickItemGrabResult) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickItemGrabResult::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickItemGrabResult_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickItemGrabResult) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickItemGrabResult::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickItemGrabResult_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

type QQuickPaintedItem struct {
	QQuickItem
}

type QQuickPaintedItem_ITF interface {
	QQuickItem_ITF
	QQuickPaintedItem_PTR() *QQuickPaintedItem
}

func (p *QQuickPaintedItem) QQuickPaintedItem_PTR() *QQuickPaintedItem {
	return p
}

func (p *QQuickPaintedItem) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQuickItem_PTR().Pointer()
	}
	return nil
}

func (p *QQuickPaintedItem) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQuickItem_PTR().SetPointer(ptr)
	}
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

func (ptr *QQuickPaintedItem) SetContentsScale(vqr float64) {
	defer qt.Recovering("QQuickPaintedItem::setContentsScale")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetContentsScale(ptr.Pointer(), C.double(vqr))
	}
}

func (ptr *QQuickPaintedItem) SetContentsSize(vqs core.QSize_ITF) {
	defer qt.Recovering("QQuickPaintedItem::setContentsSize")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetContentsSize(ptr.Pointer(), core.PointerFromQSize(vqs))
	}
}

func (ptr *QQuickPaintedItem) SetFillColor(vqc gui.QColor_ITF) {
	defer qt.Recovering("QQuickPaintedItem::setFillColor")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetFillColor(ptr.Pointer(), gui.PointerFromQColor(vqc))
	}
}

func (ptr *QQuickPaintedItem) SetRenderTarget(target QQuickPaintedItem__RenderTarget) {
	defer qt.Recovering("QQuickPaintedItem::setRenderTarget")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetRenderTarget(ptr.Pointer(), C.int(target))
	}
}

func (ptr *QQuickPaintedItem) SetTextureSize(size core.QSize_ITF) {
	defer qt.Recovering("QQuickPaintedItem::setTextureSize")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetTextureSize(ptr.Pointer(), core.PointerFromQSize(size))
	}
}

func (ptr *QQuickPaintedItem) TextureSize() *core.QSize {
	defer qt.Recovering("QQuickPaintedItem::textureSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickPaintedItem_TextureSize(ptr.Pointer()))
	}
	return nil
}

func NewQQuickPaintedItem(parent QQuickItem_ITF) *QQuickPaintedItem {
	defer qt.Recovering("QQuickPaintedItem::QQuickPaintedItem")

	return newQQuickPaintedItemFromPointer(C.QQuickPaintedItem_NewQQuickPaintedItem(PointerFromQQuickItem(parent)))
}

func (ptr *QQuickPaintedItem) Antialiasing() bool {
	defer qt.Recovering("QQuickPaintedItem::antialiasing")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Antialiasing(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_ContentsScaleChanged
func callbackQQuickPaintedItem_ContentsScaleChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::contentsScaleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsScaleChanged"); signal != nil {
		signal.(func())()
	}

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

func (ptr *QQuickPaintedItem) ContentsScaleChanged() {
	defer qt.Recovering("QQuickPaintedItem::contentsScaleChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ContentsScaleChanged(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_ContentsSizeChanged
func callbackQQuickPaintedItem_ContentsSizeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::contentsSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsSizeChanged"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickPaintedItem) ContentsSizeChanged() {
	defer qt.Recovering("QQuickPaintedItem::contentsSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ContentsSizeChanged(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_FillColorChanged
func callbackQQuickPaintedItem_FillColorChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::fillColorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fillColorChanged"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickPaintedItem) FillColorChanged() {
	defer qt.Recovering("QQuickPaintedItem::fillColorChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_FillColorChanged(ptr.Pointer())
	}
}

//export callbackQQuickPaintedItem_IsTextureProvider
func callbackQQuickPaintedItem_IsTextureProvider(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickPaintedItem::isTextureProvider")

	if signal := qt.GetSignal(C.GoString(ptrName), "isTextureProvider"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQQuickPaintedItemFromPointer(ptr).IsTextureProviderDefault()))
}

func (ptr *QQuickPaintedItem) ConnectIsTextureProvider(f func() bool) {
	defer qt.Recovering("connect QQuickPaintedItem::isTextureProvider")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isTextureProvider", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectIsTextureProvider() {
	defer qt.Recovering("disconnect QQuickPaintedItem::isTextureProvider")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isTextureProvider")
	}
}

func (ptr *QQuickPaintedItem) IsTextureProvider() bool {
	defer qt.Recovering("QQuickPaintedItem::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) IsTextureProviderDefault() bool {
	defer qt.Recovering("QQuickPaintedItem::isTextureProvider")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_IsTextureProviderDefault(ptr.Pointer()) != 0
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

//export callbackQQuickPaintedItem_Paint
func callbackQQuickPaintedItem_Paint(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::paint")

	if signal := qt.GetSignal(C.GoString(ptrName), "paint"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	}

}

func (ptr *QQuickPaintedItem) ConnectPaint(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QQuickPaintedItem::paint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paint", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectPaint(painter gui.QPainter_ITF) {
	defer qt.Recovering("disconnect QQuickPaintedItem::paint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paint")
	}
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

//export callbackQQuickPaintedItem_ReleaseResources
func callbackQQuickPaintedItem_ReleaseResources(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ReleaseResourcesDefault()
	}
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

//export callbackQQuickPaintedItem_RenderTargetChanged
func callbackQQuickPaintedItem_RenderTargetChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::renderTargetChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "renderTargetChanged"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickPaintedItem) RenderTargetChanged() {
	defer qt.Recovering("QQuickPaintedItem::renderTargetChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_RenderTargetChanged(ptr.Pointer())
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

//export callbackQQuickPaintedItem_TextureProvider
func callbackQQuickPaintedItem_TextureProvider(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickPaintedItem::textureProvider")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureProvider"); signal != nil {
		return PointerFromQSGTextureProvider(signal.(func() *QSGTextureProvider)())
	}

	return PointerFromQSGTextureProvider(NewQQuickPaintedItemFromPointer(ptr).TextureProviderDefault())
}

func (ptr *QQuickPaintedItem) ConnectTextureProvider(f func() *QSGTextureProvider) {
	defer qt.Recovering("connect QQuickPaintedItem::textureProvider")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureProvider", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTextureProvider() {
	defer qt.Recovering("disconnect QQuickPaintedItem::textureProvider")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureProvider")
	}
}

func (ptr *QQuickPaintedItem) TextureProvider() *QSGTextureProvider {
	defer qt.Recovering("QQuickPaintedItem::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickPaintedItem_TextureProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickPaintedItem) TextureProviderDefault() *QSGTextureProvider {
	defer qt.Recovering("QQuickPaintedItem::textureProvider")

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickPaintedItem_TextureProviderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickPaintedItem_TextureSizeChanged
func callbackQQuickPaintedItem_TextureSizeChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::textureSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureSizeChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickPaintedItem) ConnectTextureSizeChanged(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::textureSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectTextureSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "textureSizeChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectTextureSizeChanged() {
	defer qt.Recovering("disconnect QQuickPaintedItem::textureSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectTextureSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "textureSizeChanged")
	}
}

func (ptr *QQuickPaintedItem) TextureSizeChanged() {
	defer qt.Recovering("QQuickPaintedItem::textureSizeChanged")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_TextureSizeChanged(ptr.Pointer())
	}
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
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickPaintedItem_DestroyQQuickPaintedItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickPaintedItem_ChildMouseEventFilter
func callbackQQuickPaintedItem_ChildMouseEventFilter(ptr unsafe.Pointer, ptrName *C.char, item unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickPaintedItem::childMouseEventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "childMouseEventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QQuickItem, *core.QEvent) bool)(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickPaintedItemFromPointer(ptr).ChildMouseEventFilterDefault(NewQQuickItemFromPointer(item), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickPaintedItem) ConnectChildMouseEventFilter(f func(item *QQuickItem, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickPaintedItem::childMouseEventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childMouseEventFilter", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectChildMouseEventFilter() {
	defer qt.Recovering("disconnect QQuickPaintedItem::childMouseEventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childMouseEventFilter")
	}
}

func (ptr *QQuickPaintedItem) ChildMouseEventFilter(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickPaintedItem::childMouseEventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_ChildMouseEventFilter(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) ChildMouseEventFilterDefault(item QQuickItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickPaintedItem::childMouseEventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_ChildMouseEventFilterDefault(ptr.Pointer(), PointerFromQQuickItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_Contains
func callbackQQuickPaintedItem_Contains(ptr unsafe.Pointer, ptrName *C.char, point unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickPaintedItem::contains")

	if signal := qt.GetSignal(C.GoString(ptrName), "contains"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QPointF) bool)(core.NewQPointFFromPointer(point))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickPaintedItemFromPointer(ptr).ContainsDefault(core.NewQPointFFromPointer(point))))
}

func (ptr *QQuickPaintedItem) ConnectContains(f func(point *core.QPointF) bool) {
	defer qt.Recovering("connect QQuickPaintedItem::contains")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contains", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContains() {
	defer qt.Recovering("disconnect QQuickPaintedItem::contains")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contains")
	}
}

func (ptr *QQuickPaintedItem) Contains(point core.QPointF_ITF) bool {
	defer qt.Recovering("QQuickPaintedItem::contains")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) ContainsDefault(point core.QPointF_ITF) bool {
	defer qt.Recovering("QQuickPaintedItem::contains")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_ContainsDefault(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_DragEnterEvent
func callbackQQuickPaintedItem_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
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

//export callbackQQuickPaintedItem_DragLeaveEvent
func callbackQQuickPaintedItem_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
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

//export callbackQQuickPaintedItem_DragMoveEvent
func callbackQQuickPaintedItem_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
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

//export callbackQQuickPaintedItem_DropEvent
func callbackQQuickPaintedItem_DropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
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

//export callbackQQuickPaintedItem_FocusInEvent
func callbackQQuickPaintedItem_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
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

//export callbackQQuickPaintedItem_FocusOutEvent
func callbackQQuickPaintedItem_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
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

//export callbackQQuickPaintedItem_GeometryChanged
func callbackQQuickPaintedItem_GeometryChanged(ptr unsafe.Pointer, ptrName *C.char, newGeometry unsafe.Pointer, oldGeometry unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::geometryChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "geometryChanged"); signal != nil {
		signal.(func(*core.QRectF, *core.QRectF))(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).GeometryChangedDefault(core.NewQRectFFromPointer(newGeometry), core.NewQRectFFromPointer(oldGeometry))
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

//export callbackQQuickPaintedItem_HoverEnterEvent
func callbackQQuickPaintedItem_HoverEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).HoverEnterEventDefault(gui.NewQHoverEventFromPointer(event))
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

//export callbackQQuickPaintedItem_HoverLeaveEvent
func callbackQQuickPaintedItem_HoverLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).HoverLeaveEventDefault(gui.NewQHoverEventFromPointer(event))
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

//export callbackQQuickPaintedItem_HoverMoveEvent
func callbackQQuickPaintedItem_HoverMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).HoverMoveEventDefault(gui.NewQHoverEventFromPointer(event))
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

//export callbackQQuickPaintedItem_InputMethodEvent
func callbackQQuickPaintedItem_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
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

//export callbackQQuickPaintedItem_InputMethodQuery
func callbackQQuickPaintedItem_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QQuickPaintedItem::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickPaintedItemFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickPaintedItem) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QQuickPaintedItem::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QQuickPaintedItem::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QQuickPaintedItem) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickPaintedItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickPaintedItem_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QQuickPaintedItem) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickPaintedItem::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickPaintedItem_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

//export callbackQQuickPaintedItem_KeyPressEvent
func callbackQQuickPaintedItem_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
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

//export callbackQQuickPaintedItem_KeyReleaseEvent
func callbackQQuickPaintedItem_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
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

//export callbackQQuickPaintedItem_MouseDoubleClickEvent
func callbackQQuickPaintedItem_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickPaintedItem_MouseMoveEvent
func callbackQQuickPaintedItem_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickPaintedItem_MousePressEvent
func callbackQQuickPaintedItem_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickPaintedItem_MouseReleaseEvent
func callbackQQuickPaintedItem_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
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

//export callbackQQuickPaintedItem_MouseUngrabEvent
func callbackQQuickPaintedItem_MouseUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::mouseUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).MouseUngrabEventDefault()
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

//export callbackQQuickPaintedItem_TouchEvent
func callbackQQuickPaintedItem_TouchEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(event))
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

//export callbackQQuickPaintedItem_TouchUngrabEvent
func callbackQQuickPaintedItem_TouchUngrabEvent(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::touchUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchUngrabEvent"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).TouchUngrabEventDefault()
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

//export callbackQQuickPaintedItem_UpdatePolish
func callbackQQuickPaintedItem_UpdatePolish(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::updatePolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePolish"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).UpdatePolishDefault()
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

//export callbackQQuickPaintedItem_WheelEvent
func callbackQQuickPaintedItem_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
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

//export callbackQQuickPaintedItem_TimerEvent
func callbackQQuickPaintedItem_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQQuickPaintedItem_ChildEvent
func callbackQQuickPaintedItem_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickPaintedItem_ConnectNotify
func callbackQQuickPaintedItem_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickPaintedItem) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickPaintedItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickPaintedItem::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickPaintedItem) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickPaintedItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickPaintedItem) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickPaintedItem::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickPaintedItem_CustomEvent
func callbackQQuickPaintedItem_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickPaintedItem_DeleteLater
func callbackQQuickPaintedItem_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickPaintedItem) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickPaintedItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickPaintedItem::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickPaintedItem) DeleteLater() {
	defer qt.Recovering("QQuickPaintedItem::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickPaintedItem_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickPaintedItem) DeleteLaterDefault() {
	defer qt.Recovering("QQuickPaintedItem::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickPaintedItem_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickPaintedItem_DisconnectNotify
func callbackQQuickPaintedItem_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickPaintedItem::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickPaintedItemFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickPaintedItem) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickPaintedItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickPaintedItem::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickPaintedItem) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickPaintedItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickPaintedItem) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickPaintedItem::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickPaintedItem_EventFilter
func callbackQQuickPaintedItem_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickPaintedItem::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickPaintedItemFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickPaintedItem) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickPaintedItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickPaintedItem::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickPaintedItem) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickPaintedItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickPaintedItem::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickPaintedItem_MetaObject
func callbackQQuickPaintedItem_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickPaintedItem::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickPaintedItemFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickPaintedItem) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickPaintedItem::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickPaintedItem::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickPaintedItem) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickPaintedItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickPaintedItem_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickPaintedItem) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickPaintedItem::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickPaintedItem_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQuickRenderControl struct {
	core.QObject
}

type QQuickRenderControl_ITF interface {
	core.QObject_ITF
	QQuickRenderControl_PTR() *QQuickRenderControl
}

func (p *QQuickRenderControl) QQuickRenderControl_PTR() *QQuickRenderControl {
	return p
}

func (p *QQuickRenderControl) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickRenderControl) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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

//export callbackQQuickRenderControl_RenderRequested
func callbackQQuickRenderControl_RenderRequested(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickRenderControl::renderRequested")

	if signal := qt.GetSignal(C.GoString(ptrName), "renderRequested"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickRenderControl) RenderRequested() {
	defer qt.Recovering("QQuickRenderControl::renderRequested")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_RenderRequested(ptr.Pointer())
	}
}

//export callbackQQuickRenderControl_RenderWindow
func callbackQQuickRenderControl_RenderWindow(ptr unsafe.Pointer, ptrName *C.char, offset unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQuickRenderControl::renderWindow")

	if signal := qt.GetSignal(C.GoString(ptrName), "renderWindow"); signal != nil {
		return gui.PointerFromQWindow(signal.(func(*core.QPoint) *gui.QWindow)(core.NewQPointFromPointer(offset)))
	}

	return gui.PointerFromQWindow(NewQQuickRenderControlFromPointer(ptr).RenderWindowDefault(core.NewQPointFromPointer(offset)))
}

func (ptr *QQuickRenderControl) ConnectRenderWindow(f func(offset *core.QPoint) *gui.QWindow) {
	defer qt.Recovering("connect QQuickRenderControl::renderWindow")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "renderWindow", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectRenderWindow() {
	defer qt.Recovering("disconnect QQuickRenderControl::renderWindow")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "renderWindow")
	}
}

func (ptr *QQuickRenderControl) RenderWindow(offset core.QPoint_ITF) *gui.QWindow {
	defer qt.Recovering("QQuickRenderControl::renderWindow")

	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindow(ptr.Pointer(), core.PointerFromQPoint(offset)))
	}
	return nil
}

func (ptr *QQuickRenderControl) RenderWindowDefault(offset core.QPoint_ITF) *gui.QWindow {
	defer qt.Recovering("QQuickRenderControl::renderWindow")

	if ptr.Pointer() != nil {
		return gui.NewQWindowFromPointer(C.QQuickRenderControl_RenderWindowDefault(ptr.Pointer(), core.PointerFromQPoint(offset)))
	}
	return nil
}

func QQuickRenderControl_RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	defer qt.Recovering("QQuickRenderControl::renderWindowFor")

	return gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
}

func (ptr *QQuickRenderControl) RenderWindowFor(win QQuickWindow_ITF, offset core.QPoint_ITF) *gui.QWindow {
	defer qt.Recovering("QQuickRenderControl::renderWindowFor")

	return gui.NewQWindowFromPointer(C.QQuickRenderControl_QQuickRenderControl_RenderWindowFor(PointerFromQQuickWindow(win), core.PointerFromQPoint(offset)))
}

//export callbackQQuickRenderControl_SceneChanged
func callbackQQuickRenderControl_SceneChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickRenderControl::sceneChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneChanged"); signal != nil {
		signal.(func())()
	}

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
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickRenderControl_DestroyQQuickRenderControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickRenderControl_TimerEvent
func callbackQQuickRenderControl_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickRenderControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQQuickRenderControl_ChildEvent
func callbackQQuickRenderControl_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickRenderControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickRenderControl_ConnectNotify
func callbackQQuickRenderControl_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickRenderControl::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickRenderControlFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickRenderControl) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickRenderControl::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickRenderControl::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickRenderControl) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickRenderControl::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickRenderControl) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickRenderControl::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickRenderControl_CustomEvent
func callbackQQuickRenderControl_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickRenderControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickRenderControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickRenderControl_DeleteLater
func callbackQQuickRenderControl_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickRenderControl::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickRenderControlFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickRenderControl) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickRenderControl::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickRenderControl::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickRenderControl) DeleteLater() {
	defer qt.Recovering("QQuickRenderControl::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickRenderControl_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickRenderControl) DeleteLaterDefault() {
	defer qt.Recovering("QQuickRenderControl::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickRenderControl_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickRenderControl_DisconnectNotify
func callbackQQuickRenderControl_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickRenderControl::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickRenderControlFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickRenderControl) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickRenderControl::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickRenderControl::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickRenderControl) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickRenderControl::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickRenderControl) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickRenderControl::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickRenderControl_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickRenderControl_Event
func callbackQQuickRenderControl_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickRenderControl::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickRenderControlFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QQuickRenderControl) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickRenderControl::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectEvent() {
	defer qt.Recovering("disconnect QQuickRenderControl::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QQuickRenderControl) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickRenderControl::event")

	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickRenderControl::event")

	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickRenderControl_EventFilter
func callbackQQuickRenderControl_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickRenderControl::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickRenderControlFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickRenderControl) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickRenderControl::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickRenderControl::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickRenderControl) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickRenderControl::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickRenderControl) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickRenderControl::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickRenderControl_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickRenderControl_MetaObject
func callbackQQuickRenderControl_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickRenderControl::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickRenderControlFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickRenderControl) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickRenderControl::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickRenderControl) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickRenderControl::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickRenderControl) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickRenderControl::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickRenderControl_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickRenderControl) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickRenderControl::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickRenderControl_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQuickTextDocument struct {
	core.QObject
}

type QQuickTextDocument_ITF interface {
	core.QObject_ITF
	QQuickTextDocument_PTR() *QQuickTextDocument
}

func (p *QQuickTextDocument) QQuickTextDocument_PTR() *QQuickTextDocument {
	return p
}

func (p *QQuickTextDocument) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickTextDocument) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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

//export callbackQQuickTextDocument_TimerEvent
func callbackQQuickTextDocument_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
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

//export callbackQQuickTextDocument_ChildEvent
func callbackQQuickTextDocument_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickTextDocument_ConnectNotify
func callbackQQuickTextDocument_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextDocument) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickTextDocument::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickTextDocument::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickTextDocument) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickTextDocument::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickTextDocument) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickTextDocument::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextDocument_CustomEvent
func callbackQQuickTextDocument_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickTextDocument_DeleteLater
func callbackQQuickTextDocument_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickTextDocument::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickTextDocumentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickTextDocument) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickTextDocument::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickTextDocument::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickTextDocument) DeleteLater() {
	defer qt.Recovering("QQuickTextDocument::deleteLater")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickTextDocument) DeleteLaterDefault() {
	defer qt.Recovering("QQuickTextDocument::deleteLater")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickTextDocument_DisconnectNotify
func callbackQQuickTextDocument_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextDocument::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextDocumentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextDocument) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickTextDocument::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickTextDocument::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickTextDocument) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickTextDocument::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickTextDocument) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickTextDocument::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickTextDocument_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextDocument_Event
func callbackQQuickTextDocument_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickTextDocument::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickTextDocumentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QQuickTextDocument) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickTextDocument::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectEvent() {
	defer qt.Recovering("disconnect QQuickTextDocument::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QQuickTextDocument) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickTextDocument::event")

	if ptr.Pointer() != nil {
		return C.QQuickTextDocument_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickTextDocument) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickTextDocument::event")

	if ptr.Pointer() != nil {
		return C.QQuickTextDocument_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickTextDocument_EventFilter
func callbackQQuickTextDocument_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickTextDocument::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickTextDocumentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickTextDocument) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickTextDocument::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickTextDocument::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickTextDocument) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickTextDocument::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickTextDocument_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickTextDocument) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickTextDocument::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickTextDocument_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickTextDocument_MetaObject
func callbackQQuickTextDocument_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickTextDocument::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickTextDocumentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickTextDocument) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickTextDocument::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickTextDocument) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickTextDocument::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickTextDocument) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickTextDocument::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextDocument_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextDocument) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickTextDocument::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextDocument_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QQuickTextureFactory struct {
	core.QObject
}

type QQuickTextureFactory_ITF interface {
	core.QObject_ITF
	QQuickTextureFactory_PTR() *QQuickTextureFactory
}

func (p *QQuickTextureFactory) QQuickTextureFactory_PTR() *QQuickTextureFactory {
	return p
}

func (p *QQuickTextureFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QQuickTextureFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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

//export callbackQQuickTextureFactory_Image
func callbackQQuickTextureFactory_Image(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickTextureFactory::image")

	if signal := qt.GetSignal(C.GoString(ptrName), "image"); signal != nil {
		return gui.PointerFromQImage(signal.(func() *gui.QImage)())
	}

	return gui.PointerFromQImage(NewQQuickTextureFactoryFromPointer(ptr).ImageDefault())
}

func (ptr *QQuickTextureFactory) ConnectImage(f func() *gui.QImage) {
	defer qt.Recovering("connect QQuickTextureFactory::image")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "image", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectImage() {
	defer qt.Recovering("disconnect QQuickTextureFactory::image")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "image")
	}
}

func (ptr *QQuickTextureFactory) Image() *gui.QImage {
	defer qt.Recovering("QQuickTextureFactory::image")

	if ptr.Pointer() != nil {
		return gui.NewQImageFromPointer(C.QQuickTextureFactory_Image(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextureFactory) ImageDefault() *gui.QImage {
	defer qt.Recovering("QQuickTextureFactory::image")

	if ptr.Pointer() != nil {
		return gui.NewQImageFromPointer(C.QQuickTextureFactory_ImageDefault(ptr.Pointer()))
	}
	return nil
}

func NewQQuickTextureFactory() *QQuickTextureFactory {
	defer qt.Recovering("QQuickTextureFactory::QQuickTextureFactory")

	return newQQuickTextureFactoryFromPointer(C.QQuickTextureFactory_NewQQuickTextureFactory())
}

//export callbackQQuickTextureFactory_CreateTexture
func callbackQQuickTextureFactory_CreateTexture(ptr unsafe.Pointer, ptrName *C.char, window unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QQuickTextureFactory::createTexture")

	if signal := qt.GetSignal(C.GoString(ptrName), "createTexture"); signal != nil {
		return PointerFromQSGTexture(signal.(func(*QQuickWindow) *QSGTexture)(NewQQuickWindowFromPointer(window)))
	}

	return PointerFromQSGTexture(nil)
}

func (ptr *QQuickTextureFactory) ConnectCreateTexture(f func(window *QQuickWindow) *QSGTexture) {
	defer qt.Recovering("connect QQuickTextureFactory::createTexture")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "createTexture", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectCreateTexture(window QQuickWindow_ITF) {
	defer qt.Recovering("disconnect QQuickTextureFactory::createTexture")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "createTexture")
	}
}

func (ptr *QQuickTextureFactory) CreateTexture(window QQuickWindow_ITF) *QSGTexture {
	defer qt.Recovering("QQuickTextureFactory::createTexture")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QQuickTextureFactory_CreateTexture(ptr.Pointer(), PointerFromQQuickWindow(window)))
	}
	return nil
}

//export callbackQQuickTextureFactory_TextureByteCount
func callbackQQuickTextureFactory_TextureByteCount(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickTextureFactory::textureByteCount")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureByteCount"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(0)
}

func (ptr *QQuickTextureFactory) ConnectTextureByteCount(f func() int) {
	defer qt.Recovering("connect QQuickTextureFactory::textureByteCount")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureByteCount", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectTextureByteCount() {
	defer qt.Recovering("disconnect QQuickTextureFactory::textureByteCount")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureByteCount")
	}
}

func (ptr *QQuickTextureFactory) TextureByteCount() int {
	defer qt.Recovering("QQuickTextureFactory::textureByteCount")

	if ptr.Pointer() != nil {
		return int(C.QQuickTextureFactory_TextureByteCount(ptr.Pointer()))
	}
	return 0
}

func QQuickTextureFactory_TextureFactoryForImage(image gui.QImage_ITF) *QQuickTextureFactory {
	defer qt.Recovering("QQuickTextureFactory::textureFactoryForImage")

	return NewQQuickTextureFactoryFromPointer(C.QQuickTextureFactory_QQuickTextureFactory_TextureFactoryForImage(gui.PointerFromQImage(image)))
}

func (ptr *QQuickTextureFactory) TextureFactoryForImage(image gui.QImage_ITF) *QQuickTextureFactory {
	defer qt.Recovering("QQuickTextureFactory::textureFactoryForImage")

	return NewQQuickTextureFactoryFromPointer(C.QQuickTextureFactory_QQuickTextureFactory_TextureFactoryForImage(gui.PointerFromQImage(image)))
}

//export callbackQQuickTextureFactory_TextureSize
func callbackQQuickTextureFactory_TextureSize(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickTextureFactory::textureSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureSize"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(nil)
}

func (ptr *QQuickTextureFactory) ConnectTextureSize(f func() *core.QSize) {
	defer qt.Recovering("connect QQuickTextureFactory::textureSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureSize", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectTextureSize() {
	defer qt.Recovering("disconnect QQuickTextureFactory::textureSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureSize")
	}
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
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickTextureFactory_DestroyQQuickTextureFactory(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickTextureFactory_TimerEvent
func callbackQQuickTextureFactory_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQQuickTextureFactory_ChildEvent
func callbackQQuickTextureFactory_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickTextureFactory_ConnectNotify
func callbackQQuickTextureFactory_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextureFactory) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickTextureFactory::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickTextureFactory::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickTextureFactory) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickTextureFactory::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickTextureFactory) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickTextureFactory::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextureFactory_CustomEvent
func callbackQQuickTextureFactory_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickTextureFactory_DeleteLater
func callbackQQuickTextureFactory_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickTextureFactory::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickTextureFactory) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickTextureFactory::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickTextureFactory::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickTextureFactory) DeleteLater() {
	defer qt.Recovering("QQuickTextureFactory::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickTextureFactory_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickTextureFactory) DeleteLaterDefault() {
	defer qt.Recovering("QQuickTextureFactory::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickTextureFactory_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickTextureFactory_DisconnectNotify
func callbackQQuickTextureFactory_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickTextureFactory::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickTextureFactoryFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickTextureFactory) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickTextureFactory::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickTextureFactory::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickTextureFactory) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickTextureFactory::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickTextureFactory) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickTextureFactory::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickTextureFactory_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickTextureFactory_Event
func callbackQQuickTextureFactory_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickTextureFactory::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickTextureFactoryFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QQuickTextureFactory) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickTextureFactory::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectEvent() {
	defer qt.Recovering("disconnect QQuickTextureFactory::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QQuickTextureFactory) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickTextureFactory::event")

	if ptr.Pointer() != nil {
		return C.QQuickTextureFactory_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QQuickTextureFactory) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickTextureFactory::event")

	if ptr.Pointer() != nil {
		return C.QQuickTextureFactory_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQQuickTextureFactory_EventFilter
func callbackQQuickTextureFactory_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickTextureFactory::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickTextureFactoryFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickTextureFactory) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickTextureFactory::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickTextureFactory::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickTextureFactory) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickTextureFactory::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickTextureFactory_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickTextureFactory) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickTextureFactory::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickTextureFactory_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickTextureFactory_MetaObject
func callbackQQuickTextureFactory_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickTextureFactory::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickTextureFactoryFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickTextureFactory) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickTextureFactory::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickTextureFactory) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickTextureFactory::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickTextureFactory) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickTextureFactory::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextureFactory_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickTextureFactory) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickTextureFactory::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickTextureFactory_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

type QQuickView struct {
	QQuickWindow
}

type QQuickView_ITF interface {
	QQuickWindow_ITF
	QQuickView_PTR() *QQuickView
}

func (p *QQuickView) QQuickView_PTR() *QQuickView {
	return p
}

func (p *QQuickView) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QQuickWindow_PTR().Pointer()
	}
	return nil
}

func (p *QQuickView) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QQuickWindow_PTR().SetPointer(ptr)
	}
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

func (ptr *QQuickView) ResizeMode() QQuickView__ResizeMode {
	defer qt.Recovering("QQuickView::resizeMode")

	if ptr.Pointer() != nil {
		return QQuickView__ResizeMode(C.QQuickView_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickView) SetResizeMode(vre QQuickView__ResizeMode) {
	defer qt.Recovering("QQuickView::setResizeMode")

	if ptr.Pointer() != nil {
		C.QQuickView_SetResizeMode(ptr.Pointer(), C.int(vre))
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

func (ptr *QQuickView) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickView) MouseMoveEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MousePressEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
	}
}

func (ptr *QQuickView) MouseReleaseEvent(e gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(e))
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

//export callbackQQuickView_SetSource
func callbackQQuickView_SetSource(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::setSource")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSource"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQuickView) ConnectSetSource(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QQuickView::setSource")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSource", f)
	}
}

func (ptr *QQuickView) DisconnectSetSource(url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QQuickView::setSource")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSource")
	}
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

//export callbackQQuickView_StatusChanged
func callbackQQuickView_StatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQuickView::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQuickView__Status))(QQuickView__Status(status))
	}

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

func (ptr *QQuickView) StatusChanged(status QQuickView__Status) {
	defer qt.Recovering("QQuickView::statusChanged")

	if ptr.Pointer() != nil {
		C.QQuickView_StatusChanged(ptr.Pointer(), C.int(status))
	}
}

func (ptr *QQuickView) DestroyQQuickView() {
	defer qt.Recovering("QQuickView::~QQuickView")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickView_DestroyQQuickView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickView_ReleaseResources
func callbackQQuickView_ReleaseResources(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ReleaseResourcesDefault()
	}
}

func (ptr *QQuickView) ConnectReleaseResources(f func()) {
	defer qt.Recovering("connect QQuickView::releaseResources")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "releaseResources", f)
	}
}

func (ptr *QQuickView) DisconnectReleaseResources() {
	defer qt.Recovering("disconnect QQuickView::releaseResources")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "releaseResources")
	}
}

func (ptr *QQuickView) ReleaseResources() {
	defer qt.Recovering("QQuickView::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickView_ReleaseResources(ptr.Pointer())
	}
}

func (ptr *QQuickView) ReleaseResourcesDefault() {
	defer qt.Recovering("QQuickView::releaseResources")

	if ptr.Pointer() != nil {
		C.QQuickView_ReleaseResourcesDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_Update
func callbackQQuickView_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QQuickView) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QQuickView::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QQuickView) DisconnectUpdate() {
	defer qt.Recovering("disconnect QQuickView::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QQuickView) Update() {
	defer qt.Recovering("QQuickView::update")

	if ptr.Pointer() != nil {
		C.QQuickView_Update(ptr.Pointer())
	}
}

func (ptr *QQuickView) UpdateDefault() {
	defer qt.Recovering("QQuickView::update")

	if ptr.Pointer() != nil {
		C.QQuickView_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_SetHeight
func callbackQQuickView_SetHeight(ptr unsafe.Pointer, ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QQuickView::setHeight")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHeight"); signal != nil {
		signal.(func(int))(int(arg))
	} else {
		NewQQuickViewFromPointer(ptr).SetHeightDefault(int(arg))
	}
}

func (ptr *QQuickView) ConnectSetHeight(f func(arg int)) {
	defer qt.Recovering("connect QQuickView::setHeight")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHeight", f)
	}
}

func (ptr *QQuickView) DisconnectSetHeight() {
	defer qt.Recovering("disconnect QQuickView::setHeight")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHeight")
	}
}

func (ptr *QQuickView) SetHeight(arg int) {
	defer qt.Recovering("QQuickView::setHeight")

	if ptr.Pointer() != nil {
		C.QQuickView_SetHeight(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QQuickView) SetHeightDefault(arg int) {
	defer qt.Recovering("QQuickView::setHeight")

	if ptr.Pointer() != nil {
		C.QQuickView_SetHeightDefault(ptr.Pointer(), C.int(arg))
	}
}

//export callbackQQuickView_SetMaximumHeight
func callbackQQuickView_SetMaximumHeight(ptr unsafe.Pointer, ptrName *C.char, h C.int) {
	defer qt.Recovering("callback QQuickView::setMaximumHeight")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMaximumHeight"); signal != nil {
		signal.(func(int))(int(h))
	} else {
		NewQQuickViewFromPointer(ptr).SetMaximumHeightDefault(int(h))
	}
}

func (ptr *QQuickView) ConnectSetMaximumHeight(f func(h int)) {
	defer qt.Recovering("connect QQuickView::setMaximumHeight")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMaximumHeight", f)
	}
}

func (ptr *QQuickView) DisconnectSetMaximumHeight() {
	defer qt.Recovering("disconnect QQuickView::setMaximumHeight")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMaximumHeight")
	}
}

func (ptr *QQuickView) SetMaximumHeight(h int) {
	defer qt.Recovering("QQuickView::setMaximumHeight")

	if ptr.Pointer() != nil {
		C.QQuickView_SetMaximumHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QQuickView) SetMaximumHeightDefault(h int) {
	defer qt.Recovering("QQuickView::setMaximumHeight")

	if ptr.Pointer() != nil {
		C.QQuickView_SetMaximumHeightDefault(ptr.Pointer(), C.int(h))
	}
}

//export callbackQQuickView_SetMaximumWidth
func callbackQQuickView_SetMaximumWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) {
	defer qt.Recovering("callback QQuickView::setMaximumWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMaximumWidth"); signal != nil {
		signal.(func(int))(int(w))
	} else {
		NewQQuickViewFromPointer(ptr).SetMaximumWidthDefault(int(w))
	}
}

func (ptr *QQuickView) ConnectSetMaximumWidth(f func(w int)) {
	defer qt.Recovering("connect QQuickView::setMaximumWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMaximumWidth", f)
	}
}

func (ptr *QQuickView) DisconnectSetMaximumWidth() {
	defer qt.Recovering("disconnect QQuickView::setMaximumWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMaximumWidth")
	}
}

func (ptr *QQuickView) SetMaximumWidth(w int) {
	defer qt.Recovering("QQuickView::setMaximumWidth")

	if ptr.Pointer() != nil {
		C.QQuickView_SetMaximumWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QQuickView) SetMaximumWidthDefault(w int) {
	defer qt.Recovering("QQuickView::setMaximumWidth")

	if ptr.Pointer() != nil {
		C.QQuickView_SetMaximumWidthDefault(ptr.Pointer(), C.int(w))
	}
}

//export callbackQQuickView_SetMinimumHeight
func callbackQQuickView_SetMinimumHeight(ptr unsafe.Pointer, ptrName *C.char, h C.int) {
	defer qt.Recovering("callback QQuickView::setMinimumHeight")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMinimumHeight"); signal != nil {
		signal.(func(int))(int(h))
	} else {
		NewQQuickViewFromPointer(ptr).SetMinimumHeightDefault(int(h))
	}
}

func (ptr *QQuickView) ConnectSetMinimumHeight(f func(h int)) {
	defer qt.Recovering("connect QQuickView::setMinimumHeight")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMinimumHeight", f)
	}
}

func (ptr *QQuickView) DisconnectSetMinimumHeight() {
	defer qt.Recovering("disconnect QQuickView::setMinimumHeight")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMinimumHeight")
	}
}

func (ptr *QQuickView) SetMinimumHeight(h int) {
	defer qt.Recovering("QQuickView::setMinimumHeight")

	if ptr.Pointer() != nil {
		C.QQuickView_SetMinimumHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QQuickView) SetMinimumHeightDefault(h int) {
	defer qt.Recovering("QQuickView::setMinimumHeight")

	if ptr.Pointer() != nil {
		C.QQuickView_SetMinimumHeightDefault(ptr.Pointer(), C.int(h))
	}
}

//export callbackQQuickView_SetMinimumWidth
func callbackQQuickView_SetMinimumWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) {
	defer qt.Recovering("callback QQuickView::setMinimumWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMinimumWidth"); signal != nil {
		signal.(func(int))(int(w))
	} else {
		NewQQuickViewFromPointer(ptr).SetMinimumWidthDefault(int(w))
	}
}

func (ptr *QQuickView) ConnectSetMinimumWidth(f func(w int)) {
	defer qt.Recovering("connect QQuickView::setMinimumWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMinimumWidth", f)
	}
}

func (ptr *QQuickView) DisconnectSetMinimumWidth() {
	defer qt.Recovering("disconnect QQuickView::setMinimumWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMinimumWidth")
	}
}

func (ptr *QQuickView) SetMinimumWidth(w int) {
	defer qt.Recovering("QQuickView::setMinimumWidth")

	if ptr.Pointer() != nil {
		C.QQuickView_SetMinimumWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QQuickView) SetMinimumWidthDefault(w int) {
	defer qt.Recovering("QQuickView::setMinimumWidth")

	if ptr.Pointer() != nil {
		C.QQuickView_SetMinimumWidthDefault(ptr.Pointer(), C.int(w))
	}
}

//export callbackQQuickView_SetTitle
func callbackQQuickView_SetTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback QQuickView::setTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQQuickViewFromPointer(ptr).SetTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QQuickView) ConnectSetTitle(f func(vqs string)) {
	defer qt.Recovering("connect QQuickView::setTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setTitle", f)
	}
}

func (ptr *QQuickView) DisconnectSetTitle() {
	defer qt.Recovering("disconnect QQuickView::setTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setTitle")
	}
}

func (ptr *QQuickView) SetTitle(vqs string) {
	defer qt.Recovering("QQuickView::setTitle")

	if ptr.Pointer() != nil {
		C.QQuickView_SetTitle(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QQuickView) SetTitleDefault(vqs string) {
	defer qt.Recovering("QQuickView::setTitle")

	if ptr.Pointer() != nil {
		C.QQuickView_SetTitleDefault(ptr.Pointer(), C.CString(vqs))
	}
}

//export callbackQQuickView_SetVisible
func callbackQQuickView_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QQuickView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQQuickViewFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *QQuickView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QQuickView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QQuickView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QQuickView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

func (ptr *QQuickView) SetVisible(visible bool) {
	defer qt.Recovering("QQuickView::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QQuickView) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QQuickView::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickView_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackQQuickView_SetWidth
func callbackQQuickView_SetWidth(ptr unsafe.Pointer, ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QQuickView::setWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWidth"); signal != nil {
		signal.(func(int))(int(arg))
	} else {
		NewQQuickViewFromPointer(ptr).SetWidthDefault(int(arg))
	}
}

func (ptr *QQuickView) ConnectSetWidth(f func(arg int)) {
	defer qt.Recovering("connect QQuickView::setWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWidth", f)
	}
}

func (ptr *QQuickView) DisconnectSetWidth() {
	defer qt.Recovering("disconnect QQuickView::setWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWidth")
	}
}

func (ptr *QQuickView) SetWidth(arg int) {
	defer qt.Recovering("QQuickView::setWidth")

	if ptr.Pointer() != nil {
		C.QQuickView_SetWidth(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QQuickView) SetWidthDefault(arg int) {
	defer qt.Recovering("QQuickView::setWidth")

	if ptr.Pointer() != nil {
		C.QQuickView_SetWidthDefault(ptr.Pointer(), C.int(arg))
	}
}

//export callbackQQuickView_SetX
func callbackQQuickView_SetX(ptr unsafe.Pointer, ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QQuickView::setX")

	if signal := qt.GetSignal(C.GoString(ptrName), "setX"); signal != nil {
		signal.(func(int))(int(arg))
	} else {
		NewQQuickViewFromPointer(ptr).SetXDefault(int(arg))
	}
}

func (ptr *QQuickView) ConnectSetX(f func(arg int)) {
	defer qt.Recovering("connect QQuickView::setX")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setX", f)
	}
}

func (ptr *QQuickView) DisconnectSetX() {
	defer qt.Recovering("disconnect QQuickView::setX")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setX")
	}
}

func (ptr *QQuickView) SetX(arg int) {
	defer qt.Recovering("QQuickView::setX")

	if ptr.Pointer() != nil {
		C.QQuickView_SetX(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QQuickView) SetXDefault(arg int) {
	defer qt.Recovering("QQuickView::setX")

	if ptr.Pointer() != nil {
		C.QQuickView_SetXDefault(ptr.Pointer(), C.int(arg))
	}
}

//export callbackQQuickView_SetY
func callbackQQuickView_SetY(ptr unsafe.Pointer, ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QQuickView::setY")

	if signal := qt.GetSignal(C.GoString(ptrName), "setY"); signal != nil {
		signal.(func(int))(int(arg))
	} else {
		NewQQuickViewFromPointer(ptr).SetYDefault(int(arg))
	}
}

func (ptr *QQuickView) ConnectSetY(f func(arg int)) {
	defer qt.Recovering("connect QQuickView::setY")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setY", f)
	}
}

func (ptr *QQuickView) DisconnectSetY() {
	defer qt.Recovering("disconnect QQuickView::setY")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setY")
	}
}

func (ptr *QQuickView) SetY(arg int) {
	defer qt.Recovering("QQuickView::setY")

	if ptr.Pointer() != nil {
		C.QQuickView_SetY(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QQuickView) SetYDefault(arg int) {
	defer qt.Recovering("QQuickView::setY")

	if ptr.Pointer() != nil {
		C.QQuickView_SetYDefault(ptr.Pointer(), C.int(arg))
	}
}

//export callbackQQuickView_Alert
func callbackQQuickView_Alert(ptr unsafe.Pointer, ptrName *C.char, msec C.int) {
	defer qt.Recovering("callback QQuickView::alert")

	if signal := qt.GetSignal(C.GoString(ptrName), "alert"); signal != nil {
		signal.(func(int))(int(msec))
	} else {
		NewQQuickViewFromPointer(ptr).AlertDefault(int(msec))
	}
}

func (ptr *QQuickView) ConnectAlert(f func(msec int)) {
	defer qt.Recovering("connect QQuickView::alert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "alert", f)
	}
}

func (ptr *QQuickView) DisconnectAlert() {
	defer qt.Recovering("disconnect QQuickView::alert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "alert")
	}
}

func (ptr *QQuickView) Alert(msec int) {
	defer qt.Recovering("QQuickView::alert")

	if ptr.Pointer() != nil {
		C.QQuickView_Alert(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QQuickView) AlertDefault(msec int) {
	defer qt.Recovering("QQuickView::alert")

	if ptr.Pointer() != nil {
		C.QQuickView_AlertDefault(ptr.Pointer(), C.int(msec))
	}
}

//export callbackQQuickView_Close
func callbackQQuickView_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickView::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQQuickViewFromPointer(ptr).CloseDefault()))
}

func (ptr *QQuickView) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QQuickView::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QQuickView) DisconnectClose() {
	defer qt.Recovering("disconnect QQuickView::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QQuickView) Close() bool {
	defer qt.Recovering("QQuickView::close")

	if ptr.Pointer() != nil {
		return C.QQuickView_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickView) CloseDefault() bool {
	defer qt.Recovering("QQuickView::close")

	if ptr.Pointer() != nil {
		return C.QQuickView_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickView_FocusObject
func callbackQQuickView_FocusObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickView::focusObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusObject"); signal != nil {
		return core.PointerFromQObject(signal.(func() *core.QObject)())
	}

	return core.PointerFromQObject(NewQQuickViewFromPointer(ptr).FocusObjectDefault())
}

func (ptr *QQuickView) ConnectFocusObject(f func() *core.QObject) {
	defer qt.Recovering("connect QQuickView::focusObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusObject", f)
	}
}

func (ptr *QQuickView) DisconnectFocusObject() {
	defer qt.Recovering("disconnect QQuickView::focusObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusObject")
	}
}

func (ptr *QQuickView) FocusObject() *core.QObject {
	defer qt.Recovering("QQuickView::focusObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQuickView_FocusObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) FocusObjectDefault() *core.QObject {
	defer qt.Recovering("QQuickView::focusObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQuickView_FocusObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickView_Format
func callbackQQuickView_Format(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickView::format")

	if signal := qt.GetSignal(C.GoString(ptrName), "format"); signal != nil {
		return gui.PointerFromQSurfaceFormat(signal.(func() *gui.QSurfaceFormat)())
	}

	return gui.PointerFromQSurfaceFormat(NewQQuickViewFromPointer(ptr).FormatDefault())
}

func (ptr *QQuickView) ConnectFormat(f func() *gui.QSurfaceFormat) {
	defer qt.Recovering("connect QQuickView::format")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "format", f)
	}
}

func (ptr *QQuickView) DisconnectFormat() {
	defer qt.Recovering("disconnect QQuickView::format")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "format")
	}
}

func (ptr *QQuickView) Format() *gui.QSurfaceFormat {
	defer qt.Recovering("QQuickView::format")

	if ptr.Pointer() != nil {
		return gui.NewQSurfaceFormatFromPointer(C.QQuickView_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) FormatDefault() *gui.QSurfaceFormat {
	defer qt.Recovering("QQuickView::format")

	if ptr.Pointer() != nil {
		return gui.NewQSurfaceFormatFromPointer(C.QQuickView_FormatDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickView_Hide
func callbackQQuickView_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).HideDefault()
	}
}

func (ptr *QQuickView) ConnectHide(f func()) {
	defer qt.Recovering("connect QQuickView::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *QQuickView) DisconnectHide() {
	defer qt.Recovering("disconnect QQuickView::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *QQuickView) Hide() {
	defer qt.Recovering("QQuickView::hide")

	if ptr.Pointer() != nil {
		C.QQuickView_Hide(ptr.Pointer())
	}
}

func (ptr *QQuickView) HideDefault() {
	defer qt.Recovering("QQuickView::hide")

	if ptr.Pointer() != nil {
		C.QQuickView_HideDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_Lower
func callbackQQuickView_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QQuickView) ConnectLower(f func()) {
	defer qt.Recovering("connect QQuickView::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *QQuickView) DisconnectLower() {
	defer qt.Recovering("disconnect QQuickView::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *QQuickView) Lower() {
	defer qt.Recovering("QQuickView::lower")

	if ptr.Pointer() != nil {
		C.QQuickView_Lower(ptr.Pointer())
	}
}

func (ptr *QQuickView) LowerDefault() {
	defer qt.Recovering("QQuickView::lower")

	if ptr.Pointer() != nil {
		C.QQuickView_LowerDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_MoveEvent
func callbackQQuickView_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(ev))
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

//export callbackQQuickView_NativeEvent
func callbackQQuickView_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback QQuickView::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(C.GoString(eventType), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickViewFromPointer(ptr).NativeEventDefault(C.GoString(eventType), message, int(result))))
}

func (ptr *QQuickView) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QQuickView::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *QQuickView) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QQuickView::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *QQuickView) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QQuickView::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QQuickView_NativeEvent(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

func (ptr *QQuickView) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QQuickView::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QQuickView_NativeEventDefault(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

//export callbackQQuickView_Raise
func callbackQQuickView_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QQuickView) ConnectRaise(f func()) {
	defer qt.Recovering("connect QQuickView::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *QQuickView) DisconnectRaise() {
	defer qt.Recovering("disconnect QQuickView::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *QQuickView) Raise() {
	defer qt.Recovering("QQuickView::raise")

	if ptr.Pointer() != nil {
		C.QQuickView_Raise(ptr.Pointer())
	}
}

func (ptr *QQuickView) RaiseDefault() {
	defer qt.Recovering("QQuickView::raise")

	if ptr.Pointer() != nil {
		C.QQuickView_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_RequestActivate
func callbackQQuickView_RequestActivate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::requestActivate")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestActivate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).RequestActivateDefault()
	}
}

func (ptr *QQuickView) ConnectRequestActivate(f func()) {
	defer qt.Recovering("connect QQuickView::requestActivate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "requestActivate", f)
	}
}

func (ptr *QQuickView) DisconnectRequestActivate() {
	defer qt.Recovering("disconnect QQuickView::requestActivate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "requestActivate")
	}
}

func (ptr *QQuickView) RequestActivate() {
	defer qt.Recovering("QQuickView::requestActivate")

	if ptr.Pointer() != nil {
		C.QQuickView_RequestActivate(ptr.Pointer())
	}
}

func (ptr *QQuickView) RequestActivateDefault() {
	defer qt.Recovering("QQuickView::requestActivate")

	if ptr.Pointer() != nil {
		C.QQuickView_RequestActivateDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_RequestUpdate
func callbackQQuickView_RequestUpdate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::requestUpdate")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestUpdate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).RequestUpdateDefault()
	}
}

func (ptr *QQuickView) ConnectRequestUpdate(f func()) {
	defer qt.Recovering("connect QQuickView::requestUpdate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "requestUpdate", f)
	}
}

func (ptr *QQuickView) DisconnectRequestUpdate() {
	defer qt.Recovering("disconnect QQuickView::requestUpdate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "requestUpdate")
	}
}

func (ptr *QQuickView) RequestUpdate() {
	defer qt.Recovering("QQuickView::requestUpdate")

	if ptr.Pointer() != nil {
		C.QQuickView_RequestUpdate(ptr.Pointer())
	}
}

func (ptr *QQuickView) RequestUpdateDefault() {
	defer qt.Recovering("QQuickView::requestUpdate")

	if ptr.Pointer() != nil {
		C.QQuickView_RequestUpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_Show
func callbackQQuickView_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QQuickView) ConnectShow(f func()) {
	defer qt.Recovering("connect QQuickView::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *QQuickView) DisconnectShow() {
	defer qt.Recovering("disconnect QQuickView::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *QQuickView) Show() {
	defer qt.Recovering("QQuickView::show")

	if ptr.Pointer() != nil {
		C.QQuickView_Show(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowDefault() {
	defer qt.Recovering("QQuickView::show")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_ShowFullScreen
func callbackQQuickView_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QQuickView) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QQuickView::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *QQuickView) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QQuickView::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *QQuickView) ShowFullScreen() {
	defer qt.Recovering("QQuickView::showFullScreen")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowFullScreenDefault() {
	defer qt.Recovering("QQuickView::showFullScreen")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_ShowMaximized
func callbackQQuickView_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QQuickView) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QQuickView::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *QQuickView) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QQuickView::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *QQuickView) ShowMaximized() {
	defer qt.Recovering("QQuickView::showMaximized")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowMaximizedDefault() {
	defer qt.Recovering("QQuickView::showMaximized")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_ShowMinimized
func callbackQQuickView_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QQuickView) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QQuickView::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *QQuickView) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QQuickView::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *QQuickView) ShowMinimized() {
	defer qt.Recovering("QQuickView::showMinimized")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowMinimizedDefault() {
	defer qt.Recovering("QQuickView::showMinimized")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_ShowNormal
func callbackQQuickView_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QQuickView) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QQuickView::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *QQuickView) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QQuickView::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *QQuickView) ShowNormal() {
	defer qt.Recovering("QQuickView::showNormal")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QQuickView) ShowNormalDefault() {
	defer qt.Recovering("QQuickView::showNormal")

	if ptr.Pointer() != nil {
		C.QQuickView_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQQuickView_Size
func callbackQQuickView_Size(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickView::size")

	if signal := qt.GetSignal(C.GoString(ptrName), "size"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickViewFromPointer(ptr).SizeDefault())
}

func (ptr *QQuickView) ConnectSize(f func() *core.QSize) {
	defer qt.Recovering("connect QQuickView::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "size", f)
	}
}

func (ptr *QQuickView) DisconnectSize() {
	defer qt.Recovering("disconnect QQuickView::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "size")
	}
}

func (ptr *QQuickView) Size() *core.QSize {
	defer qt.Recovering("QQuickView::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickView_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) SizeDefault() *core.QSize {
	defer qt.Recovering("QQuickView::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickView_SizeDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickView_SurfaceType
func callbackQQuickView_SurfaceType(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickView::surfaceType")

	if signal := qt.GetSignal(C.GoString(ptrName), "surfaceType"); signal != nil {
		return C.int(signal.(func() gui.QSurface__SurfaceType)())
	}

	return C.int(NewQQuickViewFromPointer(ptr).SurfaceTypeDefault())
}

func (ptr *QQuickView) ConnectSurfaceType(f func() gui.QSurface__SurfaceType) {
	defer qt.Recovering("connect QQuickView::surfaceType")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "surfaceType", f)
	}
}

func (ptr *QQuickView) DisconnectSurfaceType() {
	defer qt.Recovering("disconnect QQuickView::surfaceType")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "surfaceType")
	}
}

func (ptr *QQuickView) SurfaceType() gui.QSurface__SurfaceType {
	defer qt.Recovering("QQuickView::surfaceType")

	if ptr.Pointer() != nil {
		return gui.QSurface__SurfaceType(C.QQuickView_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickView) SurfaceTypeDefault() gui.QSurface__SurfaceType {
	defer qt.Recovering("QQuickView::surfaceType")

	if ptr.Pointer() != nil {
		return gui.QSurface__SurfaceType(C.QQuickView_SurfaceTypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickView_TabletEvent
func callbackQQuickView_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(ev))
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

//export callbackQQuickView_TouchEvent
func callbackQQuickView_TouchEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(ev))
	} else {
		NewQQuickViewFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(ev))
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

//export callbackQQuickView_TimerEvent
func callbackQQuickView_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQQuickView_ChildEvent
func callbackQQuickView_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickView_ConnectNotify
func callbackQQuickView_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickViewFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickView) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickView::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickView) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickView::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickView) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickView::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickView_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickView) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickView::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickView_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickView_CustomEvent
func callbackQQuickView_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickView_DeleteLater
func callbackQQuickView_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickView::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickViewFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickView) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickView::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickView) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickView::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickView) DeleteLater() {
	defer qt.Recovering("QQuickView::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickView_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickView) DeleteLaterDefault() {
	defer qt.Recovering("QQuickView::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickView_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickView_DisconnectNotify
func callbackQQuickView_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickView::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickViewFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickView) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickView::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickView) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickView::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickView) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickView::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickView_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickView) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickView::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickView_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickView_EventFilter
func callbackQQuickView_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickView::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickViewFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickView) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickView::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickView) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickView::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickView) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickView::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickView_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickView) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickView::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickView_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickView_MetaObject
func callbackQQuickView_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickView::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickViewFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickView) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickView::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickView) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickView::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickView) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickView::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickView_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickView) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickView::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickView_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
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

type QQuickWidget struct {
	widgets.QWidget
}

type QQuickWidget_ITF interface {
	widgets.QWidget_ITF
	QQuickWidget_PTR() *QQuickWidget
}

func (p *QQuickWidget) QQuickWidget_PTR() *QQuickWidget {
	return p
}

func (p *QQuickWidget) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWidget_PTR().Pointer()
	}
	return nil
}

func (p *QQuickWidget) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWidget_PTR().SetPointer(ptr)
	}
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

func (ptr *QQuickWidget) ResizeMode() QQuickWidget__ResizeMode {
	defer qt.Recovering("QQuickWidget::resizeMode")

	if ptr.Pointer() != nil {
		return QQuickWidget__ResizeMode(C.QQuickWidget_ResizeMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWidget) SetResizeMode(vre QQuickWidget__ResizeMode) {
	defer qt.Recovering("QQuickWidget::setResizeMode")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetResizeMode(ptr.Pointer(), C.int(vre))
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

//export callbackQQuickWidget_DragEnterEvent
func callbackQQuickWidget_DragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(e))
	}
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

//export callbackQQuickWidget_DragLeaveEvent
func callbackQQuickWidget_DragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(e))
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

//export callbackQQuickWidget_DragMoveEvent
func callbackQQuickWidget_DragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(e))
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

//export callbackQQuickWidget_DropEvent
func callbackQQuickWidget_DropEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(e))
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

//export callbackQQuickWidget_FocusInEvent
func callbackQQuickWidget_FocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

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

//export callbackQQuickWidget_FocusOutEvent
func callbackQQuickWidget_FocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
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

//export callbackQQuickWidget_HideEvent
func callbackQQuickWidget_HideEvent(ptr unsafe.Pointer, ptrName *C.char, vqh unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(vqh))
	} else {
		NewQQuickWidgetFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(vqh))
	}
}

func (ptr *QQuickWidget) ConnectHideEvent(f func(vqh *gui.QHideEvent)) {
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

func (ptr *QQuickWidget) HideEvent(vqh gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

func (ptr *QQuickWidget) HideEventDefault(vqh gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickWidget::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
	}
}

func (ptr *QQuickWidget) InitialSize() *core.QSize {
	defer qt.Recovering("QQuickWidget::initialSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWidget_InitialSize(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWidget_KeyPressEvent
func callbackQQuickWidget_KeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(e))
	}
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

//export callbackQQuickWidget_KeyReleaseEvent
func callbackQQuickWidget_KeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(e))
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

//export callbackQQuickWidget_MouseDoubleClickEvent
func callbackQQuickWidget_MouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(e))
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

//export callbackQQuickWidget_MouseMoveEvent
func callbackQQuickWidget_MouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(e))
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

//export callbackQQuickWidget_MousePressEvent
func callbackQQuickWidget_MousePressEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(e))
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

//export callbackQQuickWidget_MouseReleaseEvent
func callbackQQuickWidget_MouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(e))
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

//export callbackQQuickWidget_SceneGraphError
func callbackQQuickWidget_SceneGraphError(ptr unsafe.Pointer, ptrName *C.char, error C.int, message *C.char) {
	defer qt.Recovering("callback QQuickWidget::sceneGraphError")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
	}

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

//export callbackQQuickWidget_SetSource
func callbackQQuickWidget_SetSource(ptr unsafe.Pointer, ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::setSource")

	if signal := qt.GetSignal(C.GoString(ptrName), "setSource"); signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QQuickWidget) ConnectSetSource(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QQuickWidget::setSource")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setSource", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetSource(url core.QUrl_ITF) {
	defer qt.Recovering("disconnect QQuickWidget::setSource")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setSource")
	}
}

func (ptr *QQuickWidget) SetSource(url core.QUrl_ITF) {
	defer qt.Recovering("QQuickWidget::setSource")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetSource(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

//export callbackQQuickWidget_ShowEvent
func callbackQQuickWidget_ShowEvent(ptr unsafe.Pointer, ptrName *C.char, vqs unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(vqs))
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(vqs))
	}
}

func (ptr *QQuickWidget) ConnectShowEvent(f func(vqs *gui.QShowEvent)) {
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

func (ptr *QQuickWidget) ShowEvent(vqs gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

func (ptr *QQuickWidget) ShowEventDefault(vqs gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickWidget::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackQQuickWidget_StatusChanged
func callbackQQuickWidget_StatusChanged(ptr unsafe.Pointer, ptrName *C.char, status C.int) {
	defer qt.Recovering("callback QQuickWidget::statusChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "statusChanged"); signal != nil {
		signal.(func(QQuickWidget__Status))(QQuickWidget__Status(status))
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

//export callbackQQuickWidget_WheelEvent
func callbackQQuickWidget_WheelEvent(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(e))
	} else {
		NewQQuickWidgetFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(e))
	}
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
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickWidget_DestroyQQuickWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickWidget_ActionEvent
func callbackQQuickWidget_ActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
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

//export callbackQQuickWidget_EnterEvent
func callbackQQuickWidget_EnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickWidget_LeaveEvent
func callbackQQuickWidget_LeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickWidget_MinimumSizeHint
func callbackQQuickWidget_MinimumSizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickWidget::minimumSizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "minimumSizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickWidgetFromPointer(ptr).MinimumSizeHintDefault())
}

func (ptr *QQuickWidget) ConnectMinimumSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QQuickWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "minimumSizeHint", f)
	}
}

func (ptr *QQuickWidget) DisconnectMinimumSizeHint() {
	defer qt.Recovering("disconnect QQuickWidget::minimumSizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "minimumSizeHint")
	}
}

func (ptr *QQuickWidget) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QQuickWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWidget_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) MinimumSizeHintDefault() *core.QSize {
	defer qt.Recovering("QQuickWidget::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWidget_MinimumSizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWidget_MoveEvent
func callbackQQuickWidget_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
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

//export callbackQQuickWidget_PaintEvent
func callbackQQuickWidget_PaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
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

//export callbackQQuickWidget_SetEnabled
func callbackQQuickWidget_SetEnabled(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QQuickWidget::setEnabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setEnabled"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetEnabledDefault(int(vbo) != 0)
	}
}

func (ptr *QQuickWidget) ConnectSetEnabled(f func(vbo bool)) {
	defer qt.Recovering("connect QQuickWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setEnabled", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetEnabled() {
	defer qt.Recovering("disconnect QQuickWidget::setEnabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setEnabled")
	}
}

func (ptr *QQuickWidget) SetEnabled(vbo bool) {
	defer qt.Recovering("QQuickWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QQuickWidget) SetEnabledDefault(vbo bool) {
	defer qt.Recovering("QQuickWidget::setEnabled")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetEnabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQQuickWidget_SetStyleSheet
func callbackQQuickWidget_SetStyleSheet(ptr unsafe.Pointer, ptrName *C.char, styleSheet *C.char) {
	defer qt.Recovering("callback QQuickWidget::setStyleSheet")

	if signal := qt.GetSignal(C.GoString(ptrName), "setStyleSheet"); signal != nil {
		signal.(func(string))(C.GoString(styleSheet))
	} else {
		NewQQuickWidgetFromPointer(ptr).SetStyleSheetDefault(C.GoString(styleSheet))
	}
}

func (ptr *QQuickWidget) ConnectSetStyleSheet(f func(styleSheet string)) {
	defer qt.Recovering("connect QQuickWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setStyleSheet", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetStyleSheet() {
	defer qt.Recovering("disconnect QQuickWidget::setStyleSheet")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setStyleSheet")
	}
}

func (ptr *QQuickWidget) SetStyleSheet(styleSheet string) {
	defer qt.Recovering("QQuickWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetStyleSheet(ptr.Pointer(), C.CString(styleSheet))
	}
}

func (ptr *QQuickWidget) SetStyleSheetDefault(styleSheet string) {
	defer qt.Recovering("QQuickWidget::setStyleSheet")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetStyleSheetDefault(ptr.Pointer(), C.CString(styleSheet))
	}
}

//export callbackQQuickWidget_SetVisible
func callbackQQuickWidget_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QQuickWidget::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
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

//export callbackQQuickWidget_SetWindowModified
func callbackQQuickWidget_SetWindowModified(ptr unsafe.Pointer, ptrName *C.char, vbo C.int) {
	defer qt.Recovering("callback QQuickWidget::setWindowModified")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowModified"); signal != nil {
		signal.(func(bool))(int(vbo) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetWindowModifiedDefault(int(vbo) != 0)
	}
}

func (ptr *QQuickWidget) ConnectSetWindowModified(f func(vbo bool)) {
	defer qt.Recovering("connect QQuickWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowModified", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetWindowModified() {
	defer qt.Recovering("disconnect QQuickWidget::setWindowModified")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowModified")
	}
}

func (ptr *QQuickWidget) SetWindowModified(vbo bool) {
	defer qt.Recovering("QQuickWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetWindowModified(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

func (ptr *QQuickWidget) SetWindowModifiedDefault(vbo bool) {
	defer qt.Recovering("QQuickWidget::setWindowModified")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetWindowModifiedDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(vbo)))
	}
}

//export callbackQQuickWidget_SetWindowTitle
func callbackQQuickWidget_SetWindowTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback QQuickWidget::setWindowTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWindowTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQQuickWidgetFromPointer(ptr).SetWindowTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QQuickWidget) ConnectSetWindowTitle(f func(vqs string)) {
	defer qt.Recovering("connect QQuickWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWindowTitle", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetWindowTitle() {
	defer qt.Recovering("disconnect QQuickWidget::setWindowTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWindowTitle")
	}
}

func (ptr *QQuickWidget) SetWindowTitle(vqs string) {
	defer qt.Recovering("QQuickWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetWindowTitle(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QQuickWidget) SetWindowTitleDefault(vqs string) {
	defer qt.Recovering("QQuickWidget::setWindowTitle")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetWindowTitleDefault(ptr.Pointer(), C.CString(vqs))
	}
}

//export callbackQQuickWidget_SizeHint
func callbackQQuickWidget_SizeHint(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickWidget::sizeHint")

	if signal := qt.GetSignal(C.GoString(ptrName), "sizeHint"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickWidgetFromPointer(ptr).SizeHintDefault())
}

func (ptr *QQuickWidget) ConnectSizeHint(f func() *core.QSize) {
	defer qt.Recovering("connect QQuickWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "sizeHint", f)
	}
}

func (ptr *QQuickWidget) DisconnectSizeHint() {
	defer qt.Recovering("disconnect QQuickWidget::sizeHint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "sizeHint")
	}
}

func (ptr *QQuickWidget) SizeHint() *core.QSize {
	defer qt.Recovering("QQuickWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWidget_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) SizeHintDefault() *core.QSize {
	defer qt.Recovering("QQuickWidget::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWidget_SizeHintDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWidget_ChangeEvent
func callbackQQuickWidget_ChangeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickWidget_Close
func callbackQQuickWidget_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickWidget::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).CloseDefault()))
}

func (ptr *QQuickWidget) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QQuickWidget::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QQuickWidget) DisconnectClose() {
	defer qt.Recovering("disconnect QQuickWidget::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QQuickWidget) Close() bool {
	defer qt.Recovering("QQuickWidget::close")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWidget) CloseDefault() bool {
	defer qt.Recovering("QQuickWidget::close")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickWidget_CloseEvent
func callbackQQuickWidget_CloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
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

//export callbackQQuickWidget_ContextMenuEvent
func callbackQQuickWidget_ContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
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

//export callbackQQuickWidget_FocusNextPrevChild
func callbackQQuickWidget_FocusNextPrevChild(ptr unsafe.Pointer, ptrName *C.char, next C.int) C.int {
	defer qt.Recovering("callback QQuickWidget::focusNextPrevChild")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusNextPrevChild"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(bool) bool)(int(next) != 0)))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).FocusNextPrevChildDefault(int(next) != 0)))
}

func (ptr *QQuickWidget) ConnectFocusNextPrevChild(f func(next bool) bool) {
	defer qt.Recovering("connect QQuickWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusNextPrevChild", f)
	}
}

func (ptr *QQuickWidget) DisconnectFocusNextPrevChild() {
	defer qt.Recovering("disconnect QQuickWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusNextPrevChild")
	}
}

func (ptr *QQuickWidget) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QQuickWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QQuickWidget) FocusNextPrevChildDefault(next bool) bool {
	defer qt.Recovering("QQuickWidget::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_FocusNextPrevChildDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

//export callbackQQuickWidget_HasHeightForWidth
func callbackQQuickWidget_HasHeightForWidth(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickWidget::hasHeightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasHeightForWidth"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).HasHeightForWidthDefault()))
}

func (ptr *QQuickWidget) ConnectHasHeightForWidth(f func() bool) {
	defer qt.Recovering("connect QQuickWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasHeightForWidth", f)
	}
}

func (ptr *QQuickWidget) DisconnectHasHeightForWidth() {
	defer qt.Recovering("disconnect QQuickWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasHeightForWidth")
	}
}

func (ptr *QQuickWidget) HasHeightForWidth() bool {
	defer qt.Recovering("QQuickWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_HasHeightForWidth(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWidget) HasHeightForWidthDefault() bool {
	defer qt.Recovering("QQuickWidget::hasHeightForWidth")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_HasHeightForWidthDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickWidget_HeightForWidth
func callbackQQuickWidget_HeightForWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) C.int {
	defer qt.Recovering("callback QQuickWidget::heightForWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "heightForWidth"); signal != nil {
		return C.int(signal.(func(int) int)(int(w)))
	}

	return C.int(NewQQuickWidgetFromPointer(ptr).HeightForWidthDefault(int(w)))
}

func (ptr *QQuickWidget) ConnectHeightForWidth(f func(w int) int) {
	defer qt.Recovering("connect QQuickWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "heightForWidth", f)
	}
}

func (ptr *QQuickWidget) DisconnectHeightForWidth() {
	defer qt.Recovering("disconnect QQuickWidget::heightForWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "heightForWidth")
	}
}

func (ptr *QQuickWidget) HeightForWidth(w int) int {
	defer qt.Recovering("QQuickWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QQuickWidget_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QQuickWidget) HeightForWidthDefault(w int) int {
	defer qt.Recovering("QQuickWidget::heightForWidth")

	if ptr.Pointer() != nil {
		return int(C.QQuickWidget_HeightForWidthDefault(ptr.Pointer(), C.int(w)))
	}
	return 0
}

//export callbackQQuickWidget_Hide
func callbackQQuickWidget_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).HideDefault()
	}
}

func (ptr *QQuickWidget) ConnectHide(f func()) {
	defer qt.Recovering("connect QQuickWidget::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *QQuickWidget) DisconnectHide() {
	defer qt.Recovering("disconnect QQuickWidget::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *QQuickWidget) Hide() {
	defer qt.Recovering("QQuickWidget::hide")

	if ptr.Pointer() != nil {
		C.QQuickWidget_Hide(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) HideDefault() {
	defer qt.Recovering("QQuickWidget::hide")

	if ptr.Pointer() != nil {
		C.QQuickWidget_HideDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_InputMethodEvent
func callbackQQuickWidget_InputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
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

//export callbackQQuickWidget_InputMethodQuery
func callbackQQuickWidget_InputMethodQuery(ptr unsafe.Pointer, ptrName *C.char, query C.int) unsafe.Pointer {
	defer qt.Recovering("callback QQuickWidget::inputMethodQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodQuery"); signal != nil {
		return core.PointerFromQVariant(signal.(func(core.Qt__InputMethodQuery) *core.QVariant)(core.Qt__InputMethodQuery(query)))
	}

	return core.PointerFromQVariant(NewQQuickWidgetFromPointer(ptr).InputMethodQueryDefault(core.Qt__InputMethodQuery(query)))
}

func (ptr *QQuickWidget) ConnectInputMethodQuery(f func(query core.Qt__InputMethodQuery) *core.QVariant) {
	defer qt.Recovering("connect QQuickWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodQuery", f)
	}
}

func (ptr *QQuickWidget) DisconnectInputMethodQuery() {
	defer qt.Recovering("disconnect QQuickWidget::inputMethodQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodQuery")
	}
}

func (ptr *QQuickWidget) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickWidget_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QQuickWidget) InputMethodQueryDefault(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QQuickWidget::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QQuickWidget_InputMethodQueryDefault(ptr.Pointer(), C.int(query)))
	}
	return nil
}

//export callbackQQuickWidget_Lower
func callbackQQuickWidget_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QQuickWidget) ConnectLower(f func()) {
	defer qt.Recovering("connect QQuickWidget::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *QQuickWidget) DisconnectLower() {
	defer qt.Recovering("disconnect QQuickWidget::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *QQuickWidget) Lower() {
	defer qt.Recovering("QQuickWidget::lower")

	if ptr.Pointer() != nil {
		C.QQuickWidget_Lower(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) LowerDefault() {
	defer qt.Recovering("QQuickWidget::lower")

	if ptr.Pointer() != nil {
		C.QQuickWidget_LowerDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_NativeEvent
func callbackQQuickWidget_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback QQuickWidget::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(C.GoString(eventType), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).NativeEventDefault(C.GoString(eventType), message, int(result))))
}

func (ptr *QQuickWidget) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QQuickWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *QQuickWidget) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QQuickWidget::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *QQuickWidget) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QQuickWidget::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_NativeEvent(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

func (ptr *QQuickWidget) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QQuickWidget::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_NativeEventDefault(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

//export callbackQQuickWidget_Raise
func callbackQQuickWidget_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QQuickWidget) ConnectRaise(f func()) {
	defer qt.Recovering("connect QQuickWidget::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *QQuickWidget) DisconnectRaise() {
	defer qt.Recovering("disconnect QQuickWidget::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *QQuickWidget) Raise() {
	defer qt.Recovering("QQuickWidget::raise")

	if ptr.Pointer() != nil {
		C.QQuickWidget_Raise(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) RaiseDefault() {
	defer qt.Recovering("QQuickWidget::raise")

	if ptr.Pointer() != nil {
		C.QQuickWidget_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_Repaint
func callbackQQuickWidget_Repaint(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::repaint")

	if signal := qt.GetSignal(C.GoString(ptrName), "repaint"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).RepaintDefault()
	}
}

func (ptr *QQuickWidget) ConnectRepaint(f func()) {
	defer qt.Recovering("connect QQuickWidget::repaint")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "repaint", f)
	}
}

func (ptr *QQuickWidget) DisconnectRepaint() {
	defer qt.Recovering("disconnect QQuickWidget::repaint")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "repaint")
	}
}

func (ptr *QQuickWidget) Repaint() {
	defer qt.Recovering("QQuickWidget::repaint")

	if ptr.Pointer() != nil {
		C.QQuickWidget_Repaint(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) RepaintDefault() {
	defer qt.Recovering("QQuickWidget::repaint")

	if ptr.Pointer() != nil {
		C.QQuickWidget_RepaintDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ResizeEvent
func callbackQQuickWidget_ResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
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

//export callbackQQuickWidget_SetDisabled
func callbackQQuickWidget_SetDisabled(ptr unsafe.Pointer, ptrName *C.char, disable C.int) {
	defer qt.Recovering("callback QQuickWidget::setDisabled")

	if signal := qt.GetSignal(C.GoString(ptrName), "setDisabled"); signal != nil {
		signal.(func(bool))(int(disable) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetDisabledDefault(int(disable) != 0)
	}
}

func (ptr *QQuickWidget) ConnectSetDisabled(f func(disable bool)) {
	defer qt.Recovering("connect QQuickWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setDisabled", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetDisabled() {
	defer qt.Recovering("disconnect QQuickWidget::setDisabled")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setDisabled")
	}
}

func (ptr *QQuickWidget) SetDisabled(disable bool) {
	defer qt.Recovering("QQuickWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetDisabled(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

func (ptr *QQuickWidget) SetDisabledDefault(disable bool) {
	defer qt.Recovering("QQuickWidget::setDisabled")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetDisabledDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(disable)))
	}
}

//export callbackQQuickWidget_SetFocus2
func callbackQQuickWidget_SetFocus2(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::setFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFocus2"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).SetFocus2Default()
	}
}

func (ptr *QQuickWidget) ConnectSetFocus2(f func()) {
	defer qt.Recovering("connect QQuickWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setFocus2", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetFocus2() {
	defer qt.Recovering("disconnect QQuickWidget::setFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setFocus2")
	}
}

func (ptr *QQuickWidget) SetFocus2() {
	defer qt.Recovering("QQuickWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetFocus2(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) SetFocus2Default() {
	defer qt.Recovering("QQuickWidget::setFocus")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetFocus2Default(ptr.Pointer())
	}
}

//export callbackQQuickWidget_SetHidden
func callbackQQuickWidget_SetHidden(ptr unsafe.Pointer, ptrName *C.char, hidden C.int) {
	defer qt.Recovering("callback QQuickWidget::setHidden")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHidden"); signal != nil {
		signal.(func(bool))(int(hidden) != 0)
	} else {
		NewQQuickWidgetFromPointer(ptr).SetHiddenDefault(int(hidden) != 0)
	}
}

func (ptr *QQuickWidget) ConnectSetHidden(f func(hidden bool)) {
	defer qt.Recovering("connect QQuickWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHidden", f)
	}
}

func (ptr *QQuickWidget) DisconnectSetHidden() {
	defer qt.Recovering("disconnect QQuickWidget::setHidden")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHidden")
	}
}

func (ptr *QQuickWidget) SetHidden(hidden bool) {
	defer qt.Recovering("QQuickWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetHidden(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

func (ptr *QQuickWidget) SetHiddenDefault(hidden bool) {
	defer qt.Recovering("QQuickWidget::setHidden")

	if ptr.Pointer() != nil {
		C.QQuickWidget_SetHiddenDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(hidden)))
	}
}

//export callbackQQuickWidget_Show
func callbackQQuickWidget_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QQuickWidget) ConnectShow(f func()) {
	defer qt.Recovering("connect QQuickWidget::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *QQuickWidget) DisconnectShow() {
	defer qt.Recovering("disconnect QQuickWidget::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *QQuickWidget) Show() {
	defer qt.Recovering("QQuickWidget::show")

	if ptr.Pointer() != nil {
		C.QQuickWidget_Show(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowDefault() {
	defer qt.Recovering("QQuickWidget::show")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowFullScreen
func callbackQQuickWidget_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QQuickWidget) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QQuickWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QQuickWidget::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *QQuickWidget) ShowFullScreen() {
	defer qt.Recovering("QQuickWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowFullScreenDefault() {
	defer qt.Recovering("QQuickWidget::showFullScreen")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowMaximized
func callbackQQuickWidget_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QQuickWidget) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QQuickWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QQuickWidget::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *QQuickWidget) ShowMaximized() {
	defer qt.Recovering("QQuickWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowMaximizedDefault() {
	defer qt.Recovering("QQuickWidget::showMaximized")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowMinimized
func callbackQQuickWidget_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QQuickWidget) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QQuickWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QQuickWidget::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *QQuickWidget) ShowMinimized() {
	defer qt.Recovering("QQuickWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowMinimizedDefault() {
	defer qt.Recovering("QQuickWidget::showMinimized")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_ShowNormal
func callbackQQuickWidget_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QQuickWidget) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QQuickWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *QQuickWidget) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QQuickWidget::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *QQuickWidget) ShowNormal() {
	defer qt.Recovering("QQuickWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) ShowNormalDefault() {
	defer qt.Recovering("QQuickWidget::showNormal")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_TabletEvent
func callbackQQuickWidget_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
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

//export callbackQQuickWidget_Update
func callbackQQuickWidget_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).UpdateDefault()
	}
}

func (ptr *QQuickWidget) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QQuickWidget::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QQuickWidget) DisconnectUpdate() {
	defer qt.Recovering("disconnect QQuickWidget::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QQuickWidget) Update() {
	defer qt.Recovering("QQuickWidget::update")

	if ptr.Pointer() != nil {
		C.QQuickWidget_Update(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) UpdateDefault() {
	defer qt.Recovering("QQuickWidget::update")

	if ptr.Pointer() != nil {
		C.QQuickWidget_UpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_UpdateMicroFocus
func callbackQQuickWidget_UpdateMicroFocus(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::updateMicroFocus")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateMicroFocus"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).UpdateMicroFocusDefault()
	}
}

func (ptr *QQuickWidget) ConnectUpdateMicroFocus(f func()) {
	defer qt.Recovering("connect QQuickWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateMicroFocus", f)
	}
}

func (ptr *QQuickWidget) DisconnectUpdateMicroFocus() {
	defer qt.Recovering("disconnect QQuickWidget::updateMicroFocus")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateMicroFocus")
	}
}

func (ptr *QQuickWidget) UpdateMicroFocus() {
	defer qt.Recovering("QQuickWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QQuickWidget_UpdateMicroFocus(ptr.Pointer())
	}
}

func (ptr *QQuickWidget) UpdateMicroFocusDefault() {
	defer qt.Recovering("QQuickWidget::updateMicroFocus")

	if ptr.Pointer() != nil {
		C.QQuickWidget_UpdateMicroFocusDefault(ptr.Pointer())
	}
}

//export callbackQQuickWidget_TimerEvent
func callbackQQuickWidget_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQQuickWidget_ChildEvent
func callbackQQuickWidget_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickWidget_ConnectNotify
func callbackQQuickWidget_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWidgetFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWidget) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickWidget) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickWidget::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickWidget) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWidget) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWidget::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWidget_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWidget_CustomEvent
func callbackQQuickWidget_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWidgetFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickWidget_DeleteLater
func callbackQQuickWidget_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWidget::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWidgetFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWidget) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickWidget) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickWidget::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickWidget) DeleteLater() {
	defer qt.Recovering("QQuickWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickWidget_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWidget) DeleteLaterDefault() {
	defer qt.Recovering("QQuickWidget::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickWidget_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickWidget_DisconnectNotify
func callbackQQuickWidget_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWidget::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWidgetFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWidget) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickWidget) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickWidget::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickWidget) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWidget) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWidget::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWidget_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWidget_EventFilter
func callbackQQuickWidget_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickWidget::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWidgetFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickWidget) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickWidget) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickWidget::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickWidget) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickWidget) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWidget::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickWidget_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickWidget_MetaObject
func callbackQQuickWidget_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickWidget::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickWidgetFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWidget) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickWidget) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickWidget::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickWidget) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWidget_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWidget) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickWidget::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWidget_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QQuickWindow::CreateTextureOption
type QQuickWindow__CreateTextureOption int64

const (
	QQuickWindow__TextureHasAlphaChannel = QQuickWindow__CreateTextureOption(0x0001)
	QQuickWindow__TextureHasMipmaps      = QQuickWindow__CreateTextureOption(0x0002)
	QQuickWindow__TextureOwnsGLTexture   = QQuickWindow__CreateTextureOption(0x0004)
	QQuickWindow__TextureCanUseAtlas     = QQuickWindow__CreateTextureOption(0x0008)
	QQuickWindow__TextureIsOpaque        = QQuickWindow__CreateTextureOption(0x0010)
)

//QQuickWindow::RenderStage
type QQuickWindow__RenderStage int64

const (
	QQuickWindow__BeforeSynchronizingStage = QQuickWindow__RenderStage(0)
	QQuickWindow__AfterSynchronizingStage  = QQuickWindow__RenderStage(1)
	QQuickWindow__BeforeRenderingStage     = QQuickWindow__RenderStage(2)
	QQuickWindow__AfterRenderingStage      = QQuickWindow__RenderStage(3)
	QQuickWindow__AfterSwapStage           = QQuickWindow__RenderStage(4)
	QQuickWindow__NoStage                  = QQuickWindow__RenderStage(5)
)

//QQuickWindow::SceneGraphError
type QQuickWindow__SceneGraphError int64

const (
	QQuickWindow__ContextNotAvailable = QQuickWindow__SceneGraphError(1)
)

type QQuickWindow struct {
	gui.QWindow
}

type QQuickWindow_ITF interface {
	gui.QWindow_ITF
	QQuickWindow_PTR() *QQuickWindow
}

func (p *QQuickWindow) QQuickWindow_PTR() *QQuickWindow {
	return p
}

func (p *QQuickWindow) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QWindow_PTR().Pointer()
	}
	return nil
}

func (p *QQuickWindow) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QWindow_PTR().SetPointer(ptr)
	}
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

//export callbackQQuickWindow_ActiveFocusItemChanged
func callbackQQuickWindow_ActiveFocusItemChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::activeFocusItemChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeFocusItemChanged"); signal != nil {
		signal.(func())()
	}

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

func (ptr *QQuickWindow) ActiveFocusItemChanged() {
	defer qt.Recovering("QQuickWindow::activeFocusItemChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ActiveFocusItemChanged(ptr.Pointer())
	}
}

//export callbackQQuickWindow_AfterAnimating
func callbackQQuickWindow_AfterAnimating(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::afterAnimating")

	if signal := qt.GetSignal(C.GoString(ptrName), "afterAnimating"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickWindow) AfterAnimating() {
	defer qt.Recovering("QQuickWindow::afterAnimating")

	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterAnimating(ptr.Pointer())
	}
}

//export callbackQQuickWindow_AfterRendering
func callbackQQuickWindow_AfterRendering(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::afterRendering")

	if signal := qt.GetSignal(C.GoString(ptrName), "afterRendering"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickWindow) AfterRendering() {
	defer qt.Recovering("QQuickWindow::afterRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterRendering(ptr.Pointer())
	}
}

//export callbackQQuickWindow_AfterSynchronizing
func callbackQQuickWindow_AfterSynchronizing(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::afterSynchronizing")

	if signal := qt.GetSignal(C.GoString(ptrName), "afterSynchronizing"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickWindow) AfterSynchronizing() {
	defer qt.Recovering("QQuickWindow::afterSynchronizing")

	if ptr.Pointer() != nil {
		C.QQuickWindow_AfterSynchronizing(ptr.Pointer())
	}
}

//export callbackQQuickWindow_BeforeRendering
func callbackQQuickWindow_BeforeRendering(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::beforeRendering")

	if signal := qt.GetSignal(C.GoString(ptrName), "beforeRendering"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickWindow) BeforeRendering() {
	defer qt.Recovering("QQuickWindow::beforeRendering")

	if ptr.Pointer() != nil {
		C.QQuickWindow_BeforeRendering(ptr.Pointer())
	}
}

//export callbackQQuickWindow_BeforeSynchronizing
func callbackQQuickWindow_BeforeSynchronizing(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::beforeSynchronizing")

	if signal := qt.GetSignal(C.GoString(ptrName), "beforeSynchronizing"); signal != nil {
		signal.(func())()
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

//export callbackQQuickWindow_ColorChanged
func callbackQQuickWindow_ColorChanged(ptr unsafe.Pointer, ptrName *C.char, vqc unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::colorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "colorChanged"); signal != nil {
		signal.(func(*gui.QColor))(gui.NewQColorFromPointer(vqc))
	}

}

func (ptr *QQuickWindow) ConnectColorChanged(f func(vqc *gui.QColor)) {
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

func (ptr *QQuickWindow) ColorChanged(vqc gui.QColor_ITF) {
	defer qt.Recovering("QQuickWindow::colorChanged")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ColorChanged(ptr.Pointer(), gui.PointerFromQColor(vqc))
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

func (ptr *QQuickWindow) ExposeEvent(vqe gui.QExposeEvent_ITF) {
	defer qt.Recovering("QQuickWindow::exposeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ExposeEvent(ptr.Pointer(), gui.PointerFromQExposeEvent(vqe))
	}
}

func (ptr *QQuickWindow) FocusInEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWindow::focusInEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

func (ptr *QQuickWindow) FocusOutEvent(ev gui.QFocusEvent_ITF) {
	defer qt.Recovering("QQuickWindow::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(ev))
	}
}

//export callbackQQuickWindow_FrameSwapped
func callbackQQuickWindow_FrameSwapped(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::frameSwapped")

	if signal := qt.GetSignal(C.GoString(ptrName), "frameSwapped"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickWindow) HasDefaultAlphaBuffer() bool {
	defer qt.Recovering("QQuickWindow::hasDefaultAlphaBuffer")

	return C.QQuickWindow_QQuickWindow_HasDefaultAlphaBuffer() != 0
}

func (ptr *QQuickWindow) HideEvent(vqh gui.QHideEvent_ITF) {
	defer qt.Recovering("QQuickWindow::hideEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(vqh))
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

func (ptr *QQuickWindow) KeyPressEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWindow::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWindow) KeyReleaseEvent(e gui.QKeyEvent_ITF) {
	defer qt.Recovering("QQuickWindow::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(e))
	}
}

func (ptr *QQuickWindow) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseGrabberItem() *QQuickItem {
	defer qt.Recovering("QQuickWindow::mouseGrabberItem")

	if ptr.Pointer() != nil {
		return NewQQuickItemFromPointer(C.QQuickWindow_MouseGrabberItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QQuickWindow::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QQuickWindow) OpenglContext() *gui.QOpenGLContext {
	defer qt.Recovering("QQuickWindow::openglContext")

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLContextFromPointer(C.QQuickWindow_OpenglContext(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWindow_OpenglContextCreated
func callbackQQuickWindow_OpenglContextCreated(ptr unsafe.Pointer, ptrName *C.char, context unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::openglContextCreated")

	if signal := qt.GetSignal(C.GoString(ptrName), "openglContextCreated"); signal != nil {
		signal.(func(*gui.QOpenGLContext))(gui.NewQOpenGLContextFromPointer(context))
	}

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

func (ptr *QQuickWindow) OpenglContextCreated(context gui.QOpenGLContext_ITF) {
	defer qt.Recovering("QQuickWindow::openglContextCreated")

	if ptr.Pointer() != nil {
		C.QQuickWindow_OpenglContextCreated(ptr.Pointer(), gui.PointerFromQOpenGLContext(context))
	}
}

//export callbackQQuickWindow_ReleaseResources
func callbackQQuickWindow_ReleaseResources(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectReleaseResources(f func()) {
	defer qt.Recovering("connect QQuickWindow::releaseResources")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "releaseResources", f)
	}
}

func (ptr *QQuickWindow) DisconnectReleaseResources() {
	defer qt.Recovering("disconnect QQuickWindow::releaseResources")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "releaseResources")
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

func (ptr *QQuickWindow) ResizeEvent(ev gui.QResizeEvent_ITF) {
	defer qt.Recovering("QQuickWindow::resizeEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(ev))
	}
}

//export callbackQQuickWindow_SceneGraphAboutToStop
func callbackQQuickWindow_SceneGraphAboutToStop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphAboutToStop")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphAboutToStop"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickWindow) SceneGraphAboutToStop() {
	defer qt.Recovering("QQuickWindow::sceneGraphAboutToStop")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphAboutToStop(ptr.Pointer())
	}
}

//export callbackQQuickWindow_SceneGraphError
func callbackQQuickWindow_SceneGraphError(ptr unsafe.Pointer, ptrName *C.char, error C.int, message *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphError")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphError"); signal != nil {
		signal.(func(QQuickWindow__SceneGraphError, string))(QQuickWindow__SceneGraphError(error), C.GoString(message))
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

func (ptr *QQuickWindow) SceneGraphError(error QQuickWindow__SceneGraphError, message string) {
	defer qt.Recovering("QQuickWindow::sceneGraphError")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphError(ptr.Pointer(), C.int(error), C.CString(message))
	}
}

//export callbackQQuickWindow_SceneGraphInitialized
func callbackQQuickWindow_SceneGraphInitialized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphInitialized")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphInitialized"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickWindow) SceneGraphInitialized() {
	defer qt.Recovering("QQuickWindow::sceneGraphInitialized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SceneGraphInitialized(ptr.Pointer())
	}
}

//export callbackQQuickWindow_SceneGraphInvalidated
func callbackQQuickWindow_SceneGraphInvalidated(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::sceneGraphInvalidated")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphInvalidated"); signal != nil {
		signal.(func())()
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

func (ptr *QQuickWindow) SetDefaultAlphaBuffer(useAlpha bool) {
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

func (ptr *QQuickWindow) ShowEvent(vqs gui.QShowEvent_ITF) {
	defer qt.Recovering("QQuickWindow::showEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(vqs))
	}
}

//export callbackQQuickWindow_Update
func callbackQQuickWindow_Update(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::update")

	if signal := qt.GetSignal(C.GoString(ptrName), "update"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QQuickWindow) ConnectUpdate(f func()) {
	defer qt.Recovering("connect QQuickWindow::update")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "update", f)
	}
}

func (ptr *QQuickWindow) DisconnectUpdate() {
	defer qt.Recovering("disconnect QQuickWindow::update")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "update")
	}
}

func (ptr *QQuickWindow) Update() {
	defer qt.Recovering("QQuickWindow::update")

	if ptr.Pointer() != nil {
		C.QQuickWindow_Update(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QQuickWindow::wheelEvent")

	if ptr.Pointer() != nil {
		C.QQuickWindow_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QQuickWindow) DestroyQQuickWindow() {
	defer qt.Recovering("QQuickWindow::~QQuickWindow")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickWindow_DestroyQQuickWindow(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickWindow_SetHeight
func callbackQQuickWindow_SetHeight(ptr unsafe.Pointer, ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QQuickWindow::setHeight")

	if signal := qt.GetSignal(C.GoString(ptrName), "setHeight"); signal != nil {
		signal.(func(int))(int(arg))
	} else {
		NewQQuickWindowFromPointer(ptr).SetHeightDefault(int(arg))
	}
}

func (ptr *QQuickWindow) ConnectSetHeight(f func(arg int)) {
	defer qt.Recovering("connect QQuickWindow::setHeight")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setHeight", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetHeight() {
	defer qt.Recovering("disconnect QQuickWindow::setHeight")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setHeight")
	}
}

func (ptr *QQuickWindow) SetHeight(arg int) {
	defer qt.Recovering("QQuickWindow::setHeight")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetHeight(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QQuickWindow) SetHeightDefault(arg int) {
	defer qt.Recovering("QQuickWindow::setHeight")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetHeightDefault(ptr.Pointer(), C.int(arg))
	}
}

//export callbackQQuickWindow_SetMaximumHeight
func callbackQQuickWindow_SetMaximumHeight(ptr unsafe.Pointer, ptrName *C.char, h C.int) {
	defer qt.Recovering("callback QQuickWindow::setMaximumHeight")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMaximumHeight"); signal != nil {
		signal.(func(int))(int(h))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMaximumHeightDefault(int(h))
	}
}

func (ptr *QQuickWindow) ConnectSetMaximumHeight(f func(h int)) {
	defer qt.Recovering("connect QQuickWindow::setMaximumHeight")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMaximumHeight", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetMaximumHeight() {
	defer qt.Recovering("disconnect QQuickWindow::setMaximumHeight")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMaximumHeight")
	}
}

func (ptr *QQuickWindow) SetMaximumHeight(h int) {
	defer qt.Recovering("QQuickWindow::setMaximumHeight")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QQuickWindow) SetMaximumHeightDefault(h int) {
	defer qt.Recovering("QQuickWindow::setMaximumHeight")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumHeightDefault(ptr.Pointer(), C.int(h))
	}
}

//export callbackQQuickWindow_SetMaximumWidth
func callbackQQuickWindow_SetMaximumWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) {
	defer qt.Recovering("callback QQuickWindow::setMaximumWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMaximumWidth"); signal != nil {
		signal.(func(int))(int(w))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMaximumWidthDefault(int(w))
	}
}

func (ptr *QQuickWindow) ConnectSetMaximumWidth(f func(w int)) {
	defer qt.Recovering("connect QQuickWindow::setMaximumWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMaximumWidth", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetMaximumWidth() {
	defer qt.Recovering("disconnect QQuickWindow::setMaximumWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMaximumWidth")
	}
}

func (ptr *QQuickWindow) SetMaximumWidth(w int) {
	defer qt.Recovering("QQuickWindow::setMaximumWidth")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QQuickWindow) SetMaximumWidthDefault(w int) {
	defer qt.Recovering("QQuickWindow::setMaximumWidth")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMaximumWidthDefault(ptr.Pointer(), C.int(w))
	}
}

//export callbackQQuickWindow_SetMinimumHeight
func callbackQQuickWindow_SetMinimumHeight(ptr unsafe.Pointer, ptrName *C.char, h C.int) {
	defer qt.Recovering("callback QQuickWindow::setMinimumHeight")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMinimumHeight"); signal != nil {
		signal.(func(int))(int(h))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMinimumHeightDefault(int(h))
	}
}

func (ptr *QQuickWindow) ConnectSetMinimumHeight(f func(h int)) {
	defer qt.Recovering("connect QQuickWindow::setMinimumHeight")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMinimumHeight", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetMinimumHeight() {
	defer qt.Recovering("disconnect QQuickWindow::setMinimumHeight")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMinimumHeight")
	}
}

func (ptr *QQuickWindow) SetMinimumHeight(h int) {
	defer qt.Recovering("QQuickWindow::setMinimumHeight")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumHeight(ptr.Pointer(), C.int(h))
	}
}

func (ptr *QQuickWindow) SetMinimumHeightDefault(h int) {
	defer qt.Recovering("QQuickWindow::setMinimumHeight")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumHeightDefault(ptr.Pointer(), C.int(h))
	}
}

//export callbackQQuickWindow_SetMinimumWidth
func callbackQQuickWindow_SetMinimumWidth(ptr unsafe.Pointer, ptrName *C.char, w C.int) {
	defer qt.Recovering("callback QQuickWindow::setMinimumWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "setMinimumWidth"); signal != nil {
		signal.(func(int))(int(w))
	} else {
		NewQQuickWindowFromPointer(ptr).SetMinimumWidthDefault(int(w))
	}
}

func (ptr *QQuickWindow) ConnectSetMinimumWidth(f func(w int)) {
	defer qt.Recovering("connect QQuickWindow::setMinimumWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setMinimumWidth", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetMinimumWidth() {
	defer qt.Recovering("disconnect QQuickWindow::setMinimumWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setMinimumWidth")
	}
}

func (ptr *QQuickWindow) SetMinimumWidth(w int) {
	defer qt.Recovering("QQuickWindow::setMinimumWidth")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumWidth(ptr.Pointer(), C.int(w))
	}
}

func (ptr *QQuickWindow) SetMinimumWidthDefault(w int) {
	defer qt.Recovering("QQuickWindow::setMinimumWidth")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetMinimumWidthDefault(ptr.Pointer(), C.int(w))
	}
}

//export callbackQQuickWindow_SetTitle
func callbackQQuickWindow_SetTitle(ptr unsafe.Pointer, ptrName *C.char, vqs *C.char) {
	defer qt.Recovering("callback QQuickWindow::setTitle")

	if signal := qt.GetSignal(C.GoString(ptrName), "setTitle"); signal != nil {
		signal.(func(string))(C.GoString(vqs))
	} else {
		NewQQuickWindowFromPointer(ptr).SetTitleDefault(C.GoString(vqs))
	}
}

func (ptr *QQuickWindow) ConnectSetTitle(f func(vqs string)) {
	defer qt.Recovering("connect QQuickWindow::setTitle")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setTitle", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetTitle() {
	defer qt.Recovering("disconnect QQuickWindow::setTitle")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setTitle")
	}
}

func (ptr *QQuickWindow) SetTitle(vqs string) {
	defer qt.Recovering("QQuickWindow::setTitle")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetTitle(ptr.Pointer(), C.CString(vqs))
	}
}

func (ptr *QQuickWindow) SetTitleDefault(vqs string) {
	defer qt.Recovering("QQuickWindow::setTitle")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetTitleDefault(ptr.Pointer(), C.CString(vqs))
	}
}

//export callbackQQuickWindow_SetVisible
func callbackQQuickWindow_SetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) {
	defer qt.Recovering("callback QQuickWindow::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
	} else {
		NewQQuickWindowFromPointer(ptr).SetVisibleDefault(int(visible) != 0)
	}
}

func (ptr *QQuickWindow) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QQuickWindow::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QQuickWindow::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

func (ptr *QQuickWindow) SetVisible(visible bool) {
	defer qt.Recovering("QQuickWindow::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QQuickWindow) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QQuickWindow::setVisible")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

//export callbackQQuickWindow_SetWidth
func callbackQQuickWindow_SetWidth(ptr unsafe.Pointer, ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QQuickWindow::setWidth")

	if signal := qt.GetSignal(C.GoString(ptrName), "setWidth"); signal != nil {
		signal.(func(int))(int(arg))
	} else {
		NewQQuickWindowFromPointer(ptr).SetWidthDefault(int(arg))
	}
}

func (ptr *QQuickWindow) ConnectSetWidth(f func(arg int)) {
	defer qt.Recovering("connect QQuickWindow::setWidth")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setWidth", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetWidth() {
	defer qt.Recovering("disconnect QQuickWindow::setWidth")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setWidth")
	}
}

func (ptr *QQuickWindow) SetWidth(arg int) {
	defer qt.Recovering("QQuickWindow::setWidth")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetWidth(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QQuickWindow) SetWidthDefault(arg int) {
	defer qt.Recovering("QQuickWindow::setWidth")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetWidthDefault(ptr.Pointer(), C.int(arg))
	}
}

//export callbackQQuickWindow_SetX
func callbackQQuickWindow_SetX(ptr unsafe.Pointer, ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QQuickWindow::setX")

	if signal := qt.GetSignal(C.GoString(ptrName), "setX"); signal != nil {
		signal.(func(int))(int(arg))
	} else {
		NewQQuickWindowFromPointer(ptr).SetXDefault(int(arg))
	}
}

func (ptr *QQuickWindow) ConnectSetX(f func(arg int)) {
	defer qt.Recovering("connect QQuickWindow::setX")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setX", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetX() {
	defer qt.Recovering("disconnect QQuickWindow::setX")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setX")
	}
}

func (ptr *QQuickWindow) SetX(arg int) {
	defer qt.Recovering("QQuickWindow::setX")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetX(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QQuickWindow) SetXDefault(arg int) {
	defer qt.Recovering("QQuickWindow::setX")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetXDefault(ptr.Pointer(), C.int(arg))
	}
}

//export callbackQQuickWindow_SetY
func callbackQQuickWindow_SetY(ptr unsafe.Pointer, ptrName *C.char, arg C.int) {
	defer qt.Recovering("callback QQuickWindow::setY")

	if signal := qt.GetSignal(C.GoString(ptrName), "setY"); signal != nil {
		signal.(func(int))(int(arg))
	} else {
		NewQQuickWindowFromPointer(ptr).SetYDefault(int(arg))
	}
}

func (ptr *QQuickWindow) ConnectSetY(f func(arg int)) {
	defer qt.Recovering("connect QQuickWindow::setY")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setY", f)
	}
}

func (ptr *QQuickWindow) DisconnectSetY() {
	defer qt.Recovering("disconnect QQuickWindow::setY")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setY")
	}
}

func (ptr *QQuickWindow) SetY(arg int) {
	defer qt.Recovering("QQuickWindow::setY")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetY(ptr.Pointer(), C.int(arg))
	}
}

func (ptr *QQuickWindow) SetYDefault(arg int) {
	defer qt.Recovering("QQuickWindow::setY")

	if ptr.Pointer() != nil {
		C.QQuickWindow_SetYDefault(ptr.Pointer(), C.int(arg))
	}
}

//export callbackQQuickWindow_Alert
func callbackQQuickWindow_Alert(ptr unsafe.Pointer, ptrName *C.char, msec C.int) {
	defer qt.Recovering("callback QQuickWindow::alert")

	if signal := qt.GetSignal(C.GoString(ptrName), "alert"); signal != nil {
		signal.(func(int))(int(msec))
	} else {
		NewQQuickWindowFromPointer(ptr).AlertDefault(int(msec))
	}
}

func (ptr *QQuickWindow) ConnectAlert(f func(msec int)) {
	defer qt.Recovering("connect QQuickWindow::alert")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "alert", f)
	}
}

func (ptr *QQuickWindow) DisconnectAlert() {
	defer qt.Recovering("disconnect QQuickWindow::alert")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "alert")
	}
}

func (ptr *QQuickWindow) Alert(msec int) {
	defer qt.Recovering("QQuickWindow::alert")

	if ptr.Pointer() != nil {
		C.QQuickWindow_Alert(ptr.Pointer(), C.int(msec))
	}
}

func (ptr *QQuickWindow) AlertDefault(msec int) {
	defer qt.Recovering("QQuickWindow::alert")

	if ptr.Pointer() != nil {
		C.QQuickWindow_AlertDefault(ptr.Pointer(), C.int(msec))
	}
}

//export callbackQQuickWindow_Close
func callbackQQuickWindow_Close(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickWindow::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).CloseDefault()))
}

func (ptr *QQuickWindow) ConnectClose(f func() bool) {
	defer qt.Recovering("connect QQuickWindow::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QQuickWindow) DisconnectClose() {
	defer qt.Recovering("disconnect QQuickWindow::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QQuickWindow) Close() bool {
	defer qt.Recovering("QQuickWindow::close")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_Close(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickWindow) CloseDefault() bool {
	defer qt.Recovering("QQuickWindow::close")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_CloseDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQQuickWindow_FocusObject
func callbackQQuickWindow_FocusObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickWindow::focusObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusObject"); signal != nil {
		return core.PointerFromQObject(signal.(func() *core.QObject)())
	}

	return core.PointerFromQObject(NewQQuickWindowFromPointer(ptr).FocusObjectDefault())
}

func (ptr *QQuickWindow) ConnectFocusObject(f func() *core.QObject) {
	defer qt.Recovering("connect QQuickWindow::focusObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusObject", f)
	}
}

func (ptr *QQuickWindow) DisconnectFocusObject() {
	defer qt.Recovering("disconnect QQuickWindow::focusObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusObject")
	}
}

func (ptr *QQuickWindow) FocusObject() *core.QObject {
	defer qt.Recovering("QQuickWindow::focusObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQuickWindow_FocusObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) FocusObjectDefault() *core.QObject {
	defer qt.Recovering("QQuickWindow::focusObject")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQuickWindow_FocusObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWindow_Format
func callbackQQuickWindow_Format(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickWindow::format")

	if signal := qt.GetSignal(C.GoString(ptrName), "format"); signal != nil {
		return gui.PointerFromQSurfaceFormat(signal.(func() *gui.QSurfaceFormat)())
	}

	return gui.PointerFromQSurfaceFormat(NewQQuickWindowFromPointer(ptr).FormatDefault())
}

func (ptr *QQuickWindow) ConnectFormat(f func() *gui.QSurfaceFormat) {
	defer qt.Recovering("connect QQuickWindow::format")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "format", f)
	}
}

func (ptr *QQuickWindow) DisconnectFormat() {
	defer qt.Recovering("disconnect QQuickWindow::format")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "format")
	}
}

func (ptr *QQuickWindow) Format() *gui.QSurfaceFormat {
	defer qt.Recovering("QQuickWindow::format")

	if ptr.Pointer() != nil {
		return gui.NewQSurfaceFormatFromPointer(C.QQuickWindow_Format(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) FormatDefault() *gui.QSurfaceFormat {
	defer qt.Recovering("QQuickWindow::format")

	if ptr.Pointer() != nil {
		return gui.NewQSurfaceFormatFromPointer(C.QQuickWindow_FormatDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWindow_Hide
func callbackQQuickWindow_Hide(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::hide")

	if signal := qt.GetSignal(C.GoString(ptrName), "hide"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).HideDefault()
	}
}

func (ptr *QQuickWindow) ConnectHide(f func()) {
	defer qt.Recovering("connect QQuickWindow::hide")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hide", f)
	}
}

func (ptr *QQuickWindow) DisconnectHide() {
	defer qt.Recovering("disconnect QQuickWindow::hide")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hide")
	}
}

func (ptr *QQuickWindow) Hide() {
	defer qt.Recovering("QQuickWindow::hide")

	if ptr.Pointer() != nil {
		C.QQuickWindow_Hide(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) HideDefault() {
	defer qt.Recovering("QQuickWindow::hide")

	if ptr.Pointer() != nil {
		C.QQuickWindow_HideDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_Lower
func callbackQQuickWindow_Lower(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::lower")

	if signal := qt.GetSignal(C.GoString(ptrName), "lower"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).LowerDefault()
	}
}

func (ptr *QQuickWindow) ConnectLower(f func()) {
	defer qt.Recovering("connect QQuickWindow::lower")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "lower", f)
	}
}

func (ptr *QQuickWindow) DisconnectLower() {
	defer qt.Recovering("disconnect QQuickWindow::lower")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "lower")
	}
}

func (ptr *QQuickWindow) Lower() {
	defer qt.Recovering("QQuickWindow::lower")

	if ptr.Pointer() != nil {
		C.QQuickWindow_Lower(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) LowerDefault() {
	defer qt.Recovering("QQuickWindow::lower")

	if ptr.Pointer() != nil {
		C.QQuickWindow_LowerDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_MoveEvent
func callbackQQuickWindow_MoveEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(ev))
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

//export callbackQQuickWindow_NativeEvent
func callbackQQuickWindow_NativeEvent(ptr unsafe.Pointer, ptrName *C.char, eventType *C.char, message unsafe.Pointer, result C.long) C.int {
	defer qt.Recovering("callback QQuickWindow::nativeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "nativeEvent"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(string, unsafe.Pointer, int) bool)(C.GoString(eventType), message, int(result))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).NativeEventDefault(C.GoString(eventType), message, int(result))))
}

func (ptr *QQuickWindow) ConnectNativeEvent(f func(eventType string, message unsafe.Pointer, result int) bool) {
	defer qt.Recovering("connect QQuickWindow::nativeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "nativeEvent", f)
	}
}

func (ptr *QQuickWindow) DisconnectNativeEvent() {
	defer qt.Recovering("disconnect QQuickWindow::nativeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "nativeEvent")
	}
}

func (ptr *QQuickWindow) NativeEvent(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QQuickWindow::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_NativeEvent(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

func (ptr *QQuickWindow) NativeEventDefault(eventType string, message unsafe.Pointer, result int) bool {
	defer qt.Recovering("QQuickWindow::nativeEvent")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_NativeEventDefault(ptr.Pointer(), C.CString(eventType), message, C.long(result)) != 0
	}
	return false
}

//export callbackQQuickWindow_Raise
func callbackQQuickWindow_Raise(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::raise")

	if signal := qt.GetSignal(C.GoString(ptrName), "raise"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).RaiseDefault()
	}
}

func (ptr *QQuickWindow) ConnectRaise(f func()) {
	defer qt.Recovering("connect QQuickWindow::raise")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "raise", f)
	}
}

func (ptr *QQuickWindow) DisconnectRaise() {
	defer qt.Recovering("disconnect QQuickWindow::raise")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "raise")
	}
}

func (ptr *QQuickWindow) Raise() {
	defer qt.Recovering("QQuickWindow::raise")

	if ptr.Pointer() != nil {
		C.QQuickWindow_Raise(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RaiseDefault() {
	defer qt.Recovering("QQuickWindow::raise")

	if ptr.Pointer() != nil {
		C.QQuickWindow_RaiseDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_RequestActivate
func callbackQQuickWindow_RequestActivate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::requestActivate")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestActivate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).RequestActivateDefault()
	}
}

func (ptr *QQuickWindow) ConnectRequestActivate(f func()) {
	defer qt.Recovering("connect QQuickWindow::requestActivate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "requestActivate", f)
	}
}

func (ptr *QQuickWindow) DisconnectRequestActivate() {
	defer qt.Recovering("disconnect QQuickWindow::requestActivate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "requestActivate")
	}
}

func (ptr *QQuickWindow) RequestActivate() {
	defer qt.Recovering("QQuickWindow::requestActivate")

	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestActivate(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RequestActivateDefault() {
	defer qt.Recovering("QQuickWindow::requestActivate")

	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestActivateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_RequestUpdate
func callbackQQuickWindow_RequestUpdate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::requestUpdate")

	if signal := qt.GetSignal(C.GoString(ptrName), "requestUpdate"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).RequestUpdateDefault()
	}
}

func (ptr *QQuickWindow) ConnectRequestUpdate(f func()) {
	defer qt.Recovering("connect QQuickWindow::requestUpdate")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "requestUpdate", f)
	}
}

func (ptr *QQuickWindow) DisconnectRequestUpdate() {
	defer qt.Recovering("disconnect QQuickWindow::requestUpdate")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "requestUpdate")
	}
}

func (ptr *QQuickWindow) RequestUpdate() {
	defer qt.Recovering("QQuickWindow::requestUpdate")

	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestUpdate(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) RequestUpdateDefault() {
	defer qt.Recovering("QQuickWindow::requestUpdate")

	if ptr.Pointer() != nil {
		C.QQuickWindow_RequestUpdateDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_Show
func callbackQQuickWindow_Show(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::show")

	if signal := qt.GetSignal(C.GoString(ptrName), "show"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowDefault()
	}
}

func (ptr *QQuickWindow) ConnectShow(f func()) {
	defer qt.Recovering("connect QQuickWindow::show")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "show", f)
	}
}

func (ptr *QQuickWindow) DisconnectShow() {
	defer qt.Recovering("disconnect QQuickWindow::show")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "show")
	}
}

func (ptr *QQuickWindow) Show() {
	defer qt.Recovering("QQuickWindow::show")

	if ptr.Pointer() != nil {
		C.QQuickWindow_Show(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowDefault() {
	defer qt.Recovering("QQuickWindow::show")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowFullScreen
func callbackQQuickWindow_ShowFullScreen(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::showFullScreen")

	if signal := qt.GetSignal(C.GoString(ptrName), "showFullScreen"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowFullScreenDefault()
	}
}

func (ptr *QQuickWindow) ConnectShowFullScreen(f func()) {
	defer qt.Recovering("connect QQuickWindow::showFullScreen")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showFullScreen", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowFullScreen() {
	defer qt.Recovering("disconnect QQuickWindow::showFullScreen")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showFullScreen")
	}
}

func (ptr *QQuickWindow) ShowFullScreen() {
	defer qt.Recovering("QQuickWindow::showFullScreen")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowFullScreen(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowFullScreenDefault() {
	defer qt.Recovering("QQuickWindow::showFullScreen")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowFullScreenDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowMaximized
func callbackQQuickWindow_ShowMaximized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::showMaximized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMaximized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowMaximizedDefault()
	}
}

func (ptr *QQuickWindow) ConnectShowMaximized(f func()) {
	defer qt.Recovering("connect QQuickWindow::showMaximized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMaximized", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowMaximized() {
	defer qt.Recovering("disconnect QQuickWindow::showMaximized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMaximized")
	}
}

func (ptr *QQuickWindow) ShowMaximized() {
	defer qt.Recovering("QQuickWindow::showMaximized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMaximized(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowMaximizedDefault() {
	defer qt.Recovering("QQuickWindow::showMaximized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMaximizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowMinimized
func callbackQQuickWindow_ShowMinimized(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::showMinimized")

	if signal := qt.GetSignal(C.GoString(ptrName), "showMinimized"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowMinimizedDefault()
	}
}

func (ptr *QQuickWindow) ConnectShowMinimized(f func()) {
	defer qt.Recovering("connect QQuickWindow::showMinimized")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showMinimized", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowMinimized() {
	defer qt.Recovering("disconnect QQuickWindow::showMinimized")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showMinimized")
	}
}

func (ptr *QQuickWindow) ShowMinimized() {
	defer qt.Recovering("QQuickWindow::showMinimized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMinimized(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowMinimizedDefault() {
	defer qt.Recovering("QQuickWindow::showMinimized")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowMinimizedDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_ShowNormal
func callbackQQuickWindow_ShowNormal(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::showNormal")

	if signal := qt.GetSignal(C.GoString(ptrName), "showNormal"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).ShowNormalDefault()
	}
}

func (ptr *QQuickWindow) ConnectShowNormal(f func()) {
	defer qt.Recovering("connect QQuickWindow::showNormal")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showNormal", f)
	}
}

func (ptr *QQuickWindow) DisconnectShowNormal() {
	defer qt.Recovering("disconnect QQuickWindow::showNormal")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showNormal")
	}
}

func (ptr *QQuickWindow) ShowNormal() {
	defer qt.Recovering("QQuickWindow::showNormal")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowNormal(ptr.Pointer())
	}
}

func (ptr *QQuickWindow) ShowNormalDefault() {
	defer qt.Recovering("QQuickWindow::showNormal")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ShowNormalDefault(ptr.Pointer())
	}
}

//export callbackQQuickWindow_Size
func callbackQQuickWindow_Size(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickWindow::size")

	if signal := qt.GetSignal(C.GoString(ptrName), "size"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(NewQQuickWindowFromPointer(ptr).SizeDefault())
}

func (ptr *QQuickWindow) ConnectSize(f func() *core.QSize) {
	defer qt.Recovering("connect QQuickWindow::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "size", f)
	}
}

func (ptr *QQuickWindow) DisconnectSize() {
	defer qt.Recovering("disconnect QQuickWindow::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "size")
	}
}

func (ptr *QQuickWindow) Size() *core.QSize {
	defer qt.Recovering("QQuickWindow::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWindow_Size(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) SizeDefault() *core.QSize {
	defer qt.Recovering("QQuickWindow::size")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QQuickWindow_SizeDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQQuickWindow_SurfaceType
func callbackQQuickWindow_SurfaceType(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QQuickWindow::surfaceType")

	if signal := qt.GetSignal(C.GoString(ptrName), "surfaceType"); signal != nil {
		return C.int(signal.(func() gui.QSurface__SurfaceType)())
	}

	return C.int(NewQQuickWindowFromPointer(ptr).SurfaceTypeDefault())
}

func (ptr *QQuickWindow) ConnectSurfaceType(f func() gui.QSurface__SurfaceType) {
	defer qt.Recovering("connect QQuickWindow::surfaceType")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "surfaceType", f)
	}
}

func (ptr *QQuickWindow) DisconnectSurfaceType() {
	defer qt.Recovering("disconnect QQuickWindow::surfaceType")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "surfaceType")
	}
}

func (ptr *QQuickWindow) SurfaceType() gui.QSurface__SurfaceType {
	defer qt.Recovering("QQuickWindow::surfaceType")

	if ptr.Pointer() != nil {
		return gui.QSurface__SurfaceType(C.QQuickWindow_SurfaceType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickWindow) SurfaceTypeDefault() gui.QSurface__SurfaceType {
	defer qt.Recovering("QQuickWindow::surfaceType")

	if ptr.Pointer() != nil {
		return gui.QSurface__SurfaceType(C.QQuickWindow_SurfaceTypeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQQuickWindow_TabletEvent
func callbackQQuickWindow_TabletEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(ev))
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

//export callbackQQuickWindow_TouchEvent
func callbackQQuickWindow_TouchEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(ev))
	} else {
		NewQQuickWindowFromPointer(ptr).TouchEventDefault(gui.NewQTouchEventFromPointer(ev))
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

//export callbackQQuickWindow_TimerEvent
func callbackQQuickWindow_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQQuickWindow_ChildEvent
func callbackQQuickWindow_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQQuickWindow_ConnectNotify
func callbackQQuickWindow_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWindowFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWindow) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickWindow::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QQuickWindow) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QQuickWindow::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QQuickWindow) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWindow::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWindow) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWindow::connectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWindow_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWindow_CustomEvent
func callbackQQuickWindow_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQQuickWindowFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQQuickWindow_DeleteLater
func callbackQQuickWindow_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QQuickWindow::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQQuickWindowFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QQuickWindow) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QQuickWindow::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QQuickWindow) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QQuickWindow::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QQuickWindow) DeleteLater() {
	defer qt.Recovering("QQuickWindow::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickWindow_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QQuickWindow) DeleteLaterDefault() {
	defer qt.Recovering("QQuickWindow::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QQuickWindow_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQQuickWindow_DisconnectNotify
func callbackQQuickWindow_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QQuickWindow::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQQuickWindowFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QQuickWindow) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QQuickWindow::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QQuickWindow) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QQuickWindow::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QQuickWindow) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWindow::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QQuickWindow) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QQuickWindow::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QQuickWindow_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQQuickWindow_EventFilter
func callbackQQuickWindow_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QQuickWindow::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQQuickWindowFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QQuickWindow) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QQuickWindow::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QQuickWindow) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QQuickWindow::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QQuickWindow) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWindow::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QQuickWindow) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QQuickWindow::eventFilter")

	if ptr.Pointer() != nil {
		return C.QQuickWindow_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQQuickWindow_MetaObject
func callbackQQuickWindow_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QQuickWindow::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQQuickWindowFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QQuickWindow) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QQuickWindow::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QQuickWindow) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QQuickWindow::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QQuickWindow) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QQuickWindow::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWindow_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickWindow) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QQuickWindow::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQuickWindow_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QSGAbstractRenderer::ClearModeBit
type QSGAbstractRenderer__ClearModeBit int64

const (
	QSGAbstractRenderer__ClearColorBuffer   = QSGAbstractRenderer__ClearModeBit(0x0001)
	QSGAbstractRenderer__ClearDepthBuffer   = QSGAbstractRenderer__ClearModeBit(0x0002)
	QSGAbstractRenderer__ClearStencilBuffer = QSGAbstractRenderer__ClearModeBit(0x0004)
)

type QSGAbstractRenderer struct {
	core.QObject
}

type QSGAbstractRenderer_ITF interface {
	core.QObject_ITF
	QSGAbstractRenderer_PTR() *QSGAbstractRenderer
}

func (p *QSGAbstractRenderer) QSGAbstractRenderer_PTR() *QSGAbstractRenderer {
	return p
}

func (p *QSGAbstractRenderer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSGAbstractRenderer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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

//export callbackQSGAbstractRenderer_SceneGraphChanged
func callbackQSGAbstractRenderer_SceneGraphChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGAbstractRenderer::sceneGraphChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "sceneGraphChanged"); signal != nil {
		signal.(func())()
	}

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

//export callbackQSGAbstractRenderer_TimerEvent
func callbackQSGAbstractRenderer_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGAbstractRenderer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
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

//export callbackQSGAbstractRenderer_ChildEvent
func callbackQSGAbstractRenderer_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGAbstractRenderer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQSGAbstractRenderer_ConnectNotify
func callbackQSGAbstractRenderer_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGAbstractRenderer::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGAbstractRenderer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGAbstractRenderer::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QSGAbstractRenderer) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGAbstractRenderer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGAbstractRenderer_CustomEvent
func callbackQSGAbstractRenderer_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGAbstractRenderer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQSGAbstractRenderer_DeleteLater
func callbackQSGAbstractRenderer_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGAbstractRenderer::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGAbstractRendererFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGAbstractRenderer) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSGAbstractRenderer::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QSGAbstractRenderer) DeleteLater() {
	defer qt.Recovering("QSGAbstractRenderer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSGAbstractRenderer_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGAbstractRenderer) DeleteLaterDefault() {
	defer qt.Recovering("QSGAbstractRenderer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSGAbstractRenderer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGAbstractRenderer_DisconnectNotify
func callbackQSGAbstractRenderer_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGAbstractRenderer::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGAbstractRendererFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGAbstractRenderer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGAbstractRenderer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QSGAbstractRenderer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGAbstractRenderer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGAbstractRenderer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGAbstractRenderer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGAbstractRenderer_Event
func callbackQSGAbstractRenderer_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGAbstractRenderer::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQSGAbstractRendererFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QSGAbstractRenderer) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSGAbstractRenderer::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectEvent() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QSGAbstractRenderer) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGAbstractRenderer::event")

	if ptr.Pointer() != nil {
		return C.QSGAbstractRenderer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGAbstractRenderer) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGAbstractRenderer::event")

	if ptr.Pointer() != nil {
		return C.QSGAbstractRenderer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGAbstractRenderer_EventFilter
func callbackQSGAbstractRenderer_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGAbstractRenderer::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQSGAbstractRendererFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QSGAbstractRenderer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSGAbstractRenderer::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QSGAbstractRenderer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGAbstractRenderer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGAbstractRenderer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGAbstractRenderer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGAbstractRenderer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGAbstractRenderer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGAbstractRenderer_MetaObject
func callbackQSGAbstractRenderer_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGAbstractRenderer::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGAbstractRendererFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGAbstractRenderer) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSGAbstractRenderer::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QSGAbstractRenderer) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSGAbstractRenderer::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QSGAbstractRenderer) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSGAbstractRenderer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGAbstractRenderer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGAbstractRenderer) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSGAbstractRenderer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGAbstractRenderer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSGBasicGeometryNode struct {
	QSGNode
}

type QSGBasicGeometryNode_ITF interface {
	QSGNode_ITF
	QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode
}

func (p *QSGBasicGeometryNode) QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode {
	return p
}

func (p *QSGBasicGeometryNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGBasicGeometryNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGNode_PTR().SetPointer(ptr)
	}
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
		ptr.SetPointer(nil)
	}
}

//export callbackQSGBasicGeometryNode_IsSubtreeBlocked
func callbackQSGBasicGeometryNode_IsSubtreeBlocked(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGBasicGeometryNode::isSubtreeBlocked")

	if signal := qt.GetSignal(C.GoString(ptrName), "isSubtreeBlocked"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGBasicGeometryNodeFromPointer(ptr).IsSubtreeBlockedDefault()))
}

func (ptr *QSGBasicGeometryNode) ConnectIsSubtreeBlocked(f func() bool) {
	defer qt.Recovering("connect QSGBasicGeometryNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked", f)
	}
}

func (ptr *QSGBasicGeometryNode) DisconnectIsSubtreeBlocked() {
	defer qt.Recovering("disconnect QSGBasicGeometryNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked")
	}
}

func (ptr *QSGBasicGeometryNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGBasicGeometryNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGBasicGeometryNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGBasicGeometryNode) IsSubtreeBlockedDefault() bool {
	defer qt.Recovering("QSGBasicGeometryNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGBasicGeometryNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGBasicGeometryNode_Preprocess
func callbackQSGBasicGeometryNode_Preprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGBasicGeometryNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGBasicGeometryNodeFromPointer(ptr).PreprocessDefault()
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

func (p *QSGClipNode) QSGClipNode_PTR() *QSGClipNode {
	return p
}

func (p *QSGClipNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGClipNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGBasicGeometryNode_PTR().SetPointer(ptr)
	}
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
		ptr.SetPointer(nil)
	}
}

//export callbackQSGClipNode_IsSubtreeBlocked
func callbackQSGClipNode_IsSubtreeBlocked(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGClipNode::isSubtreeBlocked")

	if signal := qt.GetSignal(C.GoString(ptrName), "isSubtreeBlocked"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGClipNodeFromPointer(ptr).IsSubtreeBlockedDefault()))
}

func (ptr *QSGClipNode) ConnectIsSubtreeBlocked(f func() bool) {
	defer qt.Recovering("connect QSGClipNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked", f)
	}
}

func (ptr *QSGClipNode) DisconnectIsSubtreeBlocked() {
	defer qt.Recovering("disconnect QSGClipNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked")
	}
}

func (ptr *QSGClipNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGClipNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGClipNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGClipNode) IsSubtreeBlockedDefault() bool {
	defer qt.Recovering("QSGClipNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGClipNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGClipNode_Preprocess
func callbackQSGClipNode_Preprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGClipNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGClipNodeFromPointer(ptr).PreprocessDefault()
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

func (p *QSGDynamicTexture) QSGDynamicTexture_PTR() *QSGDynamicTexture {
	return p
}

func (p *QSGDynamicTexture) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGTexture_PTR().Pointer()
	}
	return nil
}

func (p *QSGDynamicTexture) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGTexture_PTR().SetPointer(ptr)
	}
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

//export callbackQSGDynamicTexture_UpdateTexture
func callbackQSGDynamicTexture_UpdateTexture(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGDynamicTexture::updateTexture")

	if signal := qt.GetSignal(C.GoString(ptrName), "updateTexture"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QSGDynamicTexture) ConnectUpdateTexture(f func() bool) {
	defer qt.Recovering("connect QSGDynamicTexture::updateTexture")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "updateTexture", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectUpdateTexture() {
	defer qt.Recovering("disconnect QSGDynamicTexture::updateTexture")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "updateTexture")
	}
}

func (ptr *QSGDynamicTexture) UpdateTexture() bool {
	defer qt.Recovering("QSGDynamicTexture::updateTexture")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_UpdateTexture(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_Bind
func callbackQSGDynamicTexture_Bind(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGDynamicTexture::bind")

	if signal := qt.GetSignal(C.GoString(ptrName), "bind"); signal != nil {
		signal.(func())()
	} else {

	}
}

func (ptr *QSGDynamicTexture) ConnectBind(f func()) {
	defer qt.Recovering("connect QSGDynamicTexture::bind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "bind", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectBind() {
	defer qt.Recovering("disconnect QSGDynamicTexture::bind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "bind")
	}
}

func (ptr *QSGDynamicTexture) Bind() {
	defer qt.Recovering("QSGDynamicTexture::bind")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_Bind(ptr.Pointer())
	}
}

//export callbackQSGDynamicTexture_HasAlphaChannel
func callbackQSGDynamicTexture_HasAlphaChannel(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGDynamicTexture::hasAlphaChannel")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasAlphaChannel"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QSGDynamicTexture) ConnectHasAlphaChannel(f func() bool) {
	defer qt.Recovering("connect QSGDynamicTexture::hasAlphaChannel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasAlphaChannel", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectHasAlphaChannel() {
	defer qt.Recovering("disconnect QSGDynamicTexture::hasAlphaChannel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasAlphaChannel")
	}
}

func (ptr *QSGDynamicTexture) HasAlphaChannel() bool {
	defer qt.Recovering("QSGDynamicTexture::hasAlphaChannel")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_HasAlphaChannel(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_HasMipmaps
func callbackQSGDynamicTexture_HasMipmaps(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGDynamicTexture::hasMipmaps")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasMipmaps"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QSGDynamicTexture) ConnectHasMipmaps(f func() bool) {
	defer qt.Recovering("connect QSGDynamicTexture::hasMipmaps")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasMipmaps", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectHasMipmaps() {
	defer qt.Recovering("disconnect QSGDynamicTexture::hasMipmaps")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasMipmaps")
	}
}

func (ptr *QSGDynamicTexture) HasMipmaps() bool {
	defer qt.Recovering("QSGDynamicTexture::hasMipmaps")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_HasMipmaps(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_IsAtlasTexture
func callbackQSGDynamicTexture_IsAtlasTexture(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGDynamicTexture::isAtlasTexture")

	if signal := qt.GetSignal(C.GoString(ptrName), "isAtlasTexture"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGDynamicTextureFromPointer(ptr).IsAtlasTextureDefault()))
}

func (ptr *QSGDynamicTexture) ConnectIsAtlasTexture(f func() bool) {
	defer qt.Recovering("connect QSGDynamicTexture::isAtlasTexture")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isAtlasTexture", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectIsAtlasTexture() {
	defer qt.Recovering("disconnect QSGDynamicTexture::isAtlasTexture")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isAtlasTexture")
	}
}

func (ptr *QSGDynamicTexture) IsAtlasTexture() bool {
	defer qt.Recovering("QSGDynamicTexture::isAtlasTexture")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_IsAtlasTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) IsAtlasTextureDefault() bool {
	defer qt.Recovering("QSGDynamicTexture::isAtlasTexture")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_IsAtlasTextureDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_NormalizedTextureSubRect
func callbackQSGDynamicTexture_NormalizedTextureSubRect(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGDynamicTexture::normalizedTextureSubRect")

	if signal := qt.GetSignal(C.GoString(ptrName), "normalizedTextureSubRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQSGDynamicTextureFromPointer(ptr).NormalizedTextureSubRectDefault())
}

func (ptr *QSGDynamicTexture) ConnectNormalizedTextureSubRect(f func() *core.QRectF) {
	defer qt.Recovering("connect QSGDynamicTexture::normalizedTextureSubRect")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "normalizedTextureSubRect", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectNormalizedTextureSubRect() {
	defer qt.Recovering("disconnect QSGDynamicTexture::normalizedTextureSubRect")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "normalizedTextureSubRect")
	}
}

func (ptr *QSGDynamicTexture) NormalizedTextureSubRect() *core.QRectF {
	defer qt.Recovering("QSGDynamicTexture::normalizedTextureSubRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGDynamicTexture_NormalizedTextureSubRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGDynamicTexture) NormalizedTextureSubRectDefault() *core.QRectF {
	defer qt.Recovering("QSGDynamicTexture::normalizedTextureSubRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGDynamicTexture_NormalizedTextureSubRectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGDynamicTexture_RemovedFromAtlas
func callbackQSGDynamicTexture_RemovedFromAtlas(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGDynamicTexture::removedFromAtlas")

	if signal := qt.GetSignal(C.GoString(ptrName), "removedFromAtlas"); signal != nil {
		return PointerFromQSGTexture(signal.(func() *QSGTexture)())
	}

	return PointerFromQSGTexture(NewQSGDynamicTextureFromPointer(ptr).RemovedFromAtlasDefault())
}

func (ptr *QSGDynamicTexture) ConnectRemovedFromAtlas(f func() *QSGTexture) {
	defer qt.Recovering("connect QSGDynamicTexture::removedFromAtlas")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "removedFromAtlas", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectRemovedFromAtlas() {
	defer qt.Recovering("disconnect QSGDynamicTexture::removedFromAtlas")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "removedFromAtlas")
	}
}

func (ptr *QSGDynamicTexture) RemovedFromAtlas() *QSGTexture {
	defer qt.Recovering("QSGDynamicTexture::removedFromAtlas")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGDynamicTexture_RemovedFromAtlas(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGDynamicTexture) RemovedFromAtlasDefault() *QSGTexture {
	defer qt.Recovering("QSGDynamicTexture::removedFromAtlas")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGDynamicTexture_RemovedFromAtlasDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGDynamicTexture_TextureId
func callbackQSGDynamicTexture_TextureId(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGDynamicTexture::textureId")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureId"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(0)
}

func (ptr *QSGDynamicTexture) ConnectTextureId(f func() int) {
	defer qt.Recovering("connect QSGDynamicTexture::textureId")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureId", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectTextureId() {
	defer qt.Recovering("disconnect QSGDynamicTexture::textureId")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureId")
	}
}

func (ptr *QSGDynamicTexture) TextureId() int {
	defer qt.Recovering("QSGDynamicTexture::textureId")

	if ptr.Pointer() != nil {
		return int(C.QSGDynamicTexture_TextureId(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGDynamicTexture_TextureSize
func callbackQSGDynamicTexture_TextureSize(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGDynamicTexture::textureSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureSize"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(nil)
}

func (ptr *QSGDynamicTexture) ConnectTextureSize(f func() *core.QSize) {
	defer qt.Recovering("connect QSGDynamicTexture::textureSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureSize", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectTextureSize() {
	defer qt.Recovering("disconnect QSGDynamicTexture::textureSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureSize")
	}
}

func (ptr *QSGDynamicTexture) TextureSize() *core.QSize {
	defer qt.Recovering("QSGDynamicTexture::textureSize")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QSGDynamicTexture_TextureSize(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGDynamicTexture_TimerEvent
func callbackQSGDynamicTexture_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
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

//export callbackQSGDynamicTexture_ChildEvent
func callbackQSGDynamicTexture_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQSGDynamicTexture_ConnectNotify
func callbackQSGDynamicTexture_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGDynamicTexture) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGDynamicTexture::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSGDynamicTexture::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QSGDynamicTexture) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGDynamicTexture::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGDynamicTexture) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGDynamicTexture::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGDynamicTexture_CustomEvent
func callbackQSGDynamicTexture_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQSGDynamicTexture_DeleteLater
func callbackQSGDynamicTexture_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGDynamicTexture::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGDynamicTextureFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGDynamicTexture) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSGDynamicTexture::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSGDynamicTexture::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QSGDynamicTexture) DeleteLater() {
	defer qt.Recovering("QSGDynamicTexture::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSGDynamicTexture_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGDynamicTexture) DeleteLaterDefault() {
	defer qt.Recovering("QSGDynamicTexture::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSGDynamicTexture_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGDynamicTexture_DisconnectNotify
func callbackQSGDynamicTexture_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGDynamicTexture::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGDynamicTextureFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGDynamicTexture) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGDynamicTexture::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSGDynamicTexture::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QSGDynamicTexture) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGDynamicTexture::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGDynamicTexture) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGDynamicTexture::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGDynamicTexture_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGDynamicTexture_Event
func callbackQSGDynamicTexture_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGDynamicTexture::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQSGDynamicTextureFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QSGDynamicTexture) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSGDynamicTexture::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectEvent() {
	defer qt.Recovering("disconnect QSGDynamicTexture::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QSGDynamicTexture) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGDynamicTexture::event")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGDynamicTexture::event")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_EventFilter
func callbackQSGDynamicTexture_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGDynamicTexture::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQSGDynamicTextureFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QSGDynamicTexture) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSGDynamicTexture::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSGDynamicTexture::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QSGDynamicTexture) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGDynamicTexture::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGDynamicTexture) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGDynamicTexture::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGDynamicTexture_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGDynamicTexture_MetaObject
func callbackQSGDynamicTexture_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGDynamicTexture::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGDynamicTextureFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGDynamicTexture) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSGDynamicTexture::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QSGDynamicTexture) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSGDynamicTexture::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QSGDynamicTexture) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSGDynamicTexture::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGDynamicTexture_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGDynamicTexture) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSGDynamicTexture::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGDynamicTexture_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QSGEngine::CreateTextureOption
type QSGEngine__CreateTextureOption int64

const (
	QSGEngine__TextureHasAlphaChannel = QSGEngine__CreateTextureOption(0x0001)
	QSGEngine__TextureOwnsGLTexture   = QSGEngine__CreateTextureOption(0x0004)
	QSGEngine__TextureCanUseAtlas     = QSGEngine__CreateTextureOption(0x0008)
	QSGEngine__TextureIsOpaque        = QSGEngine__CreateTextureOption(0x0010)
)

type QSGEngine struct {
	core.QObject
}

type QSGEngine_ITF interface {
	core.QObject_ITF
	QSGEngine_PTR() *QSGEngine
}

func (p *QSGEngine) QSGEngine_PTR() *QSGEngine {
	return p
}

func (p *QSGEngine) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSGEngine) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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

//export callbackQSGEngine_TimerEvent
func callbackQSGEngine_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQSGEngine_ChildEvent
func callbackQSGEngine_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQSGEngine_ConnectNotify
func callbackQSGEngine_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGEngine) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QSGEngine) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSGEngine::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QSGEngine) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGEngine_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGEngine::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGEngine_CustomEvent
func callbackQSGEngine_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQSGEngine_DeleteLater
func callbackQSGEngine_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGEngine::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGEngine) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSGEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QSGEngine) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSGEngine::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QSGEngine) DeleteLater() {
	defer qt.Recovering("QSGEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QSGEngine_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGEngine) DeleteLaterDefault() {
	defer qt.Recovering("QSGEngine::deleteLater")

	if ptr.Pointer() != nil {
		C.QSGEngine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGEngine_DisconnectNotify
func callbackQSGEngine_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGEngine::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGEngine) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QSGEngine) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSGEngine::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QSGEngine) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGEngine_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGEngine::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGEngine_Event
func callbackQSGEngine_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGEngine::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQSGEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QSGEngine) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSGEngine::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QSGEngine) DisconnectEvent() {
	defer qt.Recovering("disconnect QSGEngine::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QSGEngine) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGEngine::event")

	if ptr.Pointer() != nil {
		return C.QSGEngine_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGEngine) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGEngine::event")

	if ptr.Pointer() != nil {
		return C.QSGEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGEngine_EventFilter
func callbackQSGEngine_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGEngine::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQSGEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QSGEngine) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSGEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QSGEngine) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSGEngine::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QSGEngine) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGEngine_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGEngine::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGEngine_MetaObject
func callbackQSGEngine_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGEngine::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGEngine) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSGEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QSGEngine) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSGEngine::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QSGEngine) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSGEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGEngine_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGEngine) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSGEngine::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSGFlatColorMaterial struct {
	QSGMaterial
}

type QSGFlatColorMaterial_ITF interface {
	QSGMaterial_ITF
	QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial
}

func (p *QSGFlatColorMaterial) QSGFlatColorMaterial_PTR() *QSGFlatColorMaterial {
	return p
}

func (p *QSGFlatColorMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGFlatColorMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterial_PTR().SetPointer(ptr)
	}
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

func NewQSGFlatColorMaterial() *QSGFlatColorMaterial {
	defer qt.Recovering("QSGFlatColorMaterial::QSGFlatColorMaterial")

	return newQSGFlatColorMaterialFromPointer(C.QSGFlatColorMaterial_NewQSGFlatColorMaterial())
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

//export callbackQSGFlatColorMaterial_Compare
func callbackQSGFlatColorMaterial_Compare(ptr unsafe.Pointer, ptrName *C.char, other unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGFlatColorMaterial::compare")

	if signal := qt.GetSignal(C.GoString(ptrName), "compare"); signal != nil {
		return C.int(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other)))
	}

	return C.int(NewQSGFlatColorMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other)))
}

func (ptr *QSGFlatColorMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	defer qt.Recovering("connect QSGFlatColorMaterial::compare")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "compare", f)
	}
}

func (ptr *QSGFlatColorMaterial) DisconnectCompare() {
	defer qt.Recovering("disconnect QSGFlatColorMaterial::compare")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "compare")
	}
}

func (ptr *QSGFlatColorMaterial) Compare(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGFlatColorMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGFlatColorMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

func (ptr *QSGFlatColorMaterial) CompareDefault(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGFlatColorMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGFlatColorMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

//export callbackQSGFlatColorMaterial_CreateShader
func callbackQSGFlatColorMaterial_CreateShader(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGFlatColorMaterial::createShader")

	if signal := qt.GetSignal(C.GoString(ptrName), "createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGFlatColorMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGFlatColorMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	defer qt.Recovering("connect QSGFlatColorMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "createShader", f)
	}
}

func (ptr *QSGFlatColorMaterial) DisconnectCreateShader() {
	defer qt.Recovering("disconnect QSGFlatColorMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "createShader")
	}
}

func (ptr *QSGFlatColorMaterial) CreateShader() *QSGMaterialShader {
	defer qt.Recovering("QSGFlatColorMaterial::createShader")

	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGFlatColorMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGFlatColorMaterial) CreateShaderDefault() *QSGMaterialShader {
	defer qt.Recovering("QSGFlatColorMaterial::createShader")

	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGFlatColorMaterial_CreateShaderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGFlatColorMaterial_Type
func callbackQSGFlatColorMaterial_Type(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGFlatColorMaterial::type")

	if signal := qt.GetSignal(C.GoString(ptrName), "type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGFlatColorMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGFlatColorMaterial) ConnectType(f func() *QSGMaterialType) {
	defer qt.Recovering("connect QSGFlatColorMaterial::type")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "type", f)
	}
}

func (ptr *QSGFlatColorMaterial) DisconnectType() {
	defer qt.Recovering("disconnect QSGFlatColorMaterial::type")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "type")
	}
}

func (ptr *QSGFlatColorMaterial) Type() *QSGMaterialType {
	defer qt.Recovering("QSGFlatColorMaterial::type")

	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGFlatColorMaterial_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGFlatColorMaterial) TypeDefault() *QSGMaterialType {
	defer qt.Recovering("QSGFlatColorMaterial::type")

	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGFlatColorMaterial_TypeDefault(ptr.Pointer()))
	}
	return nil
}

//QSGGeometry::DataPattern
type QSGGeometry__DataPattern int64

const (
	QSGGeometry__AlwaysUploadPattern = QSGGeometry__DataPattern(0)
	QSGGeometry__StreamPattern       = QSGGeometry__DataPattern(1)
	QSGGeometry__DynamicPattern      = QSGGeometry__DataPattern(2)
	QSGGeometry__StaticPattern       = QSGGeometry__DataPattern(3)
)

type QSGGeometry struct {
	ptr unsafe.Pointer
}

type QSGGeometry_ITF interface {
	QSGGeometry_PTR() *QSGGeometry
}

func (p *QSGGeometry) QSGGeometry_PTR() *QSGGeometry {
	return p
}

func (p *QSGGeometry) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGGeometry) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
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

func (ptr *QSGGeometry) UpdateRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF) {
	defer qt.Recovering("QSGGeometry::updateRectGeometry")

	C.QSGGeometry_QSGGeometry_UpdateRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect))
}

func QSGGeometry_UpdateTexturedRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF, textureRect core.QRectF_ITF) {
	defer qt.Recovering("QSGGeometry::updateTexturedRectGeometry")

	C.QSGGeometry_QSGGeometry_UpdateTexturedRectGeometry(PointerFromQSGGeometry(g), core.PointerFromQRectF(rect), core.PointerFromQRectF(textureRect))
}

func (ptr *QSGGeometry) UpdateTexturedRectGeometry(g QSGGeometry_ITF, rect core.QRectF_ITF, textureRect core.QRectF_ITF) {
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
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QSGGeometry_DestroyQSGGeometry(ptr.Pointer())
		ptr.SetPointer(nil)
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

func (p *QSGGeometryNode) QSGGeometryNode_PTR() *QSGGeometryNode {
	return p
}

func (p *QSGGeometryNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGBasicGeometryNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGGeometryNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGBasicGeometryNode_PTR().SetPointer(ptr)
	}
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
		ptr.SetPointer(nil)
	}
}

//export callbackQSGGeometryNode_IsSubtreeBlocked
func callbackQSGGeometryNode_IsSubtreeBlocked(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGGeometryNode::isSubtreeBlocked")

	if signal := qt.GetSignal(C.GoString(ptrName), "isSubtreeBlocked"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGGeometryNodeFromPointer(ptr).IsSubtreeBlockedDefault()))
}

func (ptr *QSGGeometryNode) ConnectIsSubtreeBlocked(f func() bool) {
	defer qt.Recovering("connect QSGGeometryNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked", f)
	}
}

func (ptr *QSGGeometryNode) DisconnectIsSubtreeBlocked() {
	defer qt.Recovering("disconnect QSGGeometryNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked")
	}
}

func (ptr *QSGGeometryNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGGeometryNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGGeometryNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGGeometryNode) IsSubtreeBlockedDefault() bool {
	defer qt.Recovering("QSGGeometryNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGGeometryNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGGeometryNode_Preprocess
func callbackQSGGeometryNode_Preprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGGeometryNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGGeometryNodeFromPointer(ptr).PreprocessDefault()
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

//QSGMaterial::Flag
type QSGMaterial__Flag int64

const (
	QSGMaterial__Blending                          = QSGMaterial__Flag(0x0001)
	QSGMaterial__RequiresDeterminant               = QSGMaterial__Flag(0x0002)
	QSGMaterial__RequiresFullMatrixExceptTranslate = QSGMaterial__Flag(0x0004 | QSGMaterial__RequiresDeterminant)
	QSGMaterial__RequiresFullMatrix                = QSGMaterial__Flag(0x0008 | QSGMaterial__RequiresFullMatrixExceptTranslate)
	QSGMaterial__CustomCompileStep                 = QSGMaterial__Flag(0x0010)
)

type QSGMaterial struct {
	ptr unsafe.Pointer
}

type QSGMaterial_ITF interface {
	QSGMaterial_PTR() *QSGMaterial
}

func (p *QSGMaterial) QSGMaterial_PTR() *QSGMaterial {
	return p
}

func (p *QSGMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
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

//export callbackQSGMaterial_Compare
func callbackQSGMaterial_Compare(ptr unsafe.Pointer, ptrName *C.char, other unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGMaterial::compare")

	if signal := qt.GetSignal(C.GoString(ptrName), "compare"); signal != nil {
		return C.int(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other)))
	}

	return C.int(NewQSGMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other)))
}

func (ptr *QSGMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	defer qt.Recovering("connect QSGMaterial::compare")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "compare", f)
	}
}

func (ptr *QSGMaterial) DisconnectCompare() {
	defer qt.Recovering("disconnect QSGMaterial::compare")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "compare")
	}
}

func (ptr *QSGMaterial) Compare(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

func (ptr *QSGMaterial) CompareDefault(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

//export callbackQSGMaterial_CreateShader
func callbackQSGMaterial_CreateShader(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGMaterial::createShader")

	if signal := qt.GetSignal(C.GoString(ptrName), "createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(nil)
}

func (ptr *QSGMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	defer qt.Recovering("connect QSGMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "createShader", f)
	}
}

func (ptr *QSGMaterial) DisconnectCreateShader() {
	defer qt.Recovering("disconnect QSGMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "createShader")
	}
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

//export callbackQSGMaterial_Type
func callbackQSGMaterial_Type(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGMaterial::type")

	if signal := qt.GetSignal(C.GoString(ptrName), "type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(nil)
}

func (ptr *QSGMaterial) ConnectType(f func() *QSGMaterialType) {
	defer qt.Recovering("connect QSGMaterial::type")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "type", f)
	}
}

func (ptr *QSGMaterial) DisconnectType() {
	defer qt.Recovering("disconnect QSGMaterial::type")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "type")
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

func (p *QSGMaterialShader) QSGMaterialShader_PTR() *QSGMaterialShader {
	return p
}

func (p *QSGMaterialShader) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGMaterialShader) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
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

//export callbackQSGMaterialShader_FragmentShader
func callbackQSGMaterialShader_FragmentShader(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QSGMaterialShader::fragmentShader")

	if signal := qt.GetSignal(C.GoString(ptrName), "fragmentShader"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSGMaterialShaderFromPointer(ptr).FragmentShaderDefault())
}

func (ptr *QSGMaterialShader) ConnectFragmentShader(f func() string) {
	defer qt.Recovering("connect QSGMaterialShader::fragmentShader")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "fragmentShader", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectFragmentShader() {
	defer qt.Recovering("disconnect QSGMaterialShader::fragmentShader")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "fragmentShader")
	}
}

func (ptr *QSGMaterialShader) FragmentShader() string {
	defer qt.Recovering("QSGMaterialShader::fragmentShader")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGMaterialShader_FragmentShader(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGMaterialShader) FragmentShaderDefault() string {
	defer qt.Recovering("QSGMaterialShader::fragmentShader")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGMaterialShader_FragmentShaderDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQSGMaterialShader_VertexShader
func callbackQSGMaterialShader_VertexShader(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QSGMaterialShader::vertexShader")

	if signal := qt.GetSignal(C.GoString(ptrName), "vertexShader"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString(NewQSGMaterialShaderFromPointer(ptr).VertexShaderDefault())
}

func (ptr *QSGMaterialShader) ConnectVertexShader(f func() string) {
	defer qt.Recovering("connect QSGMaterialShader::vertexShader")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "vertexShader", f)
	}
}

func (ptr *QSGMaterialShader) DisconnectVertexShader() {
	defer qt.Recovering("disconnect QSGMaterialShader::vertexShader")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "vertexShader")
	}
}

func (ptr *QSGMaterialShader) VertexShader() string {
	defer qt.Recovering("QSGMaterialShader::vertexShader")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGMaterialShader_VertexShader(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSGMaterialShader) VertexShaderDefault() string {
	defer qt.Recovering("QSGMaterialShader::vertexShader")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSGMaterialShader_VertexShaderDefault(ptr.Pointer()))
	}
	return ""
}

//export callbackQSGMaterialShader_Activate
func callbackQSGMaterialShader_Activate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::activate")

	if signal := qt.GetSignal(C.GoString(ptrName), "activate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).ActivateDefault()
	}
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

//export callbackQSGMaterialShader_Deactivate
func callbackQSGMaterialShader_Deactivate(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::deactivate")

	if signal := qt.GetSignal(C.GoString(ptrName), "deactivate"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).DeactivateDefault()
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

func (ptr *QSGMaterialShader) Program() *gui.QOpenGLShaderProgram {
	defer qt.Recovering("QSGMaterialShader::program")

	if ptr.Pointer() != nil {
		return gui.NewQOpenGLShaderProgramFromPointer(C.QSGMaterialShader_Program(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGMaterialShader_Compile
func callbackQSGMaterialShader_Compile(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::compile")

	if signal := qt.GetSignal(C.GoString(ptrName), "compile"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).CompileDefault()
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

//export callbackQSGMaterialShader_Initialize
func callbackQSGMaterialShader_Initialize(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGMaterialShader::initialize")

	if signal := qt.GetSignal(C.GoString(ptrName), "initialize"); signal != nil {
		signal.(func())()
	} else {
		NewQSGMaterialShaderFromPointer(ptr).InitializeDefault()
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

func (ptr *QSGMaterialShader) SetShaderSourceFile(ty gui.QOpenGLShader__ShaderTypeBit, sourceFile string) {
	defer qt.Recovering("QSGMaterialShader::setShaderSourceFile")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_SetShaderSourceFile(ptr.Pointer(), C.int(ty), C.CString(sourceFile))
	}
}

func (ptr *QSGMaterialShader) SetShaderSourceFiles(ty gui.QOpenGLShader__ShaderTypeBit, sourceFiles []string) {
	defer qt.Recovering("QSGMaterialShader::setShaderSourceFiles")

	if ptr.Pointer() != nil {
		C.QSGMaterialShader_SetShaderSourceFiles(ptr.Pointer(), C.int(ty), C.CString(strings.Join(sourceFiles, "|")))
	}
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

func (p *QSGMaterialType) QSGMaterialType_PTR() *QSGMaterialType {
	return p
}

func (p *QSGMaterialType) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGMaterialType) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
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

type QSGNode struct {
	ptr unsafe.Pointer
}

type QSGNode_ITF interface {
	QSGNode_PTR() *QSGNode
}

func (p *QSGNode) QSGNode_PTR() *QSGNode {
	return p
}

func (p *QSGNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QSGNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
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

//export callbackQSGNode_IsSubtreeBlocked
func callbackQSGNode_IsSubtreeBlocked(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGNode::isSubtreeBlocked")

	if signal := qt.GetSignal(C.GoString(ptrName), "isSubtreeBlocked"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGNodeFromPointer(ptr).IsSubtreeBlockedDefault()))
}

func (ptr *QSGNode) ConnectIsSubtreeBlocked(f func() bool) {
	defer qt.Recovering("connect QSGNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked", f)
	}
}

func (ptr *QSGNode) DisconnectIsSubtreeBlocked() {
	defer qt.Recovering("disconnect QSGNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked")
	}
}

func (ptr *QSGNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGNode) IsSubtreeBlockedDefault() bool {
	defer qt.Recovering("QSGNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
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

//export callbackQSGNode_Preprocess
func callbackQSGNode_Preprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGNodeFromPointer(ptr).PreprocessDefault()
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
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QSGNode_DestroyQSGNode(ptr.Pointer())
		ptr.SetPointer(nil)
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

func (p *QSGOpacityNode) QSGOpacityNode_PTR() *QSGOpacityNode {
	return p
}

func (p *QSGOpacityNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGOpacityNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGNode_PTR().SetPointer(ptr)
	}
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
		ptr.SetPointer(nil)
	}
}

//export callbackQSGOpacityNode_IsSubtreeBlocked
func callbackQSGOpacityNode_IsSubtreeBlocked(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGOpacityNode::isSubtreeBlocked")

	if signal := qt.GetSignal(C.GoString(ptrName), "isSubtreeBlocked"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGOpacityNodeFromPointer(ptr).IsSubtreeBlockedDefault()))
}

func (ptr *QSGOpacityNode) ConnectIsSubtreeBlocked(f func() bool) {
	defer qt.Recovering("connect QSGOpacityNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked", f)
	}
}

func (ptr *QSGOpacityNode) DisconnectIsSubtreeBlocked() {
	defer qt.Recovering("disconnect QSGOpacityNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked")
	}
}

func (ptr *QSGOpacityNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGOpacityNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGOpacityNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGOpacityNode) IsSubtreeBlockedDefault() bool {
	defer qt.Recovering("QSGOpacityNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGOpacityNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGOpacityNode_Preprocess
func callbackQSGOpacityNode_Preprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGOpacityNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGOpacityNodeFromPointer(ptr).PreprocessDefault()
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

func (p *QSGOpaqueTextureMaterial) QSGOpaqueTextureMaterial_PTR() *QSGOpaqueTextureMaterial {
	return p
}

func (p *QSGOpaqueTextureMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGOpaqueTextureMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterial_PTR().SetPointer(ptr)
	}
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

func NewQSGOpaqueTextureMaterial() *QSGOpaqueTextureMaterial {
	defer qt.Recovering("QSGOpaqueTextureMaterial::QSGOpaqueTextureMaterial")

	return newQSGOpaqueTextureMaterialFromPointer(C.QSGOpaqueTextureMaterial_NewQSGOpaqueTextureMaterial())
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

func (ptr *QSGOpaqueTextureMaterial) M_texture() *QSGTexture {
	defer qt.Recovering("QSGOpaqueTextureMaterial::m_texture")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGOpaqueTextureMaterial_M_texture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) SetM_texture(vqs QSGTexture_ITF) {
	defer qt.Recovering("QSGOpaqueTextureMaterial::setM_texture")

	if ptr.Pointer() != nil {
		C.QSGOpaqueTextureMaterial_SetM_texture(ptr.Pointer(), PointerFromQSGTexture(vqs))
	}
}

//export callbackQSGOpaqueTextureMaterial_Compare
func callbackQSGOpaqueTextureMaterial_Compare(ptr unsafe.Pointer, ptrName *C.char, other unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGOpaqueTextureMaterial::compare")

	if signal := qt.GetSignal(C.GoString(ptrName), "compare"); signal != nil {
		return C.int(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other)))
	}

	return C.int(NewQSGOpaqueTextureMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other)))
}

func (ptr *QSGOpaqueTextureMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	defer qt.Recovering("connect QSGOpaqueTextureMaterial::compare")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "compare", f)
	}
}

func (ptr *QSGOpaqueTextureMaterial) DisconnectCompare() {
	defer qt.Recovering("disconnect QSGOpaqueTextureMaterial::compare")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "compare")
	}
}

func (ptr *QSGOpaqueTextureMaterial) Compare(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGOpaqueTextureMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGOpaqueTextureMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

func (ptr *QSGOpaqueTextureMaterial) CompareDefault(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGOpaqueTextureMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGOpaqueTextureMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

//export callbackQSGOpaqueTextureMaterial_CreateShader
func callbackQSGOpaqueTextureMaterial_CreateShader(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGOpaqueTextureMaterial::createShader")

	if signal := qt.GetSignal(C.GoString(ptrName), "createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGOpaqueTextureMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGOpaqueTextureMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	defer qt.Recovering("connect QSGOpaqueTextureMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "createShader", f)
	}
}

func (ptr *QSGOpaqueTextureMaterial) DisconnectCreateShader() {
	defer qt.Recovering("disconnect QSGOpaqueTextureMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "createShader")
	}
}

func (ptr *QSGOpaqueTextureMaterial) CreateShader() *QSGMaterialShader {
	defer qt.Recovering("QSGOpaqueTextureMaterial::createShader")

	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGOpaqueTextureMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) CreateShaderDefault() *QSGMaterialShader {
	defer qt.Recovering("QSGOpaqueTextureMaterial::createShader")

	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGOpaqueTextureMaterial_CreateShaderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGOpaqueTextureMaterial_Type
func callbackQSGOpaqueTextureMaterial_Type(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGOpaqueTextureMaterial::type")

	if signal := qt.GetSignal(C.GoString(ptrName), "type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGOpaqueTextureMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGOpaqueTextureMaterial) ConnectType(f func() *QSGMaterialType) {
	defer qt.Recovering("connect QSGOpaqueTextureMaterial::type")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "type", f)
	}
}

func (ptr *QSGOpaqueTextureMaterial) DisconnectType() {
	defer qt.Recovering("disconnect QSGOpaqueTextureMaterial::type")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "type")
	}
}

func (ptr *QSGOpaqueTextureMaterial) Type() *QSGMaterialType {
	defer qt.Recovering("QSGOpaqueTextureMaterial::type")

	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGOpaqueTextureMaterial_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGOpaqueTextureMaterial) TypeDefault() *QSGMaterialType {
	defer qt.Recovering("QSGOpaqueTextureMaterial::type")

	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGOpaqueTextureMaterial_TypeDefault(ptr.Pointer()))
	}
	return nil
}

type QSGSimpleMaterial struct {
	QSGMaterial
}

type QSGSimpleMaterial_ITF interface {
	QSGMaterial_ITF
	QSGSimpleMaterial_PTR() *QSGSimpleMaterial
}

func (p *QSGSimpleMaterial) QSGSimpleMaterial_PTR() *QSGSimpleMaterial {
	return p
}

func (p *QSGSimpleMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGSimpleMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterial_PTR().SetPointer(ptr)
	}
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

type QSGSimpleMaterialShader struct {
	QSGMaterialShader
}

type QSGSimpleMaterialShader_ITF interface {
	QSGMaterialShader_ITF
	QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader
}

func (p *QSGSimpleMaterialShader) QSGSimpleMaterialShader_PTR() *QSGSimpleMaterialShader {
	return p
}

func (p *QSGSimpleMaterialShader) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterialShader_PTR().Pointer()
	}
	return nil
}

func (p *QSGSimpleMaterialShader) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterialShader_PTR().SetPointer(ptr)
	}
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

type QSGSimpleRectNode struct {
	QSGGeometryNode
}

type QSGSimpleRectNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleRectNode_PTR() *QSGSimpleRectNode
}

func (p *QSGSimpleRectNode) QSGSimpleRectNode_PTR() *QSGSimpleRectNode {
	return p
}

func (p *QSGSimpleRectNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGSimpleRectNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGGeometryNode_PTR().SetPointer(ptr)
	}
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

//export callbackQSGSimpleRectNode_IsSubtreeBlocked
func callbackQSGSimpleRectNode_IsSubtreeBlocked(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGSimpleRectNode::isSubtreeBlocked")

	if signal := qt.GetSignal(C.GoString(ptrName), "isSubtreeBlocked"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGSimpleRectNodeFromPointer(ptr).IsSubtreeBlockedDefault()))
}

func (ptr *QSGSimpleRectNode) ConnectIsSubtreeBlocked(f func() bool) {
	defer qt.Recovering("connect QSGSimpleRectNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked", f)
	}
}

func (ptr *QSGSimpleRectNode) DisconnectIsSubtreeBlocked() {
	defer qt.Recovering("disconnect QSGSimpleRectNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked")
	}
}

func (ptr *QSGSimpleRectNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGSimpleRectNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGSimpleRectNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGSimpleRectNode) IsSubtreeBlockedDefault() bool {
	defer qt.Recovering("QSGSimpleRectNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGSimpleRectNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGSimpleRectNode_Preprocess
func callbackQSGSimpleRectNode_Preprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGSimpleRectNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGSimpleRectNodeFromPointer(ptr).PreprocessDefault()
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

//QSGSimpleTextureNode::TextureCoordinatesTransformFlag
type QSGSimpleTextureNode__TextureCoordinatesTransformFlag int64

const (
	QSGSimpleTextureNode__NoTransform        = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x00)
	QSGSimpleTextureNode__MirrorHorizontally = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x01)
	QSGSimpleTextureNode__MirrorVertically   = QSGSimpleTextureNode__TextureCoordinatesTransformFlag(0x02)
)

type QSGSimpleTextureNode struct {
	QSGGeometryNode
}

type QSGSimpleTextureNode_ITF interface {
	QSGGeometryNode_ITF
	QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode
}

func (p *QSGSimpleTextureNode) QSGSimpleTextureNode_PTR() *QSGSimpleTextureNode {
	return p
}

func (p *QSGSimpleTextureNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGGeometryNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGSimpleTextureNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGGeometryNode_PTR().SetPointer(ptr)
	}
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
		ptr.SetPointer(nil)
	}
}

//export callbackQSGSimpleTextureNode_IsSubtreeBlocked
func callbackQSGSimpleTextureNode_IsSubtreeBlocked(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGSimpleTextureNode::isSubtreeBlocked")

	if signal := qt.GetSignal(C.GoString(ptrName), "isSubtreeBlocked"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGSimpleTextureNodeFromPointer(ptr).IsSubtreeBlockedDefault()))
}

func (ptr *QSGSimpleTextureNode) ConnectIsSubtreeBlocked(f func() bool) {
	defer qt.Recovering("connect QSGSimpleTextureNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked", f)
	}
}

func (ptr *QSGSimpleTextureNode) DisconnectIsSubtreeBlocked() {
	defer qt.Recovering("disconnect QSGSimpleTextureNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked")
	}
}

func (ptr *QSGSimpleTextureNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGSimpleTextureNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGSimpleTextureNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGSimpleTextureNode) IsSubtreeBlockedDefault() bool {
	defer qt.Recovering("QSGSimpleTextureNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGSimpleTextureNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGSimpleTextureNode_Preprocess
func callbackQSGSimpleTextureNode_Preprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGSimpleTextureNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGSimpleTextureNodeFromPointer(ptr).PreprocessDefault()
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

type QSGTexture struct {
	core.QObject
}

type QSGTexture_ITF interface {
	core.QObject_ITF
	QSGTexture_PTR() *QSGTexture
}

func (p *QSGTexture) QSGTexture_PTR() *QSGTexture {
	return p
}

func (p *QSGTexture) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSGTexture) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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

func NewQSGTexture() *QSGTexture {
	defer qt.Recovering("QSGTexture::QSGTexture")

	return newQSGTextureFromPointer(C.QSGTexture_NewQSGTexture())
}

//export callbackQSGTexture_Bind
func callbackQSGTexture_Bind(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGTexture::bind")

	if signal := qt.GetSignal(C.GoString(ptrName), "bind"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QSGTexture) ConnectBind(f func()) {
	defer qt.Recovering("connect QSGTexture::bind")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "bind", f)
	}
}

func (ptr *QSGTexture) DisconnectBind() {
	defer qt.Recovering("disconnect QSGTexture::bind")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "bind")
	}
}

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

//export callbackQSGTexture_HasAlphaChannel
func callbackQSGTexture_HasAlphaChannel(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGTexture::hasAlphaChannel")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasAlphaChannel"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QSGTexture) ConnectHasAlphaChannel(f func() bool) {
	defer qt.Recovering("connect QSGTexture::hasAlphaChannel")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasAlphaChannel", f)
	}
}

func (ptr *QSGTexture) DisconnectHasAlphaChannel() {
	defer qt.Recovering("disconnect QSGTexture::hasAlphaChannel")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasAlphaChannel")
	}
}

func (ptr *QSGTexture) HasAlphaChannel() bool {
	defer qt.Recovering("QSGTexture::hasAlphaChannel")

	if ptr.Pointer() != nil {
		return C.QSGTexture_HasAlphaChannel(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGTexture_HasMipmaps
func callbackQSGTexture_HasMipmaps(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGTexture::hasMipmaps")

	if signal := qt.GetSignal(C.GoString(ptrName), "hasMipmaps"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QSGTexture) ConnectHasMipmaps(f func() bool) {
	defer qt.Recovering("connect QSGTexture::hasMipmaps")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hasMipmaps", f)
	}
}

func (ptr *QSGTexture) DisconnectHasMipmaps() {
	defer qt.Recovering("disconnect QSGTexture::hasMipmaps")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hasMipmaps")
	}
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

//export callbackQSGTexture_IsAtlasTexture
func callbackQSGTexture_IsAtlasTexture(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGTexture::isAtlasTexture")

	if signal := qt.GetSignal(C.GoString(ptrName), "isAtlasTexture"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGTextureFromPointer(ptr).IsAtlasTextureDefault()))
}

func (ptr *QSGTexture) ConnectIsAtlasTexture(f func() bool) {
	defer qt.Recovering("connect QSGTexture::isAtlasTexture")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isAtlasTexture", f)
	}
}

func (ptr *QSGTexture) DisconnectIsAtlasTexture() {
	defer qt.Recovering("disconnect QSGTexture::isAtlasTexture")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isAtlasTexture")
	}
}

func (ptr *QSGTexture) IsAtlasTexture() bool {
	defer qt.Recovering("QSGTexture::isAtlasTexture")

	if ptr.Pointer() != nil {
		return C.QSGTexture_IsAtlasTexture(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTexture) IsAtlasTextureDefault() bool {
	defer qt.Recovering("QSGTexture::isAtlasTexture")

	if ptr.Pointer() != nil {
		return C.QSGTexture_IsAtlasTextureDefault(ptr.Pointer()) != 0
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

//export callbackQSGTexture_NormalizedTextureSubRect
func callbackQSGTexture_NormalizedTextureSubRect(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGTexture::normalizedTextureSubRect")

	if signal := qt.GetSignal(C.GoString(ptrName), "normalizedTextureSubRect"); signal != nil {
		return core.PointerFromQRectF(signal.(func() *core.QRectF)())
	}

	return core.PointerFromQRectF(NewQSGTextureFromPointer(ptr).NormalizedTextureSubRectDefault())
}

func (ptr *QSGTexture) ConnectNormalizedTextureSubRect(f func() *core.QRectF) {
	defer qt.Recovering("connect QSGTexture::normalizedTextureSubRect")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "normalizedTextureSubRect", f)
	}
}

func (ptr *QSGTexture) DisconnectNormalizedTextureSubRect() {
	defer qt.Recovering("disconnect QSGTexture::normalizedTextureSubRect")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "normalizedTextureSubRect")
	}
}

func (ptr *QSGTexture) NormalizedTextureSubRect() *core.QRectF {
	defer qt.Recovering("QSGTexture::normalizedTextureSubRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGTexture_NormalizedTextureSubRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTexture) NormalizedTextureSubRectDefault() *core.QRectF {
	defer qt.Recovering("QSGTexture::normalizedTextureSubRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFFromPointer(C.QSGTexture_NormalizedTextureSubRectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGTexture_RemovedFromAtlas
func callbackQSGTexture_RemovedFromAtlas(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGTexture::removedFromAtlas")

	if signal := qt.GetSignal(C.GoString(ptrName), "removedFromAtlas"); signal != nil {
		return PointerFromQSGTexture(signal.(func() *QSGTexture)())
	}

	return PointerFromQSGTexture(NewQSGTextureFromPointer(ptr).RemovedFromAtlasDefault())
}

func (ptr *QSGTexture) ConnectRemovedFromAtlas(f func() *QSGTexture) {
	defer qt.Recovering("connect QSGTexture::removedFromAtlas")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "removedFromAtlas", f)
	}
}

func (ptr *QSGTexture) DisconnectRemovedFromAtlas() {
	defer qt.Recovering("disconnect QSGTexture::removedFromAtlas")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "removedFromAtlas")
	}
}

func (ptr *QSGTexture) RemovedFromAtlas() *QSGTexture {
	defer qt.Recovering("QSGTexture::removedFromAtlas")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGTexture_RemovedFromAtlas(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTexture) RemovedFromAtlasDefault() *QSGTexture {
	defer qt.Recovering("QSGTexture::removedFromAtlas")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGTexture_RemovedFromAtlasDefault(ptr.Pointer()))
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

//export callbackQSGTexture_TextureId
func callbackQSGTexture_TextureId(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGTexture::textureId")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureId"); signal != nil {
		return C.int(signal.(func() int)())
	}

	return C.int(0)
}

func (ptr *QSGTexture) ConnectTextureId(f func() int) {
	defer qt.Recovering("connect QSGTexture::textureId")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureId", f)
	}
}

func (ptr *QSGTexture) DisconnectTextureId() {
	defer qt.Recovering("disconnect QSGTexture::textureId")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureId")
	}
}

func (ptr *QSGTexture) TextureId() int {
	defer qt.Recovering("QSGTexture::textureId")

	if ptr.Pointer() != nil {
		return int(C.QSGTexture_TextureId(ptr.Pointer()))
	}
	return 0
}

//export callbackQSGTexture_TextureSize
func callbackQSGTexture_TextureSize(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGTexture::textureSize")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureSize"); signal != nil {
		return core.PointerFromQSize(signal.(func() *core.QSize)())
	}

	return core.PointerFromQSize(nil)
}

func (ptr *QSGTexture) ConnectTextureSize(f func() *core.QSize) {
	defer qt.Recovering("connect QSGTexture::textureSize")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "textureSize", f)
	}
}

func (ptr *QSGTexture) DisconnectTextureSize() {
	defer qt.Recovering("disconnect QSGTexture::textureSize")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "textureSize")
	}
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
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSGTexture_DestroyQSGTexture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGTexture_TimerEvent
func callbackQSGTexture_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTexture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQSGTexture_ChildEvent
func callbackQSGTexture_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTexture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQSGTexture_ConnectNotify
func callbackQSGTexture_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGTexture::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTexture) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGTexture::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QSGTexture) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSGTexture::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QSGTexture) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGTexture::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGTexture_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGTexture) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGTexture::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGTexture_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTexture_CustomEvent
func callbackQSGTexture_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTexture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGTextureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQSGTexture_DeleteLater
func callbackQSGTexture_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGTexture::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTextureFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGTexture) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSGTexture::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QSGTexture) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSGTexture::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QSGTexture) DeleteLater() {
	defer qt.Recovering("QSGTexture::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSGTexture_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGTexture) DeleteLaterDefault() {
	defer qt.Recovering("QSGTexture::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSGTexture_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGTexture_DisconnectNotify
func callbackQSGTexture_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGTexture::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTexture) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGTexture::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QSGTexture) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSGTexture::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QSGTexture) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGTexture::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGTexture_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGTexture) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGTexture::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGTexture_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTexture_Event
func callbackQSGTexture_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGTexture::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQSGTextureFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QSGTexture) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSGTexture::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QSGTexture) DisconnectEvent() {
	defer qt.Recovering("disconnect QSGTexture::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QSGTexture) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGTexture::event")

	if ptr.Pointer() != nil {
		return C.QSGTexture_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGTexture) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGTexture::event")

	if ptr.Pointer() != nil {
		return C.QSGTexture_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGTexture_EventFilter
func callbackQSGTexture_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGTexture::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQSGTextureFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QSGTexture) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSGTexture::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QSGTexture) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSGTexture::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QSGTexture) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGTexture::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGTexture_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGTexture) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGTexture::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGTexture_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGTexture_MetaObject
func callbackQSGTexture_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGTexture::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGTextureFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGTexture) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSGTexture::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QSGTexture) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSGTexture::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QSGTexture) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSGTexture::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTexture_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTexture) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSGTexture::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTexture_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSGTextureMaterial struct {
	QSGOpaqueTextureMaterial
}

type QSGTextureMaterial_ITF interface {
	QSGOpaqueTextureMaterial_ITF
	QSGTextureMaterial_PTR() *QSGTextureMaterial
}

func (p *QSGTextureMaterial) QSGTextureMaterial_PTR() *QSGTextureMaterial {
	return p
}

func (p *QSGTextureMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGOpaqueTextureMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGTextureMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGOpaqueTextureMaterial_PTR().SetPointer(ptr)
	}
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

//export callbackQSGTextureMaterial_Compare
func callbackQSGTextureMaterial_Compare(ptr unsafe.Pointer, ptrName *C.char, other unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGTextureMaterial::compare")

	if signal := qt.GetSignal(C.GoString(ptrName), "compare"); signal != nil {
		return C.int(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other)))
	}

	return C.int(NewQSGTextureMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other)))
}

func (ptr *QSGTextureMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	defer qt.Recovering("connect QSGTextureMaterial::compare")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "compare", f)
	}
}

func (ptr *QSGTextureMaterial) DisconnectCompare() {
	defer qt.Recovering("disconnect QSGTextureMaterial::compare")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "compare")
	}
}

func (ptr *QSGTextureMaterial) Compare(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGTextureMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGTextureMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

func (ptr *QSGTextureMaterial) CompareDefault(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGTextureMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGTextureMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

//export callbackQSGTextureMaterial_CreateShader
func callbackQSGTextureMaterial_CreateShader(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGTextureMaterial::createShader")

	if signal := qt.GetSignal(C.GoString(ptrName), "createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGTextureMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGTextureMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	defer qt.Recovering("connect QSGTextureMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "createShader", f)
	}
}

func (ptr *QSGTextureMaterial) DisconnectCreateShader() {
	defer qt.Recovering("disconnect QSGTextureMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "createShader")
	}
}

func (ptr *QSGTextureMaterial) CreateShader() *QSGMaterialShader {
	defer qt.Recovering("QSGTextureMaterial::createShader")

	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGTextureMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureMaterial) CreateShaderDefault() *QSGMaterialShader {
	defer qt.Recovering("QSGTextureMaterial::createShader")

	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGTextureMaterial_CreateShaderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGTextureMaterial_Type
func callbackQSGTextureMaterial_Type(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGTextureMaterial::type")

	if signal := qt.GetSignal(C.GoString(ptrName), "type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGTextureMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGTextureMaterial) ConnectType(f func() *QSGMaterialType) {
	defer qt.Recovering("connect QSGTextureMaterial::type")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "type", f)
	}
}

func (ptr *QSGTextureMaterial) DisconnectType() {
	defer qt.Recovering("disconnect QSGTextureMaterial::type")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "type")
	}
}

func (ptr *QSGTextureMaterial) Type() *QSGMaterialType {
	defer qt.Recovering("QSGTextureMaterial::type")

	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGTextureMaterial_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureMaterial) TypeDefault() *QSGMaterialType {
	defer qt.Recovering("QSGTextureMaterial::type")

	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGTextureMaterial_TypeDefault(ptr.Pointer()))
	}
	return nil
}

type QSGTextureProvider struct {
	core.QObject
}

type QSGTextureProvider_ITF interface {
	core.QObject_ITF
	QSGTextureProvider_PTR() *QSGTextureProvider
}

func (p *QSGTextureProvider) QSGTextureProvider_PTR() *QSGTextureProvider {
	return p
}

func (p *QSGTextureProvider) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QSGTextureProvider) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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

//export callbackQSGTextureProvider_Texture
func callbackQSGTextureProvider_Texture(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGTextureProvider::texture")

	if signal := qt.GetSignal(C.GoString(ptrName), "texture"); signal != nil {
		return PointerFromQSGTexture(signal.(func() *QSGTexture)())
	}

	return PointerFromQSGTexture(nil)
}

func (ptr *QSGTextureProvider) ConnectTexture(f func() *QSGTexture) {
	defer qt.Recovering("connect QSGTextureProvider::texture")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "texture", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectTexture() {
	defer qt.Recovering("disconnect QSGTextureProvider::texture")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "texture")
	}
}

func (ptr *QSGTextureProvider) Texture() *QSGTexture {
	defer qt.Recovering("QSGTextureProvider::texture")

	if ptr.Pointer() != nil {
		return NewQSGTextureFromPointer(C.QSGTextureProvider_Texture(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGTextureProvider_TextureChanged
func callbackQSGTextureProvider_TextureChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGTextureProvider::textureChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "textureChanged"); signal != nil {
		signal.(func())()
	}

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

func (ptr *QSGTextureProvider) TextureChanged() {
	defer qt.Recovering("QSGTextureProvider::textureChanged")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_TextureChanged(ptr.Pointer())
	}
}

//export callbackQSGTextureProvider_TimerEvent
func callbackQSGTextureProvider_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQSGTextureProvider_ChildEvent
func callbackQSGTextureProvider_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQSGTextureProvider_ConnectNotify
func callbackQSGTextureProvider_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureProviderFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTextureProvider) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGTextureProvider::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QSGTextureProvider::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QSGTextureProvider) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGTextureProvider::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGTextureProvider) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGTextureProvider::connectNotify")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTextureProvider_CustomEvent
func callbackQSGTextureProvider_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQSGTextureProviderFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQSGTextureProvider_DeleteLater
func callbackQSGTextureProvider_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGTextureProvider::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTextureProviderFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QSGTextureProvider) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QSGTextureProvider::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QSGTextureProvider::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QSGTextureProvider) DeleteLater() {
	defer qt.Recovering("QSGTextureProvider::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSGTextureProvider_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QSGTextureProvider) DeleteLaterDefault() {
	defer qt.Recovering("QSGTextureProvider::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QSGTextureProvider_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQSGTextureProvider_DisconnectNotify
func callbackQSGTextureProvider_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QSGTextureProvider::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQSGTextureProviderFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QSGTextureProvider) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QSGTextureProvider::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QSGTextureProvider::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QSGTextureProvider) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGTextureProvider::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QSGTextureProvider) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QSGTextureProvider::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QSGTextureProvider_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQSGTextureProvider_Event
func callbackQSGTextureProvider_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGTextureProvider::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQSGTextureProviderFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QSGTextureProvider) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QSGTextureProvider::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectEvent() {
	defer qt.Recovering("disconnect QSGTextureProvider::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QSGTextureProvider) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGTextureProvider::event")

	if ptr.Pointer() != nil {
		return C.QSGTextureProvider_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QSGTextureProvider) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QSGTextureProvider::event")

	if ptr.Pointer() != nil {
		return C.QSGTextureProvider_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQSGTextureProvider_EventFilter
func callbackQSGTextureProvider_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGTextureProvider::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQSGTextureProviderFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QSGTextureProvider) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QSGTextureProvider::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QSGTextureProvider::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QSGTextureProvider) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGTextureProvider::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGTextureProvider_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QSGTextureProvider) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QSGTextureProvider::eventFilter")

	if ptr.Pointer() != nil {
		return C.QSGTextureProvider_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQSGTextureProvider_MetaObject
func callbackQSGTextureProvider_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGTextureProvider::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQSGTextureProviderFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QSGTextureProvider) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QSGTextureProvider::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QSGTextureProvider) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QSGTextureProvider::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QSGTextureProvider) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QSGTextureProvider::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTextureProvider_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGTextureProvider) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QSGTextureProvider::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QSGTextureProvider_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QSGTransformNode struct {
	QSGNode
}

type QSGTransformNode_ITF interface {
	QSGNode_ITF
	QSGTransformNode_PTR() *QSGTransformNode
}

func (p *QSGTransformNode) QSGTransformNode_PTR() *QSGTransformNode {
	return p
}

func (p *QSGTransformNode) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGNode_PTR().Pointer()
	}
	return nil
}

func (p *QSGTransformNode) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGNode_PTR().SetPointer(ptr)
	}
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
		ptr.SetPointer(nil)
	}
}

//export callbackQSGTransformNode_IsSubtreeBlocked
func callbackQSGTransformNode_IsSubtreeBlocked(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QSGTransformNode::isSubtreeBlocked")

	if signal := qt.GetSignal(C.GoString(ptrName), "isSubtreeBlocked"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQSGTransformNodeFromPointer(ptr).IsSubtreeBlockedDefault()))
}

func (ptr *QSGTransformNode) ConnectIsSubtreeBlocked(f func() bool) {
	defer qt.Recovering("connect QSGTransformNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked", f)
	}
}

func (ptr *QSGTransformNode) DisconnectIsSubtreeBlocked() {
	defer qt.Recovering("disconnect QSGTransformNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "isSubtreeBlocked")
	}
}

func (ptr *QSGTransformNode) IsSubtreeBlocked() bool {
	defer qt.Recovering("QSGTransformNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGTransformNode_IsSubtreeBlocked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSGTransformNode) IsSubtreeBlockedDefault() bool {
	defer qt.Recovering("QSGTransformNode::isSubtreeBlocked")

	if ptr.Pointer() != nil {
		return C.QSGTransformNode_IsSubtreeBlockedDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQSGTransformNode_Preprocess
func callbackQSGTransformNode_Preprocess(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QSGTransformNode::preprocess")

	if signal := qt.GetSignal(C.GoString(ptrName), "preprocess"); signal != nil {
		signal.(func())()
	} else {
		NewQSGTransformNodeFromPointer(ptr).PreprocessDefault()
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

func (p *QSGVertexColorMaterial) QSGVertexColorMaterial_PTR() *QSGVertexColorMaterial {
	return p
}

func (p *QSGVertexColorMaterial) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QSGMaterial_PTR().Pointer()
	}
	return nil
}

func (p *QSGVertexColorMaterial) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QSGMaterial_PTR().SetPointer(ptr)
	}
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

func NewQSGVertexColorMaterial() *QSGVertexColorMaterial {
	defer qt.Recovering("QSGVertexColorMaterial::QSGVertexColorMaterial")

	return newQSGVertexColorMaterialFromPointer(C.QSGVertexColorMaterial_NewQSGVertexColorMaterial())
}

//export callbackQSGVertexColorMaterial_Compare
func callbackQSGVertexColorMaterial_Compare(ptr unsafe.Pointer, ptrName *C.char, other unsafe.Pointer) C.int {
	defer qt.Recovering("callback QSGVertexColorMaterial::compare")

	if signal := qt.GetSignal(C.GoString(ptrName), "compare"); signal != nil {
		return C.int(signal.(func(*QSGMaterial) int)(NewQSGMaterialFromPointer(other)))
	}

	return C.int(NewQSGVertexColorMaterialFromPointer(ptr).CompareDefault(NewQSGMaterialFromPointer(other)))
}

func (ptr *QSGVertexColorMaterial) ConnectCompare(f func(other *QSGMaterial) int) {
	defer qt.Recovering("connect QSGVertexColorMaterial::compare")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "compare", f)
	}
}

func (ptr *QSGVertexColorMaterial) DisconnectCompare() {
	defer qt.Recovering("disconnect QSGVertexColorMaterial::compare")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "compare")
	}
}

func (ptr *QSGVertexColorMaterial) Compare(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGVertexColorMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGVertexColorMaterial_Compare(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

func (ptr *QSGVertexColorMaterial) CompareDefault(other QSGMaterial_ITF) int {
	defer qt.Recovering("QSGVertexColorMaterial::compare")

	if ptr.Pointer() != nil {
		return int(C.QSGVertexColorMaterial_CompareDefault(ptr.Pointer(), PointerFromQSGMaterial(other)))
	}
	return 0
}

//export callbackQSGVertexColorMaterial_CreateShader
func callbackQSGVertexColorMaterial_CreateShader(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGVertexColorMaterial::createShader")

	if signal := qt.GetSignal(C.GoString(ptrName), "createShader"); signal != nil {
		return PointerFromQSGMaterialShader(signal.(func() *QSGMaterialShader)())
	}

	return PointerFromQSGMaterialShader(NewQSGVertexColorMaterialFromPointer(ptr).CreateShaderDefault())
}

func (ptr *QSGVertexColorMaterial) ConnectCreateShader(f func() *QSGMaterialShader) {
	defer qt.Recovering("connect QSGVertexColorMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "createShader", f)
	}
}

func (ptr *QSGVertexColorMaterial) DisconnectCreateShader() {
	defer qt.Recovering("disconnect QSGVertexColorMaterial::createShader")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "createShader")
	}
}

func (ptr *QSGVertexColorMaterial) CreateShader() *QSGMaterialShader {
	defer qt.Recovering("QSGVertexColorMaterial::createShader")

	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGVertexColorMaterial_CreateShader(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGVertexColorMaterial) CreateShaderDefault() *QSGMaterialShader {
	defer qt.Recovering("QSGVertexColorMaterial::createShader")

	if ptr.Pointer() != nil {
		return NewQSGMaterialShaderFromPointer(C.QSGVertexColorMaterial_CreateShaderDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQSGVertexColorMaterial_Type
func callbackQSGVertexColorMaterial_Type(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QSGVertexColorMaterial::type")

	if signal := qt.GetSignal(C.GoString(ptrName), "type"); signal != nil {
		return PointerFromQSGMaterialType(signal.(func() *QSGMaterialType)())
	}

	return PointerFromQSGMaterialType(NewQSGVertexColorMaterialFromPointer(ptr).TypeDefault())
}

func (ptr *QSGVertexColorMaterial) ConnectType(f func() *QSGMaterialType) {
	defer qt.Recovering("connect QSGVertexColorMaterial::type")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "type", f)
	}
}

func (ptr *QSGVertexColorMaterial) DisconnectType() {
	defer qt.Recovering("disconnect QSGVertexColorMaterial::type")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "type")
	}
}

func (ptr *QSGVertexColorMaterial) Type() *QSGMaterialType {
	defer qt.Recovering("QSGVertexColorMaterial::type")

	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGVertexColorMaterial_Type(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSGVertexColorMaterial) TypeDefault() *QSGMaterialType {
	defer qt.Recovering("QSGVertexColorMaterial::type")

	if ptr.Pointer() != nil {
		return NewQSGMaterialTypeFromPointer(C.QSGVertexColorMaterial_TypeDefault(ptr.Pointer()))
	}
	return nil
}
