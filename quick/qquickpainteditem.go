package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

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
func callbackQQuickPaintedItemContentsScaleChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::contentsScaleChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsScaleChanged"); signal != nil {
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

//export callbackQQuickPaintedItemContentsSizeChanged
func callbackQQuickPaintedItemContentsSizeChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::contentsSizeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contentsSizeChanged"); signal != nil {
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

//export callbackQQuickPaintedItemFillColorChanged
func callbackQQuickPaintedItemFillColorChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::fillColorChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fillColorChanged"); signal != nil {
		signal.(func())()
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
func callbackQQuickPaintedItemReleaseResources(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickPaintedItem::releaseResources")

	if signal := qt.GetSignal(C.GoString(ptrName), "releaseResources"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQQuickPaintedItemRenderTargetChanged(ptrName *C.char) {
	defer qt.Recovering("callback QQuickPaintedItem::renderTargetChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "renderTargetChanged"); signal != nil {
		signal.(func())()
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
func callbackQQuickPaintedItemClassBegin(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickPaintedItem::classBegin")

	if signal := qt.GetSignal(C.GoString(ptrName), "classBegin"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQQuickPaintedItemComponentComplete(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickPaintedItem::componentComplete")

	if signal := qt.GetSignal(C.GoString(ptrName), "componentComplete"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQQuickPaintedItemDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemFocusInEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemFocusOutEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemHoverEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::hoverEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverEnterEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemHoverLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::hoverLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverLeaveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemHoverMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::hoverMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hoverMoveEvent"); signal != nil {
		signal.(func(*gui.QHoverEvent))(gui.NewQHoverEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemKeyPressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemKeyReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemMouseDoubleClickEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemMouseUngrabEvent(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickPaintedItem::mouseUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseUngrabEvent"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQQuickPaintedItemTouchEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::touchEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchEvent"); signal != nil {
		signal.(func(*gui.QTouchEvent))(gui.NewQTouchEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemTouchUngrabEvent(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickPaintedItem::touchUngrabEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "touchUngrabEvent"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQQuickPaintedItemUpdatePolish(ptrName *C.char) bool {
	defer qt.Recovering("callback QQuickPaintedItem::updatePolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "updatePolish"); signal != nil {
		signal.(func())()
		return true
	}
	return false

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
func callbackQQuickPaintedItemWheelEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

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
func callbackQQuickPaintedItemCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QQuickPaintedItem::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
