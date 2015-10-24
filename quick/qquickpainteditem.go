package quick

//#include "qquickpainteditem.h"
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

type QQuickPaintedItemITF interface {
	QQuickItemITF
	QQuickPaintedItemPTR() *QQuickPaintedItem
}

func PointerFromQQuickPaintedItem(ptr QQuickPaintedItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickPaintedItemPTR().Pointer()
	}
	return nil
}

func QQuickPaintedItemFromPointer(ptr unsafe.Pointer) *QQuickPaintedItem {
	var n = new(QQuickPaintedItem)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QQuickPaintedItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QQuickPaintedItem) QQuickPaintedItemPTR() *QQuickPaintedItem {
	return ptr
}

//QQuickPaintedItem::PerformanceHint
type QQuickPaintedItem__PerformanceHint int

var (
	QQuickPaintedItem__FastFBOResizing = QQuickPaintedItem__PerformanceHint(0x1)
)

//QQuickPaintedItem::RenderTarget
type QQuickPaintedItem__RenderTarget int

var (
	QQuickPaintedItem__Image                      = QQuickPaintedItem__RenderTarget(0)
	QQuickPaintedItem__FramebufferObject          = QQuickPaintedItem__RenderTarget(1)
	QQuickPaintedItem__InvertedYFramebufferObject = QQuickPaintedItem__RenderTarget(2)
)

func (ptr *QQuickPaintedItem) RenderTarget() QQuickPaintedItem__RenderTarget {
	if ptr.Pointer() != nil {
		return QQuickPaintedItem__RenderTarget(C.QQuickPaintedItem_RenderTarget(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickPaintedItem) SetContentsSize(v core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetContentsSize(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQSize(v)))
	}
}

func (ptr *QQuickPaintedItem) SetFillColor(v gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetFillColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(v)))
	}
}

func (ptr *QQuickPaintedItem) SetRenderTarget(target QQuickPaintedItem__RenderTarget) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetRenderTarget(C.QtObjectPtr(ptr.Pointer()), C.int(target))
	}
}

func (ptr *QQuickPaintedItem) Antialiasing() bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Antialiasing(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) ConnectContentsScaleChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectContentsScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "contentsScaleChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsScaleChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsScaleChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "contentsScaleChanged")
	}
}

//export callbackQQuickPaintedItemContentsScaleChanged
func callbackQQuickPaintedItemContentsScaleChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "contentsScaleChanged").(func())()
}

func (ptr *QQuickPaintedItem) ConnectContentsSizeChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectContentsSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "contentsSizeChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsSizeChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsSizeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "contentsSizeChanged")
	}
}

//export callbackQQuickPaintedItemContentsSizeChanged
func callbackQQuickPaintedItemContentsSizeChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "contentsSizeChanged").(func())()
}

func (ptr *QQuickPaintedItem) ConnectFillColorChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectFillColorChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "fillColorChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectFillColorChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectFillColorChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "fillColorChanged")
	}
}

//export callbackQQuickPaintedItemFillColorChanged
func callbackQQuickPaintedItemFillColorChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "fillColorChanged").(func())()
}

func (ptr *QQuickPaintedItem) IsTextureProvider() bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_IsTextureProvider(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) Mipmap() bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Mipmap(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) OpaquePainting() bool {
	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_OpaquePainting(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) Paint(painter gui.QPainterITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)))
	}
}

func (ptr *QQuickPaintedItem) PerformanceHints() QQuickPaintedItem__PerformanceHint {
	if ptr.Pointer() != nil {
		return QQuickPaintedItem__PerformanceHint(C.QQuickPaintedItem_PerformanceHints(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQuickPaintedItem) ConnectRenderTargetChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectRenderTargetChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "renderTargetChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectRenderTargetChanged() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectRenderTargetChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "renderTargetChanged")
	}
}

//export callbackQQuickPaintedItemRenderTargetChanged
func callbackQQuickPaintedItemRenderTargetChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "renderTargetChanged").(func())()
}

func (ptr *QQuickPaintedItem) ResetContentsSize() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ResetContentsSize(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QQuickPaintedItem) SetAntialiasing(enable bool) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetAntialiasing(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QQuickPaintedItem) SetMipmap(enable bool) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetMipmap(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QQuickPaintedItem) SetOpaquePainting(opaque bool) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetOpaquePainting(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(opaque)))
	}
}

func (ptr *QQuickPaintedItem) SetPerformanceHint(hint QQuickPaintedItem__PerformanceHint, enabled bool) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetPerformanceHint(C.QtObjectPtr(ptr.Pointer()), C.int(hint), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickPaintedItem) SetPerformanceHints(hints QQuickPaintedItem__PerformanceHint) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetPerformanceHints(C.QtObjectPtr(ptr.Pointer()), C.int(hints))
	}
}

func (ptr *QQuickPaintedItem) TextureProvider() *QSGTextureProvider {
	if ptr.Pointer() != nil {
		return QSGTextureProviderFromPointer(unsafe.Pointer(C.QQuickPaintedItem_TextureProvider(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQuickPaintedItem) Update(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_Update(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QQuickPaintedItem) DestroyQQuickPaintedItem() {
	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DestroyQQuickPaintedItem(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
