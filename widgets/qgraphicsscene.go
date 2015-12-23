package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QGraphicsScene_") {
		n.SetObjectName("QGraphicsScene_" + qt.Identifier())
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
	defer qt.Recovering("QGraphicsScene::backgroundBrush")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsScene_BackgroundBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) BspTreeDepth() int {
	defer qt.Recovering("QGraphicsScene::bspTreeDepth")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsScene_BspTreeDepth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) ForegroundBrush() *gui.QBrush {
	defer qt.Recovering("QGraphicsScene::foregroundBrush")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsScene_ForegroundBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) IsSortCacheEnabled() bool {
	defer qt.Recovering("QGraphicsScene::isSortCacheEnabled")

	if ptr.Pointer() != nil {
		return C.QGraphicsScene_IsSortCacheEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsScene) ItemIndexMethod() QGraphicsScene__ItemIndexMethod {
	defer qt.Recovering("QGraphicsScene::itemIndexMethod")

	if ptr.Pointer() != nil {
		return QGraphicsScene__ItemIndexMethod(C.QGraphicsScene_ItemIndexMethod(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) MinimumRenderSize() float64 {
	defer qt.Recovering("QGraphicsScene::minimumRenderSize")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScene_MinimumRenderSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) SetBackgroundBrush(brush gui.QBrush_ITF) {
	defer qt.Recovering("QGraphicsScene::setBackgroundBrush")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetBackgroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QGraphicsScene) SetBspTreeDepth(depth int) {
	defer qt.Recovering("QGraphicsScene::setBspTreeDepth")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetBspTreeDepth(ptr.Pointer(), C.int(depth))
	}
}

func (ptr *QGraphicsScene) SetFont(font gui.QFont_ITF) {
	defer qt.Recovering("QGraphicsScene::setFont")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QGraphicsScene) SetForegroundBrush(brush gui.QBrush_ITF) {
	defer qt.Recovering("QGraphicsScene::setForegroundBrush")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetForegroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QGraphicsScene) SetItemIndexMethod(method QGraphicsScene__ItemIndexMethod) {
	defer qt.Recovering("QGraphicsScene::setItemIndexMethod")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetItemIndexMethod(ptr.Pointer(), C.int(method))
	}
}

func (ptr *QGraphicsScene) SetMinimumRenderSize(minSize float64) {
	defer qt.Recovering("QGraphicsScene::setMinimumRenderSize")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetMinimumRenderSize(ptr.Pointer(), C.double(minSize))
	}
}

func (ptr *QGraphicsScene) SetPalette(palette gui.QPalette_ITF) {
	defer qt.Recovering("QGraphicsScene::setPalette")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetPalette(ptr.Pointer(), gui.PointerFromQPalette(palette))
	}
}

