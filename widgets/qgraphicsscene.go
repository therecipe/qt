package widgets

//#include "qgraphicsscene.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsScene struct {
	core.QObject
}

type QGraphicsScene_ITF interface {
	core.QObject_ITF
	QGraphicsScene_PTR() *QGraphicsScene
}

func PointerFromQGraphicsScene(ptr QGraphicsScene_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsScene_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneFromPointer(ptr unsafe.Pointer) *QGraphicsScene {
	var n = new(QGraphicsScene)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QGraphicsScene_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsScene) QGraphicsScene_PTR() *QGraphicsScene {
	return ptr
}

//QGraphicsScene::ItemIndexMethod
type QGraphicsScene__ItemIndexMethod int64

const (
	QGraphicsScene__BspTreeIndex = QGraphicsScene__ItemIndexMethod(0)
	QGraphicsScene__NoIndex      = QGraphicsScene__ItemIndexMethod(-1)
)

//QGraphicsScene::SceneLayer
type QGraphicsScene__SceneLayer int64

const (
	QGraphicsScene__ItemLayer       = QGraphicsScene__SceneLayer(0x1)
	QGraphicsScene__BackgroundLayer = QGraphicsScene__SceneLayer(0x2)
	QGraphicsScene__ForegroundLayer = QGraphicsScene__SceneLayer(0x4)
	QGraphicsScene__AllLayers       = QGraphicsScene__SceneLayer(0xffff)
)

func (ptr *QGraphicsScene) BackgroundBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsScene_BackgroundBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) BspTreeDepth() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsScene_BspTreeDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) ForegroundBrush() *gui.QBrush {
	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsScene_ForegroundBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) IsSortCacheEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_IsSortCacheEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsScene) ItemIndexMethod() QGraphicsScene__ItemIndexMethod {
	if ptr.Pointer() != nil {
		return QGraphicsScene__ItemIndexMethod(C.QGraphicsScene_ItemIndexMethod(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) MinimumRenderSize() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScene_MinimumRenderSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) SetBackgroundBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetBackgroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QGraphicsScene) SetBspTreeDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetBspTreeDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QGraphicsScene) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QGraphicsScene) SetForegroundBrush(brush gui.QBrush_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetForegroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QGraphicsScene) SetItemIndexMethod(method QGraphicsScene__ItemIndexMethod) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetItemIndexMethod(ptr.Pointer(), C.int(method))
	}
}

func (ptr *QGraphicsScene) SetMinimumRenderSize(minSize float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetMinimumRenderSize(ptr.Pointer(), C.double(minSize))
	}
}

func (ptr *QGraphicsScene) SetPalette(palette gui.QPalette_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetPalette(ptr.Pointer(), gui.PointerFromQPalette(palette))
	}
}

