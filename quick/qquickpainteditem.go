package quick

//#include "quick.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
		n.SetObjectName("QQuickPaintedItem_" + qt.RandomIdentifier())
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::contentsScale")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QQuickPaintedItem_ContentsScale(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickPaintedItem) FillColor() *gui.QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::fillColor")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QQuickPaintedItem_FillColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickPaintedItem) RenderTarget() QQuickPaintedItem__RenderTarget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::renderTarget")
		}
	}()

	if ptr.Pointer() != nil {
		return QQuickPaintedItem__RenderTarget(C.QQuickPaintedItem_RenderTarget(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickPaintedItem) SetContentsScale(v float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::setContentsScale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetContentsScale(ptr.Pointer(), C.double(v))
	}
}

func (ptr *QQuickPaintedItem) SetContentsSize(v core.QSize_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::setContentsSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetContentsSize(ptr.Pointer(), core.PointerFromQSize(v))
	}
}

func (ptr *QQuickPaintedItem) SetFillColor(v gui.QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::setFillColor")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetFillColor(ptr.Pointer(), gui.PointerFromQColor(v))
	}
}

func (ptr *QQuickPaintedItem) SetRenderTarget(target QQuickPaintedItem__RenderTarget) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::setRenderTarget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetRenderTarget(ptr.Pointer(), C.int(target))
	}
}

func (ptr *QQuickPaintedItem) Antialiasing() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::antialiasing")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Antialiasing(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) ConnectContentsScaleChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::contentsScaleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectContentsScaleChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsScaleChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsScaleChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::contentsScaleChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsScaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsScaleChanged")
	}
}

//export callbackQQuickPaintedItemContentsScaleChanged
func callbackQQuickPaintedItemContentsScaleChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::contentsScaleChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "contentsScaleChanged").(func())()
}

func (ptr *QQuickPaintedItem) ConnectContentsSizeChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::contentsSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectContentsSizeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contentsSizeChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectContentsSizeChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::contentsSizeChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectContentsSizeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contentsSizeChanged")
	}
}

//export callbackQQuickPaintedItemContentsSizeChanged
func callbackQQuickPaintedItemContentsSizeChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::contentsSizeChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "contentsSizeChanged").(func())()
}

func (ptr *QQuickPaintedItem) ConnectFillColorChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::fillColorChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectFillColorChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fillColorChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectFillColorChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::fillColorChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectFillColorChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fillColorChanged")
	}
}

//export callbackQQuickPaintedItemFillColorChanged
func callbackQQuickPaintedItemFillColorChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::fillColorChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "fillColorChanged").(func())()
}

func (ptr *QQuickPaintedItem) IsTextureProvider() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::isTextureProvider")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_IsTextureProvider(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) Mipmap() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::mipmap")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_Mipmap(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) OpaquePainting() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::opaquePainting")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQuickPaintedItem_OpaquePainting(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQuickPaintedItem) Paint(painter gui.QPainter_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::paint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QQuickPaintedItem) PerformanceHints() QQuickPaintedItem__PerformanceHint {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::performanceHints")
		}
	}()

	if ptr.Pointer() != nil {
		return QQuickPaintedItem__PerformanceHint(C.QQuickPaintedItem_PerformanceHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQuickPaintedItem) ConnectRenderTargetChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::renderTargetChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ConnectRenderTargetChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "renderTargetChanged", f)
	}
}

func (ptr *QQuickPaintedItem) DisconnectRenderTargetChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::renderTargetChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DisconnectRenderTargetChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "renderTargetChanged")
	}
}

//export callbackQQuickPaintedItemRenderTargetChanged
func callbackQQuickPaintedItemRenderTargetChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::renderTargetChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "renderTargetChanged").(func())()
}

func (ptr *QQuickPaintedItem) ResetContentsSize() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::resetContentsSize")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_ResetContentsSize(ptr.Pointer())
	}
}

func (ptr *QQuickPaintedItem) SetAntialiasing(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::setAntialiasing")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetAntialiasing(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QQuickPaintedItem) SetMipmap(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::setMipmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetMipmap(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QQuickPaintedItem) SetOpaquePainting(opaque bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::setOpaquePainting")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetOpaquePainting(ptr.Pointer(), C.int(qt.GoBoolToInt(opaque)))
	}
}

func (ptr *QQuickPaintedItem) SetPerformanceHint(hint QQuickPaintedItem__PerformanceHint, enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::setPerformanceHint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetPerformanceHint(ptr.Pointer(), C.int(hint), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QQuickPaintedItem) SetPerformanceHints(hints QQuickPaintedItem__PerformanceHint) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::setPerformanceHints")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_SetPerformanceHints(ptr.Pointer(), C.int(hints))
	}
}

func (ptr *QQuickPaintedItem) TextureProvider() *QSGTextureProvider {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::textureProvider")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQSGTextureProviderFromPointer(C.QQuickPaintedItem_TextureProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQuickPaintedItem) Update(rect core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::update")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_Update(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QQuickPaintedItem) DestroyQQuickPaintedItem() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQuickPaintedItem::~QQuickPaintedItem")
		}
	}()

	if ptr.Pointer() != nil {
		C.QQuickPaintedItem_DestroyQQuickPaintedItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
