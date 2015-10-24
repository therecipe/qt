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

type QGraphicsSceneITF interface {
	core.QObjectITF
	QGraphicsScenePTR() *QGraphicsScene
}

func PointerFromQGraphicsScene(ptr QGraphicsSceneITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsScenePTR().Pointer()
	}
	return nil
}

func QGraphicsSceneFromPointer(ptr unsafe.Pointer) *QGraphicsScene {
	var n = new(QGraphicsScene)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsScene_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsScene) QGraphicsScenePTR() *QGraphicsScene {
	return ptr
}

//QGraphicsScene::ItemIndexMethod
type QGraphicsScene__ItemIndexMethod int

var (
	QGraphicsScene__BspTreeIndex = QGraphicsScene__ItemIndexMethod(0)
	QGraphicsScene__NoIndex      = QGraphicsScene__ItemIndexMethod(-1)
)

//QGraphicsScene::SceneLayer
type QGraphicsScene__SceneLayer int

var (
	QGraphicsScene__ItemLayer       = QGraphicsScene__SceneLayer(0x1)
	QGraphicsScene__BackgroundLayer = QGraphicsScene__SceneLayer(0x2)
	QGraphicsScene__ForegroundLayer = QGraphicsScene__SceneLayer(0x4)
	QGraphicsScene__AllLayers       = QGraphicsScene__SceneLayer(0xffff)
)

func (ptr *QGraphicsScene) BspTreeDepth() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsScene_BspTreeDepth(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsScene) IsSortCacheEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_IsSortCacheEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsScene) ItemIndexMethod() QGraphicsScene__ItemIndexMethod {
	if ptr.Pointer() != nil {
		return QGraphicsScene__ItemIndexMethod(C.QGraphicsScene_ItemIndexMethod(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsScene) SetBackgroundBrush(brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetBackgroundBrush(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QGraphicsScene) SetBspTreeDepth(depth int) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetBspTreeDepth(C.QtObjectPtr(ptr.Pointer()), C.int(depth))
	}
}

func (ptr *QGraphicsScene) SetFont(font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QGraphicsScene) SetForegroundBrush(brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetForegroundBrush(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QGraphicsScene) SetItemIndexMethod(method QGraphicsScene__ItemIndexMethod) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetItemIndexMethod(C.QtObjectPtr(ptr.Pointer()), C.int(method))
	}
}

func (ptr *QGraphicsScene) SetPalette(palette gui.QPaletteITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetPalette(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPalette(palette)))
	}
}

