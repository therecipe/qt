package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsView struct {
	QAbstractScrollArea
}

type QGraphicsView_ITF interface {
	QAbstractScrollArea_ITF
	QGraphicsView_PTR() *QGraphicsView
}

func PointerFromQGraphicsView(ptr QGraphicsView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsView_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsViewFromPointer(ptr unsafe.Pointer) *QGraphicsView {
	var n = new(QGraphicsView)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsView_") {
		n.SetObjectName("QGraphicsView_" + qt.Identifier())
	}
	return n
}

func (ptr *QGraphicsView) QGraphicsView_PTR() *QGraphicsView {
	return ptr
}

//QGraphicsView::CacheModeFlag
type QGraphicsView__CacheModeFlag int64

const (
	QGraphicsView__CacheNone       = QGraphicsView__CacheModeFlag(0x0)
	QGraphicsView__CacheBackground = QGraphicsView__CacheModeFlag(0x1)
)

//QGraphicsView::DragMode
type QGraphicsView__DragMode int64

const (
	QGraphicsView__NoDrag         = QGraphicsView__DragMode(0)
	QGraphicsView__ScrollHandDrag = QGraphicsView__DragMode(1)
	QGraphicsView__RubberBandDrag = QGraphicsView__DragMode(2)
)

//QGraphicsView::OptimizationFlag
type QGraphicsView__OptimizationFlag int64

const (
	QGraphicsView__DontClipPainter           = QGraphicsView__OptimizationFlag(0x1)
	QGraphicsView__DontSavePainterState      = QGraphicsView__OptimizationFlag(0x2)
	QGraphicsView__DontAdjustForAntialiasing = QGraphicsView__OptimizationFlag(0x4)
	QGraphicsView__IndirectPainting          = QGraphicsView__OptimizationFlag(0x8)
)

//QGraphicsView::ViewportAnchor
type QGraphicsView__ViewportAnchor int64

const (
	QGraphicsView__NoAnchor         = QGraphicsView__ViewportAnchor(0)
	QGraphicsView__AnchorViewCenter = QGraphicsView__ViewportAnchor(1)
	QGraphicsView__AnchorUnderMouse = QGraphicsView__ViewportAnchor(2)
)

//QGraphicsView::ViewportUpdateMode
type QGraphicsView__ViewportUpdateMode int64

const (
	QGraphicsView__FullViewportUpdate         = QGraphicsView__ViewportUpdateMode(0)
	QGraphicsView__MinimalViewportUpdate      = QGraphicsView__ViewportUpdateMode(1)
	QGraphicsView__SmartViewportUpdate        = QGraphicsView__ViewportUpdateMode(2)
	QGraphicsView__NoViewportUpdate           = QGraphicsView__ViewportUpdateMode(3)
	QGraphicsView__BoundingRectViewportUpdate = QGraphicsView__ViewportUpdateMode(4)
)