func (ptr *QGraphicsScene) SetSceneRect(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSceneRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsScene) SetSortCacheEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSortCacheEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsScene) SetStickyFocus(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetStickyFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsScene) StickyFocus() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_StickyFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsScene) Update(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Update(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func NewQGraphicsScene(parent core.QObject_ITF) *QGraphicsScene {
	return NewQGraphicsSceneFromPointer(C.QGraphicsScene_NewQGraphicsScene(core.PointerFromQObject(parent)))
}

func NewQGraphicsScene2(sceneRect core.QRectF_ITF, parent core.QObject_ITF) *QGraphicsScene {
	return NewQGraphicsSceneFromPointer(C.QGraphicsScene_NewQGraphicsScene2(core.PointerFromQRectF(sceneRect), core.PointerFromQObject(parent)))
}

func NewQGraphicsScene3(x float64, y float64, width float64, height float64, parent core.QObject_ITF) *QGraphicsScene {
	return NewQGraphicsSceneFromPointer(C.QGraphicsScene_NewQGraphicsScene3(C.double(x), C.double(y), C.double(width), C.double(height), core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsScene) ActivePanel() *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_ActivePanel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) ActiveWindow() *QGraphicsWidget {
	if ptr.Pointer() != nil {
		return NewQGraphicsWidgetFromPointer(C.QGraphicsScene_ActiveWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) AddEllipse(rect core.QRectF_ITF, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsEllipseItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsEllipseItemFromPointer(C.QGraphicsScene_AddEllipse(ptr.Pointer(), core.PointerFromQRectF(rect), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddEllipse2(x float64, y float64, w float64, h float64, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsEllipseItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsEllipseItemFromPointer(C.QGraphicsScene_AddEllipse2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddItem(item QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_AddItem(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsScene) AddLine(line core.QLineF_ITF, pen gui.QPen_ITF) *QGraphicsLineItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsLineItemFromPointer(C.QGraphicsScene_AddLine(ptr.Pointer(), core.PointerFromQLineF(line), gui.PointerFromQPen(pen)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddLine2(x1 float64, y1 float64, x2 float64, y2 float64, pen gui.QPen_ITF) *QGraphicsLineItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsLineItemFromPointer(C.QGraphicsScene_AddLine2(ptr.Pointer(), C.double(x1), C.double(y1), C.double(x2), C.double(y2), gui.PointerFromQPen(pen)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddPath(path gui.QPainterPath_ITF, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsPathItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsPathItemFromPointer(C.QGraphicsScene_AddPath(ptr.Pointer(), gui.PointerFromQPainterPath(path), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddPixmap(pixmap gui.QPixmap_ITF) *QGraphicsPixmapItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsPixmapItemFromPointer(C.QGraphicsScene_AddPixmap(ptr.Pointer(), gui.PointerFromQPixmap(pixmap)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddPolygon(polygon gui.QPolygonF_ITF, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsPolygonItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsPolygonItemFromPointer(C.QGraphicsScene_AddPolygon(ptr.Pointer(), gui.PointerFromQPolygonF(polygon), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddRect(rect core.QRectF_ITF, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsRectItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsRectItemFromPointer(C.QGraphicsScene_AddRect(ptr.Pointer(), core.PointerFromQRectF(rect), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddRect2(x float64, y float64, w float64, h float64, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsRectItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsRectItemFromPointer(C.QGraphicsScene_AddRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddSimpleText(text string, font gui.QFont_ITF) *QGraphicsSimpleTextItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsSimpleTextItemFromPointer(C.QGraphicsScene_AddSimpleText(ptr.Pointer(), C.CString(text), gui.PointerFromQFont(font)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddText(text string, font gui.QFont_ITF) *QGraphicsTextItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsTextItemFromPointer(C.QGraphicsScene_AddText(ptr.Pointer(), C.CString(text), gui.PointerFromQFont(font)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddWidget(widget QWidget_ITF, wFlags core.Qt__WindowType) *QGraphicsProxyWidget {
	if ptr.Pointer() != nil {
		return NewQGraphicsProxyWidgetFromPointer(C.QGraphicsScene_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(wFlags)))
	}
	return nil
}

func (ptr *QGraphicsScene) Advance() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Advance(ptr.Pointer())
	}
}

func (ptr *QGraphicsScene) Clear() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Clear(ptr.Pointer())
	}
}

func (ptr *QGraphicsScene) ClearFocus() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_ClearFocus(ptr.Pointer())
	}
}

func (ptr *QGraphicsScene) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QGraphicsScene) DestroyItemGroup(group QGraphicsItemGroup_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_DestroyItemGroup(ptr.Pointer(), PointerFromQGraphicsItemGroup(group))
	}
}

func (ptr *QGraphicsScene) FocusItem() *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_FocusItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) ConnectFocusItemChanged(f func(newFocusItem *QGraphicsItem, oldFocusItem *QGraphicsItem, reason core.Qt__FocusReason)) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_ConnectFocusItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusItemChanged", f)
	}
}

func (ptr *QGraphicsScene) DisconnectFocusItemChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_DisconnectFocusItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusItemChanged")
	}
}

//export callbackQGraphicsSceneFocusItemChanged
func callbackQGraphicsSceneFocusItemChanged(ptrName *C.char, newFocusItem unsafe.Pointer, oldFocusItem unsafe.Pointer, reason C.int) {
	qt.GetSignal(C.GoString(ptrName), "focusItemChanged").(func(*QGraphicsItem, *QGraphicsItem, core.Qt__FocusReason))(NewQGraphicsItemFromPointer(newFocusItem), NewQGraphicsItemFromPointer(oldFocusItem), core.Qt__FocusReason(reason))
}

func (ptr *QGraphicsScene) HasFocus() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsScene) Height() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScene_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsScene_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QGraphicsScene) Invalidate(rect core.QRectF_ITF, layers QGraphicsScene__SceneLayer) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Invalidate(ptr.Pointer(), core.PointerFromQRectF(rect), C.int(layers))
	}
}

func (ptr *QGraphicsScene) Invalidate2(x float64, y float64, w float64, h float64, layers QGraphicsScene__SceneLayer) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Invalidate2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), C.int(layers))
	}
}

func (ptr *QGraphicsScene) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsScene) ItemAt(position core.QPointF_ITF, deviceTransform gui.QTransform_ITF) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_ItemAt(ptr.Pointer(), core.PointerFromQPointF(position), gui.PointerFromQTransform(deviceTransform)))
	}
	return nil
}

func (ptr *QGraphicsScene) ItemAt3(x float64, y float64, deviceTransform gui.QTransform_ITF) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_ItemAt3(ptr.Pointer(), C.double(x), C.double(y), gui.PointerFromQTransform(deviceTransform)))
	}
	return nil
}

func (ptr *QGraphicsScene) MouseGrabberItem() *QGraphicsItem {
	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_MouseGrabberItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) RemoveItem(item QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_RemoveItem(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsScene) Render(painter gui.QPainter_ITF, target core.QRectF_ITF, source core.QRectF_ITF, aspectRatioMode core.Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Render(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(target), core.PointerFromQRectF(source), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsScene) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QGraphicsScene) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQGraphicsSceneSelectionChanged
func callbackQGraphicsSceneSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QGraphicsScene) SendEvent(item QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_SendEvent(ptr.Pointer(), PointerFromQGraphicsItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsScene) SetActivePanel(item QGraphicsItem_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetActivePanel(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsScene) SetActiveWindow(widget QGraphicsWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetActiveWindow(ptr.Pointer(), PointerFromQGraphicsWidget(widget))
	}
}

func (ptr *QGraphicsScene) SetFocus(focusReason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetFocus(ptr.Pointer(), C.int(focusReason))
	}
}

func (ptr *QGraphicsScene) SetFocusItem(item QGraphicsItem_ITF, focusReason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetFocusItem(ptr.Pointer(), PointerFromQGraphicsItem(item), C.int(focusReason))
	}
}

func (ptr *QGraphicsScene) SetSceneRect2(x float64, y float64, w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSceneRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsScene) SetSelectionArea2(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode, deviceTransform gui.QTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSelectionArea2(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.int(mode), gui.PointerFromQTransform(deviceTransform))
	}
}

func (ptr *QGraphicsScene) SetSelectionArea3(path gui.QPainterPath_ITF, selectionOperation core.Qt__ItemSelectionOperation, mode core.Qt__ItemSelectionMode, deviceTransform gui.QTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSelectionArea3(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.int(selectionOperation), C.int(mode), gui.PointerFromQTransform(deviceTransform))
	}
}

func (ptr *QGraphicsScene) SetSelectionArea(path gui.QPainterPath_ITF, deviceTransform gui.QTransform_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSelectionArea(ptr.Pointer(), gui.PointerFromQPainterPath(path), gui.PointerFromQTransform(deviceTransform))
	}
}

func (ptr *QGraphicsScene) SetStyle(style QStyle_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

func (ptr *QGraphicsScene) Style() *QStyle {
	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QGraphicsScene_Style(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) Update2(x float64, y float64, w float64, h float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Update2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsScene) Width() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScene_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) DestroyQGraphicsScene() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_DestroyQGraphicsScene(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