func (ptr *QGraphicsScene) SetSceneRect(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSceneRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QGraphicsScene) SetSortCacheEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSortCacheEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsScene) SetStickyFocus(enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetStickyFocus(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsScene) StickyFocus() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_StickyFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsScene) Update(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Update(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func NewQGraphicsScene(parent core.QObjectITF) *QGraphicsScene {
	return QGraphicsSceneFromPointer(unsafe.Pointer(C.QGraphicsScene_NewQGraphicsScene(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQGraphicsScene2(sceneRect core.QRectFITF, parent core.QObjectITF) *QGraphicsScene {
	return QGraphicsSceneFromPointer(unsafe.Pointer(C.QGraphicsScene_NewQGraphicsScene2(C.QtObjectPtr(core.PointerFromQRectF(sceneRect)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGraphicsScene) ActivePanel() *QGraphicsItem {
	if ptr.Pointer() != nil {
		return QGraphicsItemFromPointer(unsafe.Pointer(C.QGraphicsScene_ActivePanel(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsScene) ActiveWindow() *QGraphicsWidget {
	if ptr.Pointer() != nil {
		return QGraphicsWidgetFromPointer(unsafe.Pointer(C.QGraphicsScene_ActiveWindow(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsScene) AddEllipse(rect core.QRectFITF, pen gui.QPenITF, brush gui.QBrushITF) *QGraphicsEllipseItem {
	if ptr.Pointer() != nil {
		return QGraphicsEllipseItemFromPointer(unsafe.Pointer(C.QGraphicsScene_AddEllipse(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)), C.QtObjectPtr(gui.PointerFromQPen(pen)), C.QtObjectPtr(gui.PointerFromQBrush(brush)))))
	}
	return nil
}

func (ptr *QGraphicsScene) AddItem(item QGraphicsItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_AddItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)))
	}
}

func (ptr *QGraphicsScene) AddLine(line core.QLineFITF, pen gui.QPenITF) *QGraphicsLineItem {
	if ptr.Pointer() != nil {
		return QGraphicsLineItemFromPointer(unsafe.Pointer(C.QGraphicsScene_AddLine(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLineF(line)), C.QtObjectPtr(gui.PointerFromQPen(pen)))))
	}
	return nil
}

func (ptr *QGraphicsScene) AddPath(path gui.QPainterPathITF, pen gui.QPenITF, brush gui.QBrushITF) *QGraphicsPathItem {
	if ptr.Pointer() != nil {
		return QGraphicsPathItemFromPointer(unsafe.Pointer(C.QGraphicsScene_AddPath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainterPath(path)), C.QtObjectPtr(gui.PointerFromQPen(pen)), C.QtObjectPtr(gui.PointerFromQBrush(brush)))))
	}
	return nil
}

func (ptr *QGraphicsScene) AddPixmap(pixmap gui.QPixmapITF) *QGraphicsPixmapItem {
	if ptr.Pointer() != nil {
		return QGraphicsPixmapItemFromPointer(unsafe.Pointer(C.QGraphicsScene_AddPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPixmap(pixmap)))))
	}
	return nil
}

func (ptr *QGraphicsScene) AddPolygon(polygon gui.QPolygonFITF, pen gui.QPenITF, brush gui.QBrushITF) *QGraphicsPolygonItem {
	if ptr.Pointer() != nil {
		return QGraphicsPolygonItemFromPointer(unsafe.Pointer(C.QGraphicsScene_AddPolygon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPolygonF(polygon)), C.QtObjectPtr(gui.PointerFromQPen(pen)), C.QtObjectPtr(gui.PointerFromQBrush(brush)))))
	}
	return nil
}

func (ptr *QGraphicsScene) AddRect(rect core.QRectFITF, pen gui.QPenITF, brush gui.QBrushITF) *QGraphicsRectItem {
	if ptr.Pointer() != nil {
		return QGraphicsRectItemFromPointer(unsafe.Pointer(C.QGraphicsScene_AddRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)), C.QtObjectPtr(gui.PointerFromQPen(pen)), C.QtObjectPtr(gui.PointerFromQBrush(brush)))))
	}
	return nil
}

func (ptr *QGraphicsScene) AddSimpleText(text string, font gui.QFontITF) *QGraphicsSimpleTextItem {
	if ptr.Pointer() != nil {
		return QGraphicsSimpleTextItemFromPointer(unsafe.Pointer(C.QGraphicsScene_AddSimpleText(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.QtObjectPtr(gui.PointerFromQFont(font)))))
	}
	return nil
}

func (ptr *QGraphicsScene) AddText(text string, font gui.QFontITF) *QGraphicsTextItem {
	if ptr.Pointer() != nil {
		return QGraphicsTextItemFromPointer(unsafe.Pointer(C.QGraphicsScene_AddText(C.QtObjectPtr(ptr.Pointer()), C.CString(text), C.QtObjectPtr(gui.PointerFromQFont(font)))))
	}
	return nil
}

func (ptr *QGraphicsScene) AddWidget(widget QWidgetITF, wFlags core.Qt__WindowType) *QGraphicsProxyWidget {
	if ptr.Pointer() != nil {
		return QGraphicsProxyWidgetFromPointer(unsafe.Pointer(C.QGraphicsScene_AddWidget(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(widget)), C.int(wFlags))))
	}
	return nil
}

func (ptr *QGraphicsScene) Advance() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Advance(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsScene) Clear() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsScene) ClearFocus() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_ClearFocus(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsScene) ClearSelection() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_ClearSelection(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsScene) DestroyItemGroup(group QGraphicsItemGroupITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_DestroyItemGroup(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItemGroup(group)))
	}
}

func (ptr *QGraphicsScene) FocusItem() *QGraphicsItem {
	if ptr.Pointer() != nil {
		return QGraphicsItemFromPointer(unsafe.Pointer(C.QGraphicsScene_FocusItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsScene) ConnectFocusItemChanged(f func(newFocusItem QGraphicsItemITF, oldFocusItem QGraphicsItemITF, reason core.Qt__FocusReason)) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_ConnectFocusItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "focusItemChanged", f)
	}
}

func (ptr *QGraphicsScene) DisconnectFocusItemChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_DisconnectFocusItemChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "focusItemChanged")
	}
}

