package gui

//#include "gui.h"
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
	for len(n.ObjectNameAbs()) < len("QPaintEngine_") {
		n.SetObjectNameAbs("QPaintEngine_" + qt.Identifier())
	}
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

func (ptr *QPaintEngine) ConnectDrawPolygon(f func(points *core.QPointF, pointCount int, mode QPaintEngine__PolygonDrawMode)) {
	defer qt.Recovering("connect QPaintEngine::drawPolygon")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "drawPolygon", f)
	}
}

func (ptr *QPaintEngine) DisconnectDrawPolygon() {
	defer qt.Recovering("disconnect QPaintEngine::drawPolygon")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "drawPolygon")
	}
}

//export callbackQPaintEngineDrawPolygon
func callbackQPaintEngineDrawPolygon(ptr unsafe.Pointer, ptrName *C.char, points unsafe.Pointer, pointCount C.int, mode C.int) {
	defer qt.Recovering("callback QPaintEngine::drawPolygon")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawPolygon"); signal != nil {
		signal.(func(*core.QPointF, int, QPaintEngine__PolygonDrawMode))(core.NewQPointFFromPointer(points), int(pointCount), QPaintEngine__PolygonDrawMode(mode))
	} else {
		NewQPaintEngineFromPointer(ptr).DrawPolygonDefault(core.NewQPointFFromPointer(points), int(pointCount), QPaintEngine__PolygonDrawMode(mode))
	}
}

func (ptr *QPaintEngine) DrawPolygon(points core.QPointF_ITF, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	defer qt.Recovering("QPaintEngine::drawPolygon")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPolygon(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount), C.int(mode))
	}
}

func (ptr *QPaintEngine) DrawPolygonDefault(points core.QPointF_ITF, pointCount int, mode QPaintEngine__PolygonDrawMode) {
	defer qt.Recovering("QPaintEngine::drawPolygon")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPolygonDefault(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount), C.int(mode))
	}
}

func (ptr *QPaintEngine) Begin(pdev QPaintDevice_ITF) bool {
	defer qt.Recovering("QPaintEngine::begin")

	if ptr.Pointer() != nil {
		return C.QPaintEngine_Begin(ptr.Pointer(), PointerFromQPaintDevice(pdev)) != 0
	}
	return false
}

func (ptr *QPaintEngine) ConnectDrawLines(f func(lines *core.QLineF, lineCount int)) {
	defer qt.Recovering("connect QPaintEngine::drawLines")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "drawLines", f)
	}
}

func (ptr *QPaintEngine) DisconnectDrawLines() {
	defer qt.Recovering("disconnect QPaintEngine::drawLines")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "drawLines")
	}
}

//export callbackQPaintEngineDrawLines
func callbackQPaintEngineDrawLines(ptr unsafe.Pointer, ptrName *C.char, lines unsafe.Pointer, lineCount C.int) {
	defer qt.Recovering("callback QPaintEngine::drawLines")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawLines"); signal != nil {
		signal.(func(*core.QLineF, int))(core.NewQLineFFromPointer(lines), int(lineCount))
	} else {
		NewQPaintEngineFromPointer(ptr).DrawLinesDefault(core.NewQLineFFromPointer(lines), int(lineCount))
	}
}

func (ptr *QPaintEngine) DrawLines(lines core.QLineF_ITF, lineCount int) {
	defer qt.Recovering("QPaintEngine::drawLines")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawLines(ptr.Pointer(), core.PointerFromQLineF(lines), C.int(lineCount))
	}
}

func (ptr *QPaintEngine) DrawLinesDefault(lines core.QLineF_ITF, lineCount int) {
	defer qt.Recovering("QPaintEngine::drawLines")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawLinesDefault(ptr.Pointer(), core.PointerFromQLineF(lines), C.int(lineCount))
	}
}

func (ptr *QPaintEngine) DrawPixmap(r core.QRectF_ITF, pm QPixmap_ITF, sr core.QRectF_ITF) {
	defer qt.Recovering("QPaintEngine::drawPixmap")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPixmap(ptr.Pointer(), core.PointerFromQRectF(r), PointerFromQPixmap(pm), core.PointerFromQRectF(sr))
	}
}

func (ptr *QPaintEngine) ConnectDrawPoints(f func(points *core.QPointF, pointCount int)) {
	defer qt.Recovering("connect QPaintEngine::drawPoints")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "drawPoints", f)
	}
}

func (ptr *QPaintEngine) DisconnectDrawPoints() {
	defer qt.Recovering("disconnect QPaintEngine::drawPoints")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "drawPoints")
	}
}

