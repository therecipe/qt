package widgets

//#include "qgraphicsview.h"
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

type QGraphicsViewITF interface {
	QAbstractScrollAreaITF
	QGraphicsViewPTR() *QGraphicsView
}

func PointerFromQGraphicsView(ptr QGraphicsViewITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsViewPTR().Pointer()
	}
	return nil
}

func QGraphicsViewFromPointer(ptr unsafe.Pointer) *QGraphicsView {
	var n = new(QGraphicsView)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsView_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsView) QGraphicsViewPTR() *QGraphicsView {
	return ptr
}

//QGraphicsView::CacheModeFlag
type QGraphicsView__CacheModeFlag int

var (
	QGraphicsView__CacheNone       = QGraphicsView__CacheModeFlag(0x0)
	QGraphicsView__CacheBackground = QGraphicsView__CacheModeFlag(0x1)
)

//QGraphicsView::DragMode
type QGraphicsView__DragMode int

var (
	QGraphicsView__NoDrag         = QGraphicsView__DragMode(0)
	QGraphicsView__ScrollHandDrag = QGraphicsView__DragMode(1)
	QGraphicsView__RubberBandDrag = QGraphicsView__DragMode(2)
)

//QGraphicsView::OptimizationFlag
type QGraphicsView__OptimizationFlag int

var (
	QGraphicsView__DontClipPainter           = QGraphicsView__OptimizationFlag(0x1)
	QGraphicsView__DontSavePainterState      = QGraphicsView__OptimizationFlag(0x2)
	QGraphicsView__DontAdjustForAntialiasing = QGraphicsView__OptimizationFlag(0x4)
	QGraphicsView__IndirectPainting          = QGraphicsView__OptimizationFlag(0x8)
)

//QGraphicsView::ViewportAnchor
type QGraphicsView__ViewportAnchor int

var (
	QGraphicsView__NoAnchor         = QGraphicsView__ViewportAnchor(0)
	QGraphicsView__AnchorViewCenter = QGraphicsView__ViewportAnchor(1)
	QGraphicsView__AnchorUnderMouse = QGraphicsView__ViewportAnchor(2)
)

//QGraphicsView::ViewportUpdateMode
type QGraphicsView__ViewportUpdateMode int

var (
	QGraphicsView__FullViewportUpdate         = QGraphicsView__ViewportUpdateMode(0)
	QGraphicsView__MinimalViewportUpdate      = QGraphicsView__ViewportUpdateMode(1)
	QGraphicsView__SmartViewportUpdate        = QGraphicsView__ViewportUpdateMode(2)
	QGraphicsView__NoViewportUpdate           = QGraphicsView__ViewportUpdateMode(3)
	QGraphicsView__BoundingRectViewportUpdate = QGraphicsView__ViewportUpdateMode(4)
)