//export callbackQGraphicsSceneFocusItemChanged
func callbackQGraphicsSceneFocusItemChanged(ptrName *C.char, newFocusItem unsafe.Pointer, oldFocusItem unsafe.Pointer, reason C.int) {
	qt.GetSignal(C.GoString(ptrName), "focusItemChanged").(func(*QGraphicsItem, *QGraphicsItem, core.Qt__FocusReason))(QGraphicsItemFromPointer(newFocusItem), QGraphicsItemFromPointer(oldFocusItem), core.Qt__FocusReason(reason))
}

func (ptr *QGraphicsScene) HasFocus() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_HasFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsScene) InputMethodQuery(query core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsScene_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(query)))
	}
	return ""
}

func (ptr *QGraphicsScene) Invalidate(rect core.QRectFITF, layers QGraphicsScene__SceneLayer) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Invalidate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)), C.int(layers))
	}
}

func (ptr *QGraphicsScene) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsScene) ItemAt(position core.QPointFITF, deviceTransform gui.QTransformITF) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return QGraphicsItemFromPointer(unsafe.Pointer(C.QGraphicsScene_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(position)), C.QtObjectPtr(gui.PointerFromQTransform(deviceTransform)))))
	}
	return nil
}

func (ptr *QGraphicsScene) MouseGrabberItem() *QGraphicsItem {
	if ptr.Pointer() != nil {
		return QGraphicsItemFromPointer(unsafe.Pointer(C.QGraphicsScene_MouseGrabberItem(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsScene) RemoveItem(item QGraphicsItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_RemoveItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)))
	}
}

func (ptr *QGraphicsScene) Render(painter gui.QPainterITF, target core.QRectFITF, source core.QRectFITF, aspectRatioMode core.Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_Render(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQRectF(target)), C.QtObjectPtr(core.PointerFromQRectF(source)), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsScene) ConnectSelectionChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_ConnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QGraphicsScene) DisconnectSelectionChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_DisconnectSelectionChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQGraphicsSceneSelectionChanged
func callbackQGraphicsSceneSelectionChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "selectionChanged").(func())()
}

func (ptr *QGraphicsScene) SendEvent(item QGraphicsItemITF, event core.QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsScene_SendEvent(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)), C.QtObjectPtr(core.PointerFromQEvent(event))) != 0
	}
	return false
}

func (ptr *QGraphicsScene) SetActivePanel(item QGraphicsItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetActivePanel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)))
	}
}

func (ptr *QGraphicsScene) SetActiveWindow(widget QGraphicsWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetActiveWindow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsWidget(widget)))
	}
}

func (ptr *QGraphicsScene) SetFocus(focusReason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetFocus(C.QtObjectPtr(ptr.Pointer()), C.int(focusReason))
	}
}

func (ptr *QGraphicsScene) SetFocusItem(item QGraphicsItemITF, focusReason core.Qt__FocusReason) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetFocusItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)), C.int(focusReason))
	}
}

func (ptr *QGraphicsScene) SetSelectionArea2(path gui.QPainterPathITF, mode core.Qt__ItemSelectionMode, deviceTransform gui.QTransformITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSelectionArea2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainterPath(path)), C.int(mode), C.QtObjectPtr(gui.PointerFromQTransform(deviceTransform)))
	}
}

func (ptr *QGraphicsScene) SetSelectionArea3(path gui.QPainterPathITF, selectionOperation core.Qt__ItemSelectionOperation, mode core.Qt__ItemSelectionMode, deviceTransform gui.QTransformITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSelectionArea3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainterPath(path)), C.int(selectionOperation), C.int(mode), C.QtObjectPtr(gui.PointerFromQTransform(deviceTransform)))
	}
}

func (ptr *QGraphicsScene) SetSelectionArea(path gui.QPainterPathITF, deviceTransform gui.QTransformITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSelectionArea(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainterPath(path)), C.QtObjectPtr(gui.PointerFromQTransform(deviceTransform)))
	}
}

func (ptr *QGraphicsScene) SetStyle(style QStyleITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetStyle(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQStyle(style)))
	}
}

func (ptr *QGraphicsScene) Style() *QStyle {
	if ptr.Pointer() != nil {
		return QStyleFromPointer(unsafe.Pointer(C.QGraphicsScene_Style(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsScene) DestroyQGraphicsScene() {
	if ptr.Pointer() != nil {
		C.QGraphicsScene_DestroyQGraphicsScene(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
