package gui

//#include "qpaintengine.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPaintEngine struct {
	ptr unsafe.Pointer
}

type QPaintEngine_ITF interface {
	QPaintEngine_PTR() *QPaintEngine
}

func (p *QPaintEngine) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPaintEngine) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPaintEngine(ptr QPaintEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEngine_PTR().Pointer()
	}
	return nil
}

func NewQPaintEngineFromPointer(ptr unsafe.Pointer) *QPaintEngine {
	var n = new(QPaintEngine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPaintEngine) QPaintEngine_PTR() *QPaintEngine {
	return ptr
}

//QPaintEngine::DirtyFlag
type QPaintEngine__DirtyFlag int64

const (
	QPaintEngine__DirtyPen             = QPaintEngine__DirtyFlag(0x0001)
	QPaintEngine__DirtyBrush           = QPaintEngine__DirtyFlag(0x0002)
	QPaintEngine__DirtyBrushOrigin     = QPaintEngine__DirtyFlag(0x0004)
	QPaintEngine__DirtyFont            = QPaintEngine__DirtyFlag(0x0008)
	QPaintEngine__DirtyBackground      = QPaintEngine__DirtyFlag(0x0010)
	QPaintEngine__DirtyBackgroundMode  = QPaintEngine__DirtyFlag(0x0020)
	QPaintEngine__DirtyTransform       = QPaintEngine__DirtyFlag(0x0040)
	QPaintEngine__DirtyClipRegion      = QPaintEngine__DirtyFlag(0x0080)
	QPaintEngine__DirtyClipPath        = QPaintEngine__DirtyFlag(0x0100)
	QPaintEngine__DirtyHints           = QPaintEngine__DirtyFlag(0x0200)
	QPaintEngine__DirtyCompositionMode = QPaintEngine__DirtyFlag(0x0400)
	QPaintEngine__DirtyClipEnabled     = QPaintEngine__DirtyFlag(0x0800)
	QPaintEngine__DirtyOpacity         = QPaintEngine__DirtyFlag(0x1000)
	QPaintEngine__AllDirty             = QPaintEngine__DirtyFlag(0xffff)
)

//QPaintEngine::PaintEngineFeature
type QPaintEngine__PaintEngineFeature int64

const (
	QPaintEngine__PrimitiveTransform          = QPaintEngine__PaintEngineFeature(0x00000001)
	QPaintEngine__PatternTransform            = QPaintEngine__PaintEngineFeature(0x00000002)
	QPaintEngine__PixmapTransform             = QPaintEngine__PaintEngineFeature(0x00000004)
	QPaintEngine__PatternBrush                = QPaintEngine__PaintEngineFeature(0x00000008)
	QPaintEngine__LinearGradientFill          = QPaintEngine__PaintEngineFeature(0x00000010)
	QPaintEngine__RadialGradientFill          = QPaintEngine__PaintEngineFeature(0x00000020)
	QPaintEngine__ConicalGradientFill         = QPaintEngine__PaintEngineFeature(0x00000040)
	QPaintEngine__AlphaBlend                  = QPaintEngine__PaintEngineFeature(0x00000080)
	QPaintEngine__PorterDuff                  = QPaintEngine__PaintEngineFeature(0x00000100)
	QPaintEngine__PainterPaths                = QPaintEngine__PaintEngineFeature(0x00000200)
	QPaintEngine__Antialiasing                = QPaintEngine__PaintEngineFeature(0x00000400)
	QPaintEngine__BrushStroke                 = QPaintEngine__PaintEngineFeature(0x00000800)
	QPaintEngine__ConstantOpacity             = QPaintEngine__PaintEngineFeature(0x00001000)
	QPaintEngine__MaskedBrush                 = QPaintEngine__PaintEngineFeature(0x00002000)
	QPaintEngine__PerspectiveTransform        = QPaintEngine__PaintEngineFeature(0x00004000)
	QPaintEngine__BlendModes                  = QPaintEngine__PaintEngineFeature(0x00008000)
	QPaintEngine__ObjectBoundingModeGradients = QPaintEngine__PaintEngineFeature(0x00010000)
	QPaintEngine__RasterOpModes               = QPaintEngine__PaintEngineFeature(0x00020000)
	QPaintEngine__PaintOutsidePaintEvent      = QPaintEngine__PaintEngineFeature(0x20000000)
	QPaintEngine__AllFeatures                 = QPaintEngine__PaintEngineFeature(0xffffffff)
)

//QPaintEngine::PolygonDrawMode
type QPaintEngine__PolygonDrawMode int64

const (
	QPaintEngine__OddEvenMode  = QPaintEngine__PolygonDrawMode(0)
	QPaintEngine__WindingMode  = QPaintEngine__PolygonDrawMode(1)
	QPaintEngine__ConvexMode   = QPaintEngine__PolygonDrawMode(2)
	QPaintEngine__PolylineMode = QPaintEngine__PolygonDrawMode(3)
)

//QPaintEngine::Type
type QPaintEngine__Type int64

const (
	QPaintEngine__X11           = QPaintEngine__Type(0)
	QPaintEngine__Windows       = QPaintEngine__Type(1)
	QPaintEngine__QuickDraw     = QPaintEngine__Type(2)
	QPaintEngine__CoreGraphics  = QPaintEngine__Type(3)
	QPaintEngine__MacPrinter    = QPaintEngine__Type(4)
	QPaintEngine__QWindowSystem = QPaintEngine__Type(5)
	QPaintEngine__PostScript    = QPaintEngine__Type(6)
	QPaintEngine__OpenGL        = QPaintEngine__Type(7)
	QPaintEngine__Picture       = QPaintEngine__Type(8)
	QPaintEngine__SVG           = QPaintEngine__Type(9)
	QPaintEngine__Raster        = QPaintEngine__Type(10)
	QPaintEngine__Direct3D      = QPaintEngine__Type(11)
	QPaintEngine__Pdf           = QPaintEngine__Type(12)
	QPaintEngine__OpenVG        = QPaintEngine__Type(13)
	QPaintEngine__OpenGL2       = QPaintEngine__Type(14)
	QPaintEngine__PaintBuffer   = QPaintEngine__Type(15)
	QPaintEngine__Blitter       = QPaintEngine__Type(16)
	QPaintEngine__Direct2D      = QPaintEngine__Type(17)
	QPaintEngine__User          = QPaintEngine__Type(50)
	QPaintEngine__MaxUser       = QPaintEngine__Type(100)
)

func (ptr *QPaintEngine) DrawEllipse(rect core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawEllipse(ptr.Pointer(), core.PointerFromQRectF(rect))
	}
}