func (ptr *QGraphicsScene) SetSceneRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QGraphicsScene::setSceneRect")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSceneRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsScene) SetSortCacheEnabled(enabled bool) {
	defer qt.Recovering("QGraphicsScene::setSortCacheEnabled")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSortCacheEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsScene) SetStickyFocus(enabled bool) {
	defer qt.Recovering("QGraphicsScene::setStickyFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetStickyFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsScene) StickyFocus() bool {
	defer qt.Recovering("QGraphicsScene::stickyFocus")

	if ptr.Pointer() != nil {
		return C.QGraphicsScene_StickyFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsScene) Update(rect core.QRectF_ITF) {
	defer qt.Recovering("QGraphicsScene::update")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_Update(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func NewQGraphicsScene(parent core.QObject_ITF) *QGraphicsScene {
	defer qt.Recovering("QGraphicsScene::QGraphicsScene")

	return NewQGraphicsSceneFromPointer(C.QGraphicsScene_NewQGraphicsScene(core.PointerFromQObject(parent)))
}

func NewQGraphicsScene2(sceneRect core.QRectF_ITF, parent core.QObject_ITF) *QGraphicsScene {
	defer qt.Recovering("QGraphicsScene::QGraphicsScene")

	return NewQGraphicsSceneFromPointer(C.QGraphicsScene_NewQGraphicsScene2(core.PointerFromQRectF(sceneRect), core.PointerFromQObject(parent)))
}

func NewQGraphicsScene3(x float64, y float64, width float64, height float64, parent core.QObject_ITF) *QGraphicsScene {
	defer qt.Recovering("QGraphicsScene::QGraphicsScene")

	return NewQGraphicsSceneFromPointer(C.QGraphicsScene_NewQGraphicsScene3(C.double(x), C.double(y), C.double(width), C.double(height), core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsScene) ActivePanel() *QGraphicsItem {
	defer qt.Recovering("QGraphicsScene::activePanel")

	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_ActivePanel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) ActiveWindow() *QGraphicsWidget {
	defer qt.Recovering("QGraphicsScene::activeWindow")

	if ptr.Pointer() != nil {
		return NewQGraphicsWidgetFromPointer(C.QGraphicsScene_ActiveWindow(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) AddEllipse(rect core.QRectF_ITF, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsEllipseItem {
	defer qt.Recovering("QGraphicsScene::addEllipse")

	if ptr.Pointer() != nil {
		return NewQGraphicsEllipseItemFromPointer(C.QGraphicsScene_AddEllipse(ptr.Pointer(), core.PointerFromQRectF(rect), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddEllipse2(x float64, y float64, w float64, h float64, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsEllipseItem {
	defer qt.Recovering("QGraphicsScene::addEllipse")

	if ptr.Pointer() != nil {
		return NewQGraphicsEllipseItemFromPointer(C.QGraphicsScene_AddEllipse2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddItem(item QGraphicsItem_ITF) {
	defer qt.Recovering("QGraphicsScene::addItem")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_AddItem(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsScene) AddLine(line core.QLineF_ITF, pen gui.QPen_ITF) *QGraphicsLineItem {
	defer qt.Recovering("QGraphicsScene::addLine")

	if ptr.Pointer() != nil {
		return NewQGraphicsLineItemFromPointer(C.QGraphicsScene_AddLine(ptr.Pointer(), core.PointerFromQLineF(line), gui.PointerFromQPen(pen)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddLine2(x1 float64, y1 float64, x2 float64, y2 float64, pen gui.QPen_ITF) *QGraphicsLineItem {
	defer qt.Recovering("QGraphicsScene::addLine")

	if ptr.Pointer() != nil {
		return NewQGraphicsLineItemFromPointer(C.QGraphicsScene_AddLine2(ptr.Pointer(), C.double(x1), C.double(y1), C.double(x2), C.double(y2), gui.PointerFromQPen(pen)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddPath(path gui.QPainterPath_ITF, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsPathItem {
	defer qt.Recovering("QGraphicsScene::addPath")

	if ptr.Pointer() != nil {
		return NewQGraphicsPathItemFromPointer(C.QGraphicsScene_AddPath(ptr.Pointer(), gui.PointerFromQPainterPath(path), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddPixmap(pixmap gui.QPixmap_ITF) *QGraphicsPixmapItem {
	defer qt.Recovering("QGraphicsScene::addPixmap")

	if ptr.Pointer() != nil {
		return NewQGraphicsPixmapItemFromPointer(C.QGraphicsScene_AddPixmap(ptr.Pointer(), gui.PointerFromQPixmap(pixmap)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddPolygon(polygon gui.QPolygonF_ITF, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsPolygonItem {
	defer qt.Recovering("QGraphicsScene::addPolygon")

	if ptr.Pointer() != nil {
		return NewQGraphicsPolygonItemFromPointer(C.QGraphicsScene_AddPolygon(ptr.Pointer(), gui.PointerFromQPolygonF(polygon), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddRect(rect core.QRectF_ITF, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsRectItem {
	defer qt.Recovering("QGraphicsScene::addRect")

	if ptr.Pointer() != nil {
		return NewQGraphicsRectItemFromPointer(C.QGraphicsScene_AddRect(ptr.Pointer(), core.PointerFromQRectF(rect), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddRect2(x float64, y float64, w float64, h float64, pen gui.QPen_ITF, brush gui.QBrush_ITF) *QGraphicsRectItem {
	defer qt.Recovering("QGraphicsScene::addRect")

	if ptr.Pointer() != nil {
		return NewQGraphicsRectItemFromPointer(C.QGraphicsScene_AddRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), gui.PointerFromQPen(pen), gui.PointerFromQBrush(brush)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddSimpleText(text string, font gui.QFont_ITF) *QGraphicsSimpleTextItem {
	defer qt.Recovering("QGraphicsScene::addSimpleText")

	if ptr.Pointer() != nil {
		return NewQGraphicsSimpleTextItemFromPointer(C.QGraphicsScene_AddSimpleText(ptr.Pointer(), C.CString(text), gui.PointerFromQFont(font)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddText(text string, font gui.QFont_ITF) *QGraphicsTextItem {
	defer qt.Recovering("QGraphicsScene::addText")

	if ptr.Pointer() != nil {
		return NewQGraphicsTextItemFromPointer(C.QGraphicsScene_AddText(ptr.Pointer(), C.CString(text), gui.PointerFromQFont(font)))
	}
	return nil
}

func (ptr *QGraphicsScene) AddWidget(widget QWidget_ITF, wFlags core.Qt__WindowType) *QGraphicsProxyWidget {
	defer qt.Recovering("QGraphicsScene::addWidget")

	if ptr.Pointer() != nil {
		return NewQGraphicsProxyWidgetFromPointer(C.QGraphicsScene_AddWidget(ptr.Pointer(), PointerFromQWidget(widget), C.int(wFlags)))
	}
	return nil
}

func (ptr *QGraphicsScene) Advance() {
	defer qt.Recovering("QGraphicsScene::advance")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_Advance(ptr.Pointer())
	}
}

func (ptr *QGraphicsScene) Clear() {
	defer qt.Recovering("QGraphicsScene::clear")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_Clear(ptr.Pointer())
	}
}

func (ptr *QGraphicsScene) ClearFocus() {
	defer qt.Recovering("QGraphicsScene::clearFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_ClearFocus(ptr.Pointer())
	}
}

func (ptr *QGraphicsScene) ClearSelection() {
	defer qt.Recovering("QGraphicsScene::clearSelection")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_ClearSelection(ptr.Pointer())
	}
}

func (ptr *QGraphicsScene) ConnectContextMenuEvent(f func(contextMenuEvent *QGraphicsSceneContextMenuEvent)) {
	defer qt.Recovering("connect QGraphicsScene::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQGraphicsSceneContextMenuEvent
func callbackQGraphicsSceneContextMenuEvent(ptrName *C.char, contextMenuEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*QGraphicsSceneContextMenuEvent))(NewQGraphicsSceneContextMenuEventFromPointer(contextMenuEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) DestroyItemGroup(group QGraphicsItemGroup_ITF) {
	defer qt.Recovering("QGraphicsScene::destroyItemGroup")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_DestroyItemGroup(ptr.Pointer(), PointerFromQGraphicsItemGroup(group))
	}
}

func (ptr *QGraphicsScene) ConnectDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsScene::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQGraphicsSceneDragEnterEvent
func callbackQGraphicsSceneDragEnterEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsScene::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQGraphicsSceneDragLeaveEvent
func callbackQGraphicsSceneDragLeaveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsScene::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQGraphicsSceneDragMoveEvent
func callbackQGraphicsSceneDragMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectDropEvent(f func(event *QGraphicsSceneDragDropEvent)) {
	defer qt.Recovering("connect QGraphicsScene::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQGraphicsSceneDropEvent
func callbackQGraphicsSceneDropEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*QGraphicsSceneDragDropEvent))(NewQGraphicsSceneDragDropEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectFocusInEvent(f func(focusEvent *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsScene::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQGraphicsSceneFocusInEvent
func callbackQGraphicsSceneFocusInEvent(ptrName *C.char, focusEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(focusEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) FocusItem() *QGraphicsItem {
	defer qt.Recovering("QGraphicsScene::focusItem")

	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_FocusItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) ConnectFocusItemChanged(f func(newFocusItem *QGraphicsItem, oldFocusItem *QGraphicsItem, reason core.Qt__FocusReason)) {
	defer qt.Recovering("connect QGraphicsScene::focusItemChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_ConnectFocusItemChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "focusItemChanged", f)
	}
}

func (ptr *QGraphicsScene) DisconnectFocusItemChanged() {
	defer qt.Recovering("disconnect QGraphicsScene::focusItemChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_DisconnectFocusItemChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "focusItemChanged")
	}
}

//export callbackQGraphicsSceneFocusItemChanged
func callbackQGraphicsSceneFocusItemChanged(ptrName *C.char, newFocusItem unsafe.Pointer, oldFocusItem unsafe.Pointer, reason C.int) {
	defer qt.Recovering("callback QGraphicsScene::focusItemChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusItemChanged"); signal != nil {
		signal.(func(*QGraphicsItem, *QGraphicsItem, core.Qt__FocusReason))(NewQGraphicsItemFromPointer(newFocusItem), NewQGraphicsItemFromPointer(oldFocusItem), core.Qt__FocusReason(reason))
	}

}

func (ptr *QGraphicsScene) ConnectFocusOutEvent(f func(focusEvent *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsScene::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQGraphicsSceneFocusOutEvent
func callbackQGraphicsSceneFocusOutEvent(ptrName *C.char, focusEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(focusEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) HasFocus() bool {
	defer qt.Recovering("QGraphicsScene::hasFocus")

	if ptr.Pointer() != nil {
		return C.QGraphicsScene_HasFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsScene) Height() float64 {
	defer qt.Recovering("QGraphicsScene::height")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScene_Height(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) ConnectHelpEvent(f func(helpEvent *QGraphicsSceneHelpEvent)) {
	defer qt.Recovering("connect QGraphicsScene::helpEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "helpEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectHelpEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::helpEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "helpEvent")
	}
}

//export callbackQGraphicsSceneHelpEvent
func callbackQGraphicsSceneHelpEvent(ptrName *C.char, helpEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::helpEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "helpEvent"); signal != nil {
		signal.(func(*QGraphicsSceneHelpEvent))(NewQGraphicsSceneHelpEventFromPointer(helpEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QGraphicsScene::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQGraphicsSceneInputMethodEvent
func callbackQGraphicsSceneInputMethodEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsScene::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsScene_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QGraphicsScene) Invalidate(rect core.QRectF_ITF, layers QGraphicsScene__SceneLayer) {
	defer qt.Recovering("QGraphicsScene::invalidate")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_Invalidate(ptr.Pointer(), core.PointerFromQRectF(rect), C.int(layers))
	}
}

func (ptr *QGraphicsScene) Invalidate2(x float64, y float64, w float64, h float64, layers QGraphicsScene__SceneLayer) {
	defer qt.Recovering("QGraphicsScene::invalidate")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_Invalidate2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), C.int(layers))
	}
}

func (ptr *QGraphicsScene) IsActive() bool {
	defer qt.Recovering("QGraphicsScene::isActive")

	if ptr.Pointer() != nil {
		return C.QGraphicsScene_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsScene) ItemAt(position core.QPointF_ITF, deviceTransform gui.QTransform_ITF) *QGraphicsItem {
	defer qt.Recovering("QGraphicsScene::itemAt")

	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_ItemAt(ptr.Pointer(), core.PointerFromQPointF(position), gui.PointerFromQTransform(deviceTransform)))
	}
	return nil
}

func (ptr *QGraphicsScene) ItemAt3(x float64, y float64, deviceTransform gui.QTransform_ITF) *QGraphicsItem {
	defer qt.Recovering("QGraphicsScene::itemAt")

	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_ItemAt3(ptr.Pointer(), C.double(x), C.double(y), gui.PointerFromQTransform(deviceTransform)))
	}
	return nil
}

func (ptr *QGraphicsScene) ConnectKeyPressEvent(f func(keyEvent *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsScene::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQGraphicsSceneKeyPressEvent
func callbackQGraphicsSceneKeyPressEvent(ptrName *C.char, keyEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(keyEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectKeyReleaseEvent(f func(keyEvent *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsScene::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQGraphicsSceneKeyReleaseEvent
func callbackQGraphicsSceneKeyReleaseEvent(ptrName *C.char, keyEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(keyEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectMouseDoubleClickEvent(f func(mouseEvent *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsScene::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQGraphicsSceneMouseDoubleClickEvent
func callbackQGraphicsSceneMouseDoubleClickEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) MouseGrabberItem() *QGraphicsItem {
	defer qt.Recovering("QGraphicsScene::mouseGrabberItem")

	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsScene_MouseGrabberItem(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) ConnectMouseMoveEvent(f func(mouseEvent *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsScene::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQGraphicsSceneMouseMoveEvent
func callbackQGraphicsSceneMouseMoveEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectMousePressEvent(f func(mouseEvent *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsScene::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQGraphicsSceneMousePressEvent
func callbackQGraphicsSceneMousePressEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectMouseReleaseEvent(f func(mouseEvent *QGraphicsSceneMouseEvent)) {
	defer qt.Recovering("connect QGraphicsScene::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQGraphicsSceneMouseReleaseEvent
func callbackQGraphicsSceneMouseReleaseEvent(ptrName *C.char, mouseEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*QGraphicsSceneMouseEvent))(NewQGraphicsSceneMouseEventFromPointer(mouseEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) RemoveItem(item QGraphicsItem_ITF) {
	defer qt.Recovering("QGraphicsScene::removeItem")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_RemoveItem(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsScene) Render(painter gui.QPainter_ITF, target core.QRectF_ITF, source core.QRectF_ITF, aspectRatioMode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QGraphicsScene::render")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_Render(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(target), core.PointerFromQRectF(source), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsScene) ConnectSelectionChanged(f func()) {
	defer qt.Recovering("connect QGraphicsScene::selectionChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_ConnectSelectionChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectionChanged", f)
	}
}

func (ptr *QGraphicsScene) DisconnectSelectionChanged() {
	defer qt.Recovering("disconnect QGraphicsScene::selectionChanged")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_DisconnectSelectionChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectionChanged")
	}
}

//export callbackQGraphicsSceneSelectionChanged
func callbackQGraphicsSceneSelectionChanged(ptrName *C.char) {
	defer qt.Recovering("callback QGraphicsScene::selectionChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectionChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QGraphicsScene) SendEvent(item QGraphicsItem_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsScene::sendEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsScene_SendEvent(ptr.Pointer(), PointerFromQGraphicsItem(item), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsScene) SetActivePanel(item QGraphicsItem_ITF) {
	defer qt.Recovering("QGraphicsScene::setActivePanel")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetActivePanel(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsScene) SetActiveWindow(widget QGraphicsWidget_ITF) {
	defer qt.Recovering("QGraphicsScene::setActiveWindow")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetActiveWindow(ptr.Pointer(), PointerFromQGraphicsWidget(widget))
	}
}

func (ptr *QGraphicsScene) SetFocus(focusReason core.Qt__FocusReason) {
	defer qt.Recovering("QGraphicsScene::setFocus")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetFocus(ptr.Pointer(), C.int(focusReason))
	}
}

func (ptr *QGraphicsScene) SetFocusItem(item QGraphicsItem_ITF, focusReason core.Qt__FocusReason) {
	defer qt.Recovering("QGraphicsScene::setFocusItem")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetFocusItem(ptr.Pointer(), PointerFromQGraphicsItem(item), C.int(focusReason))
	}
}

func (ptr *QGraphicsScene) SetSceneRect2(x float64, y float64, w float64, h float64) {
	defer qt.Recovering("QGraphicsScene::setSceneRect")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSceneRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsScene) SetSelectionArea2(path gui.QPainterPath_ITF, mode core.Qt__ItemSelectionMode, deviceTransform gui.QTransform_ITF) {
	defer qt.Recovering("QGraphicsScene::setSelectionArea")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSelectionArea2(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.int(mode), gui.PointerFromQTransform(deviceTransform))
	}
}

func (ptr *QGraphicsScene) SetSelectionArea3(path gui.QPainterPath_ITF, selectionOperation core.Qt__ItemSelectionOperation, mode core.Qt__ItemSelectionMode, deviceTransform gui.QTransform_ITF) {
	defer qt.Recovering("QGraphicsScene::setSelectionArea")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSelectionArea3(ptr.Pointer(), gui.PointerFromQPainterPath(path), C.int(selectionOperation), C.int(mode), gui.PointerFromQTransform(deviceTransform))
	}
}

func (ptr *QGraphicsScene) SetSelectionArea(path gui.QPainterPath_ITF, deviceTransform gui.QTransform_ITF) {
	defer qt.Recovering("QGraphicsScene::setSelectionArea")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetSelectionArea(ptr.Pointer(), gui.PointerFromQPainterPath(path), gui.PointerFromQTransform(deviceTransform))
	}
}

func (ptr *QGraphicsScene) SetStyle(style QStyle_ITF) {
	defer qt.Recovering("QGraphicsScene::setStyle")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_SetStyle(ptr.Pointer(), PointerFromQStyle(style))
	}
}

func (ptr *QGraphicsScene) Style() *QStyle {
	defer qt.Recovering("QGraphicsScene::style")

	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QGraphicsScene_Style(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsScene) Update2(x float64, y float64, w float64, h float64) {
	defer qt.Recovering("QGraphicsScene::update")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_Update2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsScene) ConnectWheelEvent(f func(wheelEvent *QGraphicsSceneWheelEvent)) {
	defer qt.Recovering("connect QGraphicsScene::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQGraphicsSceneWheelEvent
func callbackQGraphicsSceneWheelEvent(ptrName *C.char, wheelEvent unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*QGraphicsSceneWheelEvent))(NewQGraphicsSceneWheelEventFromPointer(wheelEvent))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) Width() float64 {
	defer qt.Recovering("QGraphicsScene::width")

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsScene_Width(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsScene) DestroyQGraphicsScene() {
	defer qt.Recovering("QGraphicsScene::~QGraphicsScene")

	if ptr.Pointer() != nil {
		C.QGraphicsScene_DestroyQGraphicsScene(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsScene) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsScene::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsSceneTimerEvent
func callbackQGraphicsSceneTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsScene::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsSceneChildEvent
func callbackQGraphicsSceneChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGraphicsScene) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsScene::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsScene) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsScene::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsSceneCustomEvent
func callbackQGraphicsSceneCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsScene::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