func (ptr *QGraphicsView) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGraphicsView_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsView) CacheMode() QGraphicsView__CacheModeFlag {
	if ptr.Pointer() != nil {
		return QGraphicsView__CacheModeFlag(C.QGraphicsView_CacheMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsView) DragMode() QGraphicsView__DragMode {
	if ptr.Pointer() != nil {
		return QGraphicsView__DragMode(C.QGraphicsView_DragMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsView) IsInteractive() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsView_IsInteractive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsView) OptimizationFlags() QGraphicsView__OptimizationFlag {
	if ptr.Pointer() != nil {
		return QGraphicsView__OptimizationFlag(C.QGraphicsView_OptimizationFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsView) RenderHints() gui.QPainter__RenderHint {
	if ptr.Pointer() != nil {
		return gui.QPainter__RenderHint(C.QGraphicsView_RenderHints(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsView) ResizeAnchor() QGraphicsView__ViewportAnchor {
	if ptr.Pointer() != nil {
		return QGraphicsView__ViewportAnchor(C.QGraphicsView_ResizeAnchor(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsView) RubberBandSelectionMode() core.Qt__ItemSelectionMode {
	if ptr.Pointer() != nil {
		return core.Qt__ItemSelectionMode(C.QGraphicsView_RubberBandSelectionMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsView) SetAlignment(alignment core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(alignment))
	}
}

func (ptr *QGraphicsView) SetBackgroundBrush(brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetBackgroundBrush(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QGraphicsView) SetCacheMode(mode QGraphicsView__CacheModeFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetCacheMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QGraphicsView) SetDragMode(mode QGraphicsView__DragMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetDragMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QGraphicsView) SetForegroundBrush(brush gui.QBrushITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetForegroundBrush(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQBrush(brush)))
	}
}

func (ptr *QGraphicsView) SetInteractive(allowed bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetInteractive(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(allowed)))
	}
}

func (ptr *QGraphicsView) SetOptimizationFlags(flags QGraphicsView__OptimizationFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetOptimizationFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QGraphicsView) SetRenderHints(hints gui.QPainter__RenderHint) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetRenderHints(C.QtObjectPtr(ptr.Pointer()), C.int(hints))
	}
}

func (ptr *QGraphicsView) SetResizeAnchor(anchor QGraphicsView__ViewportAnchor) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetResizeAnchor(C.QtObjectPtr(ptr.Pointer()), C.int(anchor))
	}
}

func (ptr *QGraphicsView) SetRubberBandSelectionMode(mode core.Qt__ItemSelectionMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetRubberBandSelectionMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QGraphicsView) SetSceneRect(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetSceneRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QGraphicsView) SetTransformationAnchor(anchor QGraphicsView__ViewportAnchor) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetTransformationAnchor(C.QtObjectPtr(ptr.Pointer()), C.int(anchor))
	}
}

func (ptr *QGraphicsView) SetViewportUpdateMode(mode QGraphicsView__ViewportUpdateMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetViewportUpdateMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QGraphicsView) TransformationAnchor() QGraphicsView__ViewportAnchor {
	if ptr.Pointer() != nil {
		return QGraphicsView__ViewportAnchor(C.QGraphicsView_TransformationAnchor(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsView) ViewportUpdateMode() QGraphicsView__ViewportUpdateMode {
	if ptr.Pointer() != nil {
		return QGraphicsView__ViewportUpdateMode(C.QGraphicsView_ViewportUpdateMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQGraphicsView2(scene QGraphicsSceneITF, parent QWidgetITF) *QGraphicsView {
	return QGraphicsViewFromPointer(unsafe.Pointer(C.QGraphicsView_NewQGraphicsView2(C.QtObjectPtr(PointerFromQGraphicsScene(scene)), C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func NewQGraphicsView(parent QWidgetITF) *QGraphicsView {
	return QGraphicsViewFromPointer(unsafe.Pointer(C.QGraphicsView_NewQGraphicsView(C.QtObjectPtr(PointerFromQWidget(parent)))))
}

func (ptr *QGraphicsView) CenterOn3(item QGraphicsItemITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_CenterOn3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)))
	}
}

func (ptr *QGraphicsView) CenterOn(pos core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_CenterOn(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(pos)))
	}
}

func (ptr *QGraphicsView) EnsureVisible3(item QGraphicsItemITF, xmargin int, ymargin int) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_EnsureVisible3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QGraphicsView) EnsureVisible(rect core.QRectFITF, xmargin int, ymargin int) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_EnsureVisible(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)), C.int(xmargin), C.int(ymargin))
	}
}

func (ptr *QGraphicsView) FitInView3(item QGraphicsItemITF, aspectRatioMode core.Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_FitInView3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item)), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) FitInView(rect core.QRectFITF, aspectRatioMode core.Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_FitInView(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) InputMethodQuery(query core.Qt__InputMethodQuery) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsView_InputMethodQuery(C.QtObjectPtr(ptr.Pointer()), C.int(query)))
	}
	return ""
}

func (ptr *QGraphicsView) InvalidateScene(rect core.QRectFITF, layers QGraphicsScene__SceneLayer) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_InvalidateScene(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)), C.int(layers))
	}
}

func (ptr *QGraphicsView) IsTransformed() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsView_IsTransformed(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsView) ItemAt(pos core.QPointITF) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return QGraphicsItemFromPointer(unsafe.Pointer(C.QGraphicsView_ItemAt(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pos)))))
	}
	return nil
}

func (ptr *QGraphicsView) ItemAt2(x int, y int) *QGraphicsItem {
	if ptr.Pointer() != nil {
		return QGraphicsItemFromPointer(unsafe.Pointer(C.QGraphicsView_ItemAt2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))))
	}
	return nil
}

func (ptr *QGraphicsView) Render(painter gui.QPainterITF, target core.QRectFITF, source core.QRectITF, aspectRatioMode core.Qt__AspectRatioMode) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_Render(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(core.PointerFromQRectF(target)), C.QtObjectPtr(core.PointerFromQRect(source)), C.int(aspectRatioMode))
	}
}

func (ptr *QGraphicsView) ResetCachedContent() {
	if ptr.Pointer() != nil {
		C.QGraphicsView_ResetCachedContent(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsView) ResetMatrix() {
	if ptr.Pointer() != nil {
		C.QGraphicsView_ResetMatrix(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsView) ResetTransform() {
	if ptr.Pointer() != nil {
		C.QGraphicsView_ResetTransform(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsView) Scene() *QGraphicsScene {
	if ptr.Pointer() != nil {
		return QGraphicsSceneFromPointer(unsafe.Pointer(C.QGraphicsView_Scene(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsView) SetOptimizationFlag(flag QGraphicsView__OptimizationFlag, enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetOptimizationFlag(C.QtObjectPtr(ptr.Pointer()), C.int(flag), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsView) SetRenderHint(hint gui.QPainter__RenderHint, enabled bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetRenderHint(C.QtObjectPtr(ptr.Pointer()), C.int(hint), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QGraphicsView) SetScene(scene QGraphicsSceneITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetScene(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsScene(scene)))
	}
}

func (ptr *QGraphicsView) SetTransform(matrix gui.QTransformITF, combine bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_SetTransform(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTransform(matrix)), C.int(qt.GoBoolToInt(combine)))
	}
}

func (ptr *QGraphicsView) UpdateSceneRect(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsView_UpdateSceneRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QGraphicsView) DestroyQGraphicsView() {
	if ptr.Pointer() != nil {
		C.QGraphicsView_DestroyQGraphicsView(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