func (ptr *QGraphicsView) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QGraphicsView::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsView_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) BackgroundBrush() *gui.QBrush {
	defer qt.Recovering("QGraphicsView::backgroundBrush")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsView_BackgroundBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsView) CacheMode() QGraphicsView__CacheModeFlag {
	defer qt.Recovering("QGraphicsView::cacheMode")

	if ptr.Pointer() != nil {
		return QGraphicsView__CacheModeFlag(C.QGraphicsView_CacheMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) DragMode() QGraphicsView__DragMode {
	defer qt.Recovering("QGraphicsView::dragMode")

	if ptr.Pointer() != nil {
		return QGraphicsView__DragMode(C.QGraphicsView_DragMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) ForegroundBrush() *gui.QBrush {
	defer qt.Recovering("QGraphicsView::foregroundBrush")

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsView_ForegroundBrush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsView) IsInteractive() bool {
	defer qt.Recovering("QGraphicsView::isInteractive")

	if ptr.Pointer() != nil {
		return C.QGraphicsView_IsInteractive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsView) OptimizationFlags() QGraphicsView__OptimizationFlag {
	defer qt.Recovering("QGraphicsView::optimizationFlags")

	if ptr.Pointer() != nil {
		return QGraphicsView__OptimizationFlag(C.QGraphicsView_OptimizationFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) RenderHints() gui.QPainter__RenderHint {
	defer qt.Recovering("QGraphicsView::renderHints")

	if ptr.Pointer() != nil {
		return gui.QPainter__RenderHint(C.QGraphicsView_RenderHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) ResizeAnchor() QGraphicsView__ViewportAnchor {
	defer qt.Recovering("QGraphicsView::resizeAnchor")

	if ptr.Pointer() != nil {
		return QGraphicsView__ViewportAnchor(C.QGraphicsView_ResizeAnchor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) RubberBandSelectionMode() core.Qt__ItemSelectionMode {
	defer qt.Recovering("QGraphicsView::rubberBandSelectionMode")

	if ptr.Pointer() != nil {
		return core.Qt__ItemSelectionMode(C.QGraphicsView_RubberBandSelectionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) SetAlignment(alignment core.Qt__AlignmentFlag) {
	defer qt.Recovering("QGraphicsView::setAlignment")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QGraphicsView) SetBackgroundBrush(brush gui.QBrush_ITF) {
	defer qt.Recovering("QGraphicsView::setBackgroundBrush")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetBackgroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QGraphicsView) SetCacheMode(mode QGraphicsView__CacheModeFlag) {
	defer qt.Recovering("QGraphicsView::setCacheMode")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetCacheMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsView) SetDragMode(mode QGraphicsView__DragMode) {
	defer qt.Recovering("QGraphicsView::setDragMode")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetDragMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsView) SetForegroundBrush(brush gui.QBrush_ITF) {
	defer qt.Recovering("QGraphicsView::setForegroundBrush")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetForegroundBrush(ptr.Pointer(), gui.PointerFromQBrush(brush))
	}
}

func (ptr *QGraphicsView) SetInteractive(allowed bool) {
	defer qt.Recovering("QGraphicsView::setInteractive")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetInteractive(ptr.Pointer(), C.int(qt.GoBoolToInt(allowed)))
	}
}

func (ptr *QGraphicsView) SetOptimizationFlags(flags QGraphicsView__OptimizationFlag) {
	defer qt.Recovering("QGraphicsView::setOptimizationFlags")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetOptimizationFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QGraphicsView) SetRenderHints(hints gui.QPainter__RenderHint) {
	defer qt.Recovering("QGraphicsView::setRenderHints")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetRenderHints(ptr.Pointer(), C.int(hints))
	}
}

func (ptr *QGraphicsView) SetResizeAnchor(anchor QGraphicsView__ViewportAnchor) {
	defer qt.Recovering("QGraphicsView::setResizeAnchor")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetResizeAnchor(ptr.Pointer(), C.int(anchor))
	}
}

func (ptr *QGraphicsView) SetRubberBandSelectionMode(mode core.Qt__ItemSelectionMode) {
	defer qt.Recovering("QGraphicsView::setRubberBandSelectionMode")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetRubberBandSelectionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsView) SetSceneRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QGraphicsView::setSceneRect")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetSceneRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsView) SetTransformationAnchor(anchor QGraphicsView__ViewportAnchor) {
	defer qt.Recovering("QGraphicsView::setTransformationAnchor")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetTransformationAnchor(ptr.Pointer(), C.int(anchor))
	}
}

func (ptr *QGraphicsView) SetViewportUpdateMode(mode QGraphicsView__ViewportUpdateMode) {
	defer qt.Recovering("QGraphicsView::setViewportUpdateMode")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetViewportUpdateMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QGraphicsView) TransformationAnchor() QGraphicsView__ViewportAnchor {
	defer qt.Recovering("QGraphicsView::transformationAnchor")

	if ptr.Pointer() != nil {
		return QGraphicsView__ViewportAnchor(C.QGraphicsView_TransformationAnchor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsView) ViewportUpdateMode() QGraphicsView__ViewportUpdateMode {
	defer qt.Recovering("QGraphicsView::viewportUpdateMode")

	if ptr.Pointer() != nil {
		return QGraphicsView__ViewportUpdateMode(C.QGraphicsView_ViewportUpdateMode(ptr.Pointer()))
	}
	return 0
}

func NewQGraphicsView2(scene QGraphicsScene_ITF, parent QWidget_ITF) *QGraphicsView {
	defer qt.Recovering("QGraphicsView::QGraphicsView")

	return NewQGraphicsViewFromPointer(C.QGraphicsView_NewQGraphicsView2(PointerFromQGraphicsScene(scene), PointerFromQWidget(parent)))
}

func NewQGraphicsView(parent QWidget_ITF) *QGraphicsView {
	defer qt.Recovering("QGraphicsView::QGraphicsView")

	return NewQGraphicsViewFromPointer(C.QGraphicsView_NewQGraphicsView(PointerFromQWidget(parent)))
}

func (ptr *QGraphicsView) CenterOn3(item QGraphicsItem_ITF) {
	defer qt.Recovering("QGraphicsView::centerOn")

	if ptr.Pointer() != nil {
		C.QGraphicsView_CenterOn3(ptr.Pointer(), PointerFromQGraphicsItem(item))
	}
}

func (ptr *QGraphicsView) CenterOn(pos core.QPointF_ITF) {
	defer qt.Recovering("QGraphicsView::centerOn")

	if ptr.Pointer() != nil {
		C.QGraphicsView_CenterOn(ptr.Pointer(), core.PointerFromQPointF(pos))
	}
}

func (ptr *QGraphicsView) CenterOn2(x float64, y float64) {
	defer qt.Recovering("QGraphicsView::centerOn")

	if ptr.Pointer() != nil {
		C.QGraphicsView_CenterOn2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QGraphicsView) ConnectContextMenuEvent(f func(event *gui.QContextMenuEvent)) {
	defer qt.Recovering("connect QGraphicsView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "contextMenuEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectContextMenuEvent() {
	defer qt.Recovering("disconnect QGraphicsView::contextMenuEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "contextMenuEvent")
	}
}

//export callbackQGraphicsViewContextMenuEvent
func callbackQGraphicsViewContextMenuEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::contextMenuEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "contextMenuEvent"); signal != nil {
		signal.(func(*gui.QContextMenuEvent))(gui.NewQContextMenuEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).ContextMenuEventDefault(gui.NewQContextMenuEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) ContextMenuEvent(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ContextMenuEvent(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QGraphicsView) ContextMenuEventDefault(event gui.QContextMenuEvent_ITF) {
	defer qt.Recovering("QGraphicsView::contextMenuEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ContextMenuEventDefault(ptr.Pointer(), gui.PointerFromQContextMenuEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectDragEnterEvent(f func(event *gui.QDragEnterEvent)) {
	defer qt.Recovering("connect QGraphicsView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragEnterEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectDragEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsView::dragEnterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragEnterEvent")
	}
}

//export callbackQGraphicsViewDragEnterEvent
func callbackQGraphicsViewDragEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::dragEnterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragEnterEvent"); signal != nil {
		signal.(func(*gui.QDragEnterEvent))(gui.NewQDragEnterEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).DragEnterEventDefault(gui.NewQDragEnterEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) DragEnterEvent(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QGraphicsView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_DragEnterEvent(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QGraphicsView) DragEnterEventDefault(event gui.QDragEnterEvent_ITF) {
	defer qt.Recovering("QGraphicsView::dragEnterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_DragEnterEventDefault(ptr.Pointer(), gui.PointerFromQDragEnterEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectDragLeaveEvent(f func(event *gui.QDragLeaveEvent)) {
	defer qt.Recovering("connect QGraphicsView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragLeaveEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectDragLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsView::dragLeaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragLeaveEvent")
	}
}

//export callbackQGraphicsViewDragLeaveEvent
func callbackQGraphicsViewDragLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::dragLeaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragLeaveEvent"); signal != nil {
		signal.(func(*gui.QDragLeaveEvent))(gui.NewQDragLeaveEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).DragLeaveEventDefault(gui.NewQDragLeaveEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) DragLeaveEvent(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QGraphicsView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_DragLeaveEvent(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QGraphicsView) DragLeaveEventDefault(event gui.QDragLeaveEvent_ITF) {
	defer qt.Recovering("QGraphicsView::dragLeaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_DragLeaveEventDefault(ptr.Pointer(), gui.PointerFromQDragLeaveEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectDragMoveEvent(f func(event *gui.QDragMoveEvent)) {
	defer qt.Recovering("connect QGraphicsView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dragMoveEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectDragMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsView::dragMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dragMoveEvent")
	}
}

//export callbackQGraphicsViewDragMoveEvent
func callbackQGraphicsViewDragMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::dragMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dragMoveEvent"); signal != nil {
		signal.(func(*gui.QDragMoveEvent))(gui.NewQDragMoveEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).DragMoveEventDefault(gui.NewQDragMoveEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) DragMoveEvent(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_DragMoveEvent(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QGraphicsView) DragMoveEventDefault(event gui.QDragMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsView::dragMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_DragMoveEventDefault(ptr.Pointer(), gui.PointerFromQDragMoveEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectDropEvent(f func(event *gui.QDropEvent)) {
	defer qt.Recovering("connect QGraphicsView::dropEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "dropEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectDropEvent() {
	defer qt.Recovering("disconnect QGraphicsView::dropEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "dropEvent")
	}
}

//export callbackQGraphicsViewDropEvent
func callbackQGraphicsViewDropEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::dropEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "dropEvent"); signal != nil {
		signal.(func(*gui.QDropEvent))(gui.NewQDropEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).DropEventDefault(gui.NewQDropEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) DropEvent(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QGraphicsView::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_DropEvent(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QGraphicsView) DropEventDefault(event gui.QDropEvent_ITF) {
	defer qt.Recovering("QGraphicsView::dropEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_DropEventDefault(ptr.Pointer(), gui.PointerFromQDropEvent(event))
	}
}

func (ptr *QGraphicsView) EnsureVisible3(item QGraphicsItem_ITF, xmargin int, ymargin int) {
	defer qt.Recovering("QGraphicsView::ensureVisible")

	if ptr.Pointer() != nil {
		C.QGraphicsView_EnsureVisible3(ptr.Pointer(), PointerFromQGraphicsItem(item), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QGraphicsView) EnsureVisible(rect core.QRectF_ITF, xmargin int, ymargin int) {
	defer qt.Recovering("QGraphicsView::ensureVisible")

	if ptr.Pointer() != nil {
		C.QGraphicsView_EnsureVisible(ptr.Pointer(), core.PointerFromQRectF(rect), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QGraphicsView) EnsureVisible2(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	defer qt.Recovering("QGraphicsView::ensureVisible")

	if ptr.Pointer() != nil {
		C.QGraphicsView_EnsureVisible2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QGraphicsView) Event(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsView::event")

	if ptr.Pointer() != nil {
		return C.QGraphicsView_Event(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsView) FitInView3(item QGraphicsItem_ITF, aspectRatioMode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QGraphicsView::fitInView")

	if ptr.Pointer() != nil {
		C.QGraphicsView_FitInView3(ptr.Pointer(), PointerFromQGraphicsItem(item), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) FitInView(rect core.QRectF_ITF, aspectRatioMode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QGraphicsView::fitInView")

	if ptr.Pointer() != nil {
		C.QGraphicsView_FitInView(ptr.Pointer(), core.PointerFromQRectF(rect), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) FitInView2(x float64, y float64, w float64, h float64, aspectRatioMode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QGraphicsView::fitInView")

	if ptr.Pointer() != nil {
		C.QGraphicsView_FitInView2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) ConnectFocusInEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGraphicsView::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQGraphicsViewFocusInEvent
func callbackQGraphicsViewFocusInEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).FocusInEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) FocusInEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_FocusInEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsView) FocusInEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsView::focusInEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_FocusInEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsView) FocusNextPrevChild(next bool) bool {
	defer qt.Recovering("QGraphicsView::focusNextPrevChild")

	if ptr.Pointer() != nil {
		return C.QGraphicsView_FocusNextPrevChild(ptr.Pointer(), C.int(qt.GoBoolToInt(next))) != 0
	}
	return false
}

func (ptr *QGraphicsView) ConnectFocusOutEvent(f func(event *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGraphicsView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusOutEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectFocusOutEvent() {
	defer qt.Recovering("disconnect QGraphicsView::focusOutEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusOutEvent")
	}
}

//export callbackQGraphicsViewFocusOutEvent
func callbackQGraphicsViewFocusOutEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::focusOutEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusOutEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).FocusOutEventDefault(gui.NewQFocusEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) FocusOutEvent(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_FocusOutEvent(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsView) FocusOutEventDefault(event gui.QFocusEvent_ITF) {
	defer qt.Recovering("QGraphicsView::focusOutEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_FocusOutEventDefault(ptr.Pointer(), gui.PointerFromQFocusEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectInputMethodEvent(f func(event *gui.QInputMethodEvent)) {
	defer qt.Recovering("connect QGraphicsView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "inputMethodEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectInputMethodEvent() {
	defer qt.Recovering("disconnect QGraphicsView::inputMethodEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "inputMethodEvent")
	}
}

//export callbackQGraphicsViewInputMethodEvent
func callbackQGraphicsViewInputMethodEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::inputMethodEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "inputMethodEvent"); signal != nil {
		signal.(func(*gui.QInputMethodEvent))(gui.NewQInputMethodEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).InputMethodEventDefault(gui.NewQInputMethodEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) InputMethodEvent(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_InputMethodEvent(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGraphicsView) InputMethodEventDefault(event gui.QInputMethodEvent_ITF) {
	defer qt.Recovering("QGraphicsView::inputMethodEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_InputMethodEventDefault(ptr.Pointer(), gui.PointerFromQInputMethodEvent(event))
	}
}

func (ptr *QGraphicsView) InputMethodQuery(query core.Qt__InputMethodQuery) *core.QVariant {
	defer qt.Recovering("QGraphicsView::inputMethodQuery")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QGraphicsView_InputMethodQuery(ptr.Pointer(), C.int(query)))
	}
	return nil
}

func (ptr *QGraphicsView) InvalidateScene(rect core.QRectF_ITF, layers QGraphicsScene__SceneLayer) {
	defer qt.Recovering("QGraphicsView::invalidateScene")

	if ptr.Pointer() != nil {
		C.QGraphicsView_InvalidateScene(ptr.Pointer(), core.PointerFromQRectF(rect), C.int(layers))
	}
}

func (ptr *QGraphicsView) IsTransformed() bool {
	defer qt.Recovering("QGraphicsView::isTransformed")

	if ptr.Pointer() != nil {
		return C.QGraphicsView_IsTransformed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsView) ItemAt(pos core.QPoint_ITF) *QGraphicsItem {
	defer qt.Recovering("QGraphicsView::itemAt")

	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsView_ItemAt(ptr.Pointer(), core.PointerFromQPoint(pos)))
	}
	return nil
}

func (ptr *QGraphicsView) ItemAt2(x int, y int) *QGraphicsItem {
	defer qt.Recovering("QGraphicsView::itemAt")

	if ptr.Pointer() != nil {
		return NewQGraphicsItemFromPointer(C.QGraphicsView_ItemAt2(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QGraphicsView) ConnectKeyPressEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyPressEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectKeyPressEvent() {
	defer qt.Recovering("disconnect QGraphicsView::keyPressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyPressEvent")
	}
}

//export callbackQGraphicsViewKeyPressEvent
func callbackQGraphicsViewKeyPressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::keyPressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyPressEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).KeyPressEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) KeyPressEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_KeyPressEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsView) KeyPressEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsView::keyPressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_KeyPressEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectKeyReleaseEvent(f func(event *gui.QKeyEvent)) {
	defer qt.Recovering("connect QGraphicsView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "keyReleaseEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectKeyReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsView::keyReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "keyReleaseEvent")
	}
}

//export callbackQGraphicsViewKeyReleaseEvent
func callbackQGraphicsViewKeyReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::keyReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "keyReleaseEvent"); signal != nil {
		signal.(func(*gui.QKeyEvent))(gui.NewQKeyEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).KeyReleaseEventDefault(gui.NewQKeyEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) KeyReleaseEvent(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_KeyReleaseEvent(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsView) KeyReleaseEventDefault(event gui.QKeyEvent_ITF) {
	defer qt.Recovering("QGraphicsView::keyReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_KeyReleaseEventDefault(ptr.Pointer(), gui.PointerFromQKeyEvent(event))
	}
}

func (ptr *QGraphicsView) MapFromScene(point core.QPointF_ITF) *core.QPoint {
	defer qt.Recovering("QGraphicsView::mapFromScene")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QGraphicsView_MapFromScene(ptr.Pointer(), core.PointerFromQPointF(point)))
	}
	return nil
}

func (ptr *QGraphicsView) MapFromScene5(x float64, y float64) *core.QPoint {
	defer qt.Recovering("QGraphicsView::mapFromScene")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QGraphicsView_MapFromScene5(ptr.Pointer(), C.double(x), C.double(y)))
	}
	return nil
}

func (ptr *QGraphicsView) ConnectMouseDoubleClickEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGraphicsView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectMouseDoubleClickEvent() {
	defer qt.Recovering("disconnect QGraphicsView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseDoubleClickEvent")
	}
}

//export callbackQGraphicsViewMouseDoubleClickEvent
func callbackQGraphicsViewMouseDoubleClickEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::mouseDoubleClickEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseDoubleClickEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).MouseDoubleClickEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) MouseDoubleClickEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MouseDoubleClickEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGraphicsView) MouseDoubleClickEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::mouseDoubleClickEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MouseDoubleClickEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGraphicsView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsView::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQGraphicsViewMouseMoveEvent
func callbackQGraphicsViewMouseMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).MouseMoveEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) MouseMoveEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MouseMoveEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGraphicsView) MouseMoveEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::mouseMoveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MouseMoveEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGraphicsView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QGraphicsView::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQGraphicsViewMousePressEvent
func callbackQGraphicsViewMousePressEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).MousePressEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) MousePressEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MousePressEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGraphicsView) MousePressEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::mousePressEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MousePressEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGraphicsView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QGraphicsView::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQGraphicsViewMouseReleaseEvent
func callbackQGraphicsViewMouseReleaseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).MouseReleaseEventDefault(gui.NewQMouseEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) MouseReleaseEvent(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MouseReleaseEvent(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGraphicsView) MouseReleaseEventDefault(event gui.QMouseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::mouseReleaseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MouseReleaseEventDefault(ptr.Pointer(), gui.PointerFromQMouseEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QGraphicsView::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QGraphicsView::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQGraphicsViewPaintEvent
func callbackQGraphicsViewPaintEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).PaintEventDefault(gui.NewQPaintEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) PaintEvent(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QGraphicsView::paintEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_PaintEvent(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QGraphicsView) PaintEventDefault(event gui.QPaintEvent_ITF) {
	defer qt.Recovering("QGraphicsView::paintEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_PaintEventDefault(ptr.Pointer(), gui.PointerFromQPaintEvent(event))
	}
}

func (ptr *QGraphicsView) Render(painter gui.QPainter_ITF, target core.QRectF_ITF, source core.QRect_ITF, aspectRatioMode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QGraphicsView::render")

	if ptr.Pointer() != nil {
		C.QGraphicsView_Render(ptr.Pointer(), gui.PointerFromQPainter(painter), core.PointerFromQRectF(target), core.PointerFromQRect(source), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) ResetCachedContent() {
	defer qt.Recovering("QGraphicsView::resetCachedContent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ResetCachedContent(ptr.Pointer())
	}
}

func (ptr *QGraphicsView) ResetMatrix() {
	defer qt.Recovering("QGraphicsView::resetMatrix")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ResetMatrix(ptr.Pointer())
	}
}

func (ptr *QGraphicsView) ResetTransform() {
	defer qt.Recovering("QGraphicsView::resetTransform")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ResetTransform(ptr.Pointer())
	}
}

func (ptr *QGraphicsView) ConnectResizeEvent(f func(event *gui.QResizeEvent)) {
	defer qt.Recovering("connect QGraphicsView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QGraphicsView::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQGraphicsViewResizeEvent
func callbackQGraphicsViewResizeEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).ResizeEventDefault(gui.NewQResizeEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) ResizeEvent(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QGraphicsView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ResizeEvent(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QGraphicsView) ResizeEventDefault(event gui.QResizeEvent_ITF) {
	defer qt.Recovering("QGraphicsView::resizeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ResizeEventDefault(ptr.Pointer(), gui.PointerFromQResizeEvent(event))
	}
}

func (ptr *QGraphicsView) Rotate(angle float64) {
	defer qt.Recovering("QGraphicsView::rotate")

	if ptr.Pointer() != nil {
		C.QGraphicsView_Rotate(ptr.Pointer(), C.double(angle))
	}
}

func (ptr *QGraphicsView) RubberBandRect() *core.QRect {
	defer qt.Recovering("QGraphicsView::rubberBandRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QGraphicsView_RubberBandRect(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsView) Scale(sx float64, sy float64) {
	defer qt.Recovering("QGraphicsView::scale")

	if ptr.Pointer() != nil {
		C.QGraphicsView_Scale(ptr.Pointer(), C.double(sx), C.double(sy))
	}
}

func (ptr *QGraphicsView) Scene() *QGraphicsScene {
	defer qt.Recovering("QGraphicsView::scene")

	if ptr.Pointer() != nil {
		return NewQGraphicsSceneFromPointer(C.QGraphicsView_Scene(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsView) ConnectScrollContentsBy(f func(dx int, dy int)) {
	defer qt.Recovering("connect QGraphicsView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "scrollContentsBy", f)
	}
}

func (ptr *QGraphicsView) DisconnectScrollContentsBy() {
	defer qt.Recovering("disconnect QGraphicsView::scrollContentsBy")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "scrollContentsBy")
	}
}

//export callbackQGraphicsViewScrollContentsBy
func callbackQGraphicsViewScrollContentsBy(ptr unsafe.Pointer, ptrName *C.char, dx C.int, dy C.int) {
	defer qt.Recovering("callback QGraphicsView::scrollContentsBy")

	if signal := qt.GetSignal(C.GoString(ptrName), "scrollContentsBy"); signal != nil {
		signal.(func(int, int))(int(dx), int(dy))
	} else {
		NewQGraphicsViewFromPointer(ptr).ScrollContentsByDefault(int(dx), int(dy))
	}
}

func (ptr *QGraphicsView) ScrollContentsBy(dx int, dy int) {
	defer qt.Recovering("QGraphicsView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ScrollContentsBy(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QGraphicsView) ScrollContentsByDefault(dx int, dy int) {
	defer qt.Recovering("QGraphicsView::scrollContentsBy")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ScrollContentsByDefault(ptr.Pointer(), C.int(dx), C.int(dy))
	}
}

func (ptr *QGraphicsView) SetOptimizationFlag(flag QGraphicsView__OptimizationFlag, enabled bool) {
	defer qt.Recovering("QGraphicsView::setOptimizationFlag")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetOptimizationFlag(ptr.Pointer(), C.int(flag), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsView) SetRenderHint(hint gui.QPainter__RenderHint, enabled bool) {
	defer qt.Recovering("QGraphicsView::setRenderHint")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetRenderHint(ptr.Pointer(), C.int(hint), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsView) SetScene(scene QGraphicsScene_ITF) {
	defer qt.Recovering("QGraphicsView::setScene")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetScene(ptr.Pointer(), PointerFromQGraphicsScene(scene))
	}
}

func (ptr *QGraphicsView) SetSceneRect2(x float64, y float64, w float64, h float64) {
	defer qt.Recovering("QGraphicsView::setSceneRect")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetSceneRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h))
	}
}

func (ptr *QGraphicsView) SetTransform(matrix gui.QTransform_ITF, combine bool) {
	defer qt.Recovering("QGraphicsView::setTransform")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetTransform(ptr.Pointer(), gui.PointerFromQTransform(matrix), C.int(qt.GoBoolToInt(combine)))
	}
}

func (ptr *QGraphicsView) ConnectSetupViewport(f func(widget *QWidget)) {
	defer qt.Recovering("connect QGraphicsView::setupViewport")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setupViewport", f)
	}
}

func (ptr *QGraphicsView) DisconnectSetupViewport() {
	defer qt.Recovering("disconnect QGraphicsView::setupViewport")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setupViewport")
	}
}

//export callbackQGraphicsViewSetupViewport
func callbackQGraphicsViewSetupViewport(ptr unsafe.Pointer, ptrName *C.char, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QGraphicsView::setupViewport")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupViewport"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QGraphicsView) SetupViewport(widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsView::setupViewport")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetupViewport(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsView) SetupViewportDefault(widget QWidget_ITF) {
	defer qt.Recovering("QGraphicsView::setupViewport")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetupViewportDefault(ptr.Pointer(), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsView) Shear(sh float64, sv float64) {
	defer qt.Recovering("QGraphicsView::shear")

	if ptr.Pointer() != nil {
		C.QGraphicsView_Shear(ptr.Pointer(), C.double(sh), C.double(sv))
	}
}

func (ptr *QGraphicsView) ConnectShowEvent(f func(event *gui.QShowEvent)) {
	defer qt.Recovering("connect QGraphicsView::showEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "showEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectShowEvent() {
	defer qt.Recovering("disconnect QGraphicsView::showEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "showEvent")
	}
}

//export callbackQGraphicsViewShowEvent
func callbackQGraphicsViewShowEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::showEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "showEvent"); signal != nil {
		signal.(func(*gui.QShowEvent))(gui.NewQShowEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).ShowEventDefault(gui.NewQShowEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) ShowEvent(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGraphicsView::showEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ShowEvent(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGraphicsView) ShowEventDefault(event gui.QShowEvent_ITF) {
	defer qt.Recovering("QGraphicsView::showEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ShowEventDefault(ptr.Pointer(), gui.PointerFromQShowEvent(event))
	}
}

func (ptr *QGraphicsView) SizeHint() *core.QSize {
	defer qt.Recovering("QGraphicsView::sizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QGraphicsView_SizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsView) Translate(dx float64, dy float64) {
	defer qt.Recovering("QGraphicsView::translate")

	if ptr.Pointer() != nil {
		C.QGraphicsView_Translate(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QGraphicsView) UpdateSceneRect(rect core.QRectF_ITF) {
	defer qt.Recovering("QGraphicsView::updateSceneRect")

	if ptr.Pointer() != nil {
		C.QGraphicsView_UpdateSceneRect(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QGraphicsView) ViewportEvent(event core.QEvent_ITF) bool {
	defer qt.Recovering("QGraphicsView::viewportEvent")

	if ptr.Pointer() != nil {
		return C.QGraphicsView_ViewportEvent(ptr.Pointer(), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QGraphicsView) ConnectWheelEvent(f func(event *gui.QWheelEvent)) {
	defer qt.Recovering("connect QGraphicsView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "wheelEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectWheelEvent() {
	defer qt.Recovering("disconnect QGraphicsView::wheelEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "wheelEvent")
	}
}

//export callbackQGraphicsViewWheelEvent
func callbackQGraphicsViewWheelEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::wheelEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "wheelEvent"); signal != nil {
		signal.(func(*gui.QWheelEvent))(gui.NewQWheelEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).WheelEventDefault(gui.NewQWheelEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) WheelEvent(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_WheelEvent(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QGraphicsView) WheelEventDefault(event gui.QWheelEvent_ITF) {
	defer qt.Recovering("QGraphicsView::wheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_WheelEventDefault(ptr.Pointer(), gui.PointerFromQWheelEvent(event))
	}
}

func (ptr *QGraphicsView) DestroyQGraphicsView() {
	defer qt.Recovering("QGraphicsView::~QGraphicsView")

	if ptr.Pointer() != nil {
		C.QGraphicsView_DestroyQGraphicsView(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QGraphicsView) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsView::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QGraphicsView::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQGraphicsViewChangeEvent
func callbackQGraphicsViewChangeEvent(ptr unsafe.Pointer, ptrName *C.char, ev unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
	} else {
		NewQGraphicsViewFromPointer(ptr).ChangeEventDefault(core.NewQEventFromPointer(ev))
	}
}

func (ptr *QGraphicsView) ChangeEvent(ev core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsView::changeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ChangeEvent(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QGraphicsView) ChangeEventDefault(ev core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsView::changeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ChangeEventDefault(ptr.Pointer(), core.PointerFromQEvent(ev))
	}
}

func (ptr *QGraphicsView) ConnectActionEvent(f func(event *gui.QActionEvent)) {
	defer qt.Recovering("connect QGraphicsView::actionEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "actionEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectActionEvent() {
	defer qt.Recovering("disconnect QGraphicsView::actionEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "actionEvent")
	}
}

//export callbackQGraphicsViewActionEvent
func callbackQGraphicsViewActionEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::actionEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "actionEvent"); signal != nil {
		signal.(func(*gui.QActionEvent))(gui.NewQActionEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).ActionEventDefault(gui.NewQActionEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) ActionEvent(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QGraphicsView::actionEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ActionEvent(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QGraphicsView) ActionEventDefault(event gui.QActionEvent_ITF) {
	defer qt.Recovering("QGraphicsView::actionEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ActionEventDefault(ptr.Pointer(), gui.PointerFromQActionEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectEnterEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsView::enterEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "enterEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectEnterEvent() {
	defer qt.Recovering("disconnect QGraphicsView::enterEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "enterEvent")
	}
}

//export callbackQGraphicsViewEnterEvent
func callbackQGraphicsViewEnterEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::enterEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "enterEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).EnterEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) EnterEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsView::enterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_EnterEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsView) EnterEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsView::enterEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_EnterEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectHideEvent(f func(event *gui.QHideEvent)) {
	defer qt.Recovering("connect QGraphicsView::hideEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "hideEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectHideEvent() {
	defer qt.Recovering("disconnect QGraphicsView::hideEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "hideEvent")
	}
}

//export callbackQGraphicsViewHideEvent
func callbackQGraphicsViewHideEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::hideEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "hideEvent"); signal != nil {
		signal.(func(*gui.QHideEvent))(gui.NewQHideEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).HideEventDefault(gui.NewQHideEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) HideEvent(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGraphicsView::hideEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_HideEvent(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGraphicsView) HideEventDefault(event gui.QHideEvent_ITF) {
	defer qt.Recovering("QGraphicsView::hideEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_HideEventDefault(ptr.Pointer(), gui.PointerFromQHideEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectLeaveEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "leaveEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectLeaveEvent() {
	defer qt.Recovering("disconnect QGraphicsView::leaveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "leaveEvent")
	}
}

//export callbackQGraphicsViewLeaveEvent
func callbackQGraphicsViewLeaveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::leaveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "leaveEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).LeaveEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) LeaveEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_LeaveEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsView) LeaveEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsView::leaveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_LeaveEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectMoveEvent(f func(event *gui.QMoveEvent)) {
	defer qt.Recovering("connect QGraphicsView::moveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "moveEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectMoveEvent() {
	defer qt.Recovering("disconnect QGraphicsView::moveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "moveEvent")
	}
}

//export callbackQGraphicsViewMoveEvent
func callbackQGraphicsViewMoveEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::moveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "moveEvent"); signal != nil {
		signal.(func(*gui.QMoveEvent))(gui.NewQMoveEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).MoveEventDefault(gui.NewQMoveEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) MoveEvent(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsView::moveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MoveEvent(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QGraphicsView) MoveEventDefault(event gui.QMoveEvent_ITF) {
	defer qt.Recovering("QGraphicsView::moveEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_MoveEventDefault(ptr.Pointer(), gui.PointerFromQMoveEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QGraphicsView::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QGraphicsView) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QGraphicsView::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQGraphicsViewSetVisible
func callbackQGraphicsViewSetVisible(ptr unsafe.Pointer, ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QGraphicsView::setVisible")

	if signal := qt.GetSignal(C.GoString(ptrName), "setVisible"); signal != nil {
		signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QGraphicsView) SetVisible(visible bool) {
	defer qt.Recovering("QGraphicsView::setVisible")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QGraphicsView) SetVisibleDefault(visible bool) {
	defer qt.Recovering("QGraphicsView::setVisible")

	if ptr.Pointer() != nil {
		C.QGraphicsView_SetVisibleDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QGraphicsView) ConnectCloseEvent(f func(event *gui.QCloseEvent)) {
	defer qt.Recovering("connect QGraphicsView::closeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "closeEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectCloseEvent() {
	defer qt.Recovering("disconnect QGraphicsView::closeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "closeEvent")
	}
}

//export callbackQGraphicsViewCloseEvent
func callbackQGraphicsViewCloseEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::closeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "closeEvent"); signal != nil {
		signal.(func(*gui.QCloseEvent))(gui.NewQCloseEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).CloseEventDefault(gui.NewQCloseEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) CloseEvent(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::closeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_CloseEvent(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGraphicsView) CloseEventDefault(event gui.QCloseEvent_ITF) {
	defer qt.Recovering("QGraphicsView::closeEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_CloseEventDefault(ptr.Pointer(), gui.PointerFromQCloseEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectInitPainter(f func(painter *gui.QPainter)) {
	defer qt.Recovering("connect QGraphicsView::initPainter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "initPainter", f)
	}
}

func (ptr *QGraphicsView) DisconnectInitPainter() {
	defer qt.Recovering("disconnect QGraphicsView::initPainter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "initPainter")
	}
}

//export callbackQGraphicsViewInitPainter
func callbackQGraphicsViewInitPainter(ptr unsafe.Pointer, ptrName *C.char, painter unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::initPainter")

	if signal := qt.GetSignal(C.GoString(ptrName), "initPainter"); signal != nil {
		signal.(func(*gui.QPainter))(gui.NewQPainterFromPointer(painter))
	} else {
		NewQGraphicsViewFromPointer(ptr).InitPainterDefault(gui.NewQPainterFromPointer(painter))
	}
}

func (ptr *QGraphicsView) InitPainter(painter gui.QPainter_ITF) {
	defer qt.Recovering("QGraphicsView::initPainter")

	if ptr.Pointer() != nil {
		C.QGraphicsView_InitPainter(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QGraphicsView) InitPainterDefault(painter gui.QPainter_ITF) {
	defer qt.Recovering("QGraphicsView::initPainter")

	if ptr.Pointer() != nil {
		C.QGraphicsView_InitPainterDefault(ptr.Pointer(), gui.PointerFromQPainter(painter))
	}
}

func (ptr *QGraphicsView) ConnectTabletEvent(f func(event *gui.QTabletEvent)) {
	defer qt.Recovering("connect QGraphicsView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "tabletEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectTabletEvent() {
	defer qt.Recovering("disconnect QGraphicsView::tabletEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "tabletEvent")
	}
}

//export callbackQGraphicsViewTabletEvent
func callbackQGraphicsViewTabletEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::tabletEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "tabletEvent"); signal != nil {
		signal.(func(*gui.QTabletEvent))(gui.NewQTabletEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).TabletEventDefault(gui.NewQTabletEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) TabletEvent(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QGraphicsView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_TabletEvent(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QGraphicsView) TabletEventDefault(event gui.QTabletEvent_ITF) {
	defer qt.Recovering("QGraphicsView::tabletEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_TabletEventDefault(ptr.Pointer(), gui.PointerFromQTabletEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QGraphicsView::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QGraphicsView::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQGraphicsViewTimerEvent
func callbackQGraphicsViewTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsView::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsView) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QGraphicsView::timerEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QGraphicsView::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGraphicsView::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGraphicsViewChildEvent
func callbackQGraphicsViewChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsView::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsView) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QGraphicsView::childEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QGraphicsView) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QGraphicsView::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QGraphicsView) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QGraphicsView::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQGraphicsViewCustomEvent
func callbackQGraphicsViewCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QGraphicsView::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQGraphicsViewFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QGraphicsView) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsView::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QGraphicsView) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QGraphicsView::customEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsView_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