func (ptr *QPaintEngine) DrawImage(rectangle core.QRectF_ITF, image QImage_ITF, sr core.QRectF_ITF, flags core.Qt__ImageConversionFlag) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawImage(ptr.Pointer(), core.PointerFromQRectF(rectangle), PointerFromQImage(image), core.PointerFromQRectF(sr), C.int(flags))
	}
}

func (ptr *QPaintEngine) DrawPolygon(points core.QPointF_ITF, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPolygon(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount), C.int(mode))
	}
}

func (ptr *QPaintEngine) Begin(pdev QPaintDevice_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngine_Begin(ptr.Pointer(), PointerFromQPaintDevice(pdev)) != 0
	}
	return false
}

func (ptr *QPaintEngine) DrawEllipse2(rect core.QRect_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawEllipse2(ptr.Pointer(), core.PointerFromQRect(rect))
	}
}

func (ptr *QPaintEngine) DrawLines2(lines core.QLine_ITF, lineCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawLines2(ptr.Pointer(), core.PointerFromQLine(lines), C.int(lineCount))
	}
}

func (ptr *QPaintEngine) DrawLines(lines core.QLineF_ITF, lineCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawLines(ptr.Pointer(), core.PointerFromQLineF(lines), C.int(lineCount))
	}
}

func (ptr *QPaintEngine) DrawPath(path QPainterPath_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPath(ptr.Pointer(), PointerFromQPainterPath(path))
	}
}

func (ptr *QPaintEngine) DrawPixmap(r core.QRectF_ITF, pm QPixmap_ITF, sr core.QRectF_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPixmap(ptr.Pointer(), core.PointerFromQRectF(r), PointerFromQPixmap(pm), core.PointerFromQRectF(sr))
	}
}

func (ptr *QPaintEngine) DrawPoints2(points core.QPoint_ITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPoints2(ptr.Pointer(), core.PointerFromQPoint(points), C.int(pointCount))
	}
}

func (ptr *QPaintEngine) DrawPoints(points core.QPointF_ITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPoints(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount))
	}
}

func (ptr *QPaintEngine) DrawPolygon2(points core.QPoint_ITF, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPolygon2(ptr.Pointer(), core.PointerFromQPoint(points), C.int(pointCount), C.int(mode))
	}
}

func (ptr *QPaintEngine) DrawRects2(rects core.QRect_ITF, rectCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawRects2(ptr.Pointer(), core.PointerFromQRect(rects), C.int(rectCount))
	}
}

func (ptr *QPaintEngine) DrawRects(rects core.QRectF_ITF, rectCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawRects(ptr.Pointer(), core.PointerFromQRectF(rects), C.int(rectCount))
	}
}

func (ptr *QPaintEngine) DrawTextItem(p core.QPointF_ITF, textItem QTextItem_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawTextItem(ptr.Pointer(), core.PointerFromQPointF(p), PointerFromQTextItem(textItem))
	}
}

func (ptr *QPaintEngine) DrawTiledPixmap(rect core.QRectF_ITF, pixmap QPixmap_ITF, p core.QPointF_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawTiledPixmap(ptr.Pointer(), core.PointerFromQRectF(rect), PointerFromQPixmap(pixmap), core.PointerFromQPointF(p))
	}
}

func (ptr *QPaintEngine) End() bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngine_End(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPaintEngine) HasFeature(feature QPaintEngine__PaintEngineFeature) bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngine_HasFeature(ptr.Pointer(), C.int(feature)) != 0
	}
	return false
}

func (ptr *QPaintEngine) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngine_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPaintEngine) PaintDevice() *QPaintDevice {
	if ptr.Pointer() != nil {
		return NewQPaintDeviceFromPointer(C.QPaintEngine_PaintDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintEngine) Painter() *QPainter {
	if ptr.Pointer() != nil {
		return NewQPainterFromPointer(C.QPaintEngine_Painter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintEngine) SetActive(state bool) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(state)))
	}
}

func (ptr *QPaintEngine) Type() QPaintEngine__Type {
	if ptr.Pointer() != nil {
		return QPaintEngine__Type(C.QPaintEngine_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintEngine) UpdateState(state QPaintEngineState_ITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_UpdateState(ptr.Pointer(), PointerFromQPaintEngineState(state))
	}
}

func (ptr *QPaintEngine) DestroyQPaintEngine() {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DestroyQPaintEngine(ptr.Pointer())
	}
}
