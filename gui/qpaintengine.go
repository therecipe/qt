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

type QPaintEngineITF interface {
	QPaintEnginePTR() *QPaintEngine
}

func (p *QPaintEngine) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPaintEngine) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPaintEngine(ptr QPaintEngineITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPaintEnginePTR().Pointer()
	}
	return nil
}

func QPaintEngineFromPointer(ptr unsafe.Pointer) *QPaintEngine {
	var n = new(QPaintEngine)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPaintEngine) QPaintEnginePTR() *QPaintEngine {
	return ptr
}

//QPaintEngine::DirtyFlag
type QPaintEngine__DirtyFlag int

var (
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
type QPaintEngine__PaintEngineFeature int

var (
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
type QPaintEngine__PolygonDrawMode int

var (
	QPaintEngine__OddEvenMode  = QPaintEngine__PolygonDrawMode(0)
	QPaintEngine__WindingMode  = QPaintEngine__PolygonDrawMode(1)
	QPaintEngine__ConvexMode   = QPaintEngine__PolygonDrawMode(2)
	QPaintEngine__PolylineMode = QPaintEngine__PolygonDrawMode(3)
)

//QPaintEngine::Type
type QPaintEngine__Type int

var (
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

func (ptr *QPaintEngine) DrawEllipse(rect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawEllipse(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)))
	}
}

func (ptr *QPaintEngine) DrawImage(rectangle core.QRectFITF, image QImageITF, sr core.QRectFITF, flags core.Qt__ImageConversionFlag) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawImage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.QtObjectPtr(PointerFromQImage(image)), C.QtObjectPtr(core.PointerFromQRectF(sr)), C.int(flags))
	}
}

func (ptr *QPaintEngine) DrawPolygon(points core.QPointFITF, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPolygon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(points)), C.int(pointCount), C.int(mode))
	}
}

func (ptr *QPaintEngine) Begin(pdev QPaintDeviceITF) bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngine_Begin(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPaintDevice(pdev))) != 0
	}
	return false
}

func (ptr *QPaintEngine) DrawEllipse2(rect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawEllipse2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rect)))
	}
}

func (ptr *QPaintEngine) DrawLines2(lines core.QLineITF, lineCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawLines2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLine(lines)), C.int(lineCount))
	}
}

func (ptr *QPaintEngine) DrawLines(lines core.QLineFITF, lineCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawLines(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLineF(lines)), C.int(lineCount))
	}
}

func (ptr *QPaintEngine) DrawPath(path QPainterPathITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(path)))
	}
}

func (ptr *QPaintEngine) DrawPixmap(r core.QRectFITF, pm QPixmapITF, sr core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(r)), C.QtObjectPtr(PointerFromQPixmap(pm)), C.QtObjectPtr(core.PointerFromQRectF(sr)))
	}
}

func (ptr *QPaintEngine) DrawPoints2(points core.QPointITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPoints2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(points)), C.int(pointCount))
	}
}

func (ptr *QPaintEngine) DrawPoints(points core.QPointFITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPoints(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(points)), C.int(pointCount))
	}
}

func (ptr *QPaintEngine) DrawPolygon2(points core.QPointITF, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPolygon2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(points)), C.int(pointCount), C.int(mode))
	}
}

func (ptr *QPaintEngine) DrawRects2(rects core.QRectITF, rectCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawRects2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rects)), C.int(rectCount))
	}
}

func (ptr *QPaintEngine) DrawRects(rects core.QRectFITF, rectCount int) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawRects(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rects)), C.int(rectCount))
	}
}

func (ptr *QPaintEngine) DrawTextItem(p core.QPointFITF, textItem QTextItemITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawTextItem(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(p)), C.QtObjectPtr(PointerFromQTextItem(textItem)))
	}
}

func (ptr *QPaintEngine) DrawTiledPixmap(rect core.QRectFITF, pixmap QPixmapITF, p core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawTiledPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rect)), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.QtObjectPtr(core.PointerFromQPointF(p)))
	}
}

func (ptr *QPaintEngine) End() bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngine_End(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPaintEngine) HasFeature(feature QPaintEngine__PaintEngineFeature) bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngine_HasFeature(C.QtObjectPtr(ptr.Pointer()), C.int(feature)) != 0
	}
	return false
}

func (ptr *QPaintEngine) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QPaintEngine_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPaintEngine) PaintDevice() *QPaintDevice {
	if ptr.Pointer() != nil {
		return QPaintDeviceFromPointer(unsafe.Pointer(C.QPaintEngine_PaintDevice(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPaintEngine) Painter() *QPainter {
	if ptr.Pointer() != nil {
		return QPainterFromPointer(unsafe.Pointer(C.QPaintEngine_Painter(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPaintEngine) SetActive(state bool) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_SetActive(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(state)))
	}
}

func (ptr *QPaintEngine) Type() QPaintEngine__Type {
	if ptr.Pointer() != nil {
		return QPaintEngine__Type(C.QPaintEngine_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPaintEngine) UpdateState(state QPaintEngineStateITF) {
	if ptr.Pointer() != nil {
		C.QPaintEngine_UpdateState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPaintEngineState(state)))
	}
}

func (ptr *QPaintEngine) DestroyQPaintEngine() {
	if ptr.Pointer() != nil {
		C.QPaintEngine_DestroyQPaintEngine(C.QtObjectPtr(ptr.Pointer()))
	}
}