//export callbackQPaintEngineDrawPoints
func callbackQPaintEngineDrawPoints(ptr unsafe.Pointer, ptrName *C.char, points unsafe.Pointer, pointCount C.int) {
	defer qt.Recovering("callback QPaintEngine::drawPoints")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawPoints"); signal != nil {
		signal.(func(*core.QPointF, int))(core.NewQPointFFromPointer(points), int(pointCount))
	} else {
		NewQPaintEngineFromPointer(ptr).DrawPointsDefault(core.NewQPointFFromPointer(points), int(pointCount))
	}
}

func (ptr *QPaintEngine) DrawPoints(points core.QPointF_ITF, pointCount int) {
	defer qt.Recovering("QPaintEngine::drawPoints")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPoints(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount))
	}
}

func (ptr *QPaintEngine) DrawPointsDefault(points core.QPointF_ITF, pointCount int) {
	defer qt.Recovering("QPaintEngine::drawPoints")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawPointsDefault(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount))
	}
}

func (ptr *QPaintEngine) ConnectDrawRects(f func(rects *core.QRectF, rectCount int)) {
	defer qt.Recovering("connect QPaintEngine::drawRects")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "drawRects", f)
	}
}

func (ptr *QPaintEngine) DisconnectDrawRects() {
	defer qt.Recovering("disconnect QPaintEngine::drawRects")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "drawRects")
	}
}

//export callbackQPaintEngineDrawRects
func callbackQPaintEngineDrawRects(ptr unsafe.Pointer, ptrName *C.char, rects unsafe.Pointer, rectCount C.int) {
	defer qt.Recovering("callback QPaintEngine::drawRects")

	if signal := qt.GetSignal(C.GoString(ptrName), "drawRects"); signal != nil {
		signal.(func(*core.QRectF, int))(core.NewQRectFFromPointer(rects), int(rectCount))
	} else {
		NewQPaintEngineFromPointer(ptr).DrawRectsDefault(core.NewQRectFFromPointer(rects), int(rectCount))
	}
}

func (ptr *QPaintEngine) DrawRects(rects core.QRectF_ITF, rectCount int) {
	defer qt.Recovering("QPaintEngine::drawRects")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawRects(ptr.Pointer(), core.PointerFromQRectF(rects), C.int(rectCount))
	}
}

func (ptr *QPaintEngine) DrawRectsDefault(rects core.QRectF_ITF, rectCount int) {
	defer qt.Recovering("QPaintEngine::drawRects")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DrawRectsDefault(ptr.Pointer(), core.PointerFromQRectF(rects), C.int(rectCount))
	}
}

func (ptr *QPaintEngine) End() bool {
	defer qt.Recovering("QPaintEngine::end")

	if ptr.Pointer() != nil {
		return C.QPaintEngine_End(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPaintEngine) HasFeature(feature QPaintEngine__PaintEngineFeature) bool {
	defer qt.Recovering("QPaintEngine::hasFeature")

	if ptr.Pointer() != nil {
		return C.QPaintEngine_HasFeature(ptr.Pointer(), C.int(feature)) != 0
	}
	return false
}

func (ptr *QPaintEngine) IsActive() bool {
	defer qt.Recovering("QPaintEngine::isActive")

	if ptr.Pointer() != nil {
		return C.QPaintEngine_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPaintEngine) PaintDevice() *QPaintDevice {
	defer qt.Recovering("QPaintEngine::paintDevice")

	if ptr.Pointer() != nil {
		return NewQPaintDeviceFromPointer(C.QPaintEngine_PaintDevice(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintEngine) Painter() *QPainter {
	defer qt.Recovering("QPaintEngine::painter")

	if ptr.Pointer() != nil {
		return NewQPainterFromPointer(C.QPaintEngine_Painter(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPaintEngine) SetActive(state bool) {
	defer qt.Recovering("QPaintEngine::setActive")

	if ptr.Pointer() != nil {
		C.QPaintEngine_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(state)))
	}
}

func (ptr *QPaintEngine) Type() QPaintEngine__Type {
	defer qt.Recovering("QPaintEngine::type")

	if ptr.Pointer() != nil {
		return QPaintEngine__Type(C.QPaintEngine_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPaintEngine) UpdateState(state QPaintEngineState_ITF) {
	defer qt.Recovering("QPaintEngine::updateState")

	if ptr.Pointer() != nil {
		C.QPaintEngine_UpdateState(ptr.Pointer(), PointerFromQPaintEngineState(state))
	}
}

func (ptr *QPaintEngine) DestroyQPaintEngine() {
	defer qt.Recovering("QPaintEngine::~QPaintEngine")

	if ptr.Pointer() != nil {
		C.QPaintEngine_DestroyQPaintEngine(ptr.Pointer())
	}
}

func (ptr *QPaintEngine) ObjectNameAbs() string {
	defer qt.Recovering("QPaintEngine::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QPaintEngine_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QPaintEngine) SetObjectNameAbs(name string) {
	defer qt.Recovering("QPaintEngine::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QPaintEngine_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
